# Plugin Components

* shared types and interfaces - [chainlink-ccip, ccipocr3](https://github.com/smartcontractkit/chainlink-ccip/tree/main/pkg/types/ccipocr3)
* [OCR Plugins](https://github.com/smartcontractkit/chainlink-ccip)
  * [commit](https://github.com/smartcontractkit/chainlink-ccip/tree/main/commit)
  * [execute](https://github.com/smartcontractkit/chainlink-ccip/tree/main/execute)
  * [CCIPReader](https://github.com/smartcontractkit/chainlink-ccip/blob/main/pkg/reader/ccip.go) - contract reader wrapper interface for core protocol data access.
  * [Home Chain Reader](https://github.com/smartcontractkit/chainlink-ccip/blob/main/pkg/reader/home_chain.go) - contract reader wrapper for home chain data access.
* core node integration ([CCIP Capability](https://github.com/smartcontractkit/chainlink/tree/develop/core/capabilities/ccip))
  * EVM
    * [providers (hashing, encoding, etc)](https://github.com/smartcontractkit/chainlink/tree/develop/core/capabilities/ccip/ccipevm)
    * [contract reader & writer configuration](https://github.com/smartcontractkit/chainlink/tree/develop/core/capabilities/ccip/configs/evm)
  * [Solana](https://github.com/smartcontractkit/chainlink/tree/develop/core/capabilities/ccip/ccipsolana)
* integration tests
    * [initial deploy test](https://github.com/smartcontractkit/chainlink/blob/develop/integration-tests/deployment/ccip/changeset/initial_deploy_test.go)
