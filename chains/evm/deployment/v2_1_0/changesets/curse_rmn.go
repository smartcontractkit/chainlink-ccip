package changesets

import (
	"fmt"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// CurseRMNCfg is the configuration for the CurseRMN changeset.
// It curses the given subjects on an already-deployed RMN v2.1.0 contract.
type CurseRMNCfg struct {
	ChainSel uint64
	// RMN is an address ref used to look up the deployed RMN contract in the datastore.
	RMN      datastore.AddressRef
	Subjects [][16]byte
}

func (c CurseRMNCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// CurseRMN curses one or more subjects on an RMN v2.1.0 contract.
var CurseRMN = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.SeqCurseInput,
	evm.Chain,
	CurseRMNCfg,
]{
	Sequence: sequences.RmnCurse,
	ResolveInput: func(e cldf_deployment.Environment, cfg CurseRMNCfg) (sequences.SeqCurseInput, error) {
		rmnAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.RMN, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.SeqCurseInput{}, fmt.Errorf("failed to resolve RMN ref on chain with selector %d: %w", cfg.ChainSel, err)
		}
		return sequences.SeqCurseInput{
			RMNAddress: rmnAddr,
			Subjects:   cfg.Subjects,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[CurseRMNCfg],
})

// UncurseRMNCfg is the configuration for the UncurseRMN changeset.
// It uncurses the given subjects on an already-deployed RMN v2.1.0 contract.
type UncurseRMNCfg struct {
	ChainSel uint64
	// RMN is an address ref used to look up the deployed RMN contract in the datastore.
	RMN      datastore.AddressRef
	Subjects [][16]byte
}

func (c UncurseRMNCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// UncurseRMN uncurses one or more subjects on an RMN v2.1.0 contract.
var UncurseRMN = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.SeqUncurseInput,
	evm.Chain,
	UncurseRMNCfg,
]{
	Sequence: sequences.RmnUncurse,
	ResolveInput: func(e cldf_deployment.Environment, cfg UncurseRMNCfg) (sequences.SeqUncurseInput, error) {
		rmnAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, cfg.RMN, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return sequences.SeqUncurseInput{}, fmt.Errorf("failed to resolve RMN ref on chain with selector %d: %w", cfg.ChainSel, err)
		}
		return sequences.SeqUncurseInput{
			RMNAddress: rmnAddr,
			Subjects:   cfg.Subjects,
		}, nil
	},
	ResolveDep: evm_sequences.ResolveEVMChainDep[UncurseRMNCfg],
})
