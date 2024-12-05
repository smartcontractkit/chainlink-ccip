package ccip

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
)

type OutputReader struct {
	outputDir string
}

func NewOutputReader(outputDir string) *OutputReader {
	return &OutputReader{outputDir: outputDir}
}

func (r *OutputReader) ReadNodesDetails() NodesDetails {
	byteValue := r.readFile(NodesDetailsFileName)

	var result NodesDetails

	// Unmarshal the JSON into the map
	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}

	return result
}

func (r *OutputReader) ReadChainConfigs() []devenv.ChainConfig {
	byteValue := r.readFile(ChainsConfigsFileName)

	var result []devenv.ChainConfig

	// Unmarshal the JSON into the map
	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}

	return result
}

func ToChainConfigs(cc []ChainConfig) []devenv.ChainConfig {
	chainConfigs := make([]devenv.ChainConfig, 0)
	for _, c := range cc {
		chainConfigs = append(chainConfigs, devenv.ChainConfig{
			ChainID:   c.ChainID,
			ChainName: c.ChainName,
			ChainType: c.ChainType,
			WSRPCs:    c.WSRPCs,
			HTTPRPCs:  c.HTTPRPCs,
		})
	}
	return chainConfigs
}

func (r *OutputReader) ReadAddressBook() *deployment.AddressBookMap {
	byteValue := r.readFile(AddressBookFileName)

	var result map[uint64]map[string]deployment.TypeAndVersion

	// Unmarshal the JSON into the map
	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}

	return deployment.NewMemoryAddressBookFromMap(result)
}

func (r *OutputReader) readFile(fileName string) []byte {
	file, err := os.Open(fmt.Sprintf("%s/%s", r.outputDir, fileName))
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	// Read the file's content into a byte slice
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}
	return byteValue
}
