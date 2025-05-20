/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/ping_pong_demo.json`.
 */
export type PingPongDemo = {
  "address": "PPbZmYFf5SPAM9Jhm9mNmYoCwT7icPYVKAfJoMCQovU",
  "metadata": {
    "name": "pingPongDemo",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
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
          "name": "config"
        },
        {
          "name": "ccipSendSigner",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "feeTokenProgram"
        },
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenAta",
          "writable": true
        },
        {
          "name": "ccipRouterProgram",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterNonce",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterFeeReceiver",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "docs": [
            "CHECK"
          ]
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
          "name": "config"
        },
        {
          "name": "nameVersion",
          "writable": true
        },
        {
          "name": "routerFeeBillingSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram"
        },
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenAta",
          "writable": true
        },
        {
          "name": "ccipSendSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "associatedTokenProgram"
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": []
    },
    {
      "name": "initializeConfig",
      "discriminator": [
        208,
        127,
        21,
        1,
        194,
        190,
        196,
        70
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "feeTokenMint"
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
          "name": "router",
          "type": "pubkey"
        },
        {
          "name": "counterpartChainSelector",
          "type": "u64"
        },
        {
          "name": "counterpartAddress",
          "type": "bytes"
        },
        {
          "name": "isPaused",
          "type": "bool"
        },
        {
          "name": "extraArgs",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "setCounterpart",
      "discriminator": [
        118,
        28,
        243,
        127,
        218,
        176,
        228,
        228
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "counterpartChainSelector",
          "type": "u64"
        },
        {
          "name": "counterpartAddress",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "setExtraArgs",
      "discriminator": [
        103,
        87,
        237,
        252,
        141,
        176,
        81,
        193
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "extraArgs",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "setPaused",
      "discriminator": [
        91,
        60,
        125,
        192,
        176,
        225,
        166,
        218
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "pause",
          "type": "bool"
        }
      ]
    },
    {
      "name": "startPingPong",
      "discriminator": [
        53,
        36,
        169,
        135,
        221,
        239,
        52,
        103
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "ccipSendSigner",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "feeTokenProgram"
        },
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenAta",
          "writable": true
        },
        {
          "name": "ccipRouterProgram",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterNonce",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterFeeReceiver",
          "docs": [
            "CHECK"
          ],
          "writable": true
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "docs": [
            "CHECK"
          ]
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
          "name": "systemProgram"
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
      "name": "config",
      "discriminator": [
        155,
        12,
        170,
        224,
        30,
        250,
        204,
        130
      ]
    },
    {
      "name": "nameVersion",
      "discriminator": [
        4,
        169,
        171,
        229,
        87,
        69,
        68,
        244
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "unauthorized",
      "msg": "unauthorized"
    },
    {
      "code": 6001,
      "name": "invalidMessageDataLength",
      "msg": "Invalid message data length"
    },
    {
      "code": 6002,
      "name": "invalidCounterpartAddress",
      "msg": "Invalid counterpart address"
    }
  ],
  "types": [
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
      "name": "config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "pubkey"
          },
          {
            "name": "router",
            "type": "pubkey"
          },
          {
            "name": "counterpartChainSelector",
            "type": "u64"
          },
          {
            "name": "counterpartAddress",
            "type": {
              "defined": {
                "name": "counterpartAddress"
              }
            }
          },
          {
            "name": "isPaused",
            "type": "bool"
          },
          {
            "name": "feeTokenMint",
            "type": "pubkey"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "counterpartAddress",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "bytes",
            "type": {
              "array": [
                "u8",
                64
              ]
            }
          },
          {
            "name": "len",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "nameVersion",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "version",
            "type": "string"
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
