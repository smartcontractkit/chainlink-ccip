// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package security_policy_store

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

type IPoolConfigManagerLaneKey struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
}

type IPoolConfigManagerLaneRateLimiterState struct {
	Outbound             IPoolConfigManagerTokenBucket
	Inbound              IPoolConfigManagerTokenBucket
	FastFinalityOutbound IPoolConfigManagerTokenBucket
	FastFinalityInbound  IPoolConfigManagerTokenBucket
}

type IPoolConfigManagerTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
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

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type SecurityPolicyStoreInlineRateLimitBucketArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Direction           uint8
	FastFinality        bool
	Config              RateLimiterConfig
}

type SecurityPolicyStoreInlineRateLimitConfigArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	RateLimitConfig     SecurityPolicyStoreRateLimitConfig
}

type SecurityPolicyStoreInlineSecurityConfigArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	SecurityConfig      SecurityPolicyStoreSecurityConfig
}

type SecurityPolicyStoreRateLimitConfig struct {
	Outbound             RateLimiterConfig
	Inbound              RateLimiterConfig
	FastFinalityOutbound RateLimiterConfig
	FastFinalityInbound  RateLimiterConfig
}

type SecurityPolicyStoreRateLimitPolicyBucketArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Direction           uint8
	FastFinality        bool
	Config              RateLimiterConfig
}

type SecurityPolicyStoreRateLimitPolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalDecimals       uint8
	RateLimitConfig     SecurityPolicyStoreRateLimitConfig
}

type SecurityPolicyStoreSecurityConfig struct {
	OutboundCCVs []common.Address
	InboundCCVs  []common.Address
	Hook         common.Address
}

type SecurityPolicyStoreSecurityPolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	SecurityConfig      SecurityPolicyStoreSecurityConfig
}

var SecurityPolicyStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRateLimitBucketUpdates\",\"inputs\":[{\"name\":\"bucketConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitBucketArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRateLimitConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineSecurityConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineSecurityConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRateLimitPolicyBucketUpdates\",\"inputs\":[{\"name\":\"bucketConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyBucketArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRateLimitPolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySecurityPolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.SecurityPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenProfileUpdates\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"profilesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenProfileArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenRateLimitPolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenSecurityPolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkLaneRecords\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterStateBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"states\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneRateLimiterState[]\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineRateLimitConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineSecurityConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineSecurityConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneRecordLevelBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"rateLimitLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"securityLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneVerificationBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"hooks\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVs\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyNames\",\"inputs\":[],\"outputs\":[{\"name\":\"rateLimitPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"securityPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitPolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSecurityPolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.SecurityPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPolicies\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rateLimitPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"securityPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPolicyNamesBatch\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"rateLimitPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"securityPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validateLockOrBurn\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateReleaseOrMint\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineRateLimitBucketSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineRateLimitConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineRateLimitConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineSecurityConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineSecurityConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitPolicyBucketSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"direction\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitPolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitPolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitPolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityPolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityPolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityPolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenProfileSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRateLimitPolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSecurityPolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InlineRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidInlineRateLimitBucketConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidInlineRateLimitConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitPolicyBucketConfig\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitPolicyConfig\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitPolicyDecimals\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"LocalTokenNotRegistered\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyHasNoTokens\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyHasRecords\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyInUse\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyRecordNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RateLimitRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SecurityRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHasConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPolicyNameNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60806040523461020f576156ad8038038061001981610214565b92833981019060208183031261020f578051906001600160401b03821161020f570181601f8201121561020f578051916001600160401b0383116101c8578260051b9160208061006a818601610214565b80968152019382010191821161020f57602001915b8183106101ef578333156101de57600180546001600160a01b031916331790556020906100ab82610214565b60008152600036813760408051929083016001600160401b038111848210176101c8576040528252808383015260005b8151811015610142576001906001600160a01b036100f98285610239565b5116856101058261027b565b610112575b5050016100db565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1858561010a565b50505160005b81518110156101b9576001600160a01b036101638284610239565b51169081156101a8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef848361019a600195610379565b50604051908152a101610148565b6342bcdf7f60e11b60005260046000fd5b6040516152d390816103da8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b038116810361020f5781526020928301920161007f565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101c857604052565b805182101561024d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561024d5760005260206000200190600090565b600081815260036020526040902054801561037257600019810181811161035c5760025460001981019190821161035c5780820361030b575b50505060025480156102f557600019016102cf816002610263565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61034461031c61032d936002610263565b90549060031b1c9283926002610263565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102b4565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103d357600254680100000000000000008110156101c8576103ba61032d8260018594016002556002610263565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c8062b02f1a146131465780630678bd7614612ef45780630a2ad9a914612e485780630f1e1f5614612b5e578063181f5a7714612ac85780631dce24e61461280a5780632451a6271461279c5780632fc6499d146126465780633dcd28e11461246b5780634544b021146121d457806345da866714611fb35780635a985ae114611ddf5780636389dc0d14611d475780636576b869146119fb57806373ab65491461186157806379ba5097146117bb5780637b5f0d5f146116e75780638a379c4a1461101d5780638c63fe6214610f685780638da5cb5b14610f4157806391a2749a14610db25780639993974314610cbf578063a3ed219a14610b22578063c536d9ce146109dd578063cce3dffa14610628578063d8eb845f1461056b578063f2fde38b146104be5763f40eb2431461014f57600080fd5b346104b95760603660031901126104b9576004356001600160401b0381116104b95761017f90369060040161338c565b906024356001600160401b0381116104b95761019f9036906004016133ec565b6044356001600160401b0381116104b9576101be9036906004016133ec565b9390946101c9613df1565b60005b8181106103f75750505060005b81811061030c5750505060005b8181106101ef57005b6101fa818385613af0565b906102086040830183613adb565b6102426102366102188380613b07565b9061023b61023661022c6020880188613b07565b9390943691613743565b61512b565b3691613743565b602083019235610251846138b1565b9381156102fb576102f26001600160401b036102e47ff772dd29887be2365b45f9fa9c3428a0756ba16872f56b1ff296c97f1ccb0cc1936102b1600199610299886010614d15565b50876000526012602052846040600020911690614d15565b508560005260136020526102df876040806000206102ce856138b1565b908760009216825260205220614382565b6138b1565b169360405191829182613bb2565b0390a3016101e6565b633ca4b94360e21b60005260046000fd5b6103178183856139a9565b356040600082815260116020522054156103e357806000526012602052604060002063ffffffff6002820154166103b5575461038857908161035c6001936010614bd2565b507f55cff892eb89c654e76e11e2f4d53456cd9c03f313f32e50c6b035b34fd357ec600080a2016101d9565b7fa57cd4ec0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3c95cbc10000000000000000000000000000000000000000000000000000000060005260045260246000fd5b63f90d6e6b60e01b60005260045260246000fd5b610402818385613877565b356104196020610413848688613877565b016138b1565b9080600052601260205261043c6001600160401b03604060002093168093614bd2565b156104a257908160019392600052601360205260006002604080832083908682526020522061046a81613a5f565b610475878201613a5f565b01557ffd80d25f2d85a84586e465477f73b38272015e894d7915ca43708a8ff77665c3600080a3016101cc565b637798483760e01b60005260045260245260446000fd5b600080fd5b346104b95760203660031901126104b9576001600160a01b036104df613578565b6104e7614519565b16338114610541578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104b95760403660031901126104b957610584613578565b602435906001600160401b038216918281036104b9576105a4818361446d565b505060048110156105f05715610606576105be90826141b7565b5060048110156105f057156105cf57005b6001600160a01b03906308e6ea6160e11b6000521660045260245260446000fd5b634e487b7160e01b600052602160045260246000fd5b506001600160a01b0390636282cbeb60e11b6000521660045260245260446000fd5b346104b95760403660031901126104b9576004356001600160401b0381116104b95761065890369060040161338c565b602435906001600160401b0382116104b957366023830112156104b9578160040135926001600160401b0384116104b9573660246101c08602850101116104b9576106a1613df1565b60005b8281106108fc5750505060005b828110156108fa5760006101c0820283016024810160648201916106d483613e35565b1580156108e7575b80156108d3575b80156108bf575b610854576106f78261389d565b6001600160a01b03168452600760205260408420936107158361389d565b946044830195610724876138b1565b9061072e9261431d565b6107378361389d565b6001600160a01b03168152600860205260408120610754866138b1565b6001600160401b0316825260205260408120818052806020526040822082805260205284604083209061078691613ea6565b8180528060205260408220600183526020526040822060c484016107a991613ea6565b6001825280602052604082208280526020526040822061012484016107cd91613ea6565b6001825260205260408120600182526020526040902090610184016107f191613ea6565b6107fa9061389d565b91610804906138b1565b604051918261081291613cec565b6001600160401b0316916001600160a01b0316906101807f25b551e877c0171c7b8721a3973bd6ffdb14498f117775c6e16ed739e704e2f691a36001016106b1565b6108bd836001600160401b0361087860446108716101c49761389d565b95016138b1565b6001600160a01b03604051957fd6d3ea760000000000000000000000000000000000000000000000000000000087521660048601521660248401526044830190613cec565bfd5b506108cd6101848201613e35565b156106ea565b506108e16101248201613e35565b156106e3565b506108f460c48201613e35565b156106dc565b005b8061091261090d6001938686613877565b61389d565b6001600160401b0361092a6020610413858989613877565b61094c816001600160a01b0385169485600052600760205260406000206142de565b600083815260086020908152604080832093909416808352928152838220828052808252848320825284832083815588018390558783528483208381558801839055815283822082805290528281208181558601819055858152918220828155850182905591907f2909773a1782f3824d263d8654775762ceb2c33d5919a546083db42915205c809080a3016106a4565b346104b95760203660031901126104b9576004356001600160401b0381116104b957610a0d90369060040161338c565b610a168161372c565b90610a24604051928361370b565b808252601f19610a338261372c565b01366020840137610a4381613d26565b90610a4d81613d26565b9360005b828110610a9857610a7885610a9488610a868860405195869560608752606087019061353b565b9085820360208701526137b2565b9083820360408501526137b2565b0390f35b80610b066001610ac6610ab061090d83968989613877565b610ac06020610413878b8b613877565b906141b7565b90506001600160a01b03600282015416610ae0858b6138c5565b52610aea81613c4d565b610af4858a6138c5565b52610aff84896138c5565b5001613c4d565b610b1082896138c5565b52610b1b81886138c5565b5001610a51565b346104b95760203660031901126104b9576001600160a01b03610b43613578565b16806000526009602052610b5a6040600020614d4c565b805191610b668361372c565b92610b74604051948561370b565b808452610b83601f199161372c565b0160005b818110610ca857505060005b8251811015610c1357806001600160401b03610bb1600193866138c5565b511683600052600a602052610bf26040806000206000906001600160401b03851682526020522060405192610be5846136d5565b8684526020840152613ca1565b6040820152610c0182876138c5565b52610c0c81866138c5565b5001610b93565b836040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610c4c57505050500390f35b91936001919395506020610c988192603f198a8203018652606060408a516001600160a01b0381511684526001600160401b038682015116868501520151918160408201520190613609565b9601920192018594939192610c3d565b602090610cb3613c11565b82828801015201610b87565b346104b95760203660031901126104b9576004356001600160401b0381116104b957610cef90369060040161338c565b90610cf8613df1565b6000915b808310610d0557005b610d1361090d848385613877565b926001600160a01b036020610d29838587613877565b0135941693610d45856000526005602052604060002054151590565b15610d9d57846001939495600052600660205260026040600020019082825492610d6f82856140b7565b557fbdf2e0aae6b8f6b68d06d1a004b760b404a2e6cb361bd9dc3a255a508760f14a600080a4019190610cfc565b8463dfcc7fc760e01b60005260045260246000fd5b346104b95760203660031901126104b9576004356001600160401b0381116104b957604060031982360301126104b95760405190610def82613689565b80600401356001600160401b0381116104b957610e129060043691840101613797565b82526024810135906001600160401b0382116104b9576004610e379236920101613797565b60208201908152610e46614519565b519060005b8251811015610eb157806001600160a01b03610e69600193866138c5565b5116610e7481615232565b610e80575b5001610e4b565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610e79565b505160005b81518110156108fa576001600160a01b03610ed182846138c5565b5116908115610f17577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083610f09600195614cdc565b50604051908152a101610eb6565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104b95760003660031901126104b95760206001600160a01b0360015416604051908152f35b346104b95760003660031901126104b957604051600b548082528160208101600b60005260206000209260005b818110611004575050610faa9250038261370b565b604051906010548083528260208101601060005260206000209260005b818110610feb575050610fdc9250038361370b565b610a9460405192839283613450565b8454835260019485019487945060209093019201610fc7565b8454835260019485019486945060209093019201610f95565b346104b95760603660031901126104b9576004356001600160401b0381116104b95761104d90369060040161338c565b6024356001600160401b0381116104b95761106c9036906004016133ec565b604435936001600160401b0385116104b957366023860112156104b9578460040135936001600160401b0385116104b9573660246101e08702880101116104b9576110b5613df1565b60005b8181106116115750505060005b81811061157f57600085855b808310156108fa576000926101e08102830160248101356110ff81600052600c602052604060002054151590565b1591821580611558575b61152057608481019261111b84613e35565b15801561150d575b80156114f9575b80156114e5575b61148d576044820190611143826138b1565b841561147e579261144b61143760646114306101a096604060019d9e9f6001600160401b03986111ae7fa5267cb30eb155107fe8fb4922632854c7ea8ece482e8b3a3e465b9b42f392109c8e61119a81600b614d15565b508452600d6020528b858520911690614d15565b50611456575b8b8152600f602052818120896111c9856138b1565b168252602052818120818052806020528282208280526020528d6111ef848420916138e9565b15158a70ffffffffffffffffffffffffffffffffff1960ff70ffffffffffffffffffffffffffffffff0061122760a487549501613a22565b60081b16931691161717815560016001600160801b039061124a60c48d01613a22565b920191166fffffffffffffffffffffffffffffffff19825416179055818052806020528282206001835260205282822061128660e48b016138e9565b15158a70ffffffffffffffffffffffffffffffffff1960ff70ffffffffffffffffffffffffffffffff006112bf61010487549501613a22565b60081b16931691161717815560016001600160801b03906112e36101248d01613a22565b920191166fffffffffffffffffffffffffffffffff198254161790558f82528060205282822082805260205282822061131f6101448b016138e9565b15158a70ffffffffffffffffffffffffffffffffff1960ff70ffffffffffffffffffffffffffffffff0061135861016487549501613a22565b60081b16931691161717815560016001600160801b039061137c6101848d01613a22565b920191166fffffffffffffffffffffffffffffffff198254161790558f82526020528181208f8252602052206113b56101a488016138e9565b1515815470ffffffffffffffffffffffffffffffffff1960ff70ffffffffffffffffffffffffffffffff006113ed6101c48d01613a22565b60081b1693169116171781556001600160801b038e61140f6101e48a01613a22565b920191166fffffffffffffffffffffffffffffffff198254161790556138b1565b93016139b9565b9660ff604051981688526020880190613cec565b1693a30191906110d1565b6114618589016139b9565b8c8252600e60205260ff838320911660ff198254161790556111b4565b60048a633ca4b94360e21b8152fd5b506101c4926001600160401b036114a960446108bd94016138b1565b604051947f6df5311c00000000000000000000000000000000000000000000000000000000865260048601521660248401526044830190613cec565b506114f36101a48301613e35565b15611131565b506115076101448301613e35565b1561112a565b5061151a60e48301613e35565b15611123565b60ff8288926064948452600e602052611541858460408720541692016139b9565b91631d5de0ff60e01b855260045260245216604452fd5b50818752600e60205260ff60408820541660ff611577606484016139b9565b161415611109565b61158a8183856139a9565b3560406000828152600c6020522054156103e35780600052600d602052604060002063ffffffff6002820154166103b557546103885790816115cf600193600b614bd2565b5080600052600e602052604060002060ff1981541690557fa106ef2c214ebd9b86d462fe971cb93bb5f81d3807ac052940dc6070007e521a600080a2016110c5565b61161c818385613877565b3561162d6020610413848688613877565b9080600052600d6020526116506001600160401b03604060002093168093614bd2565b156104a2576000818152600f602090815260408083208584528252808320838052808352818420835281842084815560019081018590558085528285208581558101859055908352818420848052909252808320838155820183905581835282208281558101829055939291907f38b4904e409fb3862e7e48aa073f3f2a0d0af85607884cdf07b89b5eaf1819c29080a3016110b8565b346104b95760203660031901126104b9576004356001600160401b0381116104b95761171790369060040161338c565b61172081613977565b9061172a81613977565b9260005b82811061175f5761175184610a9487604051938493604085526040850190613647565b908382036020850152613647565b806117b56117aa61177661090d6001958888613877565b6117866020610413868a8a613877565b906117a5611794838361446d565b50506117a0878c6138c5565b613ce0565b6141b7565b506117a083896138c5565b0161172e565b346104b95760003660031901126104b9576000546001600160a01b03811633036118375773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346104b95760203660031901126104b957600435806000526012602052604060002061188c81614d4c565b918251926118998461372c565b936118a7604051958661370b565b8085526118b6601f199161372c565b0160005b8181106119e457505060005b815181101561194657806001600160401b036118e4600193856138c5565b51168460005260136020526119256040806000206000906001600160401b03851682526020522060405192611918846136d5565b8784526020840152613ca1565b604082015261193482886138c5565b5261193f81876138c5565b50016118c6565b8463ffffffff600286015416604051918291604083019083526040602084015281518091526060830190602060608260051b8601019301916000905b82821061199157505050500390f35b919360019193955060206119d48192605f198a8203018652606060408a51805184526001600160401b038682015116868501520151918160408201520190613609565b9601920192018594939192611982565b6020906119ef613c11565b828289010152016118ba565b346104b95760203660031901126104b9576001600160a01b03611a1c613578565b16806000526007602052611a336040600020614d4c565b805191611a3f8361372c565b92611a4d604051948561370b565b808452611a5c601f199161372c565b0160005b818110611d1557505060005b8251811015611c9d57806001600160401b03611a8a600193866138c5565b5116611ac86000611a996139c7565b508581526008602052604081206001600160401b0384168252602052604081208180526020526040812061395f565b8360ff825460a01c1691015460405191611ae1836136d5565b151582526001600160801b038116602083015260801c6040820152611b368460406000611b0c6139c7565b5088815260086020528181206001600160401b03871682526020528181208180526020522061395f565b8460ff825460a01c1691015460405191611b4f836136d5565b151582526001600160801b038116602083015260801c6040820152611ba66000611b776139c7565b508781526008602052604081206001600160401b0386168252602052604081208782526020526040812061395f565b8560ff825460a01c1691015460405191611bbf836136d5565b151582526001600160801b038116602083015260801c6040820152611c148660406000611bea6139c7565b508a815260086020528181206001600160401b03891682526020528181208382526020522061395f565b918660ff845460a01c1693015460405193611c2e856136d5565b151584526001600160801b038116602085015260801c604084015260405193611c56856136ba565b845260208401526040830152606082015260405191611c74836136d5565b85835260208301526040820152611c8b82876138c5565b52611c9681866138c5565b5001611a6c565b8360405180916020820160208352815180915260206040840192019060005b818110611cca575050500390f35b9193509160206101c0600192611d07604088516001600160a01b0381511684526001600160401b0386820151168685015201516040830190613478565b019401910191849392611cbc565b602090604051611d24816136d5565b60008152600083820152611d366139e6565b604082015282828801015201611a60565b346104b95760203660031901126104b95760806001600160a01b03611d6a613578565b16806000526006602052604060002060405191611d86836136d5565b60ff825416835260ff611dbe600260018501549460208701958652015492604086019384526000526005602052604060002054151590565b93511691519051916040519315158452602084015260408301526060820152f35b346104b9577fffffffff00000000000000000000000000000000000000000000000000000000611e0e366135a2565b611e1a94919394613df1565b16151591611e28828561446d565b906000926004811015611f0c5715611f8a579086918680611f76575b611f6c575b611e5d876001859488948a611e6298614e74565b614f35565b5060009215611f2057604080516001600160a01b038616815260208101929092526001600160401b038316917fc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f99190a25b611ebd81846141b7565b93906004811015611f0c5715611ee55760206001600160a01b03600286015416604051908152f35b6308e6ea6160e11b83526001600160a01b03166004526001600160401b0316602452604490fd5b602484634e487b7160e01b81526021600452fd5b604080516001600160a01b038616815260208101929092526001600160401b038316917f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2611eb3565b9295508592611e49565b50611f846001838386614de1565b15611e44565b6044836001600160401b03876001600160a01b038b636282cbeb60e11b85521660045216602452fd5b346104b95760403660031901126104b9576004356001600160401b0381116104b957611fe390369060040161338c565b6024356001600160401b0381116104b9576120029036906004016133ec565b91909261200d613df1565b60005b8281106121295750505060005b81811061202657005b806120346001928486613af0565b6120416040820182613adb565b6120516102366102188380613b07565b6001600160a01b036120628361389d565b1660005260096020527f0f6e7d77fde88adeb68e2ceed23ca44c93b917ae178cd836ce8a81a41c1ad8546001600160401b036121056120ff6040600020956120c16120ac8261389d565b9760208301986120bb8a6138b1565b9161431d565b6001600160a01b036120d28261389d565b16600052600a60205261090d866040806000206120ee8b6138b1565b908860009216825260205220614382565b946138b1565b6121206001600160a01b036040519384931696169482613bb2565b0390a30161201d565b8061213a61090d6001938686613877565b6001600160401b036121526020610413858989613877565b612174816001600160a01b0385169485600052600960205260406000206142de565b82600052600a602052600060026040808320839086861682526020522061219a81613a5f565b6121a5888201613a5f565b015516907f9c8ab32bba343fe78edb9a8fc032ad81100c097e58e133bd46cd2cadd5bb633b600080a301612010565b346104b95760203660031901126104b9576004356001600160401b0381116104b9576122049036906004016133bc565b61220c613df1565b60005b81811061221857005b6122238183856138d9565b90813580600052600d60205260406000209061226560208501926001600160401b0361224e856138b1565b169060019160005201602052604060002054151590565b15612441576080840161227781613e35565b156123bc5781600052600f602052604080600020612294856138b1565b906001600160401b03600092168252602052209260608601936122b6856138e9565b1515600052602052604060002090604087013560028110156104b9576123b37fbf5d213dcd6150f8deb1dae200784d598e8895a1f9413c91733d415730b44eb6946123726001600160401b039460019b6001600160801b038d61140f60c061231f8a829d61395f565b94612329896138e9565b1515865470ffffffffffffffffffffffffffffffffff1960ff70ffffffffffffffffffffffffffffffff0061236060a08701613a22565b60081b16931691161717865501613a22565b9287600052600e6020526123a261239160ff604060002054169a6138e9565b916040519a8b5260208b01906138f6565b151560408901526060880190613924565b1693a30161220f565b6123c685936138b1565b9060408401359160028310156104b9576124306108bd936001600160401b036123f3606060e499016138e9565b93604051977f2113b40f000000000000000000000000000000000000000000000000000000008952600489015216602487015260448601906138f6565b151560648401526084830190613924565b6124526001600160401b03926138b1565b90637798483760e01b6000526004521660245260446000fd5b346104b957612479366135a2565b9190612483613df1565b61248d82856141b7565b939060048110156105f0571561261b577fffffffff000000000000000000000000000000000000000000000000000000001615156000906124ce848761446d565b9391600481101561260757156125de5792828592896001600160a01b03999760209b978460029b98906125cb575b6125c1575b8484612514959697611e5d938b87614e74565b501561257157604080516001600160a01b0394909416845260208401919091526001600160401b0391909116917f28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e9190a25b015416604051908152f35b604080516001600160a01b0394909416845260208401919091526001600160401b0391909116917fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a2612566565b9293508392612501565b506125d884878386614de1565b156124fc565b6044826001600160401b03886001600160a01b038c636282cbeb60e11b85521660045216602452fd5b602483634e487b7160e01b81526021600452fd5b6001600160401b03836001600160a01b03876308e6ea6160e11b600052166004521660245260446000fd5b346104b95760203660031901126104b9576004356001600160401b0381116104b95761267690369060040161338c565b9061267f613df1565b6000915b80831061268c57005b61269a61090d848385613877565b926001600160a01b0360206126b0838587613877565b01359416936126cc856000526005602052604060002054151590565b15610d9d578460005260066020526040600020948161272a575b6001809495960190828254926126fc8285613fa2565b557fc0d2b9fef1d1d030378ff1ee5fb1d9ffcdd755598d86e1de88cc92cb5eb36791600080a4019190612683565b61274182600052600c602052604060002054151590565b156127875781600052600e60205260ff6040600020541660ff8754169081810361276c5750506126e6565b83631d5de0ff60e01b60005260045260245260445260646000fd5b5063f90d6e6b60e01b60005260045260246000fd5b346104b95760003660031901126104b9576040516002548082526020820190600260005260206000209060005b8181106127f457610a94856127e08187038261370b565b60405191829160208352602083019061353b565b82548452602090930192600192830192016127c9565b346104b95760203660031901126104b95760043580600052600d602052604060002081600052600e60205260ff604060002054169161284882614d4c565b908151916128558361372c565b92612863604051948561370b565b808452612872601f199161372c565b0160005b818110612a8f57505060005b81518110156129fa57806001600160401b036128a0600193856138c5565b51166128e66128e160006128b26139c7565b50878152600f602052604081206001600160401b0385168252602052604081208180526020526040812061395f565b613f6b565b6129236128e185604060006128f96139c7565b508a8152600f6020528181206001600160401b03881682526020528181208180526020522061395f565b6129626128e160006129336139c7565b50898152600f602052604081206001600160401b0387168252602052604081208882526020526040812061395f565b61299f6128e187604060006129756139c7565b508c8152600f6020528181206001600160401b038a1682526020528181208382526020522061395f565b91604051936129ad856136ba565b8452602084015260408301526060820152604051916129cb836136ba565b868352602083015288604083015260608201526129e882876138c5565b526129f381866138c5565b5001612882565b838663ffffffff6002880154169160405192839260608401918452602084015260606040840152815180915260206080840192019060005b818110612a40575050500390f35b9193509160206101e0600192612a8160608851805184526001600160401b0386820151168685015260ff604082015116604085015201516060830190613478565b019401910191849392612a32565b602090604051612a9e816136ba565b6000815260008382015260006040820152612ab76139e6565b606082015282828801015201612876565b346104b95760003660031901126104b957604051612ae581613689565b601d81527f5365637572697479506f6c69637953746f726520322e312e302d646576000000602082015260405190602082528181519182602083015260005b838110612b465750508160006040809484010152601f80199101168101030190f35b60208282018101516040878401015285935001612b24565b346104b95760403660031901126104b9576004356001600160401b0381116104b957612b8e9036906004016133ec565b602435916001600160401b0383116104b957366023840112156104b9578260040135916001600160401b0383116104b95736602460c08502860101116104b957612bd6613df1565b60005b818110612d2f5750505060005b818110156108fa5760c08102830160006001600160a01b03612c0a6024840161389d565b16918215612d075782825260066020526040822091612c36846000526005602052604060002054151590565b80612ceb575b612ca557506020612c95849360447f45669b1c1be2ec7c5e8638c7b26de467615f4afbc4b701c647b23236e1222a4c94612c7860019998614c9d565b50019060ff612c86836139b9565b1660ff198254161790556139b9565b60ff60405191168152a201612be6565b60ff84606493612cbb60448488541692016139b9565b917f900b828800000000000000000000000000000000000000000000000000000000855260045260245216604452fd5b5060ff83541660ff612cff604485016139b9565b161415612c3c565b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b6001600160a01b03612d4561090d8385876139a9565b16612d5d816000526005602052604060002054151590565b15612e3457806000526006602052604060002060018101541590811591612e26575b508015612e0f575b8015612df8575b612dcb579081612d9f600193614b20565b507f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3600080a201612bd9565b7f41025a5e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b508060005260096020526040600020541515612d8e565b508060005260076020526040600020541515612d87565b600291500154151587612d7f565b63dfcc7fc760e01b60005260045260246000fd5b346104b95760203660031901126104b9576004356001600160401b0381116104b957612e789036906004016133ec565b90612e8282613977565b90612e8c83613977565b9260005b818110612ea957505050610a9460405192839283613450565b806001600160a01b03612ec261090d60019486886139a9565b1660005260066020526002604060002083810154612ee084896138c5565b520154612eed82886138c5565b5201612e90565b346104b95760203660031901126104b9576004356001600160401b0381116104b957612f249036906004016133bc565b612f2c613df1565b60005b818110612f3857005b612f438183856138d9565b906001600160a01b03612f558361389d565b1660005260076020526040600020612f7c60208401916001600160401b0361224e846138b1565b1561310f5760808301612f8e81613e35565b15613087576001600160a01b03612fa48561389d565b166000526008602052604080600020612fbc846138b1565b906001600160401b03600092168252602052206060850190612fdd826138e9565b151560005260205260406000209460408101359260028410156104b9576001600160401b037f1e0a61edb07c831e721f1d19d6285ebad9f6a7847fc7e7597d7fb2f35a7dffde9361307c8361305e61305861305260a09861090d60019f9961304d8e6001600160a01b039c61395f565b613ea6565b9a6138b1565b936138e9565b61306b604051809a6138f6565b151560208901526040880190613924565b16941692a301612f2f565b90915061309c6130968461389d565b926138b1565b9060408401359160028310156104b9576124306108bd936001600160401b036130c9606060e499016138e9565b936001600160a01b03604051987ff7c24770000000000000000000000000000000000000000000000000000000008a5216600489015216602487015260448601906138f6565b6001600160401b03906001600160a01b0361312c6130968661389d565b9163507f5f5760e11b600052166004521660245260446000fd5b346104b95760203660031901126104b9576004356001600160401b0381116104b95761317690369060040161338c565b61317f8161372c565b9161318d604051938461370b565b818352601f1961319c8361372c565b0160005b81811061337557505060005b82811061332f578360405180916020820160208352815180915260206040840192019060005b8181106131e0575050500390f35b9193509160206102806001926133216060885161323b8482516001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b6132898682015160a08601906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60408181015180516001600160801b03908116610140880152602082015163ffffffff16610160880152918101511515610180870152606081015182166101a087015260800151166101c085015201516101e08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0194019101918493926131d2565b8061335961334361090d6001948787613877565b6133536020610413858989613877565b90613d70565b61336382876138c5565b5261336e81866138c5565b50016131ac565b602090613380613836565b828288010152016131a0565b9181601f840112156104b9578235916001600160401b0383116104b9576020808501948460061b0101116104b957565b9181601f840112156104b9578235916001600160401b0383116104b95760208085019460e085020101116104b957565b9181601f840112156104b9578235916001600160401b0383116104b9576020808501948460051b0101116104b957565b906020808351928381520192019060005b81811061343a5750505090565b825184526020938401939092019160010161342d565b90916134676134759360408452604084019061341c565b91602081840391015261341c565b90565b906101206060613539936134ae8482516001600160801b0360408092805115158552826020820151166020860152015116910152565b6134e06020820151838601906001600160801b0360408092805115158552826020820151166020860152015116910152565b6040818101518051151560c087015260208101516001600160801b0390811660e08801529101511661010085015201519101906001600160801b0360408092805115158552826020820151166020860152015116910152565b565b906020808351928381520192019060005b8181106135595750505090565b82516001600160a01b031684526020938401939092019160010161354c565b600435906001600160a01b03821682036104b957565b35906001600160a01b03821682036104b957565b60809060031901126104b9576004356001600160a01b03811681036104b957906024356001600160401b03811681036104b95790604435906064357fffffffff00000000000000000000000000000000000000000000000000000000811681036104b95790565b9060406001600160a01b038161363d61362b865160608752606087019061353b565b6020870151868203602088015261353b565b9401511691015290565b90602080835192838152019201906000905b8082106136665750505090565b90919283519060048210156105f057602081600193829352019401920190613659565b604081019081106001600160401b038211176136a457604052565b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b038211176136a457604052565b606081019081106001600160401b038211176136a457604052565b60a081019081106001600160401b038211176136a457604052565b90601f801991011681019081106001600160401b038211176136a457604052565b6001600160401b0381116136a45760051b60200190565b92919061374f8161372c565b9361375d604051958661370b565b602085838152019160051b81019283116104b957905b82821061377f57505050565b6020809161378c8461358e565b815201910190613773565b9080601f830112156104b95781602061347593359101613743565b9080602083519182815201916020808360051b8301019401926000915b8383106137de57505050505090565b90919293946020806137fc600193601f19868203018752895161353b565b970193019301919392906137cf565b60405190613818826136f0565b60006080838281528260208201528260408201528260608201520152565b60405190613843826136ba565b8161384c61380b565b815261385661380b565b602082015261386361380b565b6040820152606061387261380b565b910152565b91908110156138875760061b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b03811681036104b95790565b356001600160401b03811681036104b95790565b80518210156138875760209160051b010190565b91908110156138875760e0020190565b3580151581036104b95790565b9060028210156105f05752565b359081151582036104b957565b35906001600160801b03821682036104b957565b6001600160801b036139596040809361393c81613903565b151586528361394d60208301613910565b16602087015201613910565b16910152565b9060028110156105f057600052602052604060002090565b906139818261372c565b61398e604051918261370b565b828152809261399f601f199161372c565b0190602036910137565b91908110156138875760051b0190565b3560ff811681036104b95790565b604051906139d4826136d5565b60006040838281528260208201520152565b604051906139f3826136ba565b816139fc6139c7565b8152613a066139c7565b6020820152613a136139c7565b604082015260606138726139c7565b356001600160801b03811681036104b95790565b81810292918115918404141715613a4957565b634e487b7160e01b600052601160045260246000fd5b8054906000815581613a6f575050565b6000526020600020908101905b818110613a87575050565b60008155600101613a7c565b906801000000000000000081116136a457815491818155828210613ab657505050565b600052602060002091820191015b818110613acf575050565b60008155600101613ac4565b903590605e19813603018212156104b9570190565b90821015613887576134759160051b810190613adb565b903590601e19813603018212156104b957018035906001600160401b0382116104b957602001918160051b360383136104b957565b9035601e19823603018112156104b95701602081359101916001600160401b0382116104b9578160051b360383136104b957565b9160209082815201919060005b818110613b8a5750505090565b9091926020806001926001600160a01b03613ba48861358e565b168152019401929101613b7d565b6020815260606001600160a01b03613c0a6040613c03613be5613bd58880613b3c565b8760208a01526080890191613b70565b613bf26020890189613b3c565b888303601f1901858a015290613b70565b950161358e565b1691015290565b60405190613c1e826136d5565b8160008152600060208201526040805191613c38836136d5565b60608352606060208401526000828401520152565b906040519182815491828252602082019060005260206000209260005b818110613c7f5750506135399250038361370b565b84546001600160a01b0316835260019485019487945060209093019201613c6a565b90604051613cae816136d5565b60406001600160a01b0360028395613cc581613c4d565b8552613cd360018201613c4d565b6020860152015416910152565b60048210156105f05752565b6101208061353993613cfe8185613924565b613d0e6060820160608601613924565b613d1e60c0820160c08601613924565b019101613924565b90613d308261372c565b613d3d604051918261370b565b8281528092613d4e601f199161372c565b019060005b828110613d5f57505050565b806060602080938501015201613d53565b90613d79613836565b91613d84828261446d565b9160048196929610156105f05715613de9575080613da7613dc792868686614557565b94613db48282878761472b565b94613dc1838383886147c6565b94614861565b9160405193613dd5856136ba565b845260208401526040830152606082015290565b935050505090565b33600052600360205260406000205415613e0757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b613e3e816138e9565b613e7d576001600160801b03613e5660408301613a22565b16159081613e62575090565b6001600160801b0391506020613e789101613a22565b161590565b6001600160801b0380613e9e6020613e9760408601613a22565b9401613a22565b169116111590565b9060ff825460a01c16613eb8826138e9565b15159015151480613f41575b80613f1d575b613f19576060813603126104b95761353991613f0f6040805193613eed856136d5565b613ef681613903565b8552613f0460208201613910565b602086015201613910565b60408301526148fe565b5050565b50600182015460801c6001600160801b03613f3a60408401613a22565b1614613eca565b506001600160801b036001830154166001600160801b03613f6460208401613a22565b1614613ec4565b90604051613f78816136d5565b60406001600160801b036001839582815460ff81161515875260081c166020860152015416910152565b818114613f19578015159081614073575b82151591828061405d575b61404857614009575b50613fcf5750565b600052600d6020526002604060002001600163ffffffff8254160163ffffffff8111613a495763ffffffff1663ffffffff19825416179055565b600052600d602052600260406000200160001963ffffffff8254160163ffffffff8111613a495763ffffffff1663ffffffff1982541617905538613fc7565b8363f90d6e6b60e01b60005260045260246000fd5b5083600052600c60205260406000205415613fbe565b80600052600c602052604060002054156103e35780600052600d60205263ffffffff60026040600020015416613fb357635c32f70960e01b60005260045260246000fd5b818114613f19578015159081614173575b82151591828061415d575b6140485761411e575b506140e45750565b60005260126020526002604060002001600163ffffffff8254160163ffffffff8111613a495763ffffffff1663ffffffff19825416179055565b6000526012602052600260406000200160001963ffffffff8254160163ffffffff8111613a495763ffffffff1663ffffffff19825416179055386140dc565b50836000526011602052604060002054156140d3565b806000526011602052604060002054156103e35780600052601260205263ffffffff600260406000200154166140c857635c32f70960e01b60005260045260246000fd5b6001600160a01b039092919216918260005260096020526141f660406000206001600160401b0383169060019160005201602052604060002054151590565b6142b35782600052600660205260026040600020015492836000526012602052614224826040600020614d97565b9360048510156105f0576002851461428c5760048510156105f0576003851461426d5750600052600a6020526001600160401b0360406000209116600052602052604060002090565b9150506000526013602052604060002060008052602052604060002090565b905060005260136020526001600160401b0360406000209116600052602052604060002090565b9091600052600a6020526001600160401b036040600020911660005260205260406000209060019190565b916001600160401b036142f391168093614bd2565b156142fc575050565b6001600160a01b039063507f5f5760e11b6000521660045260245260446000fd5b91906001600160a01b031680600052600560205260406000205415612e3457506001600160401b031680156143585761435591614d15565b50565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b9061438d8180613b07565b906001600160401b0382116136a4576143a68285613a93565b83600052602060002060005b83811061445257505050506143ca6020820182613b07565b9060018401906001600160401b0383116136a4576143e88383613a93565b90600052602060002060005b8381106144375750505050600261441560406001600160a01b03930161389d565b9201911673ffffffffffffffffffffffffffffffffffffffff19825416179055565b60019060206144458561389d565b94019381840155016143f4565b60019060206144608561389d565b94019381840155016143b2565b6001600160a01b03168060005260076020526144a760406000206001600160401b0384169060019160005201602052604060002054151590565b61450f57600052600660205260016040600020015480600052600d6020526144d3826040600020614d97565b9160048310156105f05760028314614508575060048210156105f05760038214614501575090600090600090565b9091600090565b9192909190565b5060019160009190565b6001600160a01b0360015416330361452d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6001600160a01b039093929361456b61380b565b501660005260086020526001600160401b03604060002091166000526020526040600020600080526020526145a460006040812061395f565b9180614692575b5050604051906145ba826136f0565b80546001600160801b03811683526001602084019263ffffffff8360801c16845260ff604086019360a01c16151583520154906001600160801b038060608601928185168452608087019460801c855261461261380b565b508180855116614643828a511661463d61463363ffffffff8c511642614abc565b858b511690613a36565b90614f28565b8082101561468b57505b1680975263ffffffff42168096525115159251169251169260405194614672866136f0565b8552602085015260408401526060830152608082015290565b905061464d565b600052600f6020526001600160401b03604060002091166000526020526040600020600080526020526146c960006040812061395f565b6146d381836151ce565b156146de57806145ab565b80915054906001600160801b036001818460081c16920154169060ff60405193614707856136f0565b82855263ffffffff4216602086015216151560408401526060830152608082015290565b6001600160a01b039093929361473f61380b565b501660005260086020526001600160401b03604060002091166000526020526040600020600080526020526147796001604060002061395f565b918061478e575050604051906145ba826136f0565b600052600f6020526001600160401b03604060002091166000526020526040600020600080526020526146c96001604060002061395f565b6001600160a01b03909392936147da61380b565b501660005260086020526001600160401b03604060002091166000526020526040600020600160005260205261481460006040812061395f565b9180614829575050604051906145ba826136f0565b600052600f6020526001600160401b0360406000209116600052602052604060002060016000526020526146c960006040812061395f565b6001600160a01b039093929361487561380b565b501660005260086020526001600160401b0360406000209116600052602052604060002060016000526020526148b06001604060002061395f565b91806148c5575050604051906145ba826136f0565b600052600f6020526001600160401b0360406000209116600052602052604060002060016000526020526146c96001604060002061395f565b815191929115614a29576001600160801b036040840151166001600160801b03602085015116106149d15761353991925b805182546020830151604090930151608090811b6fffffffffffffffffffffffffffffffff19166001600160801b039490941693841760018601557fffffffffffffffffffffff00000000000000000000000000000000000000000090911691151560a01b74ff000000000000000000000000000000000000000016919091179091174290911b73ffffffff0000000000000000000000000000000016179055565b6064836108bd604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b6001600160801b0360408401511615801590614aa6575b614a4e57613539919261492f565b6064836108bd604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b036020840151161515614a40565b91908203918211613a4957565b80548210156138875760005260206000200190600090565b80548015614b0a576000190190614af88282614ac9565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600560205260409020548015614bcb576000198101818111613a4957600454600019810191908211613a4957818103614b7a575b505050614b666004614ae1565b600052600560205260006040812055600190565b614bb3614b8b614b9c936004614ac9565b90549060031b1c9283926004614ac9565b819391549060031b91821b91600019901b19161790565b90556000526005602052604060002055388080614b59565b5050600090565b906001820191816000528260205260406000205490811515600014614c6957600019820191808311613a495781546000198101908111613a49578381614c209503614c32575b505050614ae1565b60005260205260006040812055600190565b614c52614c42614b9c9386614ac9565b90549060031b1c92839286614ac9565b905560005284602052604060002055388080614c18565b50505050600090565b805490680100000000000000008210156136a45781614b9c916001614c9994018155614ac9565b9055565b80600052600560205260406000205415600014614cd657614cbf816004614c72565b600454906000526005602052604060002055600190565b50600090565b80600052600360205260406000205415600014614cd657614cfe816002614c72565b600254906000526003602052604060002055600190565b6000828152600182016020526040902054614bcb5780614d3783600193614c72565b80549260005201602052604060002055600190565b906040519182815491828252602082019060005260206000209260005b818110614d7e5750506135399250038361370b565b8454835260019485019487945060209093019201614d69565b906001600160401b03168015614bcb576000908152600182016020526040902054614ddb576001906000805201602052604060002054614dd657600090565b600390565b50600290565b9160ff93918015614e2a57614e259350600052600f6020526001600160401b036040600020911660005260205260406000206001600052602052604060002061395f565b541690565b506001600160a01b03614e6c931660005260086020526001600160401b036040600020911660005260205260406000206001600052602052604060002061395f565b5460a01c1690565b6001600160a01b039095929493951660005260086020526001600160401b036040600020911660005260205260406000209015159081600052602052614ebe82604060002061395f565b938315614f2157614efc93600052600f6020526001600160401b0360406000209116600052602052604060002090600052602052604060002061395f565b614f0681836151ce565b15614f0f575090565b614f1b61347591613f6b565b826148fe565b5050505090565b91908201809211613a4957565b9182549060ff8260a01c16158015615123575b61511d576001600160801b0382169160018501908154614f7b63ffffffff6001600160801b0383169360801c1642614abc565b908161507f575b50508481106150405750838310614fc4575050614fa86001600160801b03928392614abc565b16166fffffffffffffffffffffffffffffffff19825416179055565b9190915460801c9283156150185781614fdc91614abc565b60001984019390818511613a49576001600160a01b0394614ffc91614f28565b636864691d60e11b600052046004526024521660445260646000fd5b6001600160a01b038383636864691d60e11b6000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116150f35761509a9261463d9160801c90613a36565b808410156150ee5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880614f82565b6150a5565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215614f48565b805160005b81811061513c57505050565b60018101808211613a49575b8281106151585750600101615130565b6001600160a01b0361516a83866138c5565b51166001600160a01b0361517e83876138c5565b51161461518d57600101615148565b6001600160a01b0361519f83866138c5565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9060ff825460a01c169181549260ff84161515901515149283615213575b50826151f757505090565b6001600160801b0391925060018091015460801c920154161490565b9092506001600160801b03806001850154169160081c161491386151ec565b6000818152600360205260409020548015614bcb576000198101818111613a4957600254600019810191908211613a495780820361528c575b5050506152786002614ae1565b600052600360205260006040812055600190565b6152ae61529d614b9c936002614ac9565b90549060031b1c9283926002614ac9565b9055600052600360205260406000205538808061526b56fea164736f6c634300081a000a",
}

