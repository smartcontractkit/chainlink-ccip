package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type SimpleConfig struct {
	Version   string                 `yaml:"version"`
	Output    OutputConfig           `yaml:"output"`
	Contracts []SimpleContractConfig `yaml:"contracts"`
}

type OutputConfig struct {
	BasePath      string `yaml:"base_path"`
	VersionPrefix string `yaml:"version_prefix"`
}

type SimpleContractConfig struct {
	SolidityPath string   `yaml:"solidity_path"`
	ContractName string   `yaml:"contract_name,omitempty"`
	Version      string   `yaml:"version,omitempty"`
	Functions    []string `yaml:"functions"`
}

// ABI structures
type ABIEntry struct {
	Type            string     `json:"type"`
	Name            string     `json:"name"`
	Inputs          []ABIParam `json:"inputs"`
	Outputs         []ABIParam `json:"outputs"`
	StateMutability string     `json:"stateMutability"`
}

type ABIParam struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType"`
}

// Contract info
type ContractInfo struct {
	Name            string
	Version         string
	PackageName     string
	OutputPath      string
	GobindingImport string
	GobindingPrefix string
	Constructor     *FunctionInfo
	Functions       map[string]*FunctionInfo
}

type FunctionInfo struct {
	Name            string
	SolidityName    string
	IsConstructor   bool
	StateMutability string
	Parameters      []ParameterInfo
	ReturnParams    []ParameterInfo
	IsWrite         bool
	IsRead          bool
	CallMethod      string
	HasOnlyOwner    bool // detected from Solidity source
}

type ParameterInfo struct {
	Name         string
	SolidityType string
	GoType       string
}

func main() {
	configPath := flag.String("config", "chains/evm/operations_gen_config_simple.yaml", "Path to config file")
	flag.Parse()

	configData, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	var config SimpleConfig
	if err := yaml.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %v\n", err)
		os.Exit(1)
	}

	for _, contractCfg := range config.Contracts {
		contractInfo, err := extractContractInfo(contractCfg, config.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting info from %s: %v\n", contractCfg.SolidityPath, err)
			os.Exit(1)
		}

		if err := generateOperationsFile(contractInfo); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating operations for %s: %v\n", contractInfo.Name, err)
			os.Exit(1)
		}

		fmt.Printf("✓ Generated operations for %s at %s\n", contractInfo.Name, contractInfo.OutputPath)
	}
}

func extractContractInfo(cfg SimpleContractConfig, output OutputConfig) (*ContractInfo, error) {
	contractName := cfg.ContractName
	if contractName == "" {
		contractName = strings.TrimSuffix(filepath.Base(cfg.SolidityPath), ".sol")
	}

	packageName := toSnakeCase(contractName)
	gobindingPath := filepath.Join("chains", "evm", "gobindings", "generated", output.VersionPrefix, packageName, packageName+".go")

	// Extract ABI from gobinding
	abi, err := extractABIFromGobinding(gobindingPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract ABI: %w", err)
	}

	// Parse ABI
	var abiEntries []ABIEntry
	if err := json.Unmarshal([]byte(abi), &abiEntries); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Extract version from typeAndVersion if not specified
	version := cfg.Version
	if version == "" {
		version = extractVersionFromSolidity(cfg.SolidityPath)
		if version == "" {
			version = "1.6.0" // default
		}
	}

	info := &ContractInfo{
		Name:            contractName,
		Version:         version,
		PackageName:     packageName,
		GobindingPrefix: packageName,
		Functions:       make(map[string]*FunctionInfo),
	}

	info.OutputPath = filepath.Join(output.BasePath, output.VersionPrefix, "operations", info.PackageName, info.PackageName+".go")
	info.GobindingImport = fmt.Sprintf("github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/%s/%s",
		output.VersionPrefix, info.GobindingPrefix)

	// Check for onlyOwner modifier in Solidity source
	ownershipFuncs := detectOwnershipModifiers(cfg.SolidityPath)

	// Extract constructor
	for _, entry := range abiEntries {
		if entry.Type == "constructor" {
			info.Constructor = parseABIFunction(entry, true, false)
			break
		}
	}

	// Extract requested functions
	for _, funcName := range cfg.Functions {
		funcInfo := findFunctionInABI(abiEntries, funcName)
		if funcInfo == nil {
			return nil, fmt.Errorf("function %s not found in ABI", funcName)
		}
		funcInfo.HasOnlyOwner = ownershipFuncs[funcName]
		info.Functions[funcName] = funcInfo
	}

	return info, nil
}

