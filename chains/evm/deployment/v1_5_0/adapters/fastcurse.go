package adapters

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	rmnsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
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

func (ca *CurseAdapter) Initialize(e cldf.Environment, selector uint64) error {
	if ca.rmnAddressCache == nil {
		ca.rmnAddressCache = make(map[uint64]common.Address)
	}
	if ca.routerAddressCache == nil {
		ca.routerAddressCache = make(map[uint64]common.Address)
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return fmt.Errorf("no EVM chain found for selector %d", selector)
	}
	if _, exists := ca.rmnAddressCache[chain.Selector]; !exists {
		rmnAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(ops.ContractType),
			Version: semver.MustParse("1.5.0"),
		}, chain.ChainSelector(), evmds.ToEVMAddress)
		if err != nil {
			return err
		}
		ca.rmnAddressCache[chain.Selector] = rmnAddr
	}
	if _, exists := ca.routerAddressCache[chain.Selector]; !exists {
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
	rmnC, err := rmn_contract.NewRMNContract(rmnAddr, chain.Client)
	if err != nil {
		return false, fmt.Errorf("failed to instantiate RMN contract at %s on chain %d: %w", rmnAddr.String(), chain.Selector, err)
	}
	// check if chain is globally cursed
	return rmnC.IsCursed(&bind.CallOpts{
		Context: e.GetContext(),
	}, subject)
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
	routerC, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return false, fmt.Errorf("failed to instantiate router contract at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	return routerC.IsChainSupported(&bind.CallOpts{
		Context: e.GetContext(),
	}, targetSel)
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
			curseID, err := generateCurseID(cfgDetailsOp.Output.Version, in.Subjects)
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

func (ca *CurseAdapter) ListConnectedChains(e cldf.Environment, selector uint64) ([]uint64, error) {
	routerAddr, ok := ca.routerAddressCache[selector]
	if !ok {
		return nil, fmt.Errorf("no router address cached for chain %d", selector)
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", selector)
	}
	// get all offRamps from router to find connected chains
	routerC, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate router contract at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	offRamps, err := routerC.GetOffRamps(&bind.CallOpts{
		Context: e.GetContext(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramps from router at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	connectedChains := make([]uint64, 0)
	for _, offRamp := range offRamps {
		if offRamp.OffRamp == (common.Address{}) {
			continue // skip uninitialized off-ramps
		}
		// if chain is non-evm, skip ( TODO: support non-evm chains later)
		_, exists := e.BlockChains.EVMChains()[offRamp.SourceChainSelector]
		if !exists {
			continue
		}
		connectedChains = append(connectedChains, offRamp.SourceChainSelector)
	}
	return connectedChains, nil
}

func (ca *CurseAdapter) DeriveCurseAdapterVersion(e cldf.Environment, selector uint64) (*semver.Version, error) {
	// fetch RMNProxy address on chain
	rmnProxyRef := datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		ChainSelector: selector,
	}
	rmnProxyAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, rmnProxyRef, selector, evmds.ToEVMAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve RMNProxy ref on chain with selector %d: %w", selector, err)
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", selector)
	}
	rmnProxyC, err := rmn_proxy_contract.NewRMNProxy(rmnProxyAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate RMNProxy contract at %s on chain %d: %w", rmnProxyAddr.String(), chain.Selector, err)
	}
	rmnAddr, err := rmnProxyC.GetARM(&bind.CallOpts{
		Context: e.GetContext(),
	})
	if err != nil {
		return nil, err
	}
	_, version, err := utils.TypeAndVersion(rmnAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to get type and version from RMN at %s on chain %d: %w", rmnAddr.String(), chain.Selector, err)
	}
	return version, nil
}

func generateCurseID(cfgVersion uint32, subjects [][16]byte) ([16]byte, error) {
	var out [16]byte

	h := sha256.New()

	// Include cfgVersion
	err := binary.Write(h, binary.BigEndian, cfgVersion)
	if err != nil {
		return [16]byte{}, err
	}

	// Include all subjects in deterministic order
	for _, s := range subjects {
		h.Write(s[:])
	}

	sum := h.Sum(nil)

	// Copy first 16 bytes of hash into output
	copy(out[:], sum[:16])

	binary.BigEndian.PutUint32(out[0:4], cfgVersion)

	return out, nil
}
