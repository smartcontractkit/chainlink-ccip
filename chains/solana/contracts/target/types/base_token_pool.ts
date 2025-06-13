/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/base_token_pool.json`.
 */
export type BaseTokenPool = {
<<<<<<< HEAD
  "address": "9vMcoeKAhpSZw1F8PQx8fhiQy1cmLyehomrZFPniPLc1",
  "metadata": {
    "name": "baseTokenPool",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
=======
  "version": "0.1.1-dev",
  "name": "base_token_pool",
>>>>>>> main
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
<<<<<<< HEAD
    },
    {
      "name": "rateLimitConfigured",
=======
    }
  ],
  "events": [
    {
      "name": "GlobalConfigUpdated",
      "fields": [
        {
          "name": "selfServedAllowed",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "Burned",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Minted",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "recipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Locked",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Released",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "recipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteChainConfigured",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "token",
          "type": {
            "defined": "RemoteAddress"
          },
          "index": false
        },
        {
          "name": "previousToken",
          "type": {
            "defined": "RemoteAddress"
          },
          "index": false
        },
        {
          "name": "poolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "previousPoolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RateLimitConfigured",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "outboundRateLimit",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        },
        {
          "name": "inboundRateLimit",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemotePoolsAppended",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "poolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "previousPoolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteChainRemoved",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RouterUpdated",
      "fields": [
        {
          "name": "oldRouter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newRouter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferRequested",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MintAuthorityTransferred",
      "fields": [
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "oldMintAuthority",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newMintAuthority",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "TokensConsumed",
      "fields": [
        {
          "name": "tokens",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigChanged",
      "fields": [
        {
          "name": "config",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidInitPoolPermissions",
      "msg": "Pool authority does not match token mint owner"
    },
    {
      "code": 6001,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 6002,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    },
    {
      "code": 6003,
      "name": "InvalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 6004,
      "name": "InvalidVersion",
      "msg": "Invalid state version"
    },
    {
      "code": 6005,
      "name": "InvalidPoolCaller",
      "msg": "Caller is not ramp on router"
    },
    {
      "code": 6006,
      "name": "InvalidSender",
      "msg": "Sender not allowed"
    },
    {
      "code": 6007,
      "name": "InvalidSourcePoolAddress",
      "msg": "Invalid source pool address"
    },
    {
      "code": 6008,
      "name": "InvalidToken",
      "msg": "Invalid token"
    },
    {
      "code": 6009,
      "name": "InvalidTokenAmountConversion",
      "msg": "Invalid token amount conversion"
    },
    {
      "code": 6010,
      "name": "AllowlistKeyAlreadyExisted",
      "msg": "Key already existed in the allowlist"
    },
    {
      "code": 6011,
      "name": "AllowlistKeyDidNotExist",
      "msg": "Key did not exist in the allowlist"
    },
    {
      "code": 6012,
      "name": "RemotePoolAddressAlreadyExisted",
      "msg": "Remote pool address already exists"
    },
    {
      "code": 6013,
      "name": "NonemptyPoolAddressesInit",
      "msg": "Expected empty pool addresses during initialization"
    },
    {
      "code": 6014,
      "name": "RLBucketOverfilled",
      "msg": "RateLimit: bucket overfilled"
    },
    {
      "code": 6015,
      "name": "RLMaxCapacityExceeded",
      "msg": "RateLimit: max capacity exceeded"
    },
    {
      "code": 6016,
      "name": "RLRateLimitReached",
      "msg": "RateLimit: rate limit reached"
    },
    {
      "code": 6017,
      "name": "RLInvalidRateLimitRate",
      "msg": "RateLimit: invalid rate limit rate"
    },
    {
      "code": 6018,
      "name": "RLDisabledNonZeroRateLimit",
      "msg": "RateLimit: disabled non-zero rate limit"
    },
    {
      "code": 6019,
      "name": "LiquidityNotAccepted",
      "msg": "Liquidity not accepted"
    }
  ]
};

export const IDL: BaseTokenPool = {
  "version": "0.1.1-dev",
  "name": "base_token_pool",
  "instructions": [],
  "types": [
    {
      "name": "BaseConfig",
>>>>>>> main
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
<<<<<<< HEAD
=======
    },
    {
      "name": "RateLimitConfig",
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
    }
  ],
  "events": [
    {
      "name": "GlobalConfigUpdated",
      "fields": [
        {
          "name": "selfServedAllowed",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "Burned",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Minted",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "recipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Locked",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Released",
      "fields": [
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "recipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteChainConfigured",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "token",
          "type": {
            "defined": "RemoteAddress"
          },
          "index": false
        },
        {
          "name": "previousToken",
          "type": {
            "defined": "RemoteAddress"
          },
          "index": false
        },
        {
          "name": "poolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "previousPoolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RateLimitConfigured",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "outboundRateLimit",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        },
        {
          "name": "inboundRateLimit",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemotePoolsAppended",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "poolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "previousPoolAddresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          },
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteChainRemoved",
      "fields": [
        {
          "name": "chainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RouterUpdated",
      "fields": [
        {
          "name": "oldRouter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newRouter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferRequested",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MintAuthorityTransferred",
      "fields": [
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "oldMintAuthority",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newMintAuthority",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "TokensConsumed",
      "fields": [
        {
          "name": "tokens",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigChanged",
      "fields": [
        {
          "name": "config",
          "type": {
            "defined": "RateLimitConfig"
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidInitPoolPermissions",
      "msg": "Pool authority does not match token mint owner"
    },
    {
      "code": 6001,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 6002,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    },
    {
      "code": 6003,
      "name": "InvalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 6004,
      "name": "InvalidVersion",
      "msg": "Invalid state version"
    },
    {
      "code": 6005,
      "name": "InvalidPoolCaller",
      "msg": "Caller is not ramp on router"
    },
    {
      "code": 6006,
      "name": "InvalidSender",
      "msg": "Sender not allowed"
    },
    {
      "code": 6007,
      "name": "InvalidSourcePoolAddress",
      "msg": "Invalid source pool address"
    },
    {
      "code": 6008,
      "name": "InvalidToken",
      "msg": "Invalid token"
    },
    {
      "code": 6009,
      "name": "InvalidTokenAmountConversion",
      "msg": "Invalid token amount conversion"
    },
    {
      "code": 6010,
      "name": "AllowlistKeyAlreadyExisted",
      "msg": "Key already existed in the allowlist"
    },
    {
      "code": 6011,
      "name": "AllowlistKeyDidNotExist",
      "msg": "Key did not exist in the allowlist"
    },
    {
      "code": 6012,
      "name": "RemotePoolAddressAlreadyExisted",
      "msg": "Remote pool address already exists"
    },
    {
      "code": 6013,
      "name": "NonemptyPoolAddressesInit",
      "msg": "Expected empty pool addresses during initialization"
    },
    {
      "code": 6014,
      "name": "RLBucketOverfilled",
      "msg": "RateLimit: bucket overfilled"
    },
    {
      "code": 6015,
      "name": "RLMaxCapacityExceeded",
      "msg": "RateLimit: max capacity exceeded"
    },
    {
      "code": 6016,
      "name": "RLRateLimitReached",
      "msg": "RateLimit: rate limit reached"
    },
    {
      "code": 6017,
      "name": "RLInvalidRateLimitRate",
      "msg": "RateLimit: invalid rate limit rate"
    },
    {
      "code": 6018,
      "name": "RLDisabledNonZeroRateLimit",
      "msg": "RateLimit: disabled non-zero rate limit"
    },
    {
      "code": 6019,
      "name": "LiquidityNotAccepted",
      "msg": "Liquidity not accepted"
>>>>>>> main
    }
  ]
};
