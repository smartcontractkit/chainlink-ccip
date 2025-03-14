package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type JdURLs struct {
	GRPCHostURL     string `json:"grpc_host_url"`
	GRCPInternalUrl string `json:"grpc_internal_url"`
	WSHostURL       string `json:"ws_host_url"`
	WSInternalURL   string `json:"ws_internal_url"`
}

type ChainURLs struct {
	HTTPHostURL     string `json:"http_host_url"`
	WSHostURL       string `json:"ws_host_url"`
	HTTPInternalURL string `json:"http_internal_url"`
	WSInternalURL   string `json:"ws_internal_url"`
}

type DonURLs struct {
	BootstrapNodes []DonURL `json:"bootstrap_nodes"`
	WorkerNodes    []DonURL `json:"worker_nodes"`
}

type DonAPICredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DonURL struct {
	HostURL        string `json:"host_url"`
	InternalURL    string `json:"internal_url"`
	P2PInternalURL string `json:"p2p_internal_url"`
	InternalIP     string `json:"internal_ip"`
}

var noOpFn = func(string) {}
var printFn = func(s string) { fmt.Println(s) }

func main() {
	// main flags
	jdFlagPtr := flag.Bool("jd", false, "Generate JD URLs")
	chainFlagPtr := flag.Bool("chain", false, "Generate chain URLs")
	donFlagPtr := flag.Bool("don", false, "Generate DON URLs")
	targetDirPtr := flag.String("target-dir", "", "Target directory (optional)")

	// chain flags
	chainTypeFlagPtr := flag.String("chain-type", "", "Chain type: evm, solana")
	chainVariantFlagPtr := flag.String("chain-variant", "", "Chain variant: besu, geth")
	chainNameFlagPtr := flag.String("chain-name", "", "Chain name")

	flag.Parse()

	if !*jdFlagPtr && !*chainFlagPtr && !*donFlagPtr {
		flag.Usage()
		return
	}

	namespace := os.Getenv("DEVSPACE_NAMESPACE")
	if namespace == "" {
		panic("DEVSPACE_NAMESPACE is not set")
	}

	ingressDomain := os.Getenv("DEVSPACE_INGRESS_DOMAIN")
	if ingressDomain == "" {
		panic("DEVSPACE_INGRESS_DOMAIN is not set")
	}

	targetDir := "."
	if targetDirPtr != nil && *targetDirPtr != "" {
		targetDir = *targetDirPtr
	}

	if jdFlagPtr != nil && *jdFlagPtr {
		fmt.Println("JD URLs:")
		jdUrls := generateJDUrls(namespace, ingressDomain)

		err := saveToFile("jd-urls.json", targetDir, jdUrls, printFn)
		if err != nil {
			panic(err)
		}
	}

	if chainFlagPtr != nil && *chainFlagPtr {
		chainID := os.Getenv("CHAIN_ID")
		if chainID == "" {
			panic("CHAIN_ID is not set")
		}

		_, err := strconv.Atoi(chainID)
		if err != nil {
			panic("CHAIN_ID is not a number")
		}

		if *chainTypeFlagPtr == "" {
			flag.Usage()
			panic("CHAIN_TYPE is not set")
		}

		if strings.EqualFold(*chainTypeFlagPtr, EVM) && *chainVariantFlagPtr == "" {
			flag.Usage()
			panic("CHAIN_VARIANT is not set")
		}

		chainUrls := generateChainUrls(namespace, ingressDomain, *chainTypeFlagPtr, chainID, chainVariantFlagPtr, chainNameFlagPtr)
		fmt.Println("Chain URLs:")

		err = saveToFile(fmt.Sprintf("chain-%s-urls.json", chainID), targetDir, chainUrls, printFn)
		if err != nil {
			panic(err)
		}
	}

	if donFlagPtr != nil && *donFlagPtr {
		bootstrapNodeCountStr := os.Getenv("DON_BOOT_NODE_COUNT")
		if bootstrapNodeCountStr == "" {
			panic("DON_BOOT_NODE_COUNT is not set")
		}

		boostrapNodeCount, err := strconv.Atoi(bootstrapNodeCountStr)
		if err != nil {
			panic("DON_BOOT_NODE_COUNT is not a number")
		}

		workerNodeCountStr := os.Getenv("DON_NODE_COUNT")
		if workerNodeCountStr == "" {
			panic("DON_NODE_COUNT is not set")
		}

		workerNodeCount, err := strconv.Atoi(workerNodeCountStr)
		if err != nil {
			panic("DON_NODE_COUNT is not a number")
		}

		donType := os.Getenv("DON_TYPE")
		if donType == "" {
			panic("DON_TYPE is not set")
		}

		donUrls := generateDONUrls(namespace, ingressDomain, donType, boostrapNodeCount, workerNodeCount)
		fmt.Println("DON URLs:")

		err = saveToFile(fmt.Sprintf("don-%s-urls.json", donType), targetDir, donUrls, printFn)
		if err != nil {
			panic(err)
		}

		if os.Getenv("DON_API_USERNAME") == "" {
			panic("DON_API_USERNAME is not set")
		}

		if os.Getenv("DON_API_PASSWORD") == "" {
			panic("DON_API_PASSWORD is not set")
		}

		err = saveToFile("don-api-credentials.json", targetDir, DonAPICredentials{
			Username: os.Getenv("DON_API_USERNAME"),
			Password: os.Getenv("DON_API_PASSWORD"),
		}, noOpFn)

		if err != nil {
			panic(err)
		}
	}
}

