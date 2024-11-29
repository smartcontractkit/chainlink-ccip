package execute

import (
	"fmt"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestObservationSizeLimits(t *testing.T) {
	maxCommitReports := 100
	maxMessages := 1100
	msgDataSize := 1000 // could be much larger than this?
	tokenDataSize := 0  // fixed size for CCTP?
	nChains := 1

	prevOutcome := exectypes.Outcome{
		State: exectypes.GetMessages,
	}

	var addr [20]byte
	commitObs := make(exectypes.CommitObservations, nChains)
	bigSeqNum := ccipocr3.SeqNum(100000)
	for i := 0; i < maxCommitReports; i++ {
		//idx := ccipocr3.ChainSelector(i % nChains)
		seqNum := bigSeqNum + ccipocr3.SeqNum(i)
		commitObs[chainA] = append(commitObs[chainA], exectypes.CommitData{
			SourceChain: chainA,
			Timestamp:   time.UnixMilli(1732035256660),
			BlockNum:    uint64(302173055),
			//MerkleRoot cciptypes.Bytes32 `json:"merkleRoot"`
			SequenceNumberRange: ccipocr3.NewSeqNumRange(seqNum, seqNum+1),

			// None are executed in this scenario with 1 message per pending commit
			//ExecutedMessages: ...

			// These fields are all empty during this observation phase.
			//Messages []cciptypes.Message `json:"messages"`
			//CostlyMessages []cciptypes.Bytes32 `json:"costlyMessages"`
			//MessageTokenData []MessageTokenData `json:"messageTokenData"`
		})
	}

	prevOutcome.PendingCommitReports = commitObs[chainA]

	msgObs := make(exectypes.MessageObservations, nChains)
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % nChains % maxCommitReports)
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
	tokenDataObs := make(exectypes.TokenDataObservations, nChains)
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % nChains)
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
	for i := 0; i < maxMessages; i++ {
		idx := ccipocr3.ChainSelector(i % nChains)
		if nil == noncesObs[idx] {
			noncesObs[idx] = make(map[string]uint64)
		}
		encAddr := typeconv.AddressBytesToString(addr[:], 123456)
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
		FChain:    make(map[ccipocr3.ChainSelector]int, nChains),
		Addresses: make(reader.ContractAddresses, len(contracts)),
	}
	set := func(contract string) {
		//type ContractAddresses map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress
		discoveryObs.Addresses[contract] =
			make(map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress, nChains)
		for i := 0; i < nChains; i++ {
			discoveryObs.Addresses[contract][ccipocr3.ChainSelector(i)] = addr[:]
		}
	}
	for _, contract := range contracts {
		set(contract)
	}

	costlyMessagesObs := make([]ccipocr3.Bytes32, maxMessages)

	maxObs := exectypes.Observation{
		CommitReports:  commitObs,
		Messages:       msgObs,
		TokenData:      tokenDataObs,
		CostlyMessages: costlyMessagesObs,
		Nonces:         noncesObs,
		Contracts:      discoveryObs,
	}

	encSize := func(obs exectypes.Observation) int {
		b, err := obs.Encode()
		require.NoError(t, err)
		return len(b)
	}

	fmt.Printf("maxObs size: %d\n", encSize(maxObs))

}
