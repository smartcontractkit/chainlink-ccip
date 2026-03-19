package adapters

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	cv "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	dsutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type EVMAggregatorConfigAdapter struct{}

var _ adapters.AggregatorConfigAdapter = (*EVMAggregatorConfigAdapter)(nil)

func (a *EVMAggregatorConfigAdapter) ScanCommitteeStates(ctx context.Context, env deployment.Environment, chainSelector uint64) ([]*adapters.CommitteeState, error) {
	refs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
		datastore.AddressRefByChainSelector(chainSelector),
	)

	if len(refs) == 0 {
		return nil, nil
	}

	evmChains := env.BlockChains.EVMChains()
	if evmChains == nil {
		return nil, fmt.Errorf("no EVM chains found in environment")
	}

	chain, ok := evmChains[chainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found in environment", chainSelector)
	}

	states := make([]*adapters.CommitteeState, 0, len(refs))
	for _, ref := range refs {
		addr := common.HexToAddress(ref.Address)
		contract, err := cv.NewCommitteeVerifier(addr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to bind CommitteeVerifier %s on chain %d: %w", ref.Address, chainSelector, err)
		}

		allConfigs, err := contract.GetAllSignatureConfigs(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("failed to get signature configs from %s on chain %d: %w", ref.Address, chainSelector, err)
		}

		sigConfigs := make([]adapters.SignatureConfig, 0, len(allConfigs))
		for _, cfg := range allConfigs {
			signers := make([]string, 0, len(cfg.Signers))
			for _, signer := range cfg.Signers {
				signers = append(signers, signer.Hex())
			}
			sigConfigs = append(sigConfigs, adapters.SignatureConfig{
				SourceChainSelector: cfg.SourceChainSelector,
				Signers:             signers,
				Threshold:           cfg.Threshold,
			})
		}

		states = append(states, &adapters.CommitteeState{
			Qualifier:        ref.Qualifier,
			ChainSelector:    chainSelector,
			Address:          ref.Address,
			SignatureConfigs: sigConfigs,
		})
	}

	return states, nil
}

func (a *EVMAggregatorConfigAdapter) ResolveVerifierAddress(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	return dsutils.FindAndFormatFirstRef(ds, chainSelector,
		func(r datastore.AddressRef) (string, error) { return r.Address, nil },
		datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
			Qualifier: qualifier,
		},
		datastore.AddressRef{
			Type:      datastore.ContractType(committee_verifier.ContractType),
			Qualifier: qualifier,
		},
	)
}
