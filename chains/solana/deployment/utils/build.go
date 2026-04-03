package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

const (
	repoURL   = "https://github.com/smartcontractkit/chainlink-ccip.git"
	anchorDir = "chains/solana/contracts"
	deployDir = "chains/solana/contracts/target/deploy"
)

// ProgramToRustFile maps contract types to their Rust source paths (relative to the anchor project root)
// used for key replacement during upgrade builds.
var ProgramToRustFile = map[cldf.ContractType]string{
	"Router":                  "programs/ccip-router/src/lib.rs",
	"CCIPCommon":              "programs/ccip-common/src/lib.rs",
	"FeeQuoter":               "programs/fee-quoter/src/lib.rs",
	"OffRamp":                 "programs/ccip-offramp/src/lib.rs",
	"BurnMintTokenPool":       "programs/burnmint-token-pool/src/lib.rs",
	"LockReleaseTokenPool":    "programs/lockrelease-token-pool/src/lib.rs",
	"RMNRemote":               "programs/rmn-remote/src/lib.rs",
	"AccessControllerProgram": "programs/access-controller/src/lib.rs",
	"ManyChainMultiSigProgram": "programs/mcm/src/lib.rs",
	"RBACTimelockProgram":     "programs/timelock/src/lib.rs",
	"CCTPTokenPool":           "programs/cctp-token-pool/src/lib.rs",
}

// ProgramToVanityPrefix maps contract types to the vanity key prefix used with `solana-keygen grind`.
var ProgramToVanityPrefix = map[cldf.ContractType]string{
	"Router":    "Ccip",
	"FeeQuoter": "FeeQ",
	"OffRamp":   "off",
	"RMNRemote": "Rmn",
}

// SolanaBuildConfig configures how Solana program artifacts are prepared.
type SolanaBuildConfig struct {
	// ContractVersion is the version tag (e.g. "solana-v0.1.2") that maps to a commit via VersionToFullCommitSHA.
	ContractVersion string
	// DestinationDir is where built/downloaded .so files and keypairs are placed.
	// Typically chain.ProgramsPath.
	DestinationDir string
	// WorkDir is the temporary working directory for git clone and build operations.
	// Each build must use a unique directory to avoid interference when running in
	// parallel. If empty, a new temp directory is created automatically.
	WorkDir string
	// LocalBuild controls the local build pipeline. If nil or BuildLocally is false,
	// artifacts are downloaded from the GitHub release instead.
	LocalBuild *LocalBuildConfig
}

type LocalBuildConfig struct {
	BuildLocally         bool
	CleanDestinationDir  bool
	CreateDestinationDir bool
	// CleanGitDir forces re-clone of the git directory. Useful for forcing key regeneration.
	CleanGitDir bool
	// GenerateVanityKeys generates vanity keypairs for programs that have configured prefixes.
	GenerateVanityKeys bool
	// UpgradeKeys maps contract types to their existing on-chain program addresses.
	// Used during upgrade builds to replace declare_id!() in Rust source so the
	// rebuilt binary's embedded key matches the deployed program.
	UpgradeKeys map[cldf.ContractType]string
}

// BuildSolana either downloads pre-built artifacts or builds them locally.
func BuildSolana(ctx context.Context, lggr logger.Logger, config SolanaBuildConfig) error {
	if config.LocalBuild == nil || !config.LocalBuild.BuildLocally {
		lggr.Info("Downloading Solana CCIP program artifacts...")
		sha, ok := VersionToShortCommitSHA[config.ContractVersion]
		if !ok {
			return fmt.Errorf("solana contract version not found: %s", config.ContractVersion)
		}
		return DownloadSolanaCCIPProgramArtifacts(ctx, config.DestinationDir, sha)
	}

	if config.WorkDir == "" {
		dir, err := os.MkdirTemp("", "solana-build-*")
		if err != nil {
			return fmt.Errorf("failed to create temp work directory: %w", err)
		}
		config.WorkDir = dir
		defer os.RemoveAll(dir)
	}

	lggr.Infow("Building Solana CCIP program artifacts locally", "workDir", config.WorkDir)
	return buildLocally(lggr, config)
}

