package execute

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"

	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func Test_USDC_Transfer(t *testing.T) {
	ctx := tests.Context(t)

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

	events := []*readerpkg.MessageSentEvent{
		newMessageSentEvent(0, 6, 1, []byte{1}),
		newMessageSentEvent(0, 6, 2, []byte{2}),
		newMessageSentEvent(0, 6, 3, []byte{3}),
	}

	attestation104 := map[string]string{
		"0x0f43587da5355551d234a2ba24dde8edfe0e385346465d6d53653b6aa642992e": `{
			"status": "complete",
			"attestation": "0x720502893578a89a8a87982982ef781c18b193"
		}`,
	}

	intTest := SetupSimpleTest(t, sourceChain, destChain)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour))
	intTest.WithUSDC(randomEthAddress, attestation104, events)
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runner.MustRunRound(ctx, t)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	outcome = runner.MustRunRound(ctx, t)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 2 - Get Messages
	outcome = runner.MustRunRound(ctx, t)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 3 - Filter
	// Messages 102-104 are executed, 105 doesn't have token data ready
	outcome = runner.MustRunRound(ctx, t)
	require.NoError(t, err)
	sequenceNumbers := testhelpers.ExtractSequenceNumbers(outcome)
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
		outcome = runner.MustRunRound(ctx, t)
	}

	sequenceNumbers = testhelpers.ExtractSequenceNumbers(outcome)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})
	//Attestation data added to the both USDC messages
	require.NotEmpty(t, outcome.Report.ChainReports[0].OffchainTokenData[2])
	require.NotEmpty(t, outcome.Report.ChainReports[0].OffchainTokenData[3])
}
