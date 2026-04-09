package changesets

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	erc20ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/erc20"
	siloedpoolops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	erc20lockboxbinding "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/erc20_lock_box"
	siloedusdcpoolbinding "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	usdcproxybinding "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	cldf_evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokenadminops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	hybridpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/hybrid_lock_release_usdc_token_pool"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	validateUSDCCLDMechanismCCTPV1        = "CCTP_V1"
	validateUSDCCLDMechanismCCTPV2        = "CCTP_V2"
	validateUSDCCLDMechanismLockRelease   = "LOCK_RELEASE"
	validateUSDCCLDMechanismCCTPV2WithCCV = "CCTP_V2_WITH_CCV"
)

type ValidateUSDCCLDRolloutConfig struct {
	Chains             map[uint64]ValidateUSDCCLDRolloutChainConfig
	HomeChainLiquidity *ValidateUSDCCLDRolloutHomeChainLiquidityConfig
}

type ValidateUSDCCLDRolloutChainConfig struct {
	USDCToken          string
	TokenAdminRegistry string
	ExpectedTokenPool  string
	USDCTokenPoolProxy string
	ExpectedProxyPools *ValidateUSDCCLDRolloutProxyPoolAddresses
	RemoteChains       map[uint64]ValidateUSDCCLDRolloutRemoteChainConfig
}

type ValidateUSDCCLDRolloutProxyPoolAddresses struct {
	CCTPV1Pool            string
	CCTPV2Pool            string
	CCTPV2PoolWithCCV     string
	SiloedLockReleasePool string
}

type ValidateUSDCCLDRolloutRemoteChainConfig struct {
	ExpectedMechanism   string
	ExpectedRemoteToken string
	ExpectedRemotePools []string
}

type ValidateUSDCCLDRolloutHomeChainLiquidityConfig struct {
	ChainSelector              uint64
	USDCToken                  string
	HybridLockReleaseTokenPool string
	SiloedUSDCTokenPool        string
	ExpectedTimelockAddress    string
	Checks                     map[uint64]ValidateUSDCCLDRolloutLiquidityLaneCheck
}

type ValidateUSDCCLDRolloutLiquidityLaneCheck struct {
	ExpectedWithdrawAmount    string
	PreHybridLocked           string
	ExpectedHybridLocked      string
	PreLockBoxBalance         string
	ExpectedLockBoxBalance    string
	ExpectedLiquidityProvider string
}

// ValidateUSDCCLDRollout validates the post-state of the EVM USDC CLD rollout.
// It is read-only and intended to be run after each CLD step to assert pool cutover
// state, proxy routing, remote lane configuration, and home-chain liquidity migration.
func ValidateUSDCCLDRollout() deployment.ChangeSetV2[ValidateUSDCCLDRolloutConfig] {
	return deployment.CreateChangeSet(
		applyValidateUSDCCLDRollout,
		verifyValidateUSDCCLDRollout,
	)
}

