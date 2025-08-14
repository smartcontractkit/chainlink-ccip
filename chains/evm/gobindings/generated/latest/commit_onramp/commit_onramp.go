// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commit_onramp

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

type CommitOnRampAllowlistConfigArgs struct {
	DestChainSelector         uint64
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []common.Address
	RemovedAllowlistedSenders []common.Address
}

type CommitOnRampDestChainConfigArgs struct {
	DestChainSelector  uint64
	VerifierAggregator common.Address
	AllowlistEnabled   bool
}

type CommitOnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

type CommitOnRampStaticConfig struct {
	RmnRemote    common.Address
	NonceManager common.Address
}

var CommitOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structCommitOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByVerifierAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c06040523461043c57612cb98038038061001981610476565b92833981019080820360e0811261043c576040811261043c5761003a610457565b82519091906001600160a01b038116810361043c578252608061005f6020850161049b565b6020840190815291603f19011261043c5760405192608084016001600160401b03811185821017610441576040526100996040820161049b565b84526100a7606082016104af565b92602085019384526100bb6080830161049b565b90604086019182526100cf60a0840161049b565b6060870190815260c084015190936001600160401b03821161043c570187601f8201121561043c578051906001600160401b0382116104415761011760208360051b01610476565b9860206060818c8681520194028301019181831161043c57602001925b8284106103cf575050505033156103be57600180546001600160a01b0319163317905580516001600160a01b03161580156103ac575b61037f57516001600160a01b0390811660808190529351811660a081905286519095911615801561039a575b8015610390575b61037f57855160028054835160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558251600380549183166001600160a01b03199283161790558451600480549190931691161790557fb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f3709560c0956020610229610457565b878152019081526040805196875290516001600160a01b039081166020880152915182169086015290511515606085015290518116608084015290511660a0820152a16000905b80518210156103425761028382826104bc565b51916001600160401b0361029782846104bc565b51511692831561032d57600084815260056020908152604091829020818401518154948401516001600160a81b0319909516600882901b610100600160a81b03161794151560ff1694851790915582516001600160a01b03909116815292151590830152929360019390917f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9190a20190610270565b8363c35aa79d60e01b60005260045260246000fd5b6040516127d290816104e7823960805181818161155201528181611e3e01526123c4015260a05181818161158b01528181611c030152611e770152f35b6306b7c75960e31b60005260046000fd5b508051151561019d565b5081516001600160a01b031615610196565b5083516001600160a01b03161561016a565b639b15e16f60e01b60005260046000fd5b60608483031261043c5760405190606082016001600160401b03811183821017610441576040528451906001600160401b038216820361043c57826020926060945261041c83880161049b565b8382015261042c604088016104af565b6040820152815201930192610134565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60408051919082016001600160401b0381118382101761044157604052565b6040519190601f01601f191682016001600160401b0381118382101761044157604052565b51906001600160a01b038216820361043c57565b5190811515820361043c57565b80518210156104d05760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816306285c6914611dd8575080631234eab0146116f9578063181f5a771461167c5780631de9a3ca1461134757806320487ded146110625780632716072b14610df05780634a7597b514610c0d5780635cb80c5d1461091a5780636def4ce71461089f5780637437ff9f1461079357806379ba5097146106aa5780638da5cb5b14610658578063972b461214610587578063c9b146b3146101b95763f2fde38b146100c457600080fd5b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45773ffffffffffffffffffffffffffffffffffffffff610110611fff565b61011861246f565b1633811461018a57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45760043567ffffffffffffffff81116101b4576102089036906004016120db565b73ffffffffffffffffffffffffffffffffffffffff60015416330361053d575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b8184101561053b576000938060051b8401358281121561053757840191608083360312610537576040519461028686611f4c565b61028f84612067565b865261029d60208501612043565b9660208701978852604085013567ffffffffffffffff8111610533576102c690369087016122f9565b9460408801958652606081013567ffffffffffffffff811161052f576102ee913691016122f9565b946060880195865267ffffffffffffffff885116825260056020526040822098511515610346818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b815151610403575b50959760010195505b84518051821015610396579061038f73ffffffffffffffffffffffffffffffffffffffff610387836001956121d5565b511688612563565b5001610357565b50509594909350600192519081516103b4575b505001929190610252565b6103f967ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d4215869251169260405191829160208352602083019061210c565b0390a285806103a9565b98939592909497989691966000146104f857600184019591875b8651805182101561049a576104478273ffffffffffffffffffffffffffffffffffffffff926121d5565b51168015610463579061045c6001928a6124d2565b500161041d565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816104ee67ffffffffffffffff8b5116925160405191829160208352602083019061210c565b0390a2908961034e565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff60045416330315610228577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45767ffffffffffffffff6105c7612050565b1680600052600560205260ff604060002054169060005260056020526001604060002001906040518083602082955493848152019060005260206000209260005b81811061063f57505061061d92500383611f84565b61063b6040519283921515835260406020840152604083019061210c565b0390f35b8454835260019485019487945060209093019201610608565b346101b45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101b45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45760005473ffffffffffffffffffffffffffffffffffffffff81163303610769577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457600060606040516107d281611f4c565b828152826020820152826040820152015260806040516107f181611f4c565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015273ffffffffffffffffffffffffffffffffffffffff60045416606082015261089d604051809273ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565bf35b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45767ffffffffffffffff6108df612050565b1660005260056020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519160ff81161515835260081c166020820152f35b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45760043567ffffffffffffffff81116101b4576109699036906004016120db565b73ffffffffffffffffffffffffffffffffffffffff6003541660005b8281101561053b576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff811680910361052f576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115610c02579085918591610bca575b5080610a0f575b5050506001915001610985565b60405194610ac960208701967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602482015283604482015260448152610a59606482611f84565b82806040998a5193610a6b8c86611f84565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d15610bc2573d90610aac82611fc5565b91610ab98b519384611f84565b82523d85602084013e5b876126f5565b90815180610b09575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a3858381610a02565b82939596979450916020919281010312610bbf57506020610b2a9101612218565b15610b3c579291908490888080610ad2565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b606090610ac3565b9150506020813d8211610bfa575b81610be560209383611f84565b81010312610bf657849051886109fb565b8380fd5b3d9150610bd8565b6040513d86823e3d90fd5b346101b45760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457610c44612050565b60243567ffffffffffffffff81116101b45760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b45760405190610c8f82611f68565b806004013567ffffffffffffffff81116101b457610cb39060043691840101612094565b8252602481013567ffffffffffffffff81116101b457610cd99060043691840101612094565b6020830152604481013567ffffffffffffffff81116101b4578101366023820112156101b4576004810135610d0d8161207c565b91610d1b6040519384611f84565b818352602060048185019360061b83010101903682116101b457602401915b818310610db8575050506040830152610d5560648201612022565b6060830152608481013567ffffffffffffffff81116101b4576080916004610d809236920101612094565b9101526044359067ffffffffffffffff82116101b457610da7610dad923690600401612094565b5061235e565b602060405160008152f35b6040833603126101b45760206040918251610dd281611f30565b610ddb86612022565b81528286013583820152815201920191610d3a565b346101b45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45760043567ffffffffffffffff81116101b457366023820112156101b4578060040135610e4a8161207c565b91610e586040519384611f84565b81835260246060602085019302820101903682116101b457602401915b818310610fce5783610e8561246f565b6000905b805182101561053b57610e9c82826121d5565b519167ffffffffffffffff610eb182846121d5565b515116928315610fa057927f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c60406001949583600052600560205260ff82600020610f748460208501519483547fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff008860081b16911617845501511515829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a20190610e89565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6060833603126101b45760405190606082019082821067ffffffffffffffff8311176110335760609260209260405261100686612067565b8152611013838701612022565b8382015261102360408701612043565b6040820152815201920191610e75565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101b45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457611099612050565b60243567ffffffffffffffff81116101b457806004018136039060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101b4576110e68461235e565b73ffffffffffffffffffffffffffffffffffffffff600254169267ffffffffffffffff604051957fd8694ccd0000000000000000000000000000000000000000000000000000000087521660048601526040602486015261119b61115e61114d848061226a565b60a060448a015260e48901916122ba565b61116b602484018561226a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8984030160648a01526122ba565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd604483013591018112156101b45781016024600482013591019367ffffffffffffffff82116101b4578160061b360385136101b4578681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160848801528181528694602090910193929160005b818110611300575050506112a460209593611274869460848573ffffffffffffffffffffffffffffffffffffffff61126760648a9901612022565b1660a4880152019061226a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8584030160c48601526122ba565b03915afa80156112f4576000906112c1575b602090604051908152f35b506020813d6020116112ec575b816112db60209383611f84565b810103126101b457602090516112b6565b3d91506112ce565b6040513d6000823e3d90fd5b91955091929360408060019273ffffffffffffffffffffffffffffffffffffffff61132a8a612022565b16815260208901356020820152019601910191879594939261122c565b346101b45760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457600060405161138481611f4c565b61138c611fff565b8152602435801515810361052f576020820190815260443573ffffffffffffffffffffffffffffffffffffffff81168103610bf657604083019081526064359173ffffffffffffffffffffffffffffffffffffffff8316830361167857606084019283526113f861246f565b73ffffffffffffffffffffffffffffffffffffffff845116158015611659575b801561164f575b6116275760c09273ffffffffffffffffffffffffffffffffffffffff85938193827fb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f3709851167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff000000000000000000000000000000000000000060045416176004556116236040519161153b83611f30565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208401526115df604051809473ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b604083019073ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565ba180f35b6004857f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b508051151561141f565b5073ffffffffffffffffffffffffffffffffffffffff82511615611418565b8480fd5b346101b45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45761063b60408051906116bd8183611f84565b601882527f436f6d6d6974566572696669657220312e372e302d6465760000000000000000602083015251918291602083526020830190611eed565b346101b45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b45760043567ffffffffffffffff81116101b457366023820112156101b457806004013567ffffffffffffffff81116101b45781019060248201903682116101b4576020818403126101b45760248101359067ffffffffffffffff82116101b4570190818303906101c082126101b4576040519161016083019083821067ffffffffffffffff83111761103357608091604052126101b4576040516117ca81611f4c565b602484013581526117dd60448501612067565b60208201526117ee60648501612067565b60408201526117ff60848501612067565b6060820152825261181260a48401612022565b926020830193845260c481013567ffffffffffffffff81116101b45782602461183d92840101612094565b604084015260e481013567ffffffffffffffff81116101b45782602461186592840101612094565b946060840195865261187a6101048301612022565b916080850192835260a0850191610124820135835261014482013560c087015261016482013567ffffffffffffffff81116101b457820160248101916080919003126101b457604051906118cd82611f4c565b6118d681612022565b8252602081013567ffffffffffffffff81116101b457866118f8918301612094565b602083015260408101359067ffffffffffffffff82116101b457611920876060938301612094565b60408401520135606082015260e086015261018481013567ffffffffffffffff81116101b4576024908201019380601f860112156101b45784356119638161207c565b956119716040519788611f84565b81875260208088019260051b820101918383116101b45760208201905b838210611daa575050505061010086019485526101a482013567ffffffffffffffff81116101b4578160246119c592850101612156565b6101208701526101c48201359167ffffffffffffffff83116101b4576119ee9201602401612156565b610140850152611a0c67ffffffffffffffff6040865101511661235e565b67ffffffffffffffff6040855101511660005260056020526040600020805473ffffffffffffffffffffffffffffffffffffffff8160081c163303611d815760ff16611d07575b50906000929173ffffffffffffffffffffffffffffffffffffffff6002541691611b356080611ab073ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60408b5101511694511694519660243590516121d5565b5101519851611b056040519a8b97889687967f3a49bb4900000000000000000000000000000000000000000000000000000000885260048801526024870152604486015260a0606486015260a4850190611eed565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016084850152611eed565b03915afa9283156112f457600093611c92575b5060009215611b8c575b61063b8367ffffffffffffffff6040519116602082015260208152611b78604082611f84565b604051918291602083526020830190611eed565b67ffffffffffffffff604073ffffffffffffffffffffffffffffffffffffffff9251015116915116604051917fea458c0c000000000000000000000000000000000000000000000000000000008352600483015260248201526020816044818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611c87578291611c3f575b50905061063b82611b52565b90506020813d602011611c7f575b81611c5a60209383611f84565b81010312610533575167ffffffffffffffff811681036105335761063b915082611c33565b3d9150611c4d565b6040513d84823e3d90fd5b90923d8082843e611ca38184611f84565b820190608083830312610bbf57611cbc60208401612218565b92604081015167ffffffffffffffff811161052f5783611cdd918301612225565b5060608101519167ffffffffffffffff8311610bbf5750611cff929101612225565b509183611b48565b855173ffffffffffffffffffffffffffffffffffffffff1660009081526002909101602052604090205415611d3c5786611a53565b73ffffffffffffffffffffffffffffffffffffffff8551167fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7ed135dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b813567ffffffffffffffff81116101b457602091611dcd87848094880101612156565b81520191019061198e565b346101b45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b457602081611e15600093611f30565b828152015260408051611e2781611f30565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015261089d8251809273ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b60005b838110611edd5750506000910152565b8181015183820152602001611ecd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611f2981518092818752878088019101611eca565b0116010190565b6040810190811067ffffffffffffffff82111761103357604052565b6080810190811067ffffffffffffffff82111761103357604052565b60a0810190811067ffffffffffffffff82111761103357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761103357604052565b67ffffffffffffffff811161103357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101b457565b359073ffffffffffffffffffffffffffffffffffffffff821682036101b457565b359081151582036101b457565b6004359067ffffffffffffffff821682036101b457565b359067ffffffffffffffff821682036101b457565b67ffffffffffffffff81116110335760051b60200190565b81601f820112156101b4578035906120ab82611fc5565b926120b96040519485611f84565b828452602083830101116101b457816000926020809301838601378301015290565b9181601f840112156101b45782359167ffffffffffffffff83116101b4576020808501948460051b0101116101b457565b906020808351928381520192019060005b81811061212a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161211d565b919060a0838203126101b4576040519061216f82611f68565b819361217a81612022565b835261218860208201612067565b6020840152604081013563ffffffff811681036101b45760408401526060810135606084015260808101359167ffffffffffffffff83116101b4576080926121d09201612094565b910152565b80518210156121e95760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b519081151582036101b457565b81601f820112156101b457805161223b81611fc5565b926122496040519485611f84565b818452602082840101116101b4576122679160208085019101611eca565b90565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101b457016020813591019167ffffffffffffffff82116101b45781360383136101b457565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9080601f830112156101b45781356123108161207c565b9261231e6040519485611f84565b81845260208085019260051b8201019283116101b457602001905b8282106123465750505090565b6020809161235384612022565b815201910190612339565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112f457600091612435575b506123fd5750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b90506020813d602011612467575b8161245060209383611f84565b810103126101b45761246190612218565b386123f5565b3d9150612443565b73ffffffffffffffffffffffffffffffffffffffff60015416330361249057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156121e95760005260206000200190600090565b600082815260018201602052604090205461255c578054906801000000000000000082101561103357826125456125108460018096018555846124ba565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b90600182019181600052826020526040600020548015156000146126ec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116126bd578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116126bd57818103612686575b50505080548015612657577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061261882826124ba565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6126a661269661251093866124ba565b90549060031b1c928392866124ba565b9055600052836020526040600020553880806125e0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b919290156127705750815115612709575090565b3b156127125790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156127835750805190602001fd5b6127c1906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611eed565b0390fdfea164736f6c634300081a000a",
}

