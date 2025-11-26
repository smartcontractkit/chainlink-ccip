// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package committee_verifier

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type BaseVerifierAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type BaseVerifierDestChainConfigArgs struct {
	Router             common.Address
	DestChainSelector  uint64
	AllowlistEnabled   bool
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type CommitteeVerifierDynamicConfig struct {
	FeeAggregator  common.Address
	AllowlistAdmin common.Address
}

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	ExecutionGasLimit   uint32
	CcipReceiveGasLimit uint32
	Finality            uint16
	CcvAndExecutorHash  [32]byte
	OnRampAddress       []byte
	OffRampAddress      []byte
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []MessageV1CodecTokenTransferV1
	Data                []byte
}

type MessageV1CodecTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	TokenReceiver      []byte
	ExtraData          []byte
}

type SignatureQuorumValidatorSignersUpdate struct {
	SourceChainSelector uint64
	Threshold           uint8
	Signers             []common.Address
}

var CommitteeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"storageLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.DestChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySignersUpdates\",\"inputs\":[{\"name\":\"sourceChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"signersUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct SignatureQuorumValidator.SignersUpdate[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSignatureConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"sourceChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"signerSets\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"thresholds\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignatureConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStorageLocation\",\"inputs\":[{\"name\":\"newLocation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"versionTag\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct CommitteeVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignatureConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationUpdated\",\"inputs\":[{\"name\":\"oldLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newLocation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"verifierVersion\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSignersSetForSource\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NonOrderedOrNonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]}]",
	Bin: "0x60c060405234610374576137668038038061001981610379565b928339810181810360608112610374576040136103745760408051919082016001600160401b0381118382101761034d576040526100568361039e565b82526100646020840161039e565b60208301908152604084015190936001600160401b03821161037457019080601f83011215610374578151916001600160401b03831161034d576100b1601f8401601f1916602001610379565b92808452602084019260208284010111610374578260206100d293016103b2565b331561036357600180546001600160a01b031916331790554660805281516001600160401b03811161034d57600654600181811c91168015610343575b602082101461032d57601f81116102c8575b506020601f821160011461024c579261018f928260809360008051602061374683398151915296600091610241575b508160011b916000199060031b1c1916176006555b604051938492604084526000604085015260606020850152518092816060860152858501906103b2565b601f01601f19168101030190a180516001600160a01b0316156102305751600780546001600160a01b03199081166001600160a01b0393841690811790925583516008805490921690841617905560408051918252925190911660208201527f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd9190a160405161337090816103d6823960805181610cba015260a051815050f35b6306b7c75960e31b60005260046000fd5b905083015138610150565b601f198216906006600052806000209160005b8181106102b0575083600080516020613746833981519152969361018f969360809660019410610297575b5050811b01600655610165565b85015160001960f88460031b161c19169055388061028a565b9192602060018192868a01518155019401920161025f565b60066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f601f830160051c81019160208410610323575b601f0160051c01905b8181106103175750610121565b6000815560010161030a565b9091508190610301565b634e487b7160e01b600052602260045260246000fd5b90607f169061010f565b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761034d57604052565b51906001600160a01b038216820361037457565b60005b8381106103c55750506000910152565b81810151838201526020016103b556fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461297057508063181f5a77146128f35780631cafd1df146124135780633a3d72b5146123625780633bbbed4b1461204a5780635cb80c5d14611d5e5780636def4ce714611c835780637437ff9f14611bc257806379ba509714611ad957806380485e2514611846578063869b7f62146116bd5780638da5cb5b1461166b578063aa8dac9814611322578063b2bd751c14610f79578063bff0ec1d14610b15578063c9b146b3146106e7578063ceac5cee1461036a578063f2fde38b1461027a578063fe163eed146102215763fec888af146100fb57600080fd5b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c57604051600060065461013b81612e0e565b80845290600181169081156101da575060011461017b575b6101778361016381850382612a48565b604051918291602083526020830190612ac3565b0390f35b600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f939250905b8082106101c057509091508101602001610163610153565b9192600181602092548385880101520191019092916101a8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b840190910191506101639050610153565b600080fd5b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760206040517f49ff34ed000000000000000000000000000000000000000000000000000000008152f35b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5773ffffffffffffffffffffffffffffffffffffffff6102c6612bba565b6102ce612e61565b1633811461034057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c573660238201121561021c576103ca903690602481600401359101612cfc565b6103d2612e61565b6040516040815260006006546103e781612e0e565b908160408501526001811690816000146106a7575060011461064a575b50908061043a83827fbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd89503602084015285612ac3565b0390a1805167ffffffffffffffff811161061b57610459600654612e0e565b601f8111610578575b50602091601f82116001146104be579181926000926104b3575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c191617600655600080f35b01519050828061047c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169260066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f9160005b85811061056057508360019510610529575b505050811b01600655005b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905582808061051e565b9192602060018192868501518155019401920161050c565b6006600052601f820160051c7ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0190602083106105f3575b601f0160051c7ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01905b8181106105e75750610462565b600081556001016105da565b7ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f91506105b0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600660009081527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f939250905b80821061068d5750909150810160600181610404565b835483830160600152600190930192602090910190610677565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660608086019190915291151560051b84019091019150829050610404565b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c57610736903690600401612b22565b73ffffffffffffffffffffffffffffffffffffffff600154163303610acb575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610ac9576000938060051b84013582811215610ac557840191608083360312610ac557604051946080860186811067ffffffffffffffff821117610a98576040526107ce84612b82565b86526107dc60208501612dca565b9660208701978852604085013567ffffffffffffffff8111610a94576108059036908701612bfe565b9460408801958652606081013567ffffffffffffffff8111610a905761082d91369101612bfe565b946060880195865267ffffffffffffffff8851168252600560205260408220985115156108a7818b907fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7eff000000000000000000000000000000000000000000000000000000000000835492151560f01b169116179055565b815151610964575b50959760010195505b845180518210156108f757906108f073ffffffffffffffffffffffffffffffffffffffff6108e883600195612d62565b5116886130aa565b50016108b8565b5050959490935060019251908151610915575b505001929190610780565b61095a67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190612c63565b0390a2858061090a565b9893959290949798969196600014610a5957600184019591875b865180518210156109fb576109a88273ffffffffffffffffffffffffffffffffffffffff92612d62565b511680156109c457906109bd6001928a61323e565b500161097e565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281610a4f67ffffffffffffffff8b51169251604051918291602083526020830190612c63565b0390a290896108af565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60085416330315610756577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461021c5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261021c5760443567ffffffffffffffff811161021c57610ba3903690600401612cce565b919060068310610f21578260041161021c577fffffffff00000000000000000000000000000000000000000000000000000000813516907f49ff34ed000000000000000000000000000000000000000000000000000000008203610f4b578360061161021c57600481013560f01c92610c1b84612de8565b8510610f2157610c30610c5b91600401612db5565b9260405160208101918252602435602482015260248152610c52604482612a48565b51902093612de8565b9360009085600611610a94578511610f1e57507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa600667ffffffffffffffff92019401911690816000526002602052604060002091825415610ef157507f0000000000000000000000000000000000000000000000000000000000000000468103610ec0575060ff60028301541692838260061c10610e9657600091825b858410610d0257005b8360061b84810460401485151715610e67576020810190818111610e6757610d35610d2f8383878d612df6565b90612eac565b9060009260408201809211610e3a57602092610d5a610d2f86946080948f8b90612df6565b60405191898352601b868401526040830152606082015282805260015afa15610e2e5773ffffffffffffffffffffffffffffffffffffffff81511691610db0838860019160005201602052604060002054151590565b15610e065773ffffffffffffffffffffffffffffffffffffffff16821115610dde5750600190930192610cf9565b807fb70ad94b0000000000000000000000000000000000000000000000000000000060049252fd5b6004827fca31867a000000000000000000000000000000000000000000000000000000008152fd5b604051903d90823e3d90fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f71253a250000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0f01ce85000000000000000000000000000000000000000000000000000000006000526004524660245260446000fd5b7fe39679910000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80fd5b7fbba6473c0000000000000000000000000000000000000000000000000000000060005260046000fd5b507fef8a07ee0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c573660238201121561021c57806004013567ffffffffffffffff811161021c57602460c0820283010136811161021c57610ff1612e61565b610ffa82612b53565b916110086040519384612a48565b825260009260240190602083015b818310611239578480855b8051831015611235576110348382612d62565b519267ffffffffffffffff602061104b8385612d62565b51015116938415611209578484526005602052604080852082518154928401517fff00ffffffffffffffffffff000000000000000000000000000000000000000090931673ffffffffffffffffffffffffffffffffffffffff919091161791151560f01b7eff0000000000000000000000000000000000000000000000000000000000001691909117815590606081015182546080830163ffffffff815116156111dd5773ffffffffffffffffffffffffffffffffffffffff7f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c946040946001999a9b979479ffffffff0000000000000000000000000000000000000000000060ff955160b01b16907fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000007dffffffff000000000000000000000000000000000000000000000000000060a087015160d01b169460a01b169116171717809455511691835192835260f01c1615156020820152a2019190611021565b6024888a7f9e720551000000000000000000000000000000000000000000000000000000008252600452fd5b602484867fc35aa79d000000000000000000000000000000000000000000000000000000008252600452fd5b5080f35b60c08336031261131e576040519060c0820182811067ffffffffffffffff8211176112f157604052833573ffffffffffffffffffffffffffffffffffffffff811681036112ed57825261128e60208501612b82565b602083015261129f60408501612dca565b604083015260608401359061ffff821682036112ed5782602092606060c09501526112cc60808701612dd7565b60808201526112dd60a08701612dd7565b60a0820152815201920191611016565b8680fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8480fd5b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760405180816020600354928381520160036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9260005b81811061165257505061139f92500382612a48565b8051906113c46113ae83612b53565b926113bc6040519485612a48565b808452612b53565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833780517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061141e6113ae83612b53565b0160005b81811061164157505081519161145061143a84612b53565b936114486040519586612a48565b808552612b53565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208501910136823760005b82518110156115315767ffffffffffffffff6114998285612d62565b5116806114a6838a612d62565b52600052600260205260406000206040518060208354918281520190836000526020600020906000905b8082106115195750505060019392826114ef60ff946002940382612a48565b6114f9858a612d62565b526115048489612d62565b500154166115128288612d62565b520161147d565b909192600160208192865481520194019201906114d0565b85878386886040519485946060860190606087525180915260808601919060005b81811061161e575050508481036020860152825180825260208201916020808360051b8301019501926000915b8383106115cb5750505050506020908483036040860152519182815201919060005b8181106115af575050500390f35b825160ff168452859450602093840193909201916001016115a1565b919396975091939460208061160a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528a51612c63565b98019301930190928897969594929361157f565b825167ffffffffffffffff16845288975060209384019390920191600101611552565b806060602080938601015201611422565b845483526001948501948694506020909301920161138a565b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461021c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760006040516116fa81612a2c565b611702612bba565b815260243573ffffffffffffffffffffffffffffffffffffffff81168103610a905760208201908152611733612e61565b73ffffffffffffffffffffffffffffffffffffffff8251161561181e578173ffffffffffffffffffffffffffffffffffffffff61181892817f781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd9551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600754161760075551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600854161760085560405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b0390a180f35b6004837f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b3461021c5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5761187d612b6b565b60243567ffffffffffffffff811161021c5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261021c576040519060a0820182811067ffffffffffffffff82111761061b57604052806004013567ffffffffffffffff811161021c576118fd9060043691840101612d33565b8252602481013567ffffffffffffffff811161021c576119239060043691840101612d33565b6020830152604481013567ffffffffffffffff811161021c5781013660238201121561021c57600481013561195781612b53565b916119656040519384612a48565b818352602060048185019360061b830101019036821161021c57602401915b818310611aa157505050604083015261199f60648201612bdd565b6060830152608481013567ffffffffffffffff811161021c5760809160046119ca9236920101612d33565b9101526044359067ffffffffffffffff821161021c576119f767ffffffffffffffff923690600401612d33565b50611a00612d51565b501680600052600560205273ffffffffffffffffffffffffffffffffffffffff6040600020541615611a745760009081526005602090815260409182902054825161ffff60a083901c16815263ffffffff60b083901c81169382019390935260d09190911c90911691810191909152606090f35b7f8a4e93c90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60408336031261021c5760206040918251611abb81612a2c565b611ac486612bdd565b81528286013583820152815201920191611984565b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760005473ffffffffffffffffffffffffffffffffffffffff81163303611b98577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760006020604051611c0181612a2c565b8281520152610177604051611c1581612a2c565b73ffffffffffffffffffffffffffffffffffffffff60075416815273ffffffffffffffffffffffffffffffffffffffff60085416602082015260405191829182919091602073ffffffffffffffffffffffffffffffffffffffff816040840195828151168552015116910152565b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5767ffffffffffffffff611cc3612b6b565b16600052600560205260406000206001815491019060405191826020825491828152019160005260206000209060005b818110611d485773ffffffffffffffffffffffffffffffffffffffff8561017788611d2081890382612a48565b604051938360ff869560f01c1615158552166020840152606060408401526060830190612c63565b8254845260209093019260019283019201611cf3565b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c57611dad903690600401612b22565b9073ffffffffffffffffffffffffffffffffffffffff600754169160005b818110611dd457005b611ddf818385612da5565b359073ffffffffffffffffffffffffffffffffffffffff821680920361021c57604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa92831561203e5760009361200b575b5082611e56575b506001915001611dcb565b8560405193611f1360208601957fa9059cbb00000000000000000000000000000000000000000000000000000000875283602482015282604482015260448152611ea1606482612a48565b600080604098895193611eb48b86612a48565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082895af13d15612003573d90611ef582612a89565b91611f028a519384612a48565b82523d6000602084013e5b86613293565b805180611f4f575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a385611e4b565b81929495969350906020918101031261021c576020015180159081150361021c57611f805792919086908880611f1b565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090611f0d565b90926020823d8211612036575b8161202560209383612a48565b81010312610f1e5750519186611e44565b3d9150612018565b6040513d6000823e3d90fd5b3461021c5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c578036036101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261021c576120c1612b97565b5060843567ffffffffffffffff811161021c576120e2903690600401612cce565b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd610124830135910181121561021c57810160048101359067ffffffffffffffff821161021c5760240190803603821361021c5760246121829167ffffffffffffffff93357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116916014811061232d575b505060601c9301612db5565b16806000526005602052604060002090815490604051907fa8d87a3b000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff86165afa90811561203e576000916122ca575b5073ffffffffffffffffffffffffffffffffffffffff16330361229c5760f01c60ff16612252575b6101776040517f49ff34ed00000000000000000000000000000000000000000000000000000000602082015260048152610163602482612a48565b6000828152600290910160205260409020541561226f5780612217565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011612325575b816122e360209383612a48565b81010312610a9457519073ffffffffffffffffffffffffffffffffffffffff82168203610f1e575073ffffffffffffffffffffffffffffffffffffffff6121ef565b3d91506122d6565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009250829060140360031b1b16168580612176565b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5767ffffffffffffffff6123a2612b6b565b166000526002602052604060002060405190818092602083549182815201908360005260206000209060005b8181106123fa578460ff6002886123e784890385612a48565b0154169061017760405192839283612cad565b82548452869450602090930192600192830192016123ce565b3461021c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5760043567ffffffffffffffff811161021c57612462903690600401612b22565b906024359067ffffffffffffffff821161021c573660238301121561021c57816004013561248f81612b53565b9261249d6040519485612a48565b8184526024602085019260051b8201019036821161021c5760248101925b82841061283b578585886124cd612e61565b6000925b80518410156126c7576124e48482612d62565b5190602082019260ff84511680159081156126b8575b506125cd5767ffffffffffffffff83511660005260026020526040600020945b8554156125565761255073ffffffffffffffffffffffffffffffffffffffff61254288612ee7565b90549060031b1c16876130aa565b5061251a565b939560408401959490929160005b875180518210156126215761258e8273ffffffffffffffffffffffffffffffffffffffff92612d62565b5116156125f7576125c073ffffffffffffffffffffffffffffffffffffffff6125b8838b51612d62565b51168861323e565b156125cd57600101612564565b7f12823a5e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fd6c62c9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50509660ff67ffffffffffffffff6001967f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b8394989995996002848651169101907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055612694828251166131de565b5051169351915116906126ac60405192839283612cad565b0390a2019290916124d1565b905060408401515110876124fa565b5060005b8181106126d457005b604067ffffffffffffffff6126f26126ed848688612da5565b612db5565b1660009081526004602052205461270c575b6001016126cb565b9067ffffffffffffffff6127246126ed848487612da5565b1660005260026020526040600020925b8354156127705761276a73ffffffffffffffffffffffffffffffffffffffff61275c86612ee7565b90549060031b1c16856130aa565b50612734565b9291600260019293017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556127bf67ffffffffffffffff6127b96126ed848789612da5565b16612f14565b506127ce6126ed828587612da5565b7f3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b8361282b60209267ffffffffffffffff6040519161280c8684612a48565b6000835260003681376000604051948594604086526040860190612c63565b9684015216930390a29050612704565b833567ffffffffffffffff811161021c57820160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261021c57604051916060830183811067ffffffffffffffff82111761061b576040526128a460248301612b82565b8352604482013560ff8116810361021c57602084015260648201359267ffffffffffffffff841161021c576128e3602094936024869536920101612bfe565b60408201528152019301926124bb565b3461021c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c5761017760408051906129348183612a48565b601b82527f436f6d6d6974746565566572696669657220312e372e302d6465760000000000602083015251918291602083526020830190612ac3565b3461021c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021c57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361021c57817ffacbd7dc0000000000000000000000000000000000000000000000000000000060209314908115612a02575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836129fb565b6040810190811067ffffffffffffffff82111761061b57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761061b57604052565b67ffffffffffffffff811161061b57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110612b0d5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612ace565b9181601f8401121561021c5782359167ffffffffffffffff831161021c576020808501948460051b01011161021c57565b67ffffffffffffffff811161061b5760051b60200190565b6004359067ffffffffffffffff8216820361021c57565b359067ffffffffffffffff8216820361021c57565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361021c57565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361021c57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361021c57565b9080601f8301121561021c578135612c1581612b53565b92612c236040519485612a48565b81845260208085019260051b82010192831161021c57602001905b828210612c4b5750505090565b60208091612c5884612bdd565b815201910190612c3e565b906020808351928381520192019060005b818110612c815750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612c74565b9060ff612cc7602092959495604085526040850190612c63565b9416910152565b9181601f8401121561021c5782359167ffffffffffffffff831161021c576020838186019501011161021c57565b929192612d0882612a89565b91612d166040519384612a48565b82948184528183011161021c578281602093846000960137010152565b9080601f8301121561021c57816020612d4e93359101612cfc565b90565b6064359061ffff8216820361021c57565b8051821015612d765760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015612d765760051b0190565b3567ffffffffffffffff8116810361021c5790565b3590811515820361021c57565b359063ffffffff8216820361021c57565b6006019081600611610e6757565b9093929384831161021c57841161021c578101920390565b90600182811c92168015612e57575b6020831014612e2857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612e1d565b73ffffffffffffffffffffffffffffffffffffffff600154163303612e8257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110612eba575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b805415612d7657600052602060002090600090565b8054821015612d765760005260206000200190600090565b60008181526004602052604090205480156130a3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610e6757600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610e6757818103613034575b5050506003548015613005577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612fc2816003612efc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61308b613045613056936003612efc565b90549060031b1c9283926003612efc565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080612f89565b5050600090565b90600182019181600052826020526040600020548015156000146131d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610e67578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610e675781810361319e575b50505080548015613005577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061315f8282612efc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6131be6131ae6130569386612efc565b90549060031b1c92839286612efc565b905560005283602052604060002055388080613127565b50505050600090565b80600052600460205260406000205415600014613238576003546801000000000000000081101561061b5761321f6130568260018594016003556003612efc565b9055600354906000526004602052604060002055600190565b50600090565b60008281526001820160205260409020546130a3578054906801000000000000000082101561061b578261327c613056846001809601855584612efc565b905580549260005201602052604060002055600190565b9192901561330e57508151156132a7575090565b3b156132b05790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156133215750805190602001fd5b61335f906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612ac3565b0390fdfea164736f6c634300081a000abea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8",
}