func buildLocally(lggr logger.Logger, config SolanaBuildConfig) error {
	commitSHA, ok := VersionToFullCommitSHA[config.ContractVersion]
	if !ok {
		return fmt.Errorf("solana contract version not found: %s", config.ContractVersion)
	}

	repoDir := filepath.Join(config.WorkDir, "repo")

	if err := cloneRepo(lggr, repoDir, commitSHA, config.LocalBuild.CleanGitDir); err != nil {
		return fmt.Errorf("error cloning repo: %w", err)
	}

	if err := replaceKeys(lggr, repoDir); err != nil {
		return fmt.Errorf("error replacing keys: %w", err)
	}

	if config.LocalBuild.GenerateVanityKeys {
		if config.LocalBuild.UpgradeKeys == nil {
			config.LocalBuild.UpgradeKeys = make(map[cldf.ContractType]string)
		}
		if err := generateVanityKeys(lggr, config.WorkDir, repoDir, config.LocalBuild.UpgradeKeys); err != nil {
			return fmt.Errorf("error generating vanity keys: %w", err)
		}
	}

	if err := replaceKeysForUpgrade(lggr, repoDir, config.LocalBuild.UpgradeKeys); err != nil {
		return fmt.Errorf("error replacing keys for upgrade: %w", err)
	}

	if err := syncRouterAndCommon(repoDir); err != nil {
		return fmt.Errorf("error syncing router and common program files: %w", err)
	}

	if err := buildProject(lggr, repoDir); err != nil {
		return fmt.Errorf("error building project: %w", err)
	}

	if config.LocalBuild.CleanDestinationDir {
		lggr.Infow("Cleaning destination dir", "dir", config.DestinationDir)
		if err := os.RemoveAll(config.DestinationDir); err != nil {
			return fmt.Errorf("error cleaning build folder: %w", err)
		}
		if err := os.MkdirAll(config.DestinationDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create build directory: %w", err)
		}
	} else if config.LocalBuild.CreateDestinationDir {
		if err := os.MkdirAll(config.DestinationDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create build directory: %w", err)
		}
	}

	deployFilePath := filepath.Join(repoDir, deployDir)
	lggr.Infow("Reading deploy directory", "path", deployFilePath)
	files, err := os.ReadDir(deployFilePath)
	if err != nil {
		return fmt.Errorf("failed to read deploy directory: %w", err)
	}

	for _, file := range files {
		src := filepath.Join(deployFilePath, file.Name())
		dst := filepath.Join(config.DestinationDir, file.Name())
		lggr.Infow("Copying artifact", "src", src, "dst", dst)
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("failed to copy file %s: %w", file.Name(), err)
		}
	}
	return nil
}

func cloneRepo(lggr logger.Logger, repoDir string, revision string, forceClean bool) error {
	if forceClean {
		lggr.Infow("Cleaning repository", "dir", repoDir)
		if err := os.RemoveAll(repoDir); err != nil {
			return fmt.Errorf("failed to clean repository: %w", err)
		}
	}
	if _, err := os.Stat(filepath.Join(repoDir, ".git")); err == nil {
		lggr.Infow("Repository already exists, resetting and fetching", "dir", repoDir)
		if _, err := RunCommand("git", []string{"reset", "--hard"}, repoDir); err != nil {
			return fmt.Errorf("failed to discard local changes: %w", err)
		}
		if _, err := RunCommand("git", []string{"fetch", "origin"}, repoDir); err != nil {
			return fmt.Errorf("failed to fetch origin: %w", err)
		}
	} else {
		lggr.Infow("Cloning repository", "url", repoURL, "revision", revision)
		if _, err := RunCommand("git", []string{"clone", repoURL, repoDir}, "."); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}
	}

	lggr.Infow("Checking out revision", "revision", revision)
	if _, err := RunCommand("git", []string{"checkout", revision}, repoDir); err != nil {
		return fmt.Errorf("failed to checkout revision %s: %w", revision, err)
	}
	return nil
}

// replaceKeys runs `make docker-update-contracts` which calls `anchor keys sync`
// to update the declare_id!() in source files to match the generated keypairs.
func replaceKeys(lggr logger.Logger, repoDir string) error {
	solanaDir := filepath.Join(repoDir, anchorDir, "..")
	lggr.Infow("Replacing keys via anchor keys sync", "dir", solanaDir)
	output, err := RunCommand("make", []string{"docker-update-contracts"}, solanaDir)
	if err != nil {
		return fmt.Errorf("anchor key replacement failed: %s %w", output, err)
	}
	return nil
}

// replaceKeysForUpgrade explicitly replaces declare_id!() macros in Rust source files
// with the keys of already-deployed programs. This ensures the rebuilt binary matches
// the on-chain program address for in-place upgrades.
func replaceKeysForUpgrade(lggr logger.Logger, repoDir string, keys map[cldf.ContractType]string) error {
	if len(keys) == 0 {
		return nil
	}
	lggr.Info("Replacing keys in Rust files for upgrade...")
	declareIDRegex := regexp.MustCompile(`declare_id!\(".*?"\);`)
	for program, key := range keys {
		filePath, exists := ProgramToRustFile[program]
		if !exists {
			return fmt.Errorf("no file path found for program %s", program)
		}

		fullPath := filepath.Join(repoDir, anchorDir, filePath)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", fullPath, err)
		}

		updatedContent := declareIDRegex.ReplaceAllString(string(content), fmt.Sprintf(`declare_id!("%s");`, key))
		if err := os.WriteFile(fullPath, []byte(updatedContent), 0644); err != nil {
			return fmt.Errorf("failed to write updated keys to file %s: %w", fullPath, err)
		}
		lggr.Infow("Updated declare_id for upgrade", "program", program, "file", filePath)
	}
	return nil
}

