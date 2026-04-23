package adapters

import (
	"github.com/Masterminds/semver/v3"

	chainsel "github.com/smartcontractkit/chain-selectors"

	cctpthroughccvtokenpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	cctpverifierops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	committeeverifierops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	executorops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	feequoterops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	usdctokenpoolproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/usdc_token_pool_proxy"
	seq1_7 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	versionedverifierresolverops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/verification"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	adapters1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	burnfromminttokenpoolv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/burn_from_mint_token_pool"
	burnminttokenpoolv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/burn_mint_token_pool"
	burnmintwithlockreleaseflagtokenpoolv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/burn_mint_with_lock_release_flag_token_pool"
	burnwithfromminttokenpoolv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/executor"
	v2feequoter "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
	lockreleasetokenpoolv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/mock_receiver_v2"
	offrampv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/offramp"
	onrampv2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func init() {
	v := semver.MustParse("2.0.0")

	// CCIP deployment registrations
	evmAdapter := evmseqV1_6.EVMAdapter{}
	evmFeesAdapterV2_0 := NewFeesAdapter(&evmAdapter)

	deploy.GetRegistry().RegisterDeployer(chainsel.FamilyEVM, v, &evmAdapter)

	fqReg := deploy.GetFQAndRampUpdaterRegistry()
	fqReg.RegisterFeeQuoterUpdater(chainsel.FamilyEVM, v, FeeQuoterUpdater[any]{})
	fqReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), adapters1_6.RampUpdateWithFQ{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), &adapters1_5.ConfigImportAdapter{})
	fqReg.RegisterConfigImporterVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})

	feeReg := fees.GetRegistry()
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, v, evmFeesAdapterV2_0)

	// CCV deployment registrations
	registerContractVerificationMetadata(v)

	laneMigratorReg := deploy.GetLaneMigratorRegistry()
	laneMigratorReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("2.0.0"), &LaneMigrator{})
	laneMigratorReg.RegisterRouterUpdater(chainsel.FamilyEVM, semver.MustParse("1.2.0"), &adapters1_2.RouterUpdater{})

	lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(chainsel.FamilyEVM, v, &ChainFamilyAdapter{})
	ccvadapters.GetChainFamilyRegistry().RegisterChainFamily(chainsel.FamilyEVM, &ChainFamilyAdapter{})
	ccvadapters.GetCommitteeVerifierContractRegistry().Register(chainsel.FamilyEVM, &EVMCommitteeVerifierContractAdapter{})
	ccvadapters.GetExecutorConfigRegistry().Register(chainsel.FamilyEVM, &EVMExecutorConfigAdapter{})
	ccvadapters.GetVerifierJobConfigRegistry().Register(chainsel.FamilyEVM, &EVMVerifierJobConfigAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().Register(chainsel.FamilyEVM, &EVMDeployChainContractsAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().RegisterLaneVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})
	ccvadapters.GetIndexerConfigRegistry().Register(chainsel.FamilyEVM, &EVMIndexerConfigAdapter{})
	ccvadapters.GetAggregatorConfigRegistry().Register(chainsel.FamilyEVM, &EVMAggregatorConfigAdapter{})
	ccvadapters.GetTokenVerifierConfigRegistry().Register(chainsel.FamilyEVM, &EVMTokenVerifierConfigAdapter{})

	tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, v, NewTokenAdapter())
	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilyEVM, v, NewFeeAggregatorAdapter())
}

