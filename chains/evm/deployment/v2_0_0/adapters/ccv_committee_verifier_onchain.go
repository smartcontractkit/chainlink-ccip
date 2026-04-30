package adapters

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	cv "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	ccvadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

// EVMCCVCommitteeVerifierOnchainAdapter implements ccvadapters.CommitteeVerifierOnchainAdapter
// for EVM chains. It is registered into the ccv adapter registry from init() so that
// ccv/deployment changesets can call it chain-family-agnostically.
type EVMCCVCommitteeVerifierOnchainAdapter struct{}

var _ ccvadapters.CommitteeVerifierOnchainAdapter = (*EVMCCVCommitteeVerifierOnchainAdapter)(nil)

func (a *EVMCCVCommitteeVerifierOnchainAdapter) ScanCommitteeStates(
	ctx context.Context,
	env deployment.Environment,
	chainSelector uint64,
) ([]*ccvadapters.CommitteeState, error) {
	refs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
		datastore.AddressRefByChainSelector(chainSelector),
	)
	if len(refs) == 0 {
		return nil, nil
	}

	evmChains := env.BlockChains.EVMChains()
	chain, ok := evmChains[chainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found in environment", chainSelector)
	}

	states := make([]*ccvadapters.CommitteeState, 0, len(refs))
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

		sigConfigs := make([]ccvadapters.SignatureConfig, 0, len(allConfigs))
		for _, cfg := range allConfigs {
			signers := make([]string, 0, len(cfg.Signers))
			for _, signer := range cfg.Signers {
				signers = append(signers, signer.Hex())
			}
			sigConfigs = append(sigConfigs, ccvadapters.SignatureConfig{
				SourceChainSelector: cfg.SourceChainSelector,
				Signers:             signers,
				Threshold:           cfg.Threshold,
			})
		}

		states = append(states, &ccvadapters.CommitteeState{
			Qualifier:        ref.Qualifier,
			ChainSelector:    chainSelector,
			Address:          ref.Address,
			SignatureConfigs: sigConfigs,
		})
	}

	return states, nil
}

func (a *EVMCCVCommitteeVerifierOnchainAdapter) ApplySignatureConfigs(
	ctx context.Context,
	env deployment.Environment,
	destChainSelector uint64,
	qualifier string,
	change ccvadapters.SignatureConfigChange,
) error {
	refs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
		datastore.AddressRefByChainSelector(destChainSelector),
		datastore.AddressRefByQualifier(qualifier),
	)
	if len(refs) == 0 {
		return fmt.Errorf("no CommitteeVerifier found for chain %d qualifier %q", destChainSelector, qualifier)
	}
	if len(refs) > 1 {
		return fmt.Errorf("multiple CommitteeVerifiers found for chain %d qualifier %q", destChainSelector, qualifier)
	}

	evmChains := env.BlockChains.EVMChains()
	chain, ok := evmChains[destChainSelector]
	if !ok {
		return fmt.Errorf("EVM chain %d not found in environment", destChainSelector)
	}

	addr := common.HexToAddress(refs[0].Address)
	contract, err := cv.NewCommitteeVerifier(addr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to bind CommitteeVerifier %s on chain %d: %w", refs[0].Address, destChainSelector, err)
	}

	sigConfigs := make([]cv.SignatureQuorumValidatorSignatureConfig, 0, len(change.NewConfigs))
	for _, c := range change.NewConfigs {
		signers := make([]common.Address, 0, len(c.Signers))
		for _, s := range c.Signers {
			if !common.IsHexAddress(s) {
				return fmt.Errorf("invalid signer address %q for source chain %d", s, c.SourceChainSelector)
			}
			signers = append(signers, common.HexToAddress(s))
		}
		sigConfigs = append(sigConfigs, cv.SignatureQuorumValidatorSignatureConfig{
			SourceChainSelector: c.SourceChainSelector,
			Threshold:           c.Threshold,
			Signers:             signers,
		})
	}

	tx, err := contract.ApplySignatureConfigs(chain.DeployerKey, change.RemovedSourceChainSelectors, sigConfigs)
	if err != nil {
		return fmt.Errorf("ApplySignatureConfigs tx failed on chain %d: %w", destChainSelector, err)
	}

	_, err = bind.WaitMined(ctx, chain.Client, tx)
	if err != nil {
		return fmt.Errorf("waiting for ApplySignatureConfigs tx on chain %d: %w", destChainSelector, err)
	}

	return nil
}