func generateDONUrls(namespace, ingressDomain, donType string, boostrapNodeCount, workerNodeCount int) DonURLs {
	bootstrapNodes := make([]DonURL, boostrapNodeCount)
	workerNodes := make([]DonURL, workerNodeCount)

	for i := range boostrapNodeCount {
		bootstrapNodes[i] = DonURL{
			HostURL:        fmt.Sprintf("http://%s-%s-bt-%d.%s:80", namespace, donType, i, ingressDomain),
			InternalURL:    fmt.Sprintf("http://%s-%s-bt-%d:80", namespace, donType, i),
			P2PInternalURL: fmt.Sprintf("http://%s-%s-bt-%d:6690", namespace, donType, i),
			InternalIP:     fmt.Sprintf("%s-bt-%d", donType, i),
		}
	}

	for i := range workerNodeCount {
		workerNodes[i] = DonURL{
			HostURL:        fmt.Sprintf("http://%s-%s-%d.%s:80", namespace, donType, i, ingressDomain),
			InternalURL:    fmt.Sprintf("http://%s-%s-%d:80", namespace, donType, i),
			P2PInternalURL: fmt.Sprintf("http://%s-%s-%d:6690", namespace, donType, i),
			InternalIP:     fmt.Sprintf("%s-%d", donType, i),
		}
	}

	return DonURLs{
		BootstrapNodes: bootstrapNodes,
		WorkerNodes:    workerNodes,
	}

}

func generateChainUrls(namespace, ingressDomain string, chainType, chainID string, chainVariant, chainName *string) ChainURLs {
	return ChainURLs{

		HTTPHostURL:     externalHTTPRPC(chainType, chainVariant, namespace, ingressDomain, chainID, chainName),
		WSHostURL:       externalChainWSRPC(chainType, chainVariant, namespace, ingressDomain, chainID, chainName),
		HTTPInternalURL: mustInternalHTTPRPC(chainType, chainVariant, chainID, chainName),
		WSInternalURL:   mustInternalWSRPC(chainType, chainVariant, chainID, chainName),
	}
}

