package adapters

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	solrouterops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type CurseAdapter struct {
	rmnRemoteConfigPDA solana.PublicKey
	rmnRemoteCursesPDA solana.PublicKey
	rmnRemote          solana.PublicKey
	routerProgramAddr  solana.PublicKey
}

func (ca *CurseAdapter) DeriveCurseAdapterVersion(e cldf.Environment, selector uint64) (*semver.Version, error) {
	return rmnremoteops.Version, nil
}

func (ca *CurseAdapter) ListConnectedChains(e cldf.Environment, selector uint64) ([]uint64, error) {
	chain, ok := e.BlockChains.SolanaChains()[selector]
	if !ok {
		return nil, fmt.Errorf("solana chain with selector %d not found in environment", selector)
	}
	if ca.routerProgramAddr == (solana.PublicKey{}) {
		return nil, fmt.Errorf("router program address not initialized")
	}
	refs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(solrouterops.DestChainType)),
		datastore.AddressRefByVersion(solrouterops.Version),
	)
	var connectedChains []uint64
	for _, ref := range refs {
		var destChain uint64
		err := json.Unmarshal([]byte(ref.Qualifier), &destChain)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal dest chain from qualifier %s: %w", ref.Qualifier, err)
		}
		connectedChains = append(connectedChains, destChain)
	}

	// for each connected chain, verify if it is enabled in router
	for _, destChain := range connectedChains {
		routerDestChainPDA, err := state.FindDestChainStatePDA(destChain, ca.routerProgramAddr)
		if err != nil {
			return nil, fmt.Errorf("failed to find dest chain state pda for remote chain %d: %w", destChain, err)
		}
		var destChainStateAccount ccip_router.DestChainState
		err = chain.GetAccountDataBorshInto(e.GetContext(), routerDestChainPDA, &destChainStateAccount)
		if err != nil {
			return nil, fmt.Errorf("failed to get dest chain state account for remote chain %d: %w", destChain, err)
		}
	}
	return connectedChains, nil
}

func NewCurseAdapter() *CurseAdapter {
	return &CurseAdapter{}
}

func (ca *CurseAdapter) Initialize(e cldf.Environment, selector uint64) error {
	_, exists := e.BlockChains.SolanaChains()[selector]
	if !exists {
		return fmt.Errorf("solana chain with selector %d not found in environment", selector)
	}
	rmnRemoteRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			ChainSelector: selector,
			Type:          datastore.ContractType(rmnremoteops.ContractType),
			Version:       rmnremoteops.Version,
		},
		selector,
		utils.ToAddress,
	)
	if err != nil {
		return fmt.Errorf("failed to find RMNRemote contract: %w", err)
	}
	ca.rmnRemote = rmnRemoteRef
	ca.rmnRemoteConfigPDA, _, err = state.FindRMNRemoteConfigPDA(ca.rmnRemote)
	if err != nil {
		return fmt.Errorf("failed to find RMNRemoteConfig PDA: %w", err)
	}
	ca.rmnRemoteCursesPDA, _, err = state.FindRMNRemoteCursesPDA(ca.rmnRemote)
	if err != nil {
		return fmt.Errorf("failed to find RMNRemoteCurses PDA: %w", err)
	}
	routerRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			ChainSelector: selector,
			Type:          datastore.ContractType(solrouterops.ContractType),
			Version:       solrouterops.Version,
		},
		selector,
		utils.ToAddress,
	)
	if err != nil {
		return fmt.Errorf("failed to find Router contract: %w", err)
	}
	ca.routerProgramAddr = routerRef
	return nil
}

func (ca *CurseAdapter) IsSubjectCursedOnChain(e cldf.Environment, selector uint64, subject api.Subject) (bool, error) {
	chain := e.BlockChains.SolanaChains()[selector]
	curseSubject := rmn_remote.CurseSubject{
		Value: subject,
	}
	return rmnremoteops.IsSubjectCursed(chain, ca.rmnRemote, curseSubject)
}

