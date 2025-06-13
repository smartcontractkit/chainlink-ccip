/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/test_token_pool.json`.
 */
export type TestTokenPool = {
  "address": "JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ",
  "metadata": {
    "name": "testTokenPool",
    "version": "0.1.1-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
    {
      "name": "acceptOwnership",
      "discriminator": [
        172,
        23,
        43,
        13,
        238,
        213,
        85,
        150
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": []
    },
    {
      "name": "appendRemotePoolAddresses",
      "discriminator": [
        172,
        57,
        83,
        55,
        70,
        112,
        26,
        197
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "pubkey"
        },
        {
          "name": "addresses",
          "type": {
            "vec": {
              "defined": {
                "name": "remoteAddress"
              }
            }
          }
        }
      ]
    },
    {
      "name": "deleteChainConfig",
      "discriminator": [
        241,
        159,
        142,
        210,
        64,
        173,
        77,
        179
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "editChainRemoteConfig",
      "discriminator": [
        149,
        112,
        186,
        72,
        116,
        217,
        159,
        175
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "pubkey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": {
              "name": "remoteConfig"
            }
          }
        }
      ]
    },
    {
      "name": "initChainRemoteConfig",
      "discriminator": [
        21,
        150,
        133,
        36,
        2,
        116,
        199,
        129
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "pubkey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": {
              "name": "remoteConfig"
            }
          }
        }
      ]
    },
    {
      "name": "initialize",
      "discriminator": [
        175,
        175,
        109,
        31,
        13,
        152,
        155,
        237
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        },
        {
          "name": "program"
        },
        {
          "name": "programData"
        }
      ],
      "args": [
        {
          "name": "poolType",
          "type": {
            "defined": {
              "name": "poolType"
            }
          }
        },
        {
          "name": "router",
          "type": "pubkey"
        },
        {
          "name": "rmnRemote",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "lockOrBurnTokens",
      "discriminator": [
        114,
        161,
        94,
        29,
        147,
        25,
        232,
        191
      ],
      "accounts": [
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "tokenProgram"
        },
        {
          "name": "mint",
          "writable": true
        },
        {
          "name": "poolSigner"
        },
        {
          "name": "poolTokenAccount",
          "writable": true
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
        },
        {
          "name": "chainConfig",
          "writable": true
        }
      ],
      "args": [
        {
          "name": "lockOrBurn",
          "type": {
            "defined": {
              "name": "lockOrBurnInV1"
            }
          }
        }
      ],
      "returns": {
        "defined": {
          "name": "lockOrBurnOutV1"
        }
      }
    },
    {
      "name": "releaseOrMintTokens",
      "discriminator": [
        92,
        100,
        150,
        198,
        252,
        63,
        164,
        228
      ],
      "accounts": [
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "offrampProgram",
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA."
          ]
        },
        {
          "name": "allowedOfframp",
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp"
          ]
        },
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "tokenProgram"
        },
        {
          "name": "mint",
          "writable": true
        },
        {
          "name": "poolSigner"
        },
        {
          "name": "poolTokenAccount",
          "writable": true
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
        },
        {
          "name": "receiverTokenAccount",
          "writable": true
        }
      ],
      "args": [
        {
          "name": "releaseOrMint",
          "type": {
            "defined": {
              "name": "releaseOrMintInV1"
            }
          }
        }
      ],
      "returns": {
        "defined": {
          "name": "releaseOrMintOutV1"
        }
      }
    },
    {
      "name": "setChainRateLimit",
      "discriminator": [
        188,
        188,
        161,
        37,
        100,
        249,
        123,
        170
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "pubkey"
        },
        {
          "name": "inbound",
          "type": {
            "defined": {
              "name": "rateLimitConfig"
            }
          }
        },
        {
          "name": "outbound",
          "type": {
            "defined": {
              "name": "rateLimitConfig"
            }
          }
        }
      ]
    },
    {
      "name": "setRouter",
      "discriminator": [
        236,
        248,
        107,
        200,
        151,
        160,
        44,
        250
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "newRouter",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "transferOwnership",
      "discriminator": [
        65,
        177,
        215,
        73,
        53,
        45,
        99,
        47
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "proposedOwner",
          "type": "pubkey"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "chainConfig",
      "discriminator": [
        13,
        177,
        233,
        141,
        212,
        29,
        148,
        56
      ]
    },
    {
      "name": "state",
      "discriminator": [
        216,
        146,
        107,
        94,
        104,
        75,
        182,
        177
      ]
    }
  ],
  "types": [
    {
      "name": "baseChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remote",
            "type": {
              "defined": {
                "name": "remoteConfig"
              }
            }
          },
          {
            "name": "inboundRateLimit",
            "type": {
              "defined": {
                "name": "rateLimitTokenBucket"
              }
            }
          },
          {
            "name": "outboundRateLimit",
            "type": {
              "defined": {
                "name": "rateLimitTokenBucket"
              }
            }
          }
        ]
      }
    },
    {
      "name": "baseConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenProgram",
            "type": "pubkey"
          },
          {
            "name": "mint",
            "type": "pubkey"
          },
          {
            "name": "decimals",
            "type": "u8"
          },
          {
            "name": "poolSigner",
            "type": "pubkey"
          },
          {
            "name": "poolTokenAccount",
            "type": "pubkey"
          },
          {
            "name": "owner",
            "type": "pubkey"
          },
          {
            "name": "proposedOwner",
            "type": "pubkey"
          },
          {
            "name": "rateLimitAdmin",
            "type": "pubkey"
          },
          {
            "name": "routerOnrampAuthority",
            "type": "pubkey"
          },
          {
            "name": "router",
            "type": "pubkey"
          },
          {
            "name": "rebalancer",
            "type": "pubkey"
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
              "vec": "pubkey"
            }
          },
          {
            "name": "rmnRemote",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "chainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "base",
            "type": {
              "defined": {
                "name": "baseChain"
              }
            }
          }
        ]
      }
    },
    {
      "name": "lockOrBurnInV1",
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
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "localToken",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "lockOrBurnOutV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destTokenAddress",
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
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
      "name": "poolType",
      "repr": {
        "kind": "rust"
      },
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "lockAndRelease"
          },
          {
            "name": "burnAndMint"
          },
          {
            "name": "wrapped"
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
      "name": "rateLimitTokenBucket",
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
              "defined": {
                "name": "rateLimitConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "releaseOrMintInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalSender",
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
            }
          },
          {
            "name": "remoteChainSelector",
            "type": "u64"
          },
          {
            "name": "receiver",
            "type": "pubkey"
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
            "type": "pubkey"
          },
          {
            "name": "sourcePoolAddress",
            "docs": [
              "@dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the",
              "expected pool address for the given remoteChainSelector."
            ],
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
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
      "name": "releaseOrMintOutV1",
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
      "name": "remoteConfig",
      "type": {
        "kind": "struct",
        "fields": [
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
            "name": "tokenAddress",
            "type": {
              "defined": {
                "name": "remoteAddress"
              }
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
      "name": "state",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "poolType",
            "type": {
              "defined": {
                "name": "poolType"
              }
            }
          },
          {
            "name": "config",
            "type": {
              "defined": {
                "name": "baseConfig"
              }
            }
          }
        ]
      }
    }
  ]
};
