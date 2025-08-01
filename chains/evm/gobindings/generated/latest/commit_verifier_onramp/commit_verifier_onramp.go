// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_verifier_onramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type CommitVerifierOnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type CommitVerifierOnRampDestChainConfigArgs struct {
	DestChainSelector  uint64
	VerifierAggregator common.Address
	AllowlistEnabled   bool
}

type CommitVerifierOnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	MessageInterceptor     common.Address
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

type CommitVerifierOnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenAmounts   []InternalEVM2AnyTokenTransfer
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

var CommitVerifierOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdminSet\",\"inputs\":[{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByVerifierAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6101006040523461053c57612cea8038038061001a81610560565b92833981019080820390610140821261053c576080821261053c5761003d610541565b9061004781610585565b82526020810151906001600160a01b038216820361053c576020830191825261007260408201610599565b926040810193845260a061008860608401610599565b6060830190815295607f19011261053c5760405160a081016001600160401b03811182821017610526576040526100c160808401610599565b81526100cf60a084016105ad565b602082019081526100e260c08501610599565b91604081019283526100f660e08601610599565b936060820194855261010b6101008701610599565b6080830190815261012087015190966001600160401b03821161053c57018a601f8201121561053c578051906001600160401b0382116105265760209b8c6060610159828660051b01610560565b9e8f8681520194028301019181831161053c57602001925b8284106104c2575050505033156104b157600180546001600160a01b0319163317905580516001600160401b031615801561049f575b801561048d575b801561047b575b61044e57516001600160401b0316608081905295516001600160a01b0390811660a08190529751811660c08190529851811660e08190528251909116158015610469575b801561045f575b61044e57815160028054855160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558451600380549183166001600160a01b03199283161790558651600480549184169183169190911790558751600580549190931691161790557fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19861012098606061029f610541565b8a8152602080820193845260408083019586529290910194855281519a8b5291516001600160a01b03908116928b019290925291518116918901919091529051811660608801529051811660808701529051151560a08601529051811660c08501529051811660e0840152905116610100820152a16000905b80518210156103ea5761032b82826105ba565b51916001600160401b0361033f82846105ba565b5151169283156103d557600084815260066020908152604091829020818401518154948401516001600160a81b0319909516600882901b610100600160a81b03161794151560ff1694851790915582516001600160a01b03909116815292151590830152929360019390917f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9190a20190610318565b8363c35aa79d60e01b60005260045260246000fd5b60405161270590816105e58239608051818181610e7601528181611a230152611ec5015260a051818181610eaf015281816113a30152611efe015260c051818181610eeb01528181611b580152611f3a015260e051818181610f270152611f760152f35b6306b7c75960e31b60005260046000fd5b5082511515610200565b5084516001600160a01b0316156101f9565b5088516001600160a01b0316156101b5565b5087516001600160a01b0316156101ae565b5086516001600160a01b0316156101a7565b639b15e16f60e01b60005260046000fd5b60608483031261053c576040519060608201906001600160401b03821183831017610526576060926020926040526104f987610585565b8152610506838801610599565b83820152610516604088016105ad565b6040820152815201930192610171565b634e487b7160e01b600052604160045260246000fd5b600080fd5b60405190608082016001600160401b0381118382101761052657604052565b6040519190601f01601f191682016001600160401b0381118382101761052657604052565b51906001600160401b038216820361053c57565b51906001600160a01b038216820361053c57565b5190811515820361053c57565b80518210156105ce5760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816306285c6914611e5e575080631234eab01461170a578063181f5a771461168d57806320487ded146112b65780632716072b1461104457806327e936f114610c215780635cb80c5d1461093e5780636def4ce7146108c35780637437ff9f1461078857806379ba50971461069f5780638da5cb5b1461064d578063972b46121461057c578063c9b146b3146101ae5763f2fde38b146100b957600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95773ffffffffffffffffffffffffffffffffffffffff610105612140565b61010d6123a2565b1633811461017f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a9576101fd903690600401612191565b73ffffffffffffffffffffffffffffffffffffffff600154163303610532575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610530576000938060051b8401358281121561052c5784019160808336031261052c576040519461027b86612049565b61028484612113565b865261029260208501612184565b9660208701978852604085013567ffffffffffffffff8111610528576102bb903690870161233d565b9460408801958652606081013567ffffffffffffffff8111610524576102e39136910161233d565b946060880195865267ffffffffffffffff88511682526006602052604082209851151561033b818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b8151516103f8575b50959760010195505b8451805182101561038b579061038473ffffffffffffffffffffffffffffffffffffffff61037c83600195612253565b511688612496565b500161034c565b50509594909350600192519081516103a9575b505001929190610247565b6103ee67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586925116926040519182916020835260208301906121c2565b0390a2858061039e565b98939592909497989691966000146104ed57600184019591875b8651805182101561048f5761043c8273ffffffffffffffffffffffffffffffffffffffff92612253565b5116801561045857906104516001928a612405565b5001610412565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816104e367ffffffffffffffff8b511692516040519182916020835260208301906121c2565b0390a29089610343565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff6005541633031561021d577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95767ffffffffffffffff6105bc6120fc565b1680600052600660205260ff604060002054169060005260066020526001604060002001906040518083602082955493848152019060005260206000209260005b81811061063457505061061292500383612081565b610630604051928392151583526040602084015260408301906121c2565b0390f35b84548352600194850194879450602090930192016105fd565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760005473ffffffffffffffffffffffffffffffffffffffff8116330361075e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957600060806040516107c781612065565b828152826020820152826040820152826060820152015260a06040516107ec81612065565b60ff60025473ffffffffffffffffffffffffffffffffffffffff81168352831c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015273ffffffffffffffffffffffffffffffffffffffff60045416606082015273ffffffffffffffffffffffffffffffffffffffff6005541660808201526108c1604051809273ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565bf35b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95767ffffffffffffffff6109036120fc565b1660005260066020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519160ff81161515835260081c166020820152f35b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a95761098d903690600401612191565b73ffffffffffffffffffffffffffffffffffffffff6004541660005b82811015610530576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff8116809103610524576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115610c16579085918591610bde575b5080610a33575b50505060019150016109a9565b610aec60405160208101967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602483015283604483015260448252610a7c606483612081565b80806040998a5194610a8e8c87612081565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208701525190828a5af13d15610bd5573d610ace816120c2565b90610adb8b519283612081565b8152809260203d92013e5b86612628565b805180610b2a575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381610a26565b610b41929495969350602080918301019101612296565b15610b525792919084908880610af4565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609150610ae6565b9150506020813d8211610c0e575b81610bf960209383612081565b81010312610c0a5784905188610a1f565b8380fd5b3d9150610bec565b6040513d86823e3d90fd5b346101a95760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576000604051610c5e81612065565b610c66612140565b81526024358015158103610524576020820190815260443573ffffffffffffffffffffffffffffffffffffffff81168103610c0a57604083019081526064359073ffffffffffffffffffffffffffffffffffffffff8216820361104057606084019182526084359273ffffffffffffffffffffffffffffffffffffffff8416840361052c5760808501938452610cfa6123a2565b73ffffffffffffffffffffffffffffffffffffffff855116158015611021575b8015611017575b610fef579273ffffffffffffffffffffffffffffffffffffffff859381809461012097827fc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f19a51167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff00000000000000000000000000000000000000006005541617600555610feb60405191610e6b83612049565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602084015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604084015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060840152610f9b604051809473ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b608083019073ffffffffffffffffffffffffffffffffffffffff60808092828151168552602081015115156020860152826040820151166040860152826060820151166060860152015116910152565ba180f35b6004867f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515610d21565b5073ffffffffffffffffffffffffffffffffffffffff83511615610d1a565b8480fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957366023820112156101a957806004013561109e81612128565b916110ac6040519384612081565b81835260246060602085019302820101903682116101a957602401915b81831061122257836110d96123a2565b6000905b8051821015610530576110f08282612253565b519167ffffffffffffffff6111058284612253565b5151169283156111f457927f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c60406001949583600052600660205260ff826000206111c88460208501519483547fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff008860081b16911617845501511515829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a201906110dd565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6060833603126101a95760405190606082019082821067ffffffffffffffff8311176112875760609260209260405261125a86612113565b8152611267838701612163565b8382015261127760408701612184565b60408201528152019201916110c9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a9576112ed6120fc565b60243567ffffffffffffffff81116101a957806004018136039060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101a95767ffffffffffffffff84169377ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115dd5760009161165e575b506116305773ffffffffffffffffffffffffffffffffffffffff6002541692604051947fd8694ccd00000000000000000000000000000000000000000000000000000000865260048601526040602486015261148461144761143684806122ae565b60a060448a015260e48901916122fe565b61145460248401856122ae565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8984030160648a01526122fe565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd604483013591018112156101a95781016024600482013591019367ffffffffffffffff82116101a9578160061b360385136101a9578681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160848801528181528694602090910193929160005b8181106115e95750505061158d6020959361155d869460848573ffffffffffffffffffffffffffffffffffffffff61155060648a9901612163565b1660a488015201906122ae565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8584030160c48601526122fe565b03915afa80156115dd576000906115aa575b602090604051908152f35b506020813d6020116115d5575b816115c460209383612081565b810103126101a9576020905161159f565b3d91506115b7565b6040513d6000823e3d90fd5b91955091929360408060019273ffffffffffffffffffffffffffffffffffffffff6116138a612163565b168152602089013560208201520196019101918795949392611515565b837ffdbd6a720000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611680915060203d602011611686575b6116788183612081565b810190612296565b856113d4565b503d61166e565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95761063060408051906116ce8183612081565b601882527f436f6d6d6974566572696669657220312e372e302d6465760000000000000000602083015251918291602083526020830190611fea565b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957366023820112156101a957806004013567ffffffffffffffff81116101a95781019060248201903682116101a9576020818403126101a95760248101359067ffffffffffffffff82116101a95701808303926101c084126101a9576040519361016085019085821067ffffffffffffffff83111761128757608091604052126101a9576040516117da81612049565b602483013581526117ed60448401612113565b60208201526117fe60648401612113565b604082015261180f60848401612113565b6060820152845261182260a48301612163565b926020850193845260c483013567ffffffffffffffff81116101a95781602461184d9286010161220c565b604086015260e483013567ffffffffffffffff81116101a9578160246118759286010161220c565b606086015261010483013567ffffffffffffffff81116101a95781602461189e9286010161220c565b608086015261012483013567ffffffffffffffff81116101a9576024908401019281601f850112156101a95783356118d581612128565b946118e36040519687612081565b81865260208087019260051b820101918483116101a95760208201905b838210611e30575050505060a0860193845261191f6101448201612163565b60c087015261016481013560e08701526101848101356101008701526101a481013567ffffffffffffffff81116101a95760249082010182601f820112156101a95780359061196d82612128565b9161197b6040519384612081565b80835260208084019160051b830101918583116101a95760208101915b838310611d3957505050506101208701526101c481013567ffffffffffffffff81116101a957602491010181601f820112156101a9578035926119da84612128565b936119e86040519586612081565b80855260208086019160051b840101928484116101a95760208101915b848310611c8a5750505050505061014084015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166000526006602052604060002080549060ff8216611c10575b5060081c73ffffffffffffffffffffffffffffffffffffffff163303611be757611a8d611a9d916024359051612253565b5160208082518301019101612296565b9160009215611ae1575b6106308367ffffffffffffffff6040519116602082015260208152611acd604082612081565b604051918291602083526020830190611fea565b67ffffffffffffffff604073ffffffffffffffffffffffffffffffffffffffff9251015116915116604051917fea458c0c000000000000000000000000000000000000000000000000000000008352600483015260248201526020816044818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611bdc578291611b94575b50905061063082611aa7565b90506020813d602011611bd4575b81611baf60209383612081565b81010312610528575167ffffffffffffffff8116810361052857610630915082611b88565b3d9150611ba2565b6040513d84823e3d90fd5b7ed135dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b835173ffffffffffffffffffffffffffffffffffffffff1660009081526002909101602052604090205415611c455784611a5c565b73ffffffffffffffffffffffffffffffffffffffff8351167fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b823567ffffffffffffffff81116101a95782019060a0600483870301126101a95760405190611cb882612065565b60208301358252604083013567ffffffffffffffff81116101a957886020611ce29286010161220c565b602083015260608301356040830152611cfd60808401612113565b606083015260a08301359167ffffffffffffffff83116101a957611d298960208096958196010161220c565b6080820152815201920191611a05565b823567ffffffffffffffff81116101a957820160e06004828b0301126101a9576040519160e0830183811067ffffffffffffffff82111761128757604052611d8360208301612163565b8352611d9160408301612163565b6020840152606082013567ffffffffffffffff81116101a957896020611db99285010161220c565b6040840152608082013567ffffffffffffffff81116101a957896020611de19285010161220c565b606084015260a0820135608084015260c08201359267ffffffffffffffff84116101a95760e083611e198c602080988198010161220c565b60a0840152013560c0820152815201920191611998565b813567ffffffffffffffff81116101a957602091611e538884809488010161220c565b815201910190611900565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957606081611e9b600093612049565b82815282602082015282604082015201526080604051611eba81612049565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660608201526108c1604051809273ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b919082519283825260005b8481106120345750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611ff5565b6080810190811067ffffffffffffffff82111761128757604052565b60a0810190811067ffffffffffffffff82111761128757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761128757604052565b67ffffffffffffffff811161128757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6004359067ffffffffffffffff821682036101a957565b359067ffffffffffffffff821682036101a957565b67ffffffffffffffff81116112875760051b60200190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b359081151582036101a957565b9181601f840112156101a95782359167ffffffffffffffff83116101a9576020808501948460051b0101116101a957565b906020808351928381520192019060005b8181106121e05750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016121d3565b81601f820112156101a957803590612223826120c2565b926122316040519485612081565b828452602083830101116101a957816000926020809301838601378301015290565b80518210156122675760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b908160209103126101a9575180151581036101a95790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101a957016020813591019167ffffffffffffffff82116101a95781360383136101a957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9080601f830112156101a957813561235481612128565b926123626040519485612081565b81845260208085019260051b8201019283116101a957602001905b82821061238a5750505090565b6020809161239784612163565b81520191019061237d565b73ffffffffffffffffffffffffffffffffffffffff6001541633036123c357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156122675760005260206000200190600090565b600082815260018201602052604090205461248f578054906801000000000000000082101561128757826124786124438460018096018555846123ed565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b906001820191816000528260205260406000205480151560001461261f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116125f0578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116125f0578181036125b9575b5050508054801561258a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061254b82826123ed565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6125d96125c961244393866123ed565b90549060031b1c928392866123ed565b905560005283602052604060002055388080612513565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b919290156126a3575081511561263c575090565b3b156126455790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156126b65750805190602001fd5b6126f4906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611fea565b0390fdfea164736f6c634300081a000a",
}

