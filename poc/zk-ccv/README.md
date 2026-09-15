# ZK CCV proof of concept

This proof of concept sends one CCIP v2 message from Sepolia to Arbitrum Sepolia. The message is verified only by the
Succinct ZK CCV. No committee signs it. The design is in `docs/zk-ccv-design.md`.

## How it works

1. `SP1Helios` on Arbitrum Sepolia anchors finalized Sepolia blocks. Each anchor is a real SP1 proof from the Succinct
   Prover Network.
2. A user sends a message on Sepolia. The message lists only the ZK CCV resolver and asks for manual execution.
3. The `zkccv` program waits until Helios anchors a block at or above the message block. It then builds a witness
   from public RPC data: a header chain from the anchored block down to the message block, and a Merkle Patricia
   Trie proof of the receipt.
4. `zkccv` submits the message and the witness to the OffRamp. `SuccinctZKVerifier` checks the witness in Solidity
   and binds it to the message id.

## Code

| Path | Content |
|---|---|
| `chains/evm/contracts/ccvs/SuccinctZKVerifier.sol` | Onchain verifier |
| `chains/evm/contracts/interfaces/IExecutionRootOracle.sol` | Getters the verifier reads from the light client |
| `chains/evm/contracts/libraries/mpt/` | RLP and trie libraries vendored from Optimism |
| `chains/evm/contracts/test/ccvs/SuccinctZKVerifier/` | Foundry tests on a witness recorded from Sepolia |
| `chains/evm/contracts/poc/zkccv/` | Foundry scripts: deploy source, deploy destination, send |
| `poc/zk-ccv/` | Go witness builder and executor |

## Deployed addresses

| Chain | Contract | Address |
|---|---|---|
| Sepolia | SuccinctZKVerifier | `0x1d87F1eabCb41BB84275A3c227cDda99c597e644` |
| Sepolia | VersionedVerifierResolver | `0x4B5571Bb9e68d8E3ef91a8481ef2e3c44D34eEcA` |
| Arbitrum Sepolia | SP1Helios | `0x304245f2c2dd2e8907334ae98bdbbeb7243a31f6` |
| Arbitrum Sepolia | SuccinctZKVerifier | `0xbb337eCc2D68643b90A7566A1c8B4314caC40668` |
| Arbitrum Sepolia | VersionedVerifierResolver | `0x245F1fe8ADA7Ae1a45d4d9025D9d4bd1a028fd76` |
| Arbitrum Sepolia | MockReceiverV2 | `0x91a574286f9ca95279cfE22077e369637e1eFB11` |

The lane is the ccv `staging_testnet` v2 lane. Its Router, OnRamp and OffRamp are used unchanged.

## Prerequisites

- Foundry, Go 1.26, Rust and Cargo.
- RPC URLs for Sepolia and Arbitrum Sepolia, and a Sepolia beacon API URL.
- One EOA private key with ETH on both testnets. The same key must hold PROVE on the Succinct Prover Network.
  Deposit PROVE through the `deposit` function of the SuccinctVApp contract on Ethereum mainnet.

## Run the proof of concept

### 1. Test the contracts

```sh
cd chains/evm
FOUNDRY_PROFILE=ccip forge test --match-path 'contracts/test/ccvs/SuccinctZKVerifier/*'
```

### 2. Run the Helios operator

Clone `github.com/succinctlabs/sp1-helios` at the main branch. Build and run the operator in execution-header mode.
Skip the deploy step if you reuse the `SP1Helios` address above.

```sh
SP1_SKIP_PROGRAM_BUILD=true cargo build --release -p sp1-helios-script --bin operator --bin genesis

# Optional: deploy a new SP1Helios. Use the Groth16 gateway as the verifier.
SP1_SKIP_PROGRAM_BUILD=true ./target/release/genesis \
  --rpc-url $ARBITRUM_SEPOLIA_RPC --private-key $KEY \
  --sp1-verifier-address 0x397A5f7f3dBd538f23DE225B51f532c34448dA9B \
  --source-chain-id 11155111 --source-consensus-rpc $SEPOLIA_BEACON_API

NETWORK_PRIVATE_KEY=$KEY SP1_HELIOS_PROOF_MODE=groth16 ./target/release/operator \
  --rpc-url $ARBITRUM_SEPOLIA_RPC --private-key $KEY \
  --contract-address 0x304245f2c2dd2e8907334ae98bdbbeb7243a31f6 \
  --source-chain-id 11155111 --source-consensus-rpc $SEPOLIA_BEACON_API \
  --loop-delay-mins 2 --commit-execution-header
```

Each update costs about one PROVE. Stop the operator when you are done.

### 3. Deploy the CCV contracts

Skip this step if you reuse the addresses above.

```sh
cd chains/evm
FOUNDRY_PROFILE=ccip forge script contracts/poc/zkccv/DeploySource.s.sol:DeploySource \
  --rpc-url $SEPOLIA_RPC --private-key $KEY --broadcast
HELIOS=0x304245f2c2dd2e8907334ae98bdbbeb7243a31f6 FOUNDRY_PROFILE=ccip forge script contracts/poc/zkccv/DeployDest.s.sol:DeployDest \
  --rpc-url $ARBITRUM_SEPOLIA_RPC --private-key $KEY --broadcast
```

### 4. Send a message

```sh
cd chains/evm
ZK_RESOLVER=0x4B5571Bb9e68d8E3ef91a8481ef2e3c44D34eEcA RECEIVER=0x91a574286f9ca95279cfE22077e369637e1eFB11 \
  FOUNDRY_PROFILE=ccip forge script contracts/poc/zkccv/Send.s.sol:Send \
  --rpc-url $SEPOLIA_RPC --private-key $KEY --broadcast
```

Note the transaction hash printed by Foundry.

### 5. Verify and execute

```sh
cd poc/zk-ccv
go build -o zkccv .
POC_PK=$KEY ./zkccv run \
  -src-rpc $SEPOLIA_RPC -dst-rpc $ARBITRUM_SEPOLIA_RPC \
  -tx $SEND_TX \
  -helios 0x304245f2c2dd2e8907334ae98bdbbeb7243a31f6 \
  -offramp 0xd5Ad61cBbdb611d18Ebee6391049f8bF862e16a0 \
  -ccv 0x245F1fe8ADA7Ae1a45d4d9025D9d4bd1a028fd76
```

The program waits for Sepolia finality and the next Helios update. This takes about 15 to 25 minutes. It then prints
the execute transaction and the execution state. State 2 is SUCCESS.

## Record a test fixture

`fixtures/sepolia_11628012.json` is the witness behind the Foundry tests. To record a new one from a past send
transaction:

```sh
./zkccv fixture -src-rpc $SEPOLIA_RPC -tx $SEND_TX -anchor-offset 3 -out fixtures/name.json
```

Then regenerate `SuccinctZKVerifierFixture.sol` from the JSON.
