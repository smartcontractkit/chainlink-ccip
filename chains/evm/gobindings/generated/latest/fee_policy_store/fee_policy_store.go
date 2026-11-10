// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_policy_store

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

type FeePolicyStoreDestOverheadConfig struct {
	DestGasOverhead   uint32
	DestBytesOverhead uint32
}

type FeePolicyStoreDestOverheadPolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	DestOverheadConfig  FeePolicyStoreDestOverheadConfig
}

type FeePolicyStoreFeeConfig struct {
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
}

type FeePolicyStoreFeePolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	FeeConfig           FeePolicyStoreFeeConfig
}

type FeePolicyStoreInlineDestOverheadConfigArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	DestOverheadConfig  FeePolicyStoreDestOverheadConfig
}

type FeePolicyStoreInlineFeeConfigArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	FeeConfig           FeePolicyStoreFeeConfig
}

type IPoolConfigManagerLaneFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

type IPoolConfigManagerLaneKey struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
}

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead            uint32
	DestBytesOverhead          uint32
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
	IsEnabled                  bool
}

type PolicyConfigPolicyRecordKey struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
}

type PolicyConfigTokenPolicyArgs struct {
	LocalToken common.Address
	PolicyName [32]byte
}

type PolicyConfigTokenProfileArgs struct {
	LocalToken                 common.Address
	LocalDecimals              uint8
	LockOrBurnFlavor           uint8
	AllowedFinalityConfig      [4]byte
	LockOrBurnAddress          common.Address
	AdditionalPoolDataProvider common.Address
}

var FeePolicyStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestOverheadPolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.DestOverheadPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeePolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.FeePolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineDestOverheadConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineDestOverheadConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineFeeConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineFeeConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenDestOverheadPolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenFeePolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenProfileUpdates\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"profilesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenProfileArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkLaneRecords\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestOverheadPolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.DestOverheadPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"fees\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneFee[]\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeePolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.FeePolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineDestOverheadConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineDestOverheadConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineFeeConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineFeeConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneRecordLevelBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"feeLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"destOverheadLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyNames\",\"inputs\":[],\"outputs\":[{\"name\":\"feePolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"destOverheadPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPolicies\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feePolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destOverheadPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPolicyNamesBatch\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"feePolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"destOverheadPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfigBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"feeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig[]\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validateLockOrBurn\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"chargeTokenFee\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"destPoolDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestOverheadPolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestOverheadPolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestOverheadPolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeePolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeePolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeePolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineDestOverheadConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineDestOverheadConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineFeeConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineFeeConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDestOverheadPolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenFeePolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenProfileSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestOverheadRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InlineRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDestOverheadPolicyConfig\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"type\":\"error\",\"name\":\"InvalidFeePolicyConfig\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"type\":\"error\",\"name\":\"InvalidInlineDestOverheadConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"type\":\"error\",\"name\":\"InvalidInlineFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"type\":\"error\",\"name\":\"LocalTokenNotRegistered\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyHasNoTokens\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyHasRecords\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyInUse\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyRecordNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PoolDataTooLarge\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destPoolDataLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenHasConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPolicyNameNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60806040523461020f57613e998038038061001981610214565b92833981019060208183031261020f578051906001600160401b03821161020f570181601f8201121561020f578051916001600160401b0383116101c8578260051b9160208061006a818601610214565b80968152019382010191821161020f57602001915b8183106101ef578333156101de57600180546001600160a01b031916331790556020906100ab82610214565b60008152600036813760408051929083016001600160401b038111848210176101c8576040528252808383015260005b8151811015610142576001906001600160a01b036100f98285610239565b5116856101058261027b565b610112575b5050016100db565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1858561010a565b50505160005b81518110156101b9576001600160a01b036101638284610239565b51169081156101a8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef848361019a600195610379565b50604051908152a101610148565b6342bcdf7f60e11b60005260046000fd5b604051613abf90816103da8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b038116810361020f5781526020928301920161007f565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101c857604052565b805182101561024d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561024d5760005260206000200190600090565b600081815260036020526040902054801561037257600019810181811161035c5760025460001981019190821161035c5780820361030b575b50505060025480156102f557600019016102cf816002610263565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61034461031c61032d936002610263565b90549060031b1c9283926002610263565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102b4565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103d357600254680100000000000000008110156101c8576103ba61032d8260018594016002556002610263565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c80630a2ad9a9146125745780630f1e1f5614612335578063181f5a77146122815780632451a627146121cd57806334dbc50014611fb957806339891a3e14611d8f578063412e593e14611bda57806351b910ca1461197a5780636389dc0d1461190757806379ba50971461183c5780637a338ed3146116865780637b5f0d5f146115955780638c63fe62146114dc5780638da5cb5b146114a857806391a2749a146112fc578063b1a26f771461114e578063bba4ef5014610e5a578063d37eeeed14610abb578063d8eb845f14610a8f578063de2425231461066d578063ec8784c114610582578063edf5e2c314610399578063eec584ea14610322578063efbe103e146102065763f2fde38b1461012f57600080fd5b346102015760206003193601126102015773ffffffffffffffffffffffffffffffffffffffff61015d61271a565b61016561316f565b163381146101d757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102015760206003193601126102015760043567ffffffffffffffff8111610201576102379036906004016126e9565b90610240612d52565b6000915b80831061024d57005b61026061025b848385612a58565b612a13565b9273ffffffffffffffffffffffffffffffffffffffff6020610283838587612a58565b013594169361029f856000526005602052604060002054151590565b156102f4578460019394956000526006602052604060002090828254926102c68285613470565b557f653518802ef88beefe781cd891f284259493bcf3231e3ab39f4b5a555a02afae600080a4019190610244565b847fdfcc7fc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102015760c06003193601126102015761033b61271a565b61034361296e565b90606435917fffffffff0000000000000000000000000000000000000000000000000000000083168303610201576084358015158103610201576020936103919360a4359360443591612c69565b604051908152f35b346102015760206003193601126102015773ffffffffffffffffffffffffffffffffffffffff6103c761271a565b168060005260076020526103de604060002061394a565b8051917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061042461040e856128db565b9461041c604051968761289a565b8086526128db565b0160005b81811061056b57505060005b82518110156104b6578067ffffffffffffffff61045360019386612a34565b511683600052600860205261049560408060002060009067ffffffffffffffff85168252602052206040519261048884612862565b8684526020840152612b71565b60408201526104a48287612a34565b526104af8186612a34565b5001610434565b8360405180916020820160208352815180915260206040840192019060005b8181106104e3575050500390f35b91935091602060c060019261055d6040885173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff8682015116868501520151604083019061ffff6060809263ffffffff815116855263ffffffff6020820151166020860152826040820151166040860152015116910152565b0194019101918493926104d5565b602090610576612b2e565b82828801015201610428565b346102015760206003193601126102015760043567ffffffffffffffff8111610201576105b39036906004016126e9565b906105bc612d52565b6000915b8083106105c957005b6105d761025b848385612a58565b9273ffffffffffffffffffffffffffffffffffffffff60206105fa838587612a58565b0135941693610616856000526005602052604060002054151590565b156102f457846001939495600052600660205283604060002001908282549261063f82856132d1565b557f7bd0030593fdd95573b972204541f9eef9cedfa1daab494d1dc9838d11cfdba4600080a40191906105c0565b346102015760606003193601126102015760043567ffffffffffffffff81116102015761069e9036906004016126e9565b60243567ffffffffffffffff8111610201576106be90369060040161262b565b60443567ffffffffffffffff8111610201576106de90369060040161273d565b9490936106e9612d52565b60005b8181106109d05750505060005b8181106108cc5750505060005b82811061070f57005b61071a818484612ad7565b3590610732602061072c838787612ad7565b01612ac2565b61073d828686612ad7565b926040840161074b81612e5a565b1561087e5781156108545760407fc145aeb57aa1bc23c7a23a7b640e48557aa6ff73c43025727740341561d9afac9161084160019761078b86600f613913565b508560005260116020526107ae67ffffffffffffffff8560002098168098613913565b50856000526012602052610809606085806000206000908b8252602052209263ffffffff6107db87612b1d565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000085541617845501612b1d565b7fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff0000000083549260201b169116179055565b61084d82518092612af8565ba301610706565b7ff292e50c0000000000000000000000000000000000000000000000000000000060005260046000fd5b6108ca9067ffffffffffffffff608494604051947f930759f200000000000000000000000000000000000000000000000000000000865260048601521660248401526044830190612af8565bfd5b6108d78183856129d4565b356040600082815260106020522054156109a357806000526011602052604060002063ffffffff600282015416610975575461094857908161091c600193600f613794565b507fdf743435d049ae540ec6d143d4d295658be6883eaedf0962d1311cf5da2610d5600080a2016106f9565b7fa57cd4ec0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3c95cbc10000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7ff90d6e6b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6109db818385612a58565b356109ec602061072c848688612a58565b90806000526011602052610a1067ffffffffffffffff604060002093168093613794565b15610a5f57908160019392600052601260205260006040808220829085825260205220557ff23ee1ed541c9a6fc556979fa67ee730c8b973567f320dfab6b0f08fa98b640d600080a3016106ec565b7f779848370000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b3461020157604060031936011261020157610ab9610aab61271a565b610ab361296e565b906131f1565b005b346102015760606003193601126102015760043567ffffffffffffffff811161020157610aec9036906004016126e9565b60243567ffffffffffffffff811161020157610b0c90369060040161262b565b60443567ffffffffffffffff811161020157610b2c9036906004016126b8565b949093610b37612d52565b60005b818110610dcb5750505060005b818110610d4f5750505060005b828110610b5d57005b610b68818484612a48565b3590610b7a602061072c838787612a48565b610b85828686612a48565b9260408401610b93816131ba565b15610d035781156108545760807f3275d66dec24a74d993e8775d56a0a359d9f1949cdbf98c1eb4c524728bf360f91600196610bd085600b613913565b5084600052600d602052610bf467ffffffffffffffff604060002097168097613913565b5084600052600e602052604080600020600090888252602052209063ffffffff610c1d84612b1d565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000835416178255610c8d610c5460608301612b1d565b83547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178355565b610c98848201612c18565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff69ffff00000000000000006bffff00000000000000000000610ce060a087549501612c18565b60501b169360401b16911617179055610cfc6040518092612bcd565ba301610b54565b6108ca9067ffffffffffffffff60c494604051947f9804ee1200000000000000000000000000000000000000000000000000000000865260048601521660248401526044830190612bcd565b610d5a8183856129d4565b3560406000828152600c6020522054156109a35780600052600d602052604060002063ffffffff6002820154166109755754610948579081610d9f600193600b613794565b507f6d8754acd53d70e89e08ddb26fe365e6a02f960df734251862d4d69b036a261a600080a201610b47565b610dd6818385612a58565b35610de7602061072c848688612a58565b9080600052600d602052610e0b67ffffffffffffffff604060002093168093613794565b15610a5f57908160019392600052600e60205260006040808220829085825260205220557faaa43a26411a6c040ef0c507c8f37898e4f20ed5da1f7e361a09dded66360e25600080a301610b3a565b346102015760406003193601126102015760043567ffffffffffffffff811161020157610e8b9036906004016126e9565b60243567ffffffffffffffff811161020157610eab9036906004016126b8565b919092610eb6612d52565b60005b8281106110ac57600085855b808310610ece57005b610edc61025b848385612a48565b92610eed602061072c838587612a48565b93610ef9828486612a48565b906040820191610f08836131ba565b1561104957608067ffffffffffffffff887fcd1632808d4a131200cb019663e9454426bc07617b7b2995874f7ca33e5af16293610f6e6001999a9b73ffffffffffffffffffffffffffffffffffffffff8816978860005260076020526040600020612e8f565b8560005260086020526040806000206000908585168252602052209063ffffffff610f9889612b1d565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000835416178255610fcf610c5460608301612b1d565b610fda858201612c18565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff69ffff00000000000000006bffff0000000000000000000061102260a087549501612c18565b60501b169360401b1691161717905561103e6040518097612bcd565b1693a3019190610ec5565b506108ca60c49267ffffffffffffffff8873ffffffffffffffffffffffffffffffffffffffff604051957fffe1fed70000000000000000000000000000000000000000000000000000000087521660048601521660248401526044830190612bcd565b806110bd61025b6001938686612a58565b67ffffffffffffffff6110d6602061072c858989612a58565b6111058173ffffffffffffffffffffffffffffffffffffffff8516948560005260076020526040600020612df4565b8260005260086020526000604080822082908585168252602052205516907f3c4fb3bc77a699c9e93a8ae0aaa4517719eac2360b1432d8c84edcd3b061b06e600080a301610eb9565b346102015760206003193601126102015760043580600052601160205260406000206111798161394a565b918251927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06111c06111aa866128db565b956111b8604051978861289a565b8087526128db565b0160005b8181106112e557505060005b8151811015611252578067ffffffffffffffff6111ef60019385612a34565b511684600052601260205261123160408060002060009067ffffffffffffffff85168252602052206040519261122484612862565b8784526020840152612a9e565b60408201526112408288612a34565b5261124b8187612a34565b50016111d0565b8463ffffffff6002860154166040519182916040830190835260406020840152815180915260206060840192019060005b818110611291575050500390f35b91935091602060806001926112d7604088518051845267ffffffffffffffff8682015116868501520151604083019063ffffffff60208092828151168552015116910152565b019401910191849392611283565b6020906112f0612a68565b828289010152016111c4565b346102015760206003193601126102015760043567ffffffffffffffff81116102015760406003198236030112610201576040519061133a826127df565b806004013567ffffffffffffffff81116102015761135e90600436918401016128f3565b825260248101359067ffffffffffffffff821161020157600461138492369201016128f3565b6020820190815261139361316f565b519060005b825181101561140b578073ffffffffffffffffffffffffffffffffffffffff6113c360019386612a34565b51166113ce816139e2565b6113da575b5001611398565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1846113d3565b505160005b8151811015610ab95773ffffffffffffffffffffffffffffffffffffffff6114388284612a34565b511690811561147e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836114706001956138da565b50604051908152a101611410565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020157600060031936011261020157602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461020157600060031936011261020157604051600b548082528160208101600b60005260206000209260005b81811061157c57505061151e9250038261289a565b60405190600f548083528260208101600f60005260206000209260005b8181106115635750506115509250038361289a565b61155f60405192839283612690565b0390f35b845483526001948501948794506020909301920161153b565b8454835260019485019486945060209093019201611509565b346102015760206003193601126102015760043567ffffffffffffffff8111610201576115c69036906004016126e9565b6115cf81612985565b906115d981612985565b9260005b82811061160e576116008461155f8760405193849360408552604085019061276e565b90838203602085015261276e565b8061164b61163b61162561025b6001958888612a58565b611635602061072c868a8a612a58565b90612f02565b506116468388612a34565b612bb2565b61168061167561165f61025b848888612a58565b61166f602061072c868a8a612a58565b90613037565b506116468389612a34565b016115dd565b346102015760206003193601126102015760043580600052600d60205260406000206116b18161394a565b918251927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116e26111aa866128db565b0160005b81811061182557505060005b8151811015611774578067ffffffffffffffff61171160019385612a34565b511684600052600e60205261175360408060002060009067ffffffffffffffff85168252602052206040519261174684612862565b8784526020840152612b71565b60408201526117628288612a34565b5261176d8187612a34565b50016116f2565b8463ffffffff6002860154166040519182916040830190835260406020840152815180915260206060840192019060005b8181106117b3575050500390f35b91935091602060c0600192611817604088518051845267ffffffffffffffff8682015116868501520151604083019061ffff6060809263ffffffff815116855263ffffffff6020820151166020860152826040820151166040860152015116910152565b0194019101918493926117a5565b602090611830612b2e565b828289010152016116e6565b346102015760006003193601126102015760005473ffffffffffffffffffffffffffffffffffffffff811633036118dd577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020157602060031936011261020157606073ffffffffffffffffffffffffffffffffffffffff61193761271a565b1680600052600660205261195d6040600020916000526005602052604060002054151590565b906001815491015490604051921515835260208301526040820152f35b346102015760406003193601126102015760043567ffffffffffffffff8111610201576119ab9036906004016126e9565b60243567ffffffffffffffff8111610201576119cb90369060040161273d565b9190926119d6612d52565b60005b828110611b3857600085855b8083106119ee57005b6119fc61025b848385612ad7565b92611a0d602061072c838587612ad7565b93611a19828486612ad7565b906040820191611a2883612e5a565b15611ad557604067ffffffffffffffff88611abe7f2d6db97987024f726ece930c43a48619971d6d1eb8dcdf871f53c3f9f30b87a994611a9060019a9b9c73ffffffffffffffffffffffffffffffffffffffff89169889600052600960205287600020612e8f565b86600052600a602052610809606086806000206000908888168252602052209263ffffffff6107db8c612b1d565b611aca83518097612af8565b1693a30191906119e5565b506108ca60849267ffffffffffffffff8873ffffffffffffffffffffffffffffffffffffffff604051957f5f8267910000000000000000000000000000000000000000000000000000000087521660048601521660248401526044830190612af8565b80611b4961025b6001938686612a58565b67ffffffffffffffff611b62602061072c858989612a58565b611b918173ffffffffffffffffffffffffffffffffffffffff8516948560005260096020526040600020612df4565b82600052600a6020526000604080822082908585168252602052205516907f5a79db4177d18f5f1b88709312519a5866f9292ee3621e520c4fd46112988d01600080a3016119d9565b346102015760206003193601126102015773ffffffffffffffffffffffffffffffffffffffff611c0861271a565b16806000526009602052611c1f604060002061394a565b8051917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611c4f61040e856128db565b0160005b818110611d7857505060005b8251811015611ce1578067ffffffffffffffff611c7e60019386612a34565b511683600052600a602052611cc060408060002060009067ffffffffffffffff851682526020522060405192611cb384612862565b8684526020840152612a9e565b6040820152611ccf8287612a34565b52611cda8186612a34565b5001611c5f565b8360405180916020820160208352815180915260206040840192019060005b818110611d0e575050500390f35b9193509160206080600192611d6a6040885173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff8682015116868501520151604083019063ffffffff60208092828151168552015116910152565b019401910191849392611d00565b602090611d83612a68565b82828801015201611c53565b346102015760406003193601126102015760043567ffffffffffffffff811161020157611dc09036906004016126e9565b906024357fffffffff000000000000000000000000000000000000000000000000000000008116809103610201571515611df9836128db565b92611e07604051948561289a565b8084527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611e34826128db565b0160005b818110611f7f57505060005b818110611ece578460405180916020820160208352815180915260206040840192019060005b818110611e78575050500390f35b91935091602060a0600192608087518051835263ffffffff85820151168584015263ffffffff604082015116604084015261ffff6060820151166060840152015115156080820152019401910191849392611e6a565b80611ee4611edf6001938588612a58565b612d96565b9115611f785761ffff908615611f685763ffffffff8281835460201c16945b54928915611f5d575460501c16915b8160405195611f2087612846565b168552818116602086015260201c166040840152166060820152826080820152611f4a8288612a34565b52611f558187612a34565b505b01611e44565b5460401c1691611f12565b63ffffffff828183541694611f03565b5050611f57565b602090604051611f8e81612846565b6000815260008382015260006040820152600060608201526000608082015282828901015201611e38565b346102015760206003193601126102015760043567ffffffffffffffff811161020157611fea9036906004016126e9565b90611ff4826128db565b91612002604051938461289a565b8083527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061202f826128db565b0160005b81811061218557505060005b8181106120ed578360405180916020820160208352815180915260206040840192019060005b818110612073575050500390f35b91935091602060e060019260c0875163ffffffff815116835263ffffffff85820151168584015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152019401910191849392612065565b806120fe611edf6001938587612a58565b90911561217e5761ffff9054915463ffffffff6040519361211e8561282a565b818116855260201c16602084015263ffffffff8116604084015263ffffffff8160201c166060840152818160401c16608084015260501c1660a08201528260c082015261216b8287612a34565b526121768186612a34565b505b0161203f565b5050612178565b6020906040516121948161282a565b60008152600083820152600060408201526000606082015260006080820152600060a0820152600060c082015282828801015201612033565b346102015760006003193601126102015760405160025490818152602081018092600260005260206000209060005b81811061226b575050508161221291038261289a565b6040519182916020830190602084525180915260408301919060005b81811061223c575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161222e565b82548452602090930192600192830192016121fc565b346102015760006003193601126102015760405161229e816127df565b601881527f466565506f6c69637953746f726520322e312e302d6465760000000000000000602082015260405190602082528181519182602083015260005b83811061231d5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b602082820181015160408784010152859350016122dd565b346102015760406003193601126102015760043567ffffffffffffffff81116102015761236690369060040161262b565b9060243567ffffffffffffffff8111610201576123879036906004016126b8565b919092612392612d52565b60005b8181106124385750505060005b8181106123ab57005b73ffffffffffffffffffffffffffffffffffffffff6123ce61025b838587612a48565b1690811561240e57816123e260019361389b565b507f936075da10ac5036063e3a526ca53d09805e57da23aa8c124651eda754b4ec13600080a2016123a2565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff61245b61025b8385876129d4565b16612473816000526005602052604060002054151590565b1561254757806000526006602052604060002080541590811591612539575b508015612522575b801561250b575b6124de5790816124b2600193613688565b507f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3600080a201612395565b7f41025a5e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600960205260406000205415156124a1565b50806000526007602052604060002054151561249a565b600191500154151587612492565b7fdfcc7fc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102015760206003193601126102015760043567ffffffffffffffff8111610201576125a590369060040161262b565b906125af82612985565b906125b983612985565b9260005b8181106125d65750505061155f60405192839283612690565b8073ffffffffffffffffffffffffffffffffffffffff6125fc61025b60019486886129d4565b16600052600660205281604060002080546126178489612a34565b5201546126248288612a34565b52016125bd565b9181601f840112156102015782359167ffffffffffffffff8311610201576020808501948460051b01011161020157565b906020808351928381520192019060005b81811061267a5750505090565b825184526020938401939092019160010161266d565b90916126a76126b59360408452604084019061265c565b91602081840391015261265c565b90565b9181601f840112156102015782359167ffffffffffffffff83116102015760208085019460c0850201011161020157565b9181601f840112156102015782359167ffffffffffffffff8311610201576020808501948460061b01011161020157565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361020157565b9181601f840112156102015782359167ffffffffffffffff8311610201576020808501948460071b01011161020157565b90602080835192838152019201906000905b80821061278d5750505090565b90919283519060048210156127b057602081600193829352019401920190612780565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6040810190811067ffffffffffffffff8211176127fb57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176127fb57604052565b60a0810190811067ffffffffffffffff8211176127fb57604052565b6060810190811067ffffffffffffffff8211176127fb57604052565b6080810190811067ffffffffffffffff8211176127fb57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176127fb57604052565b67ffffffffffffffff81116127fb5760051b60200190565b9080601f830112156102015781359061290b826128db565b92612919604051948561289a565b82845260208085019360051b82010191821161020157602001915b8183106129415750505090565b823573ffffffffffffffffffffffffffffffffffffffff8116810361020157815260209283019201612934565b6024359067ffffffffffffffff8216820361020157565b9061298f826128db565b61299c604051918261289a565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129ca82946128db565b0190602036910137565b91908110156129e45760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036102015790565b80518210156129e45760209160051b010190565b91908110156129e45760c0020190565b91908110156129e45760061b0190565b60405190612a7582612862565b8160008152600060208201526040805191612a8f836127df565b60008352600060208401520152565b90604051612aab816127df565b602063ffffffff8294548181168452821c16910152565b3567ffffffffffffffff811681036102015790565b91908110156129e45760071b0190565b359063ffffffff8216820361020157565b63ffffffff612b176020809383612b0e82612ae7565b16865201612ae7565b16910152565b3563ffffffff811681036102015790565b60405190612b3b82612862565b8160008152600060208201526040805191612b558361287e565b6000835260006020840152600082840152600060608401520152565b90604051612b7e8161287e565b606061ffff82945463ffffffff8116845263ffffffff8160201c166020850152818160401c16604085015260501c16910152565b60048210156127b05752565b359061ffff8216820361020157565b61ffff612b176060809363ffffffff612be582612ae7565b16865263ffffffff612bf960208301612ae7565b16602087015283612c0c60408301612bbe565b16604087015201612bbe565b3561ffff811681036102015790565b81810292918115918404141715612c3a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9063ffffffff94929395612c7d82846131f1565b9690965460201c1691828211612cf0575050505015612ce8577fffffffff000000000000000000000000000000000000000000000000000000006127109316612cd65761ffff612cd2925460401c1690612c27565b0490565b61ffff612cd2925460501c1690612c27565b505050600090565b6084945067ffffffffffffffff9073ffffffffffffffffffffffffffffffffffffffff604051957fb7a24c3000000000000000000000000000000000000000000000000000000000875216600486015216602484015260448301526064820152fd5b33600052600360205260406000205415612d6857565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b612dc291612dc8612da683612a13565b61166f612dba602086019261163584612ac2565b969095612a13565b91612ac2565b9260048110156127b05715159081612de1575b50929190565b905060048110156127b057151538612ddb565b9167ffffffffffffffff612e0a91168093613794565b15612e13575050565b73ffffffffffffffffffffffffffffffffffffffff907fa0febeae000000000000000000000000000000000000000000000000000000006000521660045260245260446000fd5b63ffffffff612e6882612b1d565b1615159081612e75575090565b60209150612e898263ffffffff9201612b1d565b16101590565b919073ffffffffffffffffffffffffffffffffffffffff1680600052600560205260406000205415612547575067ffffffffffffffff168015612ed857612ed591613913565b50565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff909291921691826000526007602052612f4f604060002067ffffffffffffffff83169060019160005201602052604060002054151590565b61300b578260005260066020526040600020549283600052600d602052612f7a826040600020613997565b9360048510156127b05760028514612fe35760048510156127b05760038514612fc45750600052600860205267ffffffffffffffff60406000209116600052602052604060002090565b915050600052600e602052604060002060008052602052604060002090565b9050600052600e60205267ffffffffffffffff60406000209116600052602052604060002090565b9091600052600860205267ffffffffffffffff6040600020911660005260205260406000209060019190565b73ffffffffffffffffffffffffffffffffffffffff909291921691826000526009602052613084604060002067ffffffffffffffff83169060019160005201602052604060002054151590565b61314357826000526006602052600160406000200154928360005260116020526130b2826040600020613997565b9360048510156127b0576002851461311b5760048510156127b057600385146130fc5750600052600a60205267ffffffffffffffff60406000209116600052602052604060002090565b9150506000526012602052604060002060008052602052604060002090565b9050600052601260205267ffffffffffffffff60406000209116600052602052604060002090565b9091600052600a60205267ffffffffffffffff6040600020911660005260205260406000209060019190565b73ffffffffffffffffffffffffffffffffffffffff60015416330361319057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b61271061ffff6131cc60408401612c18565b161090816131d8575090565b61271091506131ec606061ffff9201612c18565b161090565b91906131fd8184612f02565b919060048110156127b0571561327f576132178185613037565b949060048110156127b0571561322d5750509190565b9073ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff927fa9aac50f00000000000000000000000000000000000000000000000000000000600052166004521660245260446000fd5b67ffffffffffffffff9073ffffffffffffffffffffffffffffffffffffffff857f32679cbe00000000000000000000000000000000000000000000000000000000600052166004521660245260446000fd5b81811461346c57801515908161340f575b8215159182806133f9575b6133cb57613353575b506132fe5750565b60005260116020526002604060002001600163ffffffff8254160163ffffffff8111612c3a5763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b600052601160205260026040600020017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8254160163ffffffff8111612c3a5763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055386132f6565b837ff90d6e6b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50836000526010602052604060002054156132ed565b806000526010602052604060002054156109a35780600052601160205263ffffffff600260406000200154166132e2577f5c32f7090000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5050565b81811461346c578015159081613580575b82151591828061356a575b6133cb576134f2575b5061349d5750565b600052600d6020526002604060002001600163ffffffff8254160163ffffffff8111612c3a5763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b600052600d60205260026040600020017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8254160163ffffffff8111612c3a5763ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000082541617905538613495565b5083600052600c6020526040600020541561348c565b80600052600c602052604060002054156109a35780600052600d60205263ffffffff60026040600020015416613481577f5c32f7090000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156129e45760005260206000200190600090565b80548015613659577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061362a82826135dd565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260056020526040902054801561378d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111612c3a57600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211612c3a5781810361371e575b50505061370a60046135f5565b600052600560205260006040812055600190565b61377561372f6137409360046135dd565b90549060031b1c92839260046135dd565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260056020526040600020553880806136fd565b5050600090565b906001820191816000528260205260406000205490811515600014613867577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191808311612c3a5781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111612c3a57838161381e9503613830575b5050506135f5565b60005260205260006040812055600190565b61385061384061374093866135dd565b90549060031b1c928392866135dd565b905560005284602052604060002055388080613816565b50505050600090565b805490680100000000000000008210156127fb5781613740916001613897940181556135dd565b9055565b806000526005602052604060002054156000146138d4576138bd816004613870565b600454906000526005602052604060002055600190565b50600090565b806000526003602052604060002054156000146138d4576138fc816002613870565b600254906000526003602052604060002055600190565b600082815260018201602052604090205461378d578061393583600193613870565b80549260005201602052604060002055600190565b906040519182815491828252602082019060005260206000209260005b81811061397e57505061397c9250038361289a565b565b8454835260019485019487945060209093019201613967565b9067ffffffffffffffff16801561378d5760009081526001820160205260409020546139dc5760019060008052016020526040600020546139d757600090565b600390565b50600290565b600081815260036020526040902054801561378d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111612c3a57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211612c3a57808203613a78575b505050613a6460026135f5565b600052600360205260006040812055600190565b613a9a613a896137409360026135dd565b90549060031b1c92839260026135dd565b90556000526003602052604060002055388080613a5756fea164736f6c634300081a000a",
}