var CommitVerifierOnRampABI = CommitVerifierOnRampMetaData.ABI

var CommitVerifierOnRampBin = CommitVerifierOnRampMetaData.Bin

func DeployCommitVerifierOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CommitVerifierOnRampStaticConfig, dynamicConfig CommitVerifierOnRampDynamicConfig, destChainConfigArgs []CommitVerifierOnRampDestChainConfigArgs) (common.Address, *types.Transaction, *CommitVerifierOnRamp, error) {
	parsed, err := CommitVerifierOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitVerifierOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitVerifierOnRamp{address: address, abi: *parsed, CommitVerifierOnRampCaller: CommitVerifierOnRampCaller{contract: contract}, CommitVerifierOnRampTransactor: CommitVerifierOnRampTransactor{contract: contract}, CommitVerifierOnRampFilterer: CommitVerifierOnRampFilterer{contract: contract}}, nil
}

type CommitVerifierOnRamp struct {
	address common.Address
	abi     abi.ABI
	CommitVerifierOnRampCaller
	CommitVerifierOnRampTransactor
	CommitVerifierOnRampFilterer
}

type CommitVerifierOnRampCaller struct {
	contract *bind.BoundContract
}

type CommitVerifierOnRampTransactor struct {
	contract *bind.BoundContract
}

