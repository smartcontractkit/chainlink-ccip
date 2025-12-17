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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"},{\"name\":\"storageLocation\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"rmn\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteChainConfigUpdates\",\"inputs\":[{\"name\":\"remoteChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct BaseVerifier.RemoteChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasForVerification\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payloadSizeBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardVerifier.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedSendersList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageLocations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV3\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePaths\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSupportedTokens\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokensToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct LombardVerifier.SupportedTokenArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteChainConfigSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StorageLocationsUpdated\",\"inputs\":[{\"name\":\"oldLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"newLocations\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenRemoved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SupportedTokenSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"localAdapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestGasCannotBeZero\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageId\",\"inputs\":[{\"name\":\"messageMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeMessageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustTransferTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RemoteChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]}]",
	Bin: "0x60c0806040523461057857613cbf803803809161001c82856105f5565b8339810190606081830312610578578051916001600160a01b038316918284036105785760208101516001600160401b03811161057857810182601f820112156105785780519061006c82610618565b9361007a60405195866105f5565b82855260208086019360051b830101918183116105785760208101935b83851061057d576040860151899089906001600160a01b038116908a90829003610578576001549080516100ca83610618565b926100d860405194856105f5565b808452600160009081527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690602086015b8382106104d35750505060005b81811061043f57505060005b8181106102b05750507fec9f9416b098576351ada0c342c1381ca08990ee094978ddd1003ef013d0758691610176610168926040519384936040855260408501906106be565b9083820360208501526106be565b0390a1801561029f57608052331561028e57600380546001600160a01b03191633179055801561027d5760206004916040519283809263353c26b760e01b82525afa9081156102715760009161022d575b5060ff1660018103610214575060a05260405161358f90816107308239608051816128c2015260a0518181816109a5015281816114e101528181611c9301528181611f490152612ba50152f35b63398bbe0560e11b600052600160045260245260446000fd5b6020813d602011610269575b81610246602093836105f5565b8101031261026557519060ff82168203610262575060ff6101c7565b80fd5b5080fd5b3d9150610239565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b639b15e16f60e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b82518110156104295760208160051b84010151600154680100000000000000008110156103fd578060016102e79201600155610652565b919091610413578051906001600160401b0382116103fd57610309835461066d565b601f81116103c0575b50602090601f8311600114610355576001949392916000918361034a575b5050600019600383901b1c191690841b1790555b01610122565b015190508b80610330565b90601f1983169184600052816000209260005b8181106103a857509160019695949291838895931061038f575b505050811b019055610344565b015160001960f88460031b161c191690558b8080610382565b92936020600181928786015181550195019301610368565b6103ed90846000526020600020601f850160051c810191602086106103f3575b601f0160051c01906106a7565b8a610312565b90915081906103e0565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60015480156104bd57600019019061045682610652565b929092610413578261046a6001945461066d565b908161047b575b5050825501610116565b81601f6000931186146104925750555b8a80610471565b818352602083206104ad91601f0160051c81019087016106a7565b808252816020812091555561048b565b634e487b7160e01b600052603160045260246000fd5b604051600084546104e38161066d565b8084529060018116908115610555575060011461051d575b506001928261050f859460209403826105f5565b815201930191019091610109565b6000868152602081209092505b81831061053f575050810160200160016104fb565b600181602092548386880101520192019161052a565b60ff191660208581019190915291151560051b84019091019150600190506104fb565b600080fd5b84516001600160401b0381116105785782019083603f83011215610578576020820151906001600160401b0382116103fd576040516105c6601f8401601f1916602001826105f5565b8281526040848401018610610578576105ea6020949385946040868501910161062f565b815201940193610097565b601f909101601f19168101906001600160401b038211908210176103fd57604052565b6001600160401b0381116103fd5760051b60200190565b60005b8381106106425750506000910152565b8181015183820152602001610632565b60015481101561042957600160005260206000200190600090565b90600182811c9216801561069d575b602083101461068757565b634e487b7160e01b600052602260045260246000fd5b91607f169161067c565b8181106106b2575050565b600081556001016106a7565b9080602083519182815201916020808360051b8301019401926000915b8383106106ea57505050505090565b909192939460208080600193601f1986820301875289516107168151809281855285808601910161062f565b601f01601f1916010197019594919091019201906106db56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610147578063181f5a7714610142578063240028e81461013d5780632e45ca681461013857806338ff8c38146101335780633bbbed4b1461012e578063708e1f7914610129578063737037e81461012457806379ba50971461011f57806380485e251461011a57806387ae929214610115578063898068fc146101105780638da5cb5b1461010b5780639ba4393114610106578063bcb6d4f714610101578063bff0ec1d146100fc578063c4bffe2b146100f7578063d3c7c2c7146100f25763f2fde38b146100ed57600080fd5b6118df565b611853565b611790565b61131f565b61126f565b611177565b611125565b611085565b610ea7565b610c96565b610a9a565b6109fa565b61095a565b610607565b610529565b61049a565b610427565b61037d565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361020657807f83adcde100000000000000000000000000000000000000000000000000000000602092149081156101dc575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386101d1565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761025657604052565b61020b565b60c0810190811067ffffffffffffffff82111761025657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761025657604052565b604051906102c760a083610277565b565b67ffffffffffffffff811161025657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106103165750506000910152565b8181015183820152602001610306565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361036281518092818752878088019101610303565b0116010190565b90602061037a928181520190610326565b90565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576103fa60408051906103be8183610277565b601982527f4c6f6d62617264566572696669657220312e372e302d64657600000000000000602083015251918291602083526020830190610326565b0390f35b73ffffffffffffffffffffffffffffffffffffffff81160361020657565b35906102c7826103fe565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657602061049073ffffffffffffffffffffffffffffffffffffffff60043561047c816103fe565b166000526005602052604060002054151590565b6040519015158152f35b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff8111610206573660238201121561020657806004013567ffffffffffffffff81116102065736602460c083028401011161020657602461051592016119d3565b005b67ffffffffffffffff81160361020657565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065767ffffffffffffffff60043561056d81610517565b6000602060405161057d8161023a565b82815201521660005260096020526103fa60406000206001604051916105a28361023a565b8054835201546020820152604051918291829190916020806040830194805184520151910152565b90816101c09103126102065790565b9181601f840112156102065782359167ffffffffffffffff8311610206576020838186019501011161020657565b346102065760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff8111610206576106569036906004016105ca565b602435906106656044356103fe565b60843567ffffffffffffffff8111610206576106859036906004016105d9565b5050602081018035600061069882610517565b6106a182612870565b6101808401906106b18286611aea565b905015610932576106c183610517565b6101208501926106da6106d48588611b3e565b90611b8f565b60601c906106fc8167ffffffffffffffff166000526000602052604060002090565b8054909161073473ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff8116156108fc576040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152602090829060249082905afa80156108f75773ffffffffffffffffffffffffffffffffffffffff9186916108c8575b5016330361089c5760f01c60ff1661080d575b6103fa6108018989896107f98a6107f36107ed6107e78d87611aea565b90611c24565b93611ae0565b93611b3e565b929091612a32565b60405191829182610369565b61084d61085191600161083661071b8673ffffffffffffffffffffffffffffffffffffffff1690565b910160019160005201602052604060002054151590565b1590565b61085b57806107ca565b7fd0d2597600000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff16600452602490fd5b7f728fe07b00000000000000000000000000000000000000000000000000000000845233600452602484fd5b6108ea915060203d6020116108f0575b6108e28183610277565b8101906123da565b386107b7565b503d6108d8565b61202d565b7f4d1aff7e00000000000000000000000000000000000000000000000000000000865267ffffffffffffffff8216600452602486fd5b807f4f73dc4d0000000000000000000000000000000000000000000000000000000060049252fd5b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156102065782359167ffffffffffffffff8311610206576020808501948460061b01011161020657565b346102065760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff811161020657366023820112156102065780600401359067ffffffffffffffff8211610206573660248360051b83010111610206576024359067ffffffffffffffff82116102065761051592610a9160249336906004016109c9565b93909201611c62565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760025473ffffffffffffffffffffffffffffffffffffffff81163303610b59577fffffffffffffffffffffffff00000000000000000000000000000000000000006003549133828416176003551660025573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610b8f826102c9565b91610b9d6040519384610277565b829481845281830111610206578281602093846000960137010152565b9080601f830112156102065781602061037a93359101610b83565b67ffffffffffffffff81116102565760051b60200190565b81601f8201121561020657803590610c0482610bd5565b92610c126040519485610277565b82845260208085019360061b8301019181831161020657602001925b828410610c3c575050505090565b6040848303126102065760206040918251610c568161023a565b8635610c61816103fe565b81528287013583820152815201930192610c2e565b6064359061ffff8216820361020657565b359061ffff8216820361020657565b346102065760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657600435610cd181610517565b60243567ffffffffffffffff81116102065760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261020657610d176102b8565b90806004013567ffffffffffffffff811161020657610d3c9060043691840101610bba565b8252602481013567ffffffffffffffff811161020657610d629060043691840101610bba565b6020830152604481013567ffffffffffffffff811161020657610d8b9060043691840101610bed565b6040830152610d9c6064820161041c565b6060830152608481013567ffffffffffffffff8111610206576080916004610dc79236920101610bba565b91015260443567ffffffffffffffff8111610206576103fa91610df1610e00923690600401610bba565b50610dfa610c76565b50612039565b6040805161ffff909416845263ffffffff92831660208501529116908201529081906060820190565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610e5c57505050505090565b9091929394602080610e98837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610326565b97019301930191939290610e4d565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657600154610ee281610bd5565b610eef6040519182610277565b818152602081019160016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf66000935b828510610f3657604051806103fa8682610e29565b604051600083548060011c9060018116908115611031575b602083108214611004578285526020850191908115610fcf5750600114610f93575b505060019282610f8585946020940382610277565b815201920194019390610f21565b909150610fa585600052602060002090565b916000925b818410610fbb575050018282610f70565b600181602092548686015201930192610faa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682525090151560051b0190508282610f70565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f1691610f4e565b906020808351928381520192019060005b8181106110595750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161104c565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065767ffffffffffffffff6004356110c981610517565b16600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000206103fa6110fd60018354930161309f565b604051938360ff869560f01c161515855216602084015260606040840152606083019061103b565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020657602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346102065760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576004356111b281610517565b602435906044356111c16124cb565b8215611245576112407f83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9916040516111f88161023a565b818152600167ffffffffffffffff602083019688885216958660005260096020526040600020925183555191015561122f8461311d565b506040519081529081906020820190565b0390a3005b7f5a39e3030000000000000000000000000000000000000000000000000000000060005260046000fd5b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff81116102065736602382011215610206578060040135906112ca82610bd5565b916112d86040519384610277565b8083526024602084019160051b8301019136831161020657602401905b8282106113055761051584612104565b60208091833561131481610517565b8152019101906112f5565b346102065760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065760043567ffffffffffffffff81116102065761136e9036906004016105ca565b602435906044359067ffffffffffffffff8211610206576113b76113996113b29336906004016105d9565b9390926113ad6113a882611ae0565b612870565b611ae0565b612e14565b6113ca6113c4838361225e565b906122bf565b7ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160361161c575061149c6114659161144361143c6114366114306006888661226c565b90612374565b60f01c90565b61ffff1690565b9361145b611452866006612367565b600683856122a7565b9490956006612367565b9061149661148f61148961143c61143661143061148188612354565b88888b6122a7565b93612354565b9283612367565b926122a7565b92604051927fd5438eae00000000000000000000000000000000000000000000000000000000845260208460048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9182156108f7576000948593611713575b5073ffffffffffffffffffffffffffffffffffffffff8591611565604051988997889687947fa6208506000000000000000000000000000000000000000000000000000000008652600486016124a4565b0393165af19081156108f7576000906000926116ec575b50156116c257602481510361168f5760246020820151910151907ff0f3a135000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160361161c57508181036115e957005b7f6c86fa3a0000000000000000000000000000000000000000000000000000000060005260049190915260245260446000fd5b7fadaf7739000000000000000000000000000000000000000000000000000000006000527ff0f3a135000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b517fc2fdac9800000000000000000000000000000000000000000000000000000000600052602460048190525260446000fd5b7f2532cf450000000000000000000000000000000000000000000000000000000060005260046000fd5b905061170b91503d806000833e6117038183610277565b8101906123ef565b91503861157c565b8591935061174473ffffffffffffffffffffffffffffffffffffffff9160203d6020116108f0576108e28183610277565b939150611514565b602060408183019282815284518094520192019060005b8181106117705750505090565b825167ffffffffffffffff16845260209384019390920191600101611763565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576007546117cb81610bd5565b906117d96040519283610277565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061180682610bd5565b0136602084013760005b81811061182557604051806103fa858261174c565b8067ffffffffffffffff61183a6001936130ea565b90549060031b1c1661184c828661220e565b5201611810565b346102065760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610206576040516004548082526020820190600460005260206000209060005b8181106118c9576103fa856118b581870382610277565b60405191829160208352602083019061103b565b825484526020909301926001928301920161189e565b346102065760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102065773ffffffffffffffffffffffffffffffffffffffff60043561192f816103fe565b6119376124cb565b163381146119a957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff600354167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906119dc6124cb565b6119e581610bd5565b916119f36040519384610277565b81835260c060208401920281019036821161020657915b818310611a1d575050506102c790612516565b60c08336031261020657602060c091604051611a388161025b565b8535611a43816103fe565b815282860135611a5281610517565b838201526040860135611a6481611aa7565b6040820152611a7560608701610c87565b6060820152611a8660808701611ab1565b6080820152611a9760a08701611ab1565b60a0820152815201920191611a0a565b8015150361020657565b359063ffffffff8216820361020657565b90604051611acf8161023a565b602060018294805484520154910152565b3561037a81610517565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610206570180359067ffffffffffffffff821161020657602001918160051b3603831361020657565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610206570180359067ffffffffffffffff82116102065760200191813603831361020657565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611bc3575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015611c5d578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610206570190565b611bf5565b90611c6b6124cb565b60005b818110611ec1575050509060009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915b818110611cc25750505050565b611cd5611cd0828487611fcf565b611fdf565b90611cf4825173ffffffffffffffffffffffffffffffffffffffff1690565b91611d216020820193611d1b855173ffffffffffffffffffffffffffffffffffffffff1690565b90612de0565b50825173ffffffffffffffffffffffffffffffffffffffff168015611ea257925b602060405180957f095ea7b300000000000000000000000000000000000000000000000000000000825281600073ffffffffffffffffffffffffffffffffffffffff82611dd58d6004830160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9193929373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0393165af19384156108f757611e2d7f086dcdf32d9aaaee4446c7bcf02b41c0d3b4923bf9d0265b033974e09d5f05e393611e4892600197611e76575b505173ffffffffffffffffffffffffffffffffffffffff1690565b915173ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152a101611cb5565b611e969060203d8111611e9b575b611e8e8183610277565b810190612018565b611e12565b503d611e84565b50805173ffffffffffffffffffffffffffffffffffffffff1692611d42565b80611ed7611ed26001938587611fb5565b611fc5565b611efe611ef973ffffffffffffffffffffffffffffffffffffffff831661071b565b612ff9565b611f0a575b5001611c6e565b611fac7fbea12876694c4055c71f74308f752b9027cf3d554194000a366abddfc239a30691611f8573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316612cda565b60405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a138611f03565b9190811015611c5d5760051b0190565b3561037a816103fe565b9190811015611c5d5760061b0190565b60408136031261020657602060405191611ff88361023a565b8035612003816103fe565b83520135612010816103fe565b602082015290565b90816020910312610206575161037a81611aa7565b6040513d6000823e3d90fd5b9067ffffffffffffffff821680600052600060205273ffffffffffffffffffffffffffffffffffffffff60406000205416156120d757600052600060205261ffff60406000205460a01c169063ffffffff6120cc816120ac8667ffffffffffffffff166000526000602052604060002090565b5460b01c169467ffffffffffffffff166000526000602052604060002090565b5460d01c1691929190565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9061210d6124cb565b6000915b80518310156122095767ffffffffffffffff61212d848361220e565b5116926121566121518567ffffffffffffffff166000526009602052604060002090565b611ac2565b61215f8561329d565b156121d0578461219661218a600195969767ffffffffffffffff166000526009602052604060002090565b60016000918281550155565b60208281015192516040519081527f8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b9190a3019190612111565b847fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff602491166004526000fd5b509050565b8051821015611c5d5760209160051b010190565b9161225a918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b906004116102065790600490565b909291928360041161020657831161020657600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610206578411610206578101920390565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106122f3575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906002820180921161236257565b612325565b9190820180921161236257565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106123a8575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b90816020910312610206575161037a816103fe565b909160608284031261020657815192602083015161240c81611aa7565b9260408101519067ffffffffffffffff8211610206570181601f82011215610206578051612439816102c9565b926124476040519485610277565b818452602082840101116102065761037a9160208085019101610303565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b92906124bd9061037a9593604086526040860191612465565b926020818503910152612465565b73ffffffffffffffffffffffffffffffffffffffff6003541633036124ec57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060005b82518110156122095761252d818461220e565b5161254d602061253d848761220e565b51015167ffffffffffffffff1690565b67ffffffffffffffff81169182156128385761257d8267ffffffffffffffff166000526000602052604060002090565b916125e061259f835173ffffffffffffffffffffffffffffffffffffffff1690565b849073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6126406125f06040840151151590565b84547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560f01b7eff00000000000000000000000000000000000000000000000000000000000016178455565b612699612652606084015161ffff1690565b84547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1660a09190911b75ffff000000000000000000000000000000000000000016178455565b60808201906126b86126af835163ffffffff1690565b63ffffffff1690565b156128015750916127d26127c761071b7f4cef55db91890720ca3d94563535726752813bffa29490d6d41218acb6831cc99461274a6126ff60019a99985163ffffffff1690565b86547fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1660b09190911b79ffffffff0000000000000000000000000000000000000000000016178655565b6127ad61275e60a083015163ffffffff1690565b86547fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff1660d09190911b7dffffffff000000000000000000000000000000000000000000000000000016178655565b5173ffffffffffffffffffffffffffffffffffffffff1690565b915460f01c60ff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff939093168352901515602083015290a20161251a565b7f9e7205510000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f97ccaab70000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156108f757600091612947575b506129105750565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612960915060203d602011611e9b57611e8e8183610277565b38612908565b91602061037a938181520191612465565b90604051917ff0f3a1350000000000000000000000000000000000000000000000000000000060208401526024830152602482526102c7604483610277565b9190826040910312610206576020825192015190565b939061037a97969373ffffffffffffffffffffffffffffffffffffffff60e09794819388521660208701521660408501526060840152608083015260a08201528160c08201520190610326565b90604051916020830152602082526102c7604083610277565b92919493909460808401936020612a498683611b3e565b905011612c9657612a69612a636106d46040840184611b3e565b60601c90565b92612a9a61084d73ffffffffffffffffffffffffffffffffffffffff86166000526005602052604060002054151590565b612c5257612abf6121518967ffffffffffffffff166000526009602052604060002090565b805115612c1a57604096979850612afa61071b61071b612af561071b8973ffffffffffffffffffffffffffffffffffffffff1690565b61345b565b73ffffffffffffffffffffffffffffffffffffffff8116612c10575b50612b57612b4d612b48612b41612b3a612a63612b8a999a9b60208801519b611b8f565b9b87611b3e565b3691610b83565b612f54565b9335915192612977565b92875198899788977e8a1198000000000000000000000000000000000000000000000000000000008952600489016129cc565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af180156108f75761037a91600091612bdf575b50612a19565b612c01915060403d604011612c09575b612bf98183610277565b8101906129b6565b905038612bd9565b503d612bef565b9450612b57612b16565b7fa28cbf380000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff891660045260246000fd5b7f06439c6b0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b84612ca091611b3e565b90612cd66040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401612966565b0390fd5b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff928316602482015260006044808301919091528152612db2929091612d3e606484610277565b16604090600080835194612d528587610277565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602087015260208151910182855af1903d15612dd7573d612da3612d9a826102c9565b94519485610277565b83523d6000602085013e6134be565b805180612dbd575050565b81602080612dd2936102c79501019101612018565b613014565b606092506134be565b9073ffffffffffffffffffffffffffffffffffffffff8061037a9316918260005260066020521660406000205560046131ae565b612e5261071b612e388367ffffffffffffffff166000526000602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811615612f1c576040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082905afa9081156108f757600091612efd575b5015612ecf57565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b612f16915060203d602011611e9b57611e8e8183610277565b38612ec7565b7f4d1aff7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6020815111612fc357602081519101519060208110612f92575b8060031b908082046008149015171561236257610100036101008111612362571c90565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260200360031b1b1690612f6e565b612cd6906040519182917fe0d7fb0200000000000000000000000000000000000000000000000000000000835260048301610369565b61037a9080600052600660205260006040812055600461337a565b1561301b57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b906040519182815491828252602082019060005260206000209260005b8181106130d15750506102c792500383610277565b84548352600194850194879450602090930192016130bc565b600754811015611c5d57600760005260206000200190600090565b8054821015611c5d5760005260206000200190600090565b6000818152600860205260409020546131a857600754680100000000000000008110156102565761318f61315a8260018594016007556007613105565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600754906000526008602052604060002055600190565b50600090565b6000828152600182016020526040902054613203578054906801000000000000000082101561025657826131ec61315a846001809601855584613105565b905580549260005201602052604060002055600190565b5050600090565b8054801561326e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061323f8282613105565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260086020526040902054908115613203577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161236257600754927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612362578383600095613339950361333f575b505050613328600761320a565b600890600052602052604060002090565b55600190565b61332861336b91613361613357613371956007613105565b90549060031b1c90565b9283916007613105565b90612222565b5538808061331b565b6001810191806000528260205260406000205492831515600014613452577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612362578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116123625760009585836133399761340a9503613419575b50505061320a565b90600052602052604060002090565b61343961336b916134306133576134499588613105565b92839187613105565b8590600052602052604060002090565b55388080613402565b50505050600090565b806000526006602052604060002054908115806134a8575b61347b575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5060008181526005602052604090205415613473565b9192901561353957508151156134d2575090565b3b156134db5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561354c5750805190602001fd5b612cd6906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161036956fea164736f6c634300081a000a",
}

var LombardVerifierABI = LombardVerifierMetaData.ABI

var LombardVerifierBin = LombardVerifierMetaData.Bin

func DeployLombardVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, bridge common.Address, storageLocation []string, rmn common.Address) (common.Address, *types.Transaction, *LombardVerifier, error) {
	parsed, err := LombardVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardVerifierBin), backend, bridge, storageLocation, rmn)
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

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSupportedTokens(opts *bind.TransactOpts, tokensToRemove []common.Address, tokensToSet []LombardVerifierSupportedTokenArgs) (*types.Transaction, error)

	VerifyMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvData []byte) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*LombardVerifierAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardVerifierAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *LombardVerifierAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*LombardVerifierAllowListSendersRemoved, error)

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
