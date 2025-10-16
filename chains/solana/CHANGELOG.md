# Changelog - CCIP Solana Contracts

This document describes the changes introduced in the different versions of the **Chainlink CCIP Solana programs**, located in [`chains/solana/contracts/programs`](https://github.com/smartcontractkit/chainlink-ccip/tree/main/chains/solana/contracts/programs).

---

## [Unreleased] (1.6.0)
<!-- ### Added
- (Placeholder for upcoming features) -->

### Changed

- [Token Pools] Allow setting rate limit with rate and capacity set to 0 [#1290](https://github.com/smartcontractkit/chainlink-ccip/pull/1290)

<!-- ### Fixed
- (Placeholder for bug fixes) -->

---

## [0.1.2]

- Commit [`b96a80a69ad2`](https://github.com/smartcontractkit/chainlink-ccip/commit/b96a80a69ad2)
- Git Tag: [solana-v0.1.2](https://github.com/smartcontractkit/chainlink-ccip/releases/tag/solana-v0.1.2)

### Added

- [Token Pools] Modify Rate Limit Admin by the owner of the State PDA of the Token Pool.

---

## [0.1.1]

### Core Contracts & CCTP Token Pool

    1. Commit [`7f8a0f403c3a`](https://github.com/smartcontractkit/chainlink-ccip/commit/7f8a0f403c3a)
    1. Git Tag: [solana-v0.1.1-cctp](https://github.com/smartcontractkit/chainlink-ccip/releases/tag/solana-v0.1.1-cctp)

#### Added

- **Offramp**:

    1. Buffer execution
    1. Dynamic calculation of accounts

- **Router**:

    1. Updated message routing to support the extended flow with new token pools.
    1. Support for tokens with extensions

- **CCTP Token Pool**:

    1. New **Token Pool for CCTP (Circle Cross-Chain Transfer Protocol)**.

### Lock and Release + Burn and Mint Token Pools

    1. Commit: [`ee587a6c0562`](https://github.com/smartcontractkit/chainlink-ccip/commit/ee587a6c0562)
    1. Git Tag: [solana-v0.1.1](https://github.com/smartcontractkit/chainlink-ccip/releases/tag/solana-v0.1.1)

#### Added

1. Added **multisig support** when minting tokens. Tokens that have a multisig as mint authority now are able to be mint in the token pool.
1. Added a method to modify mint authority from the Signer PDA to the token multisig with validations
1. Introduced **self-served onboarding** a toggle for the global admin to configure if token pool self served is enabled or not.

---

## [0.1.0]

- Commit [`be8d09930aaa`](https://github.com/smartcontractkit/chainlink-ccip/commit/be8d09930aaa)
- Git Tag: [solana-v0.1.0](https://github.com/smartcontractkit/chainlink-ccip/releases/tag/solana-v0.1.0)

### Initial Release

This version represents the **first implementation of CCIP on Solana**.

### Added

- **Programs included** with the basic CCIP functionality on Solana:
  1. **Router**: orchestrates cross-chain message routing, includes the OnRamp (messages from Solana to any).
  1. **OffRamp**: manages receiving messages in Solana (any to Solana) and processes token price updates.
  1. **Fee Quoter**: message validation and fee calculations
  1. **Token Pools**: base contracts for token custody and transfers.
