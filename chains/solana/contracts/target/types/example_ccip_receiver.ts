/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/example_ccip_receiver.json`.
 */
export type ExampleCcipReceiver = {
  "address": "48LGpn6tPn5SjTtK2wL9uUx48JUWZdZBv11sboy2orCc",
  "metadata": {
    "name": "exampleCcipReceiver",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "docs": [
    "This program an example of a CCIP Receiver Program.",
    "Used to test CCIP Router execute."
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
      "name": "approveSender",
      "discriminator": [
        110,
        115,
        180,
        233,
        200,
        99,
        131,
        255
      ],
      "accounts": [
        {
          "name": "state"
        },
        {
          "name": "approvedSender",
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
          "name": "remoteAddress",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "ccipReceive",
      "docs": [
        "This function is called by the CCIP Offramp to execute the CCIP message.",
        "The method name needs to be ccip_receive with Anchor encoding,",
        "if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]",
        "You can send as many accounts as you need, specifying if mutable or not.",
        "But none of them could be an init, realloc or close."
      ],
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
          "name": "approvedSender"
        },
        {
          "name": "state"
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
      "docs": [
        "The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts"
      ],
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
          "name": "tokenAdmin",
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
      "name": "unapproveSender",
      "discriminator": [
        156,
        35,
        66,
        182,
        129,
        232,
        105,
        176
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "approvedSender",
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
          "name": "remoteAddress",
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
          "name": "tokenAdmin"
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
      "name": "approvedSender",
      "discriminator": [
        141,
        66,
        47,
        213,
        85,
        194,
        71,
        166
      ]
    },
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
    }
  ],
  "events": [
    {
      "name": "messageReceived",
      "discriminator": [
        231,
        68,
        47,
        77,
        173,
        241,
        157,
        166
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "onlyRouter",
      "msg": "Address is not router external execution PDA"
    },
    {
      "code": 6001,
      "name": "invalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6002,
      "name": "invalidChainAndSender",
      "msg": "Invalid combination of chain and sender"
    },
    {
      "code": 6003,
      "name": "onlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6004,
      "name": "onlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6005,
      "name": "invalidCaller",
      "msg": "Caller is not allowed"
    },
    {
      "code": 6006,
      "name": "invalidProposedOwner",
      "msg": "Proposed owner is invalid"
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
      "name": "approvedSender",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
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
      "name": "messageReceived",
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
