package commit

import (
	"encoding/json"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	reader2 "github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_maxQueryLength(t *testing.T) {
	// This test will verify that the maxQueryLength constant is set to a proper value.

	sigs := make([]*rmnpb.EcdsaSignature, estimatedMaxRmnNodesCount)
	for i := range sigs {
		sigs[i] = &rmnpb.EcdsaSignature{R: make([]byte, 32), S: make([]byte, 32)}
	}

	laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, estimatedMaxNumberOfSourceChains)
	for i := range laneUpdates {
		laneUpdates[i] = &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: math.MaxUint64,
				OnrampAddress:       make([]byte, 40),
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: math.MaxUint64,
				MaxMsgNr: math.MaxUint64,
			},
			Root: make([]byte, 32),
		}
	}

	q := Query{
		MerkleRootQuery: merkleroot.Query{
			RetryRMNSignatures: true,
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
			},
		},
		TokenPriceQuery: tokenprice.Query{},
		ChainFeeQuery:   chainfee.Query{},
	}
	b, err := q.Encode()
	require.NoError(t, err)

	// We set twice the size, for extra safety while making breaking changes between oracle versions.
	const testOffset = 10
	assert.Greater(t, maxQueryLength, 2*len(b)-testOffset)
	assert.Less(t, maxQueryLength, 2*len(b)+testOffset)
	require.Less(t, maxQueryLength, ocr3types.MaxMaxQueryLength)
}

func Test_maxObservationLength(t *testing.T) {
	const maxContractsPerChain = 6 // router/onramp/offramp/rmnHome/rmnRemote/priceRegistry

	merkleRootObs := merkleroot.Observation{
		MerkleRoots:        make([]ccipocr3.MerkleRootChain, estimatedMaxNumberOfSourceChains),
		OnRampMaxSeqNums:   make([]plugintypes.SeqNumChain, estimatedMaxNumberOfSourceChains),
		OffRampNextSeqNums: make([]plugintypes.SeqNumChain, estimatedMaxNumberOfSourceChains),
		RMNRemoteConfig: rmntypes.RemoteConfig{
			ContractAddress:  make([]byte, 20),
			ConfigDigest:     [32]byte{},
			Signers:          make([]rmntypes.RemoteSignerInfo, estimatedMaxRmnNodesCount),
			F:                math.MaxUint64,
			ConfigVersion:    math.MaxUint32,
			RmnReportVersion: [32]byte{},
		},
		FChain: make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
	}

	for i := range merkleRootObs.MerkleRoots {
		merkleRootObs.MerkleRoots[i] = ccipocr3.MerkleRootChain{
			ChainSel:      math.MaxUint64,
			OnRampAddress: make([]byte, 40),
			SeqNumsRange:  ccipocr3.NewSeqNumRange(math.MaxUint64, math.MaxUint64),
			MerkleRoot:    [32]byte{},
		}
	}

	for i := range merkleRootObs.OnRampMaxSeqNums {
		merkleRootObs.OnRampMaxSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: math.MaxUint64,
			SeqNum:   math.MaxUint64,
		}
	}

	for i := range merkleRootObs.OffRampNextSeqNums {
		merkleRootObs.OffRampNextSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: math.MaxUint64,
			SeqNum:   math.MaxUint64,
		}
	}

	for i := range merkleRootObs.RMNRemoteConfig.Signers {
		merkleRootObs.RMNRemoteConfig.Signers[i] = rmntypes.RemoteSignerInfo{
			OnchainPublicKey: make([]byte, 40),
			NodeIndex:        math.MaxUint64,
		}
	}

	maxObs := Observation{
		MerkleRootObs: merkleRootObs,
		TokenPriceObs: tokenprice.Observation{
			FeedTokenPrices: make([]ccipocr3.TokenPrice, estimatedMaxNumberOfPricedTokens),
			FeeQuoterTokenUpdates: make(map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig,
				estimatedMaxNumberOfPricedTokens),
			FChain:    make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
			Timestamp: time.Now(),
		},
		ChainFeeObs: chainfee.Observation{
			FeeComponents:     make(map[ccipocr3.ChainSelector]types.ChainFeeComponents, estimatedMaxNumberOfSourceChains),
			NativeTokenPrices: make(map[ccipocr3.ChainSelector]ccipocr3.BigInt, estimatedMaxNumberOfPricedTokens),
			ChainFeeUpdates:   make(map[ccipocr3.ChainSelector]chainfee.Update, estimatedMaxNumberOfSourceChains),
			FChain:            make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
			TimestampNow:      time.Now(),
		},
		DiscoveryObs: dt.Observation{
			FChain:    make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
			Addresses: make(reader.ContractAddresses, estimatedMaxNumberOfSourceChains*maxContractsPerChain),
		},
		FChain: make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
	}

	for i := range maxObs.TokenPriceObs.FeedTokenPrices {
		maxObs.TokenPriceObs.FeedTokenPrices[i] = ccipocr3.TokenPrice{
			TokenID: ccipocr3.UnknownEncodedAddress(strings.Repeat("x", 20)),
			Price:   ccipocr3.NewBigIntFromInt64(math.MaxInt64),
		}
	}

	b, err := maxObs.Encode()
	require.NoError(t, err)

	const testOffset = 50
	assert.Greater(t, maxObservationLength, len(b)-testOffset)
	assert.Less(t, maxObservationLength, len(b)+testOffset)
	assert.Less(t, maxObservationLength, ocr3types.MaxMaxObservationLength)
}

