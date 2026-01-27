package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Config structures
type Config struct {
	Version   string           `yaml:"version"`
	Output    OutputConfig     `yaml:"output"`
	Contracts []ContractConfig `yaml:"contracts"`
}

type OutputConfig struct {
	BasePath string `yaml:"base_path"`
}

type ContractConfig struct {
	Name      string           `yaml:"contract_name"`
	Version   string           `yaml:"version"`
	Functions []FunctionConfig `yaml:"functions"`
}

type FunctionConfig struct {
	Name   string `yaml:"name"`
	Access string `yaml:"access,omitempty"` // "owner", "public", or empty
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
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	InternalType string     `json:"internalType"`
	Components   []ABIParam `json:"components"`
}

// Contract info structures
type ContractInfo struct {
	Name          string
	Version       string
	PackageName   string
	OutputPath    string
	ABI           string
	Bytecode      string
	Constructor   *FunctionInfo
	Functions     map[string]*FunctionInfo
	FunctionOrder []string
	StructDefs    map[string]*StructDef
}

type StructDef struct {
	Name   string
	Fields []ParameterInfo
}

type FunctionInfo struct {
	Name            string
	StateMutability string
	Parameters      []ParameterInfo
	ReturnParams    []ParameterInfo
	IsWrite         bool
	CallMethod      string
	HasOnlyOwner    bool
}

type ParameterInfo struct {
	Name         string
	SolidityType string
	GoType       string
	IsStruct     bool
	StructName   string
	Components   []ParameterInfo
}

func main() {
	configPath := flag.String("config", "chains/evm/operations_gen_config.yaml", "Path to config file")
	flag.Parse()

	configData, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %v\n", err)
		os.Exit(1)
	}

	for _, contractCfg := range config.Contracts {
		info, err := extractContractInfo(contractCfg, config.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting info for %s: %v\n", contractCfg.Name, err)
			os.Exit(1)
		}

		if err := generateOperationsFile(info); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating file for %s: %v\n", contractCfg.Name, err)
			os.Exit(1)
		}

		fmt.Printf("✓ Generated operations for %s at %s\n", info.Name, info.OutputPath)
	}
}

func extractContractInfo(cfg ContractConfig, output OutputConfig) (*ContractInfo, error) {
	if cfg.Name == "" || cfg.Version == "" {
		return nil, fmt.Errorf("contract_name and version are required")
	}

	packageName := toSnakeCase(cfg.Name)
	versionPath := versionToPath(cfg.Version)

	abiString, bytecode, err := readABIAndBytecode(packageName, versionPath)
	if err != nil {
		return nil, err
	}

	var abiEntries []ABIEntry
	if err := json.Unmarshal([]byte(abiString), &abiEntries); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	info := &ContractInfo{
		Name:        cfg.Name,
		Version:     cfg.Version,
		PackageName: packageName,
		OutputPath:  filepath.Join(output.BasePath, versionPath, "operations", packageName, packageName+".go"),
		ABI:         abiString,
		Bytecode:    bytecode,
		Functions:   make(map[string]*FunctionInfo),
		StructDefs:  make(map[string]*StructDef),
	}

	extractConstructor(info, abiEntries)
	
	if err := extractFunctions(info, cfg.Functions, abiEntries); err != nil {
		return nil, err
	}

	collectAllStructDefs(info)
	return info, nil
}

func readABIAndBytecode(packageName, versionPath string) (string, string, error) {
	fileName := packageName
	abiPath := filepath.Join("chains", "evm", "abi", versionPath, fileName+".json")
	
	abiBytes, err := os.ReadFile(abiPath)
	if err != nil {
		fileNameNoUnderscore := strings.ReplaceAll(packageName, "_", "")
		abiPath = filepath.Join("chains", "evm", "abi", versionPath, fileNameNoUnderscore+".json")
		abiBytes, err = os.ReadFile(abiPath)
		if err != nil {
			return "", "", fmt.Errorf("ABI not found (tried %s.json and %s.json in %s)", 
				fileName, fileNameNoUnderscore, versionPath)
		}
		fileName = fileNameNoUnderscore
	}

	bytecodePath := filepath.Join("chains", "evm", "bytecode", versionPath, fileName+".bin")
	bytecodeBytes, err := os.ReadFile(bytecodePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read bytecode from %s: %w", bytecodePath, err)
	}

	return string(abiBytes), strings.TrimSpace(string(bytecodeBytes)), nil
}

