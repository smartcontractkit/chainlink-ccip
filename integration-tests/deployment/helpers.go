package deployment

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnmERC20Bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	mcms_types "github.com/smartcontractkit/mcms/types"

	bnmERC20Operations "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	evmrouterops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	evmofframpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	evmonrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	sequencesV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	mcmsapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func DeployMCMS(t *testing.T, e *cldf_deployment.Environment, selector uint64, qualifiers []string) {
	dReg := mcmsapi.GetRegistry()
	version := semver.MustParse("1.6.0")
	cs := mcmsapi.DeployMCMS(dReg, nil)
	fcs := mcmsapi.FinalizeDeployMCMS(dReg, nil)
	for _, qualifier := range qualifiers {
		output, err := cs.Apply(*e, mcmsapi.MCMSDeploymentConfig{
			AdapterVersion: version,
			Chains: map[uint64]mcmsapi.MCMSDeploymentConfigPerChain{
				selector: {
					Canceller:        testhelpers.SingleGroupMCMS(),
					Bypasser:         testhelpers.SingleGroupMCMS(),
					Proposer:         testhelpers.SingleGroupMCMS(),
					TimelockMinDelay: big.NewInt(0),
					Qualifier:        ptr.String(qualifier),
				},
			},
		})
		require.NoError(t, err)
		require.Greater(t, len(output.Reports), 0)
		require.NoError(t, output.DataStore.Merge(e.DataStore))
		e.DataStore = output.DataStore.Seal()
		finalizeOutput, err := fcs.Apply(*e, mcmsapi.MCMSDeploymentConfig{
			AdapterVersion: version,
			Chains: map[uint64]mcmsapi.MCMSDeploymentConfigPerChain{
				selector: {
					Canceller:        testhelpers.SingleGroupMCMS(),
					Bypasser:         testhelpers.SingleGroupMCMS(),
					Proposer:         testhelpers.SingleGroupMCMS(),
					TimelockMinDelay: big.NewInt(0),
					Qualifier:        ptr.String(qualifier),
				},
			},
		})
		require.NoError(t, err)
		require.Greater(t, len(finalizeOutput.Reports), 0)
		require.NoError(t, finalizeOutput.DataStore.Merge(e.DataStore))
		e.DataStore = finalizeOutput.DataStore.Seal()
	}
}

func SolanaTransferOwnership(t *testing.T, e *cldf_deployment.Environment, selector uint64) {
	chain := e.BlockChains.SolanaChains()[selector]
	timelockSigner := utils.GetTimelockSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.CLLQualifier,
	)
	mcmSigner := utils.GetMCMSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.ProposerManyChainMultisig,
		common_utils.CLLQualifier,
	)
	err := utils.FundSolanaAccounts(
		t.Context(),
		[]solana.PublicKey{chain.DeployerKey.PublicKey()},
		100,
		chain.Client,
	)
	require.NoError(t, err)
	err = utils.FundFromDeployerKey(
		chain,
		[]solana.PublicKey{timelockSigner, mcmSigner},
		10,
	)
	require.NoError(t, err)
	mcmsInput := mcmsapi.TransferOwnershipInput{
		ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops.ContractType),
						Version: routerops.Version,
					},
				},
				CurrentOwner:  chain.DeployerKey.PublicKey().String(),
				ProposedOwner: timelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(offrampops.ContractType),
						Version: offrampops.Version,
					},
				},
				CurrentOwner:  chain.DeployerKey.PublicKey().String(),
				ProposedOwner: timelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(fqops.ContractType),
						Version: fqops.Version,
					},
				},
				CurrentOwner:  chain.DeployerKey.PublicKey().String(),
				ProposedOwner: timelockSigner.String(),
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(rmnremoteops.ContractType),
						Version: rmnremoteops.Version,
					},
				},
				CurrentOwner:  chain.DeployerKey.PublicKey().String(),
				ProposedOwner: timelockSigner.String(),
			},
		},
		AdapterVersion: semver.MustParse("1.6.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            common_utils.CLLQualifier,
			Description:          "Transfer ownership test",
		},
	}

	transferOutput, err := mcmsapi.TransferOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(transferOutput.Reports), 0)
	require.Equal(t, 1, len(transferOutput.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *e, transferOutput.MCMSTimelockProposals, false)
	// router verify
	program := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		routerops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, timelockSigner, routerops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// offramp verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		offrampops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, timelockSigner, offrampops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// fee quoter verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		fqops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, timelockSigner, fqops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
	// rmn remote verify
	program = datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		rmnremoteops.ContractType,
		common_utils.Version_1_6_0,
		"",
	)
	require.Equal(t, timelockSigner, rmnremoteops.GetAuthority(chain, solana.MustPublicKeyFromBase58(program.Address)))
}

