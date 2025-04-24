export type CcipOfframp = {
  "version": "0.1.0-dev",
  "name": "ccip_offramp",
  "constants": [
    {
      "name": "MAX_ORACLES",
      "type": {
        "defined": "usize"
      },
      "value": "16"
    }
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initialization Flow //",
        "Initializes the CCIP Offramp, except for the config account (due to stack size limitations).",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization."
      ],
      "accounts": [
        {
          "name": "referenceAddresses",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "router",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "offrampLookupTable",
          "isMut": false,
          "isSigner": false
        },
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
      "args": []
    },
    {
      "name": "initializeConfig",
      "docs": [
        "Initializes the CCIP Offramp Config account.",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization of the config.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
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
          "name": "svmChainSelector",
          "type": "u64"
        },
        {
          "name": "enableExecutionAfter",
          "type": "i64"
        }
      ]
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
      "name": "transferOwnership",
      "docs": [
        "Transfers the ownership of the router to a new proposed owner.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the transfer.",
        "* `proposed_owner` - The public key of the new proposed owner."
      ],
      "accounts": [
        {
          "name": "config",
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
      "docs": [
        "Accepts the ownership of the router by the proposed owner.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for accepting ownership.",
        "The new owner must be a signer of the transaction."
      ],
      "accounts": [
        {
          "name": "config",
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
      "name": "setDefaultCodeVersion",
      "docs": [
        "Sets the default code version to be used. This is then used by the slim routing layer to determine",
        "which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `code_version` - The new code version to be set as default."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "codeVersion",
          "type": {
            "defined": "CodeVersion"
          }
        }
      ]
    },
    {
      "name": "updateReferenceAddresses",
      "docs": [
        "Updates reference addresses in the offramp contract, such as",
        "the CCIP router, Fee Quoter, and the Offramp Lookup Table.",
        "Only the Admin may update these addresses.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the reference addresses.",
        "* `router` - The router address to be set.",
        "* `fee_quoter` - The fee_quoter address to be set.",
        "* `offramp_lookup_table` - The offramp_lookup_table address to be set.",
        "* `rmn_remote` - The rmn_remote address to be set."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
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
          "name": "router",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "offrampLookupTable",
          "type": "publicKey"
        },
        {
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "addSourceChain",
      "docs": [
        "Adds a new source chain selector with its config to the offramp.",
        "",
        "The Admin needs to add any new chain supported.",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "",
        "# Arguments"
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "Adding a chain selector implies initializing the state for a new chain"
          ]
        },
        {
          "name": "config",
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
          "name": "newChainSelector",
          "type": "u64"
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          }
        }
      ]
    },
    {
      "name": "disableSourceChainSelector",
      "docs": [
        "Disables the source chain selector.",
        "",
        "The Admin is the only one able to disable the chain selector as source. This method is thought of as an emergency kill-switch.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for disabling the chain selector.",
        "* `source_chain_selector` - The source chain selector to be disabled."
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
          "name": "sourceChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateSourceChainConfig",
      "docs": [
        "Updates the configuration of the source chain selector.",
        "",
        "The Admin is the only one able to update the source chain config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the chain selector.",
        "* `source_chain_selector` - The source chain selector to be updated.",
        "* `source_chain_config` - The new configuration for the source chain."
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          }
        }
      ]
    },
    {
      "name": "updateSvmChainSelector",
      "docs": [
        "Updates the SVM chain selector in the offramp configuration.",
        "",
        "This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `new_chain_selector` - The new chain selector for SVM."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "newChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateEnableManualExecutionAfter",
      "docs": [
        "Updates the minimum amount of time required between a message being committed and when it can be manually executed.",
        "",
        "This is part of the OffRamp Configuration for SVM.",
        "The Admin is the only one able to update this config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `new_enable_manual_execution_after` - The new minimum amount of time required."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "newEnableManualExecutionAfter",
          "type": "i64"
        }
      ]
    },
    {
      "name": "setOcrConfig",
      "docs": [
        "Sets the OCR configuration.",
        "Only CCIP Admin can set the OCR configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for setting the OCR configuration.",
        "* `plugin_type` - The type of OCR plugin [0: Commit, 1: Execution].",
        "* `config_info` - The OCR configuration information.",
        "* `signers` - The list of signers.",
        "* `transmitters` - The list of transmitters."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
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
          "name": "pluginType",
          "type": {
            "defined": "OcrPluginType"
          }
        },
        {
          "name": "configInfo",
          "type": {
            "defined": "Ocr3ConfigInfo"
          }
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          }
        },
        {
          "name": "transmitters",
          "type": {
            "vec": "publicKey"
          }
        }
      ]
    },
    {
      "name": "commit",
      "docs": [
        "Off Ramp Flow //",
        "Commits a report to the router, containing a Merkle Root.",
        "",
        "The method name needs to be commit with Anchor encoding.",
        "",
        "This function is called by the OffChain when committing one Report to the SVM Router.",
        "In this Flow only one report is sent, the Commit Report. This is different as EVM does,",
        "this is because here all the chain state is stored in one account per Merkle Tree Root.",
        "So, to avoid having to send a dynamic size array of accounts, in this message only one Commit Report Account is sent.",
        "This message validates the signatures of the report and stores the Merkle Root in the Commit Report Account.",
        "The Report must contain an interval of messages, and the min of them must be the next sequence number expected.",
        "The max size of the interval is 64.",
        "This message emits two events: CommitReportAccepted and Transmitted.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the commit.",
        "* `report_context_byte_words` - consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number",
        "* `raw_report` - The serialized commit input report, single merkle root with RMN signatures and price updates",
        "* `rs` - slice of R components of signatures",
        "* `ss` - slice of S components of signatures",
        "* `raw_vs` - array of V components of signatures"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "commitReport",
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "rawReport",
          "type": "bytes"
        },
        {
          "name": "rs",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "ss",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "rawVs",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        }
      ]
    },
    {
      "name": "commitPriceOnly",
      "docs": [
        "Commits a report to the router, with price updates only.",
        "",
        "The method name needs to be commit with Anchor encoding.",
        "",
        "This function is called by the OffChain when committing one Report to the SVM Router,",
        "containing only price updates and no merkle root.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the commit.",
        "* `report_context_byte_words` - consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number",
        "* `raw_report` - The serialized commit input report containing the price updates,",
        "with no merkle root.",
        "* `rs` - slice of R components of signatures",
        "* `ss` - slice of S components of signatures",
        "* `raw_vs` - array of V components of signatures"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "rawReport",
          "type": "bytes"
        },
        {
          "name": "rs",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "ss",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "rawVs",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        }
      ]
    },
    {
      "name": "execute",
      "docs": [
        "Executes a message on the destination chain.",
        "",
        "The method name needs to be execute with Anchor encoding.",
        "",
        "This function is called by the OffChain when executing one Report to the SVM Router.",
        "In this Flow only one message is sent, the Execution Report. This is different as EVM does,",
        "this is because there is no try/catch mechanism to allow batch execution.",
        "This message validates that the Merkle Tree Proof of the given message is correct and is stored in the Commit Report Account.",
        "The message must be untouched to be executed.",
        "This message emits the event ExecutionStateChanged with the new state of the message.",
        "Finally, executes the CPI instruction to the receiver program in the ccip_receive message.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the execute.",
        "* `raw_execution_report` - the serialized execution report containing only one message and proofs",
        "* `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)",
        "*  consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rawExecutionReport",
          "type": "bytes"
        },
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "manuallyExecute",
      "docs": [
        "Manually executes a report to the router.",
        "",
        "When a message is not being executed, then the user can trigger the execution manually.",
        "No verification over the transmitter, but the message needs to be in some commit report.",
        "It validates that the required time has passed since the commit and then executes the report.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the execution.",
        "* `raw_execution_report` - The serialized execution report containing the message and proofs."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rawExecutionReport",
          "type": "bytes"
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "closeCommitReportAccount",
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "wsolMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "root",
          "type": "bytes"
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
            "name": "version",
            "type": "u8"
          },
          {
            "name": "defaultCodeVersion",
            "type": "u8"
          },
          {
            "name": "padding0",
            "type": {
              "array": [
                "u8",
                6
              ]
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "enableManualExecutionAfter",
            "type": "i64"
          },
          {
            "name": "padding1",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "padding2",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "ocr3",
            "type": {
              "array": [
                {
                  "defined": "Ocr3Config"
                },
                2
              ]
            }
          }
        ]
      }
    },
    {
      "name": "referenceAddresses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "offrampLookupTable",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "globalState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "latestPriceSequenceNumber",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "sourceChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "state",
            "type": {
              "defined": "SourceChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "SourceChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "commitReport",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          },
          {
            "name": "minMsgNr",
            "type": "u64"
          },
          {
            "name": "maxMsgNr",
            "type": "u64"
          },
          {
            "name": "executionStates",
            "type": "u128"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CommitInput",
      "docs": [
        "Input from an offchain node, containing the Merkle root and interval for",
        "the source chain, and optionally some price updates alongside it"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "priceUpdates",
            "type": {
              "defined": "PriceUpdates"
            }
          },
          {
            "name": "merkleRoot",
            "type": {
              "option": {
                "defined": "MerkleRoot"
              }
            }
          },
          {
            "name": "rmnSignatures",
            "type": {
              "vec": {
                "array": [
                  "u8",
                  64
                ]
              }
            }
          }
        ]
      }
    },
    {
      "name": "PriceUpdates",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenPriceUpdates",
            "type": {
              "vec": {
                "defined": "TokenPriceUpdate"
              }
            }
          },
          {
            "name": "gasPriceUpdates",
            "type": {
              "vec": {
                "defined": "GasPriceUpdate"
              }
            }
          }
        ]
      }
    },
    {
      "name": "TokenPriceUpdate",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceToken",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "GasPriceUpdate",
      "docs": [
        "Gas price for a given chain in USD; its value may contain tightly packed fields."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "usdPerUnitGas",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "MerkleRoot",
      "docs": [
        "Struct to hold a merkle root and an interval for a source chain"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "onRampAddress",
            "type": "bytes"
          },
          {
            "name": "minSeqNr",
            "type": "u64"
          },
          {
            "name": "maxSeqNr",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
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
      "name": "ConfigOcrPluginType",
      "docs": [
        "It's not possible to store enums in zero_copy accounts, so we wrap the discriminant",
        "in a struct to store in config."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "discriminant",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "ExecutionReportSingleChain",
      "docs": [
        "Report that is submitted by the execution DON at the execution phase. (including chain selector data)"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "message",
            "type": {
              "defined": "Any2SVMRampMessage"
            }
          },
          {
            "name": "offchainTokenData",
            "type": {
              "vec": "bytes"
            }
          },
          {
            "name": "proofs",
            "type": {
              "vec": {
                "array": [
                  "u8",
                  32
                ]
              }
            }
          }
        ]
      }
    },
    {
      "name": "RampMessageHeader",
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
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Any2SVMRampExtraArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "computeUnits",
            "type": "u32"
          },
          {
            "name": "isWritableBitmap",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Any2SVMRampMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "header",
            "type": {
              "defined": "RampMessageHeader"
            }
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
            "name": "tokenReceiver",
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "Any2SVMTokenTransfer"
              }
            }
          },
          {
            "name": "extraArgs",
            "type": {
              "defined": "Any2SVMRampExtraArgs"
            }
          }
        ]
      }
    },
    {
      "name": "Any2SVMTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "bytes"
          },
          {
            "name": "destTokenAddress",
            "type": "publicKey"
          },
          {
            "name": "destGasAmount",
            "type": "u32"
          },
          {
            "name": "extraData",
            "type": "bytes"
          },
          {
            "name": "amount",
            "type": {
              "defined": "CrossChainAmount"
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
      "name": "CrossChainAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "leBytes",
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
      "name": "Ocr3ConfigInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "configDigest",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "f",
            "type": "u8"
          },
          {
            "name": "n",
            "type": "u8"
          },
          {
            "name": "isSignatureVerificationEnabled",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3Config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pluginType",
            "type": {
              "defined": "ConfigOcrPluginType"
            }
          },
          {
            "name": "configInfo",
            "type": {
              "defined": "Ocr3ConfigInfo"
            }
          },
          {
            "name": "signers",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    20
                  ]
                },
                16
              ]
            }
          },
          {
            "name": "transmitters",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    32
                  ]
                },
                16
              ]
            }
          }
        ]
      }
    },
    {
      "name": "SourceChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "isRmnVerificationDisabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "onRamp",
            "type": {
              "defined": "OnRampAddress"
            }
          }
        ]
      }
    },
    {
      "name": "OnRampAddress",
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
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SourceChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minSeqNr",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "OcrPluginType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Commit"
          },
          {
            "name": "Execution"
          }
        ]
      }
    },
    {
      "name": "MessageExecutionState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Untouched"
          },
          {
            "name": "InProgress"
          },
          {
            "name": "Success"
          },
          {
            "name": "Failure"
          }
        ]
      }
    },
    {
      "name": "CodeVersion",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Default"
          },
          {
            "name": "V1"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "SourceChainConfigUpdated",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SourceChainAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferRequested",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "enableManualExecutionAfter",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "ReferenceAddressesSet",
      "fields": [
        {
          "name": "router",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "feeQuoter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "offrampLookupTable",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "rmnRemote",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "CommitReportAccepted",
      "fields": [
        {
          "name": "merkleRoot",
          "type": {
            "option": {
              "defined": "MerkleRoot"
            }
          },
          "index": false
        },
        {
          "name": "priceUpdates",
          "type": {
            "defined": "PriceUpdates"
          },
          "index": false
        }
      ]
    },
    {
      "name": "CommitReportPDAClosed",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "merkleRoot",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "SkippedAlreadyExecutedMessage",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "ExecutionStateChanged",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
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
          "name": "messageHash",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "state",
          "type": {
            "defined": "MessageExecutionState"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "ocrPluginType",
          "type": {
            "defined": "OcrPluginType"
          },
          "index": false
        },
        {
          "name": "configDigest",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          },
          "index": false
        },
        {
          "name": "transmitters",
          "type": {
            "vec": "publicKey"
          },
          "index": false
        },
        {
          "name": "f",
          "type": "u8",
          "index": false
        }
      ]
    },
    {
      "name": "Transmitted",
      "fields": [
        {
          "name": "ocrPluginType",
          "type": {
            "defined": "OcrPluginType"
          },
          "index": false
        },
        {
          "name": "configDigest",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "InvalidSequenceInterval",
      "msg": "The given sequence interval is invalid"
    },
    {
      "code": 9001,
      "name": "RootNotCommitted",
      "msg": "The given Merkle Root is missing"
    },
    {
      "code": 9002,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 9003,
      "name": "ExistingMerkleRoot",
      "msg": "The given Merkle Root is already committed"
    },
    {
      "code": 9004,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9005,
      "name": "InvalidNonce",
      "msg": "Invalid Nonce"
    },
    {
      "code": 9006,
      "name": "InvalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 9007,
      "name": "OnrampNotConfigured",
      "msg": "Onramp was not configured"
    },
    {
      "code": 9008,
      "name": "FailedToDeserializeReport",
      "msg": "Failed to deserialize report"
    },
    {
      "code": 9009,
      "name": "InvalidPluginType",
      "msg": "Invalid plugin type"
    },
    {
      "code": 9010,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9011,
      "name": "MissingExpectedPriceUpdates",
      "msg": "Commit report is missing expected price updates"
    },
    {
      "code": 9012,
      "name": "MissingExpectedMerkleRoot",
      "msg": "Commit report is missing expected merkle root"
    },
    {
      "code": 9013,
      "name": "UnexpectedMerkleRoot",
      "msg": "Commit report contains unexpected merkle root"
    },
    {
      "code": 9014,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9015,
      "name": "UnsupportedSourceChainSelector",
      "msg": "Source chain selector not supported"
    },
    {
      "code": 9016,
      "name": "UnsupportedDestinationChainSelector",
      "msg": "Destination chain selector not supported"
    },
    {
      "code": 9017,
      "name": "InvalidProof",
      "msg": "Invalid Proof for Merkle Root"
    },
    {
      "code": 9018,
      "name": "InvalidMessage",
      "msg": "Invalid message format"
    },
    {
      "code": 9019,
      "name": "ReachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 9020,
      "name": "ManualExecutionNotAllowed",
      "msg": "Manual execution not allowed"
    },
    {
      "code": 9021,
      "name": "InvalidInputsNumberOfAccounts",
      "msg": "Number of accounts is invalid"
    },
    {
      "code": 9022,
      "name": "InvalidInputsGlobalStateAccount",
      "msg": "Invalid global state account address"
    },
    {
      "code": 9023,
      "name": "InvalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 9024,
      "name": "InvalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 9025,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 9026,
      "name": "InvalidInputsSysvarAccount",
      "msg": "Invalid sysvar instructions account"
    },
    {
      "code": 9027,
      "name": "InvalidInputsFeeQuoterAccount",
      "msg": "Invalid fee quoter account"
    },
    {
      "code": 9028,
      "name": "InvalidInputsAllowedOfframpAccount",
      "msg": "Invalid offramp authorization account"
    },
    {
      "code": 9029,
      "name": "InvalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 9030,
      "name": "InvalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 9031,
      "name": "InvalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 9032,
      "name": "OfframpReleaseMintBalanceMismatch",
      "msg": "Release or mint balance mismatch"
    },
    {
      "code": 9033,
      "name": "OfframpInvalidDataLength",
      "msg": "Invalid data length"
    },
    {
      "code": 9034,
      "name": "StaleCommitReport",
      "msg": "Stale commit report"
    },
    {
      "code": 9035,
      "name": "InvalidWritabilityBitmap",
      "msg": "Invalid writability bitmap"
    },
    {
      "code": 9036,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 9037,
      "name": "Ocr3InvalidConfigFMustBePositive",
      "msg": "Invalid config: F must be positive"
    },
    {
      "code": 9038,
      "name": "Ocr3InvalidConfigTooManyTransmitters",
      "msg": "Invalid config: Too many transmitters"
    },
    {
      "code": 9039,
      "name": "Ocr3InvalidConfigNoTransmitters",
      "msg": "Invalid config: No transmitters"
    },
    {
      "code": 9040,
      "name": "Ocr3InvalidConfigTooManySigners",
      "msg": "Invalid config: Too many signers"
    },
    {
      "code": 9041,
      "name": "Ocr3InvalidConfigFIsTooHigh",
      "msg": "Invalid config: F is too high"
    },
    {
      "code": 9042,
      "name": "Ocr3InvalidConfigRepeatedOracle",
      "msg": "Invalid config: Repeated oracle address"
    },
    {
      "code": 9043,
      "name": "Ocr3WrongMessageLength",
      "msg": "Wrong message length"
    },
    {
      "code": 9044,
      "name": "Ocr3ConfigDigestMismatch",
      "msg": "Config digest mismatch"
    },
    {
      "code": 9045,
      "name": "Ocr3WrongNumberOfSignatures",
      "msg": "Wrong number signatures"
    },
    {
      "code": 9046,
      "name": "Ocr3UnauthorizedTransmitter",
      "msg": "Unauthorized transmitter"
    },
    {
      "code": 9047,
      "name": "Ocr3UnauthorizedSigner",
      "msg": "Unauthorized signer"
    },
    {
      "code": 9048,
      "name": "Ocr3NonUniqueSignatures",
      "msg": "Non unique signatures"
    },
    {
      "code": 9049,
      "name": "Ocr3OracleCannotBeZeroAddress",
      "msg": "Oracle cannot be zero address"
    },
    {
      "code": 9050,
      "name": "Ocr3StaticConfigCannotBeChanged",
      "msg": "Static config cannot be changed"
    },
    {
      "code": 9051,
      "name": "Ocr3InvalidPluginType",
      "msg": "Incorrect plugin type"
    },
    {
      "code": 9052,
      "name": "Ocr3InvalidSignature",
      "msg": "Invalid signature"
    },
    {
      "code": 9053,
      "name": "Ocr3SignaturesOutOfRegistration",
      "msg": "Signatures out of registration"
    },
    {
      "code": 9054,
      "name": "InvalidOnrampAddress",
      "msg": "Invalid onramp address"
    },
    {
      "code": 9055,
      "name": "InvalidInputsExternalExecutionSignerAccount",
      "msg": "Invalid external execution signer account"
    },
    {
      "code": 9056,
      "name": "CommitReportHasPendingMessages",
      "msg": "Commit report has pending messages"
    }
  ]
};

