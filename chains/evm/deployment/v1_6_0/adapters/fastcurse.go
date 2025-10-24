package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	rmnsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type CurseAdapter struct {
	rmnAddressCache    map[uint64]common.Address
	routerAddressCache map[uint64]common.Address
}

func NewCurseAdapter() *CurseAdapter {
	return &CurseAdapter{}
}

func (ca *CurseAdapter) Initialize(e cldf.Environment) error {
	ca.rmnAddressCache = make(map[uint64]common.Address)
	ca.routerAddressCache = make(map[uint64]common.Address)

	for _, chain := range e.BlockChains.EVMChains() {
		rmnAddr, err := rmnAddressOnChain(e, chain.Selector)
		if err != nil {
			return fmt.Errorf("failed to find RMN address on chain %d: %w", chain.Selector, err)
		}
		ca.rmnAddressCache[chain.Selector] = rmnAddr

		routerAddr, err := routerAddressOnChain(e, chain.Selector)
		if err != nil {
			return fmt.Errorf("failed to find router address on chain %d: %w", chain.Selector, err)
		}
		ca.routerAddressCache[chain.Selector] = routerAddr
	}
	return nil
}

func (ca *CurseAdapter) IsSubjectCursedOnChain(e cldf.Environment, selector uint64, subject api.Subject) (bool, error) {
	// locate rmn address on chain
	rmnAddr, ok := ca.rmnAddressCache[selector]
	if !ok {
		return false, fmt.Errorf("no RMN address cached for chain %d", selector)
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return false, fmt.Errorf("no EVM chain found for selector %d", selector)
	}

	isCursedRep, err := cldf_ops.ExecuteOperation(e.OperationsBundle, ops.IsCursed, chain, contract.FunctionInput[api.Subject]{
		ChainSelector: chain.Selector,
		Address:       rmnAddr,
		Args:          subject,
	})
	return isCursedRep.Output, err
}

func (ca *CurseAdapter) IsChainConnectedToTargetChain(e cldf.Environment, selector uint64, targetSel uint64) (bool, error) {
	// locate rmn address on chain
	routerAddr, ok := ca.routerAddressCache[selector]
	if !ok {
		return false, fmt.Errorf("no router address cached for chain %d", selector)
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return false, fmt.Errorf("no EVM chain found for selector %d", selector)
	}
	isChainSupportedRep, err := cldf_ops.ExecuteOperation(e.OperationsBundle, routerops.IsChainSupported, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       routerAddr,
		Args:          targetSel,
	})
	return isChainSupportedRep.Output, err
}

func (ca *CurseAdapter) IsCurseEnabledForChain(e cldf.Environment, selector uint64) (bool, error) {
	// locate rmn address on chain
	_, ok := ca.rmnAddressCache[selector]
	if !ok {
		return false, fmt.Errorf("no RMN address cached for chain %d", selector)
	}
	return true, nil
}

func (ca *CurseAdapter) SubjectToSelector(subject api.Subject) (uint64, error) {
	return api.GenericSubjectToSelector(subject)
}

func (ca *CurseAdapter) SelectorToSubject(selector uint64) api.Subject {
	return api.GenericSelectorToSubject(selector)
}

func (ca *CurseAdapter) Curse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"curse_rmn_remote",
		semver.MustParse("1.0.0"),
		"Cursing subjects with RMNRemote",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			rmnAddr, ok := ca.rmnAddressCache[chain.Selector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no RMN address cached for chain %d", chain.Selector)
			}
			SeqCurseInput := rmnsequences.SeqCurseInput{
				CurseInput: in,
				Addr:       rmnAddr,
			}
			seqOutput, err := cldf_ops.ExecuteSequence(b, rmnsequences.SeqCurse, chain, SeqCurseInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects on chain %d: %w", chain.Selector, err)
			}
			output.BatchOps = append(output.BatchOps, seqOutput.Output.BatchOps...)
			return output, nil
		})
}

func (ca *CurseAdapter) Uncurse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"uncurse_rmn_remote",
		semver.MustParse("1.0.0"),
		"Uncursing subjects with RMNRemote",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			rmnAddr, ok := ca.rmnAddressCache[chain.Selector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no RMN address cached for chain %d", chain.Selector)
			}
			SeqCurseInput := rmnsequences.SeqCurseInput{
				CurseInput: in,
				Addr:       rmnAddr,
			}
			seqOutput, err := cldf_ops.ExecuteSequence(b, rmnsequences.SeqUncurse, chain, SeqCurseInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects on chain %d: %w", chain.Selector, err)
			}
			output.BatchOps = append(output.BatchOps, seqOutput.Output.BatchOps...)
			return output, nil
		})
}

func routerAddressOnChain(e cldf.Environment, selector uint64) (common.Address, error) {
	routerRef := datastore.AddressRef{
		Type:    datastore.ContractType(routerops.ContractType),
		Version: semver.MustParse("1.2.0"),
	}
	routerAddrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, routerRef, selector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to resolve router ref on chain with selector %d: %w", selector, err)
	}

	return routerAddrRef, nil
}

func rmnAddressOnChain(e cldf.Environment, selector uint64) (common.Address, error) {
	rmnRef := datastore.AddressRef{
		Type:    datastore.ContractType(ops.ContractType),
		Version: semver.MustParse("1.6.0"),
	}
	rmnAddrRef, err := datastore_utils.FindAndFormatRef(e.DataStore, rmnRef, selector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to resolve RMN ref on chain with selector %d: %w", selector, err)
	}
	return rmnAddrRef, nil
}
