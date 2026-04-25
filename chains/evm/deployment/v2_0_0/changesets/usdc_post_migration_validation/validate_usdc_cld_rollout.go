package usdc_post_migration_validation

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	cldf_evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokenadminops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	burnmintpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_mint_token_pool"
	hybridpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/hybrid_lock_release_usdc_token_pool"
	usdcpoolv165ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool"
	usdcpoolcctpv2ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool_cctp_v2"
	cctpthroughccvpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	erc20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20"
	siloedpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	tokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	usdcproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/usdc_token_pool_proxy"
	erc20lockboxbinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/erc20_lock_box"
	siloedusdcpoolbinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	usdcproxybinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	validateUSDCCLDMechanismCCTPV1        = "CCTP_V1"
	validateUSDCCLDMechanismCCTPV2        = "CCTP_V2"
	validateUSDCCLDMechanismLockRelease   = "LOCK_RELEASE"
	validateUSDCCLDMechanismCCTPV2WithCCV = "CCTP_V2_WITH_CCV"

	validateUSDCCLDPoolKindLegacyCanonicalUSDCTokenPool = "LEGACY_CANONICAL_USDC_TOKEN_POOL"
	validateUSDCCLDPoolKindLegacyBurnMintTokenPool      = "LEGACY_NON_CANONICAL_BURN_MINT_TOKEN_POOL"
	validateUSDCCLDPoolKindLegacyHybridLockReleasePool  = "LEGACY_HYBRID_LOCK_RELEASE_USDC_TOKEN_POOL"
	validateUSDCCLDPoolKindUSDCTokenPoolProxy           = "USDC_TOKEN_POOL_PROXY"
	validateUSDCCLDPoolKindCCTPV1Pool                   = "CCTP_V1_POOL"
	validateUSDCCLDPoolKindCCTPV2Pool                   = "CCTP_V2_POOL"
	validateUSDCCLDPoolKindCCTPV2PoolWithCCV            = "CCTP_V2_POOL_WITH_CCV"
	validateUSDCCLDPoolKindSiloedLockReleasePool        = "SILOED_LOCK_RELEASE_POOL"
)

var (
	validateUSDCCLDVersionLegacyCanonicalUSDCTokenPool = semver.MustParse("1.6.2")
	validateUSDCCLDVersionLegacyBurnMintTokenPool      = semver.MustParse("1.5.1")
	validateUSDCCLDVersionLegacyHybridLockReleasePool  = semver.MustParse("1.6.2")
	validateUSDCCLDVersionTokenAdminRegistry           = semver.MustParse("1.5.0")
)

type ValidateUSDCCLDRolloutConfig struct {
	// Chains validates per-chain registry state, proxy wiring, and remote lane configuration.
	Chains map[uint64]ValidateUSDCCLDRolloutChainConfig
	// HomeChainLiquidity validates ETH home-chain liquidity migration into siloed lockboxes.
	HomeChainLiquidity *ValidateUSDCCLDRolloutHomeChainLiquidityConfig
}

type ValidateUSDCCLDRolloutChainConfig struct {
	USDCToken string
	// TokenAdminRegistryRef identifies the TokenAdminRegistry instance to read on this chain.
	// For this rollout it should always be {type: "TokenAdminRegistry", version: "1.5.0"}.
	TokenAdminRegistryRef cldf_datastore.AddressRef
	// ExpectedTokenPool, ExpectedTokenPoolRef, and ExpectedTokenPoolKind are alternate ways to express
	// the expected TokenAdminRegistry "active pool" on this chain for the current rollout phase.
	ExpectedTokenPool     string
	ExpectedTokenPoolRef  cldf_datastore.AddressRef
	ExpectedTokenPoolKind *ValidateUSDCCLDRolloutPoolKindConfig
	// USDCTokenPoolProxy and USDCTokenPoolProxyRef identify the proxy to inspect (if proxy validations
	// should be performed). Leaving both empty skips proxy checks on this chain.
	USDCTokenPoolProxy     string
	USDCTokenPoolProxyRef  cldf_datastore.AddressRef
	ExpectedProxyPools     *ValidateUSDCCLDRolloutProxyPoolAddresses
	ExpectedProxyPoolRefs  *ValidateUSDCCLDRolloutProxyPoolRefs
	ExpectedProxyPoolKinds *ValidateUSDCCLDRolloutProxyPoolKinds
	RemoteChains           map[uint64]ValidateUSDCCLDRolloutRemoteChainConfig
}

type ValidateUSDCCLDRolloutProxyPoolAddresses struct {
	CCTPV1Pool            string
	CCTPV2Pool            string
	CCTPV2PoolWithCCV     string
	SiloedLockReleasePool string
}

type ValidateUSDCCLDRolloutProxyPoolRefs struct {
	CCTPV1Pool            cldf_datastore.AddressRef
	CCTPV2Pool            cldf_datastore.AddressRef
	CCTPV2PoolWithCCV     cldf_datastore.AddressRef
	SiloedLockReleasePool cldf_datastore.AddressRef
}

type ValidateUSDCCLDRolloutProxyPoolKinds struct {
	CCTPV1Pool            *ValidateUSDCCLDRolloutPoolKindConfig
	CCTPV2Pool            *ValidateUSDCCLDRolloutPoolKindConfig
	CCTPV2PoolWithCCV     *ValidateUSDCCLDRolloutPoolKindConfig
	SiloedLockReleasePool *ValidateUSDCCLDRolloutPoolKindConfig
}

type ValidateUSDCCLDRolloutPoolKindConfig struct {
	// Kind identifies a well-known pool family/version used by the CLD rollout validator.
	Kind string
	// Qualifier is used to disambiguate datastore entries when multiple contracts of the same
	// type/version can exist on the same chain. This is especially relevant for legacy
	// BurnMintTokenPool v1.5.1 where multiple refs can be present.
	Qualifier string
}

