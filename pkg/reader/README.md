# CCIP Reader Package

## Overview

The `pkg/reader` package provides a comprehensive set of interfaces and implementations for reading CCIP (Cross-Chain Interoperability Protocol) data from various blockchain networks. This package serves as the primary data access layer for CCIP operations, including configuration polling, price reading, USDC message retrieval, and RMN (Risk Management Network) operations.

## Core Interfaces

### CCIPReader
The main interface for reading CCIP-related data from blockchain networks. Provides unified access to:
- Chain configurations and routing information
- Contract addresses across different chains
- Transaction and message data
- Network topology and connectivity

**Key Methods:**
- `GetChainConfig()` - Retrieves chain-specific configuration
- `GetContractAddress()` - Gets deployed contract addresses
- `GetExpectedNextSequenceNumber()` - Tracks message sequencing
- `GetSourceToDestTxData()` - Fetches cross-chain transaction data

### PriceReader
Handles token price retrieval from price feeds and fee quoters across different chains.

**Key Features:**
- **Batch Price Fetching**: Efficiently retrieves USD prices for multiple tokens using batch requests
- **Price Normalization**: Normalizes prices to e18 precision for consistent calculations
- **Multiple Price Sources**: Supports both aggregator price feeds and FeeQuoter contracts
- **Cross-Chain Price Discovery**: Fetches prices from different chains as needed

**Key Methods:**
- `GetFeedPricesUSD(tokens)` - Gets USD prices from price aggregators, normalized to e18
- `GetFeeQuoterTokenUpdates(tokens, chain)` - Gets latest token prices from FeeQuoter contracts

**Price Calculation Examples:**
```
1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
```

### USDCMessageReader
Specialized interface for retrieving CCTP (Cross-Chain Transfer Protocol) MessageSent events for USDC transfers.

**Key Features:**
- Retrieves USDC-specific cross-chain message data
- Supports CCTP v1 MessageSent events
- Maps token transfer IDs to message data
- Handles USDC domain mappings across chains

**Key Methods:**
- `MessagesByTokenID(source, dest, tokens)` - Gets USDC message data by token transfer IDs

### RMNHome
Interface for interacting with the Risk Management Network (RMN) home contract.

**Key Features:**
- **Node Management**: Tracks RMN node information and configurations
- **Config Digest Validation**: Verifies and manages configuration digests
- **Source Chain Enablement**: Manages which chains require RMN verification
- **Fault Tolerance Configuration**: Handles F-value settings for observer consensus

**Key Methods:**
- `GetRMNNodesInfo(configDigest)` - Gets node information for a configuration
- `GetRMNEnabledSourceChains(configDigest)` - Gets chains requiring RMN verification
- `GetFObserve(configDigest)` - Gets fault tolerance values per chain

### ConfigPoller
Interface for polling and caching chain configuration data with background refresh capabilities.

**Key Methods:**
- `GetChainConfig(chainSel)` - Gets cached or fresh chain configuration
- `GetOfframpSourceChainConfigs(destChain, sourceChains)` - Gets source chain configs for off-ramp

## Architecture

### Data Flow Overview
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   CCIP Plugin   │────│   CCIPReader     │────│  Chain Networks │
│   (Commit/Exec) │    │   (Main Interface)│    │  (EVM/Solana)   │
└─────────────────┘    └─────────┬────────┘    └─────────────────┘
                                 │
                ┌────────────────┼────────────────┐
                │                │                │
        ┌───────▼────────┐ ┌─────▼─────┐ ┌──────▼──────┐
        │  ConfigPoller  │ │PriceReader│ │ USDCReader  │
        │   (Config      │ │ (Prices & │ │ (USDC CCTP  │
        │   Caching)     │ │  Fees)    │ │  Messages)  │
        └────────────────┘ └───────────┘ └─────────────┘
                │                │
        ┌───────▼────────┐ ┌─────▼─────┐
        │   RMNHome      │ │ Contract  │
        │ (Risk Mgmt)    │ │ Readers   │
        └────────────────┘ └───────────┘
```

### Key Components

#### ccipChainReader
The main implementation of `CCIPReader` that coordinates all data access:
- Manages contract readers and writers for multiple chains
- Integrates ConfigPoller for configuration management
- Provides unified interface for all CCIP data operations
- Handles address book management and codec operations

#### Price Reading Architecture
The `PriceReader` uses an efficient batch processing approach:

1. **Contract Grouping**: Groups tokens by their price aggregator contracts
2. **Batch Request Creation**: Creates batch requests for `getLatestRoundData` and `getDecimals`
3. **Parallel Execution**: Executes all contract calls in parallel
4. **Price Normalization**: Normalizes all prices to e18 precision
5. **Result Mapping**: Maps normalized prices back to requested tokens

## ConfigPollerV2 Architecture

ConfigPollerV2 is an optimized version of the original ConfigPoller that provides efficient caching and batch fetching of CCIP chain configurations.

### Key Design Principles

1. **Unified Batch Operations**: All data fetching goes through a single `chainAccessor.GetAllConfigsLegacy()` call
2. **Lazy Loading**: Configurations are only fetched when first requested or during background refresh
3. **Relationship Tracking**: Automatically tracks source-destination chain relationships upon initial request
4. **Cache-First Strategy**: Always check cache first, only fetch on cache miss or during background refresh
5. **Strict Destination Chain Validation**: Only supports fetching source chain configs for the configured destination chain

### Architecture Flow

```
Three Entry Points → Single Batch Refresh → Cache Updates

