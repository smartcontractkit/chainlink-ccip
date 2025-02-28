package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

const tmplFile = "values.yaml.tmpl"
const defaultFinalityDepth = 200

type EVMChain struct {
	NetworkId     int64
	FinalityDepth int
}

type SolanaChain struct {
	ChainId int64
}

type Config struct {
	BesuChains   []EVMChain
	GethChains   []EVMChain
	SolanaChains []SolanaChain
}

func main() {
	besuChainsCountPtr := flag.Int("besu-chains-count", 0, "Number of network configs to generate")
	gethChainsCountPtr := flag.Int("geth-chains-count", 0, "Number of network configs to generate")
	solanaChainsCountPtr := flag.Int("solana-chains-count", 0, "Number of network configs to generate")
	flag.Parse()

	if *besuChainsCountPtr == 0 && *gethChainsCountPtr == 0 && *solanaChainsCountPtr == 0 {
		fmt.Println("Error: at least one chain is required, to generate config")
		flag.Usage()
		return
	}
	besuChainsCount := *besuChainsCountPtr
	gethChainsCount := *gethChainsCountPtr
	solanaChainsCount := *solanaChainsCountPtr

	c := BuildNetworkConfigs(besuChainsCount, gethChainsCount, solanaChainsCount)

	tmpl, err := template.New("values.yaml.tmpl").ParseFiles(tmplFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}

func BuildNetworkConfigs(besuChainsCount int, gethChainsCount int, solanaChainsCount int) Config {

	besuChains := BuildEVMNetworkConfigs(besuChainsCount)
	gethChains := BuildEVMNetworkConfigs(gethChainsCount)
	solanaChains := BuildSolanaNetworkConfigs(solanaChainsCount)

	c := Config{
		BesuChains:   besuChains,
		GethChains:   gethChains,
		SolanaChains: solanaChains,
	}
	return c
}

func BuildSolanaNetworkConfigs(count int) []SolanaChain {
	chains := make([]SolanaChain, 0, count)
	for i := 1; i <= count; i++ {
		chains = append(chains, SolanaChain{ChainId: int64(1000 + i)})
	}
	return chains
}

func BuildEVMNetworkConfigs(chainsCount int) []EVMChain {
	// If chainsCount is 0, return an empty slice
	if chainsCount == 0 {
		return []EVMChain{}
	}

	// Initialize the chains slice
	chains := []EVMChain{
		{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
	}

	// Add the second chain if chainsCount > 1
	if chainsCount > 1 {
		chains = append(chains, EVMChain{NetworkId: 2337, FinalityDepth: defaultFinalityDepth})
	}

	// Add subsequent chains starting from 90000000 if chainsCount > 2
	for i := 2; i < chainsCount; i++ {
		networkId := int64(90000000 + i - 1)
		chains = append(chains, EVMChain{NetworkId: networkId, FinalityDepth: defaultFinalityDepth})
	}

	return chains
}
