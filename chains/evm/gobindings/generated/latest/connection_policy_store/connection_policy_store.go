// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connection_policy_store

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

type ConnectionPolicyStoreTokenProfile struct {
	IsRegistered               bool
	LocalDecimals              uint8
	LockOrBurnFlavor           uint8
	LockOrBurnAddress          common.Address
	AllowedFinalityConfig      [4]byte
	IsOutboundPaused           bool
	IsInboundPaused            bool
	AdditionalPoolDataProvider common.Address
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

type IPoolConfigManagerLaneKey struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
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

var ConnectionPolicyStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"authorizedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyConnectionPolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.ConnectionPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineConnectionConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineConnectionConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInlineRemotePoolConfigUpdates\",\"inputs\":[{\"name\":\"inlineConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"inlineConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineRemotePoolConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"trustPolicyRemotePools\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePoolConfig\",\"type\":\"tuple\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolConfig\",\"components\":[{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemotePoolPolicyConfigUpdates\",\"inputs\":[{\"name\":\"policyConfigsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.PolicyRecordKey[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"policiesToRemove\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"policyConfigsToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyRemoteTokenAddressProviderUpdates\",\"inputs\":[{\"name\":\"providersToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemoteTokenAddressProviderArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenAddressProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySubsetTokenUpdates\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.SubsetTokenArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"tokensToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.SubsetTokenArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenConnectionPolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenProfileUpdates\",\"inputs\":[{\"name\":\"tokensToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"profilesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenProfileArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenRemotePoolPolicyUpdates\",\"inputs\":[{\"name\":\"tokenPoliciesToSet\",\"type\":\"tuple[]\",\"internalType\":\"struct PolicyConfig.TokenPolicyArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkLaneRecords\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkRequestedFinalityBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllConfiguredTokens\",\"inputs\":[{\"name\":\"startIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConnectionPolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteTokenAddressProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.ConnectionPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineConnectionConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineConnectionConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInlineRemotePoolConfigs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"inlineConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.InlineRemotePoolConfigArgs[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"trustPolicyRemotePools\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePoolConfig\",\"type\":\"tuple\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolConfig\",\"components\":[{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneConnectionBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"connections\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneConnection[]\",\"components\":[{\"name\":\"hasConnection\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isGloballyPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isTokenInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isChainOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isChainInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLaneRecordLevelBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"connectionLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"},{\"name\":\"remotePoolLevels\",\"type\":\"uint8[]\",\"internalType\":\"enum PolicyConfig.RecordLevel[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPolicyNames\",\"inputs\":[],\"outputs\":[{\"name\":\"connectionPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"remotePoolPolicyNames\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePoolPolicy\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"policyConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.RemotePoolPolicyConfigArgs[]\",\"components\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubsetTokens\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"subsetTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenConfigBatch\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.TokenConfig[]\",\"components\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"supportedChains\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"connectionPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remotePoolPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rateLimitPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"feePolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destOverheadPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"securityPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPolicies\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remotePoolPolicyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenProfiles\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"profiles\",\"type\":\"tuple[]\",\"internalType\":\"struct ConnectionPolicyStore.TokenProfile[]\",\"components\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"isOutboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isInboundPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasLaneBatch\",\"inputs\":[{\"name\":\"lanes\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.LaneKey[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"hasLanes\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveInboundLane\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"transfer\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.ResolvedInboundTransfer\",\"components\":[{\"name\":\"localAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveOutboundLane\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lane\",\"type\":\"tuple\",\"internalType\":\"struct IPoolConfigManager.OutboundLane\",\"components\":[{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChainPause\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalPause\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenPause\",\"inputs\":[{\"name\":\"localTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainPauseSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"isOutboundPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isInboundPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConnectionPolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConnectionPolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"onlySubsetTokens\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConnectionPolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalPauseSet\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineConnectionConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineConnectionConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineRemotePoolConfigRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InlineRemotePoolConfigSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"trustPolicyRemotePools\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"remotePoolConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct ConnectionPolicyStore.RemotePoolConfig\",\"components\":[{\"name\":\"useLegacyPoolDataOnSend\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remotePools\",\"type\":\"tuple[]\",\"internalType\":\"struct IPoolConfigManager.RemotePool[]\",\"components\":[{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowLegacyPoolDataOnReceive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolPolicyConfigRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolPolicyConfigSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolPolicyRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteTokenAddressProviderSet\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteTokenAddressProvider\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubsetTokenAdded\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SubsetTokenRemoved\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenConnectionPolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenPauseSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isOutboundPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isInboundPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenProfileSet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"allowedFinalityConfig\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"additionalPoolDataProvider\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemotePoolPolicySet\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newPolicyName\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRemoved\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIsPaused\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}]},{\"type\":\"error\",\"name\":\"ConnectionRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GloballyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InlineRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidAdditionalPoolData\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"additionalPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnAddress\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockOrBurnFlavor\",\"type\":\"uint8\",\"internalType\":\"enum IPoolConfigManager.LockOrBurnFlavor\"},{\"name\":\"lockOrBurnAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteToken\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolData\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"LegacyPoolDataNotAllowed\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"LegacyRemotePoolRecord\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"LocalTokenNotRegistered\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PolicyHasNoTokens\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyHasRecords\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyInUse\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PolicyRecordHasNoPools\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PolicyRecordNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PolicyRemoteTokenMismatch\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ProviderHasNoCode\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RemotePoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RemotePoolRecordNotFound\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SourceTokenMismatch\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expectedSourceToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"actualSourceToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SubsetTokenAlreadyAdded\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SubsetTokenNotFound\",\"inputs\":[{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenHasConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenIsPaused\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum ConnectionPolicyStore.PauseDirection\"}]},{\"type\":\"error\",\"name\":\"TokenNotInPolicy\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPolicyNameNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60806040523461020f5761639f8038038061001981610214565b92833981019060208183031261020f578051906001600160401b03821161020f570181601f8201121561020f578051916001600160401b0383116101c8578260051b9160208061006a818601610214565b80968152019382010191821161020f57602001915b8183106101ef578333156101de57600180546001600160a01b031916331790556020906100ab82610214565b60008152600036813760408051929083016001600160401b038111848210176101c8576040528252808383015260005b8151811015610142576001906001600160a01b036100f98285610239565b5116856101058261027b565b610112575b5050016100db565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1858561010a565b50505160005b81518110156101b9576001600160a01b036101638284610239565b51169081156101a8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef848361019a600195610379565b50604051908152a101610148565b6342bcdf7f60e11b60005260046000fd5b604051615fc590816103da8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b038116810361020f5781526020928301920161007f565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101c857604052565b805182101561024d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561024d5760005260206000200190600090565b600081815260036020526040902054801561037257600019810181811161035c5760025460001981019190821161035c5780820361030b575b50505060025480156102f557600019016102cf816002610263565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61034461031c61032d936002610263565b90549060031b1c9283926002610263565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102b4565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103d357600254680100000000000000008110156101c8576103ba61032d8260018594016002556002610263565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816306c12c9614613c895750806309adfbe614613a375780630f1e1f56146134f1578063164a111414612fa4578063181f5a7714612f455780632218f57814612bcb5780632451a62714612b7157806332ecd37514612b065780633cdd2c19146129ec57806345a9683f146127925780635ad6d90a146125445780636389dc0d146124e957806364b187ea1461216657806369a6b3db146121385780636e293a2314611f9d5780636f9b7fcc14611eb057806372d64a8114611e6357806379ba509714611ddd5780637b5f0d5f14611cf55780637ceb832214611c0f5780638c63fe6214611b405780638da5cb5b14611b1757806391a2749a146119a0578063a9cec75514611587578063af8df28b146113e1578063af9e607214610f39578063b0bf810614610b21578063bbc6fb5e14610a02578063c9a15df1146107ed578063c9d68dd0146106e3578063d6229e801461063a578063d8eb845f146105d3578063e3d9c36714610346578063f2fde38b146102bf5763fd7c89e5146101a057600080fd5b346102ba5760203660031901126102ba576004356001600160401b0381116102ba576101d0903690600401614055565b906101da82614302565b916101e860405193846142e1565b8083526101f481614302565b602084019290601f190136843760005b8281106102525783856040519182916020830190602084525180915260408301919060005b818110610237575050500390f35b82511515845285945060209384019390920191600101610229565b610284610268610263838686614454565b614409565b61027e6020610278858888614454565b01614464565b90614d59565b9060048210156102a457600191151561029d828861441d565b5201610204565b634e487b7160e01b600052602160045260246000fd5b600080fd5b346102ba5760203660031901126102ba576102d8613fb1565b6102e0615639565b6001600160a01b031633811461033557600080546001600160a01b031916821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a3005b636d6c4ee560e11b60005260046000fd5b346102ba5760203660031901126102ba57600435806000526013602052604060002061037181615e4c565b9081519161037e83614302565b9261038c60405194856142e1565b80845261039b601f1991614302565b0160005b8181106105b657505060005b81518110156104c9576001600160401b036103c6828461441d565b5116938560005260146020526103f36040806000206000906001600160401b038916825260205220615e4c565b9283519461040086614302565b9561040e60405197886142e1565b80875261041d601f1991614302565b0160005b8181106104b857505060005b855181101561047557806104436001928861441d565b51600052600760205261045960406000206144b2565b610463828a61441d565b5261046e818961441d565b500161042d565b509590935060019291946040519161048c836142c6565b888352602083015260408201526104a3828661441d565b526104ae818561441d565b50019291926103ab565b806060602080938b01015201610421565b8363ffffffff600285015416604051918291604083019083526040602084015281518091526060830190602060608260051b8601019301916000905b82821061051457505050500390f35b919390929450605f1986820301825284519060406060820192805183526001600160401b0360208201511660208401520151916060604083015282518091526080820190602060808260051b85010194019260005b82811061058b5750505050506020806001929601920192018594939192610505565b90919293946020806105a9600193607f1987820301895289516140a8565b9701950193929101610569565b6020906105c49594956147f1565b8282870101520193929361039f565b346102ba5760403660031901126102ba576105ec613fb1565b6105f4613fc7565b6105fe81836153ed565b60048110156102a4571561060e57005b63300a974760e01b60009081526001600160a01b039092166004526001600160401b0316602452604490fd5b346102ba5760403660031901126102ba576004356001600160401b0381116102ba5761066a903690600401614055565b602435906001600160e01b0319821682036102ba5760005b81811061068b57005b806106be61069f6102636001948689614454565b600090848060a01b03166106b2816159f6565b6106c4575b50856150cf565b01610682565b81526006602052604090205460301b6001600160e01b031916876106b7565b346102ba576106f136614186565b6106f9614cf3565b60005b83811061070557005b6001906001600160a01b0361071e61026383888a6143e3565b1660005260066020526040600020805460ff8160d01c1660ff8260d81c1692610749878986856150a8565b929094851515911515821490816107df575b506107d5577ff53d2e953b6972d456e89380b010f53d5adb5fcf2f9813755b77011cdebb6d4b9360ff60d81b84151560d81b169160ff60d01b9060d01b169061ffff60d01b191617179055848060a01b036107ba610263868b8d6143e3565b60408051951515865292151560208601521692a25b016106fc565b50505050506107cf565b90501515831515148c61075b565b346102ba5760203660031901126102ba576001600160a01b0361080e613fb1565b1680600052600a6020526108256040600020615e4c565b80519161083183614302565b9261083f60405194856142e1565b80845261084e601f1991614302565b0160005b8181106109bb57505060005b825181101561090857806001600160401b0361087c6001938661441d565b511683600052600b6020526040806000206000906001600160401b03841682526020522060ff6108ad825492614f15565b91604051926108bb846142ab565b828216151584526020840152604051936108d485614274565b878552602085015260081c161515604083015260608201526108f6828761441d565b52610901818661441d565b500161085e565b836040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061094157505050500390f35b919360019193955060206109ab8192603f198a820301865260c08360608b51888060a01b0381511685526001600160401b03838201511683860152604081015115156040860152015160806060850152805115156080850152015191604060a082015201906140cd565b9601920192018594939192610932565b6020906040516109ca81614274565b60008152600083820152600060408201526040516109e7816142ab565b60008152606084820152606082015282828801015201610852565b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57610a32903690600401614055565b90610a3b614cf3565b6000915b808310610a4857005b610a56610263848385614454565b926020610a64828486614454565b6001600160a01b0386169591013590610a7c866159f6565b15610b0c57856000526006602052600260406000200180549182808514159182610afb575b5050610ae65795826001959697610ab88285615539565b557fe07ad45e782b60daeb19005cd93d17910aa6611d9df3414d4929f596966ce183600080a4019190610a3f565b866320812d2f60e11b60005260045260246000fd5b610b0592506157de565b8289610aa1565b8563dfcc7fc760e01b60005260045260246000fd5b346102ba57610b2f36614387565b91610b3b939193614cf3565b60005b818110610e26575050506000915b818310610b5557005b90610b6461026384838561499d565b91610b75602061027886858561499d565b92610b8c6040610b8687868661499d565b01614793565b610ba4610b9a87868661499d565b60608101906149bf565b956020870192610bb48489614593565b610bbd8161464a565b9060005b8a828210610df75792505050610bef915060018060a01b0383169283600052600a602052604060002061572b565b80600052600b6020526040806000206000906001600160401b038a1682526020522095610c1b87615765565b610c39610c278a614793565b889060ff801983541691151516179055565b865461ff00191684151560081b61ff001617875560018701966003019560005b610c63878c614593565b9050811015610ccc5780610cc68a610ca860208f610c90600197610c8a8f610ca194614593565b906157c7565b93610c9b8580614934565b91615d24565b9201614793565b906000528a60205260406000209060ff801983541691151516179055565b01610c59565b5097929550955091966040519715158852604060208901526080880193610cf282614179565b151560408a015235601e19823603018112156102ba570192833560208501946001600160401b0382116102ba578160051b9283360387136102ba579180919392604060608d01525260a0808b01928b010195926000605e19823603015b838210610d98575050505050508660019596976001600160401b037f06ad59388432d80a26d3b9660cd30971ebc487fc2784334c556516f4ea1527a59316940390a30191610b4c565b909192939497609f198d8203018652883590828212156102ba576020809186600194019082610de46040610ddd610dd1848701806145e9565b838752838701916145c8565b9401614179565b15159101529a0196019493920190610d4f565b90610e1d8185600194610e179d9e9d610e1184898c6157c7565b80614934565b91615c8c565b01989798610bc1565b610e34610263828486614454565b610e446020610278848688614454565b9060018060a01b031680600052600b6020526040806000206000906001600160401b03851682526020522060ff815416610f175781600052600a602052610e9a6001600160401b03604060002094168094615b27565b15610eff5790610ead6001949392615765565b80600052600b602052610ed48460408060002060009086825260205220600081550161475f565b7f3260489f3856053ff7ff5e5e25474af85e3fd98c32a5556afb470aa44a645990600080a301610b3e565b5063507f5f5760e11b60005260045260245260446000fd5b6001600160401b038383636a7b66e960e11b6000526004521660245260446000fd5b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57610f69903690600401614025565b90610f7382614302565b91610f8160405193846142e1565b808352601f19610f9082614302565b0160005b81811061133b57505060005b8181106110fb57836040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610fe057505050500390f35b919390929450603f198682030182528451906101a081019180511515825260ff602082015116602083015261101d60408201516040840190614018565b63ffffffff60e01b606082015116606083015260018060a01b03608082015116608083015260018060a01b0360a08201511660a083015260c0810151926101a060c0840152835180915260206101c084019401906000905b8082106110da57505050600192602092610180808460e08796015160e085015261010081015161010085015261012081015161012085015261014081015161014085015261016081015161016085015201519101529601920192018594939192610fd1565b9091946020806001926001600160401b038951168152019601920190611075565b6111096102638284866143e3565b906001600160a01b0382169061111e826159f6565b1561132d5781600052600660205260406000209461113c828861441d565b519060018252865460ff8160a01c16602084015261116360ff8260a81c16604085016143d7565b603081901b6001600160e01b03191660608401526001600160a01b03908116608084015260018801541660a083015260008481526008602052604090206111a990615e4c565b936000526006602052600260406000200154600052600e6020526111d06040600020615e4c565b95600098895b8851811015611245576111fc6001600160401b036111f4838c61441d565b511689614d59565b9060048210156102a4578b8a8260026001951461121d575b505050016111d6565b61122d6112359161123c9361441d565b51929e615048565b9d8c61441d565b528b8a82611214565b5096979293949550979061125a82875161461a565b9661126488614302565b97611272604051998a6142e1565b808952611281601f1991614302565b013660208a013760005b87518110156112bc57806001600160401b036112a96001938b61441d565b51166112b5828c61441d565b520161128b565b5097989692939495909860005b8381106112fb575050505060c0830152600281015460e083015260030154610100909101526001905b01929192610fa0565b806001600160401b036113106001938561441d565b511661132661132083875161461a565b8861441d565b52016112c9565b9050600191509392936112f2565b604094939451906101a08201918083106001600160401b038411176113cb5760209260405260008152600083820152600060408201526000606082015260006080820152600060a0820152606060c0820152600060e08201526000610100820152600061012082015260006101408201526000610160820152600061018082015282828701015201939293610f94565b634e487b7160e01b600052604160045260246000fd5b346102ba5760203660031901126102ba5760043580600052600e602052604060002061140c81615e4c565b80519061141882614302565b9161142660405193846142e1565b808352611435601f1991614302565b0160005b81811061155b57505060005b81518110156114c357806001600160401b036114636001938561441d565b511686600052600f60205260ff6040806000206000906001600160401b038516825260205220541660405191611498836142c6565b8883526020830152151560408201526114b1828661441d565b526114bc818561441d565b5001611445565b848363ffffffff60028701541691600052601060205260018060a01b03604060002054169060405192839260608401918452602084015260606040840152815180915260206080840192019060005b818110611520575050500390f35b919350916020606060019260408751805183526001600160401b03858201511685840152015115156040820152019401910191849392611512565b60209060405161156a816142c6565b600081526000838201526000604082015282828701015201611439565b346102ba5761159536614387565b916115a1939193614cf3565b60005b8181106117fe57600085855b8083106115b957005b6115c7610263848385614571565b926115d86020610278838587614571565b936115f16115e7838587614571565b6040810190614934565b9190611606611601368584614966565b6156e9565b156117d25761162c8760018060a01b03841693846000526008602052604060002061572b565b8160005260096020526040806000206000906001600160401b038a168252602052206001600160401b0384116113cb576116668154614478565b601f8111611795575b506000601f851160011461170157976116eb7fc87011fae1db3e00505551d507555885d0f82b810c56058714645bfc51c148e9936001600160401b03936116d2888060019c9d9e6000916116f6575b508160011b916000199060031b1c19161790565b90555b60405193849360208552169660208401916145c8565b0390a30191906115b0565b90508401358f6116be565b8181526020812090601f198616815b81811061177d5750936001600160401b03936001999a9b936116eb93897fc87011fae1db3e00505551d507555885d0f82b810c56058714645bfc51c148e99810611763575b50508988811b0190556116d5565b83013560001960038b901b60f8161c191690558c80611755565b91926020600181928689013581550194019201611710565b6117c290826000526020600020601f870160051c810191602088106117c8575b601f0160051c019061491d565b8861166f565b90915081906117b5565b63995afb3360e01b60009081526001600160a01b0383166004526001600160401b038816602452604490fd5b61180c610263828486614454565b61181c6020610278848688614454565b60018060a01b038216908160005260086020526040600020926118496001600160401b0383168095615a09565b1561198757611858828261565e565b611915575b5050908160019392600052600860205261187b826040600020615b27565b508060005260096020526040806000206000908482526020522061189f8154614478565b90816118d3575b50507f5da1ddc7f5f2ac7cb20f1eca93c22a2f7f7ca74391cc0107a21af4f45b0f0c97600080a3016115a4565b81601f6000931187146118ea5750555b88806118a6565b8183526020832061190591601f0160051c810190880161491d565b80825281602081209155556118e3565b9061191f91614e03565b8051801515918261194f575b505061193857878061185d565b630f7cffcd60e41b60005260045260245260446000fd5b60209192500120816000526009602052611977604080600020600090868252602052206144b2565b602081519101201415888061192b565b505063507f5f5760e11b60005260045260245260446000fd5b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57604060031982360301126102ba57604051906119dd826142ab565b80600401356001600160401b0381116102ba57611a009060043691840101614319565b82526024810135906001600160401b0382116102ba576004611a259236920101614319565b60208201908152611a34615639565b5160005b8151811015611a9e576001906001600160a01b03611a56828561441d565b5116611a6181615f24565b611a6d575b5001611a38565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184611a66565b825160005b8151811015611b15576001600160a01b03611abe828461441d565b5116908115611b04577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083611af6600195615c23565b50604051908152a101611aa3565b6342bcdf7f60e11b60005260046000fd5b005b346102ba5760003660031901126102ba576001546040516001600160a01b039091168152602090f35b346102ba5760003660031901126102ba57604051600c548082528160208101600c60005260206000209260005b818110611bf6575050611b82925003826142e1565b604051906011548083526020830190601160005260206000209060005b818110611be057611bce85611bdc88611bba818903826142e1565b604051938493604085526040850190614240565b908382036020850152614240565b0390f35b8254845260209093019260019283019201611b9f565b8454835260019485019486945060209093019201611b6d565b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57611c3f903690600401614055565b90611c48614cf3565b6000915b808310611c5557005b611c63610263848385614454565b926020611c71828486614454565b6001600160a01b03909516940135611c88856159f6565b15611ce057846001939495600052600660205260036040600020019082825492611cb28285615424565b557fb3a957a117693885aebbdb23caa8922911f57385251d8c43e5d10e6015db37ce600080a4019190611c4c565b8463dfcc7fc760e01b60005260045260246000fd5b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57611d25903690600401614055565b611d2e8161464a565b90611d388161464a565b9260005b828110611d6d57611d5f84611bdc876040519384936040855260408501906141fe565b9083820360208501526141fe565b80611da3611d94611d846102636001958888614454565b61027e6020610278868a8a614454565b611d9e838861441d565b6148ee565b611dd7611dcd611db7610263848888614454565b611dc76020610278868a8a614454565b906153ed565b611d9e838961441d565b01611d3c565b346102ba5760003660031901126102ba576000546001600160a01b0381163303611e5257600180546001600160a01b031980821633908117909355909216600090815590916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b63015aa1e360e11b60005260046000fd5b346102ba5760403660031901126102ba576004356001600160401b03811681036102ba57611e9c611bdc91611e96613fc7565b90614812565b60405191829160208352602083019061413c565b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57611ee0903690600401614055565b611ee8614cf3565b60005b818110611ef457005b611eff818385614454565b35611f166020611f10848688614454565b01614409565b60406000838152600d602052205415611f8857907fafc9972619c33aef176a4a2930d18d6b9a2e1f370d79e37174979286d4c30627602083611f5a60019695614d1e565b600084815260108352604090819020805460a089901b8990039384169319168317905551908152a201611eeb565b5063f90d6e6b60e01b60005260045260246000fd5b346102ba5760203660031901126102ba576001600160a01b03611fbe613fb1565b16806000526008602052611fd56040600020615e4c565b805191611fe183614302565b92611fef60405194856142e1565b808452611ffe601f1991614302565b0160005b81811061212157505060005b825181101561208e57806001600160401b0361202c6001938661441d565b511683600052600960205261206d6040806000206000906001600160401b03851682526020522060405192612060846142c6565b86845260208401526144b2565b604082015261207c828761441d565b52612087818661441d565b500161200e565b836040518091602082016020835281518091526040830190602060408260051b8601019301916000905b8282106120c757505050500390f35b919360019193955060206121118192603f198a8203018652606060408a51878060a01b0381511684526001600160401b0386820151168685015201519181604082015201906140a8565b96019201920185949391926120b8565b60209061212c6147f1565b82828801015201612002565b346102ba5760203660031901126102ba5760043580151581036102ba57611b1590612161614cf3565b6147a0565b346102ba5760603660031901126102ba576004356001600160401b0381116102ba57612196903690600401614055565b906024356001600160401b0381116102ba576121b6903690600401614025565b6044356001600160401b0381116102ba576121d89094939436906004016141ce565b9390946121e3614cf3565b6000915b8083106123c85750505060005b8181106122f55750505060005b81811061220a57005b61221581838561474f565b3590612227602061027883868861474f565b916001600160401b036122406040610b8685888a61474f565b93169081156122e45780156122d3577fb778117d916c7db374b0ffb4704562ee1bbd1c0d29f15ca7faf8dd73e4b3cb46602060019561228084600c615c5c565b5083600052600e8252612297856040600020615c5c565b5083600052600f82526122c4816040806000206000908982528652209060ff801983541691151516179055565b6040519015158152a301612201565b633ca4b94360e21b60005260046000fd5b63c656089560e01b60005260046000fd5b6123008183856143e3565b3560406000828152600d6020522054156123b45780600052600e602052604060002063ffffffff60028201541661239f575461238b579081612345600193600c615b27565b508060005260106020526040600020838060a01b031981541690557fdf9c4e5139f4220985ce3b4a2959cadc2fd5187529572bbf2103ccf1594900d5600080a2016121f4565b63295f353b60e21b60005260045260246000fd5b50633c95cbc160e01b60005260045260246000fd5b63f90d6e6b60e01b60005260045260246000fd5b6123da83828495989499979699614454565b35946123ec6020610278898588614454565b9286600052600e60205261240f6001600160401b03604060002095168095615b27565b156124d05786600052600f602052600160408060002060009087825260205220019561243a87615e4c565b9560005b875181101561246e576001906124676001600160a01b0361245f838c61441d565b51168b615b27565b500161243e565b509550966001939491989299965080600052600f6020526124a38460408060002060009086825260205220600081550161475f565b7fe661c0fbd2356ac3d908ab14a46a8f7bbd4fe408766ba9f88232d759e623aecd600080a30191906121e7565b8387637798483760e01b60005260045260245260446000fd5b346102ba5760203660031901126102ba5760606001600160a01b0361250c613fb1565b168060005260066020526125246040600020916159f6565b906003600282015491015490604051921515835260208301526040820152f35b346102ba5760803660031901126102ba5761255d613fb1565b612565613fc7565b604435906001600160e01b0319821682036102ba576064356001600160401b0381116102ba57612599903690600401613fdd565b92906040519360e085018581106001600160401b038211176113cb576040526060855260208501956060875260408601946000865260608701916000835260808801936000855260a08901956000875260c08a01976000895260018060a01b0384169485600052600660205260406000209460ff601654166127815785549660ff8860d01c16612768576001600160401b0383169081600052601560205260ff6040600020541661274e57956126dc9f958f95858560ff9b956126ef9b6127029f9e9a61268d8c61267061269c9d6126869d6149d4565b9c909160301b6001600160e01b031916896150cf565b8686615151565b905288612726575b50506152b4565b8252838160a01c168a526126b5848260a81c16866143d7565b60018060a01b031686526040519a8b9a60208c525160e060208d01526101008c01906140a8565b90518a8203601f190160408c01526140a8565b9651166060880152516080870190614018565b516001600160a01b0390811660a0860152905160c085015290511660e08301520390f35b90919750600052600b6020526040600020906000526020528a60406000205416953880612695565b5063e629297160e01b600052600452600060245260446000fd5b63b68b60cb60e01b600052600452600060245260446000fd5b63068f1cbb60e31b60005260046000fd5b346102ba5760403660031901126102ba576004356001600160401b0381116102ba576127c29036906004016141ce565b6024356001600160401b0381116102ba576127e19036906004016141ce565b9290916127ec614cf3565b60005b81811061292457600085855b81831061280457005b61280f83838361474f565b3592612821602061027883868661474f565b6128316040611f1084878761474f565b9085600052600e6020526128546001600160401b03604060002092168092615a09565b1561290c5785600052600f602052604080600020600090838252602052209160018060a01b03169182600052600660205286600260406000200154036128f3578260016128a19201615c5c565b156128d85760019394957f99ca8bf1818fe2a9cc5c1f21ae12d88525e91dbcfb2f6f672ba80e5184a2c6cf600080a40191906127fb565b85630b83a80560e11b60005260045260245260445260646000fd5b86836364b903eb60e01b60005260045260245260446000fd5b85637798483760e01b60005260045260245260446000fd5b600161293182848661474f565b35612942602061027885878961474f565b6129886129556040611f1087898b61474f565b83600052600f6020526040806000206000906001600160401b03861682526020522090858060a01b031694859101615b27565b156129c857906001600160401b03600194939216907f664d2b667751d4835f78f39fd2c632484f01ef70b0b71418169fa73e4c2f0cbb600080a4016127ef565b6001600160401b0391638b0dd48360e01b6000526004521660245260445260646000fd5b346102ba576129fa36614186565b909192612a05614cf3565b60005b848110612a1157005b806001600160401b03612a2f612a2a60019489876143e3565b614464565b166000526015602052604060002080549060ff8083169260081c16612a56878983866150a8565b9190931515841515149081612af8575b50612af057612abf8183612aa7867fbe188ef34fe3f31a5ec7e7736b85eace0efffe0b9bc88bca91183b6cc57b6e3a969060ff801983541691151516179055565b9061ff00825491151560081b169061ff001916179055565b6001600160401b03612ad5612a2a868c8a6143e3565b60408051951515865292151560208601521692a25b01612a08565b505050612aea565b90501515811515148a612a66565b346102ba5760803660031901126102ba57612b1f613fc7565b6044356001600160401b03811681036102ba57606435906001600160401b03821682036102ba57612b529260043561467c565b90611bdc6040519283921515835260406020840152604083019061413c565b346102ba5760003660031901126102ba576040516002548082526020820190600260005260206000209060005b818110612bb557611bdc85611e9c818703826142e1565b8254845260209093019260019283019201612b9e565b346102ba5760603660031901126102ba576004356001600160401b0381116102ba57612bfb903690600401614055565b6024356001600160401b0381116102ba57612c1a903690600401614025565b6044356001600160401b0381116102ba57612c39903690600401614025565b949093612c44614cf3565b60005b818110612e995750505060005b818110612e1d57600085855b818310612c6957005b612c74838383614571565b3592612c866020610278838686614571565b612c9e612c94838686614571565b6040810190614593565b6001600160401b0383979297169283156122e4578115612e0457612cc18261464a565b8260005b8a828210612de357505050505081156122d35795612ce4826011615c5c565b50816000526013602052612cfc836040600020615c5c565b5081600052601460205260408060002060009085825260205220612d1f81615057565b60005b888110612dbf57505060405190876020830160208452526040820160408960051b840101988260005b828110612d8a575050505050807f3f7ae29ee595ae7323f9cf53122d3177d8ba5e32e0702ab492a57d428008d0639160019697980390a3019190612c60565b909192939a602080612db26001938f612dac90603f198c8403018b52886145e9565b906145c8565b9d01950193929101612d4b565b80612dd9612dd26001938c879d97615091565b9085615d24565b5001979197612d22565b8184612df582600196612dfc95615091565b9088615c8c565b018390612cc5565b505063466bfb0960e11b60005260045260245260446000fd5b612e288183856143e3565b356040600082815260126020522054156123b457806000526013602052604060002063ffffffff60028201541661239f575461238b579081612e6d6001936011615b27565b507f51d9277f7ad33808d961e53406ceba32abe5dc730a752f828bc34b65e29bfa0c600080a201612c54565b612ea4818385614454565b35612eb56020610278848688614454565b90806000526013602052612ed86001600160401b03604060002093168093615b27565b15612f2e579081600193926000526014602052612f0360408060002060009085825260205220615057565b7f16408cc867e3bbc60e589be5073a0c709e121c7b2a29f97cb9219ceef1a9c3a6600080a301612c47565b637798483760e01b60005260045260245260446000fd5b346102ba5760003660031901126102ba57611bdc6040805190612f6881836142e1565b601f82527f436f6e6e656374696f6e506f6c69637953746f726520322e312e302d646576006020830152519182916020835260208301906140a8565b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57612fd4903690600401614055565b90612fde82614302565b91612fec60405193846142e1565b808352601f19612ffb82614302565b0160005b818110613477575050600060ff60165416158015915b83811061311857856040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061305557505050500390f35b919360019193955060206131088192603f198a82030186528851908151151581528382015115158482015260408201511515604082015260608201511515606082015260808201511515608082015260a0820151151560a082015260c0820151151560c082015260e0820151151560e082015261010082015115156101008201526101406130f66101208401516101606101208501526101608401906140a8565b920151906101408184039101526140cd565b9601920192018594939192613046565b613126610263828688614454565b613136602061027884888a614454565b90613141838961441d565b519160018060a01b03821690816000526006602052604060002092866001600160401b038316948560005260156020526040600020908a606089015254908260ff808460d01c161515938460808c015260d81c161515918260a08b01525460ff8082161515918260c08d015260081c161515938460e08c01528d9261346f575b508115613467575b50151560208901528a9261345f575b508115613457575b50151560408601526131f28282614d59565b9360048510156102a457841561344957613263906001875260016000961460001461343457848652600960205260408620818752602052613235604087206144b2565b6101208801525b848652600b6020526040862081875260205260408620948652600a60205260408620615a09565b613401575b9061327291614fa6565b939060048110156133ed57613290575b505050506001905b01613015565b6101409099979998969193959894929401976132b160018a51950196615e4c565b978597865b8a5181101561330f578a6132d46132cd838361441d565b518b615a09565b156132e3575b506001016132b6565b99906001916133066132f5838e61441d565b519161330081615048565b9d61441d565b5290508a6132da565b5091939650919798949996939661332f61332a868a5161461a565b614eb7565b988b5b895181101561336657808b6133548261334d6001958f61441d565b519261441d565b5261335f818d61441d565b5001613332565b5097909894929a919996959396825b84811061339057505050505090600192915290878080613282565b8061339d6001928461441d565b51855260076020526133e6604086206133c1604051916133bc836142ab565b6144b2565b81528660208201526133d483875161461a565b906133df828c61441d565b528961441d565b5001613375565b634e487b7160e01b84526021600452602484fd5b60ff8354818116151561010088015261341985614f15565b61014088015260081c1661326857505050505060019061328a565b61343e8484614e03565b61012088015261323c565b50505050505060019061328a565b90508c6131e0565b91508d6131d8565b90508f6131c9565b9150386131c1565b604051906101608201918083106001600160401b038411176113cb5760209260405260008152600083820152600060408201526000606082015260006080820152600060a0820152600060c0820152600060e082015260006101008201526060610120820152606061014082015282828801015201612fff565b346102ba5760403660031901126102ba576004356001600160401b0381116102ba57613521903690600401614025565b90602435916001600160401b0383116102ba57366023840112156102ba578260040135916001600160401b0383116102ba5736602460c08502860101116102ba5761356a614cf3565b60005b82811061394557600085855b80831015611b155760009260c0810283016001600160a01b0361359e60248301614409565b1690811561393657818652600660205260408620906135bc836159f6565b6138ea576135cc60448201614431565b60405163313ce56760e01b815290602082600481885afa8992816138a9575b50613881575b50505b606481013591600583101561387d5760a4820161361081614409565b89906004860361382a576001600160a01b0316801580156137be575b61379e57505b60c484019361364861364386614409565b614d1e565b61365187615be4565b50604481019161366083614431565b9085549061379a5761ffff60a01b191660a091821b60ff60a01b161760a888901b60ff60a81b16178555999a98996001997f1a43e9c5a0da3a3ef84120fc939eac76d74623dfc040b98779c630eba40223dd979196909591949093919261376b92613757926137519261374b92613744916084906001600160a01b036136e586614409565b1660018f81901b03198b5416178a5501976136ff8961443f565b815463ffffffff60b01b191660309190911c63ffffffff60b01b1617815560016137288b614409565b9101906001808f1b03166001808f1b0319825416179055614431565b9850614409565b9361443f565b93614409565b9360ff604051961686526020860190614018565b88851b89900390811660408501526001600160e01b0319919091166060840152166080820152a2019190613579565b8c80fd5b8a6064916137b9888a630142a9c560e41b855260045261400a565b604452fd5b506040516375151b6360e01b815260048101889052602081602481855afa90811561381f578c916137f1575b501561362c565b613812915060203d8111613818575b61380a81836142e1565b810190614be0565b8c6137ea565b503d613800565b6040513d8e823e3d90fd5b8a915060038614158061386b575b6138425750613632565b8a60649161385d888a630142a9c560e41b855260045261400a565b6001600160a01b0316604452fd5b506001600160a01b0381161515613838565b8780fd5b60ff90811691168082146135f157631201705160e31b89526004859052602452604452606487fd5b9092506020813d82116138e2575b816138c4602093836142e1565b810103126138de575160ff811681036138de57918a6135eb565b8980fd5b3d91506138b7565b60ff825460a01c166044820160ff61390182614431565b16820361390f5750506135f4565b60ff8560649361391f8c94614431565b91631201705160e31b855260045260245216604452fd5b630a64406560e11b8652600486fd5b6001600160a01b0361395b6102638386866143e3565b16613965816159f6565b15613a2357806000526006602052604060002060028101541590811591613a15575b5080156139fe575b80156139e7575b6139d35790816139a7600193615a75565b507f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3600080a20161356d565b6320812d2f60e11b60005260045260246000fd5b5080600052600a6020526040600020541515613996565b50806000526008602052604060002054151561398f565b600391500154151587613987565b63dfcc7fc760e01b60005260045260246000fd5b346102ba5760203660031901126102ba576004356001600160401b0381116102ba57613a67903690600401614025565b90613a7182614302565b91613a7f60405193846142e1565b808352601f19613a8e82614302565b0160005b818110613c3a57505060005b818110613b5e578360405180916020820160208352815180915260206040840192019060005b818110613ad2575050500390f35b919350916020610100600192865180511515825260ff848201511684830152613b0360408201516040840190614018565b848060a01b03606082015116606083015263ffffffff60e01b608082015116608083015260a0810151151560a083015260c0810151151560c083015260e0858060a01b039101511660e0820152019401910191849392613ac4565b6001906001600160a01b03613b776102638386886143e3565b1660005260066020526040600020613ba2838060a01b03613b9c61026385888a6143e3565b166159f6565b9060ff81549185808060a01b03910154169160405193613bc18561428f565b15158452818160a01c166020850152613be2828260a81c16604086016143d7565b858060a01b038116606085015263ffffffff60e01b8160301b166080850152818160d01c16151560a085015260d81c16151560c083015260e0820152613c28828761441d565b52613c33818661441d565b5001613a9e565b602090604051613c498161428f565b60008152600083820152600060408201526000606082015260006080820152600060a0820152600060c0820152600060e082015282828801015201613a92565b346102ba5760c03660031901126102ba57613ca2613fb1565b90613cab613fc7565b916044356001600160401b0381116102ba57613ccb903690600401613fdd565b916064356001600160401b0381116102ba57613ceb903690600401613fdd565b909560a4359463ffffffff60e01b86168096036102ba57613d0b87614274565b6000875260208701976000895260408801956000875260608901976000895260018060a01b0387169283600052600660205260406000209060ff601654166127815781549860ff8a60d81c16613f97576001600160401b0388169687600052601560205260ff60406000205460081c16613f7d57613d9a613daa91613d908b856149d4565b8c869a929a614ae3565b9260ff8c60a01c169a8b91615889565b9691999092835115600014613f11575050505015613ef7575b600101546001600160a01b03168015159182613e74575b5050613e415750505091613df760809792613e2694608435614c1d565b8652613e0960ff8260a81c16836143d7565b6001600160a01b0316835260405194518552516020850190614018565b516001600160a01b0390811660408401529051166060820152f35b613e7090604051938493630dd0ca7960e11b8552600485015260248401526060604484015260648301906140a8565b0390fd5b6020919250604051928391639219fcd760e01b83528660048401528760248401526044830152608060648301528180613eb060848201886140a8565b03915afa908115613eeb57600091613ecc575b50158b80613dda565b613ee5915060203d6020116138185761380a81836142e1565b8b613ec3565b6040513d6000823e3d90fd5b50505063596c178360e01b60005260045260245260446000fd5b613f1c939450615151565b815160208301208151602083012003613f36575050613dc3565b613e708792613f6b8893604051958695631e56fe6160e31b8752600487015260248601526080604486015260848501906140a8565b838103600319016064850152906140a8565b8763e629297160e01b600052600452600160245260446000fd5b8563b68b60cb60e01b600052600452600160245260446000fd5b600435906001600160a01b03821682036102ba57565b602435906001600160401b03821682036102ba57565b9181601f840112156102ba578235916001600160401b0383116102ba57602083818601950101116102ba57565b60058110156102a457602452565b9060058210156102a45752565b9181601f840112156102ba578235916001600160401b0383116102ba576020808501948460051b0101116102ba57565b9181601f840112156102ba578235916001600160401b0383116102ba576020808501948460061b0101116102ba57565b60005b8381106140985750506000910152565b8181015183820152602001614088565b906020916140c181518092818552858086019101614085565b601f01601f1916010190565b9080602083519182815201916020808360051b8301019401926000915b8383106140f957505050505090565b9091929394602080600192601f19858203018652885190828061412584516040855260408501906140a8565b9301511515910152970193019301919392906140ea565b906020808351928381520192019060005b81811061415a5750505090565b82516001600160a01b031684526020938401939092019160010161414d565b359081151582036102ba57565b60606003198201126102ba57600435906001600160401b0382116102ba576141b091600401614025565b909160243560038110156102ba579060443580151581036102ba5790565b9181601f840112156102ba578235916001600160401b0383116102ba57602080850194606085020101116102ba57565b90602080835192838152019201906000905b80821061421d5750505090565b90919283519060048210156102a457602081600193829352019401920190614210565b906020808351928381520192019060005b81811061425e5750505090565b8251845260209384019390920191600101614251565b608081019081106001600160401b038211176113cb57604052565b61010081019081106001600160401b038211176113cb57604052565b604081019081106001600160401b038211176113cb57604052565b606081019081106001600160401b038211176113cb57604052565b90601f801991011681019081106001600160401b038211176113cb57604052565b6001600160401b0381116113cb5760051b60200190565b9080601f830112156102ba5781359061433182614302565b9261433f60405194856142e1565b82845260208085019360051b8201019182116102ba57602001915b8183106143675750505090565b82356001600160a01b03811681036102ba5781526020928301920161435a565b60406003198201126102ba576004356001600160401b0381116102ba57816143b191600401614055565b92909291602435906001600160401b0382116102ba576143d391600401614025565b9091565b60058210156102a45752565b91908110156143f35760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b03811681036102ba5790565b80518210156143f35760209160051b010190565b3560ff811681036102ba5790565b356001600160e01b0319811681036102ba5790565b91908110156143f35760061b0190565b356001600160401b03811681036102ba5790565b90600182811c921680156144a8575b602083101461449257565b634e487b7160e01b600052602260045260246000fd5b91607f1691614487565b90604051918260008254926144c684614478565b808452936001811690811561453457506001146144ed575b506144eb925003836142e1565b565b90506000929192526020600020906000915b8183106145185750509060206144eb92820101386144de565b60209193508060019154838589010152019101909184926144ff565b9050602092506144eb94915060ff191682840152151560051b820101386144de565b6001600160401b0381116113cb57601f01601f191660200190565b91908110156143f35760051b81013590605e19813603018212156102ba570190565b903590601e19813603018212156102ba57018035906001600160401b0382116102ba57602001918160051b360383136102ba57565b908060209392818452848401376000828201840152601f01601f1916010190565b9035601e19823603018112156102ba5701602081359101916001600160401b0382116102ba5781360383136102ba57565b9190820180921161462757565b634e487b7160e01b600052601160045260246000fd5b9190820391821161462757565b9061465482614302565b61466160405191826142e1565b8281528092614672601f1991614302565b0190602036910137565b939293600052600f6020526001600160401b0360406000209116600052602052604060002090600160ff83541692016001600160401b03815492169482861015614743576001600160401b0386911692806146d7838661461a565b11614730575b50506146e88261464a565b9460005b8381106146fb57505050509190565b8061471161470b6001938561461a565b85615a1e565b838060a01b0391549060031b1c16614729828a61441d565b52016146ec565b61473b92935061463d565b9084386146dd565b50505090915090606090565b91908110156143f3576060020190565b805490600081558161476f575050565b6000526020600020908101905b818110614787575050565b6000815560010161477c565b3580151581036102ba5790565b601654901515908160ff82161515146147ed5760ff191660ff8216176016556040519081527f971be26fcde2df81b2b63dfb7fa1d2294b524eb1896b10e5228d9c154a30af0990602090a1565b5050565b604051906147fe826142c6565b606060408360008152600060208201520152565b6001600160401b03600454911691818310156148e6576001600160401b03839116918061483f838561461a565b116148d5575b5050906148518261464a565b91600091600454925b828110614868575050505090565b614872818361461a565b6000858210156148c157600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0154600191906001600160a01b03166148ba828861441d565b520161485a565b634e487b7160e01b81526032600452602490fd5b6148df925061463d565b8138614845565b505050606090565b60048210156102a45752565b6001600160a01b0390911681526001600160401b03909116602082015260400190565b818110614928575050565b6000815560010161491d565b903590601e19813603018212156102ba57018035906001600160401b0382116102ba576020019181360383136102ba57565b92919261497282614556565b9161498060405193846142e1565b8294818452818301116102ba578281602093846000960137010152565b91908110156143f35760051b81013590607e19813603018212156102ba570190565b903590603e19813603018212156102ba570190565b91906149e08184614d59565b9060048210156102a4578115614a3e576149fa81856153ed565b9360048510156102a4578415614a14575050600190921490565b6001600160401b03925063300a974760e01b60005260018060a01b03166004521660245260446000fd5b8360018060a01b03169081600052600660205260026040600020015480600052600e602052614a7c6001600160401b03604060002093168093615a09565b614a9957509063033be34360e11b60005260045260245260446000fd5b638b0dd48360e01b60005260045260245260445260646000fd5b6001600160401b03614ae0959360609360018060a01b0316835216602082015281604082015201916145c8565b90565b939291949094614af4368585614966565b6020815191012090614b5f575b614b0b8686614fa6565b919060048110156102a45715159182614b4e575b5050614b44575090613e709160405194859463e8e962ab60e01b865260048601614ab3565b9350505050600090565b614b589250615a09565b3880614b1f565b60018060a01b038516600052600b6020526040806000206000906001600160401b038916825260205220614b968260018301615a09565b614bc5575460081c60ff16614b0157505090613e709160405194859463e8e962ab60e01b865260048601614ab3565b6000918252600301602052604090205460ff16945092505050565b908160209103126102ba575180151581036102ba5790565b9060ff8091169116039060ff821161462757565b60ff16604d811161462757600a0a90565b919060ff82169060ff811692828414614cec57828411614cc15790614c4191614bf8565b91604d60ff8416118015614c90575b614c73575050614c5f90614c0c565b908181029181830414901517156146275790565b90915063a9cb113d60e01b60005260045260245260445260646000fd5b50614c9a83614c0c565b8015614cab57600019048411614c50565b634e487b7160e01b600052601260045260246000fd5b614cca91614bf8565b91604d60ff841611614c73575050614ce190614c0c565b908115614cab570490565b5050505090565b33600052600360205260406000205415614d0957565b63d86ad9cf60e01b6000523360045260246000fd5b6001600160a01b038116908115159081614d4f575b50614d3b5750565b631935322960e21b60005260045260246000fd5b90503b1538614d33565b9060018060a01b0382166000526008602052614d8360406000206001600160401b03831690615a09565b614d9e57614d909161565e565b614d9957600090565b600290565b5050600190565b6020818303126102ba578051906001600160401b0382116102ba570181601f820112156102ba578051614dd781614556565b92614de560405194856142e1565b818452602082840101116102ba57614ae09160208085019101614085565b6001600160a01b038082166000908152600660209081526040808320600201548352601090915290205460609391169190823b15614eb1576040516304df5eb160e11b815292600092849283918291614e5f91600484016148fa565b03915afa60009181614e8c575b50614e76575b5090565b614e7f816156e9565b614e87575090565b905090565b614eaa9192503d806000833e614ea281836142e1565b810190614da5565b9038614e6c565b50505090565b90614ec182614302565b614ece60405191826142e1565b8281528092614edf601f1991614302565b019060005b828110614ef057505050565b602090604051614eff816142ab565b6060815260008382015282828501015201614ee4565b614f2160018201615e4c565b614f2b8151614eb7565b9160005b8251811015614eb15780614f456001928561441d565b5160005260076020526040600020614f5d828661441d565b516000526003840160205260ff60406000205416614f81604051926133bc846142ab565b825215156020820152614f94828761441d565b52614f9f818661441d565b5001614f2f565b60018060a01b03166000526006602052600360406000200154806000526013602052614fe060406000206001600160401b03841690615a09565b61501f57506001600160401b031660009081527f4f26c3876aa9f4b92579780beea1161a61f87ebf1ec6ee865b299e447ecba99c602052604081209091565b60005260146020526001600160401b036040600020911660005260205260406000209060029190565b60001981146146275760010190565b9061506182615e4c565b60005b815181101561508b578061508461507d6001938561441d565b5186615b27565b5001615064565b50509050565b908210156143f3576143d39160051b810190614934565b939290919360038510156102a457600185036150c75793614e87575090565b508093614e7f565b6001600160e01b0319811690811561514c576150ea81615e97565b61ffff60e01b601082811c9085901c161661514c5761ffff8360e01c16801591821561513b575b505061511b575050565b63df63778f60e01b60005260045263ffffffff60e01b1660245260446000fd5b60e01c61ffff161090503880615111565b505050565b9160609060048110156102a457600181036151da57505060018060a01b038216600052600960205260406000206001600160401b03821660005260205261519b60406000206144b2565b915b6151a6836156e9565b156151b057505090565b6001600160401b03925063995afb3360e01b60005260018060a01b03166004521660245260446000fd5b6002600094929491146151ee575b5061519d565b6001600160a01b03828116825260066020908152604080842060020154845260109091528220541680615222575b506151e8565b908092945060405180926304df5eb160e11b82528180615246888a600484016148fa565b03915afa91821561527d578092615262575b505091388061521c565b61527692503d8091833e614ea281836142e1565b3880615258565b604051903d90823e3d90fd5b916152a690614ae0949284526060602085015260608401906140a8565b9160408184039101526140a8565b9594939290946153c75760018401546060956001600160a01b03909116918261532a575b50505050509061531c60ff614ae0935460a01c16916040519460018060a01b031660208601526020855261530d6040866142e1565b60405194859360208501615289565b03601f1981018352826142e1565b60405163393e85a760e21b81526001600160a01b03891660048201526001600160401b0390921660248301526001600160e01b031990931660448201526080606482015294506000928592839182916153879160848401916145c8565b03915afa8015613eeb5760ff61531c91614ae0946000916153ac575b509338806152d8565b6153c191503d806000833e614ea281836142e1565b386153a3565b505050905060ff91505460a01c1660405190602082015260208152614ae06040826142e1565b9060018060a01b038216600052600a60205261541760406000206001600160401b03831690615a09565b614d9e57614e7291614fa6565b8181146147ed5780151590816154f5575b8215159182806154df575b6154ca5761548b575b506154515750565b60005260136020526002604060002001600163ffffffff8254160163ffffffff81116146275763ffffffff1663ffffffff19825416179055565b6000526013602052600260406000200160001963ffffffff8254160163ffffffff81116146275763ffffffff1663ffffffff1982541617905538615449565b8363f90d6e6b60e01b60005260045260246000fd5b5083600052601260205260406000205415615440565b806000526012602052604060002054156123b45780600052601360205263ffffffff6002604060002001541661543557635c32f70960e01b60005260045260246000fd5b8181146147ed5780151590816155f5575b8215159182806155df575b6154ca576155a0575b506155665750565b600052600e6020526002604060002001600163ffffffff8254160163ffffffff81116146275763ffffffff1663ffffffff19825416179055565b600052600e602052600260406000200160001963ffffffff8254160163ffffffff81116146275763ffffffff1663ffffffff198254161790553861555e565b5083600052600d60205260406000205415615555565b80600052600d602052604060002054156123b45780600052600e60205263ffffffff6002604060002001541661554a57635c32f70960e01b60005260045260246000fd5b6001546001600160a01b0316330361564d57565b6315ae3a6f60e11b60005260046000fd5b60018060a01b03169081600052600660205260026040600020015480600052600e60205261569a60406000206001600160401b03841690615a09565b156156e157600052600f6020526001600160401b036040600020911660005260205260406000209060ff825416159182156156d457505090565b614ae09250600101615a09565b505050600090565b602090805191829101209061571661570082614556565b9161570e60405193846142e1565b808352614556565b602082019190601f1901368337519020141590565b919060018060a01b031680600052600560205260406000205415613a2357506001600160401b031680156122e45761576291615c5c565b50565b906001820161577381615e4c565b9060005b82518110156157c057806157976157906001938661441d565b5184615b27565b506157a2818561441d565b5160005260038601602052604060002060ff19815416905501615777565b5050509050565b908210156143f357614ae09160051b8101906149bf565b9080600052600e6020526157f56040600020615e4c565b9160005b83518110156158655782600052600f6020526158506040806000206001600160401b03615826858961441d565b51166001600160401b03166000908152602091909152206001600160a01b03841690600101615a09565b61585c576001016157f9565b50505050600190565b50505050600090565b9080601f830112156102ba57816020614ae093359101614966565b9290606093849382156159eb57506000602083036158e8575060208183810103126102ba578035915b60ff83116158c457505060ff16929190565b613e706040519283926317397e2b60e21b84526020600485015260248401916145c8565b945090925060c083106159c657828101906060818303126159c25780359160208201356001600160401b0381116159be578161592591840161586e565b956040830135906001600160401b0382116159bb57509061594791830161586e565b93858386825115928315615982575b505050156158b257613e706040519283926317397e2b60e21b84526020600485015260248401916145c8565b6040516020810194509261599c92849261531c9287615289565b5190206159aa368385614966565b602081519101201415388681615956565b80fd5b8680fd5b8480fd5b82613e706040519283926317397e2b60e21b84526020600485015260248401916145c8565b949350606092915050565b6000526005602052604060002054151590565b60019160005201602052604060002054151590565b80548210156143f35760005260206000200190600090565b80548015615a5f576000190190615a4d8282615a1e565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600560205260409020548015615b205760001981018181116146275760045460001981019190821161462757818103615acf575b505050615abb6004615a36565b600052600560205260006040812055600190565b615b08615ae0615af1936004615a1e565b90549060031b1c9283926004615a1e565b819391549060031b91821b91600019901b19161790565b90556000526005602052604060002055388080615aae565b5050600090565b906001820191816000528260205260406000205490811515600014615865576000198201918083116146275781546000198101908111614627578381615b759503615b87575b505050615a36565b60005260205260006040812055600190565b615ba7615b97615af19386615a1e565b90549060031b1c92839286615a1e565b905560005284602052604060002055388080615b6d565b805490600160401b8210156113cb5781615af1916001615be094018155615a1e565b9055565b80600052600560205260406000205415600014615c1d57615c06816004615bbe565b600454906000526005602052604060002055600190565b50600090565b80600052600360205260406000205415600014615c1d57615c45816002615bbe565b600254906000526003602052604060002055600190565b615c668282615a09565b615b205780615c7783600193615bbe565b80549260005201602052604060002055600190565b93929190938115615d1357615ca2368387614966565b602081519101209460005b858110615cc5575050505090615cc29161441d565b52565b86615cd0828761441d565b5114615cde57600101615cad565b50613e70906001600160401b0393604051948594633dc09e5d60e01b86521660048501526040602485015260448401916145c8565b630a64406560e11b60005260046000fd5b909291615d43615d35368387614966565b602081519101208093615c5c565b508160005260076020526040600020906001600160401b0381116113cb57615d6b8254614478565b601f8111615e1a575b506000601f8211600114615db8578190615da893949596600092615dad575b50508160011b916000199060031b1c19161790565b905590565b013590503880615d93565b601f198216958382526020822091805b888110615e0257508360019596979810615de8575b505050811b01905590565b0135600019600384901b60f8161c19169055388080615ddd565b90926020600181928686013581550194019101615dc8565b615e4690836000526020600020601f840160051c810191602085106117c857601f0160051c019061491d565b38615d74565b906040519182815491828252602082019060005260206000209260005b818110615e7e5750506144eb925003836142e1565b8454835260019485019487945060209093019201615e69565b6001600160e01b031981169081156147ed5761ffff60e01b811615615f1b5760ff60015b169060f01c80615ee5575b50600103615ed15750565b633144be5b60e21b60005260045260246000fd5b60005b60108110615ef65750615ec6565b6001811b8216615f09575b600101615ee8565b91600181018091116146275791615f01565b60ff6000615ebb565b6000818152600360205260409020548015615b205760001981018181116146275760025460001981019190821161462757808203615f7e575b505050615f6a6002615a36565b600052600360205260006040812055600190565b615fa0615f8f615af1936002615a1e565b90549060031b1c9283926002615a1e565b90556000526003602052604060002055388080615f5d56fea164736f6c634300081a000a",
}

var ConnectionPolicyStoreABI = ConnectionPolicyStoreMetaData.ABI

var ConnectionPolicyStoreBin = ConnectionPolicyStoreMetaData.Bin

func DeployConnectionPolicyStore(auth *bind.TransactOpts, backend bind.ContractBackend, authorizedCallers []common.Address) (common.Address, *types.Transaction, *ConnectionPolicyStore, error) {
	parsed, err := ConnectionPolicyStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConnectionPolicyStoreBin), backend, authorizedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConnectionPolicyStore{address: address, abi: *parsed, ConnectionPolicyStoreCaller: ConnectionPolicyStoreCaller{contract: contract}, ConnectionPolicyStoreTransactor: ConnectionPolicyStoreTransactor{contract: contract}, ConnectionPolicyStoreFilterer: ConnectionPolicyStoreFilterer{contract: contract}}, nil
}

