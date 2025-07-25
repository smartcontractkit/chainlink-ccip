package testutils

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// ProjectRoot Root folder of this project
	ProjectRoot = filepath.Join(filepath.Dir(b), "/../..")
	// ContractsDir path to our contracts
	ContractsDir = filepath.Join(ProjectRoot, "target", "deploy")
	// VendorContractsDir path to vendored contract binaries
	VendorContractsDir = filepath.Join(ProjectRoot, "target", "vendor")
)
