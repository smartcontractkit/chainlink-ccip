package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	// "time"
	//
	// "github.com/pelletier/go-toml/v2"
	// "github.com/smartcontractkit/chainlink-common/pkg/config"
	// mnCfg "github.com/smartcontractkit/chainlink-framework/multinode/config"
	solcfg "github.com/smartcontractkit/chainlink-solana/pkg/solana/config"
)

const tmplFile = "values.yaml.tmpl"
const defaultFinalityDepth = 200
const SolStartingPodId = 1000

type EVMChain struct {
	NetworkId     int64
	FinalityDepth int
}

type SolanaChain struct {
	NetworkId   string
	ChainId     string
	PodId       int
	ChainConfig solcfg.TOMLConfig
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
	if count > 3 {
		panic("Up to 3 sol chains are supported")
	}
	selectors := []string{"22222222222222222222222222222222222222222222", "33333333333333333333333333333333333333333333", "44444444444444444444444444444444444444444444"}
	chains := make([]SolanaChain, 0, count)
	chainConfig := solcfg.NewDefault()
	chainConfig.Chain.ComputeUnitLimitDefault = ptr(uint32(100))
	chainConfig.MultiNode.MultiNode.VerifyChainID = ptr(false)

	for i := 0; i <= count-1; i++ {
		chains = append(chains, SolanaChain{NetworkId: selectors[i], ChainId: selectors[i], PodId: SolStartingPodId + i, ChainConfig: *chainConfig})
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

func ptr[T any](t T) *T {
	return &t
}
