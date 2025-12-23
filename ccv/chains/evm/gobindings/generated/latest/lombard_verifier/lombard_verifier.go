// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lombard_verifier

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

type BaseVerifierRemoteChainConfigArgs struct {
	Router              common.Address
	RemoteChainSelector uint64
	AllowlistEnabled    bool
	FeeUSDCents         uint16
	GasForVerification  uint32
	PayloadSizeBytes    uint32
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

type LombardVerifierDynamicConfig struct {
	FeeAggregator common.Address
}

type LombardVerifierPath struct {
	AllowedCaller [32]byte
	LChainId      [32]byte
}

type LombardVerifierSupportedTokenArgs struct {
	LocalToken   common.Address
	LocalAdapter common.Address
}

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	MessageNumber       uint64
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

var LombardVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"},{\"name\":\"storageLocation\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePaths\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSupportedTokens\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokensToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct LombardVerifier.SupportedTokenArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct LombardVerifier.DynamicConfig\",\"components\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenRemoved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"messageMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustTransferTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]}]",
	Bin: "0x60c0806040523461067c576140f1803803809161001c8285610681565b83398101908082036080811261067c5760201361067c57604051602081016001600160401b0381118282101761048957604052610058826106a4565b81526020820151926001600160a01b0384169283850361067c5760408101516001600160401b03811161067c57810182601f8201121561067c5780519061009e826106b8565b936100ac6040519586610681565b82855260208086019360051b8301019181831161067c5760208101935b83851061060457505050505060606100e191016106a4565b906001549080516100f1836106b8565b926100ff6040519485610681565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b83821061055f5750505060005b8181106104cb57505060005b81811061033c5750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d075869161019d61018f9260405193849360408552604085019061075e565b90838203602085015261075e565b0390a16001600160a01b0316801561031a57608052331561032b57600380546001600160a01b0319163317905580516001600160a01b03161561031a5751600a80546001600160a01b0319166001600160a01b039290921691821790556040519081527ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f6090602090a180156103095760206004916040519283809263353c26b760e01b82525afa9081156102fd576000916102b9575b5060ff16600181036102a0575060a05260405161392190816107d0823960805181612bcc015260a051818181610bb30152818161184101528181611fa2015281816122530152612ea80152f35b63398bbe0560e11b600052600160045260245260446000fd5b6020813d6020116102f5575b816102d260209383610681565b810103126102f157519060ff821682036102ee575060ff610253565b80fd5b5080fd5b3d91506102c5565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b82518110156104b55760208160051b84010151600154680100000000000000008110156104895780600161037392016001556106f2565b91909161049f578051906001600160401b03821161048957610395835461070d565b601f811161044c575b50602090601f83116001146103e157600194939291600091836103d6575b5050600019600383901b1c191690841b1790555b01610149565b0151905038806103bc565b90601f1983169184600052816000209260005b81811061043457509160019695949291838895931061041b575b505050811b0190556103d0565b015160001960f88460031b161c1916905538808061040e565b929360206001819287860151815501950193016103f4565b61047990846000526020600020601f850160051c8101916020861061047f575b601f0160051c0190610747565b3861039e565b909150819061046c565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60015480156105495760001901906104e2826106f2565b92909261049f57826104f66001945461070d565b9081610507575b505082550161013d565b81601f60009311861461051e5750555b38806104fd565b8183526020832061053991601f0160051c8101908701610747565b8082528160208120915555610517565b634e487b7160e01b600052603160045260246000fd5b6040516000845461056f8161070d565b80845290600181169081156105e157506001146105a9575b506001928261059b85946020940382610681565b815201930191019091610130565b6000868152602081209092505b8183106105cb57505081016020016001610587565b60018160209254838688010152019201916105b6565b60ff191660208581019190915291151560051b8401909101915060019050610587565b84516001600160401b03811161067c5782019083603f8301121561067c576020820151906001600160401b0382116104895760405161064d601f8401601f191660200182610681565b828152604084840101861061067c57610671602094938594604086850191016106cf565b8152019401936100c9565b600080fd5b601f909101601f19168101906001600160401b0382119082101761048957604052565b51906001600160a01b038216820361067c57565b6001600160401b0381116104895760051b60200190565b60005b8381106106e25750506000910152565b81810151838201526020016106d2565b6001548110156104b557600160005260206000200190600090565b90600182811c9216801561073d575b602083101461072757565b634e487b7160e01b600052602260045260246000fd5b91607f169161071c565b818110610752575050565b60008155600101610747565b9080602083519182815201916020808360051b8301019401926000915b83831061078a57505050505090565b909192939460208080600193601f1986820301875289516107b6815180928185528580860191016106cf565b601f01601f19160101970195949190910192019061077b56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610177578063181f5a7714610172578063240028e81461016d5780632e45ca681461016857806338ff8c38146101635780633bbbed4b1461015e5780635cb80c5d14610159578063708e1f7914610154578063737037e81461014f5780637437ff9f1461014a57806379ba50971461014557806380485e251461014057806387ae92921461013b578063898068fc146101365780638da5cb5b146101315780638f2aaea41461012c5780639ba4393114610127578063bcb6d4f714610122578063bff0ec1d1461011d578063c4bffe2b14610118578063d3c7c2c7146101135763f2fde38b1461010e57600080fd5b611c3f565b611bb3565b611af0565b61167f565b6115cf565b6114d7565b6113cc565b61137a565b6112da565b6110fc565b610eeb565b610cef565b610c75565b610bd7565b610b68565b6109ec565b610653565b610575565b6104e6565b610473565b6103c9565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361023657807f83adcde1000000000000000000000000000000000000000000000000000000006020921490811561020c575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610201565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761028657604052565b61023b565b6020810190811067ffffffffffffffff82111761028657604052565b60c0810190811067ffffffffffffffff82111761028657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761028657604052565b6040519061031360a0836102c3565b565b67ffffffffffffffff811161028657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106103625750506000910152565b8181015183820152602001610352565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936103ae8151809281875287808801910161034f565b0116010190565b9060206103c6928181520190610372565b90565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657610446604080519061040a81836102c3565b601982527f4c6f6d62617264566572696669657220312e372e302d64657600000000000000602083015251918291602083526020830190610372565b0390f35b73ffffffffffffffffffffffffffffffffffffffff81160361023657565b35906103138261044a565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760206104dc73ffffffffffffffffffffffffffffffffffffffff6004356104c88161044a565b166000526005602052604060002054151590565b6040519015158152f35b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff8111610236573660238201121561023657806004013567ffffffffffffffff81116102365736602460c08302840101116102365760246105619201611d33565b005b67ffffffffffffffff81160361023657565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365767ffffffffffffffff6004356105b981610563565b600060206040516105c98161026a565b828152015216600052600960205261044660406000206001604051916105ee8361026a565b8054835201546020820152604051918291829190916020806040830194805184520151910152565b90816101c09103126102365790565b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b346102365760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff8111610236576106a2903690600401610616565b602435906106b160443561044a565b60843567ffffffffffffffff8111610236576106d1903690600401610625565b505060208101803560006106e482610563565b6106ed82612b7a565b6101808401906106fd8286611e4a565b9050156109935761070d83610563565b61012085019273ffffffffffffffffffffffffffffffffffffffff61073d6107358689611e9e565b810190611eef565b169061075d8167ffffffffffffffff166000526000602052604060002090565b8054909161079573ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff81161561095d576040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152602090829060249082905afa80156109585773ffffffffffffffffffffffffffffffffffffffff918691610929575b501633036108fd5760f01c60ff1661086e575b61044661086289898961085a8a61085461084e6108488d87611e4a565b90611f33565b93611e40565b93611e9e565b929091612d3c565b604051918291826103b5565b6108ae6108b291600161089761077c8673ffffffffffffffffffffffffffffffffffffffff1690565b910160019160005201602052604060002054151590565b1590565b6108bc578061082b565b7fd0d2597600000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff16600452602490fd5b7f728fe07b00000000000000000000000000000000000000000000000000000000845233600452602484fd5b61094b915060203d602011610951575b61094381836102c3565b8101906126e4565b38610818565b503d610939565b612337565b7f4d1aff7e00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff8216600452602486fd5b807f4f73dc4d0000000000000000000000000000000000000000000000000000000060049252fd5b9181601f840112156102365782359167ffffffffffffffff8311610236576020808501948460051b01011161023657565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff811161023657610a3b9036906004016109bb565b9073ffffffffffffffffffffffffffffffffffffffff600a54169160005b818110610a6257005b610a7861077c610a738385876122bf565b6122cf565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa8015610958576001948892600092610b38575b5081610aec575b5050505001610a59565b81610b1c7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610b2c9461326d565b6040519081529081906020820190565b0390a338858180610ae2565b610b5a91925060203d8111610b61575b610b5281836102c3565b810190612fdd565b9038610adb565b503d610b48565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102365760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff811161023657610c269036906004016109bb565b6024359167ffffffffffffffff831161023657366023840112156102365782600401359167ffffffffffffffff8311610236573660248460061b86010111610236576024610561940191611f71565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576000604051610cb28161028b565b52610446604051610cc28161028b565b600a5473ffffffffffffffffffffffffffffffffffffffff16908190526040519081529081906020820190565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760025473ffffffffffffffffffffffffffffffffffffffff81163303610dae577fffffffffffffffffffffffff00000000000000000000000000000000000000006003549133828416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610de482610315565b91610df260405193846102c3565b829481845281830111610236578281602093846000960137010152565b9080601f83011215610236578160206103c693359101610dd8565b67ffffffffffffffff81116102865760051b60200190565b81601f8201121561023657803590610e5982610e2a565b92610e6760405194856102c3565b82845260208085019360061b8301019181831161023657602001925b828410610e91575050505090565b6040848303126102365760206040918251610eab8161026a565b8635610eb68161044a565b81528287013583820152815201930192610e83565b6064359061ffff8216820361023657565b359061ffff8216820361023657565b346102365760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657600435610f2681610563565b60243567ffffffffffffffff81116102365760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261023657610f6c610304565b90806004013567ffffffffffffffff811161023657610f919060043691840101610e0f565b8252602481013567ffffffffffffffff811161023657610fb79060043691840101610e0f565b6020830152604481013567ffffffffffffffff811161023657610fe09060043691840101610e42565b6040830152610ff160648201610468565b6060830152608481013567ffffffffffffffff811161023657608091600461101c9236920101610e0f565b91015260443567ffffffffffffffff81116102365761044691611046611055923690600401610e0f565b5061104f610ecb565b50612343565b6040805161ffff909416845263ffffffff92831660208501529116908201529081906060820190565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106110b157505050505090565b90919293946020806110ed837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610372565b970193019301919392906110a2565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760015461113781610e2a565b61114460405191826102c3565b818152602081019160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66000935b82851061118b5760405180610446868261107e565b604051600083548060011c9060018116908115611286575b60208310821461125957828552602085019190811561122457506001146111e8575b5050600192826111da859460209403826102c3565b815201920194019390611176565b9091506111fa85600052602060002090565b916000925b8184106112105750500182826111c5565b6001816020925486860152019301926111ff565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682525090151560051b01905082826111c5565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f16916111a3565b906020808351928381520192019060005b8181106112ae5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016112a1565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365767ffffffffffffffff60043561131e81610563565b16600052600060205273ffffffffffffffffffffffffffffffffffffffff6040600020610446611352600183549301613431565b604051938360ff869560f01c1615158552166020840152606060408401526060830190611290565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760006040516114098161028b565b6004356114158161044a565b815261141f6127d5565b73ffffffffffffffffffffffffffffffffffffffff815116156114af5751600a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527ff6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f6090602090a180f35b6004827f8579befe000000000000000000000000000000000000000000000000000000008152fd5b346102365760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043561151281610563565b602435906044356115216127d5565b82156115a5576115a07f83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9916040516115588161026a565b818152600167ffffffffffffffff602083019688885216958660005260096020526040600020925183555191015561158f846134af565b506040519081529081906020820190565b0390a3005b7f5a39e3030000000000000000000000000000000000000000000000000000000060005260046000fd5b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff811161023657366023820112156102365780600401359061162a82610e2a565b9161163860405193846102c3565b8083526024602084019160051b8301019136831161023657602401905b828210611665576105618461240e565b60208091833561167481610563565b815201910190611655565b346102365760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365760043567ffffffffffffffff8111610236576116ce903690600401610616565b602435906044359067ffffffffffffffff8211610236576117176116f9611712933690600401610625565b93909261170d61170882611e40565b612b7a565b611e40565b613088565b61172a6117248383612568565b906125c9565b7ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160361197c57506117fc6117c5916117a361179c61179661179060068886612576565b9061267e565b60f01c90565b61ffff1690565b936117bb6117b2866006612671565b600683856125b1565b9490956006612671565b906117f66117ef6117e961179c6117966117906117e18861265e565b88888b6125b1565b9361265e565b9283612671565b926125b1565b92604051927fd5438eae00000000000000000000000000000000000000000000000000000000845260208460048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215610958576000948593611a73575b5073ffffffffffffffffffffffffffffffffffffffff85916118c5604051988997889687947fa6208506000000000000000000000000000000000000000000000000000000008652600486016127ae565b0393165af190811561095857600090600092611a4c575b5015611a225760248151036119ef5760246020820151910151907ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160361197c575081810361194957005b7f6c86fa3a0000000000000000000000000000000000000000000000000000000060005260049190915260245260446000fd5b7fadaf7739000000000000000000000000000000000000000000000000000000006000527ff0f3a135000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b517fc2fdac9800000000000000000000000000000000000000000000000000000000600052602460048190525260446000fd5b7f2532cf450000000000000000000000000000000000000000000000000000000060005260046000fd5b9050611a6b91503d806000833e611a6381836102c3565b8101906126f9565b9150386118dc565b85919350611aa473ffffffffffffffffffffffffffffffffffffffff9160203d6020116109515761094381836102c3565b939150611874565b602060408183019282815284518094520192019060005b818110611ad05750505090565b825167ffffffffffffffff16845260209384019390920191600101611ac3565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023657600754611b2b81610e2a565b90611b3960405192836102c3565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611b6682610e2a565b0136602084013760005b818110611b8557604051806104468582611aac565b8067ffffffffffffffff611b9a60019361347c565b90549060031b1c16611bac8286612518565b5201611b70565b346102365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610236576040516004548082526020820190600460005260206000209060005b818110611c295761044685611c15818703826102c3565b604051918291602083526020830190611290565b8254845260209093019260019283019201611bfe565b346102365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102365773ffffffffffffffffffffffffffffffffffffffff600435611c8f8161044a565b611c976127d5565b16338114611d0957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90611d3c6127d5565b611d4581610e2a565b91611d5360405193846102c3565b81835260c060208401920281019036821161023657915b818310611d7d5750505061031390612820565b60c08336031261023657602060c091604051611d98816102a7565b8535611da38161044a565b815282860135611db281610563565b838201526040860135611dc481611e07565b6040820152611dd560608701610edc565b6060820152611de660808701611e11565b6080820152611df760a08701611e11565b60a0820152815201920191611d6a565b8015150361023657565b359063ffffffff8216820361023657565b90604051611e2f8161026a565b602060018294805484520154910152565b356103c681610563565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610236570180359067ffffffffffffffff821161023657602001918160051b3603831361023657565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610236570180359067ffffffffffffffff82116102365760200191813603831361023657565b9081602091031261023657356103c68161044a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015611f6c578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610236570190565b611f04565b90611f7a6127d5565b60005b8181106121d0575050509060009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915b818110611fd15750505050565b611fe4611fdf8284876122d9565b6122e9565b90612003825173ffffffffffffffffffffffffffffffffffffffff1690565b91612030602082019361202a855173ffffffffffffffffffffffffffffffffffffffff1690565b90613054565b50825173ffffffffffffffffffffffffffffffffffffffff1680156121b157925b602060405180957f095ea7b300000000000000000000000000000000000000000000000000000000825281600073ffffffffffffffffffffffffffffffffffffffff826120e48d6004830160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9193929373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0393165af19384156109585761213c7f086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e39361215792600197612185575b505173ffffffffffffffffffffffffffffffffffffffff1690565b915173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152a101611fc4565b6121a59060203d81116121aa575b61219d81836102c3565b810190612322565b612121565b503d612193565b50805173ffffffffffffffffffffffffffffffffffffffff1692612051565b806121e1610a7360019385876122bf565b61220861220373ffffffffffffffffffffffffffffffffffffffff831661077c565b6132cf565b612214575b5001611f7d565b6122b67fbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a3069161228f73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316612fec565b60405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a13861220d565b9190811015611f6c5760051b0190565b356103c68161044a565b9190811015611f6c5760061b0190565b604081360312610236576020604051916123028361026a565b803561230d8161044a565b8352013561231a8161044a565b602082015290565b9081602091031261023657516103c681611e07565b6040513d6000823e3d90fd5b9067ffffffffffffffff821680600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000205416156123e157600052600060205261ffff60406000205460a01c169063ffffffff6123d6816123b68667ffffffffffffffff166000526000602052604060002090565b5460b01c169467ffffffffffffffff166000526000602052604060002090565b5460d01c1691929190565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b906124176127d5565b6000915b80518310156125135767ffffffffffffffff6124378483612518565b51169261246061245b8567ffffffffffffffff166000526009602052604060002090565b611e22565b6124698561362f565b156124da57846124a0612494600195969767ffffffffffffffff166000526009602052604060002090565b60016000918281550155565b60208281015192516040519081527f8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b9190a301919061241b565b847fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff602491166004526000fd5b509050565b8051821015611f6c5760209160051b010190565b91612564918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b906004116102365790600490565b909291928360041161023657831161023657600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610236578411610236578101920390565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106125fd575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906002820180921161266c57565b61262f565b9190820180921161266c57565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106126b2575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b9081602091031261023657516103c68161044a565b909160608284031261023657815192602083015161271681611e07565b9260408101519067ffffffffffffffff8211610236570181601f8201121561023657805161274381610315565b9261275160405194856102c3565b81845260208284010111610236576103c6916020808501910161034f565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b92906127c7906103c6959360408652604086019161276f565b92602081850391015261276f565b73ffffffffffffffffffffffffffffffffffffffff6003541633036127f657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060005b8251811015612513576128378184612518565b5161285760206128478487612518565b51015167ffffffffffffffff1690565b67ffffffffffffffff8116918215612b42576128878267ffffffffffffffff166000526000602052604060002090565b916128ea6128a9835173ffffffffffffffffffffffffffffffffffffffff1690565b849073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61294a6128fa6040840151151590565b84547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f01b7eff00000000000000000000000000000000000000000000000000000000000016178455565b6129a361295c606084015161ffff1690565b84547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff000000000000000000000000000000000000000016178455565b60808201906129c26129b9835163ffffffff1690565b63ffffffff1690565b15612b0b575091612adc612ad161077c7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc994612a54612a0960019a99985163ffffffff1690565b86547fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1660b09190911b79ffffffff0000000000000000000000000000000000000000000016178655565b612ab7612a6860a083015163ffffffff1690565b86547fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff1660d09190911b7dffffffff000000000000000000000000000000000000000000000000000016178655565b5173ffffffffffffffffffffffffffffffffffffffff1690565b915460f01c60ff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff939093168352901515602083015290a201612824565b7f9e7205510000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f97ccaab70000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561095857600091612c51575b50612c1a5750565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612c6a915060203d6020116121aa5761219d81836102c3565b38612c12565b9160206103c693818152019161276f565b90604051917ff0f3a1350000000000000000000000000000000000000000000000000000000060208401526024830152602482526103136044836102c3565b9190826040910312610236576020825192015190565b93906103c697969373ffffffffffffffffffffffffffffffffffffffff60e09794819388521660208701521660408501526060840152608083015260a08201528160c08201520190610372565b90604051916020830152602082526103136040836102c3565b929493909460808401936020612d528683611e9e565b905011612f9957612d6c61077c6107356040840184611e9e565b92612d9d6108ae73ffffffffffffffffffffffffffffffffffffffff86166000526005602052604060002054151590565b612f5557612dc261245b8967ffffffffffffffff166000526009602052604060002090565b805115612f1d57604096979850612dfd61077c61077c612df861077c8973ffffffffffffffffffffffffffffffffffffffff1690565b6137ed565b73ffffffffffffffffffffffffffffffffffffffff8116612f13575b50612e5a612e50612e4b612e44612e3d612e8d98999a60208701519a810190611eef565b9b87611e9e565b3691610dd8565b6131c8565b9335915192612c81565b92875198899788977e8a119800000000000000000000000000000000000000000000000000000000895260048901612cd6565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af18015610958576103c691600091612ee2575b50612d23565b612f04915060403d604011612f0c575b612efc81836102c3565b810190612cc0565b905038612edc565b503d612ef2565b9450612e5a612e19565b7fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff891660045260246000fd5b7f06439c6b0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b84612fa391611e9e565b90612fd96040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401612c70565b0390fd5b90816020910312610236575190565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff909216602483015260006044808401919091528252610313919061304f6064836102c3565b613375565b9073ffffffffffffffffffffffffffffffffffffffff806103c6931691826000526006602052166040600020556004613540565b6130c661077c6130ac8367ffffffffffffffff166000526000602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811615613190576040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082905afa90811561095857600091613171575b501561314357565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b61318a915060203d6020116121aa5761219d81836102c3565b3861313b565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b602081511161323757602081519101519060208110613206575b8060031b908082046008149015171561266c5761010003610100811161266c571c90565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b16906131e2565b612fd9906040519182917fe0d7fb02000000000000000000000000000000000000000000000000000000008352600483016103b5565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9290921660248301526044808301939093529181526103139161304f6064836102c3565b6103c69080600052600660205260006040812055600461370c565b156132f157565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9073ffffffffffffffffffffffffffffffffffffffff61340392166040906000808351946133a385876102c3565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602087015260208151910182855af1903d15613428573d6133f46133eb82610315565b945194856102c3565b83523d6000602085013e613850565b80518061340e575050565b81602080613423936103139501019101612322565b6132ea565b60609250613850565b906040519182815491828252602082019060005260206000209260005b818110613463575050610313925003836102c3565b845483526001948501948794506020909301920161344e565b600754811015611f6c57600760005260206000200190600090565b8054821015611f6c5760005260206000200190600090565b60008181526008602052604090205461353a5760075468010000000000000000811015610286576135216134ec8260018594016007556007613497565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600754906000526008602052604060002055600190565b50600090565b60008281526001820160205260409020546135955780549068010000000000000000821015610286578261357e6134ec846001809601855584613497565b905580549260005201602052604060002055600190565b5050600090565b80548015613600577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906135d18282613497565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260086020526040902054908115613595577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161266c57600754927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840193841161266c5783836000956136cb95036136d1575b5050506136ba600761359c565b600890600052602052604060002090565b55600190565b6136ba6136fd916136f36136e9613703956007613497565b90549060031b1c90565b9283916007613497565b9061252c565b553880806136ad565b60018101918060005282602052604060002054928315156000146137e4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff840184811161266c578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff850194851161266c5760009585836136cb9761379c95036137ab575b50505061359c565b90600052602052604060002090565b6137cb6136fd916137c26136e96137db9588613497565b92839187613497565b8590600052602052604060002090565b55388080613794565b50505050600090565b8060005260066020526040600020549081158061383a575b61380d575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5060008181526005602052604090205415613805565b919290156138cb5750815115613864575090565b3b1561386d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156138de5750805190602001fd5b612fd9906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016103b556fea164736f6c634300081a000a",
}

