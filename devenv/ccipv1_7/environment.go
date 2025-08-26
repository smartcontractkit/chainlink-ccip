package ccipv1_7

import (
	"fmt"

	"github.com/smartcontractkit/devenv/ccipv17/services"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/fake"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/jd"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/s3provider"

	ns "github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"
)

type Cfg struct {
	Indexer         *services.IndexerInput `toml:"indexer" validate:"required"`
	CCIPv17         *CCIPv17               `toml:"ccipv17" validate:"required"`
	StorageProvider *s3provider.Input      `toml:"storage_provider" validate:"required"`
	FakeServer      *fake.Input            `toml:"fake_server"      validate:"required"`
	JD              *jd.Input              `toml:"jd"`
	Blockchains     []*blockchain.Input    `toml:"blockchains"      validate:"required"`
	NodeSets        []*ns.Input            `toml:"nodesets"         validate:"required"`
}

// verifyEnvironment internal function describing how to verify your environment is working.
func verifyEnvironment(in *Cfg) error {
	if !in.CCIPv17.Verify {
		return nil
	}
	Plog.Info().Msg("Verifying environment")
	// CCIPv17 verification, check that example transfer works
	return nil
}

// NewEnvironment creates a new datafeeds environment either locally in Docker or remotely in K8s.
func NewEnvironment() (*Cfg, error) {
	if err := framework.DefaultNetwork(nil); err != nil {
		return nil, err
	}
	in, err := Load[Cfg]()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	track := NewTimeTracker(Plog)
	for _, b := range in.Blockchains {
		_, err = blockchain.NewBlockchainNetwork(b)
		if err != nil {
			return nil, fmt.Errorf("failed to create blockchain network: %w", err)
		}
	}
	s3Out, err := s3provider.NewMinioFactory().NewFrom(in.StorageProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage provider: %w", err)
	}
	_ = s3Out
	_, err = fake.NewDockerFakeDataProvider(in.FakeServer)
	if err != nil {
		return nil, fmt.Errorf("failed to create fake data provider: %w", err)
	}
	_, err = services.NewIndexer(in.Indexer)
	if err != nil {
		return nil, fmt.Errorf("failed to create an example service: %w", err)
	}
	_, err = jd.NewJD(in.JD)
	if err != nil {
		return nil, fmt.Errorf("failed to create job distributor: %w", err)
	}
	track.Record("[infra] deploying blockchains")
	if err := DefaultProductConfiguration(in, ConfigureNodesNetwork); err != nil {
		return nil, fmt.Errorf("failed to setup default CLDF orchestration: %w", err)
	}
	track.Record("[changeset] configured nodes network")
	_, err = ns.NewSharedDBNodeSet(in.NodeSets[0], nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new shared db node set: %w", err)
	}
	track.Record("[infra] deployed CL nodes")
	if err := DefaultProductConfiguration(in, ConfigureProductContractsJobs); err != nil {
		return nil, fmt.Errorf("failed to setup default CLDF orchestration: %w", err)
	}
	track.Record("[changeset] deployed product contracts")
	track.Print()
	return in, Store[Cfg](in)
}
