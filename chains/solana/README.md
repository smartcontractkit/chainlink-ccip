# CCIP Solana Onchain

- `contracts`: solana programs (rust) + tests (go) built on the anchor framework
- `gobindings`: auto-generated go bindings for contracts using `anchor-go`
- `scripts`: various scripts for generating artifacts

## Dependencies

- rust: https://www.rust-lang.org/tools/install
- go: https://go.dev/doc/install
- solana: https://docs.anza.xyz/cli/install
- anchor: https://www.anchor-lang.com/docs/installation

## Development

```bash
# install anchor-go if needed
go install github.com/gagliardetto/anchor-go@v0.2.3

# build contracts + IDL
anchor build

# go bindings need to be regenerated if contract changes were made
./scripts/anchor-go-gen.sh

# test contracts
go test ./... -v -count=1 -failfast
```
