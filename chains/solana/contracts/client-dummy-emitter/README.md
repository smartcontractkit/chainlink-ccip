# CCIP Dummy Events Emitter

This client emits dummy events from various CCIP contracts on the Solana blockchain.

## Usage

### Running the Client

You can run the client using `go run` or the built binary.

#### Using `go run`

```sh
go run main.go --contract=<contract-name> --result=<OK|REVERT> --keypair=/path/to/solana/keypair.json --num-tx=2
```

#### Using Built Binary

1. Build the client:
```sh
go build -o dummy-emitter main.go
```

2. Run the client:
```sh
./dummy-emitter --contract=<contract-name> --result=<OK|REVERT> --keypair=/path/to/solana/keypair.json --num-tx=2
```

### Flags
- `contract`: The contract name to emit events from (see Supported Contracts below)
- `result`: The result type (`OK` or `REVERT`)
- `keypair`: The path to the Solana keypair file. Needs to be funded first: [Solana Faucet](https://faucet.solana.com/)
- `num-tx`: The number of transactions to send

### Supported Contracts

#### CCIP Router
```sh
go run main.go --contract=ccip-router --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```
[View on Explorer](https://explorer.solana.com/address/7sDY5A5S5NZe1zcqEuZybW6ZxAna1NWUZxU4ypdn8UQU?cluster=devnet)

#### CCIP Offramp
```sh
go run main.go --contract=ccip-offramp --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```
[View on Explorer](https://explorer.solana.com/address/7h44xjUiJHH5wJCNpewaEDmgYLbUd7DDp6URuBKEenMT?cluster=devnet)

#### Fee Quoter
```sh
go run main.go --contract=fee-quoter --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```
[View on Explorer](https://explorer.solana.com/address/9ykZ4KUXJUtACe5Cg3UuTM14t5bxk1Amf6uaawGpGR5d?cluster=devnet)

#### MCM
```sh
go run main.go --contract=mcm --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```
[View on Explorer](https://explorer.solana.com/address/EqaAbT4NkoDU7WeKTHK9DrJEZ6xgSmZzoufpZQ7GPQE6?cluster=devnet)

#### Timelock
```sh
go run main.go --contract=timelock --result=OK --keypair=/path/to/solana/keypair.json --num-tx=2
```
[View on Explorer](https://explorer.solana.com/address/8hNnreBcZRQgcWnKaEkzsQwZA6B9ngjXFhSkToVu8V67?cluster=devnet)

### Testing Reverts

To test event emission followed by a revert, use `--result=REVERT` instead of `--result=OK`.:


### See Results
All emitted events can be viewed within each transaction under the Transaction History Tab on the Solana Explorer for each program (links above).

## Development

To add support for a new contract:

1. Add the program ID to the `programIDs` map in `main.go`
2. Create the corresponding dummy emitter program, build and deploy it to the Solana devnet.
```sh
anchor build
anchor deploy -p program-name --provider.cluster devnet
```