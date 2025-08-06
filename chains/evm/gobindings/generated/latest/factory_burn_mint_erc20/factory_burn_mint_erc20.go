// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory_burn_mint_erc20

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

var FactoryBurnMintERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBurners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHyperEVMLinker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBurner\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMinter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHyperEVMLinker\",\"inputs\":[{\"name\":\"newHyperEVMLinker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessGranted\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessRevoked\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HyperEVMLinkerSet\",\"inputs\":[{\"name\":\"hyperEVMLinker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessGranted\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessRevoked\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LinkerAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotBurner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotMinter\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c060405234610487576128a5803803806100198161048c565b92833981019060c0818303126104875780516001600160401b03811161048757826100459183016104b1565b602082015190926001600160401b038211610487576100659183016104b1565b604082015160ff811681036104875760608301519160a060808501519401519460018060a01b038616809603610487578051906001600160401b0382116103845760035490600182811c9216801561047d575b60208310146103645781601f84931161040d575b50602090601f83116001146103a55760009261039a575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116103845760045490600182811c9216801561037a575b60208310146103645781601f8493116102f4575b50602090601f831160011461028c57600092610281575b50508160011b916000199060031b1c1916176004555b3315610270573360018060a01b0319600654161760065560805260a0528160018060a01b03196007541617600755806101b9575b604051612388908161051d82396080518161138c015260a05181818161050101526111bc0152f35b811561022b57600254908082018092116102155760207fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9160009360025584845283825260408420818154019055604051908152a33880610191565b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606490fd5b639b15e16f60e01b60005260046000fd5b015190503880610147565b600460009081528281209350601f198516905b8181106102dc57509084600195949392106102c3575b505050811b0160045561015d565b015160001960f88460031b161c191690553880806102b5565b9293602060018192878601518155019501930161029f565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c8101916020851061035a575b90601f859493920160051c01905b81811061034b5750610130565b6000815584935060010161033e565b9091508190610330565b634e487b7160e01b600052602260045260246000fd5b91607f169161011c565b634e487b7160e01b600052604160045260246000fd5b0151905038806100e3565b600360009081528281209350601f198516905b8181106103f557509084600195949392106103dc575b505050811b016003556100f9565b015160001960f88460031b161c191690553880806103ce565b929360206001819287860151815501950193016103b8565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c81019160208510610473575b90601f859493920160051c01905b81811061046457506100cc565b60008155849350600101610457565b9091508190610449565b91607f16916100b8565b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761038457604052565b81601f82011215610487578051906001600160401b038211610384576104e0601f8301601f191660200161048c565b92828452602083830101116104875760005b82811061050757505060206000918301015290565b806020809284010151828287010152016104f256fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146117855750806306fdde03146116a8578063095ea7b3146114eb57806318160ddd146114af57806323b872dd146113b0578063313ce56714611354578063395093511461131857806340c10f191461114557806342966c68146110ee5780634334614a146110895780634f5632f814610ff85780636618846314610ad65780636b32810b14610f6357806370a0823114610efe57806379ba509714610e1557806379cc679014610adb57806386fe8b4314610d745780638da5cb5b14610d225780638fd6a6ac14610cd057806395d89b4114610b735780639dc29fac14610adb578063a457c2d714610ad6578063a8fa343c14610a2b578063a9059cbb146107c9578063aa271e1a1461075a578063c2e3273d146106c9578063c630948d14610626578063c64d0ebc14610595578063cf30d40814610524578063d5abeb01146104cb578063d73dd62314610486578063dd62ed3e146103f6578063dfc6a39b14610326578063f2fde38b146102365763f81094f3146101a057600080fd5b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff6101ec61193c565b6101f4611dee565b166101fe81612196565b61020457005b60207fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e991604051908152a1005b600080fd5b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff61028261193c565b61028a611dee565b163381146102fc57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055573ffffffffffffffffffffffffffffffffffffffff600654167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761035d61193c565b610365611dee565b73ffffffffffffffffffffffffffffffffffffffff81169081156103cc577f8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f557f58dc48d21bdfdfd38a14ec5c284648779a185303f99274a36e229247ad554b00600080a2005b7f4496fa750000000000000000000000000000000000000000000000000000000060005260046000fd5b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761042d61193c565b73ffffffffffffffffffffffffffffffffffffffff61044a61195f565b9116600052600160205273ffffffffffffffffffffffffffffffffffffffff604060002091166000526020526020604060002054604051908152f35b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610231576104c96104c061193c565b60243590611bf4565b005b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760207f8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f5473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff6105e161193c565b6105e9611dee565b166105f381612321565b6105f957005b60207f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad91604051908152a1005b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff61067261193c565b61067a611dee565b16610684816122c1565b61069a575b610691611dee565b6105f381612321565b7fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea6020604051838152a1610689565b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff61071561193c565b61071d611dee565b16610727816122c1565b61072d57005b60207fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea91604051908152a1005b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760206107bf73ffffffffffffffffffffffffffffffffffffffff6107ab61193c565b166000526009602052604060002054151590565b6040519015158152f35b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761080061193c565b73ffffffffffffffffffffffffffffffffffffffff166024353082146102315733156109a75781156109235733600052600060205260406000205481811061089f5781903360005260006020520360406000205581600052600060205260406000208181540190556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157610a6261193c565b610a6a611dee565b73ffffffffffffffffffffffffffffffffffffffff80600754921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600755167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3005b611982565b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157610b1261193c565b602435610b2c33600052600b602052604060002054151590565b15610b45576104c991610b40823383611c93565b611e39565b7fc820b10b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760405160006004548060011c90600181168015610cc6575b602083108114610c9957828552908115610c575750600114610bf7575b610bf383610be781850382611b48565b604051918291826118d4565b0390f35b91905060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b916000905b808210610c3d57509091508101602001610be7610bd7565b919260018160209254838588010152019101909291610c25565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150610be79050610bd7565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f1691610bba565b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157602073ffffffffffffffffffffffffffffffffffffffff60065416604051908152f35b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157604051806020600a54918281520190600a6000527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a89060005b818110610dff57610bf385610df381870382611b48565b60405191829182611af8565b8254845260209093019260019283019201610ddc565b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760055473ffffffffffffffffffffffffffffffffffffffff81163303610ed4577fffffffffffffffffffffffff00000000000000000000000000000000000000006006549133828416176006551660055573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff610f4a61193c565b1660005260006020526020604060002054604051908152f35b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760405180602060085491828152019060086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee39060005b818110610fe257610bf385610df381870382611b48565b8254845260209093019260019283019201610fcb565b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315773ffffffffffffffffffffffffffffffffffffffff61104461193c565b61104c611dee565b1661105681612000565b61105c57005b60207f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c91604051908152a1005b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760206107bf73ffffffffffffffffffffffffffffffffffffffff6110da61193c565b16600052600b602052604060002054151590565b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761113433600052600b602052604060002054151590565b15610b45576104c960043533611e39565b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761117c61193c565b60243590611197336000526009602052604060002054151590565b156112ea5773ffffffffffffffffffffffffffffffffffffffff1690308214610231577f000000000000000000000000000000000000000000000000000000000000000080151590816112d5575b5061129c57811561123e577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef602082611222600094600254611bb8565b60025584845283825260408420818154019055604051908152a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b6112a890600254611bb8565b7fcbbf11130000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90506112e382600254611bb8565b11836111e5565b7fe2c8c9d5000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760206107bf6104c061193c565b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102315760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610231576113e761193c565b6113ef61195f565b73ffffffffffffffffffffffffffffffffffffffff60443591611413833386611c93565b16913083146102315773ffffffffffffffffffffffffffffffffffffffff169081156109a75782156109235781600052600060205260406000205481811061089f57817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9260209285600052600084520360406000205584600052600082526040600020818154019055604051908152a3602060405160018152f35b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610231576020600254604051908152f35b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315761152261193c565b73ffffffffffffffffffffffffffffffffffffffff166024353082146102315733156116255781156115a157336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b346102315760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102315760405160006003548060011c9060018116801561177b575b602083108114610c9957828552908115610c57575060011461171b57610bf383610be781850382611b48565b91905060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b916000905b80821061176157509091508101602001610be7610bd7565b919260018160209254838588010152019101909291611749565b91607f16916116ef565b346102315760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023157600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361023157817f36372b0700000000000000000000000000000000000000000000000000000000602093149081156118aa575b8115611880575b8115611856575b811561182c575b5015158152f35b7f8fd6a6ac0000000000000000000000000000000000000000000000000000000091501483611825565b7f06e27847000000000000000000000000000000000000000000000000000000008114915061181e565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150611817565b7fe6599b4d0000000000000000000000000000000000000000000000000000000081149150611810565b9190916020815282519283602083015260005b8481106119265750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b80602080928401015160408286010152016118e7565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361023157565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361023157565b346102315760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610231576119b961193c565b6024359060009133835260016020526040832073ffffffffffffffffffffffffffffffffffffffff83168452602052604083205490808210611a745773ffffffffffffffffffffffffffffffffffffffff9103911691308314611a715733156116255782156115a15760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a360206001604051908152f35b80fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b602060408183019282815284518094520192019060005b818110611b1c5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611b0f565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611b8957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b91908201809211611bc557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90611c3073ffffffffffffffffffffffffffffffffffffffff913360005260016020526040600020838516600052602052604060002054611bb8565b9116903082146102315733156116255781156115a157336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3600190565b73ffffffffffffffffffffffffffffffffffffffff909291921690816000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff8416600052602052604060002054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d0e575b50505050565b808210611d905773ffffffffffffffffffffffffffffffffffffffff91039216913083146102315781156116255782156115a15760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a338808080611d08565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b73ffffffffffffffffffffffffffffffffffffffff600654163303611e0f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff168015611f355780600052600060205260406000205491808310611eb1576020817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926000958587528684520360408620558060025403600255604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152fd5b8054821015611fd15760005260206000200190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000818152600b6020526040902054801561218f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611bc557600a54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611bc557818103612120575b505050600a5480156120f1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016120ae81600a611fb9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600a55600052600b60205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61217761213161214293600a611fb9565b90549060031b1c928392600a611fb9565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600b602052604060002055388080612075565b5050600090565b600081815260096020526040902054801561218f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611bc557600854907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611bc557818103612287575b50505060085480156120f1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612244816008611fb9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600855600052600960205260006040812055600190565b6122a9612298612142936008611fb9565b90549060031b1c9283926008611fb9565b9055600052600960205260406000205538808061220b565b8060005260096020526040600020541560001461231b5760085468010000000000000000811015611b89576123026121428260018594016008556008611fb9565b9055600854906000526009602052604060002055600190565b50600090565b80600052600b6020526040600020541560001461231b57600a5468010000000000000000811015611b8957612362612142826001859401600a55600a611fb9565b9055600a5490600052600b60205260406000205560019056fea164736f6c634300081a000a",
}