type ValidateUSDCCLDRolloutRemoteChainConfig struct {
	ExpectedMechanism   string
	ExpectedRemoteToken string
	ExpectedRemotePools []string
	// ExpectedRemotePoolKinds resolves expected remote pool addresses from datastore scoped to the
	// remote chain selector. This avoids hardcoding addresses and prevents cross-chain ref collisions.
	ExpectedRemotePoolKinds  []ValidateUSDCCLDRolloutPoolKindConfig
	LegacyRemotePoolRef      cldf_datastore.AddressRef
	CurrentRemotePoolRef     cldf_datastore.AddressRef
	RequireLegacyRemotePool  bool
	RequireCurrentRemotePool bool
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
	// ExpectedWithdrawAmount is the raw USDC amount (6 decimals) withdrawn from the hybrid pool for
	// this lane during the migration step. It is optional, but when provided must be > 0.
	//
	// It is intentionally uint64 to align with the migration changeset input type.
	ExpectedWithdrawAmount    *uint64
	PreHybridLocked           *uint64
	ExpectedHybridLocked      *uint64
	PreLockBoxBalance         *uint64
	ExpectedLockBoxBalance    *uint64
	ExpectedLiquidityProvider string
}

// ValidateUSDCCLDRollout validates the post-state of the EVM USDC CLD rollout.
// It is read-only and intended to be run after each CLD step to assert pool cutover
// state, proxy routing, remote lane configuration, and home-chain liquidity migration.
//
// Key properties:
//   - Non-destructive: does not submit transactions and does not write to the datastore.
//   - Chain-scoped datastore reads: every ref resolution is explicitly scoped by chain selector, so
//     identical (type, version, qualifier) refs on other chains cannot collide.
//   - Aggregated reporting: this changeset attempts to run all independent validations and reports
//     every misconfiguration it can observe in a single error message (it only skips checks that are
//     strictly blocked by an upstream read/bind failure).
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

	// Step 0: validate input shape. This is intentionally strict because a validator with silently
	// ignored fields can be misleading during a mainnet rollout.
	for chainSelector, chainCfg := range cfg.Chains {
		verifyEVMChainSelector(v, chainSelector, "chains")
		if _, ok := e.BlockChains.EVMChains()[chainSelector]; !ok {
			v.addf("chains[%d]: chain selector is not available in environment", chainSelector)
		}
		verifyHexAddress(v, fmt.Sprintf("chains[%d].USDCToken", chainSelector), chainCfg.USDCToken, true)
		verifyTokenAdminRegistryRef(v, fmt.Sprintf("chains[%d].TokenAdminRegistryRef", chainSelector), chainCfg.TokenAdminRegistryRef)
		verifyHexAddress(v, fmt.Sprintf("chains[%d].ExpectedTokenPool", chainSelector), chainCfg.ExpectedTokenPool, false)
		verifyHexAddress(v, fmt.Sprintf("chains[%d].USDCTokenPoolProxy", chainSelector), chainCfg.USDCTokenPoolProxy, false)
		verifyProxyPools(v, fmt.Sprintf("chains[%d].ExpectedProxyPools", chainSelector), chainCfg.ExpectedProxyPools)
		verifyAddressRef(v, fmt.Sprintf("chains[%d].ExpectedTokenPoolRef", chainSelector), chainCfg.ExpectedTokenPoolRef)
		verifyAddressRef(v, fmt.Sprintf("chains[%d].USDCTokenPoolProxyRef", chainSelector), chainCfg.USDCTokenPoolProxyRef)
		verifyProxyPoolRefs(v, fmt.Sprintf("chains[%d].ExpectedProxyPoolRefs", chainSelector), chainCfg.ExpectedProxyPoolRefs)
		verifyPoolKindConfig(v, fmt.Sprintf("chains[%d].ExpectedTokenPoolKind", chainSelector), chainCfg.ExpectedTokenPoolKind)
		verifyProxyPoolKinds(v, fmt.Sprintf("chains[%d].ExpectedProxyPoolKinds", chainSelector), chainCfg.ExpectedProxyPoolKinds)

		for remoteSelector, remoteCfg := range chainCfg.RemoteChains {
			verifyEVMChainSelector(v, remoteSelector, fmt.Sprintf("chains[%d].RemoteChains", chainSelector))
			if remoteCfg.ExpectedMechanism != "" && !isSupportedMechanism(remoteCfg.ExpectedMechanism) {
				v.addf("chains[%d].RemoteChains[%d].ExpectedMechanism: unsupported mechanism %q", chainSelector, remoteSelector, remoteCfg.ExpectedMechanism)
			}
			verifyHexAddress(v, fmt.Sprintf("chains[%d].RemoteChains[%d].ExpectedRemoteToken", chainSelector, remoteSelector), remoteCfg.ExpectedRemoteToken, false)
			for i, pool := range remoteCfg.ExpectedRemotePools {
				verifyHexAddress(v, fmt.Sprintf("chains[%d].RemoteChains[%d].ExpectedRemotePools[%d]", chainSelector, remoteSelector, i), pool, true)
			}
			for i := range remoteCfg.ExpectedRemotePoolKinds {
				verifyPoolKindConfig(v, fmt.Sprintf("chains[%d].RemoteChains[%d].ExpectedRemotePoolKinds[%d]", chainSelector, remoteSelector, i), &remoteCfg.ExpectedRemotePoolKinds[i])
			}
			verifyAddressRef(v, fmt.Sprintf("chains[%d].RemoteChains[%d].LegacyRemotePoolRef", chainSelector, remoteSelector), remoteCfg.LegacyRemotePoolRef)
			verifyAddressRef(v, fmt.Sprintf("chains[%d].RemoteChains[%d].CurrentRemotePoolRef", chainSelector, remoteSelector), remoteCfg.CurrentRemotePoolRef)
		}
	}

	if cfg.HomeChainLiquidity != nil {
		home := cfg.HomeChainLiquidity
		// Home-chain liquidity migration only applies on Ethereum mainnet/Sepolia in this rollout.
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
			verifyOptionalUint64(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedWithdrawAmount", remoteSelector), laneCheck.ExpectedWithdrawAmount, true)
			verifyOptionalUint64(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].PreHybridLocked", remoteSelector), laneCheck.PreHybridLocked, false)
			verifyOptionalUint64(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedHybridLocked", remoteSelector), laneCheck.ExpectedHybridLocked, false)
			verifyOptionalUint64(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].PreLockBoxBalance", remoteSelector), laneCheck.PreLockBoxBalance, false)
			verifyOptionalUint64(v, fmt.Sprintf("homeChainLiquidity.Checks[%d].ExpectedLockBoxBalance", remoteSelector), laneCheck.ExpectedLockBoxBalance, false)
		}
	}

	return v.err()
}