func verifyValidateUSDCCLDRollout(e deployment.Environment, cfg ValidateUSDCCLDRolloutConfig) error {
	v := &validationCollector{}
	if len(cfg.Chains) == 0 && cfg.HomeChainLiquidity == nil {
		v.addf("at least one chain check or a home-chain liquidity check must be provided")
	}

	for chainSelector, chainCfg := range cfg.Chains {
		verifyEVMChainSelector(v, chainSelector, "chains")
		if _, ok := e.BlockChains.EVMChains()[chainSelector]; !ok {
			v.addf("chains[%d]: chain selector is not available in environment", chainSelector)
		}
		verifyHexAddress(v, fmt.Sprintf("chains[%d].USDCToken", chainSelector), chainCfg.USDCToken, true)
		verifyHexAddress(v, fmt.Sprintf("chains[%d].TokenAdminRegistry", chainSelector), chainCfg.TokenAdminRegistry, true)
		verifyHexAddress(v, fmt.Sprintf("chains[%d].ExpectedTokenPool", chainSelector), chainCfg.ExpectedTokenPool, false)
		verifyHexAddress(v, fmt.Sprintf("chains[%d].USDCTokenPoolProxy", chainSelector), chainCfg.USDCTokenPoolProxy, false)
		verifyProxyPools(v, fmt.Sprintf("chains[%d].ExpectedProxyPools", chainSelector), chainCfg.ExpectedProxyPools)

		for remoteSelector, remoteCfg := range chainCfg.RemoteChains {
			verifyEVMChainSelector(v, remoteSelector, fmt.Sprintf("chains[%d].RemoteChains", chainSelector))
			if remoteCfg.ExpectedMechanism != "" && !isSupportedMechanism(remoteCfg.ExpectedMechanism) {
				v.addf("chains[%d].RemoteChains[%d].ExpectedMechanism: unsupported mechanism %q", chainSelector, remoteSelector, remoteCfg.ExpectedMechanism)
			}
			verifyHexAddress(v, fmt.Sprintf("chains[%d].RemoteChains[%d].ExpectedRemoteToken", chainSelector, remoteSelector), remoteCfg.ExpectedRemoteToken, false)
			for i, pool := range remoteCfg.ExpectedRemotePools {
				verifyHexAddress(v, fmt.Sprintf("chains[%d].RemoteChains[%d].ExpectedRemotePools[%d]", chainSelector, remoteSelector, i), pool, true)
			}
		}
	}

	if cfg.HomeChainLiquidity != nil {
		home := cfg.HomeChainLiquidity
		verifyEVMChainSelector(v, home.ChainSelector, "homeChainLiquidity")
		if _, ok := e.BlockChains.EVMChains()[home.ChainSelector]; !ok {
			v.addf("homeChainLiquidity: chain selector %d is not available in environment", home.ChainSelector)
		}
		if home.ChainSelector != chain_selectors.ETHEREUM_MAINNET.Selector && home.ChainSelector != chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector {
			v.addf("homeChainLiquidity: chain selector %d must be Ethereum mainnet or Sepolia", home.ChainSelector)
		}
		verifyHexAddress(v, "homeChainLiquidity.USDCToken", home.USDCToken, true)
		verifyHexAddress(v, "homeChainLiquidity.HybridLockReleaseTokenPool", home.HybridLockReleaseTokenPool, true)
		verifyHexAddress(v, "homeChainLiquidity.SiloedUSDCTokenPool", home.SiloedUSDCTokenPool, true)
		verifyHexAddress(v, "homeChainLiquidity.ExpectedTimelockAddress", home.ExpectedTimelockAddress, false)
		if len(home.Checks) == 0 {
			v.addf("homeChainLiquidity.Checks: at least one lane check must be provided")
		}
		for remoteSelector, laneCheck := range home.Checks {
			verifyEVMChainSelector(v, remoteSelector, "homeChainLiquidity.Checks")
			verifyHexAddress(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedLiquidityProvider", remoteSelector), laneCheck.ExpectedLiquidityProvider, false)
			verifyOptionalBigInt(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedWithdrawAmount", remoteSelector), laneCheck.ExpectedWithdrawAmount, true)
			verifyOptionalBigInt(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].PreHybridLocked", remoteSelector), laneCheck.PreHybridLocked, false)
			verifyOptionalBigInt(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedHybridLocked", remoteSelector), laneCheck.ExpectedHybridLocked, false)
			verifyOptionalBigInt(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].PreLockBoxBalance", remoteSelector), laneCheck.PreLockBoxBalance, false)
			verifyOptionalBigInt(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedLockBoxBalance", remoteSelector), laneCheck.ExpectedLockBoxBalance, false)
		}
	}

	return v.err()
}

func applyValidateUSDCCLDRollout(e deployment.Environment, cfg ValidateUSDCCLDRolloutConfig) (deployment.ChangesetOutput, error) {
	v := &validationCollector{}

	for chainSelector, chainCfg := range cfg.Chains {
		chain, ok := e.BlockChains.EVMChains()[chainSelector]
		if !ok {
			v.addf("chains[%d]: evm chain not available in environment", chainSelector)
			continue
		}
		validateChainState(e, chain, chainCfg, v)
	}

	if cfg.HomeChainLiquidity != nil {
		home := cfg.HomeChainLiquidity
		chain, ok := e.BlockChains.EVMChains()[home.ChainSelector]
		if !ok {
			v.addf("homeChainLiquidity: evm chain %d not available in environment", home.ChainSelector)
		} else {
			validateHomeChainLiquidityState(e, chain, *home, v)
		}
	}

	if err := v.err(); err != nil {
		return deployment.ChangesetOutput{}, err
	}
	return deployment.ChangesetOutput{}, nil
}

func validateChainState(
	e deployment.Environment,
	chain cldf_evm.Chain,
	cfg ValidateUSDCCLDRolloutChainConfig,
	v *validationCollector,
) {
	callOpts := &bind.CallOpts{Context: e.GetContext()}
	usdcToken := common.HexToAddress(cfg.USDCToken)

	tokenConfigReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tokenadminops.GetTokenConfig,
		chain,
		cldf_evm_contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cfg.TokenAdminRegistry),
			Args:          usdcToken,
		},
	)
	if err != nil {
		v.addf("chains[%d]: failed to read token config from token admin registry %s: %v", chain.Selector, cfg.TokenAdminRegistry, err)
		return
	}

	if cfg.ExpectedTokenPool != "" && tokenConfigReport.Output.TokenPool != common.HexToAddress(cfg.ExpectedTokenPool) {
		v.addf(
			"chains[%d]: token admin registry active pool mismatch, expected %s got %s",
			chain.Selector,
			cfg.ExpectedTokenPool,
			tokenConfigReport.Output.TokenPool.Hex(),
		)
	}

	if cfg.USDCTokenPoolProxy == "" {
		return
	}

	proxy, err := usdcproxybinding.NewUSDCTokenPoolProxy(common.HexToAddress(cfg.USDCTokenPoolProxy), chain.Client)
	if err != nil {
		v.addf("chains[%d]: failed to bind USDCTokenPoolProxy %s: %v", chain.Selector, cfg.USDCTokenPoolProxy, err)
		return
	}

	proxyToken, err := proxy.GetToken(callOpts)
	if err != nil {
		v.addf("chains[%d]: failed to read proxy token from %s: %v", chain.Selector, cfg.USDCTokenPoolProxy, err)
		return
	}
	if proxyToken != usdcToken {
		v.addf("chains[%d]: proxy token mismatch, expected %s got %s", chain.Selector, usdcToken.Hex(), proxyToken.Hex())
	}

	if supported, err := proxy.IsSupportedToken(callOpts, usdcToken); err != nil {
		v.addf("chains[%d]: failed to check proxy supported token on %s: %v", chain.Selector, cfg.USDCTokenPoolProxy, err)
	} else if !supported {
		v.addf("chains[%d]: proxy %s does not report USDC %s as supported", chain.Selector, cfg.USDCTokenPoolProxy, usdcToken.Hex())
	}

	pools, err := proxy.GetPools(callOpts)
	if err != nil {
		v.addf("chains[%d]: failed to read proxy pools from %s: %v", chain.Selector, cfg.USDCTokenPoolProxy, err)
		return
	}

	validateExpectedProxyPools(chain.Selector, cfg.ExpectedProxyPools, pools, v)

	for remoteSelector, remoteCfg := range cfg.RemoteChains {
		validateRemoteChainState(callOpts, chain.Selector, proxy, pools, remoteSelector, remoteCfg, v)
	}
}

func validateRemoteChainState(
	callOpts *bind.CallOpts,
	chainSelector uint64,
	proxy *usdcproxybinding.USDCTokenPoolProxy,
	pools usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	remoteSelector uint64,
	cfg ValidateUSDCCLDRolloutRemoteChainConfig,
	v *validationCollector,
) {
	if cfg.ExpectedMechanism == "" {
		return
	}

	mechanism, err := proxy.GetLockOrBurnMechanism(callOpts, remoteSelector)
	if err != nil {
		v.addf("chains[%d].remoteChains[%d]: failed to read proxy mechanism: %v", chainSelector, remoteSelector, err)
		return
	}

	expectedMechanism, err := mechanismUint8(cfg.ExpectedMechanism)
	if err != nil {
		v.addf("chains[%d].remoteChains[%d]: %v", chainSelector, remoteSelector, err)
		return
	}
	if mechanism != expectedMechanism {
		v.addf("chains[%d].remoteChains[%d]: mechanism mismatch, expected %s got %d", chainSelector, remoteSelector, cfg.ExpectedMechanism, mechanism)
	}

	if supported, err := proxy.IsSupportedChain(callOpts, remoteSelector); err != nil {
		v.addf("chains[%d].remoteChains[%d]: failed to check proxy supported chain: %v", chainSelector, remoteSelector, err)
	} else if !supported {
		v.addf("chains[%d].remoteChains[%d]: proxy does not report the chain as supported", chainSelector, remoteSelector)
	}

	backingPool := backingPoolForMechanism(pools, cfg.ExpectedMechanism)
	if backingPool == (common.Address{}) {
		v.addf("chains[%d].remoteChains[%d]: proxy backing pool for mechanism %s is not configured", chainSelector, remoteSelector, cfg.ExpectedMechanism)
	}

	if cfg.ExpectedRemoteToken != "" {
		remoteTokenBytes, err := proxy.GetRemoteToken(callOpts, remoteSelector)
		if err != nil {
			v.addf("chains[%d].remoteChains[%d]: failed to read remote token from proxy: %v", chainSelector, remoteSelector, err)
		} else {
			remoteToken, decodeErr := decodeEVMAddressBytes(remoteTokenBytes)
			if decodeErr != nil {
				v.addf("chains[%d].remoteChains[%d]: failed to decode remote token bytes: %v", chainSelector, remoteSelector, decodeErr)
			} else if remoteToken != common.HexToAddress(cfg.ExpectedRemoteToken) {
				v.addf("chains[%d].remoteChains[%d]: remote token mismatch, expected %s got %s", chainSelector, remoteSelector, cfg.ExpectedRemoteToken, remoteToken.Hex())
			}
		}
	}

	if len(cfg.ExpectedRemotePools) > 0 {
		remotePoolsBytes, err := proxy.GetRemotePools(callOpts, remoteSelector)
		if err != nil {
			v.addf("chains[%d].remoteChains[%d]: failed to read remote pools from proxy: %v", chainSelector, remoteSelector, err)
		} else {
			actualPools, decodeErr := decodeEVMAddressList(remotePoolsBytes)
			if decodeErr != nil {
				v.addf("chains[%d].remoteChains[%d]: failed to decode remote pool bytes: %v", chainSelector, remoteSelector, decodeErr)
			} else {
				for _, expectedPool := range cfg.ExpectedRemotePools {
					if !containsAddress(actualPools, common.HexToAddress(expectedPool)) {
						v.addf("chains[%d].remoteChains[%d]: expected remote pool %s missing from proxy config", chainSelector, remoteSelector, expectedPool)
					}
				}
			}
		}
	}
}

func validateHomeChainLiquidityState(
	e deployment.Environment,
	chain cldf_evm.Chain,
	cfg ValidateUSDCCLDRolloutHomeChainLiquidityConfig,
	v *validationCollector,
) {
	callOpts := &bind.CallOpts{Context: e.GetContext()}
	usdcToken := common.HexToAddress(cfg.USDCToken)
	siloedPoolAddr := common.HexToAddress(cfg.SiloedUSDCTokenPool)
	hybridPoolAddr := common.HexToAddress(cfg.HybridLockReleaseTokenPool)

	siloedPool, err := siloedusdcpoolbinding.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	if err != nil {
		v.addf("homeChainLiquidity: failed to bind siloed pool %s: %v", cfg.SiloedUSDCTokenPool, err)
		return
	}
	siloedPoolToken, err := siloedPool.GetToken(callOpts)
	if err != nil {
		v.addf("homeChainLiquidity: failed to read siloed pool token from %s: %v", cfg.SiloedUSDCTokenPool, err)
		return
	}
	if siloedPoolToken != usdcToken {
		v.addf("homeChainLiquidity: siloed pool token mismatch, expected %s got %s", usdcToken.Hex(), siloedPoolToken.Hex())
	}

	lockBoxConfigReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		siloedpoolops.GetAllLockBoxConfigs,
		chain,
		cldf_evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       siloedPoolAddr,
			Args:          struct{}{},
		},
	)
	if err != nil {
		v.addf("homeChainLiquidity: failed to read lockbox configs from %s: %v", cfg.SiloedUSDCTokenPool, err)
		return
	}

	lockBoxes := make(map[uint64]common.Address, len(lockBoxConfigReport.Output))
	seenLockBoxes := make(map[common.Address]uint64, len(lockBoxConfigReport.Output))
	for _, lockBoxCfg := range lockBoxConfigReport.Output {
		if prevSelector, exists := seenLockBoxes[lockBoxCfg.LockBox]; exists && prevSelector != lockBoxCfg.RemoteChainSelector {
			v.addf(
				"homeChainLiquidity: lockbox %s is assigned to multiple remote selectors (%d and %d)",
				lockBoxCfg.LockBox.Hex(),
				prevSelector,
				lockBoxCfg.RemoteChainSelector,
			)
		}
		lockBoxes[lockBoxCfg.RemoteChainSelector] = lockBoxCfg.LockBox
		seenLockBoxes[lockBoxCfg.LockBox] = lockBoxCfg.RemoteChainSelector
	}

	for remoteSelector, laneCheck := range cfg.Checks {
		lockBoxAddr, ok := lockBoxes[remoteSelector]
		if !ok {
			v.addf("homeChainLiquidity.checks[%d]: no lockbox configured on siloed pool %s", remoteSelector, cfg.SiloedUSDCTokenPool)
			continue
		}

		lockReleaseReport, err := cldf_ops.ExecuteOperation(
			e.OperationsBundle,
			hybridpoolops.ShouldUseLockRelease,
			chain,
			cldf_evm_contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       hybridPoolAddr,
				Args:          remoteSelector,
			},
		)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to read hybrid lock-release mode: %v", remoteSelector, err)
			continue
		}
		if !lockReleaseReport.Output {
			v.addf("homeChainLiquidity.checks[%d]: hybrid pool %s is not configured for lock-release", remoteSelector, cfg.HybridLockReleaseTokenPool)
		}

		lockBox, err := erc20lockboxbinding.NewERC20LockBox(lockBoxAddr, chain.Client)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to bind lockbox %s: %v", remoteSelector, lockBoxAddr.Hex(), err)
			continue
		}
		lockBoxToken, err := lockBox.GetToken(callOpts)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to read lockbox token: %v", remoteSelector, err)
			continue
		}
		if lockBoxToken != usdcToken {
			v.addf("homeChainLiquidity.checks[%d]: lockbox token mismatch, expected %s got %s", remoteSelector, usdcToken.Hex(), lockBoxToken.Hex())
		}

		lockBoxAuthorizedCallers, err := lockBox.GetAllAuthorizedCallers(callOpts)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to read lockbox authorized callers: %v", remoteSelector, err)
			continue
		}
		if !containsAddress(lockBoxAuthorizedCallers, siloedPoolAddr) {
			v.addf("homeChainLiquidity.checks[%d]: siloed pool %s is not authorized on lockbox %s", remoteSelector, siloedPoolAddr.Hex(), lockBoxAddr.Hex())
		}
		if cfg.ExpectedTimelockAddress != "" && !containsAddress(lockBoxAuthorizedCallers, common.HexToAddress(cfg.ExpectedTimelockAddress)) {
			v.addf("homeChainLiquidity.checks[%d]: timelock %s is not authorized on lockbox %s", remoteSelector, cfg.ExpectedTimelockAddress, lockBoxAddr.Hex())
		}

		if laneCheck.ExpectedLiquidityProvider != "" {
			liquidityProviderReport, err := cldf_ops.ExecuteOperation(
				e.OperationsBundle,
				hybridpoolops.GetLiquidityProvider,
				chain,
				cldf_evm_contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       hybridPoolAddr,
					Args:          remoteSelector,
				},
			)
			if err != nil {
				v.addf("homeChainLiquidity.checks[%d]: failed to read hybrid pool liquidity provider: %v", remoteSelector, err)
			} else if liquidityProviderReport.Output != common.HexToAddress(laneCheck.ExpectedLiquidityProvider) {
				v.addf(
					"homeChainLiquidity.checks[%d]: liquidity provider mismatch, expected %s got %s",
					remoteSelector,
					laneCheck.ExpectedLiquidityProvider,
					liquidityProviderReport.Output.Hex(),
				)
			}
			if !containsAddress(lockBoxAuthorizedCallers, common.HexToAddress(laneCheck.ExpectedLiquidityProvider)) {
				v.addf("homeChainLiquidity.checks[%d]: liquidity provider %s is not authorized on lockbox %s", remoteSelector, laneCheck.ExpectedLiquidityProvider, lockBoxAddr.Hex())
			}
		}

		lockedTokensReport, err := cldf_ops.ExecuteOperation(
			e.OperationsBundle,
			hybridpoolops.GetLockedTokensForChain,
			chain,
			cldf_evm_contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       hybridPoolAddr,
				Args:          remoteSelector,
			},
		)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to read hybrid pool locked tokens: %v", remoteSelector, err)
			continue
		}

		lockBoxBalanceReport, err := cldf_ops.ExecuteOperation(
			e.OperationsBundle,
			erc20ops.BalanceOf,
			chain,
			cldf_evm_contract.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       usdcToken,
				Args:          lockBoxAddr,
			},
		)
		if err != nil {
			v.addf("homeChainLiquidity.checks[%d]: failed to read lockbox USDC balance: %v", remoteSelector, err)
			continue
		}

		expectedHybridLocked, hybridErr := deriveExpectedHybridLocked(laneCheck)
		validateDerivedLiquidityValue(
			v,
			fmt.Sprintf("homeChainLiquidity.checks[%d].ExpectedHybridLocked", remoteSelector),
			lockedTokensReport.Output,
			expectedHybridLocked,
			hybridErr,
		)

		expectedLockBoxBalance, lockBoxErr := deriveExpectedLockBoxBalance(laneCheck)
		validateDerivedLiquidityValue(
			v,
			fmt.Sprintf("homeChainLiquidity.checks[%d].ExpectedLockBoxBalance", remoteSelector),
			lockBoxBalanceReport.Output,
			expectedLockBoxBalance,
			lockBoxErr,
		)
	}
}

