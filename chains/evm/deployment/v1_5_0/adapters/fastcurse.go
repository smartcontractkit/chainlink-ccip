package adapters

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	rmnsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	// based on RMN contract code
	// The curse vote address used in an OwnerUnvoteToCurseRequest to lift a curse
	Lift_Curse_Vote_Addr = utils.ZeroAddress
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
		rmnAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(ops.ContractType),
			Version: semver.MustParse("1.6.0"),
		}, chain.ChainSelector(), evmds.ToEVMAddress)
		if err != nil {
			return err
		}
		ca.rmnAddressCache[chain.Selector] = rmnAddr

		routerAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(routerops.ContractType),
			Version: semver.MustParse("1.2.0"),
		}, chain.ChainSelector(), evmds.ToEVMAddress)
		if err != nil {
			return err
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

func (ca *CurseAdapter) Curse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"curse_rmn",
		semver.MustParse("1.0.0"),
		"Cursing subjects with RMN",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			rmnAddr, ok := ca.rmnAddressCache[chain.Selector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no RMN address cached for chain %d", chain.Selector)
			}
			// form the curse ID
			// get config version
			cfgDetailsOp, err := cldf_ops.ExecuteOperation(b, ops.GetConfigDetails, chain, contract.FunctionInput[any]{
				Address:       rmnAddr,
				ChainSelector: chain.Selector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get config details for RMN at %s on chain %d: %w", rmnAddr.String(), chain.Selector, err)
			}
			curseID, err := generateCurseID(cfgDetailsOp.Output.Version)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to generate curse ID for RMN at %s on chain %d: %w", rmnAddr.String(), chain.Selector, err)
			}
			SeqCurseInput := rmnsequences.SeqCurseInput{
				CurseInput: in,
				Addr:       rmnAddr,
				CurseID:    curseID,
			}
			seqOutput, err := cldf_ops.ExecuteSequence(b, rmnsequences.SeqCurse, chain, SeqCurseInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects on chain %d: %w", chain.Selector, err)
			}
			output.BatchOps = append(output.BatchOps, seqOutput.Output.BatchOps...)
			return output, nil
		})
}

func (ca *CurseAdapter) SubjectToSelector(subject api.Subject) (uint64, error) {
	return api.GenericSubjectToSelector(subject)
}

func (ca *CurseAdapter) SelectorToSubject(selector uint64) api.Subject {
	return api.GenericSelectorToSubject(selector)
}

func (ca *CurseAdapter) Uncurse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"uncurse_rmn",
		semver.MustParse("1.0.0"),
		"Uncursing subjects with RMN",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			chain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			rmnAddr, ok := ca.rmnAddressCache[chain.Selector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("no RMN address cached for chain %d", chain.Selector)
			}
			// get the curse details
			requests := make([]rmn_contract.RMNOwnerUnvoteToCurseRequest, 0)
			for _, subject := range in.Subjects {
				curseProgressRep, err := cldf_ops.ExecuteOperation(b, ops.GetCurseProgress, chain, contract.FunctionInput[api.Subject]{
					Address:       rmnAddr,
					ChainSelector: chain.Selector,
					Args:          subject,
				})
				if err != nil {
					return sequences.OnChainOutput{},
						fmt.Errorf("failed to get curse progress for subject %x on chain %d: %w", subject, chain.Selector, err)
				}
				for i, cp := range curseProgressRep.Output.CurseVoteAddrs {
					requests = append(requests, rmn_contract.RMNOwnerUnvoteToCurseRequest{
						CurseVoteAddr: cp,
						Unit: rmn_contract.RMNUnvoteToCurseRequest{
							Subject:    subject,
							CursesHash: curseProgressRep.Output.CursesHashes[i],
						},
					})
				}
			}
			SeqCurseInput := rmnsequences.SeqUncurseInput{
				Addr:     rmnAddr,
				Requests: requests,
			}
			seqOutput, err := cldf_ops.ExecuteSequence(b, rmnsequences.SeqUncurse, chain, SeqCurseInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to curse subjects on chain %d: %w", chain.Selector, err)
			}
			output.BatchOps = append(output.BatchOps, seqOutput.Output.BatchOps...)
			return output, nil
		})
}

func generateCurseID(cfgVersion uint32) ([16]byte, error) {
	var out [16]byte

	_, err := rand.Read(out[4:])
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to generate random bytes for curse ID: %w", err)
	}
	binary.BigEndian.PutUint32(out[0:4], cfgVersion)

	return out, nil
}
