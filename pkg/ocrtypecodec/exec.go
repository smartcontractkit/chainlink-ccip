package ocrtypecodec

import (
	"encoding/json"
	"math/big"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ExecCodec is an interface for encoding and decoding OCR related exec plugin types.
type ExecCodec interface {
	EncodeObservation(observation exectypes.Observation) ([]byte, error)
	DecodeObservation(data []byte) (exectypes.Observation, error)

	EncodeOutcome(outcome exectypes.Outcome) ([]byte, error)
	DecodeOutcome(data []byte) (exectypes.Outcome, error)
}

type ExecCodecProto struct{}

func (e ExecCodecProto) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	// Encode CommitReports
	commitReports := make(map[uint64]*ocrtypecodecpb.CommitObservations, len(observation.CommitReports))
	for chainSel, commits := range observation.CommitReports {
		commitData := make([]*ocrtypecodecpb.CommitData, len(commits))
		for i, commit := range commits {
			commitData[i] = &ocrtypecodecpb.CommitData{
				SourceChain:   uint64(commit.SourceChain),
				OnRampAddress: commit.OnRampAddress,
				Timestamp:     uint64(commit.Timestamp.Unix()),
				BlockNum:      commit.BlockNum,
				MerkleRoot:    commit.MerkleRoot[:],
				SequenceNumberRange: &ocrtypecodecpb.SeqNumRange{
					MinMsgNr: uint64(commit.SequenceNumberRange.Start()),
					MaxMsgNr: uint64(commit.SequenceNumberRange.End()),
				},
				ExecutedMessages: encodeSeqNums(commit.ExecutedMessages),
				Messages:         encodeMessages(commit.Messages),
				Hashes:           encodeBytes32Slice(commit.Hashes),
				CostlyMessages:   encodeBytes32Slice(commit.CostlyMessages),
				MessageTokenData: encodeMessageTokenData(commit.MessageTokenData),
			}
		}
		commitReports[uint64(chainSel)] = &ocrtypecodecpb.CommitObservations{
			CommitData: commitData,
		}
	}

	// Encode Messages
	messages := make(map[uint64]*ocrtypecodecpb.SeqNumToMessage, len(observation.Messages))
	for chainSel, seqNums := range observation.Messages {
		seqNumToMsg := make(map[uint64]*ocrtypecodecpb.Message, len(seqNums))
		for seqNum, msg := range seqNums {
			seqNumToMsg[uint64(seqNum)] = encodeMessage(msg)
		}
		messages[uint64(chainSel)] = &ocrtypecodecpb.SeqNumToMessage{
			Messages: seqNumToMsg,
		}
	}

	// Encode Hashes
	messageHashes := make(map[uint64]*ocrtypecodecpb.SeqNumToBytes, len(observation.Hashes))
	for chainSel, hashMap := range observation.Hashes {
		seqNumToBytes := &ocrtypecodecpb.SeqNumToBytes{SeqNumToBytes: make(map[uint64][]byte, len(hashMap))}
		for seqNum, hash := range hashMap {
			seqNumToBytes.SeqNumToBytes[uint64(seqNum)] = hash[:]
		}
		messageHashes[uint64(chainSel)] = seqNumToBytes
	}

	// Encode TokenDataObservations
	tokenDataObservations := make(map[uint64]*ocrtypecodecpb.SeqNumToTokenData, len(observation.TokenData))
	for chainSel, tokenMap := range observation.TokenData {
		seqNumToTokenData := &ocrtypecodecpb.SeqNumToTokenData{
			TokenData: make(map[uint64]*ocrtypecodecpb.MessageTokenData),
		}
		for seqNum, tokenData := range tokenMap {
			seqNumToTokenData.TokenData[uint64(seqNum)] = encodeMessageTokenDataEntry(tokenData)
		}
		tokenDataObservations[uint64(chainSel)] = seqNumToTokenData
	}

	// Encode Costly Messages
	costlyMessages := encodeBytes32Slice(observation.CostlyMessages)

	// Encode Nonces
	nonceObservations := make(map[uint64]*ocrtypecodecpb.StringAddrToNonce, len(observation.Nonces))
	for chainSel, nonceMap := range observation.Nonces {
		addrToNonce := make(map[string]uint64, len(nonceMap))
		for addr, nonce := range nonceMap {
			addrToNonce[addr] = nonce
		}
		nonceObservations[uint64(chainSel)] = &ocrtypecodecpb.StringAddrToNonce{Nonces: addrToNonce}
	}

	// Encode Contracts (DiscoveryObservation)
	contracts := &ocrtypecodecpb.DiscoveryObservation{
		FChain: encodeFChain(observation.Contracts.FChain),
		ContractNames: &ocrtypecodecpb.ContractNameChainAddresses{
			Addresses: encodeContractAddresses(observation.Contracts.Addresses),
		},
	}

	// Encode FChain
	fChain := encodeFChain(observation.FChain)

	pbObs := &ocrtypecodecpb.ExecObservation{
		CommitReports:         commitReports,
		SeqNumsToMessages:     messages,
		MessageHashes:         messageHashes,
		TokenDataObservations: &ocrtypecodecpb.TokenDataObservations{TokenData: tokenDataObservations},
		CostlyMessages:        costlyMessages,
		Nonces:                nonceObservations,
		Contracts:             contracts,
		FChain:                fChain,
	}

	return proto.Marshal(pbObs)
}

