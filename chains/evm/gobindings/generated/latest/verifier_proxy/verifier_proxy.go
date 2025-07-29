// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy

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

type InternalEVM2AnyCommitVerifierMessage struct {
	Header             InternalHeader
	Sender             common.Address
	Data               []byte
	Receiver           []byte
	DestChainExtraArgs []byte
	VerifierExtraArgs  [][]byte
	FeeToken           common.Address
	FeeTokenAmount     *big.Int
	FeeValueJuels      *big.Int
	TokenAmounts       []InternalEVMTokenTransfer
	RequiredVerifiers  []InternalRequiredVerifier
}

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	SourcePoolAddress  common.Address
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
	DestExecData       []byte
	RequiredVerifierId [32]byte
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalRequiredVerifier struct {
	VerifierId [32]byte
	Payload    []byte
	FeeAmount  *big.Int
	GasLimit   uint64
	ExtraArgs  []byte
}

type VerifierProxyDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
}

type VerifierProxyDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type VerifierProxyStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
	VerifierRegistry   common.Address
}

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierId\",\"inputs\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MalformedVerifierExtraArgs\",\"inputs\":[{\"name\":\"verifierExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x610100604052346104cd5761325a8038038061001a81610507565b92833981019080820361010081126104cd57608081126104cd5761003c6104e8565b906100468361052c565b82526020830151906001600160a01b03821682036104cd576020830191825261007160408501610540565b604084019081526060610085818701610540565b85820190815292607f1901126104cd5760405192606084016001600160401b038111858210176104d2576040526100be60808701610540565b845260a08601519485151586036104cd57602085019586526100e260c08801610540565b6040860190815260e088015190976001600160401b0382116104cd570188601f820112156104cd578051906001600160401b0382116104d25761012a60208360051b01610507565b996020808c858152019360061b830101918183116104cd57602001925b8284106104705750505050331561045f57600180546001600160a01b0319163317905580516001600160401b031615801561044d575b801561043b575b61040e57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0529051811660e052815116158015610429575b801561041f575b61040e5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f33306e6a63669cb4611ac048b14a2f5fc09e420dcb3dddbaedd3700085564f819260e0926000606061023f6104e8565b8281526020810183905260408101839052015260805160a05160c05186516001600160401b0390931695926001600160a01b0390811692918116911660606102856104e8565b8881526020808201938452604080830195865292909101948552815198895291516001600160a01b039081169289019290925291518116918701919091529051811660608601529051811660808501529051151560a084015290511660c0820152a16000905b80518210156103b8576102fe8282610554565b51916001600160401b036103128284610554565b5151169283156103a357600084815260046020908152604091829020928101518354600160401b600160e01b0319811682851b600160401b600160e01b03161790945582516001600160401b0390941684526001600160a01b031690830152929360019390917ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed09190a201906102eb565b8363c35aa79d60e01b60005260045260246000fd5b604051612cdb908161057f823960805181818161044801528181610b1b015261236f015260a051818181611b8901526123a8015260c0518181816123e4015261257a015260e051818181610ca501526124200152f35b6306b7c75960e31b60005260046000fd5b50815115156101c4565b5082516001600160a01b0316156101bd565b5082516001600160a01b031615610184565b5081516001600160a01b03161561017d565b639b15e16f60e01b60005260046000fd5b6040848303126104cd5760408051919082016001600160401b038111838210176104d2576040526104a08561052c565b82526020850151906001600160a01b03821682036104cd5782602092836040950152815201930192610147565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b038111838210176104d257604052565b6040519190601f01601f191682016001600160401b038111838210176104d257604052565b51906001600160401b03821682036104cd57565b51906001600160a01b03821682036104cd57565b80518210156105685760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80630242cf6014611f6057806306285c6914611ed8578063181f5a7714611e5757806320487ded14611a9c57806348a98aa414611a1b5780635cb80c5d146117185780636def4ce7146116995780637437ff9f146115be57806379ba5097146114d55780638da5cb5b146114835780639041be3d1461140357806390423fa21461118b578063df0aa9e914610218578063f2fde38b146101285763fbca3b74146100c157600080fd5b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576100f861223c565b507f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235773ffffffffffffffffffffffffffffffffffffffff6101746122f3565b61017c612bb7565b163381146101ee57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235761024f61223c565b67ffffffffffffffff602435116101235760243560040160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60243536030112610123576064359173ffffffffffffffffffffffffffffffffffffffff831683036101235760025460ff8160a01c1661116157740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff82161760025567ffffffffffffffff8216600052600460205260406000209373ffffffffffffffffffffffffffffffffffffffff8116156111375784549173ffffffffffffffffffffffffffffffffffffffff8360401c16330361110d5760006103c29161036d60846024350188612602565b73ffffffffffffffffffffffffffffffffffffffff6040939293518096819582947f11fa42270000000000000000000000000000000000000000000000000000000084526020600485015260248401916124b0565b0392165afa9283156108e95760009687928895611071575b5067ffffffffffffffff82169167ffffffffffffffff83146110425767ffffffffffffffff60017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009401169283911617905560405190610439826121ab565b6000825267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083015267ffffffffffffffff86166040830152606082015261049c610493602480350188612602565b91909780612602565b9390956104ad6064602435016125e1565b6040519490929060006104c16020886121e3565b86526000805b81811061102b575050602098604051976104e18b8a6121e3565b600089526040519c8d67ffffffffffffffff6101608281810110920111176108f557604061053d8f9273ffffffffffffffffffffffffffffffffffffffff9a8f8f8d906105479a61016089018852885216908601523691612750565b9101523691612750565b60608b015260808a015260a08901521660c087015260443560e08701526000610100870152610120860152610140850152604460243501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd60243536030181121561012357602435019460048601359567ffffffffffffffff871161012357602481018760061b803603821361012357600486916105e58b612224565b9a6105f36040519c8d6121e3565b8b52828b01940101019036821161012357915b818310610ff4575050508551957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061065661064089612224565b9861064e6040519a8b6121e3565b808a52612224565b018460005b828110610fde5750505060005b8151811015610a2b5761067b8183612787565b5190610685612717565b508582015115610a015773ffffffffffffffffffffffffffffffffffffffff6106b08184511661251b565b169182158015610969575b61092457868101519273ffffffffffffffffffffffffffffffffffffffff8251166040518060a081011067ffffffffffffffff60a0830111176108f557898b73ffffffffffffffffffffffffffffffffffffffff60009461079499828e8e60a08901604052885267ffffffffffffffff87890196168652816040890191168152606088019283526080880193845267ffffffffffffffff6040519d8e998a997f9a4575b9000000000000000000000000000000000000000000000000000000008b5260048b01525160a060248b015260c48a01906122b0565b965116604488015251166064860152516084850152511660a4830152038183855af19182156108e957600092610848575b6001945073ffffffffffffffffffffffffffffffffffffffff815116928980825192015192015192604051946107fa866121c7565b85528a85015260408401526060830152608082015260405161081c88826121e3565b6000815260a0820152600060c0820152610836828b612787565b52610841818a612787565b5001610668565b91503d93846000823e61085b85826121e3565b888186810103126101235780519467ffffffffffffffff86116101235760408683018284010312610123576040519161089383612173565b8681015167ffffffffffffffff8111610123576108b7908383019089840101612653565b83528a878201015167ffffffffffffffff8111610123576001976108df938301920101612653565b89820152916107c5565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff9051167fbf16aab60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201528781602481875afa9081156108e9576000916109d4575b50156106bb565b6109f49150883d8a116109fa575b6109ec81836121e3565b810190612448565b8b6109cd565b503d6109e2565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b8786888773ffffffffffffffffffffffffffffffffffffffff6002541691600060405180947f70a2ce5400000000000000000000000000000000000000000000000000000000825267ffffffffffffffff87166004830152604060248301528180610a99604482018b612841565b03915afa9283156108e957600093610f9e575b5060005b8551811015610adc5780610ac660019286612787565b5160a0610ad3838a612787565b51015201610ab0565b50839085610120820152604051838101907f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f7321825267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015267ffffffffffffffff8416606082015230608082015260808152610b6560a0826121e3565b51902073ffffffffffffffffffffffffffffffffffffffff848301511667ffffffffffffffff6060845101511673ffffffffffffffffffffffffffffffffffffffff60c08501511660e0850151906040519288840194855260408401526060830152608082015260808152610bdb60a0826121e3565b519020610c9d606084015186815191012091610c1e604086015188815191012088610c1e610c4a6101208a01516040519283918583019586526040830190612841565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826121e3565b51902060808801518a81519101209060a0890151926040519788968d88019a60008c5260408901526060880152608087015260a086015260c085015260e0840152610100808401526101208301906127ca565b5190208151527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169060005b8460a08301518051831015610f1a5782610cf791612787565b515110610eca57610d0c8160a0840151612787565b5185815191015190868110610e9b575b506040517feeb7b2480000000000000000000000000000000000000000000000000000000081528160048201528681602481885afa80156108e95773ffffffffffffffffffffffffffffffffffffffff91600091610e6e575b5016908115610e415750906000610dd78193604051610d9a81610c1e898d830161293f565b6040519586809481937f1234eab00000000000000000000000000000000000000000000000000000000083526040600484015260448301906122b0565b87602483015203925af19182156108e9578692610dfa575b506001019050610cde565b9091503d806000833e610e0d81836121e3565b81019186828403126101235781519167ffffffffffffffff8311610123578793600193610e3a9201612653565b5090610def565b7f81fa16780000000000000000000000000000000000000000000000000000000060005260045260246000fd5b610e8e9150883d8a11610e94575b610e8681836121e3565b8101906124ef565b88610d75565b503d610e7c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90870360031b1b1686610d1c565b610ed99060a086930151612787565b5190610f166040519283927f313c6972000000000000000000000000000000000000000000000000000000008452600484015260248301906122b0565b0390fd5b50828567ffffffffffffffff60608351015116907f7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee56167ffffffffffffffff60405192169180610f69868261293f565b0390a37fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff600254166002555151604051908152f35b9092503d806000833e610fb181836121e3565b8101828282031261012357815167ffffffffffffffff811161012357610fd79201612698565b9185610aac565b610fe6612717565b82828c01015201859061065b565b6040833603126101235785604091825161100d81612173565b61101686612316565b81528286013583820152815201920191610606565b602090611036612717565b82828b010152016104c7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b97925093503d8083893e61108581896121e3565b870160808882031261110957602088015167ffffffffffffffff811161110557816110b1918a01612653565b92604089015167ffffffffffffffff811161110157826110d2918b01612653565b9860608101519167ffffffffffffffff83116110fe57506110f4929101612698565b91969193886103da565b80fd5b5080fd5b8380fd5b8280fd5b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760006040516111c88161218f565b6111d06122f3565b8152602435801515810361110957602082019081526044359073ffffffffffffffffffffffffffffffffffffffff821682036111055760408301918252611215612bb7565b73ffffffffffffffffffffffffffffffffffffffff8351161580156113e4575b80156113da575b6113b2579173ffffffffffffffffffffffffffffffffffffffff60e0927f33306e6a63669cb4611ac048b14a2f5fc09e420dcb3dddbaedd3700085564f8194828451167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff000000000000000000000000000000000000000060035416176003556113ae611326612337565b91611376604051809473ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b608083019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565ba180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b508051151561123c565b5073ffffffffffffffffffffffffffffffffffffffff82511615611235565b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235767ffffffffffffffff61144361223c565b166000526004602052600167ffffffffffffffff604060002054160167ffffffffffffffff81116110425760209067ffffffffffffffff60405191168152f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760005473ffffffffffffffffffffffffffffffffffffffff81163303611594577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576000604080516115fc8161218f565b828152826020820152015260606040516116158161218f565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152611697604051809273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565bf35b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235767ffffffffffffffff6116d961223c565b1660005260046020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519167ffffffffffffffff81168352831c166020820152f35b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760043567ffffffffffffffff8111610123573660238201121561012357806004013567ffffffffffffffff8111610123573660248260051b840101116101235773ffffffffffffffffffffffffffffffffffffffff6003541660005b82811015611a195760009073ffffffffffffffffffffffffffffffffffffffff6117d460248360051b8801016125e1565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115611a0e5790859185916119da575b508061182f575b50505060019150016117a3565b6118e860405160208101967fa9059cbb000000000000000000000000000000000000000000000000000000008852846024830152836044830152604482526118786064836121e3565b80806040998a519461188a8c876121e3565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208701525190828a5af13d156119d1573d6118ca81612253565b906118d78b5192836121e3565b8152809260203d92013e5b86612c02565b805180611926575b505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381611822565b61193d929495969350602080918301019101612448565b1561194e57929190849088806118f0565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091506118e2565b9150506020813d8211611a06575b816119f5602093836121e3565b81010312611105578490518861181b565b3d91506119e8565b6040513d86823e3d90fd5b005b346101235760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357611a5261223c565b5060243573ffffffffffffffffffffffffffffffffffffffff8116810361012357611a7e60209161251b565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101235760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357611ad361223c565b60243567ffffffffffffffff811161012357806004018136039060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101235767ffffffffffffffff84169377ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156108e957600091611e38575b50611e0a5773ffffffffffffffffffffffffffffffffffffffff6002541692604051947fd8694ccd000000000000000000000000000000000000000000000000000000008652600486015260406024860152611c6a611c2d611c1c8480612460565b60a060448a015260e48901916124b0565b611c3a6024840185612460565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8984030160648a01526124b0565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd604483013591018112156101235781016024600482013591019367ffffffffffffffff8211610123578160061b36038513610123578681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160848801528181528694602090910193929160005b818110611dc357505050611d7360209593611d43869460848573ffffffffffffffffffffffffffffffffffffffff611d3660648a9901612316565b1660a48801520190612460565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8584030160c48601526124b0565b03915afa80156108e957600090611d90575b602090604051908152f35b506020813d602011611dbb575b81611daa602093836121e3565b810103126101235760209051611d85565b3d9150611d9d565b91955091929360408060019273ffffffffffffffffffffffffffffffffffffffff611ded8a612316565b168152602089013560208201520196019101918795949392611cfb565b837ffdbd6a720000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611e51915060203d6020116109fa576109ec81836121e3565b85611bba565b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012357611ed46040805190611e9881836121e3565b601782527f566572696669657250726f787920312e372e302d6465760000000000000000006020830152519182916020835260208301906122b0565b0390f35b346101235760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610123576080611f11612337565b611697604051809273ffffffffffffffffffffffffffffffffffffffff6060809267ffffffffffffffff8151168552826020820151166020860152826040820151166040860152015116910152565b346101235760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101235760043567ffffffffffffffff81116101235736602382011215610123578060040135611fba81612224565b91611fc860405193846121e3565b8183526024602084019260061b8201019036821161012357602401915b81831061210f5783611ff5612bb7565b6000905b8051821015611a195761200c8282612787565b519167ffffffffffffffff6120218284612787565b5151169283156120e157927ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed060406001949583600052600460205273ffffffffffffffffffffffffffffffffffffffff6020836000209201518254927bffffffffffffffffffffffffffffffffffffffff000000000000000082861b167fffffffff0000000000000000000000000000000000000000ffffffffffffffff851617905567ffffffffffffffff845193168352166020820152a20190611ff9565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b604083360312610123576040519061212682612173565b833567ffffffffffffffff8116810361012357825260208401359073ffffffffffffffffffffffffffffffffffffffff821682036101235782602092836040950152815201920191611fe5565b6040810190811067ffffffffffffffff8211176108f557604052565b6060810190811067ffffffffffffffff8211176108f557604052565b6080810190811067ffffffffffffffff8211176108f557604052565b60e0810190811067ffffffffffffffff8211176108f557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108f557604052565b67ffffffffffffffff81116108f55760051b60200190565b6004359067ffffffffffffffff8216820361012357565b67ffffffffffffffff81116108f557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106122a05750506000910152565b8181015183820152602001612290565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936122ec8151809281875287808801910161228d565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361012357565b359073ffffffffffffffffffffffffffffffffffffffff8216820361012357565b60006060604051612347816121ab565b8281528260208201528260408201520152604051612364816121ab565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015290565b90816020910312610123575180151581036101235790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561012357016020813591019167ffffffffffffffff821161012357813603831361012357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90816020910312610123575173ffffffffffffffffffffffffffffffffffffffff811681036101235790565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156108e95773ffffffffffffffffffffffffffffffffffffffff916000916125c457501690565b6125dd915060203d602011610e9457610e8681836121e3565b1690565b3573ffffffffffffffffffffffffffffffffffffffff811681036101235790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610123570180359067ffffffffffffffff82116101235760200191813603831361012357565b81601f8201121561012357805161266981612253565b9261267760405194856121e3565b8184526020828401011161012357612695916020808501910161228d565b90565b9080601f830112156101235781516126af81612224565b926126bd60405194856121e3565b81845260208085019260051b820101918383116101235760208201905b8382106126e957505050505090565b815167ffffffffffffffff81116101235760209161270c87848094880101612653565b8152019101906126da565b60405190612724826121c7565b600060c08382815282602082015260606040820152606080820152826080820152606060a08201520152565b92919261275c82612253565b9161276a60405193846121e3565b829481845281830111610123578281602093846000960137010152565b805182101561279b5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9080602083519182815201916020808360051b8301019401926000915b8383106127f657505050505090565b9091929394602080612832837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516122b0565b970193019301919392906127e7565b9080602083519182815201916020808360051b8301019401926000915b83831061286d57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff825116815273ffffffffffffffffffffffffffffffffffffffff83830151168382015260c08061292a61290e6128fc604087015160e0604088015260e08701906122b0565b606087015186820360608801526122b0565b6080860151608086015260a086015185820360a08701526122b0565b9301519101529701930193019193929061285e565b906020825267ffffffffffffffff60608251805160208601528260208201511660408601528260408201511682860152015116608083015273ffffffffffffffffffffffffffffffffffffffff60208201511660a0830152610140612acb612a5f612a2a6129f56129c160408701516101c060c08a01526101e08901906122b0565b60608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08983030160e08a01526122b0565b60808601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0888303016101008901526122b0565b60a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0878303016101208801526127ca565b73ffffffffffffffffffffffffffffffffffffffff60c0850151168386015260e08401516101608601526101008401516101808601526101208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868303016101a0870152612841565b910151916101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082840301910152815180825260208201916020808360051b8301019401926000915b838310612b2457505050505090565b9091929394602080612ba8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080612b798584015160a08785015260a08401906122b0565b926040810151604084015267ffffffffffffffff606082015116606084015201519060808184039101526122b0565b97019301930191939290612b15565b73ffffffffffffffffffffffffffffffffffffffff600154163303612bd857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91929015612c7d5750815115612c16575090565b3b15612c1f5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015612c905750805190602001fd5b610f16906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906122b056fea164736f6c634300081a000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig VerifierProxyStaticConfig, dynamicConfig VerifierProxyDynamicConfig, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{address: address, abi: *parsed, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_VerifierProxy *VerifierProxySession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(VerifierProxyDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyDynamicConfig)).(*VerifierProxyDynamicConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCaller) GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(VerifierProxyStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyStaticConfig)).(*VerifierProxyStaticConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxySession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxySession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_VerifierProxy *VerifierProxySession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_VerifierProxy *VerifierProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