func applyValidateUSDCCLDRollout(e deployment.Environment, cfg ValidateUSDCCLDRolloutConfig) (deployment.ChangesetOutput, error) {
	v := &validationCollector{}
	defer func() {
		if v.err() == nil {
			e.Logger.Infof(
				"✅ ValidateUSDCCLDRollout: completed successfully (chains=%d homeChainLiquidity=%v)",
				len(cfg.Chains),
				cfg.HomeChainLiquidity != nil,
			)
		}
	}()

	// Step 1: validate per-chain post-state (TokenAdminRegistry, proxy wiring, lane configuration).
	for chainSelector, chainCfg := range cfg.Chains {
		chain, ok := e.BlockChains.EVMChains()[chainSelector]
		if !ok {
			v.addf("chains[%d]: evm chain not available in environment", chainSelector)
			continue
		}
		validateChainState(e, chain, chainCfg, v)
	}

	if cfg.HomeChainLiquidity != nil {
		// Step 2: validate home-chain liquidity migration to siloed lockboxes.
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
	errsBefore := len(v.errs)
	defer func() {
		if len(v.errs) == errsBefore {
			e.Logger.Infof("✅ ValidateUSDCCLDRollout: chain %d: all configured checks for this chain passed", chain.Selector)
		}
	}()

	callOpts := &bind.CallOpts{Context: e.GetContext()}
	usdcToken := common.HexToAddress(cfg.USDCToken)

	// Step 1: resolve TokenAdminRegistry address from datastore (chain-scoped).
	tokenAdminRegistryAddr, err := resolveExpectedAddress(e, chain.Selector, "", cfg.TokenAdminRegistryRef)
	if err != nil {
		v.addf("chains[%d]: failed to resolve TokenAdminRegistry: %v", chain.Selector, err)
	}
	if tokenAdminRegistryAddr == nil {
		v.addf("chains[%d]: TokenAdminRegistryRef must resolve to an address", chain.Selector)
		// Continue running proxy validations even if TAR cannot be read.
	}

	// Step 1a: resolve the expected "active pool" for TokenAdminRegistry on this chain. You can
	// express this expectation as a direct address, as a datastore ref, or as a higher-level pool kind.
	// All resolution paths are chain-scoped using chain.Selector.
	expectedTokenPoolAddr, err := resolveExpectedAddress(e, chain.Selector, cfg.ExpectedTokenPool, cfg.ExpectedTokenPoolRef)
	if err != nil {
		v.addf("chains[%d]: failed to resolve expected token pool: %v", chain.Selector, err)
	}
	if expectedTokenPoolAddr == nil && cfg.ExpectedTokenPoolKind != nil {
		expectedTokenPoolAddr, err = resolvePoolKindAddress(e, chain.Selector, *cfg.ExpectedTokenPoolKind)
		if err != nil {
			v.addf("chains[%d]: failed to resolve expected token pool kind %s: %v", chain.Selector, cfg.ExpectedTokenPoolKind.Kind, err)
		}
	}

	if tokenAdminRegistryAddr != nil {
		tokenConfigReport, err := cldf_ops.ExecuteOperation(
			e.OperationsBundle,
			tokenadminops.GetTokenConfig,
			chain,
			cldf_evm_contract.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       *tokenAdminRegistryAddr,
				Args:          usdcToken,
			},
		)
		if err != nil {
			v.addf("chains[%d]: failed to read token config from token admin registry %s: %v", chain.Selector, tokenAdminRegistryAddr.Hex(), err)
		} else if expectedTokenPoolAddr != nil && tokenConfigReport.Output.TokenPool != *expectedTokenPoolAddr {
			v.addf(
				"chains[%d]: token admin registry active pool mismatch, expected %s got %s",
				chain.Selector,
				expectedTokenPoolAddr.Hex(),
				tokenConfigReport.Output.TokenPool.Hex(),
			)
		} else if expectedTokenPoolAddr != nil {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: chain %d: TokenAdminRegistry active pool matches expected address %s",
				chain.Selector,
				expectedTokenPoolAddr.Hex(),
			)
		}
	}

	proxyAddr, err := resolveExpectedAddress(e, chain.Selector, cfg.USDCTokenPoolProxy, cfg.USDCTokenPoolProxyRef)
	if err != nil {
		v.addf("chains[%d]: failed to resolve USDCTokenPoolProxy: %v", chain.Selector, err)
	}
	if proxyAddr == nil {
		// Proxy checks are optional. Some phases only validate TAR without validating proxy internals.
		return
	}

	// Step 1b: validate the proxy local wiring (token + supported token).
	proxy, err := usdcproxybinding.NewUSDCTokenPoolProxy(*proxyAddr, chain.Client)
	if err != nil {
		v.addf("chains[%d]: failed to bind USDCTokenPoolProxy %s: %v", chain.Selector, proxyAddr.Hex(), err)
		return
	}

	if supported, err := proxy.IsSupportedToken(callOpts, usdcToken); err != nil {
		v.addf("chains[%d]: failed to check proxy supported token on %s: %v", chain.Selector, proxyAddr.Hex(), err)
	} else if !supported {
		v.addf("chains[%d]: proxy %s does not report USDC %s as supported", chain.Selector, proxyAddr.Hex(), usdcToken.Hex())
	} else {
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: chain %d: proxy %s reports USDC %s as a supported token",
			chain.Selector,
			proxyAddr.Hex(),
			usdcToken.Hex(),
		)
	}

	// Step 1c: read proxy backing pool addresses once (GetPools), then:
	// - validate backing pool wiring (optional)
	// - validate every remote lane independently so one bad lane does not hide others
	var pools *usdcproxybinding.USDCTokenPoolProxyPoolAddresses
	actualPools, err := proxy.GetPools(callOpts)
	if err != nil {
		v.addf("chains[%d]: failed to read proxy pools from %s: %v", chain.Selector, proxyAddr.Hex(), err)
	} else {
		pools = &actualPools
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: chain %d: read proxy backing pools from %s (cctpV1=%s cctpV2=%s cctpV2WithCCV=%s siloedLockRelease=%s)",
			chain.Selector,
			proxyAddr.Hex(),
			actualPools.CctpV1Pool.Hex(),
			actualPools.CctpV2Pool.Hex(),
			actualPools.CctpV2PoolWithCCV.Hex(),
			actualPools.SiloedLockReleasePool.Hex(),
		)
		validateExpectedProxyPools(e, chain.Selector, cfg.ExpectedProxyPools, actualPools, v)
		validateExpectedProxyPoolRefs(e, chain.Selector, cfg.ExpectedProxyPoolRefs, actualPools, v)
		validateExpectedProxyPoolKinds(e, chain.Selector, cfg.ExpectedProxyPoolKinds, actualPools, v)
	}

	for remoteSelector, remoteCfg := range cfg.RemoteChains {
		validateRemoteChainState(e, chain, callOpts, chain.Selector, proxy, pools, remoteSelector, remoteCfg, v)
	}
}