var LombardVerifierABI = LombardVerifierMetaData.ABI

var LombardVerifierBin = LombardVerifierMetaData.Bin

func DeployLombardVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig LombardVerifierDynamicConfig, bridge common.Address, storageLocation []string, rmn common.Address) (common.Address, *types.Transaction, *LombardVerifier, error) {
	parsed, err := LombardVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardVerifierBin), backend, dynamicConfig, bridge, storageLocation, rmn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LombardVerifier{address: address, abi: *parsed, LombardVerifierCaller: LombardVerifierCaller{contract: contract}, LombardVerifierTransactor: LombardVerifierTransactor{contract: contract}, LombardVerifierFilterer: LombardVerifierFilterer{contract: contract}}, nil
}

type LombardVerifier struct {
	address common.Address
	abi     abi.ABI
	LombardVerifierCaller
	LombardVerifierTransactor
	LombardVerifierFilterer
}

type LombardVerifierCaller struct {
	contract *bind.BoundContract
}

type LombardVerifierTransactor struct {
	contract *bind.BoundContract
}

type LombardVerifierFilterer struct {
	contract *bind.BoundContract
}

type LombardVerifierSession struct {
	Contract     *LombardVerifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LombardVerifierCallerSession struct {
	Contract *LombardVerifierCaller
	CallOpts bind.CallOpts
}

type LombardVerifierTransactorSession struct {
	Contract     *LombardVerifierTransactor
	TransactOpts bind.TransactOpts
}

type LombardVerifierRaw struct {
	Contract *LombardVerifier
}

type LombardVerifierCallerRaw struct {
	Contract *LombardVerifierCaller
}

type LombardVerifierTransactorRaw struct {
	Contract *LombardVerifierTransactor
}

func NewLombardVerifier(address common.Address, backend bind.ContractBackend) (*LombardVerifier, error) {
	abi, err := abi.JSON(strings.NewReader(LombardVerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLombardVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LombardVerifier{address: address, abi: abi, LombardVerifierCaller: LombardVerifierCaller{contract: contract}, LombardVerifierTransactor: LombardVerifierTransactor{contract: contract}, LombardVerifierFilterer: LombardVerifierFilterer{contract: contract}}, nil
}

func NewLombardVerifierCaller(address common.Address, caller bind.ContractCaller) (*LombardVerifierCaller, error) {
	contract, err := bindLombardVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierCaller{contract: contract}, nil
}

func NewLombardVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*LombardVerifierTransactor, error) {
	contract, err := bindLombardVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierTransactor{contract: contract}, nil
}

func NewLombardVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*LombardVerifierFilterer, error) {
	contract, err := bindLombardVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierFilterer{contract: contract}, nil
}

func bindLombardVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LombardVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LombardVerifier *LombardVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardVerifier.Contract.LombardVerifierCaller.contract.Call(opts, result, method, params...)
}