var SecurityPolicyStoreABI = SecurityPolicyStoreMetaData.ABI

var SecurityPolicyStoreBin = SecurityPolicyStoreMetaData.Bin

func DeploySecurityPolicyStore(auth *bind.TransactOpts, backend bind.ContractBackend, authorizedCallers []common.Address) (common.Address, *types.Transaction, *SecurityPolicyStore, error) {
	parsed, err := SecurityPolicyStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SecurityPolicyStoreBin), backend, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecurityPolicyStore{address: address, abi: *parsed, SecurityPolicyStoreCaller: SecurityPolicyStoreCaller{contract: contract}, SecurityPolicyStoreTransactor: SecurityPolicyStoreTransactor{contract: contract}, SecurityPolicyStoreFilterer: SecurityPolicyStoreFilterer{contract: contract}}, nil
}

type SecurityPolicyStore struct {
	address common.Address
	abi     abi.ABI
	SecurityPolicyStoreCaller
	SecurityPolicyStoreTransactor
	SecurityPolicyStoreFilterer
}

type SecurityPolicyStoreCaller struct {
	contract *bind.BoundContract
}

type SecurityPolicyStoreTransactor struct {
	contract *bind.BoundContract
}

type SecurityPolicyStoreFilterer struct {
	contract *bind.BoundContract
}

