package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

const tmplFile = "values.yaml.tmpl"
const defaultFinalityDepth = 200

type chain struct {
	NetworkId     int64
	FinalityDepth int
}
type Config struct {
	Chains []chain
}

func main() {

	chainsCountPtr := flag.Int("chains-count", 0, "Number of network configs to generate")
	flag.Parse()

	if chainsCountPtr == nil || *chainsCountPtr == 0 {
		fmt.Println("Error: chains-count is required for network-generated profile")
		flag.Usage()
		return
	}

	chainsCount := *chainsCountPtr

	c := BuildNetworkConfigs(chainsCount)

	tmpl, err := template.New("values.yaml.tmpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}

func BuildNetworkConfigs(chainsCount int) Config {
	chains := []chain{
		{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
	}

	if chainsCount > 1 {
		chains = append(chains, chain{NetworkId: 2337, FinalityDepth: defaultFinalityDepth})

		for i := 1; i < chainsCount-1; i++ {

			chains = append(chains, chain{NetworkId: int64(90000000 + i), FinalityDepth: defaultFinalityDepth})
		}
	}

	// Depending on how many variations of config we need, we can hardcode it here,
	// or read from some config files
	c := Config{
		Chains: chains,
	}
	return c
}