var FactoryBurnMintERC20ABI = FactoryBurnMintERC20MetaData.ABI

var FactoryBurnMintERC20Bin = FactoryBurnMintERC20MetaData.Bin

func DeployFactoryBurnMintERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, newOwner common.Address) (common.Address, *types.Transaction, *FactoryBurnMintERC20, error) {
	parsed, err := FactoryBurnMintERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryBurnMintERC20Bin), backend, name, symbol, decimals_, maxSupply_, preMint, newOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FactoryBurnMintERC20{address: address, abi: *parsed, FactoryBurnMintERC20Caller: FactoryBurnMintERC20Caller{contract: contract}, FactoryBurnMintERC20Transactor: FactoryBurnMintERC20Transactor{contract: contract}, FactoryBurnMintERC20Filterer: FactoryBurnMintERC20Filterer{contract: contract}}, nil
}

type FactoryBurnMintERC20 struct {
	address common.Address
	abi     abi.ABI
	FactoryBurnMintERC20Caller
	FactoryBurnMintERC20Transactor
	FactoryBurnMintERC20Filterer
}

type FactoryBurnMintERC20Caller struct {
	contract *bind.BoundContract
}

type FactoryBurnMintERC20Transactor struct {
	contract *bind.BoundContract
}

type FactoryBurnMintERC20Filterer struct {
	contract *bind.BoundContract
}

