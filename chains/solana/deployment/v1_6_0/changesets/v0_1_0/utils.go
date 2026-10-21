package v0_1_0

import (
	"fmt"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	pdasol "github.com/smartcontractkit/cld-changesets/pkg/family/solana"
	"github.com/smartcontractkit/mcms"
	"github.com/smartcontractkit/mcms/sdk"
	mcmsSolana "github.com/smartcontractkit/mcms/sdk/solana"
	mcmsTypes "github.com/smartcontractkit/mcms/types"

	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldfproposalutils "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/mcms/proposalutils"

	proposeutils "github.com/smartcontractkit/cld-changesets/legacy/mcms/proposeutils"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

func ValidateMCMSConfigSolana(
	e cldf.Environment,
	mcms *cldfproposalutils.TimelockConfig,
	chain cldf_solana.Chain,
	chainState solanastateview.CCIPChainState,
	tokenAddress solana.PublicKey,
	tokenPoolMetadata string,
	contractsToValidate map[cldf.ContractType]bool) error {
	if mcms != nil {
		if err := solanastateview.ValidateSolanaTimelockConfig(e, chain.Selector, mcms); err != nil {
			return fmt.Errorf("failed to validate MCMS config: %w", err)
		}
	}
	if contractsToValidate[shared.FeeQuoter] {
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.FeeQuoter, shared.FeeQuoter, tokenAddress); err != nil {
			return fmt.Errorf("failed to validate ownership for fee quoter: %w", err)
		}
	}
	if contractsToValidate[shared.Router] {
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.Router, shared.Router, tokenAddress); err != nil {
			return fmt.Errorf("failed to validate ownership for router: %w", err)
		}
	}
	if contractsToValidate[shared.OffRamp] {
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.OffRamp, shared.OffRamp, tokenAddress); err != nil {
			return fmt.Errorf("failed to validate ownership for off ramp: %w", err)
		}
	}
	if contractsToValidate[shared.RMNRemote] {
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.RMNRemote, shared.RMNRemote, tokenAddress); err != nil {
			return fmt.Errorf("failed to validate ownership for rmnremote: %w", err)
		}
	}
	if !tokenAddress.IsZero() {
		metadata := shared.CLLMetadata
		if tokenPoolMetadata != "" {
			metadata = tokenPoolMetadata
		}
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.BurnMintTokenPools[metadata], shared.BurnMintTokenPool, tokenAddress); contractsToValidate[shared.BurnMintTokenPool] && err != nil {
			return fmt.Errorf("failed to validate ownership for burnmint: %w", err)
		}
		if err := solanastateview.ValidateOwnershipSolana(&e, chain, mcms != nil, chainState.LockReleaseTokenPools[metadata], shared.LockReleaseTokenPool, tokenAddress); contractsToValidate[shared.LockReleaseTokenPool] && err != nil {
			return fmt.Errorf("failed to validate ownership for lockrelease: %w", err)
		}
	}

	return nil
}

func buildProposalCommon(
	e cldf.Environment,
	chainSelector uint64,
	description string,
	minDelay time.Duration,
	batches []mcmsTypes.BatchOperation) (*mcms.TimelockProposal, error) {
	timelocks := map[uint64]string{}
	proposers := map[uint64]string{}
	inspectors := map[uint64]sdk.Inspector{}

	chain := e.BlockChains.SolanaChains()[chainSelector]
	mcmState, err := solanastateview.LoadMCMSWithTimelockFromDataStore(&e, chain, shared.DefaultMCMSQualifier)
	if err != nil {
		return nil, fmt.Errorf("failed to load mcms with timelock state: %w", err)
	}

	timelocks[chainSelector] = mcmsSolana.ContractAddress(
		mcmState.TimelockProgram,
		mcmsSolana.PDASeed(mcmState.TimelockSeed),
	)
	proposers[chainSelector] = mcmsSolana.ContractAddress(mcmState.McmProgram, mcmsSolana.PDASeed(mcmState.ProposerMcmSeed))
	inspectors[chainSelector] = mcmsSolana.NewInspector(chain.Client)

	proposal, err := proposeutils.BuildProposalFromBatchesV2(
		e,
		timelocks,
		proposers,
		inspectors,
		batches,
		description,
		cldfproposalutils.TimelockConfig{MinDelay: minDelay})
	if err != nil {
		return nil, fmt.Errorf("failed to build proposal: %w", err)
	}
	return proposal, nil
}

func BuildProposalsForTxns(
	e cldf.Environment,
	chainSelector uint64,
	description string,
	minDelay time.Duration,
	txns []mcmsTypes.Transaction) (*mcms.TimelockProposal, error) {
	batches := []mcmsTypes.BatchOperation{
		{
			ChainSelector: mcmsTypes.ChainSelector(chainSelector),
			Transactions:  txns,
		},
	}
	return buildProposalCommon(e, chainSelector, description, minDelay, batches)
}

func BuildMCMSTxn(ixn solana.Instruction, programID string, contractType cldf.ContractType) (*mcmsTypes.Transaction, error) {
	data, err := ixn.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to extract data: %w", err)
	}
	for _, account := range ixn.Accounts() {
		if account.IsSigner {
			account.IsSigner = false
		}
	}
	tx, err := mcmsSolana.NewTransaction(
		programID,
		data,
		big.NewInt(0),        // e.g. value
		ixn.Accounts(),       // pass along needed accounts
		string(contractType), // some string identifying the target
		[]string{},           // any relevant metadata
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	return &tx, nil
}

func FetchTimelockSigner(e cldf.Environment, chainSelector uint64) (solana.PublicKey, error) {
	mcmState, err := solanastateview.LoadMCMSWithTimelockFromDataStore(&e, e.BlockChains.SolanaChains()[chainSelector], shared.DefaultMCMSQualifier)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to load mcm state: %w", err)
	}
	timelockSignerPDA := pdasol.GetTimelockSignerPDA(mcmState.TimelockProgram, mcmState.TimelockSeed)
	return timelockSignerPDA, nil
}

// returns the authority for the given ixn based on program ownership
func GetAuthorityForIxn(
	e *cldf.Environment,
	chain cldf_solana.Chain,
	chainState solanastateview.CCIPChainState,
	contractType cldf.ContractType,
	tokenAddress solana.PublicKey, // used for burnmint and lockrelease
	tokenMetadata string, // used for burnmint and lockrelease
) solana.PublicKey {
	timelockSigner, err := FetchTimelockSigner(*e, chain.Selector)
	if err != nil {
		return chain.DeployerKey.PublicKey()
	}
	if solanastateview.IsSolanaProgramOwnedByTimelock(e, chain, chainState, contractType, tokenAddress, tokenMetadata) {
		return timelockSigner
	}
	return chain.DeployerKey.PublicKey()
}
