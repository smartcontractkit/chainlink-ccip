export type CctpTokenPool = {
  "version": "0.1.0-dev",
  "name": "cctp_token_pool",
  "instructions": [
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
      "name": "typeVersion",
      "docs": [
        "Returns the program type (name) and version.",
        "Used by offchain code to easily determine which program & version is being interacted with.",
        "",
        "# Arguments",
        "* `ctx` - The context"
      ],
      "accounts": [],
      "args": [],
      "returns": "string"
    },
    {
      "name": "setRmnRemote",
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
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
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
          }
        }
      ]
    },
    {
      "name": "editChainRemoteConfig",
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
          }
        }
      ]
    },
    {
      "name": "editChainRemoteConfigCctp",
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
            "defined": "CctpChain"
          }
        }
      ]
    },
    {
      "name": "appendRemotePoolAddresses",
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
          "name": "addresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          }
        }
      ]
    },
    {
      "name": "setChainRateLimit",
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
          "name": "inbound",
          "type": {
            "defined": "RateLimitConfig"
          }
        },
        {
          "name": "outbound",
          "type": {
            "defined": "RateLimitConfig"
          }
        }
      ]
    },
    {
      "name": "deleteChainConfig",
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
        }
      ]
    },
    {
      "name": "configureAllowList",
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
        }
      ],
      "args": [
        {
          "name": "add",
          "type": {
            "vec": "publicKey"
          }
        },
        {
          "name": "enabled",
          "type": "bool"
        }
      ]
    },
    {
      "name": "removeFromAllowList",
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
        }
      ],
      "args": [
        {
          "name": "remove",
          "type": {
            "vec": "publicKey"
          }
        }
      ]
    },
    {
      "name": "releaseOrMintTokens",
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA."
          ]
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp"
          ]
        },
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "receiverTokenAccount",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "releaseOrMint",
          "type": {
            "defined": "ReleaseOrMintInV1"
          }
        }
      ],
      "returns": {
        "defined": "ReleaseOrMintOutV1"
      }
    },
    {
      "name": "lockOrBurnTokens",
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "hold a balance to pay for the rent of initializing the CCTP MessageSentEvent account"
          ]
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cctpAuthorityPda",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageTransmitterAccount",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpTokenMessengerAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpTokenMinterAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpLocalToken",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageTransmitter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is CCTP's MessageTransmitter program, which",
            "is invoked CCTP's TokenMessengerMinter by this program."
          ]
        },
        {
          "name": "cctpTokenMessengerMinter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is CCTP's TokenMessengerMinter program, which",
            "is invoked by this program."
          ]
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cctpEventAuthority",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpRemoteTokenMessengerKey",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageSentEvent",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is the account in which CCTP will store the event. It is not a PDA of CCTP,",
            "but CCTP will initialize it and become the owner for it."
          ]
        }
      ],
      "args": [
        {
          "name": "lockOrBurn",
          "type": {
            "defined": "LockOrBurnInV1"
          }
        }
      ],
      "returns": {
        "defined": "LockOrBurnOutV1"
      }
    },
    {
      "name": "reclaimEventAccount",
      "accounts": [
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "cctpMessageTransmitter",
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
        }
      ],
      "args": [
        {
          "name": "mint",
          "type": "publicKey"
        },
        {
          "name": "originalSender",
          "type": "publicKey"
        },
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "msgNonce",
          "type": "u64"
        },
        {
          "name": "attestation",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "deriveAccountsReleaseOrMintTokens",
      "accounts": [],
      "args": [
        {
          "name": "stage",
          "type": "string"
        },
        {
          "name": "releaseOrMint",
          "type": {
            "defined": "ReleaseOrMintInV1"
          }
        }
      ],
      "returns": {
        "defined": "DeriveAccountsResponse"
      }
    },
    {
      "name": "deriveAccountsLockOrBurnTokens",
      "accounts": [],
      "args": [
        {
          "name": "stage",
          "type": "string"
        },
        {
          "name": "lockOrBurn",
          "type": {
            "defined": "LockOrBurnInV1"
          }
        }
      ],
      "returns": {
        "defined": "DeriveAccountsResponse"
      }
    }
  ],
  "accounts": [
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
              "defined": "BaseConfig"
            }
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
              "defined": "BaseChain"
            }
          },
          {
            "name": "cctp",
            "type": {
              "defined": "CctpChain"
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CctpMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "data",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "MessageAndAttestation",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "message",
            "type": {
              "defined": "CctpMessage"
            }
          },
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "DepositForBurnWithCallerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "mintRecipient",
            "type": "publicKey"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "CctpChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domainId",
            "type": "u32"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "OnrampDeriveStage",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "RetrieveChainConfig"
          },
          {
            "name": "BuildDynamicAccounts"
          }
        ]
      }
    },
    {
      "name": "OfframpDeriveStage",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "RetrieveChainConfig"
          },
          {
            "name": "BuildDynamicAccounts"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "RemoteChainCctpConfigChanged",
      "fields": [
        {
          "name": "config",
          "type": {
            "defined": "CctpChain"
          },
          "index": false
        }
      ]
    },
    {
      "name": "CcipCctpMessageSentEvent",
      "fields": [
        {
          "name": "originalSender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "msgTotalNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "eventAddress",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "cctpNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageSentBytes",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 12000,
      "name": "InvalidTokenData",
      "msg": "Invalid token data"
    },
    {
      "code": 12001,
      "name": "InvalidReceiver",
      "msg": "Invalid receiver"
    },
    {
      "code": 12002,
      "name": "InvalidSourceDomain",
      "msg": "Invalid source domain"
    },
    {
      "code": 12003,
      "name": "InvalidDestDomain",
      "msg": "Invalid destination domain"
    },
    {
      "code": 12004,
      "name": "InvalidNonce",
      "msg": "Invalid nonce"
    },
    {
      "code": 12005,
      "name": "MalformedCctpMessage",
      "msg": "CCTP message is malformed or too short"
    },
    {
      "code": 12006,
      "name": "InvalidTokenMessengerMinter",
      "msg": "Invalid Token Messenger Minter"
    },
    {
      "code": 12007,
      "name": "InvalidMessageTransmitter",
      "msg": "Invalid Message Transmitter"
    },
    {
      "code": 12008,
      "name": "InvalidMessageSentEventAccount",
      "msg": "Invalid Message Sent Event Account"
    },
    {
      "code": 12009,
      "name": "InvalidTokenPoolExtraData",
      "msg": "Invalid Token Pool Extra Data"
    },
    {
      "code": 12010,
      "name": "FailedCctpCpi",
      "msg": "Failed CCTP CPI"
    }
  ]
};

