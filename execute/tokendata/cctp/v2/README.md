# CCTPv2TokenDataObserver Design

## CCTPv2
Cirlce's CCTPv2 (Cross Chain Transfer Protocol v2) allows CCIP to transfer USDC cross-chain without fees or the use of
custom (liquid) token pools. This simplifies and reduces the cost of using CCIP to transfer USDC. 

On the source chain, USDC is burned via a call to `depositForBurn` on CCTPv2's `TokenMessengerV2` contract.
On the destination chain, USDC is minted via a call to `receiveMessage` on CCTPv2's `MessageTransmitterV2` contract.
The CCTPv2 smart contracts are detailed here:
https://developers.circle.com/cctp/evm-smart-contracts

`receiveMessage` takes `message` and `attestation` as args, which both must be fetched from the CCTPv2 HTTP API:
https://developers.circle.com/api-reference/cctp/all/get-messages-v-2

## CCTPv2TokenDataObserver
`CCTPv2TokenDataObserver` fetches and outputs attestations for USDC transfers that were made using Circle's CCTPv2 
protocol.

`CCTPv2TokenDataObserver.Observe()`:
- accepts a collection of CCIP Messages
- determines if any of these CCIP Messages contain USDC token transfers that were made using Circle's CCTPv2 protocol.
- fetches attestations for these CCTPv2 USDC transfers
- outputs `TokenData` that pairs attestations with the corresponding token transfer

Downstream, these attestations will be sent to the CCTPv2 on-chain contracts, which will mint USDC
on the destination chain, completing the token transfer. 

## High Level Overview
CCTPv2TokenDataObserver's `Observe` function maps:

```
exectypes.MessageObservations -> exectypes.TokenDataObservations
```
these types are:
```
type MessageObservations   map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message
type TokenDataObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData
```
Essentially, `Observe` takes a collection of `Message`s and maps each one to a `MessageTokenData`.
`Message` has the field:
```
TokenAmounts []RampTokenAmount
```
and `MessageTokenData` is type:
```
type MessageTokenData struct {
	TokenData []TokenData
}
```
Importantly, `Observe` must be structure-preserving. The output `TokenDataObservations` must have
the same structure as the input `MessageObservations`: both maps must have the same chain selector,
keys, each chain selector key must have the same `SeqNum`s, each `SeqNum` must map to the same size
list.

So the core of what `Observe` does is the transform: `RampTokenAmount -> TokenData` for each 
`RampTokenAmount` in the input MessageObservations. `TokenData` is a generic container, and can
be generated from an `tokendata.AttestationStatus`:
```
type AttestationStatus struct {
	// ID is the identifier of the message that the attestation was fetched for
	ID cciptypes.Bytes
	// MessageBody is the raw message content
	MessageBody cciptypes.Bytes
	// Attestation is the attestation data fetched from the API, encoded in bytes
	Attestation cciptypes.Bytes
	// Error is the error that occurred during fetching the attestation data
	Error error
}
```
This reduces the transform to `RampTokenAmount -> AttestationStatus`, or:
```
RampTokenAmount -> (ID, MessageBody, Attestation)
```
The `ID` is just the CCIP Message ID, which is already present. `MessageBody` and `Attestation`
need to be fetched from the CCTPv2 API:
```
https://iris-api-sandbox.circle.com/v2/messages/{sourceDomainId}?transactionHash={$txHash}
```
So `sourceDomainId` and `txHash` are required to call the CCTPv2 HTTP API. `txHash` is a field of CCIP `Message`:
`Message.Header.TxHeader`. There are two ways to get `sourceDomainId`: we can either define a static map from 
`ChainSelector -> sourceDomainId`, or parse `sourceDomainId` from `RampTokenAmount`. The static map approach is 
error-prone: we would need to remember to update it as CCTP supports more chains, thus we opt for the latter approach. 

