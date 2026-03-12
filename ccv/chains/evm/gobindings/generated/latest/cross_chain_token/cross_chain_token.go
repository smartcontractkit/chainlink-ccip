// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross_chain_token

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

type BaseERC20ConstructorParams struct {
	Name      string
	Symbol    string
	MaxSupply *big.Int
	PreMint   *big.Int
	Decimals  uint8
	CcipAdmin common.Address
}

var CrossChainTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"args\",\"type\":\"tuple\",\"internalType\":\"struct BaseERC20.ConstructorParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"burnMintRoleAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURN_MINT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyCCIPAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c0604052346106d157612df380380380610019816106d6565b92833981016060828203126106d15781516001600160401b0381116106d157820160c0818303126106d1576040519160c083016001600160401b038111848210176105c85760405281516001600160401b0381116106d1578161007d9184016106fb565b83526020820151906001600160401b0382116106d15761009e9183016106fb565b90816020840152604081015191604084019283526060820151916060850192835260808101519060ff821682036106d15760a06100e2916080880193845201610766565b9160a0860192835261010260406100fb60208a01610766565b9801610766565b6001600160a01b0381166106cb575033955b518051906001600160401b0382116105c85760035490600182811c921680156106c1575b60208310146105a85781601f849311610651575b50602090601f83116001146105e9576000926105de575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116105c857600454600181811c911680156105be575b60208210146105a857601f8111610543575b50602090601f83116001146104d75760ff939291600091836104cc575b50508160011b916000199060031b1c1916176004555b5116608052825160a052516001600160a01b0316806104c6575033915b815190816103e8575b5050600580546001600160a01b039384166001600160a01b03198216811790925590921690507f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a36001600160a01b038116156103d257600780546001600160d01b031690556102789061077a565b506001600160a01b0381166103c2575b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6600081815260066020527f3195c024b2ddd6d9b8f6c836aa52f67fe69376c8903d009b80229b3ce4425f528054600080516020612dd383398151915291829055909290917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9080a47f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848600081815260066020527f42d20fd6db25ea5a8e33f43724ad72f2ebd9488257fa78c86176b8175fc383fb8054600080516020612dd383398151915291829055909290917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9080a460405161256a908161086982396080518161142c015260a05181818161033001526110e30152f35b6103cb906107c5565b5038610288565b636116401160e11b600052600060045260246000fd5b5180151590816104bc575b506104a85750516001600160a01b0382169081156104925730821461047d57600254818101809111610467576002557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef602060009284845283825260408420818154019055604051908152a3388080610206565b634e487b7160e01b600052601160045260246000fd5b50630bc2c5df60e11b60005260045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b63cbbf111360e01b60005260045260246000fd5b90508111386103f3565b916101fd565b0151905038806101ca565b90601f198316916004600052816000209260005b81811061052b575091600193918560ff97969410610512575b505050811b016004556101e0565b015160001960f88460031b161c19169055388080610504565b929360206001819287860151815501950193016104eb565b60046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c8101916020851061059e575b601f0160051c01905b81811061059257506101ad565b60008155600101610585565b909150819061057c565b634e487b7160e01b600052602260045260246000fd5b90607f169061019b565b634e487b7160e01b600052604160045260246000fd5b015190503880610163565b600360009081528281209350601f198516905b8181106106395750908460019594939210610620575b505050811b01600355610179565b015160001960f88460031b161c19169055388080610612565b929360206001819287860151815501950193016105fc565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c810191602085106106b7575b90601f859493920160051c01905b8181106106a8575061014c565b6000815584935060010161069b565b909150819061068d565b91607f1691610138565b95610114565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176105c857604052565b81601f820112156106d1578051906001600160401b0382116105c85761072a601f8301601f19166020016106d6565b92828452602083830101116106d15760005b82811061075157505060206000918301015290565b8060208092840101518282870101520161073c565b51906001600160a01b03821682036106d157565b600854906001600160a01b0382166107b4576001600160a01b03199091166001600160a01b038216176008556107b19060006107d9565b90565b631fe1e13d60e11b60005260046000fd5b6107b190600080516020612dd38339815191525b60008181526006602090815260408083206001600160a01b038616845290915290205460ff166108615760008181526006602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b505060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146119e957508063022d63fb146119ad57806306fdde03146118d0578063095ea7b31461178e5780630aa6220b146116a857806318160ddd1461166c578063181f5a77146115b657806323b872dd14611560578063248a9ca31461150d578063282c51f3146114b45780632f2ff15d14611450578063313ce567146113f457806336568abe1461126557806340c10f191461105157806342966c681461100e578063634e93da14610eb7578063649a5ec714610c8757806370a0823114610c2257806379cc67901461095657806384ef8ffc14610bd05780638da5cb5b14610bd05780638fd6a6ac14610b7e57806391d1485414610b0557806395d89b41146109ac5780639dc29fac14610956578063a1eda53c146108d1578063a217fddf14610897578063a8fa343c146107ec578063a9059cbb1461079d578063c630948d146106ac578063c91ddc2014610653578063cc8463c81461060a578063cefc1429146104cc578063cf6eefb714610441578063d5391393146103e8578063d547741f14610353578063d5abeb01146102fa578063d602b9fd146102615763dd62ed3e146101cc57600080fd5b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57610203611bd1565b73ffffffffffffffffffffffffffffffffffffffff610220611bf4565b9116600052600160205273ffffffffffffffffffffffffffffffffffffffff604060002091166000526020526020604060002054604051908152f35b600080fd5b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57610298611c9e565b600780547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff166102d357005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1005b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760043561038d611bf4565b81156103be57816103b76103b26103bc94600052600660205260016040600020015490565b611d95565b6122ea565b005b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760206040517f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57604065ffffffffffff6104a66007549065ffffffffffff73ffffffffffffffffffffffffffffffffffffffff83169260a01c1690565b73ffffffffffffffffffffffffffffffffffffffff849392935193168352166020820152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760075473ffffffffffffffffffffffffffffffffffffffff1633036105dc5760075460a081901c65ffffffffffff169073ffffffffffffffffffffffffffffffffffffffff16811580156105d2575b6105a4576105799061057373ffffffffffffffffffffffffffffffffffffffff6008541661227c565b506121a0565b50600780547fffffffffffff0000000000000000000000000000000000000000000000000000169055005b507f19ca5ebb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b504282101561054a565b7fc22c8022000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576020610643611c65565b65ffffffffffff60405191168152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760206040517fcfd2b420c3d2b6ebd6af82f6e29c095b45a072b8d1b5d9eda2a56dcb850acaa68152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576103bc6106e6611bd1565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a660005260066020527f3195c024b2ddd6d9b8f6c836aa52f67fe69376c8903d009b80229b3ce4425f525461073a90611d95565b61074381612149565b507f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84860005260066020527f42d20fd6db25ea5a8e33f43724ad72f2ebd9488257fa78c86176b8175fc383fb5461079890611d95565b612176565b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576107e16107d7611bd1565b6024359033611f1f565b602060405160018152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57610823611bd1565b61082b611c9e565b73ffffffffffffffffffffffffffffffffffffffff80600554921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600555167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3005b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57602060405160008152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576008548060d01c908115158061094c575b156109425760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b0390f35b5050600080610922565b5042821015610911565b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576103bc610990611bd1565b6024359061099c611d0a565b6109a7823383611e02565b61204f565b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760405160006004548060011c90600181168015610afb575b602083108114610ace57828552908115610a8c5750600114610a2c575b61093e83610a2081850382611c17565b60405191829182611b69565b91905060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b916000905b808210610a7257509091508101602001610a20610a10565b919260018160209254838588010152019101909291610a5a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150610a209050610a10565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b91607f16916109f3565b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57610b3c611bf4565b600435600052600660205273ffffffffffffffffffffffffffffffffffffffff60406000209116600052602052602060ff604060002054166040519015158152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57602073ffffffffffffffffffffffffffffffffffffffff60055416604051908152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57602073ffffffffffffffffffffffffffffffffffffffff60085416604051908152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5773ffffffffffffffffffffffffffffffffffffffff610c6e611bd1565b1660005260006020526020604060002054604051908152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760043565ffffffffffff81169081810361025c57610cd2611c9e565b610cdb42612360565b9165ffffffffffff610ceb611c65565b1680821115610e4e57507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9265ffffffffffff826206978080610d3895109118026206978018169061212b565b906008548060d01c80610dca575b50506008805473ffffffffffffffffffffffffffffffffffffffff1660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff9283168152919092166020820152a1005b421115610e235779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006007549260301b169116176007555b8380610d46565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a1610e1c565b0365ffffffffffff8111610e88577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92610d38919061212b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57610eee611bd1565b610ef6611c9e565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020610f33610f2542612360565b610f2d611c65565b9061212b565b65ffffffffffff73ffffffffffffffffffffffffffffffffffffffff610f7c6007549065ffffffffffff73ffffffffffffffffffffffffffffffffffffffff83169260a01c1690565b9690501694600754867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b169216171760075516610fe4575b65ffffffffffff60405191168152a2005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1610fd3565b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57611045611d0a565b6103bc6004353361204f565b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57611088611bd1565b3360009081527f3195c024b2ddd6d9b8f6c836aa52f67fe69376c8903d009b80229b3ce4425f516020526040902054602435919060ff16156112135773ffffffffffffffffffffffffffffffffffffffff163081146111e6577f000000000000000000000000000000000000000000000000000000000000000080151590816111d1575b50611198578015611169576020917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6000938261114c8694600254611c58565b60025584845283825260408420818154019055604051908152a380f35b7fec442f0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b6111a482600254611c58565b7fcbbf11130000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90506111df83600254611c58565b118361110c565b7f17858bbe0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fe2517d3f00000000000000000000000000000000000000000000000000000000600052336004527f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a660245260446000fd5b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760043561129f611bf4565b8115806113bd575b6112fc575b3373ffffffffffffffffffffffffffffffffffffffff8216036112d2576103bc916122ea565b7f6697b2320000000000000000000000000000000000000000000000000000000060005260046000fd5b60075465ffffffffffff60a082901c169073ffffffffffffffffffffffffffffffffffffffff16158015906113ad575b801561139b575b61136557507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff600754166007556112ac565b65ffffffffffff907f19ca5ebb000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b504265ffffffffffff82161015611333565b5065ffffffffffff81161561132c565b5073ffffffffffffffffffffffffffffffffffffffff6008541673ffffffffffffffffffffffffffffffffffffffff8216146112a7565b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760043561148a611bf4565b81156103be57816114af6103b26103bc94600052600660205260016040600020015490565b612208565b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760206040517f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8488152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576020611558600435600052600660205260016040600020015490565b604051908152f35b3461025c5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576107e161159a611bd1565b6115a2611bf4565b604435916115b1833383611e02565b611f1f565b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57604051604081019080821067ffffffffffffffff83111761163d5761093e91604052601981527f43726f7373436861696e546f6b656e20322e302e302d64657600000000000000602082015260405191829182611b69565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576020600254604051908152f35b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576116df611c9e565b6008548060d01c8061170a575b6008805473ffffffffffffffffffffffffffffffffffffffff169055005b4211156117635779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006007549260301b169116176007555b80806116ec565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a161175c565b3461025c5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576117c5611bd1565b73ffffffffffffffffffffffffffffffffffffffff166024353082146118a257331561187357811561184457336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b7f94280d6200000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7fe602df0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b507f17858bbe0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c5760405160006003548060011c906001811680156119a3575b602083108114610ace57828552908115610a8c57506001146119435761093e83610a2081850382611c17565b91905060036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b916000905b80821061198957509091508101602001610a20610a10565b919260018160209254838588010152019101909291611971565b91607f1691611917565b3461025c5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c576020604051620697808152f35b3461025c5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261025c57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361025c57817f314987860000000000000000000000000000000000000000000000000000000060209314908115611b0e575b8115611ab3575b8115611a89575b5015158152f35b7fe6599b4d0000000000000000000000000000000000000000000000000000000091501483611a82565b90507f36372b070000000000000000000000000000000000000000000000000000000081148015611ae5575b90611a7b565b507f8fd6a6ac000000000000000000000000000000000000000000000000000000008114611adf565b90507f7965db0b0000000000000000000000000000000000000000000000000000000081148015611b40575b90611a74565b507f01ffc9a7000000000000000000000000000000000000000000000000000000008114611b3a565b9190916020815282519283602083015260005b848110611bbb5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201611b7c565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361025c57565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361025c57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761163d57604052565b91908201809211610e8857565b6008548060d01c8015159081611c94575b5015611c8a5760a01c65ffffffffffff1690565b5060075460d01c90565b9050421138611c76565b3360009081527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8602052604090205460ff1615611cd757565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b3360009081527f42d20fd6db25ea5a8e33f43724ad72f2ebd9488257fa78c86176b8175fc383fa602052604090205460ff1615611d4357565b7fe2517d3f00000000000000000000000000000000000000000000000000000000600052336004527f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84860245260446000fd5b806000526006602052604060002073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615611dd15750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b73ffffffffffffffffffffffffffffffffffffffff9092919216806000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff8416600052602052604060002054927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8410611e7c575b50505050565b828410611ed35773ffffffffffffffffffffffffffffffffffffffff16903082146118a257801561187357811561184457600052600160205260406000209060005260205260406000209103905538808080611e76565b8373ffffffffffffffffffffffffffffffffffffffff84927ffb8f41b2000000000000000000000000000000000000000000000000000000006000521660045260245260445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff169081156120205773ffffffffffffffffffffffffffffffffffffffff1691821561116957308314611ff2576000828152806020526040812054828110611fbf5791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b6064937fe450d38c0000000000000000000000000000000000000000000000000000000083949352600452602452604452fd5b827f17858bbe0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f96c6fd1e00000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff1680156120205730156120fc576000918183528260205260408320548181106120ca57817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926020928587528684520360408620558060025403600255604051908152a3565b83927fe450d38c0000000000000000000000000000000000000000000000000000000060649552600452602452604452fd5b7f17858bbe00000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b9065ffffffffffff8091169116019065ffffffffffff8211610e8857565b612173907f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66123aa565b90565b612173907f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486123aa565b6008549073ffffffffffffffffffffffffffffffffffffffff82166103be57612173917fffffffffffffffffffffffff000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff831691161760085560006123aa565b908115612219575b612173916123aa565b6008549173ffffffffffffffffffffffffffffffffffffffff83166103be577fffffffffffffffffffffffff000000000000000000000000000000000000000090921673ffffffffffffffffffffffffffffffffffffffff821617600855612210565b6121739073ffffffffffffffffffffffffffffffffffffffff6008541673ffffffffffffffffffffffffffffffffffffffff8216146122bd575b6000612489565b7fffffffffffffffffffffffff0000000000000000000000000000000000000000600854166008556122b6565b9061217391801580612329575b15612489577fffffffffffffffffffffffff000000000000000000000000000000000000000060085416600855612489565b5073ffffffffffffffffffffffffffffffffffffffff6008541673ffffffffffffffffffffffffffffffffffffffff8316146122f7565b65ffffffffffff81116123785765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b806000526006602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260ff604060002054161560001461248257806000526006602052604060002073ffffffffffffffffffffffffffffffffffffffff8316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b806000526006602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260ff6040600020541660001461248257806000526006602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a460019056fea164736f6c634300081a000acfd2b420c3d2b6ebd6af82f6e29c095b45a072b8d1b5d9eda2a56dcb850acaa6",
}