type SecurityPolicyStoreSession struct {
	Contract     *SecurityPolicyStore
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SecurityPolicyStoreCallerSession struct {
	Contract *SecurityPolicyStoreCaller
	CallOpts bind.CallOpts
}

type SecurityPolicyStoreTransactorSession struct {
	Contract     *SecurityPolicyStoreTransactor
	TransactOpts bind.TransactOpts
}

type SecurityPolicyStoreRaw struct {
	Contract *SecurityPolicyStore
}

type SecurityPolicyStoreCallerRaw struct {
	Contract *SecurityPolicyStoreCaller
}

type SecurityPolicyStoreTransactorRaw struct {
	Contract *SecurityPolicyStoreTransactor
}

func NewSecurityPolicyStore(address common.Address, backend bind.ContractBackend) (*SecurityPolicyStore, error) {
	abi, err := abi.JSON(strings.NewReader(SecurityPolicyStoreABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSecurityPolicyStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStore{address: address, abi: abi, SecurityPolicyStoreCaller: SecurityPolicyStoreCaller{contract: contract}, SecurityPolicyStoreTransactor: SecurityPolicyStoreTransactor{contract: contract}, SecurityPolicyStoreFilterer: SecurityPolicyStoreFilterer{contract: contract}}, nil
}

func NewSecurityPolicyStoreCaller(address common.Address, caller bind.ContractCaller) (*SecurityPolicyStoreCaller, error) {
	contract, err := bindSecurityPolicyStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreCaller{contract: contract}, nil
}

func NewSecurityPolicyStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*SecurityPolicyStoreTransactor, error) {
	contract, err := bindSecurityPolicyStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreTransactor{contract: contract}, nil
}

func NewSecurityPolicyStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*SecurityPolicyStoreFilterer, error) {
	contract, err := bindSecurityPolicyStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreFilterer{contract: contract}, nil
}

func bindSecurityPolicyStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SecurityPolicyStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecurityPolicyStore.Contract.SecurityPolicyStoreCaller.contract.Call(opts, result, method, params...)
}

