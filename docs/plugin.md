

This document covers all changes made to the CCIP Offchain Plugin to enable CCIPv2 on Solana. Most of these changes are applicable and relevant for any new chain family we support for CCIP.

## Overview

CCIP offchain code is split into two parts:

### Core CCIP Plugin

- Located in the `chainlink-ccip` repo
- Designed to be chain-family agnostic but initially tailored for EVM
- Changes were made to support non-EVM chains like Solana

### CCIP Capability

- Plugin inside the main `chainlink` repo
- Facilitates communication between the core CCIP plugin and the blockchain
- Chain-family specific
- Bulk of Solana support changes are here

---

## Chain Selectors

A global identification system is required for all chains (mainnet, testnet, devnet, etc.).

**Properties:**

- `ChainSelector (uint64)`: Globally unique chain ID
- `ChainName (string)`: Human-readable name
- `Family (string)`: Chain family (e.g., evm, solana, starknet)
- `ChainID (string)`: Unique within the chain family

**Solana-specific identifiers:**

- `Family`: `solana`
- `ChainID`: Base-58 encoded genesis hash
- `ChainName`: `solana-mainnet`, `solana-testnet`, `solana-devnet`

---

## Home Chain

- CCIPv2 uses a central chain called the *Home Chain* to store configs.
- Ethereum is currently used as the Home Chain.
- Home Chain dependencies remain EVM-specific by design.
- Solana-specific configs are also stored on the Home Chain.

---

## CCIP Job Delegate

- Handles job instantiation during Chainlink Node startup.
- Previously depended on `evmConfigs`.
- Now accepts a generic config interface.
- Made chain agnostic.

---

## Plugin Oracle Creator

- Detects new CCIP DONs and initializes instances on the Core Node.

**Requirements:**

- DON configuration
- Chain-specific configuration and modules

**Refactor Summary:**

- Previously located in `core.capabilities.ccip.oraclecreator`
- Now chain-family agnostic
- Uses a `plugin` struct per chain:

```go
type plugin struct {
   CommitPluginCodec   cciptypes.CommitPluginCodec
   ExecutePluginCodec  cciptypes.ExecutePluginCodec
   ExtraArgsCodec      cciptypes.ExtraDataCodec
   MessageHasher       func(lggr logger.Logger) cciptypes.MessageHasher
   TokenDataEncoder    cciptypes.TokenDataEncoder
   GasEstimateProvider cciptypes.EstimateProvider
   RMNCrypto           func(lggr logger.Logger) cciptypes.RMNCrypto
}