var CrossChainTokenABI = CrossChainTokenMetaData.ABI

var CrossChainTokenBin = CrossChainTokenMetaData.Bin

func DeployCrossChainToken(auth *bind.TransactOpts, backend bind.ContractBackend, args BaseERC20ConstructorParams, burnMintRoleAdmin common.Address, owner common.Address) (common.Address, *types.Transaction, *CrossChainToken, error) {
	parsed, err := CrossChainTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainTokenBin), backend, args, burnMintRoleAdmin, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChainToken{address: address, abi: *parsed, CrossChainTokenCaller: CrossChainTokenCaller{contract: contract}, CrossChainTokenTransactor: CrossChainTokenTransactor{contract: contract}, CrossChainTokenFilterer: CrossChainTokenFilterer{contract: contract}}, nil
}

type CrossChainToken struct {
	address common.Address
	abi     abi.ABI
	CrossChainTokenCaller
	CrossChainTokenTransactor
	CrossChainTokenFilterer
}

type CrossChainTokenCaller struct {
	contract *bind.BoundContract
}

type CrossChainTokenTransactor struct {
	contract *bind.BoundContract
}

type CrossChainTokenFilterer struct {
	contract *bind.BoundContract
}

type CrossChainTokenSession struct {
	Contract     *CrossChainToken
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CrossChainTokenCallerSession struct {
	Contract *CrossChainTokenCaller
	CallOpts bind.CallOpts
}

type CrossChainTokenTransactorSession struct {
	Contract     *CrossChainTokenTransactor
	TransactOpts bind.TransactOpts
}

type CrossChainTokenRaw struct {
	Contract *CrossChainToken
}

type CrossChainTokenCallerRaw struct {
	Contract *CrossChainTokenCaller
}

type CrossChainTokenTransactorRaw struct {
	Contract *CrossChainTokenTransactor
}

func NewCrossChainToken(address common.Address, backend bind.ContractBackend) (*CrossChainToken, error) {
	abi, err := abi.JSON(strings.NewReader(CrossChainTokenABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCrossChainToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainToken{address: address, abi: abi, CrossChainTokenCaller: CrossChainTokenCaller{contract: contract}, CrossChainTokenTransactor: CrossChainTokenTransactor{contract: contract}, CrossChainTokenFilterer: CrossChainTokenFilterer{contract: contract}}, nil
}

func NewCrossChainTokenCaller(address common.Address, caller bind.ContractCaller) (*CrossChainTokenCaller, error) {
	contract, err := bindCrossChainToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenCaller{contract: contract}, nil
}

func NewCrossChainTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainTokenTransactor, error) {
	contract, err := bindCrossChainToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenTransactor{contract: contract}, nil
}

func NewCrossChainTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainTokenFilterer, error) {
	contract, err := bindCrossChainToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenFilterer{contract: contract}, nil
}

func bindCrossChainToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossChainTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CrossChainToken *CrossChainTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainToken.Contract.CrossChainTokenCaller.contract.Call(opts, result, method, params...)
}