func needsRemotePoolValidation(cfg ValidateUSDCCLDRolloutRemoteChainConfig) bool {
	return len(cfg.ExpectedRemotePools) > 0 ||
		len(cfg.ExpectedRemotePoolKinds) > 0 ||
		cfg.RequireLegacyRemotePool ||
		cfg.RequireCurrentRemotePool ||
		!datastore_utils.IsAddressRefEmpty(cfg.LegacyRemotePoolRef) ||
		!datastore_utils.IsAddressRefEmpty(cfg.CurrentRemotePoolRef)
}

func validateRemoteChainState(
	e deployment.Environment,
	chain cldf_evm.Chain,
	callOpts *bind.CallOpts,
	chainSelector uint64,
	proxy *usdcproxybinding.USDCTokenPoolProxy,
	pools *usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	remoteSelector uint64,
	cfg ValidateUSDCCLDRolloutRemoteChainConfig,
	v *validationCollector,
) {
	if cfg.ExpectedMechanism == "" {
		// Remote lane validations are opt-in per remoteSelector.
		return
	}

	errsBefore := len(v.errs)
	defer func() {
		if len(v.errs) == errsBefore {
			e.Logger.Infof(
				"✅ ValidateUSDCCLDRollout: chain %d -> remote %d: lane validation passed (mechanism=%s)",
				chainSelector,
				remoteSelector,
				cfg.ExpectedMechanism,
			)
		}
	}()

	// Step 2a: mechanism must match what CLD configured on the proxy for this remote chain.
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

	// Step 2b: ensure the proxy has a non-zero backing pool address configured for this mechanism.
	// This check is skipped only if GetPools failed earlier.
	if pools != nil {
		backingPool := backingPoolForMechanism(*pools, cfg.ExpectedMechanism)
		if backingPool == (common.Address{}) {
			v.addf("chains[%d].remoteChains[%d]: proxy backing pool for mechanism %s is not configured", chainSelector, remoteSelector, cfg.ExpectedMechanism)
		}
	} else {
		v.addf("chains[%d].remoteChains[%d]: skipped proxy backing-pool validation because GetPools failed", chainSelector, remoteSelector)
	}

	if cfg.ExpectedRemoteToken != "" {
		// Step 2c: ensure the proxy’s remote token entry matches expectation.
		remoteTokenBytes, err := proxy.GetRemoteToken(callOpts, remoteSelector)
		if err != nil {
			v.addf("chains[%d].remoteChains[%d]: failed to read remote token from proxy: %v", chainSelector, remoteSelector, err)
		} else {
			remoteToken, decodeErr := decodeEVMAddressBytes(remoteTokenBytes)
			if decodeErr != nil {
				v.addf("chains[%d].remoteChains[%d]: failed to decode remote token bytes: %v", chainSelector, remoteSelector, decodeErr)
			} else if remoteToken != common.HexToAddress(cfg.ExpectedRemoteToken) {
				v.addf("chains[%d].remoteChains[%d]: remote token mismatch, expected %s got %s", chainSelector, remoteSelector, cfg.ExpectedRemoteToken, remoteToken.Hex())
			} else {
				e.Logger.Infof(
					"ValidateUSDCCLDRollout: chain %d -> remote %d: remote token matches expected %s",
					chainSelector,
					remoteSelector,
					cfg.ExpectedRemoteToken,
				)
			}
		}
	}

	if needsRemotePoolValidation(cfg) {
		remotePoolsBytes, proxyErr := proxy.GetRemotePools(callOpts, remoteSelector)
		if proxyErr != nil {
			backingPool := common.Address{}
			if pools != nil {
				backingPool = backingPoolForMechanism(*pools, cfg.ExpectedMechanism)
			}
			if backingPool == (common.Address{}) {
				v.addf(
					"chains[%d].remoteChains[%d]: failed to read remote pools from proxy: %v (subpool fallback unavailable: backing pool for mechanism %s is unset or proxy GetPools failed)",
					chainSelector,
					remoteSelector,
					proxyErr,
					cfg.ExpectedMechanism,
				)
			} else {
				subpoolReport, subpoolErr := cldf_ops.ExecuteOperation(
					e.OperationsBundle,
					tokenpoolops.GetRemotePools,
					chain,
					cldf_evm_contract.FunctionInput[uint64]{
						ChainSelector: chain.Selector,
						Address:       backingPool,
						Args:          remoteSelector,
					},
				)
				if subpoolErr != nil {
					v.addf(
						"chains[%d].remoteChains[%d]: proxy GetRemotePools failed: %v; backing subpool %s GetRemotePools failed: %v",
						chainSelector,
						remoteSelector,
						proxyErr,
						backingPool.Hex(),
						subpoolErr,
					)
				} else {
					remotePoolsBytes = subpoolReport.Output
					e.Logger.Infof(
						"ValidateUSDCCLDRollout: chain %d -> remote %d: read remote pools from backing subpool %s (proxy GetRemotePools failed: %v)",
						chainSelector,
						remoteSelector,
						backingPool.Hex(),
						proxyErr,
					)
					proxyErr = nil
				}
			}
		}

		if proxyErr == nil {
			actualPools, decodeErr := decodeEVMAddressList(remotePoolsBytes)
			if decodeErr != nil {
				v.addf("chains[%d].remoteChains[%d]: failed to decode remote pool bytes: %v", chainSelector, remoteSelector, decodeErr)
			} else {
				for _, expectedPool := range cfg.ExpectedRemotePools {
					if !containsAddress(actualPools, common.HexToAddress(expectedPool)) {
						v.addf("chains[%d].remoteChains[%d]: expected remote pool %s missing from remote pool list", chainSelector, remoteSelector, expectedPool)
					}
				}
				for _, expectedPoolKind := range cfg.ExpectedRemotePoolKinds {
					expectedPoolAddr, resolveErr := resolvePoolKindAddress(e, remoteSelector, expectedPoolKind)
					if resolveErr != nil {
						v.addf("chains[%d].remoteChains[%d]: failed to resolve remote pool kind %s: %v", chainSelector, remoteSelector, expectedPoolKind.Kind, resolveErr)
						continue
					}
					if !containsAddress(actualPools, *expectedPoolAddr) {
						v.addf("chains[%d].remoteChains[%d]: expected remote pool kind %s (%s) missing from remote pool list", chainSelector, remoteSelector, expectedPoolKind.Kind, expectedPoolAddr.Hex())
					}
				}
				validateExpectedRemotePoolRef(e, chainSelector, remoteSelector, "legacy", cfg.RequireLegacyRemotePool, cfg.LegacyRemotePoolRef, actualPools, v)
				validateExpectedRemotePoolRef(e, chainSelector, remoteSelector, "current", cfg.RequireCurrentRemotePool, cfg.CurrentRemotePoolRef, actualPools, v)
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
	errsBefore := len(v.errs)
	defer func() {
		if len(v.errs) == errsBefore {
			e.Logger.Infof(
				"✅ ValidateUSDCCLDRollout: homeChainLiquidity chain %d: all configured liquidity checks passed",
				cfg.ChainSelector,
			)
		}
	}()

	callOpts := &bind.CallOpts{Context: e.GetContext()}
	usdcToken := common.HexToAddress(cfg.USDCToken)
	siloedPoolAddr := common.HexToAddress(cfg.SiloedUSDCTokenPool)
	hybridPoolAddr := common.HexToAddress(cfg.HybridLockReleaseTokenPool)

	// Step 3a: validate siloed pool base wiring and discover lockbox topology for all remotes.
	siloedPool, err := siloedusdcpoolbinding.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	if err != nil {
		v.addf("homeChainLiquidity: failed to bind siloed pool %s: %v", cfg.SiloedUSDCTokenPool, err)
		return
	}
	siloedPoolToken, err := siloedPool.GetToken(callOpts)
	if err != nil {
		v.addf("homeChainLiquidity: failed to read siloed pool token from %s: %v", cfg.SiloedUSDCTokenPool, err)
	} else if siloedPoolToken != usdcToken {
		v.addf("homeChainLiquidity: siloed pool token mismatch, expected %s got %s", usdcToken.Hex(), siloedPoolToken.Hex())
	} else {
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: homeChainLiquidity: siloed pool %s token matches USDC %s",
			cfg.SiloedUSDCTokenPool,
			usdcToken.Hex(),
		)
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
	}

	lockBoxes := make(map[uint64]common.Address)
	if err == nil {
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: homeChainLiquidity: loaded %d lockbox config(s) from siloed pool %s",
			len(lockBoxConfigReport.Output),
			cfg.SiloedUSDCTokenPool,
		)
		lockBoxes = make(map[uint64]common.Address, len(lockBoxConfigReport.Output))
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
	}

	for remoteSelector, laneCheck := range cfg.Checks {
		errsBeforeLane := len(v.errs)
		// Step 3b: validate each migrated lane. This is designed to be run twice during the rollout:
		// - first with a tiny (1 USDC) withdraw amount as a canary
		// - then with a larger percentage (e.g. 60%) and later with the remaining balance
		//
		// We check:
		// - hybrid pool is configured for lock-release for the remote selector
		// - a lockbox exists on the siloed pool for the remote selector
		// - lockbox authorizations include siloed pool and expected timelock/LP
		// - liquidity conservation: hybrid locked amount decreases, lockbox balance increases
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
		} else {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: homeChainLiquidity.checks[%d]: hybrid pool is in lock-release mode for remote chain",
				remoteSelector,
			)
		}

		lockBoxAddr, ok := lockBoxes[remoteSelector]
		if !ok {
			v.addf("homeChainLiquidity.checks[%d]: no lockbox configured on siloed pool %s", remoteSelector, cfg.SiloedUSDCTokenPool)
		}

		lockBoxAuthorizedCallers := []common.Address(nil)
		if ok {
			lockBox, lockBoxErr := erc20lockboxbinding.NewERC20LockBox(lockBoxAddr, chain.Client)
			if lockBoxErr != nil {
				v.addf("homeChainLiquidity.checks[%d]: failed to bind lockbox %s: %v", remoteSelector, lockBoxAddr.Hex(), lockBoxErr)
			} else {
				lockBoxToken, lockBoxTokenErr := lockBox.GetToken(callOpts)
				if lockBoxTokenErr != nil {
					v.addf("homeChainLiquidity.checks[%d]: failed to read lockbox token: %v", remoteSelector, lockBoxTokenErr)
				} else if lockBoxToken != usdcToken {
					v.addf("homeChainLiquidity.checks[%d]: lockbox token mismatch, expected %s got %s", remoteSelector, usdcToken.Hex(), lockBoxToken.Hex())
				}

				lockBoxAuthorizedCallers, lockBoxErr = lockBox.GetAllAuthorizedCallers(callOpts)
				if lockBoxErr != nil {
					v.addf("homeChainLiquidity.checks[%d]: failed to read lockbox authorized callers: %v", remoteSelector, lockBoxErr)
				} else {
					if !containsAddress(lockBoxAuthorizedCallers, siloedPoolAddr) {
						v.addf("homeChainLiquidity.checks[%d]: siloed pool %s is not authorized on lockbox %s", remoteSelector, siloedPoolAddr.Hex(), lockBoxAddr.Hex())
					}
					if cfg.ExpectedTimelockAddress != "" && !containsAddress(lockBoxAuthorizedCallers, common.HexToAddress(cfg.ExpectedTimelockAddress)) {
						v.addf("homeChainLiquidity.checks[%d]: timelock %s is not authorized on lockbox %s", remoteSelector, cfg.ExpectedTimelockAddress, lockBoxAddr.Hex())
					}
				}
			}
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
			if ok && len(lockBoxAuthorizedCallers) > 0 && !containsAddress(lockBoxAuthorizedCallers, common.HexToAddress(laneCheck.ExpectedLiquidityProvider)) {
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

		if !ok {
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
			e,
			v,
			fmt.Sprintf("homeChainLiquidity.checks[%d].ExpectedHybridLocked", remoteSelector),
			lockedTokensReport.Output,
			expectedHybridLocked,
			hybridErr,
		)

		expectedLockBoxBalance, lockBoxErr := deriveExpectedLockBoxBalance(laneCheck)
		validateDerivedLiquidityValue(
			e,
			v,
			fmt.Sprintf("homeChainLiquidity.checks[%d].ExpectedLockBoxBalance", remoteSelector),
			lockBoxBalanceReport.Output,
			expectedLockBoxBalance,
			lockBoxErr,
		)

		if len(v.errs) == errsBeforeLane {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: homeChainLiquidity.checks[%d]: lane liquidity assertions completed for this remote",
				remoteSelector,
			)
		}
	}
}

