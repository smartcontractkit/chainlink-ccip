// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory_burn_mint_erc20

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

var FactoryBurnMintERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBurners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBurner\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMinter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessGranted\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessRevoked\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessGranted\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessRevoked\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotBurner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotMinter\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c0604052346104b55761280e80380380610019816104ba565b92833981019060c0818303126104b55780516001600160401b0381116104b557826100459183016104df565b602082015190926001600160401b0382116104b5576100659183016104df565b604082015160ff811681036104b55760608301519160a060808501519401519460018060a01b0386168096036104b5578051906001600160401b0382116103b25760035490600182811c921680156104ab575b60208310146103925781601f84931161043b575b50602090601f83116001146103d3576000926103c8575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116103b25760045490600182811c921680156103a8575b60208310146103925781601f849311610322575b50602090601f83116001146102ba576000926102af575b50508160011b916000199060031b1c1916176004555b331561029e573360018060a01b031960065416176006556080528060a0528260018060a01b031960075416176007558082119081610294575b5061028057806101c9575b6040516122c3908161054b823960805181611240015260a05181818161042601526110700152f35b811561023b57600254908082018092116102255760207fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9160009360025584845283825260408420818154019055604051908152a338806101a1565b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606490fd5b63cbbf111360e01b60005260045260246000fd5b9050151538610196565b639b15e16f60e01b60005260046000fd5b015190503880610147565b600460009081528281209350601f198516905b81811061030a57509084600195949392106102f1575b505050811b0160045561015d565b015160001960f88460031b161c191690553880806102e3565b929360206001819287860151815501950193016102cd565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610388575b90601f859493920160051c01905b8181106103795750610130565b6000815584935060010161036c565b909150819061035e565b634e487b7160e01b600052602260045260246000fd5b91607f169161011c565b634e487b7160e01b600052604160045260246000fd5b0151905038806100e3565b600360009081528281209350601f198516905b818110610423575090846001959493921061040a575b505050811b016003556100f9565b015160001960f88460031b161c191690553880806103fc565b929360206001819287860151815501950193016103e6565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c810191602085106104a1575b90601f859493920160051c01905b81811061049257506100cc565b60008155849350600101610485565b9091508190610477565b91607f16916100b8565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176103b257604052565b81601f820112156104b5578051906001600160401b0382116103b25761050e601f8301601f19166020016104ba565b92828452602083830101116104b55760005b82811061053557505060206000918301015290565b8060208092840101518282870101520161052056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146116ef5750806306fdde0314611612578063095ea7b31461145557806318160ddd14611419578063181f5a771461136357806323b872dd14611264578063313ce5671461120857806339509351146111cc57806340c10f1914610ff957806342966c6814610fa25780634334614a14610f3d5780634f5632f814610eac578063661884631461098a5780636b32810b14610e1757806370a0823114610db257806379ba509714610cc957806379cc67901461098f57806386fe8b4314610c285780638da5cb5b14610bd65780638fd6a6ac14610b8457806395d89b4114610a275780639dc29fac1461098f578063a457c2d71461098a578063a8fa343c146108df578063a9059cbb1461067d578063aa271e1a1461060e578063c2e3273d1461057d578063c630948d146104da578063c64d0ebc14610449578063d5abeb01146103f0578063d73dd623146103ab578063dd62ed3e1461031b578063f2fde38b1461022b5763f81094f31461019557600080fd5b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff6101e16118a6565b6101e9611d29565b166101f3816120d1565b6101f957005b60207fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e991604051908152a1005b600080fd5b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff6102776118a6565b61027f611d29565b163381146102f157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055573ffffffffffffffffffffffffffffffffffffffff600654167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576103526118a6565b73ffffffffffffffffffffffffffffffffffffffff61036f6118c9565b9116600052600160205273ffffffffffffffffffffffffffffffffffffffff604060002091166000526020526020604060002054604051908152f35b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576103ee6103e56118a6565b60243590611b2f565b005b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff6104956118a6565b61049d611d29565b166104a78161225c565b6104ad57005b60207f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad91604051908152a1005b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff6105266118a6565b61052e611d29565b16610538816121fc565b61054e575b610545611d29565b6104a78161225c565b7fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea6020604051838152a161053d565b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff6105c96118a6565b6105d1611d29565b166105db816121fc565b6105e157005b60207fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea91604051908152a1005b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657602061067373ffffffffffffffffffffffffffffffffffffffff61065f6118a6565b166000526009602052604060002054151590565b6040519015158152f35b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576106b46118a6565b73ffffffffffffffffffffffffffffffffffffffff1660243530821461022657331561085b5781156107d7573360005260006020526040600020548181106107535781903360005260006020520360406000205581600052600060205260406000208181540190556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576109166118a6565b61091e611d29565b73ffffffffffffffffffffffffffffffffffffffff80600754921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600755167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3005b6118ec565b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576109c66118a6565b6024356109e033600052600b602052604060002054151590565b156109f9576103ee916109f4823383611bce565b611d74565b7fc820b10b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760405160006004548060011c90600181168015610b7a575b602083108114610b4d57828552908115610b0b5750600114610aab575b610aa783610a9b81850382611ab2565b6040519182918261183e565b0390f35b91905060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b916000905b808210610af157509091508101602001610a9b610a8b565b919260018160209254838588010152019101909291610ad9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150610a9b9050610a8b565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f1691610a6e565b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657602073ffffffffffffffffffffffffffffffffffffffff60065416604051908152f35b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657604051806020600a54918281520190600a6000527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a89060005b818110610cb357610aa785610ca781870382611ab2565b60405191829182611a62565b8254845260209093019260019283019201610c90565b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760055473ffffffffffffffffffffffffffffffffffffffff81163303610d88577fffffffffffffffffffffffff00000000000000000000000000000000000000006006549133828416176006551660055573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff610dfe6118a6565b1660005260006020526020604060002054604051908152f35b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760405180602060085491828152019060086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee39060005b818110610e9657610aa785610ca781870382611ab2565b8254845260209093019260019283019201610e7f565b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265773ffffffffffffffffffffffffffffffffffffffff610ef86118a6565b610f00611d29565b16610f0a81611f3b565b610f1057005b60207f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c91604051908152a1005b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657602061067373ffffffffffffffffffffffffffffffffffffffff610f8e6118a6565b16600052600b602052604060002054151590565b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657610fe833600052600b602052604060002054151590565b156109f9576103ee60043533611d74565b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576110306118a6565b6024359061104b336000526009602052604060002054151590565b1561119e5773ffffffffffffffffffffffffffffffffffffffff1690308214610226577f00000000000000000000000000000000000000000000000000000000000000008015159081611189575b506111505781156110f2577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020826110d6600094600254611af3565b60025584845283825260408420818154019055604051908152a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b61115c90600254611af3565b7fcbbf11130000000000000000000000000000000000000000000000000000000060005260045260246000fd5b905061119782600254611af3565b1183611099565b7fe2c8c9d5000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760206106736103e56118a6565b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102265760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265761129b6118a6565b6112a36118c9565b73ffffffffffffffffffffffffffffffffffffffff604435916112c7833386611bce565b16913083146102265773ffffffffffffffffffffffffffffffffffffffff1690811561085b5782156107d75781600052600060205260406000205481811061075357817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9260209285600052600084520360406000205584600052600082526040600020818154019055604051908152a3602060405160018152f35b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657604051604081019080821067ffffffffffffffff8311176113ea57610aa791604052601a81527f466163746f72794275726e4d696e74455243323020312e362e3200000000000060208201526040519182918261183e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576020600254604051908152f35b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265761148c6118a6565b73ffffffffffffffffffffffffffffffffffffffff1660243530821461022657331561158f57811561150b57336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b346102265760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102265760405160006003548060011c906001811680156116e5575b602083108114610b4d57828552908115610b0b575060011461168557610aa783610a9b81850382611ab2565b91905060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b916000905b8082106116cb57509091508101602001610a9b610a8b565b9192600181602092548385880101520191019092916116b3565b91607f1691611659565b346102265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022657600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361022657817f36372b070000000000000000000000000000000000000000000000000000000060209314908115611814575b81156117ea575b81156117c0575b8115611796575b5015158152f35b7f8fd6a6ac000000000000000000000000000000000000000000000000000000009150148361178f565b7f06e278470000000000000000000000000000000000000000000000000000000081149150611788565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150611781565b7fe6599b4d000000000000000000000000000000000000000000000000000000008114915061177a565b9190916020815282519283602083015260005b8481106118905750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201611851565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361022657565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361022657565b346102265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610226576119236118a6565b6024359060009133835260016020526040832073ffffffffffffffffffffffffffffffffffffffff831684526020526040832054908082106119de5773ffffffffffffffffffffffffffffffffffffffff91039116913083146119db57331561158f57821561150b5760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a360206001604051908152f35b80fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b602060408183019282815284518094520192019060005b818110611a865750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611a79565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176113ea57604052565b91908201809211611b0057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90611b6b73ffffffffffffffffffffffffffffffffffffffff913360005260016020526040600020838516600052602052604060002054611af3565b91169030821461022657331561158f57811561150b57336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3600190565b73ffffffffffffffffffffffffffffffffffffffff909291921690816000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff8416600052602052604060002054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c49575b50505050565b808210611ccb5773ffffffffffffffffffffffffffffffffffffffff910392169130831461022657811561158f57821561150b5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a338808080611c43565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b73ffffffffffffffffffffffffffffffffffffffff600654163303611d4a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff168015611e705780600052600060205260406000205491808310611dec576020817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926000958587528684520360408620558060025403600255604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152fd5b8054821015611f0c5760005260206000200190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000818152600b602052604090205480156120ca577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611b0057600a54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611b005781810361205b575b505050600a54801561202c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611fe981600a611ef4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600a55600052600b60205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6120b261206c61207d93600a611ef4565b90549060031b1c928392600a611ef4565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600b602052604060002055388080611fb0565b5050600090565b60008181526009602052604090205480156120ca577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611b0057600854907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611b00578181036121c2575b505050600854801561202c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161217f816008611ef4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600855600052600960205260006040812055600190565b6121e46121d361207d936008611ef4565b90549060031b1c9283926008611ef4565b90556000526009602052604060002055388080612146565b8060005260096020526040600020541560001461225657600854680100000000000000008110156113ea5761223d61207d8260018594016008556008611ef4565b9055600854906000526009602052604060002055600190565b50600090565b80600052600b6020526040600020541560001461225657600a54680100000000000000008110156113ea5761229d61207d826001859401600a55600a611ef4565b9055600a5490600052600b60205260406000205560019056fea164736f6c634300081a000a",
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

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FactoryBurnMintERC20.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20Session) TypeAndVersion() (string, error) {
	return _FactoryBurnMintERC20.Contract.TypeAndVersion(&_FactoryBurnMintERC20.CallOpts)
}

func (_FactoryBurnMintERC20 *FactoryBurnMintERC20CallerSession) TypeAndVersion() (string, error) {
	return _FactoryBurnMintERC20.Contract.TypeAndVersion(&_FactoryBurnMintERC20.CallOpts)
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

	GetMinters(opts *bind.CallOpts) ([]common.Address, error)

	IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error)

	IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

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

	Address() common.Address
}