var CommitOnRampABI = CommitOnRampMetaData.ABI

var CommitOnRampBin = CommitOnRampMetaData.Bin

func DeployCommitOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CommitOnRampStaticConfig, dynamicConfig CommitOnRampDynamicConfig, destChainConfigArgs []CommitOnRampDestChainConfigArgs) (common.Address, *types.Transaction, *CommitOnRamp, error) {
	parsed, err := CommitOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitOnRamp{address: address, abi: *parsed, CommitOnRampCaller: CommitOnRampCaller{contract: contract}, CommitOnRampTransactor: CommitOnRampTransactor{contract: contract}, CommitOnRampFilterer: CommitOnRampFilterer{contract: contract}}, nil
}

type CommitOnRamp struct {
	address common.Address
	abi     abi.ABI
	CommitOnRampCaller
	CommitOnRampTransactor
	CommitOnRampFilterer
}

type CommitOnRampCaller struct {
	contract *bind.BoundContract
}

type CommitOnRampTransactor struct {
	contract *bind.BoundContract
}

type CommitOnRampFilterer struct {
	contract *bind.BoundContract
}

type CommitOnRampSession struct {
	Contract     *CommitOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CommitOnRampCallerSession struct {
	Contract *CommitOnRampCaller
	CallOpts bind.CallOpts
}

type CommitOnRampTransactorSession struct {
	Contract     *CommitOnRampTransactor
	TransactOpts bind.TransactOpts
}

type CommitOnRampRaw struct {
	Contract *CommitOnRamp
}

type CommitOnRampCallerRaw struct {
	Contract *CommitOnRampCaller
}

type CommitOnRampTransactorRaw struct {
	Contract *CommitOnRampTransactor
}

func NewCommitOnRamp(address common.Address, backend bind.ContractBackend) (*CommitOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(CommitOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCommitOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitOnRamp{address: address, abi: abi, CommitOnRampCaller: CommitOnRampCaller{contract: contract}, CommitOnRampTransactor: CommitOnRampTransactor{contract: contract}, CommitOnRampFilterer: CommitOnRampFilterer{contract: contract}}, nil
}

func NewCommitOnRampCaller(address common.Address, caller bind.ContractCaller) (*CommitOnRampCaller, error) {
	contract, err := bindCommitOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampCaller{contract: contract}, nil
}

func NewCommitOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitOnRampTransactor, error) {
	contract, err := bindCommitOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampTransactor{contract: contract}, nil
}

func NewCommitOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitOnRampFilterer, error) {
	contract, err := bindCommitOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampFilterer{contract: contract}, nil
}

func bindCommitOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CommitOnRamp *CommitOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOnRamp.Contract.CommitOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_CommitOnRamp *CommitOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.CommitOnRampTransactor.contract.Transfer(opts)
}

func (_CommitOnRamp *CommitOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.CommitOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_CommitOnRamp *CommitOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_CommitOnRamp *CommitOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.contract.Transfer(opts)
}

func (_CommitOnRamp *CommitOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_CommitOnRamp *CommitOnRampCaller) GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

	error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getAllowedSendersList", destChainSelector)

	outstruct := new(GetAllowedSendersList)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfiguredAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_CommitOnRamp *CommitOnRampSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _CommitOnRamp.Contract.GetAllowedSendersList(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetAllowedSendersList(destChainSelector uint64) (GetAllowedSendersList,

	error) {
	return _CommitOnRamp.Contract.GetAllowedSendersList(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistEnabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.VerifierAggregator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CommitOnRamp *CommitOnRampSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitOnRamp.Contract.GetDestChainConfig(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CommitOnRamp.Contract.GetDestChainConfig(&_CommitOnRamp.CallOpts, destChainSelector)
}

func (_CommitOnRamp *CommitOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (CommitOnRampDynamicConfig, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CommitOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitOnRampDynamicConfig)).(*CommitOnRampDynamicConfig)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetDynamicConfig() (CommitOnRampDynamicConfig, error) {
	return _CommitOnRamp.Contract.GetDynamicConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetDynamicConfig() (CommitOnRampDynamicConfig, error) {
	return _CommitOnRamp.Contract.GetDynamicConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee(&_CommitOnRamp.CallOpts, destChainSelector, message)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee(&_CommitOnRamp.CallOpts, destChainSelector, message)
}

func (_CommitOnRamp *CommitOnRampCaller) GetFee0(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getFee0", destChainSelector, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetFee0(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee0(&_CommitOnRamp.CallOpts, destChainSelector, arg1, arg2)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetFee0(destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error) {
	return _CommitOnRamp.Contract.GetFee0(&_CommitOnRamp.CallOpts, destChainSelector, arg1, arg2)
}

func (_CommitOnRamp *CommitOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (CommitOnRampStaticConfig, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CommitOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitOnRampStaticConfig)).(*CommitOnRampStaticConfig)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) GetStaticConfig() (CommitOnRampStaticConfig, error) {
	return _CommitOnRamp.Contract.GetStaticConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) GetStaticConfig() (CommitOnRampStaticConfig, error) {
	return _CommitOnRamp.Contract.GetStaticConfig(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) Owner() (common.Address, error) {
	return _CommitOnRamp.Contract.Owner(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) Owner() (common.Address, error) {
	return _CommitOnRamp.Contract.Owner(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CommitOnRamp *CommitOnRampSession) TypeAndVersion() (string, error) {
	return _CommitOnRamp.Contract.TypeAndVersion(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampCallerSession) TypeAndVersion() (string, error) {
	return _CommitOnRamp.Contract.TypeAndVersion(&_CommitOnRamp.CallOpts)
}

func (_CommitOnRamp *CommitOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_CommitOnRamp *CommitOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOnRamp.Contract.AcceptOwnership(&_CommitOnRamp.TransactOpts)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CommitOnRamp.Contract.AcceptOwnership(&_CommitOnRamp.TransactOpts)
}

func (_CommitOnRamp *CommitOnRampTransactor) ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []CommitOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "applyAllowlistUpdates", allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []CommitOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyAllowlistUpdates(&_CommitOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ApplyAllowlistUpdates(allowlistConfigArgsItems []CommitOnRampAllowlistConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyAllowlistUpdates(&_CommitOnRamp.TransactOpts, allowlistConfigArgsItems)
}

func (_CommitOnRamp *CommitOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CommitOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CommitOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CommitOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ApplyDestChainConfigUpdates(&_CommitOnRamp.TransactOpts, destChainConfigArgs)
}

func (_CommitOnRamp *CommitOnRampTransactor) ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "forwardToVerifier", rawMessage, verifierIndex)
}

func (_CommitOnRamp *CommitOnRampSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, rawMessage, verifierIndex)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) ForwardToVerifier(rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.ForwardToVerifier(&_CommitOnRamp.TransactOpts, rawMessage, verifierIndex)
}

func (_CommitOnRamp *CommitOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampSession) SetDynamicConfig(dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.SetDynamicConfig(&_CommitOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) SetDynamicConfig(dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.SetDynamicConfig(&_CommitOnRamp.TransactOpts, dynamicConfig)
}

func (_CommitOnRamp *CommitOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_CommitOnRamp *CommitOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.TransferOwnership(&_CommitOnRamp.TransactOpts, to)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.TransferOwnership(&_CommitOnRamp.TransactOpts, to)
}

func (_CommitOnRamp *CommitOnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CommitOnRamp *CommitOnRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.WithdrawFeeTokens(&_CommitOnRamp.TransactOpts, feeTokens)
}

func (_CommitOnRamp *CommitOnRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CommitOnRamp.Contract.WithdrawFeeTokens(&_CommitOnRamp.TransactOpts, feeTokens)
}

type CommitOnRampAllowListSendersAddedIterator struct {
	Event *CommitOnRampAllowListSendersAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampAllowListSendersAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampAllowListSendersAdded)
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
		it.Event = new(CommitOnRampAllowListSendersAdded)
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

func (it *CommitOnRampAllowListSendersAddedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampAllowListSendersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampAllowListSendersAdded struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampAllowListSendersAddedIterator{contract: _CommitOnRamp.contract, event: "AllowListSendersAdded", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "AllowListSendersAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampAllowListSendersAdded)
				if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseAllowListSendersAdded(log types.Log) (*CommitOnRampAllowListSendersAdded, error) {
	event := new(CommitOnRampAllowListSendersAdded)
	if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampAllowListSendersRemovedIterator struct {
	Event *CommitOnRampAllowListSendersRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampAllowListSendersRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampAllowListSendersRemoved)
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
		it.Event = new(CommitOnRampAllowListSendersRemoved)
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

func (it *CommitOnRampAllowListSendersRemovedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampAllowListSendersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampAllowListSendersRemoved struct {
	DestChainSelector uint64
	Senders           []common.Address
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampAllowListSendersRemovedIterator{contract: _CommitOnRamp.contract, event: "AllowListSendersRemoved", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "AllowListSendersRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampAllowListSendersRemoved)
				if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseAllowListSendersRemoved(log types.Log) (*CommitOnRampAllowListSendersRemoved, error) {
	event := new(CommitOnRampAllowListSendersRemoved)
	if err := _CommitOnRamp.contract.UnpackLog(event, "AllowListSendersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampConfigSetIterator struct {
	Event *CommitOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampConfigSet)
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
		it.Event = new(CommitOnRampConfigSet)
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

func (it *CommitOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampConfigSet struct {
	StaticConfig  CommitOnRampStaticConfig
	DynamicConfig CommitOnRampDynamicConfig
	Raw           types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CommitOnRampConfigSetIterator, error) {

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CommitOnRampConfigSetIterator{contract: _CommitOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampConfigSet)
				if err := _CommitOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseConfigSet(log types.Log) (*CommitOnRampConfigSet, error) {
	event := new(CommitOnRampConfigSet)
	if err := _CommitOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampDestChainConfigSetIterator struct {
	Event *CommitOnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampDestChainConfigSet)
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
		it.Event = new(CommitOnRampDestChainConfigSet)
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

func (it *CommitOnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampDestChainConfigSet struct {
	DestChainSelector uint64
	Router            common.Address
	AllowlistEnabled  bool
	Raw               types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampDestChainConfigSetIterator{contract: _CommitOnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampDestChainConfigSet)
				if err := _CommitOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseDestChainConfigSet(log types.Log) (*CommitOnRampDestChainConfigSet, error) {
	event := new(CommitOnRampDestChainConfigSet)
	if err := _CommitOnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampFeeTokenWithdrawnIterator struct {
	Event *CommitOnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampFeeTokenWithdrawn)
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
		it.Event = new(CommitOnRampFeeTokenWithdrawn)
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

func (it *CommitOnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CommitOnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampFeeTokenWithdrawnIterator{contract: _CommitOnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampFeeTokenWithdrawn)
				if err := _CommitOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CommitOnRampFeeTokenWithdrawn, error) {
	event := new(CommitOnRampFeeTokenWithdrawn)
	if err := _CommitOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampOwnershipTransferRequestedIterator struct {
	Event *CommitOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampOwnershipTransferRequested)
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
		it.Event = new(CommitOnRampOwnershipTransferRequested)
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

func (it *CommitOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampOwnershipTransferRequestedIterator{contract: _CommitOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampOwnershipTransferRequested)
				if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*CommitOnRampOwnershipTransferRequested, error) {
	event := new(CommitOnRampOwnershipTransferRequested)
	if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CommitOnRampOwnershipTransferredIterator struct {
	Event *CommitOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CommitOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitOnRampOwnershipTransferred)
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
		it.Event = new(CommitOnRampOwnershipTransferred)
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

func (it *CommitOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CommitOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CommitOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CommitOnRamp *CommitOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CommitOnRampOwnershipTransferredIterator{contract: _CommitOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CommitOnRamp *CommitOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CommitOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CommitOnRampOwnershipTransferred)
				if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitOnRamp *CommitOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*CommitOnRampOwnershipTransferred, error) {
	event := new(CommitOnRampOwnershipTransferred)
	if err := _CommitOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CommitOnRamp *CommitOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CommitOnRamp.abi.Events["AllowListSendersAdded"].ID:
		return _CommitOnRamp.ParseAllowListSendersAdded(log)
	case _CommitOnRamp.abi.Events["AllowListSendersRemoved"].ID:
		return _CommitOnRamp.ParseAllowListSendersRemoved(log)
	case _CommitOnRamp.abi.Events["ConfigSet"].ID:
		return _CommitOnRamp.ParseConfigSet(log)
	case _CommitOnRamp.abi.Events["DestChainConfigSet"].ID:
		return _CommitOnRamp.ParseDestChainConfigSet(log)
	case _CommitOnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _CommitOnRamp.ParseFeeTokenWithdrawn(log)
	case _CommitOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _CommitOnRamp.ParseOwnershipTransferRequested(log)
	case _CommitOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _CommitOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CommitOnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitOnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f370")
}

func (CommitOnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c")
}

func (CommitOnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CommitOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CommitOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CommitOnRamp *CommitOnRamp) Address() common.Address {
	return _CommitOnRamp.address
}

type CommitOnRampInterface interface {
	GetAllowedSendersList(opts *bind.CallOpts, destChainSelector uint64) (GetAllowedSendersList,

		error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CommitOnRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFee0(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, arg2 []byte) (*big.Int, error)

	GetStaticConfig(opts *bind.CallOpts) (CommitOnRampStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowlistUpdates(opts *bind.TransactOpts, allowlistConfigArgsItems []CommitOnRampAllowlistConfigArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CommitOnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardToVerifier(opts *bind.TransactOpts, rawMessage []byte, verifierIndex *big.Int) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CommitOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitOnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitOnRampAllowListSendersRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CommitOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CommitOnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitOnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CommitOnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CommitOnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CommitOnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CommitOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CommitOnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CommitOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CommitOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CommitOnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
