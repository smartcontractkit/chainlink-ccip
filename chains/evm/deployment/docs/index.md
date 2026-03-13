---
title: "EVM Implementation"
sidebar_label: "Overview"
sidebar_position: 1
---

# EVM Deployment Implementation

This documentation covers the EVM-specific implementation of the CCIP Deployment Tooling API. The EVM adapter is the most comprehensive reference implementation, supporting contract versions from v1.0.0 through v1.6.5.

For the shared interfaces this implementation fulfills, see [Interfaces Reference](../../../../deployment/docs/interfaces.md). For the shared types, see [Types Reference](../../../../deployment/docs/types.md).

---

## Package Layout

```
chains/evm/deployment/
├── utils/
│   ├── common.go                     # Address/version utilities
│   ├── datastore/
│   │   └── datastore.go              # EVM address format conversions
│   └── operations/
│       └── contract/
│           ├── read.go               # NewRead operation pattern
│           ├── write.go              # NewWrite operation pattern
│           ├── deploy.go             # NewDeploy operation pattern
│           └── function.go           # FunctionInput struct
├── v1_0_0/                           # Core infrastructure (MCMS, Router basics)
│   ├── adapters/
│   ├── operations/
│   └── sequences/
├── v1_2_0/                           # Router enhancements
│   ├── adapters/
│   └── operations/
├── v1_5_0/                           # Token support
│   ├── adapters/
│   ├── changesets/
│   ├── operations/
│   └── sequences/
├── v1_5_1/                           # Token pool patch
│   ├── operations/
│   └── sequences/
├── v1_6_0/                           # Complete implementation
│   ├── adapters/                     # Registration + specialized adapters
│   ├── operations/                   # All contract operations
│   ├── sequences/                    # All sequence compositions
│   └── testadapter/                  # Test utilities
├── v1_6_1/                           # Operations + sequences
├── v1_6_2/                           # Operations only
├── v1_6_3/                           # Operations only
└── v1_6_5/                           # Operations only
```

## Version Support

| Version | Scope | Contents |
|---------|-------|----------|
| **v1.0.0** | Core infrastructure | MCMS deployment, Router, WETH, LINK |
| **v1.2.0** | Router updates | Router operations for lane migration |
| **v1.5.0** | Token support | Token pools, token admin registry |
| **v1.5.1** | Token pool patch | LockRelease pool enhancements |
| **v1.6.0** | Complete implementation | Full adapter, all contracts, all sequences |
| **v1.6.1 -- v1.6.5** | Incremental updates | Additional operations and sequences |

The v1.6.0 version is the primary implementation that registers the `EVMAdapter` with all shared registries. Newer versions add operations and sequences that build upon the v1.6.0 foundation.

## Documentation

- [EVMAdapter Reference](adapter.md) -- adapter struct, interface implementations, registration
- [Operations Reference](operations.md) -- operation framework and all contract operations
- [Sequences Reference](sequences.md) -- all sequences and EVM-specific changesets
