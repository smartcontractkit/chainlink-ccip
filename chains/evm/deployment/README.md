# CCIP Deployments for EVM Chains

This module defines tooling for CCIP contracts on EVM-compatible chains. It provides a structured approach for deploying and configuring CCIP contracts through operations, sequences, and changesets.

## Core Components

- **Operations**: Produce a single-side effect (deploy contract, call function)
- **Sequences**: Ordered collections of operations that represent a complete workflow
- **Changesets**: Integration of sequence(s) with a deployment environment (MCMS, datastore, etc.)

Consumers can use the level of granularity they require.
- Want to execute MCMS proposals in chainlink-deployments (or a similar deployment environment)? Use a changeset.
- Want to complete a full operational story without integrating with a full-fledged deployment environment? Use a sequence.
- Want to make a single contract call? Use an operation. 

## Hierarchy

```
deployment/
├── utils/
│   ├── changesets/      # Utilities for building changesets
│   └── operations/      # Utilities for building operations
│       ├── call/
│       └── deployment/
├── v1_7_0/              # CCIP 1.7.0 operations, sequences, & changesets
│   ├── changesets/
│   ├── sequences/
│   └── operations/
├── v1_6_0/
└── ...
```

## Development Guide

***TODO***

## North Star

***TODO***