func validateExpectedProxyPools(
	chainSelector uint64,
	expected *ValidateUSDCCLDRolloutProxyPoolAddresses,
	actual usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	v *validationCollector,
) {
	if expected == nil {
		return
	}
	if expected.CCTPV1Pool != "" && actual.CctpV1Pool != common.HexToAddress(expected.CCTPV1Pool) {
		v.addf("chains[%d]: proxy CCTP V1 pool mismatch, expected %s got %s", chainSelector, expected.CCTPV1Pool, actual.CctpV1Pool.Hex())
	}
	if expected.CCTPV2Pool != "" && actual.CctpV2Pool != common.HexToAddress(expected.CCTPV2Pool) {
		v.addf("chains[%d]: proxy CCTP V2 pool mismatch, expected %s got %s", chainSelector, expected.CCTPV2Pool, actual.CctpV2Pool.Hex())
	}
	if expected.CCTPV2PoolWithCCV != "" && actual.CctpV2PoolWithCCV != common.HexToAddress(expected.CCTPV2PoolWithCCV) {
		v.addf("chains[%d]: proxy CCTP V2 with CCV pool mismatch, expected %s got %s", chainSelector, expected.CCTPV2PoolWithCCV, actual.CctpV2PoolWithCCV.Hex())
	}
	if expected.SiloedLockReleasePool != "" && actual.SiloedLockReleasePool != common.HexToAddress(expected.SiloedLockReleasePool) {
		v.addf("chains[%d]: proxy siloed lock-release pool mismatch, expected %s got %s", chainSelector, expected.SiloedLockReleasePool, actual.SiloedLockReleasePool.Hex())
	}
}