func Test_maxOutcomeLength(t *testing.T) {
	maxOutc := Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType:                     merkleroot.OutcomeType(math.MaxInt),
			RangesSelectedForReport:         make([]plugintypes.ChainRange, estimatedMaxNumberOfSourceChains),
			RootsToReport:                   make([]ccipocr3.MerkleRootChain, estimatedMaxNumberOfSourceChains),
			OffRampNextSeqNums:              make([]plugintypes.SeqNumChain, estimatedMaxNumberOfSourceChains),
			ReportTransmissionCheckAttempts: math.MaxUint64,
			RMNReportSignatures:             make([]ccipocr3.RMNECDSASignature, estimatedMaxRmnNodesCount),
			RMNRemoteCfg: rmntypes.RemoteConfig{
				ContractAddress:  make([]byte, 20),
				ConfigDigest:     [32]byte{},
				Signers:          make([]rmntypes.RemoteSignerInfo, estimatedMaxRmnNodesCount),
				F:                math.MaxUint64,
				ConfigVersion:    math.MaxUint32,
				RmnReportVersion: [32]byte{},
			},
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: make([]ccipocr3.TokenPrice, estimatedMaxNumberOfPricedTokens),
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: make([]ccipocr3.GasPriceChain, estimatedMaxNumberOfSourceChains),
		},
	}

	for i := range maxOutc.MerkleRootOutcome.RangesSelectedForReport {
		maxOutc.MerkleRootOutcome.RangesSelectedForReport[i] = plugintypes.ChainRange{
			ChainSel:    math.MaxUint64,
			SeqNumRange: ccipocr3.NewSeqNumRange(math.MaxUint64, math.MaxUint64),
		}
	}

	for i := range maxOutc.MerkleRootOutcome.RootsToReport {
		maxOutc.MerkleRootOutcome.RootsToReport[i] = ccipocr3.MerkleRootChain{
			ChainSel:      math.MaxUint64,
			OnRampAddress: make([]byte, 40),
			SeqNumsRange:  ccipocr3.NewSeqNumRange(math.MaxUint64, math.MaxUint64),
			MerkleRoot:    [32]byte{},
		}
	}

	for i := range maxOutc.MerkleRootOutcome.OffRampNextSeqNums {
		maxOutc.MerkleRootOutcome.OffRampNextSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: math.MaxUint64,
			SeqNum:   math.MaxUint64,
		}
	}

	for i := range maxOutc.MerkleRootOutcome.RMNRemoteCfg.Signers {
		maxOutc.MerkleRootOutcome.RMNRemoteCfg.Signers[i] = rmntypes.RemoteSignerInfo{
			OnchainPublicKey: make([]byte, 40),
			NodeIndex:        math.MaxUint64,
		}
	}

	for i := range maxOutc.TokenPriceOutcome.TokenPrices {
		maxOutc.TokenPriceOutcome.TokenPrices[i] = ccipocr3.TokenPrice{
			TokenID: ccipocr3.UnknownEncodedAddress(strings.Repeat("x", 20)),
			Price:   ccipocr3.NewBigIntFromInt64(math.MaxInt64),
		}
	}

	for i := range maxOutc.ChainFeeOutcome.GasPrices {
		maxOutc.ChainFeeOutcome.GasPrices[i] = ccipocr3.GasPriceChain{
			ChainSel: math.MaxUint64,
			GasPrice: ccipocr3.NewBigIntFromInt64(math.MaxInt64),
		}
	}

	b, err := maxOutc.Encode()
	require.NoError(t, err)

	const testOffset = 10
	assert.Greater(t, maxOutcomeLength, len(b)-testOffset)
	assert.Less(t, maxOutcomeLength, len(b)+testOffset)
	assert.Less(t, maxOutcomeLength, ocr3types.MaxMaxOutcomeLength)
}

