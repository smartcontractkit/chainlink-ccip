// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package advanced_pool_hooks_extractor

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

type IPolicyEngineParameter struct {
	Name  [32]byte
	Value []byte
}

type IPolicyEnginePayload struct {
	Selector [4]byte
	Sender   common.Address
	Data     []byte
	Context  []byte
}

var AdvancedPoolHooksExtractorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PARAM_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_AMOUNT_POST_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_FROM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_REMOTE_CHAIN_SELECTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_REQUESTED_FINALITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_SOURCE_DENOMINATED_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_SOURCE_POOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_SOURCE_POOL_DATA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_TO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARAM_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extract\",\"inputs\":[{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"struct IPolicyEngine.Payload\",\"components\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPolicyEngine.Parameter[]\",\"components\":[{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"UnsupportedSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]}]",
	Bin: "0x6080806040523460155761130b908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461052c5780632eb866711461041f57806333c05ead146103c65780638572e6f81461036d5780638709a94a146103145780638a893fd5146102bb578063a4b616f314610262578063be2719ea14610209578063c26cd1ee146101b0578063dac0d49614610157578063e0d169c0146100fe5763e4528bf7146100a057600080fd5b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f0d2d49551f0c0301537208b1e18ac6b2eaad1a8e62061a2579a6123e92cf51378152f35b600080fd5b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f9a6c1917fcf35001153b9951e537209ccb8ebcc5be8c133daa6285a42081b0c48152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f0d172009e3817c908f5f9657cc7c6d88fd284af3f37b66446e02da140dc3da818152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517fcaed9fe748826c17a6bbf34cda465187a44e04fe0ef52519bc6b07e8fd57121b8152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d88152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517fe32332318510df2a33cbbddd86b6f0111e2fb7e55391c4925c5fadaeca9fb4298152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f0c548cc8fd8090ef28614d6a1c6269108d2b4c6d3e100ebab8ebba82671a5d398152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f89c4783cb6cc307f98e95f2d5d5d8647bdb3d4bdd087209374f187b38e0988958152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f1b56e27094b67facb247d55c7c05912fc4cbffd28f63f412fcdd194991f8db488152f35b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760206040517f9b9b0454cadcb5884dd3faa6ba975da4d2459aa3f11d31291a25a8358f84946d8152f35b346100f95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f95760043567ffffffffffffffff81116100f95760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100f957610498906004016106c5565b6040518091602082016020835281518091526040830190602060408260051b8601019301916000905b8282106104d057505050500390f35b9193602061051c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186526040838a518051845201519181858201520190610639565b96019201920185949391926104c1565b346100f95760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100f9576105a9604080519061056d81836105f8565b602082527f416476616e636564506f6f6c486f6f6b73457874726163746f7220322e302e30602083015251918291602083526020830190610639565b0390f35b6040810190811067ffffffffffffffff8211176105c957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176105c957604052565b919082519283825260005b8481106106835750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610644565b357fffffffff00000000000000000000000000000000000000000000000000000000811681036100f95790565b9060607fa8027c0f000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061071385610698565b1614610ce557507f63711574000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061076584610698565b16146107c3577fffffffff0000000000000000000000000000000000000000000000000000000061079583610698565b7fa519a14f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b908060406107d2920190611117565b8101906060818303126100f957803567ffffffffffffffff81116100f957810191610100838203126100f95760405192610100840184811067ffffffffffffffff8211176105c957604052803567ffffffffffffffff81116100f9578261083a918301611168565b8452610848602082016111dd565b926020850193845261085c604083016111f2565b906040860191825260608601946060840135865261087c608085016111f2565b916080880192835260a085013567ffffffffffffffff81116100f957866108a4918701611168565b9460a0890195865260c081013567ffffffffffffffff81116100f957876108cc918301611168565b9660c08a0197885260e082013567ffffffffffffffff81116100f9576108f29201611168565b60e089015261090360408201611213565b9360405198610140610915818c6105f8565b60098b527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00160005b818110610cbf5750509173ffffffffffffffffffffffffffffffffffffffff8095949267ffffffffffffffff945160405190610979826105ad565b7f45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8825260208201526109aa8d611240565b526109b48c611240565b505116604051906020820152602081526109cf6040826105f8565b604051906109dc826105ad565b7f1b56e27094b67facb247d55c7c05912fc4cbffd28f63f412fcdd194991f8db4882526020820152610a0d8b61127c565b52610a178a61127c565b506020604051910135602082015260208152610a346040826105f8565b60405190610a41826105ad565b7f89c4783cb6cc307f98e95f2d5d5d8647bdb3d4bdd087209374f187b38e09889582526020820152610a728a61128c565b52610a7c8961128c565b50511660405190602082015260208152610a976040826105f8565b60405190610aa4826105ad565b7f0d2d49551f0c0301537208b1e18ac6b2eaad1a8e62061a2579a6123e92cf513782526020820152610ad58861129c565b52610adf8761129c565b50511660405190602082015260208152610afa6040826105f8565b60405190610b07826105ad565b7f9b9b0454cadcb5884dd3faa6ba975da4d2459aa3f11d31291a25a8358f84946d82526020820152610b38866112ac565b52610b42856112ac565b507fffffffff000000000000000000000000000000000000000000000000000000006040519116602082015260208152610b7d6040826105f8565b60405190610b8a826105ad565b7f9a6c1917fcf35001153b9951e537209ccb8ebcc5be8c133daa6285a42081b0c482526020820152610bbb856112bc565b52610bc5846112bc565b505160405190610bd4826105ad565b7fe32332318510df2a33cbbddd86b6f0111e2fb7e55391c4925c5fadaeca9fb42982526020820152610c05846112cc565b52610c0f836112cc565b505160405190610c1e826105ad565b7f0d172009e3817c908f5f9657cc7c6d88fd284af3f37b66446e02da140dc3da8182526020820152610c4f836112dc565b52610c59826112dc565b505160405190602082015260208152610c736040826105f8565b60405190610c80826105ad565b7fcaed9fe748826c17a6bbf34cda465187a44e04fe0ef52519bc6b07e8fd57121b82526020820152610cb1826112ed565b52610cbb816112ed565b5090565b808c6020809360405192610cd2846105ad565b600084526060838501520101520161093e565b91806040610cf4920190611117565b81016080828203126100f957813567ffffffffffffffff81116100f95782019160a0838303126100f9576040519060a0820182811067ffffffffffffffff8211176105c957604052833567ffffffffffffffff81116100f95783610d59918601611168565b8252610d67602085016111dd565b9060208301918252610d7b604086016111f2565b9460408401958652610d986080898601928a8101358452016111f2565b9360808101948552610dac60208401611213565b9560408401359067ffffffffffffffff82116100f957610dcd918501611168565b5060405196610100610ddf818a6105f8565b600789527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00160005b8181106110f35750509173ffffffffffffffffffffffffffffffffffffffff96979899918767ffffffffffffffff969594511660405190602082015260208152610e536040826105f8565b60405190610e60826105ad565b7f45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d882526020820152610e918b611240565b52610e9b8a611240565b505160405190610eaa826105ad565b7f1b56e27094b67facb247d55c7c05912fc4cbffd28f63f412fcdd194991f8db4882526020820152610edb8a61127c565b52610ee58961127c565b505160405190602082015260208152610eff6040826105f8565b60405190610f0c826105ad565b7f89c4783cb6cc307f98e95f2d5d5d8647bdb3d4bdd087209374f187b38e09889582526020820152610f3d8961128c565b52610f478861128c565b50604051910135602082015260208152610f626040826105f8565b60405190610f6f826105ad565b7f0c548cc8fd8090ef28614d6a1c6269108d2b4c6d3e100ebab8ebba82671a5d3982526020820152610fa08761129c565b52610faa8661129c565b50511660405190602082015260208152610fc56040826105f8565b60405190610fd2826105ad565b7f0d2d49551f0c0301537208b1e18ac6b2eaad1a8e62061a2579a6123e92cf513782526020820152611003856112ac565b5261100d846112ac565b505116604051906020820152602081526110286040826105f8565b60405190611035826105ad565b7f9b9b0454cadcb5884dd3faa6ba975da4d2459aa3f11d31291a25a8358f84946d82526020820152611066836112bc565b52611070826112bc565b507fffffffff0000000000000000000000000000000000000000000000000000000060405191166020820152602081526110ab6040826105f8565b604051906110b8826105ad565b7f9a6c1917fcf35001153b9951e537209ccb8ebcc5be8c133daa6285a42081b0c4825260208201526110e9826112cc565b52610cbb816112cc565b602090604051611102816105ad565b600081528d8382015282828d01015201610e08565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100f9570180359067ffffffffffffffff82116100f9576020019181360383136100f957565b81601f820112156100f95780359067ffffffffffffffff82116105c957604051926111bb601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856105f8565b828452602083830101116100f957816000926020809301838601378301015290565b359067ffffffffffffffff821682036100f957565b359073ffffffffffffffffffffffffffffffffffffffff821682036100f957565b35907fffffffff00000000000000000000000000000000000000000000000000000000821682036100f957565b80511561124d5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80516001101561124d5760400190565b80516002101561124d5760600190565b80516003101561124d5760800190565b80516004101561124d5760a00190565b80516005101561124d5760c00190565b80516006101561124d5760e00190565b80516007101561124d576101000190565b80516008101561124d57610120019056fea164736f6c634300081a000a",
}

