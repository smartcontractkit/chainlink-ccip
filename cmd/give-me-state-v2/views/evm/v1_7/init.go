package v1_7

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	committee_verifier "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	erc20 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	fee_quoter "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	offramp "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/offramp"
	onramp "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/onramp"
)

// Parsed ABIs for v1.7 contracts
var (
	FeeQuoterABI         abi.ABI
	CommitteeVerifierABI abi.ABI
	OnRampABI            abi.ABI
	OffRampABI           abi.ABI
	ERC20ABI             abi.ABI
)

func init() {
	// Parse the FeeQuoter ABI once at startup
	parsed, err := fee_quoter.FeeQuoterMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse FeeQuoter ABI: %v", err))
	}
	FeeQuoterABI = *parsed

	// Parse the CommitteeVerifier ABI once at startup
	parsedCV, err := committee_verifier.CommitteeVerifierMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse CommitteeVerifier ABI: %v", err))
	}
	CommitteeVerifierABI = *parsedCV

	// Parse the OnRamp ABI once at startup
	parsedOnRamp, err := onramp.OnRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse OnRamp ABI: %v", err))
	}
	OnRampABI = *parsedOnRamp

	// Parse the OffRamp ABI once at startup
	parsedOffRamp, err := offramp.OffRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse OffRamp ABI: %v", err))
	}
	OffRampABI = *parsedOffRamp

	// Parse the ERC20 ABI once at startup
	parsedERC20, err := erc20.FactoryBurnMintERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse ERC20 ABI: %v", err))
	}
	ERC20ABI = *parsedERC20

	// Register v1.7 views
	views.Register("evm", "FeeQuoter", "1.7.0", ViewFeeQuoter)
	views.Register("evm", "CommitteeVerifier", "1.7.0", ViewCommitteeVerifier)
	views.Register("evm", "OnRamp", "1.7.0", ViewOnRamp)
	views.Register("evm", "OffRamp", "1.7.0", ViewOffRamp)
}