func validateDerivedLiquidityValue(v *validationCollector, field string, actual *big.Int, expected *big.Int, err error) {
	if err != nil {
		v.addf("%s: %v", field, err)
		return
	}
	if expected == nil {
		return
	}
	if actual.Cmp(expected) != 0 {
		v.addf("%s mismatch, expected %s got %s", field, expected.String(), actual.String())
	}
}

func deriveExpectedHybridLocked(check ValidateUSDCCLDRolloutLiquidityLaneCheck) (*big.Int, error) {
	if check.ExpectedHybridLocked != "" {
		return parseBigInt(check.ExpectedHybridLocked)
	}
	if check.PreHybridLocked == "" || check.ExpectedWithdrawAmount == "" {
		return nil, nil
	}
	pre, err := parseBigInt(check.PreHybridLocked)
	if err != nil {
		return nil, fmt.Errorf("invalid PreHybridLocked %q: %w", check.PreHybridLocked, err)
	}
	withdraw, err := parseBigInt(check.ExpectedWithdrawAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid ExpectedWithdrawAmount %q: %w", check.ExpectedWithdrawAmount, err)
	}
	if pre.Cmp(withdraw) < 0 {
		return nil, fmt.Errorf("pre hybrid locked amount %s is smaller than expected withdraw amount %s", pre.String(), withdraw.String())
	}
	return new(big.Int).Sub(pre, withdraw), nil
}

