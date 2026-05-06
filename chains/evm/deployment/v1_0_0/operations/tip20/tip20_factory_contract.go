package tip20

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var FactoryContractType deployment.ContractType = "TIP20Factory"

// TIP20FactoryABI is a minimal ABI for the Tempo TIP-20 factory precompile (ITIP20Factory).
const TIP20FactoryABI = `[{"type":"function","name":"createToken","inputs":[{"name":"name","type":"string"},{"name":"symbol","type":"string"},{"name":"currency","type":"string"},{"name":"quoteToken","type":"address"},{"name":"admin","type":"address"},{"name":"salt","type":"bytes32"}],"outputs":[{"name":"","type":"address"}],"stateMutability":"nonpayable"},{"type":"function","name":"getTokenAddress","inputs":[{"name":"sender","type":"address"},{"name":"salt","type":"bytes32"}],"outputs":[{"name":"","type":"address"}],"stateMutability":"pure"},{"type":"function","name":"isTIP20","inputs":[{"name":"token","type":"address"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"}]`

// TIP20Factory binds calls to the on-chain TIP-20 factory precompile.
type TIP20Factory struct {
	address  common.Address
	contract *bind.BoundContract
}

func NewTIP20Factory(address common.Address, backend bind.ContractBackend) (*TIP20Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(TIP20FactoryABI))
	if err != nil {
		return nil, err
	}
	return &TIP20Factory{
		address:  address,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (f *TIP20Factory) Address() common.Address {
	return f.address
}

func (f *TIP20Factory) CreateToken(opts *bind.TransactOpts, name, symbol, currency string, quoteToken, admin common.Address, salt [32]byte) (*types.Transaction, error) {
	return f.contract.Transact(opts, "createToken", name, symbol, currency, quoteToken, admin, salt)
}

func (f *TIP20Factory) GetTokenAddress(opts *bind.CallOpts, sender common.Address, salt [32]byte) (common.Address, error) {
	var out []any
	err := f.contract.Call(opts, &out, "getTokenAddress", sender, salt)
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

func (f *TIP20Factory) IsTIP20(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []any
	err := f.contract.Call(opts, &out, "isTIP20", token)
	if err != nil {
		return false, err
	}
	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

// CreateTokenArgs are the arguments to TIP20Factory.createToken.
type CreateTokenArgs struct {
	Name       string
	Symbol     string
	Currency   string
	QuoteToken common.Address
	Admin      common.Address
	Salt       [32]byte
}

// GetTokenAddressArgs are the arguments to TIP20Factory.getTokenAddress.
type GetTokenAddressArgs struct {
	Sender common.Address
	Salt   [32]byte
}

var CreateToken = contract.NewWrite(contract.WriteParams[CreateTokenArgs, *TIP20Factory]{
	Name:            "tip20-factory:create-token",
	Version:         Version,
	Description:     "Creates a TIP-20 token via the Tempo TIP-20 factory precompile",
	ContractType:    FactoryContractType,
	ContractABI:     TIP20FactoryABI,
	NewContract:     NewTIP20Factory,
	IsAllowedCaller: contract.AllCallersAllowed[*TIP20Factory, CreateTokenArgs],
	Validate: func(args CreateTokenArgs) error {
		if args.Name == "" {
			return errors.New("name is required")
		}
		if args.Symbol == "" {
			return errors.New("symbol is required")
		}
		if args.Admin == (common.Address{}) {
			return errors.New("admin is required")
		}
		return nil
	},
	CallContract: func(f *TIP20Factory, opts *bind.TransactOpts, args CreateTokenArgs) (*types.Transaction, error) {
		return f.CreateToken(opts, args.Name, args.Symbol, args.Currency, args.QuoteToken, args.Admin, args.Salt)
	},
})

var GetTokenAddress = contract.NewRead(contract.ReadParams[GetTokenAddressArgs, common.Address, *TIP20Factory]{
	Name:         "tip20-factory:get-token-address",
	Version:      Version,
	Description:  "Predicts the TIP-20 token address for a sender and salt",
	ContractType: FactoryContractType,
	NewContract:  NewTIP20Factory,
	CallContract: func(f *TIP20Factory, opts *bind.CallOpts, args GetTokenAddressArgs) (common.Address, error) {
		return f.GetTokenAddress(opts, args.Sender, args.Salt)
	},
})

var IsTIP20 = contract.NewRead(contract.ReadParams[common.Address, bool, *TIP20Factory]{
	Name:         "tip20-factory:is-tip20",
	Version:      Version,
	Description:  "Returns whether an address is a valid TIP-20 token",
	ContractType: FactoryContractType,
	NewContract:  NewTIP20Factory,
	CallContract: func(f *TIP20Factory, opts *bind.CallOpts, token common.Address) (bool, error) {
		return f.IsTIP20(opts, token)
	},
})