// Helper Functions

func encodeSeqNums(seqNums []cciptypes.SeqNum) []uint64 {
	result := make([]uint64, len(seqNums))
	for i, num := range seqNums {
		result[i] = uint64(num)
	}
	return result
}

func encodeMessages(messages []cciptypes.Message) []*ocrtypecodecpb.Message {
	result := make([]*ocrtypecodecpb.Message, len(messages))
	for i, msg := range messages {
		result[i] = encodeMessage(msg)
	}
	return result
}

func encodeMessage(msg cciptypes.Message) *ocrtypecodecpb.Message {
	return &ocrtypecodecpb.Message{
		Header:         encodeMessageHeader(msg.Header),
		Sender:         msg.Sender,
		Data:           msg.Data,
		Receiver:       msg.Receiver,
		ExtraArgs:      msg.ExtraArgs,
		FeeToken:       msg.FeeToken,
		FeeTokenAmount: msg.FeeTokenAmount.Bytes(),
		FeeValueJuels:  msg.FeeValueJuels.Bytes(),
		TokenAmounts:   encodeRampTokenAmounts(msg.TokenAmounts),
	}
}

func encodeMessageHeader(header cciptypes.RampMessageHeader) *ocrtypecodecpb.RampMessageHeader {
	return &ocrtypecodecpb.RampMessageHeader{
		MessageId:           header.MessageID[:],
		SourceChainSelector: uint64(header.SourceChainSelector),
		DestChainSelector:   uint64(header.DestChainSelector),
		SequenceNumber:      uint64(header.SequenceNumber),
		Nonce:               header.Nonce,
		MsgHash:             header.MsgHash[:],
		OnRamp:              header.OnRamp,
	}
}

func encodeRampTokenAmounts(tokenAmounts []cciptypes.RampTokenAmount) []*ocrtypecodecpb.RampTokenAmount {
	result := make([]*ocrtypecodecpb.RampTokenAmount, len(tokenAmounts))
	for i, token := range tokenAmounts {
		result[i] = &ocrtypecodecpb.RampTokenAmount{
			SourcePoolAddress: token.SourcePoolAddress,
			DestTokenAddress:  token.DestTokenAddress,
			ExtraData:         token.ExtraData,
			Amount:            token.Amount.Bytes(),
			DestExecData:      token.DestExecData,
		}
	}
	return result
}

func encodeBytes32Slice(slice []cciptypes.Bytes32) [][]byte {
	result := make([][]byte, len(slice))
	for i, val := range slice {
		result[i] = val[:]
	}
	return result
}

func encodeMessageTokenData(data []exectypes.MessageTokenData) []*ocrtypecodecpb.MessageTokenData {
	result := make([]*ocrtypecodecpb.MessageTokenData, len(data))
	for i, item := range data {
		result[i] = encodeMessageTokenDataEntry(item)
	}
	return result
}

func encodeMessageTokenDataEntry(data exectypes.MessageTokenData) *ocrtypecodecpb.MessageTokenData {
	tokenData := make([]*ocrtypecodecpb.TokenData, len(data.TokenData))
	for i, td := range data.TokenData {
		tokenData[i] = &ocrtypecodecpb.TokenData{
			Ready: td.Ready,
			Data:  td.Data,
		}
	}
	return &ocrtypecodecpb.MessageTokenData{TokenData: tokenData}
}

func encodeFChain(fChain map[cciptypes.ChainSelector]int) map[uint64]int32 {
	result := make(map[uint64]int32, len(fChain))
	for k, v := range fChain {
		result[uint64(k)] = int32(v)
	}
	return result
}

func encodeContractAddresses(addresses reader.ContractAddresses) map[string]*ocrtypecodecpb.ChainAddressMap {
	result := make(map[string]*ocrtypecodecpb.ChainAddressMap, len(addresses))
	for name, chains := range addresses {
		chainAddrs := make(map[uint64][]byte, len(chains))
		for chain, addr := range chains {
			chainAddrs[uint64(chain)] = addr[:]
		}
		result[name] = &ocrtypecodecpb.ChainAddressMap{ChainAddresses: chainAddrs}
	}
	return result
}