type ConnectionPolicyStore struct {
	address common.Address
	abi     abi.ABI
	ConnectionPolicyStoreCaller
	ConnectionPolicyStoreTransactor
	ConnectionPolicyStoreFilterer
}

type ConnectionPolicyStoreCaller struct {
	contract *bind.BoundContract
}

type ConnectionPolicyStoreTransactor struct {
	contract *bind.BoundContract
}

type ConnectionPolicyStoreFilterer struct {
	contract *bind.BoundContract
}

type ConnectionPolicyStoreSession struct {
	Contract     *ConnectionPolicyStore
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ConnectionPolicyStoreCallerSession struct {
	Contract *ConnectionPolicyStoreCaller
	CallOpts bind.CallOpts
}

type ConnectionPolicyStoreTransactorSession struct {
	Contract     *ConnectionPolicyStoreTransactor
	TransactOpts bind.TransactOpts
}

type ConnectionPolicyStoreRaw struct {
	Contract *ConnectionPolicyStore
}

type ConnectionPolicyStoreCallerRaw struct {
	Contract *ConnectionPolicyStoreCaller
}

type ConnectionPolicyStoreTransactorRaw struct {
	Contract *ConnectionPolicyStoreTransactor
}

func NewConnectionPolicyStore(address common.Address, backend bind.ContractBackend) (*ConnectionPolicyStore, error) {
	abi, err := abi.JSON(strings.NewReader(ConnectionPolicyStoreABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindConnectionPolicyStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStore{address: address, abi: abi, ConnectionPolicyStoreCaller: ConnectionPolicyStoreCaller{contract: contract}, ConnectionPolicyStoreTransactor: ConnectionPolicyStoreTransactor{contract: contract}, ConnectionPolicyStoreFilterer: ConnectionPolicyStoreFilterer{contract: contract}}, nil
}

func NewConnectionPolicyStoreCaller(address common.Address, caller bind.ContractCaller) (*ConnectionPolicyStoreCaller, error) {
	contract, err := bindConnectionPolicyStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreCaller{contract: contract}, nil
}

func NewConnectionPolicyStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectionPolicyStoreTransactor, error) {
	contract, err := bindConnectionPolicyStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTransactor{contract: contract}, nil
}

func NewConnectionPolicyStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectionPolicyStoreFilterer, error) {
	contract, err := bindConnectionPolicyStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreFilterer{contract: contract}, nil
}

func bindConnectionPolicyStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConnectionPolicyStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectionPolicyStore.Contract.ConnectionPolicyStoreCaller.contract.Call(opts, result, method, params...)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ConnectionPolicyStoreTransactor.contract.Transfer(opts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ConnectionPolicyStoreTransactor.contract.Transact(opts, method, params...)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectionPolicyStore.Contract.contract.Call(opts, result, method, params...)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.contract.Transfer(opts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.contract.Transact(opts, method, params...)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "checkLaneRecords", localToken, remoteChainSelector)

	if err != nil {
		return err
	}

	return err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _ConnectionPolicyStore.Contract.CheckLaneRecords(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) CheckLaneRecords(localToken common.Address, remoteChainSelector uint64) error {
	return _ConnectionPolicyStore.Contract.CheckLaneRecords(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) CheckRequestedFinalityBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) error {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "checkRequestedFinalityBatch", lanes, requestedFinalityConfig)

	if err != nil {
		return err
	}

	return err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) CheckRequestedFinalityBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) error {
	return _ConnectionPolicyStore.Contract.CheckRequestedFinalityBatch(&_ConnectionPolicyStore.CallOpts, lanes, requestedFinalityConfig)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) CheckRequestedFinalityBatch(lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) error {
	return _ConnectionPolicyStore.Contract.CheckRequestedFinalityBatch(&_ConnectionPolicyStore.CallOpts, lanes, requestedFinalityConfig)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _ConnectionPolicyStore.Contract.GetAllAuthorizedCallers(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _ConnectionPolicyStore.Contract.GetAllAuthorizedCallers(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getAllConfiguredTokens", startIndex, maxCount)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _ConnectionPolicyStore.Contract.GetAllConfiguredTokens(&_ConnectionPolicyStore.CallOpts, startIndex, maxCount)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetAllConfiguredTokens(startIndex uint64, maxCount uint64) ([]common.Address, error) {
	return _ConnectionPolicyStore.Contract.GetAllConfiguredTokens(&_ConnectionPolicyStore.CallOpts, startIndex, maxCount)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetConnectionPolicy(opts *bind.CallOpts, policyName [32]byte) (GetConnectionPolicy,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getConnectionPolicy", policyName)

	outstruct := new(GetConnectionPolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RemoteTokenAddressProvider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PolicyConfigs = *abi.ConvertType(out[2], new([]ConnectionPolicyStoreConnectionPolicyConfigArgs)).(*[]ConnectionPolicyStoreConnectionPolicyConfigArgs)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetConnectionPolicy(policyName [32]byte) (GetConnectionPolicy,

	error) {
	return _ConnectionPolicyStore.Contract.GetConnectionPolicy(&_ConnectionPolicyStore.CallOpts, policyName)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetConnectionPolicy(policyName [32]byte) (GetConnectionPolicy,

	error) {
	return _ConnectionPolicyStore.Contract.GetConnectionPolicy(&_ConnectionPolicyStore.CallOpts, policyName)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetInlineConnectionConfigs(opts *bind.CallOpts, localToken common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getInlineConnectionConfigs", localToken)

	if err != nil {
		return *new([]ConnectionPolicyStoreInlineConnectionConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectionPolicyStoreInlineConnectionConfigArgs)).(*[]ConnectionPolicyStoreInlineConnectionConfigArgs)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetInlineConnectionConfigs(localToken common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	return _ConnectionPolicyStore.Contract.GetInlineConnectionConfigs(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetInlineConnectionConfigs(localToken common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error) {
	return _ConnectionPolicyStore.Contract.GetInlineConnectionConfigs(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetInlineRemotePoolConfigs(opts *bind.CallOpts, localToken common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getInlineRemotePoolConfigs", localToken)

	if err != nil {
		return *new([]ConnectionPolicyStoreInlineRemotePoolConfigArgs), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectionPolicyStoreInlineRemotePoolConfigArgs)).(*[]ConnectionPolicyStoreInlineRemotePoolConfigArgs)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetInlineRemotePoolConfigs(localToken common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	return _ConnectionPolicyStore.Contract.GetInlineRemotePoolConfigs(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetInlineRemotePoolConfigs(localToken common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error) {
	return _ConnectionPolicyStore.Contract.GetInlineRemotePoolConfigs(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetLaneConnectionBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getLaneConnectionBatch", lanes)

	if err != nil {
		return *new([]IPoolConfigManagerLaneConnection), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerLaneConnection)).(*[]IPoolConfigManagerLaneConnection)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetLaneConnectionBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	return _ConnectionPolicyStore.Contract.GetLaneConnectionBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetLaneConnectionBatch(lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error) {
	return _ConnectionPolicyStore.Contract.GetLaneConnectionBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getLaneRecordLevelBatch", lanes)

	outstruct := new(GetLaneRecordLevelBatch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConnectionLevels = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.RemotePoolLevels = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _ConnectionPolicyStore.Contract.GetLaneRecordLevelBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetLaneRecordLevelBatch(lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

	error) {
	return _ConnectionPolicyStore.Contract.GetLaneRecordLevelBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getPolicyNames")

	outstruct := new(GetPolicyNames)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConnectionPolicyNames = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.RemotePoolPolicyNames = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _ConnectionPolicyStore.Contract.GetPolicyNames(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetPolicyNames() (GetPolicyNames,

	error) {
	return _ConnectionPolicyStore.Contract.GetPolicyNames(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetRemotePoolPolicy(opts *bind.CallOpts, policyName [32]byte) (GetRemotePoolPolicy,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getRemotePoolPolicy", policyName)

	outstruct := new(GetRemotePoolPolicy)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PolicyConfigs = *abi.ConvertType(out[1], new([]ConnectionPolicyStoreRemotePoolPolicyConfigArgs)).(*[]ConnectionPolicyStoreRemotePoolPolicyConfigArgs)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetRemotePoolPolicy(policyName [32]byte) (GetRemotePoolPolicy,

	error) {
	return _ConnectionPolicyStore.Contract.GetRemotePoolPolicy(&_ConnectionPolicyStore.CallOpts, policyName)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetRemotePoolPolicy(policyName [32]byte) (GetRemotePoolPolicy,

	error) {
	return _ConnectionPolicyStore.Contract.GetRemotePoolPolicy(&_ConnectionPolicyStore.CallOpts, policyName)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetSubsetTokens(opts *bind.CallOpts, policyName [32]byte, remoteChainSelector uint64, startIndex uint64, maxCount uint64) (GetSubsetTokens,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getSubsetTokens", policyName, remoteChainSelector, startIndex, maxCount)

	outstruct := new(GetSubsetTokens)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OnlySubsetTokens = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SubsetTokens = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetSubsetTokens(policyName [32]byte, remoteChainSelector uint64, startIndex uint64, maxCount uint64) (GetSubsetTokens,

	error) {
	return _ConnectionPolicyStore.Contract.GetSubsetTokens(&_ConnectionPolicyStore.CallOpts, policyName, remoteChainSelector, startIndex, maxCount)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetSubsetTokens(policyName [32]byte, remoteChainSelector uint64, startIndex uint64, maxCount uint64) (GetSubsetTokens,

	error) {
	return _ConnectionPolicyStore.Contract.GetSubsetTokens(&_ConnectionPolicyStore.CallOpts, policyName, remoteChainSelector, startIndex, maxCount)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetTokenConfigBatch(opts *bind.CallOpts, localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getTokenConfigBatch", localTokens)

	if err != nil {
		return *new([]IPoolConfigManagerTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolConfigManagerTokenConfig)).(*[]IPoolConfigManagerTokenConfig)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetTokenConfigBatch(localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	return _ConnectionPolicyStore.Contract.GetTokenConfigBatch(&_ConnectionPolicyStore.CallOpts, localTokens)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetTokenConfigBatch(localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error) {
	return _ConnectionPolicyStore.Contract.GetTokenConfigBatch(&_ConnectionPolicyStore.CallOpts, localTokens)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

	error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getTokenPolicies", localToken)

	outstruct := new(GetTokenPolicies)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConnectionPolicyName = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.RemotePoolPolicyName = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _ConnectionPolicyStore.Contract.GetTokenPolicies(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetTokenPolicies(localToken common.Address) (GetTokenPolicies,

	error) {
	return _ConnectionPolicyStore.Contract.GetTokenPolicies(&_ConnectionPolicyStore.CallOpts, localToken)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) GetTokenProfiles(opts *bind.CallOpts, localTokens []common.Address) ([]ConnectionPolicyStoreTokenProfile, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "getTokenProfiles", localTokens)

	if err != nil {
		return *new([]ConnectionPolicyStoreTokenProfile), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectionPolicyStoreTokenProfile)).(*[]ConnectionPolicyStoreTokenProfile)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) GetTokenProfiles(localTokens []common.Address) ([]ConnectionPolicyStoreTokenProfile, error) {
	return _ConnectionPolicyStore.Contract.GetTokenProfiles(&_ConnectionPolicyStore.CallOpts, localTokens)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) GetTokenProfiles(localTokens []common.Address) ([]ConnectionPolicyStoreTokenProfile, error) {
	return _ConnectionPolicyStore.Contract.GetTokenProfiles(&_ConnectionPolicyStore.CallOpts, localTokens)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) HasLaneBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]bool, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "hasLaneBatch", lanes)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) HasLaneBatch(lanes []IPoolConfigManagerLaneKey) ([]bool, error) {
	return _ConnectionPolicyStore.Contract.HasLaneBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) HasLaneBatch(lanes []IPoolConfigManagerLaneKey) ([]bool, error) {
	return _ConnectionPolicyStore.Contract.HasLaneBatch(&_ConnectionPolicyStore.CallOpts, lanes)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) Owner() (common.Address, error) {
	return _ConnectionPolicyStore.Contract.Owner(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) Owner() (common.Address, error) {
	return _ConnectionPolicyStore.Contract.Owner(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) ResolveInboundLane(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourcePoolAddress []byte, sourcePoolData []byte, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte) (IPoolConfigManagerResolvedInboundTransfer, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "resolveInboundLane", localToken, remoteChainSelector, sourcePoolAddress, sourcePoolData, sourceDenominatedAmount, requestedFinalityConfig)

	if err != nil {
		return *new(IPoolConfigManagerResolvedInboundTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolConfigManagerResolvedInboundTransfer)).(*IPoolConfigManagerResolvedInboundTransfer)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ResolveInboundLane(localToken common.Address, remoteChainSelector uint64, sourcePoolAddress []byte, sourcePoolData []byte, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte) (IPoolConfigManagerResolvedInboundTransfer, error) {
	return _ConnectionPolicyStore.Contract.ResolveInboundLane(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector, sourcePoolAddress, sourcePoolData, sourceDenominatedAmount, requestedFinalityConfig)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) ResolveInboundLane(localToken common.Address, remoteChainSelector uint64, sourcePoolAddress []byte, sourcePoolData []byte, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte) (IPoolConfigManagerResolvedInboundTransfer, error) {
	return _ConnectionPolicyStore.Contract.ResolveInboundLane(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector, sourcePoolAddress, sourcePoolData, sourceDenominatedAmount, requestedFinalityConfig)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) ResolveOutboundLane(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, requestedFinalityConfig [4]byte, tokenArgs []byte) (IPoolConfigManagerOutboundLane, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "resolveOutboundLane", localToken, remoteChainSelector, requestedFinalityConfig, tokenArgs)

	if err != nil {
		return *new(IPoolConfigManagerOutboundLane), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolConfigManagerOutboundLane)).(*IPoolConfigManagerOutboundLane)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ResolveOutboundLane(localToken common.Address, remoteChainSelector uint64, requestedFinalityConfig [4]byte, tokenArgs []byte) (IPoolConfigManagerOutboundLane, error) {
	return _ConnectionPolicyStore.Contract.ResolveOutboundLane(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector, requestedFinalityConfig, tokenArgs)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) ResolveOutboundLane(localToken common.Address, remoteChainSelector uint64, requestedFinalityConfig [4]byte, tokenArgs []byte) (IPoolConfigManagerOutboundLane, error) {
	return _ConnectionPolicyStore.Contract.ResolveOutboundLane(&_ConnectionPolicyStore.CallOpts, localToken, remoteChainSelector, requestedFinalityConfig, tokenArgs)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConnectionPolicyStore.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) TypeAndVersion() (string, error) {
	return _ConnectionPolicyStore.Contract.TypeAndVersion(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreCallerSession) TypeAndVersion() (string, error) {
	return _ConnectionPolicyStore.Contract.TypeAndVersion(&_ConnectionPolicyStore.CallOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "acceptOwnership")
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.AcceptOwnership(&_ConnectionPolicyStore.TransactOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.AcceptOwnership(&_ConnectionPolicyStore.TransactOpts)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_ConnectionPolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyAuthorizedCallerUpdates(&_ConnectionPolicyStore.TransactOpts, authorizedCallerArgs)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyConnectionPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyConnectionPolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyConnectionPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyConnectionPolicyConfigUpdates(&_ConnectionPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyConnectionPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyConnectionPolicyConfigUpdates(&_ConnectionPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyInlineConnectionConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyInlineConnectionConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyInlineConnectionConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyInlineConnectionConfigUpdates(&_ConnectionPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyInlineConnectionConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyInlineConnectionConfigUpdates(&_ConnectionPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyInlineRemotePoolConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyInlineRemotePoolConfigUpdates", inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyInlineRemotePoolConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyInlineRemotePoolConfigUpdates(&_ConnectionPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyInlineRemotePoolConfigUpdates(inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyInlineRemotePoolConfigUpdates(&_ConnectionPolicyStore.TransactOpts, inlineConfigsToRemove, inlineConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyRemotePoolPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyRemotePoolPolicyConfigUpdates", policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyRemotePoolPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyRemotePoolPolicyConfigUpdates(&_ConnectionPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyRemotePoolPolicyConfigUpdates(policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyRemotePoolPolicyConfigUpdates(&_ConnectionPolicyStore.TransactOpts, policyConfigsToRemove, policiesToRemove, policyConfigsToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyRemoteTokenAddressProviderUpdates(opts *bind.TransactOpts, providersToSet []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyRemoteTokenAddressProviderUpdates", providersToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyRemoteTokenAddressProviderUpdates(providersToSet []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyRemoteTokenAddressProviderUpdates(&_ConnectionPolicyStore.TransactOpts, providersToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyRemoteTokenAddressProviderUpdates(providersToSet []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyRemoteTokenAddressProviderUpdates(&_ConnectionPolicyStore.TransactOpts, providersToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplySubsetTokenUpdates(opts *bind.TransactOpts, tokensToRemove []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applySubsetTokenUpdates", tokensToRemove, tokensToAdd)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplySubsetTokenUpdates(tokensToRemove []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplySubsetTokenUpdates(&_ConnectionPolicyStore.TransactOpts, tokensToRemove, tokensToAdd)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplySubsetTokenUpdates(tokensToRemove []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplySubsetTokenUpdates(&_ConnectionPolicyStore.TransactOpts, tokensToRemove, tokensToAdd)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyTokenConnectionPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyTokenConnectionPolicyUpdates", tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyTokenConnectionPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenConnectionPolicyUpdates(&_ConnectionPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyTokenConnectionPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenConnectionPolicyUpdates(&_ConnectionPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyTokenProfileUpdates", tokensToRemove, profilesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenProfileUpdates(&_ConnectionPolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyTokenProfileUpdates(tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenProfileUpdates(&_ConnectionPolicyStore.TransactOpts, tokensToRemove, profilesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) ApplyTokenRemotePoolPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "applyTokenRemotePoolPolicyUpdates", tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) ApplyTokenRemotePoolPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenRemotePoolPolicyUpdates(&_ConnectionPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) ApplyTokenRemotePoolPolicyUpdates(tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.ApplyTokenRemotePoolPolicyUpdates(&_ConnectionPolicyStore.TransactOpts, tokenPoliciesToSet)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "setChainPause", remoteChainSelectors, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) SetChainPause(remoteChainSelectors []uint64, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetChainPause(&_ConnectionPolicyStore.TransactOpts, remoteChainSelectors, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) SetChainPause(remoteChainSelectors []uint64, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetChainPause(&_ConnectionPolicyStore.TransactOpts, remoteChainSelectors, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) SetGlobalPause(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "setGlobalPause", isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) SetGlobalPause(isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetGlobalPause(&_ConnectionPolicyStore.TransactOpts, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) SetGlobalPause(isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetGlobalPause(&_ConnectionPolicyStore.TransactOpts, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) SetTokenPause(opts *bind.TransactOpts, localTokens []common.Address, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "setTokenPause", localTokens, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) SetTokenPause(localTokens []common.Address, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetTokenPause(&_ConnectionPolicyStore.TransactOpts, localTokens, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) SetTokenPause(localTokens []common.Address, direction uint8, isPaused bool) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.SetTokenPause(&_ConnectionPolicyStore.TransactOpts, localTokens, direction, isPaused)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ConnectionPolicyStore.contract.Transact(opts, "transferOwnership", to)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.TransferOwnership(&_ConnectionPolicyStore.TransactOpts, to)
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConnectionPolicyStore.Contract.TransferOwnership(&_ConnectionPolicyStore.TransactOpts, to)
}

type ConnectionPolicyStoreAuthorizedCallerAddedIterator struct {
	Event *ConnectionPolicyStoreAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreAuthorizedCallerAdded)
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
		it.Event = new(ConnectionPolicyStoreAuthorizedCallerAdded)
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

func (it *ConnectionPolicyStoreAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*ConnectionPolicyStoreAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreAuthorizedCallerAddedIterator{contract: _ConnectionPolicyStore.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreAuthorizedCallerAdded)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseAuthorizedCallerAdded(log types.Log) (*ConnectionPolicyStoreAuthorizedCallerAdded, error) {
	event := new(ConnectionPolicyStoreAuthorizedCallerAdded)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreAuthorizedCallerRemovedIterator struct {
	Event *ConnectionPolicyStoreAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreAuthorizedCallerRemoved)
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
		it.Event = new(ConnectionPolicyStoreAuthorizedCallerRemoved)
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

func (it *ConnectionPolicyStoreAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*ConnectionPolicyStoreAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreAuthorizedCallerRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreAuthorizedCallerRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*ConnectionPolicyStoreAuthorizedCallerRemoved, error) {
	event := new(ConnectionPolicyStoreAuthorizedCallerRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreChainPauseSetIterator struct {
	Event *ConnectionPolicyStoreChainPauseSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreChainPauseSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreChainPauseSet)
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
		it.Event = new(ConnectionPolicyStoreChainPauseSet)
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

func (it *ConnectionPolicyStoreChainPauseSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreChainPauseSet struct {
	RemoteChainSelector uint64
	IsOutboundPaused    bool
	IsInboundPaused     bool
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*ConnectionPolicyStoreChainPauseSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "ChainPauseSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreChainPauseSetIterator{contract: _ConnectionPolicyStore.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreChainPauseSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "ChainPauseSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreChainPauseSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseChainPauseSet(log types.Log) (*ConnectionPolicyStoreChainPauseSet, error) {
	event := new(ConnectionPolicyStoreChainPauseSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator struct {
	Event *ConnectionPolicyStoreConnectionPolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreConnectionPolicyConfigRemoved)
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
		it.Event = new(ConnectionPolicyStoreConnectionPolicyConfigRemoved)
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

func (it *ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreConnectionPolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterConnectionPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "ConnectionPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "ConnectionPolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchConnectionPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "ConnectionPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreConnectionPolicyConfigRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyConfigRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseConnectionPolicyConfigRemoved(log types.Log) (*ConnectionPolicyStoreConnectionPolicyConfigRemoved, error) {
	event := new(ConnectionPolicyStoreConnectionPolicyConfigRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreConnectionPolicyConfigSetIterator struct {
	Event *ConnectionPolicyStoreConnectionPolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreConnectionPolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreConnectionPolicyConfigSet)
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
		it.Event = new(ConnectionPolicyStoreConnectionPolicyConfigSet)
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

func (it *ConnectionPolicyStoreConnectionPolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreConnectionPolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreConnectionPolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	OnlySubsetTokens    bool
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterConnectionPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreConnectionPolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "ConnectionPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreConnectionPolicyConfigSetIterator{contract: _ConnectionPolicyStore.contract, event: "ConnectionPolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchConnectionPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "ConnectionPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreConnectionPolicyConfigSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyConfigSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseConnectionPolicyConfigSet(log types.Log) (*ConnectionPolicyStoreConnectionPolicyConfigSet, error) {
	event := new(ConnectionPolicyStoreConnectionPolicyConfigSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreConnectionPolicyRemovedIterator struct {
	Event *ConnectionPolicyStoreConnectionPolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreConnectionPolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreConnectionPolicyRemoved)
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
		it.Event = new(ConnectionPolicyStoreConnectionPolicyRemoved)
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

func (it *ConnectionPolicyStoreConnectionPolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreConnectionPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreConnectionPolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterConnectionPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreConnectionPolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "ConnectionPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreConnectionPolicyRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "ConnectionPolicyRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchConnectionPolicyRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "ConnectionPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreConnectionPolicyRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseConnectionPolicyRemoved(log types.Log) (*ConnectionPolicyStoreConnectionPolicyRemoved, error) {
	event := new(ConnectionPolicyStoreConnectionPolicyRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "ConnectionPolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreGlobalPauseSetIterator struct {
	Event *ConnectionPolicyStoreGlobalPauseSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreGlobalPauseSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreGlobalPauseSet)
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
		it.Event = new(ConnectionPolicyStoreGlobalPauseSet)
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

func (it *ConnectionPolicyStoreGlobalPauseSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreGlobalPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreGlobalPauseSet struct {
	IsPaused bool
	Raw      types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterGlobalPauseSet(opts *bind.FilterOpts) (*ConnectionPolicyStoreGlobalPauseSetIterator, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "GlobalPauseSet")
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreGlobalPauseSetIterator{contract: _ConnectionPolicyStore.contract, event: "GlobalPauseSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchGlobalPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreGlobalPauseSet) (event.Subscription, error) {

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "GlobalPauseSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreGlobalPauseSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "GlobalPauseSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseGlobalPauseSet(log types.Log) (*ConnectionPolicyStoreGlobalPauseSet, error) {
	event := new(ConnectionPolicyStoreGlobalPauseSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "GlobalPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreInlineConnectionConfigRemovedIterator struct {
	Event *ConnectionPolicyStoreInlineConnectionConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreInlineConnectionConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreInlineConnectionConfigRemoved)
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
		it.Event = new(ConnectionPolicyStoreInlineConnectionConfigRemoved)
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

func (it *ConnectionPolicyStoreInlineConnectionConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreInlineConnectionConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreInlineConnectionConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterInlineConnectionConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineConnectionConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "InlineConnectionConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreInlineConnectionConfigRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "InlineConnectionConfigRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchInlineConnectionConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineConnectionConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "InlineConnectionConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreInlineConnectionConfigRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineConnectionConfigRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseInlineConnectionConfigRemoved(log types.Log) (*ConnectionPolicyStoreInlineConnectionConfigRemoved, error) {
	event := new(ConnectionPolicyStoreInlineConnectionConfigRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineConnectionConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreInlineConnectionConfigSetIterator struct {
	Event *ConnectionPolicyStoreInlineConnectionConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreInlineConnectionConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreInlineConnectionConfigSet)
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
		it.Event = new(ConnectionPolicyStoreInlineConnectionConfigSet)
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

func (it *ConnectionPolicyStoreInlineConnectionConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreInlineConnectionConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreInlineConnectionConfigSet struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	RemoteToken         []byte
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterInlineConnectionConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineConnectionConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "InlineConnectionConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreInlineConnectionConfigSetIterator{contract: _ConnectionPolicyStore.contract, event: "InlineConnectionConfigSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchInlineConnectionConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineConnectionConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "InlineConnectionConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreInlineConnectionConfigSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineConnectionConfigSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseInlineConnectionConfigSet(log types.Log) (*ConnectionPolicyStoreInlineConnectionConfigSet, error) {
	event := new(ConnectionPolicyStoreInlineConnectionConfigSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineConnectionConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator struct {
	Event *ConnectionPolicyStoreInlineRemotePoolConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreInlineRemotePoolConfigRemoved)
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
		it.Event = new(ConnectionPolicyStoreInlineRemotePoolConfigRemoved)
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

func (it *ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreInlineRemotePoolConfigRemoved struct {
	LocalToken          common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterInlineRemotePoolConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "InlineRemotePoolConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "InlineRemotePoolConfigRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchInlineRemotePoolConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineRemotePoolConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "InlineRemotePoolConfigRemoved", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreInlineRemotePoolConfigRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineRemotePoolConfigRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseInlineRemotePoolConfigRemoved(log types.Log) (*ConnectionPolicyStoreInlineRemotePoolConfigRemoved, error) {
	event := new(ConnectionPolicyStoreInlineRemotePoolConfigRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineRemotePoolConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreInlineRemotePoolConfigSetIterator struct {
	Event *ConnectionPolicyStoreInlineRemotePoolConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreInlineRemotePoolConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreInlineRemotePoolConfigSet)
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
		it.Event = new(ConnectionPolicyStoreInlineRemotePoolConfigSet)
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

func (it *ConnectionPolicyStoreInlineRemotePoolConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreInlineRemotePoolConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreInlineRemotePoolConfigSet struct {
	LocalToken             common.Address
	RemoteChainSelector    uint64
	TrustPolicyRemotePools bool
	RemotePoolConfig       ConnectionPolicyStoreRemotePoolConfig
	Raw                    types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterInlineRemotePoolConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineRemotePoolConfigSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "InlineRemotePoolConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreInlineRemotePoolConfigSetIterator{contract: _ConnectionPolicyStore.contract, event: "InlineRemotePoolConfigSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchInlineRemotePoolConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineRemotePoolConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "InlineRemotePoolConfigSet", localTokenRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreInlineRemotePoolConfigSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineRemotePoolConfigSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseInlineRemotePoolConfigSet(log types.Log) (*ConnectionPolicyStoreInlineRemotePoolConfigSet, error) {
	event := new(ConnectionPolicyStoreInlineRemotePoolConfigSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "InlineRemotePoolConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreOwnershipTransferRequestedIterator struct {
	Event *ConnectionPolicyStoreOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreOwnershipTransferRequested)
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
		it.Event = new(ConnectionPolicyStoreOwnershipTransferRequested)
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

func (it *ConnectionPolicyStoreOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConnectionPolicyStoreOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreOwnershipTransferRequestedIterator{contract: _ConnectionPolicyStore.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreOwnershipTransferRequested)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseOwnershipTransferRequested(log types.Log) (*ConnectionPolicyStoreOwnershipTransferRequested, error) {
	event := new(ConnectionPolicyStoreOwnershipTransferRequested)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreOwnershipTransferredIterator struct {
	Event *ConnectionPolicyStoreOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreOwnershipTransferred)
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
		it.Event = new(ConnectionPolicyStoreOwnershipTransferred)
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

func (it *ConnectionPolicyStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConnectionPolicyStoreOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreOwnershipTransferredIterator{contract: _ConnectionPolicyStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreOwnershipTransferred)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseOwnershipTransferred(log types.Log) (*ConnectionPolicyStoreOwnershipTransferred, error) {
	event := new(ConnectionPolicyStoreOwnershipTransferred)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator struct {
	Event *ConnectionPolicyStoreRemotePoolPolicyConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreRemotePoolPolicyConfigRemoved)
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
		it.Event = new(ConnectionPolicyStoreRemotePoolPolicyConfigRemoved)
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

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreRemotePoolPolicyConfigRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterRemotePoolPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "RemotePoolPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "RemotePoolPolicyConfigRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchRemotePoolPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "RemotePoolPolicyConfigRemoved", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreRemotePoolPolicyConfigRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyConfigRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseRemotePoolPolicyConfigRemoved(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyConfigRemoved, error) {
	event := new(ConnectionPolicyStoreRemotePoolPolicyConfigRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator struct {
	Event *ConnectionPolicyStoreRemotePoolPolicyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreRemotePoolPolicyConfigSet)
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
		it.Event = new(ConnectionPolicyStoreRemotePoolPolicyConfigSet)
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

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreRemotePoolPolicyConfigSet struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	RemotePoolAddresses [][]byte
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterRemotePoolPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "RemotePoolPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator{contract: _ConnectionPolicyStore.contract, event: "RemotePoolPolicyConfigSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchRemotePoolPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "RemotePoolPolicyConfigSet", policyNameRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreRemotePoolPolicyConfigSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyConfigSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseRemotePoolPolicyConfigSet(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyConfigSet, error) {
	event := new(ConnectionPolicyStoreRemotePoolPolicyConfigSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreRemotePoolPolicyRemovedIterator struct {
	Event *ConnectionPolicyStoreRemotePoolPolicyRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreRemotePoolPolicyRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreRemotePoolPolicyRemoved)
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
		it.Event = new(ConnectionPolicyStoreRemotePoolPolicyRemoved)
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

func (it *ConnectionPolicyStoreRemotePoolPolicyRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreRemotePoolPolicyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreRemotePoolPolicyRemoved struct {
	PolicyName [32]byte
	Raw        types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterRemotePoolPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreRemotePoolPolicyRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "RemotePoolPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreRemotePoolPolicyRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "RemotePoolPolicyRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchRemotePoolPolicyRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyRemoved, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "RemotePoolPolicyRemoved", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreRemotePoolPolicyRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseRemotePoolPolicyRemoved(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyRemoved, error) {
	event := new(ConnectionPolicyStoreRemotePoolPolicyRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemotePoolPolicyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator struct {
	Event *ConnectionPolicyStoreRemoteTokenAddressProviderSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreRemoteTokenAddressProviderSet)
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
		it.Event = new(ConnectionPolicyStoreRemoteTokenAddressProviderSet)
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

func (it *ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreRemoteTokenAddressProviderSet struct {
	PolicyName                 [32]byte
	RemoteTokenAddressProvider common.Address
	Raw                        types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterRemoteTokenAddressProviderSet(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "RemoteTokenAddressProviderSet", policyNameRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator{contract: _ConnectionPolicyStore.contract, event: "RemoteTokenAddressProviderSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchRemoteTokenAddressProviderSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemoteTokenAddressProviderSet, policyName [][32]byte) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "RemoteTokenAddressProviderSet", policyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreRemoteTokenAddressProviderSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemoteTokenAddressProviderSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseRemoteTokenAddressProviderSet(log types.Log) (*ConnectionPolicyStoreRemoteTokenAddressProviderSet, error) {
	event := new(ConnectionPolicyStoreRemoteTokenAddressProviderSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "RemoteTokenAddressProviderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreSubsetTokenAddedIterator struct {
	Event *ConnectionPolicyStoreSubsetTokenAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreSubsetTokenAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreSubsetTokenAdded)
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
		it.Event = new(ConnectionPolicyStoreSubsetTokenAdded)
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

func (it *ConnectionPolicyStoreSubsetTokenAddedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreSubsetTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreSubsetTokenAdded struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalToken          common.Address
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterSubsetTokenAdded(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (*ConnectionPolicyStoreSubsetTokenAddedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "SubsetTokenAdded", policyNameRule, remoteChainSelectorRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreSubsetTokenAddedIterator{contract: _ConnectionPolicyStore.contract, event: "SubsetTokenAdded", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchSubsetTokenAdded(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreSubsetTokenAdded, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "SubsetTokenAdded", policyNameRule, remoteChainSelectorRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreSubsetTokenAdded)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "SubsetTokenAdded", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseSubsetTokenAdded(log types.Log) (*ConnectionPolicyStoreSubsetTokenAdded, error) {
	event := new(ConnectionPolicyStoreSubsetTokenAdded)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "SubsetTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreSubsetTokenRemovedIterator struct {
	Event *ConnectionPolicyStoreSubsetTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreSubsetTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreSubsetTokenRemoved)
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
		it.Event = new(ConnectionPolicyStoreSubsetTokenRemoved)
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

func (it *ConnectionPolicyStoreSubsetTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreSubsetTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreSubsetTokenRemoved struct {
	PolicyName          [32]byte
	RemoteChainSelector uint64
	LocalToken          common.Address
	Raw                 types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterSubsetTokenRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (*ConnectionPolicyStoreSubsetTokenRemovedIterator, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "SubsetTokenRemoved", policyNameRule, remoteChainSelectorRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreSubsetTokenRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "SubsetTokenRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchSubsetTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreSubsetTokenRemoved, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (event.Subscription, error) {

	var policyNameRule []interface{}
	for _, policyNameItem := range policyName {
		policyNameRule = append(policyNameRule, policyNameItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "SubsetTokenRemoved", policyNameRule, remoteChainSelectorRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreSubsetTokenRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "SubsetTokenRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseSubsetTokenRemoved(log types.Log) (*ConnectionPolicyStoreSubsetTokenRemoved, error) {
	event := new(ConnectionPolicyStoreSubsetTokenRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "SubsetTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreTokenConnectionPolicySetIterator struct {
	Event *ConnectionPolicyStoreTokenConnectionPolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreTokenConnectionPolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreTokenConnectionPolicySet)
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
		it.Event = new(ConnectionPolicyStoreTokenConnectionPolicySet)
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

func (it *ConnectionPolicyStoreTokenConnectionPolicySetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreTokenConnectionPolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreTokenConnectionPolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterTokenConnectionPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*ConnectionPolicyStoreTokenConnectionPolicySetIterator, error) {

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

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "TokenConnectionPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTokenConnectionPolicySetIterator{contract: _ConnectionPolicyStore.contract, event: "TokenConnectionPolicySet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchTokenConnectionPolicySet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenConnectionPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "TokenConnectionPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreTokenConnectionPolicySet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenConnectionPolicySet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseTokenConnectionPolicySet(log types.Log) (*ConnectionPolicyStoreTokenConnectionPolicySet, error) {
	event := new(ConnectionPolicyStoreTokenConnectionPolicySet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenConnectionPolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreTokenPauseSetIterator struct {
	Event *ConnectionPolicyStoreTokenPauseSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreTokenPauseSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreTokenPauseSet)
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
		it.Event = new(ConnectionPolicyStoreTokenPauseSet)
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

func (it *ConnectionPolicyStoreTokenPauseSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreTokenPauseSet struct {
	LocalToken       common.Address
	IsOutboundPaused bool
	IsInboundPaused  bool
	Raw              types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenPauseSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "TokenPauseSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTokenPauseSetIterator{contract: _ConnectionPolicyStore.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenPauseSet, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "TokenPauseSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreTokenPauseSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseTokenPauseSet(log types.Log) (*ConnectionPolicyStoreTokenPauseSet, error) {
	event := new(ConnectionPolicyStoreTokenPauseSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreTokenProfileSetIterator struct {
	Event *ConnectionPolicyStoreTokenProfileSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreTokenProfileSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreTokenProfileSet)
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
		it.Event = new(ConnectionPolicyStoreTokenProfileSet)
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

func (it *ConnectionPolicyStoreTokenProfileSetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreTokenProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreTokenProfileSet struct {
	LocalToken                 common.Address
	LocalDecimals              uint8
	LockOrBurnFlavor           uint8
	LockOrBurnAddress          common.Address
	AllowedFinalityConfig      [4]byte
	AdditionalPoolDataProvider common.Address
	Raw                        types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenProfileSetIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTokenProfileSetIterator{contract: _ConnectionPolicyStore.contract, event: "TokenProfileSet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "TokenProfileSet", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreTokenProfileSet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseTokenProfileSet(log types.Log) (*ConnectionPolicyStoreTokenProfileSet, error) {
	event := new(ConnectionPolicyStoreTokenProfileSet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreTokenRemotePoolPolicySetIterator struct {
	Event *ConnectionPolicyStoreTokenRemotePoolPolicySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreTokenRemotePoolPolicySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreTokenRemotePoolPolicySet)
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
		it.Event = new(ConnectionPolicyStoreTokenRemotePoolPolicySet)
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

func (it *ConnectionPolicyStoreTokenRemotePoolPolicySetIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreTokenRemotePoolPolicySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreTokenRemotePoolPolicySet struct {
	LocalToken    common.Address
	OldPolicyName [32]byte
	NewPolicyName [32]byte
	Raw           types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterTokenRemotePoolPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*ConnectionPolicyStoreTokenRemotePoolPolicySetIterator, error) {

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

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "TokenRemotePoolPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTokenRemotePoolPolicySetIterator{contract: _ConnectionPolicyStore.contract, event: "TokenRemotePoolPolicySet", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchTokenRemotePoolPolicySet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenRemotePoolPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "TokenRemotePoolPolicySet", localTokenRule, oldPolicyNameRule, newPolicyNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreTokenRemotePoolPolicySet)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenRemotePoolPolicySet", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseTokenRemotePoolPolicySet(log types.Log) (*ConnectionPolicyStoreTokenRemotePoolPolicySet, error) {
	event := new(ConnectionPolicyStoreTokenRemotePoolPolicySet)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenRemotePoolPolicySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConnectionPolicyStoreTokenRemovedIterator struct {
	Event *ConnectionPolicyStoreTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConnectionPolicyStoreTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectionPolicyStoreTokenRemoved)
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
		it.Event = new(ConnectionPolicyStoreTokenRemoved)
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

func (it *ConnectionPolicyStoreTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *ConnectionPolicyStoreTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConnectionPolicyStoreTokenRemoved struct {
	LocalToken common.Address
	Raw        types.Log
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenRemovedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.FilterLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectionPolicyStoreTokenRemovedIterator{contract: _ConnectionPolicyStore.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _ConnectionPolicyStore.contract.WatchLogs(opts, "TokenRemoved", localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConnectionPolicyStoreTokenRemoved)
				if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

func (_ConnectionPolicyStore *ConnectionPolicyStoreFilterer) ParseTokenRemoved(log types.Log) (*ConnectionPolicyStoreTokenRemoved, error) {
	event := new(ConnectionPolicyStoreTokenRemoved)
	if err := _ConnectionPolicyStore.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetConnectionPolicy struct {
	TokenCount                 uint32
	RemoteTokenAddressProvider common.Address
	PolicyConfigs              []ConnectionPolicyStoreConnectionPolicyConfigArgs
}
type GetLaneRecordLevelBatch struct {
	ConnectionLevels []uint8
	RemotePoolLevels []uint8
}
type GetPolicyNames struct {
	ConnectionPolicyNames [][32]byte
	RemotePoolPolicyNames [][32]byte
}
type GetRemotePoolPolicy struct {
	TokenCount    uint32
	PolicyConfigs []ConnectionPolicyStoreRemotePoolPolicyConfigArgs
}
type GetSubsetTokens struct {
	OnlySubsetTokens bool
	SubsetTokens     []common.Address
}
type GetTokenPolicies struct {
	IsRegistered         bool
	ConnectionPolicyName [32]byte
	RemotePoolPolicyName [32]byte
}

func (ConnectionPolicyStoreAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (ConnectionPolicyStoreAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (ConnectionPolicyStoreChainPauseSet) Topic() common.Hash {
	return common.HexToHash("0xbe188ef34fe3f31a5ec7e7736b85eace0efffe0b9bc88bca91183b6cc57b6e3a")
}

func (ConnectionPolicyStoreConnectionPolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0xe661c0fbd2356ac3d908ab14a46a8f7bbd4fe408766ba9f88232d759e623aecd")
}

func (ConnectionPolicyStoreConnectionPolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb778117d916c7db374b0ffb4704562ee1bbd1c0d29f15ca7faf8dd73e4b3cb46")
}

func (ConnectionPolicyStoreConnectionPolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0xdf9c4e5139f4220985ce3b4a2959cadc2fd5187529572bbf2103ccf1594900d5")
}

func (ConnectionPolicyStoreGlobalPauseSet) Topic() common.Hash {
	return common.HexToHash("0x971be26fcde2df81b2b63dfb7fa1d2294b524eb1896b10e5228d9c154a30af09")
}

func (ConnectionPolicyStoreInlineConnectionConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x5da1ddc7f5f2ac7cb20f1eca93c22a2f7f7ca74391cc0107a21af4f45b0f0c97")
}

func (ConnectionPolicyStoreInlineConnectionConfigSet) Topic() common.Hash {
	return common.HexToHash("0xc87011fae1db3e00505551d507555885d0f82b810c56058714645bfc51c148e9")
}

func (ConnectionPolicyStoreInlineRemotePoolConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x3260489f3856053ff7ff5e5e25474af85e3fd98c32a5556afb470aa44a645990")
}

func (ConnectionPolicyStoreInlineRemotePoolConfigSet) Topic() common.Hash {
	return common.HexToHash("0x06ad59388432d80a26d3b9660cd30971ebc487fc2784334c556516f4ea1527a5")
}

func (ConnectionPolicyStoreOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ConnectionPolicyStoreOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ConnectionPolicyStoreRemotePoolPolicyConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0x16408cc867e3bbc60e589be5073a0c709e121c7b2a29f97cb9219ceef1a9c3a6")
}

func (ConnectionPolicyStoreRemotePoolPolicyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f7ae29ee595ae7323f9cf53122d3177d8ba5e32e0702ab492a57d428008d063")
}

func (ConnectionPolicyStoreRemotePoolPolicyRemoved) Topic() common.Hash {
	return common.HexToHash("0x51d9277f7ad33808d961e53406ceba32abe5dc730a752f828bc34b65e29bfa0c")
}

func (ConnectionPolicyStoreRemoteTokenAddressProviderSet) Topic() common.Hash {
	return common.HexToHash("0xafc9972619c33aef176a4a2930d18d6b9a2e1f370d79e37174979286d4c30627")
}

func (ConnectionPolicyStoreSubsetTokenAdded) Topic() common.Hash {
	return common.HexToHash("0x99ca8bf1818fe2a9cc5c1f21ae12d88525e91dbcfb2f6f672ba80e5184a2c6cf")
}

func (ConnectionPolicyStoreSubsetTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x664d2b667751d4835f78f39fd2c632484f01ef70b0b71418169fa73e4c2f0cbb")
}

func (ConnectionPolicyStoreTokenConnectionPolicySet) Topic() common.Hash {
	return common.HexToHash("0xe07ad45e782b60daeb19005cd93d17910aa6611d9df3414d4929f596966ce183")
}

func (ConnectionPolicyStoreTokenPauseSet) Topic() common.Hash {
	return common.HexToHash("0xf53d2e953b6972d456e89380b010f53d5adb5fcf2f9813755b77011cdebb6d4b")
}

func (ConnectionPolicyStoreTokenProfileSet) Topic() common.Hash {
	return common.HexToHash("0x1a43e9c5a0da3a3ef84120fc939eac76d74623dfc040b98779c630eba40223dd")
}

func (ConnectionPolicyStoreTokenRemotePoolPolicySet) Topic() common.Hash {
	return common.HexToHash("0xb3a957a117693885aebbdb23caa8922911f57385251d8c43e5d10e6015db37ce")
}

func (ConnectionPolicyStoreTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3")
}

func (_ConnectionPolicyStore *ConnectionPolicyStore) Address() common.Address {
	return _ConnectionPolicyStore.address
}

type ConnectionPolicyStoreInterface interface {
	CheckLaneRecords(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64) error

	CheckRequestedFinalityBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey, requestedFinalityConfig [4]byte) error

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllConfiguredTokens(opts *bind.CallOpts, startIndex uint64, maxCount uint64) ([]common.Address, error)

	GetConnectionPolicy(opts *bind.CallOpts, policyName [32]byte) (GetConnectionPolicy,

		error)

	GetInlineConnectionConfigs(opts *bind.CallOpts, localToken common.Address) ([]ConnectionPolicyStoreInlineConnectionConfigArgs, error)

	GetInlineRemotePoolConfigs(opts *bind.CallOpts, localToken common.Address) ([]ConnectionPolicyStoreInlineRemotePoolConfigArgs, error)

	GetLaneConnectionBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]IPoolConfigManagerLaneConnection, error)

	GetLaneRecordLevelBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) (GetLaneRecordLevelBatch,

		error)

	GetPolicyNames(opts *bind.CallOpts) (GetPolicyNames,

		error)

	GetRemotePoolPolicy(opts *bind.CallOpts, policyName [32]byte) (GetRemotePoolPolicy,

		error)

	GetSubsetTokens(opts *bind.CallOpts, policyName [32]byte, remoteChainSelector uint64, startIndex uint64, maxCount uint64) (GetSubsetTokens,

		error)

	GetTokenConfigBatch(opts *bind.CallOpts, localTokens []common.Address) ([]IPoolConfigManagerTokenConfig, error)

	GetTokenPolicies(opts *bind.CallOpts, localToken common.Address) (GetTokenPolicies,

		error)

	GetTokenProfiles(opts *bind.CallOpts, localTokens []common.Address) ([]ConnectionPolicyStoreTokenProfile, error)

	HasLaneBatch(opts *bind.CallOpts, lanes []IPoolConfigManagerLaneKey) ([]bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ResolveInboundLane(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourcePoolAddress []byte, sourcePoolData []byte, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte) (IPoolConfigManagerResolvedInboundTransfer, error)

	ResolveOutboundLane(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, requestedFinalityConfig [4]byte, tokenArgs []byte) (IPoolConfigManagerOutboundLane, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyConnectionPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreConnectionPolicyConfigArgs) (*types.Transaction, error)

	ApplyInlineConnectionConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineConnectionConfigArgs) (*types.Transaction, error)

	ApplyInlineRemotePoolConfigUpdates(opts *bind.TransactOpts, inlineConfigsToRemove []IPoolConfigManagerLaneKey, inlineConfigsToSet []ConnectionPolicyStoreInlineRemotePoolConfigArgs) (*types.Transaction, error)

	ApplyRemotePoolPolicyConfigUpdates(opts *bind.TransactOpts, policyConfigsToRemove []PolicyConfigPolicyRecordKey, policiesToRemove [][32]byte, policyConfigsToSet []ConnectionPolicyStoreRemotePoolPolicyConfigArgs) (*types.Transaction, error)

	ApplyRemoteTokenAddressProviderUpdates(opts *bind.TransactOpts, providersToSet []ConnectionPolicyStoreRemoteTokenAddressProviderArgs) (*types.Transaction, error)

	ApplySubsetTokenUpdates(opts *bind.TransactOpts, tokensToRemove []ConnectionPolicyStoreSubsetTokenArgs, tokensToAdd []ConnectionPolicyStoreSubsetTokenArgs) (*types.Transaction, error)

	ApplyTokenConnectionPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	ApplyTokenProfileUpdates(opts *bind.TransactOpts, tokensToRemove []common.Address, profilesToSet []PolicyConfigTokenProfileArgs) (*types.Transaction, error)

	ApplyTokenRemotePoolPolicyUpdates(opts *bind.TransactOpts, tokenPoliciesToSet []PolicyConfigTokenPolicyArgs) (*types.Transaction, error)

	SetChainPause(opts *bind.TransactOpts, remoteChainSelectors []uint64, direction uint8, isPaused bool) (*types.Transaction, error)

	SetGlobalPause(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error)

	SetTokenPause(opts *bind.TransactOpts, localTokens []common.Address, direction uint8, isPaused bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*ConnectionPolicyStoreAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*ConnectionPolicyStoreAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*ConnectionPolicyStoreAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*ConnectionPolicyStoreAuthorizedCallerRemoved, error)

	FilterChainPauseSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*ConnectionPolicyStoreChainPauseSetIterator, error)

	WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreChainPauseSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseChainPauseSet(log types.Log) (*ConnectionPolicyStoreChainPauseSet, error)

	FilterConnectionPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreConnectionPolicyConfigRemovedIterator, error)

	WatchConnectionPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseConnectionPolicyConfigRemoved(log types.Log) (*ConnectionPolicyStoreConnectionPolicyConfigRemoved, error)

	FilterConnectionPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreConnectionPolicyConfigSetIterator, error)

	WatchConnectionPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseConnectionPolicyConfigSet(log types.Log) (*ConnectionPolicyStoreConnectionPolicyConfigSet, error)

	FilterConnectionPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreConnectionPolicyRemovedIterator, error)

	WatchConnectionPolicyRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreConnectionPolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseConnectionPolicyRemoved(log types.Log) (*ConnectionPolicyStoreConnectionPolicyRemoved, error)

	FilterGlobalPauseSet(opts *bind.FilterOpts) (*ConnectionPolicyStoreGlobalPauseSetIterator, error)

	WatchGlobalPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreGlobalPauseSet) (event.Subscription, error)

	ParseGlobalPauseSet(log types.Log) (*ConnectionPolicyStoreGlobalPauseSet, error)

	FilterInlineConnectionConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineConnectionConfigRemovedIterator, error)

	WatchInlineConnectionConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineConnectionConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineConnectionConfigRemoved(log types.Log) (*ConnectionPolicyStoreInlineConnectionConfigRemoved, error)

	FilterInlineConnectionConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineConnectionConfigSetIterator, error)

	WatchInlineConnectionConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineConnectionConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineConnectionConfigSet(log types.Log) (*ConnectionPolicyStoreInlineConnectionConfigSet, error)

	FilterInlineRemotePoolConfigRemoved(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineRemotePoolConfigRemovedIterator, error)

	WatchInlineRemotePoolConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineRemotePoolConfigRemoved, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineRemotePoolConfigRemoved(log types.Log) (*ConnectionPolicyStoreInlineRemotePoolConfigRemoved, error)

	FilterInlineRemotePoolConfigSet(opts *bind.FilterOpts, localToken []common.Address, remoteChainSelector []uint64) (*ConnectionPolicyStoreInlineRemotePoolConfigSetIterator, error)

	WatchInlineRemotePoolConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreInlineRemotePoolConfigSet, localToken []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInlineRemotePoolConfigSet(log types.Log) (*ConnectionPolicyStoreInlineRemotePoolConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConnectionPolicyStoreOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ConnectionPolicyStoreOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConnectionPolicyStoreOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ConnectionPolicyStoreOwnershipTransferred, error)

	FilterRemotePoolPolicyConfigRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreRemotePoolPolicyConfigRemovedIterator, error)

	WatchRemotePoolPolicyConfigRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyConfigRemoved, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolPolicyConfigRemoved(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyConfigRemoved, error)

	FilterRemotePoolPolicyConfigSet(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64) (*ConnectionPolicyStoreRemotePoolPolicyConfigSetIterator, error)

	WatchRemotePoolPolicyConfigSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyConfigSet, policyName [][32]byte, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolPolicyConfigSet(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyConfigSet, error)

	FilterRemotePoolPolicyRemoved(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreRemotePoolPolicyRemovedIterator, error)

	WatchRemotePoolPolicyRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemotePoolPolicyRemoved, policyName [][32]byte) (event.Subscription, error)

	ParseRemotePoolPolicyRemoved(log types.Log) (*ConnectionPolicyStoreRemotePoolPolicyRemoved, error)

	FilterRemoteTokenAddressProviderSet(opts *bind.FilterOpts, policyName [][32]byte) (*ConnectionPolicyStoreRemoteTokenAddressProviderSetIterator, error)

	WatchRemoteTokenAddressProviderSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreRemoteTokenAddressProviderSet, policyName [][32]byte) (event.Subscription, error)

	ParseRemoteTokenAddressProviderSet(log types.Log) (*ConnectionPolicyStoreRemoteTokenAddressProviderSet, error)

	FilterSubsetTokenAdded(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (*ConnectionPolicyStoreSubsetTokenAddedIterator, error)

	WatchSubsetTokenAdded(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreSubsetTokenAdded, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (event.Subscription, error)

	ParseSubsetTokenAdded(log types.Log) (*ConnectionPolicyStoreSubsetTokenAdded, error)

	FilterSubsetTokenRemoved(opts *bind.FilterOpts, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (*ConnectionPolicyStoreSubsetTokenRemovedIterator, error)

	WatchSubsetTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreSubsetTokenRemoved, policyName [][32]byte, remoteChainSelector []uint64, localToken []common.Address) (event.Subscription, error)

	ParseSubsetTokenRemoved(log types.Log) (*ConnectionPolicyStoreSubsetTokenRemoved, error)

	FilterTokenConnectionPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*ConnectionPolicyStoreTokenConnectionPolicySetIterator, error)

	WatchTokenConnectionPolicySet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenConnectionPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenConnectionPolicySet(log types.Log) (*ConnectionPolicyStoreTokenConnectionPolicySet, error)

	FilterTokenPauseSet(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenPauseSetIterator, error)

	WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenPauseSet, localToken []common.Address) (event.Subscription, error)

	ParseTokenPauseSet(log types.Log) (*ConnectionPolicyStoreTokenPauseSet, error)

	FilterTokenProfileSet(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenProfileSetIterator, error)

	WatchTokenProfileSet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenProfileSet, localToken []common.Address) (event.Subscription, error)

	ParseTokenProfileSet(log types.Log) (*ConnectionPolicyStoreTokenProfileSet, error)

	FilterTokenRemotePoolPolicySet(opts *bind.FilterOpts, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (*ConnectionPolicyStoreTokenRemotePoolPolicySetIterator, error)

	WatchTokenRemotePoolPolicySet(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenRemotePoolPolicySet, localToken []common.Address, oldPolicyName [][32]byte, newPolicyName [][32]byte) (event.Subscription, error)

	ParseTokenRemotePoolPolicySet(log types.Log) (*ConnectionPolicyStoreTokenRemotePoolPolicySet, error)

	FilterTokenRemoved(opts *bind.FilterOpts, localToken []common.Address) (*ConnectionPolicyStoreTokenRemovedIterator, error)

	WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *ConnectionPolicyStoreTokenRemoved, localToken []common.Address) (event.Subscription, error)

	ParseTokenRemoved(log types.Log) (*ConnectionPolicyStoreTokenRemoved, error)

	Address() common.Address
}
