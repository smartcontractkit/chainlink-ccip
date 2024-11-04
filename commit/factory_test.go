package commit

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	reader2 "github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_maxQueryLength(t *testing.T) {
	// This test will verify that the maxQueryLength constant is set to a proper value.

	// Estimate the maximum number of source chains we are going to ever have.
	// This value should be tweaked after we are close to supporting that many chains.
	const estimatedMaxNumberOfSourceChains = 1000

	// Estimate the maximum number of RMN report signers we are going to ever have.
	// This value is defined in RMNRemote contract as `f`.
	// This value should be tweaked if necessary in order to define new limits.
	const estimatedMaxRmnReportSigners = 256

	sigs := make([]*rmnpb.EcdsaSignature, estimatedMaxRmnReportSigners)
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
	assert.Equal(t, 2*len(b), maxQueryLength)
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
			lggr: lggr,
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
