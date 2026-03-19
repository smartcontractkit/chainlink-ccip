package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	mock_recv_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_receiver_v2"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func TestDeployChainContracts_Idempotency(t *testing.T) {
	tests := []struct {
		desc              string
		existingAddresses []datastore.AddressRef
	}{
		{
			desc: "full deployment",
		},
		{
			desc: "partial deployment",
			existingAddresses: []datastore.AddressRef{
				{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(link.ContractType),
					Version:       link.Version,
					Address:       common.HexToAddress("0x01").Hex(),
				},
				{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(weth.ContractType),
					Version:       weth.Version,
					Address:       common.HexToAddress("0x02").Hex(),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSelector := uint64(5009297550715157269)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")

			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSelector], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
				ChainSelector:  chainSelector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{e.BlockChains.EVMChains()[chainSelector].DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")
			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSelector],
				sequences.DeployChainContractsInput{
					ChainSelector:     chainSelector,
					ExistingAddresses: test.existingAddresses,
					CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
					ContractParams:    testsetup.CreateBasicContractParams(),
					DeployTestRouter:  true,
					DeployerKeyOwned:  true,
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.BatchOps, 2, "Expected 2 batch operations")

			exists := map[deployment.ContractType]bool{
				rmn_remote.ContractType:                 false,
				router.ContractType:                     false,
				executor.ContractType:                   false,
				link.ContractType:                       false,
				weth.ContractType:                       false,
				committee_verifier.ContractType:         false,
				onramp.ContractType:                     false,
				offramp.ContractType:                    false,
				fee_quoter.ContractType:                 false,
				sequences.CommitteeVerifierResolverType: false,
				rmn_proxy.ContractType:                  false,
				token_admin_registry.ContractType:       false,
				mock_receiver.ContractType:              false,
				sequences.ExecutorProxyType:             false,
				router.TestRouterContractType:           false,
			}
			for _, addr := range report.Output.Addresses {
				exists[deployment.ContractType(addr.Type)] = true
			}
			for ctype, found := range exists {
				require.True(t, found, "Expected contract of type %s to be deployed", ctype)
			}

			for _, existing := range test.existingAddresses {
				found := false
				for _, addr := range report.Output.Addresses {
					if addr.Type == existing.Type {
						require.Equal(t, existing.Address, addr.Address, "Expected existing address to be reused for %s", existing.Type)
						found = true
						break
					}
				}
				require.True(t, found, "Expected to find existing address for %s", existing.Type)
			}
		})
	}
}

func TestDeployChainContracts_MultipleDeployments(t *testing.T) {
	t.Run("sequential deployments", func(t *testing.T) {
		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4949039107694359620, 6433500567565415381}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")
		evmChains := e.BlockChains.EVMChains()

		// Deploy to each chain sequentially using the same bundle
		var allReports []operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]
		for _, evmChain := range evmChains {
			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
				ChainSelector:  evmChain.Selector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{evmChain.DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")
			input := sequences.DeployChainContractsInput{
				ChainSelector:     evmChain.Selector,
				ExistingAddresses: nil,
				ContractParams:    testsetup.CreateBasicContractParams(),
				CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
				DeployerKeyOwned:  true,
			}

			report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, evmChain, input)
			require.NoError(t, err, "Failed to execute sequence for chain %d", evmChain.Selector)
			require.NotEmpty(t, report.Output.Addresses, "Expected operation reports for chain %d", evmChain.Selector)

			allReports = append(allReports, report)
		}

		// Verify all deployments succeeded
		require.Len(t, allReports, len(evmChains), "Expected reports for all chains")

		for _, report := range allReports {
			require.NotEmpty(t, report.Output.Addresses, "Expected addresses")
		}
	})

	t.Run("concurrent deployments", func(t *testing.T) {
		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4949039107694359620, 6433500567565415381}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")
		evmChains := e.BlockChains.EVMChains()

		// Deploy to all chains concurrently using the same bundle
		type deployResult struct {
			chainSelector uint64
			report        operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]
			err           error
		}

		resultChan := make(chan deployResult, len(evmChains))

		// Launch concurrent deployments
		for _, evmChain := range evmChains {
			go func(chainSel uint64) {
				evmChain := evmChains[chainSel]
				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
					ChainSelector:  evmChain.Selector,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{evmChain.DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				input := sequences.DeployChainContractsInput{
					ChainSelector:     chainSel,
					ExistingAddresses: nil,
					ContractParams:    testsetup.CreateBasicContractParams(),
					CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
					DeployerKeyOwned:  true,
				}

				report, execErr := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, evmChain, input)
				resultChan <- deployResult{chainSel, report, execErr}
			}(evmChain.Selector)
		}

		// Collect all results
		var results []deployResult
		for i := 0; i < len(evmChains); i++ {
			result := <-resultChan
			results = append(results, result)
		}

		// Verify all deployments succeeded
		require.Len(t, results, len(evmChains), "Expected results for all chains")

		for _, result := range results {
			require.NoError(t, result.err, "Failed to execute sequence for chain %d", result.chainSelector)
			require.NotEmpty(t, result.report.Output.Addresses, "Expected addresses for chain %d", result.chainSelector)
		}
	})
}

