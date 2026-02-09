package v1_5

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	erc20 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/rmn_proxy_contract"
	token_admin_registry "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_admin_registry"
	token_pool "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
)

// Parsed ABIs for v1.5 contracts
// Note: Some ABIs use ccv/chains/evm bindings (local), OnRamp/OffRamp use manual decoding
var (
	TokenAdminRegistryABI abi.ABI
	TokenPoolABI          abi.ABI
	RMNProxyABI           abi.ABI
	ERC20ABI              abi.ABI
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
