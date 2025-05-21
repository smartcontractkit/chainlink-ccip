# Chainlink CCIP Smart Contracts

## Installation

```sh
# via pnpm
$ pnpm add @chainlink/contracts-ccip
# via npm
$ npm install @chainlink/contracts-ccip --save
```

### Directory Structure

```sh
@chainlink/contracts-ccip
├── contracts # Solidity contracts
└── abi # ABI json output
```

### Usage

The solidity smart contracts themselves can be imported via the `contracts` directory of `@chainlink/contracts-ccip`:

```solidity
import '@chainlink/contracts-ccip/contracts/applications/CCIPReceiver.sol';
```

### Remapping

This repository uses [Solidity remappings](https://docs.soliditylang.org/en/v0.8.20/using-the-compiler.html#compiler-remapping) to resolve imports.
The remapping is defined in the `remappings.txt` file in the root of the repository.

### Changesets

We use [changesets](https://github.com/changesets/changesets) to manage versioning the contracts.

Every PR that modifies any configuration or code, should most likely accompanied by a changeset file.

To install `changesets`:

1. Install `pnpm` if it is not already installed - [docs](https://pnpm.io/installation).
2. Run `pnpm install`.

Either after or before you create a commit, run the `pnpm changeset` command in the `chains/evm` directory to create an accompanying changeset entry which will reflect on the CHANGELOG for the next release.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## License

The CCIP repo is licensed under the [BUSL-1.1](./src/v0.8/ccip/LICENSE.md) license, however, there are a few exceptions

- `contracts/applications/*` is licensed under the [MIT](./src/v0.8/ccip/LICENSE-MIT.md) license
- `contracts/interfaces/*` is licensed under the [MIT](./src/v0.8/ccip/LICENSE-MIT.md) license
- `contracts/libraries/{Client.sol, Internal.sol}` is licensed under the [MIT](./src/v0.8/ccip/LICENSE-MIT.md) license
