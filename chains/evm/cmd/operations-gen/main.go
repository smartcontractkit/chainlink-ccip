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
	BasePath string `yaml:"base_path"`
}

type SimpleContractConfig struct {
	ContractName string   `yaml:"contract_name"`
	Version      string   `yaml:"version"`
	Functions    []string `yaml:"functions"`
}

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
	Components   []ABIParam `json:"components"` // For tuple types (structs)
}

type ContractInfo struct {
	Name            string
	Version         string
	PackageName     string
	OutputPath      string
	GobindingImport string
	GobindingPrefix string
	ABI             string // JSON ABI string
	Bytecode        string // Hex bytecode string
	Constructor     *FunctionInfo
	Functions       map[string]*FunctionInfo
	FunctionOrder   []string              // Preserve order from config
	StructDefs      map[string]*StructDef // Local struct definitions (key is struct name)
}

type StructDef struct {
	Name            string // Local name (e.g., "DestChainConfigArgs")
	GethwrapperType string // Full gethwrapper type (e.g., "fee_quoter.FeeQuoterDestChainConfigArgs")
	Fields          []ParameterInfo
	NeedsConversion bool // True if this struct needs a conversion function
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
	Name            string
	SolidityType    string
	GoType          string
	IsStruct        bool            // True if this is a custom struct type
	StructName      string          // Name of the struct (without package prefix)
	GethwrapperType string          // Full gethwrapper type if IsStruct (e.g., "fee_quoter.FeeQuoterDestChainConfigArgs")
	Components      []ParameterInfo // For struct fields
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
			fmt.Fprintf(os.Stderr, "Error extracting info for %s: %v\n", contractCfg.ContractName, err)
			os.Exit(1)
		}

		if err := generateOperationsFile(contractInfo); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating operations for %s: %v\n", contractInfo.Name, err)
			os.Exit(1)
		}

		fmt.Printf("✓ Generated operations for %s at %s\n", contractInfo.Name, contractInfo.OutputPath)
	}
}

// versionToPath converts a semver version to a directory path format
// e.g., "1.6.0" -> "v1_6_0"
func versionToPath(version string) string {
	return "v" + strings.ReplaceAll(version, ".", "_")
}

func findGobindingPath(packageName, contractVersion string) (path string, actualPackageName string, version string, err error) {
	baseDir := filepath.Join("chains", "evm", "gobindings", "generated")

	// Try in order of preference:
	// 1. Contract version converted to directory name (e.g., 1.5.0 -> v1_5_0)
	// 2. "latest"

	var candidateVersions []string

	if contractVersion != "" {
		candidateVersions = append(candidateVersions, versionToPath(contractVersion))
	}

	candidateVersions = append(candidateVersions, "latest")

	// Also try package name without underscores (e.g., "onramp" instead of "on_ramp")
	packageNameNoUnderscore := strings.ReplaceAll(packageName, "_", "")

	for _, ver := range candidateVersions {
		// Try with original package name
		candidate := filepath.Join(baseDir, ver, packageName, packageName+".go")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, packageName, ver, nil
		}

		// Try without underscores
		if packageNameNoUnderscore != packageName {
			candidate = filepath.Join(baseDir, ver, packageNameNoUnderscore, packageNameNoUnderscore+".go")
			if _, err := os.Stat(candidate); err == nil {
				return candidate, packageNameNoUnderscore, ver, nil
			}
		}
	}

	return "", "", "", fmt.Errorf("gobinding not found for package %s (tried versions: %s, with/without underscores)",
		packageName, strings.Join(candidateVersions, ", "))
}

