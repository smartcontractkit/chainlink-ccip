# Chainlink CCIP Smart Contracts

> [!IMPORTANT] 
> Since v1.6.0 of the CCIP contracts, the contracts have been moved to a new repository:
> [chainlink-ccip](https://github.com/smartcontractkit/chainlink-ccip).
> The EVM contracts exist in the `chains/evm` directory of the repository.

## Installation

Chainlink-ccip relies on chainlink-evm, below are the instructions to install both.
To find the correct version of chainlink-evm to use for any given version of chainlink-ccip, 
please refer to the `package.json` file in `/chains/evm`.
It contains an NPM dependency on `@chainlink/contracts` with the correct version for the given version of chainlink-ccip.

NOTE: while other versions of chainlink-evm may work, we only test against the version specified in the `package.json` file.
Audits are also only done against the version specified in the `package.json` file. 
No guarantees are made for other versions.

#### Foundry (git)

> [!WARNING]
> When installing via git, the ref defaults to master when no tag is given.

When installing through git, it is recommended to use a specific version tag to avoid breaking changes.
The corresponding git tag will be `contracts-ccip-v<version>` for chainlink-ccip, 
and `contracts-v<version>` for chainlink-evm.

```sh
$ forge install smartcontractkit/chainlink-evm@contracts-v<version>
$ forge install smartcontractkit/chainlink-ccip@contracts-ccip-v<version>
```

Add the following remappings

```
@chainlink/contracts/=lib/smartcontractkit/chainlink-evm/contracts/
@chainlink/contracts-ccip/contracts/=lib/smartcontractkit/chainlink-ccip/chains/evm/contracts/
```
#### NPM

```sh
# pnpm
$ pnpm add @chainlink/contracts
$ pnpm add @chainlink/contracts-ccip
```
```sh
# npm
$ npm install @chainlink/contracts --save
$ npm install @chainlink/contracts-ccip --save
```

Add the following remappings

```
@chainlink/contracts/=node_modules/@chainlink/contracts/
@chainlink/contracts-ccip/contracts/=node_modules/@chainlink/contracts-ccip/contracts/
```

### Directory Structure

```sh
@chainlink/contracts-ccip
├── contracts # Solidity contracts
├── scripts # Compilation script
└── abi # ABI json output
```

### Usage

> [!WARNING]
> Contracts in `dev/` directories or with a typeAndVersion ending in `-dev` are under active development
> and are likely unaudited.
> Please refrain from using these in production applications.


The contracts can be imported via `@chainlink/contracts-ccip/contracts`:

```solidity
import {CCIPReceiver} from '@chainlink/contracts-ccip/contracts/applications/CCIPReceiver.sol';
```

#### Getting started with CCIP

To get started with CCIP, please refer to the [CCIP documentation](https://docs.chain.link/ccip).

The MockRouter contract is a good starting point when developing dapps that use CCIP.
It is a simplified same-chain entry and exit point for CCIP messages.
It lives in `contracts/test/mocks/MockRouter.sol`.


### Remapping

This repository uses [Solidity remappings](https://docs.soliditylang.org/en/v0.8.20/using-the-compiler.html#compiler-remapping) to resolve imports, 
which are defined in the `remappings.txt` file.

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
