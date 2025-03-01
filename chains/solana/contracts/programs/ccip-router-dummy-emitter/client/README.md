# CCIP Router Dummy Emitter

This client emits dummy events from the `ccip-router` contract on the Solana blockchain.

## Usage

### Running the Client

You can run the client using `go run` or the built binary.

#### Using `go run`

```sh
go run main.go --contract=ccip-router --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```

#### Using Built Binary

```sh
go build -o dummy-emitter main.go
```

```sh
./dummy-emitter --contract=ccip-router --result=REVERT --keypair=/path/to/solana/keypair.json --num-tx=2
```

### Flags
- contract: The contract name to emit events from (e.g., ccip-router).
- result: The result type (OK or REVERT).
- keypair: The path to the Solana keypair file. Needs to be funded first: [Solana Faucet](https://faucet.solana.com/)
- num-tx: The number of transactions to send.


### See Results on Solana Explorer
Results are shown under Transaction History Tab on the Solana Explorer for the [contract](https://explorer.solana.com/address/7sDY5A5S5NZe1zcqEuZybW6ZxAna1NWUZxU4ypdn8UQU?cluster=devnet).