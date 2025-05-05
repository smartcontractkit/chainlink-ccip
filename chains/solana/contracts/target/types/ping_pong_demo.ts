export type PingPongDemo = {
  "version": "0.1.0-dev",
  "name": "ping_pong_demo",
  "instructions": [
    {
      "name": "initializeConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
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
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
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
      "name": "initialize",
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "nameVersion",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "routerFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipSendSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "associatedTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
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
      "accounts": [
        {
          "name": "clock",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [],
      "returns": "string"
    },
    {
      "name": "setCounterpart",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
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
      "name": "setPaused",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
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
      "name": "setExtraArgs",
      "accounts": [
        {
          "name": "config",
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
          "name": "extraArgs",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "startPingPong",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "ccipSendSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipRouterProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterNonce",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeReceiver",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "isMut": false,
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
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "ccipReceive",
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
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSendSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipRouterProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterNonce",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeReceiver",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "isMut": false,
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
    }
  ],
  "accounts": [
    {
      "name": "config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "counterpartChainSelector",
            "type": "u64"
          },
          {
            "name": "counterpartAddress",
            "type": {
              "defined": "CounterpartAddress"
            }
          },
          {
            "name": "isPaused",
            "type": "bool"
          },
          {
            "name": "feeTokenMint",
            "type": "publicKey"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
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
    }
  ],
  "types": [
    {
      "name": "CounterpartAddress",
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
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    },
    {
      "code": 6001,
      "name": "InvalidMessageDataLength",
      "msg": "Invalid message data length"
    },
    {
      "code": 6002,
      "name": "InvalidCounterpartAddress",
      "msg": "Invalid counterpart address"
    }
  ]
};

export const IDL: PingPongDemo = {
  "version": "0.1.0-dev",
  "name": "ping_pong_demo",
  "instructions": [
    {
      "name": "initializeConfig",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
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
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
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
      "name": "initialize",
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "nameVersion",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "routerFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipSendSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "associatedTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
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
      "accounts": [
        {
          "name": "clock",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [],
      "returns": "string"
    },
    {
      "name": "setCounterpart",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
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
      "name": "setPaused",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
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
      "name": "setExtraArgs",
      "accounts": [
        {
          "name": "config",
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
          "name": "extraArgs",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "startPingPong",
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "ccipSendSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipRouterProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterNonce",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeReceiver",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "isMut": false,
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
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "ccipReceive",
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
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "ccipSendSigner",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeTokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAta",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "ccipRouterProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterDestChainState",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterNonce",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeReceiver",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "ccipRouterFeeBillingSigner",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK"
          ]
        },
        {
          "name": "feeQuoterLinkTokenConfig",
          "isMut": false,
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
    }
  ],
  "accounts": [
    {
      "name": "config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "counterpartChainSelector",
            "type": "u64"
          },
          {
            "name": "counterpartAddress",
            "type": {
              "defined": "CounterpartAddress"
            }
          },
          {
            "name": "isPaused",
            "type": "bool"
          },
          {
            "name": "feeTokenMint",
            "type": "publicKey"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
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
    }
  ],
  "types": [
    {
      "name": "CounterpartAddress",
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
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    },
    {
      "code": 6001,
      "name": "InvalidMessageDataLength",
      "msg": "Invalid message data length"
    },
    {
      "code": 6002,
      "name": "InvalidCounterpartAddress",
      "msg": "Invalid counterpart address"
    }
  ]
};
