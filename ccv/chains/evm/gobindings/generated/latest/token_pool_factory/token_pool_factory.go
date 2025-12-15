// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_pool_factory

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

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type TokenPoolFactoryRemoteChainConfig struct {
	RemotePoolFactory   common.Address
	RemoteRouter        common.Address
	RemoteRMNProxy      common.Address
	RemoteLockBox       common.Address
	RemoteTokenDecimals uint8
}

type TokenPoolFactoryRemoteTokenPoolInfo struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	RemotePoolInitCode  []byte
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	PoolType            uint8
	RemoteTokenAddress  []byte
	RemoteTokenInitCode []byte
	RateLimiterConfig   RateLimiterConfig
}

var TokenPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"tokenAdminModule\",\"type\":\"address\",\"internalType\":\"contract RegistryModuleOwnerCustom\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ccipRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenAndTokenPool\",\"inputs\":[{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"tokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTokenPoolWithExistingToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localPoolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenPools\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPoolFactory.RemoteTokenPoolInfo[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"poolType\",\"type\":\"uint8\",\"internalType\":\"enum TokenPoolFactory.PoolType\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteTokenInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"tokenPoolInitCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RemoteChainConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteChainConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct TokenPoolFactory.RemoteChainConfig\",\"components\":[{\"name\":\"remotePoolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteRMNProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteLockBox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidLockBoxChainSelector\",\"inputs\":[{\"name\":\"lockBoxSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockBoxToken\",\"inputs\":[{\"name\":\"lockBoxToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101003461012957601f612cde38819003918201601f19168301916001600160401b0383118484101761012e578084926080946040528339810103126101295780516001600160a01b0381169190828103610129576020820151906001600160a01b03821680830361012957610083606061007c60408701610144565b9501610144565b9415908115610120575b50801561010f575b80156100fe575b6100ed5760a05260c05260805260e052604051612b85908161015982396080518181816114ab0152611b5e015260a05181610424015260c0518161039f015260e0518181816114d30152611b8a0152f35b63f6b2911f60e01b60005260046000fd5b506001600160a01b0384161561009c565b506001600160a01b03831615610095565b9050153861008d565b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101295756fe6080604052600436101561001257600080fd5b6000803560e01c8063181f5a771461068a5780635d31fe08146101dc57639588fa5a1461003e57600080fd5b346101d55760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576004359073ffffffffffffffffffffffffffffffffffffffff82168092036101d55761009761084d565b6044359160028310156101d55760643567ffffffffffffffff81116101d8576100c4903690600401610817565b9290916084359067ffffffffffffffff82116101d55760206101b78873ffffffffffffffffffffffffffffffffffffffff898989896101063660048c016108d2565b9590946101a7610114610900565b60405160c4358d82019081523360601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015291939160ff9161018681603484015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610770565b519020956040519b6101978d610709565b8c52168c8b015260408a01610944565b1660608701526080860152610b45565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b80fd5b5080fd5b50346101d55760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043567ffffffffffffffff81116101d85761022c903690600401610817565b9061023561084d565b9060443560028110156106865760643567ffffffffffffffff81116106825761026290369060040161085d565b916084359267ffffffffffffffff841161067e576101a79561028b6103289536906004016108d2565b93909273ffffffffffffffffffffffffffffffffffffffff6102ab610900565b61030160405160208101906102f78161015a3360c43586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b519020938461097f565b9a8b9860ff846040519b6103148d610709565b169c8d8c521660208b015260408a01610944565b92813b156101d5576040517fc630948d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152818160248183875af180156105f55761066e575b509073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610620578280916024604051809481937f96ea2f7a0000000000000000000000000000000000000000000000000000000083528760048401525af1801561061557908391610659575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610620576040517f156194da000000000000000000000000000000000000000000000000000000008152826004820152838160248183865af1801561063957908491610644575b5050803b15610620576040517f4e847fc700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015286166024820152838160448183865af1801561063957908491610624575b5050803b15610620576040517fddadfa8e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201523360248201529083908290604490829084905af1801561061557908391610600575b5050803b156101d8578180916024604051809481937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af180156105f5576105e0575b50506040805173ffffffffffffffffffffffffffffffffffffffff928316815292909116602083015290f35b0390f35b6105eb828092610770565b6101d557806105b0565b6040513d84823e3d90fd5b8161060a91610770565b6101d8578138610565565b6040513d85823e3d90fd5b8280fd5b8161062e91610770565b6106205782386104f8565b6040513d86823e3d90fd5b8161064e91610770565b61062057823861048e565b8161066391610770565b6101d857813861040b565b8161067891610770565b38610386565b8680fd5b8580fd5b8480fd5b50346101d557807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557506105dc6040516106cb604082610770565b601a81527f546f6b656e506f6f6c466163746f727920312e362e302d64657600000000000060208201526040519182916020835260208301906107d4565b60a0810190811067ffffffffffffffff82111761072557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761072557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761072557604052565b60005b8381106107c45750506000910152565b81810151838201526020016107b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610810815180928187528780880191016107b1565b0116010190565b9181601f840112156108485782359167ffffffffffffffff8311610848576020808501948460051b01011161084857565b600080fd5b6024359060ff8216820361084857565b81601f820112156108485780359067ffffffffffffffff821161072557604051926108b0601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200185610770565b8284526020838301011161084857816000926020809301838601378301015290565b9181601f840112156108485782359167ffffffffffffffff8311610848576020838186019501011161084857565b60a4359073ffffffffffffffffffffffffffffffffffffffff8216820361084857565b359073ffffffffffffffffffffffffffffffffffffffff8216820361084857565b60028210156109505752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908051156109ef576020815191016000f5903d15198215166109e35773ffffffffffffffffffffffffffffffffffffffff8216156109b957565b7fb06ebf3d0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040513d6000823e3d90fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116107255760051b60200190565b60405190610a3e82610754565b60006040838281528260208201520152565b60405190610a5d82610709565b81600081526060602082015260606040820152610a78610a31565b60608201526080610a87610a31565b910152565b35906fffffffffffffffffffffffffffffffff8216820361084857565b805115610ab65760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015610ab65760209160051b010190565b90816020910312610848575173ffffffffffffffffffffffffffffffffffffffff811681036108485790565b90816020910312610848575167ffffffffffffffff811681036108485790565b939093600095610b5486610a19565b94610b626040519687610770565b8686527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610b8f88610a19565b01885b818110611c0157505087946040955b88811015611333578060051b8501357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe418636030181121561132f57608087015181870136036101c0811261132b57895192610100840184811067ffffffffffffffff8211176112fe578b528089013567ffffffffffffffff811681036112f25784526020818a01013567ffffffffffffffff81116112f257610c48903690838c010161085d565b60208501528a818a01013567ffffffffffffffff81116112f257610c71903690838c010161085d565b8b85015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08301126112f6578a518e610cab82610709565b610cb96060848d0101610923565b8252610cc96080848d0101610923565b6020830152610cdc60a0848d0101610923565b8d830152610cee60c0848d0101610923565b60608301525060e0828b01013560ff811681036112fa5760808201526060850152610100818a01013560028110156112f2576080850152610120818a01013567ffffffffffffffff81116112f257610d4b903690838c010161085d565b60a0850152610140818a01013567ffffffffffffffff81116112f2577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea0610d998b606093853692010161085d565b9360c0870194855201126112f6578a5190610db382610754565b898101610160013580151590036112f2576101a081610de4610180828e610160610df1978201013588520101610a8c565b60208501528b0101610a8c565b8b82015260e0840152610e02610a50565b5060608301519073ffffffffffffffffffffffffffffffffffffffff6060830151169160a0850151511561128f575b505060808301516002811015611262576001149081611259575b50611136575b60208201515115610f1f575b508751610e6a8982610770565b600181528b5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08a018110610f0e575090600192916020820151610ead82610aa9565b52610eb781610aa9565b5067ffffffffffffffff8251169160e060a0820151910151918b5193610edc85610709565b845260208401528a8301528060608301526080820152610efc828b610ae5565b52610f07818a610ae5565b5001610ba1565b806060602080938501015201610e70565b888201516060830151610f3f60a085015160208082518301019101610af9565b6080850151906002821015611108578f928d61106e969594600160209573ffffffffffffffffffffffffffffffffffffffff61104996169250146000146110a5576080838101518385015187860151606096870151955173ffffffffffffffffffffffffffffffffffffffff958616818b0190815260ff9094166020850152600060408501529185169683019690965294831691810191909152911660a0820152610fee90829060c00161015a565b8d519283918161100781850197888151938492016107b1565b830161101b825180938580850191016107b1565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610770565b51902073ffffffffffffffffffffffffffffffffffffffff6060850151511691611c1e565b73ffffffffffffffffffffffffffffffffffffffff8951911660208201526020815261109a8982610770565b602082015238610e5d565b6080838101518484015194870151935173ffffffffffffffffffffffffffffffffffffffff93841681890190815260ff929092166020830152600060408301529483166060820152929091169082015261110390829060a00161015a565b610fee565b5060248f7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b6111fa6112328d73ffffffffffffffffffffffffffffffffffffffff61116960a087015160208082518301019101610af9565b166112288d602067ffffffffffffffff8951169373ffffffffffffffffffffffffffffffffffffffff60608b015151169550610f24948351936111ae84880186610770565b86855283850196611c558839805191848301938452818301528082526111d5606083610770565b519889946111eb858701988992519283916107b1565b850191518093858401906107b1565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610770565b8351902084611c1e565b905073ffffffffffffffffffffffffffffffffffffffff6060808501510191169052610e51565b90501538610e4b565b60248e7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b73ffffffffffffffffffffffffffffffffffffffff6112ba9251602081519101209151169084611c1e565b73ffffffffffffffffffffffffffffffffffffffff8b5191166020820152602081526112e68b82610770565b60a08401523880610e31565b8e80fd5b8d80fd5b8f80fd5b60248f7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8c80fd5b8a80fd5b509197959650929150838201516002811015611bd457600103611b2557606082015173ffffffffffffffffffffffffffffffffffffffff168061189b575073ffffffffffffffffffffffffffffffffffffffff82511673ffffffffffffffffffffffffffffffffffffffff61144461140461143e896114328a602060808b015198610f24948351936113c784880186610770565b86855283850196611c558839805191848301938452818301528082526113ee606083610770565b519788946111eb858701988992519283916107b1565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610770565b30908351902086611c1e565b9361097f565b911673ffffffffffffffffffffffffffffffffffffffff821603611873576115699392916020806080611564945b8651838801518b5173ffffffffffffffffffffffffffffffffffffffff92831681870190815260ff9092166020830152600060408301527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f0000000000000000000000000000000000000000000000000000000000000000831660808301529290911660a082015261151090829060c00161015a565b955b015194611536878a51988996858801378501918383018c81528151948592016107b1565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610770565b61097f565b9273ffffffffffffffffffffffffffffffffffffffff84169060209083516115918382610770565b85815285368137833b15610682579091859085519384917fe8a1da170000000000000000000000000000000000000000000000000000000083526044830188600485015285518091528160648501960190855b81811061184f575050507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8386030160248401528151908186528086019581808460051b83010194019686915b8483106116df57505050505081929350038183865af180156116d2576116be575b50803b15610620579082809260248351809581937ff2fde38b0000000000000000000000000000000000000000000000000000000083523360048401525af19081156116b557506116a257505090565b6116ad828092610770565b6101d5575090565b513d84823e3d90fd5b836116cb91949294610770565b9138611652565b50505051903d90823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0919497508093965085929503018352848751918c8c61012083019167ffffffffffffffff86511684528486015192610120868601528351809152610140850190866101408260051b8801019501925b8181106118005750505050839260c0608061177c8895856117ed9660019b015190868303908701526107d4565b946117b9606082015160608601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01519101906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b9801930193018a95929693889592611631565b9295968091945061183c867ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec08a600196990301885288516107d4565b960194019101908f928b9695949261174f565b825167ffffffffffffffff168852968301968b9650889550918301916001016115e4565b6004867ff6b2911f000000000000000000000000000000000000000000000000000000008152fd5b84517f21df0da7000000000000000000000000000000000000000000000000000000008152602081600481855afa908115611a1e578791611b06575b5073ffffffffffffffffffffffffffffffffffffffff80855116911603611a285784517f44faef45000000000000000000000000000000000000000000000000000000008152602081600481855afa908115611a1e579067ffffffffffffffff9188916119ff575b501661195957611569939291602080608061156494611472565b846004602088938351928380927f44faef450000000000000000000000000000000000000000000000000000000082525afa9182156116b55760248467ffffffffffffffff8185876119d0575b507f247245fe00000000000000000000000000000000000000000000000000000000835216600452fd5b6119f2915060203d6020116119f8575b6119ea8183610770565b810190610b25565b846119a6565b503d6119e0565b611a18915060203d6020116119f8576119ea8183610770565b3861193f565b86513d89823e3d90fd5b85600484602088948551938480927f21df0da70000000000000000000000000000000000000000000000000000000082525afa918215611afc57604494508392611ab3575b50517f4101f29a00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff91821660045216602452fd5b73ffffffffffffffffffffffffffffffffffffffff919250611aed829160203d602011611af5575b611ae58183610770565b810190610af9565b929150611a6d565b503d611adb565b84513d85823e3d90fd5b611b1f915060203d602011611af557611ae58183610770565b386118d7565b8151602080840151865173ffffffffffffffffffffffffffffffffffffffff9384168184015260ff9091166040820152600060608201527f000000000000000000000000000000000000000000000000000000000000000083166080808301919091527f000000000000000000000000000000000000000000000000000000000000000090931660a082015261156995949361156493909291829190611bce8160c0810161015a565b95611512565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b602090611c0f999799610a50565b82828b01015201979597610b92565b60559173ffffffffffffffffffffffffffffffffffffffff93600b92604051926040840152602083015281520160ff815320169056fe60c0346100ea57601f610f2438819003918201601f19168301916001600160401b038311848410176100ef5780849260409485528339810103126100ea5780516001600160a01b03811691908290036100ea5760200151906001600160401b03821682036100ea5733156100d957600180546001600160a01b0319163317905580156100c85760805260a052604051610e1e9081610106823960805181818161033e01528181610590015281816107920152610a87015260a0518181816106870152610b450152f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806250549914610724578063181f5a77146106ab57806321df0da71461040a57806344faef45146106485780636170c4b11461054a57806379ba5097146104615780638da5cb5b1461040f5780639608b2321461040a578063a6801258146103bf578063bd028e7c146101b45763f2fde38b1461009457600080fd5b346101af5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af576100cb610941565b73ffffffffffffffffffffffffffffffffffffffff60015416908133036101855773ffffffffffffffffffffffffffffffffffffffff169033821461015b57817fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000557fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101af5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5760043567ffffffffffffffff81116101af57366023820112156101af5780600401359067ffffffffffffffff82116101af576024810190602436918460061b0101116101af5773ffffffffffffffffffffffffffffffffffffffff6001541633036103915760005b82811061025457005b61025f818484610aeb565b359073ffffffffffffffffffffffffffffffffffffffff82168092036101af578115610367576020610292828686610aeb565b0135918215158093036101af576001928160005260026020528060ff604060002054161515036102c5575b50500161024b565b81600052600260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527fb2cc4dde7f9044ba1999f7843e2f9cd1e4ce506f8cc2e16de26ce982bf113fa6602073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692a384806102bd565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101af5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5760206104006103fb610941565b610aab565b6040519015158152f35b610a3c565b346101af5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101af5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5760005473ffffffffffffffffffffffffffffffffffffffff81163303610520577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101af5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5761058161092a565b61058e6024358092610b2a565b7f0000000000000000000000000000000000000000000000000000000000000000906106036040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152826064820152606481526105fd608482610964565b83610bd3565b6040519081527f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62602073ffffffffffffffffffffffffffffffffffffffff33941692a3005b346101af5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af57602060405167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101af5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5761072060408051906106ec8183610964565b601682527f45524332304c6f636b426f7820312e372e302d64657600000000000000000000602083015251918291826109d4565b0390f35b346101af5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af5761075b61092a565b6024356044359173ffffffffffffffffffffffffffffffffffffffff83168093036101af578161078a91610b2a565b8115610900577f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff8216916040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481875afa9081156108f4576000916108c2575b5080831161089157507f2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b639891610888602092604051907fa9059cbb000000000000000000000000000000000000000000000000000000008583015287602483015283604483015260448252610883606483610964565b610bd3565b604051908152a3005b827fcf4791810000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b90506020813d6020116108ec575b816108dd60209383610964565b810103126101af57518561080e565b3d91506108d0565b6040513d6000823e3d90fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b6004359067ffffffffffffffff821682036101af57565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101af57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176109a557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9190916020815282519283602083015260005b848110610a265750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b80602080928401015160408286010152016109e7565b346101af5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101af57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff80600154169116908114908115610ad5575090565b9050600052600260205260ff6040600020541690565b9190811015610afb5760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015610ba95767ffffffffffffffff1667ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168103610b7c5750610b7533610aab565b1561039157565b7f29cf73780000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff16604091600080845192610bfc8685610964565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610d40573d9067ffffffffffffffff82116109a557610c909360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160192610c8187519485610964565b83523d6000602085013e610d49565b805180610c9c57505050565b81602091810103126101af57602001518015908115036101af57610cbd5750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91610c90926060915b91929015610dc45750815115610d5d575090565b3b15610d665790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610dd75750805190602001fd5b610e0d906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016109d4565b0390fdfea164736f6c634300081a000aa164736f6c634300081a000a",
}