func SolanaTransferMCMSContracts(t *testing.T, e *cldf_deployment.Environment, selector uint64, qualifier string, testTransferBack bool) {
	chain := e.BlockChains.SolanaChains()[selector]
	timelockSigner := utils.GetTimelockSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		qualifier,
	)
	mcmSigner := utils.GetMCMSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.ProposerManyChainMultisig,
		qualifier,
	)
	err := utils.FundSolanaAccounts(
		t.Context(),
		[]solana.PublicKey{chain.DeployerKey.PublicKey()},
		100,
		chain.Client,
	)
	require.NoError(t, err)
	err = utils.FundFromDeployerKey(
		chain,
		[]solana.PublicKey{timelockSigner, mcmSigner},
		10,
	)
	require.NoError(t, err)
	mcmsRefs := utils.GetAllMCMS(
		chain,
		qualifier,
		e.DataStore.Addresses().Filter(),
	)
	mcmsInput := mcmsapi.TransferOwnershipInput{
		ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain.Selector,
				ContractRef:   mcmsRefs,
				CurrentOwner:  chain.DeployerKey.PublicKey().String(),
				ProposedOwner: timelockSigner.String(),
			},
		},
		AdapterVersion: semver.MustParse("1.6.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            qualifier,
			Description:          "Transfer ownership test",
		},
	}

	transferOutput, err := mcmsapi.TransferOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(transferOutput.Reports), 0)
	require.Equal(t, 1, len(transferOutput.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *e, transferOutput.MCMSTimelockProposals, false)
	if testTransferBack {
		mcmsRefs = mcmsRefs[len(mcmsRefs)-4:] // get only mcms refs
		// transfer back to mcmSigner
		mcmsInputBack := mcmsapi.TransferOwnershipInput{
			ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
				{
					ChainSelector: chain.Selector,
					ContractRef:   mcmsRefs,
					CurrentOwner:  timelockSigner.String(),
					ProposedOwner: chain.DeployerKey.PublicKey().String(),
				},
			},
			AdapterVersion: semver.MustParse("1.6.0"),
			MCMS: mcms.Input{
				OverridePreviousRoot: false,
				ValidUntil:           3759765795,
				TimelockDelay:        mcms_types.MustParseDuration("1s"),
				TimelockAction:       mcms_types.TimelockActionSchedule,
				Qualifier:            qualifier,
				Description:          "Transfer ownership back test",
			},
		}
		transferBackOutput, err := mcmsapi.TransferOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInputBack)
		require.NoError(t, err)
		require.Greater(t, len(transferBackOutput.Reports), 0)
		require.Equal(t, 1, len(transferBackOutput.MCMSTimelockProposals))

		testhelpers.ProcessTimelockProposals(t, *e, transferBackOutput.MCMSTimelockProposals, false)

		acceptBackOutput, err := mcmsapi.AcceptOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInputBack)
		require.NoError(t, err)
		require.Greater(t, len(acceptBackOutput.Reports), 0)
		require.Equal(t, 0, len(acceptBackOutput.MCMSTimelockProposals))
	}
}