### SourceTokenDataPayloadV2
For CCTPv2 enabled USDC transfers, `RampTokenAmount.ExtraData (Bytes)` will decode to `SourceTokenDataPayloadV2`:
```
type SourceTokenDataPayloadV2 struct {
	// SourceDomain is the Circle domain ID of the source chain
	SourceDomain uint32
	
	// DepositHash is the content-addressable hash that uniquely identifies a CCTPv2 transfer.
	// It's calculated as keccak256(abi.encode(sourceDomain, amount, destinationDomain,
	// mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold))
	DepositHash [32]byte
}
```
This gives us `sourceDomainId`, so we can now call the CCTPv2 HTTP API. Let's go through the logic from the top. Given:
```
messages map[cciptypes.SeqNum]cciptypes.Message
```
we can parse `SourceTokenDataPayloadV2`s:
```
// Map SeqNum to a map from index of Message.TokenAmounts to SourceTokenDataPayloadV2
// Not that not every RampTokenAmount in Message.TokenAmounts is a CCTPv2 USDC transfer
v2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2
```
The `sourceDomainId` must be the same across all Messages from the same source chain. The messages may have 
different `TxHash`s, but some might share `TxHash`s. For efficiency, we don't want to make duplicate requests to the
CCTPv2 API. This means we can't make one request per message: we need to extract all unique tx hashes and make API calls
for each unique tx hash. This gives the following helper functions:
```
getSourceDomainID(v2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2) uint32
getTxHashes(messages map[cciptypes.SeqNum]cciptypes.Message) map[TxHash][]SeqNum
getCCTPv2Messages(sourceDomainID uint32, txHash string) []CCTPv2Messages
```
Where `CCTPv2Message` contains both `Attestation` and `MessageBody`, which are needed to populate the 
corresponding `tokendata.AttestationStatus` fields. The main challenge is to assign a CCTPv2Message/attestation to each 
`SourceTokenDataPayloadV2`. We use `DepositHash` to do this.

### DepositHash
In an ideal world, on the source chain when we make a CCTPv2 USDC transfer by calling `depositForBurn` on CCTPv2's
`TokenMessengerV2` contract, `depositForBurn` would return a unique ID or nonce, and we would put this nonce in
`SourceTokenDataPayloadV2`, and then use `SourceTokenDataPayloadV2.nonce` to fetch the exact attestation for this
transfer. 

However, `depositForBurn` does not return a unique ID. Instead, on-chain we hash all the params of the call to
`depositForBurn` and add this `DepositHash` to `SourceTokenDataPayloadV2`. `DepositHash` is calculated as:
```
keccak256(abi.encode(sourceDomain, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold))
```
`CCTPv2Message` (fetched from the API) has the same fields, so a `SourceTokenDataPayloadV2` matches a `CCTPv2Message`
when each of their `DepositHash`s are the same and **they are part of the same tx (e.g. share the same TxHash)**. 

**IMPORTANT**: `DepositHash` may not be unique, e.g. there could be multiple CCTPv2 USDC transfers within the same
tx that have the exact same params. In this case of duplicate `DepositHash`s in the same tx, attestations for this
`DepositHash` are **fungible**. For example, if we have:
```
tokenPayloadA with DepositHashX
tokenPayloadB with DepositHashX
cctpV2Message1 with attestation1 and DepositHashX
cctpV2Message2 with attestation2 and DepositHashX
```
then `attestation1` can be assigned to **either** `tokenPayloadA` or `tokenPayloadB`, and `attestation2` assigned to the 
remaining `tokenPayload`. 

### Putting it all together
With this understanding we can write the high-level design of `Observe()`. For each chain:
1. Extract data from `messages`: v2TokenPayloads, sourceDomainID, txHashes
2. Fetch attestations from the CCTPv2 API and assign an attestation to each v2TokenPayload
3. Convert each attestation to `TokenData`

Step 2 is implemented by `assignAttestationsToV2TokenPayloads` which, for each txHash:
- fetch `CCTPv2Messages` from the API
- converts `CCTPv2Messages` to a map from `DepositHash` to `[]AttestationStatus`
  - Again note that `DepositHash` is not unique and may map to multiple `AttestationStatus`
  - `AttestationStatus` acts here as a container for `MessageBody` and `Attestation`
- iterate over each `v2TokenPayload`
  - look up `v2TokenPayload.DepositHash` in the `DepositHash` to `[]AttestationStatus` map
  - if the lookup returns a non-empty list, destructively pop() an `AttestationStatus`
    - **The popped `AttestationStatus` is now assigned to this specific `v2TokenPayload`**
    - The pop() being destructive ensures the `AttestationStatus` won't be assigned again
- At the end of this iteration, each `v2TokenPayload` has been assigned a `AttestationStatus` (for this specific TxHash)

Step 3 converts these `AttestationStatus` into `TokenData` and ensures structure is preserved.