type CommitVerifierOnRampFilterer struct {
	contract *bind.BoundContract
}

type CommitVerifierOnRampSession struct {
	Contract     *CommitVerifierOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitVerifierOnRampCallerSession struct {
	Contract *CommitVerifierOnRampCaller
	CallOpts bind.CallOpts
}

type CommitVerifierOnRampTransactorSession struct {
	Contract     *CommitVerifierOnRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitVerifierOnRampRaw struct {
	Contract *CommitVerifierOnRamp
}

type CommitVerifierOnRampCallerRaw struct {
	Contract *CommitVerifierOnRampCaller
}

type CommitVerifierOnRampTransactorRaw struct {
	Contract *CommitVerifierOnRampTransactor
}

func NewCommitVerifierOnRamp(address common.Address, backend bind.ContractBackend) (*CommitVerifierOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitVerifierOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitVerifierOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRamp{address: address, abi: abi, CommitVerifierOnRampCaller: CommitVerifierOnRampCaller{contract: contract}, CommitVerifierOnRampTransactor: CommitVerifierOnRampTransactor{contract: contract}, CommitVerifierOnRampFilterer: CommitVerifierOnRampFilterer{contract: contract}}, nil
}

func NewCommitVerifierOnRampCaller(address common.Address, caller bind.ContractCaller) (*CommitVerifierOnRampCaller, error) {
	contract, err := bindCommitVerifierOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampCaller{contract: contract}, nil
}

func NewCommitVerifierOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitVerifierOnRampTransactor, error) {
	contract, err := bindCommitVerifierOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampTransactor{contract: contract}, nil
}

func NewCommitVerifierOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitVerifierOnRampFilterer, error) {
	contract, err := bindCommitVerifierOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampFilterer{contract: contract}, nil
}

func bindCommitVerifierOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitVerifierOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitVerifierOnRamp.Contract.CommitVerifierOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.CommitVerifierOnRampTransactor.contract.Transfer(opts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.CommitVerifierOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitVerifierOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.contract.Transfer(opts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

	error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	outstruct := new(GetAllowedSendersList)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfiguredAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _CommitVerifierOnRamp.Contract.GetAllowedSendersList(&_CommitVerifierOnRamp.CallOpts, destChainSelector)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _CommitVerifierOnRamp.Contract.GetAllowedSendersList(&_CommitVerifierOnRamp.CallOpts, destChainSelector)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.VerifierAggregator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitVerifierOnRamp.Contract.GetDestChainConfig(&_CommitVerifierOnRamp.CallOpts, destChainSelector)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitVerifierOnRamp.Contract.GetDestChainConfig(&_CommitVerifierOnRamp.CallOpts, destChainSelector)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitVerifierOnRampDynamicConfig, error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitVerifierOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitVerifierOnRampDynamicConfig)).(*CommitVerifierOnRampDynamicConfig)

	return out0, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) GetDynamicConfig() (CommitVerifierOnRampDynamicConfig, error) {
	return _CommitVerifierOnRamp.Contract.GetDynamicConfig(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) GetDynamicConfig() (CommitVerifierOnRampDynamicConfig, error) {
	return _CommitVerifierOnRamp.Contract.GetDynamicConfig(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CommitVerifierOnRamp.Contract.GetFee(&_CommitVerifierOnRamp.CallOpts, destChainSelector, message)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CommitVerifierOnRamp.Contract.GetFee(&_CommitVerifierOnRamp.CallOpts, destChainSelector, message)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (CommitVerifierOnRampStaticConfig, error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CommitVerifierOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitVerifierOnRampStaticConfig)).(*CommitVerifierOnRampStaticConfig)

	return out0, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) GetStaticConfig() (CommitVerifierOnRampStaticConfig, error) {
	return _CommitVerifierOnRamp.Contract.GetStaticConfig(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) GetStaticConfig() (CommitVerifierOnRampStaticConfig, error) {
	return _CommitVerifierOnRamp.Contract.GetStaticConfig(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) Owner() (common.Address, error) {
	return _CommitVerifierOnRamp.Contract.Owner(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) Owner() (common.Address, error) {
	return _CommitVerifierOnRamp.Contract.Owner(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitVerifierOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) TypeAndVersion() (string, error) {
	return _CommitVerifierOnRamp.Contract.TypeAndVersion(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitVerifierOnRamp.Contract.TypeAndVersion(&_CommitVerifierOnRamp.CallOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.AcceptOwnership(&_CommitVerifierOnRamp.TransactOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.AcceptOwnership(&_CommitVerifierOnRamp.TransactOpts)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []CommitVerifierOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []CommitVerifierOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ApplyAllowlistUpdates(&_CommitVerifierOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []CommitVerifierOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ApplyAllowlistUpdates(&_CommitVerifierOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CommitVerifierOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CommitVerifierOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitVerifierOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CommitVerifierOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitVerifierOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "forwardToVerifier", rawMessage, verifierIndex)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ForwardToVerifier(&_CommitVerifierOnRamp.TransactOpts, rawMessage, verifierIndex)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.ForwardToVerifier(&_CommitVerifierOnRamp.TransactOpts, rawMessage, verifierIndex)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitVerifierOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) SetDynamicConfig(dynamicConfig CommitVerifierOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.SetDynamicConfig(&_CommitVerifierOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) SetDynamicConfig(dynamicConfig CommitVerifierOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.SetDynamicConfig(&_CommitVerifierOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.TransferOwnership(&_CommitVerifierOnRamp.TransactOpts, to)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.TransferOwnership(&_CommitVerifierOnRamp.TransactOpts, to)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.WithdrawFeeTokens(&_CommitVerifierOnRamp.TransactOpts, feeTokens)
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitVerifierOnRamp.Contract.WithdrawFeeTokens(&_CommitVerifierOnRamp.TransactOpts, feeTokens)
}

type CommitVerifierOnRampAllowListAdminSetIterator struct {
	Event *CommitVerifierOnRampAllowListAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampAllowListAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampAllowListAdminSet)
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
		it.Event = new(CommitVerifierOnRampAllowListAdminSet)
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

func (it *CommitVerifierOnRampAllowListAdminSetIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampAllowListAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampAllowListAdminSet struct {
	AllowlistAdmin common.Address
	Raw            types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*CommitVerifierOnRampAllowListAdminSetIterator, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampAllowListAdminSetIterator{contract: _CommitVerifierOnRamp.contract, event: "AllowListAdminSet", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error) {

	var allowlistAdminRule []interface{}
	for _, allowlistAdminItem := range allowlistAdmin {
		allowlistAdminRule = append(allowlistAdminRule, allowlistAdminItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "AllowListAdminSet", allowlistAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampAllowListAdminSet)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseAllowListAdminSet(log types.Log) (*CommitVerifierOnRampAllowListAdminSet, error) {
	event := new(CommitVerifierOnRampAllowListAdminSet)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampAllowListSendersAddedIterator struct {
	Event *CommitVerifierOnRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampAllowListSendersAdded)
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
		it.Event = new(CommitVerifierOnRampAllowListSendersAdded)
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

func (it *CommitVerifierOnRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampAllowListSendersAddedIterator{contract: _CommitVerifierOnRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampAllowListSendersAdded)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitVerifierOnRampAllowListSendersAdded, error) {
	event := new(CommitVerifierOnRampAllowListSendersAdded)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampAllowListSendersRemovedIterator struct {
	Event *CommitVerifierOnRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampAllowListSendersRemoved)
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
		it.Event = new(CommitVerifierOnRampAllowListSendersRemoved)
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

func (it *CommitVerifierOnRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampAllowListSendersRemovedIterator{contract: _CommitVerifierOnRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampAllowListSendersRemoved)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitVerifierOnRampAllowListSendersRemoved, error) {
	event := new(CommitVerifierOnRampAllowListSendersRemoved)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampCCIPMessageSentIterator struct {
	Event *CommitVerifierOnRampCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampCCIPMessageSent)
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
		it.Event = new(CommitVerifierOnRampCCIPMessageSent)
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

func (it *CommitVerifierOnRampCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampCCIPMessageSent struct {
	DestChainSelector uint64
	Message           InternalEVM2AnyRampMessage
	Raw               types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampCCIPMessageSentIterator{contract: _CommitVerifierOnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampCCIPMessageSent, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampCCIPMessageSent)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseCCIPMessageSent(log types.Log) (*CommitVerifierOnRampCCIPMessageSent, error) {
	event := new(CommitVerifierOnRampCCIPMessageSent)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampConfigSetIterator struct {
	Event *CommitVerifierOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampConfigSet)
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
		it.Event = new(CommitVerifierOnRampConfigSet)
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

func (it *CommitVerifierOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampConfigSet struct {
	StaticConfig  CommitVerifierOnRampStaticConfig
	DynamicConfig CommitVerifierOnRampDynamicConfig
	Raw           types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitVerifierOnRampConfigSetIterator, error) {

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampConfigSetIterator{contract: _CommitVerifierOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampConfigSet)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseConfigSet(log types.Log) (*CommitVerifierOnRampConfigSet, error) {
	event := new(CommitVerifierOnRampConfigSet)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampDestChainConfigSetIterator struct {
	Event *CommitVerifierOnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampDestChainConfigSet)
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
		it.Event = new(CommitVerifierOnRampDestChainConfigSet)
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

func (it *CommitVerifierOnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampDestChainConfigSetIterator{contract: _CommitVerifierOnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampDestChainConfigSet)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseDestChainConfigSet(log types.Log) (*CommitVerifierOnRampDestChainConfigSet, error) {
	event := new(CommitVerifierOnRampDestChainConfigSet)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampFeeTokenWithdrawnIterator struct {
	Event *CommitVerifierOnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampFeeTokenWithdrawn)
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
		it.Event = new(CommitVerifierOnRampFeeTokenWithdrawn)
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

func (it *CommitVerifierOnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CommitVerifierOnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampFeeTokenWithdrawnIterator{contract: _CommitVerifierOnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampFeeTokenWithdrawn)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitVerifierOnRampFeeTokenWithdrawn, error) {
	event := new(CommitVerifierOnRampFeeTokenWithdrawn)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampOwnershipTransferRequestedIterator struct {
	Event *CommitVerifierOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampOwnershipTransferRequested)
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
		it.Event = new(CommitVerifierOnRampOwnershipTransferRequested)
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

func (it *CommitVerifierOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampOwnershipTransferRequestedIterator{contract: _CommitVerifierOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampOwnershipTransferRequested)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitVerifierOnRampOwnershipTransferRequested, error) {
	event := new(CommitVerifierOnRampOwnershipTransferRequested)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitVerifierOnRampOwnershipTransferredIterator struct {
	Event *CommitVerifierOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitVerifierOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitVerifierOnRampOwnershipTransferred)
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
		it.Event = new(CommitVerifierOnRampOwnershipTransferred)
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

func (it *CommitVerifierOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitVerifierOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitVerifierOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitVerifierOnRampOwnershipTransferredIterator{contract: _CommitVerifierOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitVerifierOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitVerifierOnRampOwnershipTransferred)
				if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitVerifierOnRamp *CommitVerifierOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitVerifierOnRampOwnershipTransferred, error) {
	event := new(CommitVerifierOnRampOwnershipTransferred)
	if err := _CommitVerifierOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetAllowedSendersList struct {
	IsEnabled           bool
	ConfiguredAddresses []common.Address
}
type GetDestChainConfig struct {
	AllowlistEnabled   bool
	VerifierAggregator common.Address
}

func (_CommitVerifierOnRamp *CommitVerifierOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CommitVerifierOnRamp.abi.Events["AllowListAdminSet"].ID:
		return _CommitVerifierOnRamp.ParseAllowListAdminSet(log)
	case _CommitVerifierOnRamp.abi.Events["AllowListSendersAdded"].ID:
		return _CommitVerifierOnRamp.ParseAllowListSendersAdded(log)
	case _CommitVerifierOnRamp.abi.Events["AllowListSendersRemoved"].ID:
		return _CommitVerifierOnRamp.ParseAllowListSendersRemoved(log)
	case _CommitVerifierOnRamp.abi.Events["CCIPMessageSent"].ID:
		return _CommitVerifierOnRamp.ParseCCIPMessageSent(log)
	case _CommitVerifierOnRamp.abi.Events["ConfigSet"].ID:
		return _CommitVerifierOnRamp.ParseConfigSet(log)
	case _CommitVerifierOnRamp.abi.Events["DestChainConfigSet"].ID:
		return _CommitVerifierOnRamp.ParseDestChainConfigSet(log)
	case _CommitVerifierOnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _CommitVerifierOnRamp.ParseFeeTokenWithdrawn(log)
	case _CommitVerifierOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _CommitVerifierOnRamp.ParseOwnershipTransferRequested(log)
	case _CommitVerifierOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _CommitVerifierOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CommitVerifierOnRampAllowListAdminSet) Topic() common.Hash {
	return common.HexToHash("0xb8c9b44ae5b5e3afb195f67391d9ff50cb904f9c0fa5fd520e497a97c1aa5a1e")
}

func (CommitVerifierOnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitVerifierOnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitVerifierOnRampCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x8cd775d4a25bd349439a70817fd110144d6ab229ae1b9f54a1e5777d2041bfed")
}

func (CommitVerifierOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc7372d2d886367d7bb1b0e0708a5436f2c91d6963de210eb2dc1ec2ecd6d21f1")
}

func (CommitVerifierOnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitVerifierOnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitVerifierOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitVerifierOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CommitVerifierOnRamp *CommitVerifierOnRamp) Address() common.Address {
	return _CommitVerifierOnRamp.address
}

type CommitVerifierOnRampInterface interface {
	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitVerifierOnRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetStaticConfig(opts *bind.CallOpts) (CommitVerifierOnRampStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []CommitVerifierOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CommitVerifierOnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitVerifierOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListAdminSet(opts *bind.FilterOpts, allowlistAdmin []common.Address) (*CommitVerifierOnRampAllowListAdminSetIterator, error)

	WatchAllowListAdminSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListAdminSet, allowlistAdmin []common.Address) (event.Subscription, error)

	ParseAllowListAdminSet(log types.Log) (*CommitVerifierOnRampAllowListAdminSet, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitVerifierOnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitVerifierOnRampAllowListSendersRemoved, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampCCIPMessageSent, destChainSelector []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*CommitVerifierOnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitVerifierOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitVerifierOnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitVerifierOnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CommitVerifierOnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitVerifierOnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitVerifierOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitVerifierOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitVerifierOnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
