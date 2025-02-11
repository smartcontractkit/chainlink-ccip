package devspace

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

// KubernetesManifest represents a single manifest with Kind and Metadata
type KubernetesManifest struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name   string `yaml:"name"`
		Labels struct {
			AppKubernetesInstance string `yaml:"app.kubernetes.io/instance"`
			AppKubernetesName     string `yaml:"app.kubernetes.io/name"`
			Release               string `yaml:"Release"`
			App                   string `yaml:"app"`
		}
	} `yaml:"metadata"`
}

// splitYAML splits the input YAML data into separate Kubernetes manifests
func splitYAML(content []byte) [][]byte {
	var manifests [][]byte
	parts := bytes.Split(content, []byte("\n---\n"))

	for _, part := range parts {
		if len(bytes.TrimSpace(part)) > 0 {
			manifests = append(manifests, part)
		}
	}
	return manifests
}

// removeEscapedLines filters out lines that start with the ESC character
func removeEscapedLines(input io.Reader) string {
	var result bytes.Buffer
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && utf8.RuneStart(line[0]) && line[0] == '\x1b' {
			continue // Skip lines that start with ESC (ASCII 27)
		}
		result.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	return result.String()
}

// SplitRender processes the output from "devspace <some command> --render" and saves individual manifests
// in subdirectories grouped by kubernetes instance name. This is similar as for example `helm template` command
// with --output-dir flag
func SplitRender(outputDir string) {
	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Read and filter input from stdin
	cleanedInput := removeEscapedLines(os.Stdin)

	// Split the YAML input into separate manifests
	manifests := splitYAML([]byte(cleanedInput))

	for _, manifest := range manifests {
		var km KubernetesManifest
		if err := yaml.Unmarshal(manifest, &km); err != nil {
			log.Printf("Skipping invalid manifest: %v", err)
			continue
		}

		// Validate required fields
		if km.Kind == "" || km.Metadata.Name == "" {
			log.Println("Skipping document due to missing Kind or Metadata.Name")
			continue
		}

		// Create Subdirectory directory for each App
		appName := getInstanceDir(km)

		appNameDir := filepath.Join(outputDir, appName)
		if err := os.MkdirAll(appNameDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}

		// Create a filename like "ConfigMap-my-config.yaml"
		manifestFilePath := fmt.Sprintf("%s/%s-%s.yaml", appNameDir, km.Metadata.Name, km.Kind)

		// Write the manifest to a new file
		if err := os.WriteFile(manifestFilePath, manifest, 0o600); err != nil {
			log.Printf("Failed to write file %s: %v", manifestFilePath, err)
		} else {
			fmt.Printf("Created: %s\n", manifestFilePath)
		}
	}
}

func getInstanceDir(km KubernetesManifest) string {
	if km.Metadata.Labels.AppKubernetesInstance != "" {
		return km.Metadata.Labels.AppKubernetesInstance
	} else if km.Metadata.Labels.App != "" {
		return km.Metadata.Labels.App
	} else if km.Metadata.Labels.Release != "" {
		return km.Metadata.Labels.Release
	} else {
		return "unknown"
	}
}
