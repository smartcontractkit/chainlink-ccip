export type RedirectingCcipReceiver = {
  "version": "0.1.0-dev",
  "name": "redirecting_ccip_receiver",
  "docs": [
    "This program implements a simple token redirection: It will redirect tokens",
    "to the ATA of a user described in the message's `data` field."
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts"
      ],
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenAdmin",
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
      "name": "ccipReceive",
      "docs": [
        "This function is called by the CCIP Offramp to execute the CCIP message.",
        "The method name needs to be ccip_receive with Anchor encoding,",
        "if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]",
        "You can send as many accounts as you need, specifying if mutable or not.",
        "But none of them could be an init, realloc or close."
      ],
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
          "name": "state",
          "isMut": false,
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
          "name": "tokenAdmin",
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
      "name": "setBehavior",
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
          "name": "behavior",
          "type": {
            "defined": "Behavior"
          }
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
    }
  ],
  "accounts": [
    {
      "name": "state",
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
          },
          {
            "name": "behavior",
            "type": {
              "defined": "Behavior"
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
    }
  ],
  "types": [
    {
      "name": "Any2SVMMessage",
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
                "defined": "SVMTokenAmount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "SVMTokenAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Behavior",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Normal"
          },
          {
            "name": "ExtraCUs"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "TokensRedirected",
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
        },
        {
          "name": "tokenAmount",
          "type": {
            "defined": "SVMTokenAmount"
          },
          "index": false
        },
        {
          "name": "destination",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "OnlyRouter",
      "msg": "Address is not router external execution PDA"
    },
    {
      "code": 6001,
      "name": "InvalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6002,
      "name": "InvalidChainAndSender",
      "msg": "Invalid combination of chain and sender"
    },
    {
      "code": 6003,
      "name": "OnlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6004,
      "name": "OnlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6005,
      "name": "InvalidCaller",
      "msg": "Caller is not allowed"
    },
    {
      "code": 6006,
      "name": "InvalidProposedOwner",
      "msg": "Proposed owner is invalid"
    },
    {
      "code": 6007,
      "name": "UnsupportedNumberOfTokens",
      "msg": "This redirecting receiver supports only one token transfer"
    },
    {
      "code": 6008,
      "name": "InvalidMint",
      "msg": "The provided mint account doesn't correspond to the transferred token"
    },
    {
      "code": 6009,
      "name": "InvalidMessage",
      "msg": "The message provided is not the address of the final receiver in b58 format"
    }
  ]
};

export const IDL: RedirectingCcipReceiver = {
  "version": "0.1.0-dev",
  "name": "redirecting_ccip_receiver",
  "docs": [
    "This program implements a simple token redirection: It will redirect tokens",
    "to the ATA of a user described in the message's `data` field."
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts"
      ],
      "accounts": [
        {
          "name": "state",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenAdmin",
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
      "name": "ccipReceive",
      "docs": [
        "This function is called by the CCIP Offramp to execute the CCIP message.",
        "The method name needs to be ccip_receive with Anchor encoding,",
        "if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]",
        "You can send as many accounts as you need, specifying if mutable or not.",
        "But none of them could be an init, realloc or close."
      ],
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
          "name": "state",
          "isMut": false,
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
          "name": "tokenAdmin",
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
      "name": "setBehavior",
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
          "name": "behavior",
          "type": {
            "defined": "Behavior"
          }
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
    }
  ],
  "accounts": [
    {
      "name": "state",
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
          },
          {
            "name": "behavior",
            "type": {
              "defined": "Behavior"
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
    }
  ],
  "types": [
    {
      "name": "Any2SVMMessage",
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
                "defined": "SVMTokenAmount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "SVMTokenAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Behavior",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Normal"
          },
          {
            "name": "ExtraCUs"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "TokensRedirected",
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
        },
        {
          "name": "tokenAmount",
          "type": {
            "defined": "SVMTokenAmount"
          },
          "index": false
        },
        {
          "name": "destination",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "OnlyRouter",
      "msg": "Address is not router external execution PDA"
    },
    {
      "code": 6001,
      "name": "InvalidRouter",
      "msg": "Invalid router address"
    },
    {
      "code": 6002,
      "name": "InvalidChainAndSender",
      "msg": "Invalid combination of chain and sender"
    },
    {
      "code": 6003,
      "name": "OnlyOwner",
      "msg": "Address is not owner"
    },
    {
      "code": 6004,
      "name": "OnlyProposedOwner",
      "msg": "Address is not proposed_owner"
    },
    {
      "code": 6005,
      "name": "InvalidCaller",
      "msg": "Caller is not allowed"
    },
    {
      "code": 6006,
      "name": "InvalidProposedOwner",
      "msg": "Proposed owner is invalid"
    },
    {
      "code": 6007,
      "name": "UnsupportedNumberOfTokens",
      "msg": "This redirecting receiver supports only one token transfer"
    },
    {
      "code": 6008,
      "name": "InvalidMint",
      "msg": "The provided mint account doesn't correspond to the transferred token"
    },
    {
      "code": 6009,
      "name": "InvalidMessage",
      "msg": "The message provided is not the address of the final receiver in b58 format"
    }
  ]
};
