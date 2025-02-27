export type ExampleBurnmintTokenPool = {
  "version": "0.1.0-dev",
  "name": "example_burnmint_token_pool",
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
      "name": "transferOwnership",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
          "name": "authority",
          "isMut": false,
          "isSigner": true
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
          "isMut": true,
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
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfigAndCurses",
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
          "isMut": true,
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
          "isMut": false,
          "isSigner": false
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
          "name": "rmnRemoteConfigAndCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false
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
          }
        ]
      }
    }
  ]
};

export const IDL: ExampleBurnmintTokenPool = {
  "version": "0.1.0-dev",
  "name": "example_burnmint_token_pool",
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
      "name": "transferOwnership",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
          "name": "authority",
          "isMut": false,
          "isSigner": true
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
          "isMut": true,
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
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfigAndCurses",
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
          "isMut": true,
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
          "isMut": false,
          "isSigner": false
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
          "name": "rmnRemoteConfigAndCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false
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
          }
        ]
      }
    }
  ]
};
