/*
* Package cmd implements the git-charts utility for managing Helm chart git references.
*
* Key assumptions:
* 1. All DevSpace yaml files reside in ./dependencies/* directory
* 2. Only processes files with deployments and profiles sections
* 3. Only handles charts referenced via CHAINLINK_HELM_REGISTRY_URI
* 5. Each chart must have a Chart.yaml with valid semver version
* 6. Git history of Chart.yaml must be available to find version matches
* 7. The git-charts profile must exist and use the replace directive for deployments
* 8. Chart references in git-charts profile must use git/subPath/revision format
*
* The tool matches chart versions to git commits and updates revision fields in
* DevSpace yaml files while preserving the rest of the file structure.
**/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const (
	reposRootKey = "repo-root"
	dryRunKey    = "dry-run"
)

type chartYaml struct {
	Version string `yaml:"version"`
}

type GitChartsProfile struct {
	Name  string `yaml:"name"`
	Merge struct {
		Deployments map[string]struct {
			Helm struct {
				Chart struct {
					Git      string `yaml:"git"`
					SubPath  string `yaml:"subPath"`
					Revision string `yaml:"revision"`
				} `yaml:"chart"`
			} `yaml:"helm"`
		} `yaml:"deployments"`
	} `yaml:"merge"`
}
type devspaceYaml struct {
	Deployments map[string]struct {
		Helm struct {
			Chart struct {
				Version string `yaml:"version"`
				Name    string `yaml:"name"`
			} `yaml:"chart"`
		} `yaml:"helm"`
	} `yaml:"deployments"`

	Profiles []GitChartsProfile `yaml:"profiles"`
}

//nolint:gochecknoinits
func init() {
	flags := gitChartsCmd.Flags()
	flags.String(reposRootKey, "..", "Root directory containing all required repos")
	flags.Bool(dryRunKey, false, "Show what changes would be made without making them")
	_ = viper.BindPFlags(flags)
	rootCmd.AddCommand(gitChartsCmd)
}

var gitChartsCmd = &cobra.Command{
	Use:   "git-charts [chart-path] [version]",
	Short: "Find git commit for a specific chart version",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		reposRoot := viper.GetString(reposRootKey)
		dryRun := viper.GetBool(dryRunKey)

		processDevspaceFiles(reposRoot, dryRun)
	},
}

// isDevspaceFile checks if a file is a devspace yaml by looking for required fields
func isDevspaceFile(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Debug("Failed to read file", "path", path, "error", err)
		return false
	}

	// Check file extension
	if !strings.HasSuffix(strings.ToLower(path), ".yaml") && !strings.HasSuffix(strings.ToLower(path), ".yml") {
		return false
	}

	// Quick check for required fields without full unmarshal
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		logger.Debug("Failed to parse YAML", "path", path, "error", err)
		return false
	}

	_, hasDeployments := result["deployments"]
	_, hasProfiles := result["profiles"]

	return hasDeployments && hasProfiles
}

func processDevspaceFiles(reposRoot string, dryRun bool) {
	depFolder := "./dependencies"
	err := filepath.Walk(depFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isDevspaceFile(path) {
			logger.Info("Processing devspace file", "path", path)
			processDevspaceYaml(path, reposRoot, dryRun)
		}
		return nil
	})
	if err != nil {
		logger.Error("Error walking dependencies folder", "error", err)
		os.Exit(1)
	}
}

