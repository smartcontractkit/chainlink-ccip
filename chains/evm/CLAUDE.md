# CLAUDE.md

## Project Overview

This project focuses on CCIP (Cross-Chain Interoperability Protocol) development with agents.

## Development Guidelines

- [`context/ccip-solidity-agent.md`](context/ccip-solidity-agent.md) - Agent mindset, always read it, read it before anything else
- [`context/ccip-solidity-protocol.md`](context/ccip-solidity-protocol.md) - Learn how CCIP protocol works
- [`context/ccip-solidity-style.md`](context/ccip-solidity-style.md) - Solidity style guide for CCIP
- [`context/ccip-solidity-test-pattern.md`](context/ccip-solidity-test-pattern.md) - Learn how to write tests for CCIP contracts

Always index the above when touching CCIP Solidity, and THINK about a plan before implementing any changes.

## Commands

Chainlink CCIP EVM uses Foundry. You can reference the [Foundry Book](https://getfoundry.sh/forge/overview) when you need to read the manual.
All Foundry commands should be run with ccip profile by default

```bash
export FOUNDRY_PROFILE=ccip
```

Format

```bash
forge fmt
```

Build

```bash
forge build
```

Run tests

```bash
forge test
```

Run specific test

```bash
forge test --match-test testName
forge test --match-contract Benchmark
```
