package weth

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/weth9"
)

var ContractType cldf_deployment.ContractType = "WETH"

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "weth:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys the WETH9 contract",
	ContractMetadata: weth9.WETH9MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(weth9.WETH9Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

// =============================================================================
// WETH Withdraw Operation (Unwrap WETH to Native ETH)
// =============================================================================

// WithdrawInput specifies the amount of WETH to unwrap to native ETH.
// The native ETH will be sent to msg.sender (the MCMS timelock).
type WithdrawInput struct {
	Amount *big.Int // Amount in wei to unwrap
}

// Withdraw unwraps WETH to native ETH. The ETH is sent to msg.sender.
// This operation should be called by the MCMS timelock after receiving WETH
// from fee withdrawals.
var Withdraw = contract.NewWrite(contract.WriteParams[WithdrawInput, *weth9.WETH9]{
	Name:            "weth:withdraw",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Unwraps WETH to native ETH, sending ETH to msg.sender",
	ContractType:    ContractType,
	ContractABI:     weth9.WETH9ABI,
	NewContract:     weth9.NewWETH9,
	// Always return false - withdraw is meant to be called by MCMS timelock as part of
	// atomic batch operations (e.g., sweep-and-unwrap). The timelock receives WETH from
	// earlier transactions in the batch, so we can't execute directly with deployer key.
	IsAllowedCaller: func(_ *weth9.WETH9, _ *bind.CallOpts, _ common.Address, _ WithdrawInput) (bool, error) { return false, nil },
	Validate:        func(args WithdrawInput) error { return nil },
	CallContract: func(weth *weth9.WETH9, opts *bind.TransactOpts, args WithdrawInput) (*types.Transaction, error) {
		return weth.Withdraw(opts, args.Amount)
	},
})

// =============================================================================
// WETH Balance Read Operation
// =============================================================================

// BalanceOf reads the WETH balance of an account
var BalanceOf = contract.NewRead(contract.ReadParams[common.Address, *big.Int, *weth9.WETH9]{
	Name:         "weth:balance-of",
	Version:      semver.MustParse("1.0.0"),
	Description:  "Reads the WETH balance of an account",
	ContractType: ContractType,
	NewContract:  weth9.NewWETH9,
	CallContract: func(weth *weth9.WETH9, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
		return weth.BalanceOf(opts, account)
	},
})

// =============================================================================
// Native ETH Transfer Helper
// =============================================================================

// CreateNativeETHTransfer creates an MCMS transaction that transfers native ETH.
// This is used after unwrapping WETH to send the native ETH to the treasury.
// The MCMS timelock will execute this as a simple value transfer.
func CreateNativeETHTransfer(chainSelector uint64, to common.Address, amount *big.Int) (mcms_types.BatchOperation, error) {
	if amount == nil || amount.Sign() <= 0 {
		return mcms_types.BatchOperation{}, fmt.Errorf("amount must be positive, got %v", amount)
	}
	if to == (common.Address{}) {
		return mcms_types.BatchOperation{}, fmt.Errorf("recipient address cannot be zero")
	}

	// Create AdditionalFields with the ETH value as a number (not quoted string)
	// MCMS expects {"value": 123} not {"value": "123"}
	additionalFields := json.RawMessage(fmt.Sprintf(`{"value": %s}`, amount.String()))

	return mcms_types.BatchOperation{
		ChainSelector: mcms_types.ChainSelector(chainSelector),
		Transactions: []mcms_types.Transaction{
			{
				OperationMetadata: mcms_types.OperationMetadata{
					ContractType: "NativeETHTransfer",
				},
				To:               to.Hex(),
				Data:             []byte{}, // Empty but non-nil for MCMS validation
				AdditionalFields: additionalFields,
			},
		},
	}, nil
}

// CreateNativeETHTransferTx creates a single Transaction for native ETH transfer.
// Unlike CreateNativeETHTransfer (which returns BatchOperation), this returns a
// Transaction that can be combined with other transactions in a single atomic batch.
// Use this when building multi-step atomic operations.
func CreateNativeETHTransferTx(to common.Address, amount *big.Int) (mcms_types.Transaction, error) {
	if amount == nil || amount.Sign() <= 0 {
		return mcms_types.Transaction{}, fmt.Errorf("amount must be positive, got %v", amount)
	}
	if to == (common.Address{}) {
		return mcms_types.Transaction{}, fmt.Errorf("recipient address cannot be zero")
	}

	// Create AdditionalFields with the ETH value as a number (not quoted string)
	// MCMS expects {"value": 123} not {"value": "123"}
	additionalFields := json.RawMessage(fmt.Sprintf(`{"value": %s}`, amount.String()))

	return mcms_types.Transaction{
		OperationMetadata: mcms_types.OperationMetadata{
			ContractType: "NativeETHTransfer",
		},
		To:               to.Hex(),
		Data:             []byte{}, // Empty but non-nil for MCMS validation
		AdditionalFields: additionalFields,
	}, nil
}