func processDevspaceYaml(path string, reposRoot string, dryRun bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error("Failed to read file", "path", path, "error", err)
		return
	}

	var config devspaceYaml
	// We don't want to exit on error, because we're checking
	// the values afterwards. It's OK to partially decode the file.
	_ = yaml.Unmarshal(data, &config)

	if len(config.Deployments) == 0 {
		logger.Debug("No deployments found in devspace file, skipping", "path", path)
		return
	}

	var gitCharts GitChartsProfile
	for _, profile := range config.Profiles {
		if profile.Name == "git-charts" {
			gitCharts = profile
			break
		}
	}
	if gitCharts.Name == "" {
		logger.Debug("No git-charts profile found, skipping", "path", path)
		return
	}

	// Validate the git-charts profile fields
	if gitCharts.Merge.Deployments == nil {
		logger.Error("No deployments found in git-charts profile", "path", path)
		return
	}

	logger.Debug("Processing DevSpace configuration", "path", path, "deployments", len(config.Deployments), "repo_root", reposRoot)

	// Process each deployment that has a chart version
	yamlString := string(data)
	newContent := yamlString

	for deployName, deploy := range config.Deployments {
		if deploy.Helm.Chart.Version == "" || deploy.Helm.Chart.Name == "" {
			logger.Debug("Skipping deployment without version or name",
				"name", deployName,
				"chart", deploy.Helm.Chart.Name,
				"version", deploy.Helm.Chart.Version,
				"file", path)
			continue
		}

		override, exists := gitCharts.Merge.Deployments[deployName]
		if !exists {
			logger.Debug("No matching git-charts profile deployment found",
				"deployment", deployName,
				"path", path)
			continue
		}

		logger.Debug("Found matching git-charts profile deployment",
			"deployment", deployName,
			"git", override.Helm.Chart.Git,
			"subPath", override.Helm.Chart.SubPath,
			"revision", override.Helm.Chart.Revision,
			"path", path)

		if override.Helm.Chart.Git == "" {
			logger.Error("Missing git field in deployment", "deployment", deployName, "path", path)
			continue
		}

		logger.Debug("Processing deployment", "name", deployName, "chart", deploy.Helm.Chart.Name, "version", deploy.Helm.Chart.Version,
			"git", override.Helm.Chart.Git, "subPath", override.Helm.Chart.SubPath, "old_revision", override.Helm.Chart.Revision,
		)

		// chart repo is the last subpath of the git URL
		gitStrings := strings.Split(override.Helm.Chart.Git, "/")
		chartsRepo := gitStrings[len(gitStrings)-1]
		chartPath := filepath.Join(reposRoot, chartsRepo, override.Helm.Chart.SubPath)

		if _, err := os.Stat(chartPath); os.IsNotExist(err) {
			logger.Error("Chart directory not found",
				"chart", chartPath,
				"deployment", deployName,
				"file", path)
			continue
		}

		commit := findChartVersionQuiet(chartPath, deploy.Helm.Chart.Version)
		if commit == "" {
			logger.Error("Failed to find matching commit",
				"chart", chartPath,
				"version", deploy.Helm.Chart.Version,
				"deployment", deployName,
				"file", path)
			continue
		}

		resolvedVersion := getVersionAtCommit(chartPath, commit)
		logger.Info("Found commit for chart version",
			"chart", chartPath,
			"deployment", deployName,
			"target_version", deploy.Helm.Chart.Version,
			"resolved_version", resolvedVersion,
			"commit", commit,
			"file", path)

		// Update the revision field by directly modifying the YAML document
		newContent = updateYAMLRevision(newContent, deployName, commit)
	}

	if newContent == yamlString {
		logger.Info("No changes needed for any deployments", "file", path)
		return
	}

	if dryRun {
		logger.Info("Would update revisions (dry-run)", "file", path)
		return
	}

	if err := os.WriteFile(path, []byte(newContent), 0o600); err != nil {
		logger.Error("Failed to write updated config", "error", err, "file", path)
	} else {
		logger.Info("Updated revisions in file", "path", path)
	}
}