var TokenPoolFactoryABI = TokenPoolFactoryMetaData.ABI

var TokenPoolFactoryBin = TokenPoolFactoryMetaData.Bin

func DeployTokenPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address, tokenAdminModule common.Address, rmnProxy common.Address, ccipRouter common.Address) (common.Address, *types.Transaction, *TokenPoolFactory, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenPoolFactoryBin), backend, tokenAdminRegistry, tokenAdminModule, rmnProxy, ccipRouter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPoolFactory{address: address, abi: *parsed, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

type TokenPoolFactory struct {
	address common.Address
	abi     abi.ABI
	TokenPoolFactoryCaller
	TokenPoolFactoryTransactor
	TokenPoolFactoryFilterer
}

type TokenPoolFactoryCaller struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryTransactor struct {
	contract *bind.BoundContract
}

type TokenPoolFactoryFilterer struct {
	contract *bind.BoundContract
}

type TokenPoolFactorySession struct {
	Contract     *TokenPoolFactory
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryCallerSession struct {
	Contract *TokenPoolFactoryCaller
	CallOpts bind.CallOpts
}

type TokenPoolFactoryTransactorSession struct {
	Contract     *TokenPoolFactoryTransactor
	TransactOpts bind.TransactOpts
}

type TokenPoolFactoryRaw struct {
	Contract *TokenPoolFactory
}

type TokenPoolFactoryCallerRaw struct {
	Contract *TokenPoolFactoryCaller
}

type TokenPoolFactoryTransactorRaw struct {
	Contract *TokenPoolFactoryTransactor
}

func NewTokenPoolFactory(address common.Address, backend bind.ContractBackend) (*TokenPoolFactory, error) {
	abi, err := abi.JSON(strings.NewReader(TokenPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactory{address: address, abi: abi, TokenPoolFactoryCaller: TokenPoolFactoryCaller{contract: contract}, TokenPoolFactoryTransactor: TokenPoolFactoryTransactor{contract: contract}, TokenPoolFactoryFilterer: TokenPoolFactoryFilterer{contract: contract}}, nil
}

func NewTokenPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenPoolFactoryCaller, error) {
	contract, err := bindTokenPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryCaller{contract: contract}, nil
}

func NewTokenPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPoolFactoryTransactor, error) {
	contract, err := bindTokenPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryTransactor{contract: contract}, nil
}

func NewTokenPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPoolFactoryFilterer, error) {
	contract, err := bindTokenPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryFilterer{contract: contract}, nil
}

func bindTokenPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.TokenPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.TokenPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transfer(opts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.contract.Transact(opts, method, params...)
}

func (_TokenPoolFactory *TokenPoolFactoryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenPoolFactory.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TokenPoolFactory *TokenPoolFactorySession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryCallerSession) TypeAndVersion() (string, error) {
	return _TokenPoolFactory.Contract.TypeAndVersion(&_TokenPoolFactory.CallOpts)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenAndTokenPool", remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, lockBox, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, lockBox, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenAndTokenPool(remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenAndTokenPool(&_TokenPoolFactory.TransactOpts, remoteTokenPools, localTokenDecimals, localPoolType, tokenInitCode, tokenPoolInitCode, lockBox, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactor) DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.contract.Transact(opts, "deployTokenPoolWithExistingToken", token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, lockBox, salt)
}

func (_TokenPoolFactory *TokenPoolFactorySession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, lockBox, salt)
}

