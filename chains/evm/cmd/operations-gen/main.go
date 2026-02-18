package main

import (
	"bytes"
	_ "embed"
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

//go:embed operations.tmpl
var operationsTemplate string

const (
	// anyType is the fallback Go type for unknown Solidity types
	anyType = "any"
)

var (
	// typeMap maps Solidity types to their Go equivalents
	typeMap = map[string]string{
		"address": "common.Address",
		"string":  "string",
		"bool":    "bool",
		"bytes":   "[]byte",
		"bytes32": "[32]byte",
		"bytes16": "[16]byte",
		"bytes4":  "[4]byte",
		"uint8":   "uint8",
		"uint16":  "uint16",
		"uint32":  "uint32",
		"uint64":  "uint64",
		"uint96":  "*big.Int",
		"uint128": "*big.Int",
		"uint160": "*big.Int",
		"uint192": "*big.Int",
		"uint224": "*big.Int",
		"uint256": "*big.Int",
		"int8":    "int8",
		"int16":   "int16",
		"int32":   "int32",
		"int64":   "int64",
		"int96":   "*big.Int",
		"int128":  "*big.Int",
		"int160":  "*big.Int",
		"int192":  "*big.Int",
		"int224":  "*big.Int",
		"int256":  "*big.Int",
	}

	// nameOverrides provides special case naming for specific contracts
	nameOverrides = map[string]string{
		"OnRamp":  "onramp",
		"OffRamp": "offramp",
	}
)

// Config structures
type Config struct {
	Version   string           `yaml:"version"`
	Input     InputConfig      `yaml:"input"`
	Output    OutputConfig     `yaml:"output"`
	Contracts []ContractConfig `yaml:"contracts"`
}

type InputConfig struct {
	BasePath string `yaml:"base_path"`
}

type OutputConfig struct {
	BasePath string `yaml:"base_path"`
}

type ContractConfig struct {
	Name         string           `yaml:"contract_name"`
	Version      string           `yaml:"version"`
	PackageName  string           `yaml:"package_name,omitempty"`  // Optional: override package name
	ABIFile      string           `yaml:"abi_file,omitempty"`      // Optional: override ABI file name
	NoDeployment bool             `yaml:"no_deployment,omitempty"` // Optional: skip bytecode and deploy operation
	Functions    []FunctionConfig `yaml:"functions"`
}

type FunctionConfig struct {
	Name   string `yaml:"name"`
	Access string `yaml:"access,omitempty"` // "owner", "public", or empty
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
	Components   []ABIParam `json:"components"`
}

type ContractInfo struct {
	Name          string
	Version       string
	PackageName   string
	OutputPath    string
	ABI           string
	Bytecode      string
	NoDeployment  bool
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
	CallMethod      string // The method name or full signature for overloaded functions
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

type TemplateData struct {
	PackageName       string
	PackageNameHyphen string
	ContractType      string
	Version           string
	ABI               string
	Bytecode          string
	NeedsBigInt       bool
	HasWriteOps       bool
	NoDeployment      bool
	Constructor       *ConstructorData
	StructDefs        []StructDefData
	ArgStructs        []ArgStructData
	Operations        []OperationData
	ContractMethods   []ContractMethodData
}

type ConstructorData struct {
	Parameters []ParameterData
}

type StructDefData struct {
	Name   string
	Fields []ParameterData
}

type ArgStructData struct {
	Name   string
	Fields []ParameterData
}

type ParameterData struct {
	GoName string
	GoType string
}

type OperationData struct {
	Name          string
	MethodName    string
	OpName        string
	ArgsType      string
	CallArgs      string
	IsWrite       bool
	AccessControl string // Only for writes
	ReturnType    string // Only for reads
}

type WriteOpData struct {
	Name          string
	MethodName    string
	OpName        string
	ArgsType      string
	CallArgs      string
	AccessControl string
}

type ReadOpData struct {
	Name       string
	MethodName string
	OpName     string
	ArgsType   string
	ReturnType string
	CallArgs   string
}

type ContractMethodData struct {
	Name       string
	MethodName string
	Params     string
	Returns    string
	MethodBody string
}

func main() {
	configPath := flag.String("config", "operations_gen_config.yaml", "Path to config file")
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

	// Get the directory containing the config file
	configDir := filepath.Dir(*configPath)
	absConfigDir, err := filepath.Abs(configDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving config directory: %v\n", err)
		os.Exit(1)
	}

	// Make input and output paths relative to config file location
	config.Input.BasePath = filepath.Join(absConfigDir, config.Input.BasePath)
	config.Output.BasePath = filepath.Join(absConfigDir, config.Output.BasePath)

	for _, contractCfg := range config.Contracts {
		info, err := extractContractInfo(contractCfg, config.Input, config.Output)
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

func extractContractInfo(cfg ContractConfig, input InputConfig, output OutputConfig) (*ContractInfo, error) {
	if cfg.Name == "" || cfg.Version == "" {
		return nil, fmt.Errorf("contract_name and version are required")
	}

	// Use explicit package name if provided, otherwise derive from contract name
	packageName := cfg.PackageName
	if packageName == "" {
		packageName = toSnakeCase(cfg.Name)
	}
	versionPath := versionToPath(cfg.Version)

	abiString, bytecode, err := readABIAndBytecode(cfg, packageName, versionPath, input.BasePath)
	if err != nil {
		return nil, err
	}

	var abiEntries []ABIEntry
	if err := json.Unmarshal([]byte(abiString), &abiEntries); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	info := &ContractInfo{
		Name:         cfg.Name,
		Version:      cfg.Version,
		PackageName:  packageName,
		OutputPath:   filepath.Join(output.BasePath, versionPath, "operations", packageName, packageName+".go"),
		ABI:          abiString,
		Bytecode:     bytecode,
		NoDeployment: cfg.NoDeployment,
		Functions:    make(map[string]*FunctionInfo),
		StructDefs:   make(map[string]*StructDef),
	}

	extractConstructor(info, abiEntries)

	if err := extractFunctions(info, cfg.Functions, abiEntries); err != nil {
		return nil, err
	}

	collectAllStructDefs(info)
	return info, nil
}

func readABIAndBytecode(
	cfg ContractConfig,
	packageName,
	versionPath,
	basePath string) (abiString string, bytecode string, err error) {
	// Determine ABI filename:
	// 1. Use explicit abi_file if provided
	// 2. Use package_name if provided (useful for usdc_token_pool_cctp_v2)
	// 3. Otherwise derive from contract_name
	var abiFileName string
	if cfg.ABIFile != "" {
		abiFileName = cfg.ABIFile
	} else {
		// Use the package name for ABI lookup - it matches the actual file names better
		abiFileName = packageName + ".json"
	}

	abiPath := filepath.Join(basePath, "abi", versionPath, abiFileName)
	abiBytes, err := os.ReadFile(abiPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read ABI from %s: %w", abiPath, err)
	}

	// Skip bytecode reading if no_deployment is true
	if cfg.NoDeployment {
		return string(abiBytes), "", nil
	}

	// Bytecode file matches ABI file name (without .json, with .bin)
	bytecodeFileName := strings.TrimSuffix(abiFileName, ".json") + ".bin"
	bytecodePath := filepath.Join(basePath, "bytecode", versionPath, bytecodeFileName)
	bytecodeBytes, err := os.ReadFile(bytecodePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read bytecode from %s: %w", bytecodePath, err)
	}

	return string(abiBytes), strings.TrimSpace(string(bytecodeBytes)), nil
}

func extractConstructor(info *ContractInfo, abiEntries []ABIEntry) {
	for _, entry := range abiEntries {
		if entry.Type == "constructor" {
			info.Constructor = parseABIFunction(entry, true, info.PackageName)
			break
		}
	}
}

func extractFunctions(info *ContractInfo, funcConfigs []FunctionConfig, abiEntries []ABIEntry) error {
	for _, funcCfg := range funcConfigs {
		funcInfos := findFunctionInABI(abiEntries, funcCfg.Name, info.PackageName)
		if funcInfos == nil {
			return fmt.Errorf("function %s not found in ABI", funcCfg.Name)
		}

		for _, funcInfo := range funcInfos {
			switch funcCfg.Access {
			case "owner":
				funcInfo.HasOnlyOwner = true
			case "public":
				funcInfo.HasOnlyOwner = false
			default:
				return fmt.Errorf("unknown access control '%s' for function %s (use 'owner' or 'public')",
					funcCfg.Access, funcCfg.Name)
			}

			// Use the potentially suffixed name as the key
			info.Functions[funcInfo.Name] = funcInfo
			info.FunctionOrder = append(info.FunctionOrder, funcInfo.Name)
		}
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

		if !funcInfo.IsWrite && len(funcInfo.ReturnParams) > 1 {
			structName := multiReturnStructName(funcInfo.Name)
			if _, exists := info.StructDefs[structName]; !exists {
				info.StructDefs[structName] = &StructDef{
					Name:   structName,
					Fields: funcInfo.ReturnParams,
				}
			}
		}
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

func findFunctionInABI(entries []ABIEntry, funcName string, packageName string) []*FunctionInfo {
	var candidates []ABIEntry
	for _, entry := range entries {
		if entry.Type == "function" && strings.EqualFold(entry.Name, funcName) {
			candidates = append(candidates, entry)
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	// Parse all overloads
	var funcInfos []*FunctionInfo
	for i, candidate := range candidates {
		funcInfo := parseABIFunction(candidate, false, packageName)

		// For overloaded functions, follow Geth's naming convention:
		// - First overload (i=0): no suffix (e.g., "Curse", "curse")
		// - Second overload (i=1): suffix "0" (e.g., "Curse0", "curse0")
		// - Third overload (i=2): suffix "1" (e.g., "Curse1", "curse1")
		if len(candidates) > 1 && i > 0 {
			suffix := fmt.Sprintf("%d", i-1)
			funcInfo.Name = funcInfo.Name + suffix
			funcInfo.CallMethod = funcInfo.CallMethod + suffix
		}

		funcInfos = append(funcInfos, funcInfo)
	}

	return funcInfos
}

func parseABIFunction(entry ABIEntry, _ bool, packageName string) *FunctionInfo {
	funcInfo := &FunctionInfo{
		Name:            capitalize(entry.Name),
		StateMutability: entry.StateMutability,
		CallMethod:      entry.Name,
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

//nolint:unparam
func parseABIParam(param ABIParam, packageName string) ParameterInfo {
	goType := solidityToGoType(param.Type, param.InternalType)

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

func solidityToGoType(solidityType, _ string) string {
	baseType := strings.TrimSuffix(solidityType, "[]")
	if goType, ok := typeMap[baseType]; ok {
		if strings.HasSuffix(solidityType, "[]") {
			return "[]" + goType
		}
		return goType
	}

	if strings.HasPrefix(baseType, "tuple") {
		return anyType
	}

	return anyType
}

func extractStructName(internalType string) string {
	if internalType == "" {
		return ""
	}

	parts := strings.Split(internalType, ".")
	structName := parts[len(parts)-1]
	structName = strings.TrimSuffix(structName, "[]")

	return structName
}

func versionToPath(version string) string {
	return "v" + strings.ReplaceAll(version, ".", "_")
}

func toSnakeCase(s string) string {
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
	data := prepareTemplateData(info)

	t, err := template.New("operations").Parse(operationsTemplate)
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

	if err := os.WriteFile(info.OutputPath, formatted, 0600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func prepareTemplateData(info *ContractInfo) TemplateData {
	data := TemplateData{
		PackageName:       info.PackageName,
		PackageNameHyphen: toKebabCase(info.PackageName),
		ContractType:      info.Name,
		Version:           info.Version,
		ABI:               info.ABI,
		Bytecode:          info.Bytecode,
		NeedsBigInt:       checkNeedsBigInt(info),
		NoDeployment:      info.NoDeployment,
	}

	if info.Constructor != nil {
		data.Constructor = &ConstructorData{
			Parameters: prepareParameters(info.Constructor.Parameters),
		}
	}

	for _, name := range info.FunctionOrder {
		funcInfo := info.Functions[name]
		data.ContractMethods = append(data.ContractMethods, prepareContractMethod(funcInfo, funcInfo.IsWrite))

		// Add to unified operations list
		if funcInfo.IsWrite {
			data.HasWriteOps = true
			writeOp := prepareWriteOp(funcInfo)
			data.Operations = append(data.Operations, OperationData{
				Name:          writeOp.Name,
				MethodName:    writeOp.MethodName,
				OpName:        writeOp.OpName,
				ArgsType:      writeOp.ArgsType,
				CallArgs:      writeOp.CallArgs,
				IsWrite:       true,
				AccessControl: writeOp.AccessControl,
			})
			if len(funcInfo.Parameters) > 1 {
				data.ArgStructs = append(data.ArgStructs, ArgStructData{
					Name:   funcInfo.Name + "Args",
					Fields: prepareParameters(funcInfo.Parameters),
				})
			}
		} else {
			readOp := prepareReadOp(funcInfo)
			data.Operations = append(data.Operations, OperationData{
				Name:       readOp.Name,
				MethodName: readOp.MethodName,
				OpName:     readOp.OpName,
				ArgsType:   readOp.ArgsType,
				CallArgs:   readOp.CallArgs,
				IsWrite:    false,
				ReturnType: readOp.ReturnType,
			})
			// Generate Args struct for read operations with multiple parameters
			if len(funcInfo.Parameters) > 1 {
				data.ArgStructs = append(data.ArgStructs, ArgStructData{
					Name:   funcInfo.Name + "Args",
					Fields: prepareParameters(funcInfo.Parameters),
				})
			}
		}
	}

	var structNames []string
	for name := range info.StructDefs {
		structNames = append(structNames, name)
	}
	sort.Strings(structNames)
	for _, name := range structNames {
		structDef := info.StructDefs[name]
		data.StructDefs = append(data.StructDefs, StructDefData{
			Name:   structDef.Name,
			Fields: prepareParameters(structDef.Fields),
		})
	}

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

func prepareContractMethod(funcInfo *FunctionInfo, isWrite bool) ContractMethodData {
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
			paramName := p.Name
			if len(paramName) > 0 {
				paramName = strings.ToLower(paramName[:1]) + paramName[1:]
			}
			if paramName == "" {
				paramName = fmt.Sprintf("arg%d", len(methodArgs))
			}
			params += fmt.Sprintf(", %s %s", paramName, p.GoType)
			methodArgs = append(methodArgs, paramName)
		}
	}

	returns := "(*types.Transaction, error)"
	returnType := anyType
	if !isWrite {
		if len(funcInfo.ReturnParams) == 1 {
			returnType = funcInfo.ReturnParams[0].GoType
		} else if len(funcInfo.ReturnParams) > 1 {
			returnType = multiReturnStructName(funcInfo.Name)
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
		if len(funcInfo.ReturnParams) > 1 {
			methodBody = buildMultiReturnMethodBody(funcInfo, callArgsStr, returnType)
		} else {
			methodBody = fmt.Sprintf(
				`var out []any
	err := c.contract.Call(opts, &out, "%s"%s)
	if err != nil {
		var zero %s
		return zero, err
	}
	return *abi.ConvertType(out[0], new(%s)).(*%s), nil`,
				funcInfo.CallMethod, callArgsStr, returnType, returnType, returnType,
			)
		}
	}

	return ContractMethodData{
		Name:       funcInfo.Name,
		MethodName: funcInfo.CallMethod,
		Params:     params,
		Returns:    returns,
		MethodBody: methodBody,
	}
}

func prepareParameters(params []ParameterInfo) []ParameterData {
	var result []ParameterData
	for i, param := range params {
		name := capitalize(param.Name)
		if name == "" {
			name = fmt.Sprintf("Field%d", i)
		}
		result = append(result, ParameterData{
			GoName: name,
			GoType: param.GoType,
		})
	}
	return result
}

// buildCallArgs extracts common logic for building argument types and call arguments
func buildCallArgs(funcInfo *FunctionInfo, argsPrefix string) (argsType string, callArgs string) {
	if len(funcInfo.Parameters) == 0 {
		return "struct{}", ""
	}

	if len(funcInfo.Parameters) == 1 {
		return funcInfo.Parameters[0].GoType, ", " + argsPrefix
	}

	// Multiple parameters - use Args struct
	argsType = funcInfo.Name + "Args"
	var callArgsList []string
	for _, p := range funcInfo.Parameters {
		callArgsList = append(callArgsList, argsPrefix+"."+capitalize(p.Name))
	}
	callArgs = ", " + strings.Join(callArgsList, ", ")
	return argsType, callArgs
}

func prepareWriteOp(funcInfo *FunctionInfo) WriteOpData {
	argsType, callArgs := buildCallArgs(funcInfo, "args")

	accessControl := "AllCallersAllowed"
	if funcInfo.HasOnlyOwner {
		accessControl = "OnlyOwner"
	}

	return WriteOpData{
		Name:          funcInfo.Name,
		MethodName:    funcInfo.CallMethod,
		OpName:        toKebabCase(funcInfo.Name),
		ArgsType:      argsType,
		CallArgs:      callArgs,
		AccessControl: accessControl,
	}
}

func multiReturnStructName(funcName string) string {
	return funcName + "Result"
}

func buildMultiReturnMethodBody(funcInfo *FunctionInfo, callArgsStr, returnType string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "var out []any\n")
	fmt.Fprintf(&b, "\terr := c.contract.Call(opts, &out, \"%s\"%s)\n", funcInfo.CallMethod, callArgsStr)
	fmt.Fprintf(&b, "\toutstruct := new(%s)\n", returnType)
	fmt.Fprintf(&b, "\tif err != nil {\n")
	fmt.Fprintf(&b, "\t\treturn *outstruct, err\n")
	fmt.Fprintf(&b, "\t}\n\n")
	for i, p := range funcInfo.ReturnParams {
		fieldName := capitalize(p.Name)
		if fieldName == "" {
			fieldName = fmt.Sprintf("Field%d", i)
		}
		fmt.Fprintf(&b, "\toutstruct.%s = *abi.ConvertType(out[%d], new(%s)).(*%s)\n",
			fieldName, i, p.GoType, p.GoType)
	}
	fmt.Fprintf(&b, "\n\treturn *outstruct, nil")
	return b.String()
}

func prepareReadOp(funcInfo *FunctionInfo) ReadOpData {
	argsType, callArgs := buildCallArgs(funcInfo, "args")

	returnType := anyType
	if len(funcInfo.ReturnParams) == 1 {
		returnType = funcInfo.ReturnParams[0].GoType
	} else if len(funcInfo.ReturnParams) > 1 {
		returnType = multiReturnStructName(funcInfo.Name)
	}

	return ReadOpData{
		Name:       funcInfo.Name,
		MethodName: funcInfo.CallMethod,
		OpName:     toKebabCase(funcInfo.Name),
		ArgsType:   argsType,
		ReturnType: returnType,
		CallArgs:   callArgs,
	}
}
