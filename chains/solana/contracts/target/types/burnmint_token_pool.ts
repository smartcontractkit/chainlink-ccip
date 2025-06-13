/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/burnmint_token_pool.json`.
 */
export type BurnmintTokenPool = {
<<<<<<< HEAD
  "address": "41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB",
  "metadata": {
    "name": "burnmintTokenPool",
    "version": "0.1.0-dev",
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
=======
  "version": "0.1.1-dev",
  "name": "burnmint_token_pool",
  "instructions": [
    {
      "name": "initGlobalConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "updateGlobalConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "selfServedAllowed",
          "type": "bool"
        }
      ]
    },
    {
      "name": "initialize",
>>>>>>> main
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
<<<<<<< HEAD
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
      "name": "configureAllowList",
      "discriminator": [
        18,
        180,
        102,
        187,
        209,
        0,
        130,
        191
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
        }
      ],
      "args": [
        {
          "name": "add",
          "type": {
            "vec": "pubkey"
          }
        },
        {
          "name": "enabled",
          "type": "bool"
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
=======
          "name": "programData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
>>>>>>> main
        }
      ],
      "args": [
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
      "name": "initializeStateVersion",
      "discriminator": [
        54,
        186,
        181,
        26,
        2,
        198,
        200,
        158
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        }
      ],
      "args": [
        {
          "name": "mint",
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
          "name": "state"
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
          "name": "state"
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
      "name": "removeFromAllowList",
      "discriminator": [
        44,
        46,
        123,
        213,
        40,
        11,
        107,
        18
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
        }
      ],
      "args": [
        {
          "name": "remove",
          "type": {
            "vec": "pubkey"
          }
        }
      ]
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
    },
    {
      "name": "transferMintAuthorityToMultisig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "newMultisigMintAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "typeVersion",
      "docs": [
        "Returns the program type (name) and version.",
        "Used by offchain code to easily determine which program & version is being interacted with.",
        "",
        "# Arguments",
        "* `ctx` - The context"
      ],
      "discriminator": [
        129,
        251,
        8,
        243,
        122,
        229,
        252,
        164
      ],
      "accounts": [
        {
          "name": "clock"
        }
      ],
      "args": [],
      "returns": "string"
    }
  ],
  "accounts": [
    {
<<<<<<< HEAD
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
=======
      "name": "poolConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "selfServedAllowed",
            "type": "bool"
          }
        ]
      }
>>>>>>> main
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
<<<<<<< HEAD
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
=======
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidMultisig",
      "msg": "Invalid Multisig Mint"
    },
    {
      "code": 6001,
      "name": "MintAuthorityAlreadySet",
      "msg": "Mint Authority already set"
    },
    {
      "code": 6002,
      "name": "FixedMintToken",
      "msg": "Token with no Mint Authority"
    },
    {
      "code": 6003,
      "name": "InvalidToken2022Multisig",
      "msg": "Invalid Multisig Account Data for Token 2022"
    },
    {
      "code": 6004,
      "name": "InvalidSPLTokenMultisig",
      "msg": "Invalid Multisig Account Data for SPL Token"
    },
    {
      "code": 6005,
      "name": "PoolSignerNotInMultisig",
      "msg": "Token Pool Signer PDA must be signer of the Multisig"
    },
    {
      "code": 6006,
      "name": "MultisigMustHaveMoreThanOneSigner",
      "msg": "Multisig must have more than one signer"
    },
    {
      "code": 6007,
      "name": "InvalidMultisigOwner",
      "msg": "Multisig Owner must match Token Program ID"
    }
  ]
};

export const IDL: BurnmintTokenPool = {
  "version": "0.1.1-dev",
  "name": "burnmint_token_pool",
  "instructions": [
    {
      "name": "initGlobalConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "updateGlobalConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "selfServedAllowed",
          "type": "bool"
        }
      ]
    },
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "router",
          "type": "publicKey"
        },
        {
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "transferMintAuthorityToMultisig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "newMultisigMintAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "typeVersion",
      "docs": [
        "Returns the program type (name) and version.",
        "Used by offchain code to easily determine which program & version is being interacted with.",
        "",
        "# Arguments",
        "* `ctx` - The context"
      ],
      "accounts": [
        {
          "name": "clock",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [],
      "returns": "string"
    },
    {
      "name": "transferOwnership",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "proposedOwner",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": []
    },
    {
      "name": "setRouter",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "newRouter",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "initializeStateVersion",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "mint",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "initChainRemoteConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "mint",
          "type": "publicKey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": "RemoteConfig"
>>>>>>> main
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
      "name": "poolConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "selfServedAllowed",
            "type": "bool"
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
            "name": "version",
            "type": "u8"
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
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidMultisig",
      "msg": "Invalid Multisig Mint"
    },
    {
      "code": 6001,
      "name": "MintAuthorityAlreadySet",
      "msg": "Mint Authority already set"
    },
    {
      "code": 6002,
      "name": "FixedMintToken",
      "msg": "Token with no Mint Authority"
    },
    {
      "code": 6003,
      "name": "InvalidToken2022Multisig",
      "msg": "Invalid Multisig Account Data for Token 2022"
    },
    {
      "code": 6004,
      "name": "InvalidSPLTokenMultisig",
      "msg": "Invalid Multisig Account Data for SPL Token"
    },
    {
      "code": 6005,
      "name": "PoolSignerNotInMultisig",
      "msg": "Token Pool Signer PDA must be signer of the Multisig"
    },
    {
      "code": 6006,
      "name": "MultisigMustHaveMoreThanOneSigner",
      "msg": "Multisig must have more than one signer"
    },
    {
      "code": 6007,
      "name": "InvalidMultisigOwner",
      "msg": "Multisig Owner must match Token Program ID"
    }
  ]
};