func (e ExecCodecProto) DecodeObservation(data []byte) (exectypes.Observation, error) {
	pbObs := &ocrtypecodecpb.ExecObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return exectypes.Observation{}, err
	}

	// Decode CommitReports
	commitReports := make(exectypes.CommitObservations, len(pbObs.CommitReports))
	for chainSel, commitObs := range pbObs.CommitReports {
		commitData := make([]exectypes.CommitData, len(commitObs.CommitData))
		for i, commit := range commitObs.CommitData {
			commitData[i] = exectypes.CommitData{
				SourceChain:   cciptypes.ChainSelector(commit.SourceChain),
				OnRampAddress: commit.OnRampAddress,
				Timestamp:     time.Unix(int64(commit.Timestamp), 0),
				BlockNum:      commit.BlockNum,
				MerkleRoot:    cciptypes.Bytes32(commit.MerkleRoot),
				SequenceNumberRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(commit.SequenceNumberRange.MinMsgNr),
					cciptypes.SeqNum(commit.SequenceNumberRange.MaxMsgNr),
				),
				ExecutedMessages: decodeSeqNums(commit.ExecutedMessages),
				Messages:         decodeMessages(commit.Messages),
				Hashes:           decodeBytes32Slice(commit.Hashes),
				CostlyMessages:   decodeBytes32Slice(commit.CostlyMessages),
				MessageTokenData: decodeMessageTokenData(commit.MessageTokenData),
			}
		}
		commitReports[cciptypes.ChainSelector(chainSel)] = commitData
	}

	// Decode Messages
	messages := make(exectypes.MessageObservations, len(pbObs.SeqNumsToMessages))
	for chainSel, msgMap := range pbObs.SeqNumsToMessages {
		innerMap := make(map[cciptypes.SeqNum]cciptypes.Message, len(msgMap.Messages))
		for seqNum, msg := range msgMap.Messages {
			innerMap[cciptypes.SeqNum(seqNum)] = decodeMessage(msg)
		}
		messages[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	// Decode Hashes
	hashes := make(exectypes.MessageHashes, len(pbObs.MessageHashes))
	for chainSel, hashMap := range pbObs.MessageHashes {
		innerMap := make(map[cciptypes.SeqNum]cciptypes.Bytes32, len(hashMap.SeqNumToBytes))
		for seqNum, hash := range hashMap.SeqNumToBytes {
			innerMap[cciptypes.SeqNum(seqNum)] = cciptypes.Bytes32(hash)
		}
		hashes[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	// Decode TokenDataObservations
	tokenDataObservations := make(exectypes.TokenDataObservations, len(pbObs.TokenDataObservations.TokenData))
	for chainSel, tokenMap := range pbObs.TokenDataObservations.TokenData {
		innerMap := make(map[cciptypes.SeqNum]exectypes.MessageTokenData, len(tokenMap.TokenData))
		for seqNum, tokenData := range tokenMap.TokenData {
			innerMap[cciptypes.SeqNum(seqNum)] = decodeMessageTokenDataEntry(tokenData)
		}
		tokenDataObservations[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	// Decode Costly Messages
	costlyMessages := decodeBytes32Slice(pbObs.CostlyMessages)

	// Decode Nonces
	nonces := make(exectypes.NonceObservations, len(pbObs.Nonces))
	for chainSel, nonceMap := range pbObs.Nonces {
		innerMap := make(map[string]uint64, len(nonceMap.Nonces))
		for addr, nonce := range nonceMap.Nonces {
			innerMap[addr] = nonce
		}
		nonces[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	// Decode Contracts (DiscoveryObservation)
	contracts := discoverytypes.Observation{
		FChain:    decodeFChain(pbObs.Contracts.FChain),
		Addresses: decodeContractAddresses(pbObs.Contracts.ContractNames),
	}

	// Decode FChain
	fChain := decodeFChain(pbObs.FChain)

	return exectypes.Observation{
		CommitReports:  commitReports,
		Messages:       messages,
		Hashes:         hashes,
		TokenData:      tokenDataObservations,
		CostlyMessages: costlyMessages,
		Nonces:         nonces,
		Contracts:      contracts,
		FChain:         fChain,
	}, nil
}

// Helper Functions

func decodeContractAddresses(addresses *ocrtypecodecpb.ContractNameChainAddresses) reader.ContractAddresses {
	result := make(reader.ContractAddresses, len(addresses.Addresses))

	for name, chainMap := range addresses.Addresses {
		chainAddrs := make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress, len(chainMap.ChainAddresses))
		for chain, addr := range chainMap.ChainAddresses {
			chainAddrs[cciptypes.ChainSelector(chain)] = addr
		}
		result[name] = chainAddrs
	}
	return result
}

func decodeMessageTokenData(data []*ocrtypecodecpb.MessageTokenData) []exectypes.MessageTokenData {
	result := make([]exectypes.MessageTokenData, len(data))
	for i, item := range data {
		result[i] = decodeMessageTokenDataEntry(item)
	}
	return result
}

func decodeSeqNums(seqNums []uint64) []cciptypes.SeqNum {
	result := make([]cciptypes.SeqNum, len(seqNums))
	for i, num := range seqNums {
		result[i] = cciptypes.SeqNum(num)
	}
	return result
}

func decodeMessages(messages []*ocrtypecodecpb.Message) []cciptypes.Message {
	result := make([]cciptypes.Message, len(messages))
	for i, msg := range messages {
		result[i] = decodeMessage(msg)
	}
	return result
}

func decodeMessage(msg *ocrtypecodecpb.Message) cciptypes.Message {
	return cciptypes.Message{
		Header:         decodeMessageHeader(msg.Header),
		Sender:         msg.Sender,
		Data:           msg.Data,
		Receiver:       msg.Receiver,
		ExtraArgs:      msg.ExtraArgs,
		FeeToken:       msg.FeeToken,
		FeeTokenAmount: cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeTokenAmount)),
		FeeValueJuels:  cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeValueJuels)),
		TokenAmounts:   decodeRampTokenAmounts(msg.TokenAmounts),
	}
}

func decodeMessageHeader(header *ocrtypecodecpb.RampMessageHeader) cciptypes.RampMessageHeader {
	return cciptypes.RampMessageHeader{
		MessageID:           cciptypes.Bytes32(header.MessageId),
		SourceChainSelector: cciptypes.ChainSelector(header.SourceChainSelector),
		DestChainSelector:   cciptypes.ChainSelector(header.DestChainSelector),
		SequenceNumber:      cciptypes.SeqNum(header.SequenceNumber),
		Nonce:               header.Nonce,
		MsgHash:             cciptypes.Bytes32(header.MsgHash),
		OnRamp:              header.OnRamp,
	}
}

func decodeRampTokenAmounts(tokenAmounts []*ocrtypecodecpb.RampTokenAmount) []cciptypes.RampTokenAmount {
	result := make([]cciptypes.RampTokenAmount, len(tokenAmounts))
	for i, token := range tokenAmounts {
		result[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: token.SourcePoolAddress,
			DestTokenAddress:  token.DestTokenAddress,
			ExtraData:         token.ExtraData,
			Amount:            cciptypes.NewBigInt(big.NewInt(0).SetBytes(token.Amount)),
			DestExecData:      token.DestExecData,
		}
	}
	return result
}

func decodeBytes32Slice(slice [][]byte) []cciptypes.Bytes32 {
	result := make([]cciptypes.Bytes32, len(slice))
	for i, val := range slice {
		result[i] = cciptypes.Bytes32(val)
	}
	return result
}

func decodeMessageTokenDataEntry(data *ocrtypecodecpb.MessageTokenData) exectypes.MessageTokenData {
	tokenData := make([]exectypes.TokenData, len(data.TokenData))
	for i, td := range data.TokenData {
		tokenData[i] = exectypes.TokenData{
			Ready: td.Ready,
			Data:  td.Data,
		}
	}
	return exectypes.MessageTokenData{TokenData: tokenData}
}

func decodeFChain(fChain map[uint64]int32) map[cciptypes.ChainSelector]int {
	result := make(map[cciptypes.ChainSelector]int, len(fChain))
	for k, v := range fChain {
		result[cciptypes.ChainSelector(k)] = int(v)
	}
	return result
}

func (e ExecCodecProto) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	return nil, nil
}

func (e ExecCodecProto) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	return exectypes.Outcome{}, nil
}

func NewExecCodecProto() *ExecCodecProto {
	return &ExecCodecProto{}
}

type ExecCodecJSON struct{}

func NewExecCodecJSON() *ExecCodecJSON {
	return &ExecCodecJSON{}
}

func (*ExecCodecJSON) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	return json.Marshal(observation)
}

func (*ExecCodecJSON) DecodeObservation(data []byte) (exectypes.Observation, error) {
	if len(data) == 0 {
		return exectypes.Observation{}, nil
	}
	obs := exectypes.Observation{}
	err := json.Unmarshal(data, &obs)
	return obs, err
}

func (*ExecCodecJSON) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	// We sort again here in case construction is not via the constructor.
	return json.Marshal(exectypes.NewSortedOutcome(outcome.State, outcome.CommitReports, outcome.Report))
}

func (*ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}
	o := exectypes.Outcome{}
	err := json.Unmarshal(data, &o)
	return o, err
}
