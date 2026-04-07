package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	v1_0_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokensapi.TokenAdapter = &EVMTokenBase{}

// EVMTokenBase provides version-agnostic EVM token adapter methods that are
// shared across all pool versions (v1.5.1, v1.6.0, v1.6.1, v2.0.0).
// It is also registered at v1.0.0 so callers that only need token deployment
// can obtain a valid adapter without importing a pool-version package.
type EVMTokenBase struct{}

func (a *EVMTokenBase) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", ref.Address)
	}

	return common.HexToAddress(ref.Address).Bytes(), nil
}

func (a *EVMTokenBase) DeployToken() *cldf_ops.Sequence[tokensapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return v1_0_0_seq.DeployToken
}

func (a *EVMTokenBase) DeployTokenVerify(e deployment.Environment, input tokensapi.DeployTokenInput) error {
	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
		ChainSelector: input.ChainSelector,
		Type:          datastore.ContractType(input.Type),
		Qualifier:     input.Symbol,
	}, input.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		e.OperationsBundle.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	if err := utils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if err := utils.ValidateEVMAddress(input.ExternalAdmin, "ExternalAdmin"); err != nil {
		return err
	}

	if input.Decimals > 18 {
		return fmt.Errorf("EVM tokens cannot have more than 18 decimals, got %d", input.Decimals)
	}

	if input.PreMint != nil && input.Supply != nil && *input.Supply != 0 && *input.PreMint > *input.Supply {
		return fmt.Errorf("pre-mint amount cannot be greater than max supply, got pre-mint %d and supply %d", *input.PreMint, *input.Supply)
	}

	return nil
}

func (a *EVMTokenBase) DeriveTokenPoolCounterpart(_ deployment.Environment, _ uint64, tokenPool []byte, _ []byte) ([]byte, error) {
	return tokenPool, nil
}

// UpdateAuthorities transfers token pool ownership to the timelock via MCMS.
// It creates a self-contained EVMTransferOwnershipAdapter within the sequence
// closure so it works correctly regardless of how the embedding struct is initialized.
func (a *EVMTokenBase) UpdateAuthorities() *cldf_ops.Sequence[tokensapi.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return cldf_ops.NewSequence(
		"evm-base:update-authorities",
		cciputils.Version_1_0_0,
		"Transfer token pool ownership to timelock on EVM chain",
		func(b cldf_ops.Bundle, e *deployment.Environment, input tokensapi.UpdateAuthoritiesInput) (sequences.OnChainOutput, error) {
			chain, ok := e.BlockChains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			timelockRef, err := datastore_utils.FindAndFormatRef(
				e.DataStore,
				datastore.AddressRef{
					Type:          datastore.ContractType(cciputils.RBACTimelock),
					ChainSelector: chain.Selector,
					Qualifier:     cciputils.CLLQualifier,
				},
				chain.Selector,
				datastore_utils.FullRef,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find timelock address for chain %d: %w", input.ChainSelector, err)
			}

			adapter := &EVMTransferOwnershipAdapter{}
			if err := adapter.InitializeTimelockAddress(*e, mcms.Input{Qualifier: cciputils.CLLQualifier}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize timelock address for chain %d: %w", input.ChainSelector, err)
			}

			ownershipInput := deployops.TransferOwnershipPerChainInput{
				ChainSelector: chain.Selector,
				CurrentOwner:  chain.DeployerKey.From.Hex(),
				ProposedOwner: timelockRef.Address,
				ContractRef:   []datastore.AddressRef{input.TokenPoolRef},
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, e.BlockChains,
				adapter.SequenceTransferOwnershipViaMCMS(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership on chain %d: %w", input.ChainSelector, err)
			}

			result, err = sequences.RunAndMergeSequence(b, e.BlockChains,
				adapter.SequenceAcceptOwnership(), ownershipInput, result)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership on chain %d: %w", input.ChainSelector, err)
			}

			return result, nil
		})
}

func (a *EVMTokenBase) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

// Pool-specific stubs -- these are overridden by per-version adapters (v1.5.1, v1.6.1, v2.0.0).
// EVMTokenBase is registered at v1.0.0 so callers that only need token deployment (DeployToken,
// DeployTokenVerify) can obtain a valid adapter without importing a pool-version package.

func (a *EVMTokenBase) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) DeriveTokenAddress(_ deployment.Environment, _ uint64, _ datastore.AddressRef) ([]byte, error) {
	return nil, fmt.Errorf("DeriveTokenAddress is not implemented on EVMTokenBase; use a pool-version adapter")
}

func (a *EVMTokenBase) DeriveTokenDecimals(_ deployment.Environment, _ uint64, _ datastore.AddressRef, _ []byte) (uint8, error) {
	return 0, fmt.Errorf("DeriveTokenDecimals is not implemented on EVMTokenBase; use a pool-version adapter")
}

func (a *EVMTokenBase) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokensapi.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *EVMTokenBase) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
