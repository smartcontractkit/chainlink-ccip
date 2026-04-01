package verification

import (
	"github.com/Masterminds/semver/v3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var contractMetadata = map[cldf.ContractType]map[semver.Version]rawContractInfo{}

// RegisterContractMetadataForTest seeds contractMetadata so NeedsVerification and
// LoadSolidityContractMetadata succeed in tests. Not for production use.
func RegisterContractMetadataForTest(contractType cldf.ContractType, version semver.Version, solidityStandardJSON string, bytecode string, name string) {
	if contractMetadata[contractType] == nil {
		contractMetadata[contractType] = map[semver.Version]rawContractInfo{}
	}
	contractMetadata[contractType][version] = rawContractInfo{
		SolidityStandardJSONInput: solidityStandardJSON,
		Bytecode:                  bytecode,
		Name:                      name,
	}
}