export const IDL: CctpTokenPool = {
  "version": "0.1.0-dev",
  "name": "cctp_token_pool",
  "instructions": [
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
      "name": "typeVersion",
      "docs": [
        "Returns the program type (name) and version.",
        "Used by offchain code to easily determine which program & version is being interacted with.",
        "",
        "# Arguments",
        "* `ctx` - The context"
      ],
      "accounts": [],
      "args": [],
      "returns": "string"
    },
    {
      "name": "setRmnRemote",
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
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
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
          }
        }
      ]
    },
    {
      "name": "editChainRemoteConfig",
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
          }
        }
      ]
    },
    {
      "name": "editChainRemoteConfigCctp",
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
            "defined": "CctpChain"
          }
        }
      ]
    },
    {
      "name": "appendRemotePoolAddresses",
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
          "name": "addresses",
          "type": {
            "vec": {
              "defined": "RemoteAddress"
            }
          }
        }
      ]
    },
    {
      "name": "setChainRateLimit",
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
          "name": "inbound",
          "type": {
            "defined": "RateLimitConfig"
          }
        },
        {
          "name": "outbound",
          "type": {
            "defined": "RateLimitConfig"
          }
        }
      ]
    },
    {
      "name": "deleteChainConfig",
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
        }
      ]
    },
    {
      "name": "configureAllowList",
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
        }
      ],
      "args": [
        {
          "name": "add",
          "type": {
            "vec": "publicKey"
          }
        },
        {
          "name": "enabled",
          "type": "bool"
        }
      ]
    },
    {
      "name": "removeFromAllowList",
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
        }
      ],
      "args": [
        {
          "name": "remove",
          "type": {
            "vec": "publicKey"
          }
        }
      ]
    },
    {
      "name": "releaseOrMintTokens",
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA."
          ]
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp"
          ]
        },
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "receiverTokenAccount",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "releaseOrMint",
          "type": {
            "defined": "ReleaseOrMintInV1"
          }
        }
      ],
      "returns": {
        "defined": "ReleaseOrMintOutV1"
      }
    },
    {
      "name": "lockOrBurnTokens",
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "hold a balance to pay for the rent of initializing the CCTP MessageSentEvent account"
          ]
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cctpAuthorityPda",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageTransmitterAccount",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpTokenMessengerAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpTokenMinterAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpLocalToken",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageTransmitter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is CCTP's MessageTransmitter program, which",
            "is invoked CCTP's TokenMessengerMinter by this program."
          ]
        },
        {
          "name": "cctpTokenMessengerMinter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is CCTP's TokenMessengerMinter program, which",
            "is invoked by this program."
          ]
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cctpEventAuthority",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpRemoteTokenMessengerKey",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK this is not read by the pool, just forwarded to CCTP"
          ]
        },
        {
          "name": "cctpMessageSentEvent",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK this is the account in which CCTP will store the event. It is not a PDA of CCTP,",
            "but CCTP will initialize it and become the owner for it."
          ]
        }
      ],
      "args": [
        {
          "name": "lockOrBurn",
          "type": {
            "defined": "LockOrBurnInV1"
          }
        }
      ],
      "returns": {
        "defined": "LockOrBurnOutV1"
      }
    },
    {
      "name": "reclaimEventAccount",
      "accounts": [
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "cctpMessageTransmitter",
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
        }
      ],
      "args": [
        {
          "name": "mint",
          "type": "publicKey"
        },
        {
          "name": "originalSender",
          "type": "publicKey"
        },
        {
          "name": "remoteChainSelector",
          "type": "u64"
        },
        {
          "name": "msgNonce",
          "type": "u64"
        },
        {
          "name": "attestation",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "deriveAccountsReleaseOrMintTokens",
      "accounts": [],
      "args": [
        {
          "name": "stage",
          "type": "string"
        },
        {
          "name": "releaseOrMint",
          "type": {
            "defined": "ReleaseOrMintInV1"
          }
        }
      ],
      "returns": {
        "defined": "DeriveAccountsResponse"
      }
    },
    {
      "name": "deriveAccountsLockOrBurnTokens",
      "accounts": [],
      "args": [
        {
          "name": "stage",
          "type": "string"
        },
        {
          "name": "lockOrBurn",
          "type": {
            "defined": "LockOrBurnInV1"
          }
        }
      ],
      "returns": {
        "defined": "DeriveAccountsResponse"
      }
    }
  ],
  "accounts": [
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
              "defined": "BaseConfig"
            }
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
              "defined": "BaseChain"
            }
          },
          {
            "name": "cctp",
            "type": {
              "defined": "CctpChain"
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CctpMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "data",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "MessageAndAttestation",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "message",
            "type": {
              "defined": "CctpMessage"
            }
          },
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "DepositForBurnWithCallerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "mintRecipient",
            "type": "publicKey"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "CctpChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domainId",
            "type": "u32"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "OnrampDeriveStage",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "RetrieveChainConfig"
          },
          {
            "name": "BuildDynamicAccounts"
          }
        ]
      }
    },
    {
      "name": "OfframpDeriveStage",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "RetrieveChainConfig"
          },
          {
            "name": "BuildDynamicAccounts"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "RemoteChainCctpConfigChanged",
      "fields": [
        {
          "name": "config",
          "type": {
            "defined": "CctpChain"
          },
          "index": false
        }
      ]
    },
    {
      "name": "CcipCctpMessageSentEvent",
      "fields": [
        {
          "name": "originalSender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "msgTotalNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "eventAddress",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "cctpNonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageSentBytes",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 12000,
      "name": "InvalidTokenData",
      "msg": "Invalid token data"
    },
    {
      "code": 12001,
      "name": "InvalidReceiver",
      "msg": "Invalid receiver"
    },
    {
      "code": 12002,
      "name": "InvalidSourceDomain",
      "msg": "Invalid source domain"
    },
    {
      "code": 12003,
      "name": "InvalidDestDomain",
      "msg": "Invalid destination domain"
    },
    {
      "code": 12004,
      "name": "InvalidNonce",
      "msg": "Invalid nonce"
    },
    {
      "code": 12005,
      "name": "MalformedCctpMessage",
      "msg": "CCTP message is malformed or too short"
    },
    {
      "code": 12006,
      "name": "InvalidTokenMessengerMinter",
      "msg": "Invalid Token Messenger Minter"
    },
    {
      "code": 12007,
      "name": "InvalidMessageTransmitter",
      "msg": "Invalid Message Transmitter"
    },
    {
      "code": 12008,
      "name": "InvalidMessageSentEventAccount",
      "msg": "Invalid Message Sent Event Account"
    },
    {
      "code": 12009,
      "name": "InvalidTokenPoolExtraData",
      "msg": "Invalid Token Pool Extra Data"
    },
    {
      "code": 12010,
      "name": "FailedCctpCpi",
      "msg": "Failed CCTP CPI"
    }
  ]
};
