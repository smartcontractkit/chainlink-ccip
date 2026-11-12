// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool_config_manager

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

type ConnectionPolicyStoreConnectionPolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	OnlySubsetTokens    bool
}

type ConnectionPolicyStoreInlineConnectionConfigArgs struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	RemoteToken         []byte
}

type ConnectionPolicyStoreInlineRemotePoolConfigArgs struct {
	LocalToken             common.Address
	RemoteChainSelector    uint64
	TrustPolicyRemotePools bool
	RemotePoolConfig       ConnectionPolicyStoreRemotePoolConfig
}

type ConnectionPolicyStoreRemotePoolConfig struct {
	UseLegacyPoolDataOnSend bool
	RemotePools             []IPoolConfigManagerRemotePool
}

type ConnectionPolicyStoreRemotePoolPolicyConfigArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	RemotePoolAddresses [][]byte
}

type ConnectionPolicyStoreRemoteTokenAddressProviderArgs struct {
	PolicyName                 [32]byte
	RemoteTokenAddressProvider common.Address
}

type ConnectionPolicyStoreSubsetTokenArgs struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalToken          common.Address
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

type IPoolConfigManagerLaneConnection struct {
	HasConnection           bool
	IsOutboundPaused        bool
	IsInboundPaused         bool
	IsGloballyPaused        bool
	IsTokenOutboundPaused   bool
	IsTokenInboundPaused    bool
	IsChainOutboundPaused   bool
	IsChainInboundPaused    bool
	UseLegacyPoolDataOnSend bool
	RemoteToken             []byte
	RemotePools             []IPoolConfigManagerRemotePool
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

type IPoolConfigManagerLaneRateLimiterState struct {
	Outbound             IPoolConfigManagerTokenBucket
	Inbound              IPoolConfigManagerTokenBucket
	FastFinalityOutbound IPoolConfigManagerTokenBucket
	FastFinalityInbound  IPoolConfigManagerTokenBucket
}

type IPoolConfigManagerOutboundLane struct {
	RemoteToken       []byte
	DestPoolData      []byte
	LocalDecimals     uint8
	LockOrBurnFlavor  uint8
	LockOrBurnAddress common.Address
	FeeAmount         *big.Int
	Hook              common.Address
}

type IPoolConfigManagerRemotePool struct {
	RemotePoolAddress            []byte
	AllowLegacyPoolDataOnReceive bool
}

type IPoolConfigManagerResolvedInboundTransfer struct {
	LocalAmount       *big.Int
	LockOrBurnFlavor  uint8
	LockOrBurnAddress common.Address
	Hook              common.Address
}

type IPoolConfigManagerTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

type IPoolConfigManagerTokenConfig struct {
	IsRegistered               bool
	LocalDecimals              uint8
	LockOrBurnFlavor           uint8
	AllowedFinalityConfig      [4]byte
	LockOrBurnAddress          common.Address
	AdditionalPoolDataProvider common.Address
	SupportedChains            []uint64
	ConnectionPolicyName       [32]byte
	RemotePoolPolicyName       [32]byte
	RateLimitPolicyName        [32]byte
	FeePolicyName              [32]byte
	DestOverheadPolicyName     [32]byte
	SecurityPolicyName         [32]byte
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

var PoolConfigManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"connectionStore\",\"type\":\"address\",\"internalType\":\"contract ConnectionPolicyStore\"},{\"name\":\"feeStore\",\"type\":\"address\",\"internalType\":\"contract FeePolicyStore\"},{\"name\":\"securityStore\",\"type\":\"address\",\"internalType\":\"contract SecurityPolicyStore\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONFIG_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONNECTION_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEST_OVERHEAD_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEE_COLLECTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POLICY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMIT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REMOTE_POOL_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SECURITY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNPAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyConnectionPolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.ConnectionPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestOverheadPolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.DestOverheadPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFeePolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.FeePolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineConnectionConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineConnectionConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineDestOverheadConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineDestOverheadConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineFeeConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineFeeConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRateLimitBucketUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitBucketArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRateLimitConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRemotePoolConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineRemotePoolConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"trustPolicyRemotePools\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePoolConfig\",\"type\":\"tuple\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolConfig\",\"components\":[{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineSecurityConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineSecurityConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRateLimitPolicyBucketUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyBucketArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRateLimitPolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemotePoolPolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteTokenAddressProviderUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemoteTokenAddressProviderArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenAddressProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySecurityPolicyConfigUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.SecurityPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySubsetTokenUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.SubsetTokenArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"tokensToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.SubsetTokenArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenConnectionPolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenDestOverheadPolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenFeePolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenProfileUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenProfileArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenRateLimitPolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenRemotePoolPolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenSecurityPolicyUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllConfiguredTokens\",\"inputs\":[{\"name\":\"startIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConnectionPolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.ConnectionPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterStateBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"states\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneRateLimiterState[]\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestOverheadPolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.DestOverheadPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"fees\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneFee[]\",\"components\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeePolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.FeePolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineConnectionConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineConnectionConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineDestOverheadConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineDestOverheadConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destOverheadConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.DestOverheadConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineFeeConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct FeePolicyStore.InlineFeeConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct FeePolicyStore.FeeConfig\",\"components\":[{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineRateLimitConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineRateLimitConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineRemotePoolConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineRemotePoolConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"trustPolicyRemotePools\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePoolConfig\",\"type\":\"tuple\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolConfig\",\"components\":[{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineSecurityConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.InlineSecurityConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneConnectionBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"connections\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneConnection[]\",\"components\":[{\"name\":\"hasConnection\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isGloballyPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isChainOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isChainInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneRecordLevelBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"connectionLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"remotePoolLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"rateLimitLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"feeLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"destOverheadLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"securityLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneVerificationBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"hooks\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVs\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyNames\",\"inputs\":[],\"outputs\":[{\"name\":\"connectionPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"remotePoolPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"rateLimitPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"feePolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"destOverheadPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"securityPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitPolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.RateLimitPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rateLimitConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.RateLimitConfig\",\"components\":[{\"name\":\"outbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityOutbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"fastFinalityInbound\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePoolPolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSecurityPolicy\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct SecurityPolicyStore.SecurityPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"securityConfig\",\"type\":\"tuple\",\"internalType\":\"struct SecurityPolicyStore.SecurityConfig\",\"components\":[{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"connectionStore\",\"type\":\"address\",\"internalType\":\"contract ConnectionPolicyStore\"},{\"name\":\"feeStore\",\"type\":\"address\",\"internalType\":\"contract FeePolicyStore\"},{\"name\":\"securityStore\",\"type\":\"address\",\"internalType\":\"contract SecurityPolicyStore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubsetTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAdminRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenConfigBatch\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.TokenConfig[]\",\"components\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"supportedChains\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"connectionPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remotePoolPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rateLimitPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"feePolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destOverheadPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"securityPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfigBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"feeConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig[]\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseChains\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseGlobally\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseTokens\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpauseChains\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseGlobally\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTokens\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateLockOrBurn\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"chargeTokenFee\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"lane\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.OutboundLane\",\"components\":[{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateReleaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"transfer\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.ResolvedInboundTransfer\",\"components\":[{\"name\":\"localAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contract IMultiTokenPool\"},{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract ITokenAdminRegistry\"},{\"name\":\"connectionStore\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract ConnectionPolicyStore\"},{\"name\":\"feeStore\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract FeePolicyStore\"},{\"name\":\"securityStore\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract SecurityPolicyStore\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTokenPool\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolConfigManagerMismatch\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolConfigManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x6101003461033f57601f615f3538819003918201601f19168301916001600160401b038311848410176103445780849260a09460405283398101031261033f5780516001600160a01b03811680820361033f576020830151926001600160a01b03841680850361033f5760408201516001600160a01b038116919082810361033f5760608401516001600160a01b0381169490919085830361033f57608001516001600160a01b0381169081810361033f57811561032957600180546001600160d01b03169055600254916001600160a01b038316610318576100ee9260018060a01b0319161760025561035a565b5085158015610310575b8015610308575b8015610300575b6102ef577fecdad3df04e9354405258a50ba2fb49a154c3fd0cef28c42959c045eafcf629b97608097885260a05260c05260e052604051938452602084015260408301526060820152a1604051615b2c90816103e98239608051818181612f5e01528181613fba0152615662015260a0518181816104f9015281816108ae01528181610b2501528181610b6d01528181610f47015281816112d2015281816113d30152818161150901528181611ac601528181611b7601528181611f3a015281816121df015281816122d5015281816124ea015281816125df0152818161268d01528181612c1901528181612ea101528181612f9e015281816133d00152818161391301528181613e1601528181613fdf0152818161414401528181614e7401526154c8015260c05181818161057701528181610a6e01528181610d00015281816114660152818161156b01528181611f880152818161233a015281816130630152818161343301528181613e4a0152818161400701528181614df40152615531015260e0518181816105e1015281816109c4015281816115d201528181611bdb01528181611dba01528181611fd7015281816123a00152818161255901528181612c8501528181613e7a01528181614031015281816141a901528181614b6901526155960152f35b630a64406560e11b60005260046000fd5b508415610106565b5083156100ff565b5082156100f8565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381166000908152600080516020615f15833981519152602052604090205460ff166103e2576001600160a01b03166000818152600080516020615f1583398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c908162b02f1a146140ec5750806301ffc9a71461407b578063022d63fb1461405d57806306285c6914613f9b5780630678bd7614613f6a5780630aa6220b14613ed55780630f1e1f5614613d4d5780630fb066b314613d24578063164a1114146138b6578063181f5a77146138595780631dce24e6146125455780632218f578146137f857806322985246146137bd578063248a9ca3146137875780632a0276f81461375e5780632f2ff15d146137155780632fc6499d146136da57806332ecd375146136ae57806334dbc5001461337457806336568abe1461328f57806339891a3e14612f8d5780633bab381f14612f48578063412e593e14610a515780634544b02114612f0d57806345a9683f14612e0e57806345da866714612df357806347baacdd14612dca5780634b910d1d14612afe57806351b910ca14612aac57806362a2a47c14612a71578063634e93da1461299b578063649a5ec7146127fe57806364b187ea1461278b5780636576b86914611d9d5780636e293a23146112b55780636f9b7fcc1461275057806371765b561461267257806372d64a811461258657806373ab65491461254557806377a14de9146124ce5780637a338ed3146114525780637b5f0d5f146122795780637ceb83221461223e57806382654ec7146121c157806384ef8ffc14611eec57806388a18721146121985780638a379c4a146120fc5780638c63fe6214611f155780638da5cb5b14611eec57806391d1485414611ea15780639993974314611e66578063a1eda53c14611e03578063a217fddf14611de7578063a3ed219a14611d9d578063a4d19feb14611d62578063a4eaa81714611b3b578063a9cec75514611a73578063af8df28b14610b11578063af9e6072146114ae578063b0bf810614611493578063b1a26f7714611452578063bba4ef5014611400578063bbc6fb5e14611363578063c2840e6014611328578063c34aea40146112ff578063c9a15df1146112b5578063cc8463c81461128a578063cce3dffa1461120b578063cefc14291461115f578063cf6eefb714611124578063d37eeeed1461106a578063d547741f14611009578063d602b9fd14610fa4578063d7fdecf414610f2b578063d894ac4214610d72578063dbe1065514610d49578063de24252314610c30578063e05dfe9714610b52578063e3d9c36714610b11578063e63ab1e914610ad6578063ec8784c114610a9b578063edf5e2c314610a51578063efbe103e14610a13578063f40eb24314610952578063f80cd24b14610890578063fb1bb9de146108555763ff910457146103dd57600080fd5b346106e15760c03660031901126106e1576103f66145cc565b906103ff6145f6565b604435916064356001600160e01b031981169283820361085057608435956001600160401b03871161084c573660238801121561084c578660040135926001600160401b03841161073857366024858a0101116107385760a435978815158099036108485795809798866104f5988160c060405161047c816148c0565b60608152606060208201528260408201528260608201528260808201528260a082015201526104aa86615636565b604051809a8192632d6b6c8560e11b835260018060a01b0389169a8b60048501526001600160401b038b16978860248601526044850152608060648501526024608485019201614d41565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa97881561083d57879861074f575b5060208801988951519260405197637762c27560e11b8952600489015260248801528a60448801526064870152608486015260a485015260208460c48160018060a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa93841561074457859461070c575b5060a086019380855288039788116106f857906105dc60209392604051998a948594633dcd28e160e01b865260048601614d62565b0381857f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19182156106ec5780926106a9575b50610662945060c083019160018060a01b031682526040519485946020865261064f855160e0602089015261010088019061451f565b9051868203601f1901604088015261451f565b9260ff604082015116606086015261068260608201516080870190614748565b608001516001600160a01b0390811660a0860152905160c085015290511660e08301520390f35b9091506020853d6020116106e4575b816106c5602093836148f6565b810103126106e157506106da61066294614d2d565b9038610619565b80fd5b3d91506106b8565b604051903d90823e3d90fd5b634e487b7160e01b85526011600452602485fd5b9093506020813d60201161073c575b81610728602093836148f6565b81010312610738575192386105a7565b8480fd5b3d915061071b565b6040513d87823e3d90fd5b9097503d8088833e61076181836148f6565b810190602081830312610839578051906001600160401b038211610831570160e0818303126108395760405191610797836148c0565b81516001600160401b03811161083557816107b3918401614bb5565b835260208201516001600160401b03811161083557906107d4918301614bb5565b60208301526107e560408201615192565b604083015260608101516005811015610831576108259160c091606085015261081060808201614d2d565b608085015260a081015160a085015201614d2d565b60c08201529638610531565b8880fd5b8980fd5b8780fd5b6040513d89823e3d90fd5b8580fd5b8380fd5b600080fd5b50346106e157806003193601126106e15760206040517f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a8152f35b50346106e157806108a03661479b565b906108ac9392936152fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561084c57610918948461090d93819560405198899687958694630c9d68dd60e41b865260606004870152606486019161500e565b916024840190614ea1565b82604483015203925af180156109455761092f5780f35b610938916148f6565b806000126106e157388180f35b50604051903d90823e3d90fd5b50346106e15761096136614544565b5050336000908152600080516020615a40833981519152602052604090205460ff161593506109f792505050575b61099836614af9565b6109a560405191826148f6565b813682526109f3576109f19036600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b005b5080fd5b610a0e600080516020615ac08339815191526153df565b61098f565b50346106e15760203660031901126106e1576004356001600160401b0381116109f357610a449036906004016143f9565b5050610a4e6151a0565b80f35b50346106e15760203660031901126106e157610a6b6145cc565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e15760203660031901126106e1576004356001600160401b0381116109f357610acc9036906004016143f9565b5050610a4e614d9f565b50346106e157806003193601126106e15760206040517f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8152f35b50346106e15760203660031901126106e1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e157610b613661479b565b9091610b6b61534c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561073857928491604051948592633cdd2c1960e01b84528060648501606060048701525260848401929060005b818110610bfe57505050838383610be68294849996506024840190614ea1565b6001604483015203925af180156109455761092f5780f35b92945092819550602080916001600160401b03610c1c600195614638565b168152019501910190879492879492610bc6565b50346106e15760603660031901126106e1576004356001600160401b0381116109f357610c619036906004016143f9565b50506024356001600160401b0381116109f357610c8290369060040161449c565b50506044356001600160401b0381116109f357610ca390369060040161476b565b5050336000908152600080516020615a40833981519152602052604090205460ff1615610d2d575b610cd436614af9565b610ce160405191826148f6565b813682526109f3576109f19036600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b610d44600080516020615b008339815191526153df565b610ccb565b50346106e157806003193601126106e1576020604051600080516020615ac08339815191528152f35b50346106e15760603660031901126106e1576004356001600160a01b03811691908290036106e1576024356001600160401b0381116109f357610db990369060040161449c565b6044356001600160a01b038116929190839003610850573360009081527f92227d48d66cabccdfe59bda1e8cb95f286820615128bd3f8666d1f5ec8b70aa602052604090205460ff1615610ef2576040516302005f3f60e41b8152602081600481895afa908115610744578591610eb8575b506001600160a01b0316308103610ea1575090839185803b1561084c57839283610e749360405197889586948593630639185360e51b855260406004860152604485019161500e565b90602483015203925af1801561094557610e8b5780f35b610e94916148f6565b806000126106e157818180f35b63674f3ea160e11b85526004869052602452604484fd5b90506020813d602011610eea575b81610ed3602093836148f6565b8101031261073857610ee490614d2d565b38610e2b565b3d9150610ec6565b63e2517d3f60e01b600052336004527f2dca0f5ce7e75a4b43fe2b0d6f5d0b7a2bf92ecf89f8f0aa17b8308b6703882160245260446000fd5b50346106e157806003193601126106e157610f446152fc565b807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b15610fa1578180916024604051809581936369a6b3db60e01b83528160048401525af180156109455761092f5780f35b50fd5b50346106e157806003193601126106e157610fbd615280565b65ffffffffffff610fcc615221565b90508160d01b6001541660015516610fe15780f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a180f35b50346106e15760403660031901126106e1576004356110266145b6565b90801561105b57908161105261104d61105794600052600060205260016040600020015490565b6152c1565b615842565b5080f35b631fe1e13d60e11b8352600483fd5b50346106e15760603660031901126106e1576004356001600160401b0381116109f35761109b9036906004016143f9565b50506024356001600160401b0381116109f3576110bc90369060040161449c565b50506044356001600160401b0381116109f3576110dd9036906004016144cc565b5050336000908152600080516020615a40833981519152602052604090205460ff161561110d57610cd436614af9565b610d44600080516020615a808339815191526153df565b50346106e157806003193601126106e157604065ffffffffffff611146615221565b83516001600160a01b0390921682529091166020820152f35b50346106e157806003193601126106e157611178615221565b506001600160a01b031633036111f85765ffffffffffff611197615221565b909116801580156111ee575b6111dc57506002546111c891906111c2906001600160a01b031661580b565b50615781565b50600180546001600160d01b031916905580f35b6319ca5ebb60e01b8352600452602482fd5b50428110156111a3565b636116401160e11b815233600452602490fd5b50346106e15760403660031901126106e1576004356001600160401b0381116109f35761123c9036906004016143f9565b50506024356001600160401b0381116109f357366023820112156109f3578060040135906001600160401b038211611286576024906101c03693020101116106e157610a4e614b14565b8280fd5b50346106e157806003193601126106e15760206112a56151e8565b65ffffffffffff60405191168152f35b50346106e15760203660031901126106e1576112cf6145cc565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e157806003193601126106e1576020604051600080516020615b008339815191528152f35b50346106e157806003193601126106e15760206040517f74f7a545c65c11839a48d7453738b30c295408df2d944516167556759ddc6d068152f35b50346106e15760203660031901126106e1576004356001600160401b0381116109f3576113949036906004016143f9565b505061139e615240565b6113a736614af9565b6113b460405191826148f6565b813682526109f3576109f19036600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b50346106e15760403660031901126106e1576004356001600160401b0381116109f3576114319036906004016143f9565b50506024356001600160401b0381116109f357610a449036906004016144cc565b50346106e15760203660031901126106e1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e1576114a2366146fc565b50505050610a4e614fc1565b50346106e15760203660031901126106e1576004356001600160401b0381116109f3576114df90369060040161449c565b6040516357cf303960e11b8152602060048201529183838061150560248201868661500e565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa928315611a68578493611832575b50604051630a2ad9a960e01b81526020600482015284818061156760248201878761500e565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908582156106ec576115ce9381928294611810575b5060405180958192630a2ad9a960e01b83526020600484015288602484019161500e565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa92831561180557869087946117de575b50865b858110611769578688604051918291602083016020845282518091526040840191602060408360051b870101940192815b83831061164d5786860387f35b919395509193603f198782030183528551906101a081019180511515825260ff602082015116602083015261168a60408201516040840190614748565b63ffffffff60e01b606082015116606083015260018060a01b03608082015116608083015260018060a01b0360a08201511660a083015260c0810151926101a060c0840152835180915260206101c0840194019085905b80821061174857505050600192602092610180808460e08796015160e0850152610100810151610100850152610120810151610120850152610140810151610140850152610160810151610160850152015191015297019301930190928695949293611640565b9091946020806001926001600160401b0389511681520196019201906116e1565b8061177660019284614aa4565b51610120611784838b614aa4565b5101526117918185614aa4565b5161014061179f838b614aa4565b5101526117ac8186614aa4565b516101606117ba838b614aa4565b5101526117c78187614aa4565b516101806117d5838b614aa4565b5101520161160f565b90506117fd9193503d8088833e6117f581836148f6565b8101906150ab565b92903861160c565b6040513d88823e3d90fd5b9093506118289192503d8084833e6117f581836148f6565b91909192386115aa565b9092503d8085833e61184481836148f6565b810190602081830312610738578051906001600160401b038211610848570181601f820112156107385780519061187a82614917565b9261188860405194856148f6565b82845260208085019360051b830101918183116108395760208101935b8385106118b85750505050509138611541565b84516001600160401b0381116108355782016101a0818503601f19011261083557604051906101a082018281106001600160401b03821117611a52576040526119036020820161492e565b825261191160408201615192565b602083015260608101516005811015611a4e57604083015260808101516001600160e01b031981168103611a4e57606083015261195060a08201614d2d565b608083015261196160c08201614d2d565b60a083015260e08101516001600160401b038111611a4e5760209082010185601f82011215611a4e57805161199581614917565b916119a360405193846148f6565b81835260208084019260051b82010190888211611a4a57602001915b818310611a2657505050916101a060209492859460c084015261010081015160e084015261012081015161010084015261014081015161012084015261016081015161014084015261018081015161016084015201516101808201528152019401936118a5565b82516001600160401b0381168103611a46578152602092830192016119bf565b8f80fd5b8e80fd5b8b80fd5b634e487b7160e01b600052604160045260246000fd5b6040513d86823e3d90fd5b50346106e157611a82366146fc565b9250929050611a8f615240565b611a9836614af9565b90611aa660405192836148f6565b803683526106e15750611af39036600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b5060005b818110611b0057005b80611b35611b19611b146001948688615170565b614c8b565b611b2f6020611b2985888a615170565b01614c9f565b906154c6565b01611af7565b50346106e157611b4a36614689565b5060405163fd7c89e560e01b8152602060048201529291828480611b726024820186866149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa938415611d2d578394611d38575b5082611bd7916040518093819263629b6ce760e11b8352602060048401528660248401916149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215611d2d578380928194611ca5575b50805b828110611c6057611c4086611c5c87611c4e8860405195869560608752606087019061464c565b90858203602087015261484c565b90838203604085015261484c565b0390f35b80611c6d60019289614aa4565b5115611c7a575b01611c19565b82611c858289614aa4565b526060611c928287614aa4565b526060611c9f8288614aa4565b52611c74565b9450925090503d8083853e611cba81856148f6565b8301906060848303126112865783516001600160401b03811161084c5782611ce3918601614eae565b9160208501516001600160401b0381116107385781611d039187016150f2565b946040810151906001600160401b03821161084857611d239291016150f2565b9193919238611c16565b6040513d85823e3d90fd5b611bd7919450611d5a84913d8084833e611d5281836148f6565b81019061493b565b949150611bae565b50346106e157806003193601126106e15760206040517f82db594318110a04b6349ce48645aa69f0892751bc893d15e61d9e2b9c4630f58152f35b50346106e15760203660031901126106e157611db76145cc565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e157806003193601126106e157602090604051908152f35b50346106e157806003193601126106e1576002548060d01c9182151580611e5c575b15611e53575060a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b91505080611e37565b5042831015611e25565b50346106e15760203660031901126106e1576004356001600160401b0381116109f357611e979036906004016143f9565b5050610a4e614cb3565b50346106e15760403660031901126106e1576040611ebd6145b6565b91600435815280602052209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b50346106e157806003193601126106e1576002546040516001600160a01b039091168152602090f35b50346106e157806003193601126106e157604051634631ff3160e11b815281816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156120f157829183916120d5575b50604051634631ff3160e11b815283816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015611a6857849185916120b9575b50604051634631ff3160e11b81529185836004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa928315611805579086611c5c9594939261204398918095612089575b50509061205f61206d9261205161207b96976040519b8c9b60c08d5260c08d0190614818565b908b820360208d0152614818565b9089820360408b0152614818565b908782036060890152614818565b908582036080870152614818565b9083820360a0850152614818565b61207b955061206d9392506120ae61205f92612051923d8091833e6117f581836148f6565b96509293905061201d565b90506120cf91503d8086833e6117f581836148f6565b38611fc1565b90506120eb91503d8084833e6117f581836148f6565b38611f73565b6040513d84823e3d90fd5b50346106e15760603660031901126106e1576004356001600160401b0381116109f35761212d9036906004016143f9565b50506024356001600160401b0381116109f35761214e90369060040161449c565b50506044356001600160401b0381116109f357366023820112156109f3578060040135906001600160401b038211611286576024906101e03693020101116106e157610a4e614c33565b50346106e157806003193601126106e1576020604051600080516020615aa08339815191528152f35b50346106e157806121d13661479b565b906121dd93929361534c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561084c57610be6948461090d93819560405198899687958694630c9d68dd60e41b865260606004870152606486019161500e565b50346106e15760203660031901126106e1576004356001600160401b0381116109f35761226f9036906004016143f9565b5050610a4e614fc1565b50346106e15760203660031901126106e1576004356001600160401b0381116109f3576122aa9036906004016143f9565b604051637b5f0d5f60e01b81526020600482015291908383806122d16024820185876149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215611a6857849385936124ac575b50604051637b5f0d5f60e01b815260206004820152858082806123366024820188886149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9283156109455761239c948293839561248a575b50604051637b5f0d5f60e01b815260206004820152958692839260248401916149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa928315611805579086611c5c9594939261240c98918095612452575b5050906124286124369261241a61244496976040519b8c9b60c08d5260c08d01906147d6565b908b820360208d01526147d6565b9089820360408b01526147d6565b9087820360608901526147d6565b9085820360808701526147d6565b9083820360a08501526147d6565b612444955061243693925061247f6124289261241a923d8091833e61247781836148f6565b810190614f7a565b9650929390506123e6565b9094506124a29193503d8084833e61247781836148f6565b9290929338612378565b9092506124c49193503d8086833e61247781836148f6565b929092913861230f565b50346106e157806003193601126106e1576124e761534c565b807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b15610fa1578180916024604051809581936369a6b3db60e01b8352600160048401525af180156109455761092f5780f35b50346106e15760203660031901126106e1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615419565b50346106e15760403660031901126106e157600435906001600160401b038216809203610850576001600160401b036125bd6145f6565b6040516372d64a8160e01b8152600481019490945216602483015280826044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156106ec57809161262d575b60405160208082528190611c5c9082018561464c565b90503d8082843e61263e81846148f6565b8201916020818403126109f3578051916001600160401b0383116106e157509161266c91611c5c9301614eae565b38612617565b50346106e1576126813661479b565b909161268b6152fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561073857928491604051948592633cdd2c1960e01b84528060648501606060048701525260848401929060005b81811061271e575050508383836127068294849996506024840190614ea1565b6000604483015203925af180156109455761092f5780f35b92945092819550602080916001600160401b0361273c600195614638565b1681520195019101908794928794926126e6565b50346106e15760203660031901126106e1576004356001600160401b0381116109f3576127819036906004016143f9565b5050610a4e614e3d565b50346106e15760603660031901126106e1576004356001600160401b0381116109f3576127bc9036906004016143f9565b50506024356001600160401b0381116109f3576127dd90369060040161449c565b50506044356001600160401b0381116109f3576127819036906004016146cc565b50346106e15760203660031901126106e15760043565ffffffffffff81168082036112865761282b615280565b61283442615880565b9065ffffffffffff6128446151e8565b168082111561294c57507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b929165ffffffffffff826206978080612892951091180262069780181690615704565b906002548060d01c806128ef575b5050600280546001600160a01b031660a083901b65ffffffffffff60a01b161760d084901b6001600160d01b0319161790556040805165ffffffffffff9283168152919092166020820152a180f35b42111561292157600180546001600160d01b031660309290921b6001600160d01b0319169190911790555b38806128a0565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a161291a565b0365ffffffffffff8111612987577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92916128929190615704565b634e487b7160e01b84526011600452602484fd5b50346106e15760203660031901126106e1576129b56145cc565b6129bd615280565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed660206129fa6129ec42615880565b6129f46151e8565b90615704565b65ffffffffffff612a09615221565b6001805460d085901b166001600160a01b0390981697881785851660a01b179055919091169050612a47575b65ffffffffffff60405191168152a280f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1612a35565b50346106e157806003193601126106e15760206040517f2dca0f5ce7e75a4b43fe2b0d6f5d0b7a2bf92ecf89f8f0aa17b8308b670388218152f35b50346106e15760403660031901126106e1576004356001600160401b0381116109f357612add9036906004016143f9565b50506024356001600160401b0381116109f357610acc90369060040161476b565b50346106e15760403660031901126106e157600435906001600160401b0382116106e157816004019061010060031984360301126106e157602435926001600160e01b031984168085036108505782604051612b59816148a5565b8181528160208201528160408201526060015260848201612b7990614c8b565b90612b8660248401614c9f565b92612b9083615636565b612b9d60a4820187614cfb565b92909660c48301612bad91614cfb565b604051630360964b60e11b81526001600160a01b03871660048201526001600160401b038816602482015260c06044820152988995612bf19160c488019190614d41565b90600319868303016064870152612c0792614d41565b606490920135608484015260a48301527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691869103815a93608094fa938415611d2d578394612d4c575b5090602091612c808551966040519788948594635a985ae160e01b865260048601614d62565b0381847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19081156106ec578091612d0f575b6001600160a01b0382166060840190815260405184518152602080860151608093918791612ced91850190614748565b6040908101516001600160a01b03908116918401919091529051166060820152f35b90506020833d602011612d44575b81612d2a602093836148f6565b810103126106e15750612d3e608092614d2d565b38612cbd565b3d9150612d1d565b9093506080813d608011612dc2575b81612d68608093836148f6565b810103126112865760405190612d7d826148a5565b80518252602081015160058110156107385760608392612db59260208097960152612daa60408201614d2d565b604085015201614d2d565b6060820152939091612c5a565b3d9150612d5b565b50346106e157806003193601126106e1576020604051600080516020615a208339815191528152f35b50346106e157612e02366146fc565b50505050610a4e614cb3565b50346106e15760403660031901126106e1576004356001600160401b0381116109f357612e3f9036906004016146cc565b50506024356001600160401b0381116109f357612e609036906004016146cc565b9091612e6a615465565b612e7336614af9565b90612e8160405192836148f6565b803683526106e15750612ece9036600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b5060005b818110612edb57005b80612f07612ef76040612ef16001958789614c7b565b01614c8b565b611b2f6020611b2985888a614c7b565b01612ed2565b50346106e15760203660031901126106e1576004356001600160401b0381116109f357612f3e90369060040161446c565b5050610a4e614c33565b50346106e157806003193601126106e1576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346106e157612f9c36614689565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169392909190843b1561084c576040516301ac453d60e71b8152848180612ff187878760048501614c09565b0381895afa8015610744578590613277575b9050604051809663fd7c89e560e01b825260206004830152818061302b6024820188886149bc565b03915afa948515611a68578495613257575b5081849161305f6040519586938493631cc48d1f60e11b855260048501614c09565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215611d2d578392613163575b50829391935b81811061312c57505050604051918291602083016020845282518091526020604085019301915b8181106130d6575050500390f35b91935091602060a0600192608087518051835263ffffffff85820151168584015263ffffffff604082015116604084015261ffff60608201511660608401520151151560808201520194019101918493926130c8565b8061313b600192859795614aa4565b511561314b575b019391936130a1565b613153614ace565b61315d8286614aa4565b52613142565b9091503d8084833e61317581836148f6565b81019060208183031261084c578051906001600160401b038211610738570181601f8201121561084c578051906131ab82614917565b926131b960405194856148f6565b828452602060a081860194028301019181831161325357602001925b8284106131e75750505050903861309b565b60a08483031261325357602060a091604051613202816148db565b86518152613211838801614a2a565b8382015261322160408801614a2a565b604082015261323260608801614bfa565b60608201526132436080880161492e565b60808201528152019301926131d5565b8680fd5b8491955061326f83913d8085833e611d5281836148f6565b95915061303d565b613280916148f6565b8360001261084c573884613003565b50346106e15760403660031901126106e1576004356132ac6145b6565b9080158061335d575b6132e3575b336001600160a01b038316036132d4579061105791615842565b63334bd91960e11b8352600483fd5b6132eb615221565b906001600160a01b03161580159061334d575b801561333b575b61332157506001805465ffffffffffff60a01b191690556132ba565b6319ca5ebb60e01b845265ffffffffffff16600452602483fd5b504265ffffffffffff82161015613305565b5065ffffffffffff8116156132fe565b506002546001600160a01b038381169116146132b5565b50346106e15760203660031901126106e1576004356001600160401b0381116109f3576133a59036906004016143f9565b60405163fd7c89e560e01b81526020600482015291838084806133cc6024820187876149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9182156109455761342f948293613692575b50604051809581926234dbc560e81b8352602060048401528660248401916149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa928315611a68578493613579575b50835b82811061351c578385604051918291602083016020845282518091526020604085019301915b8181106134a2575050500390f35b91935091602060e060019260c0875163ffffffff815116835263ffffffff85820151168584015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152019401910191849392613494565b8061352960019284614aa4565b5115613536575b0161346e565b604051613542816148c0565b8681528660208201528660408201528660608201528660808201528660a08201528660c08201526135738287614aa4565b52613530565b9092503d8085833e61358b81836148f6565b810190602081830312610738578051906001600160401b038211610848570181601f82011215610738578051906135c182614917565b926135cf60405194856148f6565b828452602060e081860194028301019181831161083957602001925b8284106135fd5750505050913861346b565b60e08483031261083957602060e091604051613618816148c0565b61362187614a2a565b815261362e838801614a2a565b8382015261363e60408801614a2a565b604082015261364f60608801614a2a565b606082015261366060808801614bfa565b608082015261367160a08801614bfa565b60a082015261368260c0880161492e565b60c08201528152019301926135eb565b6136a79193503d8084833e611d5281836148f6565b913861340c565b50346106e15760803660031901126106e1576136c86145f6565b506136d161460c565b506112cf614622565b50346106e15760203660031901126106e1576004356001600160401b0381116109f35761370b9036906004016143f9565b5050610a4e614b14565b50346106e15760403660031901126106e1576004356137326145b6565b90801561105b57908161375961104d61105794600052600060205260016040600020015490565b6157c9565b50346106e157806003193601126106e1576020604051600080516020615a808339815191528152f35b50346106e15760203660031901126106e15760206137b5600435600052600060205260016040600020015490565b604051908152f35b50346106e157806003193601126106e15760206040517ffb5864e8ff833c3cb2d2d08505e82ff02a43554c74a35d4f5a64e852612783118152f35b50346106e15761380736614544565b5050336000908152600080516020615a40833981519152602052604090205460ff1615935061383d92505050576113a736614af9565b613854600080516020615a208339815191526153df565b61139e565b50346106e157806003193601126106e157611c5c604080519061387c81836148f6565b601b82527a2837b7b621b7b73334b3a6b0b730b3b2b910191718971816b232bb60291b60208301525191829160208352602083019061451f565b50346106e15760203660031901126106e1576004356001600160401b0381116109f35790806138ec61390f9336906004016143f9565b604051630592844560e21b815260206004820152948592839260248401916149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610945578192613aa0575b509060405190602082016020835281518091526040830193602060408360051b860101930194815b8383106139805785850386f35b90919293603f19868203018352865180511515825260208101511515602083015260408101511515604083015260608101511515606083015260808101511515608083015260a0810151151560a083015260c0810151151560c083015260e0810151151560e08301526101008101511515610100830152610140613a1761012083015161016061012086015261016085019061451f565b91015191610140818303910152815180825260208201906020808260051b850101940192855b828110613a5f5750505050506020806001929801930193019195939290613973565b9091929394602080600192601f198682030188528851908280613a8b845160408552604085019061451f565b93015115159101529701950193929101613a3d565b9091503d8083833e613ab281836148f6565b810190602081830312611286578051906001600160401b03821161084c570181601f82011215611286578051613ae781614917565b92613af560405194856148f6565b81845260208085019260051b840101928184116108485760208101925b848410613b25575050505050903861394b565b83516001600160401b038111610839578201610160818503601f190112610839576040519061016082018281106001600160401b03821117611a5257604052613b706020820161492e565b8252613b7e6040820161492e565b6020830152613b8f6060820161492e565b6040830152613ba06080820161492e565b6060830152613bb160a0820161492e565b6080830152613bc260c0820161492e565b60a0830152613bd360e0820161492e565b60c0830152613be5610100820161492e565b60e0830152613bf7610120820161492e565b6101008301526101408101516001600160401b03811161083557856020613c2092840101614bb5565b6101208301526101608101516001600160401b03811161083557602091010184601f82011215610831578051613c5581614917565b91613c6360405193846148f6565b81835260208084019260051b82010190878211611a4e5760208101925b828410613c9f5750505050610140820152815260209384019301613b12565b83516001600160401b038111613d20578201906040828b03601f190112613d205760405190604082018281106001600160401b03821117611a52576040526020830151916001600160401b038311611a4657613d11604085613d098f602080999881990101614bb5565b84520161492e565b83820152815201930192613c80565b8d80fd5b50346106e157806003193601126106e1576020604051600080516020615a608339815191528152f35b50346106e15760403660031901126106e1576004356001600160401b0381116109f357613d7e90369060040161449c565b50506024356001600160401b0381116109f35790613da36109f19236906004016144cc565b50507f82db594318110a04b6349ce48645aa69f0892751bc893d15e61d9e2b9c4630f581526020818152604080832033845290915281205460ff1615613ea7575b613ded36614af9565b90613dfb60405192836148f6565b368252503660006020830137600036820160200152613e43817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b50613e77817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b613ed07f74f7a545c65c11839a48d7453738b30c295408df2d944516167556759ddc6d066152c1565b613de4565b50346106e157806003193601126106e157613eee615280565b6002548060d01c80613f0e575b600280546001600160a01b031690558280f35b421115613f4057600180546001600160d01b031660309290921b6001600160d01b0319169190911790555b3880613efb565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58180a1613f39565b50346106e15760203660031901126106e1576004356001600160401b0381116109f35761370b90369060040161446c565b50346106e157806003193601126106e157604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000811660208301527f00000000000000000000000000000000000000000000000000000000000000008116928201929092527f00000000000000000000000000000000000000000000000000000000000000009091166060820152608090f35b50346106e157806003193601126106e1576020604051620697808152f35b50346106e15760203660031901126106e15760043563ffffffff60e01b81168091036109f3576020906318a4c3c360e11b81149081156140c1575b506040519015158152f35b637965db0b60e01b8114915081156140db575b50826140b6565b6301ffc9a760e01b149050826140d4565b82346106e15760203660031901126106e1576004356001600160401b0381116109f35761411d9036906004016143f9565b63fd7c89e560e01b845260206004850152908284806141406024820186866149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa938415611d2d5783946143d5575b5090826141a592604051809481926258178d60e11b8352602060048401528560248401916149bc565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215611d2d5783926142e5575b50829391935b81811061427957505050604051918291602083016020845282518091526020604085019301915b81811061421c575050500390f35b91935091602061028060019261426b6060885161423a848251614429565b61424b8682015160a0860190614429565b61425e6040820151610140860190614429565b01516101e0830190614429565b01940191019184939261420e565b80614288600192859795614aa4565b5115614298575b019391936141e7565b6040516142a4816148a5565b6142ac614ace565b81526142b6614ace565b60208201526142c3614ace565b60408201526142d0614ace565b60608201526142df8286614aa4565b5261428f565b9091503d8084833e6142f781836148f6565b81019060208183031261084c578051906001600160401b038211610738570181601f8201121561084c5780519061432d82614917565b9261433b60405194856148f6565b828452602061028081860194028301019181831161325357602001925b82841061436a575050505090846141e1565b6102808483031261325357602061028091604051614387816148a5565b6143918588614a3b565b81526143a08560a08901614a3b565b838201526143b2856101408901614a3b565b60408201526143c5856101e08901614a3b565b6060820152815201930192614358565b6141a5929194506143f084913d8084833e611d5281836148f6565b9491925061417c565b9181601f84011215610850578235916001600160401b038311610850576020808501948460061b01011161085057565b80516001600160801b03908116835260208083015163ffffffff169084015260408083015115159084015260608083015182169084015260809182015116910152565b9181601f84011215610850578235916001600160401b0383116108505760208085019460e0850201011161085057565b9181601f84011215610850578235916001600160401b038311610850576020808501948460051b01011161085057565b9181601f84011215610850578235916001600160401b0383116108505760208085019460c0850201011161085057565b60005b83811061450f5750506000910152565b81810151838201526020016144ff565b90602091614538815180928185528580860191016144fc565b601f01601f1916010190565b6060600319820112610850576004356001600160401b038111610850578161456e916004016143f9565b929092916024356001600160401b03811161085057816145909160040161449c565b92909291604435906001600160401b038211610850576145b29160040161449c565b9091565b602435906001600160a01b038216820361085057565b600435906001600160a01b038216820361085057565b35906001600160a01b038216820361085057565b602435906001600160401b038216820361085057565b604435906001600160401b038216820361085057565b606435906001600160401b038216820361085057565b35906001600160401b038216820361085057565b906020808351928381520192019060005b81811061466a5750505090565b82516001600160a01b031684526020938401939092019160010161465d565b604060031982011261085057600435906001600160401b038211610850576146b3916004016143f9565b90916024356001600160e01b0319811681036108505790565b9181601f84011215610850578235916001600160401b038311610850576020808501946060850201011161085057565b6040600319820112610850576004356001600160401b0381116108505781614726916004016143f9565b92909291602435906001600160401b038211610850576145b29160040161449c565b9060058210156147555752565b634e487b7160e01b600052602160045260246000fd5b9181601f84011215610850578235916001600160401b038311610850576020808501948460071b01011161085057565b604060031982011261085057600435906001600160401b038211610850576147c59160040161449c565b909160243560038110156108505790565b90602080835192838152019201906000905b8082106147f55750505090565b9091928351906004821015614755576020816001938293520194019201906147e8565b906020808351928381520192019060005b8181106148365750505090565b8251845260209384019390920191600101614829565b9080602083519182815201916020808360051b8301019401926000915b83831061487857505050505090565b9091929394602080614896600193601f19868203018752895161464c565b97019301930191939290614869565b608081019081106001600160401b03821117611a5257604052565b60e081019081106001600160401b03821117611a5257604052565b60a081019081106001600160401b03821117611a5257604052565b90601f801991011681019081106001600160401b03821117611a5257604052565b6001600160401b038111611a525760051b60200190565b5190811515820361085057565b602081830312610850578051906001600160401b03821161085057019080601f8301121561085057815161496e81614917565b9261497c60405194856148f6565b81845260208085019260051b82010192831161085057602001905b8282106149a45750505090565b602080916149b18461492e565b815201910190614997565b9160209082815201919060005b8181106149d65750505090565b909192604080600192838060a01b036149ee886145e2565b1681526001600160401b03614a0560208901614638565b1660208201520194019291016149c9565b51906001600160801b038216820361085057565b519063ffffffff8216820361085057565b91908260a091031261085057604051614a53816148db565b6080614a9f818395614a6481614a16565b8552614a7260208201614a2a565b6020860152614a836040820161492e565b6040860152614a9460608201614a16565b606086015201614a16565b910152565b8051821015614ab85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b60405190614adb826148db565b60006080838281528260208201528260408201528260608201520152565b6001600160401b038111611a5257601f01601f191660200190565b336000908152600080516020615a40833981519152602052604090205460ff1615614b99575b614b96614b4636614af9565b614b5360405191826148f6565b36815236600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b50565b614bb0600080516020615aa083398151915261539c565b614b3a565b81601f82011215610850578051614bcb81614af9565b92614bd960405194856148f6565b8184526020828401011161085057614bf791602080850191016144fc565b90565b519061ffff8216820361085057565b91602091614c22919594956040855260408501916149bc565b6001600160e01b0319909416910152565b336000908152600080516020615a40833981519152602052604090205460ff1615614c6457614b96614b4636614af9565b614bb0600080516020615aa08339815191526153df565b9190811015614ab8576060020190565b356001600160a01b03811681036108505790565b356001600160401b03811681036108505790565b336000908152600080516020615a40833981519152602052604090205460ff1615614ce457614b96614b4636614af9565b614bb0600080516020615ac083398151915261539c565b903590601e198136030182121561085057018035906001600160401b0382116108505760200191813603831361085057565b51906001600160a01b038216820361085057565b908060209392818452848401376000828201840152601f01601f1916010190565b6001600160a01b0390911681526001600160401b0391909116602082015260408101919091526001600160e01b0319909116606082015260800190565b336000908152600080516020615a40833981519152602052604090205460ff1615614e21575b614b96614dd136614af9565b614dde60405191826148f6565b36815236600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b614e38600080516020615b0083398151915261539c565b614dc5565b614e45615465565b614b96614e5136614af9565b614e5e60405191826148f6565b36815236600060208301376000368201602001527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316615768565b9060038210156147555752565b9080601f83011215610850578151614ec581614917565b92614ed360405194856148f6565b81845260208085019260051b82010192831161085057602001905b828210614efb5750505090565b60208091614f0884614d2d565b815201910190614eee565b9080601f8301121561085057815190614f2b82614917565b92614f3960405194856148f6565b82845260208085019360051b82010191821161085057602001915b818310614f615750505090565b8251600481101561085057815260209283019201614f54565b9190916040818403126108505780516001600160401b0381116108505783614fa3918301614f13565b9260208201516001600160401b03811161085057614bf79201614f13565b336000908152600080516020615a40833981519152602052604090205460ff1615614ff257614b96614e5136614af9565b615009600080516020615a2083398151915261539c565b614e45565b9160209082815201919060005b8181106150285750505090565b909192602080600192838060a01b03615040886145e2565b16815201940192910161501b565b9080601f8301121561085057815161506581614917565b9261507360405194856148f6565b81845260208085019260051b82010192831161085057602001905b82821061509b5750505090565b815181526020918201910161508e565b9190916040818403126108505780516001600160401b03811161085057836150d491830161504e565b9260208201516001600160401b03811161085057614bf7920161504e565b9080601f8301121561085057815161510981614917565b9261511760405194856148f6565b81845260208085019260051b820101918383116108505760208201905b83821061514357505050505090565b81516001600160401b0381116108505760209161516587848094880101614eae565b815201910190615134565b9190811015614ab85760051b81013590605e1981360301821215610850570190565b519060ff8216820361085057565b336000908152600080516020615a40833981519152602052604090205460ff16156151d157614b96614dd136614af9565b614e38600080516020615a8083398151915261539c565b6002548060d01c8015159081615217575b501561520d5760a01c65ffffffffffff1690565b5060015460d01c90565b90504211386151f9565b6001546001600160a01b0381169160a09190911c65ffffffffffff1690565b336000908152600080516020615a40833981519152602052604090205460ff161561526757565b61527e600080516020615a6083398151915261539c565b565b336000908152600080516020615ae0833981519152602052604090205460ff16156152a757565b63e2517d3f60e01b60005233600452600060245260446000fd5b60008181526020818152604080832033845290915290205460ff16156152e45750565b63e2517d3f60e01b6000523360045260245260446000fd5b336000908152600080516020615ae0833981519152602052604090205460ff161561532357565b61527e7f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a6152c1565b336000908152600080516020615ae0833981519152602052604090205460ff161561537357565b61527e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6152c1565b3360009081527f576e6b4a1c65e6c1c7a304850c027b050999e591be341911f940a2264d79cea2602052604090205460ff16156153d65750565b61527e906152c1565b3360009081527f0c04d32de920659c4e17776678bb1549241e67cf1288598a77d82b72bdd31f1e602052604090205460ff16156153d65750565b61545d9060008061542936614af9565b61543660405191826148f6565b36815260208101903683833782602036830101525190845afa615457615738565b916158b1565b602081519101f35b336000908152600080516020615a40833981519152602052604090205460ff161561548c57565b61527e600080516020615a608339815191526153df565b6001600160a01b0390911681526001600160401b03909116602082015260400190565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169190823b156108505760405163d8eb845f60e01b8152928390816155188585600484016154a3565b9181600097889403915afa8015611a6857615622575b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561084c5783604051809263d8eb845f60e01b825281806155808888600484016154a3565b03915afa8015611a685790849161560d575b50507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561084c579183916155e99360405180958194829363d8eb845f60e01b8452600484016154a3565b03915afa80156120f1576155fb575050565b6156068280926148f6565b6106e15750565b81615617916148f6565b611286578238615592565b8361562f919492946148f6565b913861552e565b60405163bbe4f6db60e01b81526001600160a01b039182166004820181905291602090829060249082907f0000000000000000000000000000000000000000000000000000000000000000165afa9081156156f8576000916156be575b506001600160a01b031633036156a65750565b63b4e0bca560e01b6000523360045260245260446000fd5b90506020813d6020116156f0575b816156d9602093836148f6565b81010312610850576156ea90614d2d565b38615693565b3d91506156cc565b6040513d6000823e3d90fd5b9065ffffffffffff8091169116019065ffffffffffff821161572257565b634e487b7160e01b600052601160045260246000fd5b3d15615763573d9061574982614af9565b9161575760405193846148f6565b82523d6000602084013e565b606090565b600080614bf79360208151910182855af1615457615738565b600254906001600160a01b0382166157b8576001600160a01b03199091166001600160a01b03821617600255614bf7906000615912565b631fe1e13d60e11b60005260046000fd5b9081156157da575b614bf791615912565b600254916001600160a01b0383166157b8576001600160a01b03199092166001600160a01b038216176002556157d1565b600254614bf791906001600160a01b0380831691161461582d575b600061599d565b600280546001600160a01b0319169055615826565b90614bf791801580615869575b1561599d57600280546001600160a01b031916905561599d565b506002546001600160a01b0383811691161461584f565b65ffffffffffff81116158985765ffffffffffff1690565b6306dfcc6560e41b600052603060045260245260446000fd5b906158d757508051156158c657805190602001fd5b63d6bda27560e01b60005260046000fd5b81511580615909575b6158e8575090565b639996b31560e01b60009081526001600160a01b0391909116600452602490fd5b50803b156158e0565b6000818152602081815260408083206001600160a01b038616845290915290205460ff16615996576000818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5050600090565b6000818152602081815260408083206001600160a01b038616845290915290205460ff1615615996576000818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fec38c2c90d317c225486715a2db5bc45232e7a65ea457e3c21efa1859a07bf678d4970c86de16ecd1df76427ebca809ee11eb7383bea0477ebfa92c227a1bfa4bf5b90f0c9764b1e15f233c1f5445e8a0bdf08c3e8ff2c00e7ef65c16a43b3f19f33fe78eb7c840e8bf68670029904b6f0da8e79346941c278a4e7676d473df15d662d5cdacd53bb3a71de494b20e05699f449a0c7faa264ed3102723f683b29c4698baa05b306e3e5e3fa66d29891e203a1418ef5bee962e2c9b109f129e8920ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb54f5df22126fddd649eec0166666c6741d9f12c91810b8d3a853eb3d621d35cbba164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

var PoolConfigManagerABI = PoolConfigManagerMetaData.ABI

var PoolConfigManagerBin = PoolConfigManagerMetaData.Bin

func DeployPoolConfigManager(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address, connectionStore common.Address, feeStore common.Address, securityStore common.Address, admin common.Address) (common.Address, *types.Transaction, *PoolConfigManager, error) {
	parsed, err := PoolConfigManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolConfigManagerBin), backend, tokenAdminRegistry, connectionStore, feeStore, securityStore, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolConfigManager{address: address, abi: *parsed, PoolConfigManagerCaller: PoolConfigManagerCaller{contract: contract}, PoolConfigManagerTransactor: PoolConfigManagerTransactor{contract: contract}, PoolConfigManagerFilterer: PoolConfigManagerFilterer{contract: contract}}, nil
}

type PoolConfigManager struct {
	address common.Address
	abi     abi.ABI
	PoolConfigManagerCaller
	PoolConfigManagerTransactor
	PoolConfigManagerFilterer
}

type PoolConfigManagerCaller struct {
	contract *bind.BoundContract
}

type PoolConfigManagerTransactor struct {
	contract *bind.BoundContract
}

type PoolConfigManagerFilterer struct {
	contract *bind.BoundContract
}

type PoolConfigManagerSession struct {
	Contract     *PoolConfigManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type PoolConfigManagerCallerSession struct {
	Contract *PoolConfigManagerCaller
	CallOpts bind.CallOpts
}

type PoolConfigManagerTransactorSession struct {
	Contract     *PoolConfigManagerTransactor
	TransactOpts bind.TransactOpts
}

type PoolConfigManagerRaw struct {
	Contract *PoolConfigManager
}

type PoolConfigManagerCallerRaw struct {
	Contract *PoolConfigManagerCaller
}

type PoolConfigManagerTransactorRaw struct {
	Contract *PoolConfigManagerTransactor
}

func NewPoolConfigManager(address common.Address, backend bind.ContractBackend) (*PoolConfigManager, error) {
	abi, err := abi.JSON(strings.NewReader(PoolConfigManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindPoolConfigManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManager{address: address, abi: abi, PoolConfigManagerCaller: PoolConfigManagerCaller{contract: contract}, PoolConfigManagerTransactor: PoolConfigManagerTransactor{contract: contract}, PoolConfigManagerFilterer: PoolConfigManagerFilterer{contract: contract}}, nil
}

func NewPoolConfigManagerCaller(address common.Address, caller bind.ContractCaller) (*PoolConfigManagerCaller, error) {
	contract, err := bindPoolConfigManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerCaller{contract: contract}, nil
}

func NewPoolConfigManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolConfigManagerTransactor, error) {
	contract, err := bindPoolConfigManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerTransactor{contract: contract}, nil
}

func NewPoolConfigManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolConfigManagerFilterer, error) {
	contract, err := bindPoolConfigManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerFilterer{contract: contract}, nil
}

func bindPoolConfigManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolConfigManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_PoolConfigManager *PoolConfigManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolConfigManager.Contract.PoolConfigManagerCaller.contract.Call(opts, result, method, params...)
}

func (_PoolConfigManager *PoolConfigManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PoolConfigManagerTransactor.contract.Transfer(opts)
}

func (_PoolConfigManager *PoolConfigManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PoolConfigManagerTransactor.contract.Transact(opts, method, params...)
}

func (_PoolConfigManager *PoolConfigManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolConfigManager.Contract.contract.Call(opts, result, method, params...)
}

func (_PoolConfigManager *PoolConfigManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.contract.Transfer(opts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.contract.Transact(opts, method, params...)
}

func (_PoolConfigManager *PoolConfigManagerCaller) CONFIGROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "CONFIG_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) CONFIGROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.CONFIGROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) CONFIGROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.CONFIGROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) CONNECTIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "CONNECTION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) CONNECTIONROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.CONNECTIONROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) CONNECTIONROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.CONNECTIONROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.DEFAULTADMINROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.DEFAULTADMINROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) DESTOVERHEADROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "DEST_OVERHEAD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) DESTOVERHEADROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.DESTOVERHEADROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) DESTOVERHEADROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.DESTOVERHEADROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) FEECOLLECTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "FEE_COLLECTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) FEECOLLECTORROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.FEECOLLECTORROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) FEECOLLECTORROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.FEECOLLECTORROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) FEEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "FEE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) FEEROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.FEEROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) FEEROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.FEEROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) PAUSERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.PAUSERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.PAUSERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) POLICYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "POLICY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) POLICYROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.POLICYROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) POLICYROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.POLICYROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) RATELIMITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "RATE_LIMIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) RATELIMITROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.RATELIMITROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) RATELIMITROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.RATELIMITROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) REMOTEPOOLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "REMOTE_POOL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) REMOTEPOOLROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.REMOTEPOOLROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) REMOTEPOOLROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.REMOTEPOOLROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) SECURITYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "SECURITY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) SECURITYROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.SECURITYROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) SECURITYROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.SECURITYROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) TOKENMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "TOKEN_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) TOKENMANAGERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.TOKENMANAGERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) TOKENMANAGERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.TOKENMANAGERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "UNPAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) UNPAUSERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.UNPAUSERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) UNPAUSERROLE() ([32]byte, error) {
	return _PoolConfigManager.Contract.UNPAUSERROLE(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) DefaultAdmin() (common.Address, error) {
	return _PoolConfigManager.Contract.DefaultAdmin(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) DefaultAdmin() (common.Address, error) {
	return _PoolConfigManager.Contract.DefaultAdmin(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) DefaultAdminDelay() (*big.Int, error) {
	return _PoolConfigManager.Contract.DefaultAdminDelay(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _PoolConfigManager.Contract.DefaultAdminDelay(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _PoolConfigManager.Contract.DefaultAdminDelayIncreaseWait(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _PoolConfigManager.Contract.DefaultAdminDelayIncreaseWait(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getAllConfiguredTokens", startIndex, maxCount)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _PoolConfigManager.Contract.GetAllConfiguredTokens(&_PoolConfigManager.CallOpts, startIndex, maxCount)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _PoolConfigManager.Contract.GetAllConfiguredTokens(&_PoolConfigManager.CallOpts, startIndex, maxCount)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetConnectionPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, common.Address, []ConnectionPolicyStoreConnectionPolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getConnectionPolicy", arg0)

	if err != nil {
		return *new(uint32), *new(common.Address), *new([]ConnectionPolicyStoreConnectionPolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]ConnectionPolicyStoreConnectionPolicyConfigArgs)).(*[]ConnectionPolicyStoreConnectionPolicyConfigArgs)

	return out0, out1, out2, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetConnectionPolicy(arg0 [32]byte) (uint32, common.Address, []ConnectionPolicyStoreConnectionPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetConnectionPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetConnectionPolicy(arg0 [32]byte) (uint32, common.Address, []ConnectionPolicyStoreConnectionPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetConnectionPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetCurrentRateLimiterStateBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getCurrentRateLimiterStateBatch", lanes)

	if err != nil {
		return *new([]IPoolConfigManagerLaneRateLimiterState), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneRateLimiterState)).(*[]IPoolConfigManagerLaneRateLimiterState)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetCurrentRateLimiterStateBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	return _PoolConfigManager.Contract.GetCurrentRateLimiterStateBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetCurrentRateLimiterStateBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error) {
	return _PoolConfigManager.Contract.GetCurrentRateLimiterStateBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetDestOverheadPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []FeePolicyStoreDestOverheadPolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getDestOverheadPolicy", arg0)

	if err != nil {
		return *new(uint32), *new([]FeePolicyStoreDestOverheadPolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]FeePolicyStoreDestOverheadPolicyConfigArgs)).(*[]FeePolicyStoreDestOverheadPolicyConfigArgs)

	return out0, out1, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetDestOverheadPolicy(arg0 [32]byte) (uint32, []FeePolicyStoreDestOverheadPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetDestOverheadPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetDestOverheadPolicy(arg0 [32]byte) (uint32, []FeePolicyStoreDestOverheadPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetDestOverheadPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetFeeBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getFeeBatch", lanes, requestedFinalityConfig)

	if err != nil {
		return *new([]IPoolConfigManagerLaneFee), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneFee)).(*[]IPoolConfigManagerLaneFee)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetFeeBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	return _PoolConfigManager.Contract.GetFeeBatch(&_PoolConfigManager.CallOpts, lanes, requestedFinalityConfig)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetFeeBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error) {
	return _PoolConfigManager.Contract.GetFeeBatch(&_PoolConfigManager.CallOpts, lanes, requestedFinalityConfig)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetFeePolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []FeePolicyStoreFeePolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getFeePolicy", arg0)

	if err != nil {
		return *new(uint32), *new([]FeePolicyStoreFeePolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]FeePolicyStoreFeePolicyConfigArgs)).(*[]FeePolicyStoreFeePolicyConfigArgs)

	return out0, out1, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetFeePolicy(arg0 [32]byte) (uint32, []FeePolicyStoreFeePolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetFeePolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetFeePolicy(arg0 [32]byte) (uint32, []FeePolicyStoreFeePolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetFeePolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineConnectionConfigs(opts *bind.CallOpts, arg0 common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineConnectionConfigs", arg0)

	if err != nil {
		return *new([]ConnectionPolicyStoreInlineConnectionConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectionPolicyStoreInlineConnectionConfigArgs)).(*[]ConnectionPolicyStoreInlineConnectionConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineConnectionConfigs(arg0 common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineConnectionConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineConnectionConfigs(arg0 common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineConnectionConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineDestOverheadConfigs(opts *bind.CallOpts, arg0 common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineDestOverheadConfigs", arg0)

	if err != nil {
		return *new([]FeePolicyStoreInlineDestOverheadConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeePolicyStoreInlineDestOverheadConfigArgs)).(*[]FeePolicyStoreInlineDestOverheadConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineDestOverheadConfigs(arg0 common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineDestOverheadConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineDestOverheadConfigs(arg0 common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineDestOverheadConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineFeeConfigs(opts *bind.CallOpts, arg0 common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineFeeConfigs", arg0)

	if err != nil {
		return *new([]FeePolicyStoreInlineFeeConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]FeePolicyStoreInlineFeeConfigArgs)).(*[]FeePolicyStoreInlineFeeConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineFeeConfigs(arg0 common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineFeeConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineFeeConfigs(arg0 common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineFeeConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineRateLimitConfigs(opts *bind.CallOpts, arg0 common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineRateLimitConfigs", arg0)

	if err != nil {
		return *new([]SecurityPolicyStoreInlineRateLimitConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]SecurityPolicyStoreInlineRateLimitConfigArgs)).(*[]SecurityPolicyStoreInlineRateLimitConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineRateLimitConfigs(arg0 common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineRateLimitConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineRateLimitConfigs(arg0 common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineRateLimitConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineRemotePoolConfigs(opts *bind.CallOpts, arg0 common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineRemotePoolConfigs", arg0)

	if err != nil {
		return *new([]ConnectionPolicyStoreInlineRemotePoolConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectionPolicyStoreInlineRemotePoolConfigArgs)).(*[]ConnectionPolicyStoreInlineRemotePoolConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineRemotePoolConfigs(arg0 common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineRemotePoolConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineRemotePoolConfigs(arg0 common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineRemotePoolConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetInlineSecurityConfigs(opts *bind.CallOpts, arg0 common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getInlineSecurityConfigs", arg0)

	if err != nil {
		return *new([]SecurityPolicyStoreInlineSecurityConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]SecurityPolicyStoreInlineSecurityConfigArgs)).(*[]SecurityPolicyStoreInlineSecurityConfigArgs)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetInlineSecurityConfigs(arg0 common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineSecurityConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetInlineSecurityConfigs(arg0 common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error) {
	return _PoolConfigManager.Contract.GetInlineSecurityConfigs(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetLaneConnectionBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getLaneConnectionBatch", lanes)

	if err != nil {
		return *new([]IPoolConfigManagerLaneConnection), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneConnection)).(*[]IPoolConfigManagerLaneConnection)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetLaneConnectionBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	return _PoolConfigManager.Contract.GetLaneConnectionBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetLaneConnectionBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	return _PoolConfigManager.Contract.GetLaneConnectionBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getLaneRecordLevelBatch", lanes)

	outstruct := new(GetLaneRecordLevelBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConnectionLevels = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.RemotePoolLevels = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)
	outstruct.RateLimitLevels = *abi.ConvertType(out[2], new([]uint8)).(*[]uint8)
	outstruct.FeeLevels = *abi.ConvertType(out[3], new([]uint8)).(*[]uint8)
	outstruct.DestOverheadLevels = *abi.ConvertType(out[4], new([]uint8)).(*[]uint8)
	outstruct.SecurityLevels = *abi.ConvertType(out[5], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _PoolConfigManager.Contract.GetLaneRecordLevelBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _PoolConfigManager.Contract.GetLaneRecordLevelBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetLaneVerificationBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, arg1 [4]byte) (GetLaneVerificationBatch,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getLaneVerificationBatch", lanes, arg1)

	outstruct := new(GetLaneVerificationBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hooks = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OutboundCCVs = *abi.ConvertType(out[1], new([][]common.Address)).(*[][]common.Address)
	outstruct.InboundCCVs = *abi.ConvertType(out[2], new([][]common.Address)).(*[][]common.Address)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetLaneVerificationBatch(lanes []IPoolConfigManagerLaneKey, arg1 [4]byte) (GetLaneVerificationBatch,

	error) {
	return _PoolConfigManager.Contract.GetLaneVerificationBatch(&_PoolConfigManager.CallOpts, lanes, arg1)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetLaneVerificationBatch(lanes []IPoolConfigManagerLaneKey, arg1 [4]byte) (GetLaneVerificationBatch,

	error) {
	return _PoolConfigManager.Contract.GetLaneVerificationBatch(&_PoolConfigManager.CallOpts, lanes, arg1)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getPolicyNames")

	outstruct := new(GetPolicyNames)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConnectionPolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.RemotePoolPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.RateLimitPolicyNames = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	outstruct.FeePolicyNames = *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)
	outstruct.DestOverheadPolicyNames = *abi.ConvertType(out[4], new([][32]byte)).(*[][32]byte)
	outstruct.SecurityPolicyNames = *abi.ConvertType(out[5], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _PoolConfigManager.Contract.GetPolicyNames(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _PoolConfigManager.Contract.GetPolicyNames(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetRateLimitPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, uint8, []SecurityPolicyStoreRateLimitPolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getRateLimitPolicy", arg0)

	if err != nil {
		return *new(uint32), *new(uint8), *new([]SecurityPolicyStoreRateLimitPolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new([]SecurityPolicyStoreRateLimitPolicyConfigArgs)).(*[]SecurityPolicyStoreRateLimitPolicyConfigArgs)

	return out0, out1, out2, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetRateLimitPolicy(arg0 [32]byte) (uint32, uint8, []SecurityPolicyStoreRateLimitPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetRateLimitPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetRateLimitPolicy(arg0 [32]byte) (uint32, uint8, []SecurityPolicyStoreRateLimitPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetRateLimitPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetRemotePoolPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []ConnectionPolicyStoreRemotePoolPolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getRemotePoolPolicy", arg0)

	if err != nil {
		return *new(uint32), *new([]ConnectionPolicyStoreRemotePoolPolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]ConnectionPolicyStoreRemotePoolPolicyConfigArgs)).(*[]ConnectionPolicyStoreRemotePoolPolicyConfigArgs)

	return out0, out1, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetRemotePoolPolicy(arg0 [32]byte) (uint32, []ConnectionPolicyStoreRemotePoolPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetRemotePoolPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetRemotePoolPolicy(arg0 [32]byte) (uint32, []ConnectionPolicyStoreRemotePoolPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetRemotePoolPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolConfigManager.Contract.GetRoleAdmin(&_PoolConfigManager.CallOpts, role)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolConfigManager.Contract.GetRoleAdmin(&_PoolConfigManager.CallOpts, role)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetSecurityPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []SecurityPolicyStoreSecurityPolicyConfigArgs, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getSecurityPolicy", arg0)

	if err != nil {
		return *new(uint32), *new([]SecurityPolicyStoreSecurityPolicyConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]SecurityPolicyStoreSecurityPolicyConfigArgs)).(*[]SecurityPolicyStoreSecurityPolicyConfigArgs)

	return out0, out1, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetSecurityPolicy(arg0 [32]byte) (uint32, []SecurityPolicyStoreSecurityPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetSecurityPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetSecurityPolicy(arg0 [32]byte) (uint32, []SecurityPolicyStoreSecurityPolicyConfigArgs, error) {
	return _PoolConfigManager.Contract.GetSecurityPolicy(&_PoolConfigManager.CallOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getStaticConfig")

	outstruct := new(GetStaticConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAdminRegistry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ConnectionStore = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeStore = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SecurityStore = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetStaticConfig() (GetStaticConfig,

	error) {
	return _PoolConfigManager.Contract.GetStaticConfig(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetStaticConfig() (GetStaticConfig,

	error) {
	return _PoolConfigManager.Contract.GetStaticConfig(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetSubsetTokens(opts *bind.CallOpts, arg0 [32]byte, arg1 uint64, arg2 uint64, arg3 uint64) (bool, []common.Address, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getSubsetTokens", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(bool), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetSubsetTokens(arg0 [32]byte, arg1 uint64, arg2 uint64, arg3 uint64) (bool, []common.Address, error) {
	return _PoolConfigManager.Contract.GetSubsetTokens(&_PoolConfigManager.CallOpts, arg0, arg1, arg2, arg3)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetSubsetTokens(arg0 [32]byte, arg1 uint64, arg2 uint64, arg3 uint64) (bool, []common.Address, error) {
	return _PoolConfigManager.Contract.GetSubsetTokens(&_PoolConfigManager.CallOpts, arg0, arg1, arg2, arg3)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetTokenAdminRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getTokenAdminRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetTokenAdminRegistry() (common.Address, error) {
	return _PoolConfigManager.Contract.GetTokenAdminRegistry(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetTokenAdminRegistry() (common.Address, error) {
	return _PoolConfigManager.Contract.GetTokenAdminRegistry(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetTokenConfigBatch(opts *bind.CallOpts, localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getTokenConfigBatch", localTokens)

	if err != nil {
		return *new([]IPoolConfigManagerTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerTokenConfig)).(*[]IPoolConfigManagerTokenConfig)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetTokenConfigBatch(localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	return _PoolConfigManager.Contract.GetTokenConfigBatch(&_PoolConfigManager.CallOpts, localTokens)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetTokenConfigBatch(localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	return _PoolConfigManager.Contract.GetTokenConfigBatch(&_PoolConfigManager.CallOpts, localTokens)
}

func (_PoolConfigManager *PoolConfigManagerCaller) GetTokenTransferFeeConfigBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "getTokenTransferFeeConfigBatch", lanes)

	if err != nil {
		return *new([]IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolV2TokenTransferFeeConfig)).(*[]IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) GetTokenTransferFeeConfigBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	return _PoolConfigManager.Contract.GetTokenTransferFeeConfigBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) GetTokenTransferFeeConfigBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error) {
	return _PoolConfigManager.Contract.GetTokenTransferFeeConfigBatch(&_PoolConfigManager.CallOpts, lanes)
}

func (_PoolConfigManager *PoolConfigManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolConfigManager.Contract.HasRole(&_PoolConfigManager.CallOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolConfigManager.Contract.HasRole(&_PoolConfigManager.CallOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) Owner() (common.Address, error) {
	return _PoolConfigManager.Contract.Owner(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) Owner() (common.Address, error) {
	return _PoolConfigManager.Contract.Owner(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _PoolConfigManager.Contract.PendingDefaultAdmin(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _PoolConfigManager.Contract.PendingDefaultAdmin(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_PoolConfigManager *PoolConfigManagerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _PoolConfigManager.Contract.PendingDefaultAdminDelay(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _PoolConfigManager.Contract.PendingDefaultAdminDelay(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolConfigManager.Contract.SupportsInterface(&_PoolConfigManager.CallOpts, interfaceId)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolConfigManager.Contract.SupportsInterface(&_PoolConfigManager.CallOpts, interfaceId)
}

func (_PoolConfigManager *PoolConfigManagerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolConfigManager.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_PoolConfigManager *PoolConfigManagerSession) TypeAndVersion() (string, error) {
	return _PoolConfigManager.Contract.TypeAndVersion(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerCallerSession) TypeAndVersion() (string, error) {
	return _PoolConfigManager.Contract.TypeAndVersion(&_PoolConfigManager.CallOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_PoolConfigManager *PoolConfigManagerSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.AcceptDefaultAdminTransfer(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.AcceptDefaultAdminTransfer(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyConnectionPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyConnectionPolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyConnectionPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyConnectionPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyConnectionPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyConnectionPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyDestOverheadPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyDestOverheadPolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyDestOverheadPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyDestOverheadPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyDestOverheadPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyDestOverheadPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyFeePolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyFeePolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyFeePolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyFeePolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyFeePolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyFeePolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineConnectionConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineConnectionConfigUpdates", arg0, inlineConfigsToSet)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineConnectionConfigUpdates(arg0 []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineConnectionConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, inlineConfigsToSet)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineConnectionConfigUpdates(arg0 []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineConnectionConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, inlineConfigsToSet)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineDestOverheadConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineDestOverheadConfigUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineDestOverheadConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineDestOverheadConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineDestOverheadConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineDestOverheadConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineFeeConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineFeeConfigUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineFeeConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineFeeConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineFeeConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineFeeConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineRateLimitBucketUpdates(opts *bind.TransactOpts, arg0 []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineRateLimitBucketUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineRateLimitBucketUpdates(arg0 []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRateLimitBucketUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineRateLimitBucketUpdates(arg0 []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRateLimitBucketUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineRateLimitConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineRateLimitConfigUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineRateLimitConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRateLimitConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineRateLimitConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRateLimitConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineRemotePoolConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineRemotePoolConfigUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineRemotePoolConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRemotePoolConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineRemotePoolConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineRemotePoolConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyInlineSecurityConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyInlineSecurityConfigUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyInlineSecurityConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineSecurityConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyInlineSecurityConfigUpdates(arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyInlineSecurityConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyRateLimitPolicyBucketUpdates(opts *bind.TransactOpts, arg0 []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyRateLimitPolicyBucketUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyRateLimitPolicyBucketUpdates(arg0 []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRateLimitPolicyBucketUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyRateLimitPolicyBucketUpdates(arg0 []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRateLimitPolicyBucketUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyRateLimitPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyRateLimitPolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyRateLimitPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRateLimitPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyRateLimitPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRateLimitPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyRemotePoolPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyRemotePoolPolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyRemotePoolPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRemotePoolPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyRemotePoolPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRemotePoolPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyRemoteTokenAddressProviderUpdates(opts *bind.TransactOpts, arg0 []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyRemoteTokenAddressProviderUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyRemoteTokenAddressProviderUpdates(arg0 []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRemoteTokenAddressProviderUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyRemoteTokenAddressProviderUpdates(arg0 []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyRemoteTokenAddressProviderUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplySecurityPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applySecurityPolicyConfigUpdates", arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplySecurityPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplySecurityPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplySecurityPolicyConfigUpdates(arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplySecurityPolicyConfigUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1, arg2)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplySubsetTokenUpdates(opts *bind.TransactOpts, arg0 []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applySubsetTokenUpdates", arg0, tokensToAdd)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplySubsetTokenUpdates(arg0 []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplySubsetTokenUpdates(&_PoolConfigManager.TransactOpts, arg0, tokensToAdd)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplySubsetTokenUpdates(arg0 []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplySubsetTokenUpdates(&_PoolConfigManager.TransactOpts, arg0, tokensToAdd)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenConnectionPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenConnectionPolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenConnectionPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenConnectionPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenConnectionPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenConnectionPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenDestOverheadPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenDestOverheadPolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenDestOverheadPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenDestOverheadPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenDestOverheadPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenDestOverheadPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenFeePolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenFeePolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenFeePolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenFeePolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenFeePolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenFeePolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenProfileUpdates(opts *bind.TransactOpts, arg0 []common.Address, arg1 []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenProfileUpdates", arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenProfileUpdates(arg0 []common.Address, arg1 []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenProfileUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenProfileUpdates(arg0 []common.Address, arg1 []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenProfileUpdates(&_PoolConfigManager.TransactOpts, arg0, arg1)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenRateLimitPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenRateLimitPolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenRateLimitPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenRateLimitPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenRateLimitPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenRateLimitPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenRemotePoolPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenRemotePoolPolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenRemotePoolPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenRemotePoolPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenRemotePoolPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenRemotePoolPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ApplyTokenSecurityPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "applyTokenSecurityPolicyUpdates", arg0)
}

func (_PoolConfigManager *PoolConfigManagerSession) ApplyTokenSecurityPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenSecurityPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ApplyTokenSecurityPolicyUpdates(arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ApplyTokenSecurityPolicyUpdates(&_PoolConfigManager.TransactOpts, arg0)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_PoolConfigManager *PoolConfigManagerSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.BeginDefaultAdminTransfer(&_PoolConfigManager.TransactOpts, newAdmin)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.BeginDefaultAdminTransfer(&_PoolConfigManager.TransactOpts, newAdmin)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_PoolConfigManager *PoolConfigManagerSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.CancelDefaultAdminTransfer(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.CancelDefaultAdminTransfer(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_PoolConfigManager *PoolConfigManagerSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ChangeDefaultAdminDelay(&_PoolConfigManager.TransactOpts, newDelay)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ChangeDefaultAdminDelay(&_PoolConfigManager.TransactOpts, newDelay)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "grantRole", role, account)
}

func (_PoolConfigManager *PoolConfigManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.GrantRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.GrantRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) PauseChains(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "pauseChains", remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerSession) PauseChains(remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseChains(&_PoolConfigManager.TransactOpts, remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) PauseChains(remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseChains(&_PoolConfigManager.TransactOpts, remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) PauseGlobally(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "pauseGlobally")
}

func (_PoolConfigManager *PoolConfigManagerSession) PauseGlobally() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseGlobally(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) PauseGlobally() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseGlobally(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) PauseTokens(opts *bind.TransactOpts, localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "pauseTokens", localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerSession) PauseTokens(localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseTokens(&_PoolConfigManager.TransactOpts, localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) PauseTokens(localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.PauseTokens(&_PoolConfigManager.TransactOpts, localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "renounceRole", role, account)
}

func (_PoolConfigManager *PoolConfigManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RenounceRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RenounceRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "revokeRole", role, account)
}

func (_PoolConfigManager *PoolConfigManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RevokeRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RevokeRole(&_PoolConfigManager.TransactOpts, role, account)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_PoolConfigManager *PoolConfigManagerSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RollbackDefaultAdminDelay(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.RollbackDefaultAdminDelay(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) UnpauseChains(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "unpauseChains", remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerSession) UnpauseChains(remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseChains(&_PoolConfigManager.TransactOpts, remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) UnpauseChains(remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseChains(&_PoolConfigManager.TransactOpts, remoteChainSelectors, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) UnpauseGlobally(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "unpauseGlobally")
}

func (_PoolConfigManager *PoolConfigManagerSession) UnpauseGlobally() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseGlobally(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) UnpauseGlobally() (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseGlobally(&_PoolConfigManager.TransactOpts)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) UnpauseTokens(opts *bind.TransactOpts, localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "unpauseTokens", localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerSession) UnpauseTokens(localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseTokens(&_PoolConfigManager.TransactOpts, localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) UnpauseTokens(localTokens []common.Address, direction uint8) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.UnpauseTokens(&_PoolConfigManager.TransactOpts, localTokens, direction)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ValidateLockOrBurn(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, tokenArgs []byte, chargeTokenFee bool) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "validateLockOrBurn", localToken, remoteChainSelector, amount, requestedFinalityConfig, tokenArgs, chargeTokenFee)
}

func (_PoolConfigManager *PoolConfigManagerSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, tokenArgs []byte, chargeTokenFee bool) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ValidateLockOrBurn(&_PoolConfigManager.TransactOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig, tokenArgs, chargeTokenFee)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ValidateLockOrBurn(localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, tokenArgs []byte, chargeTokenFee bool) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ValidateLockOrBurn(&_PoolConfigManager.TransactOpts, localToken, remoteChainSelector, amount, requestedFinalityConfig, tokenArgs, chargeTokenFee)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) ValidateReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "validateReleaseOrMint", releaseOrMintIn, requestedFinalityConfig)
}

func (_PoolConfigManager *PoolConfigManagerSession) ValidateReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ValidateReleaseOrMint(&_PoolConfigManager.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) ValidateReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.ValidateReleaseOrMint(&_PoolConfigManager.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_PoolConfigManager *PoolConfigManagerTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, pool common.Address, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.contract.Transact(opts, "withdrawFeeTokens", pool, feeTokens, recipient)
}

func (_PoolConfigManager *PoolConfigManagerSession) WithdrawFeeTokens(pool common.Address, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.WithdrawFeeTokens(&_PoolConfigManager.TransactOpts, pool, feeTokens, recipient)
}

func (_PoolConfigManager *PoolConfigManagerTransactorSession) WithdrawFeeTokens(pool common.Address, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _PoolConfigManager.Contract.WithdrawFeeTokens(&_PoolConfigManager.TransactOpts, pool, feeTokens, recipient)
}

type PoolConfigManagerDefaultAdminDelayChangeCanceledIterator struct {
	Event *PoolConfigManagerDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerDefaultAdminDelayChangeCanceled)
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
		it.Event = new(PoolConfigManagerDefaultAdminDelayChangeCanceled)
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

func (it *PoolConfigManagerDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerDefaultAdminDelayChangeCanceledIterator{contract: _PoolConfigManager.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerDefaultAdminDelayChangeCanceled)
				if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*PoolConfigManagerDefaultAdminDelayChangeCanceled, error) {
	event := new(PoolConfigManagerDefaultAdminDelayChangeCanceled)
	if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerDefaultAdminDelayChangeScheduledIterator struct {
	Event *PoolConfigManagerDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerDefaultAdminDelayChangeScheduled)
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
		it.Event = new(PoolConfigManagerDefaultAdminDelayChangeScheduled)
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

func (it *PoolConfigManagerDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerDefaultAdminDelayChangeScheduledIterator{contract: _PoolConfigManager.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerDefaultAdminDelayChangeScheduled)
				if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*PoolConfigManagerDefaultAdminDelayChangeScheduled, error) {
	event := new(PoolConfigManagerDefaultAdminDelayChangeScheduled)
	if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerDefaultAdminTransferCanceledIterator struct {
	Event *PoolConfigManagerDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerDefaultAdminTransferCanceled)
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
		it.Event = new(PoolConfigManagerDefaultAdminTransferCanceled)
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

func (it *PoolConfigManagerDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerDefaultAdminTransferCanceledIterator{contract: _PoolConfigManager.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerDefaultAdminTransferCanceled)
				if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*PoolConfigManagerDefaultAdminTransferCanceled, error) {
	event := new(PoolConfigManagerDefaultAdminTransferCanceled)
	if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerDefaultAdminTransferScheduledIterator struct {
	Event *PoolConfigManagerDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerDefaultAdminTransferScheduled)
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
		it.Event = new(PoolConfigManagerDefaultAdminTransferScheduled)
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

func (it *PoolConfigManagerDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*PoolConfigManagerDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerDefaultAdminTransferScheduledIterator{contract: _PoolConfigManager.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerDefaultAdminTransferScheduled)
				if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*PoolConfigManagerDefaultAdminTransferScheduled, error) {
	event := new(PoolConfigManagerDefaultAdminTransferScheduled)
	if err := _PoolConfigManager.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerRoleAdminChangedIterator struct {
	Event *PoolConfigManagerRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerRoleAdminChanged)
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
		it.Event = new(PoolConfigManagerRoleAdminChanged)
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

func (it *PoolConfigManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PoolConfigManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerRoleAdminChangedIterator{contract: _PoolConfigManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerRoleAdminChanged)
				if err := _PoolConfigManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseRoleAdminChanged(log types.Log) (*PoolConfigManagerRoleAdminChanged, error) {
	event := new(PoolConfigManagerRoleAdminChanged)
	if err := _PoolConfigManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerRoleGrantedIterator struct {
	Event *PoolConfigManagerRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerRoleGranted)
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
		it.Event = new(PoolConfigManagerRoleGranted)
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

func (it *PoolConfigManagerRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolConfigManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerRoleGrantedIterator{contract: _PoolConfigManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerRoleGranted)
				if err := _PoolConfigManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseRoleGranted(log types.Log) (*PoolConfigManagerRoleGranted, error) {
	event := new(PoolConfigManagerRoleGranted)
	if err := _PoolConfigManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerRoleRevokedIterator struct {
	Event *PoolConfigManagerRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerRoleRevoked)
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
		it.Event = new(PoolConfigManagerRoleRevoked)
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

func (it *PoolConfigManagerRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolConfigManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerRoleRevokedIterator{contract: _PoolConfigManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerRoleRevoked)
				if err := _PoolConfigManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseRoleRevoked(log types.Log) (*PoolConfigManagerRoleRevoked, error) {
	event := new(PoolConfigManagerRoleRevoked)
	if err := _PoolConfigManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PoolConfigManagerStaticConfigSetIterator struct {
	Event *PoolConfigManagerStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PoolConfigManagerStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolConfigManagerStaticConfigSet)
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
		it.Event = new(PoolConfigManagerStaticConfigSet)
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

func (it *PoolConfigManagerStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *PoolConfigManagerStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PoolConfigManagerStaticConfigSet struct {
	TokenAdminRegistry common.Address
	ConnectionStore    common.Address
	FeeStore           common.Address
	SecurityStore      common.Address
	Raw                types.Log
}

func (_PoolConfigManager *PoolConfigManagerFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*PoolConfigManagerStaticConfigSetIterator, error) {

	logs, sub, err := _PoolConfigManager.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &PoolConfigManagerStaticConfigSetIterator{contract: _PoolConfigManager.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_PoolConfigManager *PoolConfigManagerFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _PoolConfigManager.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PoolConfigManagerStaticConfigSet)
				if err := _PoolConfigManager.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_PoolConfigManager *PoolConfigManagerFilterer) ParseStaticConfigSet(log types.Log) (*PoolConfigManagerStaticConfigSet, error) {
	event := new(PoolConfigManagerStaticConfigSet)
	if err := _PoolConfigManager.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetLaneRecordLevelBatch struct {
	ConnectionLevels   []uint8
	RemotePoolLevels   []uint8
	RateLimitLevels    []uint8
	FeeLevels          []uint8
	DestOverheadLevels []uint8
	SecurityLevels     []uint8
}
type GetLaneVerificationBatch struct {
	Hooks        []common.Address
	OutboundCCVs [][]common.Address
	InboundCCVs  [][]common.Address
}
type GetPolicyNames struct {
	ConnectionPolicyNames   [][32]byte
	RemotePoolPolicyNames   [][32]byte
	RateLimitPolicyNames    [][32]byte
	FeePolicyNames          [][32]byte
	DestOverheadPolicyNames [][32]byte
	SecurityPolicyNames     [][32]byte
}
type GetStaticConfig struct {
	TokenAdminRegistry common.Address
	ConnectionStore    common.Address
	FeeStore           common.Address
	SecurityStore      common.Address
}
type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (PoolConfigManagerDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (PoolConfigManagerDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (PoolConfigManagerDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (PoolConfigManagerDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (PoolConfigManagerRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (PoolConfigManagerRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (PoolConfigManagerRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (PoolConfigManagerStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xecdad3df04e9354405258a50ba2fb49a154c3fd0cef28c42959c045eafcf629b")
}

func (_PoolConfigManager *PoolConfigManager) Address() common.Address {
	return _PoolConfigManager.address
}

type PoolConfigManagerInterface interface {
	CONFIGROLE(opts *bind.CallOpts) ([32]byte, error)

	CONNECTIONROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DESTOVERHEADROLE(opts *bind.CallOpts) ([32]byte, error)

	FEECOLLECTORROLE(opts *bind.CallOpts) ([32]byte, error)

	FEEROLE(opts *bind.CallOpts) ([32]byte, error)

	PAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

	POLICYROLE(opts *bind.CallOpts) ([32]byte, error)

	RATELIMITROLE(opts *bind.CallOpts) ([32]byte, error)

	REMOTEPOOLROLE(opts *bind.CallOpts) ([32]byte, error)

	SECURITYROLE(opts *bind.CallOpts) ([32]byte, error)

	TOKENMANAGERROLE(opts *bind.CallOpts) ([32]byte, error)

	UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error)

	GetConnectionPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, common.Address, []ConnectionPolicyStoreConnectionPolicyConfigArgs, error)

	GetCurrentRateLimiterStateBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneRateLimiterState, error)

	GetDestOverheadPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []FeePolicyStoreDestOverheadPolicyConfigArgs, error)

	GetFeeBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) ([]IPoolConfigManagerLaneFee, error)

	GetFeePolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []FeePolicyStoreFeePolicyConfigArgs, error)

	GetInlineConnectionConfigs(opts *bind.CallOpts, arg0 common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error)

	GetInlineDestOverheadConfigs(opts *bind.CallOpts, arg0 common.Address) ([]FeePolicyStoreInlineDestOverheadConfigArgs, error)

	GetInlineFeeConfigs(opts *bind.CallOpts, arg0 common.Address) ([]FeePolicyStoreInlineFeeConfigArgs, error)

	GetInlineRateLimitConfigs(opts *bind.CallOpts, arg0 common.Address) ([]SecurityPolicyStoreInlineRateLimitConfigArgs, error)

	GetInlineRemotePoolConfigs(opts *bind.CallOpts, arg0 common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error)

	GetInlineSecurityConfigs(opts *bind.CallOpts, arg0 common.Address) ([]SecurityPolicyStoreInlineSecurityConfigArgs, error)

	GetLaneConnectionBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error)

	GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

		error)

	GetLaneVerificationBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, arg1 [4]byte) (GetLaneVerificationBatch,

		error)

	GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

		error)

	GetRateLimitPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, uint8, []SecurityPolicyStoreRateLimitPolicyConfigArgs, error)

	GetRemotePoolPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []ConnectionPolicyStoreRemotePoolPolicyConfigArgs, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetSecurityPolicy(opts *bind.CallOpts, arg0 [32]byte) (uint32, []SecurityPolicyStoreSecurityPolicyConfigArgs, error)

	GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

		error)

	GetSubsetTokens(opts *bind.CallOpts, arg0 [32]byte, arg1 uint64, arg2 uint64, arg3 uint64) (bool, []common.Address, error)

	GetTokenAdminRegistry(opts *bind.CallOpts) (common.Address, error)

	GetTokenConfigBatch(opts *bind.CallOpts, localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error)

	GetTokenTransferFeeConfigBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolV2TokenTransferFeeConfig, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyConnectionPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error)

	ApplyDestOverheadPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreDestOverheadPolicyConfigArgs) (*types.Transaction, error)

	ApplyFeePolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []FeePolicyStoreFeePolicyConfigArgs) (*types.Transaction, error)

	ApplyInlineConnectionConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error)

	ApplyInlineDestOverheadConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineDestOverheadConfigArgs) (*types.Transaction, error)

	ApplyInlineFeeConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []FeePolicyStoreInlineFeeConfigArgs) (*types.Transaction, error)

	ApplyInlineRateLimitBucketUpdates(opts *bind.TransactOpts, arg0 []SecurityPolicyStoreInlineRateLimitBucketArgs) (*types.Transaction, error)

	ApplyInlineRateLimitConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineRateLimitConfigArgs) (*types.Transaction, error)

	ApplyInlineRemotePoolConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error)

	ApplyInlineSecurityConfigUpdates(opts *bind.TransactOpts, arg0 []IPoolConfigManagerLaneKey, arg1 []SecurityPolicyStoreInlineSecurityConfigArgs) (*types.Transaction, error)

	ApplyRateLimitPolicyBucketUpdates(opts *bind.TransactOpts, arg0 []SecurityPolicyStoreRateLimitPolicyBucketArgs) (*types.Transaction, error)

	ApplyRateLimitPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreRateLimitPolicyConfigArgs) (*types.Transaction, error)

	ApplyRemotePoolPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error)

	ApplyRemoteTokenAddressProviderUpdates(opts *bind.TransactOpts, arg0 []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error)

	ApplySecurityPolicyConfigUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigPolicyRecordKey, arg1 [][32]byte, arg2 []SecurityPolicyStoreSecurityPolicyConfigArgs) (*types.Transaction, error)

	ApplySubsetTokenUpdates(opts *bind.TransactOpts, arg0 []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error)

	ApplyTokenConnectionPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenDestOverheadPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenFeePolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenProfileUpdates(opts *bind.TransactOpts, arg0 []common.Address, arg1 []PolicyConfigTokenProfileArgs) (*types.Transaction, error)

	ApplyTokenRateLimitPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenRemotePoolPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenSecurityPolicyUpdates(opts *bind.TransactOpts, arg0 []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	PauseChains(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error)

	PauseGlobally(opts *bind.TransactOpts) (*types.Transaction, error)

	PauseTokens(opts *bind.TransactOpts, localTokens []common.Address, direction uint8) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	UnpauseChains(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8) (*types.Transaction, error)

	UnpauseGlobally(opts *bind.TransactOpts) (*types.Transaction, error)

	UnpauseTokens(opts *bind.TransactOpts, localTokens []common.Address, direction uint8) (*types.Transaction, error)

	ValidateLockOrBurn(opts *bind.TransactOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, requestedFinalityConfig [4]byte, tokenArgs []byte, chargeTokenFee bool) (*types.Transaction, error)

	ValidateReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, pool common.Address, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*PoolConfigManagerDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*PoolConfigManagerDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*PoolConfigManagerDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*PoolConfigManagerDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*PoolConfigManagerDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*PoolConfigManagerDefaultAdminTransferScheduled, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PoolConfigManagerRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*PoolConfigManagerRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolConfigManagerRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*PoolConfigManagerRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolConfigManagerRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*PoolConfigManagerRoleRevoked, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*PoolConfigManagerStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *PoolConfigManagerStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*PoolConfigManagerStaticConfigSet, error)

	Address() common.Address
}
