export type PingPongDemo = {
  "version": "0.1.0-dev",
  "name": "ping_pong_demo",
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initialize the global config account.",
        "Call this just once."
      ],
      "accounts": [
        {
          "name": "globalConfig",
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
        }
      ]
    },
    {
      "name": "initializeChain",
      "docs": [
        "Initialize the chain's config account.",
        "Call this once for each chain you want to ping-pong with."
      ],
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
      "name": "initializeFeeToken",
      "docs": [
        "Initializes the ATA for the fee token and approve the Router for transferring from it.",
        "Call this once for each token you want to pay CCIP fees with."
      ],
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
      "args": [
        {
          "name": "counterpartChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "setCounterpart",
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "pause",
          "type": "bool"
        }
      ]
    },
    {
      "name": "setExtraArgs",
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "counterpartChainSelector",
          "type": "u64"
        },
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "tokenPoolsSigner",
          "isMut": true,
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
          "name": "counterpartChainSelector",
          "type": "u64"
        }
      ]
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
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
          "name": "tokenPoolsSigner",
          "isMut": true,
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
      "name": "globalConfig",
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
      "name": "initialize",
      "docs": [
        "Initialize the global config account.",
        "Call this just once."
      ],
      "accounts": [
        {
          "name": "globalConfig",
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
        }
      ]
    },
    {
      "name": "initializeChain",
      "docs": [
        "Initialize the chain's config account.",
        "Call this once for each chain you want to ping-pong with."
      ],
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
      "name": "initializeFeeToken",
      "docs": [
        "Initializes the ATA for the fee token and approve the Router for transferring from it.",
        "Call this once for each token you want to pay CCIP fees with."
      ],
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
      "args": [
        {
          "name": "counterpartChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "setCounterpart",
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "pause",
          "type": "bool"
        }
      ]
    },
    {
      "name": "setExtraArgs",
      "accounts": [
        {
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "counterpartChainSelector",
          "type": "u64"
        },
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
        },
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
          "name": "tokenPoolsSigner",
          "isMut": true,
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
          "name": "counterpartChainSelector",
          "type": "u64"
        }
      ]
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
          "name": "globalConfig",
          "isMut": false,
          "isSigner": false
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
          "name": "tokenPoolsSigner",
          "isMut": true,
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
      "name": "globalConfig",
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