func extractContractInfo(cfg SimpleContractConfig, output OutputConfig) (*ContractInfo, error) {
	contractName := cfg.ContractName
	version := cfg.Version

	if contractName == "" {
		return nil, fmt.Errorf("contract_name is required")
	}
	if version == "" {
		return nil, fmt.Errorf("version is required")
	}

	packageName := toSnakeCase(contractName)

	// Find the correct gobinding path using version from config
	gobindingPath, actualPackageName, gobindingVersion, err := findGobindingPath(packageName, version)
	if err != nil {
		return nil, fmt.Errorf("failed to find gobinding: %w", err)
	}

	// Use the actual package name that was found (e.g., "onramp" not "on_ramp")
	packageName = actualPackageName

	// Extract ABI and Bytecode from gobinding
	abiString, bytecode, err := extractABIAndBytecodeFromGobinding(gobindingPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract ABI and bytecode: %w", err)
	}

	// Parse ABI
	var abiEntries []ABIEntry
	if err := json.Unmarshal([]byte(abiString), &abiEntries); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	info := &ContractInfo{
		Name:            contractName,
		Version:         version,
		PackageName:     packageName,
		GobindingPrefix: packageName,
		ABI:             abiString,
		Bytecode:        bytecode,
		Functions:       make(map[string]*FunctionInfo),
		StructDefs:      make(map[string]*StructDef),
	}

	// Convert version to path format (e.g., "1.6.0" -> "v1_6_0")
	versionPath := versionToPath(version)
	info.OutputPath = filepath.Join(output.BasePath, versionPath, "operations", info.PackageName, info.PackageName+".go")
	info.GobindingImport = fmt.Sprintf("github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/%s/%s",
		gobindingVersion, info.GobindingPrefix)

	// Detect ownership pattern from ABI (presence of owner(), transferOwnership(), etc.)
	isOwnable := detectOwnableFromABI(abiEntries)

	// Extract constructor
	for _, entry := range abiEntries {
		if entry.Type == "constructor" {
			info.Constructor = parseABIFunction(entry, true, false, 0, info.GobindingPrefix)
			break
		}
	}

	// Extract requested functions (preserve order from config)
	info.FunctionOrder = cfg.Functions // Store the order
	for _, funcName := range cfg.Functions {
		funcInfo := findFunctionInABI(abiEntries, funcName, info.GobindingPrefix)
		if funcInfo == nil {
			return nil, fmt.Errorf("function %s not found in ABI", funcName)
		}

		// If contract is Ownable and this is a write function, mark it as owner-only
		// For read functions or non-ownable contracts, allow all callers
		if isOwnable && funcInfo.IsWrite {
			funcInfo.HasOnlyOwner = true
		}

		info.Functions[funcName] = funcInfo
	}

	// Collect all struct definitions from constructor and functions
	if info.Constructor != nil {
		collectStructDefs(info.Constructor.Parameters, info.StructDefs)
	}
	for _, funcInfo := range info.Functions {
		collectStructDefs(funcInfo.Parameters, info.StructDefs)
		collectStructDefs(funcInfo.ReturnParams, info.StructDefs)
	}

	return info, nil
}

func collectStructDefs(params []ParameterInfo, structDefs map[string]*StructDef) {
	for _, param := range params {
		if param.IsStruct && param.StructName != "" {
			// Add struct definition if not already present
			if _, exists := structDefs[param.StructName]; !exists {
				structDefs[param.StructName] = &StructDef{
					Name:            param.StructName,
					GethwrapperType: param.GethwrapperType,
					Fields:          param.Components,
				}
			}
			// Recursively collect nested struct definitions
			collectStructDefs(param.Components, structDefs)
		}
	}
}

