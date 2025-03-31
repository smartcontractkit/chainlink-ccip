export type TestCcipInvalidReceiver = {
  "version": "0.0.0-dev",
  "name": "test_ccip_invalid_receiver",
  "instructions": [
    {
      "name": "ccipReceive",
      "accounts": [
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA. Must be second."
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
          "name": "counter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    },
    {
      "name": "addOfframp",
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "poolProxyReleaseOrMint",
      "accounts": [
        {
          "name": "testPool",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "receiverTokenAccount",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
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
      "returns": "bytes"
    },
    {
      "name": "poolProxyLockOrBurn",
      "accounts": [
        {
          "name": "testPool",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
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
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
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
      "returns": "bytes"
    },
    {
      "name": "receiverProxyExecute",
      "accounts": [
        {
          "name": "testReceiver",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "counter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": "u8"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "ReleaseOrMintInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalSender",
            "type": "bytes"
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
            "type": "bytes"
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
    }
  ]
};

export const IDL: TestCcipInvalidReceiver = {
  "version": "0.0.0-dev",
  "name": "test_ccip_invalid_receiver",
  "instructions": [
    {
      "name": "ccipReceive",
      "accounts": [
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA. Must be second."
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
          "name": "counter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    },
    {
      "name": "addOfframp",
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "poolProxyReleaseOrMint",
      "accounts": [
        {
          "name": "testPool",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "poolTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "receiverTokenAccount",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
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
      "returns": "bytes"
    },
    {
      "name": "poolProxyLockOrBurn",
      "accounts": [
        {
          "name": "testPool",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "poolSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
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
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "chainConfig",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
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
      "returns": "bytes"
    },
    {
      "name": "receiverProxyExecute",
      "accounts": [
        {
          "name": "testReceiver",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "counter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": "u8"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "ReleaseOrMintInV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalSender",
            "type": "bytes"
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
            "type": "bytes"
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
    }
  ]
};
