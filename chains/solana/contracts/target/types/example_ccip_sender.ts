export type ExampleCcipSender = {
  "version": "0.1.0-dev",
  "name": "example_ccip_sender",
  "docs": [
    "This program an example of a CCIP Sender Program.",
    "Used to test CCIP Router ccip_send."
  ],
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
        }
      ]
    },
    {
      "name": "ccipSend",
      "accounts": [
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSender",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authorityFeeTokenAta",
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
          "name": "ccipRouter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipDestChainState",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipSenderNonce",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenUserAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterLinkTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemoteConfig",
          "isMut": false,
          "isSigner": false
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
              "defined": "SVMTokenAmount"
            }
          }
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "feeToken",
          "type": "publicKey"
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "updateRouter",
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
      "name": "withdrawTokens",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "programTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "toTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSender",
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
          "name": "amount",
          "type": "u64"
        },
        {
          "name": "decimals",
          "type": "u8"
        }
      ]
    },
    {
      "name": "initChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
      "name": "updateChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
      "name": "removeChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
          "name": "chainSelector",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "baseState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
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
    }
  ],
  "types": [
    {
      "name": "SenderError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "TransferTokenDuplicated"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "MessageSent",
      "fields": [
        {
          "name": "messageId",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6001,
      "name": "OnlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6002,
      "name": "OnlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6003,
      "name": "InvalidProposedOwner",
      "msg": "Proposed owner is invalid"
    }
  ]
};

export const IDL: ExampleCcipSender = {
  "version": "0.1.0-dev",
  "name": "example_ccip_sender",
  "docs": [
    "This program an example of a CCIP Sender Program.",
    "Used to test CCIP Router ccip_send."
  ],
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
        }
      ]
    },
    {
      "name": "ccipSend",
      "accounts": [
        {
          "name": "state",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "chainConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSender",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authorityFeeTokenAta",
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
          "name": "ccipRouter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipDestChainState",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipSenderNonce",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenUserAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipFeeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipFeeQuoterLinkTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipRmnRemoteConfig",
          "isMut": false,
          "isSigner": false
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
              "defined": "SVMTokenAmount"
            }
          }
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "feeToken",
          "type": "publicKey"
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "updateRouter",
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
      "name": "withdrawTokens",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "programTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "toTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSender",
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
          "name": "amount",
          "type": "u64"
        },
        {
          "name": "decimals",
          "type": "u8"
        }
      ]
    },
    {
      "name": "initChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
      "name": "updateChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
      "name": "removeChainConfig",
      "accounts": [
        {
          "name": "state",
          "isMut": true,
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
          "name": "chainSelector",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "baseState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
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
    }
  ],
  "types": [
    {
      "name": "SenderError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "TransferTokenDuplicated"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "MessageSent",
      "fields": [
        {
          "name": "messageId",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6001,
      "name": "OnlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6002,
      "name": "OnlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6003,
      "name": "InvalidProposedOwner",
      "msg": "Proposed owner is invalid"
    }
  ]
};
