package ccv

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	cv "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// OnChainCommitteeState represents the actual on-chain state of a committee verifier.
type OnChainCommitteeState struct {
	// Qualifier is the committee qualifier from the datastore.
	Qualifier string
	// ChainSelector is the chain where this verifier is deployed.
	ChainSelector uint64
	// Address is the CommitteeVerifier contract address (ContractType, not ResolverType).
	Address string
	// SignatureConfigs contains the on-chain signature configurations per source chain.
	SignatureConfigs []OnChainSignatureConfig
}

// OnChainSignatureConfig represents the signature configuration read from the contract.
type OnChainSignatureConfig struct {
	// SourceChainSelector is the source chain this config applies to.
	SourceChainSelector uint64
	// Signers are the authorized signer addresses.
	Signers []common.Address
	// Threshold is the minimum number of signatures required.
	Threshold uint8
}

// OnChainTopology represents the complete on-chain state of all committee verifiers.
type OnChainTopology struct {
	// Committees maps qualifier to committee states.
	Committees map[string][]*OnChainCommitteeState
}

// ScanOnChainTopology scans all CommitteeVerifier contracts from the datastore
// and reads their on-chain state (signers and thresholds).
// Scanning is parallelized for performance with large numbers of chains.
func ScanOnChainTopology(
	ctx context.Context,
	env deployment.Environment,
) (*OnChainTopology, error) {
	ds := env.DataStore

	// Find all CommitteeVerifier addresses in the datastore
	// Note: We use ContractType (CommitteeVerifier), not ResolverType (CommitteeVerifierResolver)
	// because the actual verifier contract has GetAllSignatureConfigs() method
	refs := ds.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
	)

	if len(refs) == 0 {
		return nil, fmt.Errorf("no CommitteeVerifier contracts found in datastore")
	}

	var mu sync.Mutex
	topology := &OnChainTopology{
		Committees: make(map[string][]*OnChainCommitteeState),
	}

	g, ctx := errgroup.WithContext(ctx)

	for _, ref := range refs {
		chainFamily, err := chainsel.GetSelectorFamily(ref.ChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for selector %d: %w", ref.ChainSelector, err)
		}
		switch chainFamily {
		case chainsel.FamilyEVM:
			g.Go(func() error {
				state, err := scanCommitteeVerifier(ctx, env, ref)
				if err != nil {
					return fmt.Errorf("failed to scan committee verifier %s on chain %d: %w",
						ref.Address, ref.ChainSelector, err)
				}

				mu.Lock()
				topology.Committees[state.Qualifier] = append(topology.Committees[state.Qualifier], state)
				mu.Unlock()
				return nil
			})
		default:
			// Skip other chain families for now
			env.Logger.Warnw("skipping CommitteeVerifier scan on unsupported chain family", "family", chainFamily, "selector", ref.ChainSelector)
			continue
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return topology, nil
}

// scanCommitteeVerifier reads the on-chain state of a single CommitteeVerifier contract.
func scanCommitteeVerifier(
	ctx context.Context,
	env deployment.Environment,
	ref datastore.AddressRef,
) (*OnChainCommitteeState, error) {
	evmChains := env.BlockChains.EVMChains()
	if evmChains == nil {
		return nil, fmt.Errorf("no EVM chains found in environment")
	}

	chain, ok := evmChains[ref.ChainSelector]
	if !ok {
		return nil, fmt.Errorf("chain %d not found in environment", ref.ChainSelector)
	}

	addr := common.HexToAddress(ref.Address)
	contract, err := cv.NewCommitteeVerifier(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CommitteeVerifier contract: %w", err)
	}

	callOpts := &bind.CallOpts{Context: ctx}

	// Get all signature configs from the contract
	allConfigs, err := contract.GetAllSignatureConfigs(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get all signature configs: %w", err)
	}

	sigConfigs := make([]OnChainSignatureConfig, 0, len(allConfigs))
	for _, cfg := range allConfigs {
		sigConfigs = append(sigConfigs, OnChainSignatureConfig{
			SourceChainSelector: cfg.SourceChainSelector,
			Signers:             cfg.Signers,
			Threshold:           cfg.Threshold,
		})
	}

	return &OnChainCommitteeState{
		Qualifier:        ref.Qualifier,
		ChainSelector:    ref.ChainSelector,
		Address:          ref.Address,
		SignatureConfigs: sigConfigs,
	}, nil
}
