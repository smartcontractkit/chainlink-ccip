# Billing

Before diving into billing let's do a recap of some parts of important parts of the system.

Note: During the rest of the documentation we'll consider ChainA as source chain and chainB as destinationChain.


## On Chain

We have multiple contracts interactions that happen to send message from ChainA to ChainB. These contracts are deployed on each chain that need to support CCIP
1. Router: Non changing initial contract that users/dApps call to send a message
2. OnRamp: Forwards Router's send request
3. FeeQuoter: OnRamp calls to estimate the fees for sending a message <<< Important
4. OffRamp: Receives the messages (Offchain part plays the part here to send the message to the other chain's OffRamp)

So if ChainA is sending a message to ChainB the flow looks like

1. ChainA: Router -> OnRamp -> FeeQuoter
2. OffChain using OffChain Reporting 3 (OCR3) processes send events emitted from OnRamp in Commit Plugin send a transaction commit on ChainB's OffRamp
3. ChanB: OffRamp -> FeeQuoter

## OffChain

We have 2 plugins (Commit and Execute).

Commit is the one responsible for responsible for reporting gas prices, and token prices that don't come from keystone.

Execute is responsible for fee boosting [Add Link] during the actual execution of messages.

## Fee Structure

To send a message from ChainA to ChainB we need to account for multiple fees. For in detailed doc check [billing](https://docs.chain.link/ccip/billing)

1. Network/Premium fees.
2. ChainB Transaction fees (execution costs + data availability cost on the destination chain)

Of-course the details of the calculation will depend on the message being sent and whether it has tokens to send or not, data availability..etc.
For in details look on how the fees is calculated you can check FeeQuoter's `getValidatedFee` (TODO: Put link)

To be able to pay in one of the available fee tokens on ChainA we need to arrive at a quote in USD and convert the USD amount into corresponding fee token amount.

So the components we need to calculate the final price are:
1. ChainA fee token price, usually LINK and the native token of the chain. This is what the user pays in the end. TokenPriceProcessor will update them. [TODO: link]
2. ChainB Fee/Gas price. FeeChainProcessor will update them. [TODO: link]
3. ChainB native token price (to be able to calculate the fees denominated in USD


## Token Prices

During commit round. Token prices are updated by TokenPriceProcessor. 

For each round these are the steps:
1. Observation:  
   a. Fetches the token prices from USD feed for tokens we [configure](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L89-L91) during the plugin initiation.   
   b. Fetches current token prices stored in FeeQuoter
2. Outcome:  
Cross-check values from 1a and 1b. and posts the tokens that needs updating in the Outcome. The prices from the feed (1a) will be used when:  
   a. If the token price on FeeQuoter is not available.  
   b. If the token price on FeeQuoter is stale by checking against when it was last updated and the [TokenPriceBatchWriteFrequency](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L87).  
   c. If the token price on FeeQuoter is not within the [PriceDeviationThreshold](https://github.com/smartcontractkit/chainlink-ccip/blob/f03ff5183eb8323ba8e0a13dc58d1da13b755307/pluginconfig/commit.go#L41), this is per chain configuration.
   

### Fee Token


## Gas Prices

## Fee Boosting

Assigning inflight messages that were previously skipped due to being underpaid an increasing weight for execution as time passes

## Aggregate Rate Limiting

