package execute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sel "github.com/smartcontractkit/chain-selectors"

	logger2 "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func runRoundAndGetOutcome(ctx context.Context, ocrTypeCodec ocrtypecodec.ExecCodec,
	t *testing.T, r *testhelpers.OCR3Runner[[]byte]) exectypes.Outcome {
	result, err := r.RunRound(ctx)
	require.NoError(t, err)
	outcome, err := ocrTypeCodec.DecodeOutcome(result.Outcome)
	require.NoError(t, err)
	return outcome
}

func Test_LBTC_USDC_Transfer(t *testing.T) {
	ocrTypeCodec := ocrtypecodec.DefaultExecCodec
	ctx := t.Context()

	sourceChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
	destChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET_BASE_1.Selector)

	usdcAddress := "0x3765b189a8fe4a0bc34457835f01c9d178dbea60"
	usdcAddressBytes, err := cciptypes.NewUnknownAddressFromHex(usdcAddress)
	require.NoError(t, err)

	lbtcAddress := "0xc791ec14ad1d566425f006eec12a300343164ab1"
	lbtcAddressBytes, err := cciptypes.NewUnknownAddressFromHex(lbtcAddress)
	require.NoError(t, err)
	lbtcMessageHash1 := internal.MustDecode("0xa9165956caf08b3da46db4cccdd58098b2c3a90e57372f3f28d7d46672e2091b")
	lbtcMessageHash2 := internal.MustDecode("0xc317b01e5a87000f8c51517227ea9ff07c9f4da646e8209c56424dc85ff50fe7")
	lbtcMessageHash3 := internal.MustDecode("0x5f9b38941ce144fad0b6890a3e15bc67c7b51bb89751ab6f672f088f44e36b91")

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
		makeMsgWithMetadata(106, sourceChain, destChain, false, withTokens(cciptypes.RampTokenAmount{
			SourcePoolAddress: lbtcAddressBytes,
			ExtraData:         lbtcMessageHash1,
		})),
		makeMsgWithMetadata(107, sourceChain, destChain, false, withTokens(cciptypes.RampTokenAmount{
			SourcePoolAddress: lbtcAddressBytes,
			ExtraData:         lbtcMessageHash2,
		})),
		makeMsgWithMetadata(108, sourceChain, destChain, false,
			withTokens(cciptypes.RampTokenAmount{
				SourcePoolAddress: usdcAddressBytes,
				ExtraData:         readerpkg.NewSourceTokenDataPayload(3, 0).ToBytes(),
			}, cciptypes.RampTokenAmount{
				SourcePoolAddress: lbtcAddressBytes,
				ExtraData:         lbtcMessageHash3,
			}),
		),
	}

	events := []*readerpkg.MessageSentEvent{
		newMessageSentEvent(0, 6, 1, []byte{1}),
		newMessageSentEvent(0, 6, 2, []byte{2}),
		newMessageSentEvent(0, 6, 3, []byte{3}),
	}

	usdcAttestation104_108 := map[string]string{
		"0x0f43587da5355551d234a2ba24dde8edfe0e385346465d6d53653b6aa642992e": `{
			"status": "complete",
			"attestation": "0x100001"
		}`,
		"0x2b235443d276ec7dd517dcf34cca9dcd34f33542ccb6f305828d98e777404b63": `{
			"status": "complete",
			"attestation": "0x100003"
		}`,
	}

	lbtcAttestation106_108 := map[string]string{
		lbtcMessageHash1.String(): `{
            "message_hash": "0xa9165956caf08b3da46db4cccdd58098b2c3a90e57372f3f28d7d46672e2091b",
			"status": "NOTARIZATION_STATUS_SESSION_APPROVED",
			"attestation": "0x200001"
		}`,
		lbtcMessageHash3.String(): `{
            "message_hash": "0x5f9b38941ce144fad0b6890a3e15bc67c7b51bb89751ab6f672f088f44e36b91",
			"status": "NOTARIZATION_STATUS_SESSION_APPROVED",
			"attestation": "0x200003"
		}`,
	}

	intTest := SetupSimpleTest(t, logger2.Nop(), []cciptypes.ChainSelector{sourceChain}, destChain)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), 1, sourceChain)
	intTest.WithUSDC(usdcAddress, usdcAttestation104_108, events, sourceChain)
	intTest.WithLBTC(lbtcAddress, lbtcAttestation106_108, sourceChain)
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 2 - Get Messages
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 3 - Filter
	// Messages 102-104,106,108 are executed, 105 and 107 don't have token data ready
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.NoError(t, err)
	require.Len(t, outcome.Reports[0].ChainReports, 1)
	sequenceNumbers := extractSequenceNumbers(outcome.Reports[0].ChainReports[0].Messages)
	assert.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 106, 108})
	//Attestation data added to the USDC
	assert.Equal(t, internal.MustDecodeRaw("0x100001"), outcome.Reports[0].ChainReports[0].OffchainTokenData[2][0])
	//Attestation data added to the LBTC
	assert.Equal(t, internal.MustDecodeRaw("0x200001"), outcome.Reports[0].ChainReports[0].OffchainTokenData[3][0])
	//Attestation data added to the USDC+LBTC
	assert.Equal(t, internal.MustDecodeRaw("0x100003"), outcome.Reports[0].ChainReports[0].OffchainTokenData[4][0])
	assert.Equal(t, internal.MustDecodeRaw("0x200003"), outcome.Reports[0].ChainReports[0].OffchainTokenData[4][1])

	intTest.usdcServer.AddResponse(
		"0x70ef528624085241badbff913575c0ab50241e7cb6db183a5614922ab0bcba5d",
		`{
			"status": "complete",
			"attestation": "0x100002"
		}`)

	intTest.lbtcServer.AddResponse(
		lbtcMessageHash2.String(),
		`{
            "message_hash": "0xc317b01e5a87000f8c51517227ea9ff07c9f4da646e8209c56424dc85ff50fe7",
			"status": "NOTARIZATION_STATUS_SESSION_APPROVED",
			"attestation": "0x200002"
		}`)

	// Run 3 more rounds to get all attestations
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}

	require.Len(t, outcome.Reports, 1)
	sequenceNumbers = extractSequenceNumbers(outcome.Reports[0].ChainReports[0].Messages)
	// 102, 103 and 104 are in the inflight message cache.
	assert.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{105, 107})
	//Attestation data added to the remaining USDC messages
	assert.Equal(t, internal.MustDecodeRaw("0x100002"), outcome.Reports[0].ChainReports[0].OffchainTokenData[0][0])
	//Attestation data added to the remaining LBTC messages
	assert.Equal(t, internal.MustDecodeRaw("0x200002"), outcome.Reports[0].ChainReports[0].OffchainTokenData[1][0])
}
