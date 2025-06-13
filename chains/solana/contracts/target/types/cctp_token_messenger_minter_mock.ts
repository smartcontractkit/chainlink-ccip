export type CctpTokenMessengerMinterMock = {
  "version": "0.1.0",
  "name": "cctp_token_messenger_minter_mock",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "upgradeAuthority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgramData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "InitializeParams"
          }
        }
      ]
    },
    {
      "name": "addRemoteTokenMessenger",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "AddRemoteTokenMessengerParams"
          }
        }
      ]
    },
    {
      "name": "depositForBurn",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "burnTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "burnTokenMint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "DepositForBurnParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "depositForBurnWithCaller",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "burnTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "burnTokenMint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "DepositForBurnWithCallerParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "handleReceiveMessage",
      "accounts": [
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenPair",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "recipientTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "custodyTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "HandleReceiveMessageParams"
          }
        }
      ]
    },
    {
      "name": "setTokenController",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "SetTokenControllerParams"
          }
        }
      ]
    },
    {
      "name": "pause",
      "accounts": [
        {
          "name": "pauser",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "PauseParams"
          }
        }
      ]
    },
    {
      "name": "unpause",
      "accounts": [
        {
          "name": "pauser",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "UnpauseParams"
          }
        }
      ]
    },
    {
      "name": "updatePauser",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "UpdatePauserParams"
          }
        }
      ]
    },
    {
      "name": "setMaxBurnAmountPerMessage",
      "accounts": [
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "SetMaxBurnAmountPerMessageParams"
          }
        }
      ]
    },
    {
      "name": "addLocalToken",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "custodyTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "localTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "AddLocalTokenParams"
          }
        }
      ]
    },
    {
      "name": "linkTokenPair",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenPair",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "LinkTokenPairParams"
          }
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "tokenMessenger",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "pendingOwner",
            "type": "publicKey"
          },
          {
            "name": "localMessageTransmitter",
            "type": "publicKey"
          },
          {
            "name": "messageBodyVersion",
            "type": "u32"
          },
          {
            "name": "authorityBump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "remoteTokenMessenger",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domain",
            "type": "u32"
          },
          {
            "name": "tokenMessenger",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "tokenMinter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          },
          {
            "name": "pauser",
            "type": "publicKey"
          },
          {
            "name": "paused",
            "type": "bool"
          },
          {
            "name": "bump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "tokenPair",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "remoteToken",
            "type": "publicKey"
          },
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "bump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "localToken",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "custody",
            "type": "publicKey"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "burnLimitPerMessage",
            "type": "u64"
          },
          {
            "name": "messagesSent",
            "type": "u64"
          },
          {
            "name": "messagesReceived",
            "type": "u64"
          },
          {
            "name": "amountSent",
            "type": "u128"
          },
          {
            "name": "amountReceived",
            "type": "u128"
          },
          {
            "name": "bump",
            "type": "u8"
          },
          {
            "name": "custodyBump",
            "type": "u8"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "InitializeParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          },
          {
            "name": "localMessageTransmitter",
            "type": "publicKey"
          },
          {
            "name": "messageBodyVersion",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "AddRemoteTokenMessengerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domain",
            "type": "u32"
          },
          {
            "name": "tokenMessenger",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "DepositForBurnParams",
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
      "name": "HandleReceiveMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "sender",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          },
          {
            "name": "authorityBump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "SetTokenControllerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "PauseParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "UnpauseParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "UpdatePauserParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newPauser",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "SetMaxBurnAmountPerMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "burnLimitPerMessage",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "AddLocalTokenParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "LinkTokenPairParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "remoteToken",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "DepositForBurn",
      "fields": [
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "burnToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "depositor",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mintRecipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "destinationDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "destinationTokenMessenger",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "destinationCaller",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MintAndWithdraw",
      "fields": [
        {
          "name": "mintRecipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mintToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteTokenMessengerAdded",
      "fields": [
        {
          "name": "domain",
          "type": "u32",
          "index": false
        },
        {
          "name": "tokenMessenger",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "SetTokenController",
      "fields": [
        {
          "name": "tokenController",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "PauserChanged",
      "fields": [
        {
          "name": "newAddress",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "SetBurnLimitPerMessage",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "burnLimitPerMessage",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "LocalTokenAdded",
      "fields": [
        {
          "name": "custody",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "LocalTokenRemoved",
      "fields": [
        {
          "name": "custody",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "TokenPairLinked",
      "fields": [
        {
          "name": "localToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "remoteToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Pause",
      "fields": []
    },
    {
      "name": "Unpause",
      "fields": []
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidAuthority",
      "msg": "Invalid authority"
    },
    {
      "code": 6001,
      "name": "InvalidTokenMinterState",
      "msg": "Invalid token minter state"
    },
    {
      "code": 6002,
      "name": "ProgramPaused",
      "msg": "Program paused"
    },
    {
      "code": 6003,
      "name": "InvalidTokenPairState",
      "msg": "Invalid token pair state"
    },
    {
      "code": 6004,
      "name": "InvalidLocalTokenState",
      "msg": "Invalid local token state"
    },
    {
      "code": 6005,
      "name": "InvalidPauser",
      "msg": "Invalid pauser"
    },
    {
      "code": 6006,
      "name": "InvalidTokenController",
      "msg": "Invalid token controller"
    },
    {
      "code": 6007,
      "name": "BurnAmountExceeded",
      "msg": "Burn amount exceeded"
    },
    {
      "code": 6008,
      "name": "InvalidAmount",
      "msg": "Invalid amount"
    },
    {
      "code": 6009,
      "name": "InvalidDestinationDomain",
      "msg": "Invalid destination domain"
    },
    {
      "code": 6010,
      "name": "InvalidTokenPair",
      "msg": "Invalid token pair"
    },
    {
      "code": 6011,
      "name": "MalformedMessage",
      "msg": "Malformed message"
    },
    {
      "code": 6012,
      "name": "InvalidMessageBodyVersion",
      "msg": "Invalid message body version"
    }
  ]
};

export const IDL: CctpTokenMessengerMinterMock = {
  "version": "0.1.0",
  "name": "cctp_token_messenger_minter_mock",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "upgradeAuthority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgramData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "InitializeParams"
          }
        }
      ]
    },
    {
      "name": "addRemoteTokenMessenger",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "AddRemoteTokenMessengerParams"
          }
        }
      ]
    },
    {
      "name": "depositForBurn",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "burnTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "burnTokenMint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "DepositForBurnParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "depositForBurnWithCaller",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "burnTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "burnTokenMint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMessengerMinterProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "DepositForBurnWithCallerParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "handleReceiveMessage",
      "accounts": [
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "remoteTokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenPair",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "recipientTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "custodyTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "HandleReceiveMessageParams"
          }
        }
      ]
    },
    {
      "name": "setTokenController",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "SetTokenControllerParams"
          }
        }
      ]
    },
    {
      "name": "pause",
      "accounts": [
        {
          "name": "pauser",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "PauseParams"
          }
        }
      ]
    },
    {
      "name": "unpause",
      "accounts": [
        {
          "name": "pauser",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "UnpauseParams"
          }
        }
      ]
    },
    {
      "name": "updatePauser",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMessenger",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenMinter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "UpdatePauserParams"
          }
        }
      ]
    },
    {
      "name": "setMaxBurnAmountPerMessage",
      "accounts": [
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "SetMaxBurnAmountPerMessageParams"
          }
        }
      ]
    },
    {
      "name": "addLocalToken",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "localToken",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "custodyTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "localTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "AddLocalTokenParams"
          }
        }
      ]
    },
    {
      "name": "linkTokenPair",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "tokenController",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "tokenMinter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenPair",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "eventAuthority",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "LinkTokenPairParams"
          }
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "tokenMessenger",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "pendingOwner",
            "type": "publicKey"
          },
          {
            "name": "localMessageTransmitter",
            "type": "publicKey"
          },
          {
            "name": "messageBodyVersion",
            "type": "u32"
          },
          {
            "name": "authorityBump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "remoteTokenMessenger",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domain",
            "type": "u32"
          },
          {
            "name": "tokenMessenger",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "tokenMinter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          },
          {
            "name": "pauser",
            "type": "publicKey"
          },
          {
            "name": "paused",
            "type": "bool"
          },
          {
            "name": "bump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "tokenPair",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "remoteToken",
            "type": "publicKey"
          },
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "bump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "localToken",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "custody",
            "type": "publicKey"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "burnLimitPerMessage",
            "type": "u64"
          },
          {
            "name": "messagesSent",
            "type": "u64"
          },
          {
            "name": "messagesReceived",
            "type": "u64"
          },
          {
            "name": "amountSent",
            "type": "u128"
          },
          {
            "name": "amountReceived",
            "type": "u128"
          },
          {
            "name": "bump",
            "type": "u8"
          },
          {
            "name": "custodyBump",
            "type": "u8"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "InitializeParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          },
          {
            "name": "localMessageTransmitter",
            "type": "publicKey"
          },
          {
            "name": "messageBodyVersion",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "AddRemoteTokenMessengerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "domain",
            "type": "u32"
          },
          {
            "name": "tokenMessenger",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "DepositForBurnParams",
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
      "name": "HandleReceiveMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "sender",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          },
          {
            "name": "authorityBump",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "SetTokenControllerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenController",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "PauseParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "UnpauseParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "UpdatePauserParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newPauser",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "SetMaxBurnAmountPerMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "burnLimitPerMessage",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "AddLocalTokenParams",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "LinkTokenPairParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "localToken",
            "type": "publicKey"
          },
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "remoteToken",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "DepositForBurn",
      "fields": [
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "burnToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "depositor",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mintRecipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "destinationDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "destinationTokenMessenger",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "destinationCaller",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MintAndWithdraw",
      "fields": [
        {
          "name": "mintRecipient",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "amount",
          "type": "u64",
          "index": false
        },
        {
          "name": "mintToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "RemoteTokenMessengerAdded",
      "fields": [
        {
          "name": "domain",
          "type": "u32",
          "index": false
        },
        {
          "name": "tokenMessenger",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "SetTokenController",
      "fields": [
        {
          "name": "tokenController",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "PauserChanged",
      "fields": [
        {
          "name": "newAddress",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "SetBurnLimitPerMessage",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "burnLimitPerMessage",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "LocalTokenAdded",
      "fields": [
        {
          "name": "custody",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "LocalTokenRemoved",
      "fields": [
        {
          "name": "custody",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "mint",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "TokenPairLinked",
      "fields": [
        {
          "name": "localToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "remoteDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "remoteToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "Pause",
      "fields": []
    },
    {
      "name": "Unpause",
      "fields": []
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidAuthority",
      "msg": "Invalid authority"
    },
    {
      "code": 6001,
      "name": "InvalidTokenMinterState",
      "msg": "Invalid token minter state"
    },
    {
      "code": 6002,
      "name": "ProgramPaused",
      "msg": "Program paused"
    },
    {
      "code": 6003,
      "name": "InvalidTokenPairState",
      "msg": "Invalid token pair state"
    },
    {
      "code": 6004,
      "name": "InvalidLocalTokenState",
      "msg": "Invalid local token state"
    },
    {
      "code": 6005,
      "name": "InvalidPauser",
      "msg": "Invalid pauser"
    },
    {
      "code": 6006,
      "name": "InvalidTokenController",
      "msg": "Invalid token controller"
    },
    {
      "code": 6007,
      "name": "BurnAmountExceeded",
      "msg": "Burn amount exceeded"
    },
    {
      "code": 6008,
      "name": "InvalidAmount",
      "msg": "Invalid amount"
    },
    {
      "code": 6009,
      "name": "InvalidDestinationDomain",
      "msg": "Invalid destination domain"
    },
    {
      "code": 6010,
      "name": "InvalidTokenPair",
      "msg": "Invalid token pair"
    },
    {
      "code": 6011,
      "name": "MalformedMessage",
      "msg": "Malformed message"
    },
    {
      "code": 6012,
      "name": "InvalidMessageBodyVersion",
      "msg": "Invalid message body version"
    }
  ]
};
