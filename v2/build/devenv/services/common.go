package services

import (
	"os"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go"
)

const (
	AppPathInsideContainer = "/app"
)

// CwdSourcePath returns source path for current working directory
func CwdSourcePath(sourcePath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(wd), sourcePath), nil
}

// GoSourcePathMounts returns default Golang cache/build-cache and dev-image mounts
func GoSourcePathMounts(sourcePath string, containerDirTarget string) testcontainers.ContainerMounts {
	return testcontainers.Mounts(
		testcontainers.BindMount(
			sourcePath,
			testcontainers.ContainerMountTarget(containerDirTarget),
		),
		testcontainers.VolumeMount(
			"go-mod-cache",
			"/go/pkg/mod",
		),
		testcontainers.VolumeMount(
			"go-build-cache",
			"/root/.cache/go-build",
		),
	)
}