func TestDeployChainContracts_MultipleCommitteeVerifiersAndMultipleMockReceiverConfigs(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	// Configure two committee verifiers with different qualifiers and request both for MockReceiver
	params := testsetup.CreateBasicContractParams()
	params.CommitteeVerifiers = []sequences.CommitteeVerifierParams{
		{
			Version:          committee_verifier.Version,
			FeeAggregator:    common.HexToAddress("0x01"),
			StorageLocations: []string{"https://test.chain.link.fake"},
			Qualifier:        "alpha",
		},
		{
			Version:          committee_verifier.Version,
			FeeAggregator:    common.HexToAddress("0x01"),
			StorageLocations: []string{"https://test.chain.link.fake"},
			Qualifier:        "beta",
		},
	}
	params.MockReceivers = []sequences.MockReceiverParams{
		{
			Version: mock_receiver.Version,
			RequiredVerifiers: []datastore.AddressRef{
				{
					ChainSelector: chainSelector,
					Type:          datastore.ContractType(committee_verifier.ContractType),
					Version:       committee_verifier.Version,
					Qualifier:     "alpha",
				},
				{
					ChainSelector: chainSelector,
					Type:          datastore.ContractType(committee_verifier.ContractType),
					Version:       committee_verifier.Version,
					Qualifier:     "beta",
				},
			},
			Qualifier: "q1",
		},
		{
			Version: mock_receiver.Version,
			RequiredVerifiers: []datastore.AddressRef{
				{
					Type:      datastore.ContractType(committee_verifier.ContractType),
					Version:   committee_verifier.Version,
					Qualifier: "alpha",
				},
			},
			OptionalVerifiers: []datastore.AddressRef{
				{
					Type:      datastore.ContractType(committee_verifier.ContractType),
					Version:   committee_verifier.Version,
					Qualifier: "beta",
				},
			},
			OptionalThreshold: 1,
			Qualifier:         "q2",
		},
	}

	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSelector], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{e.BlockChains.EVMChains()[chainSelector].DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployChainContractsInput{
			ChainSelector:     chainSelector,
			ExistingAddresses: nil,
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			ContractParams:    params,
			DeployerKeyOwned:  true,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	// Assert mock receiver properties
	ds := datastore.NewMemoryDataStore()
	for _, addr := range report.Output.Addresses {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	sealed := ds.Seal()

	q1ReceiverRef, err := datastore_utils.FindAndFormatRef(sealed, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(mock_receiver.ContractType),
		Version:       mock_receiver.Version,
		Qualifier:     "q1",
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	q2ReceiverRef, err := datastore_utils.FindAndFormatRef(sealed, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(mock_receiver.ContractType),
		Version:       mock_receiver.Version,
		Qualifier:     "q2",
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	q1Receiver, err := mock_recv_bindings.NewMockReceiverV2(q1ReceiverRef, e.BlockChains.EVMChains()[chainSelector].Client)
	require.NoError(t, err)

	q1Result, err := q1Receiver.GetCCVsAndMinBlockConfirmations(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, chainSelector, []byte{})
	require.NoError(t, err)
	require.Len(t, q1Result.RequiredVerifier, 2)
	require.Len(t, q1Result.OptionalVerifiers, 0)
	require.Equal(t, uint8(0), q1Result.Threshold)

	q2Receiver, err := mock_recv_bindings.NewMockReceiverV2(q2ReceiverRef, e.BlockChains.EVMChains()[chainSelector].Client)
	require.NoError(t, err)

	q2Result, err := q2Receiver.GetCCVsAndMinBlockConfirmations(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, chainSelector, []byte{})
	require.NoError(t, err)
	require.Len(t, q2Result.RequiredVerifier, 1)
	require.Len(t, q2Result.OptionalVerifiers, 1)
	require.Equal(t, uint8(1), q2Result.Threshold)
}

func singleSignerMCMSConfig(signer common.Address) (mcms_types.Config, error) {
	return mcms_types.NewConfig(1, []common.Address{signer}, nil)
}

// deployMCMSInstanceForTest deploys one MCMS instance (proposer, bypasser, canceller,
// timelock, call proxy) for the given qualifier and returns the timelock address and
// the full list of deployed address refs.
func deployMCMSInstanceForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	deployer common.Address,
	qualifier string,
) (timelockAddr common.Address, addresses []datastore.AddressRef) {
	t.Helper()
	mcmsCfg, err := singleSignerMCMSConfig(deployer)
	require.NoError(t, err)

	qualifierPtr := &qualifier

	proposerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.ProposerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	require.NotEmpty(t, proposerReport.Output.Addresses)
	addresses = append(addresses, proposerReport.Output.Addresses...)

	bypasserReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.BypasserManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	require.NotEmpty(t, bypasserReport.Output.Addresses)
	addresses = append(addresses, bypasserReport.Output.Addresses...)

	cancellerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.CancellerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	require.NotEmpty(t, cancellerReport.Output.Addresses)
	addresses = append(addresses, cancellerReport.Output.Addresses...)

	timelockRef, err := contract_utils.MaybeDeployContract(b, mcms_ops.OpDeployTimelock, chain, contract_utils.DeployInput[mcms_ops.OpDeployTimelockInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.RBACTimelock, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployTimelockInput{
			TimelockMinDelay: big.NewInt(0),
			Proposers:        []common.Address{common.HexToAddress(proposerReport.Output.Addresses[0].Address)},
			Bypassers:        []common.Address{common.HexToAddress(bypasserReport.Output.Addresses[0].Address)},
			Cancellers:       []common.Address{common.HexToAddress(cancellerReport.Output.Addresses[0].Address)},
			Admin:            deployer,
			Executors:        []common.Address{},
		},
	}, nil)
	require.NoError(t, err)
	timelockAddr = common.HexToAddress(timelockRef.Address)
	addresses = append(addresses, timelockRef)

	callProxyRef, err := contract_utils.MaybeDeployContract(b, mcms_ops.OpDeployCallProxy, chain, contract_utils.DeployInput[mcms_ops.OpDeployCallProxyInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.CallProxy, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployCallProxyInput{
			TimelockAddress: timelockAddr,
		},
	}, nil)
	require.NoError(t, err)
	addresses = append(addresses, callProxyRef)

	return timelockAddr, addresses
}

// deployAllMCMSForTest deploys both CLLCCIP and RMNMCMS instances and returns the
// CLL and RMN timelock addresses along with the combined list of all MCMS address refs.
func deployAllMCMSForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	deployer common.Address,
) (cllTimelockAddr, rmnTimelockAddr common.Address, addresses []datastore.AddressRef) {
	t.Helper()

	cllTimelockAddr, cllAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.CLLQualifier)
	addresses = append(addresses, cllAddrs...)

	rmnTimelockAddr, rmnAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)
	addresses = append(addresses, rmnAddrs...)

	return cllTimelockAddr, rmnTimelockAddr, addresses
}