func validateExpectedProxyPools(
	e deployment.Environment,
	chainSelector uint64,
	expected *ValidateUSDCCLDRolloutProxyPoolAddresses,
	actual usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	v *validationCollector,
) {
	if expected == nil {
		return
	}
	if expected.CCTPV1Pool != "" {
		if actual.CctpV1Pool != common.HexToAddress(expected.CCTPV1Pool) {
			v.addf("chains[%d]: proxy CCTP V1 pool mismatch, expected %s got %s", chainSelector, expected.CCTPV1Pool, actual.CctpV1Pool.Hex())
		} else {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: chain %d: proxy CCTP V1 backing pool matches expected %s",
				chainSelector,
				expected.CCTPV1Pool,
			)
		}
	}
	if expected.CCTPV2Pool != "" {
		if actual.CctpV2Pool != common.HexToAddress(expected.CCTPV2Pool) {
			v.addf("chains[%d]: proxy CCTP V2 pool mismatch, expected %s got %s", chainSelector, expected.CCTPV2Pool, actual.CctpV2Pool.Hex())
		} else {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: chain %d: proxy CCTP V2 backing pool matches expected %s",
				chainSelector,
				expected.CCTPV2Pool,
			)
		}
	}
	if expected.CCTPV2PoolWithCCV != "" {
		if actual.CctpV2PoolWithCCV != common.HexToAddress(expected.CCTPV2PoolWithCCV) {
			v.addf("chains[%d]: proxy CCTP V2 with CCV pool mismatch, expected %s got %s", chainSelector, expected.CCTPV2PoolWithCCV, actual.CctpV2PoolWithCCV.Hex())
		} else {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: chain %d: proxy CCTP V2 with CCV backing pool matches expected %s",
				chainSelector,
				expected.CCTPV2PoolWithCCV,
			)
		}
	}
	if expected.SiloedLockReleasePool != "" {
		if actual.SiloedLockReleasePool != common.HexToAddress(expected.SiloedLockReleasePool) {
			v.addf("chains[%d]: proxy siloed lock-release pool mismatch, expected %s got %s", chainSelector, expected.SiloedLockReleasePool, actual.SiloedLockReleasePool.Hex())
		} else {
			e.Logger.Infof(
				"ValidateUSDCCLDRollout: chain %d: proxy siloed lock-release backing pool matches expected %s",
				chainSelector,
				expected.SiloedLockReleasePool,
			)
		}
	}
}

