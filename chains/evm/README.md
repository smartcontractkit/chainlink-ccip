# Chainlink CCIP Smart Contracts

> [!IMPORTANT] Since v1.6.0 of the Chainlink contracts, the contracts have been moved to a new repository:
> [chainlink-ccip](https://github.com/smartcontractkit/chainlink-ccip). 
> The EVM contracts exist in the `chains/evm` directory of the repository.

## Installation

#### Foundry (git)

> [!WARNING]
> When installing via git, the ref defaults to master when no tag is given.

When installing through git, it is recommended to use a specific version tag to avoid breaking changes.
The recommended `chainlink-evm` version for any `chainlink-ccip` version can be found in the `package.json` NPM
dependency of the `chainlink/contracts` package. 
The corresponding git tag will be `contracts-v<version>`.

```sh
forge install smartcontractkit/chainlink-evm@<version_tag>
forge install smartcontractkit/chainlink-ccip@<version_tag>
```

Add the following remapping to your `remappings.txt` file:
```
@chainlink/contracts/=lib/smartcontractkit/chainlink-evm/contracts/
@chainlink/contracts-ccip/contracts/=lib/smartcontractkit/chainlink-ccip/chains/evm/contracts/
```


#### NPM
```sh
# pnpm
pnpm add @chainlink/contracts-ccip
```

```sh
# npm
npm install @chainlink/contracts-ccip --save
```

Add the following remapping to your `remappings.txt` file:
```
@chainlink/contracts/=node_modules/@chainlink/contracts/
@chainlink/contracts-ccip/contracts/=node_modules/@chainlink/contracts-ccip/contracts
```
### Directory Structure

```sh
@chainlink/contracts-ccip
├── contracts # Solidity contracts
├── scripts # Compilation script
└── abi # ABI json
```

## Usage

The solidity smart contracts can be imported via `@chainlink/contracts-ccip/contracts`:

```solidity
import {CCIPReceiver} from  '@chainlink/contracts-ccip/contracts/applications/CCIPReceiver.sol';
```

#### Getting started with CCIP

To get started with CCIP, please refer to the [CCIP documentation](https://docs.chain.link/ccip).

The MockRouter contract is a good starting point when developing dapps that use CCIP.
It is a simplified same-chain entry and exit point for CCIP messages.
It lives in the `contracts/test/mocks/MockRouter.sol` file.

### Remapping

This repository uses [Solidity remappings](https://docs.soliditylang.org/en/v0.8.20/using-the-compiler.html#compiler-remapping) to resolve imports.
The remapping is defined in the `remappings.txt` file in the root of the repository.
Please see the Installation section above for the correct remappings based on your installation method.

If required, you can remap dependencies used within CCIP contracts, e.g. Openzeppelin contracts, 
by adding the following to your `remappings.txt` file:

```
@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/=node_modules/@openzeppelin/contracts/
@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/=node_modules/@openzeppelin/contracts/
```

This allows you to use a wide range of versions of Openzeppelin in your project without conflicts.

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