// sealAddressRefs builds a sealed read-only DataStore from one or more address ref
// slices. Duplicate refs (e.g. when combining MCMS + output refs) are silently ignored.
func sealAddressRefs(t *testing.T, refGroups ...[]datastore.AddressRef) datastore.DataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	for _, group := range refGroups {
		for _, ref := range group {
			_ = ds.Addresses().Add(ref)
		}
	}
	return ds.Seal()
}

// requireContractOwner loads the Ownable contract at addr and asserts its owner is one of validOwners.
func requireContractOwner(
	t *testing.T,
	chain evm.Chain,
	addr common.Address,
	ct deployment.ContractType,
	qualifier string,
	validOwners []common.Address,
) {
	t.Helper()
	_, ownable, err := mcms_seq.LoadOwnableContract(addr, chain.Client)
	require.NoError(t, err, "LoadOwnableContract for %s/%s at %s", ct, qualifier, addr)
	owner, err := ownable.Owner(nil)
	require.NoError(t, err, "Owner() for %s/%s at %s", ct, qualifier, addr)
	for _, vo := range validOwners {
		if owner == vo {
			return
		}
	}
	require.Failf(t, "unexpected owner",
		"%s/%s at %s: owner %s not in valid set %v", ct, qualifier, addr, owner, validOwners)
}

