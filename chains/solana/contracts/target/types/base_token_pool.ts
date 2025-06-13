/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/base_token_pool.json`.
 */
export type BaseTokenPool = {
  "address": "9vMcoeKAhpSZw1F8PQx8fhiQy1cmLyehomrZFPniPLc1",
  "metadata": {
    "name": "baseTokenPool",
    "version": "0.1.1-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [],
  "events": [
    {
      "name": "burned",
      "discriminator": [
        207,
        37,
        251,
        154,
        239,
        229,
        14,
        67
      ]
    },
    {
      "name": "configChanged",
      "discriminator": [
        147,
        25,
        86,
        98,
        98,
        77,
        78,
        192
      ]
    },
    {
      "name": "globalConfigUpdated",
      "discriminator": [
        232,
        238,
        158,
        123,
        210,
        172,
        159,
        46
      ]
    },
    {
      "name": "locked",
      "discriminator": [
        188,
        53,
        118,
        62,
        64,
        12,
        198,
        84
      ]
    },
    {
      "name": "mintAuthorityTransferred",
      "discriminator": [
        23,
        136,
        155,
        223,
        27,
        166,
        51,
        85
      ]
    },
    {
      "name": "minted",
      "discriminator": [
        174,
        131,
        21,
        57,
        88,
        117,
        114,
        121
      ]
    },
    {
      "name": "ownershipTransferRequested",
      "discriminator": [
        79,
        54,
        99,
        123,
        57,
        244,
        134,
        35
      ]
    },
    {
      "name": "ownershipTransferred",
      "discriminator": [
        172,
        61,
        205,
        183,
        250,
        50,
        38,
        98
      ]
    },
    {
      "name": "rateLimitConfigured",
      "discriminator": [
        249,
        210,
        194,
        93,
        236,
        75,
        175,
        59
      ]
    },
    {
      "name": "released",
      "discriminator": [
        232,
        229,
        255,
        136,
        101,
        189,
        15,
        220
      ]
    },
    {
      "name": "remoteChainConfigured",
      "discriminator": [
        231,
        252,
        78,
        228,
        152,
        49,
        233,
        226
      ]
    },
    {
      "name": "remoteChainRemoved",
      "discriminator": [
        4,
        212,
        235,
        138,
        165,
        232,
        75,
        32
      ]
    },
    {
      "name": "remotePoolsAppended",
      "discriminator": [
        248,
        177,
        249,
        167,
        14,
        247,
        25,
        223
      ]
    },
    {
      "name": "routerUpdated",
      "discriminator": [
        230,
        116,
        235,
        209,
        74,
        144,
        208,
        95
      ]
    },
    {
      "name": "tokensConsumed",
      "discriminator": [
        126,
        8,
        242,
        245,
        121,
        78,
        210,
        0
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "invalidInitPoolPermissions",
      "msg": "Pool authority does not match token mint owner"
    },
    {
      "code": 6001,
      "name": "invalidRmnRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 6002,
      "name": "unauthorized",
      "msg": "unauthorized"
    },
    {
      "code": 6003,
      "name": "invalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 6004,
      "name": "invalidVersion",
      "msg": "Invalid state version"
    },
    {
      "code": 6005,
      "name": "invalidPoolCaller",
      "msg": "Caller is not ramp on router"
    },
    {
      "code": 6006,
      "name": "invalidSender",
      "msg": "Sender not allowed"
    },
    {
      "code": 6007,
      "name": "invalidSourcePoolAddress",
      "msg": "Invalid source pool address"
    },
    {
      "code": 6008,
      "name": "invalidToken",
      "msg": "Invalid token"
    },
    {
      "code": 6009,
      "name": "invalidTokenAmountConversion",
      "msg": "Invalid token amount conversion"
    },
    {
      "code": 6010,
      "name": "allowlistKeyAlreadyExisted",
      "msg": "Key already existed in the allowlist"
    },
    {
      "code": 6011,
      "name": "allowlistKeyDidNotExist",
      "msg": "Key did not exist in the allowlist"
    },
    {
      "code": 6012,
      "name": "remotePoolAddressAlreadyExisted",
      "msg": "Remote pool address already exists"
    },
    {
      "code": 6013,
      "name": "nonemptyPoolAddressesInit",
      "msg": "Expected empty pool addresses during initialization"
    },
    {
      "code": 6014,
      "name": "rlBucketOverfilled",
      "msg": "RateLimit: bucket overfilled"
    },
    {
      "code": 6015,
      "name": "rlMaxCapacityExceeded",
      "msg": "RateLimit: max capacity exceeded"
    },
    {
      "code": 6016,
      "name": "rlRateLimitReached",
      "msg": "RateLimit: rate limit reached"
    },
    {
      "code": 6017,
      "name": "rlInvalidRateLimitRate",
      "msg": "RateLimit: invalid rate limit rate"
    },
    {
      "code": 6018,
      "name": "rlDisabledNonZeroRateLimit",
      "msg": "RateLimit: disabled non-zero rate limit"
    },
    {
      "code": 6019,
      "name": "liquidityNotAccepted",
      "msg": "Liquidity not accepted"
    }
  ],
  "types": [
    {
      "name": "burned",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sender",
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "configChanged",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "config",
            "type": {
              "defined": {
                "name": "rateLimitConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "globalConfigUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "selfServedAllowed",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "locked",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sender",
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "mintAuthorityTransferred",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "mint",
            "type": "pubkey"
          },
          {
            "name": "oldMintAuthority",
            "type": "pubkey"
          },
          {
            "name": "newMintAuthority",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "minted",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sender",
            "type": "pubkey"
          },
          {
            "name": "recipient",
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "ownershipTransferRequested",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "from",
            "type": "pubkey"
          },
          {
            "name": "to",
            "type": "pubkey"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "ownershipTransferred",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "from",
            "type": "pubkey"
          },
          {
            "name": "to",
            "type": "pubkey"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "rateLimitConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "enabled",
            "type": "bool"
          },
          {
            "name": "capacity",
            "type": "u64"
          },
          {
            "name": "rate",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "rateLimitConfigured",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "outboundRateLimit",
            "type": {
              "defined": {
                "name": "rateLimitConfig"
              }
            }
          },
          {
            "name": "inboundRateLimit",
            "type": {
              "defined": {
                "name": "rateLimitConfig"
              }
            }
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "released",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sender",
            "type": "pubkey"
          },
          {
            "name": "recipient",
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "remoteAddress",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "address",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "remoteChainConfigured",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "token",
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
            }
          },
          {
            "name": "previousToken",
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
            }
          },
          {
            "name": "poolAddresses",
            "type": {
              "vec": {
                "defined": {
                  "name": "remoteAddress"
                }
              }
            }
          },
          {
            "name": "previousPoolAddresses",
            "type": {
              "vec": {
                "defined": {
                  "name": "remoteAddress"
                }
              }
            }
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "remoteChainRemoved",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "remotePoolsAppended",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "poolAddresses",
            "type": {
              "vec": {
                "defined": {
                  "name": "remoteAddress"
                }
              }
            }
          },
          {
            "name": "previousPoolAddresses",
            "type": {
              "vec": {
                "defined": {
                  "name": "remoteAddress"
                }
              }
            }
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "routerUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "oldRouter",
            "type": "pubkey"
          },
          {
            "name": "newRouter",
            "type": "pubkey"
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "tokensConsumed",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokens",
            "type": "u64"
          }
        ]
      }
    }
  ]
};