func saveToFile(filename, targetDir string, data any, afterSaveHook func(string)) error {
	filePath := filepath.Join(targetDir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()
	marshalled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(marshalled)
	if err != nil {
		return err
	}

	afterSaveHook(string(marshalled))
	fmt.Printf("File %s saved\n", filePath)

	return nil
}

func generateJDUrls(namespace, ingressDomain string) JdURLs {
	return JdURLs{
		GRPCHostURL:     fmt.Sprintf("%s-job-distributor-grpc.%s:443", namespace, ingressDomain),
		WSHostURL:       "", // empty on purpose, since it is not exposed to the outside world
		GRCPInternalUrl: "job-distributor:80",
		WSInternalURL:   "job-distributor-noderpc-lb:80",
	}
}

type ChainType = string

const (
	EVM    ChainType = "evm"
	SOLANA ChainType = "solana"
)

type ChainVariant = string

const (
	Besu ChainVariant = "besu"
	Geth ChainVariant = "geth"
)

func externalChainWSRPC(chainType ChainType, chainVariant *ChainVariant, namespace, ingressDomain, chainID string, chainName *string) string {
	var u string
	if strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "besu") {
		u = fmt.Sprintf("wss://chain-%s-rpc.%s/ws/", *chainName, ingressDomain)
	} else {
		u = fmt.Sprintf("wss://%s-%s-%s-ws.%s", namespace, mustChainTypeHostNamePart(chainType, *chainVariant), chainID, ingressDomain)
	}
	return u
}

func externalHTTPRPC(chainType ChainType, chainVariant *ChainVariant, namespace, ingressDomain, chainID string, chainName *string) string {
	var u string
	if strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "besu") {
		u = fmt.Sprintf("https://chain-%s-rpc.%s", *chainName, ingressDomain)
	} else {
		u = fmt.Sprintf("https://%s-%s-%s-http.%s:443", namespace, mustChainTypeHostNamePart(chainType, *chainVariant), chainID, ingressDomain)
	}
	return u
}

func mustInternalWSRPC(chainType ChainType, chainVariant *ChainVariant, chainID string, chainName *string) string {
	var u string

	switch {
	case strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "besu"):
		u = fmt.Sprintf("ws://%s-node-rpc-1.chain-%s.svc.cluster.local:8546", strings.ToLower(*chainVariant), *chainName)
	case strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "geth"):
		u = fmt.Sprintf("ws://%s-%s-ws:8546", strings.ToLower(*chainVariant), chainID)
	case strings.EqualFold(chainType, SOLANA):
		u = fmt.Sprintf("ws://%s-%s:8545", strings.ToLower(chainType), chainID)
	default:
		panic(fmt.Sprintf("unsupported chain type: %s and variant: %s", chainType, *chainVariant))
	}
	return u
}

func mustInternalHTTPRPC(chainType ChainType, chainVariant *ChainVariant, chainID string, chainName *string) string {
	var u string
	switch {
	case strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "besu"):
		u = fmt.Sprintf("http://%s-node-rpc-1.chain-%s.svc.cluster.local:8545", string(*chainVariant), *chainName)
	case strings.EqualFold(chainType, EVM) && strings.EqualFold(*chainVariant, "geth"):
		u = fmt.Sprintf("http://%s-%s:8544", strings.ToLower(string(*chainVariant)), chainID)
	case strings.EqualFold(chainType, SOLANA):
		u = fmt.Sprintf("http://%s-%s:8544", strings.ToLower(chainType), chainID)
	default:
		panic(fmt.Sprintf("unsupported chain type: %s and variant: %v", chainType, *chainVariant))

	}
	return u
}

func mustChainTypeHostNamePart(chainType ChainType, chainVariant ChainVariant) string {
	if strings.EqualFold(chainType, EVM) && strings.EqualFold(chainVariant, "besu") {
		return "besu"
	} else if strings.EqualFold(chainType, EVM) && strings.EqualFold(chainVariant, "geth") {
		return "geth"
	} else if strings.EqualFold(chainType, SOLANA) {
		return "solana"
	}
	panic(fmt.Sprintf("unsupported chain type: %s and variant: %s", chainType, chainVariant))
}
