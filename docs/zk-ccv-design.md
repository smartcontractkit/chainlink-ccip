# ZK CCV Design: Succinct SP1 Helios Verifier

Status: DRAFT for review.
Scope: one new onchain verifier contract and one new offchain verifier service.
Non-scope: no changes to OnRamp, OffRamp, Router, executor, indexer core, or aggregator.

This document defines how a CCIP v2 message can be verified with a zero-knowledge light client
instead of a signer committee. Implementation plans should be derived from this document.

## 1. Summary

CCIP v2 lets each message choose its Cross-Chain Verifiers. This design adds a ZK CCV.
The ZK CCV proves on the destination chain that a `CCIPMessageSent` event happened on the
source chain at finality. It does this without trusting any signer set.

The proof has two layers:

1. A batch layer. Succinct runs the SP1 Helios light client. It submits one SP1 proof per
   update to the `SP1Helios` contract on the destination chain. Each update stores the
   finalized execution block hash and receipts root for one source block.
2. A per-message layer. Our offchain verifier builds a witness from public source-chain
   data. The witness proves that the message's receipt is inside a receipts trie whose
   root chains up to a Helios-anchored block. Our onchain verifier checks this witness
   in Solidity. No ZK proof is created per message.

Anyone with a source-chain RPC node can rebuild the witness. The only trusted inputs are
the Ethereum sync committee and the SP1 proof system.

## 2. Trust model

| Layer | What is trusted | Who produces it |
|---|---|---|
| Finality anchor | `executionBlockHashes[N]` and `executionReceiptsRoots[N]` on destination | Succinct, via SP1 proof checked by `SP1Helios` |
| Header ancestry | Hash chain from an anchored block down to the message block | Anyone, from public RPC; checked in Solidity |
| Receipt inclusion | Merkle Patricia Trie proof of the receipt | Anyone, from public RPC; checked in Solidity |
| Event binding | Decoded log matches the `messageId` under execution | Checked in Solidity |
| Delivery | Message executes only if `verifyMessage` does not revert | Destination OffRamp |

Offchain witness bytes are candidate evidence only. The destination contract has the final say.

## 3. Components

```
 SOURCE CHAIN (Ethereum)                      DESTINATION CHAIN
┌──────────────────────────┐                 ┌──────────────────────────────────┐
│ Router ─> OnRamp         │                 │            OffRamp               │
│            │             │                 │               │ verifyMessage    │
│            │ forwardTo   │                 │               v                  │
│            v Verifier    │                 │  VersionedVerifierResolver       │
│  Resolver ─> ZKVerifier  │                 │               │ by versionTag    │
│  (source instance)       │                 │               v                  │
│            │             │                 │          ZKVerifier ──> SP1Helios│
│    emits CCIPMessageSent │                 │  (MPT + header checks)  (roots)  │
└────────────┼─────────────┘                 └──────────────────▲───────────────┘
             │                                                  │ update
             │ finalized events                                 │ (SP1 proof)
             v                                                  │
┌──────────────────────────┐    verifierResults   ┌─────────────┴───────────┐
│ ZK verifier service      │──> Postgres + REST   │ Succinct Helios operator│
│ (chainlink-ccv pipeline) │         │            │ (batch SP1 proving)     │
└──────────────────────────┘         v            └─────────────────────────┘
                              Indexer ─> Executor ─> OffRamp.execute
                              (both unchanged)
```

New components are the `ZKVerifier` contract and the ZK verifier service.
`SP1Helios` is Succinct's contract. All other components are unchanged.

## 4. SP1 Helios anchor

`SP1Helios` after sp1-helios PR #56 provides:

- `updateExecutionHeader(bytes proof, ExecutionHeaderProofOutputs po)`. It verifies an SP1
  proof of Ethereum sync-committee finality. It then stores
  `executionBlockHashes[po.executionBlockNumber]` and
  `executionReceiptsRoots[po.executionBlockNumber]`.
- Public mappings `executionBlockHashes(uint256)` and `executionReceiptsRoots(uint256)`.

Two properties of this contract shape the design:

1. Each update anchors exactly one execution block: the finalized head at that update.
2. Heads must be beacon checkpoint slots. Checkpoints occur every 32 slots.

So most message blocks are never directly anchored. The witness must bridge from an
anchored block back to the message block. It does this with a header ancestry chain.
Block header H at height M commits to `parentHash`. Walking parent hashes from an anchored
block hash reaches the message block header. That header contains the receipts root for
the message block. This closes the gap without any change to Succinct's contract.

Correction to the PRD: the PRD assumes Helios stores roots for every block, and that only
`executionReceiptsRoots[N]` is needed. The merged code does not do that. The ancestry
chain is the required addition. When the message block happens to be anchored, the chain
is empty and verification uses the stored receipts root directly.

