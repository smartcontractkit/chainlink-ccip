package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	// hash of chainlink-ccip commit to verify
	COMMIT_HASH = "f93a56f0edc533b3a50c07f55182753b4c4b5b69"

	// RPC URL based on environment (cluster)
	SOLANA_RPC_URL = "https://api.devnet.solana.com"
	CLUSTER        = "devnet"

	// funded keypair
	KEYPAIR_PATH = "$HOME/.config/solana/id_devnet.json"
)

const (
	// URL of the repository to verify (FIXED)
	REPO_URL          = "https://github.com/smartcontractkit/chainlink-ccip"
	GITHUB_REPO       = "smartcontractkit/chainlink-ccip"
	GITHUB_BRANCH_DIR = "chains/solana/contracts/Anchor.toml"
	// path to the directory containing the contracts (FIXED)
	MOUNT_PATH = "chains/solana/contracts"
)

// Fetches Anchor.toml from the github repo
func fetchToml() []byte {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", GITHUB_REPO, COMMIT_HASH, GITHUB_BRANCH_DIR)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch Anchor.toml: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	return body
}

// Downloads target/idl/<program>.json into ./idl/
func fetchIDLs(programs map[string]string) error {
	baseURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/chains/solana/contracts/target/idl", GITHUB_REPO, COMMIT_HASH)

	if err := os.MkdirAll("idl", 0o755); err != nil {
		return fmt.Errorf("creating idl dir: %w", err)
	}

	for program := range programs {
		url := fmt.Sprintf("%s/%s.json", baseURL, program)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("❌ Failed to download IDL for %s: %v", program, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("⚠️ Skipping %s: IDL not found (status %d)", program, resp.StatusCode)
			continue
		}

		outPath := filepath.Join("idl", fmt.Sprintf("%s.json", program))
		outFile, err := os.Create(outPath)
		if err != nil {
			log.Printf("❌ Could not write %s: %v", outPath, err)
			continue
		}

		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			log.Printf("❌ Failed writing IDL for %s: %v", program, err)
		} else {
			fmt.Printf("✅ Downloaded IDL for %s → %s\n", program, outPath)
		}
		outFile.Close()
	}

	return nil
}

// Updates each IDL file with the program address
func patchIDLsWithAddresses(programs map[string]string) error {
	for program, address := range programs {
		idlPath := filepath.Join("idl", program+".json")

		data, err := os.ReadFile(idlPath)
		if err != nil {
			log.Printf("❌ Could not read %s: %v", idlPath, err)
			continue
		}

		var idl map[string]interface{}
		if err := json.Unmarshal(data, &idl); err != nil {
			log.Printf("❌ Invalid JSON in %s: %v", idlPath, err)
			continue
		}

		// Inject metadata.address
		idl["metadata"] = map[string]interface{}{
			"address": address,
		}

		// Marshal and write back to file
		out, err := json.MarshalIndent(idl, "", "  ")
		if err != nil {
			log.Printf("❌ Failed to re-encode JSON for %s: %v", idlPath, err)
			continue
		}

		if err := os.WriteFile(idlPath, out, 0644); err != nil {
			log.Printf("❌ Could not write updated IDL to %s: %v", idlPath, err)
			continue
		}

		fmt.Printf("✅ Patched %s with address %s\n", program, address)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 || (os.Args[1] != "idl" && os.Args[1] != "verify") {
		fmt.Println("Please supply either 'idl' or 'verify' as an argument")
		os.Exit(1)
	}

	mode := os.Args[1]

	anchorData := struct {
		Programs struct {
			Localnet map[string]string
		}
	}{}
	anchorBytes := fetchToml()
	err := toml.Unmarshal(anchorBytes, &anchorData)
	if err != nil {
		log.Fatal("Failed to unmarshal anchor toml")
	}

	if mode == "verify" {
		// print the verify commands
		for libName, address := range anchorData.Programs.Localnet {
			fmt.Printf("🔍 Verifying %s at %s \n", libName, address)

			verifyCmd := []string{
				"solana-verify", "verify-from-repo", REPO_URL,
				"--commit-hash", COMMIT_HASH,
				"--url", SOLANA_RPC_URL,
				"--program-id", address,
				"--mount-path", MOUNT_PATH,
				"--library-name", libName,
				"--keypair", KEYPAIR_PATH,
				"--skip-prompt", "--remote",
			}

			fmt.Println("[DRY RUN]", strings.Join(verifyCmd, " "))
			fmt.Println()
		}
		return
	}

	// fetch the idls
	err = fetchIDLs(anchorData.Programs.Localnet)
	if err != nil {
		log.Fatalf("❌ Failed to fetch IDLs: %v", err)
	}

	// update the idl files for idl init
	err = patchIDLsWithAddresses(anchorData.Programs.Localnet)
	if err != nil {
		log.Fatalf("❌ Failed to patch IDLs: %v", err)
	}

	// print the idl init commands
	for program, address := range anchorData.Programs.Localnet {
		fmt.Printf("🔍 Initializing IDL for %s\n", program)

		idlInitCmd := []string{
			"anchor", "idl", "init", "-f",
			filepath.Join("target", "idl", program+".json"),
			"--provider.cluster", CLUSTER,
			"--provider.wallet", KEYPAIR_PATH,
			address,
		}

		fmt.Println("[DRY RUN]", strings.Join(idlInitCmd, " "))
		fmt.Println()
	}
}