var FeePolicyStoreABI = FeePolicyStoreMetaData.ABI

var FeePolicyStoreBin = FeePolicyStoreMetaData.Bin

func DeployFeePolicyStore(auth *bind.TransactOpts, backend bind.ContractBackend, authorizedCallers []common.Address) (common.Address, *types.Transaction, *FeePolicyStore, error) {
	parsed, err := FeePolicyStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeePolicyStoreBin), backend, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeePolicyStore{address: address, abi: *parsed, FeePolicyStoreCaller: FeePolicyStoreCaller{contract: contract}, FeePolicyStoreTransactor: FeePolicyStoreTransactor{contract: contract}, FeePolicyStoreFilterer: FeePolicyStoreFilterer{contract: contract}}, nil
}

type FeePolicyStore struct {
	address common.Address
	abi     abi.ABI
	FeePolicyStoreCaller
	FeePolicyStoreTransactor
	FeePolicyStoreFilterer
}

type FeePolicyStoreCaller struct {
	contract *bind.BoundContract
}

type FeePolicyStoreTransactor struct {
	contract *bind.BoundContract
}

type FeePolicyStoreFilterer struct {
	contract *bind.BoundContract
}

type FeePolicyStoreSession struct {
	Contract     *FeePolicyStore
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeePolicyStoreCallerSession struct {
	Contract *FeePolicyStoreCaller
	CallOpts bind.CallOpts
}

type FeePolicyStoreTransactorSession struct {
	Contract     *FeePolicyStoreTransactor
	TransactOpts bind.TransactOpts
}

type FeePolicyStoreRaw struct {
	Contract *FeePolicyStore
}

type FeePolicyStoreCallerRaw struct {
	Contract *FeePolicyStoreCaller
}

type FeePolicyStoreTransactorRaw struct {
	Contract *FeePolicyStoreTransactor
}

func NewFeePolicyStore(address common.Address, backend bind.ContractBackend) (*FeePolicyStore, error) {
	abi, err := abi.JSON(strings.NewReader(FeePolicyStoreABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeePolicyStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStore{address: address, abi: abi, FeePolicyStoreCaller: FeePolicyStoreCaller{contract: contract}, FeePolicyStoreTransactor: FeePolicyStoreTransactor{contract: contract}, FeePolicyStoreFilterer: FeePolicyStoreFilterer{contract: contract}}, nil
}

func NewFeePolicyStoreCaller(address common.Address, caller bind.ContractCaller) (*FeePolicyStoreCaller, error) {
	contract, err := bindFeePolicyStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreCaller{contract: contract}, nil
}

func NewFeePolicyStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*FeePolicyStoreTransactor, error) {
	contract, err := bindFeePolicyStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreTransactor{contract: contract}, nil
}

func NewFeePolicyStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*FeePolicyStoreFilterer, error) {
	contract, err := bindFeePolicyStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreFilterer{contract: contract}, nil
}

func bindFeePolicyStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeePolicyStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FeePolicyStore *FeePolicyStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeePolicyStore.Contract.FeePolicyStoreCaller.contract.Call(opts, result, method, params...)
}