func TestDeployChainContracts_DefaultOwnershipTransfer(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	// Pre-deploy both MCMS instances (CLLCCIP and RMNMCMS) before running DeployChainContracts.
	cllTimelockAddr, rmnTimelockAddr, mcmsAddresses := deployAllMCMSForTest(t, e.OperationsBundle, chain, deployer)

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{deployer}},
		}, nil)
	require.NoError(t, err)

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:     chainSelector,
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			ContractParams:    testsetup.CreateBasicContractParams(),
			DeployTestRouter:  true,
			ExistingAddresses: mcmsAddresses,
		},
	)
	require.NoError(t, err, "ExecuteSequence with default ownership transfer should not error")

	outputDS := sealAddressRefs(t, report.Output.Addresses)
	mcmsDS := sealAddressRefs(t, mcmsAddresses)

	// Verify single-instance product contracts are deployed and ownership transferred.
	// For Ownable2Step contracts, owner is still the deployer (pending transfer not yet
	// accepted). For one-step Ownable contracts, owner is the timelock directly.
	singleProductTypes := []deployment.ContractType{
		rmn_remote.ContractType,
		router.ContractType,
		token_admin_registry.ContractType,
		fee_quoter.ContractType,
		offramp.ContractType,
		onramp.ContractType,
	}
	for _, ct := range singleProductTypes {
		addr, err := datastore_utils.FindAndFormatRef(outputDS, datastore.AddressRef{
			Type: datastore.ContractType(ct),
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Expected product contract %s to be deployed", ct)
		requireContractOwner(t, chain, addr, ct, "", []common.Address{deployer, cllTimelockAddr})
	}

	// Verify multi-instance executor contracts (with qualifiers) deployed and ownership transferred.
	for _, q := range []string{"default", "custom"} {
		for _, ct := range []deployment.ContractType{executor.ContractType, sequences.ExecutorProxyType} {
			addr, err := datastore_utils.FindAndFormatRef(outputDS, datastore.AddressRef{
				Type:      datastore.ContractType(ct),
				Qualifier: q,
			}, chainSelector, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err, "Expected %s/%s to be deployed", ct, q)
			requireContractOwner(t, chain, addr, ct, q, []common.Address{deployer, cllTimelockAddr})
		}
	}

	// Verify MCM contracts (Proposer, Bypasser, Canceller) have ownership transferred.
	// CLL MCMs must be owned by CLL timelock; RMN MCMs may also be owned by RMN timelock.
	mcmTypes := []deployment.ContractType{
		common_utils.ProposerManyChainMultisig,
		common_utils.BypasserManyChainMultisig,
		common_utils.CancellerManyChainMultisig,
	}
	for _, qualifier := range []string{common_utils.CLLQualifier, common_utils.RMNTimelockQualifier} {
		validOwners := []common.Address{deployer, cllTimelockAddr}
		if qualifier == common_utils.RMNTimelockQualifier {
			validOwners = append(validOwners, rmnTimelockAddr)
		}
		for _, ct := range mcmTypes {
			addr, err := datastore_utils.FindAndFormatRef(mcmsDS, datastore.AddressRef{
				Type:      datastore.ContractType(ct),
				Qualifier: qualifier,
			}, chainSelector, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err, "Expected to find MCM %s/%s in mcmsAddresses", ct, qualifier)
			requireContractOwner(t, chain, addr, ct, qualifier, validOwners)
		}
	}

	// Verify CLLCCIP timelock is self-governed and deployer is no longer admin.
	cllTimelock, err := bindings.NewRBACTimelock(cllTimelockAddr, chain.Client)
	require.NoError(t, err)

	cllSelfAdmin, err := cllTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, cllTimelockAddr)
	require.NoError(t, err)
	require.True(t, cllSelfAdmin, "CLLCCIP timelock should be admin of itself")

	cllDeployerAdmin, err := cllTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, deployer)
	require.NoError(t, err)
	require.False(t, cllDeployerAdmin, "Deployer should no longer be admin of CLLCCIP timelock")

	// Verify RMNMCMS timelock is governed by CLLCCIP timelock and deployer is no longer admin.
	rmnTimelock, err := bindings.NewRBACTimelock(rmnTimelockAddr, chain.Client)
	require.NoError(t, err)

	rmnCllAdmin, err := rmnTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, cllTimelockAddr)
	require.NoError(t, err)
	require.True(t, rmnCllAdmin, "CLLCCIP timelock should be admin of RMNMCMS timelock")

	rmnDeployerAdmin, err := rmnTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, deployer)
	require.NoError(t, err)
	require.False(t, rmnDeployerAdmin, "Deployer should no longer be admin of RMNMCMS timelock")
}