func (_SecurityPolicyStore *SecurityPolicyStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.SecurityPolicyStoreTransactor.contract.Transfer(opts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.SecurityPolicyStoreTransactor.contract.Transact(opts, method, params...)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecurityPolicyStore.Contract.contract.Call(opts, result, method, params...)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.contract.Transfer(opts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.contract.Transact(opts, method, params...)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "checkLaneRecords", localToken, remoteChainSelector)

	if err != nil {
		return err
	}

	return err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _SecurityPolicyStore.Contract.CheckLaneRecords(&_SecurityPolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _SecurityPolicyStore.Contract.CheckLaneRecords(&_SecurityPolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SecurityPolicyStore.Contract.GetAllAuthorizedCallers(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SecurityPolicyStore.Contract.GetAllAuthorizedCallers(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetCurrentRateLimiterStateBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getCurrentRateLimiterStateBatch", lanes)

	if err != nil {
		return *new([]IPoolConfigManagerLaneRateLimiterState), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneRateLimiterState)).(*[]IPoolConfigManagerLaneRateLimiterState)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetCurrentRateLimiterStateBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	return _SecurityPolicyStore.Contract.GetCurrentRateLimiterStateBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetCurrentRateLimiterStateBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	return _SecurityPolicyStore.Contract.GetCurrentRateLimiterStateBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetInlineRateLimitConfigs(opts *bind.CallOpts, localToken common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getInlineRateLimitConfigs", localToken)

	if err != nil {
		return *new([]SecurityPolicyStoreInlineRateLimitConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]SecurityPolicyStoreInlineRateLimitConfigArgs)).(*[]SecurityPolicyStoreInlineRateLimitConfigArgs)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetInlineRateLimitConfigs(localToken common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	return _SecurityPolicyStore.Contract.GetInlineRateLimitConfigs(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetInlineRateLimitConfigs(localToken common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	return _SecurityPolicyStore.Contract.GetInlineRateLimitConfigs(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetInlineSecurityConfigs(opts *bind.CallOpts, localToken common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getInlineSecurityConfigs", localToken)

	if err != nil {
		return *new([]SecurityPolicyStoreInlineSecurityConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]SecurityPolicyStoreInlineSecurityConfigArgs)).(*[]SecurityPolicyStoreInlineSecurityConfigArgs)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetInlineSecurityConfigs(localToken common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	return _SecurityPolicyStore.Contract.GetInlineSecurityConfigs(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetInlineSecurityConfigs(localToken common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	return _SecurityPolicyStore.Contract.GetInlineSecurityConfigs(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getLaneRecordLevelBatch", lanes)

	outstruct := new(GetLaneRecordLevelBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RateLimitLevels = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.SecurityLevels = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetLaneRecordLevelBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetLaneRecordLevelBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetLaneVerificationBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneVerificationBatch,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getLaneVerificationBatch", lanes)

	outstruct := new(GetLaneVerificationBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hooks = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OutboundCCVs = *abi.ConvertType(out[1], new([][]common.Address)).(*[][]common.Address)
	outstruct.InboundCCVs = *abi.ConvertType(out[2], new([][]common.Address)).(*[][]common.Address)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetLaneVerificationBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneVerificationBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetLaneVerificationBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetLaneVerificationBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneVerificationBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetLaneVerificationBatch(&_SecurityPolicyStore.CallOpts, lanes)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getPolicyNames")

	outstruct := new(GetPolicyNames)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RateLimitPolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.SecurityPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _SecurityPolicyStore.Contract.GetPolicyNames(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _SecurityPolicyStore.Contract.GetPolicyNames(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetRateLimitPolicy(opts *bind.CallOpts, policyName [32]byte) (GetRateLimitPolicy,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getRateLimitPolicy", policyName)

	outstruct := new(GetRateLimitPolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LocalDecimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.PolicyConfigs = *abi.ConvertType(out[2], new([]SecurityPolicyStoreRateLimitPolicyConfigArgs)).(*[]SecurityPolicyStoreRateLimitPolicyConfigArgs)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetRateLimitPolicy(policyName [32]byte) (GetRateLimitPolicy,

	error) {
	return _SecurityPolicyStore.Contract.GetRateLimitPolicy(&_SecurityPolicyStore.CallOpts, policyName)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetRateLimitPolicy(policyName [32]byte) (GetRateLimitPolicy,

	error) {
	return _SecurityPolicyStore.Contract.GetRateLimitPolicy(&_SecurityPolicyStore.CallOpts, policyName)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetSecurityPolicy(opts *bind.CallOpts, policyName [32]byte) (GetSecurityPolicy,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getSecurityPolicy", policyName)

	outstruct := new(GetSecurityPolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PolicyConfigs = *abi.ConvertType(out[1], new([]SecurityPolicyStoreSecurityPolicyConfigArgs)).(*[]SecurityPolicyStoreSecurityPolicyConfigArgs)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetSecurityPolicy(policyName [32]byte) (GetSecurityPolicy,

	error) {
	return _SecurityPolicyStore.Contract.GetSecurityPolicy(&_SecurityPolicyStore.CallOpts, policyName)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetSecurityPolicy(policyName [32]byte) (GetSecurityPolicy,

	error) {
	return _SecurityPolicyStore.Contract.GetSecurityPolicy(&_SecurityPolicyStore.CallOpts, policyName)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getTokenPolicies", localToken)

	outstruct := new(GetTokenPolicies)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LocalDecimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.RateLimitPolicyName = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.SecurityPolicyName = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _SecurityPolicyStore.Contract.GetTokenPolicies(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _SecurityPolicyStore.Contract.GetTokenPolicies(&_SecurityPolicyStore.CallOpts, localToken)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) GetTokenPolicyNamesBatch(opts *bind.CallOpts, localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "getTokenPolicyNamesBatch", localTokens)

	outstruct := new(GetTokenPolicyNamesBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RateLimitPolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.SecurityPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) GetTokenPolicyNamesBatch(localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetTokenPolicyNamesBatch(&_SecurityPolicyStore.CallOpts, localTokens)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) GetTokenPolicyNamesBatch(localTokens []common.Address) (GetTokenPolicyNamesBatch,

	error) {
	return _SecurityPolicyStore.Contract.GetTokenPolicyNamesBatch(&_SecurityPolicyStore.CallOpts, localTokens)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) Owner() (common.Address, error) {
	return _SecurityPolicyStore.Contract.Owner(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) Owner() (common.Address, error) {
	return _SecurityPolicyStore.Contract.Owner(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SecurityPolicyStore.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) TypeAndVersion() (string, error) {
	return _SecurityPolicyStore.Contract.TypeAndVersion(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreCallerSession) TypeAndVersion() (string, error) {
	return _SecurityPolicyStore.Contract.TypeAndVersion(&_SecurityPolicyStore.CallOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "acceptOwnership")
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) AcceptOwnership() (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.AcceptOwnership(&_SecurityPolicyStore.TransactOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.AcceptOwnership(&_SecurityPolicyStore.TransactOpts)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_SecurityPolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_SecurityPolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyInlineRateLimitBucketUpdates(opts *bind.TransactOpts, bucketConfigsToSet []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyInlineRateLimitBucketUpdates", bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyInlineRateLimitBucketUpdates(bucketConfigsToSet []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineRateLimitBucketUpdates(&_SecurityPolicyStore.TransactOpts, bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyInlineRateLimitBucketUpdates(bucketConfigsToSet []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineRateLimitBucketUpdates(&_SecurityPolicyStore.TransactOpts, bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyInlineRateLimitConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyInlineRateLimitConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyInlineRateLimitConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineRateLimitConfigUpdates(&_SecurityPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyInlineRateLimitConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineRateLimitConfigUpdates(&_SecurityPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyInlineSecurityConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyInlineSecurityConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyInlineSecurityConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineSecurityConfigUpdates(&_SecurityPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyInlineSecurityConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyInlineSecurityConfigUpdates(&_SecurityPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyRateLimitPolicyBucketUpdates(opts *bind.TransactOpts, bucketConfigsToSet []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyRateLimitPolicyBucketUpdates", bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyRateLimitPolicyBucketUpdates(bucketConfigsToSet []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyRateLimitPolicyBucketUpdates(&_SecurityPolicyStore.TransactOpts, bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyRateLimitPolicyBucketUpdates(bucketConfigsToSet []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyRateLimitPolicyBucketUpdates(&_SecurityPolicyStore.TransactOpts, bucketConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyRateLimitPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyRateLimitPolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyRateLimitPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyRateLimitPolicyConfigUpdates(&_SecurityPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyRateLimitPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyRateLimitPolicyConfigUpdates(&_SecurityPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplySecurityPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applySecurityPolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplySecurityPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplySecurityPolicyConfigUpdates(&_SecurityPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplySecurityPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplySecurityPolicyConfigUpdates(&_SecurityPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyTokenProfileUpdates", tokensToRemove, profilesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenProfileUpdates(&_SecurityPolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenProfileUpdates(&_SecurityPolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyTokenRateLimitPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyTokenRateLimitPolicyUpdates", tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyTokenRateLimitPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenRateLimitPolicyUpdates(&_SecurityPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyTokenRateLimitPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenRateLimitPolicyUpdates(&_SecurityPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ApplyTokenSecurityPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "applyTokenSecurityPolicyUpdates", tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ApplyTokenSecurityPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenSecurityPolicyUpdates(&_SecurityPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ApplyTokenSecurityPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ApplyTokenSecurityPolicyUpdates(&_SecurityPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "transferOwnership", to)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.TransferOwnership(&_SecurityPolicyStore.TransactOpts, to)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.TransferOwnership(&_SecurityPolicyStore.TransactOpts, to)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ValidateLockOrBurn(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "validateLockOrBurn", localToken, remoteChainSelector, amount, requestedFinalityConfig)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ValidateLockOrBurn(&_SecurityPolicyStore.TransactOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ValidateLockOrBurn(&_SecurityPolicyStore.TransactOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactor) ValidateReleaseOrMint(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.contract.Transact(opts, "validateReleaseOrMint", localToken, remoteChainSelector, localAmount, requestedFinalityConfig)
}

func (_SecurityPolicyStore *SecurityPolicyStoreSession) ValidateReleaseOrMint(localToken common.Address, remoteChainSelector uint64, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ValidateReleaseOrMint(&_SecurityPolicyStore.TransactOpts, localToken, remoteChainSelector, localAmount, requestedFinalityConfig)
}

func (_SecurityPolicyStore *SecurityPolicyStoreTransactorSession) ValidateReleaseOrMint(localToken common.Address, remoteChainSelector uint64, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _SecurityPolicyStore.Contract.ValidateReleaseOrMint(&_SecurityPolicyStore.TransactOpts, localToken, remoteChainSelector, localAmount, requestedFinalityConfig)
}

type SecurityPolicyStoreAuthorizedCallerAddedIterator struct {
	Event *SecurityPolicyStoreAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreAuthorizedCallerAdded)
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
		it.Event = new(SecurityPolicyStoreAuthorizedCallerAdded)
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

func (it *SecurityPolicyStoreAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SecurityPolicyStoreAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreAuthorizedCallerAddedIterator{contract: _SecurityPolicyStore.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreAuthorizedCallerAdded)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseAuthorizedCallerAdded(log types.Log) (*SecurityPolicyStoreAuthorizedCallerAdded, error) {
	event := new(SecurityPolicyStoreAuthorizedCallerAdded)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreAuthorizedCallerRemovedIterator struct {
	Event *SecurityPolicyStoreAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreAuthorizedCallerRemoved)
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
		it.Event = new(SecurityPolicyStoreAuthorizedCallerRemoved)
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

func (it *SecurityPolicyStoreAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SecurityPolicyStoreAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreAuthorizedCallerRemovedIterator{contract: _SecurityPolicyStore.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreAuthorizedCallerRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*SecurityPolicyStoreAuthorizedCallerRemoved, error) {
	event := new(SecurityPolicyStoreAuthorizedCallerRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator struct {
	Event *SecurityPolicyStoreFastFinalityInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreFastFinalityInboundRateLimitConsumed)
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
		it.Event = new(SecurityPolicyStoreFastFinalityInboundRateLimitConsumed)
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

func (it *SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreFastFinalityInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator{contract: _SecurityPolicyStore.contract, event: "FastFinalityInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreFastFinalityInboundRateLimitConsumed)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreFastFinalityInboundRateLimitConsumed, error) {
	event := new(SecurityPolicyStoreFastFinalityInboundRateLimitConsumed)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator struct {
	Event *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed)
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
		it.Event = new(SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed)
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

func (it *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator{contract: _SecurityPolicyStore.contract, event: "FastFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed, error) {
	event := new(SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInboundRateLimitConsumedIterator struct {
	Event *SecurityPolicyStoreInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInboundRateLimitConsumed)
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
		it.Event = new(SecurityPolicyStoreInboundRateLimitConsumed)
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

func (it *SecurityPolicyStoreInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInboundRateLimitConsumedIterator{contract: _SecurityPolicyStore.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInboundRateLimitConsumed)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreInboundRateLimitConsumed, error) {
	event := new(SecurityPolicyStoreInboundRateLimitConsumed)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInlineRateLimitBucketSetIterator struct {
	Event *SecurityPolicyStoreInlineRateLimitBucketSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInlineRateLimitBucketSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInlineRateLimitBucketSet)
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
		it.Event = new(SecurityPolicyStoreInlineRateLimitBucketSet)
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

func (it *SecurityPolicyStoreInlineRateLimitBucketSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInlineRateLimitBucketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInlineRateLimitBucketSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Direction           uint8
	FastFinality        bool
	Config              RateLimiterConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInlineRateLimitBucketSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitBucketSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InlineRateLimitBucketSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInlineRateLimitBucketSetIterator{contract: _SecurityPolicyStore.contract, event: "InlineRateLimitBucketSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInlineRateLimitBucketSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitBucketSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InlineRateLimitBucketSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInlineRateLimitBucketSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitBucketSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInlineRateLimitBucketSet(log types.Log) (*SecurityPolicyStoreInlineRateLimitBucketSet, error) {
	event := new(SecurityPolicyStoreInlineRateLimitBucketSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitBucketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInlineRateLimitConfigRemovedIterator struct {
	Event *SecurityPolicyStoreInlineRateLimitConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInlineRateLimitConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInlineRateLimitConfigRemoved)
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
		it.Event = new(SecurityPolicyStoreInlineRateLimitConfigRemoved)
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

func (it *SecurityPolicyStoreInlineRateLimitConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInlineRateLimitConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInlineRateLimitConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInlineRateLimitConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InlineRateLimitConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInlineRateLimitConfigRemovedIterator{contract: _SecurityPolicyStore.contract, event: "InlineRateLimitConfigRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInlineRateLimitConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InlineRateLimitConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInlineRateLimitConfigRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitConfigRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInlineRateLimitConfigRemoved(log types.Log) (*SecurityPolicyStoreInlineRateLimitConfigRemoved, error) {
	event := new(SecurityPolicyStoreInlineRateLimitConfigRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInlineRateLimitConfigSetIterator struct {
	Event *SecurityPolicyStoreInlineRateLimitConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInlineRateLimitConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInlineRateLimitConfigSet)
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
		it.Event = new(SecurityPolicyStoreInlineRateLimitConfigSet)
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

func (it *SecurityPolicyStoreInlineRateLimitConfigSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInlineRateLimitConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInlineRateLimitConfigSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	RateLimitConfig     SecurityPolicyStoreRateLimitConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInlineRateLimitConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InlineRateLimitConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInlineRateLimitConfigSetIterator{contract: _SecurityPolicyStore.contract, event: "InlineRateLimitConfigSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInlineRateLimitConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InlineRateLimitConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInlineRateLimitConfigSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitConfigSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInlineRateLimitConfigSet(log types.Log) (*SecurityPolicyStoreInlineRateLimitConfigSet, error) {
	event := new(SecurityPolicyStoreInlineRateLimitConfigSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineRateLimitConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInlineSecurityConfigRemovedIterator struct {
	Event *SecurityPolicyStoreInlineSecurityConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInlineSecurityConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInlineSecurityConfigRemoved)
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
		it.Event = new(SecurityPolicyStoreInlineSecurityConfigRemoved)
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

func (it *SecurityPolicyStoreInlineSecurityConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInlineSecurityConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInlineSecurityConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInlineSecurityConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineSecurityConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InlineSecurityConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInlineSecurityConfigRemovedIterator{contract: _SecurityPolicyStore.contract, event: "InlineSecurityConfigRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInlineSecurityConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineSecurityConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InlineSecurityConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInlineSecurityConfigRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineSecurityConfigRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInlineSecurityConfigRemoved(log types.Log) (*SecurityPolicyStoreInlineSecurityConfigRemoved, error) {
	event := new(SecurityPolicyStoreInlineSecurityConfigRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineSecurityConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreInlineSecurityConfigSetIterator struct {
	Event *SecurityPolicyStoreInlineSecurityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreInlineSecurityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreInlineSecurityConfigSet)
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
		it.Event = new(SecurityPolicyStoreInlineSecurityConfigSet)
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

func (it *SecurityPolicyStoreInlineSecurityConfigSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreInlineSecurityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreInlineSecurityConfigSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	SecurityConfig      SecurityPolicyStoreSecurityConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterInlineSecurityConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineSecurityConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "InlineSecurityConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreInlineSecurityConfigSetIterator{contract: _SecurityPolicyStore.contract, event: "InlineSecurityConfigSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchInlineSecurityConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineSecurityConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "InlineSecurityConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreInlineSecurityConfigSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineSecurityConfigSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseInlineSecurityConfigSet(log types.Log) (*SecurityPolicyStoreInlineSecurityConfigSet, error) {
	event := new(SecurityPolicyStoreInlineSecurityConfigSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "InlineSecurityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreOutboundRateLimitConsumedIterator struct {
	Event *SecurityPolicyStoreOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreOutboundRateLimitConsumed)
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
		it.Event = new(SecurityPolicyStoreOutboundRateLimitConsumed)
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

func (it *SecurityPolicyStoreOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreOutboundRateLimitConsumedIterator{contract: _SecurityPolicyStore.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreOutboundRateLimitConsumed)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreOutboundRateLimitConsumed, error) {
	event := new(SecurityPolicyStoreOutboundRateLimitConsumed)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreOwnershipTransferRequestedIterator struct {
	Event *SecurityPolicyStoreOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreOwnershipTransferRequested)
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
		it.Event = new(SecurityPolicyStoreOwnershipTransferRequested)
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

func (it *SecurityPolicyStoreOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SecurityPolicyStoreOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreOwnershipTransferRequestedIterator{contract: _SecurityPolicyStore.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreOwnershipTransferRequested)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseOwnershipTransferRequested(log types.Log) (*SecurityPolicyStoreOwnershipTransferRequested, error) {
	event := new(SecurityPolicyStoreOwnershipTransferRequested)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreOwnershipTransferredIterator struct {
	Event *SecurityPolicyStoreOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreOwnershipTransferred)
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
		it.Event = new(SecurityPolicyStoreOwnershipTransferred)
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

func (it *SecurityPolicyStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SecurityPolicyStoreOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreOwnershipTransferredIterator{contract: _SecurityPolicyStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreOwnershipTransferred)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseOwnershipTransferred(log types.Log) (*SecurityPolicyStoreOwnershipTransferred, error) {
	event := new(SecurityPolicyStoreOwnershipTransferred)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreRateLimitPolicyBucketSetIterator struct {
	Event *SecurityPolicyStoreRateLimitPolicyBucketSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreRateLimitPolicyBucketSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreRateLimitPolicyBucketSet)
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
		it.Event = new(SecurityPolicyStoreRateLimitPolicyBucketSet)
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

func (it *SecurityPolicyStoreRateLimitPolicyBucketSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreRateLimitPolicyBucketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreRateLimitPolicyBucketSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalDecimals       uint8
	Direction           uint8
	FastFinality        bool
	Config              RateLimiterConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterRateLimitPolicyBucketSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyBucketSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "RateLimitPolicyBucketSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreRateLimitPolicyBucketSetIterator{contract: _SecurityPolicyStore.contract, event: "RateLimitPolicyBucketSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchRateLimitPolicyBucketSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyBucketSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "RateLimitPolicyBucketSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreRateLimitPolicyBucketSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyBucketSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseRateLimitPolicyBucketSet(log types.Log) (*SecurityPolicyStoreRateLimitPolicyBucketSet, error) {
	event := new(SecurityPolicyStoreRateLimitPolicyBucketSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyBucketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator struct {
	Event *SecurityPolicyStoreRateLimitPolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreRateLimitPolicyConfigRemoved)
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
		it.Event = new(SecurityPolicyStoreRateLimitPolicyConfigRemoved)
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

func (it *SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreRateLimitPolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterRateLimitPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "RateLimitPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator{contract: _SecurityPolicyStore.contract, event: "RateLimitPolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchRateLimitPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "RateLimitPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreRateLimitPolicyConfigRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyConfigRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseRateLimitPolicyConfigRemoved(log types.Log) (*SecurityPolicyStoreRateLimitPolicyConfigRemoved, error) {
	event := new(SecurityPolicyStoreRateLimitPolicyConfigRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreRateLimitPolicyConfigSetIterator struct {
	Event *SecurityPolicyStoreRateLimitPolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreRateLimitPolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreRateLimitPolicyConfigSet)
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
		it.Event = new(SecurityPolicyStoreRateLimitPolicyConfigSet)
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

func (it *SecurityPolicyStoreRateLimitPolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreRateLimitPolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreRateLimitPolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalDecimals       uint8
	RateLimitConfig     SecurityPolicyStoreRateLimitConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterRateLimitPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "RateLimitPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreRateLimitPolicyConfigSetIterator{contract: _SecurityPolicyStore.contract, event: "RateLimitPolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchRateLimitPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "RateLimitPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreRateLimitPolicyConfigSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyConfigSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseRateLimitPolicyConfigSet(log types.Log) (*SecurityPolicyStoreRateLimitPolicyConfigSet, error) {
	event := new(SecurityPolicyStoreRateLimitPolicyConfigSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreRateLimitPolicyRemovedIterator struct {
	Event *SecurityPolicyStoreRateLimitPolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreRateLimitPolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreRateLimitPolicyRemoved)
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
		it.Event = new(SecurityPolicyStoreRateLimitPolicyRemoved)
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

func (it *SecurityPolicyStoreRateLimitPolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreRateLimitPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreRateLimitPolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterRateLimitPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*SecurityPolicyStoreRateLimitPolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "RateLimitPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreRateLimitPolicyRemovedIterator{contract: _SecurityPolicyStore.contract, event: "RateLimitPolicyRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchRateLimitPolicyRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "RateLimitPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreRateLimitPolicyRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseRateLimitPolicyRemoved(log types.Log) (*SecurityPolicyStoreRateLimitPolicyRemoved, error) {
	event := new(SecurityPolicyStoreRateLimitPolicyRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "RateLimitPolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreSecurityPolicyConfigRemovedIterator struct {
	Event *SecurityPolicyStoreSecurityPolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreSecurityPolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreSecurityPolicyConfigRemoved)
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
		it.Event = new(SecurityPolicyStoreSecurityPolicyConfigRemoved)
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

func (it *SecurityPolicyStoreSecurityPolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreSecurityPolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreSecurityPolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterSecurityPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreSecurityPolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "SecurityPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreSecurityPolicyConfigRemovedIterator{contract: _SecurityPolicyStore.contract, event: "SecurityPolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchSecurityPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "SecurityPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreSecurityPolicyConfigRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyConfigRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseSecurityPolicyConfigRemoved(log types.Log) (*SecurityPolicyStoreSecurityPolicyConfigRemoved, error) {
	event := new(SecurityPolicyStoreSecurityPolicyConfigRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreSecurityPolicyConfigSetIterator struct {
	Event *SecurityPolicyStoreSecurityPolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreSecurityPolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreSecurityPolicyConfigSet)
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
		it.Event = new(SecurityPolicyStoreSecurityPolicyConfigSet)
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

func (it *SecurityPolicyStoreSecurityPolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreSecurityPolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreSecurityPolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	SecurityConfig      SecurityPolicyStoreSecurityConfig
	Raw                 types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterSecurityPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreSecurityPolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "SecurityPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreSecurityPolicyConfigSetIterator{contract: _SecurityPolicyStore.contract, event: "SecurityPolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchSecurityPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "SecurityPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreSecurityPolicyConfigSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyConfigSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseSecurityPolicyConfigSet(log types.Log) (*SecurityPolicyStoreSecurityPolicyConfigSet, error) {
	event := new(SecurityPolicyStoreSecurityPolicyConfigSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreSecurityPolicyRemovedIterator struct {
	Event *SecurityPolicyStoreSecurityPolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreSecurityPolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreSecurityPolicyRemoved)
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
		it.Event = new(SecurityPolicyStoreSecurityPolicyRemoved)
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

func (it *SecurityPolicyStoreSecurityPolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreSecurityPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreSecurityPolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterSecurityPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*SecurityPolicyStoreSecurityPolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "SecurityPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreSecurityPolicyRemovedIterator{contract: _SecurityPolicyStore.contract, event: "SecurityPolicyRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchSecurityPolicyRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "SecurityPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreSecurityPolicyRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseSecurityPolicyRemoved(log types.Log) (*SecurityPolicyStoreSecurityPolicyRemoved, error) {
	event := new(SecurityPolicyStoreSecurityPolicyRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "SecurityPolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreTokenProfileSetIterator struct {
	Event *SecurityPolicyStoreTokenProfileSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreTokenProfileSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreTokenProfileSet)
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
		it.Event = new(SecurityPolicyStoreTokenProfileSet)
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

func (it *SecurityPolicyStoreTokenProfileSetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreTokenProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreTokenProfileSet struct {
	LocalToken    common.Address
	LocalDecimals uint8
	Raw           types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*SecurityPolicyStoreTokenProfileSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreTokenProfileSetIterator{contract: _SecurityPolicyStore.contract, event: "TokenProfileSet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreTokenProfileSet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseTokenProfileSet(log types.Log) (*SecurityPolicyStoreTokenProfileSet, error) {
	event := new(SecurityPolicyStoreTokenProfileSet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreTokenRateLimitPolicySetIterator struct {
	Event *SecurityPolicyStoreTokenRateLimitPolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreTokenRateLimitPolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreTokenRateLimitPolicySet)
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
		it.Event = new(SecurityPolicyStoreTokenRateLimitPolicySet)
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

func (it *SecurityPolicyStoreTokenRateLimitPolicySetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreTokenRateLimitPolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreTokenRateLimitPolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterTokenRateLimitPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*SecurityPolicyStoreTokenRateLimitPolicySetIterator, error) {

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

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "TokenRateLimitPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreTokenRateLimitPolicySetIterator{contract: _SecurityPolicyStore.contract, event: "TokenRateLimitPolicySet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchTokenRateLimitPolicySet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenRateLimitPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "TokenRateLimitPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreTokenRateLimitPolicySet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenRateLimitPolicySet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseTokenRateLimitPolicySet(log types.Log) (*SecurityPolicyStoreTokenRateLimitPolicySet, error) {
	event := new(SecurityPolicyStoreTokenRateLimitPolicySet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenRateLimitPolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreTokenRemovedIterator struct {
	Event *SecurityPolicyStoreTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreTokenRemoved)
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
		it.Event = new(SecurityPolicyStoreTokenRemoved)
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

func (it *SecurityPolicyStoreTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreTokenRemoved struct {
	LocalToken common.Address
	Raw        types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*SecurityPolicyStoreTokenRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreTokenRemovedIterator{contract: _SecurityPolicyStore.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreTokenRemoved)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseTokenRemoved(log types.Log) (*SecurityPolicyStoreTokenRemoved, error) {
	event := new(SecurityPolicyStoreTokenRemoved)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SecurityPolicyStoreTokenSecurityPolicySetIterator struct {
	Event *SecurityPolicyStoreTokenSecurityPolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SecurityPolicyStoreTokenSecurityPolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityPolicyStoreTokenSecurityPolicySet)
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
		it.Event = new(SecurityPolicyStoreTokenSecurityPolicySet)
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

func (it *SecurityPolicyStoreTokenSecurityPolicySetIterator) Error() error {
	return it.fail
}

func (it *SecurityPolicyStoreTokenSecurityPolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SecurityPolicyStoreTokenSecurityPolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) FilterTokenSecurityPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*SecurityPolicyStoreTokenSecurityPolicySetIterator, error) {

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

	logs, sub, err := _SecurityPolicyStore.contract.FilterLogs(opts, "TokenSecurityPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &SecurityPolicyStoreTokenSecurityPolicySetIterator{contract: _SecurityPolicyStore.contract, event: "TokenSecurityPolicySet", logs: logs, sub: sub}, nil
}

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) WatchTokenSecurityPolicySet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenSecurityPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SecurityPolicyStore.contract.WatchLogs(opts, "TokenSecurityPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SecurityPolicyStoreTokenSecurityPolicySet)
				if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenSecurityPolicySet", log); err != nil {
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

func (_SecurityPolicyStore *SecurityPolicyStoreFilterer) ParseTokenSecurityPolicySet(log types.Log) (*SecurityPolicyStoreTokenSecurityPolicySet, error) {
	event := new(SecurityPolicyStoreTokenSecurityPolicySet)
	if err := _SecurityPolicyStore.contract.UnpackLog(event, "TokenSecurityPolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetLaneRecordLevelBatch struct {
	RateLimitLevels []uint8
	SecurityLevels  []uint8
}
type GetLaneVerificationBatch struct {
	Hooks        []common.Address
	OutboundCCVs [][]common.Address
	InboundCCVs  [][]common.Address
}
type GetPolicyNames struct {
	RateLimitPolicyNames [][32]byte
	SecurityPolicyNames  [][32]byte
}
type GetRateLimitPolicy struct {
	TokenCount    uint32
	LocalDecimals uint8
	PolicyConfigs []SecurityPolicyStoreRateLimitPolicyConfigArgs
}
type GetSecurityPolicy struct {
	TokenCount    uint32
	PolicyConfigs []SecurityPolicyStoreSecurityPolicyConfigArgs
}
type GetTokenPolicies struct {
	IsRegistered        bool
	LocalDecimals       uint8
	RateLimitPolicyName [32]byte
	SecurityPolicyName  [32]byte
}
type GetTokenPolicyNamesBatch struct {
	RateLimitPolicyNames [][32]byte
	SecurityPolicyNames  [][32]byte
}

func (SecurityPolicyStoreAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (SecurityPolicyStoreAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (SecurityPolicyStoreFastFinalityInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f9")
}

func (SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e")
}

func (SecurityPolicyStoreInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SecurityPolicyStoreInlineRateLimitBucketSet) Topic() common.Hash {
	return common.HexToHash("0x1e0a61edb07c831e721f1d19d6285ebad9f6a7847fc7e7597d7fb2f35a7dffde")
}

func (SecurityPolicyStoreInlineRateLimitConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x2909773a1782f3824d263d8654775762ceb2c33d5919a546083db42915205c80")
}

func (SecurityPolicyStoreInlineRateLimitConfigSet) Topic() common.Hash {
	return common.HexToHash("0x25b551e877c0171c7b8721a3973bd6ffdb14498f117775c6e16ed739e704e2f6")
}

func (SecurityPolicyStoreInlineSecurityConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x9c8ab32bba343fe78edb9a8fc032ad81100c097e58e133bd46cd2cadd5bb633b")
}

func (SecurityPolicyStoreInlineSecurityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0f6e7d77fde88adeb68e2ceed23ca44c93b917ae178cd836ce8a81a41c1ad854")
}

func (SecurityPolicyStoreOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (SecurityPolicyStoreOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SecurityPolicyStoreOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SecurityPolicyStoreRateLimitPolicyBucketSet) Topic() common.Hash {
	return common.HexToHash("0xbf5d213dcd6150f8deb1dae200784d598e8895a1f9413c91733d415730b44eb6")
}

func (SecurityPolicyStoreRateLimitPolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x38b4904e409fb3862e7e48aa073f3f2a0d0af85607884cdf07b89b5eaf1819c2")
}

func (SecurityPolicyStoreRateLimitPolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa5267cb30eb155107fe8fb4922632854c7ea8ece482e8b3a3e465b9b42f39210")
}

func (SecurityPolicyStoreRateLimitPolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0xa106ef2c214ebd9b86d462fe971cb93bb5f81d3807ac052940dc6070007e521a")
}

func (SecurityPolicyStoreSecurityPolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0xfd80d25f2d85a84586e465477f73b38272015e894d7915ca43708a8ff77665c3")
}

func (SecurityPolicyStoreSecurityPolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf772dd29887be2365b45f9fa9c3428a0756ba16872f56b1ff296c97f1ccb0cc1")
}

func (SecurityPolicyStoreSecurityPolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0x55cff892eb89c654e76e11e2f4d53456cd9c03f313f32e50c6b035b34fd357ec")
}

func (SecurityPolicyStoreTokenProfileSet) Topic() common.Hash {
	return common.HexToHash("0x45669b1c1be2ec7c5e8638c7b26de467615f4afbc4b701c647b23236e1222a4c")
}

func (SecurityPolicyStoreTokenRateLimitPolicySet) Topic() common.Hash {
	return common.HexToHash("0xc0d2b9fef1d1d030378ff1ee5fb1d9ffcdd755598d86e1de88cc92cb5eb36791")
}

func (SecurityPolicyStoreTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3")
}

func (SecurityPolicyStoreTokenSecurityPolicySet) Topic() common.Hash {
	return common.HexToHash("0xbdf2e0aae6b8f6b68d06d1a004b760b404a2e6cb361bd9dc3a255a508760f14a")
}

func (_SecurityPolicyStore *SecurityPolicyStore) Address() common.Address {
	return _SecurityPolicyStore.address
}

type SecurityPolicyStoreInterface interface {
	CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetCurrentRateLimiterStateBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error)

	GetInlineRateLimitConfigs(opts *bind.CallOpts, localToken common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error)

	GetInlineSecurityConfigs(opts *bind.CallOpts, localToken common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error)

	GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

		error)

	GetLaneVerificationBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneVerificationBatch,

		error)

	GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

		error)

	GetRateLimitPolicy(opts *bind.CallOpts, policyName [32]byte) (GetRateLimitPolicy,

		error)

	GetSecurityPolicy(opts *bind.CallOpts, policyName [32]byte) (GetSecurityPolicy,

		error)

	GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

		error)

	GetTokenPolicyNamesBatch(opts *bind.CallOpts, localTokens []common.Address) (GetTokenPolicyNamesBatch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyInlineRateLimitBucketUpdates(opts *bind.TransactOpts, bucketConfigsToSet []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error)

	ApplyInlineRateLimitConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error)

	ApplyInlineSecurityConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error)

	ApplyRateLimitPolicyBucketUpdates(opts *bind.TransactOpts, bucketConfigsToSet []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error)

	ApplyRateLimitPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error)

	ApplySecurityPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error)

	ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error)

	ApplyTokenRateLimitPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenSecurityPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	ValidateLockOrBurn(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	ValidateReleaseOrMint(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, localAmount *big.Int, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SecurityPolicyStoreAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*SecurityPolicyStoreAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SecurityPolicyStoreAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*SecurityPolicyStoreAuthorizedCallerRemoved, error)

	FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreFastFinalityInboundRateLimitConsumedIterator, error)

	WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreFastFinalityInboundRateLimitConsumed, error)

	FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreFastFinalityOutboundRateLimitConsumedIterator, error)

	WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreFastFinalityOutboundRateLimitConsumed, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreInboundRateLimitConsumed, error)

	FilterInlineRateLimitBucketSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitBucketSetIterator, error)

	WatchInlineRateLimitBucketSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitBucketSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineRateLimitBucketSet(log types.Log) (*SecurityPolicyStoreInlineRateLimitBucketSet, error)

	FilterInlineRateLimitConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitConfigRemovedIterator, error)

	WatchInlineRateLimitConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineRateLimitConfigRemoved(log types.Log) (*SecurityPolicyStoreInlineRateLimitConfigRemoved, error)

	FilterInlineRateLimitConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineRateLimitConfigSetIterator, error)

	WatchInlineRateLimitConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineRateLimitConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineRateLimitConfigSet(log types.Log) (*SecurityPolicyStoreInlineRateLimitConfigSet, error)

	FilterInlineSecurityConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineSecurityConfigRemovedIterator, error)

	WatchInlineSecurityConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineSecurityConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineSecurityConfigRemoved(log types.Log) (*SecurityPolicyStoreInlineSecurityConfigRemoved, error)

	FilterInlineSecurityConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*SecurityPolicyStoreInlineSecurityConfigSetIterator, error)

	WatchInlineSecurityConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreInlineSecurityConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineSecurityConfigSet(log types.Log) (*SecurityPolicyStoreInlineSecurityConfigSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SecurityPolicyStoreOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*SecurityPolicyStoreOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SecurityPolicyStoreOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SecurityPolicyStoreOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SecurityPolicyStoreOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SecurityPolicyStoreOwnershipTransferred, error)

	FilterRateLimitPolicyBucketSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyBucketSetIterator, error)

	WatchRateLimitPolicyBucketSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyBucketSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitPolicyBucketSet(log types.Log) (*SecurityPolicyStoreRateLimitPolicyBucketSet, error)

	FilterRateLimitPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyConfigRemovedIterator, error)

	WatchRateLimitPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitPolicyConfigRemoved(log types.Log) (*SecurityPolicyStoreRateLimitPolicyConfigRemoved, error)

	FilterRateLimitPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreRateLimitPolicyConfigSetIterator, error)

	WatchRateLimitPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitPolicyConfigSet(log types.Log) (*SecurityPolicyStoreRateLimitPolicyConfigSet, error)

	FilterRateLimitPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*SecurityPolicyStoreRateLimitPolicyRemovedIterator, error)

	WatchRateLimitPolicyRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreRateLimitPolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseRateLimitPolicyRemoved(log types.Log) (*SecurityPolicyStoreRateLimitPolicyRemoved, error)

	FilterSecurityPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreSecurityPolicyConfigRemovedIterator, error)

	WatchSecurityPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSecurityPolicyConfigRemoved(log types.Log) (*SecurityPolicyStoreSecurityPolicyConfigRemoved, error)

	FilterSecurityPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*SecurityPolicyStoreSecurityPolicyConfigSetIterator, error)

	WatchSecurityPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSecurityPolicyConfigSet(log types.Log) (*SecurityPolicyStoreSecurityPolicyConfigSet, error)

	FilterSecurityPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*SecurityPolicyStoreSecurityPolicyRemovedIterator, error)

	WatchSecurityPolicyRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreSecurityPolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseSecurityPolicyRemoved(log types.Log) (*SecurityPolicyStoreSecurityPolicyRemoved, error)

	FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*SecurityPolicyStoreTokenProfileSetIterator, error)

	WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error)

	ParseTokenProfileSet(log types.Log) (*SecurityPolicyStoreTokenProfileSet, error)

	FilterTokenRateLimitPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*SecurityPolicyStoreTokenRateLimitPolicySetIterator, error)

	WatchTokenRateLimitPolicySet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenRateLimitPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenRateLimitPolicySet(log types.Log) (*SecurityPolicyStoreTokenRateLimitPolicySet, error)

	FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*SecurityPolicyStoreTokenRemovedIterator, error)

	WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error)

	ParseTokenRemoved(log types.Log) (*SecurityPolicyStoreTokenRemoved, error)

	FilterTokenSecurityPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*SecurityPolicyStoreTokenSecurityPolicySetIterator, error)

	WatchTokenSecurityPolicySet(opts *bind.WatchOpts, sink chan<- *SecurityPolicyStoreTokenSecurityPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenSecurityPolicySet(log types.Log) (*SecurityPolicyStoreTokenSecurityPolicySet, error)

	Address() common.Address
}