func validateExpectedProxyPoolRefs(
	e deployment.Environment,
	chainSelector uint64,
	expected *ValidateUSDCCLDRolloutProxyPoolRefs,
	actual usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	v *validationCollector,
) {
	if expected == nil {
		return
	}
	validateResolvedProxyPoolRef(e, chainSelector, "CCTPV1", expected.CCTPV1Pool, actual.CctpV1Pool, v)
	validateResolvedProxyPoolRef(e, chainSelector, "CCTPV2", expected.CCTPV2Pool, actual.CctpV2Pool, v)
	validateResolvedProxyPoolRef(e, chainSelector, "CCTPV2WithCCV", expected.CCTPV2PoolWithCCV, actual.CctpV2PoolWithCCV, v)
	validateResolvedProxyPoolRef(e, chainSelector, "SiloedLockRelease", expected.SiloedLockReleasePool, actual.SiloedLockReleasePool, v)
}

func validateExpectedProxyPoolKinds(
	e deployment.Environment,
	chainSelector uint64,
	expected *ValidateUSDCCLDRolloutProxyPoolKinds,
	actual usdcproxybinding.USDCTokenPoolProxyPoolAddresses,
	v *validationCollector,
) {
	if expected == nil {
		return
	}
	validateResolvedProxyPoolKind(e, chainSelector, "CCTPV1", expected.CCTPV1Pool, actual.CctpV1Pool, v)
	validateResolvedProxyPoolKind(e, chainSelector, "CCTPV2", expected.CCTPV2Pool, actual.CctpV2Pool, v)
	validateResolvedProxyPoolKind(e, chainSelector, "CCTPV2WithCCV", expected.CCTPV2PoolWithCCV, actual.CctpV2PoolWithCCV, v)
	validateResolvedProxyPoolKind(e, chainSelector, "SiloedLockRelease", expected.SiloedLockReleasePool, actual.SiloedLockReleasePool, v)
}