func (_VerifierProxy *VerifierProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

type VerifierProxyCCIPMessageSentIterator struct {
	Event *VerifierProxyCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyCCIPMessageSent)
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
		it.Event = new(VerifierProxyCCIPMessageSent)
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

func (it *VerifierProxyCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyCommitVerifierMessage
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCCIPMessageSentIterator{contract: _VerifierProxy.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyCCIPMessageSent)
				if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error) {
	event := new(VerifierProxyCCIPMessageSent)
	if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyConfigSetIterator struct {
	Event *VerifierProxyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyConfigSet)
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
		it.Event = new(VerifierProxyConfigSet)
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

func (it *VerifierProxyConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyConfigSet struct {
	StaticConfig  VerifierProxyStaticConfig
	DynamicConfig VerifierProxyDynamicConfig
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyConfigSetIterator{contract: _VerifierProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error) {
	event := new(VerifierProxyConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyDestChainConfigSetIterator struct {
	Event *VerifierProxyDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyDestChainConfigSet)
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
		it.Event = new(VerifierProxyDestChainConfigSet)
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

func (it *VerifierProxyDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyDestChainConfigSetIterator{contract: _VerifierProxy.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyDestChainConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error) {
	event := new(VerifierProxyDestChainConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyFeeTokenWithdrawnIterator struct {
	Event *VerifierProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyFeeTokenWithdrawn)
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
		it.Event = new(VerifierProxyFeeTokenWithdrawn)
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

func (it *VerifierProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFeeTokenWithdrawnIterator{contract: _VerifierProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyFeeTokenWithdrawn)
				if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error) {
	event := new(VerifierProxyFeeTokenWithdrawn)
	if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
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
		it.Event = new(VerifierProxyOwnershipTransferRequested)
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

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
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
		it.Event = new(VerifierProxyOwnershipTransferred)
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

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	SequenceNumber uint64
	Router         common.Address
}

func (_VerifierProxy *VerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierProxy.abi.Events["CCIPMessageSent"].ID:
		return _VerifierProxy.ParseCCIPMessageSent(log)
	case _VerifierProxy.abi.Events["ConfigSet"].ID:
		return _VerifierProxy.ParseConfigSet(log)
	case _VerifierProxy.abi.Events["DestChainConfigSet"].ID:
		return _VerifierProxy.ParseDestChainConfigSet(log)
	case _VerifierProxy.abi.Events["FeeTokenWithdrawn"].ID:
		return _VerifierProxy.ParseFeeTokenWithdrawn(log)
	case _VerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierProxy.ParseOwnershipTransferRequested(log)
	case _VerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _VerifierProxy.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierProxyCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee561")
}

func (VerifierProxyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x33306e6a63669cb4611ac048b14a2f5fc09e420dcb3dddbaedd3700085564f81")
}

func (VerifierProxyDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed0")
}

func (VerifierProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