func (ca *CurseAdapter) IsChainConnectedToTargetChain(e cldf.Environment, selector uint64, targetSel uint64) (bool, error) {
	// get offRamp address
	offRampAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(offramp.ContractType),
		Version:       offramp.Version,
	}, selector, utils.ToAddress)
	if err != nil {
		return false, err
	}
	pda, _, err := state.FindOfframpSourceChainPDA(targetSel, offRampAddr)
	if err != nil {
		return false, fmt.Errorf("failed to find offramp source chain pda: %w", err)
	}

	var chainStateAccount ccip_offramp.SourceChain
	if err = e.BlockChains.SolanaChains()[selector].GetAccountDataBorshInto(e.GetContext(), pda, &chainStateAccount); err != nil {
		return false, nil
	}

	return chainStateAccount.Config.IsEnabled, nil
}

func (ca *CurseAdapter) IsCurseEnabledForChain(e cldf.Environment, selector uint64) (bool, error) {
	return ca.rmnRemote != solana.PublicKey{}, nil
}

func (ca *CurseAdapter) SubjectToSelector(subject api.Subject) (uint64, error) {
	if subject == api.GlobalCurseSubject() {
		return 0, nil
	}

	return binary.LittleEndian.Uint64(subject[:]), nil
}

func (ca *CurseAdapter) SelectorToSubject(selector uint64) api.Subject {
	var b api.Subject
	binary.LittleEndian.PutUint64(b[0:], selector)
	return b
}

func (ca *CurseAdapter) Curse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"curse_rmn_remote",
		semver.MustParse("1.0.0"),
		"Cursing subjects with RMNRemote",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			for _, subject := range in.Subjects {
				// Solana uses little endian to encode the subject so we expect the last 8 bytes to be 0
				if !bytes.Equal(subject[8:], []byte{0, 0, 0, 0, 0, 0, 0, 0}) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid subject format for Solana RMNRemote curse: %v", subject)
				}
			}
			opInput := rmnremoteops.CurseInput{
				RMNRemote:          ca.rmnRemote,
				RMNRemoteCursePDA:  ca.rmnRemoteCursesPDA,
				RMNRemoteConfigPDA: ca.rmnRemoteConfigPDA,
				Subjects:           in.Subjects,
			}
			chain, exists := chains.SolanaChains()[in.ChainSelector]
			if !exists {
				return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not found", in.ChainSelector)
			}
			opOutput, err := cldf_ops.ExecuteOperation(b, rmnremoteops.Curse, chain, opInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to curse with RMNRemote at %s on chain %d: %w", ca.rmnRemote.String(), in.ChainSelector, err)
			}
			return opOutput.Output, nil
		})
}

func (ca *CurseAdapter) Uncurse() *cldf_ops.Sequence[api.CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"curse_rmn_remote",
		semver.MustParse("1.0.0"),
		"Cursing subjects with RMNRemote",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.CurseInput) (output sequences.OnChainOutput, err error) {
			for _, subject := range in.Subjects {
				// Solana uses little endian to encode the subject so we expect the last 8 bytes to be 0
				if !bytes.Equal(subject[8:], []byte{0, 0, 0, 0, 0, 0, 0, 0}) {
					return sequences.OnChainOutput{}, fmt.Errorf("invalid subject format for Solana RMNRemote curse: %v", subject)
				}
			}
			opInput := rmnremoteops.CurseInput{
				RMNRemote:          ca.rmnRemote,
				RMNRemoteCursePDA:  ca.rmnRemoteCursesPDA,
				RMNRemoteConfigPDA: ca.rmnRemoteConfigPDA,
				Subjects:           in.Subjects,
			}
			chain, exists := chains.SolanaChains()[in.ChainSelector]
			if !exists {
				return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not found", in.ChainSelector)
			}
			opOutput, err := cldf_ops.ExecuteOperation(b, rmnremoteops.Uncurse, chain, opInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to uncurse with RMNRemote at %s on chain %d: %w", ca.rmnRemote.String(), in.ChainSelector, err)
			}
			return opOutput.Output, nil
		})
}