func registerContractVerificationMetadata(v *semver.Version) {
	verification.RegisterContractMetadata(feequoterops.ContractType, v, v2feequoter.SolidityStandardInput, v2feequoter.FeeQuoterBin, "contracts/FeeQuoter.sol:FeeQuoter")
	verification.RegisterContractMetadata(onrampops.ContractType, v, onrampv2.SolidityStandardInput, onrampv2.OnRampBin, "contracts/onRamp/OnRamp.sol:OnRamp")
	verification.RegisterContractMetadata(offrampops.ContractType, v, offrampv2.SolidityStandardInput, offrampv2.OffRampBin, "contracts/offRamp/OffRamp.sol:OffRamp")
	verification.RegisterContractMetadata(utils.BurnMintTokenPool, v, burnminttokenpoolv2.SolidityStandardInput, burnminttokenpoolv2.BurnMintTokenPoolBin, "contracts/pools/BurnMintTokenPool.sol:BurnMintTokenPool")
	verification.RegisterContractMetadata(utils.BurnWithFromMintTokenPool, v, burnwithfromminttokenpoolv2.SolidityStandardInput, burnwithfromminttokenpoolv2.BurnWithFromMintTokenPoolBin, "contracts/pools/BurnWithFromMintTokenPool.sol:BurnWithFromMintTokenPool")
	verification.RegisterContractMetadata(utils.BurnMintWithLockReleaseFlag, v, burnmintwithlockreleaseflagtokenpoolv2.SolidityStandardInput, burnmintwithlockreleaseflagtokenpoolv2.BurnMintWithLockReleaseFlagTokenPoolBin, "contracts/pools/USDC/BurnMintWithLockReleaseFlagTokenPool.sol:BurnMintWithLockReleaseFlagTokenPool")
	verification.RegisterContractMetadata(utils.BurnFromMintTokenPool, v, burnfromminttokenpoolv2.SolidityStandardInput, burnfromminttokenpoolv2.BurnFromMintTokenPoolBin, "contracts/pools/BurnFromMintTokenPool.sol:BurnFromMintTokenPool")
	verification.RegisterContractMetadata(utils.LockReleaseTokenPool, v, lockreleasetokenpoolv2.SolidityStandardInput, lockreleasetokenpoolv2.LockReleaseTokenPoolBin, "contracts/pools/LockReleaseTokenPool.sol:LockReleaseTokenPool")
	verification.RegisterContractMetadata(versionedverifierresolverops.CommitteeVerifierResolverType, v, versioned_verifier_resolver.SolidityStandardInput, versioned_verifier_resolver.VersionedVerifierResolverBin, "contracts/ccvs/VersionedVerifierResolver.sol:VersionedVerifierResolver")
	verification.RegisterContractMetadata(versionedverifierresolverops.CCTPVerifierResolverType, v, versioned_verifier_resolver.SolidityStandardInput, versioned_verifier_resolver.VersionedVerifierResolverBin, "contracts/ccvs/VersionedVerifierResolver.sol:VersionedVerifierResolver")
	verification.RegisterContractMetadata(cctpverifierops.ContractType, v, cctp_verifier.SolidityStandardInput, cctp_verifier.CCTPVerifierBin, "contracts/ccvs/CCTPVerifier.sol:CCTPVerifier")
	verification.RegisterContractMetadata(cctpthroughccvtokenpoolops.ContractType, v, cctp_through_ccv_token_pool.SolidityStandardInput, cctp_through_ccv_token_pool.CCTPThroughCCVTokenPoolBin, "contracts/pools/USDC/CCTPThroughCCVTokenPool.sol:CCTPThroughCCVTokenPool")
	verification.RegisterContractMetadata(usdctokenpoolproxyops.ContractType, v, usdc_token_pool_proxy.SolidityStandardInput, usdc_token_pool_proxy.USDCTokenPoolProxyBin, "contracts/pools/USDC/USDCTokenPoolProxy.sol:USDCTokenPoolProxy")
	verification.RegisterContractMetadata(mock_receiver.ContractType, v, mock_receiver_v2.SolidityStandardInput, mock_receiver_v2.MockReceiverV2Bin, "contracts/test/mocks/MockReceiverV2.sol:MockReceiverV2")
	verification.RegisterContractMetadata(executorops.ContractType, v, executor.SolidityStandardInput, executor.ExecutorBin, "contracts/executor/Executor.sol:Executor")
	verification.RegisterContractMetadata(seq1_7.ExecutorProxyType, v, proxy.SolidityStandardInput, proxy.ProxyBin, "contracts/Proxy.sol:Proxy")
	verification.RegisterContractMetadata(committeeverifierops.ContractType, v, committee_verifier.SolidityStandardInput, committee_verifier.CommitteeVerifierBin, "contracts/ccvs/CommitteeVerifier.sol:CommitteeVerifier")
}
