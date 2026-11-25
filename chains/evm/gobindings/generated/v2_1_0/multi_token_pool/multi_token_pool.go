// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multi_token_pool

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

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead            uint32
	DestBytesOverhead          uint32
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
	IsEnabled                  bool
}

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolLockOrBurnOutV1 struct {
	DestTokenAddress []byte
	DestPoolData     []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender          []byte
	RemoteChainSelector     uint64
	Receiver                common.Address
	SourceDenominatedAmount *big.Int
	LocalToken              common.Address
	SourcePoolAddress       []byte
	SourcePoolData          []byte
	OffchainTokenData       []byte
}

type PoolReleaseOrMintOutV1 struct {
	DestinationAmount *big.Int
}

var MultiTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolConfigManager\",\"inputs\":[],\"outputs\":[{\"name\":\"poolConfigManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"supported\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolConfigManager\",\"inputs\":[{\"name\":\"poolConfigManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolConfigManagerSet\",\"inputs\":[{\"name\":\"oldPoolConfigManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newPoolConfigManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterSet\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotPoolConfigManager\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LaneFeeNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"LocalTokenRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedLockOrBurnFlavor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
}

var MultiTokenPoolABI = MultiTokenPoolMetaData.ABI

type MultiTokenPool struct {
	address common.Address
	abi     abi.ABI
	MultiTokenPoolCaller
	MultiTokenPoolTransactor
	MultiTokenPoolFilterer
}

type MultiTokenPoolCaller struct {
	contract *bind.BoundContract
}

type MultiTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type MultiTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type MultiTokenPoolSession struct {
	Contract     *MultiTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolCallerSession struct {
	Contract *MultiTokenPoolCaller
	CallOpts bind.CallOpts
}

type MultiTokenPoolTransactorSession struct {
	Contract     *MultiTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type MultiTokenPoolRaw struct {
	Contract *MultiTokenPool
}

type MultiTokenPoolCallerRaw struct {
	Contract *MultiTokenPoolCaller
}

type MultiTokenPoolTransactorRaw struct {
	Contract *MultiTokenPoolTransactor
}

func NewMultiTokenPool(address common.Address, backend bind.ContractBackend) (*MultiTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(MultiTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMultiTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPool{address: address, abi: abi, MultiTokenPoolCaller: MultiTokenPoolCaller{contract: contract}, MultiTokenPoolTransactor: MultiTokenPoolTransactor{contract: contract}, MultiTokenPoolFilterer: MultiTokenPoolFilterer{contract: contract}}, nil
}

func NewMultiTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*MultiTokenPoolCaller, error) {
	contract, err := bindMultiTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolCaller{contract: contract}, nil
}

func NewMultiTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTokenPoolTransactor, error) {
	contract, err := bindMultiTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolTransactor{contract: contract}, nil
}

func NewMultiTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiTokenPoolFilterer, error) {
	contract, err := bindMultiTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolFilterer{contract: contract}, nil
}

func bindMultiTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MultiTokenPool *MultiTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPool.Contract.MultiTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPool *MultiTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.MultiTokenPoolTransactor.contract.Transfer(opts)
}

func (_MultiTokenPool *MultiTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.MultiTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_MultiTokenPool *MultiTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_MultiTokenPool *MultiTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.contract.Transfer(opts)
}

