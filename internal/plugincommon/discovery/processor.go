package discovery

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// Ensure that PluginProcessor is implemented.
var _ plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome] = &ContractDiscoveryProcessor{}

// ContractDiscoveryProcessor is a plugin processor for discovering contracts.
type ContractDiscoveryProcessor struct {
	lggr      logger.Logger
	reader    *reader.CCIPReader
	homechain reader.HomeChain
	dest      cciptypes.ChainSelector
	fRoleDON  int
}

func NewContractDiscoveryProcessor(
	lggr logger.Logger,
	reader *reader.CCIPReader,
	homechain reader.HomeChain,
	dest cciptypes.ChainSelector,
	fRoleDON int,
) *ContractDiscoveryProcessor {
	return &ContractDiscoveryProcessor{
		lggr:      lggr,
		reader:    reader,
		homechain: homechain,
		dest:      dest,
		fRoleDON:  fRoleDON,
	}
}

// Query is not needed for this processor.
func (cdp *ContractDiscoveryProcessor) Query(_ context.Context, _ dt.Outcome) (dt.Query, error) {
	return nil, nil
}

// Observation reads contract addresses from the destination chain.
// In the future this should be extended to omit observations unless one of the nodes requests addresses.
func (cdp *ContractDiscoveryProcessor) Observation(
	ctx context.Context, _ dt.Outcome, _ dt.Query,
) (dt.Observation, error) {
	contracts, err := (*cdp.reader).DiscoverContracts(ctx, cdp.dest)
	if err != nil {
		if errors.Is(err, reader.ErrContractReaderNotFound) {
			// Not a dest reader, no observations will be made.
			// Processor is not disabled because the outcome phase will bind observed contracts.
			return dt.Observation{}, nil
		}
		return dt.Observation{}, fmt.Errorf("unable to discover contracts: %w", err)
	}

	fChain, err := cdp.homechain.GetFChain()
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to get fchain: %w", err)
	}

	return dt.Observation{
		FChain:           fChain,
		OnRamp:           contracts[consts.ContractNameOnRamp],
		DestNonceManager: contracts[consts.ContractNameNonceManager][cdp.dest],
	}, nil
}

func (cdp *ContractDiscoveryProcessor) ValidateObservation(
	_ dt.Outcome, _ dt.Query, _ plugincommon.AttributedObservation[dt.Observation],
) error {
	// TODO: make sure all observations come from dest readers.
	return nil
}

// Outcome comes to consensus on the contract addresses and updates the chainreader. It doesn't actually
// return an Outcome.
func (cdp *ContractDiscoveryProcessor) Outcome(
	_ dt.Outcome, _ dt.Query, aos []plugincommon.AttributedObservation[dt.Observation],
) (dt.Outcome, error) {
	// come to consensus on the onramp addresses and update the chainreader.

	// fChain consensus - uses the role DON F value because all nodes can observe the home chain.
	fChainObs := make(map[cciptypes.ChainSelector][]int)
	for _, ao := range aos {
		for chainSel, f := range ao.Observation.FChain {
			fChainObs[chainSel] = append(fChainObs[chainSel], f)
		}
	}
	fMin := make(map[cciptypes.ChainSelector]int)
	for chain := range fChainObs {
		fMin[chain] = cdp.fRoleDON
	}

	fChain := plugincommon.GetConsensusMap(cdp.lggr, "fChain", fChainObs, fMin)

	// onramp address consensus
	onrampAddrs := make(map[cciptypes.ChainSelector][][]byte)
	for _, ao := range aos {
		for chain, addr := range ao.Observation.OnRamp {
			onrampAddrs[chain] = append(onrampAddrs[chain], addr)
		}
	}

	// nonce manager address consensus
	var nonceManagerAddrs [][]byte
	for _, ao := range aos {
		nonceManagerAddrs = append(
			nonceManagerAddrs,
			ao.Observation.DestNonceManager,
		)
	}

	nonceManagerConsensus, err := getConsensusNonceManager("nonceManager", nonceManagerAddrs, fChain[cdp.dest])
	if err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to reach consensus on nonce manager: %w", err)
	}

	// call Sync to bind contracts.
	contracts := make(map[string]map[cciptypes.ChainSelector][]byte)
	contracts[consts.ContractNameOnRamp] = plugincommon.GetConsensusMap(cdp.lggr, "onramp", onrampAddrs, fChain)
	contracts[consts.ContractNameNonceManager] = map[cciptypes.ChainSelector][]byte{
		cdp.dest: nonceManagerConsensus,
	}
	if err := (*cdp.reader).Sync(context.Background(), contracts); err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return dt.Outcome{}, nil
}

func getConsensusNonceManager(
	objectName string,
	nonceManagerAddrs [][]byte,
	f int,
) ([]byte, error) {
	if len(nonceManagerAddrs) < 2*f+1 {
		return nil, fmt.Errorf(
			"could not reach consensus on %s, not enough observations (want %d got %d)",
			objectName, 2*f+1, len(nonceManagerAddrs))
	}
	keyer := func(addr []byte) string { return hex.EncodeToString(addr) }
	unkeyer := func(key string) ([]byte, error) { return hex.DecodeString(key) }
	votes := make(map[string]int)
	for _, addr := range nonceManagerAddrs {
		votes[keyer(addr)]++
	}

	// Find the most common nonce manager address.
	maxVotes := 0
	var maxAddr string
	for addr, vote := range votes {
		if vote > maxVotes {
			maxVotes = vote
			maxAddr = addr
		}
	}

	// Check if the most common nonce manager address was observed at least f+1 times.
	if maxVotes < f+1 {
		return nil, fmt.Errorf(
			"could not reach consensus on %s, no single address was observed more than %d times",
			objectName, f)
	}

	// Return the most common nonce manager address.
	nonceManager, err := unkeyer(maxAddr)
	if err != nil {
		return nil, fmt.Errorf("could not decode nonce manager address: %w", err)
	}

	return nonceManager, nil
}
