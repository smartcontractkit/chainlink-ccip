package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

const templatesPath = "../../templates/"

type rmnNode struct {
	NodeID            int `yaml:"nodeId"`
	RageproxyKeystore string
	Passphrase        string
	RMNKeystore       string
}
type ValuesTmplConfig struct {
	RMNNodes         []*rmnNode `yaml:"rmnNodes"`
	Provider         string
	SharedConfigToml string
	LocalConfigToml  string
}

func main() {
	// Define command-line flags
	provider := flag.String("provider", "", "CRIB Provider")
	rmnNodesCount := flag.Int("nodes-count", 0, "Number of nodes to generate")
	templateFile := flag.String("template", "base.yaml.tmpl", "template file to use for rendering configuration")
	ccipEnvStateDir := flag.String("ccip-env-state-dir", ".tmp", "directory with ccip-env-state")

	// Parse flags
	flag.Parse()

	if rmnNodesCount == nil || *rmnNodesCount == 0 {
		fmt.Println("Error: nodes-count is required.")
		flag.Usage()
		return
	}

	GenerateConfig(*templateFile, *rmnNodesCount, *provider, *ccipEnvStateDir)
}
func GenerateConfig(templateFile string, nodesCount int, provider string, ccipEnvStateDir string) {
	tmplConfig := BuildTemplateConfig(nodesCount, provider, ccipEnvStateDir)

	tmpl, err := template.New(templateFile).Funcs(template.FuncMap{
		"indent": indent,
	}).ParseFiles(path.Join(templatesPath, templateFile))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, tmplConfig)
	if err != nil {
		panic(err)
	}
}

func BuildTemplateConfig(nodesCount int, provider string, ccipEnvStateDir string) ValuesTmplConfig {
	nodes := make([]*rmnNode, 0)
	for i := 0; i < nodesCount; i++ {
		nodes = append(nodes, &rmnNode{NodeID: i})
	}

	dataSubDir := "data"
	rmnSubDir := "rmn"
	sharedConfigToml := readFile(path.Join(ccipEnvStateDir, rmnSubDir, "rmn-shared-config.toml"))
	localConfigToml := readFile(path.Join(ccipEnvStateDir, rmnSubDir, "rmn-local-config.toml"))
	keystore := readFile(path.Join(ccipEnvStateDir, dataSubDir, "rmn-node-identities.json"))
	// Unmarshal the JSON keystore into an array of map
	var keystoreMap []map[string]interface{}
	if err := json.Unmarshal([]byte(keystore), &keystoreMap); err != nil {
		fmt.Println("Error unmarshalling keystore JSON:", err)
		os.Exit(2)
	}

	for i, node := range nodes {
		node.RageproxyKeystore = keystoreMap[i]["RageProxyKeystore"].(string)
		node.Passphrase = keystoreMap[i]["Passphrase"].(string)
		node.RMNKeystore = keystoreMap[i]["RMNKeystore"].(string)
	}

	tmplConfig := ValuesTmplConfig{
		SharedConfigToml: sharedConfigToml,
		LocalConfigToml:  localConfigToml,
		RMNNodes:         nodes,
		Provider:         provider,
	}

	return tmplConfig
}

// indent adds a specified number of spaces before each line of a string
func indent(spaces int, input string) string {
	prefix := strings.Repeat(" ", spaces)
	return prefix + strings.ReplaceAll(input, "\n", "\n"+prefix)
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(2)
	}
	defer file.Close()

	// Read the file's content into a byte slice
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
	}
	return string(byteValue)
}
