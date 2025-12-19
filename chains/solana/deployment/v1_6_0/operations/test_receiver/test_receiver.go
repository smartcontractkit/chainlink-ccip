package test_receiver

import (
	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "TestReceiver"
var Version *semver.Version = semver.MustParse("1.6.0")
var ProgramName = "test_ccip_receiver"

type Params struct {
	TokenPrivKey  solana.PrivateKey
	TokenDecimals uint8
}

var Deploy = operations.NewOperation(
	"receiver:deploy",
	Version,
	"Deploys the Receiver program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName)
	},
)
