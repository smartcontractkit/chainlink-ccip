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

	if *gethChainsCountPtr == 0 && *solanaChainsCountPtr == 0 {
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
		panic(err)
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
	chains := make([]SolanaChain, 0)
	for i := 1; i <= count; i++ {
		chains = append(chains, SolanaChain{ChainId: int64(1000 + i)})
	}
	return chains
}

func BuildEVMNetworkConfigs(chainsCount int) []EVMChain {
	chains := make([]EVMChain{}, 0)

	// Generate chains starting from NetworkId 1337
	for i := 0; i < chainsCount; i++ {
		networkId := int64(1337 + i) // The first chain will have NetworkId 1337
		chains = append(chains, EVMChain{NetworkId: networkId, FinalityDepth: defaultFinalityDepth})
	}

	return chains
}