var CommitteeVerifierABI = CommitteeVerifierMetaData.ABI

var CommitteeVerifierBin = CommitteeVerifierMetaData.Bin

func DeployCommitteeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig CommitteeVerifierDynamicConfig, storageLocation string) (common.Address, *types.Transaction, *CommitteeVerifier, error) {
	parsed, err := CommitteeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitteeVerifierBin), backend, dynamicConfig, storageLocation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitteeVerifier{address: address, abi: *parsed, CommitteeVerifierCaller: CommitteeVerifierCaller{contract: contract}, CommitteeVerifierTransactor: CommitteeVerifierTransactor{contract: contract}, CommitteeVerifierFilterer: CommitteeVerifierFilterer{contract: contract}}, nil
}

type CommitteeVerifier struct {
	address common.Address
	abi     abi.ABI
	CommitteeVerifierCaller
	CommitteeVerifierTransactor
	CommitteeVerifierFilterer
}

type CommitteeVerifierCaller struct {
	contract *bind.BoundContract
}

type CommitteeVerifierTransactor struct {
	contract *bind.BoundContract
}

type CommitteeVerifierFilterer struct {
	contract *bind.BoundContract
}

type CommitteeVerifierSession struct {
	Contract     *CommitteeVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitteeVerifierCallerSession struct {
	Contract *CommitteeVerifierCaller
	CallOpts bind.CallOpts
}

type CommitteeVerifierTransactorSession struct {
	Contract     *CommitteeVerifierTransactor
	TransactOpts bind.TransactOpts
}

type CommitteeVerifierRaw struct {
	Contract *CommitteeVerifier
}

type CommitteeVerifierCallerRaw struct {
	Contract *CommitteeVerifierCaller
}

type CommitteeVerifierTransactorRaw struct {
	Contract *CommitteeVerifierTransactor
}

func NewCommitteeVerifier(address common.Address, backend bind.ContractBackend) (*CommitteeVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(CommitteeVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitteeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifier{address: address, abi: abi, CommitteeVerifierCaller: CommitteeVerifierCaller{contract: contract}, CommitteeVerifierTransactor: CommitteeVerifierTransactor{contract: contract}, CommitteeVerifierFilterer: CommitteeVerifierFilterer{contract: contract}}, nil
}

func NewCommitteeVerifierCaller(address common.Address, caller bind.ContractCaller) (*CommitteeVerifierCaller, error) {
	contract, err := bindCommitteeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierCaller{contract: contract}, nil
}

func NewCommitteeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeVerifierTransactor, error) {
	contract, err := bindCommitteeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierTransactor{contract: contract}, nil
}

func NewCommitteeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeVerifierFilterer, error) {
	contract, err := bindCommitteeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierFilterer{contract: contract}, nil
}

func bindCommitteeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitteeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeVerifier.Contract.CommitteeVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.CommitteeVerifierTransactor.contract.Transfer(opts)
}

func (_CommitteeVerifier *CommitteeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.CommitteeVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitteeVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.contract.Transfer(opts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "forwardToVerifier", message, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _CommitteeVerifier.Contract.ForwardToVerifier(&_CommitteeVerifier.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) ForwardToVerifier(message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return _CommitteeVerifier.Contract.ForwardToVerifier(&_CommitteeVerifier.CallOpts, message, arg1, arg2, arg3, arg4)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetAllSignatureConfigs(opts *bind.CallOpts) (GetAllSignatureConfigs,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getAllSignatureConfigs")

	outstruct := new(GetAllSignatureConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainSelectors = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.SignerSets = *abi.ConvertType(out[1], new([][]common.Address)).(*[][]common.Address)
	outstruct.Thresholds = *abi.ConvertType(out[2], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetAllSignatureConfigs() (GetAllSignatureConfigs,

	error) {
	return _CommitteeVerifier.Contract.GetAllSignatureConfigs(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetAllSignatureConfigs() (GetAllSignatureConfigs,

	error) {
	return _CommitteeVerifier.Contract.GetAllSignatureConfigs(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitteeVerifier.Contract.GetDestChainConfig(&_CommitteeVerifier.CallOpts, destChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitteeVerifier.Contract.GetDestChainConfig(&_CommitteeVerifier.CallOpts, destChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitteeVerifierDynamicConfig, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitteeVerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitteeVerifierDynamicConfig)).(*CommitteeVerifierDynamicConfig)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetDynamicConfig() (CommitteeVerifierDynamicConfig, error) {
	return _CommitteeVerifier.Contract.GetDynamicConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetDynamicConfig() (CommitteeVerifierDynamicConfig, error) {
	return _CommitteeVerifier.Contract.GetDynamicConfig(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _CommitteeVerifier.Contract.GetFee(&_CommitteeVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getSignatureConfig", sourceChainSelector)

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetSignatureConfig(sourceChainSelector uint64) ([]common.Address, uint8, error) {
	return _CommitteeVerifier.Contract.GetSignatureConfig(&_CommitteeVerifier.CallOpts, sourceChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetSignatureConfig(sourceChainSelector uint64) ([]common.Address, uint8, error) {
	return _CommitteeVerifier.Contract.GetSignatureConfig(&_CommitteeVerifier.CallOpts, sourceChainSelector)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) GetStorageLocation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "getStorageLocation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) GetStorageLocation() (string, error) {
	return _CommitteeVerifier.Contract.GetStorageLocation(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) GetStorageLocation() (string, error) {
	return _CommitteeVerifier.Contract.GetStorageLocation(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) Owner() (common.Address, error) {
	return _CommitteeVerifier.Contract.Owner(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) Owner() (common.Address, error) {
	return _CommitteeVerifier.Contract.Owner(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeVerifier.Contract.SupportsInterface(&_CommitteeVerifier.CallOpts, interfaceId)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CommitteeVerifier.Contract.SupportsInterface(&_CommitteeVerifier.CallOpts, interfaceId)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) TypeAndVersion() (string, error) {
	return _CommitteeVerifier.Contract.TypeAndVersion(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) TypeAndVersion() (string, error) {
	return _CommitteeVerifier.Contract.TypeAndVersion(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "verifyMessage", message, messageHash, ccvData)

	if err != nil {
		return err
	}

	return err

}

func (_CommitteeVerifier *CommitteeVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitteeVerifier.Contract.VerifyMessage(&_CommitteeVerifier.CallOpts, message, messageHash, ccvData)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) VerifyMessage(message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error {
	return _CommitteeVerifier.Contract.VerifyMessage(&_CommitteeVerifier.CallOpts, message, messageHash, ccvData)
}

func (_CommitteeVerifier *CommitteeVerifierCaller) VersionTag(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CommitteeVerifier.contract.Call(opts, &out, "versionTag")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CommitteeVerifier *CommitteeVerifierSession) VersionTag() ([4]byte, error) {
	return _CommitteeVerifier.Contract.VersionTag(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierCallerSession) VersionTag() ([4]byte, error) {
	return _CommitteeVerifier.Contract.VersionTag(&_CommitteeVerifier.CallOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_CommitteeVerifier *CommitteeVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptOwnership(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.AcceptOwnership(&_CommitteeVerifier.TransactOpts)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyAllowlistUpdates(&_CommitteeVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyAllowlistUpdates(&_CommitteeVerifier.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyDestChainConfigUpdates(&_CommitteeVerifier.TransactOpts, destChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplyDestChainConfigUpdates(&_CommitteeVerifier.TransactOpts, destChainConfigArgs)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) ApplySignersUpdates(opts *bind.TransactOpts, sourceChainSelectorsToRemove []uint64, signersUpdates []SignatureQuorumValidatorSignersUpdate) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "applySignersUpdates", sourceChainSelectorsToRemove, signersUpdates)
}

func (_CommitteeVerifier *CommitteeVerifierSession) ApplySignersUpdates(sourceChainSelectorsToRemove []uint64, signersUpdates []SignatureQuorumValidatorSignersUpdate) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplySignersUpdates(&_CommitteeVerifier.TransactOpts, sourceChainSelectorsToRemove, signersUpdates)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) ApplySignersUpdates(sourceChainSelectorsToRemove []uint64, signersUpdates []SignatureQuorumValidatorSignersUpdate) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.ApplySignersUpdates(&_CommitteeVerifier.TransactOpts, sourceChainSelectorsToRemove, signersUpdates)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierSession) SetDynamicConfig(dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetDynamicConfig(&_CommitteeVerifier.TransactOpts, dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) SetDynamicConfig(dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.SetDynamicConfig(&_CommitteeVerifier.TransactOpts, dynamicConfig)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitteeVerifier *CommitteeVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferOwnership(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.TransferOwnership(&_CommitteeVerifier.TransactOpts, to)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "updateStorageLocation", newLocation)
}

func (_CommitteeVerifier *CommitteeVerifierSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.UpdateStorageLocation(&_CommitteeVerifier.TransactOpts, newLocation)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) UpdateStorageLocation(newLocation string) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.UpdateStorageLocation(&_CommitteeVerifier.TransactOpts, newLocation)
}

func (_CommitteeVerifier *CommitteeVerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitteeVerifier *CommitteeVerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.WithdrawFeeTokens(&_CommitteeVerifier.TransactOpts, feeTokens)
}

func (_CommitteeVerifier *CommitteeVerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitteeVerifier.Contract.WithdrawFeeTokens(&_CommitteeVerifier.TransactOpts, feeTokens)
}

type CommitteeVerifierAllowListSendersAddedIterator struct {
	Event *CommitteeVerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierAllowListSendersAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierAllowListSendersAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierAllowListSendersAddedIterator{contract: _CommitteeVerifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierAllowListSendersAdded)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitteeVerifierAllowListSendersAdded, error) {
	event := new(CommitteeVerifierAllowListSendersAdded)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierAllowListSendersRemovedIterator struct {
	Event *CommitteeVerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierAllowListSendersRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierAllowListSendersRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierAllowListSendersRemovedIterator{contract: _CommitteeVerifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierAllowListSendersRemoved)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitteeVerifierAllowListSendersRemoved, error) {
	event := new(CommitteeVerifierAllowListSendersRemoved)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierConfigSetIterator struct {
	Event *CommitteeVerifierConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierConfigSet struct {
	DynamicConfig CommitteeVerifierDynamicConfig
	Raw           types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierConfigSetIterator, error) {

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierConfigSetIterator{contract: _CommitteeVerifier.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseConfigSet(log types.Log) (*CommitteeVerifierConfigSet, error) {
	event := new(CommitteeVerifierConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierDestChainConfigSetIterator struct {
	Event *CommitteeVerifierDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierDestChainConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierDestChainConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierDestChainConfigSetIterator{contract: _CommitteeVerifier.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierDestChainConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseDestChainConfigSet(log types.Log) (*CommitteeVerifierDestChainConfigSet, error) {
	event := new(CommitteeVerifierDestChainConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierFeeTokenWithdrawnIterator struct {
	Event *CommitteeVerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierFeeTokenWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierFeeTokenWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeVerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierFeeTokenWithdrawnIterator{contract: _CommitteeVerifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierFeeTokenWithdrawn)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitteeVerifierFeeTokenWithdrawn, error) {
	event := new(CommitteeVerifierFeeTokenWithdrawn)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierOwnershipTransferRequestedIterator struct {
	Event *CommitteeVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierOwnershipTransferRequestedIterator{contract: _CommitteeVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierOwnershipTransferRequested)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitteeVerifierOwnershipTransferRequested, error) {
	event := new(CommitteeVerifierOwnershipTransferRequested)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierOwnershipTransferredIterator struct {
	Event *CommitteeVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierOwnershipTransferredIterator{contract: _CommitteeVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierOwnershipTransferred)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*CommitteeVerifierOwnershipTransferred, error) {
	event := new(CommitteeVerifierOwnershipTransferred)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierSignatureConfigSetIterator struct {
	Event *CommitteeVerifierSignatureConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierSignatureConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierSignatureConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierSignatureConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierSignatureConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierSignatureConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierSignatureConfigSet struct {
	SourceChainSelector uint64
	Signers             []common.Address
	Threshold           uint8
	Raw                 types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterSignatureConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CommitteeVerifierSignatureConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "SignatureConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierSignatureConfigSetIterator{contract: _CommitteeVerifier.contract, event: "SignatureConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierSignatureConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "SignatureConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierSignatureConfigSet)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseSignatureConfigSet(log types.Log) (*CommitteeVerifierSignatureConfigSet, error) {
	event := new(CommitteeVerifierSignatureConfigSet)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "SignatureConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitteeVerifierStorageLocationUpdatedIterator struct {
	Event *CommitteeVerifierStorageLocationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitteeVerifierStorageLocationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitteeVerifierStorageLocationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(CommitteeVerifierStorageLocationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *CommitteeVerifierStorageLocationUpdatedIterator) Error() error {
	return it.fail
}

func (it *CommitteeVerifierStorageLocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitteeVerifierStorageLocationUpdated struct {
	OldLocation string
	NewLocation string
	Raw         types.Log
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CommitteeVerifierStorageLocationUpdatedIterator, error) {

	logs, sub, err := _CommitteeVerifier.contract.FilterLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return &CommitteeVerifierStorageLocationUpdatedIterator{contract: _CommitteeVerifier.contract, event: "StorageLocationUpdated", logs: logs, sub: sub}, nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationUpdated) (event.Subscription, error) {

	logs, sub, err := _CommitteeVerifier.contract.WatchLogs(opts, "StorageLocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitteeVerifierStorageLocationUpdated)
				if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_CommitteeVerifier *CommitteeVerifierFilterer) ParseStorageLocationUpdated(log types.Log) (*CommitteeVerifierStorageLocationUpdated, error) {
	event := new(CommitteeVerifierStorageLocationUpdated)
	if err := _CommitteeVerifier.contract.UnpackLog(event, "StorageLocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetAllSignatureConfigs struct {
	SourceChainSelectors []uint64
	SignerSets           [][]common.Address
	Thresholds           []uint8
}
type GetDestChainConfig struct {
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}
type GetFee struct {
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}

func (CommitteeVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitteeVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitteeVerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0x781b4fc361184bd997c249fbc855854e7d6daf1c72a585b5598c778e23dc35cd")
}

func (CommitteeVerifierDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitteeVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitteeVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitteeVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CommitteeVerifierSignatureConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3780850db9abcff2be2b607bfbc2b86c9c131d50e456bf09dbaf923039ad4b83")
}

func (CommitteeVerifierStorageLocationUpdated) Topic() common.Hash {
	return common.HexToHash("0xbea2c78e36ed08c4b0076b01d186a4c2194d4109169fad20958c761b40908bd8")
}

func (_CommitteeVerifier *CommitteeVerifier) Address() common.Address {
	return _CommitteeVerifier.address
}

type CommitteeVerifierInterface interface {
	ForwardToVerifier(opts *bind.CallOpts, message MessageV1CodecMessageV1, arg1 [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) ([]byte, error)

	GetAllSignatureConfigs(opts *bind.CallOpts) (GetAllSignatureConfigs,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitteeVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetSignatureConfig(opts *bind.CallOpts, sourceChainSelector uint64) ([]common.Address, uint8, error)

	GetStorageLocation(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyMessage(opts *bind.CallOpts, message MessageV1CodecMessageV1, messageHash [32]byte, ccvData []byte) error

	VersionTag(opts *bind.CallOpts) ([4]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []BaseVerifierAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []BaseVerifierDestChainConfigArgs) (*types.Transaction, error)

	ApplySignersUpdates(opts *bind.TransactOpts, sourceChainSelectorsToRemove []uint64, signersUpdates []SignatureQuorumValidatorSignersUpdate) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitteeVerifierDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStorageLocation(opts *bind.TransactOpts, newLocation string) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitteeVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitteeVerifierAllowListSendersRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitteeVerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitteeVerifierConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitteeVerifierDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitteeVerifierDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CommitteeVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitteeVerifierFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitteeVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitteeVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitteeVerifierOwnershipTransferred, error)

	FilterSignatureConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CommitteeVerifierSignatureConfigSetIterator, error)

	WatchSignatureConfigSet(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierSignatureConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSignatureConfigSet(log types.Log) (*CommitteeVerifierSignatureConfigSet, error)

	FilterStorageLocationUpdated(opts *bind.FilterOpts) (*CommitteeVerifierStorageLocationUpdatedIterator, error)

	WatchStorageLocationUpdated(opts *bind.WatchOpts, sink chan<- *CommitteeVerifierStorageLocationUpdated) (event.Subscription, error)

	ParseStorageLocationUpdated(log types.Log) (*CommitteeVerifierStorageLocationUpdated, error)

	Address() common.Address
}
