export type BaseTokenPool = {
  "version": "0.1.0-dev",
  "name": "base_token_pool",
  "instructions": [],
  "types": [
    {
      "name": "BaseConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenProgram",
            "type": "publicKey"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "decimals",
            "type": "u8"
          },
          {
            "name": "poolSigner",
            "type": "publicKey"
          },
          {
            "name": "poolTokenAccount",
            "type": "publicKey"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "rateLimitAdmin",
            "type": "publicKey"
          },
          {
            "name": "routerOnrampAuthority",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "rebalancer",
            "type": "publicKey"
          },
          {
            "name": "canAcceptLiquidity",
            "type": "bool"
          },
          {
            "name": "listEnabled",
            "type": "bool"
          },
          {
            "name": "allowList",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "BaseChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remote",
            "type": {
              "defined": "RemoteConfig"
            }
          },
          {
            "name": "inboundRateLimit",
            "type": {
              "defined": "RateLimitTokenBucket"
            }
          },
          {
            "name": "outboundRateLimit",
            "type": {
              "defined": "RateLimitTokenBucket"
            }
          }
        ]
      }
    },
    {
      "name": "RemoteConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "poolAddresses",
            "type": {
              "vec": {
                "defined": "RemoteAddress"
              }
            }
          },
          {
            "name": "tokenAddress",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "decimals",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "RemoteAddress",
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
      "name": "LockOrBurnInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "receiver",
            "type": "bytes"
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "originalSender",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "localToken",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "LockOrBurnOutV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destTokenAddress",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "destPoolData",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReleaseOrMintInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalSender",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "sourcePoolAddress",
            "docs": [
              "@dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the",
              "expected pool address for the given remoteChainSelector."
            ],
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "sourcePoolData",
            "type": "bytes"
          },
          {
            "name": "offchainTokenData",
            "docs": [
              "@dev WARNING: offchainTokenData is untrusted data."
            ],
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReleaseOrMintOutV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationAmount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "RateLimitTokenBucket",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokens",
            "type": "u64"
          },
          {
            "name": "lastUpdated",
            "type": "u64"
          },
          {
            "name": "cfg",
            "type": {
              "defined": "RateLimitConfig"
            }
          }
        ]
      }
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
  "version": "0.1.0-dev",
  "name": "base_token_pool",
  "instructions": [],
  "types": [
    {
      "name": "BaseConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenProgram",
            "type": "publicKey"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "decimals",
            "type": "u8"
          },
          {
            "name": "poolSigner",
            "type": "publicKey"
          },
          {
            "name": "poolTokenAccount",
            "type": "publicKey"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "rateLimitAdmin",
            "type": "publicKey"
          },
          {
            "name": "routerOnrampAuthority",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "rebalancer",
            "type": "publicKey"
          },
          {
            "name": "canAcceptLiquidity",
            "type": "bool"
          },
          {
            "name": "listEnabled",
            "type": "bool"
          },
          {
            "name": "allowList",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "BaseChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remote",
            "type": {
              "defined": "RemoteConfig"
            }
          },
          {
            "name": "inboundRateLimit",
            "type": {
              "defined": "RateLimitTokenBucket"
            }
          },
          {
            "name": "outboundRateLimit",
            "type": {
              "defined": "RateLimitTokenBucket"
            }
          }
        ]
      }
    },
    {
      "name": "RemoteConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "poolAddresses",
            "type": {
              "vec": {
                "defined": "RemoteAddress"
              }
            }
          },
          {
            "name": "tokenAddress",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "decimals",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "RemoteAddress",
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
      "name": "LockOrBurnInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "receiver",
            "type": "bytes"
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "originalSender",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "localToken",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "LockOrBurnOutV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destTokenAddress",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "destPoolData",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReleaseOrMintInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalSender",
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "sourcePoolAddress",
            "docs": [
              "@dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the",
              "expected pool address for the given remoteChainSelector."
            ],
            "type": {
              "defined": "RemoteAddress"
            }
          },
          {
            "name": "sourcePoolData",
            "type": "bytes"
          },
          {
            "name": "offchainTokenData",
            "docs": [
              "@dev WARNING: offchainTokenData is untrusted data."
            ],
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReleaseOrMintOutV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationAmount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "RateLimitTokenBucket",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokens",
            "type": "u64"
          },
          {
            "name": "lastUpdated",
            "type": "u64"
          },
          {
            "name": "cfg",
            "type": {
              "defined": "RateLimitConfig"
            }
          }
        ]
      }
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