## 5. Onchain: ZKVerifier contract

Location: `chains/evm/contracts/ccvs/SuccinctZKVerifier.sol` plus the RLP and Merkle Patricia
Trie libraries vendored from Optimism under `chains/evm/contracts/libraries/mpt/`. The verifier
reads the light client through `IExecutionRootOracle`, see section 9.

The contract extends `BaseVerifier` and implements `ICrossChainVerifierV1`. It is deployed
behind `VersionedVerifierResolver` on both chains. The resolver address is the stable CCV
address that users, pools, and lanes reference. Inbound routing uses the 4-byte version
tag at the start of `verifierResults`, as with all existing CCVs.

Per-source-chain config, set by owner:

| Field | Purpose |
|---|---|
| `rootOracle` | Address of the `SP1Helios` root oracle for this source chain |
| `onRamp` | Expected emitter address of `CCIPMessageSent` on the source chain |
| `maxHeaderChainLength` | Upper bound on ancestry headers, bounds gas and calldata |

### forwardToVerifier (source side)

Same pattern as `CommitteeVerifier.forwardToVerifier`:

1. Assert the destination chain is not cursed by RMN.
2. Assert the sender passes the allowlist when the allowlist is enabled.
3. Return `versionTag()` as the verifier blob.

### verifyMessage (destination side)

The function is `view`-safe: it reads state and reverts on failure, it writes nothing.

1. Assert the source chain is not cursed by RMN.
2. Check the 4-byte version tag prefix. Revert on mismatch.
3. Decode the witness envelope. See section 6.
4. Resolve the anchor. Read `executionBlockHashes[anchorBlockNumber]` from the configured
   `SP1Helios`. Revert if zero.
5. Walk the ancestry chain. Require `keccak256(headers[0]) == anchoredBlockHash`. For each
   next header, require `keccak256(headers[i+1]) == headers[i].parentHash`. The last
   header is the message block header. Extract its `receiptsRoot`. If the chain is empty,
   read `executionReceiptsRoots[anchorBlockNumber]` instead and require it non-zero.
6. Verify receipt inclusion. Key is `rlp(txIndex)`. Value is the EIP-2718 receipt
   envelope. Verify the Merkle Patricia Trie proof against `receiptsRoot`.
7. Decode the receipt. Require success status. Select the log at `logIndex`, an index
   into this receipt's log list.
8. Bind the event. Require `log.address == config.onRamp`. Require `log.topics[0]` is the
   `CCIPMessageSent` topic. Decode the event's `messageId`. Require it equals the
   `messageId` argument.

Step 8 is sufficient binding. The OffRamp recomputes `messageId` as the keccak256 hash of
the encoded message it is executing. Equality with the proven event therefore binds the
full message content, per the `ICrossChainVerifierV1` contract obligations and
`CCV_INVARIANTS.md`.

### getFee

Inherited from `BaseVerifier`. Config per destination: flat `feeUSDCents`,
`gasForVerification` from benchmarks, and `payloadSizeBytes` sized for the worst-case
witness. Witness size is dominated by the ancestry chain: roughly 650 bytes per header
plus 1 to 3 KB of trie nodes. With one Helios update per checkpoint the worst case is
about 32 headers, near 22 KB.

## 6. verifierResults encoding

Length-prefixed envelope, following the `LombardVerifier` forward-compatibility pattern:

```
[versionTag            4 bytes ]
[anchorBlockNumber     8 bytes ]  block anchored in SP1Helios
[headerCount           2 bytes ]  0 when the message block is the anchor block
headerCount times:
  [headerLen           2 bytes ]
  [header RLP          variable]  ordered anchor block first, message block last
[txIndex               2 bytes ]  transaction index in the message block
[logIndex              1 byte  ]  log index within the receipt
[proofNodeCount        1 byte  ]
proofNodeCount times:
  [nodeLen             2 bytes ]
  [trie node RLP       variable]
```

The receipt bytes are not carried separately. The final trie node contains the receipt as
the proven leaf value. The envelope is prover-agnostic: nothing in it is specific to SP1.

## 7. Offchain: ZK verifier service

Location: `chainlink-ccv`, new package `verifier/pkg/zk/` and binary `cmd/verifier/zk/`.
The service reuses the standard coordinator pipeline. Model it on the token verifier, not
the committee verifier. There is no aggregator and no quorum. One valid witness suffices.

Pipeline stages, all existing:

1. `SourceReaderService` discovers `CCIPMessageSent` events, filters by our resolver
   address, and holds tasks until the message's finality gate passes.
2. `TaskVerifier` calls our `vtypes.Verifier.VerifyMessages` implementation.
3. `StorageWriter` persists results through `storage.NewCCVWriter` into Postgres.

`VerifyMessages` per task:

