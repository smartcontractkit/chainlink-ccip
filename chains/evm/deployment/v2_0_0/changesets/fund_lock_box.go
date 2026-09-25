package changesets
package changesets

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

// LockBoxDeposit is a single deposit into a lockbox bucket.
type LockBoxDeposit struct {
	// RemoteChainSelector identifies the bucket to fund. Zero designates the unsiloed (shared)
	// bucket; any other value designates the silo for that remote chain.
	RemoteChainSelector uint64 `json:"remoteChainSelector" yaml:"remoteChainSelector"`
	// Amount is the amount to deposit, in raw base units.
	Amount *big.Int `json:"amount" yaml:"amount"`
}

// ChainLockBoxFunding groups every lockbox funding entry for one chain under that chain's selector.
// A chain legitimately holds several lockboxes , one per token, per silo group, or per remote
// chain and nesting them keeps the selector written once per chain.
type ChainLockBoxFunding struct {
	Selector uint64 `json:"selector" yaml:"selector"`
	// LockBoxAddress is the ERC20LockBox to fund. Named by address rather than by datastore
	// qualifier: qualifiers are a side-effect of whichever sequence deployed the lockbox
	// ("<poolQualifier>-silo(<minRemoteSelector>)" per silo group), so they are not a stable
	// addressing scheme. A wrong address fails on-chain when the read or write reverts.
	LockBoxAddress common.Address `json:"lockBoxAddress" yaml:"lockBoxAddress"`
	// TokenAddress is the token to deposit into the lockbox.
	TokenAddress common.Address `json:"tokenAddress" yaml:"tokenAddress"`
	// Deposits are the individual deposits to make. Each entry targets one bucket: a siloed bucket
	// (RemoteChainSelector set to the remote chain) or the unsiloed shared bucket
	// (RemoteChainSelector zero).
	Deposits []LockBoxDeposit `json:"deposits" yaml:"deposits"`
	// UsePlainTransfer, when true, transfers tokens directly to the lockbox via ERC20.transfer
	// instead of using the lockbox's deposit() function. This bypasses the Deposit event emission
	// and the per-bucket accounting. Use only as a break-glass option.
	UsePlainTransfer bool `json:"usePlainTransfer,omitempty" yaml:"usePlainTransfer,omitempty"`
}

type FundLockBoxCfg struct {
	// Version is not a field: this package is v2_0_0, so the target lockbox version is already
	// pinned by the changeset the operator picked.
	Input []ChainLockBoxFunding `json:"input" yaml:"input"`
}

// FundLockBox funds v2.0 ERC20LockBoxes directly, without a legacy pool to migrate from. It tops up
// a lockbox's siloed (per-chain-selector) and unsiloed (shared) buckets.
//
// This is an EVM-only changeset: lockboxes and siloed pools are EVM concepts, so there is no
// chain-agnostic interface to route through.
//
// All writes are proposal-only: the timelock must be an authorized caller on the lockbox to
// deposit, so the changeset authorizes it (idempotently) and then approves + deposits. The tokens
// are expected to be held by the timelock, since the deposit pulls from the caller's balance.
var FundLockBox = func(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) cldf_deployment.ChangeSetV2[cs_changesets.WithMCMS[FundLockBoxCfg]] {
	return cldf_deployment.CreateChangeSet(
		applyFundLockBox(mcmsRegistry),
		verifyFundLockBox,
	)
}

func verifyFundLockBox(
	e cldf_deployment.Environment,
	input cs_changesets.WithMCMS[FundLockBoxCfg],
) error {
	cfg := input.Cfg
	if len(cfg.Input) == 0 {
		return fmt.Errorf("at least one entry is required in input")
	}

	evmChains := e.BlockChains.EVMChains()
	seenChains := common_utils.NewSet[uint64]()

	for _, funding := range cfg.Input {
		sel := funding.Selector
		if _, err := chain_selectors.GetSelectorFamily(sel); err != nil {
			return fmt.Errorf("invalid chain selector %d: %w", sel, err)
		}
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		if seenChains.Add(sel) {
			return fmt.Errorf(
				"duplicate entry for chain %d: merge its lockboxes into a single entry", sel)
		}
		if funding.LockBoxAddress == (common.Address{}) {
			return fmt.Errorf("zero lockbox address for chain %d", sel)
		}
		if funding.TokenAddress == (common.Address{}) {
			return fmt.Errorf("zero token address for chain %d", sel)
		}
		if len(funding.Deposits) == 0 {
			return fmt.Errorf("no deposits listed for chain %d", sel)
		}

		seenBuckets := common_utils.NewSet[uint64]()
		for i, deposit := range funding.Deposits {
			if deposit.Amount == nil || deposit.Amount.Sign() <= 0 {
				return fmt.Errorf("deposits[%d] for chain %d: amount must be positive", i, sel)
			}
			if seenBuckets.Add(deposit.RemoteChainSelector) {
				return fmt.Errorf(
					"duplicate remote chain selector %d in deposits for chain %d: merge them into a single entry",
					deposit.RemoteChainSelector, sel)
			}
		}
	}

	return nil
}

func applyFundLockBox(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) func(cldf_deployment.Environment, cs_changesets.WithMCMS[FundLockBoxCfg]) (cldf_deployment.ChangesetOutput, error) {
	return func(
		e cldf_deployment.Environment,
		input cs_changesets.WithMCMS[FundLockBoxCfg],
	) (cldf_deployment.ChangesetOutput, error) {
		cfg := input.Cfg

		batchOps := make([]mcms_types.BatchOperation, 0, len(cfg.Input))
		reports := make([]cldf_ops.Report[any, any], 0, len(cfg.Input))

		mcmsReader, ok := mcmsRegistry.GetMCMSReader(chain_selectors.FamilyEVM)
		if !ok {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family '%s'", chain_selectors.FamilyEVM)
		}

		for _, funding := range cfg.Input {
			sel := funding.Selector

			timelockRef, err := mcmsReader.GetTimelockRef(e, sel, input.MCMS)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to get timelock address from MCMS config for chain %d: %w", sel, err)
			}

			deposits := make([]evm_tokens.LockBoxDeposit, 0, len(funding.Deposits))
			for _, deposit := range funding.Deposits {
				deposits = append(deposits, evm_tokens.LockBoxDeposit{
					RemoteChainSelector: deposit.RemoteChainSelector,
					Amount:              deposit.Amount,
				})
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle, evm_tokens.FundLockBox, e.BlockChains,
				evm_tokens.FundLockBoxInput{
					ChainSelector:    sel,
					LockBoxAddress:   funding.LockBoxAddress.Hex(),
					TokenAddress:     funding.TokenAddress.Hex(),
					TimelockAddress:  timelockRef.Address,
					Deposits:         deposits,
					UsePlainTransfer: funding.UsePlainTransfer,
				},
			)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to fund lockbox %s on chain %d: %w", funding.LockBoxAddress, sel, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return cs_changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}