func extractABIFromGobinding(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Find the ABI string in the MetaData
	re := regexp.MustCompile(`ABI: "((?:[^"\\]|\\.)*)"`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		return "", fmt.Errorf("ABI not found in gobinding")
	}

	// Unescape the JSON string
	abi := matches[1]
	abi = strings.ReplaceAll(abi, `\"`, `"`)
	abi = strings.ReplaceAll(abi, `\\`, `\`)

	return abi, nil
}

func extractVersionFromSolidity(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	re := regexp.MustCompile(`typeAndVersion.*?=.*?"[^"]*\s+(\d+\.\d+\.\d+)"`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func detectOwnershipModifiers(path string) map[string]bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	source := string(data)
	result := make(map[string]bool)

	// Find functions with onlyOwner modifier
	re := regexp.MustCompile(`function\s+(\w+)\s*\([^)]*\)[^{]*\bonlyOwner\b`)
	matches := re.FindAllStringSubmatch(source, -1)
	for _, match := range matches {
		if len(match) > 1 {
			result[match[1]] = true
		}
	}

	return result
}

func findFunctionInABI(entries []ABIEntry, funcName string) *FunctionInfo {
	// Find all functions with this name (may be overloaded)
	var candidates []ABIEntry
	for _, entry := range entries {
		if entry.Type == "function" && entry.Name == funcName {
			candidates = append(candidates, entry)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	// Prefer array version if multiple exist
	var selected ABIEntry
	hasArrayVersion := false
	for _, c := range candidates {
		for _, param := range c.Inputs {
			if strings.Contains(param.Type, "[]") {
				selected = c
				hasArrayVersion = true
				break
			}
		}
		if hasArrayVersion {
			break
		}
	}
	if !hasArrayVersion {
		selected = candidates[0]
	}

	return parseABIFunction(selected, false, hasArrayVersion)
}

func parseABIFunction(entry ABIEntry, isConstructor bool, isOverloaded bool) *FunctionInfo {
	info := &FunctionInfo{
		SolidityName:    entry.Name,
		IsConstructor:   isConstructor,
		StateMutability: entry.StateMutability,
	}

	// Parse parameters
	for _, param := range entry.Inputs {
		info.Parameters = append(info.Parameters, ParameterInfo{
			Name:         param.Name,
			SolidityType: param.Type,
			GoType:       solidityToGoType(param.Type, param.InternalType),
		})
	}

	// Parse return parameters
	for _, param := range entry.Outputs {
		info.ReturnParams = append(info.ReturnParams, ParameterInfo{
			Name:         param.Name,
			SolidityType: param.Type,
			GoType:       solidityToGoType(param.Type, param.InternalType),
		})
	}

	// Determine operation type
	info.IsRead = info.StateMutability == "view" || info.StateMutability == "pure"
	info.IsWrite = !info.IsRead && !isConstructor

	// Generate Go name
	if !isConstructor {
		info.Name = capitalize(info.SolidityName)
		// Handle overloaded functions
		info.CallMethod = info.Name
		if isOverloaded {
			info.CallMethod = info.Name + "0"
		}
	}

	return info
}

func solidityToGoType(solidityType, internalType string) string {
	// Check for custom types first
	if strings.HasPrefix(internalType, "contract ") {
		return "common.Address"
	}

	// Handle bytes16 special case (often used for chain selectors/subjects)
	if solidityType == "bytes16" {
		// Check if it's a custom type like fastcurse.Subject
		if strings.Contains(internalType, "Subject") {
			return "fastcurse.Subject"
		}
		return "[16]byte"
	}

	// Handle array types
	if strings.HasSuffix(solidityType, "[]") {
		elemType := strings.TrimSuffix(solidityType, "[]")
		return "[]" + solidityToGoType(elemType, strings.TrimSuffix(internalType, "[]"))
	}

	// Map basic types
	typeMap := map[string]string{
		"uint8":   "uint8",
		"uint16":  "uint16",
		"uint32":  "uint32",
		"uint64":  "uint64",
		"uint256": "*big.Int",
		"int8":    "int8",
		"int16":   "int16",
		"int32":   "int32",
		"int64":   "int64",
		"int256":  "*big.Int",
		"address": "common.Address",
		"bool":    "bool",
		"string":  "string",
		"bytes":   "[]byte",
		"bytes32": "[32]byte",
	}

	if goType, ok := typeMap[solidityType]; ok {
		return goType
	}

	return "interface{}"
}

func toSnakeCase(s string) string {
	// Handle acronyms properly: RMNRemote -> rmn_remote
	var result []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if i > 0 && r >= 'A' && r <= 'Z' {
			// Check if this is the start of a new word
			// Add underscore if previous char was lowercase OR if next char is lowercase
			prevLower := runes[i-1] >= 'a' && runes[i-1] <= 'z'
			nextLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
			if prevLower || nextLower {
				result = append(result, '_')
			}
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func toKebabCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '-')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func makeVarName(contractName string) string {
	if len(contractName) == 0 {
		return ""
	}
	return strings.ToLower(string(contractName[0])) + contractName[1:]
}

func generateOperationsFile(info *ContractInfo) error {
	tmpl := `package {{.PackageName}}

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"{{.GobindingImport}}"
{{- if .NeedsFastcurse}}
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
{{- end}}
)

var ContractType cldf_deployment.ContractType = "{{.ContractType}}"
var Version = semver.MustParse("{{.Version}}")

{{- if .Constructor}}

type ConstructorArgs struct {
{{- range .Constructor.Parameters}}
	{{.GoName}} {{.GoType}}
{{- end}}
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "{{.PackageNameHyphen}}:deploy",
	Version:          semver.MustParse("{{.Version}}"),
	Description:      "Deploys the {{.ContractType}} contract",
	ContractMetadata: {{.GobindingPrefix}}.{{.ContractType}}MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("{{.Version}}")).String(): {
			EVM: common.FromHex({{.GobindingPrefix}}.{{.ContractType}}Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
{{- end}}

{{- range .WriteOps}}

type {{.ArgsStructName}} struct {
{{- range .Parameters}}
	{{.GoName}} {{.GoType}}
{{- end}}
}

var {{.Name}} = contract.NewWrite(contract.WriteParams[{{.ArgsStructName}}, *{{$.GobindingPrefix}}.{{$.ContractType}}]{
	Name:            "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:         semver.MustParse("{{$.Version}}"),
	Description:     "{{.Description}}",
	ContractType:    ContractType,
	ContractABI:     {{$.GobindingPrefix}}.{{$.ContractType}}ABI,
	NewContract:     {{$.GobindingPrefix}}.New{{$.ContractType}},
	IsAllowedCaller: contract.{{.AccessControl}}[*{{$.GobindingPrefix}}.{{$.ContractType}}, {{.ArgsStructName}}],
	Validate:        func({{.ArgsStructName}}) error { return nil },
	CallContract: func({{$.ContractVarName}} *{{$.GobindingPrefix}}.{{$.ContractType}}, opts *bind.TransactOpts, args {{.ArgsStructName}}) (*types.Transaction, error) {
		return {{$.ContractVarName}}.{{.CallMethod}}(opts{{.CallArgs}})
	},
})
{{- end}}

{{- range .ReadOps}}

var {{.Name}} = contract.NewRead(contract.ReadParams[{{.ArgsType}}, {{.ReturnType}}, *{{$.GobindingPrefix}}.{{$.ContractType}}]{
	Name:         "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:      semver.MustParse("{{$.Version}}"),
	Description:  "{{.Description}}",
	ContractType: ContractType,
	NewContract:  {{$.GobindingPrefix}}.New{{$.ContractType}},
	CallContract: func({{$.ContractVarName}} *{{$.GobindingPrefix}}.{{$.ContractType}}, opts *bind.CallOpts, args {{.ArgsType}}) ({{.ReturnType}}, error) {
		return {{$.ContractVarName}}.{{.CallMethod}}(opts, args)
	},
})
{{- end}}
`

	data := prepareTemplateData(info)

	t, err := template.New("operations").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("template execution error: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Generated code (unformatted):\n%s\n", buf.String())
		return fmt.Errorf("formatting failed: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(info.OutputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(info.OutputPath, formatted, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func prepareTemplateData(info *ContractInfo) map[string]interface{} {
	data := map[string]interface{}{
		"PackageName":       info.PackageName,
		"PackageNameHyphen": strings.ReplaceAll(info.PackageName, "_", "-"),
		"ContractType":      info.Name,
		"Version":           info.Version,
		"GobindingImport":   info.GobindingImport,
		"GobindingPrefix":   info.GobindingPrefix,
		"ContractVarName":   makeVarName(info.Name),
	}

	// Get sorted function names for deterministic ordering
	funcNames := make([]string, 0, len(info.Functions))
	for name := range info.Functions {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)

	// Check if we need fastcurse import
	needsFastcurse := false
	for _, name := range funcNames {
		funcInfo := info.Functions[name]
		for _, param := range funcInfo.Parameters {
			if strings.Contains(param.GoType, "fastcurse") {
				needsFastcurse = true
				break
			}
		}
		for _, param := range funcInfo.ReturnParams {
			if strings.Contains(param.GoType, "fastcurse") {
				needsFastcurse = true
				break
			}
		}
		if needsFastcurse {
			break
		}
	}
	data["NeedsFastcurse"] = needsFastcurse

	// Add constructor
	if info.Constructor != nil {
		data["Constructor"] = map[string]interface{}{
			"Parameters": prepareParameters(info.Constructor.Parameters),
		}
	}

	// Prepare write operations (in sorted order)
	var writeOps []map[string]interface{}
	for _, name := range funcNames {
		funcInfo := info.Functions[name]
		if funcInfo.IsWrite {
			writeOps = append(writeOps, prepareWriteOp(funcInfo))
		}
	}
	data["WriteOps"] = writeOps

	// Prepare read operations (in sorted order)
	var readOps []map[string]interface{}
	for _, name := range funcNames {
		funcInfo := info.Functions[name]
		if funcInfo.IsRead {
			readOps = append(readOps, prepareReadOp(funcInfo))
		}
	}
	data["ReadOps"] = readOps

	return data
}

func prepareParameters(params []ParameterInfo) []map[string]string {
	var result []map[string]string
	for _, p := range params {
		result = append(result, map[string]string{
			"GoName": capitalize(p.Name),
			"GoType": p.GoType,
		})
	}
	return result
}

func prepareWriteOp(funcInfo *FunctionInfo) map[string]interface{} {
	argsStructName := funcInfo.Name + "Args"

	accessControl := "AllCallersAllowed"
	if funcInfo.HasOnlyOwner {
		accessControl = "OnlyOwner"
	}

	var callArgs string
	if len(funcInfo.Parameters) > 0 {
		var args []string
		for _, p := range funcInfo.Parameters {
			args = append(args, "args."+capitalize(p.Name))
		}
		callArgs = ", " + strings.Join(args, ", ")
	}

	return map[string]interface{}{
		"Name":           funcInfo.Name,
		"OpName":         toKebabCase(funcInfo.SolidityName),
		"ArgsStructName": argsStructName,
		"Parameters":     prepareParameters(funcInfo.Parameters),
		"Description":    fmt.Sprintf("Calls %s on the contract", funcInfo.SolidityName),
		"AccessControl":  accessControl,
		"CallMethod":     funcInfo.CallMethod,
		"CallArgs":       callArgs,
	}
}

func prepareReadOp(funcInfo *FunctionInfo) map[string]interface{} {
	argsType := "struct{}"
	if len(funcInfo.Parameters) == 1 {
		argsType = funcInfo.Parameters[0].GoType
	}

	returnType := "interface{}"
	if len(funcInfo.ReturnParams) == 1 {
		returnType = funcInfo.ReturnParams[0].GoType
	}

	return map[string]interface{}{
		"Name":        funcInfo.Name,
		"OpName":      toKebabCase(funcInfo.SolidityName),
		"ArgsType":    argsType,
		"ReturnType":  returnType,
		"Description": fmt.Sprintf("Calls %s on the contract", funcInfo.SolidityName),
		"CallMethod":  funcInfo.CallMethod,
	}
}
