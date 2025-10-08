package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Determine the correct paths based on the current working directory
	gobindingsDir, bytecodeDir, err := findProjectDirs()
	if err != nil {
		return fmt.Errorf("failed to find project directories: %w", err)
	}

	fmt.Printf("Using gobindings dir: %s\n", gobindingsDir)
	fmt.Printf("Using bytecode dir: %s\n", bytecodeDir)

	// Get all versioned directories (exclude "latest" because it's unaudited)
	entries, err := os.ReadDir(gobindingsDir)
	if err != nil {
		return fmt.Errorf("failed to read gobindings directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "latest" {
			continue
		}

		versionDir := filepath.Join(gobindingsDir, entry.Name())
		fmt.Printf("Processing version: %s\n", entry.Name())

		if err := processVersionDir(versionDir, gobindingsDir, bytecodeDir); err != nil {
			return fmt.Errorf("failed to process version %s: %w", entry.Name(), err)
		}
	}

	return nil
}

func findProjectDirs() (gobindingsDir, bytecodeDir string, err error) {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Try to find the gobindings directory by walking up from current directory
	dir := cwd
	for {
		// Check if "generated" subdirectory exists (this indicates we're in gobindings)
		generatedPath := filepath.Join(dir, "generated")
		if info, err := os.Stat(generatedPath); err == nil && info.IsDir() {
			// We found the gobindings directory
			gobindingsDir = generatedPath
			bytecodeDir = filepath.Join(filepath.Dir(dir), "bytecode")
			return gobindingsDir, bytecodeDir, nil
		}

		// Check if "gobindings/generated" exists (we might be in chains/evm)
		gobindingsPath := filepath.Join(dir, "gobindings", "generated")
		if info, err := os.Stat(gobindingsPath); err == nil && info.IsDir() {
			gobindingsDir = gobindingsPath
			bytecodeDir = filepath.Join(dir, "bytecode")
			return gobindingsDir, bytecodeDir, nil
		}

		// Check if "chains/evm/gobindings/generated" exists (we might be in project root)
		chainsEvmPath := filepath.Join(dir, "chains", "evm", "gobindings", "generated")
		if info, err := os.Stat(chainsEvmPath); err == nil && info.IsDir() {
			gobindingsDir = chainsEvmPath
			bytecodeDir = filepath.Join(dir, "chains", "evm", "bytecode")
			return gobindingsDir, bytecodeDir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root
			break
		}
		dir = parent
	}

	return "", "", fmt.Errorf("could not find gobindings directory. Please run from within the chainlink-ccip project")
}

func processVersionDir(versionDir, gobindingsDir, bytecodeDir string) error {
	return filepath.Walk(versionDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .go files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Extract bytecode from the file
		bytecode, err := extractBytecode(path)
		if err != nil {
			fmt.Printf("  Warning: failed to extract bytecode from %s: %v\n", path, err)
			return nil // Continue processing other files
		}

		if bytecode == "" {
			// No bytecode found in this file, skip
			return nil
		}

		// Determine the output path
		relPath, err := filepath.Rel(gobindingsDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Replace .go with .bin
		relPath = strings.TrimSuffix(relPath, ".go") + ".bin"
		outputPath := filepath.Join(bytecodeDir, relPath)

		// Create the output directory
		outputDir := filepath.Dir(outputPath)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", outputDir, err)
		}

		// Write the bytecode to the file
		if err := os.WriteFile(outputPath, []byte(bytecode), 0600); err != nil {
			return fmt.Errorf("failed to write bytecode to %s: %w", outputPath, err)
		}

		fmt.Printf("  ✓ Extracted bytecode: %s\n", relPath)
		return nil
	})
}

func extractBytecode(filePath string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("failed to parse file: %w", err)
	}

	var bytecode string

	// Walk through the AST to find the MetaData variable with Bin field
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for variable declarations
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			return true
		}

		if bin := extractBinFromVarDecl(genDecl); bin != "" {
			bytecode = bin
			return false // Stop walking, we found what we need
		}

		return true
	})

	return bytecode, nil
}

// extractBinFromVarDecl extracts the Bin field from a variable declaration if it's a MetaData variable
func extractBinFromVarDecl(genDecl *ast.GenDecl) string {
	for _, spec := range genDecl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok || len(valueSpec.Values) == 0 {
			continue
		}

		// Check if this is a MetaData variable
		if !isMetaDataVar(valueSpec) {
			continue
		}

		// Extract the Bin field from the composite literal
		if bin := extractBinFromValue(valueSpec.Values[0]); bin != "" {
			return bin
		}
	}
	return ""
}

// isMetaDataVar checks if a value spec is a MetaData variable
func isMetaDataVar(valueSpec *ast.ValueSpec) bool {
	for _, name := range valueSpec.Names {
		if strings.HasSuffix(name.Name, "MetaData") {
			return true
		}
	}
	return false
}

// extractBinFromValue extracts the Bin field from a value expression
func extractBinFromValue(value ast.Expr) string {
	// Look for composite literal (&bind.MetaData{...})
	unaryExpr, ok := value.(*ast.UnaryExpr)
	if !ok {
		return ""
	}

	compositeLit, ok := unaryExpr.X.(*ast.CompositeLit)
	if !ok {
		return ""
	}

	return extractBinFromCompositeLit(compositeLit)
}

// extractBinFromCompositeLit extracts the Bin field from a composite literal
func extractBinFromCompositeLit(compositeLit *ast.CompositeLit) string {
	for _, elt := range compositeLit.Elts {
		kvExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		key, ok := kvExpr.Key.(*ast.Ident)
		if !ok || key.Name != "Bin" {
			continue
		}

		// Extract the string literal value
		if basicLit, ok := kvExpr.Value.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
			// Remove the surrounding quotes
			return strings.Trim(basicLit.Value, `"`)
		}
	}
	return ""
}
