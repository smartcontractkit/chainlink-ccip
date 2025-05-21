/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/ccip_offramp.json`.
 */
export type CcipOfframp = {
  "address": "offqSMQWgQud6WJz694LRzkeN5kMYpCHTpXQr3Rkcjm",
  "metadata": {
    "name": "ccipOfframp",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
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
          "name": "config",
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
      "name": "addSourceChain",
      "docs": [
        "Adds a new source chain selector with its config to the offramp.",
        "",
        "The Admin needs to add any new chain supported.",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "",
        "# Arguments"
      ],
      "discriminator": [
        26,
        58,
        148,
        88,
        190,
        27,
        2,
        144
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "docs": [
            "Adding a chain selector implies initializing the state for a new chain"
          ],
          "writable": true
        },
        {
          "name": "config"
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
          "name": "newChainSelector",
          "type": "u64"
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": {
              "name": "sourceChainConfig"
            }
          }
        }
      ]
    },
    {
      "name": "closeCommitReportAccount",
      "discriminator": [
        109,
        145,
        129,
        64,
        226,
        172,
        61,
        106
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "commitReport",
          "writable": true
        },
        {
          "name": "referenceAddresses"
        },
        {
          "name": "wsolMint"
        },
        {
          "name": "feeTokenReceiver",
          "writable": true
        },
        {
          "name": "feeBillingSigner"
        },
        {
          "name": "tokenProgram"
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
      "discriminator": [
        223,
        140,
        142,
        165,
        229,
        208,
        156,
        74
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "referenceAddresses"
        },
        {
          "name": "sourceChain",
          "writable": true
        },
        {
          "name": "commitReport",
          "writable": true
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
          "name": "sysvarInstructions"
        },
        {
          "name": "feeBillingSigner"
        },
        {
          "name": "feeQuoter"
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig"
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
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
      "discriminator": [
        186,
        145,
        195,
        227,
        207,
        211,
        226,
        134
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "referenceAddresses"
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
          "name": "sysvarInstructions"
        },
        {
          "name": "feeBillingSigner"
        },
        {
          "name": "feeQuoter"
        },
        {
          "name": "feeQuoterAllowedPriceUpdater",
          "docs": [
            "so that it can authorize the call made by this offramp"
          ]
        },
        {
          "name": "feeQuoterConfig"
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
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
      "discriminator": [
        58,
        101,
        54,
        252,
        248,
        31,
        226,
        121
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
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
      "discriminator": [
        130,
        221,
        242,
        154,
        13,
        193,
        189,
        29
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "referenceAddresses"
        },
        {
          "name": "sourceChain"
        },
        {
          "name": "commitReport",
          "writable": true
        },
        {
          "name": "offramp"
        },
        {
          "name": "allowedOfframp",
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions"
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
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
          "name": "referenceAddresses",
          "writable": true
        },
        {
          "name": "router"
        },
        {
          "name": "feeQuoter"
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "offrampLookupTable"
        },
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
        },
        {
          "name": "program"
        },
        {
          "name": "programData"
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
      "discriminator": [
        238,
        219,
        224,
        11,
        226,
        248,
        47,
        192
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "referenceAddresses"
        },
        {
          "name": "sourceChain"
        },
        {
          "name": "commitReport",
          "writable": true
        },
        {
          "name": "offramp"
        },
        {
          "name": "allowedOfframp",
          "docs": [
            "CHECK PDA of the router program verifying the signer is an allowed offramp.",
            "If PDA does not exist, the router doesn't allow this offramp. This is just used",
            "so that token pools and receivers can then check that the caller is an actual offramp that",
            "has been registered in the router as such for that source chain."
          ]
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
          "name": "sysvarInstructions"
        },
        {
          "name": "rmnRemote"
        },
        {
          "name": "rmnRemoteCurses"
        },
        {
          "name": "rmnRemoteConfig"
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
      "discriminator": [
        47,
        151,
        233,
        254,
        121,
        82,
        206,
        152
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "codeVersion",
          "type": {
            "defined": {
              "name": "codeVersion"
            }
          }
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
      "discriminator": [
        4,
        131,
        107,
        110,
        250,
        158,
        244,
        200
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
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
          "name": "pluginType",
          "type": {
            "defined": {
              "name": "ocrPluginType"
            }
          }
        },
        {
          "name": "configInfo",
          "type": {
            "defined": {
              "name": "ocr3ConfigInfo"
            }
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
            "vec": "pubkey"
          }
        }
      ]
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
          "name": "config",
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
      "discriminator": [
        157,
        236,
        73,
        92,
        84,
        197,
        152,
        105
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
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
      "discriminator": [
        119,
        179,
        218,
        249,
        217,
        184,
        181,
        9
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "referenceAddresses",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "router",
          "type": "pubkey"
        },
        {
          "name": "feeQuoter",
          "type": "pubkey"
        },
        {
          "name": "offrampLookupTable",
          "type": "pubkey"
        },
        {
          "name": "rmnRemote",
          "type": "pubkey"
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
      "discriminator": [
        52,
        85,
        37,
        124,
        209,
        140,
        181,
        104
      ],
      "accounts": [
        {
          "name": "sourceChain",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
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
            "defined": {
              "name": "sourceChainConfig"
            }
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
      "discriminator": [
        164,
        212,
        71,
        101,
        166,
        113,
        26,
        93
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "newChainSelector",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "commitReport",
      "discriminator": [
        46,
        231,
        247,
        231,
        174,
        68,
        34,
        26
      ]
    },
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
      "name": "globalState",
      "discriminator": [
        163,
        46,
        74,
        168,
        216,
        123,
        133,
        98
      ]
    },
    {
      "name": "referenceAddresses",
      "discriminator": [
        99,
        5,
        216,
        212,
        250,
        75,
        74,
        12
      ]
    },
    {
      "name": "sourceChain",
      "discriminator": [
        242,
        235,
        220,
        98,
        252,
        121,
        191,
        216
      ]
    }
  ],
  "events": [
    {
      "name": "commitReportAccepted",
      "discriminator": [
        44,
        46,
        77,
        237,
        70,
        187,
        170,
        133
      ]
    },
    {
      "name": "commitReportPdaClosed",
      "discriminator": [
        69,
        240,
        72,
        149,
        174,
        18,
        236,
        46
      ]
    },
    {
      "name": "configSet",
      "discriminator": [
        15,
        104,
        59,
        16,
        236,
        241,
        8,
        6
      ]
    },
    {
      "name": "executionStateChanged",
      "discriminator": [
        185,
        176,
        140,
        112,
        239,
        78,
        31,
        249
      ]
    },
    {
      "name": "ocrConfigSet",
      "discriminator": [
        136,
        69,
        4,
        187,
        69,
        241,
        211,
        31
      ]
    },
    {
      "name": "ownershipTransferRequested",
      "discriminator": [
        79,
        54,
        99,
        123,
        57,
        244,
        134,
        35
      ]
    },
    {
      "name": "ownershipTransferred",
      "discriminator": [
        172,
        61,
        205,
        183,
        250,
        50,
        38,
        98
      ]
    },
    {
      "name": "referenceAddressesSet",
      "discriminator": [
        146,
        234,
        139,
        115,
        99,
        143,
        216,
        191
      ]
    },
    {
      "name": "skippedAlreadyExecutedMessage",
      "discriminator": [
        124,
        136,
        216,
        231,
        25,
        232,
        5,
        239
      ]
    },
    {
      "name": "sourceChainAdded",
      "discriminator": [
        98,
        127,
        170,
        88,
        67,
        55,
        230,
        8
      ]
    },
    {
      "name": "sourceChainConfigUpdated",
      "discriminator": [
        31,
        205,
        106,
        132,
        10,
        220,
        181,
        30
      ]
    },
    {
      "name": "transmitted",
      "discriminator": [
        144,
        94,
        142,
        170,
        49,
        110,
        67,
        189
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "invalidSequenceInterval",
      "msg": "The given sequence interval is invalid"
    },
    {
      "code": 9001,
      "name": "rootNotCommitted",
      "msg": "The given Merkle Root is missing"
    },
    {
      "code": 9002,
      "name": "invalidRmnRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 9003,
      "name": "existingMerkleRoot",
      "msg": "The given Merkle Root is already committed"
    },
    {
      "code": 9004,
      "name": "unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9005,
      "name": "invalidNonce",
      "msg": "Invalid Nonce"
    },
    {
      "code": 9006,
      "name": "invalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 9007,
      "name": "onrampNotConfigured",
      "msg": "Onramp was not configured"
    },
    {
      "code": 9008,
      "name": "failedToDeserializeReport",
      "msg": "Failed to deserialize report"
    },
    {
      "code": 9009,
      "name": "invalidPluginType",
      "msg": "Invalid plugin type"
    },
    {
      "code": 9010,
      "name": "invalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9011,
      "name": "missingExpectedPriceUpdates",
      "msg": "Commit report is missing expected price updates"
    },
    {
      "code": 9012,
      "name": "missingExpectedMerkleRoot",
      "msg": "Commit report is missing expected merkle root"
    },
    {
      "code": 9013,
      "name": "unexpectedMerkleRoot",
      "msg": "Commit report contains unexpected merkle root"
    },
    {
      "code": 9014,
      "name": "redundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9015,
      "name": "unsupportedSourceChainSelector",
      "msg": "Source chain selector not supported"
    },
    {
      "code": 9016,
      "name": "unsupportedDestinationChainSelector",
      "msg": "Destination chain selector not supported"
    },
    {
      "code": 9017,
      "name": "invalidProof",
      "msg": "Invalid Proof for Merkle Root"
    },
    {
      "code": 9018,
      "name": "invalidMessage",
      "msg": "Invalid message format"
    },
    {
      "code": 9019,
      "name": "reachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 9020,
      "name": "manualExecutionNotAllowed",
      "msg": "Manual execution not allowed"
    },
    {
      "code": 9021,
      "name": "invalidInputsNumberOfAccounts",
      "msg": "Number of accounts is invalid"
    },
    {
      "code": 9022,
      "name": "invalidInputsGlobalStateAccount",
      "msg": "Invalid global state account address"
    },
    {
      "code": 9023,
      "name": "invalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 9024,
      "name": "invalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 9025,
      "name": "invalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 9026,
      "name": "invalidInputsSysvarAccount",
      "msg": "Invalid sysvar instructions account"
    },
    {
      "code": 9027,
      "name": "invalidInputsFeeQuoterAccount",
      "msg": "Invalid fee quoter account"
    },
    {
      "code": 9028,
      "name": "invalidInputsAllowedOfframpAccount",
      "msg": "Invalid offramp authorization account"
    },
    {
      "code": 9029,
      "name": "invalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 9030,
      "name": "invalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 9031,
      "name": "invalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 9032,
      "name": "offrampReleaseMintBalanceMismatch",
      "msg": "Release or mint balance mismatch"
    },
    {
      "code": 9033,
      "name": "offrampInvalidDataLength",
      "msg": "Invalid data length"
    },
    {
      "code": 9034,
      "name": "staleCommitReport",
      "msg": "Stale commit report"
    },
    {
      "code": 9035,
      "name": "invalidWritabilityBitmap",
      "msg": "Invalid writability bitmap"
    },
    {
      "code": 9036,
      "name": "invalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 9037,
      "name": "ocr3InvalidConfigFMustBePositive",
      "msg": "Invalid config: F must be positive"
    },
    {
      "code": 9038,
      "name": "ocr3InvalidConfigTooManyTransmitters",
      "msg": "Invalid config: Too many transmitters"
    },
    {
      "code": 9039,
      "name": "ocr3InvalidConfigNoTransmitters",
      "msg": "Invalid config: No transmitters"
    },
    {
      "code": 9040,
      "name": "ocr3InvalidConfigTooManySigners",
      "msg": "Invalid config: Too many signers"
    },
    {
      "code": 9041,
      "name": "ocr3InvalidConfigFIsTooHigh",
      "msg": "Invalid config: F is too high"
    },
    {
      "code": 9042,
      "name": "ocr3InvalidConfigRepeatedOracle",
      "msg": "Invalid config: Repeated oracle address"
    },
    {
      "code": 9043,
      "name": "ocr3WrongMessageLength",
      "msg": "Wrong message length"
    },
    {
      "code": 9044,
      "name": "ocr3ConfigDigestMismatch",
      "msg": "Config digest mismatch"
    },
    {
      "code": 9045,
      "name": "ocr3WrongNumberOfSignatures",
      "msg": "Wrong number signatures"
    },
    {
      "code": 9046,
      "name": "ocr3UnauthorizedTransmitter",
      "msg": "Unauthorized transmitter"
    },
    {
      "code": 9047,
      "name": "ocr3UnauthorizedSigner",
      "msg": "Unauthorized signer"
    },
    {
      "code": 9048,
      "name": "ocr3NonUniqueSignatures",
      "msg": "Non unique signatures"
    },
    {
      "code": 9049,
      "name": "ocr3OracleCannotBeZeroAddress",
      "msg": "Oracle cannot be zero address"
    },
    {
      "code": 9050,
      "name": "ocr3StaticConfigCannotBeChanged",
      "msg": "Static config cannot be changed"
    },
    {
      "code": 9051,
      "name": "ocr3InvalidPluginType",
      "msg": "Incorrect plugin type"
    },
    {
      "code": 9052,
      "name": "ocr3InvalidSignature",
      "msg": "Invalid signature"
    },
    {
      "code": 9053,
      "name": "ocr3SignaturesOutOfRegistration",
      "msg": "Signatures out of registration"
    },
    {
      "code": 9054,
      "name": "invalidOnrampAddress",
      "msg": "Invalid onramp address"
    },
    {
      "code": 9055,
      "name": "invalidInputsExternalExecutionSignerAccount",
      "msg": "Invalid external execution signer account"
    },
    {
      "code": 9056,
      "name": "commitReportHasPendingMessages",
      "msg": "Commit report has pending messages"
    }
  ],
  "types": [
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
    },
    {
      "name": "commitReportAccepted",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "merkleRoot",
            "type": {
              "option": {
                "defined": {
                  "name": "merkleRoot"
                }
              }
            }
          },
          {
            "name": "priceUpdates",
            "type": {
              "defined": {
                "name": "priceUpdates"
              }
            }
          }
        ]
      }
    },
    {
      "name": "commitReportPdaClosed",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
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
      "name": "config",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
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
            "type": "pubkey"
          },
          {
            "name": "proposedOwner",
            "type": "pubkey"
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
                  "defined": {
                    "name": "ocr3Config"
                  }
                },
                2
              ]
            }
          }
        ]
      }
    },
    {
      "name": "configOcrPluginType",
      "docs": [
        "It's not possible to store enums in zero_copy accounts, so we wrap the discriminant",
        "in a struct to store in config."
      ],
      "repr": {
        "kind": "c"
      },
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
      "name": "configSet",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "enableManualExecutionAfter",
            "type": "i64"
          }
        ]
      }
    },
    {
      "name": "executionStateChanged",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
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
            "name": "messageHash",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "state",
            "type": {
              "defined": {
                "name": "messageExecutionState"
              }
            }
          }
        ]
      }
    },
    {
      "name": "gasPriceUpdate",
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
      "name": "merkleRoot",
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
      "name": "messageExecutionState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "untouched"
          },
          {
            "name": "inProgress"
          },
          {
            "name": "success"
          },
          {
            "name": "failure"
          }
        ]
      }
    },
    {
      "name": "ocr3Config",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pluginType",
            "type": {
              "defined": {
                "name": "configOcrPluginType"
              }
            }
          },
          {
            "name": "configInfo",
            "type": {
              "defined": {
                "name": "ocr3ConfigInfo"
              }
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
      "name": "ocr3ConfigInfo",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
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
      "name": "ocrConfigSet",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "ocrPluginType",
            "type": {
              "defined": {
                "name": "ocrPluginType"
              }
            }
          },
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
              "vec": "pubkey"
            }
          },
          {
            "name": "f",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "ocrPluginType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "commit"
          },
          {
            "name": "execution"
          }
        ]
      }
    },
    {
      "name": "onRampAddress",
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
      "name": "ownershipTransferRequested",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "from",
            "type": "pubkey"
          },
          {
            "name": "to",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "ownershipTransferred",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "from",
            "type": "pubkey"
          },
          {
            "name": "to",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "priceUpdates",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenPriceUpdates",
            "type": {
              "vec": {
                "defined": {
                  "name": "tokenPriceUpdate"
                }
              }
            }
          },
          {
            "name": "gasPriceUpdates",
            "type": {
              "vec": {
                "defined": {
                  "name": "gasPriceUpdate"
                }
              }
            }
          }
        ]
      }
    },
    {
      "name": "referenceAddresses",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "router",
            "type": "pubkey"
          },
          {
            "name": "feeQuoter",
            "type": "pubkey"
          },
          {
            "name": "offrampLookupTable",
            "type": "pubkey"
          },
          {
            "name": "rmnRemote",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "referenceAddressesSet",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "router",
            "type": "pubkey"
          },
          {
            "name": "feeQuoter",
            "type": "pubkey"
          },
          {
            "name": "offrampLookupTable",
            "type": "pubkey"
          },
          {
            "name": "rmnRemote",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "skippedAlreadyExecutedMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
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
              "defined": {
                "name": "sourceChainState"
              }
            }
          },
          {
            "name": "config",
            "type": {
              "defined": {
                "name": "sourceChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "sourceChainAdded",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sourceChainConfig",
            "type": {
              "defined": {
                "name": "sourceChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "sourceChainConfig",
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
              "defined": {
                "name": "codeVersion"
              }
            }
          },
          {
            "name": "onRamp",
            "type": {
              "defined": {
                "name": "onRampAddress"
              }
            }
          }
        ]
      }
    },
    {
      "name": "sourceChainConfigUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sourceChainConfig",
            "type": {
              "defined": {
                "name": "sourceChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "sourceChainState",
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
      "name": "tokenPriceUpdate",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceToken",
            "type": "pubkey"
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
      "name": "transmitted",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "ocrPluginType",
            "type": {
              "defined": {
                "name": "ocrPluginType"
              }
            }
          },
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
            "name": "sequenceNumber",
            "type": "u64"
          }
        ]
      }
    }
  ]
};
