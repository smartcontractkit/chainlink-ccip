# Context for invariants

## Legacy support

There are 5 chain families that already supported CCIP v1: EVM, Solana, Aptos, Sui and TON. All net-new implementations do not need to support any legacy v1 support like IPoolV1 (the legcay TokenPool interface) or legacy ExtraArgs. When IPoolV2 is mentioned, it means the CCIP v2 pool interface, even on net-new chains.

This is also true for the receiver interface. When a v2 receiver is mentioned, it means a receiver for CCIP v2 and is applicable to every CCIP v2 chain.


## Zero address sentinel

The zero address is used to signal "use the default CCV(s)". Even if a chain does not technically have a zero address, it must still support passing a value to signal this. 

## Instant finality chains

Some chains have instant finality and cannot re-org. In that case, the sending side of CCIP may always assume 0 as finality value and doesn't have to validate 0 with the pool, ccvs and executor. This means those contract also do not have to contian finality related config for sending messages. This is ONLY allowed on Canton and only on the sending side. 


# Missing features

Missing features should be flagged as FAIL unless context is provided on why they're not implemented. Partially implemented interfaces are a FAIL. A missing validation check is a FAIL. Further invariants relying on a missing feature are also a FAIL.