func validateResolvedProxyPoolRef(
	e deployment.Environment,
	chainSelector uint64,
	label string,
	ref cldf_datastore.AddressRef,
	actual common.Address,
	v *validationCollector,
) {
	if datastore_utils.IsAddressRefEmpty(ref) {
		return
	}
	resolvedRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
	if err != nil {
		v.addf("chains[%d]: failed to resolve %s proxy pool ref %s: %v", chainSelector, label, datastore_utils.SprintRef(ref), err)
		return
	}
	expected := common.HexToAddress(resolvedRef.Address)
	if actual != expected {
		v.addf("chains[%d]: proxy %s pool mismatch, expected %s got %s", chainSelector, label, expected.Hex(), actual.Hex())
	} else {
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: chain %d: proxy %s backing pool matches datastore ref (%s)",
			chainSelector,
			label,
			expected.Hex(),
		)
	}
}

func validateResolvedProxyPoolKind(
	e deployment.Environment,
	chainSelector uint64,
	label string,
	spec *ValidateUSDCCLDRolloutPoolKindConfig,
	actual common.Address,
	v *validationCollector,
) {
	if spec == nil {
		return
	}
	expected, err := resolvePoolKindAddress(e, chainSelector, *spec)
	if err != nil {
		v.addf("chains[%d]: failed to resolve %s proxy pool kind %s: %v", chainSelector, label, spec.Kind, err)
		return
	}
	if actual != *expected {
		v.addf("chains[%d]: proxy %s pool mismatch, expected %s got %s", chainSelector, label, expected.Hex(), actual.Hex())
	} else {
		e.Logger.Infof(
			"ValidateUSDCCLDRollout: chain %d: proxy %s backing pool matches resolved pool kind %s",
			chainSelector,
			label,
			spec.Kind,
		)
	}
}

func validateExpectedRemotePoolRef(
	e deployment.Environment,
	chainSelector uint64,
	remoteSelector uint64,
	label string,
	required bool,
	ref cldf_datastore.AddressRef,
	actualPools []common.Address,
	v *validationCollector,
) {
	if datastore_utils.IsAddressRefEmpty(ref) {
		if required {
			v.addf("chains[%d].remoteChains[%d]: %s remote pool ref is required", chainSelector, remoteSelector, label)
		}
		return
	}
	// Important: resolve against remoteSelector so we only search the remote chain's datastore scope.
	// This prevents accidentally selecting a same-type/version ref from the local chain.
	resolvedRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, remoteSelector, datastore_utils.FullRef)
	if err != nil {
		v.addf("chains[%d].remoteChains[%d]: failed to resolve %s remote pool ref %s: %v", chainSelector, remoteSelector, label, datastore_utils.SprintRef(ref), err)
		return
	}
	expected := common.HexToAddress(resolvedRef.Address)
	if !containsAddress(actualPools, expected) {
		v.addf("chains[%d].remoteChains[%d]: %s remote pool %s missing from remote pool list", chainSelector, remoteSelector, label, expected.Hex())
	}
}

func validateDerivedLiquidityValue(e deployment.Environment, v *validationCollector, field string, actual *big.Int, expected *big.Int, err error) {
	if err != nil {
		v.addf("%s: %v", field, err)
		return
	}
	if expected == nil {
		return
	}
	if actual.Cmp(expected) != 0 {
		v.addf("%s mismatch, expected %s got %s", field, expected.String(), actual.String())
	} else {
		e.Logger.Infof("ValidateUSDCCLDRollout: %s matches expected %s", field, expected.String())
	}
}

func deriveExpectedHybridLocked(check ValidateUSDCCLDRolloutLiquidityLaneCheck) (*big.Int, error) {
	if check.ExpectedHybridLocked != nil {
		return new(big.Int).SetUint64(*check.ExpectedHybridLocked), nil
	}
	if check.PreHybridLocked == nil || check.ExpectedWithdrawAmount == nil {
		return nil, nil
	}
	pre := new(big.Int).SetUint64(*check.PreHybridLocked)
	withdraw := new(big.Int).SetUint64(*check.ExpectedWithdrawAmount)
	if pre.Cmp(withdraw) < 0 {
		return nil, fmt.Errorf("pre hybrid locked amount %s is smaller than expected withdraw amount %s", pre.String(), withdraw.String())
	}
	return new(big.Int).Sub(pre, withdraw), nil
}

