package aggregator_config

import (
	"maps"

	model "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

// ToModelCommittee converts the operation's Committee type to model.Committee.
func (c *Committee) ToModelCommittee() *model.Committee {
	if c == nil {
		return nil
	}

	quorumConfigs := make(map[model.SourceSelector]*model.QuorumConfig)
	for selector, qc := range c.QuorumConfigs {
		signers := make([]model.Signer, 0, len(qc.Signers))
		for _, s := range qc.Signers {
			signers = append(signers, model.Signer{
				Address: s.Address,
			})
		}
		quorumConfigs[selector] = &model.QuorumConfig{
			SourceVerifierAddress: qc.SourceVerifierAddress,
			Signers:               signers,
			Threshold:             qc.Threshold,
		}
	}

	destVerifiers := make(map[model.DestinationSelector]string, len(c.DestinationVerifiers))
	maps.Copy(destVerifiers, c.DestinationVerifiers)

	return &model.Committee{
		QuorumConfigs:        quorumConfigs,
		DestinationVerifiers: destVerifiers,
	}
}