func (_TokenPoolFactory *TokenPoolFactoryTransactorSession) DeployTokenPoolWithExistingToken(token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error) {
	return _TokenPoolFactory.Contract.DeployTokenPoolWithExistingToken(&_TokenPoolFactory.TransactOpts, token, localTokenDecimals, localPoolType, remoteTokenPools, tokenPoolInitCode, lockBox, salt)
}

type TokenPoolFactoryRemoteChainConfigUpdatedIterator struct {
	Event *TokenPoolFactoryRemoteChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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
		it.Event = new(TokenPoolFactoryRemoteChainConfigUpdated)
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

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolFactoryRemoteChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolFactoryRemoteChainConfigUpdated struct {
	RemoteChainSelector uint64
	RemoteChainConfig   TokenPoolFactoryRemoteChainConfig
	Raw                 types.Log
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.FilterLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolFactoryRemoteChainConfigUpdatedIterator{contract: _TokenPoolFactory.contract, event: "RemoteChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_TokenPoolFactory *TokenPoolFactoryFilterer) WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolFactory.contract.WatchLogs(opts, "RemoteChainConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolFactoryRemoteChainConfigUpdated)
				if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
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

func (_TokenPoolFactory *TokenPoolFactoryFilterer) ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error) {
	event := new(TokenPoolFactoryRemoteChainConfigUpdated)
	if err := _TokenPoolFactory.contract.UnpackLog(event, "RemoteChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (TokenPoolFactoryRemoteChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xcf2e104173e7782dc2782d45728a7c097f4abfd93ed53dbf6c39da81c1a8f33c")
}

func (_TokenPoolFactory *TokenPoolFactory) Address() common.Address {
	return _TokenPoolFactory.address
}

type TokenPoolFactoryInterface interface {
	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployTokenAndTokenPool(opts *bind.TransactOpts, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, localTokenDecimals uint8, localPoolType uint8, tokenInitCode []byte, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error)

	DeployTokenPoolWithExistingToken(opts *bind.TransactOpts, token common.Address, localTokenDecimals uint8, localPoolType uint8, remoteTokenPools []TokenPoolFactoryRemoteTokenPoolInfo, tokenPoolInitCode []byte, lockBox common.Address, salt [32]byte) (*types.Transaction, error)

	FilterRemoteChainConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolFactoryRemoteChainConfigUpdatedIterator, error)

	WatchRemoteChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolFactoryRemoteChainConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemoteChainConfigUpdated(log types.Log) (*TokenPoolFactoryRemoteChainConfigUpdated, error)

	Address() common.Address
}
