# CCIP Solidity Protocol Documentation

## Overview

The Chainlink Cross-Chain Interoperability Protocol (CCIP) is a trust-minimized cross-chain messaging protocol that enables secure communication and token transfers between blockchain networks. This document provides a comprehensive understanding of the CCIP onchain architecture, its core contracts, and how they interact with offchain components.

## Table of Contents

1. [Core Architecture](#core-architecture)
2. [Key Components](#key-components)
3. [Message Flow](#message-flow)
4. [Contract Deep Dive](#contract-deep-dive)
5. [Token Transfer Mechanisms](#token-transfer-mechanisms)
6. [Security & Risk Management](#security--risk-management)
7. [Setup & Configuration](#setup--configuration)

## Core Architecture

CCIP operates through a combination of onchain contracts and offchain Distributed Oracle Networks (DONs). The protocol ensures secure, ordered, and verifiable cross-chain communication through a two-phase process:

1. **Commit Phase**: Messages are collected, merkle roots are generated, and committed onchain
2. **Execute Phase**: Messages are validated against committed roots and executed on the destination chain

## Key Components

### 1. Router (`Router.sol`)

The Router serves as the primary entry point for CCIP on each chain:

- **For Sending**: Accepts `ccipSend()` calls from users
- **For Receiving**: Delivers messages to receiver contracts via `ccipReceive()`
- **Configuration**: Maintains mappings of destination chains to OnRamps and source chains to OffRamps

Key functions:

- `ccipSend(uint64 destChainSelector, Client.EVM2AnyMessage message)`: Send cross-chain message
- `routeMessage(Client.Any2EVMMessage message, ...)`: Route incoming message to receiver

### 2. OnRamp (`onRamp/OnRamp.sol`)

Handles outbound messages on the source chain:

- **Message Processing**: Validates and formats messages for cross-chain transfer
- **Fee Calculation**: Queries FeeQuoter for message fees
- **Nonce Management**: Obtains sequential nonces from NonceManager
- **Event Emission**: Emits `CCIPMessageSent` events for offchain DON observation

Key components:

- Static Config: Chain selector, RMN remote, NonceManager, TokenAdminRegistry
- Dynamic Config: FeeQuoter, message interceptor, fee aggregator
- Destination Chain Config: Enabled status, max message size

### 3. OffRamp (`offRamp/OffRamp.sol`)

Handles inbound messages on the destination chain:

- **Message Validation**: Verifies messages against committed merkle roots
- **Execution Management**: Tracks message execution status
- **OCR3 Integration**: Implements commit and execution plugins
- **Nonce Updates**: Increments receiver nonces via NonceManager

Key features:

- Dual OCR3 configuration (Commit and Execution plugins)
- Source chain configuration management
- Message state tracking (untouched → in-progress → success/failure)

### 4. FeeQuoter (`FeeQuoter.sol`)

Manages fee calculation and pricing:

- **Fee Components**: Base fee + per-token transfer fees
- **Price Updates**: Maintains gas and token prices
- **Premium Calculation**: Applies configurable premiums
- **Destination Chain Config**: Stores fee parameters per destination

Key functions:

- `getTokenTransferFee()`: Calculate token-specific transfer fees
- `getValidatedFee()`: Get and validate total message fee
- `updatePrices()`: Update gas/token prices (called by OffRamp)

### 5. NonceManager (`NonceManager.sol`)

Ensures message ordering and prevents replay attacks:

- **Outbound Nonces**: Sequential per sender per destination
- **Inbound Nonces**: Tracks per sender per source
- **Authorized Callers**: Only OnRamp/OffRamp can update nonces
- **Nonce Validation**: Ensures strict ordering

### 6. TokenAdminRegistry (`tokenAdminRegistry/TokenAdminRegistry.sol`)

Central registry for token pool management:

- **Pool Registration**: Maps tokens to their official pools
- **Access Control**: Admin/owner roles for pool configuration
- **Pool Factory Integration**: Supports factory-deployed pools

### 7. Token Pools

Handle token transfers across chains with different mechanisms:

#### Base Types:

- **LockReleaseTokenPool**: Locks tokens on source, releases on destination
- **BurnMintTokenPool**: Burns tokens on source, mints on destination
- **BurnMintTokenPoolAbstract**: Base contract for burn/mint pools

#### Specialized Pools:

- **USDC Pools**: Integration with Circle's CCTP
- **Fast Transfer Pools**: Optimized for specific use cases
- **Siloed Pools**: Isolated liquidity management

## Message Flow

### 1. Sending a Message

```
User → Router.ccipSend() → OnRamp.forwardFromRouter() → Event Emission
```

1. User calls `Router.ccipSend()` with destination and message
2. Router validates destination chain support
3. Router transfers tokens to appropriate pools
4. Router forwards to OnRamp via `forwardFromRouter()`
5. OnRamp gets nonce from NonceManager
6. OnRamp calculates fees via FeeQuoter
7. OnRamp calls TokenPools to lockOrBurn tokens
8. OnRamp emits `CCIPMessageSent` event

### 2. Committing Messages (Offchain → Onchain)

```
DON Observes Events → Generate Merkle Root → Create Commit Report → OffRamp.commit()
```

1. Offchain DON nodes observe OnRamp events
2. Commit plugin aggregates messages and creates merkle roots
3. DON reaches consensus on commit report
4. Commit report sent to OffRamp
5. RMN provides additional validation

### 3. Executing Messages

```
DON Creates Execution Report → OffRamp.execute() → Router.routeMessage() → Receiver.ccipReceive()
```

1. Execution plugin identifies messages ready for execution
2. Creates execution report with merkle proofs
3. OffRamp validates proofs against committed roots
4. OffRamp updates nonces via NonceManager
5. OffRamp calls TokenPools to releaseOrMint tokens to receiver
6. OffRamp calls Router to deliver message
7. Router invokes receiver's `ccipReceive()`

## Contract Deep Dive

### Router Implementation Pseudo Code

```solidity
contract Router {
    // Destination chain selector → OnRamp address
    mapping(uint64 => address) private s_onRamps;

    // Source chain selector → OffRamp address
    mapping(uint64 => address) private s_offRamps;

    function ccipSend(
        uint64 destChainSelector,
        Client.EVM2AnyMessage calldata message
    ) external returns (bytes32 messageId) {
        // Validate destination chain
        address onRamp = s_onRamps[destChainSelector];

        // Transfer tokens to pools
        for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
            IERC20(token).safeTransferFrom(msg.sender, pool, amount);
        }

        // Forward to OnRamp
        return IOnRamp(onRamp).forwardFromRouter(message, feeTokenAmount);
    }
}
```

### OnRamp Message Processing Pseudo Code

```solidity
contract OnRamp {
    function forwardFromRouter(
        Client.EVM2AnyMessage calldata message,
        uint256 feeTokenAmount
    ) external returns (bytes32) {
        // Get nonce
        uint64 nonce = INonceManager(s_nonceManager)
            .getOutboundNonce(message.sender, destChainSelector);

        // Calculate fees
        uint256 fee = IFeeQuoter(s_feeQuoter)
            .getValidatedFee(destChainSelector, message);

        // Build internal message
        Internal.EVM2AnyRampMessage memory rampMessage = Internal.EVM2AnyRampMessage({
            header: Internal.RampMessageHeader({
                messageId: keccak256(abi.encode(message)),
                sourceChainSelector: s_chainSelector,
                destChainSelector: destChainSelector,
                sequenceNumber: ++s_sequenceNumber,
                nonce: nonce
            }),
            sender: message.sender,
            data: message.data,
            tokenAmounts: message.tokenAmounts,
            extraArgs: message.extraArgs
        });

        // Emit for offchain observation
        emit CCIPMessageSent(rampMessage);

        return rampMessage.header.messageId;
    }
}
```

### OffRamp Execution Pseudo Code

```solidity
contract OffRamp {
    function execute(
        ExecutionReport memory report,
        bytes memory signatures
    ) external {
        // Verify OCR signatures
        _verifyOCRSignatures(report, signatures);

        // Process each message
        for (uint256 i = 0; i < report.messages.length; ++i) {
            Any2EVMRampMessage memory message = report.messages[i];

            // Verify merkle proof
            _verifyMerkleProof(message, report.proofs[i]);

            // Update message state
            s_messageStatus[message.header.messageId] = MSG_STATE_IN_PROGRESS;

            // Release tokens from pools
            _releaseTokens(message.tokenAmounts, message.receiver);

            // Update nonce
            INonceManager(s_nonceManager).incrementInboundNonce(
                message.header.sourceChainSelector,
                message.sender
            );

            // Route to receiver
            IRouter(s_router).routeMessage(
                _convertToClient(message),
                message.header.sourceChainSelector,
                message.receiver,
                message.gasLimit
            );
        }
    }
}
```

## Token Transfer Mechanisms

CCIP uses a unified interface for all token pools with two main functions:

### Pool Interface

```solidity
interface IPoolV1 {
    function lockOrBurn(
        Pool.LockOrBurnInV1 calldata lockOrBurnIn
    ) external returns (Pool.LockOrBurnOutV1 memory);

    function releaseOrMint(
        Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
    ) external returns (Pool.ReleaseOrMintOutV1 memory);
}
```

### Pool Data Structures

```solidity
struct LockOrBurnInV1 {
    bytes receiver;                // Recipient on destination chain
    uint64 remoteChainSelector;    // Destination chain ID
    address originalSender;        // Original transaction sender
    uint256 amount;               // Amount to lock/burn
    address localToken;           // Token address on source chain
}

struct ReleaseOrMintInV1 {
    bytes originalSender;         // Original sender from source chain
    uint64 remoteChainSelector;   // Source chain ID
    address receiver;             // Recipient on destination chain
    uint256 sourceDenominatedAmount; // Amount in source token decimals
    address localToken;           // Token address on destination chain
    bytes sourcePoolAddress;      // Source pool address (for validation)
    bytes sourcePoolData;         // Additional data from source pool
    bytes offchainTokenData;      // Offchain data (untrusted)
}
```

### Lock & Release Mechanism

Used for wrapped tokens or native protocol tokens that maintain liquidity:

```solidity
contract LockReleaseTokenPool is TokenPool {
    // Source chain: Lock tokens (no override needed, base class handles transfer)
    // The base TokenPool.lockOrBurn() handles validation and emits events
    // Tokens are transferred to the pool via Router before this is called

    // Destination chain: Release tokens
    function _releaseOrMint(address receiver, uint256 amount) internal override {
        i_token.safeTransfer(receiver, amount);
    }

    // Liquidity management functions
    function provideLiquidity(uint256 amount) external;
    function withdrawLiquidity(uint256 amount) external;
}
```

### Burn & Mint Mechanism

Used for tokens with cross-chain mint/burn capabilities:

```solidity
contract BurnMintTokenPool is BurnMintTokenPoolAbstract {
    // Source chain: Burn tokens
    function _lockOrBurn(uint256 amount) internal virtual override {
        IBurnMintERC20(address(i_token)).burn(amount);
    }

    // Destination chain: Mint tokens (inherited from BurnMintTokenPoolAbstract)
    // function _releaseOrMint is inherited from BurnMintTokenPoolAbstract
}

abstract contract BurnMintTokenPoolAbstract is TokenPool {
    function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
        IBurnMintERC20(address(i_token)).mint(receiver, amount);
    }
}
```

### Specialized Burn Variants

- **BurnFromMintTokenPool**: Uses `burnFrom(msg.sender, amount)`
- **BurnWithFromMintTokenPool**: Uses `burn(msg.sender, amount)`
- **BurnToAddressMintTokenPool**: Uses configurable burn address

### Token Data Flow Between Chains

The CCIP protocol carefully orchestrates how data flows from source to destination pools:

#### 1. Source Chain (OnRamp) Processing

When `lockOrBurn` is called on the source pool, it returns:

```solidity
struct LockOrBurnOutV1 {
    bytes destTokenAddress;    // Destination token address
    bytes destPoolData;        // Pool-specific data for destination
}
```

The OnRamp processes this return data:

```solidity
// In OnRamp._lockOrBurnSingleToken()
Pool.LockOrBurnOutV1 memory poolReturnData = sourcePool.lockOrBurn(lockOrBurnIn);

// Store in message
Internal.EVM2AnyTokenTransfer memory tokenTransfer = Internal.EVM2AnyTokenTransfer({
    sourcePoolAddress: address(sourcePool),
    destTokenAddress: poolReturnData.destTokenAddress,
    extraData: poolReturnData.destPoolData,  // Passed to destination pool
    amount: amount,
    destExecData: ""  // Set by FeeQuoter
});
```

#### 2. Gas Limit Configuration

The FeeQuoter determines gas limits for token operations:

```solidity
// Per-token configuration
struct TokenTransferFeeConfig {
    uint32 destGasOverhead;     // Gas for releaseOrMint execution
    uint16 destBytesOverhead;   // Max bytes for destPoolData
    bool isEnabled;
}

// Processing in FeeQuoter.processPoolReturnData()
uint32 destGasAmount = tokenTransferFeeConfig.isEnabled
    ? tokenTransferFeeConfig.destGasOverhead
    : destChainConfig.defaultTokenDestGasOverhead;

// Encode gas amount for destination
destExecDataPerToken[i] = abi.encode(destGasAmount);
```

#### 3. Destination Chain (OffRamp) Processing

The OffRamp orchestrates a secure token release/mint with balance verification:

```solidity
// In OffRamp._releaseOrMintSingleToken()
// The destGasAmount is used for all three operations sequentially:

// 1. Pre-transfer balance check
(uint256 balancePre, uint256 gasLeft) = _getBalanceOfReceiver(
    receiver,
    localToken,
    sourceTokenAmount.destGasAmount
);

// 2. Execute releaseOrMint with remaining gas
(bool success, bytes memory returnData, uint256 gasUsedReleaseOrMint) =
    CallWithExactGas._callWithExactGasSafeReturnData(
        abi.encodeCall(
            IPoolV1.releaseOrMint,
            Pool.ReleaseOrMintInV1({
                originalSender: originalSender,
                receiver: receiver,
                sourceDenominatedAmount: sourceTokenAmount.amount,
                localToken: localToken,
                remoteChainSelector: sourceChainSelector,
                sourcePoolAddress: sourceTokenAmount.sourcePoolAddress,
                sourcePoolData: sourceTokenAmount.extraData,  // destPoolData from source
                offchainTokenData: offchainTokenData          // From execution DON
            })
        ),
        localPoolAddress,
        gasLeft,  // Remaining gas after balance check
        i_gasForCallExactCheck,
        Internal.MAX_RET_BYTES
    );

// 3. Post-transfer balance check with remaining gas
(uint256 balancePost, ) = _getBalanceOfReceiver(
    receiver,
    localToken,
    gasLeft - gasUsedReleaseOrMint
);

// Verify the balance increased by exactly the amount returned by the pool
if (balancePost < balancePre || balancePost - balancePre != localAmount) {
    revert ReleaseOrMintBalanceMismatch(localAmount, balancePre, balancePost);
}
```

This three-step process ensures:

- Token pools cannot lie about amounts transferred
- Malicious tokens cannot grief the system
- Gas usage is controlled and predictable

#### 4. Offchain Token Data

The execution DON can provide additional data for each token transfer:

```solidity
struct ExecutionReport {
    Internal.Any2EVMRampMessage[] messages;
    bytes[][] offchainTokenData;  // Per-message, per-token data
    // ...
}
```

This allows for:

- Attestations for wrapped tokens (e.g., USDC CCTP attestations)
- Additional validation data
- Protocol-specific parameters

#### 5. Gas Limit Overrides

Manual execution supports gas limit overrides:

```solidity
struct GasLimitOverride {
    uint256 receiverExecutionGasLimit;
    uint256[] tokenGasOverrides;  // Per-token gas overrides
}

// Override validation
if (tokenGasOverride != 0) {
    if (tokenGasOverride < existingGasLimit) revert InvalidManualExecutionGasLimit();
    destGasAmount = tokenGasOverride;
}
```

### Rate Limiting

All token pools implement configurable rate limiting:

```solidity
struct RateLimiter.Config {
    bool isEnabled;
    uint128 capacity;     // Maximum tokens (bucket size)
    uint128 rate;         // Refill rate per second
}

// Applied per remote chain for both inbound and outbound transfers
struct ChainUpdate {
    uint64 remoteChainSelector;
    RateLimiter.Config outboundRateLimiterConfig;
    RateLimiter.Config inboundRateLimiterConfig;
}
```

## Security & Risk Management

### 1. Risk Management Network (RMN)

- Independent validation layer for cross-chain messages
- Can "curse" (pause) the system in case of anomalies
- Provides additional signatures for high-value transfers

### 2. Access Control

- **Authorized Callers**: Critical functions restricted to specific contracts
- **Admin/Owner Roles**: Separate roles for configuration and ownership
- **Timelock Integration**: Major changes can go through timelocks

### 3. Message Validation

- **Nonce Ordering**: Strict sequential nonce enforcement
- **Merkle Proofs**: Cryptographic proof of message inclusion
- **Source Verification**: OnRamp address validation

### 4. Emergency Controls

- **Pause Mechanisms**: Can pause specific chains or entire system
- **Rate Limiting**: Configurable limits on token transfers
- **Manual Execution**: Admin can force-execute stuck messages

## Setup & Configuration

### 1. Deploying a New CCIP Chain Contract

```solidity
// 1. Deploy core infrastructure on both chains
Router sourceRouter = new Router(wrappedNative, rmnRemote);
Router destRouter = new Router(wrappedNative, rmnRemote);

// 2. Deploy NonceManagers
NonceManager sourceNonceManager = new NonceManager();
NonceManager destNonceManager = new NonceManager();

// 3. Deploy TokenAdminRegistry
TokenAdminRegistry tokenAdminRegistry = new TokenAdminRegistry();

// 4. Deploy OnRamp on source
OnRamp onRamp = new OnRamp(staticConfig, dynamicConfig, destChainConfigs);

// 5. Deploy OffRamp on destination
OffRamp offRamp = new OffRamp(staticConfig, dynamicConfig, sourceChainConfigs);

// 6. Configure routers
sourceRouter.applyRampUpdates(onRampUpdates, ...);
destRouter.applyRampUpdates(..., offRampUpdates);

// 7. Set up OCR3 for OffRamp
offRamp.setOCR3Configs(ocrConfigs);

// 8. Configure access control
nonceManager.applyAuthorizedCallerUpdates(onRamp/offRamp);
feeQuoter.applyAuthorizedCallerUpdates(offRamp);
```

### 2. Adding Token Support

```solidity
// 1. Deploy token pools on both chains
LockReleaseTokenPool sourcePool = new LockReleaseTokenPool(token, ...);
BurnMintTokenPool destPool = new BurnMintTokenPool(token, ...);

// 2. Configure TokenAdminRegistry
tokenAdminRegistry.proposeAdministrator(token, admin);
tokenAdminRegistry.acceptAdminRole(token);
tokenAdminRegistry.setPool(token, pool);

// 3. Configure cross-chain pool mapping
TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
chainUpdates[0] = TokenPool.ChainUpdate({
    remoteChainSelector: remoteChain,
    remotePoolAddresses: [remotePool],
    remoteTokenAddress: abi.encode(remoteToken),
    outboundRateLimiterConfig: rateLimitConfig,
    inboundRateLimiterConfig: rateLimitConfig
});
pool.applyChainUpdates(chainUpdates);
```

### 3. OCR3 Configuration

OffRamp uses two separate OCR3 instances:

```solidity
// Commit Plugin Configuration
MultiOCR3Base.OCRConfigArgs({
    ocrPluginType: uint8(Internal.OCRPluginType.Commit),
    configDigest: commitConfigDigest,
    F: F,  // Byzantine fault tolerance
    n: n,  // Total oracles
    isSignatureVerificationEnabled: true,
    signers: commitSigners,
    transmitters: commitTransmitters
})

// Execution Plugin Configuration
MultiOCR3Base.OCRConfigArgs({
    ocrPluginType: uint8(Internal.OCRPluginType.Execution),
    configDigest: execConfigDigest,
    F: F,
    n: n,
    isSignatureVerificationEnabled: true,
    signers: execSigners,
    transmitters: execTransmitters
})
```

## Best Practices

1. **Always verify chain selectors** before deployment
2. **Set appropriate rate limits** based on token risk profile
3. **Configure gas limits** appropriately for destination chain execution
4. **Implement proper error handling** in receiver contracts
5. **Use message IDs** for tracking and debugging
6. **Set up monitoring** for RMN cursing events

## Conclusion

CCIP provides a robust, secure, and flexible framework for cross-chain communication. The protocol's layered architecture with separate commit and execution phases, combined with multiple security mechanisms, ensures reliable message delivery while maintaining decentralization. Understanding the contract interactions and setup procedures is crucial for successfully integrating with CCIP.

For more details on the offchain components and DON architecture, refer to the [CCIP Protocol Documentation](https://github.com/smartcontractkit/chainlink-ccip/blob/main/docs/ccip_protocol.md).