type FactoryBurnMintERC20Session struct {
	Contract     *FactoryBurnMintERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FactoryBurnMintERC20CallerSession struct {
	Contract *FactoryBurnMintERC20Caller
	CallOpts bind.CallOpts
}

type FactoryBurnMintERC20TransactorSession struct {
	Contract     *FactoryBurnMintERC20Transactor
	TransactOpts bind.TransactOpts
}

type FactoryBurnMintERC20Raw struct {
	Contract *FactoryBurnMintERC20
}

type FactoryBurnMintERC20CallerRaw struct {
	Contract *FactoryBurnMintERC20Caller
}

type FactoryBurnMintERC20TransactorRaw struct {
	Contract *FactoryBurnMintERC20Transactor
}

func NewFactoryBurnMintERC20(address common.Address, backend bind.ContractBackend) (*FactoryBurnMintERC20, error) {
	abi, err := abi.JSON(strings.NewReader(FactoryBurnMintERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFactoryBurnMintERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20{address: address, abi: abi, FactoryBurnMintERC20Caller: FactoryBurnMintERC20Caller{contract: contract}, FactoryBurnMintERC20Transactor: FactoryBurnMintERC20Transactor{contract: contract}, FactoryBurnMintERC20Filterer: FactoryBurnMintERC20Filterer{contract: contract}}, nil
}

func NewFactoryBurnMintERC20Caller(address common.Address, caller bind.ContractCaller) (*FactoryBurnMintERC20Caller, error) {
	contract, err := bindFactoryBurnMintERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20Caller{contract: contract}, nil
}

func NewFactoryBurnMintERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*FactoryBurnMintERC20Transactor, error) {
	contract, err := bindFactoryBurnMintERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20Transactor{contract: contract}, nil
}

func NewFactoryBurnMintERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*FactoryBurnMintERC20Filterer, error) {
	contract, err := bindFactoryBurnMintERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20Filterer{contract: contract}, nil
}

func bindFactoryBurnMintERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryBurnMintERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryBurnMintERC20.Contract.FactoryBurnMintERC20Caller.contract.Call(opts, result, method, params...)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.FactoryBurnMintERC20Transactor.contract.Transfer(opts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.FactoryBurnMintERC20Transactor.contract.Transact(opts, method, params...)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryBurnMintERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.contract.Transfer(opts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.contract.Transact(opts, method, params...)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.Allowance(&_FactoryBurnMintERC20.CallOpts, owner, spender)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.Allowance(&_FactoryBurnMintERC20.CallOpts, owner, spender)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.BalanceOf(&_FactoryBurnMintERC20.CallOpts, account)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.BalanceOf(&_FactoryBurnMintERC20.CallOpts, account)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Decimals() (uint8, error) {
	return _FactoryBurnMintERC20.Contract.Decimals(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) Decimals() (uint8, error) {
	return _FactoryBurnMintERC20.Contract.Decimals(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GetBurners() ([]common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetBurners(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) GetBurners() ([]common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetBurners(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GetCCIPAdmin() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetCCIPAdmin(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) GetCCIPAdmin() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetCCIPAdmin(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) GetHyperEVMLinker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "getHyperEVMLinker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GetHyperEVMLinker() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetHyperEVMLinker(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) GetHyperEVMLinker() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetHyperEVMLinker(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GetMinters() ([]common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetMinters(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) GetMinters() ([]common.Address, error) {
	return _FactoryBurnMintERC20.Contract.GetMinters(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "isBurner", burner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) IsBurner(burner common.Address) (bool, error) {
	return _FactoryBurnMintERC20.Contract.IsBurner(&_FactoryBurnMintERC20.CallOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) IsBurner(burner common.Address) (bool, error) {
	return _FactoryBurnMintERC20.Contract.IsBurner(&_FactoryBurnMintERC20.CallOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "isMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) IsMinter(minter common.Address) (bool, error) {
	return _FactoryBurnMintERC20.Contract.IsMinter(&_FactoryBurnMintERC20.CallOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) IsMinter(minter common.Address) (bool, error) {
	return _FactoryBurnMintERC20.Contract.IsMinter(&_FactoryBurnMintERC20.CallOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) MaxSupply() (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.MaxSupply(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) MaxSupply() (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.MaxSupply(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Name() (string, error) {
	return _FactoryBurnMintERC20.Contract.Name(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) Name() (string, error) {
	return _FactoryBurnMintERC20.Contract.Name(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Owner() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.Owner(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) Owner() (common.Address, error) {
	return _FactoryBurnMintERC20.Contract.Owner(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FactoryBurnMintERC20.Contract.SupportsInterface(&_FactoryBurnMintERC20.CallOpts, interfaceId)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FactoryBurnMintERC20.Contract.SupportsInterface(&_FactoryBurnMintERC20.CallOpts, interfaceId)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Symbol() (string, error) {
	return _FactoryBurnMintERC20.Contract.Symbol(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) Symbol() (string, error) {
	return _FactoryBurnMintERC20.Contract.Symbol(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) TotalSupply() (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.TotalSupply(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _FactoryBurnMintERC20.Contract.TotalSupply(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "acceptOwnership")
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) AcceptOwnership() (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.AcceptOwnership(&_FactoryBurnMintERC20.TransactOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.AcceptOwnership(&_FactoryBurnMintERC20.TransactOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "approve", spender, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Approve(&_FactoryBurnMintERC20.TransactOpts, spender, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Approve(&_FactoryBurnMintERC20.TransactOpts, spender, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "burn", amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Burn(&_FactoryBurnMintERC20.TransactOpts, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Burn(&_FactoryBurnMintERC20.TransactOpts, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "burn0", account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Burn0(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Burn0(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "burnFrom", account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.BurnFrom(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.BurnFrom(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.DecreaseAllowance(&_FactoryBurnMintERC20.TransactOpts, spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.DecreaseAllowance(&_FactoryBurnMintERC20.TransactOpts, spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.DecreaseApproval(&_FactoryBurnMintERC20.TransactOpts, spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.DecreaseApproval(&_FactoryBurnMintERC20.TransactOpts, spender, subtractedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "grantBurnRole", burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantBurnRole(&_FactoryBurnMintERC20.TransactOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantBurnRole(&_FactoryBurnMintERC20.TransactOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantMintAndBurnRoles(&_FactoryBurnMintERC20.TransactOpts, burnAndMinter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantMintAndBurnRoles(&_FactoryBurnMintERC20.TransactOpts, burnAndMinter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "grantMintRole", minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantMintRole(&_FactoryBurnMintERC20.TransactOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.GrantMintRole(&_FactoryBurnMintERC20.TransactOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.IncreaseAllowance(&_FactoryBurnMintERC20.TransactOpts, spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.IncreaseAllowance(&_FactoryBurnMintERC20.TransactOpts, spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.IncreaseApproval(&_FactoryBurnMintERC20.TransactOpts, spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.IncreaseApproval(&_FactoryBurnMintERC20.TransactOpts, spender, addedValue)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "mint", account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Mint(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Mint(&_FactoryBurnMintERC20.TransactOpts, account, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "revokeBurnRole", burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.RevokeBurnRole(&_FactoryBurnMintERC20.TransactOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.RevokeBurnRole(&_FactoryBurnMintERC20.TransactOpts, burner)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "revokeMintRole", minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.RevokeMintRole(&_FactoryBurnMintERC20.TransactOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.RevokeMintRole(&_FactoryBurnMintERC20.TransactOpts, minter)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.SetCCIPAdmin(&_FactoryBurnMintERC20.TransactOpts, newAdmin)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.SetCCIPAdmin(&_FactoryBurnMintERC20.TransactOpts, newAdmin)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) SetHyperEVMLinker(opts *bind.TransactOpts, newHyperEVMLinker common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "setHyperEVMLinker", newHyperEVMLinker)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) SetHyperEVMLinker(newHyperEVMLinker common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.SetHyperEVMLinker(&_FactoryBurnMintERC20.TransactOpts, newHyperEVMLinker)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) SetHyperEVMLinker(newHyperEVMLinker common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.SetHyperEVMLinker(&_FactoryBurnMintERC20.TransactOpts, newHyperEVMLinker)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "transfer", to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Transfer(&_FactoryBurnMintERC20.TransactOpts, to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.Transfer(&_FactoryBurnMintERC20.TransactOpts, to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.TransferFrom(&_FactoryBurnMintERC20.TransactOpts, from, to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.TransferFrom(&_FactoryBurnMintERC20.TransactOpts, from, to, amount)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.contract.Transact(opts, "transferOwnership", to)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.TransferOwnership(&_FactoryBurnMintERC20.TransactOpts, to)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FactoryBurnMintERC20.Contract.TransferOwnership(&_FactoryBurnMintERC20.TransactOpts, to)
}

type FactoryBurnMintERC20ApprovalIterator struct {
	Event *FactoryBurnMintERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20Approval)
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
		it.Event = new(FactoryBurnMintERC20Approval)
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

func (it *FactoryBurnMintERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FactoryBurnMintERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20ApprovalIterator{contract: _FactoryBurnMintERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20Approval)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseApproval(log types.Log) (*FactoryBurnMintERC20Approval, error) {
	event := new(FactoryBurnMintERC20Approval)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20BurnAccessGrantedIterator struct {
	Event *FactoryBurnMintERC20BurnAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20BurnAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20BurnAccessGranted)
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
		it.Event = new(FactoryBurnMintERC20BurnAccessGranted)
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

func (it *FactoryBurnMintERC20BurnAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20BurnAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20BurnAccessGranted struct {
	Burner common.Address
	Raw    types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterBurnAccessGranted(opts *bind.FilterOpts) (*FactoryBurnMintERC20BurnAccessGrantedIterator, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "BurnAccessGranted")
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20BurnAccessGrantedIterator{contract: _FactoryBurnMintERC20.contract, event: "BurnAccessGranted", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20BurnAccessGranted) (event.Subscription, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "BurnAccessGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20BurnAccessGranted)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseBurnAccessGranted(log types.Log) (*FactoryBurnMintERC20BurnAccessGranted, error) {
	event := new(FactoryBurnMintERC20BurnAccessGranted)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20BurnAccessRevokedIterator struct {
	Event *FactoryBurnMintERC20BurnAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20BurnAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20BurnAccessRevoked)
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
		it.Event = new(FactoryBurnMintERC20BurnAccessRevoked)
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

func (it *FactoryBurnMintERC20BurnAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20BurnAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20BurnAccessRevoked struct {
	Burner common.Address
	Raw    types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterBurnAccessRevoked(opts *bind.FilterOpts) (*FactoryBurnMintERC20BurnAccessRevokedIterator, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "BurnAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20BurnAccessRevokedIterator{contract: _FactoryBurnMintERC20.contract, event: "BurnAccessRevoked", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20BurnAccessRevoked) (event.Subscription, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "BurnAccessRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20BurnAccessRevoked)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseBurnAccessRevoked(log types.Log) (*FactoryBurnMintERC20BurnAccessRevoked, error) {
	event := new(FactoryBurnMintERC20BurnAccessRevoked)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20CCIPAdminTransferredIterator struct {
	Event *FactoryBurnMintERC20CCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20CCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20CCIPAdminTransferred)
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
		it.Event = new(FactoryBurnMintERC20CCIPAdminTransferred)
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

func (it *FactoryBurnMintERC20CCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20CCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20CCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*FactoryBurnMintERC20CCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20CCIPAdminTransferredIterator{contract: _FactoryBurnMintERC20.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20CCIPAdminTransferred)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseCCIPAdminTransferred(log types.Log) (*FactoryBurnMintERC20CCIPAdminTransferred, error) {
	event := new(FactoryBurnMintERC20CCIPAdminTransferred)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20HyperEVMLinkerSetIterator struct {
	Event *FactoryBurnMintERC20HyperEVMLinkerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20HyperEVMLinkerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20HyperEVMLinkerSet)
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
		it.Event = new(FactoryBurnMintERC20HyperEVMLinkerSet)
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

func (it *FactoryBurnMintERC20HyperEVMLinkerSetIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20HyperEVMLinkerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20HyperEVMLinkerSet struct {
	HyperEVMLinker common.Address
	Raw            types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterHyperEVMLinkerSet(opts *bind.FilterOpts, hyperEVMLinker []common.Address) (*FactoryBurnMintERC20HyperEVMLinkerSetIterator, error) {

	var hyperEVMLinkerRule []interface{}
	for _, hyperEVMLinkerItem := range hyperEVMLinker {
		hyperEVMLinkerRule = append(hyperEVMLinkerRule, hyperEVMLinkerItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "HyperEVMLinkerSet", hyperEVMLinkerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20HyperEVMLinkerSetIterator{contract: _FactoryBurnMintERC20.contract, event: "HyperEVMLinkerSet", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchHyperEVMLinkerSet(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20HyperEVMLinkerSet, hyperEVMLinker []common.Address) (event.Subscription, error) {

	var hyperEVMLinkerRule []interface{}
	for _, hyperEVMLinkerItem := range hyperEVMLinker {
		hyperEVMLinkerRule = append(hyperEVMLinkerRule, hyperEVMLinkerItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "HyperEVMLinkerSet", hyperEVMLinkerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20HyperEVMLinkerSet)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "HyperEVMLinkerSet", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseHyperEVMLinkerSet(log types.Log) (*FactoryBurnMintERC20HyperEVMLinkerSet, error) {
	event := new(FactoryBurnMintERC20HyperEVMLinkerSet)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "HyperEVMLinkerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20MintAccessGrantedIterator struct {
	Event *FactoryBurnMintERC20MintAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20MintAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20MintAccessGranted)
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
		it.Event = new(FactoryBurnMintERC20MintAccessGranted)
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

func (it *FactoryBurnMintERC20MintAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20MintAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20MintAccessGranted struct {
	Minter common.Address
	Raw    types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterMintAccessGranted(opts *bind.FilterOpts) (*FactoryBurnMintERC20MintAccessGrantedIterator, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "MintAccessGranted")
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20MintAccessGrantedIterator{contract: _FactoryBurnMintERC20.contract, event: "MintAccessGranted", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20MintAccessGranted) (event.Subscription, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "MintAccessGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20MintAccessGranted)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseMintAccessGranted(log types.Log) (*FactoryBurnMintERC20MintAccessGranted, error) {
	event := new(FactoryBurnMintERC20MintAccessGranted)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20MintAccessRevokedIterator struct {
	Event *FactoryBurnMintERC20MintAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20MintAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20MintAccessRevoked)
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
		it.Event = new(FactoryBurnMintERC20MintAccessRevoked)
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

func (it *FactoryBurnMintERC20MintAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20MintAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20MintAccessRevoked struct {
	Minter common.Address
	Raw    types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterMintAccessRevoked(opts *bind.FilterOpts) (*FactoryBurnMintERC20MintAccessRevokedIterator, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "MintAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20MintAccessRevokedIterator{contract: _FactoryBurnMintERC20.contract, event: "MintAccessRevoked", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20MintAccessRevoked) (event.Subscription, error) {

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "MintAccessRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20MintAccessRevoked)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseMintAccessRevoked(log types.Log) (*FactoryBurnMintERC20MintAccessRevoked, error) {
	event := new(FactoryBurnMintERC20MintAccessRevoked)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20OwnershipTransferRequestedIterator struct {
	Event *FactoryBurnMintERC20OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20OwnershipTransferRequested)
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
		it.Event = new(FactoryBurnMintERC20OwnershipTransferRequested)
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

func (it *FactoryBurnMintERC20OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20OwnershipTransferRequestedIterator{contract: _FactoryBurnMintERC20.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20OwnershipTransferRequested)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseOwnershipTransferRequested(log types.Log) (*FactoryBurnMintERC20OwnershipTransferRequested, error) {
	event := new(FactoryBurnMintERC20OwnershipTransferRequested)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20OwnershipTransferredIterator struct {
	Event *FactoryBurnMintERC20OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20OwnershipTransferred)
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
		it.Event = new(FactoryBurnMintERC20OwnershipTransferred)
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

func (it *FactoryBurnMintERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20OwnershipTransferredIterator{contract: _FactoryBurnMintERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20OwnershipTransferred)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseOwnershipTransferred(log types.Log) (*FactoryBurnMintERC20OwnershipTransferred, error) {
	event := new(FactoryBurnMintERC20OwnershipTransferred)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FactoryBurnMintERC20TransferIterator struct {
	Event *FactoryBurnMintERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FactoryBurnMintERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBurnMintERC20Transfer)
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
		it.Event = new(FactoryBurnMintERC20Transfer)
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

func (it *FactoryBurnMintERC20TransferIterator) Error() error {
	return it.fail
}

func (it *FactoryBurnMintERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FactoryBurnMintERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FactoryBurnMintERC20TransferIterator{contract: _FactoryBurnMintERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FactoryBurnMintERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FactoryBurnMintERC20Transfer)
				if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Filterer) ParseTransfer(log types.Log) (*FactoryBurnMintERC20Transfer, error) {
	event := new(FactoryBurnMintERC20Transfer)
	if err := _FactoryBurnMintERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FactoryBurnMintERC20.abi.Events["Approval"].ID:
		return _FactoryBurnMintERC20.ParseApproval(log)
	case _FactoryBurnMintERC20.abi.Events["BurnAccessGranted"].ID:
		return _FactoryBurnMintERC20.ParseBurnAccessGranted(log)
	case _FactoryBurnMintERC20.abi.Events["BurnAccessRevoked"].ID:
		return _FactoryBurnMintERC20.ParseBurnAccessRevoked(log)
	case _FactoryBurnMintERC20.abi.Events["CCIPAdminTransferred"].ID:
		return _FactoryBurnMintERC20.ParseCCIPAdminTransferred(log)
	case _FactoryBurnMintERC20.abi.Events["HyperEVMLinkerSet"].ID:
		return _FactoryBurnMintERC20.ParseHyperEVMLinkerSet(log)
	case _FactoryBurnMintERC20.abi.Events["MintAccessGranted"].ID:
		return _FactoryBurnMintERC20.ParseMintAccessGranted(log)
	case _FactoryBurnMintERC20.abi.Events["MintAccessRevoked"].ID:
		return _FactoryBurnMintERC20.ParseMintAccessRevoked(log)
	case _FactoryBurnMintERC20.abi.Events["OwnershipTransferRequested"].ID:
		return _FactoryBurnMintERC20.ParseOwnershipTransferRequested(log)
	case _FactoryBurnMintERC20.abi.Events["OwnershipTransferred"].ID:
		return _FactoryBurnMintERC20.ParseOwnershipTransferred(log)
	case _FactoryBurnMintERC20.abi.Events["Transfer"].ID:
		return _FactoryBurnMintERC20.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FactoryBurnMintERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (FactoryBurnMintERC20BurnAccessGranted) Topic() common.Hash {
	return common.HexToHash("0x92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad")
}

func (FactoryBurnMintERC20BurnAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0x0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c")
}

func (FactoryBurnMintERC20CCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (FactoryBurnMintERC20HyperEVMLinkerSet) Topic() common.Hash {
	return common.HexToHash("0x58dc48d21bdfdfd38a14ec5c284648779a185303f99274a36e229247ad554b00")
}

func (FactoryBurnMintERC20MintAccessGranted) Topic() common.Hash {
	return common.HexToHash("0xe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea")
}

func (FactoryBurnMintERC20MintAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0xed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9")
}

func (FactoryBurnMintERC20OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FactoryBurnMintERC20OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FactoryBurnMintERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20) Address() common.Address {
	return _FactoryBurnMintERC20.address
}

type FactoryBurnMintERC20Interface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetBurners(opts *bind.CallOpts) ([]common.Address, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetHyperEVMLinker(opts *bind.CallOpts) (common.Address, error)

	GetMinters(opts *bind.CallOpts) ([]common.Address, error)

	IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error)

	IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetHyperEVMLinker(opts *bind.TransactOpts, newHyperEVMLinker common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FactoryBurnMintERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*FactoryBurnMintERC20Approval, error)

	FilterBurnAccessGranted(opts *bind.FilterOpts) (*FactoryBurnMintERC20BurnAccessGrantedIterator, error)

	WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20BurnAccessGranted) (event.Subscription, error)

	ParseBurnAccessGranted(log types.Log) (*FactoryBurnMintERC20BurnAccessGranted, error)

	FilterBurnAccessRevoked(opts *bind.FilterOpts) (*FactoryBurnMintERC20BurnAccessRevokedIterator, error)

	WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20BurnAccessRevoked) (event.Subscription, error)

	ParseBurnAccessRevoked(log types.Log) (*FactoryBurnMintERC20BurnAccessRevoked, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*FactoryBurnMintERC20CCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*FactoryBurnMintERC20CCIPAdminTransferred, error)

	FilterHyperEVMLinkerSet(opts *bind.FilterOpts, hyperEVMLinker []common.Address) (*FactoryBurnMintERC20HyperEVMLinkerSetIterator, error)

	WatchHyperEVMLinkerSet(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20HyperEVMLinkerSet, hyperEVMLinker []common.Address) (event.Subscription, error)

	ParseHyperEVMLinkerSet(log types.Log) (*FactoryBurnMintERC20HyperEVMLinkerSet, error)

	FilterMintAccessGranted(opts *bind.FilterOpts) (*FactoryBurnMintERC20MintAccessGrantedIterator, error)

	WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20MintAccessGranted) (event.Subscription, error)

	ParseMintAccessGranted(log types.Log) (*FactoryBurnMintERC20MintAccessGranted, error)

	FilterMintAccessRevoked(opts *bind.FilterOpts) (*FactoryBurnMintERC20MintAccessRevokedIterator, error)

	WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20MintAccessRevoked) (event.Subscription, error)

	ParseMintAccessRevoked(log types.Log) (*FactoryBurnMintERC20MintAccessRevoked, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FactoryBurnMintERC20OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FactoryBurnMintERC20OwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FactoryBurnMintERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *FactoryBurnMintERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*FactoryBurnMintERC20Transfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
