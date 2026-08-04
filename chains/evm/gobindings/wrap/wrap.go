package main

import (
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
)

func main() {
	project := os.Args[1]
	contract := os.Args[2]
	pkgName := os.Args[3]

	var outDirSuffix string
	if len(os.Args) >= 5 {
		outDirSuffix = os.Args[4] + "/latest"
	} else {
		outDirSuffix = "../generated/latest"
	}

	projectRoot := "../solc/" + project
	abiGenPath := "../scripts/abigen"

	wrap.GenWrapper(projectRoot, contract, pkgName, outDirSuffix, abiGenPath)
}