// updateYAMLRevision updates only the revision field inside a profile while preserving the rest of the document
func updateYAMLRevision(yamlContent, deploymentName, newRevision string) string {
	lines := strings.Split(yamlContent, "\n")
	inGitChartsProfile := false
	inReplace := false
	inDeployments := false
	inTargetDeployment := false
	profileIndent := 0

	logger.Debug("Looking for git-charts profile to update revision",
		"deployment", deploymentName,
		"revision", newRevision)

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		indent := len(line) - len(strings.TrimLeft(line, " "))

		// Track when we enter a profile
		if strings.HasPrefix(trimmed, "- name: git-charts") {
			inGitChartsProfile = true
			profileIndent = indent
			logger.Debug("Found git-charts profile", "line", i)
			continue
		}

		// Exit profile handling
		if inGitChartsProfile && strings.HasPrefix(trimmed, "-") && indent == profileIndent {
			inGitChartsProfile = false
			inReplace = false
			inDeployments = false
			inTargetDeployment = false
			continue
		}

		// Track merge section in git-charts profile
		if inGitChartsProfile && strings.HasPrefix(trimmed, "merge:") {
			inReplace = true
			continue
		}

		// Track deployments section in replace
		if inReplace && strings.HasPrefix(trimmed, "deployments:") {
			inDeployments = true
			continue
		}

		// Track specific deployment
		if inDeployments && strings.HasPrefix(trimmed, deploymentName+":") {
			inTargetDeployment = true
			logger.Debug("Found target deployment in profile", "deployment", deploymentName)
			continue
		}

		// Handle chart section and update/insert revision
		if inTargetDeployment && strings.HasPrefix(trimmed, "chart:") {
			chartIndent := indent
			// Look ahead to find version or add revision after chart
			for j := i + 1; j < len(lines); j++ {
				nextLine := lines[j]
				nextTrimmed := strings.TrimSpace(nextLine)
				nextIndent := len(nextLine) - len(strings.TrimLeft(nextLine, " "))

				// Exit if we've left the chart section
				if nextIndent <= chartIndent {
					// No version found, add revision at end of chart section
					revisionLine := fmt.Sprintf("%srevision: %s", strings.Repeat(" ", chartIndent+2), newRevision)
					lines = append(lines[:j], append([]string{revisionLine}, lines[j:]...)...)
					logger.Debug("Added new revision", "at_line", j)
					break
				}

				// Update existing revision
				if strings.HasPrefix(nextTrimmed, "revision:") {
					lines[j] = fmt.Sprintf("%srevision: %s", strings.Repeat(" ", nextIndent), newRevision)
					oldRevision := strings.TrimSpace(strings.TrimPrefix(nextTrimmed, "revision:"))
					logger.Debug("Updated existing revision", "old_revision", oldRevision, "new_revision", newRevision, "at_line", j)
					break
				}

				// Add revision after version
				if strings.HasPrefix(nextTrimmed, "version:") {
					revisionLine := fmt.Sprintf("%srevision: %s", strings.Repeat(" ", nextIndent), newRevision)
					lines = append(lines[:j+1], append([]string{revisionLine}, lines[j+1:]...)...)
					logger.Debug("Inserted revision after version", "at_line", j+1)
					break
				}
			}
			break
		}
	}

	return strings.Join(lines, "\n")
}

func getCurrentVersion(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		logger.Error("Failed to read Chart.yaml", "file", file, "error", err)
		return ""
	}
	var chart chartYaml
	if err := yaml.Unmarshal(data, &chart); err != nil {
		logger.Error("Failed to unmarshal Chart.yaml", "file", file, "error", err)
		// return ""
	}
	return chart.Version
}

func getCurrentCommitInDir(dir string) string {
	cmd := getCommand("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getGitHistory(dir string) []string {
	cmd := getCommand("git", "log", "--format=%H", "--", "Chart.yaml")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

func getVersionAtCommit(dir, commit string) string {
	cmd := getCommand("git", "show", fmt.Sprintf("%s:%s", commit, "./Chart.yaml"))
	cmd.Dir = dir
	data, err := cmd.Output()
	if err != nil {
		logger.Error("Failed to run git show", "dir", dir, "commit", commit, "error", err)
		return ""
	}

	var chart chartYaml
	if err := yaml.Unmarshal(data, &chart); err != nil {
		logger.Error("Failed to unmarshal Chart.yaml", "dir", dir, "commit", commit, "error", err)
		return ""
	}
	return chart.Version
}

// New helper that returns the commit without printing or exiting
func findChartVersionQuiet(chartPath string, targetVersion string) string {
	chartFile := filepath.Join(chartPath, "Chart.yaml")
	if _, err := os.Stat(chartFile); os.IsNotExist(err) {
		logger.Error("Chart.yaml not found", "path", chartFile, "chart_dir", chartPath)
		return ""
	}

	// Parse target version as constraint
	constraint, err := semver.NewConstraint(targetVersion)
	if err != nil {
		logger.Error("Invalid version constraint",
			"version", targetVersion,
			"error", err,
			"chart", chartPath)
		return ""
	}

	// Check current version first
	currentVersion := getCurrentVersion(chartFile)
	if v, err := semver.NewVersion(currentVersion); err == nil {
		if constraint.Check(v) {
			return getCurrentCommitInDir(chartPath)
		}
	}

	// Search through git history
	commits := getGitHistory(chartPath)
	for _, commit := range commits {
		version := getVersionAtCommit(chartPath, commit)
		if v, err := semver.NewVersion(version); err == nil {
			if constraint.Check(v) {
				return commit
			}
		}
	}

	logger.Error("No matching version found",
		"constraint", targetVersion,
		"chart", chartPath,
		"current_version", currentVersion)
	return ""
}

func getCommand(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
