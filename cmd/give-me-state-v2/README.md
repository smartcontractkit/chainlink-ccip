# give-me-state-v2

A high-performance tool for reading on-chain state from Chainlink CCIP contracts across multiple EVM (and Solana/Aptos) chains in parallel. Generates a comprehensive JSON snapshot of all contract configurations and state.

## Quick Start

### 1. Install Dependencies and Run Example

```bash
cd cmd/give-me-state-v2
go mod download
go run .
```

### 2. Prepare Input Files

You need two files:

**Network config (YAML)** -- defines which chains to query and their RPC endpoints:

```yaml
networks:
  - type: mainnet
    chain_selector: 5009297550715157269   # ethereum-mainnet
    rpcs:
      - rpc_name: MyRPC
        http_url: https://ethereum-rpc.publicnode.com
```

See `example.yaml` for a working config with public RPCs for Ethereum, Base, and Optimism.

**Address refs (JSON)** -- lists the contracts to query:

```json
[
    {
        "address": "0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D",
        "chainSelector": 5009297550715157269,
        "labels": [],
        "qualifier": "0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D-Router",
        "type": "Router",
        "version": "1.2.0"
    }
]
```

See `example_address_refs.json` for a small set of contracts on Ethereum, Base, and Optimism.

### 3. Run With Input Files

```bash
# Using the example files (small run, public RPCs, ~3 chains / 12 contracts)
go run .

# Full mainnet run (requires private RPCs in your YAML)
go run . -network mainnet.yaml -addresses mainnet_address_refs.json -output state.json
```

## CLI Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-network` | `testnet.yaml` | Path to network config YAML |
| `-addresses` | `address_refs.json` | Path to address refs JSON |
| `-output` | stdout | Output file path (omit to print to stdout) |
| `-timeout` | `30m` | Overall timeout for all operations |
| `-workers` | `12` | Worker goroutines per RPC endpoint |
| `-format` | `false` | Format output to match legacy state.json structure |
| `-live` | `true` | Show live progress bar and RPC stats during run |
| `-nops` | `false` | Include node operator data from the Job Distributor (requires `.env`) |

## Node Operators / Job Distributor (Optional)

To include node operator data in the output, pass the `-nops` flag. This connects to the Job Distributor (JD) gRPC service and fetches:
- **Nodes** -- ID, name, CSA public key, connection status, labels, version
- **Chain Configs** -- per-node chain assignments with account/admin addresses, OCR2 key bundles, P2P keys

The data appears under a top-level `"nodeOperators"` key in the output JSON. If the JD connection fails, the tool logs a warning and continues without JD data.

### Setup

JD connection details are read from environment variables. The easiest way is to create a `.env` file in this directory (see `env.sample` for the template):

```bash
cp env.sample .env
# Edit .env with your actual credentials
```

When `-nops` is provided, the tool automatically loads `.env` from the current working directory. Variables that are already exported in your shell take precedence over values in the file.

| Environment Variable | Required | Description |
|----------------------|----------|-------------|
| `JD_GRPC_URL` | Yes | gRPC endpoint address (e.g. `jd.example.com:443`) |
| `JD_TLS` | No | Set to `true` to enable TLS (default: `false`) |
| `JD_COGNITO_CLIENT_ID` | No | AWS Cognito app client ID (omit for insecure/no-auth) |
| `JD_COGNITO_CLIENT_SECRET` | No | AWS Cognito app client secret |
| `JD_USERNAME` | No | Cognito username |
| `JD_PASSWORD` | No | Cognito password |
| `JD_AWS_REGION` | No | AWS region for Cognito (e.g. `us-west-2`) |

### Example

```bash
# Run with node operator data
go run . -network testnet.yaml -addresses address_refs.json -nops -output state.json
```

## How It Works

1. **Typed Orchestrators** -- EVM, Solana, and Aptos each have their own orchestrator that knows how to talk to that chain type. Each orchestrator handles caching, deduplication, and request routing.

2. **Multicall3 Batching (EVM)** -- On startup, the EVM orchestrator probes each chain for the [Multicall3](https://www.multicall3.com/) contract (`0xcA11bde05977b3631167028862bE2a173976CA11`). When available, individual `eth_call` requests are automatically batched into `tryAggregate` calls, reducing RPC round-trips by ~87%.

3. **Cache & Dedup** -- Identical calls (same target + calldata) are deduplicated. If two views request the same data, only one RPC call is made.

4. **Generic Engine** -- All HTTP requests flow through a shared engine with per-endpoint health scoring (Laplace-smoothed), automatic retry with exponential backoff, and rate-limit cooldown.

5. **Views** -- Each contract type + version has a registered view function that knows which on-chain methods to call and how to decode the results using Go ABI bindings.

## Output

The tool produces a JSON file with the state of every queried contract, organized by chain selector. At the end of a run, it prints statistics showing:

- **RPC Efficiency** -- logical calls vs actual HTTP calls, cache hit rate, multicall batch sizes
- **Throughput** -- calls per second
- **Per-chain RPC pressure** -- which chains generated the most traffic, and how much was saved by caching/batching

## Supported Contract Types

The tool includes views for CCIP contracts across multiple versions:

- **v1.6**: OnRamp, OffRamp, FeeQuoter, NonceManager, CCIPHome, RMNHome, RMNRemote
- **v1.5**: CommitStore, EVM2EVMOnRamp, EVM2EVMOffRamp, RMN
- **v1.2**: Router, PriceRegistry
- **v1.0**: ARMProxy, CapabilitiesRegistry, LinkToken
- **Token Pools**: v1.5.1, v1.6, v1.6.1 (LockRelease, BurnMint, USDC, etc.)
- **MCMS**: ManyChainMultiSig, RBACTimelock
