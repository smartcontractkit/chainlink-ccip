package v1_5

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	commit_store "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/commit_store"
	evm_2_evm_offramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
	evm_2_evm_onramp "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	rmn_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"

	erc20 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/rmn_proxy_contract"
	token_admin_registry "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_admin_registry"
	token_pool "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
)

// Parsed ABIs for v1.5 contracts
var (
	TokenAdminRegistryABI abi.ABI
	TokenPoolABI          abi.ABI
	RMNProxyABI           abi.ABI
	ERC20ABI              abi.ABI
	CommitStoreABI        abi.ABI
	RMNContractABI        abi.ABI
	EVM2EVMOnRampABI      abi.ABI
	EVM2EVMOffRampABI     abi.ABI
)

func init() {
	// Parse the TokenAdminRegistry ABI once at startup
	parsed, err := token_admin_registry.TokenAdminRegistryMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse TokenAdminRegistry ABI: %v", err))
	}
	TokenAdminRegistryABI = *parsed

	// Parse the TokenPool ABI once at startup
	parsedPool, err := token_pool.TokenPoolMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse TokenPool ABI: %v", err))
	}
	TokenPoolABI = *parsedPool

	// Parse the RMNProxy ABI once at startup
	parsedProxy, err := rmn_proxy.RMNProxyMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse RMNProxy ABI: %v", err))
	}
	RMNProxyABI = *parsedProxy

	// Parse the ERC20 ABI once at startup
	parsedERC20, err := erc20.FactoryBurnMintERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse ERC20 ABI: %v", err))
	}
	ERC20ABI = *parsedERC20

	// Parse the CommitStore ABI once at startup
	parsedCS, err := commit_store.CommitStoreMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse CommitStore v1.5 ABI: %v", err))
	}
	CommitStoreABI = *parsedCS

	// Parse the RMNContract ABI once at startup
	parsedRMN, err := rmn_contract.RMNContractMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse RMNContract v1.5 ABI: %v", err))
	}
	RMNContractABI = *parsedRMN

	// Parse the EVM2EVMOnRamp ABI once at startup
	parsedOnRamp, err := evm_2_evm_onramp.EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse EVM2EVMOnRamp v1.5 ABI: %v", err))
	}
	EVM2EVMOnRampABI = *parsedOnRamp

	// Parse the EVM2EVMOffRamp ABI once at startup
	parsedOffRamp, err := evm_2_evm_offramp.EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse EVM2EVMOffRamp v1.5 ABI: %v", err))
	}
	EVM2EVMOffRampABI = *parsedOffRamp

	// Register v1.5 views
	views.Register("evm", "TokenAdminRegistry", "1.5.0", ViewTokenAdminRegistry)
	views.Register("evm", "OnRamp", "1.5.0", ViewOnRamp)
	views.Register("evm", "OffRamp", "1.5.0", ViewOffRamp)
	views.Register("evm", "RegistryModuleOwnerCustom", "1.5.0", ViewRegistryModuleOwnerCustom)

	// CommitStore and RMN views
	views.Register("evm", "CommitStore", "1.5.0", ViewCommitStore)
	views.Register("evm", "RMN", "1.5.0", ViewRMN)

	// Token Pool views (v1.5.0)
	views.Register("evm", "BurnMintTokenPool", "1.5.0", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.5.0", ViewLockReleaseTokenPool)
	views.Register("evm", "BurnMintTokenPoolAndProxy", "1.5.0", ViewBurnMintTokenPoolAndProxy)
	views.Register("evm", "LockReleaseTokenPoolAndProxy", "1.5.0", ViewLockReleaseTokenPoolAndProxy)
}