1. Read the destination `SP1Helios` over RPC. Find the smallest anchored block number
   at or above the task's block number. If none exists yet, return a retriable error
   with a delay. The pipeline retries until Helios sync catches up.
2. Fetch all receipts of the message block from the source RPC. Build the receipts trie
   locally and extract the proof for the message's transaction index. Derive the
   transaction index and receipt-local log index from the task's transaction hash.
   Correction to the PRD: `eth_getProof` serves state proofs only. There is no receipt
   proof RPC. Building the trie from the block's receipts is the standard method.
3. Fetch headers from the anchor block down to the message block. RLP-encode each header
   and check the hash chain locally before publishing.
4. Encode the envelope from section 6.
5. Optional preflight: `eth_call` the destination `verifyMessage` with the envelope.
6. Call `commit.CreateVerifierNodeResult` with the envelope bytes and version tag, and
   return the result to the pipeline.

Result serving: expose the stored results through the token-verifier REST API shape.
Register the service in the indexer config as a `rest` reader with our resolver addresses
as issuer addresses. The indexer and executor then work unchanged.

Storage decision, diverging from the PRD: the PRD names a `gcp://` locator. No code
implements URI-scheme storage in either repo today. v1 uses the working path: Postgres
behind the REST API. A bucket-backed `CCVStorage` implementation can be added later
without changing the envelope or the contract. `getStorageLocations()` returns the REST
endpoint URL for third-party executors.

Hosted builder option: a Succinct-hosted witness builder API can replace steps 2 and 3
behind a config switch, in the way the CCTP verifier wraps its attestation service. Same
bytes, same onchain trust model. Not required for v1.

## 8. Multi-CCV composition

Nothing special is needed. A message lists the ZK CCV resolver in `extraArgs.ccvs`, or a
pool or lane requires it. The OffRamp enforces quorum across all required CCVs. The ZK
result is one entry in `verifierResults[]` next to committee signatures. The executor
orders results by CCV address and is agnostic to their content.

## 9. Prover-agnostic anchor

The contract reads the root oracle through a minimal interface:

```solidity
interface IExecutionRootOracle {
  function executionBlockHashes(uint256 blockNumber) external view returns (bytes32);
  function executionReceiptsRoots(uint256 blockNumber) external view returns (bytes32);
}
```

`SP1Helios` already satisfies this interface without changes. A future Risc Zero or
Boundless anchor satisfies it by implementing the same two getters. Such a deployment
gets its own version tag and config entry. The envelope and the verification algorithm
stay identical. This meets the adaptability goal in the PRD.

## 10. Direction constraint and rollout

Helios proves Ethereum consensus. Therefore v1 verifies messages whose source is
Ethereum. Lanes into Ethereum keep committee-only verification until a light client for
the other chain exists. Each destination chain needs one `SP1Helios` deployment, an
operator posting updates, and one `ZKVerifier` deployment behind a resolver.

## 11. Risks and open questions

- Helios update cadence bounds both latency and witness size. Sparse updates mean long
  ancestry chains and large calldata. This needs a per-chain cadence agreement with
  Succinct and a benchmark before mainnet.
- Sync-committee security is weaker than full consensus verification. Security must sign
  off on this trust level, or require committee CCV alongside ZK on sensitive lanes.
  Lido's stated preference is both together, which composes naturally.
- The guardian role on `SP1Helios` can rotate verification keys. Key management and
  upgrade policy for the Helios deployment must be agreed with Succinct.
- Trie and RLP library choice needs an audit plan. Candidates: Optimism's
  `Lib_RLPReader` and `MerkleTrie` lineage, or Succinct's own library if offered.
- `logIndex` and `txIndex` widths cap at 255 logs per receipt and 65535 transactions per
  block. Confirm these bounds hold for all target chains before locking the envelope.

## 12. Testing

- Foundry: fixture-driven tests with recorded mainnet and Sepolia blocks. Cases: valid
  witness, empty ancestry, long ancestry, wrong root, truncated proof, wrong log index,
  wrong emitter, messageId mismatch, version mismatch, unanchored block, cursed chain.
- Gas benchmarks for `verifyMessage` as a function of ancestry length.
- Go: witness builder round-trip against the Foundry fixtures, so both sides prove the
  same bytes.
- E2E: devenv lane with a mocked or real Helios operator, ZK-only message and
  committee-plus-ZK message.

## 13. Proof of concept

`poc/zk-ccv/README.md` documents a working run on the ccv staging testnet lane Sepolia to Arbitrum Sepolia. One
message was verified only by the ZK CCV and executed with state SUCCESS. The anchor came from real SP1 Groth16 proofs
posted to `SP1Helios` by the upstream operator in execution-header mode. The witness carried 22 headers and cost about
1.35M gas to verify on Arbitrum Sepolia.