export const IDL: CcipOfframp = {
  "version": "0.1.0-dev",
  "name": "ccip_offramp",
  "constants": [
    {
      "name": "MAX_ORACLES",
      "type": {
        "defined": "usize"
      },
      "value": "16"
    }
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initialization Flow //",
        "Initializes the CCIP Offramp, except for the config account (due to stack size limitations).",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization."
      ],
      "accounts": [
        {
          "name": "referenceAddresses",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "router",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "offrampLookupTable",
          "isMut": false,
          "isSigner": false
        },
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
      "args": []
    },
    {
      "name": "initializeConfig",
      "docs": [
        "Initializes the CCIP Offramp Config account.",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization of the config.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
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
          "name": "svmChainSelector",
          "type": "u64"
        },
        {
          "name": "enableExecutionAfter",
          "type": "i64"
        }
      ]
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
      "name": "transferOwnership",
      "docs": [
        "Transfers the ownership of the router to a new proposed owner.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the transfer.",
        "* `proposed_owner` - The public key of the new proposed owner."
      ],
      "accounts": [
        {
          "name": "config",
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
      "docs": [
        "Accepts the ownership of the router by the proposed owner.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for accepting ownership.",
        "The new owner must be a signer of the transaction."
      ],
      "accounts": [
        {
          "name": "config",
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
      "name": "setDefaultCodeVersion",
      "docs": [
        "Sets the default code version to be used. This is then used by the slim routing layer to determine",
        "which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.",
        "",
        "Shared func signature with other programs",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `code_version` - The new code version to be set as default."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "codeVersion",
          "type": {
            "defined": "CodeVersion"
          }
        }
      ]
    },
    {
      "name": "updateReferenceAddresses",
      "docs": [
        "Updates reference addresses in the offramp contract, such as",
        "the CCIP router, Fee Quoter, and the Offramp Lookup Table.",
        "Only the Admin may update these addresses.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the reference addresses.",
        "* `router` - The router address to be set.",
        "* `fee_quoter` - The fee_quoter address to be set.",
        "* `offramp_lookup_table` - The offramp_lookup_table address to be set.",
        "* `rmn_remote` - The rmn_remote address to be set."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
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
          "name": "router",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "offrampLookupTable",
          "type": "publicKey"
        },
        {
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "addSourceChain",
      "docs": [
        "Adds a new source chain selector with its config to the offramp.",
        "",
        "The Admin needs to add any new chain supported.",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "",
        "# Arguments"
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "Adding a chain selector implies initializing the state for a new chain"
          ]
        },
        {
          "name": "config",
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
          "name": "newChainSelector",
          "type": "u64"
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          }
        }
      ]
    },
    {
      "name": "disableSourceChainSelector",
      "docs": [
        "Disables the source chain selector.",
        "",
        "The Admin is the only one able to disable the chain selector as source. This method is thought of as an emergency kill-switch.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for disabling the chain selector.",
        "* `source_chain_selector` - The source chain selector to be disabled."
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
          "name": "sourceChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateSourceChainConfig",
      "docs": [
        "Updates the configuration of the source chain selector.",
        "",
        "The Admin is the only one able to update the source chain config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the chain selector.",
        "* `source_chain_selector` - The source chain selector to be updated.",
        "* `source_chain_config` - The new configuration for the source chain."
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          }
        }
      ]
    },
    {
      "name": "updateSvmChainSelector",
      "docs": [
        "Updates the SVM chain selector in the offramp configuration.",
        "",
        "This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `new_chain_selector` - The new chain selector for SVM."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "newChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateEnableManualExecutionAfter",
      "docs": [
        "Updates the minimum amount of time required between a message being committed and when it can be manually executed.",
        "",
        "This is part of the OffRamp Configuration for SVM.",
        "The Admin is the only one able to update this config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `new_enable_manual_execution_after` - The new minimum amount of time required."
      ],
      "accounts": [
        {
          "name": "config",
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
          "name": "newEnableManualExecutionAfter",
          "type": "i64"
        }
      ]
    },
    {
      "name": "setOcrConfig",
      "docs": [
        "Sets the OCR configuration.",
        "Only CCIP Admin can set the OCR configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for setting the OCR configuration.",
        "* `plugin_type` - The type of OCR plugin [0: Commit, 1: Execution].",
        "* `config_info` - The OCR configuration information.",
        "* `signers` - The list of signers.",
        "* `transmitters` - The list of transmitters."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
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
          "name": "pluginType",
          "type": {
            "defined": "OcrPluginType"
          }
        },
        {
          "name": "configInfo",
          "type": {
            "defined": "Ocr3ConfigInfo"
          }
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          }
        },
        {
          "name": "transmitters",
          "type": {
            "vec": "publicKey"
          }
        }
      ]
    },
    {
      "name": "commit",
      "docs": [
        "Off Ramp Flow //",
        "Commits a report to the router, containing a Merkle Root.",
        "",
        "The method name needs to be commit with Anchor encoding.",
        "",
        "This function is called by the OffChain when committing one Report to the SVM Router.",
        "In this Flow only one report is sent, the Commit Report. This is different as EVM does,",
        "this is because here all the chain state is stored in one account per Merkle Tree Root.",
        "So, to avoid having to send a dynamic size array of accounts, in this message only one Commit Report Account is sent.",
        "This message validates the signatures of the report and stores the Merkle Root in the Commit Report Account.",
        "The Report must contain an interval of messages, and the min of them must be the next sequence number expected.",
        "The max size of the interval is 64.",
        "This message emits two events: CommitReportAccepted and Transmitted.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the commit.",
        "* `report_context_byte_words` - consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number",
        "* `raw_report` - The serialized commit input report, single merkle root with RMN signatures and price updates",
        "* `rs` - slice of R components of signatures",
        "* `ss` - slice of S components of signatures",
        "* `raw_vs` - array of V components of signatures"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "commitReport",
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "rawReport",
          "type": "bytes"
        },
        {
          "name": "rs",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "ss",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "rawVs",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        }
      ]
    },
    {
      "name": "commitPriceOnly",
      "docs": [
        "Commits a report to the router, with price updates only.",
        "",
        "The method name needs to be commit with Anchor encoding.",
        "",
        "This function is called by the OffChain when committing one Report to the SVM Router,",
        "containing only price updates and no merkle root.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the commit.",
        "* `report_context_byte_words` - consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number",
        "* `raw_report` - The serialized commit input report containing the price updates,",
        "with no merkle root.",
        "* `rs` - slice of R components of signatures",
        "* `ss` - slice of S components of signatures",
        "* `raw_vs` - array of V components of signatures"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "rawReport",
          "type": "bytes"
        },
        {
          "name": "rs",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "ss",
          "type": {
            "vec": {
              "array": [
                "u8",
                32
              ]
            }
          }
        },
        {
          "name": "rawVs",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        }
      ]
    },
    {
      "name": "execute",
      "docs": [
        "Executes a message on the destination chain.",
        "",
        "The method name needs to be execute with Anchor encoding.",
        "",
        "This function is called by the OffChain when executing one Report to the SVM Router.",
        "In this Flow only one message is sent, the Execution Report. This is different as EVM does,",
        "this is because there is no try/catch mechanism to allow batch execution.",
        "This message validates that the Merkle Tree Proof of the given message is correct and is stored in the Commit Report Account.",
        "The message must be untouched to be executed.",
        "This message emits the event ExecutionStateChanged with the new state of the message.",
        "Finally, executes the CPI instruction to the receiver program in the ccip_receive message.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the execute.",
        "* `raw_execution_report` - the serialized execution report containing only one message and proofs",
        "* `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)",
        "*  consists of:",
        "* report_context_byte_words[0]: ConfigDigest",
        "* report_context_byte_words[1]: 24 byte padding, 8 byte sequence number"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rawExecutionReport",
          "type": "bytes"
        },
        {
          "name": "reportContextByteWords",
          "type": {
            "array": [
              {
                "array": [
                  "u8",
                  32
                ]
              },
              2
            ]
          }
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "manuallyExecute",
      "docs": [
        "Manually executes a report to the router.",
        "",
        "When a message is not being executed, then the user can trigger the execution manually.",
        "No verification over the transmitter, but the message needs to be in some commit report.",
        "It validates that the required time has passed since the commit and then executes the report.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the execution.",
        "* `raw_execution_report` - The serialized execution report containing the message and proofs."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "sourceChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "allowedOfframp",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemote",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteCurses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "rmnRemoteConfig",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rawExecutionReport",
          "type": "bytes"
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "closeCommitReportAccount",
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "commitReport",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "wsolMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "root",
          "type": "bytes"
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
            "name": "version",
            "type": "u8"
          },
          {
            "name": "defaultCodeVersion",
            "type": "u8"
          },
          {
            "name": "padding0",
            "type": {
              "array": [
                "u8",
                6
              ]
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "enableManualExecutionAfter",
            "type": "i64"
          },
          {
            "name": "padding1",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "padding2",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "ocr3",
            "type": {
              "array": [
                {
                  "defined": "Ocr3Config"
                },
                2
              ]
            }
          }
        ]
      }
    },
    {
      "name": "referenceAddresses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "offrampLookupTable",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "globalState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "latestPriceSequenceNumber",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "sourceChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "state",
            "type": {
              "defined": "SourceChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "SourceChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "commitReport",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          },
          {
            "name": "minMsgNr",
            "type": "u64"
          },
          {
            "name": "maxMsgNr",
            "type": "u64"
          },
          {
            "name": "executionStates",
            "type": "u128"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CommitInput",
      "docs": [
        "Input from an offchain node, containing the Merkle root and interval for",
        "the source chain, and optionally some price updates alongside it"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "priceUpdates",
            "type": {
              "defined": "PriceUpdates"
            }
          },
          {
            "name": "merkleRoot",
            "type": {
              "option": {
                "defined": "MerkleRoot"
              }
            }
          },
          {
            "name": "rmnSignatures",
            "type": {
              "vec": {
                "array": [
                  "u8",
                  64
                ]
              }
            }
          }
        ]
      }
    },
    {
      "name": "PriceUpdates",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenPriceUpdates",
            "type": {
              "vec": {
                "defined": "TokenPriceUpdate"
              }
            }
          },
          {
            "name": "gasPriceUpdates",
            "type": {
              "vec": {
                "defined": "GasPriceUpdate"
              }
            }
          }
        ]
      }
    },
    {
      "name": "TokenPriceUpdate",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceToken",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "GasPriceUpdate",
      "docs": [
        "Gas price for a given chain in USD; its value may contain tightly packed fields."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "usdPerUnitGas",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "MerkleRoot",
      "docs": [
        "Struct to hold a merkle root and an interval for a source chain"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "onRampAddress",
            "type": "bytes"
          },
          {
            "name": "minSeqNr",
            "type": "u64"
          },
          {
            "name": "maxSeqNr",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
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
      "name": "ConfigOcrPluginType",
      "docs": [
        "It's not possible to store enums in zero_copy accounts, so we wrap the discriminant",
        "in a struct to store in config."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "discriminant",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "ExecutionReportSingleChain",
      "docs": [
        "Report that is submitted by the execution DON at the execution phase. (including chain selector data)"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "message",
            "type": {
              "defined": "Any2SVMRampMessage"
            }
          },
          {
            "name": "offchainTokenData",
            "type": {
              "vec": "bytes"
            }
          },
          {
            "name": "proofs",
            "type": {
              "vec": {
                "array": [
                  "u8",
                  32
                ]
              }
            }
          }
        ]
      }
    },
    {
      "name": "RampMessageHeader",
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
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Any2SVMRampExtraArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "computeUnits",
            "type": "u32"
          },
          {
            "name": "isWritableBitmap",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "Any2SVMRampMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "header",
            "type": {
              "defined": "RampMessageHeader"
            }
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
            "name": "tokenReceiver",
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "Any2SVMTokenTransfer"
              }
            }
          },
          {
            "name": "extraArgs",
            "type": {
              "defined": "Any2SVMRampExtraArgs"
            }
          }
        ]
      }
    },
    {
      "name": "Any2SVMTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "bytes"
          },
          {
            "name": "destTokenAddress",
            "type": "publicKey"
          },
          {
            "name": "destGasAmount",
            "type": "u32"
          },
          {
            "name": "extraData",
            "type": "bytes"
          },
          {
            "name": "amount",
            "type": {
              "defined": "CrossChainAmount"
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
      "name": "CrossChainAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "leBytes",
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
      "name": "Ocr3ConfigInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "configDigest",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "f",
            "type": "u8"
          },
          {
            "name": "n",
            "type": "u8"
          },
          {
            "name": "isSignatureVerificationEnabled",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3Config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pluginType",
            "type": {
              "defined": "ConfigOcrPluginType"
            }
          },
          {
            "name": "configInfo",
            "type": {
              "defined": "Ocr3ConfigInfo"
            }
          },
          {
            "name": "signers",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    20
                  ]
                },
                16
              ]
            }
          },
          {
            "name": "transmitters",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    32
                  ]
                },
                16
              ]
            }
          }
        ]
      }
    },
    {
      "name": "SourceChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "isRmnVerificationDisabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "onRamp",
            "type": {
              "defined": "OnRampAddress"
            }
          }
        ]
      }
    },
    {
      "name": "OnRampAddress",
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
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SourceChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minSeqNr",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "OcrPluginType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Commit"
          },
          {
            "name": "Execution"
          }
        ]
      }
    },
    {
      "name": "MessageExecutionState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Untouched"
          },
          {
            "name": "InProgress"
          },
          {
            "name": "Success"
          },
          {
            "name": "Failure"
          }
        ]
      }
    },
    {
      "name": "CodeVersion",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Default"
          },
          {
            "name": "V1"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "SourceChainConfigUpdated",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SourceChainAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferRequested",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OwnershipTransferred",
      "fields": [
        {
          "name": "from",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "enableManualExecutionAfter",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "ReferenceAddressesSet",
      "fields": [
        {
          "name": "router",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "feeQuoter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "offrampLookupTable",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "rmnRemote",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "CommitReportAccepted",
      "fields": [
        {
          "name": "merkleRoot",
          "type": {
            "option": {
              "defined": "MerkleRoot"
            }
          },
          "index": false
        },
        {
          "name": "priceUpdates",
          "type": {
            "defined": "PriceUpdates"
          },
          "index": false
        }
      ]
    },
    {
      "name": "CommitReportPDAClosed",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "merkleRoot",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "SkippedAlreadyExecutedMessage",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "ExecutionStateChanged",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
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
          "name": "messageHash",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "state",
          "type": {
            "defined": "MessageExecutionState"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "ocrPluginType",
          "type": {
            "defined": "OcrPluginType"
          },
          "index": false
        },
        {
          "name": "configDigest",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          },
          "index": false
        },
        {
          "name": "transmitters",
          "type": {
            "vec": "publicKey"
          },
          "index": false
        },
        {
          "name": "f",
          "type": "u8",
          "index": false
        }
      ]
    },
    {
      "name": "Transmitted",
      "fields": [
        {
          "name": "ocrPluginType",
          "type": {
            "defined": "OcrPluginType"
          },
          "index": false
        },
        {
          "name": "configDigest",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "InvalidSequenceInterval",
      "msg": "The given sequence interval is invalid"
    },
    {
      "code": 9001,
      "name": "RootNotCommitted",
      "msg": "The given Merkle Root is missing"
    },
    {
      "code": 9002,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 9003,
      "name": "ExistingMerkleRoot",
      "msg": "The given Merkle Root is already committed"
    },
    {
      "code": 9004,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9005,
      "name": "InvalidNonce",
      "msg": "Invalid Nonce"
    },
    {
      "code": 9006,
      "name": "InvalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 9007,
      "name": "OnrampNotConfigured",
      "msg": "Onramp was not configured"
    },
    {
      "code": 9008,
      "name": "FailedToDeserializeReport",
      "msg": "Failed to deserialize report"
    },
    {
      "code": 9009,
      "name": "InvalidPluginType",
      "msg": "Invalid plugin type"
    },
    {
      "code": 9010,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9011,
      "name": "MissingExpectedPriceUpdates",
      "msg": "Commit report is missing expected price updates"
    },
    {
      "code": 9012,
      "name": "MissingExpectedMerkleRoot",
      "msg": "Commit report is missing expected merkle root"
    },
    {
      "code": 9013,
      "name": "UnexpectedMerkleRoot",
      "msg": "Commit report contains unexpected merkle root"
    },
    {
      "code": 9014,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9015,
      "name": "UnsupportedSourceChainSelector",
      "msg": "Source chain selector not supported"
    },
    {
      "code": 9016,
      "name": "UnsupportedDestinationChainSelector",
      "msg": "Destination chain selector not supported"
    },
    {
      "code": 9017,
      "name": "InvalidProof",
      "msg": "Invalid Proof for Merkle Root"
    },
    {
      "code": 9018,
      "name": "InvalidMessage",
      "msg": "Invalid message format"
    },
    {
      "code": 9019,
      "name": "ReachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 9020,
      "name": "ManualExecutionNotAllowed",
      "msg": "Manual execution not allowed"
    },
    {
      "code": 9021,
      "name": "InvalidInputsNumberOfAccounts",
      "msg": "Number of accounts is invalid"
    },
    {
      "code": 9022,
      "name": "InvalidInputsGlobalStateAccount",
      "msg": "Invalid global state account address"
    },
    {
      "code": 9023,
      "name": "InvalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 9024,
      "name": "InvalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 9025,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 9026,
      "name": "InvalidInputsSysvarAccount",
      "msg": "Invalid sysvar instructions account"
    },
    {
      "code": 9027,
      "name": "InvalidInputsFeeQuoterAccount",
      "msg": "Invalid fee quoter account"
    },
    {
      "code": 9028,
      "name": "InvalidInputsAllowedOfframpAccount",
      "msg": "Invalid offramp authorization account"
    },
    {
      "code": 9029,
      "name": "InvalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 9030,
      "name": "InvalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 9031,
      "name": "InvalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 9032,
      "name": "OfframpReleaseMintBalanceMismatch",
      "msg": "Release or mint balance mismatch"
    },
    {
      "code": 9033,
      "name": "OfframpInvalidDataLength",
      "msg": "Invalid data length"
    },
    {
      "code": 9034,
      "name": "StaleCommitReport",
      "msg": "Stale commit report"
    },
    {
      "code": 9035,
      "name": "InvalidWritabilityBitmap",
      "msg": "Invalid writability bitmap"
    },
    {
      "code": 9036,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 9037,
      "name": "Ocr3InvalidConfigFMustBePositive",
      "msg": "Invalid config: F must be positive"
    },
    {
      "code": 9038,
      "name": "Ocr3InvalidConfigTooManyTransmitters",
      "msg": "Invalid config: Too many transmitters"
    },
    {
      "code": 9039,
      "name": "Ocr3InvalidConfigNoTransmitters",
      "msg": "Invalid config: No transmitters"
    },
    {
      "code": 9040,
      "name": "Ocr3InvalidConfigTooManySigners",
      "msg": "Invalid config: Too many signers"
    },
    {
      "code": 9041,
      "name": "Ocr3InvalidConfigFIsTooHigh",
      "msg": "Invalid config: F is too high"
    },
    {
      "code": 9042,
      "name": "Ocr3InvalidConfigRepeatedOracle",
      "msg": "Invalid config: Repeated oracle address"
    },
    {
      "code": 9043,
      "name": "Ocr3WrongMessageLength",
      "msg": "Wrong message length"
    },
    {
      "code": 9044,
      "name": "Ocr3ConfigDigestMismatch",
      "msg": "Config digest mismatch"
    },
    {
      "code": 9045,
      "name": "Ocr3WrongNumberOfSignatures",
      "msg": "Wrong number signatures"
    },
    {
      "code": 9046,
      "name": "Ocr3UnauthorizedTransmitter",
      "msg": "Unauthorized transmitter"
    },
    {
      "code": 9047,
      "name": "Ocr3UnauthorizedSigner",
      "msg": "Unauthorized signer"
    },
    {
      "code": 9048,
      "name": "Ocr3NonUniqueSignatures",
      "msg": "Non unique signatures"
    },
    {
      "code": 9049,
      "name": "Ocr3OracleCannotBeZeroAddress",
      "msg": "Oracle cannot be zero address"
    },
    {
      "code": 9050,
      "name": "Ocr3StaticConfigCannotBeChanged",
      "msg": "Static config cannot be changed"
    },
    {
      "code": 9051,
      "name": "Ocr3InvalidPluginType",
      "msg": "Incorrect plugin type"
    },
    {
      "code": 9052,
      "name": "Ocr3InvalidSignature",
      "msg": "Invalid signature"
    },
    {
      "code": 9053,
      "name": "Ocr3SignaturesOutOfRegistration",
      "msg": "Signatures out of registration"
    },
    {
      "code": 9054,
      "name": "InvalidOnrampAddress",
      "msg": "Invalid onramp address"
    },
    {
      "code": 9055,
      "name": "InvalidInputsExternalExecutionSignerAccount",
      "msg": "Invalid external execution signer account"
    },
    {
      "code": 9056,
      "name": "CommitReportHasPendingMessages",
      "msg": "Commit report has pending messages"
    }
  ]
};
