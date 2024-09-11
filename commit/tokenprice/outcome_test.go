package tokenprice

import (
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/shared"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/stretchr/testify/assert"
)

var ts = time.Now().UTC()

var feedTokenPricesMap = map[types.Account]cciptypes.TokenPrice{
	tokenA: {TokenID: tokenA, Price: cbi100},
	tokenB: {TokenID: tokenB, Price: cbi200},
	tokenC: {TokenID: tokenC, Price: cbi100},
	tokenD: {TokenID: tokenD, Price: cbi200},
}

var feedTokenPrices = []cciptypes.TokenPrice{
	feedTokenPricesMap[tokenA],
	feedTokenPricesMap[tokenB],
	feedTokenPricesMap[tokenC],
	feedTokenPricesMap[tokenD],
}

var feeQuoterUpdates = map[types.Account]shared.TimestampedBig{
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
	TokenInfo: map[types.Account]pluginconfig.TokenInfo{
		tokenA: {DeviationPPB: cbi(1)},
		tokenB: {DeviationPPB: cbi(2)},
		tokenC: {DeviationPPB: cbi(3)},
		tokenD: {DeviationPPB: cbi(4)},
	},
	TokenPriceChainSelector: uint64(feedChainSel),
}

func TestGetConsensusObservation(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr: lggr,
		cfg: pluginconfig.CommitPluginConfig{
			DestChain:      destChainSel,
			OffchainConfig: offChainCfg,
		},
		bigF: 1,
	}

	// 3 oracles, same observations, will pass destChain 2f+1 and fail feedChain 2f+1
	aos := []shared.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
	}

	consensusObs, err := p.getConsensusObservation(aos)
	assert.NoError(t, err)
	assert.Equal(t, fChains[destChainSel], fChains[destChainSel])
	assert.Equal(t, fChains[feedChainSel], fChains[feedChainSel])

	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	// Only FeeQuoter will have consensus because we have
	assert.Len(t, consensusObs.FeeQuoterTokenUpdates, 3)
	assert.Len(t, consensusObs.FeedTokenPrices, 0)

	// Same but with 5 oracles, will have consensus on both feedprice and feequoter
	aos = []shared.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
		{OracleID: 4, Observation: obs},
		{OracleID: 5, Observation: obs},
	}

	consensusObs, err = p.getConsensusObservation(aos)
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
		lggr: lggr,
		cfg: pluginconfig.CommitPluginConfig{
			DestChain:      destChainSel,
			OffchainConfig: offChainCfg,
		},
		bigF: 1,
	}

	conObs := ConsensusObservation{
		FeedTokenPrices:       feedTokenPricesMap,
		FeeQuoterTokenUpdates: feeQuoterUpdates,
		Timestamp:             ts,
	}

	// tokenA Will be updated because of time
	// tokenB will be updated because of deviation
	// tokenC will be updated because it's not available on feeQuoter
	//tokenD will not be updated because it's same price and time is not passed
	tokenPrices := p.selectTokensForUpdate(conObs)
	assert.Len(t, tokenPrices, 3)
	assert.Equal(t, conObs.FeedTokenPrices[tokenA], tokenPrices[0])
	assert.Equal(t, conObs.FeedTokenPrices[tokenB], tokenPrices[1])
	assert.Equal(t, conObs.FeedTokenPrices[tokenC], tokenPrices[2])
}

// Test Plugin Outcome method returns the correct token prices
func TestOutcome(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr: lggr,
		cfg: pluginconfig.CommitPluginConfig{
			DestChain:      destChainSel,
			OffchainConfig: offChainCfg,
		},
		bigF: 1,
	}

	outcome, err := p.Outcome(Outcome{}, Query{}, []shared.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
		{OracleID: 4, Observation: obs},
		{OracleID: 5, Observation: obs},
	})

	expectedOutcome := []cciptypes.TokenPrice{
		feedTokenPricesMap[tokenA],
		feedTokenPricesMap[tokenB],
		feedTokenPricesMap[tokenC],
		// tokenD is not updated because it's the same price and time is not passed
	}

	assert.NoError(t, err)
	assert.Len(t, outcome.TokenPrices, 3)
	assert.Equal(t, expectedOutcome, outcome.TokenPrices)
}
