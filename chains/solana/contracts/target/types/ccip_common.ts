/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/ccip_common.json`.
 */
export type CcipCommon = {
  "address": "Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C",
  "metadata": {
    "name": "ccipCommon",
    "version": "0.1.0-dev",
    "spec": "0.1.0"
  },
  "instructions": [],
  "errors": [
    {
      "code": 10000,
      "name": "invalidSequenceInterval",
      "msg": "The given sequence interval is invalid"
    },
    {
      "code": 10001,
      "name": "invalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 10002,
      "name": "invalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 10003,
      "name": "invalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 10004,
      "name": "invalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 10005,
      "name": "invalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 10006,
      "name": "invalidInputsPoolSignerAccounts",
      "msg": "Invalid pool signer account"
    },
    {
      "code": 10007,
      "name": "invalidChainFamilySelector",
      "msg": "Invalid chain family selector"
    },
    {
      "code": 10008,
      "name": "invalidEncoding",
      "msg": "Invalid encoding"
    },
    {
      "code": 10009,
      "name": "invalidEvmAddress",
      "msg": "Invalid EVM address"
    },
    {
      "code": 10010,
      "name": "invalidSvmAddress",
      "msg": "Invalid SVM address"
    }
  ]
};
