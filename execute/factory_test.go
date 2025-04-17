package execute

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

/*
200 reports, 200 messages
-------------------------
Total size of observation: 1056507
200 CommitReports: 60635
200 Messages: 583635
TokenData: 112035
Nonces: 12035
Contracts: 274962

900 reports, 100 messages
-------------------------
Total size of observation: 937407
900 CommitReports: 302235
100 Messages: 291835
TokenData: 56035
Nonces: 6035
Contracts: 274962

100 reports, 100 messages
-------------------------
Total size of observation: 665507
100 CommitReports: 30335
100 Messages: 291835
TokenData: 56035
Nonces: 6035
Contracts: 274962

100 reports, 200 messages
-------------------------
Total size of observation: 1025407
100 CommitReports: 30335
200 Messages: 582835
TokenData: 112035
Nonces: 12035
Contracts: 274962

100 reports, 350 messages
-------------------------
Total size of observation: 1565257
100 CommitReports: 30335
350 Messages: 1019335
TokenData: 196035
Nonces: 21035
Contracts: 274962
*/
func TestObservationSize(t *testing.T) {
	t.Skip("This test is for estimating message sizes, not for running in CI")
	ocrTypeCodec := ocrtypecodec.DefaultExecCodec

	maxCommitReports := 100
	maxMessages := 1100
	msgDataSize := 1000 // could be much larger than this?
	tokenDataSize := 0  // fixed size for CCTP?

	var addr [20]byte

	commitObs := make(exectypes.CommitObservations, estimatedMaxNumberOfSourceChains)
	bigSeqNum := ccipocr3.SeqNum(100000)
	for i := 0; i < maxCommitReports; i++ {
		idx := ccipocr3.ChainSelector(i % estimatedMaxNumberOfSourceChains)
		seqNum := bigSeqNum + ccipocr3.SeqNum(i)
		commitObs[idx] = append(commitObs[idx], exectypes.CommitData{
			SourceChain: ccipocr3.ChainSelector(123456),
			Timestamp:   time.UnixMilli(1732035256660),
			BlockNum:    uint64(302173055),
			//MerkleRoot cciptypes.Bytes32 `json:"merkleRoot"`
			SequenceNumberRange: ccipocr3.NewSeqNumRange(seqNum, seqNum+1),

			// None are executed in this scenario with 1 message per pending commit
			//ExecutedMessages: ...

			// These fields are all empty during this observation phase.
			//Messages []cciptypes.Message `json:"messages"`
			//MessageTokenData []MessageTokenData `json:"messageTokenData"`
		})
	}

	msgObs := make(exectypes.MessageObservations, estimatedMaxNumberOfSourceChains)
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % estimatedMaxNumberOfSourceChains % maxCommitReports)
		if nil == msgObs[idx] {
			msgObs[idx] = make(map[ccipocr3.SeqNum]ccipocr3.Message)
		}
		data := make([]byte, msgDataSize)
		var extraArgs [100]byte // this is too large?
		msgObs[idx][bigSeqNum+ccipocr3.SeqNum(i)] = ccipocr3.Message{
			Header:         ccipocr3.RampMessageHeader{},
			Sender:         addr[:],
			Data:           data[:],
			Receiver:       addr[:],
			ExtraArgs:      extraArgs[:],
			FeeToken:       addr[:],
			FeeTokenAmount: ccipocr3.BigInt{},
			FeeValueJuels:  ccipocr3.BigInt{},
			TokenAmounts: []ccipocr3.RampTokenAmount{
				{
					SourcePoolAddress: addr[:],
					DestTokenAddress:  addr[:],
					ExtraData:         nil,
					Amount:            ccipocr3.NewBigInt(big.NewInt(100000)),
					DestExecData:      nil,
				},
			},
		}
	}

	// This could be bigger, since each message could send multiple tokens.
	tokenDataObs := make(exectypes.TokenDataObservations, estimatedMaxNumberOfSourceChains)
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % estimatedMaxNumberOfSourceChains)
		if nil == tokenDataObs[idx] {
			tokenDataObs[idx] = make(map[ccipocr3.SeqNum]exectypes.MessageTokenData)
		}
		data := make([]byte, tokenDataSize)
		tokenDataObs[idx][bigSeqNum+ccipocr3.SeqNum(i)] = exectypes.MessageTokenData{
			TokenData: []exectypes.TokenData{
				{
					Ready:     true,
					Data:      data[:],
					Error:     nil,
					Supported: true,
				},
			},
		}
	}

	// separate sender for each message
	noncesObs := make(exectypes.NonceObservations, maxMessages)
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % estimatedMaxNumberOfSourceChains)
		if nil == noncesObs[idx] {
			noncesObs[idx] = make(map[string]uint64)
		}
		encAddr, err := mockAddrCodec.AddressBytesToString(addr[:], idx)
		require.NoError(t, err)
		noncesObs[idx][encAddr] = uint64(bigSeqNum + ccipocr3.SeqNum(i))
	}

	contracts := []string{
		consts.ContractNameFeeQuoter,
		consts.ContractNameOnRamp,
		consts.ContractNameOffRamp,
		consts.ContractNameRMNRemote,
		consts.ContractNameRouter,
		consts.ContractNameNonceManager,
		//consts.ContractNameRMNHome, // I don't think we look this one up
	}
	discoveryObs := dt.Observation{
		FChain:    make(map[ccipocr3.ChainSelector]int, estimatedMaxNumberOfSourceChains),
		Addresses: make(reader.ContractAddresses, len(contracts)),
	}
	set := func(contract string) {
		//type ContractAddresses map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress
		discoveryObs.Addresses[contract] =
			make(map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress, estimatedMaxNumberOfSourceChains)
		for i := 0; i < estimatedMaxNumberOfSourceChains; i++ {
			discoveryObs.Addresses[contract][ccipocr3.ChainSelector(i)] = addr[:]
		}
	}
	for _, contract := range contracts {
		set(contract)
	}

	maxObs := exectypes.Observation{
		CommitReports: commitObs,
		Messages:      msgObs,
		TokenData:     tokenDataObs,
		Nonces:        noncesObs,
		Contracts:     discoveryObs,
	}

	encSize := func(obs exectypes.Observation) int {
		b, err := ocrTypeCodec.EncodeObservation(obs)
		require.NoError(t, err)
		return len(b)
	}

	msgSum := 0
	for _, msgs := range msgObs {
		msgSum += len(msgs)
	}
	fmt.Println("Total size of observation:", encSize(maxObs))
	fmt.Printf("%d CommitReports: %d\n", len(maxObs.CommitReports),
		encSize(exectypes.Observation{CommitReports: commitObs}))
	fmt.Printf("%d Messages: %d\n", msgSum, encSize(exectypes.Observation{Messages: msgObs}))
	fmt.Printf("TokenData: %d\n", encSize(exectypes.Observation{TokenData: tokenDataObs}))
	fmt.Printf("Nonces: %d\n", encSize(exectypes.Observation{Nonces: noncesObs}))
	fmt.Printf("Contracts: %d\n", encSize(exectypes.Observation{Contracts: discoveryObs}))

	b, err := ocrTypeCodec.EncodeObservation(maxObs)
	require.NoError(t, err)
	assert.Greater(t, lenientMaxObservationLength, len(b))
	assert.LessOrEqual(t, lenientMaxObservationLength, ocr3types.MaxMaxObservationLength)
}
