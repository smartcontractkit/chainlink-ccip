# Billing

Before diving into billing let's do a recap of some parts of important parts of the system.


## On Chain

We have multiple contracts interactions that happens to send message from ChainA to ChainB. These contracts are deployed on each chain that need to support CCIP
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

Commit is the one responsible for initial billing.

Execute is responsible for fee boosting [Add Link] during the actual execution of messages.

## Fee Structure

To send a message from ChainA to ChainB we need to account for multiple fees

1. ChainA transaction fees
2. ChainB Transaction fees (execution fees on the destination chain)

Of-course the details of the calculation will depend on the message being sent and whether it has tokens to send or not, data availability..etc.
For in details look on how the fees is calculated you can check FeeQuoter's `getValidatedFee` (TODO: Put link)

In the end the user pays in one of the available fee tokens on ChainA. 

To be able to do this we need to convert whatever the estimated fees the user will pay on ChainB 



## Token Prices

### Fee Token

## Gas Prices

## Fee Boosting

## Aggregate Rate Limiting

