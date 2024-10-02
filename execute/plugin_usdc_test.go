package execute

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"

	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func Test_USDC_Transfer(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	sourceChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
	destChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET_BASE_1.Selector)

	addressBytes, err := hex.DecodeString(strings.TrimPrefix(randomEthAddress, "0x"))
	require.NoError(t, err)

	messages := []inmem.MessagesWithMetadata{
		makeMsg(102, sourceChain, destChain, false),
		makeMsg(103, sourceChain, destChain, false),
		makeMsgWithToken(104, sourceChain, destChain, false, []cciptypes.RampTokenAmount{
			{
				SourcePoolAddress: addressBytes,
				ExtraData:         readerpkg.NewSourceTokenDataPayload(1, 0).ToBytes(),
			},
		}),
		makeMsgWithToken(105, sourceChain, destChain, false, []cciptypes.RampTokenAmount{
			{
				SourcePoolAddress: addressBytes,
				ExtraData:         readerpkg.NewSourceTokenDataPayload(2, 0).ToBytes(),
			},
		}),
	}

	intTest, runner := SetupSimpleTest(ctx, t, lggr, sourceChain, destChain)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour))
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runner.MustRunRound(t, ctx)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	outcome = runner.MustRunRound(t, ctx)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 2 - Get Messages
	outcome = runner.MustRunRound(t, ctx)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 3 - Filter
	// Messages 102-104 are executed, 105 doesn't have token data ready
	outcome = runner.MustRunRound(t, ctx)
	require.NoError(t, err)
	sequenceNumbers := slicelib.Map(outcome.Report.ChainReports[0].Messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104})
	//Attestation data added to the USDC
	require.NotEmpty(t, outcome.Report.ChainReports[0].OffchainTokenData[2])

	intTest.server.AddResponse(
		"0x70ef528624085241badbff913575c0ab50241e7cb6db183a5614922ab0bcba5d",
		`{
			"status": "complete",
			"attestation": "0x720502893578a89a8a87982982ef781c18b194"
		}`)

	// Run 3 more rounds to get all attestations
	for i := 0; i < 3; i++ {
		outcome = runner.MustRunRound(t, ctx)
	}

	sequenceNumbers = slicelib.Map(outcome.Report.ChainReports[0].Messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})
	//Attestation data added to the both USDC messages
	require.NotEmpty(t, outcome.Report.ChainReports[0].OffchainTokenData[2])
	require.NotEmpty(t, outcome.Report.ChainReports[0].OffchainTokenData[3])
}
