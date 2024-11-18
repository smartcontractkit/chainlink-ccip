package testing

import (
	"io/fs"
	"log"
	"os"
)

// MockKubeConfigFile is a helper function that returns the path to a tempfile
// containing the desired content and permissions
func MockKubeConfigFile(content []byte, perm fs.FileMode) *os.File {
	tempFile, err := os.CreateTemp("", "config")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}

	if err := os.WriteFile(tempFile.Name(), content, perm); err != nil {
		log.Fatal(err)
	}

	return tempFile
}