// syncRouterAndCommon ensures the ccip-common lib.rs declare_id matches the router's,
// since ccip-common is a shared crate that must carry the router's program ID.
func syncRouterAndCommon(repoDir string) error {
	routerFile := filepath.Join(repoDir, anchorDir, ProgramToRustFile["Router"])
	commonFile := filepath.Join(repoDir, anchorDir, ProgramToRustFile["CCIPCommon"])

	file, err := os.Open(routerFile)
	if err != nil {
		return fmt.Errorf("error opening router file: %w", err)
	}
	defer file.Close()

	declareRegex := regexp.MustCompile(`declare_id!\("(.+?)"\);`)
	scanner := bufio.NewScanner(file)
	var declareID string
	for scanner.Scan() {
		if match := declareRegex.FindStringSubmatch(scanner.Text()); match != nil {
			declareID = match[0]
			break
		}
	}
	if declareID == "" {
		return fmt.Errorf("declare_id not found in router file")
	}

	commonContent, err := os.ReadFile(commonFile)
	if err != nil {
		return fmt.Errorf("error reading common file: %w", err)
	}
	updatedContent := declareRegex.ReplaceAllString(string(commonContent), declareID)
	return os.WriteFile(commonFile, []byte(updatedContent), 0644)
}

func generateVanityKeys(lggr logger.Logger, workDir string, repoDir string, keys map[cldf.ContractType]string) error {
	lggr.Info("Generating vanity keys...")
	jsonFilePattern := regexp.MustCompile(`Wrote keypair to (.*\.json)`)
	for program, prefix := range ProgramToVanityPrefix {
		if _, exists := keys[program]; exists {
			lggr.Infow("Vanity key already exists, skipping", "program", program)
			continue
		}

		output, err := RunCommand("solana-keygen", []string{"grind", "--starts-with", prefix + ":1"}, workDir)
		if err != nil {
			return fmt.Errorf("failed to generate vanity key for %s: %w", program, err)
		}

		scanner := bufio.NewScanner(strings.NewReader(output))
		var jsonFilePath string
		for scanner.Scan() {
			if matches := jsonFilePattern.FindStringSubmatch(scanner.Text()); len(matches) > 1 {
				jsonFilePath = matches[1]
				break
			}
		}
		if jsonFilePath == "" {
			return fmt.Errorf("failed to parse vanity key output for %s", program)
		}

		// keygen writes relative to workDir; resolve from there
		if !filepath.IsAbs(jsonFilePath) {
			jsonFilePath = filepath.Join(workDir, jsonFilePath)
		}

		fileName := filepath.Base(jsonFilePath)
		keys[program] = strings.TrimSuffix(fileName, ".json")

		destination := filepath.Join(repoDir, deployDir, programTypeToDeployName(program)+"-keypair.json")
		if err := os.Rename(jsonFilePath, destination); err != nil {
			return fmt.Errorf("failed to move vanity key from %s to %s: %w", jsonFilePath, destination, err)
		}
		lggr.Infow("Generated vanity key", "program", program, "key", keys[program])
	}
	return nil
}

func buildProject(lggr logger.Logger, repoDir string) error {
	solanaDir := filepath.Join(repoDir, anchorDir, "..")
	lggr.Infow("Building project", "dir", solanaDir)
	output, err := RunCommand("make", []string{"docker-build-contracts"}, solanaDir)
	if err != nil {
		return fmt.Errorf("anchor build failed: %s %w", output, err)
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// programTypeToDeployName maps a contract type to its compiled program base name.
var programDeployNames = map[cldf.ContractType]string{
	"Router":                   "ccip_router",
	"FeeQuoter":                "fee_quoter",
	"OffRamp":                  "ccip_offramp",
	"BurnMintTokenPool":        "burnmint_token_pool",
	"LockReleaseTokenPool":     "lockrelease_token_pool",
	"RMNRemote":                "rmn_remote",
	"AccessControllerProgram":  "access_controller",
	"ManyChainMultiSigProgram": "mcm",
	"RBACTimelockProgram":      "timelock",
	"CCTPTokenPool":            "cctp_token_pool",
}

func programTypeToDeployName(ct cldf.ContractType) string {
	if name, ok := programDeployNames[ct]; ok {
		return name
	}
	return string(ct)
}