func (_MultiTokenPool *MultiTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getFee", localToken, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.IsEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetFee(localToken common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _MultiTokenPool.Contract.GetFee(&_MultiTokenPool.CallOpts, localToken, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetFee(localToken common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _MultiTokenPool.Contract.GetFee(&_MultiTokenPool.CallOpts, localToken, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetPoolConfigManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getPoolConfigManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetPoolConfigManager() (common.Address, error) {
	return _MultiTokenPool.Contract.GetPoolConfigManager(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetPoolConfigManager() (common.Address, error) {
	return _MultiTokenPool.Contract.GetPoolConfigManager(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getRemoteToken", localToken, remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetRemoteToken(localToken common.Address, remoteChainSelector uint64) ([]byte, error) {
	return _MultiTokenPool.Contract.GetRemoteToken(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetRemoteToken(localToken common.Address, remoteChainSelector uint64) ([]byte, error) {
	return _MultiTokenPool.Contract.GetRemoteToken(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetRemoteToken0(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getRemoteToken0", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetRemoteToken0(arg0 uint64) ([]byte, error) {
	return _MultiTokenPool.Contract.GetRemoteToken0(&_MultiTokenPool.CallOpts, arg0)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetRemoteToken0(arg0 uint64) ([]byte, error) {
	return _MultiTokenPool.Contract.GetRemoteToken0(&_MultiTokenPool.CallOpts, arg0)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, arg2 *big.Int, requestedFinalityConfig [4]byte, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, arg2, requestedFinalityConfig, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, arg2 *big.Int, requestedFinalityConfig [4]byte, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _MultiTokenPool.Contract.GetRequiredCCVs(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector, arg2, requestedFinalityConfig, arg4, direction)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, arg2 *big.Int, requestedFinalityConfig [4]byte, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _MultiTokenPool.Contract.GetRequiredCCVs(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector, arg2, requestedFinalityConfig, arg4, direction)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _MultiTokenPool.Contract.GetRmnProxy(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _MultiTokenPool.Contract.GetRmnProxy(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetRouter() (common.Address, error) {
	return _MultiTokenPool.Contract.GetRouter(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _MultiTokenPool.Contract.GetRouter(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts, localToken common.Address) (uint8, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getTokenDecimals", localToken)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetTokenDecimals(localToken common.Address) (uint8, error) {
	return _MultiTokenPool.Contract.GetTokenDecimals(&_MultiTokenPool.CallOpts, localToken)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetTokenDecimals(localToken common.Address) (uint8, error) {
	return _MultiTokenPool.Contract.GetTokenDecimals(&_MultiTokenPool.CallOpts, localToken)
}

func (_MultiTokenPool *MultiTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", localToken, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MultiTokenPool.Contract.GetTokenTransferFeeConfig(&_MultiTokenPool.CallOpts, localToken, destChainSelector, arg2, arg3)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MultiTokenPool.Contract.GetTokenTransferFeeConfig(&_MultiTokenPool.CallOpts, localToken, destChainSelector, arg2, arg3)
}

func (_MultiTokenPool *MultiTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "isSupportedChain", localToken, remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) IsSupportedChain(localToken common.Address, remoteChainSelector uint64) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedChain(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) IsSupportedChain(localToken common.Address, remoteChainSelector uint64) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedChain(&_MultiTokenPool.CallOpts, localToken, remoteChainSelector)
}

func (_MultiTokenPool *MultiTokenPoolCaller) IsSupportedChain0(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "isSupportedChain0", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) IsSupportedChain0(arg0 uint64) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedChain0(&_MultiTokenPool.CallOpts, arg0)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) IsSupportedChain0(arg0 uint64) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedChain0(&_MultiTokenPool.CallOpts, arg0)
}

func (_MultiTokenPool *MultiTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedToken(&_MultiTokenPool.CallOpts, token)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MultiTokenPool.Contract.IsSupportedToken(&_MultiTokenPool.CallOpts, token)
}

func (_MultiTokenPool *MultiTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) Owner() (common.Address, error) {
	return _MultiTokenPool.Contract.Owner(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) Owner() (common.Address, error) {
	return _MultiTokenPool.Contract.Owner(&_MultiTokenPool.CallOpts)
}

func (_MultiTokenPool *MultiTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MultiTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MultiTokenPool *MultiTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MultiTokenPool.Contract.SupportsInterface(&_MultiTokenPool.CallOpts, interfaceId)
}

func (_MultiTokenPool *MultiTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MultiTokenPool.Contract.SupportsInterface(&_MultiTokenPool.CallOpts, interfaceId)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_MultiTokenPool *MultiTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPool.Contract.AcceptOwnership(&_MultiTokenPool.TransactOpts)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MultiTokenPool.Contract.AcceptOwnership(&_MultiTokenPool.TransactOpts)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_MultiTokenPool *MultiTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.LockOrBurn(&_MultiTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.LockOrBurn(&_MultiTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_MultiTokenPool *MultiTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.LockOrBurn0(&_MultiTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.LockOrBurn0(&_MultiTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn, requestedFinalityConfig)
}

func (_MultiTokenPool *MultiTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.ReleaseOrMint(&_MultiTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.ReleaseOrMint(&_MultiTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn)
}

func (_MultiTokenPool *MultiTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.ReleaseOrMint0(&_MultiTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.ReleaseOrMint0(&_MultiTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) SetPoolConfigManager(opts *bind.TransactOpts, poolConfigManager common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "setPoolConfigManager", poolConfigManager)
}

func (_MultiTokenPool *MultiTokenPoolSession) SetPoolConfigManager(poolConfigManager common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.SetPoolConfigManager(&_MultiTokenPool.TransactOpts, poolConfigManager)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) SetPoolConfigManager(poolConfigManager common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.SetPoolConfigManager(&_MultiTokenPool.TransactOpts, poolConfigManager)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "setRouter", router)
}

func (_MultiTokenPool *MultiTokenPoolSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.SetRouter(&_MultiTokenPool.TransactOpts, router)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.SetRouter(&_MultiTokenPool.TransactOpts, router)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_MultiTokenPool *MultiTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.TransferOwnership(&_MultiTokenPool.TransactOpts, to)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.TransferOwnership(&_MultiTokenPool.TransactOpts, to)
}

func (_MultiTokenPool *MultiTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_MultiTokenPool *MultiTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.WithdrawFeeTokens(&_MultiTokenPool.TransactOpts, feeTokens, recipient)
}

func (_MultiTokenPool *MultiTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MultiTokenPool.Contract.WithdrawFeeTokens(&_MultiTokenPool.TransactOpts, feeTokens, recipient)
}

type MultiTokenPoolFeeTokenWithdrawnIterator struct {
	Event *MultiTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(MultiTokenPoolFeeTokenWithdrawn)
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

func (it *MultiTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*MultiTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolFeeTokenWithdrawnIterator{contract: _MultiTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolFeeTokenWithdrawn)
				if err := _MultiTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*MultiTokenPoolFeeTokenWithdrawn, error) {
	event := new(MultiTokenPoolFeeTokenWithdrawn)
	if err := _MultiTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolLockedOrBurnedIterator struct {
	Event *MultiTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolLockedOrBurned)
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
		it.Event = new(MultiTokenPoolLockedOrBurned)
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

func (it *MultiTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*MultiTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolLockedOrBurnedIterator{contract: _MultiTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolLockedOrBurned, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolLockedOrBurned)
				if err := _MultiTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*MultiTokenPoolLockedOrBurned, error) {
	event := new(MultiTokenPoolLockedOrBurned)
	if err := _MultiTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolOwnershipTransferRequestedIterator struct {
	Event *MultiTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolOwnershipTransferRequested)
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
		it.Event = new(MultiTokenPoolOwnershipTransferRequested)
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

func (it *MultiTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolOwnershipTransferRequestedIterator{contract: _MultiTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolOwnershipTransferRequested)
				if err := _MultiTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolOwnershipTransferRequested, error) {
	event := new(MultiTokenPoolOwnershipTransferRequested)
	if err := _MultiTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolOwnershipTransferredIterator struct {
	Event *MultiTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolOwnershipTransferred)
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
		it.Event = new(MultiTokenPoolOwnershipTransferred)
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

func (it *MultiTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolOwnershipTransferredIterator{contract: _MultiTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolOwnershipTransferred)
				if err := _MultiTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolOwnershipTransferred, error) {
	event := new(MultiTokenPoolOwnershipTransferred)
	if err := _MultiTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolPoolConfigManagerSetIterator struct {
	Event *MultiTokenPoolPoolConfigManagerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolPoolConfigManagerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolPoolConfigManagerSet)
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
		it.Event = new(MultiTokenPoolPoolConfigManagerSet)
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

func (it *MultiTokenPoolPoolConfigManagerSetIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolPoolConfigManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolPoolConfigManagerSet struct {
	OldPoolConfigManager common.Address
	NewPoolConfigManager common.Address
	Raw                  types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterPoolConfigManagerSet(opts *bind.FilterOpts) (*MultiTokenPoolPoolConfigManagerSetIterator, error) {

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "PoolConfigManagerSet")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolPoolConfigManagerSetIterator{contract: _MultiTokenPool.contract, event: "PoolConfigManagerSet", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchPoolConfigManagerSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolPoolConfigManagerSet) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "PoolConfigManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolPoolConfigManagerSet)
				if err := _MultiTokenPool.contract.UnpackLog(event, "PoolConfigManagerSet", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParsePoolConfigManagerSet(log types.Log) (*MultiTokenPoolPoolConfigManagerSet, error) {
	event := new(MultiTokenPoolPoolConfigManagerSet)
	if err := _MultiTokenPool.contract.UnpackLog(event, "PoolConfigManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolReleasedOrMintedIterator struct {
	Event *MultiTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolReleasedOrMinted)
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
		it.Event = new(MultiTokenPoolReleasedOrMinted)
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

func (it *MultiTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*MultiTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolReleasedOrMintedIterator{contract: _MultiTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolReleasedOrMinted, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolReleasedOrMinted)
				if err := _MultiTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*MultiTokenPoolReleasedOrMinted, error) {
	event := new(MultiTokenPoolReleasedOrMinted)
	if err := _MultiTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MultiTokenPoolRouterSetIterator struct {
	Event *MultiTokenPoolRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MultiTokenPoolRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTokenPoolRouterSet)
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
		it.Event = new(MultiTokenPoolRouterSet)
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

func (it *MultiTokenPoolRouterSetIterator) Error() error {
	return it.fail
}

func (it *MultiTokenPoolRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MultiTokenPoolRouterSet struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_MultiTokenPool *MultiTokenPoolFilterer) FilterRouterSet(opts *bind.FilterOpts) (*MultiTokenPoolRouterSetIterator, error) {

	logs, sub, err := _MultiTokenPool.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &MultiTokenPoolRouterSetIterator{contract: _MultiTokenPool.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_MultiTokenPool *MultiTokenPoolFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolRouterSet) (event.Subscription, error) {

	logs, sub, err := _MultiTokenPool.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MultiTokenPoolRouterSet)
				if err := _MultiTokenPool.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_MultiTokenPool *MultiTokenPoolFilterer) ParseRouterSet(log types.Log) (*MultiTokenPoolRouterSet, error) {
	event := new(MultiTokenPoolRouterSet)
	if err := _MultiTokenPool.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (MultiTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (MultiTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (MultiTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MultiTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MultiTokenPoolPoolConfigManagerSet) Topic() common.Hash {
	return common.HexToHash("0x68a57a9ddc0e982542db2e451d4bd7a0648cc6a66ba2b10e620c24daa5756808")
}

func (MultiTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (MultiTokenPoolRouterSet) Topic() common.Hash {
	return common.HexToHash("0x8ab4cea230ab04f96aaa9a71fbc4bfa98097c7b9f07e06be33868b500f810028")
}

func (_MultiTokenPool *MultiTokenPool) Address() common.Address {
	return _MultiTokenPool.address
}

type MultiTokenPoolInterface interface {
	GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

		error)

	GetPoolConfigManager(opts *bind.CallOpts) (common.Address, error)

	GetRemoteToken(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) ([]byte, error)

	GetRemoteToken0(opts *bind.CallOpts, arg0 uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, arg2 *big.Int, requestedFinalityConfig [4]byte, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts, localToken common.Address) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsSupportedChain(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) (bool, error)

	IsSupportedChain0(opts *bind.CallOpts, arg0 uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	SetPoolConfigManager(opts *bind.TransactOpts, poolConfigManager common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*MultiTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*MultiTokenPoolFeeTokenWithdrawn, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*MultiTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolLockedOrBurned, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*MultiTokenPoolLockedOrBurned, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MultiTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MultiTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MultiTokenPoolOwnershipTransferred, error)

	FilterPoolConfigManagerSet(opts *bind.FilterOpts) (*MultiTokenPoolPoolConfigManagerSetIterator, error)

	WatchPoolConfigManagerSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolPoolConfigManagerSet) (event.Subscription, error)

	ParsePoolConfigManagerSet(log types.Log) (*MultiTokenPoolPoolConfigManagerSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64, token []common.Address) (*MultiTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolReleasedOrMinted, remoteChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*MultiTokenPoolReleasedOrMinted, error)

	FilterRouterSet(opts *bind.FilterOpts) (*MultiTokenPoolRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *MultiTokenPoolRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*MultiTokenPoolRouterSet, error)

	Address() common.Address
}