func deriveExpectedLockBoxBalance(check ValidateUSDCCLDRolloutLiquidityLaneCheck) (*big.Int, error) {
	if check.ExpectedLockBoxBalance != "" {
		return parseBigInt(check.ExpectedLockBoxBalance)
	}
	if check.PreLockBoxBalance == "" || check.ExpectedWithdrawAmount == "" {
		return nil, nil
	}
	pre, err := parseBigInt(check.PreLockBoxBalance)
	if err != nil {
		return nil, fmt.Errorf("invalid PreLockBoxBalance %q: %w", check.PreLockBoxBalance, err)
	}
	withdraw, err := parseBigInt(check.ExpectedWithdrawAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid ExpectedWithdrawAmount %q: %w", check.ExpectedWithdrawAmount, err)
	}
	return new(big.Int).Add(pre, withdraw), nil
}

func verifyProxyPools(v *validationCollector, field string, pools *ValidateUSDCCLDRolloutProxyPoolAddresses) {
	if pools == nil {
		return
	}
	verifyHexAddress(v, field+".CCTPV1Pool", pools.CCTPV1Pool, false)
	verifyHexAddress(v, field+".CCTPV2Pool", pools.CCTPV2Pool, false)
	verifyHexAddress(v, field+".CCTPV2PoolWithCCV", pools.CCTPV2PoolWithCCV, false)
	verifyHexAddress(v, field+".SiloedLockReleasePool", pools.SiloedLockReleasePool, false)
}