func TestDeployChainContracts_DefaultOwnershipTransfer_Idempotent(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	cllTimelockAddr, _, mcmsAddresses := deployAllMCMSForTest(t, e.OperationsBundle, chain, deployer)

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{deployer}},
		}, nil)
	require.NoError(t, err)

	input := sequences.DeployChainContractsInput{
		ChainSelector:     chainSelector,
		CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
		ContractParams:    testsetup.CreateBasicContractParams(),
		DeployTestRouter:  true,
		ExistingAddresses: mcmsAddresses,
	}

	// First run: deploys contracts, transfers ownership, sets up governance.
	firstReport, err := operations.ExecuteSequence(
		e.OperationsBundle, sequences.DeployChainContracts, chain, input,
	)
	require.NoError(t, err, "First run should succeed")

	// Second run: pass first run's output as existing addresses alongside MCMS refs.
	// Everything should be idempotent: contracts reused, ownership already transferred,
	// timelocks already self-governed.
	input.ExistingAddresses = append(mcmsAddresses, firstReport.Output.Addresses...)
	secondReport, err := operations.ExecuteSequence(
		e.OperationsBundle, sequences.DeployChainContracts, chain, input,
	)
	require.NoError(t, err, "Second run (idempotent) should succeed")

	// Verify the same addresses were returned (contracts reused, not redeployed).
	require.Equal(t, len(firstReport.Output.Addresses), len(secondReport.Output.Addresses),
		"Idempotent run should return the same number of addresses")
	firstAddrs := make(map[string]bool)
	for _, ref := range firstReport.Output.Addresses {
		firstAddrs[ref.Address] = true
	}
	for _, ref := range secondReport.Output.Addresses {
		require.True(t, firstAddrs[ref.Address],
			"Address %s (type=%s) from second run should match first run", ref.Address, ref.Type)
	}

	// Verify ownership is still pointing to CLL timelock after second run.
	outputDS := sealAddressRefs(t, secondReport.Output.Addresses)
	singleProductTypes := []deployment.ContractType{
		rmn_remote.ContractType,
		router.ContractType,
		token_admin_registry.ContractType,
		fee_quoter.ContractType,
		offramp.ContractType,
		onramp.ContractType,
	}
	for _, ct := range singleProductTypes {
		addr, err := datastore_utils.FindAndFormatRef(outputDS, datastore.AddressRef{
			Type: datastore.ContractType(ct),
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Expected product contract %s after idempotent run", ct)
		requireContractOwner(t, chain, addr, ct, "", []common.Address{deployer, cllTimelockAddr})
	}
}

func TestDeployChainContracts_DeployerKeyOwned(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	// Pre-deploy MCMS so addresses exist, but run with DeployerKeyOwned: true to skip transfer.
	cllTimelockAddr, _, mcmsAddresses := deployAllMCMSForTest(t, e.OperationsBundle, chain, deployer)

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{deployer}},
		}, nil)
	require.NoError(t, err)

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:     chainSelector,
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			ContractParams:    testsetup.CreateBasicContractParams(),
			DeployTestRouter:  true,
			ExistingAddresses: mcmsAddresses,
			DeployerKeyOwned:  true,
		},
	)
	require.NoError(t, err, "ExecuteSequence with DeployerKeyOwned should succeed")

	outputDS := sealAddressRefs(t, report.Output.Addresses)

	// Verify product contracts are deployed but still owned by deployer (no transfer).
	singleProductTypes := []deployment.ContractType{
		rmn_remote.ContractType,
		router.ContractType,
		token_admin_registry.ContractType,
		fee_quoter.ContractType,
		offramp.ContractType,
		onramp.ContractType,
	}
	for _, ct := range singleProductTypes {
		addr, err := datastore_utils.FindAndFormatRef(outputDS, datastore.AddressRef{
			Type: datastore.ContractType(ct),
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err, "Expected product contract %s to be deployed", ct)
		requireContractOwner(t, chain, addr, ct, "", []common.Address{deployer})
	}

	// Verify deployer is still admin of CLL timelock (governance unchanged).
	cllTimelock, err := bindings.NewRBACTimelock(cllTimelockAddr, chain.Client)
	require.NoError(t, err)

	deployerStillAdmin, err := cllTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, deployer)
	require.NoError(t, err)
	require.True(t, deployerStillAdmin, "Deployer should still be admin of CLL timelock when DeployerKeyOwned is true")

	timelockSelfAdmin, err := cllTimelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, cllTimelockAddr)
	require.NoError(t, err)
	require.False(t, timelockSelfAdmin, "CLL timelock should NOT be self-governed when DeployerKeyOwned is true")
}

