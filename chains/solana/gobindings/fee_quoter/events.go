// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

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

type DestChainAddedEventData struct {
	DestChainSelector uint64
	DestChainConfig   DestChainConfig
}

var DestChainAddedEventDataDiscriminator = [8]byte{59, 154, 48, 81, 230, 41, 80, 200}

func (obj DestChainAddedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DestChainAddedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DestChainConfig` param:
	err = encoder.Encode(obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainAddedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DestChainAddedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[59 154 48 81 230 41 80 200]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DestChainConfig`:
	err = decoder.Decode(&obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

func (*DestChainAddedEventData) isEventData() {}

type DestChainConfigUpdatedEventData struct {
	DestChainSelector uint64
	DestChainConfig   DestChainConfig
}

var DestChainConfigUpdatedEventDataDiscriminator = [8]byte{3, 141, 73, 190, 73, 231, 51, 80}

func (obj DestChainConfigUpdatedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DestChainConfigUpdatedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DestChainConfig` param:
	err = encoder.Encode(obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainConfigUpdatedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DestChainConfigUpdatedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[3 141 73 190 73 231 51 80]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DestChainConfig`:
	err = decoder.Decode(&obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

func (*DestChainConfigUpdatedEventData) isEventData() {}

type FeeTokenAddedEventData struct {
	FeeToken ag_solanago.PublicKey
	Enabled  bool
}

var FeeTokenAddedEventDataDiscriminator = [8]byte{181, 180, 252, 21, 215, 79, 93, 237}

func (obj FeeTokenAddedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FeeTokenAddedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeTokenAddedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FeeTokenAddedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[181 180 252 21 215 79 93 237]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

func (*FeeTokenAddedEventData) isEventData() {}

type FeeTokenDisabledEventData struct {
	FeeToken ag_solanago.PublicKey
}

var FeeTokenDisabledEventDataDiscriminator = [8]byte{34, 139, 66, 75, 30, 17, 45, 151}

func (obj FeeTokenDisabledEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FeeTokenDisabledEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeTokenDisabledEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FeeTokenDisabledEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[34 139 66 75 30 17 45 151]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (*FeeTokenDisabledEventData) isEventData() {}

type FeeTokenEnabledEventData struct {
	FeeToken ag_solanago.PublicKey
}

var FeeTokenEnabledEventDataDiscriminator = [8]byte{106, 180, 145, 189, 113, 180, 21, 15}

func (obj FeeTokenEnabledEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FeeTokenEnabledEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeTokenEnabledEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FeeTokenEnabledEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[106 180 145 189 113 180 21 15]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (*FeeTokenEnabledEventData) isEventData() {}

type FeeTokenRemovedEventData struct {
	FeeToken ag_solanago.PublicKey
}

var FeeTokenRemovedEventDataDiscriminator = [8]byte{40, 31, 230, 252, 183, 150, 147, 201}

func (obj FeeTokenRemovedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FeeTokenRemovedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FeeTokenRemovedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FeeTokenRemovedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[40 31 230 252 183 150 147 201]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	return nil
}

func (*FeeTokenRemovedEventData) isEventData() {}

type OwnershipTransferRequestedEventData struct {
	From ag_solanago.PublicKey
	To   ag_solanago.PublicKey
}

var OwnershipTransferRequestedEventDataDiscriminator = [8]byte{79, 54, 99, 123, 57, 244, 134, 35}

func (obj OwnershipTransferRequestedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(OwnershipTransferRequestedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `From` param:
	err = encoder.Encode(obj.From)
	if err != nil {
		return err
	}
	// Serialize `To` param:
	err = encoder.Encode(obj.To)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OwnershipTransferRequestedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(OwnershipTransferRequestedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[79 54 99 123 57 244 134 35]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `From`:
	err = decoder.Decode(&obj.From)
	if err != nil {
		return err
	}
	// Deserialize `To`:
	err = decoder.Decode(&obj.To)
	if err != nil {
		return err
	}
	return nil
}

func (*OwnershipTransferRequestedEventData) isEventData() {}

type OwnershipTransferredEventData struct {
	From ag_solanago.PublicKey
	To   ag_solanago.PublicKey
}

var OwnershipTransferredEventDataDiscriminator = [8]byte{172, 61, 205, 183, 250, 50, 38, 98}

func (obj OwnershipTransferredEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(OwnershipTransferredEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `From` param:
	err = encoder.Encode(obj.From)
	if err != nil {
		return err
	}
	// Serialize `To` param:
	err = encoder.Encode(obj.To)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OwnershipTransferredEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(OwnershipTransferredEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[172 61 205 183 250 50 38 98]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `From`:
	err = decoder.Decode(&obj.From)
	if err != nil {
		return err
	}
	// Deserialize `To`:
	err = decoder.Decode(&obj.To)
	if err != nil {
		return err
	}
	return nil
}

func (*OwnershipTransferredEventData) isEventData() {}

type PriceUpdaterAddedEventData struct {
	PriceUpdater ag_solanago.PublicKey
}

var PriceUpdaterAddedEventDataDiscriminator = [8]byte{87, 31, 151, 133, 151, 187, 97, 186}

func (obj PriceUpdaterAddedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PriceUpdaterAddedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `PriceUpdater` param:
	err = encoder.Encode(obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PriceUpdaterAddedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PriceUpdaterAddedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[87 31 151 133 151 187 97 186]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PriceUpdater`:
	err = decoder.Decode(&obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}

func (*PriceUpdaterAddedEventData) isEventData() {}

type PriceUpdaterRemovedEventData struct {
	PriceUpdater ag_solanago.PublicKey
}

var PriceUpdaterRemovedEventDataDiscriminator = [8]byte{225, 194, 40, 213, 212, 39, 76, 148}

func (obj PriceUpdaterRemovedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PriceUpdaterRemovedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `PriceUpdater` param:
	err = encoder.Encode(obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PriceUpdaterRemovedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PriceUpdaterRemovedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[225 194 40 213 212 39 76 148]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PriceUpdater`:
	err = decoder.Decode(&obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}

func (*PriceUpdaterRemovedEventData) isEventData() {}

type TokenTransferFeeConfigUpdatedEventData struct {
	DestChainSelector      uint64
	Token                  ag_solanago.PublicKey
	TokenTransferFeeConfig TokenTransferFeeConfig
}

var TokenTransferFeeConfigUpdatedEventDataDiscriminator = [8]byte{253, 199, 166, 1, 178, 150, 242, 253}

func (obj TokenTransferFeeConfigUpdatedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenTransferFeeConfigUpdatedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `TokenTransferFeeConfig` param:
	err = encoder.Encode(obj.TokenTransferFeeConfig)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenTransferFeeConfigUpdatedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TokenTransferFeeConfigUpdatedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[253 199 166 1 178 150 242 253]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `TokenTransferFeeConfig`:
	err = decoder.Decode(&obj.TokenTransferFeeConfig)
	if err != nil {
		return err
	}
	return nil
}

func (*TokenTransferFeeConfigUpdatedEventData) isEventData() {}

type UsdPerTokenUpdatedEventData struct {
	Token     ag_solanago.PublicKey
	Value     [28]uint8
	Timestamp int64
}

var UsdPerTokenUpdatedEventDataDiscriminator = [8]byte{67, 154, 252, 56, 104, 14, 192, 219}

func (obj UsdPerTokenUpdatedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(UsdPerTokenUpdatedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Value` param:
	err = encoder.Encode(obj.Value)
	if err != nil {
		return err
	}
	// Serialize `Timestamp` param:
	err = encoder.Encode(obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UsdPerTokenUpdatedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(UsdPerTokenUpdatedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[67 154 252 56 104 14 192 219]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Value`:
	err = decoder.Decode(&obj.Value)
	if err != nil {
		return err
	}
	// Deserialize `Timestamp`:
	err = decoder.Decode(&obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (*UsdPerTokenUpdatedEventData) isEventData() {}

type UsdPerUnitGasUpdatedEventData struct {
	DestChain uint64
	Value     [28]uint8
	Timestamp int64
}

var UsdPerUnitGasUpdatedEventDataDiscriminator = [8]byte{174, 255, 2, 41, 197, 110, 31, 40}

func (obj UsdPerUnitGasUpdatedEventData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(UsdPerUnitGasUpdatedEventDataDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DestChain` param:
	err = encoder.Encode(obj.DestChain)
	if err != nil {
		return err
	}
	// Serialize `Value` param:
	err = encoder.Encode(obj.Value)
	if err != nil {
		return err
	}
	// Serialize `Timestamp` param:
	err = encoder.Encode(obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UsdPerUnitGasUpdatedEventData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(UsdPerUnitGasUpdatedEventDataDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[174 255 2 41 197 110 31 40]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DestChain`:
	err = decoder.Decode(&obj.DestChain)
	if err != nil {
		return err
	}
	// Deserialize `Value`:
	err = decoder.Decode(&obj.Value)
	if err != nil {
		return err
	}
	// Deserialize `Timestamp`:
	err = decoder.Decode(&obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (*UsdPerUnitGasUpdatedEventData) isEventData() {}

var eventTypes = map[[8]byte]reflect.Type{
	DestChainAddedEventDataDiscriminator:                reflect.TypeOf(DestChainAddedEventData{}),
	DestChainConfigUpdatedEventDataDiscriminator:        reflect.TypeOf(DestChainConfigUpdatedEventData{}),
	FeeTokenAddedEventDataDiscriminator:                 reflect.TypeOf(FeeTokenAddedEventData{}),
	FeeTokenDisabledEventDataDiscriminator:              reflect.TypeOf(FeeTokenDisabledEventData{}),
	FeeTokenEnabledEventDataDiscriminator:               reflect.TypeOf(FeeTokenEnabledEventData{}),
	FeeTokenRemovedEventDataDiscriminator:               reflect.TypeOf(FeeTokenRemovedEventData{}),
	OwnershipTransferRequestedEventDataDiscriminator:    reflect.TypeOf(OwnershipTransferRequestedEventData{}),
	OwnershipTransferredEventDataDiscriminator:          reflect.TypeOf(OwnershipTransferredEventData{}),
	PriceUpdaterAddedEventDataDiscriminator:             reflect.TypeOf(PriceUpdaterAddedEventData{}),
	PriceUpdaterRemovedEventDataDiscriminator:           reflect.TypeOf(PriceUpdaterRemovedEventData{}),
	TokenTransferFeeConfigUpdatedEventDataDiscriminator: reflect.TypeOf(TokenTransferFeeConfigUpdatedEventData{}),
	UsdPerTokenUpdatedEventDataDiscriminator:            reflect.TypeOf(UsdPerTokenUpdatedEventData{}),
	UsdPerUnitGasUpdatedEventDataDiscriminator:          reflect.TypeOf(UsdPerUnitGasUpdatedEventData{}),
}
var eventNames = map[[8]byte]string{
	DestChainAddedEventDataDiscriminator:                "DestChainAdded",
	DestChainConfigUpdatedEventDataDiscriminator:        "DestChainConfigUpdated",
	FeeTokenAddedEventDataDiscriminator:                 "FeeTokenAdded",
	FeeTokenDisabledEventDataDiscriminator:              "FeeTokenDisabled",
	FeeTokenEnabledEventDataDiscriminator:               "FeeTokenEnabled",
	FeeTokenRemovedEventDataDiscriminator:               "FeeTokenRemoved",
	OwnershipTransferRequestedEventDataDiscriminator:    "OwnershipTransferRequested",
	OwnershipTransferredEventDataDiscriminator:          "OwnershipTransferred",
	PriceUpdaterAddedEventDataDiscriminator:             "PriceUpdaterAdded",
	PriceUpdaterRemovedEventDataDiscriminator:           "PriceUpdaterRemoved",
	TokenTransferFeeConfigUpdatedEventDataDiscriminator: "TokenTransferFeeConfigUpdated",
	UsdPerTokenUpdatedEventDataDiscriminator:            "UsdPerTokenUpdated",
	UsdPerUnitGasUpdatedEventDataDiscriminator:          "UsdPerUnitGasUpdated",
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
