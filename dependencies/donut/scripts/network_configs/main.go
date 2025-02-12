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
	GethChains   []EVMChain
	SolanaChains []SolanaChain
}

func main() {

	gethChainsCountPtr := flag.Int("geth-chains-count", 0, "Number of network configs to generate")
	solanaChainsCountPtr := flag.Int("solana-chains-count", 0, "Number of network configs to generate")
	flag.Parse()

	if *gethChainsCountPtr == 0 && *solanaChainsCountPtr == 0 {
		fmt.Println("Error: at least one chain is required, to generate config")
		flag.Usage()
		return
	}

	gethChainsCount := *gethChainsCountPtr
	solanaChainsCount := *solanaChainsCountPtr

	c := BuildNetworkConfigs(gethChainsCount, solanaChainsCount)

	tmpl, err := template.New("values.yaml.tmpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}

func BuildNetworkConfigs(gethChainsCount int, solanaChainsCount int) Config {
	gethChains := BuildGethNetworkConfigs(gethChainsCount)
	solanaChains := BuildSolanaNetworkConfigs(solanaChainsCount)

	c := Config{
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

func BuildGethNetworkConfigs(chainsCount int) []EVMChain {
	chains := []EVMChain{
		{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
	}

	if chainsCount > 1 {
		chains = append(chains, EVMChain{NetworkId: 2337, FinalityDepth: defaultFinalityDepth})

		for i := 1; i < chainsCount-1; i++ {

			chains = append(chains, EVMChain{NetworkId: int64(90000000 + i), FinalityDepth: defaultFinalityDepth})
		}
	}
	return chains
}