func EVMTransferOwnership(t *testing.T, e *cldf_deployment.Environment, selector uint64) {
	chain := e.BlockChains.EVMChains()[selector]
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range e.DataStore.Addresses().Filter() {
		if addrRef.Type == datastore.ContractType(common_utils.RBACTimelock) {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}

	// Add the timelock as an authorized caller (price updater) on the FeeQuoter.
	// This must happen while the deployer is still the owner (before ownership transfer).
	// The FeeQuoter's updatePrices requires msg.sender to be in the authorizedCallers set.
	if timelockAddr, ok := timelockAddrs[selector]; ok {
		for _, addrRef := range e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(selector),
			datastore.AddressRefByType(datastore.ContractType(fq163ops.ContractType)),
		) {
			fq, err := fq163ops.NewFeeQuoterContract(common.HexToAddress(addrRef.Address), chain.Client)
			require.NoError(t, err, "failed to create FeeQuoter contract instance")
			tx, err := fq.ApplyAuthorizedCallerUpdates(chain.DeployerKey, fq163ops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{common.HexToAddress(timelockAddr)},
			})
			require.NoError(t, err, "failed to add timelock as FeeQuoter price updater")
			_, err = chain.Confirm(tx)
			require.NoError(t, err, "failed to confirm FeeQuoter authorized caller update")
			t.Logf("Added timelock %s as authorized caller on FeeQuoter %s", timelockAddr, addrRef.Address)
		}
	}

	mcmsInput := mcmsapi.TransferOwnershipInput{
		ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(fq163ops.ContractType),
						Version: fq163ops.Version,
					},
				},
				ProposedOwner: timelockAddrs[chain.Selector],
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(evmrouterops.ContractType),
						Version: evmrouterops.Version,
					},
				},
				ProposedOwner: timelockAddrs[chain.Selector],
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(evmofframpops.ContractType),
						Version: evmofframpops.Version,
					},
				},
				ProposedOwner: timelockAddrs[chain.Selector],
			},
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(evmonrampops.ContractType),
						Version: evmonrampops.Version,
					},
				},
				ProposedOwner: timelockAddrs[chain.Selector],
			},
		},
		AdapterVersion: semver.MustParse("1.6.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            common_utils.CLLQualifier,
			Description:          "Transfer ownership test",
		},
	}
	transferOutput, err := mcmsapi.TransferOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(transferOutput.Reports), 0)
	require.Equal(t, 1, len(transferOutput.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, transferOutput.MCMSTimelockProposals, false)
}

func MergeAddresses(t *testing.T, env *cldf_deployment.Environment, ds datastore.MutableDataStore) {
	t.Helper()

	if ds != nil {
		require.NoError(t, ds.Merge(env.DataStore))
		env.DataStore = ds.Seal()
	}
}

func NewDefaultDeploymentConfigForSolana(version *semver.Version) mcmsapi.ContractDeploymentConfigPerChain {
	return mcmsapi.ContractDeploymentConfigPerChain{
		Version:                      version,
		MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
		TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
		NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
		LinkPremiumMultiplier:        9e17, // 0.9 ETH
		TokenPrivKey:                 solana.NewWallet().PrivateKey.String(),
		TokenDecimals:                9,
	}
}

func NewDefaultDeploymentConfigForEVM(version *semver.Version) mcmsapi.ContractDeploymentConfigPerChain {
	return mcmsapi.ContractDeploymentConfigPerChain{
		Version:                                 version,
		MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
		TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
		NativeTokenPremiumMultiplier:            1e18, // 1.0 ETH
		LinkPremiumMultiplier:                   9e17, // 0.9 ETH
		PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
		GasForCallExactCheck:                    uint16(5000),
		TokenDecimals:                           18,
	}
}

func NewLaneChainDefinitionForV2(chainSelector, remoteChainSelector uint64) lanes.ChainDefinition {
	return lanes.ChainDefinition{
		Selector: chainSelector,
		CommitteeVerifiers: []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(committee_verifier.ContractType),
						Version:       committee_verifier.Version,
						Qualifier:     "alpha",
					},
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(sequencesV2_0_0.CommitteeVerifierResolverType),
						Version:       common_utils.Version_2_0_0,
					},
				},
				RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
					remoteChainSelector: testsetupV2_0_0.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
		DefaultInboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
				Qualifier:     "alpha",
			},
		},
		DefaultOutboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
				Qualifier:     "alpha",
			},
		},
		DefaultExecutor: datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(sequencesV2_0_0.ExecutorProxyType),
			Version:       executor.Version,
			Qualifier:     "default",
		},
		FeeQuoterDestChainConfig: testsetupV2_0_0.CreateBasicFeeQuoterDestChainConfig(),
		ExecutorDestChainConfig:  testsetupV2_0_0.CreateBasicExecutorDestChainConfig(),
		AddressBytesLength:       20,
		BaseExecutionGasCost:     80_000,
	}
}

func NewDefaultInputForMCMS(desc string, overrides ...func(*mcms.Input)) mcms.Input {
	in := mcms.Input{
		OverridePreviousRoot: false,
		ValidUntil:           math.MaxUint32,
		TimelockDelay:        mcms_types.MustParseDuration("1s"),
		TimelockAction:       mcms_types.TimelockActionSchedule,
		Qualifier:            common_utils.CLLQualifier,
		Description:          desc,
	}
	for _, override := range overrides {
		override(&in)
	}
	return in
}

