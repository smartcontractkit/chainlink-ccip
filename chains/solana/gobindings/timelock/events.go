// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"encoding/base64"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_rpc "github.com/gagliardetto/solana-go/rpc"
	ag_base58 "github.com/mr-tron/base58"
	"reflect"
	"strings"
)

type BypasserCallExecutedEventData struct {
	Index  uint64
	Target ag_solanago.PublicKey
	Data   []byte
}

var BypasserCallExecutedEventDataDiscriminator = [8]byte{61, 41, 96, 207, 16, 173, 99, 75}

func (obj BypasserCallExecutedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BypasserCallExecutedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Target` param:
	err = encoder.Encode(obj.Target)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BypasserCallExecutedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BypasserCallExecutedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[61 41 96 207 16 173 99 75]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Target`:
	err = decoder.Decode(&obj.Target)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (*BypasserCallExecutedEventData) isEventData() {}

type CallExecutedEventData struct {
	Id     [32]uint8
	Index  uint64
	Target ag_solanago.PublicKey
	Data   []byte
}

var CallExecutedEventDataDiscriminator = [8]byte{237, 120, 238, 142, 189, 37, 65, 128}

func (obj CallExecutedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CallExecutedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Target` param:
	err = encoder.Encode(obj.Target)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CallExecutedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CallExecutedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[237 120 238 142 189 37 65 128]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Target`:
	err = decoder.Decode(&obj.Target)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (*CallExecutedEventData) isEventData() {}

type CallScheduledEventData struct {
	Id          [32]uint8
	Index       uint64
	Target      ag_solanago.PublicKey
	Predecessor [32]uint8
	Salt        [32]uint8
	Delay       uint64
	Data        []byte
}

var CallScheduledEventDataDiscriminator = [8]byte{191, 85, 90, 167, 132, 223, 184, 57}

func (obj CallScheduledEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CallScheduledEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Target` param:
	err = encoder.Encode(obj.Target)
	if err != nil {
		return err
	}
	// Serialize `Predecessor` param:
	err = encoder.Encode(obj.Predecessor)
	if err != nil {
		return err
	}
	// Serialize `Salt` param:
	err = encoder.Encode(obj.Salt)
	if err != nil {
		return err
	}
	// Serialize `Delay` param:
	err = encoder.Encode(obj.Delay)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CallScheduledEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CallScheduledEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[191 85 90 167 132 223 184 57]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Target`:
	err = decoder.Decode(&obj.Target)
	if err != nil {
		return err
	}
	// Deserialize `Predecessor`:
	err = decoder.Decode(&obj.Predecessor)
	if err != nil {
		return err
	}
	// Deserialize `Salt`:
	err = decoder.Decode(&obj.Salt)
	if err != nil {
		return err
	}
	// Deserialize `Delay`:
	err = decoder.Decode(&obj.Delay)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (*CallScheduledEventData) isEventData() {}

type CancelledEventData struct {
	Id [32]uint8
}

var CancelledEventDataDiscriminator = [8]byte{136, 23, 42, 65, 143, 233, 234, 46}

func (obj CancelledEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CancelledEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CancelledEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CancelledEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[136 23 42 65 143 233 234 46]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	return nil
}

func (*CancelledEventData) isEventData() {}

type FunctionSelectorBlockedEventData struct {
	Selector [8]uint8
}

var FunctionSelectorBlockedEventDataDiscriminator = [8]byte{67, 101, 36, 217, 222, 85, 191, 71}

func (obj FunctionSelectorBlockedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FunctionSelectorBlockedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Selector` param:
	err = encoder.Encode(obj.Selector)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FunctionSelectorBlockedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FunctionSelectorBlockedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[67 101 36 217 222 85 191 71]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Selector`:
	err = decoder.Decode(&obj.Selector)
	if err != nil {
		return err
	}
	return nil
}

func (*FunctionSelectorBlockedEventData) isEventData() {}

type FunctionSelectorUnblockedEventData struct {
	Selector [8]uint8
}

var FunctionSelectorUnblockedEventDataDiscriminator = [8]byte{189, 124, 164, 141, 141, 189, 0, 218}

func (obj FunctionSelectorUnblockedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FunctionSelectorUnblockedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Selector` param:
	err = encoder.Encode(obj.Selector)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FunctionSelectorUnblockedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FunctionSelectorUnblockedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[189 124 164 141 141 189 0 218]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Selector`:
	err = decoder.Decode(&obj.Selector)
	if err != nil {
		return err
	}
	return nil
}

func (*FunctionSelectorUnblockedEventData) isEventData() {}

type MinDelayChangeEventData struct {
	OldDuration uint64
	NewDuration uint64
}

var MinDelayChangeEventDataDiscriminator = [8]byte{186, 71, 244, 116, 244, 76, 230, 254}

func (obj MinDelayChangeEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MinDelayChangeEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `OldDuration` param:
	err = encoder.Encode(obj.OldDuration)
	if err != nil {
		return err
	}
	// Serialize `NewDuration` param:
	err = encoder.Encode(obj.NewDuration)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MinDelayChangeEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MinDelayChangeEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[186 71 244 116 244 76 230 254]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `OldDuration`:
	err = decoder.Decode(&obj.OldDuration)
	if err != nil {
		return err
	}
	// Deserialize `NewDuration`:
	err = decoder.Decode(&obj.NewDuration)
	if err != nil {
		return err
	}
	return nil
}

func (*MinDelayChangeEventData) isEventData() {}

var eventTypes = map[[8]byte]reflect.Type{
	BypasserCallExecutedEventDataDiscriminator:      reflect.TypeOf(BypasserCallExecutedEventData{}),
	CallExecutedEventDataDiscriminator:              reflect.TypeOf(CallExecutedEventData{}),
	CallScheduledEventDataDiscriminator:             reflect.TypeOf(CallScheduledEventData{}),
	CancelledEventDataDiscriminator:                 reflect.TypeOf(CancelledEventData{}),
	FunctionSelectorBlockedEventDataDiscriminator:   reflect.TypeOf(FunctionSelectorBlockedEventData{}),
	FunctionSelectorUnblockedEventDataDiscriminator: reflect.TypeOf(FunctionSelectorUnblockedEventData{}),
	MinDelayChangeEventDataDiscriminator:            reflect.TypeOf(MinDelayChangeEventData{}),
}
var eventNames = map[[8]byte]string{
	BypasserCallExecutedEventDataDiscriminator:      "BypasserCallExecuted",
	CallExecutedEventDataDiscriminator:              "CallExecuted",
	CallScheduledEventDataDiscriminator:             "CallScheduled",
	CancelledEventDataDiscriminator:                 "Cancelled",
	FunctionSelectorBlockedEventDataDiscriminator:   "FunctionSelectorBlocked",
	FunctionSelectorUnblockedEventDataDiscriminator: "FunctionSelectorUnblocked",
	MinDelayChangeEventDataDiscriminator:            "MinDelayChange",
}
var (
	_ *strings.Builder = nil
)
var (
	_ *base64.Encoding = nil
)
var (
	_ *ag_binary.Decoder = nil
)
var (
	_ *ag_rpc.GetTransactionResult = nil
)
var (
	_ *ag_base58.Alphabet = nil
)

type Event struct {
	Name string
	Data EventData
}

type EventData interface {
	UnmarshalWithDecoder(decoder *ag_binary.Decoder) error
	isEventData()
}

const eventLogPrefix = "Program data: "

func DecodeEvents(txData *ag_rpc.GetTransactionResult, targetProgramId ag_solanago.PublicKey, getAddressTables func(altAddresses []ag_solanago.PublicKey) (tables map[ag_solanago.PublicKey]ag_solanago.PublicKeySlice, err error)) (evts []*Event, err error) {
	var tx *ag_solanago.Transaction
	if tx, err = txData.Transaction.GetTransaction(); err != nil {
		return
	}

	altAddresses := make([]ag_solanago.PublicKey, len(tx.Message.AddressTableLookups))
	for i, alt := range tx.Message.AddressTableLookups {
		altAddresses[i] = alt.AccountKey
	}
	if len(altAddresses) > 0 {
		var tables map[ag_solanago.PublicKey]ag_solanago.PublicKeySlice
		if tables, err = getAddressTables(altAddresses); err != nil {
			return
		}
		tx.Message.SetAddressTables(tables)
		if err = tx.Message.ResolveLookups(); err != nil {
			return
		}
	}

	var base64Binaries [][]byte
	logMessageEventBinaries, err := decodeEventsFromLogMessage(txData.Meta.LogMessages)
	if err != nil {
		return
	}

	emitedCPIEventBinaries, err := decodeEventsFromEmitCPI(txData.Meta.InnerInstructions, tx.Message.AccountKeys, targetProgramId)
	if err != nil {
		return
	}

	base64Binaries = append(base64Binaries, logMessageEventBinaries...)
	base64Binaries = append(base64Binaries, emitedCPIEventBinaries...)
	evts, err = parseEvents(base64Binaries)
	return
}

func decodeEventsFromLogMessage(logMessages []string) (eventBinaries [][]byte, err error) {
	for _, log := range logMessages {
		if strings.HasPrefix(log, eventLogPrefix) {
			eventBase64 := log[len(eventLogPrefix):]

			var eventBinary []byte
			if eventBinary, err = base64.StdEncoding.DecodeString(eventBase64); err != nil {
				err = fmt.Errorf("failed to decode logMessage event: %s", eventBase64)
				return
			}
			eventBinaries = append(eventBinaries, eventBinary)
		}
	}
	return
}

func decodeEventsFromEmitCPI(InnerInstructions []ag_rpc.InnerInstruction, accountKeys ag_solanago.PublicKeySlice, targetProgramId ag_solanago.PublicKey) (eventBinaries [][]byte, err error) {
	for _, parsedIx := range InnerInstructions {
		for _, ix := range parsedIx.Instructions {
			if accountKeys[ix.ProgramIDIndex] != targetProgramId {
				continue
			}

			var ixData []byte
			if ixData, err = ag_base58.Decode(ix.Data.String()); err != nil {
				return
			}
			eventBase64 := base64.StdEncoding.EncodeToString(ixData[8:])
			var eventBinary []byte
			if eventBinary, err = base64.StdEncoding.DecodeString(eventBase64); err != nil {
				return
			}
			eventBinaries = append(eventBinaries, eventBinary)
		}
	}
	return
}

func parseEvents(base64Binaries [][]byte) (evts []*Event, err error) {
	decoder := ag_binary.NewDecoderWithEncoding(nil, ag_binary.EncodingBorsh)

	for _, eventBinary := range base64Binaries {
		eventDiscriminator := ag_binary.TypeID(eventBinary[:8])
		if eventType, ok := eventTypes[eventDiscriminator]; ok {
			eventData := reflect.New(eventType).Interface().(EventData)
			decoder.Reset(eventBinary)
			if err = eventData.UnmarshalWithDecoder(decoder); err != nil {
				err = fmt.Errorf("failed to unmarshal event %s: %w", eventType.String(), err)
				return
			}
			evts = append(evts, &Event{
				Name: eventNames[eventDiscriminator],
				Data: eventData,
			})
		}
	}
	return
}
