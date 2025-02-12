package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"text/template"
)

const templatesPath = "../../templates/"

type chain struct {
	NetworkId int64 `yaml:"networkId"`
	BlockTime int   `yaml:"blockTime"`
}
type ValuesTmplConfig struct {
	Chains   []chain `yaml:"chains"`
	Provider string
	Product  string
	UseCase  string
}

func main() {
	// Define command-line flags
	chainOverridesDir := flag.String("chain-overrides-dir", "../../values/chain-overrides", "Directory with chain overrides")
	chainOverridesFileName := flag.String("chain-overrides-file", "default.yaml", "Chain overrides file name")
	provider := flag.String("provider", "", "CRIB Provider")
	product := flag.String("product", "chainlink", "Product, chainlink or ccip")
	useCase := flag.String("use-case", "default", "Use case for an environment")
	chainsCount := flag.Int("chains-count", 0, "Number of chains to generate")
	templateFile := flag.String("template", "base.yaml.tmpl", "template file to use for rendering configuration")

	// Parse flags
	flag.Parse()

	// Validate input
	if *chainOverridesDir == "" || *chainOverridesFileName == "" {
		fmt.Println("Error: Both string parameters are required.")
		flag.Usage()
		return
	}

	if chainsCount == nil || *chainsCount == 0 {
		fmt.Println("Error: chains-count is required.")
		flag.Usage()
		return
	}

	filePath := path.Join(*chainOverridesDir, *chainOverridesFileName)
	GenerateConfig(filePath, *templateFile, *chainsCount, *provider, *product, *useCase)
}
func GenerateConfig(overridesFilePath string, templateFile string, chainsCount int, provider string, product string, useCase string) {
	tmplConfig, err := BuildTemplateConfig(overridesFilePath, chainsCount, provider, product, useCase)

	tmpl, err := template.New(templateFile).ParseFiles(path.Join(templatesPath, templateFile))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, tmplConfig)
	if err != nil {
		panic(err)
	}
}

func BuildTemplateConfig(overridesFilePath string, chainsCount int, provider string, product string, useCase string) (ValuesTmplConfig, error) {
	data, err := os.ReadFile(overridesFilePath)
	if err != nil {
		fmt.Printf("Error reading chain overrides file: %s : %s\n", overridesFilePath, err)
		os.Exit(1)
	}

	var overrides ValuesTmplConfig
	// Unmarshal YAML into the struct
	err = yaml.Unmarshal(data, &overrides)
	if err != nil {
		fmt.Printf("Failed to unmarshal YAML: %v\n", err)
		os.Exit(1)
	}

	if len(overrides.Chains) > chainsCount {
		fmt.Printf("Invalid overrides config, number of chains in the override file: %d  "+
			"can't be larger that chainsCount: %d", len(overrides.Chains), chainsCount)
		os.Exit(1)
	}

	chains := []chain{
		{NetworkId: 1001},
	}

	if chainsCount > 1 {
		chains = append(chains, chain{NetworkId: 1002})

		for i := 1; i < chainsCount-1; i++ {
			chains = append(chains, chain{NetworkId: int64(90000000 + i)})
		}
	}

	for i := 0; i < chainsCount; i++ {
		for _, override := range overrides.Chains {
			if override.NetworkId == chains[i].NetworkId {
				chains[i].BlockTime = override.BlockTime
			}
		}
	}

	tmplConfig := ValuesTmplConfig{
		Chains:   chains,
		Provider: provider,
		Product:  product,
		UseCase:  useCase,
	}

	return tmplConfig, err
}
