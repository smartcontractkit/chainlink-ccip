package deployment

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	evmrouterops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	evmfqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	evmofframpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	evmonrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	mcmsapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"

	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

func DeployMCMS(t *testing.T, e *cldf_deployment.Environment, selector uint64) {
	// For EVM only, set the timelock admin
	var timelockAdmin common.Address
	chain1, ok := e.BlockChains.EVMChains()[selector]
	if ok {
		timelockAdmin = chain1.DeployerKey.From
	}
	dReg := mcmsapi.GetRegistry()
	version := semver.MustParse("1.6.0")
	cs := mcmsapi.DeployMCMS(dReg)
	output, err := cs.Apply(*e, mcmsapi.MCMSDeploymentConfig{
		AdapterVersion: version,
		Chains: map[uint64]mcmsapi.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(common_utils.CLLQualifier),
				TimelockAdmin:    timelockAdmin,
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	output.DataStore.Merge(e.DataStore)
	e.DataStore = output.DataStore.Seal()
}

func SolanaTransferOwnership(t *testing.T, e *cldf_deployment.Environment, selector uint64) {
	chain := e.BlockChains.SolanaChains()[selector]
	timelockSigner := utils.GetTimelockSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.CLLQualifier)
	mcmSigner := utils.GetMCMSignerPDA(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.CLLQualifier,
	)
	timelockCompositeAddress := utils.GetTimelockCompositeAddress(
		e.DataStore.Addresses().Filter(),
		chain.Selector,
		common_utils.CLLQualifier)
	err := utils.FundSolanaAccounts(
		context.Background(),
		[]solana.PublicKey{timelockSigner, mcmSigner},
		100,
		chain.Client,
	)
	require.NoError(t, err)
	t.Logf("Timelock Signer PDA: %s", timelockSigner.String())
	t.Logf("Timelock Composite Address: %s", timelockCompositeAddress)
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
	require.Equal(t, 0, len(transferOutput.MCMSTimelockProposals))

	acceptOutput, err := mcmsapi.AcceptOwnershipChangeset(mcmsapi.GetTransferOwnershipRegistry(), mcmsreaderapi.GetRegistry()).Apply(*e, mcmsInput)
	require.NoError(t, err)
	require.Greater(t, len(acceptOutput.Reports), 0)
	require.Equal(t, 1, len(acceptOutput.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *e, acceptOutput.MCMSTimelockProposals, false)
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

func EVMTransferOwnership(t *testing.T, e *cldf_deployment.Environment, selector uint64) {
	chain := e.BlockChains.EVMChains()[selector]
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range e.DataStore.Addresses().Filter() {
		if addrRef.Type == datastore.ContractType(common_utils.RBACTimelock) {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}
	mcmsInput := mcmsapi.TransferOwnershipInput{
		ChainInputs: []mcmsapi.TransferOwnershipPerChainInput{
			{
				ChainSelector: chain.Selector,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(evmfqops.ContractType),
						Version: evmfqops.Version,
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
