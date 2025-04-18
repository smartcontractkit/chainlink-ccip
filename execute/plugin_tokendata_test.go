package execute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sel "github.com/smartcontractkit/chain-selectors"

	logger2 "github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func runRoundAndGetOutcome(ctx context.Context, ocrTypeCodec ocrtypecodec.ExecCodec,
	t *testing.T, r *testhelpers.OCR3Runner[[]byte]) exectypes.Outcome {
	result, err := r.RunRound(ctx)
	require.NoError(t, err)
	outcome, err := ocrTypeCodec.DecodeOutcome(result.Outcome)
	require.NoError(t, err)
	return outcome
}

func Test_USDC_Transfer(t *testing.T) {
	ocrTypeCodec := ocrtypecodec.DefaultExecCodec
	ctx := tests.Context(t)

	sourceChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
	destChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET_BASE_1.Selector)

	usdcAddress := "0x3765b189a8fe4a0bc34457835f01c9d178dbea60"
	usdcAddressBytes, err := cciptypes.NewUnknownAddressFromHex(usdcAddress)
	require.NoError(t, err)

	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(102, sourceChain, destChain, false),
		makeMsgWithMetadata(103, sourceChain, destChain, false),
		makeMsgWithMetadata(104, sourceChain, destChain, false, withTokens(cciptypes.RampTokenAmount{
			SourcePoolAddress: usdcAddressBytes,
			ExtraData:         readerpkg.NewSourceTokenDataPayload(1, 0).ToBytes(),
		})),
		makeMsgWithMetadata(105, sourceChain, destChain, false, withTokens(cciptypes.RampTokenAmount{
			SourcePoolAddress: usdcAddressBytes,
			ExtraData:         readerpkg.NewSourceTokenDataPayload(2, 0).ToBytes(),
		})),
		makeMsgWithMetadata(106, sourceChain, destChain, false,
			withTokens(cciptypes.RampTokenAmount{
				SourcePoolAddress: usdcAddressBytes,
				ExtraData:         readerpkg.NewSourceTokenDataPayload(3, 0).ToBytes(),
			}),
		),
	}

	events := []*readerpkg.MessageSentEvent{
		newMessageSentEvent(0, 6, 1, []byte{1}),
		newMessageSentEvent(0, 6, 2, []byte{2}),
		newMessageSentEvent(0, 6, 3, []byte{3}),
	}

	usdcAttestation104_106 := map[string]string{
		"0x0f43587da5355551d234a2ba24dde8edfe0e385346465d6d53653b6aa642992e": `{
			"status": "complete",
			"attestation": "0x100001"
		}`,
		"0x2b235443d276ec7dd517dcf34cca9dcd34f33542ccb6f305828d98e777404b63": `{
			"status": "complete",
			"attestation": "0x100003"
		}`,
	}

	intTest := SetupSimpleTest(t, logger2.Test(t), []cciptypes.ChainSelector{sourceChain}, destChain)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), 1, sourceChain)
	intTest.WithUSDC(usdcAddress, usdcAttestation104_106, events, sourceChain)
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 2 - Get Messages
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 3 - Filter
	// Messages 102-104,106 are executed, 105 doesn't have token data ready
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 1)
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	assert.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 106})
	//Attestation data added to the USDC
	assert.Equal(t, internal.MustDecodeRaw("0x100001"), outcome.Report.ChainReports[0].OffchainTokenData[2][0])
	assert.Equal(t, internal.MustDecodeRaw("0x100003"), outcome.Report.ChainReports[0].OffchainTokenData[3][0])

	intTest.usdcServer.AddResponse(
		"0x70ef528624085241badbff913575c0ab50241e7cb6db183a5614922ab0bcba5d",
		`{
			"status": "complete",
			"attestation": "0x100002"
		}`)

	// Run 3 more rounds to get all attestations
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}

	require.Len(t, outcome.Report.ChainReports, 1)
	sequenceNumbers = extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	// 102, 103 and 104 are in the inflight message cache.
	assert.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{105})
	//Attestation data added to the remaining USDC messages
	assert.Equal(t, internal.MustDecodeRaw("0x100002"), outcome.Report.ChainReports[0].OffchainTokenData[0][0])
}
