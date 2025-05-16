# Billing

Before diving into billing let's do a recap of some important parts of the system.

## On Chain

We have multiple contracts interactions that happen to send message from `sourceChain` to `destinationChain`. These contracts are deployed on each chain that need to support CCIP
1. Router: Immutable entry point contract that users/dapps call to get quotes and send messages
2. OnRamp: Forwards Router's send request
3. FeeQuoter: OnRamp calls to estimate the fees for sending a message <<< Important
4. OffRamp: Receives the messages (Offchain part plays the part here to send the message to the other chain's OffRamp)

So if `sourceChain` is sending a message to `destinationChain` the flow looks like

1. `sourceChain`: Router -> OnRamp -> FeeQuoter
2. OffChain using OCR3 processes send events emitted from OnRamp in Commit Plugin send a transaction commit on `destinationChain`'s OffRamp
3. `destinationChain`: OffRamp ---ocr3 commit-plugin--> FeeQuoter

## OffChain

We have 2 plugins (Commit and Execute).

Commit is the one responsible for reporting gas prices, and token prices that don't come from keystone.

## Fee Structure

To send a message from `sourceChain` to `destinationChain` we need to account for multiple fees. For more details [billing documentation](https://docs.chain.link/ccip/billing)

1. Network/Premium fees.
2. `destinationChain` Transaction fees (execution costs + data availability cost on the destination chain)

Of course the details of the calculation will depend on the message being sent and whether it has tokens to send or not, data availability..etc.
For in details look on how the fees is calculated you can check FeeQuoter's [`getValidatedFee`](https://github.com/smartcontractkit/chainlink-ccip/blob/0259b25b503036131b4c483531f07446078abefa/chains/evm/contracts/FeeQuoter.sol#L558)

To be able to pay in one of the available fee tokens on `sourceChain` we need to arrive at a quote in USD and convert the USD amount into corresponding fee token amount.

So the components we need to calculate the final price are:
1. `sourceChain` fee token price in USD, usually LINK and the native token of the chain. This is what the user pays in the end. TokenPriceProcessor will update them 
2. `destinationChain` Fee/Gas price. This comes in native token. FeeChainProcessor will update them.
3. `destinationChain` native token price in USD (to be able to calculate the fees denominated in USD


## Token Prices Processor

TokenPrices [processor](https://github.com/smartcontractkit/chainlink-ccip/blob/f151a0cb6f3838be4e8290c7f5695d58d065a18a/commit/tokenprice/processor.go#L20-L20) is part of the Commit Plugin. It is responsible for updating the token prices on the destination chain FeeQuoter. 
This is done to be able to calculate the fees in USD.

For each round these are the steps:
1. **Observation:**  [source code](https://github.com/smartcontractkit/chainlink-ccip/blob/f151a0cb6f3838be4e8290c7f5695d58d065a18a/commit/tokenprice/observation.go)
   a. Fetches the token prices from USD feed for tokens we [configure](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L89-L91) during the plugin initiation.   
   b. Fetches current token prices stored in **destination chain FeeQuoter** (the chain the current node is supposed to commit to).
2. **Outcome:**  [source code](https://github.com/smartcontractkit/chainlink-ccip/blob/f151a0cb6f3838be4e8290c7f5695d58d065a18a/commit/tokenprice/outcome.go)
Cross-check values from 1a and 1b. and posts the tokens that needs updating in the Outcome. The prices from the feed (1a) will be used when:  
   a. If the token price on FeeQuoter is not available.  
   b. If the token price on FeeQuoter is stale by checking against when it was last updated and the [TokenPriceBatchWriteFrequency](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L87).  
   c. If the token price on FeeQuoter is not within the [PriceDeviationThreshold](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L41), this is per chain configuration.
   

## Gas Prices / Chain fees Processor

ChainFeeProcessor is part of the Commit Plugin. It is responsible for updating the chain fees on the destination chain FeeQuoter.

For each round, these are the steps:

1. **Observation:** [source code](https://github.com/smartcontractkit/chainlink-ccip/blob/0259b25b503036131b4c483531f07446078abefa/commit/chainfee/observation.go)
   a. Fetches the current gas prices from RPCs (Using chain writer). Prices are in native chains' token.
   b. Fetch native token price from **source chains' FeeQuoters**. These are all the source chains that have a [DestChainConfig](https://github.com/smartcontractkit/chainlink/blob/12af1de88238e0e918177d6b5622070417f48adf/contracts/src/v0.8/ccip/onRamp/OnRamp.sol#L406-L414) with the current chain as the destination chain. (and that the current node can read from as not all nodes can read from all chains)
   c. Fetches current gas prices stored in **destination chain FeeQuoter** (the chain the current node is supposed to commit to).

2. **Outcome:** [source code](https://github.com/smartcontractkit/chainlink-ccip/blob/f151a0cb6f3838be4e8290c7f5695d58d065a18a/commit/chainfee/outcome.go)
   Cross-check values from 1a and 1b, and posts the gas prices that need updating in the Outcome. The prices from the feed (1a) will be used when:
   a. If the gas price on FeeQuoter is not available.
   b. If the gas price on FeeQuoter is stale by checking against when it was last updated and the [RemoteGasPriceBatchWriteFrequency](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L80).
   c. If the gas price on FeeQuoter is not within the [PriceDeviationThreshold](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L31-L32), this is per chain and per fee component (execution, data availability) config.

One more thing that is done is to calculate the gas price in USD using the native token price from 1b. This is done to be able to calculate the fees in USD. For details on the calculation and the representation onchain please check the [code](https://github.com/smartcontractkit/chainlink-ccip/blob/5c54ab8396e3409cefef84dfa29d27920fc0ca46/commit/chainfee/outcome.go#L35-L81) with the comments.

## Aggregate Rate Limiting
