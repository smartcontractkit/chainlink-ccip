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
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

type CommitVerifierOnRampStaticConfig struct {
	RmnRemote    common.Address
	NonceManager common.Address
}

var CommitVerifierOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowlistUpdates\",\"inputs\":[{\"name\":\"allowlistConfigArgsItems\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.AllowlistConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedAllowlistedSenders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCommitVerifierOnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardToVerifier\",\"inputs\":[{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedSendersList\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configuredAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verifierAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListSendersAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListSendersRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"senders\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitVerifierOnRamp.StaticConfig\",\"components\":[{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCommitVerifierOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAllowListRequest\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByVerifierAggregator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwnerOrAllowlistAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c06040523461043c57612aca8038038061001981610476565b92833981019080820360e0811261043c576040811261043c5761003a610457565b82519091906001600160a01b038116810361043c578252608061005f6020850161049b565b6020840190815291603f19011261043c5760405192608084016001600160401b03811185821017610441576040526100996040820161049b565b84526100a7606082016104af565b92602085019384526100bb6080830161049b565b90604086019182526100cf60a0840161049b565b6060870190815260c084015190936001600160401b03821161043c570187601f8201121561043c578051906001600160401b0382116104415761011760208360051b01610476565b9860206060818c8681520194028301019181831161043c57602001925b8284106103cf575050505033156103be57600180546001600160a01b0319163317905580516001600160a01b03161580156103ac575b61037f57516001600160a01b0390811660808190529351811660a081905286519095911615801561039a575b8015610390575b61037f57855160028054835160ff60a01b90151560a01b166001600160a01b039384166001600160a81b0319909216919091171790558251600380549183166001600160a01b03199283161790558451600480549190931691161790557fb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f3709560c0956020610229610457565b878152019081526040805196875290516001600160a01b039081166020880152915182169086015290511515606085015290518116608084015290511660a0820152a16000905b80518210156103425761028382826104bc565b51916001600160401b0361029782846104bc565b51511692831561032d57600084815260056020908152604091829020818401518154948401516001600160a81b0319909516600882901b610100600160a81b03161794151560ff1694851790915582516001600160a01b03909116815292151590830152929360019390917f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c9190a20190610270565b8363c35aa79d60e01b60005260045260246000fd5b6040516125e390816104e7823960805181818161136401528181611cea01526121d5015260a05181818161139d01528181611a140152611d230152f35b6306b7c75960e31b60005260046000fd5b508051151561019d565b5081516001600160a01b031615610196565b5083516001600160a01b03161561016a565b639b15e16f60e01b60005260046000fd5b60608483031261043c5760405190606082016001600160401b03811183821017610441576040528451906001600160401b038216820361043c57826020926060945261041c83880161049b565b8382015261042c604088016104af565b6040820152815201930192610134565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60408051919082016001600160401b0381118382101761044157604052565b6040519190601f01601f191682016001600160401b0381118382101761044157604052565b51906001600160a01b038216820361043c57565b5190811515820361043c57565b80518210156104d05760209160051b010190565b634e487b7160e01b600052603260045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816306285c6914611c84575080631234eab01461150b578063181f5a771461148e5780631de9a3ca1461115957806320487ded14610e745780632716072b14610c025780635cb80c5d1461090f5780636def4ce7146108945780637437ff9f1461078857806379ba50971461069f5780638da5cb5b1461064d578063972b46121461057c578063c9b146b3146101ae5763f2fde38b146100b957600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95773ffffffffffffffffffffffffffffffffffffffff610105611e8f565b61010d612280565b1633811461017f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a9576101fd903690600401611f24565b73ffffffffffffffffffffffffffffffffffffffff600154163303610532575b906000917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81823603015b81841015610530576000938060051b8401358281121561052c5784019160808336031261052c576040519461027b86611df8565b61028484611ef7565b865261029260208501611ed3565b9660208701978852604085013567ffffffffffffffff8111610528576102bb903690870161210a565b9460408801958652606081013567ffffffffffffffff8111610524576102e39136910161210a565b946060880195865267ffffffffffffffff88511682526005602052604082209851151561033b818b9060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b8151516103f8575b50959760010195505b8451805182101561038b579061038473ffffffffffffffffffffffffffffffffffffffff61037c83600195611fe6565b511688612374565b500161034c565b50509594909350600192519081516103a9575b505001929190610247565b6103ee67ffffffffffffffff7fc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d42158692511692604051918291602083526020830190611f55565b0390a2858061039e565b98939592909497989691966000146104ed57600184019591875b8651805182101561048f5761043c8273ffffffffffffffffffffffffffffffffffffffff92611fe6565b5116801561045857906104516001928a6122e3565b5001610412565b60248a67ffffffffffffffff8e51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b50509690929550600191939897947f330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc32816104e367ffffffffffffffff8b51169251604051918291602083526020830190611f55565b0390a29089610343565b60248767ffffffffffffffff8b51167f463258ff000000000000000000000000000000000000000000000000000000008252600452fd5b8280fd5b5080fd5b8580fd5b005b73ffffffffffffffffffffffffffffffffffffffff6004541633031561021d577f905d7d9b0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95767ffffffffffffffff6105bc611ee0565b1680600052600560205260ff604060002054169060005260056020526001604060002001906040518083602082955493848152019060005260206000209260005b81811061063457505061061292500383611e14565b61063060405192839215158352604060208401526040830190611f55565b0390f35b84548352600194850194879450602090930192016105fd565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760005473ffffffffffffffffffffffffffffffffffffffff8116330361075e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957600060606040516107c781611df8565b828152826020820152826040820152015260806040516107e681611df8565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015273ffffffffffffffffffffffffffffffffffffffff600454166060820152610892604051809273ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565bf35b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95767ffffffffffffffff6108d4611ee0565b1660005260056020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519160ff81161515835260081c166020820152f35b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a95761095e903690600401611f24565b73ffffffffffffffffffffffffffffffffffffffff6003541660005b82811015610530576000908060051b85013573ffffffffffffffffffffffffffffffffffffffff8116809103610524576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115610bf7579085918591610bbf575b5080610a04575b505050600191500161097a565b60405194610abe60208701967fa9059cbb00000000000000000000000000000000000000000000000000000000885284602482015283604482015260448152610a4e606482611e14565b82806040998a5193610a608c86611e14565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460208601525190828a5af13d15610bb7573d90610aa182611e55565b91610aae8b519384611e14565b82523d85602084013e5b87612506565b90815180610afe575b50505060207f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9160019651908152a38583816109f7565b82939596979450916020919281010312610bb457506020610b1f9101612029565b15610b31579291908490888080610ac7565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b606090610ab8565b9150506020813d8211610bef575b81610bda60209383611e14565b81010312610beb57849051886109f0565b8380fd5b3d9150610bcd565b6040513d86823e3d90fd5b346101a95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957366023820112156101a9578060040135610c5c81611f0c565b91610c6a6040519384611e14565b81835260246060602085019302820101903682116101a957602401915b818310610de05783610c97612280565b6000905b805182101561053057610cae8282611fe6565b519167ffffffffffffffff610cc38284611fe6565b515116928315610db257927f57bf8e83dfad9b024aa6d338f551b28f7496a0eef9fac94c960d594605c3211c60406001949583600052600560205260ff82600020610d868460208501519483547fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff008860081b16911617845501511515829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b541673ffffffffffffffffffffffffffffffffffffffff83519216825215156020820152a20190610c9b565b837fc35aa79d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6060833603126101a95760405190606082019082821067ffffffffffffffff831117610e4557606092602092604052610e1886611ef7565b8152610e25838701611eb2565b83820152610e3560408701611ed3565b6040820152815201920191610c87565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957610eab611ee0565b60243567ffffffffffffffff81116101a957806004018136039060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101a957610ef88461216f565b73ffffffffffffffffffffffffffffffffffffffff600254169267ffffffffffffffff604051957fd8694ccd00000000000000000000000000000000000000000000000000000000875216600486015260406024860152610fad610f70610f5f848061207b565b60a060448a015260e48901916120cb565b610f7d602484018561207b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8984030160648a01526120cb565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd604483013591018112156101a95781016024600482013591019367ffffffffffffffff82116101a9578160061b360385136101a9578681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0160848801528181528694602090910193929160005b818110611112575050506110b660209593611086869460848573ffffffffffffffffffffffffffffffffffffffff61107960648a9901611eb2565b1660a4880152019061207b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc8584030160c48601526120cb565b03915afa8015611106576000906110d3575b602090604051908152f35b506020813d6020116110fe575b816110ed60209383611e14565b810103126101a957602090516110c8565b3d91506110e0565b6040513d6000823e3d90fd5b91955091929360408060019273ffffffffffffffffffffffffffffffffffffffff61113c8a611eb2565b16815260208901356020820152019601910191879594939261103e565b346101a95760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957600060405161119681611df8565b61119e611e8f565b81526024358015158103610524576020820190815260443573ffffffffffffffffffffffffffffffffffffffff81168103610beb57604083019081526064359173ffffffffffffffffffffffffffffffffffffffff8316830361148a576060840192835261120a612280565b73ffffffffffffffffffffffffffffffffffffffff84511615801561146b575b8015611461575b6114395760c09273ffffffffffffffffffffffffffffffffffffffff85938193827fb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f3709851167fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006002549260a01b1691161760025551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff000000000000000000000000000000000000000060045416176004556114356040519161134d83611ddc565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016835273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208401526113f1604051809473ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b604083019073ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565ba180f35b6004857f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515611231565b5073ffffffffffffffffffffffffffffffffffffffff8251161561122a565b8480fd5b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95761063060408051906114cf8183611e14565b601882527f436f6d6d6974566572696669657220312e372e302d6465760000000000000000602083015251918291602083526020830190611d99565b346101a95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a95760043567ffffffffffffffff81116101a957366023820112156101a957806004013567ffffffffffffffff81116101a95781019060248201903682116101a9576020818403126101a95760248101359067ffffffffffffffff82116101a95701918281039061018082126101a9576040519161012083019083821067ffffffffffffffff831117610e4557608091604052126101a9576040516115dc81611df8565b602485013581526115ef60448601611ef7565b602082015261160060648601611ef7565b604082015261161160848601611ef7565b6060820152825261162460a48501611eb2565b926020830193845260c485013567ffffffffffffffff81116101a95781602461164f92880101611f9f565b604084015260e485013567ffffffffffffffff81116101a95781602461167792880101611f9f565b946060840195865261168c6101048201611eb2565b916080850192835260a0850191610124810135835261014481013560c087015261016481013567ffffffffffffffff81116101a9578101602481019060e0908703126101a9576040519060e0820182811067ffffffffffffffff821117610e45576040526116f981611eb2565b825261170760208201611eb2565b6020830152604081013567ffffffffffffffff81116101a9578461172c918301611f9f565b6040830152606081013567ffffffffffffffff81116101a95784611751918301611f9f565b60608301526080810135608083015260a08101359067ffffffffffffffff82116101a9576117838560c0938301611f9f565b60a0840152013560c082015260e087015261018481013567ffffffffffffffff81116101a957602491010181601f820112156101a9578035946117c586611f0c565b956117d36040519788611e14565b80875260208088019160051b840101928484116101a95760208101915b848310611bbb57505050505050610100840192835261181d67ffffffffffffffff6040865101511661216f565b67ffffffffffffffff6040855101511660005260056020526040600020805473ffffffffffffffffffffffffffffffffffffffff8160081c163303611b925760ff16611b18575b50906000929173ffffffffffffffffffffffffffffffffffffffff600254169161194660a06118c173ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60408b510151169451169451966024359051611fe6565b51015198516119166040519a8b97889687967f3a49bb4900000000000000000000000000000000000000000000000000000000885260048801526024870152604486015260a0606486015260a4850190611d99565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc848303016084850152611d99565b03915afa92831561110657600093611aa3575b506000921561199d575b6106308367ffffffffffffffff6040519116602082015260208152611989604082611e14565b604051918291602083526020830190611d99565b67ffffffffffffffff604073ffffffffffffffffffffffffffffffffffffffff9251015116915116604051917fea458c0c000000000000000000000000000000000000000000000000000000008352600483015260248201526020816044818573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115611a98578291611a50575b50905061063082611963565b90506020813d602011611a90575b81611a6b60209383611e14565b81010312610528575167ffffffffffffffff8116810361052857610630915082611a44565b3d9150611a5e565b6040513d84823e3d90fd5b90923d8082843e611ab48184611e14565b820190608083830312610bb457611acd60208401612029565b92604081015167ffffffffffffffff81116105245783611aee918301612036565b5060608101519167ffffffffffffffff8311610bb45750611b10929101612036565b509183611959565b855173ffffffffffffffffffffffffffffffffffffffff1660009081526002909101602052604090205415611b4d5786611864565b73ffffffffffffffffffffffffffffffffffffffff8551167fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7ed135dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b823567ffffffffffffffff81116101a95782019060c0600483870301126101a9576040519060c0820182811067ffffffffffffffff821117610e4557604052602083013560028110156101a9578252611c1660408401611eb2565b602083015260608301356040830152611c3160808401611ef7565b606083015260a083013563ffffffff811681036101a957608083015260c08301359167ffffffffffffffff83116101a957611c7489602080969581960101611f9f565b60a08201528152019201916117f0565b346101a95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101a957602081611cc1600093611ddc565b828152015260408051611cd381611ddc565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201526108928251809273ffffffffffffffffffffffffffffffffffffffff60208092828151168552015116910152565b60005b838110611d895750506000910152565b8181015183820152602001611d79565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611dd581518092818752878088019101611d76565b0116010190565b6040810190811067ffffffffffffffff821117610e4557604052565b6080810190811067ffffffffffffffff821117610e4557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610e4557604052565b67ffffffffffffffff8111610e4557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b359073ffffffffffffffffffffffffffffffffffffffff821682036101a957565b359081151582036101a957565b6004359067ffffffffffffffff821682036101a957565b359067ffffffffffffffff821682036101a957565b67ffffffffffffffff8111610e455760051b60200190565b9181601f840112156101a95782359167ffffffffffffffff83116101a9576020808501948460051b0101116101a957565b906020808351928381520192019060005b818110611f735750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611f66565b81601f820112156101a957803590611fb682611e55565b92611fc46040519485611e14565b828452602083830101116101a957816000926020809301838601378301015290565b8051821015611ffa5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b519081151582036101a957565b81601f820112156101a957805161204c81611e55565b9261205a6040519485611e14565b818452602082840101116101a9576120789160208085019101611d76565b90565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101a957016020813591019167ffffffffffffffff82116101a95781360383136101a957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9080601f830112156101a957813561212181611f0c565b9261212f6040519485611e14565b81845260208085019260051b8201019283116101a957602001905b8282106121575750505090565b6020809161216484611eb2565b81520191019061214a565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561110657600091612246575b5061220e5750565b67ffffffffffffffff907ffdbd6a72000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b90506020813d602011612278575b8161226160209383611e14565b810103126101a95761227290612029565b38612206565b3d9150612254565b73ffffffffffffffffffffffffffffffffffffffff6001541633036122a157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015611ffa5760005260206000200190600090565b600082815260018201602052604090205461236d5780549068010000000000000000821015610e4557826123566123218460018096018555846122cb565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b90600182019181600052826020526040600020548015156000146124fd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116124ce578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116124ce57818103612497575b50505080548015612468577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061242982826122cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6124b76124a761232193866122cb565b90549060031b1c928392866122cb565b9055600052836020526040600020553880806123f1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50505050600090565b91929015612581575081511561251a575090565b3b156125235790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156125945750805190602001fd5b6125d2906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190611d99565b0390fdfea164736f6c634300081a000a",
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
	case _CommitVerifierOnRamp.abi.Events["AllowListSendersAdded"].ID:
		return _CommitVerifierOnRamp.ParseAllowListSendersAdded(log)
	case _CommitVerifierOnRamp.abi.Events["AllowListSendersRemoved"].ID:
		return _CommitVerifierOnRamp.ParseAllowListSendersRemoved(log)
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

func (CommitVerifierOnRampAllowListSendersAdded) Topic() common.Hash {
	return common.HexToHash("0x330939f6eafe8bb516716892fe962ff19770570838686e6579dbc1cc51fc3281")
}

func (CommitVerifierOnRampAllowListSendersRemoved) Topic() common.Hash {
	return common.HexToHash("0xc237ec1921f855ccd5e9a5af9733f2d58943a5a8501ec5988e305d7a4d421586")
}

func (CommitVerifierOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb4dd79c1c2d7bf2a53542e064fc5cee59a3c2f854a0d1aa4eacf7b668dd7f370")
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

	FilterAllowListSendersAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersAddedIterator, error)

	WatchAllowListSendersAdded(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersAdded(log types.Log) (*CommitVerifierOnRampAllowListSendersAdded, error)

	FilterAllowListSendersRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*CommitVerifierOnRampAllowListSendersRemovedIterator, error)

	WatchAllowListSendersRemoved(opts *bind.WatchOpts, sink chan<- *CommitVerifierOnRampAllowListSendersRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseAllowListSendersRemoved(log types.Log) (*CommitVerifierOnRampAllowListSendersRemoved, error)

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
