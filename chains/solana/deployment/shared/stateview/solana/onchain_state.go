package solana

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	solstate "github.com/smartcontractkit/cld-changesets/legacy/pkg/family/solana"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

// CCIPOnChainState is the Solana part of chainlink's multi-family
// chainlink/deployment/ccip/shared/stateview.CCIPOnChainState. The Solana changesets only read
// SolChains, so the EVM and other families' state, which would pull in their bindings, is left
// out.
type CCIPOnChainState struct {
	SolChains map[uint64]CCIPChainState
}

// LoadOnchainStateSolana loads the CCIP state of every Solana chain in the environment from
// the environment datastore.
//
// Reproduced from chainlink/deployment/ccip/shared/stateview.LoadOnchainStateSolana.
func LoadOnchainStateSolana(e cldf.Environment) (CCIPOnChainState, error) {
	state := CCIPOnChainState{
		SolChains: make(map[uint64]CCIPChainState),
	}
	if e.DataStore == nil {
		return state, errors.New("datastore not available for solana state loading")
	}
	allRefs, err := e.DataStore.Addresses().Fetch()
	if err != nil {
		return state, fmt.Errorf("failed to fetch address refs from datastore: %w", err)
	}
	for chainSelector, chain := range e.BlockChains.SolanaChains() {
		var chainRefs []datastore.AddressRef
		for _, ref := range allRefs {
			if ref.ChainSelector == chainSelector {
				chainRefs = append(chainRefs, ref)
			}
		}
		chainState, err := LoadChainStateSolana(chain, chainRefs)
		if err != nil {
			return state, err
		}
		state.SolChains[chainSelector] = chainState
	}
	return state, nil
}

// SupportedChains returns the Solana chains in the state. Chainlink's version also lists the
// other families' chains; callers here only look up Solana selectors, which both return.
func (c CCIPOnChainState) SupportedChains() map[uint64]struct{} {
	chains := make(map[uint64]struct{})
	for chain := range c.SolChains {
		chains[chain] = struct{}{}
	}
	return chains
}

// SolanaMCMSStateByChain returns the default-qualifier MCMS bundle of every Solana chain in the
// environment, for proposal aggregation.
//
// Reproduced from chainlink/deployment/ccip/shared/stateview.CCIPOnChainState.SolanaMCMSStateByChain,
// including its early return: the first chain whose bundle can't be loaded ends the scan.
func (c CCIPOnChainState) SolanaMCMSStateByChain(e cldf.Environment) map[uint64]solstate.MCMSWithTimelockState {
	mcmsStateByChain := make(map[uint64]solstate.MCMSWithTimelockState)
	for chainSelector := range e.BlockChains.SolanaChains() {
		if e.DataStore == nil {
			return mcmsStateByChain
		}
		refs, err := e.DataStore.Addresses().Fetch()
		if err != nil {
			return mcmsStateByChain
		}
		bundle, err := shared.MCMSBundleRefs(refs, chainSelector, shared.DefaultMCMSQualifier)
		if err != nil || len(bundle) == 0 {
			return mcmsStateByChain
		}
		mcmState, err := solstate.MaybeLoadMCMSWithTimelockChainStateV2(bundle)
		if err != nil {
			return mcmsStateByChain
		}
		mcmsStateByChain[chainSelector] = *mcmState
	}
	return mcmsStateByChain
}