func verifyEVMChainSelector(v *validationCollector, selector uint64, field string) {
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		v.addf("%s: invalid chain selector %d: %v", field, selector, err)
		return
	}
	if family != chain_selectors.FamilyEVM {
		v.addf("%s: chain selector %d must be EVM, got family %s", field, selector, family)
	}
}

func verifyHexAddress(v *validationCollector, field, value string, required bool) {
	if value == "" {
		if required {
			v.addf("%s: value is required", field)
		}
		return
	}
	if !common.IsHexAddress(value) {
		v.addf("%s: invalid address %q", field, value)
	}
}

func verifyOptionalBigInt(v *validationCollector, field, value string, mustBePositive bool) {
	if value == "" {
		return
	}
	n, err := parseBigInt(value)
	if err != nil {
		v.addf("%s: %v", field, err)
		return
	}
	if mustBePositive && n.Sign() <= 0 {
		v.addf("%s: must be greater than zero", field)
	}
	if !mustBePositive && n.Sign() < 0 {
		v.addf("%s: must be zero or greater", field)
	}
}

func parseBigInt(value string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(strings.TrimSpace(value), 10)
	if !ok {
		return nil, fmt.Errorf("invalid decimal integer %q", value)
	}
	return n, nil
}

func mechanismUint8(mechanism string) (uint8, error) {
	switch mechanism {
	case validateUSDCCLDMechanismCCTPV1:
		return 1, nil
	case validateUSDCCLDMechanismCCTPV2:
		return 2, nil
	case validateUSDCCLDMechanismLockRelease:
		return 3, nil
	case validateUSDCCLDMechanismCCTPV2WithCCV:
		return 4, nil
	default:
		return 0, fmt.Errorf("unsupported mechanism %q", mechanism)
	}
}