┌─────────────────────┐    ┌──────────────────────────────────┐    ┌──────────────────────────┐
│ GetChainConfig()    │    │ GetOfframpSourceChainConfigs()   │    │ Background Ticker        │
│                     │    │                                  │    │                          │
└──────────┬──────────┘    └──────────────┬───────────────────┘    └─────────┬────────────────┘
           │                              │                                  │
           │ 1. Check cache               │ 1. Validate destChain            │
           │                              │ 2. Filter out destChain          │
           │                              │ 3. Track source chains           │
           │                              │ 4. Check cache                   │
           ▼                              ▼                                  │
   ┌───────────────┐             ┌───────────────┐                           │
   │ Cache Hit?    │             │ Cache Hit?    │                           │
   │ Return data   │             │ Return data   │                           │
   └───┬───────────┘             └───┬───────────┘                           │
       │ Cache Miss                  │ Cache Miss                            │
       ▼                             ▼                                       │
   ┌──────────────────────────────────────────────┐                          │
   │        batchRefreshChainAndSourceConfigs()   │◄─────────────────────────┘
   │                                              │
   └─────────────────────┬────────────────────────┘
                         │
                         ▼
   ┌─────────────────────────────────────────────────────────────────────┐
   │                 Chain Accessor Batch Call                           │
   │   accessor.GetAllConfigsLegacy(ctx, destChain, knownSourceChains)   │
   └─────────────────────┬───────────────────────────────────────────────┘
                         │
                         ▼
   ┌─────────────────────────────────────────────────────────────────────┐
   │                      Update Caches                                  │
   │   • ChainConfigSnapshot → chainConfigData                           │
   │   • SourceChainConfigs → staticSourceChainConfigs                   │
   └─────────────────────────────────────────────────────────────────────┘
```

### Core Data Structures

#### knownSourceChains
```go
knownSourceChains map[cciptypes.ChainSelector]struct{}
```
- **Purpose**: Tracks which source chains send messages to the configured destination chain
- **Usage**: Enables efficient batch operations during background refresh
- **Population**: Automatically populated when `GetOfframpSourceChainConfigs()` is called
- **Storage**: Uses `struct{}` as value type for memory efficiency (set semantics)

#### chainCaches
```go
chainCaches map[cciptypes.ChainSelector]*chainCache
```
- **Purpose**: Stores cached configuration data per chain
- **Structure**: Each cache contains both chain config and source chain configs with separate locks
- **Lifecycle**: Created on-demand, updated via batch refresh

#### chainCache Structure
```go
type chainCache struct {
    // Chain config specific lock and data
    chainConfigMu      sync.RWMutex
    chainConfigData    cciptypes.ChainConfigSnapshot
    chainConfigRefresh time.Time

    // Source chain config specific lock and data  
    sourceChainMu            sync.RWMutex
    staticSourceChainConfigs map[cciptypes.ChainSelector]StaticSourceChainConfig
    sourceChainRefresh       time.Time
}
```

### Method Details

#### GetChainConfig(chainSel)
1. **Accessor Validation**: Verify chain accessor exists for the requested chain
2. **Cache Check**: Look for existing `ChainConfigSnapshot` in cache
3. **Cache Hit**: Return cached data immediately with debug logging showing cache age
4. **Cache Miss**: Trigger `batchRefreshChainAndSourceConfigs()` and return fresh data
5. **Lock Management**: Use RLock for reads, no locking during network I/O

#### GetOfframpSourceChainConfigs(destChain, sourceChains)
1. **Destination Chain Validation**: Ensure `destChain` matches configured `destChainSelector`
2. **Chain Filtering**: Remove destination chain from source chains list
3. **Track Relationships**: Add all valid source chains to `knownSourceChains` for future batch operations
4. **Cache Check**: Look for existing `StaticSourceChainConfig` entries
5. **Partial Cache Hit**: Return cached entries, identify missing ones
6. **Cache Miss**: Trigger `batchRefreshChainAndSourceConfigs()` for missing data
7. **Result Assembly**: Combine cached and newly fetched data

#### batchRefreshChainAndSourceConfigs(chainSel)
1. **Determine Context**: Check if fetching for destination chain vs source chain
2. **Get Known Source Chains**: For destination chain, retrieve all tracked source chains
3. **Network Call**: Call `accessor.GetAllConfigsLegacy(ctx, destChain, sourceChainSelectors)`
4. **Update Chain Config Cache**: Atomically update chain configuration with write lock
5. **Update Source Config Cache**: For destination chain, update source chain configs
6. **Error Handling**: Log unexpected source configs returned for non-destination chains
7. **Performance Logging**: Log latency for monitoring

## Key Benefits

### Simplified Architecture
- **Single Fetch Method**: Unified data fetching through consistent interfaces
- **Modular Design**: Clear separation of concerns between different data types
- **Consistent Error Handling**: All network operations follow the same pattern

### Improved Efficiency
- **Batch Operations**: Minimize network calls through intelligent batching
- **Intelligent Caching**: Cache-first strategies with background refresh
- **Relationship Discovery**: Automatic learning of chain relationships
- **Memory Efficient**: Optimized data structures and set semantics

### Better Error Handling
- **Graceful Degradation**: Can return cached data on fetch failures
- **Health Monitoring**: Tracks consecutive failures for operational visibility
- **Consistent Error Patterns**: Unified error handling across all components
- **Atomic Operations**: Thread-safe operations with proper locking

### Enhanced Safety
- **Strict Validation**: Input validation and constraint enforcement
- **Lock-Free I/O**: Network calls never hold locks to prevent blocking
- **Separate Granular Locks**: Independent locking for different data types
- **Deadlock Prevention**: Careful lock ordering and duration management