func DeployChainContractsV2_0_0(t *testing.T, e *cldf_deployment.Environment, cumulativeDS *datastore.MemoryDataStore, chainSel uint64) {
	t.Helper()

	create2FactoryRef, err := evm_contract.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel],
		evm_contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version),
			ChainSelector:  chainSel,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
			},
		},
		nil,
	)
	require.NoError(t, err)

	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequencesV2_0_0.DeployChainContracts,
		e.BlockChains.EVMChains()[chainSel],
		sequencesV2_0_0.DeployChainContractsInput{
			ChainSelector:    chainSel,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetupV2_0_0.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)

	for _, addr := range chainReport.Output.Addresses {
		require.NoError(t, cumulativeDS.Addresses().Add(addr))
	}
}

func RequireBigIntsEqual(t *testing.T, want, got *big.Int, msg string) {
	t.Helper()
	require.NotNil(t, got, msg)
	require.Zero(t, want.Cmp(got), "%s: want %s got %s", msg, want.String(), got.String())
}

// RequireBigIntsApprox asserts |want-got| <= tolerance. Used for rate-limit values where GenerateTPRLConfigs'
// float64 x1.1 premium introduces sub-token rounding noise that differs across pool-version code paths.
func RequireBigIntsApprox(t *testing.T, want, got *big.Int, tolerance int64, msg string) {
	t.Helper()
	require.NotNil(t, got, msg)
	diff := new(big.Int).Abs(new(big.Int).Sub(want, got))
	require.LessOrEqualf(t, diff.Cmp(big.NewInt(tolerance)), 0, "%s: want ~%s got %s (diff %s > tolerance %d)", msg, want, got, diff, tolerance)
}

func NewRandHex(t *testing.T, nBytes int) string {
	t.Helper()
	data := make([]byte, nBytes)
	_, err := rand.Read(data)
	require.NoError(t, err)
	return hex.EncodeToString(data)
}

func DeployBurnMintPoolEVM(t *testing.T, env *cldf_deployment.Environment, selector uint64, version *semver.Version, token common.Address) datastore.AddressRef {
	t.Helper()

	poolQualifier := "Pool-" + NewRandHex(t, 32)
	poolType := common_utils.BurnMintTokenPool.String()

	output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: version,
		MCMS:                NewDefaultInputForMCMS(fmt.Sprintf("Deploy %s", poolQualifier)),
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selector: {
				TokenPoolVersion: version,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenRef:           &datastore.AddressRef{Address: token.Hex()},
					TokenPoolQualifier: poolQualifier,
					PoolType:           poolType,
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	dsFilter := datastore.AddressRef{ChainSelector: selector, Qualifier: poolQualifier, Type: datastore.ContractType(poolType)}
	ref, err := datastore_utils.FindAndFormatRef(env.DataStore, dsFilter, selector, datastore_utils.FullRef)
	require.NoError(t, err)

	return ref
}

func DeployBurnMintTokenEVM(t *testing.T, env *cldf_deployment.Environment, selector uint64, admin string) *bnmERC20Bindings.BurnMintERC20 {
	t.Helper()

	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)

	symbol := "Token-" + NewRandHex(t, 32)
	output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: common_utils.Version_1_0_0,
		MCMS:                NewDefaultInputForMCMS(fmt.Sprintf("Deploy %s", symbol)),
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selector: {
				DeployTokenInput: &tokensapi.DeployTokenInput{
					ExternalAdmin: admin,
					CCIPAdmin:     admin,
					Decimals:      uint8(18),
					Supply:        nil, // unlimited supply
					Symbol:        symbol,
					Name:          symbol + " Token",
					Type:          bnmERC20Operations.ContractType,
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	dsFilter := datastore.AddressRef{ChainSelector: selector, Qualifier: symbol, Type: datastore.ContractType(bnmERC20Operations.ContractType)}
	ref, err := datastore_utils.FindAndFormatRef(env.DataStore, dsFilter, selector, datastore_utils.FullRef)
	require.NoError(t, err)

	require.True(t, common.IsHexAddress(ref.Address))
	token, err := bnmERC20Bindings.NewBurnMintERC20(common.HexToAddress(ref.Address), chain.Client)
	require.NoError(t, err)

	return token
}