func (_CrossChainToken *CrossChainTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainToken.Contract.CrossChainTokenTransactor.contract.Transfer(opts)
}

func (_CrossChainToken *CrossChainTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainToken.Contract.CrossChainTokenTransactor.contract.Transact(opts, method, params...)
}

func (_CrossChainToken *CrossChainTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainToken.Contract.contract.Call(opts, result, method, params...)
}

func (_CrossChainToken *CrossChainTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainToken.Contract.contract.Transfer(opts)
}

func (_CrossChainToken *CrossChainTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainToken.Contract.contract.Transact(opts, method, params...)
}

func (_CrossChainToken *CrossChainTokenCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) BURNERROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.BURNERROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) BURNERROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.BURNERROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) BURNMINTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "BURN_MINT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) BURNMINTADMINROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.BURNMINTADMINROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) BURNMINTADMINROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.BURNMINTADMINROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.DEFAULTADMINROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.DEFAULTADMINROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) MINTERROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.MINTERROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _CrossChainToken.Contract.MINTERROLE(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossChainToken.Contract.Allowance(&_CrossChainToken.CallOpts, owner, spender)
}

func (_CrossChainToken *CrossChainTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossChainToken.Contract.Allowance(&_CrossChainToken.CallOpts, owner, spender)
}