func extractConstructor(info *ContractInfo, abiEntries []ABIEntry) {
	for _, entry := range abiEntries {
		if entry.Type == "constructor" {
			info.Constructor = parseABIFunction(entry, true, info.PackageName, false, 0)
			break
		}
	}
}

func extractFunctions(info *ContractInfo, funcConfigs []FunctionConfig, abiEntries []ABIEntry) error {
	functionNames := make([]string, len(funcConfigs))
	for i, cfg := range funcConfigs {
		functionNames[i] = cfg.Name
	}
	info.FunctionOrder = functionNames

	for _, funcCfg := range funcConfigs {
		funcInfo := findFunctionInABI(abiEntries, funcCfg.Name, info.PackageName)
		if funcInfo == nil {
			return fmt.Errorf("function %s not found in ABI", funcCfg.Name)
		}

		switch funcCfg.Access {
		case "owner":
			funcInfo.HasOnlyOwner = true
		case "public", "":
			funcInfo.HasOnlyOwner = false
		default:
			return fmt.Errorf("unknown access control '%s' for function %s (use 'owner' or 'public')", 
				funcCfg.Access, funcCfg.Name)
		}

		info.Functions[funcCfg.Name] = funcInfo
	}
	
	return nil
}

func collectAllStructDefs(info *ContractInfo) {
	if info.Constructor != nil {
		collectStructDefs(info.Constructor.Parameters, info.StructDefs)
	}
	for _, funcInfo := range info.Functions {
		collectStructDefs(funcInfo.Parameters, info.StructDefs)
		collectStructDefs(funcInfo.ReturnParams, info.StructDefs)
	}
}

func collectStructDefs(params []ParameterInfo, structDefs map[string]*StructDef) {
	for _, param := range params {
		if param.IsStruct && param.StructName != "" {
			if _, exists := structDefs[param.StructName]; !exists {
				structDefs[param.StructName] = &StructDef{
					Name:   param.StructName,
					Fields: param.Components,
				}
			}
			collectStructDefs(param.Components, structDefs)
		}
	}
}

