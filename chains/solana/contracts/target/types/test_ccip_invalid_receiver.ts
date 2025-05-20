/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/test_ccip_invalid_receiver.json`.
 */
export type TestCcipInvalidReceiver = {
  "address": "FmyF3oW69MSAhyPSiZ69C4RKBdCPv5vAFTScisV7Me2j",
  "metadata": {
    "name": "testCcipInvalidReceiver",
    "version": "0.0.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
    {
      "name": "addOfframp",
      "discriminator": [
        164,
        255,
        154,
        96,
        204,
        239,
        24,
        2
      ],
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "ccipReceive",
      "discriminator": [
        11,
        244,
        9,
        249,
        44,
        83,
        47,
        245
      ],
      "accounts": [
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "offrampProgram",
          "docs": [
            "CHECK offramp program: exists only to derive the allowed offramp PDA",
            "and the authority PDA. Must be second."
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
          "name": "counter",
          "writable": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": {
              "name": "any2SvmMessage"
            }
          }
        }
      ]
    },
    {
      "name": "poolProxyLockOrBurn",
      "discriminator": [
        123,
        20,
        147,
        83,
        195,
        25,
        120,
        101
      ],
      "accounts": [
        {
          "name": "testPool",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "tokenProgram",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "writable": true
        },
        {
          "name": "poolSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "poolTokenAccount",
          "writable": true
        },
        {
          "name": "rmnRemote",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "chainConfig",
          "docs": [
            "CHECK"
          ],
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
      "returns": "bytes"
    },
    {
      "name": "poolProxyReleaseOrMint",
      "discriminator": [
        22,
        95,
        117,
        49,
        2,
        121,
        100,
        188
      ],
      "accounts": [
        {
          "name": "testPool",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram"
        },
        {
          "name": "allowedOfframp",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "state",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "tokenProgram",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "mint",
          "writable": true
        },
        {
          "name": "poolSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "poolTokenAccount",
          "writable": true
        },
        {
          "name": "chainConfig",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "rmnRemote",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteCurses",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "rmnRemoteConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "receiverTokenAccount",
          "docs": [
            "CHECK"
          ],
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
      "returns": "bytes"
    },
    {
      "name": "receiverProxyExecute",
      "discriminator": [
        99,
        169,
        76,
        14,
        44,
        89,
        147,
        67
      ],
      "accounts": [
        {
          "name": "testReceiver",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "cpiSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "offrampProgram",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "allowedOfframp",
          "docs": [
            "CHECK"
          ]
        }
      ],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": {
              "name": "any2SvmMessage"
            }
          }
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "discriminator": [
        247,
        97,
        179,
        16,
        207,
        36,
        236,
        132
      ]
    },
    {
      "name": "counter",
      "discriminator": [
        255,
        176,
        4,
        245,
        188,
        253,
        124,
        25
      ]
    }
  ],
  "types": [
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "any2SvmMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "messageId",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sender",
            "type": "bytes"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": {
                  "name": "svmTokenAmount"
                }
              }
            }
          }
        ]
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
      "name": "releaseOrMintInV1",
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
      "name": "svmTokenAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    }
  ]
};