func (_CrossChainToken *CrossChainTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossChainToken.Contract.BalanceOf(&_CrossChainToken.CallOpts, account)
}

func (_CrossChainToken *CrossChainTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossChainToken.Contract.BalanceOf(&_CrossChainToken.CallOpts, account)
}

func (_CrossChainToken *CrossChainTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) Decimals() (uint8, error) {
	return _CrossChainToken.Contract.Decimals(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) Decimals() (uint8, error) {
	return _CrossChainToken.Contract.Decimals(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) DefaultAdmin() (common.Address, error) {
	return _CrossChainToken.Contract.DefaultAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) DefaultAdmin() (common.Address, error) {
	return _CrossChainToken.Contract.DefaultAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossChainToken.Contract.DefaultAdminDelay(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossChainToken.Contract.DefaultAdminDelay(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossChainToken.Contract.DefaultAdminDelayIncreaseWait(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossChainToken.Contract.DefaultAdminDelayIncreaseWait(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) GetCCIPAdmin() (common.Address, error) {
	return _CrossChainToken.Contract.GetCCIPAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _CrossChainToken.Contract.GetCCIPAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossChainToken.Contract.GetRoleAdmin(&_CrossChainToken.CallOpts, role)
}

func (_CrossChainToken *CrossChainTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossChainToken.Contract.GetRoleAdmin(&_CrossChainToken.CallOpts, role)
}

func (_CrossChainToken *CrossChainTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossChainToken.Contract.HasRole(&_CrossChainToken.CallOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossChainToken.Contract.HasRole(&_CrossChainToken.CallOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) MaxSupply() (*big.Int, error) {
	return _CrossChainToken.Contract.MaxSupply(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) MaxSupply() (*big.Int, error) {
	return _CrossChainToken.Contract.MaxSupply(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) Name() (string, error) {
	return _CrossChainToken.Contract.Name(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) Name() (string, error) {
	return _CrossChainToken.Contract.Name(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) Owner() (common.Address, error) {
	return _CrossChainToken.Contract.Owner(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) Owner() (common.Address, error) {
	return _CrossChainToken.Contract.Owner(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_CrossChainToken *CrossChainTokenSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _CrossChainToken.Contract.PendingDefaultAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _CrossChainToken.Contract.PendingDefaultAdmin(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_CrossChainToken *CrossChainTokenSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _CrossChainToken.Contract.PendingDefaultAdminDelay(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _CrossChainToken.Contract.PendingDefaultAdminDelay(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossChainToken.Contract.SupportsInterface(&_CrossChainToken.CallOpts, interfaceId)
}

func (_CrossChainToken *CrossChainTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossChainToken.Contract.SupportsInterface(&_CrossChainToken.CallOpts, interfaceId)
}

func (_CrossChainToken *CrossChainTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) Symbol() (string, error) {
	return _CrossChainToken.Contract.Symbol(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) Symbol() (string, error) {
	return _CrossChainToken.Contract.Symbol(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) TotalSupply() (*big.Int, error) {
	return _CrossChainToken.Contract.TotalSupply(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CrossChainToken.Contract.TotalSupply(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainToken.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainToken *CrossChainTokenSession) TypeAndVersion() (string, error) {
	return _CrossChainToken.Contract.TypeAndVersion(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenCallerSession) TypeAndVersion() (string, error) {
	return _CrossChainToken.Contract.TypeAndVersion(&_CrossChainToken.CallOpts)
}

func (_CrossChainToken *CrossChainTokenTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_CrossChainToken *CrossChainTokenSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossChainToken.Contract.AcceptDefaultAdminTransfer(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossChainToken.Contract.AcceptDefaultAdminTransfer(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "approve", spender, value)
}

func (_CrossChainToken *CrossChainTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Approve(&_CrossChainToken.TransactOpts, spender, value)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Approve(&_CrossChainToken.TransactOpts, spender, value)
}

func (_CrossChainToken *CrossChainTokenTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_CrossChainToken *CrossChainTokenSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.BeginDefaultAdminTransfer(&_CrossChainToken.TransactOpts, newAdmin)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.BeginDefaultAdminTransfer(&_CrossChainToken.TransactOpts, newAdmin)
}

func (_CrossChainToken *CrossChainTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "burn", amount)
}

func (_CrossChainToken *CrossChainTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Burn(&_CrossChainToken.TransactOpts, amount)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Burn(&_CrossChainToken.TransactOpts, amount)
}

func (_CrossChainToken *CrossChainTokenTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "burn0", account, amount)
}

func (_CrossChainToken *CrossChainTokenSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Burn0(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Burn0(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "burnFrom", account, amount)
}

func (_CrossChainToken *CrossChainTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.BurnFrom(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.BurnFrom(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_CrossChainToken *CrossChainTokenSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossChainToken.Contract.CancelDefaultAdminTransfer(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossChainToken.Contract.CancelDefaultAdminTransfer(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_CrossChainToken *CrossChainTokenSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.ChangeDefaultAdminDelay(&_CrossChainToken.TransactOpts, newDelay)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.ChangeDefaultAdminDelay(&_CrossChainToken.TransactOpts, newDelay)
}

func (_CrossChainToken *CrossChainTokenTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_CrossChainToken *CrossChainTokenSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.GrantMintAndBurnRoles(&_CrossChainToken.TransactOpts, burnAndMinter)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.GrantMintAndBurnRoles(&_CrossChainToken.TransactOpts, burnAndMinter)
}

func (_CrossChainToken *CrossChainTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "grantRole", role, account)
}

func (_CrossChainToken *CrossChainTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.GrantRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.GrantRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "mint", account, amount)
}

func (_CrossChainToken *CrossChainTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Mint(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Mint(&_CrossChainToken.TransactOpts, account, amount)
}

func (_CrossChainToken *CrossChainTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "renounceRole", role, account)
}

func (_CrossChainToken *CrossChainTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.RenounceRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.RenounceRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "revokeRole", role, account)
}

func (_CrossChainToken *CrossChainTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.RevokeRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.RevokeRole(&_CrossChainToken.TransactOpts, role, account)
}

func (_CrossChainToken *CrossChainTokenTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_CrossChainToken *CrossChainTokenSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossChainToken.Contract.RollbackDefaultAdminDelay(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossChainToken.Contract.RollbackDefaultAdminDelay(&_CrossChainToken.TransactOpts)
}

func (_CrossChainToken *CrossChainTokenTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_CrossChainToken *CrossChainTokenSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.SetCCIPAdmin(&_CrossChainToken.TransactOpts, newAdmin)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainToken.Contract.SetCCIPAdmin(&_CrossChainToken.TransactOpts, newAdmin)
}

func (_CrossChainToken *CrossChainTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "transfer", to, value)
}

func (_CrossChainToken *CrossChainTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Transfer(&_CrossChainToken.TransactOpts, to, value)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.Transfer(&_CrossChainToken.TransactOpts, to, value)
}

func (_CrossChainToken *CrossChainTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_CrossChainToken *CrossChainTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.TransferFrom(&_CrossChainToken.TransactOpts, from, to, value)
}

func (_CrossChainToken *CrossChainTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainToken.Contract.TransferFrom(&_CrossChainToken.TransactOpts, from, to, value)
}

type CrossChainTokenApprovalIterator struct {
	Event *CrossChainTokenApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenApproval)
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
		it.Event = new(CrossChainTokenApproval)
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

func (it *CrossChainTokenApprovalIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossChainTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenApprovalIterator{contract: _CrossChainToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossChainTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenApproval)
				if err := _CrossChainToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseApproval(log types.Log) (*CrossChainTokenApproval, error) {
	event := new(CrossChainTokenApproval)
	if err := _CrossChainToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenCCIPAdminTransferredIterator struct {
	Event *CrossChainTokenCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenCCIPAdminTransferred)
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
		it.Event = new(CrossChainTokenCCIPAdminTransferred)
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

func (it *CrossChainTokenCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*CrossChainTokenCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenCCIPAdminTransferredIterator{contract: _CrossChainToken.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainTokenCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenCCIPAdminTransferred)
				if err := _CrossChainToken.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseCCIPAdminTransferred(log types.Log) (*CrossChainTokenCCIPAdminTransferred, error) {
	event := new(CrossChainTokenCCIPAdminTransferred)
	if err := _CrossChainToken.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenDefaultAdminDelayChangeCanceledIterator struct {
	Event *CrossChainTokenDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenDefaultAdminDelayChangeCanceled)
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
		it.Event = new(CrossChainTokenDefaultAdminDelayChangeCanceled)
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

func (it *CrossChainTokenDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenDefaultAdminDelayChangeCanceledIterator{contract: _CrossChainToken.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenDefaultAdminDelayChangeCanceled)
				if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossChainTokenDefaultAdminDelayChangeCanceled, error) {
	event := new(CrossChainTokenDefaultAdminDelayChangeCanceled)
	if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenDefaultAdminDelayChangeScheduledIterator struct {
	Event *CrossChainTokenDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenDefaultAdminDelayChangeScheduled)
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
		it.Event = new(CrossChainTokenDefaultAdminDelayChangeScheduled)
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

func (it *CrossChainTokenDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenDefaultAdminDelayChangeScheduledIterator{contract: _CrossChainToken.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenDefaultAdminDelayChangeScheduled)
				if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossChainTokenDefaultAdminDelayChangeScheduled, error) {
	event := new(CrossChainTokenDefaultAdminDelayChangeScheduled)
	if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenDefaultAdminTransferCanceledIterator struct {
	Event *CrossChainTokenDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenDefaultAdminTransferCanceled)
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
		it.Event = new(CrossChainTokenDefaultAdminTransferCanceled)
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

func (it *CrossChainTokenDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenDefaultAdminTransferCanceledIterator{contract: _CrossChainToken.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenDefaultAdminTransferCanceled)
				if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*CrossChainTokenDefaultAdminTransferCanceled, error) {
	event := new(CrossChainTokenDefaultAdminTransferCanceled)
	if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenDefaultAdminTransferScheduledIterator struct {
	Event *CrossChainTokenDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenDefaultAdminTransferScheduled)
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
		it.Event = new(CrossChainTokenDefaultAdminTransferScheduled)
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

func (it *CrossChainTokenDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossChainTokenDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenDefaultAdminTransferScheduledIterator{contract: _CrossChainToken.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenDefaultAdminTransferScheduled)
				if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*CrossChainTokenDefaultAdminTransferScheduled, error) {
	event := new(CrossChainTokenDefaultAdminTransferScheduled)
	if err := _CrossChainToken.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenRoleAdminChangedIterator struct {
	Event *CrossChainTokenRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenRoleAdminChanged)
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
		it.Event = new(CrossChainTokenRoleAdminChanged)
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

func (it *CrossChainTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossChainTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenRoleAdminChangedIterator{contract: _CrossChainToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenRoleAdminChanged)
				if err := _CrossChainToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseRoleAdminChanged(log types.Log) (*CrossChainTokenRoleAdminChanged, error) {
	event := new(CrossChainTokenRoleAdminChanged)
	if err := _CrossChainToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenRoleGrantedIterator struct {
	Event *CrossChainTokenRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenRoleGranted)
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
		it.Event = new(CrossChainTokenRoleGranted)
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

func (it *CrossChainTokenRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossChainTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenRoleGrantedIterator{contract: _CrossChainToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenRoleGranted)
				if err := _CrossChainToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseRoleGranted(log types.Log) (*CrossChainTokenRoleGranted, error) {
	event := new(CrossChainTokenRoleGranted)
	if err := _CrossChainToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenRoleRevokedIterator struct {
	Event *CrossChainTokenRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenRoleRevoked)
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
		it.Event = new(CrossChainTokenRoleRevoked)
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

func (it *CrossChainTokenRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossChainTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenRoleRevokedIterator{contract: _CrossChainToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenRoleRevoked)
				if err := _CrossChainToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseRoleRevoked(log types.Log) (*CrossChainTokenRoleRevoked, error) {
	event := new(CrossChainTokenRoleRevoked)
	if err := _CrossChainToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainTokenTransferIterator struct {
	Event *CrossChainTokenTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainTokenTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainTokenTransfer)
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
		it.Event = new(CrossChainTokenTransfer)
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

func (it *CrossChainTokenTransferIterator) Error() error {
	return it.fail
}

func (it *CrossChainTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_CrossChainToken *CrossChainTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainTokenTransferIterator{contract: _CrossChainToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_CrossChainToken *CrossChainTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossChainTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainTokenTransfer)
				if err := _CrossChainToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_CrossChainToken *CrossChainTokenFilterer) ParseTransfer(log types.Log) (*CrossChainTokenTransfer, error) {
	event := new(CrossChainTokenTransfer)
	if err := _CrossChainToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (CrossChainTokenApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (CrossChainTokenCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (CrossChainTokenDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (CrossChainTokenDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (CrossChainTokenDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (CrossChainTokenDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (CrossChainTokenRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (CrossChainTokenRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (CrossChainTokenRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (CrossChainTokenTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_CrossChainToken *CrossChainToken) Address() common.Address {
	return _CrossChainToken.address
}

type CrossChainTokenInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	BURNMINTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossChainTokenApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossChainTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*CrossChainTokenApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*CrossChainTokenCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainTokenCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*CrossChainTokenCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossChainTokenDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossChainTokenDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossChainTokenDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*CrossChainTokenDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossChainTokenDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossChainTokenDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*CrossChainTokenDefaultAdminTransferScheduled, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossChainTokenRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*CrossChainTokenRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossChainTokenRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*CrossChainTokenRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossChainTokenRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossChainTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*CrossChainTokenRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainTokenTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossChainTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*CrossChainTokenTransfer, error)

	Address() common.Address
}