func Test_maxReportLength(t *testing.T) {
	rep := ccipocr3.CommitPluginReport{
		MerkleRoots: make([]ccipocr3.MerkleRootChain, estimatedMaxNumberOfSourceChains),
		PriceUpdates: ccipocr3.PriceUpdates{
			TokenPriceUpdates: make([]ccipocr3.TokenPrice, estimatedMaxNumberOfPricedTokens),
			GasPriceUpdates:   make([]ccipocr3.GasPriceChain, estimatedMaxNumberOfSourceChains),
		},
		RMNSignatures: make([]ccipocr3.RMNECDSASignature, estimatedMaxRmnNodesCount),
	}

	for i := range rep.MerkleRoots {
		rep.MerkleRoots[i] = ccipocr3.MerkleRootChain{
			ChainSel:      math.MaxUint64,
			OnRampAddress: make([]byte, 40),
			SeqNumsRange:  ccipocr3.NewSeqNumRange(math.MaxUint64, math.MaxUint64),
			MerkleRoot:    [32]byte{},
		}
	}

	for i := range rep.PriceUpdates.TokenPriceUpdates {
		rep.PriceUpdates.TokenPriceUpdates[i] = ccipocr3.TokenPrice{
			TokenID: ccipocr3.UnknownEncodedAddress(strings.Repeat("x", 20)),
			Price:   ccipocr3.NewBigIntFromInt64(math.MaxInt64),
		}
	}

	for i := range rep.PriceUpdates.GasPriceUpdates {
		rep.PriceUpdates.GasPriceUpdates[i] = ccipocr3.GasPriceChain{
			ChainSel: math.MaxUint64,
			GasPrice: ccipocr3.NewBigIntFromInt64(math.MaxInt64),
		}
	}

	// Chain specific encoding are more compact than JSON. We measure using JSON encoding.
	b, err := json.Marshal(rep)
	require.NoError(t, err)

	const testOffset = 10
	assert.Greater(t, maxReportLength, len(b)-testOffset)
	assert.Less(t, maxReportLength, len(b)+testOffset)
	assert.Less(t, maxReportLength, ocr3types.MaxMaxReportLength)
}

func TestPluginFactory_NewReportingPlugin(t *testing.T) {
	t.Run("basic checks for the happy flow", func(t *testing.T) {
		ctx := tests.Context(t)
		lggr := logger.Test(t)

		offChainConfig := pluginconfig.CommitOffchainConfig{
			MaxMerkleTreeSize: 123,
		}
		b, err := json.Marshal(offChainConfig)
		require.NoError(t, err)

		p := &PluginFactory{
			baseLggr: lggr,
			ocrConfig: reader.OCR3ConfigWithMeta{
				Version:      1,
				ConfigDigest: [32]byte{1, 2, 3},
				Config: reader2.OCR3Config{
					OfframpAddress: []byte{1, 2, 3},
					OffchainConfig: b,
					ChainSelector:  1,
				},
			},
		}

		plugin, pluginInfo, err := p.NewReportingPlugin(ctx, ocr3types.ReportingPluginConfig{
			OffchainConfig: b,
		})
		require.NoError(t, err)

		pluginCommit, is := plugin.(*Plugin)
		require.True(t, is)
		pluginOffchainConfig := pluginCommit.offchainCfg

		require.Equal(t, uint(5), pluginOffchainConfig.MaxReportTransmissionCheckAttempts)          // default is used
		require.Equal(t, merklemulti.MaxNumberTreeLeaves, pluginOffchainConfig.NewMsgScanBatchSize) // default is used
		require.Equal(t, offChainConfig.MaxMerkleTreeSize, pluginOffchainConfig.MaxMerkleTreeSize)  // override

		require.Equal(t, maxQueryLength, pluginInfo.Limits.MaxQueryLength)
	})
}
