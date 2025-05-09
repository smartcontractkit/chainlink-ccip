package tokenprice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var ts = time.Now().UTC()

var feedTokenPricesMap = map[cciptypes.UnknownEncodedAddress]cciptypes.TokenPrice{
	tokenA: {TokenID: tokenA, Price: cbi100},
	tokenB: {TokenID: tokenB, Price: cbi200},
	tokenC: {TokenID: tokenC, Price: cbi100},
	tokenD: {TokenID: tokenD, Price: cbi200},
}

var feedTokenPrices = cciptypes.TokenPriceMap{
	tokenA: feedTokenPricesMap[tokenA].Price,
	tokenB: feedTokenPricesMap[tokenB].Price,
	tokenC: feedTokenPricesMap[tokenC].Price,
	tokenD: feedTokenPricesMap[tokenD].Price,
}

var feeQuoterUpdates = map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
	tokenA: {Timestamp: ts.Add(-2 * time.Minute), Value: cbi100},     // Update because of time
	tokenB: {Timestamp: ts, Value: cbi100},                           // update because of deviation
	tokenD: {Timestamp: ts, Value: feedTokenPricesMap[tokenD].Price}, // no update, same price and timestamp
}
var fChains = map[cciptypes.ChainSelector]int{
	destChainSel: 1,
	feedChainSel: 2,
}
var obs = Observation{
	FeedTokenPrices:       feedTokenPrices,
	FeeQuoterTokenUpdates: feeQuoterUpdates,
	FChain:                fChains,
	Timestamp:             ts,
}

var offChainCfg = pluginconfig.CommitOffchainConfig{
	TokenPriceBatchWriteFrequency: *commonconfig.MustNewDuration(time.Minute),
	TokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
		tokenA: {DeviationPPB: cbi(1)},
		tokenB: {DeviationPPB: cbi(2)},
		tokenC: {DeviationPPB: cbi(3)},
		tokenD: {DeviationPPB: cbi(4)},
	},
	PriceFeedChainSelector: feedChainSel,
}

func TestGetConsensusObservation(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr:        lggr,
		destChain:   destChainSel,
		offChainCfg: offChainCfg,
		fRoleDON:    1,
	}

	// 3 oracles, same observations, will pass destChain 2f+1 and fail feedChain 2f+1
	aos := []plugincommon.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
	}

	consensusObs, err := p.getConsensusObservation(lggr, aos)
	assert.NoError(t, err)
	assert.Equal(t, fChains[destChainSel], fChains[destChainSel])
	assert.Equal(t, fChains[feedChainSel], fChains[feedChainSel])

	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	// Only FeeQuoter will have consensus because we have
	assert.Len(t, consensusObs.FeeQuoterTokenUpdates, 3)
	assert.Len(t, consensusObs.FeedTokenPrices, 0)

	// Same but with 5 oracles, will have consensus on both feedprice and feequoter
	aos = []plugincommon.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
		{OracleID: 4, Observation: obs},
		{OracleID: 5, Observation: obs},
	}

	consensusObs, err = p.getConsensusObservation(lggr, aos)
	assert.NoError(t, err)
	assert.Equal(t, fChains[destChainSel], consensusObs.FChain[destChainSel])
	assert.Equal(t, fChains[feedChainSel], consensusObs.FChain[feedChainSel])

	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	assert.Len(t, consensusObs.FeeQuoterTokenUpdates, 3)
	assert.Len(t, consensusObs.FeedTokenPrices, 4)
}

func TestSelectTokensForUpdate(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr:        lggr,
		destChain:   destChainSel,
		offChainCfg: offChainCfg,
		fRoleDON:    1,
	}

	conObs := ConsensusObservation{
		FeedTokenPrices:       feedTokenPricesMap,
		FeeQuoterTokenUpdates: feeQuoterUpdates,
		Timestamp:             ts,
	}

	// tokenA Will be updated because of time
	// tokenB will be updated because of deviation
	// tokenC will be updated because it's not available on feeQuoter
	// tokenD will not be updated because it's same price and time is not passed
	tokenPrices := p.selectTokensForUpdate(lggr, conObs)
	assert.Len(t, tokenPrices, 3)
	assert.Equal(t, conObs.FeedTokenPrices[tokenA].Price, tokenPrices[tokenA])
	assert.Equal(t, conObs.FeedTokenPrices[tokenB].Price, tokenPrices[tokenB])
	assert.Equal(t, conObs.FeedTokenPrices[tokenC].Price, tokenPrices[tokenC])
}

// Test Plugin Outcome method returns the correct token prices
func TestOutcome(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)
	p := &processor{
		lggr:            lggr,
		destChain:       destChainSel,
		offChainCfg:     offChainCfg,
		fRoleDON:        1,
		metricsReporter: plugincommon.NoopReporter{},
	}

	outcome, err := p.Outcome(ctx, Outcome{}, Query{}, []plugincommon.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
		{OracleID: 4, Observation: obs},
		{OracleID: 5, Observation: obs},
	})

	expectedOutcome := cciptypes.TokenPriceMap{
		tokenA: feedTokenPricesMap[tokenA].Price,
		tokenB: feedTokenPricesMap[tokenB].Price,
		tokenC: feedTokenPricesMap[tokenC].Price,
		// tokenD is not updated because it's the same price and time is not passed
	}

	assert.NoError(t, err)
	assert.Len(t, outcome.TokenPrices, 3)
	assert.Equal(t, expectedOutcome, outcome.TokenPrices)
}

// TestOutcome_EmptyObservations tests the Outcome method when observations only contain minimal data.
func TestOutcome_EmptyObservations(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)
	numOracles := 5 // Need enough oracles for consensus

	p := &processor{
		lggr:            lggr,
		destChain:       destChainSel,
		offChainCfg:     offChainCfg,
		fRoleDON:        fChains[destChainSel], // Use f from fChains for the destination chain
		metricsReporter: plugincommon.NoopReporter{},
	}

	// Prepare attributed observations with only minimal data
	aos := make([]plugincommon.AttributedObservation[Observation], numOracles)
	for i := 0; i < numOracles; i++ {
		obs := Observation{
			FChain:    fChains,
			Timestamp: ts,
			// FeedTokenPrices and FeeQuoterTokenUpdates are nil/empty
		}
		aos[i] = plugincommon.AttributedObservation[Observation]{
			Observation: obs,
			OracleID:    commontypes.OracleID(i),
		}
	}

	// Call Outcome
	outcome, err := p.Outcome(ctx, Outcome{}, Query{}, aos)

	// Assertions
	assert.NoError(t, err)
	assert.Empty(t, outcome.TokenPrices, "Expected TokenPrices to be empty when observations have no price data")
}
