/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/example_ccip_sender.json`.
 */
export type ExampleCcipSender = {
  "address": "4LfBQWYaU6zQZbDyYjX8pbY4qjzrhoumUFYZEZEqMNhJ",
  "metadata": {
    "name": "exampleCcipSender",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "docs": [
    "This program an example of a CCIP Sender Program.",
    "Used to test CCIP Router ccip_send."
  ],
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
          "name": "authority",
          "signer": true
        }
      ],
      "args": []
    },
    {
      "name": "ccipSend",
      "discriminator": [
        108,
        216,
        134,
        191,
        249,
        234,
        33,
        84
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "chainConfig"
        },
        {
          "name": "ccipSender",
          "writable": true
        },
        {
          "name": "authorityFeeTokenAta",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "systemProgram"
        },
        {
          "name": "ccipRouter"
        },
        {
          "name": "ccipConfig"
        },
        {
          "name": "ccipDestChainState",
          "writable": true
        },
        {
          "name": "ccipSenderNonce",
          "writable": true
        },
        {
          "name": "ccipFeeTokenProgram"
        },
        {
          "name": "ccipFeeTokenMint"
        },
        {
          "name": "ccipFeeTokenUserAta",
          "writable": true
        },
        {
          "name": "ccipFeeTokenReceiver",
          "writable": true
        },
        {
          "name": "ccipFeeBillingSigner"
        },
        {
          "name": "ccipFeeQuoter"
        },
        {
          "name": "ccipFeeQuoterConfig"
        },
        {
          "name": "ccipFeeQuoterDestChain"
        },
        {
          "name": "ccipFeeQuoterBillingTokenConfig"
        },
        {
          "name": "ccipFeeQuoterLinkTokenConfig"
        },
        {
          "name": "ccipRmnRemote"
        },
        {
          "name": "ccipRmnRemoteCurses"
        },
        {
          "name": "ccipRmnRemoteConfig"
        }
      ],
      "args": [
        {
          "name": "destChainSelector",
          "type": "u64"
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
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "feeToken",
          "type": "pubkey"
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "initChainConfig",
      "discriminator": [
        21,
        94,
        4,
        115,
        130,
        211,
        10,
        229
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
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
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "recipient",
          "type": "bytes"
        },
        {
          "name": "extraArgsBytes",
          "type": "bytes"
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
          "name": "router",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "removeChainConfig",
      "discriminator": [
        212,
        220,
        209,
        147,
        118,
        171,
        38,
        48
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
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
          "name": "chainSelector",
          "type": "u64"
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
      "name": "updateChainConfig",
      "discriminator": [
        192,
        127,
        91,
        206,
        38,
        245,
        41,
        121
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
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
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "recipient",
          "type": "bytes"
        },
        {
          "name": "extraArgsBytes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "updateRouter",
      "discriminator": [
        32,
        109,
        12,
        153,
        101,
        129,
        64,
        70
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
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
      "name": "withdrawTokens",
      "discriminator": [
        2,
        4,
        225,
        61,
        19,
        182,
        106,
        170
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "programTokenAccount",
          "writable": true
        },
        {
          "name": "toTokenAccount",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "tokenProgram"
        },
        {
          "name": "ccipSender"
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "amount",
          "type": "u64"
        },
        {
          "name": "decimals",
          "type": "u8"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "baseState",
      "discriminator": [
        46,
        139,
        13,
        192,
        80,
        181,
        96,
        46
      ]
    },
    {
      "name": "remoteChainConfig",
      "discriminator": [
        248,
        170,
        246,
        200,
        84,
        101,
        138,
        67
      ]
    }
  ],
  "events": [
    {
      "name": "messageSent",
      "discriminator": [
        116,
        70,
        224,
        76,
        128,
        28,
        110,
        55
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "invalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6001,
      "name": "onlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6002,
      "name": "onlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6003,
      "name": "invalidProposedOwner",
      "msg": "Proposed owner is invalid"
    }
  ],
  "types": [
    {
      "name": "baseState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "pubkey"
          },
          {
            "name": "proposedOwner",
            "type": "pubkey"
          },
          {
            "name": "router",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "messageSent",
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
          }
        ]
      }
    },
    {
      "name": "remoteChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "recipient",
            "type": "bytes"
          },
          {
            "name": "extraArgsBytes",
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
