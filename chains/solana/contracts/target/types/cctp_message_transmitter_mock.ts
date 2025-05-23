export type CctpMessageTransmitterMock = {
  "version": "0.1.0",
  "name": "cctp_message_transmitter_mock",
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
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgramData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "messageTransmitterProgram",
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
      "name": "transferOwnership",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "TransferOwnershipParams"
          }
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "accounts": [
        {
          "name": "pendingOwner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "AcceptOwnershipParams"
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
          "name": "messageTransmitter",
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
      "name": "updateAttesterManager",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "UpdateAttesterManagerParams"
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
          "name": "messageTransmitter",
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
          "name": "messageTransmitter",
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
      "name": "setMaxMessageBodySize",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "SetMaxMessageBodySizeParams"
          }
        }
      ]
    },
    {
      "name": "enableAttester",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "EnableAttesterParams"
          }
        }
      ]
    },
    {
      "name": "disableAttester",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "DisableAttesterParams"
          }
        }
      ]
    },
    {
      "name": "setSignatureThreshold",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "SetSignatureThresholdParams"
          }
        }
      ]
    },
    {
      "name": "sendMessage",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "SendMessageParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "sendMessageWithCaller",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "SendMessageWithCallerParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "replaceMessage",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "ReplaceMessageParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "receiveMessage",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "caller",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "usedNonces",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "receiver",
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
            "defined": "ReceiveMessageParams"
          }
        }
      ]
    },
    {
      "name": "reclaimEventAccount",
      "accounts": [
        {
          "name": "payee",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "ReclaimEventAccountParams"
          }
        }
      ]
    },
    {
      "name": "getNoncePda",
      "accounts": [
        {
          "name": "messageTransmitter",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "GetNoncePDAParams"
          }
        }
      ],
      "returns": "publicKey"
    },
    {
      "name": "isNonceUsed",
      "accounts": [
        {
          "name": "usedNonces",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "IsNonceUsedParams"
          }
        }
      ],
      "returns": "bool"
    }
  ],
  "accounts": [
    {
      "name": "messageTransmitter",
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
            "name": "attesterManager",
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
            "name": "localDomain",
            "type": "u32"
          },
          {
            "name": "version",
            "type": "u32"
          },
          {
            "name": "signatureThreshold",
            "type": "u32"
          },
          {
            "name": "enabledAttesters",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "maxMessageBodySize",
            "type": "u64"
          },
          {
            "name": "nextAvailableNonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "usedNonces",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "firstNonce",
            "type": "u64"
          },
          {
            "name": "usedNonces",
            "type": {
              "array": [
                "u64",
                100
              ]
            }
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
            "name": "localDomain",
            "type": "u32"
          },
          {
            "name": "attester",
            "type": "publicKey"
          },
          {
            "name": "maxMessageBodySize",
            "type": "u64"
          },
          {
            "name": "version",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "TransferOwnershipParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newOwner",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "AcceptOwnershipParams",
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
      "name": "UpdateAttesterManagerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newAttesterManager",
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
      "name": "SetMaxMessageBodySizeParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newMaxMessageBodySize",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "EnableAttesterParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newAttester",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "DisableAttesterParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "attester",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "SetSignatureThresholdParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newSignatureThreshold",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SendMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "recipient",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "SendMessageWithCallerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "recipient",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "ReplaceMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalMessage",
            "type": "bytes"
          },
          {
            "name": "originalAttestation",
            "type": "bytes"
          },
          {
            "name": "newMessageBody",
            "type": "bytes"
          },
          {
            "name": "newDestinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "ReceiveMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "message",
            "type": "bytes"
          },
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReclaimEventAccountParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "GetNoncePDAParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "nonce",
            "type": "u64"
          },
          {
            "name": "sourceDomain",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "IsNonceUsedParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "OwnershipTransferStarted",
      "fields": [
        {
          "name": "previousOwner",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newOwner",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "previousOwner",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newOwner",
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
      "name": "AttesterManagerUpdated",
      "fields": [
        {
          "name": "previousAttesterManager",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAttesterManager",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MessageReceived",
      "fields": [
        {
          "name": "caller",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "messageBody",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "SignatureThresholdUpdated",
      "fields": [
        {
          "name": "oldSignatureThreshold",
          "type": "u32",
          "index": false
        },
        {
          "name": "newSignatureThreshold",
          "type": "u32",
          "index": false
        }
      ]
    },
    {
      "name": "AttesterEnabled",
      "fields": [
        {
          "name": "attester",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AttesterDisabled",
      "fields": [
        {
          "name": "attester",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MaxMessageBodySizeUpdated",
      "fields": [
        {
          "name": "newMaxMessageBodySize",
          "type": "u64",
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
  ]
};

export const IDL: CctpMessageTransmitterMock = {
  "version": "0.1.0",
  "name": "cctp_message_transmitter_mock",
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
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitterProgramData",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "messageTransmitterProgram",
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
      "name": "transferOwnership",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "TransferOwnershipParams"
          }
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "accounts": [
        {
          "name": "pendingOwner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "AcceptOwnershipParams"
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
          "name": "messageTransmitter",
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
      "name": "updateAttesterManager",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "UpdateAttesterManagerParams"
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
          "name": "messageTransmitter",
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
          "name": "messageTransmitter",
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
      "name": "setMaxMessageBodySize",
      "accounts": [
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "SetMaxMessageBodySizeParams"
          }
        }
      ]
    },
    {
      "name": "enableAttester",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "EnableAttesterParams"
          }
        }
      ]
    },
    {
      "name": "disableAttester",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "DisableAttesterParams"
          }
        }
      ]
    },
    {
      "name": "setSignatureThreshold",
      "accounts": [
        {
          "name": "attesterManager",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
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
            "defined": "SetSignatureThresholdParams"
          }
        }
      ]
    },
    {
      "name": "sendMessage",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "SendMessageParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "sendMessageWithCaller",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "SendMessageWithCallerParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "replaceMessage",
      "accounts": [
        {
          "name": "eventRentPayer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderAuthorityPda",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "senderProgram",
          "isMut": false,
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
          "name": "params",
          "type": {
            "defined": "ReplaceMessageParams"
          }
        }
      ],
      "returns": "u64"
    },
    {
      "name": "receiveMessage",
      "accounts": [
        {
          "name": "payer",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "caller",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "authorityPda",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "messageTransmitter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "usedNonces",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "receiver",
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
            "defined": "ReceiveMessageParams"
          }
        }
      ]
    },
    {
      "name": "reclaimEventAccount",
      "accounts": [
        {
          "name": "payee",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "messageTransmitter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "messageSentEventData",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "ReclaimEventAccountParams"
          }
        }
      ]
    },
    {
      "name": "getNoncePda",
      "accounts": [
        {
          "name": "messageTransmitter",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "GetNoncePDAParams"
          }
        }
      ],
      "returns": "publicKey"
    },
    {
      "name": "isNonceUsed",
      "accounts": [
        {
          "name": "usedNonces",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "params",
          "type": {
            "defined": "IsNonceUsedParams"
          }
        }
      ],
      "returns": "bool"
    }
  ],
  "accounts": [
    {
      "name": "messageTransmitter",
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
            "name": "attesterManager",
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
            "name": "localDomain",
            "type": "u32"
          },
          {
            "name": "version",
            "type": "u32"
          },
          {
            "name": "signatureThreshold",
            "type": "u32"
          },
          {
            "name": "enabledAttesters",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "maxMessageBodySize",
            "type": "u64"
          },
          {
            "name": "nextAvailableNonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "usedNonces",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "remoteDomain",
            "type": "u32"
          },
          {
            "name": "firstNonce",
            "type": "u64"
          },
          {
            "name": "usedNonces",
            "type": {
              "array": [
                "u64",
                100
              ]
            }
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
            "name": "localDomain",
            "type": "u32"
          },
          {
            "name": "attester",
            "type": "publicKey"
          },
          {
            "name": "maxMessageBodySize",
            "type": "u64"
          },
          {
            "name": "version",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "TransferOwnershipParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newOwner",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "AcceptOwnershipParams",
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
      "name": "UpdateAttesterManagerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newAttesterManager",
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
      "name": "SetMaxMessageBodySizeParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newMaxMessageBodySize",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "EnableAttesterParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newAttester",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "DisableAttesterParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "attester",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "SetSignatureThresholdParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "newSignatureThreshold",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SendMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "recipient",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "SendMessageWithCallerParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destinationDomain",
            "type": "u32"
          },
          {
            "name": "recipient",
            "type": "publicKey"
          },
          {
            "name": "messageBody",
            "type": "bytes"
          },
          {
            "name": "destinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "ReplaceMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "originalMessage",
            "type": "bytes"
          },
          {
            "name": "originalAttestation",
            "type": "bytes"
          },
          {
            "name": "newMessageBody",
            "type": "bytes"
          },
          {
            "name": "newDestinationCaller",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "ReceiveMessageParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "message",
            "type": "bytes"
          },
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "ReclaimEventAccountParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "attestation",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "GetNoncePDAParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "nonce",
            "type": "u64"
          },
          {
            "name": "sourceDomain",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "IsNonceUsedParams",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "OwnershipTransferStarted",
      "fields": [
        {
          "name": "previousOwner",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newOwner",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "previousOwner",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newOwner",
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
      "name": "AttesterManagerUpdated",
      "fields": [
        {
          "name": "previousAttesterManager",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAttesterManager",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MessageReceived",
      "fields": [
        {
          "name": "caller",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "sourceDomain",
          "type": "u32",
          "index": false
        },
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "sender",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "messageBody",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "SignatureThresholdUpdated",
      "fields": [
        {
          "name": "oldSignatureThreshold",
          "type": "u32",
          "index": false
        },
        {
          "name": "newSignatureThreshold",
          "type": "u32",
          "index": false
        }
      ]
    },
    {
      "name": "AttesterEnabled",
      "fields": [
        {
          "name": "attester",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AttesterDisabled",
      "fields": [
        {
          "name": "attester",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "MaxMessageBodySizeUpdated",
      "fields": [
        {
          "name": "newMaxMessageBodySize",
          "type": "u64",
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
  ]
};