func deriveExpectedLockBoxBalance(check ValidateUSDCCLDRolloutLiquidityLaneCheck) (*big.Int, error) {
	if check.ExpectedLockBoxBalance != nil {
		return new(big.Int).SetUint64(*check.ExpectedLockBoxBalance), nil
	}
	if check.PreLockBoxBalance == nil || check.ExpectedWithdrawAmount == nil {
		return nil, nil
	}
	pre := new(big.Int).SetUint64(*check.PreLockBoxBalance)
	withdraw := new(big.Int).SetUint64(*check.ExpectedWithdrawAmount)
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

func verifyProxyPoolRefs(v *validationCollector, field string, pools *ValidateUSDCCLDRolloutProxyPoolRefs) {
	if pools == nil {
		return
	}
	verifyAddressRef(v, field+".CCTPV1Pool", pools.CCTPV1Pool)
	verifyAddressRef(v, field+".CCTPV2Pool", pools.CCTPV2Pool)
	verifyAddressRef(v, field+".CCTPV2PoolWithCCV", pools.CCTPV2PoolWithCCV)
	verifyAddressRef(v, field+".SiloedLockReleasePool", pools.SiloedLockReleasePool)
}

func verifyProxyPoolKinds(v *validationCollector, field string, pools *ValidateUSDCCLDRolloutProxyPoolKinds) {
	if pools == nil {
		return
	}
	verifyPoolKindConfig(v, field+".CCTPV1Pool", pools.CCTPV1Pool)
	verifyPoolKindConfig(v, field+".CCTPV2Pool", pools.CCTPV2Pool)
	verifyPoolKindConfig(v, field+".CCTPV2PoolWithCCV", pools.CCTPV2PoolWithCCV)
	verifyPoolKindConfig(v, field+".SiloedLockReleasePool", pools.SiloedLockReleasePool)
}

func verifyPoolKindConfig(v *validationCollector, field string, spec *ValidateUSDCCLDRolloutPoolKindConfig) {
	if spec == nil {
		return
	}
	if !isSupportedPoolKind(spec.Kind) {
		v.addf("%s.Kind: unsupported pool kind %q", field, spec.Kind)
	}
}

func verifyAddressRef(v *validationCollector, field string, ref cldf_datastore.AddressRef) {
	if datastore_utils.IsAddressRefEmpty(ref) {
		return
	}
	if ref.Address != "" && !common.IsHexAddress(ref.Address) {
		v.addf("%s.Address: invalid address %q", field, ref.Address)
	}
}

func verifyTokenAdminRegistryRef(v *validationCollector, field string, ref cldf_datastore.AddressRef) {
	if datastore_utils.IsAddressRefEmpty(ref) {
		v.addf("%s: value is required", field)
		return
	}
	verifyAddressRef(v, field, ref)

	// In this rollout TAR is fixed to TokenAdminRegistry v1.5.0.
	if ref.Type != cldf_datastore.ContractType("TokenAdminRegistry") {
		v.addf("%s.Type: expected %q, got %q", field, "TokenAdminRegistry", ref.Type)
	}
	if ref.Version == nil || !ref.Version.Equal(validateUSDCCLDVersionTokenAdminRegistry) {
		if ref.Version == nil {
			v.addf("%s.Version: expected %q, got <nil>", field, "1.5.0")
		} else {
			v.addf("%s.Version: expected %q, got %q", field, "1.5.0", ref.Version.String())
		}
	}
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

func verifyOptionalUint64(v *validationCollector, field string, value *uint64, mustBePositive bool) {
	if value == nil {
		return
	}
	if mustBePositive && *value == 0 {
		v.addf("%s: must be greater than zero", field)
	}
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

func resolveExpectedAddress(
	e deployment.Environment,
	chainSelector uint64,
	directAddress string,
	ref cldf_datastore.AddressRef,
) (*common.Address, error) {
	if directAddress != "" {
		addr := common.HexToAddress(directAddress)
		return &addr, nil
	}
	if datastore_utils.IsAddressRefEmpty(ref) {
		return nil, nil
	}
	// FindAndFormatRef forcibly scopes the search by overriding ref.ChainSelector with chainSelector.
	// This keeps lookups deterministic even if the global datastore contains similar refs on other chains.
	resolvedRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, err
	}
	addr := common.HexToAddress(resolvedRef.Address)
	return &addr, nil
}

func isSupportedPoolKind(kind string) bool {
	_, err := poolKindRef(kind, "")
	return err == nil
}

func resolvePoolKindAddress(
	e deployment.Environment,
	chainSelector uint64,
	spec ValidateUSDCCLDRolloutPoolKindConfig,
) (*common.Address, error) {
	ref, err := poolKindRef(spec.Kind, spec.Qualifier)
	if err != nil {
		return nil, err
	}
	// Pool-kind resolution is also chain-scoped, with Qualifier used only as an additional filter
	// within that chain when multiple refs of the same type/version exist.
	resolvedRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, err
	}
	addr := common.HexToAddress(resolvedRef.Address)
	return &addr, nil
}

func poolKindRef(kind string, qualifier string) (cldf_datastore.AddressRef, error) {
	// These kinds are intentionally pinned to the rollout's expected pool types/versions.
	// If the rollout bumps versions, update these mappings so the validator stays honest.
	switch kind {
	case validateUSDCCLDPoolKindLegacyCanonicalUSDCTokenPool:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType("USDCTokenPool"),
			Version: validateUSDCCLDVersionLegacyCanonicalUSDCTokenPool,
		}, nil
	case validateUSDCCLDPoolKindLegacyBurnMintTokenPool:
		return cldf_datastore.AddressRef{
			Type:      cldf_datastore.ContractType(burnmintpoolops.ContractType),
			Version:   validateUSDCCLDVersionLegacyBurnMintTokenPool,
			Qualifier: qualifier,
		}, nil
	case validateUSDCCLDPoolKindLegacyHybridLockReleasePool:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(hybridpoolops.ContractType),
			Version: validateUSDCCLDVersionLegacyHybridLockReleasePool,
		}, nil
	case validateUSDCCLDPoolKindUSDCTokenPoolProxy:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(usdcproxyops.ContractType),
			Version: usdcproxyops.Version,
		}, nil
	case validateUSDCCLDPoolKindCCTPV1Pool:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(usdcpoolv165ops.ContractType),
			Version: usdcpoolv165ops.Version,
		}, nil
	case validateUSDCCLDPoolKindCCTPV2Pool:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(usdcpoolcctpv2ops.ContractType),
			Version: usdcpoolcctpv2ops.Version,
		}, nil
	case validateUSDCCLDPoolKindCCTPV2PoolWithCCV:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(cctpthroughccvpoolops.ContractType),
			Version: cctpthroughccvpoolops.Version,
		}, nil
	case validateUSDCCLDPoolKindSiloedLockReleasePool:
		return cldf_datastore.AddressRef{
			Type:    cldf_datastore.ContractType(siloedpoolops.ContractType),
			Version: siloedpoolops.Version,
		}, nil
	default:
		return cldf_datastore.AddressRef{}, fmt.Errorf("unsupported pool kind %q", kind)
	}
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
	lines := make([]string, 0, len(v.errs))
	for _, msg := range v.errs {
		prefix := "❌"
		// Some checks explicitly record when a sub-check was skipped due to an upstream failure.
		if strings.Contains(strings.ToLower(msg), "skipped") {
			prefix = "⚠️"
		}
		lines = append(lines, fmt.Sprintf("%s %s", prefix, msg))
	}
	return fmt.Errorf("post-validation failed:\n%s", strings.Join(lines, "\n"))
}
