package v1_6

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	ccip_home "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/ccip_home"
	fee_quoter "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	nonce_manager "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/nonce_manager"
	offramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	onramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	rmn_home "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	rmn_remote "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"

	token_pool "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
)

// Parsed ABIs for v1.6 contracts
var (
	OnRampABI       abi.ABI
	OffRampABI      abi.ABI
	FeeQuoterABI    abi.ABI
	NonceManagerABI abi.ABI
	CCIPHomeABI     abi.ABI
	RMNHomeABI      abi.ABI
	RMNRemoteABI    abi.ABI
	TokenPoolABI    abi.ABI
)

func init() {
	parsed, err := onramp.OnRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse OnRamp v1.6 ABI: %v", err))
	}
	OnRampABI = *parsed

	parsed, err = offramp.OffRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse OffRamp v1.6 ABI: %v", err))
	}
	OffRampABI = *parsed

	parsed, err = fee_quoter.FeeQuoterMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse FeeQuoter v1.6 ABI: %v", err))
	}
	FeeQuoterABI = *parsed

	parsed, err = nonce_manager.NonceManagerMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse NonceManager v1.6 ABI: %v", err))
	}
	NonceManagerABI = *parsed

	parsed, err = ccip_home.CCIPHomeMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse CCIPHome v1.6 ABI: %v", err))
	}
	CCIPHomeABI = *parsed

	parsed, err = rmn_home.RMNHomeMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse RMNHome v1.6 ABI: %v", err))
	}
	RMNHomeABI = *parsed

	parsed, err = rmn_remote.RMNRemoteMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse RMNRemote v1.6 ABI: %v", err))
	}
	RMNRemoteABI = *parsed

	parsed, err = token_pool.TokenPoolMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse TokenPool ABI: %v", err))
	}
	TokenPoolABI = *parsed

	// Register v1.6 views
	views.Register("evm", "FeeQuoter", "1.6.0", ViewFeeQuoter)
	views.Register("evm", "FeeQuoter", "1.6.3", ViewFeeQuoter) // Same ABI as 1.6.0
	views.Register("evm", "OnRamp", "1.6.0", ViewOnRamp)
	views.Register("evm", "OffRamp", "1.6.0", ViewOffRamp)
	views.Register("evm", "RegistryModuleOwnerCustom", "1.6.0", ViewRegistryModuleOwnerCustom)
	views.Register("evm", "NonceManager", "1.6.0", ViewNonceManager)
	views.Register("evm", "RMNRemote", "1.6.0", ViewRMNRemote)
	views.Register("evm", "BurnMintTokenPool", "1.6.0", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.6.0", ViewLockReleaseTokenPool)

	// Home contracts (deployed only on Ethereum mainnet)
	views.Register("evm", "CCIPHome", "1.6.0", ViewCCIPHome)
	views.Register("evm", "RMNHome", "1.6.0", ViewRMNHome)
}