func findFunctionInABI(entries []ABIEntry, funcName string, packageName string) *FunctionInfo {
	var candidates []ABIEntry
	for _, entry := range entries {
		if entry.Type == "function" && strings.EqualFold(entry.Name, funcName) {
			candidates = append(candidates, entry)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	if len(candidates) == 1 {
		return parseABIFunction(candidates[0], false, packageName, false, 0)
	}

	for i, candidate := range candidates {
		if len(candidate.Inputs) > 0 {
			needsSuffix := i > 0
			return parseABIFunction(candidate, false, packageName, needsSuffix, i)
		}
	}

	return parseABIFunction(candidates[0], false, packageName, false, 0)
}

func parseABIFunction(entry ABIEntry, isConstructor bool, packageName string, needsSuffix bool, suffixIndex int) *FunctionInfo {

	callMethod := entry.Name
	if needsSuffix {
		callMethod = fmt.Sprintf("%s%d", entry.Name, suffixIndex)
	}

	funcInfo := &FunctionInfo{
		Name:            capitalize(entry.Name),
		StateMutability: entry.StateMutability,
		CallMethod:      callMethod,
		IsWrite:         entry.StateMutability != "view" && entry.StateMutability != "pure",
	}

	for _, input := range entry.Inputs {
		funcInfo.Parameters = append(funcInfo.Parameters, parseABIParam(input, packageName))
	}

	for _, output := range entry.Outputs {
		funcInfo.ReturnParams = append(funcInfo.ReturnParams, parseABIParam(output, packageName))
	}

	return funcInfo
}

func parseABIParam(param ABIParam, packageName string) ParameterInfo {
	goType := solidityToGoType(param.Type, param.InternalType, packageName)
	
	paramInfo := ParameterInfo{
		Name:         param.Name,
		SolidityType: param.Type,
		GoType:       goType,
	}

	if strings.HasPrefix(param.Type, "tuple") {
		structName := extractStructName(param.InternalType)
		if structName != "" {
			paramInfo.IsStruct = true
			paramInfo.StructName = structName
			
			if strings.HasSuffix(param.Type, "[]") {
				paramInfo.GoType = "[]" + structName
			} else {
				paramInfo.GoType = structName
			}

			for _, comp := range param.Components {
				paramInfo.Components = append(paramInfo.Components, parseABIParam(comp, packageName))
			}
		}
	}

	return paramInfo
}

func solidityToGoType(solidityType, internalType, packageName string) string {
	typeMap := map[string]string{
		"address":   "common.Address",
		"string":    "string",
		"bool":      "bool",
		"bytes":     "[]byte",
		"bytes32":   "[32]byte",
		"bytes16":   "[16]byte",
		"bytes4":    "[4]byte",
		"uint8":     "uint8",
		"uint16":    "uint16",
		"uint32":    "uint32",
		"uint64":    "uint64",
		"uint96":    "*big.Int",
		"uint128":   "*big.Int",
		"uint160":   "*big.Int",
		"uint192":   "*big.Int",
		"uint224":   "*big.Int",
		"uint256":   "*big.Int",
		"int8":      "int8",
		"int16":     "int16",
		"int32":     "int32",
		"int64":     "int64",
		"int96":     "*big.Int",
		"int128":    "*big.Int",
		"int160":    "*big.Int",
		"int192":    "*big.Int",
		"int224":    "*big.Int",
		"int256":    "*big.Int",
	}

	baseType := strings.TrimSuffix(solidityType, "[]")
	if goType, ok := typeMap[baseType]; ok {
		if strings.HasSuffix(solidityType, "[]") {
			return "[]" + goType
		}
		return goType
	}

	if strings.HasPrefix(baseType, "tuple") {
		return "interface{}"
	}

	return "interface{}"
}

func extractStructName(internalType string) string {
	if internalType == "" {
		return ""
	}
	
	parts := strings.Split(internalType, ".")
	structName := parts[len(parts)-1]
	structName = strings.TrimSuffix(structName, "[]")
	
	if structName == "" || strings.HasPrefix(structName, "I") {
		return ""
	}
	
	return structName
}

func versionToPath(version string) string {
	return "v" + strings.ReplaceAll(version, ".", "_")
}

func toSnakeCase(s string) string {
	nameOverrides := map[string]string{
		"OnRamp":  "onramp",
		"OffRamp": "offramp",
	}
	
	if override, ok := nameOverrides[s]; ok {
		return override
	}
	
	var result []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if i > 0 && r >= 'A' && r <= 'Z' {
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
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func toKebabCase(s string) string {
	return strings.ReplaceAll(toSnakeCase(s), "_", "-")
}

func generateOperationsFile(info *ContractInfo) error {
	tmpl := `package {{.PackageName}}

import (
{{- if .NeedsBigInt}}
	"math/big"
{{end}}
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var ContractType cldf_deployment.ContractType = "{{.ContractType}}"
var Version = semver.MustParse("{{.Version}}")

const {{.ContractType}}ABI = ` + "`" + `{{.ABI}}` + "`" + `
const {{.ContractType}}Bin = "{{.Bytecode}}"

type {{.ContractType}}Contract struct {
	address  common.Address
	abi      abi.ABI
	backend  bind.ContractBackend
	contract *bind.BoundContract
}

func New{{.ContractType}}Contract(address common.Address, backend bind.ContractBackend) (*{{.ContractType}}Contract, error) {
	parsed, err := abi.JSON(strings.NewReader({{.ContractType}}ABI))
	if err != nil {
		return nil, err
	}
	return &{{.ContractType}}Contract{
		address:  address,
		abi:      parsed,
		backend:  backend,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (c *{{.ContractType}}Contract) Address() common.Address {
	return c.address
}

func (c *{{.ContractType}}Contract) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}
{{range .ContractMethods}}

func (c *{{$.ContractType}}Contract) {{.Name}}({{.Params}}) {{.Returns}} {
	{{.MethodBody}}
}
{{- end}}
{{range .StructDefs}}

type {{.Name}} struct {
{{- range .Fields}}
	{{.GoName}} {{.GoType}}
{{- end}}
}
{{- end}}
{{range .WriteArgStructs}}

type {{.Name}} struct {
{{- range .Fields}}
	{{.GoName}} {{.GoType}}
{{- end}}
}
{{- end}}
{{if .Constructor}}

type ConstructorArgs struct {
{{- range .Constructor.Parameters}}
	{{.GoName}} {{.GoType}}
{{- end}}
}
{{end}}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:        "{{.PackageNameHyphen}}:deploy",
	Version:     Version,
	Description: "Deploys the {{.ContractType}} contract",
	ContractMetadata: &bind.MetaData{
		ABI: {{.ContractType}}ABI,
		Bin: {{.ContractType}}Bin,
	},
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex({{.ContractType}}Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
{{range .WriteOps}}

var {{.Name}} = contract.NewWrite(contract.WriteParams[{{.ArgsType}}, *{{$.ContractType}}Contract]{
	Name:            "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:         Version,
	Description:     "Calls {{.MethodName}} on the contract",
	ContractType:    ContractType,
	ContractABI:     {{$.ContractType}}ABI,
	NewContract:     New{{$.ContractType}}Contract,
	IsAllowedCaller: contract.{{.AccessControl}}[*{{$.ContractType}}Contract, {{.ArgsType}}],
	Validate:        func({{.ArgsType}}) error { return nil },
	CallContract: func(c *{{$.ContractType}}Contract, opts *bind.TransactOpts, args {{.ArgsType}}) (*types.Transaction, error) {
		return c.{{.Name}}(opts{{.CallArgs}})
	},
})
{{- end}}
{{range .ReadOps}}

var {{.Name}} = contract.NewRead(contract.ReadParams[{{.ArgsType}}, {{.ReturnType}}, *{{$.ContractType}}Contract]{
	Name:         "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:      Version,
	Description:  "Calls {{.MethodName}} on the contract",
	ContractType: ContractType,
	NewContract:  New{{$.ContractType}}Contract,
	CallContract: func(c *{{$.ContractType}}Contract, opts *bind.CallOpts, args {{.ArgsType}}) ({{.ReturnType}}, error) {
		return c.{{.Name}}(opts{{.CallArgs}})
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
		return fmt.Errorf("formatting error: %w\n%s", err, buf.String())
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
		"PackageNameHyphen": toKebabCase(info.Name),
		"ContractType":      info.Name,
		"Version":           info.Version,
		"ABI":               info.ABI,
		"Bytecode":          info.Bytecode,
	}

	needsBigInt := checkNeedsBigInt(info)
	data["NeedsBigInt"] = needsBigInt

	if info.Constructor != nil {
		data["Constructor"] = map[string]interface{}{
			"Parameters": prepareParameters(info.Constructor.Parameters),
		}
	}

	var writeOps, readOps []map[string]interface{}
	var contractMethods []map[string]interface{}
	var writeArgStructs []map[string]interface{}

	for _, name := range info.FunctionOrder {
		funcInfo := info.Functions[name]
		contractMethods = append(contractMethods, prepareContractMethod(funcInfo, funcInfo.IsWrite))
		
		if funcInfo.IsWrite {
			writeOps = append(writeOps, prepareWriteOp(funcInfo))
			if len(funcInfo.Parameters) > 1 {
				writeArgStructs = append(writeArgStructs, map[string]interface{}{
					"Name":   funcInfo.Name + "Args",
					"Fields": prepareParameters(funcInfo.Parameters),
				})
			}
		} else {
			readOps = append(readOps, prepareReadOp(funcInfo))
		}
	}

	var structDefs []map[string]interface{}
	var structNames []string
	for name := range info.StructDefs {
		structNames = append(structNames, name)
	}
	sort.Strings(structNames)
	for _, name := range structNames {
		structDef := info.StructDefs[name]
		structDefs = append(structDefs, map[string]interface{}{
			"Name":   structDef.Name,
			"Fields": prepareParameters(structDef.Fields),
		})
	}

	data["StructDefs"] = structDefs
	data["WriteArgStructs"] = writeArgStructs
	data["WriteOps"] = writeOps
	data["ReadOps"] = readOps
	data["ContractMethods"] = contractMethods

	return data
}

func checkNeedsBigInt(info *ContractInfo) bool {
	check := func(params []ParameterInfo) bool {
		for _, p := range params {
			if strings.Contains(p.GoType, "*big.Int") {
				return true
			}
		}
		return false
	}

	for _, funcInfo := range info.Functions {
		if check(funcInfo.Parameters) || check(funcInfo.ReturnParams) {
			return true
		}
	}

	if info.Constructor != nil && check(info.Constructor.Parameters) {
		return true
	}

	for _, structDef := range info.StructDefs {
		if check(structDef.Fields) {
			return true
		}
	}

	return false
}

func prepareContractMethod(funcInfo *FunctionInfo, isWrite bool) map[string]interface{} {
	optsType := "*bind.CallOpts"
	if isWrite {
		optsType = "*bind.TransactOpts"
	}

	params := fmt.Sprintf("opts %s", optsType)
	var methodArgs []string
	
	if len(funcInfo.Parameters) == 1 {
		params += fmt.Sprintf(", args %s", funcInfo.Parameters[0].GoType)
		methodArgs = []string{"args"}
	} else if len(funcInfo.Parameters) > 1 {
		for _, p := range funcInfo.Parameters {
			paramName := strings.ToLower(p.Name[:1]) + p.Name[1:]
			if paramName == "" {
				paramName = "arg" + fmt.Sprint(len(methodArgs))
			}
			params += fmt.Sprintf(", %s %s", paramName, p.GoType)
			methodArgs = append(methodArgs, paramName)
		}
	}

	returns := "(*types.Transaction, error)"
	returnType := "interface{}"
	if !isWrite {
		if len(funcInfo.ReturnParams) == 1 {
			returnType = funcInfo.ReturnParams[0].GoType
		}
		returns = fmt.Sprintf("(%s, error)", returnType)
	}

	var methodBody string
	if isWrite {
		if len(methodArgs) > 0 {
			methodBody = fmt.Sprintf("return c.contract.Transact(opts, \"%s\", %s)", 
				funcInfo.CallMethod, strings.Join(methodArgs, ", "))
		} else {
			methodBody = fmt.Sprintf("return c.contract.Transact(opts, \"%s\")", funcInfo.CallMethod)
		}
	} else {
		callArgsStr := ""
		if len(methodArgs) > 0 {
			callArgsStr = ", " + strings.Join(methodArgs, ", ")
		}
		methodBody = fmt.Sprintf(`var out []interface{}
	err := c.contract.Call(opts, &out, "%s"%s)
	if err != nil {
		var zero %s
		return zero, err
	}
	return *abi.ConvertType(out[0], new(%s)).(*%s), nil`, funcInfo.CallMethod, callArgsStr, returnType, returnType, returnType)
	}

	return map[string]interface{}{
		"Name":       funcInfo.Name,
		"MethodName": funcInfo.CallMethod,
		"Params":     params,
		"Returns":    returns,
		"MethodBody": methodBody,
	}
}

func prepareParameters(params []ParameterInfo) []map[string]string {
	var result []map[string]string
	for _, param := range params {
		result = append(result, map[string]string{
			"GoName": capitalize(param.Name),
			"GoType": param.GoType,
		})
	}
	return result
}

func prepareWriteOp(funcInfo *FunctionInfo) map[string]interface{} {
	argsType := "struct{}"
	var callArgsList []string

	if len(funcInfo.Parameters) == 1 {
		argsType = funcInfo.Parameters[0].GoType
		callArgsList = []string{"args"}
	} else if len(funcInfo.Parameters) > 1 {
		argsType = funcInfo.Name + "Args"
		for _, p := range funcInfo.Parameters {
			callArgsList = append(callArgsList, "args."+capitalize(p.Name))
		}
	}

	callArgs := ""
	if len(callArgsList) > 0 {
		callArgs = ", " + strings.Join(callArgsList, ", ")
	}

	accessControl := "AllCallersAllowed"
	if funcInfo.HasOnlyOwner {
		accessControl = "OnlyOwner"
	}

	return map[string]interface{}{
		"Name":          funcInfo.Name,
		"MethodName":    funcInfo.CallMethod,
		"OpName":        toKebabCase(funcInfo.Name),
		"ArgsType":      argsType,
		"CallArgs":      callArgs,
		"AccessControl": accessControl,
	}
}

func prepareReadOp(funcInfo *FunctionInfo) map[string]interface{} {
	argsType := "struct{}"
	callArgs := ""

	if len(funcInfo.Parameters) == 1 {
		argsType = funcInfo.Parameters[0].GoType
		callArgs = ", args"
	} else if len(funcInfo.Parameters) > 1 {
		argsType = funcInfo.Name + "Args"
		var parts []string
		for _, p := range funcInfo.Parameters {
			parts = append(parts, ", args."+capitalize(p.Name))
		}
		callArgs = strings.Join(parts, "")
	}

	returnType := "interface{}"
	if len(funcInfo.ReturnParams) == 1 {
		returnType = funcInfo.ReturnParams[0].GoType
	}

	return map[string]interface{}{
		"Name":       funcInfo.Name,
		"MethodName": funcInfo.CallMethod,
		"OpName":     toKebabCase(funcInfo.Name),
		"ArgsType":   argsType,
		"ReturnType": returnType,
		"CallArgs":   callArgs,
	}
}