func isSupportedMechanism(mechanism string) bool {
	_, err := mechanismUint8(mechanism)
	return err == nil
}

func backingPoolForMechanism(pools usdcproxybinding.USDCTokenPoolProxyPoolAddresses, mechanism string) common.Address {
	switch mechanism {
	case validateUSDCCLDMechanismCCTPV1:
		return pools.CctpV1Pool
	case validateUSDCCLDMechanismCCTPV2:
		return pools.CctpV2Pool
	case validateUSDCCLDMechanismLockRelease:
		return pools.SiloedLockReleasePool
	case validateUSDCCLDMechanismCCTPV2WithCCV:
		return pools.CctpV2PoolWithCCV
	default:
		return common.Address{}
	}
}

func decodeEVMAddressBytes(data []byte) (common.Address, error) {
	switch len(data) {
	case common.AddressLength:
		return common.BytesToAddress(data), nil
	case 32:
		return common.BytesToAddress(data[32-common.AddressLength:]), nil
	default:
		if len(data) > common.AddressLength {
			return common.BytesToAddress(data[len(data)-common.AddressLength:]), nil
		}
		return common.Address{}, fmt.Errorf("expected 20-byte or 32-byte address encoding, got %d bytes", len(data))
	}
}

func decodeEVMAddressList(values [][]byte) ([]common.Address, error) {
	out := make([]common.Address, 0, len(values))
	for _, value := range values {
		addr, err := decodeEVMAddressBytes(value)
		if err != nil {
			return nil, err
		}
		out = append(out, addr)
	}
	return out, nil
}

func containsAddress(values []common.Address, target common.Address) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

type validationCollector struct {
	errs []string
}

func (v *validationCollector) addf(format string, args ...any) {
	v.errs = append(v.errs, fmt.Sprintf(format, args...))
}

func (v *validationCollector) err() error {
	if len(v.errs) == 0 {
		return nil
	}
	return fmt.Errorf("post-validation failed:\n- %s", strings.Join(v.errs, "\n- "))
}
