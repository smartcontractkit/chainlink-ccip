package main

import (
	"fmt"
	"os"
	"path/filepath"

	gethwrappers "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generation"
)

var (
	rootDir = "../solc/"
)

func main() {
	project := os.Args[1]
	className := os.Args[2]
	pkgName := os.Args[3]

	var outDirSuffix string
	if len(os.Args) >= 5 {
		outDirSuffix = os.Args[4]
	}

	abiPath := rootDir + project + "/" + className + "/" + className + ".sol/" + className + ".abi.json"
	binPath := rootDir + project + "/" + className + "/" + className + ".sol/" + className + ".bin"
	metadataPath := rootDir + project + "/" + className + "/build/build.json"

	GenWrapper(abiPath, binPath, metadataPath, className, pkgName, outDirSuffix)
}

// GenWrapper generates a contract wrapper for the given contract.
//
// abiPath is the path to the contract's ABI JSON file.
//
// binPath is the path to the contract's binary file, typically with .bin extension.
//
// className is the name of the generated contract class.
//
// pkgName is the name of the package the contract will be generated in. Try
// to follow idiomatic Go package naming conventions where possible.
//
// outDirSuffixInput is the directory suffix to generate the wrapper in. If not provided, the
// wrapper will be generated in the default location. The default location is
// <project>/generated/<pkgName>/<pkgName>.go. The suffix will take place after
// the <project>/generated, so the overridden location would be
// <project>/generated/<outDirSuffixInput>/<pkgName>/<pkgName>.go.
func GenWrapper(abiPath, binPath, metadataPath, className, pkgName, outDirSuffixInput string) {
	fmt.Println("Generating", pkgName, "contract wrapper")

	cwd, err := os.Getwd() // gethwrappers directory
	if err != nil {
		gethwrappers.Exit("could not get working directory", err)
	}
	outDir := filepath.Join(cwd, "generated", outDirSuffixInput, pkgName)
	if mkdErr := os.MkdirAll(outDir, 0700); err != nil {
		gethwrappers.Exit(
			fmt.Sprintf("failed to create wrapper dir, outDirSuffixInput: %s (could be empty)", outDirSuffixInput),
			mkdErr)
	}
	outPath := filepath.Join(outDir, pkgName+".go")
	metadataOutPath := filepath.Join(outDir, pkgName+"_metadata.go")

	gethwrappers.Abigen(gethwrappers.AbigenArgs{
		Bin:         binPath,
		ABI:         abiPath,
		Metadata:    metadataPath,
		Out:         outPath,
		MetadataOut: metadataOutPath,
		Type:        className,
		Pkg:         pkgName,
	})

	// Build succeeded, so update the versions db with the new contract data
	versions, err := gethwrappers.ReadVersionsDB()
	if err != nil {
		gethwrappers.Exit("could not read current versions database", err)
	}
	versions.GethVersion = gethwrappers.GethVersion
	versions.ContractVersions[pkgName] = gethwrappers.ContractVersion{
		Hash:       gethwrappers.VersionHash(abiPath, binPath),
		AbiPath:    abiPath,
		BinaryPath: binPath,
	}
	if err := gethwrappers.WriteVersionsDB(versions); err != nil {
		gethwrappers.Exit("could not save versions db", err)
	}
}
