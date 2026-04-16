package tip20

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TIP20TokenABI is a minimal ABI for TIP-20 token role management (ITIP20 ISSUER_ROLE + ITIP20RolesAuth).
// Role layout follows TIP20RolesAuth: https://github.com/tempoxyz/tempo/blob/main/tips/ref-impls/src/abstracts/TIP20RolesAuth.sol
// hasRole is a public mapping, so the getter is hasRole(address,bytes32), not OpenZeppelin AccessControl order.
// ISSUER_ROLE: https://github.com/tempoxyz/tempo-std/blob/master/src/interfaces/ITIP20.sol
const TIP20TokenABI = `[{"type":"function","name":"grantRole","inputs":[{"name":"role","type":"bytes32"},{"name":"account","type":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"hasRole","inputs":[{"name":"account","type":"address"},{"name":"role","type":"bytes32"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"},{"type":"function","name":"ISSUER_ROLE","inputs":[],"outputs":[{"name":"","type":"bytes32"}],"stateMutability":"view"}]`

// DefaultAdminRole is TIP20RolesAuth.DEFAULT_ADMIN_ROLE (bytes32(0)).
var DefaultAdminRole [32]byte

// TIP20Token binds calls to an on-chain TIP-20 token (mint/burn issuer role and access checks).
type TIP20Token struct {
	address  common.Address
	contract *bind.BoundContract
}

func NewTIP20Token(address common.Address, backend bind.ContractBackend) (*TIP20Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TIP20TokenABI))
	if err != nil {
		return nil, err
	}
	return &TIP20Token{
		address:  address,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (t *TIP20Token) Address() common.Address {
	return t.address
}

func (t *TIP20Token) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "grantRole", role, account)
}

func (t *TIP20Token) HasRole(opts *bind.CallOpts, account common.Address, role [32]byte) (bool, error) {
	var out []any
	err := t.contract.Call(opts, &out, "hasRole", account, role)
	if err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

func (t *TIP20Token) ISSUERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []any
	err := t.contract.Call(opts, &out, "ISSUER_ROLE")
	if err != nil {
		return [32]byte{}, err
	}
	return *abi.ConvertType(out[0], new([32]byte)).(*[32]byte), nil
}
