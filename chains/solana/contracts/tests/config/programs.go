package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gagliardetto/solana-go"
)

// Structure to match Anchor.toml format
type Config struct {
	Programs map[string]map[string]string `toml:"programs"`
}

var AnchorToml = parseAnchorToml()

func parseAnchorToml() Config {
	var config Config
	startDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return config
	}

	dir := startDir
	var tomlPath string
	for {
		tomlPath = filepath.Join(dir, "Anchor.toml")
		if _, err := os.Stat(tomlPath); err == nil {
			break
		}

		// Move up one level
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			fmt.Println("Anchor.toml not found")
			return config
		}
		dir = parentDir
	}

	if _, err := toml.DecodeFile(tomlPath, &config); err != nil {
		fmt.Println("Error reading Anchor.toml:", err)
		return config
	}

	return config
}

// GetProgramID retrieves the program ID dynamically
func GetProgramID(programName string) solana.PublicKey {
	programID, exists := AnchorToml.Programs["localnet"][programName]
	if !exists {
		fmt.Printf("Program ID %s not found in Anchor.toml", programName)
		return solana.PublicKey{}
	}

	return solana.MustPublicKeyFromBase58(programID)
}
