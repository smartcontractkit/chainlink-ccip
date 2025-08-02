# Go Bindings Module

This module provides Go bindings for multiple versions of the smart contracts.

## Overview

Each released version of the contract bindings is included in this repository. By importing the latest version of this module, you gain access to **all existing contract versions**, ensuring backward compatibility and version flexibility.

## Releasing a New Version

To release a new version (i.e., update `latest` to point to a fixed contract version):

1. **Generate or copy the new bindings** into the appropriate versioned directory.
2. **Update all `alias.go` files** to point to the new version directory. These alias files are used to route the `latest` import path to the correct versioned package.

This approach keeps the code modular and simplifies integration across different contract versions while maintaining a clean upgrade path.