func (_FeePolicyStore *FeePolicyStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.FeePolicyStoreTransactor.contract.Transfer(opts)
}

func (_FeePolicyStore *FeePolicyStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.FeePolicyStoreTransactor.contract.Transact(opts, method, params...)
}

func (_FeePolicyStore *FeePolicyStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeePolicyStore.Contract.contract.Call(opts, result, method, params...)
}

func (_FeePolicyStore *FeePolicyStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.contract.Transfer(opts)
}

func (_FeePolicyStore *FeePolicyStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.contract.Transact(opts, method, params...)
}

func (_FeePolicyStore *FeePolicyStoreCaller) CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "checkLaneRecords", localToken, remoteChainSelector)

	if err != nil {
		return err
	}

	return err

}

func (_FeePolicyStore *FeePolicyStoreSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _FeePolicyStore.Contract.CheckLaneRecords(&_FeePolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _FeePolicyStore.Contract.CheckLaneRecords(&_FeePolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeePolicyStore.Contract.GetAllAuthorizedCallers(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeePolicyStore.Contract.GetAllAuthorizedCallers(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetDestOverheadPolicy(opts *bind.CallOpts, policyName [32]byte) (GetDestOverheadPolicy,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getDestOverheadPolicy", policyName)

	outstruct := new(GetDestOverheadPolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PolicyConfigs = *abi.ConvertType(out[1], new([]FeePolicyStoreDestOverheadPolicyConfigArgs)).(*[]FeePolicyStoreDestOverheadPolicyConfigArgs)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetDestOverheadPolicy(policyName [32]byte) (GetDestOverheadPolicy,

	error) {
	return _FeePolicyStore.Contract.GetDestOverheadPolicy(&_FeePolicyStore.CallOpts, policyName)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetDestOverheadPolicy(policyName [32]byte) (GetDestOverheadPolicy,

	error) {
	return _FeePolicyStore.Contract.GetDestOverheadPolicy(&_FeePolicyStore.CallOpts, policyName)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetFeeBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getFeeBatch", lanes, requestedFinalityConfig)

	if err != nil {
		return *new([]IPoolConfigManagerLaneFee), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneFee)).(*[]IPoolConfigManagerLaneFee)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetFeeBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	return _FeePolicyStore.Contract.GetFeeBatch(&_FeePolicyStore.CallOpts, lanes, requestedFinalityConfig)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetFeeBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	return _FeePolicyStore.Contract.GetFeeBatch(&_FeePolicyStore.CallOpts, lanes, requestedFinalityConfig)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetFeePolicy(opts *bind.CallOpts, policyName [32]byte) (GetFeePolicy,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getFeePolicy", policyName)

	outstruct := new(GetFeePolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PolicyConfigs = *abi.ConvertType(out[1], new([]FeePolicyStoreFeePolicyConfigArgs)).(*[]FeePolicyStoreFeePolicyConfigArgs)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetFeePolicy(policyName [32]byte) (GetFeePolicy,

	error) {
	return _FeePolicyStore.Contract.GetFeePolicy(&_FeePolicyStore.CallOpts, policyName)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetFeePolicy(policyName [32]byte) (GetFeePolicy,

	error) {
	return _FeePolicyStore.Contract.GetFeePolicy(&_FeePolicyStore.CallOpts, policyName)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetInlineDestOverheadConfigs(opts *bind.CallOpts, localToken common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getInlineDestOverheadConfigs", localToken)

	if err != nil {
		return *new([]FeePolicyStoreInlineDestOverheadConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeePolicyStoreInlineDestOverheadConfigArgs)).(*[]FeePolicyStoreInlineDestOverheadConfigArgs)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetInlineDestOverheadConfigs(localToken common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	return _FeePolicyStore.Contract.GetInlineDestOverheadConfigs(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetInlineDestOverheadConfigs(localToken common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	return _FeePolicyStore.Contract.GetInlineDestOverheadConfigs(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetInlineFeeConfigs(opts *bind.CallOpts, localToken common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getInlineFeeConfigs", localToken)

	if err != nil {
		return *new([]FeePolicyStoreInlineFeeConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeePolicyStoreInlineFeeConfigArgs)).(*[]FeePolicyStoreInlineFeeConfigArgs)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetInlineFeeConfigs(localToken common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	return _FeePolicyStore.Contract.GetInlineFeeConfigs(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetInlineFeeConfigs(localToken common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	return _FeePolicyStore.Contract.GetInlineFeeConfigs(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getLaneRecordLevelBatch", lanes)

	outstruct := new(GetLaneRecordLevelBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeLevels = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.DestOverheadLevels = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _FeePolicyStore.Contract.GetLaneRecordLevelBatch(&_FeePolicyStore.CallOpts, lanes)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _FeePolicyStore.Contract.GetLaneRecordLevelBatch(&_FeePolicyStore.CallOpts, lanes)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getPolicyNames")

	outstruct := new(GetPolicyNames)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeePolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.DestOverheadPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _FeePolicyStore.Contract.GetPolicyNames(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _FeePolicyStore.Contract.GetPolicyNames(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getTokenPolicies", localToken)

	outstruct := new(GetTokenPolicies)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.FeePolicyName = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.DestOverheadPolicyName = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _FeePolicyStore.Contract.GetTokenPolicies(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _FeePolicyStore.Contract.GetTokenPolicies(&_FeePolicyStore.CallOpts, localToken)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetTokenPolicyNamesBatch(opts *bind.CallOpts, localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getTokenPolicyNamesBatch", localTokens)

	outstruct := new(GetTokenPolicyNamesBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeePolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.DestOverheadPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetTokenPolicyNamesBatch(localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	return _FeePolicyStore.Contract.GetTokenPolicyNamesBatch(&_FeePolicyStore.CallOpts, localTokens)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetTokenPolicyNamesBatch(localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	return _FeePolicyStore.Contract.GetTokenPolicyNamesBatch(&_FeePolicyStore.CallOpts, localTokens)
}

func (_FeePolicyStore *FeePolicyStoreCaller) GetTokenTransferFeeConfigBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "getTokenTransferFeeConfigBatch", lanes)

	if err != nil {
		return *new([]IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolV2TokenTransferFeeConfig)).(*[]IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) GetTokenTransferFeeConfigBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	return _FeePolicyStore.Contract.GetTokenTransferFeeConfigBatch(&_FeePolicyStore.CallOpts, lanes)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) GetTokenTransferFeeConfigBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	return _FeePolicyStore.Contract.GetTokenTransferFeeConfigBatch(&_FeePolicyStore.CallOpts, lanes)
}

func (_FeePolicyStore *FeePolicyStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) Owner() (common.Address, error) {
	return _FeePolicyStore.Contract.Owner(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) Owner() (common.Address, error) {
	return _FeePolicyStore.Contract.Owner(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) TypeAndVersion() (string, error) {
	return _FeePolicyStore.Contract.TypeAndVersion(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) TypeAndVersion() (string, error) {
	return _FeePolicyStore.Contract.TypeAndVersion(&_FeePolicyStore.CallOpts)
}

func (_FeePolicyStore *FeePolicyStoreCaller) ValidateLockOrBurn(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, chargeTokenFee bool, destPoolDataLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeePolicyStore.contract.Call(opts, &out, "validateLockOrBurn", localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee, destPoolDataLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeePolicyStore *FeePolicyStoreSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, chargeTokenFee bool, destPoolDataLength *big.Int) (*big.Int, error) {
	return _FeePolicyStore.Contract.ValidateLockOrBurn(&_FeePolicyStore.CallOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee, destPoolDataLength)
}

func (_FeePolicyStore *FeePolicyStoreCallerSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, chargeTokenFee bool, destPoolDataLength *big.Int) (*big.Int, error) {
	return _FeePolicyStore.Contract.ValidateLockOrBurn(&_FeePolicyStore.CallOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig, chargeTokenFee, destPoolDataLength)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "acceptOwnership")
}

func (_FeePolicyStore *FeePolicyStoreSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeePolicyStore.Contract.AcceptOwnership(&_FeePolicyStore.TransactOpts)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeePolicyStore.Contract.AcceptOwnership(&_FeePolicyStore.TransactOpts)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_FeePolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_FeePolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyDestOverheadPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyDestOverheadPolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyDestOverheadPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyDestOverheadPolicyConfigUpdates(&_FeePolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyDestOverheadPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyDestOverheadPolicyConfigUpdates(&_FeePolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyFeePolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyFeePolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyFeePolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyFeePolicyConfigUpdates(&_FeePolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyFeePolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyFeePolicyConfigUpdates(&_FeePolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyInlineDestOverheadConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyInlineDestOverheadConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyInlineDestOverheadConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyInlineDestOverheadConfigUpdates(&_FeePolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyInlineDestOverheadConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyInlineDestOverheadConfigUpdates(&_FeePolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyInlineFeeConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyInlineFeeConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyInlineFeeConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyInlineFeeConfigUpdates(&_FeePolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyInlineFeeConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyInlineFeeConfigUpdates(&_FeePolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyTokenDestOverheadPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyTokenDestOverheadPolicyUpdates", tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyTokenDestOverheadPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenDestOverheadPolicyUpdates(&_FeePolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyTokenDestOverheadPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenDestOverheadPolicyUpdates(&_FeePolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyTokenFeePolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyTokenFeePolicyUpdates", tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyTokenFeePolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenFeePolicyUpdates(&_FeePolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyTokenFeePolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenFeePolicyUpdates(&_FeePolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "applyTokenProfileUpdates", tokensToRemove, profilesToSet)
}

func (_FeePolicyStore *FeePolicyStoreSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenProfileUpdates(&_FeePolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.ApplyTokenProfileUpdates(&_FeePolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_FeePolicyStore *FeePolicyStoreTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeePolicyStore.contract.Transact(opts, "transferOwnership", to)
}

func (_FeePolicyStore *FeePolicyStoreSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.TransferOwnership(&_FeePolicyStore.TransactOpts, to)
}

func (_FeePolicyStore *FeePolicyStoreTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeePolicyStore.Contract.TransferOwnership(&_FeePolicyStore.TransactOpts, to)
}

type FeePolicyStoreAuthorizedCallerAddedIterator struct {
	Event *FeePolicyStoreAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreAuthorizedCallerAdded)
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
		it.Event = new(FeePolicyStoreAuthorizedCallerAdded)
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

func (it *FeePolicyStoreAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeePolicyStoreAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreAuthorizedCallerAddedIterator{contract: _FeePolicyStore.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreAuthorizedCallerAdded)
				if err := _FeePolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseAuthorizedCallerAdded(log types.Log) (*FeePolicyStoreAuthorizedCallerAdded, error) {
	event := new(FeePolicyStoreAuthorizedCallerAdded)
	if err := _FeePolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreAuthorizedCallerRemovedIterator struct {
	Event *FeePolicyStoreAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreAuthorizedCallerRemoved)
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
		it.Event = new(FeePolicyStoreAuthorizedCallerRemoved)
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

func (it *FeePolicyStoreAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeePolicyStoreAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreAuthorizedCallerRemovedIterator{contract: _FeePolicyStore.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreAuthorizedCallerRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*FeePolicyStoreAuthorizedCallerRemoved, error) {
	event := new(FeePolicyStoreAuthorizedCallerRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreDestOverheadPolicyConfigRemovedIterator struct {
	Event *FeePolicyStoreDestOverheadPolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreDestOverheadPolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreDestOverheadPolicyConfigRemoved)
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
		it.Event = new(FeePolicyStoreDestOverheadPolicyConfigRemoved)
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

func (it *FeePolicyStoreDestOverheadPolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreDestOverheadPolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreDestOverheadPolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterDestOverheadPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreDestOverheadPolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "DestOverheadPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreDestOverheadPolicyConfigRemovedIterator{contract: _FeePolicyStore.contract, event: "DestOverheadPolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchDestOverheadPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "DestOverheadPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreDestOverheadPolicyConfigRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyConfigRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseDestOverheadPolicyConfigRemoved(log types.Log) (*FeePolicyStoreDestOverheadPolicyConfigRemoved, error) {
	event := new(FeePolicyStoreDestOverheadPolicyConfigRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreDestOverheadPolicyConfigSetIterator struct {
	Event *FeePolicyStoreDestOverheadPolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreDestOverheadPolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreDestOverheadPolicyConfigSet)
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
		it.Event = new(FeePolicyStoreDestOverheadPolicyConfigSet)
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

func (it *FeePolicyStoreDestOverheadPolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreDestOverheadPolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreDestOverheadPolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	DestOverheadConfig  FeePolicyStoreDestOverheadConfig
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterDestOverheadPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreDestOverheadPolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "DestOverheadPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreDestOverheadPolicyConfigSetIterator{contract: _FeePolicyStore.contract, event: "DestOverheadPolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchDestOverheadPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "DestOverheadPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreDestOverheadPolicyConfigSet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyConfigSet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseDestOverheadPolicyConfigSet(log types.Log) (*FeePolicyStoreDestOverheadPolicyConfigSet, error) {
	event := new(FeePolicyStoreDestOverheadPolicyConfigSet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreDestOverheadPolicyRemovedIterator struct {
	Event *FeePolicyStoreDestOverheadPolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreDestOverheadPolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreDestOverheadPolicyRemoved)
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
		it.Event = new(FeePolicyStoreDestOverheadPolicyRemoved)
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

func (it *FeePolicyStoreDestOverheadPolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreDestOverheadPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreDestOverheadPolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterDestOverheadPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*FeePolicyStoreDestOverheadPolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "DestOverheadPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreDestOverheadPolicyRemovedIterator{contract: _FeePolicyStore.contract, event: "DestOverheadPolicyRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchDestOverheadPolicyRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "DestOverheadPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreDestOverheadPolicyRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseDestOverheadPolicyRemoved(log types.Log) (*FeePolicyStoreDestOverheadPolicyRemoved, error) {
	event := new(FeePolicyStoreDestOverheadPolicyRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "DestOverheadPolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreFeePolicyConfigRemovedIterator struct {
	Event *FeePolicyStoreFeePolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreFeePolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreFeePolicyConfigRemoved)
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
		it.Event = new(FeePolicyStoreFeePolicyConfigRemoved)
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

func (it *FeePolicyStoreFeePolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreFeePolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreFeePolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterFeePolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreFeePolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "FeePolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreFeePolicyConfigRemovedIterator{contract: _FeePolicyStore.contract, event: "FeePolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchFeePolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "FeePolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreFeePolicyConfigRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyConfigRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseFeePolicyConfigRemoved(log types.Log) (*FeePolicyStoreFeePolicyConfigRemoved, error) {
	event := new(FeePolicyStoreFeePolicyConfigRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreFeePolicyConfigSetIterator struct {
	Event *FeePolicyStoreFeePolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreFeePolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreFeePolicyConfigSet)
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
		it.Event = new(FeePolicyStoreFeePolicyConfigSet)
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

func (it *FeePolicyStoreFeePolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreFeePolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreFeePolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	FeeConfig           FeePolicyStoreFeeConfig
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterFeePolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreFeePolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "FeePolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreFeePolicyConfigSetIterator{contract: _FeePolicyStore.contract, event: "FeePolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchFeePolicyConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "FeePolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreFeePolicyConfigSet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyConfigSet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseFeePolicyConfigSet(log types.Log) (*FeePolicyStoreFeePolicyConfigSet, error) {
	event := new(FeePolicyStoreFeePolicyConfigSet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreFeePolicyRemovedIterator struct {
	Event *FeePolicyStoreFeePolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreFeePolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreFeePolicyRemoved)
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
		it.Event = new(FeePolicyStoreFeePolicyRemoved)
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

func (it *FeePolicyStoreFeePolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreFeePolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreFeePolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterFeePolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*FeePolicyStoreFeePolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "FeePolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreFeePolicyRemovedIterator{contract: _FeePolicyStore.contract, event: "FeePolicyRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchFeePolicyRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "FeePolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreFeePolicyRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseFeePolicyRemoved(log types.Log) (*FeePolicyStoreFeePolicyRemoved, error) {
	event := new(FeePolicyStoreFeePolicyRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "FeePolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreInlineDestOverheadConfigRemovedIterator struct {
	Event *FeePolicyStoreInlineDestOverheadConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreInlineDestOverheadConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreInlineDestOverheadConfigRemoved)
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
		it.Event = new(FeePolicyStoreInlineDestOverheadConfigRemoved)
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

func (it *FeePolicyStoreInlineDestOverheadConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreInlineDestOverheadConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreInlineDestOverheadConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterInlineDestOverheadConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineDestOverheadConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "InlineDestOverheadConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreInlineDestOverheadConfigRemovedIterator{contract: _FeePolicyStore.contract, event: "InlineDestOverheadConfigRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchInlineDestOverheadConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineDestOverheadConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "InlineDestOverheadConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreInlineDestOverheadConfigRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "InlineDestOverheadConfigRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseInlineDestOverheadConfigRemoved(log types.Log) (*FeePolicyStoreInlineDestOverheadConfigRemoved, error) {
	event := new(FeePolicyStoreInlineDestOverheadConfigRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "InlineDestOverheadConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreInlineDestOverheadConfigSetIterator struct {
	Event *FeePolicyStoreInlineDestOverheadConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreInlineDestOverheadConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreInlineDestOverheadConfigSet)
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
		it.Event = new(FeePolicyStoreInlineDestOverheadConfigSet)
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

func (it *FeePolicyStoreInlineDestOverheadConfigSetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreInlineDestOverheadConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreInlineDestOverheadConfigSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	DestOverheadConfig  FeePolicyStoreDestOverheadConfig
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterInlineDestOverheadConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineDestOverheadConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "InlineDestOverheadConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreInlineDestOverheadConfigSetIterator{contract: _FeePolicyStore.contract, event: "InlineDestOverheadConfigSet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchInlineDestOverheadConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineDestOverheadConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "InlineDestOverheadConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreInlineDestOverheadConfigSet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "InlineDestOverheadConfigSet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseInlineDestOverheadConfigSet(log types.Log) (*FeePolicyStoreInlineDestOverheadConfigSet, error) {
	event := new(FeePolicyStoreInlineDestOverheadConfigSet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "InlineDestOverheadConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreInlineFeeConfigRemovedIterator struct {
	Event *FeePolicyStoreInlineFeeConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreInlineFeeConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreInlineFeeConfigRemoved)
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
		it.Event = new(FeePolicyStoreInlineFeeConfigRemoved)
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

func (it *FeePolicyStoreInlineFeeConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreInlineFeeConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreInlineFeeConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterInlineFeeConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineFeeConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "InlineFeeConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreInlineFeeConfigRemovedIterator{contract: _FeePolicyStore.contract, event: "InlineFeeConfigRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchInlineFeeConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineFeeConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "InlineFeeConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreInlineFeeConfigRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "InlineFeeConfigRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseInlineFeeConfigRemoved(log types.Log) (*FeePolicyStoreInlineFeeConfigRemoved, error) {
	event := new(FeePolicyStoreInlineFeeConfigRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "InlineFeeConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreInlineFeeConfigSetIterator struct {
	Event *FeePolicyStoreInlineFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreInlineFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreInlineFeeConfigSet)
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
		it.Event = new(FeePolicyStoreInlineFeeConfigSet)
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

func (it *FeePolicyStoreInlineFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreInlineFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreInlineFeeConfigSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	FeeConfig           FeePolicyStoreFeeConfig
	Raw                 types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterInlineFeeConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineFeeConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "InlineFeeConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreInlineFeeConfigSetIterator{contract: _FeePolicyStore.contract, event: "InlineFeeConfigSet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchInlineFeeConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineFeeConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "InlineFeeConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreInlineFeeConfigSet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "InlineFeeConfigSet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseInlineFeeConfigSet(log types.Log) (*FeePolicyStoreInlineFeeConfigSet, error) {
	event := new(FeePolicyStoreInlineFeeConfigSet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "InlineFeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreOwnershipTransferRequestedIterator struct {
	Event *FeePolicyStoreOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreOwnershipTransferRequested)
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
		it.Event = new(FeePolicyStoreOwnershipTransferRequested)
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

func (it *FeePolicyStoreOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeePolicyStoreOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreOwnershipTransferRequestedIterator{contract: _FeePolicyStore.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreOwnershipTransferRequested)
				if err := _FeePolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseOwnershipTransferRequested(log types.Log) (*FeePolicyStoreOwnershipTransferRequested, error) {
	event := new(FeePolicyStoreOwnershipTransferRequested)
	if err := _FeePolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreOwnershipTransferredIterator struct {
	Event *FeePolicyStoreOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreOwnershipTransferred)
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
		it.Event = new(FeePolicyStoreOwnershipTransferred)
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

func (it *FeePolicyStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeePolicyStoreOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreOwnershipTransferredIterator{contract: _FeePolicyStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreOwnershipTransferred)
				if err := _FeePolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseOwnershipTransferred(log types.Log) (*FeePolicyStoreOwnershipTransferred, error) {
	event := new(FeePolicyStoreOwnershipTransferred)
	if err := _FeePolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreTokenDestOverheadPolicySetIterator struct {
	Event *FeePolicyStoreTokenDestOverheadPolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreTokenDestOverheadPolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreTokenDestOverheadPolicySet)
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
		it.Event = new(FeePolicyStoreTokenDestOverheadPolicySet)
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

func (it *FeePolicyStoreTokenDestOverheadPolicySetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreTokenDestOverheadPolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreTokenDestOverheadPolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterTokenDestOverheadPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*FeePolicyStoreTokenDestOverheadPolicySetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var oldPolicyNameRule []interface{}
	for _, oldPolicyNameItem := range oldPolicyName {
		oldPolicyNameRule = append(oldPolicyNameRule, oldPolicyNameItem)
	}
	var newPolicyNameRule []interface{}
	for _, newPolicyNameItem := range newPolicyName {
		newPolicyNameRule = append(newPolicyNameRule, newPolicyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "TokenDestOverheadPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreTokenDestOverheadPolicySetIterator{contract: _FeePolicyStore.contract, event: "TokenDestOverheadPolicySet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchTokenDestOverheadPolicySet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenDestOverheadPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var oldPolicyNameRule []interface{}
	for _, oldPolicyNameItem := range oldPolicyName {
		oldPolicyNameRule = append(oldPolicyNameRule, oldPolicyNameItem)
	}
	var newPolicyNameRule []interface{}
	for _, newPolicyNameItem := range newPolicyName {
		newPolicyNameRule = append(newPolicyNameRule, newPolicyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "TokenDestOverheadPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreTokenDestOverheadPolicySet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "TokenDestOverheadPolicySet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseTokenDestOverheadPolicySet(log types.Log) (*FeePolicyStoreTokenDestOverheadPolicySet, error) {
	event := new(FeePolicyStoreTokenDestOverheadPolicySet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "TokenDestOverheadPolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreTokenFeePolicySetIterator struct {
	Event *FeePolicyStoreTokenFeePolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreTokenFeePolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreTokenFeePolicySet)
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
		it.Event = new(FeePolicyStoreTokenFeePolicySet)
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

func (it *FeePolicyStoreTokenFeePolicySetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreTokenFeePolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreTokenFeePolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterTokenFeePolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*FeePolicyStoreTokenFeePolicySetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var oldPolicyNameRule []interface{}
	for _, oldPolicyNameItem := range oldPolicyName {
		oldPolicyNameRule = append(oldPolicyNameRule, oldPolicyNameItem)
	}
	var newPolicyNameRule []interface{}
	for _, newPolicyNameItem := range newPolicyName {
		newPolicyNameRule = append(newPolicyNameRule, newPolicyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "TokenFeePolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreTokenFeePolicySetIterator{contract: _FeePolicyStore.contract, event: "TokenFeePolicySet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchTokenFeePolicySet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenFeePolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var oldPolicyNameRule []interface{}
	for _, oldPolicyNameItem := range oldPolicyName {
		oldPolicyNameRule = append(oldPolicyNameRule, oldPolicyNameItem)
	}
	var newPolicyNameRule []interface{}
	for _, newPolicyNameItem := range newPolicyName {
		newPolicyNameRule = append(newPolicyNameRule, newPolicyNameItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "TokenFeePolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreTokenFeePolicySet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "TokenFeePolicySet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseTokenFeePolicySet(log types.Log) (*FeePolicyStoreTokenFeePolicySet, error) {
	event := new(FeePolicyStoreTokenFeePolicySet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "TokenFeePolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreTokenProfileSetIterator struct {
	Event *FeePolicyStoreTokenProfileSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreTokenProfileSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreTokenProfileSet)
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
		it.Event = new(FeePolicyStoreTokenProfileSet)
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

func (it *FeePolicyStoreTokenProfileSetIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreTokenProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreTokenProfileSet struct {
	LocalToken common.Address
	Raw        types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*FeePolicyStoreTokenProfileSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreTokenProfileSetIterator{contract: _FeePolicyStore.contract, event: "TokenProfileSet", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreTokenProfileSet)
				if err := _FeePolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseTokenProfileSet(log types.Log) (*FeePolicyStoreTokenProfileSet, error) {
	event := new(FeePolicyStoreTokenProfileSet)
	if err := _FeePolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeePolicyStoreTokenRemovedIterator struct {
	Event *FeePolicyStoreTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeePolicyStoreTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePolicyStoreTokenRemoved)
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
		it.Event = new(FeePolicyStoreTokenRemoved)
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

func (it *FeePolicyStoreTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *FeePolicyStoreTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeePolicyStoreTokenRemoved struct {
	LocalToken common.Address
	Raw        types.Log
}

func (_FeePolicyStore *FeePolicyStoreFilterer) FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*FeePolicyStoreTokenRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _FeePolicyStore.contract.FilterLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeePolicyStoreTokenRemovedIterator{contract: _FeePolicyStore.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

func (_FeePolicyStore *FeePolicyStoreFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _FeePolicyStore.contract.WatchLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeePolicyStoreTokenRemoved)
				if err := _FeePolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

func (_FeePolicyStore *FeePolicyStoreFilterer) ParseTokenRemoved(log types.Log) (*FeePolicyStoreTokenRemoved, error) {
	event := new(FeePolicyStoreTokenRemoved)
	if err := _FeePolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestOverheadPolicy struct {
	TokenCount    uint32
	PolicyConfigs []FeePolicyStoreDestOverheadPolicyConfigArgs
}
type GetFeePolicy struct {
	TokenCount    uint32
	PolicyConfigs []FeePolicyStoreFeePolicyConfigArgs
}
type GetLaneRecordLevelBatch struct {
	FeeLevels          []uint8
	DestOverheadLevels []uint8
}
type GetPolicyNames struct {
	FeePolicyNames          [][32]byte
	DestOverheadPolicyNames [][32]byte
}
type GetTokenPolicies struct {
	IsRegistered           bool
	FeePolicyName          [32]byte
	DestOverheadPolicyName [32]byte
}
type GetTokenPolicyNamesBatch struct {
	FeePolicyNames          [][32]byte
	DestOverheadPolicyNames [][32]byte
}

func (FeePolicyStoreAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (FeePolicyStoreAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (FeePolicyStoreDestOverheadPolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0xf23ee1ed541c9a6fc556979fa67ee730c8b973567f320dfab6b0f08fa98b640d")
}

func (FeePolicyStoreDestOverheadPolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc145aeb57aa1bc23c7a23a7b640e48557aa6ff73c43025727740341561d9afac")
}

func (FeePolicyStoreDestOverheadPolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0xdf743435d049ae540ec6d143d4d295658be6883eaedf0962d1311cf5da2610d5")
}

func (FeePolicyStoreFeePolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0xaaa43a26411a6c040ef0c507c8f37898e4f20ed5da1f7e361a09dded66360e25")
}

func (FeePolicyStoreFeePolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3275d66dec24a74d993e8775d56a0a359d9f1949cdbf98c1eb4c524728bf360f")
}

func (FeePolicyStoreFeePolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0x6d8754acd53d70e89e08ddb26fe365e6a02f960df734251862d4d69b036a261a")
}

func (FeePolicyStoreInlineDestOverheadConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x5a79db4177d18f5f1b88709312519a5866f9292ee3621e520c4fd46112988d01")
}

func (FeePolicyStoreInlineDestOverheadConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2d6db97987024f726ece930c43a48619971d6d1eb8dcdf871f53c3f9f30b87a9")
}

func (FeePolicyStoreInlineFeeConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x3c4fb3bc77a699c9e93a8ae0aaa4517719eac2360b1432d8c84edcd3b061b06e")
}

func (FeePolicyStoreInlineFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0xcd1632808d4a131200cb019663e9454426bc07617b7b2995874f7ca33e5af162")
}

func (FeePolicyStoreOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeePolicyStoreOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FeePolicyStoreTokenDestOverheadPolicySet) Topic() common.Hash {
	return common.HexToHash("0x7bd0030593fdd95573b972204541f9eef9cedfa1daab494d1dc9838d11cfdba4")
}

func (FeePolicyStoreTokenFeePolicySet) Topic() common.Hash {
	return common.HexToHash("0x653518802ef88beefe781cd891f284259493bcf3231e3ab39f4b5a555a02afae")
}

func (FeePolicyStoreTokenProfileSet) Topic() common.Hash {
	return common.HexToHash("0x936075da10ac5036063e3a526ca53d09805e57da23aa8c124651eda754b4ec13")
}

func (FeePolicyStoreTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3")
}

func (_FeePolicyStore *FeePolicyStore) Address() common.Address {
	return _FeePolicyStore.address
}

type FeePolicyStoreInterface interface {
	CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetDestOverheadPolicy(opts *bind.CallOpts, policyName [32]byte) (GetDestOverheadPolicy,

		error)

	GetFeeBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error)

	GetFeePolicy(opts *bind.CallOpts, policyName [32]byte) (GetFeePolicy,

		error)

	GetInlineDestOverheadConfigs(opts *bind.CallOpts, localToken common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error)

	GetInlineFeeConfigs(opts *bind.CallOpts, localToken common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error)

	GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

		error)

	GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

		error)

	GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

		error)

	GetTokenPolicyNamesBatch(opts *bind.CallOpts, localTokens []common.Address) (GetTokenPolicyNamesBatch,

		error)

	GetTokenTransferFeeConfigBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ValidateLockOrBurn(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, chargeTokenFee bool, destPoolDataLength *big.Int) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestOverheadPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error)

	ApplyFeePolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error)

	ApplyInlineDestOverheadConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error)

	ApplyInlineFeeConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error)

	ApplyTokenDestOverheadPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenFeePolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeePolicyStoreAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*FeePolicyStoreAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeePolicyStoreAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*FeePolicyStoreAuthorizedCallerRemoved, error)

	FilterDestOverheadPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreDestOverheadPolicyConfigRemovedIterator, error)

	WatchDestOverheadPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseDestOverheadPolicyConfigRemoved(log types.Log) (*FeePolicyStoreDestOverheadPolicyConfigRemoved, error)

	FilterDestOverheadPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreDestOverheadPolicyConfigSetIterator, error)

	WatchDestOverheadPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseDestOverheadPolicyConfigSet(log types.Log) (*FeePolicyStoreDestOverheadPolicyConfigSet, error)

	FilterDestOverheadPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*FeePolicyStoreDestOverheadPolicyRemovedIterator, error)

	WatchDestOverheadPolicyRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreDestOverheadPolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseDestOverheadPolicyRemoved(log types.Log) (*FeePolicyStoreDestOverheadPolicyRemoved, error)

	FilterFeePolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreFeePolicyConfigRemovedIterator, error)

	WatchFeePolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFeePolicyConfigRemoved(log types.Log) (*FeePolicyStoreFeePolicyConfigRemoved, error)

	FilterFeePolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*FeePolicyStoreFeePolicyConfigSetIterator, error)

	WatchFeePolicyConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFeePolicyConfigSet(log types.Log) (*FeePolicyStoreFeePolicyConfigSet, error)

	FilterFeePolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*FeePolicyStoreFeePolicyRemovedIterator, error)

	WatchFeePolicyRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreFeePolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseFeePolicyRemoved(log types.Log) (*FeePolicyStoreFeePolicyRemoved, error)

	FilterInlineDestOverheadConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineDestOverheadConfigRemovedIterator, error)

	WatchInlineDestOverheadConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineDestOverheadConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineDestOverheadConfigRemoved(log types.Log) (*FeePolicyStoreInlineDestOverheadConfigRemoved, error)

	FilterInlineDestOverheadConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineDestOverheadConfigSetIterator, error)

	WatchInlineDestOverheadConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineDestOverheadConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineDestOverheadConfigSet(log types.Log) (*FeePolicyStoreInlineDestOverheadConfigSet, error)

	FilterInlineFeeConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineFeeConfigRemovedIterator, error)

	WatchInlineFeeConfigRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineFeeConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineFeeConfigRemoved(log types.Log) (*FeePolicyStoreInlineFeeConfigRemoved, error)

	FilterInlineFeeConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*FeePolicyStoreInlineFeeConfigSetIterator, error)

	WatchInlineFeeConfigSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreInlineFeeConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineFeeConfigSet(log types.Log) (*FeePolicyStoreInlineFeeConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeePolicyStoreOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeePolicyStoreOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeePolicyStoreOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeePolicyStoreOwnershipTransferred, error)

	FilterTokenDestOverheadPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*FeePolicyStoreTokenDestOverheadPolicySetIterator, error)

	WatchTokenDestOverheadPolicySet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenDestOverheadPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenDestOverheadPolicySet(log types.Log) (*FeePolicyStoreTokenDestOverheadPolicySet, error)

	FilterTokenFeePolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*FeePolicyStoreTokenFeePolicySetIterator, error)

	WatchTokenFeePolicySet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenFeePolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenFeePolicySet(log types.Log) (*FeePolicyStoreTokenFeePolicySet, error)

	FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*FeePolicyStoreTokenProfileSetIterator, error)

	WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error)

	ParseTokenProfileSet(log types.Log) (*FeePolicyStoreTokenProfileSet, error)

	FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*FeePolicyStoreTokenRemovedIterator, error)

	WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeePolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error)

	ParseTokenRemoved(log types.Log) (*FeePolicyStoreTokenRemoved, error)

	Address() common.Address
}
