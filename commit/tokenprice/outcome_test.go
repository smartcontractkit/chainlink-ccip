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

func TestGetConsensusObservation(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr: lggr,
		cfg: pluginconfig.CommitPluginConfig{
			DestChain: destChainSel,
			OffchainConfig: pluginconfig.CommitOffchainConfig{
				TokenPriceBatchWriteFrequency: *commonconfig.MustNewDuration(time.Minute),
				TokenInfo: map[types.Account]pluginconfig.TokenInfo{
					tokenA: {DeviationPPB: cbi(1)},
					tokenB: {DeviationPPB: cbi(1)},
					tokenC: {DeviationPPB: cbi(1)},
					tokenD: {DeviationPPB: cbi(1)},
				},
				TokenPriceChainSelector: uint64(feedChainSel),
			},
		},
		bigF: 1,
	}

	ts := time.Now().UTC()
	obs := Observation{
		FeedTokenPrices: []cciptypes.TokenPrice{
			{TokenID: tokenA, Price: cbi100},
			{TokenID: tokenB, Price: cbi200},
			{TokenID: tokenC, Price: cbi100},
			{TokenID: tokenD, Price: cbi200},
		},
		FeeQuoterTokenUpdates: map[types.Account]shared.TimestampedBig{
			tokenA: {Timestamp: ts, Value: cbi100},
			tokenB: {Timestamp: ts, Value: cbi100},
		},
		FChain: map[cciptypes.ChainSelector]int{
			destChainSel: 1,
			feedChainSel: 2,
		},
		Timestamp: ts,
	}

	// 3 oracles, same observations, will pass destChain 2f+1 and fail feedChain 2f+1
	aos := []shared.AttributedObservation[Observation]{
		{OracleID: 1, Observation: obs},
		{OracleID: 2, Observation: obs},
		{OracleID: 3, Observation: obs},
	}

	consensusObs, err := p.getConsensusObservation(aos)
	assert.NoError(t, err)
	assert.Equal(t, 1, consensusObs.FChain[destChainSel])
	assert.Equal(t, 2, consensusObs.FChain[feedChainSel])

	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	// Only FeeQuoter will have consensus because we have
	assert.Len(t, consensusObs.FeeQuoterTokenUpdates, 2)
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
	assert.Equal(t, 1, consensusObs.FChain[destChainSel])
	assert.Equal(t, 2, consensusObs.FChain[feedChainSel])

	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	assert.Len(t, consensusObs.FeeQuoterTokenUpdates, 2)
	assert.Len(t, consensusObs.FeedTokenPrices, 4)
}

func TestSelectTokensForUpdate(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr: lggr,
		cfg: pluginconfig.CommitPluginConfig{
			OffchainConfig: pluginconfig.CommitOffchainConfig{
				TokenPriceBatchWriteFrequency: *commonconfig.MustNewDuration(time.Minute),
				TokenInfo: map[types.Account]pluginconfig.TokenInfo{
					tokenA: {DeviationPPB: cbi(1)},
					tokenB: {DeviationPPB: cbi(1)},
					tokenC: {DeviationPPB: cbi(1)},
					tokenD: {DeviationPPB: cbi(1)},
				},
			},
		},
		bigF: 1,
	}

	ts := time.Now().UTC()
	obs := ConsensusObservation{
		FeedTokenPrices: map[types.Account]cciptypes.TokenPrice{
			tokenA: {TokenID: tokenA, Price: cbi100},
			tokenB: {TokenID: tokenB, Price: cbi200},
			tokenC: {TokenID: tokenC, Price: cbi(300)},
			tokenD: {TokenID: tokenC, Price: cbi(400)},
		},
		FeeQuoterTokenUpdates: map[types.Account]shared.TimestampedBig{
			// tokenA Will be updated because of time
			tokenA: {Timestamp: ts.Add(-2 * time.Minute), Value: cbi100},
			// tokenB will be updated because of deviation
			tokenB: {Timestamp: ts, Value: cbi100},
			// tokenC will be updated because it's not available on feeQuoter

			//tokenD will not be updated because it's same price and time is not passed
			tokenD: {Timestamp: ts, Value: cbi(400)},
		},
		Timestamp: ts,
	}

	tokenPrices := p.selectTokensForUpdate(obs)
	assert.Len(t, tokenPrices, 3)
	assert.Equal(t, obs.FeedTokenPrices[tokenA], tokenPrices[0])
	assert.Equal(t, obs.FeedTokenPrices[tokenB], tokenPrices[1])
	assert.Equal(t, obs.FeedTokenPrices[tokenC], tokenPrices[2])
}