func extractABIAndBytecodeFromGobinding(path string) (string, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	content := string(data)

	// Find the ABI string in the MetaData
	abiRe := regexp.MustCompile(`ABI: "((?:[^"\\]|\\.)*)"`)
	abiMatches := abiRe.FindStringSubmatch(content)
	if len(abiMatches) < 2 {
		return "", "", fmt.Errorf("ABI not found in gobinding")
	}

	// Unescape the JSON string
	abi := abiMatches[1]
	abi = strings.ReplaceAll(abi, `\"`, `"`)
	abi = strings.ReplaceAll(abi, `\\`, `\`)

	// Find the Bin string (bytecode)
	binRe := regexp.MustCompile(`Bin: "(0x[0-9a-fA-F]*)"`)
	binMatches := binRe.FindStringSubmatch(content)
	if len(binMatches) < 2 {
		return "", "", fmt.Errorf("Bin not found in gobinding")
	}

	bytecode := binMatches[1]

	return abi, bytecode, nil
}

func extractVersionFromSolidity(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	// Match both constant declarations and function returns:
	// string public constant override typeAndVersion = "ContractName 1.6.0";
	// function typeAndVersion() ... { return "ContractName 1.5.0"; }
	// Use (?s) to make . match newlines, [\s\S]*? to match anything including newlines
	re := regexp.MustCompile(`(?s)typeAndVersion[\s\S]*?(?:=|return)\s*"[^"]+\s+(\d+\.\d+\.\d+)"`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// detectOwnableFromABI checks if the contract follows the Ownable pattern
// by looking for standard ownership functions in the ABI
func detectOwnableFromABI(entries []ABIEntry) bool {
	hasOwner := false
	hasTransferOwnership := false

	for _, entry := range entries {
		if entry.Type != "function" {
			continue
		}

		// Check for owner() function
		if entry.Name == "owner" && len(entry.Inputs) == 0 {
			hasOwner = true
		}

		// Check for transferOwnership(address) function
		if entry.Name == "transferOwnership" && len(entry.Inputs) == 1 {
			hasTransferOwnership = true
		}

		// If we found both, it's an Ownable contract
		if hasOwner && hasTransferOwnership {
			return true
		}
	}

	return false
}

func detectOwnershipModifiers(path string) map[string]bool {
	result := make(map[string]bool)
	visited := make(map[string]bool)

	// Recursively check this file and all inherited contracts
	detectOwnershipModifiersRecursive(path, result, visited)

	return result
}

func detectOwnershipModifiersRecursive(path string, result map[string]bool, visited map[string]bool) {
	// Avoid infinite loops
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	if visited[absPath] {
		return
	}
	visited[absPath] = true

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	source := string(data)

	// Find functions with onlyOwner modifier in this file
	re := regexp.MustCompile(`function\s+(\w+)\s*\([^)]*\)[^{]*\bonlyOwner\b`)
	matches := re.FindAllStringSubmatch(source, -1)
	for _, match := range matches {
		if len(match) > 1 {
			result[match[1]] = true
		}
	}

	// Find inherited contracts and process them
	// Match: contract ContractName is Parent1, Parent2 {
	inheritRe := regexp.MustCompile(`contract\s+\w+\s+is\s+([\w\s,]+)\s*\{`)
	inheritMatches := inheritRe.FindStringSubmatch(source)
	if len(inheritMatches) > 1 {
		parents := strings.Split(inheritMatches[1], ",")
		for _, parent := range parents {
			parent = strings.TrimSpace(parent)
			if parent == "" {
				continue
			}

			// Find the import for this parent
			parentPath := findImportPathForContract(source, parent, filepath.Dir(path))
			if parentPath != "" {
				detectOwnershipModifiersRecursive(parentPath, result, visited)
			}
		}
	}
}

func findImportPathForContract(source, contractName, baseDir string) string {
	// Find import statements: import {ContractName} from "path";
	// or: import "path";

	// Try: import {X, ContractName, Y} from "path";
	re1 := regexp.MustCompile(`import\s*\{[^}]*\b` + regexp.QuoteMeta(contractName) + `\b[^}]*\}\s*from\s*"([^"]+)"`)
	matches := re1.FindStringSubmatch(source)
	if len(matches) > 1 {
		return resolveImportPath(matches[1], baseDir)
	}

	// Try: import "path/ContractName.sol";
	re2 := regexp.MustCompile(`import\s*"([^"]*` + regexp.QuoteMeta(contractName) + `\.sol)"`)
	matches = re2.FindStringSubmatch(source)
	if len(matches) > 1 {
		return resolveImportPath(matches[1], baseDir)
	}

	return ""
}

func resolveImportPath(importPath, baseDir string) string {
	// Handle relative imports
	if strings.HasPrefix(importPath, ".") {
		resolved := filepath.Join(baseDir, importPath)
		if _, err := os.Stat(resolved); err == nil {
			return resolved
		}
	}

	// Handle absolute imports from project root
	// Try from chains/evm/contracts
	contractsBase := "chains/evm/contracts"
	resolved := filepath.Join(contractsBase, importPath)
	if _, err := os.Stat(resolved); err == nil {
		return resolved
	}

	// Handle @chainlink imports (look in node_modules or contracts)
	if strings.HasPrefix(importPath, "@chainlink/") {
		// These are usually in node_modules or vendored, skip for now
		return ""
	}

	// Handle @openzeppelin imports (skip)
	if strings.HasPrefix(importPath, "@openzeppelin/") {
		return ""
	}

	return ""
}

func findFunctionInABI(entries []ABIEntry, funcName string, gobindingPrefix string) *FunctionInfo {
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

	// Check if truly overloaded (multiple versions exist)
	isOverloaded := len(candidates) > 1

	// Select the best version when overloaded
	var selected ABIEntry
	var selectedIndex int
	if isOverloaded {
		// Preference order:
		// 1. Version with array parameters (more flexible)
		// 2. Version with MORE parameters (more specific)
		// 3. First candidate

		var arrayVersion ABIEntry
		var arrayIndex int
		var maxParams int
		var maxParamVersion ABIEntry
		var maxParamIndex int

		for i, c := range candidates {
			// Check for array parameters
			hasArray := false
			for _, param := range c.Inputs {
				if strings.Contains(param.Type, "[]") {
					hasArray = true
					break
				}
			}
			if hasArray && arrayVersion.Name == "" {
				arrayVersion = c
				arrayIndex = i
			}

			// Track version with most parameters
			if len(c.Inputs) > maxParams {
				maxParams = len(c.Inputs)
				maxParamVersion = c
				maxParamIndex = i
			}
		}

		// Select in priority order
		if arrayVersion.Name != "" {
			selected = arrayVersion
			selectedIndex = arrayIndex
		} else if maxParamVersion.Name != "" {
			selected = maxParamVersion
			selectedIndex = maxParamIndex
		} else {
			selected = candidates[0]
			selectedIndex = 0
		}
	} else {
		selected = candidates[0]
		selectedIndex = 0
	}

	// Only add suffix if we selected a non-first overload
	// (abigen adds suffixes to 2nd, 3rd, etc. versions: func0, func1, ...)
	needsSuffix := isOverloaded && selectedIndex > 0

	return parseABIFunction(selected, false, needsSuffix, selectedIndex, gobindingPrefix)
}

func parseABIFunction(entry ABIEntry, isConstructor bool, needsSuffix bool, suffixIndex int, gobindingPrefix string) *FunctionInfo {
	info := &FunctionInfo{
		SolidityName:    entry.Name,
		IsConstructor:   isConstructor,
		StateMutability: entry.StateMutability,
	}

	// Parse parameters
	for _, param := range entry.Inputs {
		info.Parameters = append(info.Parameters, parseABIParam(param, gobindingPrefix))
	}

	// Parse return parameters
	for _, param := range entry.Outputs {
		info.ReturnParams = append(info.ReturnParams, parseABIParam(param, gobindingPrefix))
	}

	// Determine operation type
	info.IsRead = info.StateMutability == "view" || info.StateMutability == "pure"
	info.IsWrite = !info.IsRead && !isConstructor

	// Generate Go name
	if !isConstructor {
		info.Name = capitalize(info.SolidityName)
		// Determine call method (handle overloaded functions)
		info.CallMethod = info.Name
		if needsSuffix {
			// abigen adds numeric suffixes: first overload gets no suffix, second gets "0", third gets "1", etc.
			// But we're given the index in our selected candidates, so we use that
			if suffixIndex > 0 {
				info.CallMethod = fmt.Sprintf("%s%d", info.Name, suffixIndex-1)
			}
		}
	}

	return info
}

func solidityToGoType(solidityType, internalType, gobindingPrefix string) string {
	// Check for custom types first
	if strings.HasPrefix(internalType, "contract ") {
		return "common.Address"
	}

	// Handle struct types (tuple in Solidity ABI)
	if solidityType == "tuple" || solidityType == "tuple[]" {
		// Extract the Go type name from internalType
		// Format can be either "struct ContractName.StructName" or "structContractName.StructName" (no space)
		structPrefix := "struct "
		if strings.HasPrefix(internalType, "struct") && !strings.HasPrefix(internalType, "struct ") {
			structPrefix = "struct"
		}

		if strings.HasPrefix(internalType, structPrefix) {
			// Handle both "structFeeQuoter.Foo[]" and "struct FeeQuoter.Foo"
			typeStr := strings.TrimPrefix(internalType, structPrefix)
			// Remove trailing [] if present
			typeStr = strings.TrimSuffix(typeStr, "[]")

			parts := strings.Split(typeStr, ".")
			if len(parts) == 2 {
				// Convert "FeeQuoter.DestChainConfigArgs" to "fee_quoter.FeeQuoterDestChainConfigArgs"
				goTypeName := gobindingPrefix + "." + parts[0] + parts[1]
				if solidityType == "tuple[]" {
					return "[]" + goTypeName
				}
				return goTypeName
			}
		}
		// Fallback
		return "interface{}"
	}

	// Handle array types (must come after tuple check)
	if strings.HasSuffix(solidityType, "[]") {
		elemType := strings.TrimSuffix(solidityType, "[]")
		return "[]" + solidityToGoType(elemType, strings.TrimSuffix(internalType, "[]"), gobindingPrefix)
	}

	// Map basic types
	typeMap := map[string]string{
		"uint8":   "uint8",
		"uint16":  "uint16",
		"uint24":  "uint32", // Go doesn't have uint24, use uint32
		"uint32":  "uint32",
		"uint48":  "uint64", // Go doesn't have uint48, use uint64
		"uint64":  "uint64",
		"uint96":  "*big.Int", // Larger than uint64
		"uint128": "*big.Int",
		"uint160": "*big.Int",
		"uint192": "*big.Int",
		"uint224": "*big.Int",
		"uint256": "*big.Int",
		"int8":    "int8",
		"int16":   "int16",
		"int24":   "int32",
		"int32":   "int32",
		"int48":   "int64",
		"int64":   "int64",
		"int96":   "*big.Int",
		"int128":  "*big.Int",
		"int160":  "*big.Int",
		"int192":  "*big.Int",
		"int224":  "*big.Int",
		"int256":  "*big.Int",
		"address": "common.Address",
		"bool":    "bool",
		"string":  "string",
		"bytes":   "[]byte",
		"bytes4":  "[4]byte",
		"bytes16": "[16]byte",
		"bytes32": "[32]byte",
	}

	if goType, ok := typeMap[solidityType]; ok {
		return goType
	}

	return "interface{}"
}

func parseABIParam(param ABIParam, gobindingPrefix string) ParameterInfo {
	paramInfo := ParameterInfo{
		Name:         param.Name,
		SolidityType: param.Type,
	}

	// Check if this is a struct type (tuple in ABI)
	baseType := strings.TrimSuffix(param.Type, "[]")
	isArray := strings.HasSuffix(param.Type, "[]")

	if baseType == "tuple" && len(param.Components) > 0 {
		// Extract struct name from internalType
		// Format: "struct ContractName.StructName" or "struct ContractName.StructName[]"
		structPrefix := "struct "
		if strings.HasPrefix(param.InternalType, "struct") && !strings.HasPrefix(param.InternalType, "struct ") {
			structPrefix = "struct"
		}

		internalType := strings.TrimPrefix(param.InternalType, structPrefix)
		internalType = strings.TrimSuffix(internalType, "[]")

		parts := strings.Split(internalType, ".")
		if len(parts) == 2 {
			// Local name: just the struct name (e.g., "DestChainConfigArgs")
			localStructName := parts[1]

			// Gethwrapper type: prefix.ContractNameStructName (e.g., "fee_quoter.FeeQuoterDestChainConfigArgs")
			gethwrapperType := gobindingPrefix + "." + parts[0] + parts[1]

			paramInfo.IsStruct = true
			paramInfo.StructName = localStructName
			paramInfo.GethwrapperType = gethwrapperType

			// Recursively parse struct fields
			for _, comp := range param.Components {
				paramInfo.Components = append(paramInfo.Components, parseABIParam(comp, gobindingPrefix))
			}

			// Set GoType
			if isArray {
				paramInfo.GoType = "[]" + localStructName
			} else {
				paramInfo.GoType = localStructName
			}
		} else {
			// Fallback to solidityToGoType
			paramInfo.GoType = solidityToGoType(param.Type, param.InternalType, gobindingPrefix)
		}
	} else {
		// Not a struct, use solidityToGoType
		paramInfo.GoType = solidityToGoType(param.Type, param.InternalType, gobindingPrefix)
	}

	return paramInfo
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

// {{.ContractType}}ABI is the ABI for the {{.ContractType}} contract
const {{.ContractType}}ABI = ` + "`" + `{{.ABI}}` + "`" + `

// {{.ContractType}}Bin is the bytecode for the {{.ContractType}} contract
const {{.ContractType}}Bin = "{{.Bytecode}}"

// {{.ContractType}}Contract is a local wrapper for {{.ContractType}} contract interactions
type {{.ContractType}}Contract struct {
	address  common.Address
	abi      abi.ABI
	backend  bind.ContractBackend
	contract *bind.BoundContract
}

// New{{.ContractType}}Contract creates a new {{.ContractType}}Contract wrapper
func New{{.ContractType}}Contract(address common.Address, backend bind.ContractBackend) (*{{.ContractType}}Contract, error) {
	parsed, err := abi.JSON(strings.NewReader({{.ContractType}}ABI))
	if err != nil {
		return nil, err
	}
	contract := bind.NewBoundContract(address, parsed, backend, backend, backend)
	return &{{.ContractType}}Contract{
		address:  address,
		abi:      parsed,
		backend:  backend,
		contract: contract,
	}, nil
}

// Address returns the contract address
func (c *{{.ContractType}}Contract) Address() common.Address {
	return c.address
}

// Owner returns the contract owner (if applicable)
func (c *{{.ContractType}}Contract) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

{{- range .ContractMethods}}

// {{.Name}} calls {{.MethodName}} on the contract
func (c *{{$.ContractType}}Contract) {{.Name}}({{.Params}}) {{.Returns}} {
	{{.MethodBody}}
}
{{- end}}

{{- range .StructDefs}}

// {{.Name}} represents {{.GethwrapperType}} structure
type {{.Name}} struct {
{{- range .Fields}}
	{{.GoName}} {{.GoType}}
{{- end}}
}
{{- end}}

{{- if .Constructor}}

type ConstructorArgs struct {
{{- range .Constructor.Parameters}}
	{{.GoName}} {{.GoType}}
{{- end}}
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "{{.PackageNameHyphen}}:deploy",
	Version:          Version,
	Description:      "Deploys the {{.ContractType}} contract",
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
{{- end}}

{{- range .WriteOps}}
{{- if and (not .UseSingleArg) (not .UseNoArgs)}}

type {{.ArgsStructName}} struct {
{{- range .Parameters}}
	{{.GoName}} {{.GoType}}
{{- end}}
}
{{- end}}

var {{.Name}} = contract.NewWrite(contract.WriteParams[{{.ArgsType}}, *{{$.ContractType}}Contract]{
	Name:            "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:         Version,
	Description:     "{{.Description}}",
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

{{- range .ReadOps}}

var {{.Name}} = contract.NewRead(contract.ReadParams[{{.ArgsType}}, {{.ReturnType}}, *{{$.ContractType}}Contract]{
	Name:         "{{$.PackageNameHyphen}}:{{.OpName}}",
	Version:      Version,
	Description:  "{{.Description}}",
	ContractType: ContractType,
	NewContract:  New{{$.ContractType}}Contract,
	CallContract: func(c *{{$.ContractType}}Contract, opts *bind.CallOpts, args {{.ArgsType}}) ({{.ReturnType}}, error) {
		return c.{{.Name}}(opts, {{.CallArg}})
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
		_, err := fmt.Fprintf(os.Stderr, "Generated code (unformatted):\n%s\n", buf.String())
		if err != nil {
			return err
		}
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
		"ABI":               info.ABI,
		"Bytecode":          info.Bytecode,
	}

	// Use the order from config (already preserved in FunctionOrder)
	funcNames := info.FunctionOrder

	// Check if we need big.Int import
	needsBigInt := false

	// Check functions
	for _, name := range funcNames {
		funcInfo := info.Functions[name]
		for _, param := range funcInfo.Parameters {
			if strings.Contains(param.GoType, "*big.Int") {
				needsBigInt = true
			}
		}
		for _, param := range funcInfo.ReturnParams {
			if strings.Contains(param.GoType, "*big.Int") {
				needsBigInt = true
			}
		}
	}

	// Check structs
	for _, structDef := range info.StructDefs {
		for _, field := range structDef.Fields {
			if strings.Contains(field.GoType, "*big.Int") {
				needsBigInt = true
				break
			}
		}
	}

	// Check constructor
	if info.Constructor != nil {
		for _, param := range info.Constructor.Parameters {
			if strings.Contains(param.GoType, "*big.Int") {
				needsBigInt = true
				break
			}
		}
	}

	data["NeedsBigInt"] = needsBigInt

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

	// Prepare struct definitions
	var structDefs []map[string]interface{}
	for _, structDef := range info.StructDefs {
		structDefs = append(structDefs, map[string]interface{}{
			"Name":            structDef.Name,
			"GethwrapperType": structDef.GethwrapperType,
			"Fields":          prepareParameters(structDef.Fields),
		})
	}
	data["StructDefs"] = structDefs

	// Prepare contract wrapper methods
	var contractMethods []map[string]interface{}
	for _, name := range funcNames {
		funcInfo := info.Functions[name]
		if funcInfo.IsWrite {
			contractMethods = append(contractMethods, prepareContractMethod(funcInfo, true))
		} else if funcInfo.IsRead {
			contractMethods = append(contractMethods, prepareContractMethod(funcInfo, false))
		}
	}
	data["ContractMethods"] = contractMethods

	return data
}

func prepareContractMethod(funcInfo *FunctionInfo, isWrite bool) map[string]interface{} {
	// Params: opts *bind.TransactOpts, args <Type>
	// or: opts *bind.CallOpts, args <Type>
	optsType := "*bind.TransactOpts"
	if !isWrite {
		optsType = "*bind.CallOpts"
	}

	argsParam := ""
	if len(funcInfo.Parameters) == 1 {
		argsType := funcInfo.Parameters[0].GoType
		argsParam = fmt.Sprintf("args %s", argsType)
	} else if len(funcInfo.Parameters) > 1 {
		argsType := funcInfo.Name + "Args"
		argsParam = fmt.Sprintf("args %s", argsType)
	}

	params := fmt.Sprintf("opts %s", optsType)
	if argsParam != "" {
		params = fmt.Sprintf("%s, %s", params, argsParam)
	}

	// Returns
	returns := "(*types.Transaction, error)"
	returnType := "interface{}"
	if !isWrite {
		if len(funcInfo.ReturnParams) == 1 {
			returnType = funcInfo.ReturnParams[0].GoType
		}
		returns = fmt.Sprintf("(%s, error)", returnType)
	}

	// Build method body
	var methodBody string
	if isWrite {
		// Build call args for Transact - pass local types directly
		var callArgsList []string
		if len(funcInfo.Parameters) > 0 {
			if len(funcInfo.Parameters) == 1 {
				callArgsList = append(callArgsList, "args")
			} else {
				// Multiple parameters - extract fields
				for _, p := range funcInfo.Parameters {
					callArgsList = append(callArgsList, "args."+capitalize(p.Name))
				}
			}
		}
		callArgsStr := strings.Join(callArgsList, ", ")
		if callArgsStr != "" {
			methodBody = fmt.Sprintf("return c.contract.Transact(opts, \"%s\", %s)", funcInfo.CallMethod, callArgsStr)
		} else {
			methodBody = fmt.Sprintf("return c.contract.Transact(opts, \"%s\")", funcInfo.CallMethod)
		}
	} else {
		// Read operation - use Call
		var callArgsList []string
		if len(funcInfo.Parameters) > 0 {
			if len(funcInfo.Parameters) == 1 {
				callArgsList = append(callArgsList, "args")
			}
		}

		callArgsStr := ""
		if len(callArgsList) > 0 {
			callArgsStr = ", " + strings.Join(callArgsList, ", ")
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
	useSingleArg := len(funcInfo.Parameters) == 1
	useNoArgs := len(funcInfo.Parameters) == 0
	argsType := argsStructName

	accessControl := "AllCallersAllowed"
	if funcInfo.HasOnlyOwner {
		accessControl = "OnlyOwner"
	}

	var callArgs string
	if len(funcInfo.Parameters) > 0 {
		if useSingleArg {
			// For single parameter
			param := funcInfo.Parameters[0]
			argsType = param.GoType
			callArgs = ", args"
		} else {
			// For multiple parameters, just pass args (conversion in wrapper)
			callArgs = ", args"
		}
	} else {
		// For zero parameters, use struct{}
		argsType = "struct{}"
	}

	return map[string]interface{}{
		"Name":           funcInfo.Name,
		"OpName":         toKebabCase(funcInfo.SolidityName),
		"ArgsStructName": argsStructName,
		"ArgsType":       argsType,
		"UseSingleArg":   useSingleArg,
		"UseNoArgs":      useNoArgs,
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
		param := funcInfo.Parameters[0]
		argsType = param.GoType
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
		"CallArg":     "args",
		"Description": fmt.Sprintf("Calls %s on the contract", funcInfo.SolidityName),
		"CallMethod":  funcInfo.CallMethod,
	}
}