func (_LombardVerifier *LombardVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.Contract.LombardVerifierTransactor.contract.Transfer(opts)
}

func (_LombardVerifier *LombardVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardVerifier.Contract.LombardVerifierTransactor.contract.Transact(opts, method, params...)
}

func (_LombardVerifier *LombardVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardVerifier.Contract.contract.Call(opts, result, method, params...)
}

func (_LombardVerifier *LombardVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.Contract.contract.Transfer(opts)
}

func (_LombardVerifier *LombardVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardVerifier.Contract.contract.Transact(opts, method, params...)
}

func (_LombardVerifier *LombardVerifierCaller) GetDynamicConfig(opts *bind.CallOpts) (LombardVerifierDynamicConfig, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(LombardVerifierDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(LombardVerifierDynamicConfig)).(*LombardVerifierDynamicConfig)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetDynamicConfig() (LombardVerifierDynamicConfig, error) {
	return _LombardVerifier.Contract.GetDynamicConfig(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetDynamicConfig() (LombardVerifierDynamicConfig, error) {
	return _LombardVerifier.Contract.GetDynamicConfig(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.GasForVerification = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PayloadSizeBytes = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_LombardVerifier *LombardVerifierSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _LombardVerifier.Contract.GetFee(&_LombardVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

	error) {
	return _LombardVerifier.Contract.GetFee(&_LombardVerifier.CallOpts, destChainSelector, arg1, arg2, arg3)
}

func (_LombardVerifier *LombardVerifierCaller) GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardVerifierPath, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getPath", remoteChainSelector)

	if err != nil {
		return *new(LombardVerifierPath), err
	}

	out0 := *abi.ConvertType(out[0], new(LombardVerifierPath)).(*LombardVerifierPath)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetPath(remoteChainSelector uint64) (LombardVerifierPath, error) {
	return _LombardVerifier.Contract.GetPath(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetPath(remoteChainSelector uint64) (LombardVerifierPath, error) {
	return _LombardVerifier.Contract.GetPath(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCaller) GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getRemoteChainConfig", remoteChainSelector)

	outstruct := new(GetRemoteChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllowedSendersList = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_LombardVerifier *LombardVerifierSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _LombardVerifier.Contract.GetRemoteChainConfig(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetRemoteChainConfig(remoteChainSelector uint64) (GetRemoteChainConfig,

	error) {
	return _LombardVerifier.Contract.GetRemoteChainConfig(&_LombardVerifier.CallOpts, remoteChainSelector)
}

func (_LombardVerifier *LombardVerifierCaller) GetStorageLocations(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getStorageLocations")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetStorageLocations() ([]string, error) {
	return _LombardVerifier.Contract.GetStorageLocations(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetStorageLocations() ([]string, error) {
	return _LombardVerifier.Contract.GetStorageLocations(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetSupportedChains() ([]uint64, error) {
	return _LombardVerifier.Contract.GetSupportedChains(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LombardVerifier.Contract.GetSupportedChains(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) GetSupportedTokens() ([]common.Address, error) {
	return _LombardVerifier.Contract.GetSupportedTokens(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _LombardVerifier.Contract.GetSupportedTokens(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) IBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "i_bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) IBridge() (common.Address, error) {
	return _LombardVerifier.Contract.IBridge(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) IBridge() (common.Address, error) {
	return _LombardVerifier.Contract.IBridge(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardVerifier.Contract.IsSupportedToken(&_LombardVerifier.CallOpts, token)
}

func (_LombardVerifier *LombardVerifierCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardVerifier.Contract.IsSupportedToken(&_LombardVerifier.CallOpts, token)
}

func (_LombardVerifier *LombardVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) Owner() (common.Address, error) {
	return _LombardVerifier.Contract.Owner(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) Owner() (common.Address, error) {
	return _LombardVerifier.Contract.Owner(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardVerifier.Contract.SupportsInterface(&_LombardVerifier.CallOpts, interfaceId)
}

func (_LombardVerifier *LombardVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardVerifier.Contract.SupportsInterface(&_LombardVerifier.CallOpts, interfaceId)
}

func (_LombardVerifier *LombardVerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LombardVerifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LombardVerifier *LombardVerifierSession) TypeAndVersion() (string, error) {
	return _LombardVerifier.Contract.TypeAndVersion(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierCallerSession) TypeAndVersion() (string, error) {
	return _LombardVerifier.Contract.TypeAndVersion(&_LombardVerifier.CallOpts)
}

func (_LombardVerifier *LombardVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "acceptOwnership")
}

func (_LombardVerifier *LombardVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardVerifier.Contract.AcceptOwnership(&_LombardVerifier.TransactOpts)
}

func (_LombardVerifier *LombardVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardVerifier.Contract.AcceptOwnership(&_LombardVerifier.TransactOpts)
}

func (_LombardVerifier *LombardVerifierTransactor) ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "applyRemoteChainConfigUpdates", remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyRemoteChainConfigUpdates(&_LombardVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierTransactorSession) ApplyRemoteChainConfigUpdates(remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ApplyRemoteChainConfigUpdates(&_LombardVerifier.TransactOpts, remoteChainConfigArgs)
}

func (_LombardVerifier *LombardVerifierTransactor) ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "forwardToVerifier", message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ForwardToVerifier(&_LombardVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierTransactorSession) ForwardToVerifier(message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.ForwardToVerifier(&_LombardVerifier.TransactOpts, message, messageId, arg2, arg3, arg4)
}

func (_LombardVerifier *LombardVerifierTransactor) RemovePaths(opts *bind.TransactOpts, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "removePaths", remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierSession) RemovePaths(remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.Contract.RemovePaths(&_LombardVerifier.TransactOpts, remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierTransactorSession) RemovePaths(remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _LombardVerifier.Contract.RemovePaths(&_LombardVerifier.TransactOpts, remoteChainSelectors)
}

func (_LombardVerifier *LombardVerifierTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_LombardVerifier *LombardVerifierSession) SetDynamicConfig(dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetDynamicConfig(&_LombardVerifier.TransactOpts, dynamicConfig)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetDynamicConfig(dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetDynamicConfig(&_LombardVerifier.TransactOpts, dynamicConfig)
}

func (_LombardVerifier *LombardVerifierTransactor) SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "setPath", remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierTransactorSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.SetPath(&_LombardVerifier.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardVerifier *LombardVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "transferOwnership", to)
}

func (_LombardVerifier *LombardVerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.TransferOwnership(&_LombardVerifier.TransactOpts, to)
}

func (_LombardVerifier *LombardVerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.TransferOwnership(&_LombardVerifier.TransactOpts, to)
}

func (_LombardVerifier *LombardVerifierTransactor) UpdateSupportedTokens(opts *bind.TransactOpts, tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "updateSupportedTokens", tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierSession) UpdateSupportedTokens(tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.UpdateSupportedTokens(&_LombardVerifier.TransactOpts, tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierTransactorSession) UpdateSupportedTokens(tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error) {
	return _LombardVerifier.Contract.UpdateSupportedTokens(&_LombardVerifier.TransactOpts, tokensToRemove, tokensToSet)
}

func (_LombardVerifier *LombardVerifierTransactor) VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "verifyMessage", message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierSession) VerifyMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.VerifyMessage(&_LombardVerifier.TransactOpts, message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierTransactorSession) VerifyMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error) {
	return _LombardVerifier.Contract.VerifyMessage(&_LombardVerifier.TransactOpts, message, messageId, ccvData)
}

func (_LombardVerifier *LombardVerifierTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_LombardVerifier *LombardVerifierSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.WithdrawFeeTokens(&_LombardVerifier.TransactOpts, feeTokens)
}

func (_LombardVerifier *LombardVerifierTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _LombardVerifier.Contract.WithdrawFeeTokens(&_LombardVerifier.TransactOpts, feeTokens)
}

type LombardVerifierAllowListSendersAddedIterator struct {
	Event *LombardVerifierAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierAllowListSendersAdded)
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
		it.Event = new(LombardVerifierAllowListSendersAdded)
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

func (it *LombardVerifierAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierAllowListSendersAddedIterator{contract: _LombardVerifier.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierAllowListSendersAdded)
				if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseAllowListSendersAdded(log types.Log) (*LombardVerifierAllowListSendersAdded, error) {
	event := new(LombardVerifierAllowListSendersAdded)
	if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierAllowListSendersRemovedIterator struct {
	Event *LombardVerifierAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierAllowListSendersRemoved)
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
		it.Event = new(LombardVerifierAllowListSendersRemoved)
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

func (it *LombardVerifierAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierAllowListSendersRemovedIterator{contract: _LombardVerifier.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierAllowListSendersRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseAllowListSendersRemoved(log types.Log) (*LombardVerifierAllowListSendersRemoved, error) {
	event := new(LombardVerifierAllowListSendersRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierDynamicConfigSetIterator struct {
	Event *LombardVerifierDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierDynamicConfigSet)
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
		it.Event = new(LombardVerifierDynamicConfigSet)
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

func (it *LombardVerifierDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierDynamicConfigSet struct {
	DynamicConfig LombardVerifierDynamicConfig
	Raw           types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardVerifierDynamicConfigSetIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierDynamicConfigSetIterator{contract: _LombardVerifier.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierDynamicConfigSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseDynamicConfigSet(log types.Log) (*LombardVerifierDynamicConfigSet, error) {
	event := new(LombardVerifierDynamicConfigSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierFeeTokenWithdrawnIterator struct {
	Event *LombardVerifierFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierFeeTokenWithdrawn)
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
		it.Event = new(LombardVerifierFeeTokenWithdrawn)
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

func (it *LombardVerifierFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardVerifierFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierFeeTokenWithdrawnIterator{contract: _LombardVerifier.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierFeeTokenWithdrawn)
				if err := _LombardVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseFeeTokenWithdrawn(log types.Log) (*LombardVerifierFeeTokenWithdrawn, error) {
	event := new(LombardVerifierFeeTokenWithdrawn)
	if err := _LombardVerifier.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierOwnershipTransferRequestedIterator struct {
	Event *LombardVerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierOwnershipTransferRequested)
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
		it.Event = new(LombardVerifierOwnershipTransferRequested)
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

func (it *LombardVerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierOwnershipTransferRequestedIterator{contract: _LombardVerifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierOwnershipTransferRequested)
				if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*LombardVerifierOwnershipTransferRequested, error) {
	event := new(LombardVerifierOwnershipTransferRequested)
	if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierOwnershipTransferredIterator struct {
	Event *LombardVerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierOwnershipTransferred)
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
		it.Event = new(LombardVerifierOwnershipTransferred)
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

func (it *LombardVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierOwnershipTransferredIterator{contract: _LombardVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierOwnershipTransferred)
				if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*LombardVerifierOwnershipTransferred, error) {
	event := new(LombardVerifierOwnershipTransferred)
	if err := _LombardVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierPathRemovedIterator struct {
	Event *LombardVerifierPathRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierPathRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierPathRemoved)
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
		it.Event = new(LombardVerifierPathRemoved)
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

func (it *LombardVerifierPathRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierPathRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierPathRemoved struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierPathRemovedIterator{contract: _LombardVerifier.contract, event: "PathRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierPathRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "PathRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParsePathRemoved(log types.Log) (*LombardVerifierPathRemoved, error) {
	event := new(LombardVerifierPathRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "PathRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierPathSetIterator struct {
	Event *LombardVerifierPathSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierPathSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierPathSet)
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
		it.Event = new(LombardVerifierPathSet)
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

func (it *LombardVerifierPathSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierPathSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierPathSet struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierPathSetIterator{contract: _LombardVerifier.contract, event: "PathSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierPathSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "PathSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParsePathSet(log types.Log) (*LombardVerifierPathSet, error) {
	event := new(LombardVerifierPathSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "PathSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierRemoteChainConfigSetIterator struct {
	Event *LombardVerifierRemoteChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierRemoteChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierRemoteChainConfigSet)
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
		it.Event = new(LombardVerifierRemoteChainConfigSet)
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

func (it *LombardVerifierRemoteChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierRemoteChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierRemoteChainConfigSet struct {
	RemoteChainSelector uint64
	Router              common.Address
	AllowlistEnabled    bool
	Raw                 types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardVerifierRemoteChainConfigSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardVerifierRemoteChainConfigSetIterator{contract: _LombardVerifier.contract, event: "RemoteChainConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "RemoteChainConfigSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierRemoteChainConfigSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseRemoteChainConfigSet(log types.Log) (*LombardVerifierRemoteChainConfigSet, error) {
	event := new(LombardVerifierRemoteChainConfigSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "RemoteChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierStorageLocationsUpdatedIterator struct {
	Event *LombardVerifierStorageLocationsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierStorageLocationsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierStorageLocationsUpdated)
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
		it.Event = new(LombardVerifierStorageLocationsUpdated)
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

func (it *LombardVerifierStorageLocationsUpdatedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierStorageLocationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierStorageLocationsUpdated struct {
	OldLocations []string
	NewLocations []string
	Raw          types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*LombardVerifierStorageLocationsUpdatedIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierStorageLocationsUpdatedIterator{contract: _LombardVerifier.contract, event: "StorageLocationsUpdated", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *LombardVerifierStorageLocationsUpdated) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "StorageLocationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierStorageLocationsUpdated)
				if err := _LombardVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseStorageLocationsUpdated(log types.Log) (*LombardVerifierStorageLocationsUpdated, error) {
	event := new(LombardVerifierStorageLocationsUpdated)
	if err := _LombardVerifier.contract.UnpackLog(event, "StorageLocationsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierSupportedTokenRemovedIterator struct {
	Event *LombardVerifierSupportedTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierSupportedTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierSupportedTokenRemoved)
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
		it.Event = new(LombardVerifierSupportedTokenRemoved)
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

func (it *LombardVerifierSupportedTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierSupportedTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierSupportedTokenRemoved struct {
	Token common.Address
	Raw   types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterSupportedTokenRemoved(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenRemovedIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "SupportedTokenRemoved")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierSupportedTokenRemovedIterator{contract: _LombardVerifier.contract, event: "SupportedTokenRemoved", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchSupportedTokenRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "SupportedTokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierSupportedTokenRemoved)
				if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenRemoved", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseSupportedTokenRemoved(log types.Log) (*LombardVerifierSupportedTokenRemoved, error) {
	event := new(LombardVerifierSupportedTokenRemoved)
	if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardVerifierSupportedTokenSetIterator struct {
	Event *LombardVerifierSupportedTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardVerifierSupportedTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardVerifierSupportedTokenSet)
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
		it.Event = new(LombardVerifierSupportedTokenSet)
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

func (it *LombardVerifierSupportedTokenSetIterator) Error() error {
	return it.fail
}

func (it *LombardVerifierSupportedTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardVerifierSupportedTokenSet struct {
	LocalToken   common.Address
	LocalAdapter common.Address
	Raw          types.Log
}

func (_LombardVerifier *LombardVerifierFilterer) FilterSupportedTokenSet(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenSetIterator, error) {

	logs, sub, err := _LombardVerifier.contract.FilterLogs(opts, "SupportedTokenSet")
	if err != nil {
		return nil, err
	}
	return &LombardVerifierSupportedTokenSetIterator{contract: _LombardVerifier.contract, event: "SupportedTokenSet", logs: logs, sub: sub}, nil
}

func (_LombardVerifier *LombardVerifierFilterer) WatchSupportedTokenSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenSet) (event.Subscription, error) {

	logs, sub, err := _LombardVerifier.contract.WatchLogs(opts, "SupportedTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardVerifierSupportedTokenSet)
				if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenSet", log); err != nil {
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

func (_LombardVerifier *LombardVerifierFilterer) ParseSupportedTokenSet(log types.Log) (*LombardVerifierSupportedTokenSet, error) {
	event := new(LombardVerifierSupportedTokenSet)
	if err := _LombardVerifier.contract.UnpackLog(event, "SupportedTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetFee struct {
	FeeUSDCents        uint16
	GasForVerification uint32
	PayloadSizeBytes   uint32
}
type GetRemoteChainConfig struct {
	AllowlistEnabled   bool
	Router             common.Address
	AllowedSendersList []common.Address
}

func (LombardVerifierAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (LombardVerifierAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (LombardVerifierDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf6c55d191ab03af25fd3025708a62a6038eb78ea86b0afb256fc3df66c860f60")
}

func (LombardVerifierFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (LombardVerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LombardVerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LombardVerifierPathRemoved) Topic() common.Hash {
	return common.HexToHash("0x8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b")
}

func (LombardVerifierPathSet) Topic() common.Hash {
	return common.HexToHash("0x83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9")
}

func (LombardVerifierRemoteChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc9")
}

func (LombardVerifierStorageLocationsUpdated) Topic() common.Hash {
	return common.HexToHash("0xec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d07586")
}

func (LombardVerifierSupportedTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0xbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a306")
}

func (LombardVerifierSupportedTokenSet) Topic() common.Hash {
	return common.HexToHash("0x086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e3")
}

func (_LombardVerifier *LombardVerifier) Address() common.Address {
	return _LombardVerifier.address
}

type LombardVerifierInterface interface {
	GetDynamicConfig(opts *bind.CallOpts) (LombardVerifierDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte, arg3 uint16) (GetFee,

		error)

	GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardVerifierPath, error)

	GetRemoteChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (GetRemoteChainConfig,

		error)

	GetStorageLocations(opts *bind.CallOpts) ([]string, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	IBridge(opts *bind.CallOpts) (common.Address, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRemoteChainConfigUpdates(opts *bind.TransactOpts, remoteChainConfigArgs []BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error)

	RemovePaths(opts *bind.TransactOpts, remoteChainSelectors []uint64) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig LombardVerifierDynamicConfig) (*types.Transaction, error)

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSupportedTokens(opts *bind.TransactOpts, tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*LombardVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*LombardVerifierAllowListSendersRemoved, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardVerifierDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LombardVerifierDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardVerifierFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardVerifierFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LombardVerifierFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LombardVerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardVerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardVerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LombardVerifierOwnershipTransferred, error)

	FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathRemovedIterator, error)

	WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathRemoved(log types.Log) (*LombardVerifierPathRemoved, error)

	FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardVerifierPathSetIterator, error)

	WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathSet(log types.Log) (*LombardVerifierPathSet, error)

	FilterRemoteChainConfigSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardVerifierRemoteChainConfigSetIterator, error)

	WatchRemoteChainConfigSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierRemoteChainConfigSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigSet(log types.Log) (*LombardVerifierRemoteChainConfigSet, error)

	FilterStorageLocationsUpdated(opts *bind.FilterOpts) (*LombardVerifierStorageLocationsUpdatedIterator, error)

	WatchStorageLocationsUpdated(opts *bind.WatchOpts, sink chan<- *LombardVerifierStorageLocationsUpdated) (event.Subscription, error)

	ParseStorageLocationsUpdated(log types.Log) (*LombardVerifierStorageLocationsUpdated, error)

	FilterSupportedTokenRemoved(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenRemovedIterator, error)

	WatchSupportedTokenRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenRemoved) (event.Subscription, error)

	ParseSupportedTokenRemoved(log types.Log) (*LombardVerifierSupportedTokenRemoved, error)

	FilterSupportedTokenSet(opts *bind.FilterOpts) (*LombardVerifierSupportedTokenSetIterator, error)

	WatchSupportedTokenSet(opts *bind.WatchOpts, sink chan<- *LombardVerifierSupportedTokenSet) (event.Subscription, error)

	ParseSupportedTokenSet(log types.Log) (*LombardVerifierSupportedTokenSet, error)

	Address() common.Address
}