func TestDeployChainContracts_DefaultTransfer_FailsWithoutMCMS(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{deployer}},
		}, nil)
	require.NoError(t, err)

	t.Run("no MCMS at all", func(t *testing.T) {
		_, err = operations.ExecuteSequence(
			e.OperationsBundle,
			sequences.DeployChainContracts,
			chain,
			sequences.DeployChainContractsInput{
				ChainSelector:  chainSelector,
				CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
				ContractParams: testsetup.CreateBasicContractParams(),
			},
		)
		require.Error(t, err, "Expected error when default ownership transfer finds no MCMS in ExistingAddresses")
		require.Contains(t, err.Error(), common_utils.CLLQualifier)
	})

	t.Run("only CLLCCIP deployed, missing RMNMCMS", func(t *testing.T) {
		_, cllAddresses := deployMCMSInstanceForTest(t, e.OperationsBundle, chain, deployer, common_utils.CLLQualifier)

		_, err = operations.ExecuteSequence(
			e.OperationsBundle,
			sequences.DeployChainContracts,
			chain,
			sequences.DeployChainContractsInput{
				ChainSelector:     chainSelector,
				CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
				ContractParams:    testsetup.CreateBasicContractParams(),
				ExistingAddresses: cllAddresses,
			},
		)
		require.Error(t, err, "Expected error when default ownership transfer finds RMNMCMS not in ExistingAddresses")
		require.Contains(t, err.Error(), common_utils.RMNTimelockQualifier)
	})

	t.Run("timelocks present but MCM contracts missing", func(t *testing.T) {
		_, _, allAddrs := deployAllMCMSForTest(t, e.OperationsBundle, chain, deployer)

		// Keep only timelock and call proxy refs, stripping out Proposer/Bypasser/Canceller.
		var timelockOnlyAddrs []datastore.AddressRef
		for _, ref := range allAddrs {
			ct := deployment.ContractType(ref.Type)
			if ct == common_utils.RBACTimelock || ct == common_utils.CallProxy {
				timelockOnlyAddrs = append(timelockOnlyAddrs, ref)
			}
		}

		_, err = operations.ExecuteSequence(
			e.OperationsBundle,
			sequences.DeployChainContracts,
			chain,
			sequences.DeployChainContractsInput{
				ChainSelector:     chainSelector,
				CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
				ContractParams:    testsetup.CreateBasicContractParams(),
				ExistingAddresses: timelockOnlyAddrs,
			},
		)
		require.Error(t, err, "Expected error when timelocks exist but MCM contracts are missing")
		require.Contains(t, err.Error(), "ownership transfer requires MCM contract")
	})
}
