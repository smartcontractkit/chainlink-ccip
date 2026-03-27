package sequences

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	laneapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

func init() {
	v, err := semver.NewVersion("1.6.0")
	if err != nil {
		panic(err)
	}
	adapter := &SolanaAdapter{}
	laneapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, v, adapter)
	deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilySolana, v, adapter)
	deployapi.GetUpgraderRegistry().RegisterUpgrader(chain_selectors.FamilySolana, v, adapter)
	deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilySolana, v, adapter)
	mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilySolana, adapter)
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilySolana, v, adapter)
}

type SolanaAdapter struct {
	timelockAddr map[uint64]solana.PublicKey
}

func (a *SolanaAdapter) GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return a.GetRouterAddress(ds, chainSelector)
}

func (a *SolanaAdapter) GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(offramp.ContractType),
		Version:       offramp.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       fee_quoter.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (a *SolanaAdapter) GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// PrepareArtifacts implements deploy.ArtifactPreparer.
// It ensures Solana program .so files are available in ProgramsPath before deployment.
// If ChainSpecific contains a *utils.SolanaBuildConfig, it uses that to drive the
// build/download. Otherwise, it checks that artifacts already exist (the in-memory test path).
func (a *SolanaAdapter) PrepareArtifacts(e cldf.Environment, chainSelector uint64, cfg deployapi.ContractDeploymentConfigPerChain) error {
	chain, ok := e.BlockChains.SolanaChains()[chainSelector]
	if !ok {
		return fmt.Errorf("solana chain not found for selector %d", chainSelector)
	}

	buildCfg, _ := cfg.ChainSpecific.(*utils.SolanaBuildConfig)
	if buildCfg == nil {
		// No explicit build config — check if artifacts already exist on disk.
		// In-memory tests preload programs before genesis, so this is a no-op.
		if _, err := os.Stat(chain.ProgramsPath); err != nil {
			return fmt.Errorf("no build config provided and programs path does not exist: %s", chain.ProgramsPath)
		}
		// Verify at least one .so file exists
		entries, err := os.ReadDir(chain.ProgramsPath)
		if err != nil {
			return fmt.Errorf("failed to read programs path: %w", err)
		}
		for _, entry := range entries {
			if filepath.Ext(entry.Name()) == ".so" {
				e.Logger.Infow("Artifacts already present, skipping build", "path", chain.ProgramsPath)
				return nil
			}
		}
		return fmt.Errorf("no .so artifacts found in %s and no build config provided", chain.ProgramsPath)
	}

	// Override destination to the chain's programs path if not explicitly set
	if buildCfg.DestinationDir == "" {
		buildCfg.DestinationDir = chain.ProgramsPath
	}

	return utils.BuildSolana(e.GetContext(), e.Logger, *buildCfg)
}

func (a *SolanaAdapter) GetRMNRemoteAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(rmnremoteops.ContractType),
		Version:       rmnremoteops.Version,
	}, chainSelector, utils.ToByteArray)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// svmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector SVM")) = 0x1e10bdc4.
var svmFamilySelector = [4]byte{0x1e, 0x10, 0xbd, 0xc4}

func (a *SolanaAdapter) GetFeeQuoterDestChainConfig() laneapi.FeeQuoterDestChainConfig {
	return laneapi.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   16,
		ChainFamilySelector:         binary.BigEndian.Uint32(svmFamilySelector[:]),
		DefaultTokenFeeUSDCents:     25,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		V1Params: &laneapi.FeeQuoterV1Params{
			MaxNumberOfTokensPerMsg:           10,
			DestGasPerPayloadByteHigh:         40,
			DestGasPerPayloadByteThreshold:    3000,
			DestDataAvailabilityOverheadGas:   100,
			DestGasPerDataAvailabilityByte:    16,
			DestDataAvailabilityMultiplierBps: 1,
			GasMultiplierWeiPerEth:            11e17,
		},
		V2Params: &laneapi.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: 90,
			USDPerUnitGas:            big.NewInt(1e6),
		},
	}
}

func (a *SolanaAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(4e12)
}

func (a *SolanaAdapter) GetChainFamilySelector() [4]byte {
	return svmFamilySelector
}