var AdvancedPoolHooksExtractorABI = AdvancedPoolHooksExtractorMetaData.ABI

var AdvancedPoolHooksExtractorBin = AdvancedPoolHooksExtractorMetaData.Bin

func DeployAdvancedPoolHooksExtractor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdvancedPoolHooksExtractor, error) {
	parsed, err := AdvancedPoolHooksExtractorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdvancedPoolHooksExtractorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdvancedPoolHooksExtractor{address: address, abi: *parsed, AdvancedPoolHooksExtractorCaller: AdvancedPoolHooksExtractorCaller{contract: contract}, AdvancedPoolHooksExtractorTransactor: AdvancedPoolHooksExtractorTransactor{contract: contract}, AdvancedPoolHooksExtractorFilterer: AdvancedPoolHooksExtractorFilterer{contract: contract}}, nil
}

type AdvancedPoolHooksExtractor struct {
	address common.Address
	abi     abi.ABI
	AdvancedPoolHooksExtractorCaller
	AdvancedPoolHooksExtractorTransactor
	AdvancedPoolHooksExtractorFilterer
}

type AdvancedPoolHooksExtractorCaller struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksExtractorTransactor struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksExtractorFilterer struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksExtractorSession struct {
	Contract     *AdvancedPoolHooksExtractor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksExtractorCallerSession struct {
	Contract *AdvancedPoolHooksExtractorCaller
	CallOpts bind.CallOpts
}

type AdvancedPoolHooksExtractorTransactorSession struct {
	Contract     *AdvancedPoolHooksExtractorTransactor
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksExtractorRaw struct {
	Contract *AdvancedPoolHooksExtractor
}

type AdvancedPoolHooksExtractorCallerRaw struct {
	Contract *AdvancedPoolHooksExtractorCaller
}

type AdvancedPoolHooksExtractorTransactorRaw struct {
	Contract *AdvancedPoolHooksExtractorTransactor
}

func NewAdvancedPoolHooksExtractor(address common.Address, backend bind.ContractBackend) (*AdvancedPoolHooksExtractor, error) {
	abi, err := abi.JSON(strings.NewReader(AdvancedPoolHooksExtractorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAdvancedPoolHooksExtractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksExtractor{address: address, abi: abi, AdvancedPoolHooksExtractorCaller: AdvancedPoolHooksExtractorCaller{contract: contract}, AdvancedPoolHooksExtractorTransactor: AdvancedPoolHooksExtractorTransactor{contract: contract}, AdvancedPoolHooksExtractorFilterer: AdvancedPoolHooksExtractorFilterer{contract: contract}}, nil
}

func NewAdvancedPoolHooksExtractorCaller(address common.Address, caller bind.ContractCaller) (*AdvancedPoolHooksExtractorCaller, error) {
	contract, err := bindAdvancedPoolHooksExtractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksExtractorCaller{contract: contract}, nil
}

func NewAdvancedPoolHooksExtractorTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedPoolHooksExtractorTransactor, error) {
	contract, err := bindAdvancedPoolHooksExtractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksExtractorTransactor{contract: contract}, nil
}

func NewAdvancedPoolHooksExtractorFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedPoolHooksExtractorFilterer, error) {
	contract, err := bindAdvancedPoolHooksExtractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksExtractorFilterer{contract: contract}, nil
}

func bindAdvancedPoolHooksExtractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdvancedPoolHooksExtractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooksExtractor.Contract.AdvancedPoolHooksExtractorCaller.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooksExtractor.Contract.AdvancedPoolHooksExtractorTransactor.contract.Transfer(opts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooksExtractor.Contract.AdvancedPoolHooksExtractorTransactor.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooksExtractor.Contract.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooksExtractor.Contract.contract.Transfer(opts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooksExtractor.Contract.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMAMOUNT() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMAMOUNT(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMAMOUNT() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMAMOUNT(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMAMOUNTPOSTFEE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_AMOUNT_POST_FEE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMAMOUNTPOSTFEE() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMAMOUNTPOSTFEE(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMAMOUNTPOSTFEE() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMAMOUNTPOSTFEE(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMFROM(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_FROM")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMFROM() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMFROM(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMFROM() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMFROM(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMREMOTECHAINSELECTOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_REMOTE_CHAIN_SELECTOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMREMOTECHAINSELECTOR() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMREMOTECHAINSELECTOR(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMREMOTECHAINSELECTOR() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMREMOTECHAINSELECTOR(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMREQUESTEDFINALITY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_REQUESTED_FINALITY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMREQUESTEDFINALITY() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMREQUESTEDFINALITY(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMREQUESTEDFINALITY() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMREQUESTEDFINALITY(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMSOURCEDENOMINATEDAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_SOURCE_DENOMINATED_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMSOURCEDENOMINATEDAMOUNT() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEDENOMINATEDAMOUNT(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMSOURCEDENOMINATEDAMOUNT() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEDENOMINATEDAMOUNT(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMSOURCEPOOLADDRESS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_SOURCE_POOL_ADDRESS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMSOURCEPOOLADDRESS() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEPOOLADDRESS(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMSOURCEPOOLADDRESS() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEPOOLADDRESS(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMSOURCEPOOLDATA(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_SOURCE_POOL_DATA")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMSOURCEPOOLDATA() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEPOOLDATA(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMSOURCEPOOLDATA() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMSOURCEPOOLDATA(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMTO(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_TO")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMTO() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMTO(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMTO() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMTO(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) PARAMTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "PARAM_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) PARAMTOKEN() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMTOKEN(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) PARAMTOKEN() ([32]byte, error) {
	return _AdvancedPoolHooksExtractor.Contract.PARAMTOKEN(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) Extract(opts *bind.CallOpts, payload IPolicyEnginePayload) ([]IPolicyEngineParameter, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "extract", payload)

	if err != nil {
		return *new([]IPolicyEngineParameter), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPolicyEngineParameter)).(*[]IPolicyEngineParameter)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) Extract(payload IPolicyEnginePayload) ([]IPolicyEngineParameter, error) {
	return _AdvancedPoolHooksExtractor.Contract.Extract(&_AdvancedPoolHooksExtractor.CallOpts, payload)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) Extract(payload IPolicyEnginePayload) ([]IPolicyEngineParameter, error) {
	return _AdvancedPoolHooksExtractor.Contract.Extract(&_AdvancedPoolHooksExtractor.CallOpts, payload)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdvancedPoolHooksExtractor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorSession) TypeAndVersion() (string, error) {
	return _AdvancedPoolHooksExtractor.Contract.TypeAndVersion(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractorCallerSession) TypeAndVersion() (string, error) {
	return _AdvancedPoolHooksExtractor.Contract.TypeAndVersion(&_AdvancedPoolHooksExtractor.CallOpts)
}

func (_AdvancedPoolHooksExtractor *AdvancedPoolHooksExtractor) Address() common.Address {
	return _AdvancedPoolHooksExtractor.address
}

type AdvancedPoolHooksExtractorInterface interface {
	PARAMAMOUNT(opts *bind.CallOpts) ([32]byte, error)

	PARAMAMOUNTPOSTFEE(opts *bind.CallOpts) ([32]byte, error)

	PARAMFROM(opts *bind.CallOpts) ([32]byte, error)

	PARAMREMOTECHAINSELECTOR(opts *bind.CallOpts) ([32]byte, error)

	PARAMREQUESTEDFINALITY(opts *bind.CallOpts) ([32]byte, error)

	PARAMSOURCEDENOMINATEDAMOUNT(opts *bind.CallOpts) ([32]byte, error)

	PARAMSOURCEPOOLADDRESS(opts *bind.CallOpts) ([32]byte, error)

	PARAMSOURCEPOOLDATA(opts *bind.CallOpts) ([32]byte, error)

	PARAMTO(opts *bind.CallOpts) ([32]byte, error)

	PARAMTOKEN(opts *bind.CallOpts) ([32]byte, error)

	Extract(opts *bind.CallOpts, payload IPolicyEnginePayload) ([]IPolicyEngineParameter, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Address() common.Address
}
