export type FeeQuoter = {
  "version": "0.1.0-dev",
  "name": "fee_quoter",
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initializes the Fee Quoter.",
        "",
        "The initialization is responsibility of Admin, nothing more than calling this method should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `default_gas_limit` - The default gas limit for other destination chains.",
        "* `default_allow_out_of_order_execution` - Whether out-of-order execution is allowed by default for other destination chains.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "linkTokenMint",
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
          "name": "maxFeeJuelsPerMsg",
          "type": "u128"
        },
        {
          "name": "onramp",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "transferOwnership",
      "docs": [
        "Transfers the ownership of the fee quoter to a new proposed owner.",
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
          "name": "newOwner",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "docs": [
        "Accepts the ownership of the fee quoter by the proposed owner.",
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
      "name": "addBillingTokenConfig",
      "docs": [
        "Adds a billing token configuration.",
        "Only CCIP Admin can add a billing token configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding the billing token configuration.",
        "* `config` - The billing token configuration to be added."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
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
          "name": "config",
          "type": {
            "defined": "BillingTokenConfig"
          }
        }
      ]
    },
    {
      "name": "updateBillingTokenConfig",
      "docs": [
        "Updates the billing token configuration.",
        "Only CCIP Admin can update a billing token configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the billing token configuration.",
        "* `config` - The new billing token configuration."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
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
          "name": "config",
          "type": {
            "defined": "BillingTokenConfig"
          }
        }
      ]
    },
    {
      "name": "addDestChain",
      "docs": [
        "Adds a new destination chain selector to the fee quoter.",
        "",
        "The Admin needs to add any new chain supported.",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding the chain selector.",
        "* `chain_selector` - The new chain selector to be added.",
        "* `dest_chain_config` - The configuration for the chain as destination."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          }
        }
      ]
    },
    {
      "name": "disableDestChain",
      "docs": [
        "Disables the destination chain selector.",
        "",
        "The Admin is the only one able to disable the chain selector as destination. This method is thought of as an emergency kill-switch.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for disabling the chain selector.",
        "* `chain_selector` - The destination chain selector to be disabled."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "chainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateDestChainConfig",
      "docs": [
        "Updates the configuration of the destination chain selector.",
        "",
        "The Admin is the only one able to update the destination chain config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the chain selector.",
        "* `chain_selector` - The destination chain selector to be updated.",
        "* `dest_chain_config` - The new configuration for the destination chain."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          }
        }
      ]
    },
    {
      "name": "setTokenTransferFeeConfig",
      "docs": [
        "Sets the token transfer fee configuration for a particular token when it's transferred to a particular dest chain.",
        "It is an upsert, initializing the per-chain-per-token config account if it doesn't exist",
        "and overwriting it if it does.",
        "",
        "Only the Admin can perform this operation.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for setting the token billing configuration.",
        "* `chain_selector` - The chain selector.",
        "* `mint` - The public key of the token mint.",
        "* `cfg` - The token transfer fee configuration."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "perChainPerTokenConfig",
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
          "name": "mint",
          "type": "publicKey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": "TokenTransferFeeConfig"
          }
        }
      ]
    },
    {
      "name": "addPriceUpdater",
      "docs": [
        "Add a price updater address to the list of allowed price updaters.",
        "On price updates, the fee quoter will check the that caller is allowed.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `price_updater` - The price updater address."
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "priceUpdater",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "removePriceUpdater",
      "docs": [
        "Remove a price updater address from the list of allowed price updaters.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `price_updater` - The price updater address."
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "priceUpdater",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "getFee",
      "docs": [
        "Calculates the fee for sending a message to the destination chain.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the fee calculation.",
        "* `dest_chain_selector` - The chain selector for the destination chain.",
        "* `message` - The message to be sent.",
        "",
        "# Additional accounts",
        "",
        "In addition to the fixed amount of accounts defined in the `GetFee` context,",
        "the following accounts must be provided:",
        "",
        "* First, the billing token config accounts for each token sent with the message, sequentially.",
        "For each token with no billing config account (i.e. tokens that cannot be possibly used as fee",
        "tokens, which also have no BPS fees enabled) the ZERO address must be provided instead.",
        "* Then, the per chain / per token config of every token sent with the message, sequentially",
        "in the same order.",
        "",
        "# Returns",
        "",
        "GetFeeResult struct with:",
        "- the fee token mint address,",
        "- the fee amount of said token,",
        "- the fee value in juels,",
        "- additional data required when performing the cross-chain transfer of tokens in that message",
        "- deserialized and processed extra args"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "linkTokenConfig",
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
          "name": "message",
          "type": {
            "defined": "SVM2AnyMessage"
          }
        }
      ],
      "returns": {
        "defined": "GetFeeResult"
      }
    },
    {
      "name": "updatePrices",
      "docs": [
        "Updates prices for tokens and gas. This method may only be called by an allowed price updater.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts always required for the price updates",
        "* `token_updates` - Vector of token price updates",
        "* `gas_updates` - Vector of gas price updates",
        "",
        "# Additional accounts",
        "",
        "In addition to the fixed amount of accounts defined in the `UpdatePrices` context,",
        "the following accounts must be provided:",
        "",
        "* First, the billing token config accounts for each token whose price is being updated, in the same order",
        "as the token_updates vector.",
        "* Then, the dest chain accounts of every chain whose gas price is being updated, in the same order as the",
        "gas_updates vector."
      ],
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "allowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "was added by the owner as an allowed price updater. The constraints enforced guarantee that it is the right PDA",
            "and that it was initialized."
          ]
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "tokenUpdates",
          "type": {
            "vec": {
              "defined": "TokenPriceUpdate"
            }
          }
        },
        {
          "name": "gasUpdates",
          "type": {
            "vec": {
              "defined": "GasPriceUpdate"
            }
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
            "name": "version",
            "type": "u8"
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
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "publicKey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          }
        ]
      }
    },
    {
      "name": "destChain",
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
              "defined": "DestChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "DestChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "billingTokenConfigWrapper",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "config",
            "type": {
              "defined": "BillingTokenConfig"
            }
          }
        ]
      }
    },
    {
      "name": "perChainPerTokenConfig",
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
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "tokenTransferConfig",
            "type": {
              "defined": "TokenTransferFeeConfig"
            }
          }
        ]
      }
    },
    {
      "name": "allowedPriceUpdater",
      "type": {
        "kind": "struct",
        "fields": []
      }
    }
  ],
  "types": [
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
      "name": "GenericExtraArgsV2",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "gasLimit",
            "type": "u128"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "SVMExtraArgsV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "computeUnits",
            "type": "u32"
          },
          {
            "name": "accountIsWritableBitmap",
            "type": "u64"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          },
          {
            "name": "tokenReceiver",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "accounts",
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
      "name": "SVM2AnyMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "receiver",
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
          },
          {
            "name": "feeToken",
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
      "name": "TokenTransferAdditionalData",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destBytesOverhead",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "GetFeeResult",
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
          },
          {
            "name": "juels",
            "type": "u128"
          },
          {
            "name": "tokenTransferAdditionalData",
            "type": {
              "vec": {
                "defined": "TokenTransferAdditionalData"
              }
            }
          },
          {
            "name": "processedExtraArgs",
            "type": {
              "defined": "ProcessedExtraArgs"
            }
          }
        ]
      }
    },
    {
      "name": "ProcessedExtraArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "bytes",
            "type": "bytes"
          },
          {
            "name": "gasLimit",
            "type": "u128"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          },
          {
            "name": "tokenReceiver",
            "type": {
              "option": "bytes"
            }
          }
        ]
      }
    },
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "maxNumberOfTokensPerMsg",
            "type": "u16"
          },
          {
            "name": "maxDataBytes",
            "type": "u32"
          },
          {
            "name": "maxPerMsgGasLimit",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteBase",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteHigh",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteThreshold",
            "type": "u32"
          },
          {
            "name": "destDataAvailabilityOverheadGas",
            "type": "u32"
          },
          {
            "name": "destGasPerDataAvailabilityByte",
            "type": "u16"
          },
          {
            "name": "destDataAvailabilityMultiplierBps",
            "type": "u16"
          },
          {
            "name": "defaultTokenFeeUsdcents",
            "type": "u16"
          },
          {
            "name": "defaultTokenDestGasOverhead",
            "type": "u32"
          },
          {
            "name": "defaultTxGasLimit",
            "type": "u32"
          },
          {
            "name": "gasMultiplierWeiPerEth",
            "type": "u64"
          },
          {
            "name": "networkFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "gasPriceStalenessThreshold",
            "type": "u32"
          },
          {
            "name": "enforceOutOfOrder",
            "type": "bool"
          },
          {
            "name": "chainFamilySelector",
            "type": {
              "array": [
                "u8",
                4
              ]
            }
          }
        ]
      }
    },
    {
      "name": "DestChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "usdPerUnitGas",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          }
        ]
      }
    },
    {
      "name": "TimestampedPackedU224",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          }
        ]
      }
    },
    {
      "name": "BillingTokenConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "enabled",
            "type": "bool"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          },
          {
            "name": "premiumMultiplierWeiPerEth",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "TokenTransferFeeConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "maxFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "deciBps",
            "type": "u16"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destBytesOverhead",
            "type": "u32"
          },
          {
            "name": "isEnabled",
            "type": "bool"
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
      "name": "ConfigSet",
      "fields": [
        {
          "name": "maxFeeJuelsPerMsg",
          "type": "u128",
          "index": false
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "linkTokenLocalDecimals",
          "type": "u8",
          "index": false
        },
        {
          "name": "onramp",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "defaultCodeVersion",
          "type": {
            "defined": "CodeVersion"
          },
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenAdded",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "enabled",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenEnabled",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenDisabled",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenRemoved",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "DestChainAdded",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "DestChainConfigUpdated",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
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
      "name": "UsdPerUnitGasUpdated",
      "fields": [
        {
          "name": "destChain",
          "type": "u64",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "UsdPerTokenUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "TokenPriceUpdateIgnored",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "TokenTransferFeeConfigUpdated",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "tokenTransferFeeConfig",
          "type": {
            "defined": "TokenTransferFeeConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "PremiumMultiplierWeiPerEthUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "premiumMultiplierWeiPerEth",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "PriceUpdaterAdded",
      "fields": [
        {
          "name": "priceUpdater",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "PriceUpdaterRemoved",
      "fields": [
        {
          "name": "priceUpdater",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 8000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 8001,
      "name": "InvalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 8002,
      "name": "ZeroGasLimit",
      "msg": "Gas limit is zero"
    },
    {
      "code": 8003,
      "name": "DefaultGasLimitExceedsMaximum",
      "msg": "Default gas limit exceeds the maximum"
    },
    {
      "code": 8004,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 8005,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 8006,
      "name": "InvalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 8007,
      "name": "InvalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 8008,
      "name": "InvalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 8009,
      "name": "InvalidInputsMintOwner",
      "msg": "Mint account input has an invalid owner"
    },
    {
      "code": 8010,
      "name": "InvalidInputsTokenConfigAccount",
      "msg": "Token config account is invalid"
    },
    {
      "code": 8011,
      "name": "InvalidInputsMissingExtraArgs",
      "msg": "Missing extra args in message to SVM receiver"
    },
    {
      "code": 8012,
      "name": "InvalidInputsMissingDataAfterExtraArgs",
      "msg": "Missing data after extra args tag"
    },
    {
      "code": 8013,
      "name": "InvalidInputsDestChainStateAccount",
      "msg": "Destination chain state account is invalid"
    },
    {
      "code": 8014,
      "name": "InvalidInputsPerChainPerTokenConfig",
      "msg": "Per chain per token config account is invalid"
    },
    {
      "code": 8015,
      "name": "InvalidInputsBillingTokenConfig",
      "msg": "Billing token config account is invalid"
    },
    {
      "code": 8016,
      "name": "InvalidInputsAccountCount",
      "msg": "Number of accounts provided is incorrect"
    },
    {
      "code": 8017,
      "name": "InvalidInputsNoUpdates",
      "msg": "No price or gas update provided"
    },
    {
      "code": 8018,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 8019,
      "name": "DestinationChainDisabled",
      "msg": "Destination chain disabled"
    },
    {
      "code": 8020,
      "name": "FeeTokenDisabled",
      "msg": "Fee token disabled"
    },
    {
      "code": 8021,
      "name": "MessageTooLarge",
      "msg": "Message exceeds maximum data size"
    },
    {
      "code": 8022,
      "name": "UnsupportedNumberOfTokens",
      "msg": "Message contains an unsupported number of tokens"
    },
    {
      "code": 8023,
      "name": "InvalidEVMAddress",
      "msg": "Invalid EVM address"
    },
    {
      "code": 8024,
      "name": "InvalidEncoding",
      "msg": "Invalid encoding"
    },
    {
      "code": 8025,
      "name": "InvalidTokenPrice",
      "msg": "Invalid token price"
    },
    {
      "code": 8026,
      "name": "StaleGasPrice",
      "msg": "Stale gas price"
    },
    {
      "code": 8027,
      "name": "InvalidInputsMissingTokenConfig",
      "msg": "Inputs are missing token configuration"
    },
    {
      "code": 8028,
      "name": "MessageFeeTooHigh",
      "msg": "Message fee is too high"
    },
    {
      "code": 8029,
      "name": "MessageGasLimitTooHigh",
      "msg": "Message gas limit too high"
    },
    {
      "code": 8030,
      "name": "ExtraArgOutOfOrderExecutionMustBeTrue",
      "msg": "Extra arg out of order execution must be true"
    },
    {
      "code": 8031,
      "name": "InvalidExtraArgsTag",
      "msg": "Invalid extra args tag"
    },
    {
      "code": 8032,
      "name": "InvalidExtraArgsAccounts",
      "msg": "Invalid amount of accounts in extra args"
    },
    {
      "code": 8033,
      "name": "InvalidExtraArgsWritabilityBitmap",
      "msg": "Invalid writability bitmap in extra args"
    },
    {
      "code": 8034,
      "name": "InvalidChainFamilySelector",
      "msg": "Invalid chain family selector"
    },
    {
      "code": 8035,
      "name": "InvalidTokenReceiver",
      "msg": "Invalid token receiver"
    },
    {
      "code": 8036,
      "name": "InvalidSVMAddress",
      "msg": "Invalid SVM address"
    },
    {
      "code": 8037,
      "name": "UnauthorizedPriceUpdater",
      "msg": "The caller is not an authorized price updater"
    },
    {
      "code": 8038,
      "name": "InvalidLinkDecimals",
      "msg": "The LINK mint uses an unsupported number of decimals"
    },
    {
      "code": 8039,
      "name": "InvalidTokenTransferFeeMaxMin",
      "msg": "Minimum token transfer fee exceeds maximum"
    },
    {
      "code": 8040,
      "name": "InvalidTokenTransferFeeDestBytesOverhead",
      "msg": "Insufficient dest bytes overhead on transfer fee config"
    },
    {
      "code": 8041,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    }
  ]
};

export const IDL: FeeQuoter = {
  "version": "0.1.0-dev",
  "name": "fee_quoter",
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initializes the Fee Quoter.",
        "",
        "The initialization is responsibility of Admin, nothing more than calling this method should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `default_gas_limit` - The default gas limit for other destination chains.",
        "* `default_allow_out_of_order_execution` - Whether out-of-order execution is allowed by default for other destination chains.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "linkTokenMint",
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
          "name": "maxFeeJuelsPerMsg",
          "type": "u128"
        },
        {
          "name": "onramp",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "transferOwnership",
      "docs": [
        "Transfers the ownership of the fee quoter to a new proposed owner.",
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
          "name": "newOwner",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "docs": [
        "Accepts the ownership of the fee quoter by the proposed owner.",
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
      "name": "addBillingTokenConfig",
      "docs": [
        "Adds a billing token configuration.",
        "Only CCIP Admin can add a billing token configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding the billing token configuration.",
        "* `config` - The billing token configuration to be added."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenReceiver",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
          "isSigner": false
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
          "name": "config",
          "type": {
            "defined": "BillingTokenConfig"
          }
        }
      ]
    },
    {
      "name": "updateBillingTokenConfig",
      "docs": [
        "Updates the billing token configuration.",
        "Only CCIP Admin can update a billing token configuration.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the billing token configuration.",
        "* `config` - The new billing token configuration."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
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
          "name": "config",
          "type": {
            "defined": "BillingTokenConfig"
          }
        }
      ]
    },
    {
      "name": "addDestChain",
      "docs": [
        "Adds a new destination chain selector to the fee quoter.",
        "",
        "The Admin needs to add any new chain supported.",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding the chain selector.",
        "* `chain_selector` - The new chain selector to be added.",
        "* `dest_chain_config` - The configuration for the chain as destination."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          }
        }
      ]
    },
    {
      "name": "disableDestChain",
      "docs": [
        "Disables the destination chain selector.",
        "",
        "The Admin is the only one able to disable the chain selector as destination. This method is thought of as an emergency kill-switch.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for disabling the chain selector.",
        "* `chain_selector` - The destination chain selector to be disabled."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "chainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "updateDestChainConfig",
      "docs": [
        "Updates the configuration of the destination chain selector.",
        "",
        "The Admin is the only one able to update the destination chain config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the chain selector.",
        "* `chain_selector` - The destination chain selector to be updated.",
        "* `dest_chain_config` - The new configuration for the destination chain."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
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
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          }
        }
      ]
    },
    {
      "name": "setTokenTransferFeeConfig",
      "docs": [
        "Sets the token transfer fee configuration for a particular token when it's transferred to a particular dest chain.",
        "It is an upsert, initializing the per-chain-per-token config account if it doesn't exist",
        "and overwriting it if it does.",
        "",
        "Only the Admin can perform this operation.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for setting the token billing configuration.",
        "* `chain_selector` - The chain selector.",
        "* `mint` - The public key of the token mint.",
        "* `cfg` - The token transfer fee configuration."
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "perChainPerTokenConfig",
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
          "name": "mint",
          "type": "publicKey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": "TokenTransferFeeConfig"
          }
        }
      ]
    },
    {
      "name": "addPriceUpdater",
      "docs": [
        "Add a price updater address to the list of allowed price updaters.",
        "On price updates, the fee quoter will check the that caller is allowed.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `price_updater` - The price updater address."
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "priceUpdater",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "removePriceUpdater",
      "docs": [
        "Remove a price updater address from the list of allowed price updaters.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `price_updater` - The price updater address."
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "priceUpdater",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "getFee",
      "docs": [
        "Calculates the fee for sending a message to the destination chain.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the fee calculation.",
        "* `dest_chain_selector` - The chain selector for the destination chain.",
        "* `message` - The message to be sent.",
        "",
        "# Additional accounts",
        "",
        "In addition to the fixed amount of accounts defined in the `GetFee` context,",
        "the following accounts must be provided:",
        "",
        "* First, the billing token config accounts for each token sent with the message, sequentially.",
        "For each token with no billing config account (i.e. tokens that cannot be possibly used as fee",
        "tokens, which also have no BPS fees enabled) the ZERO address must be provided instead.",
        "* Then, the per chain / per token config of every token sent with the message, sequentially",
        "in the same order.",
        "",
        "# Returns",
        "",
        "GetFeeResult struct with:",
        "- the fee token mint address,",
        "- the fee amount of said token,",
        "- the fee value in juels,",
        "- additional data required when performing the cross-chain transfer of tokens in that message",
        "- deserialized and processed extra args"
      ],
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "billingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "linkTokenConfig",
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
          "name": "message",
          "type": {
            "defined": "SVM2AnyMessage"
          }
        }
      ],
      "returns": {
        "defined": "GetFeeResult"
      }
    },
    {
      "name": "updatePrices",
      "docs": [
        "Updates prices for tokens and gas. This method may only be called by an allowed price updater.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts always required for the price updates",
        "* `token_updates` - Vector of token price updates",
        "* `gas_updates` - Vector of gas price updates",
        "",
        "# Additional accounts",
        "",
        "In addition to the fixed amount of accounts defined in the `UpdatePrices` context,",
        "the following accounts must be provided:",
        "",
        "* First, the billing token config accounts for each token whose price is being updated, in the same order",
        "as the token_updates vector.",
        "* Then, the dest chain accounts of every chain whose gas price is being updated, in the same order as the",
        "gas_updates vector."
      ],
      "accounts": [
        {
          "name": "authority",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "allowedPriceUpdater",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "was added by the owner as an allowed price updater. The constraints enforced guarantee that it is the right PDA",
            "and that it was initialized."
          ]
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "tokenUpdates",
          "type": {
            "vec": {
              "defined": "TokenPriceUpdate"
            }
          }
        },
        {
          "name": "gasUpdates",
          "type": {
            "vec": {
              "defined": "GasPriceUpdate"
            }
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
            "name": "version",
            "type": "u8"
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
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "publicKey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          }
        ]
      }
    },
    {
      "name": "destChain",
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
              "defined": "DestChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "DestChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "billingTokenConfigWrapper",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "config",
            "type": {
              "defined": "BillingTokenConfig"
            }
          }
        ]
      }
    },
    {
      "name": "perChainPerTokenConfig",
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
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "tokenTransferConfig",
            "type": {
              "defined": "TokenTransferFeeConfig"
            }
          }
        ]
      }
    },
    {
      "name": "allowedPriceUpdater",
      "type": {
        "kind": "struct",
        "fields": []
      }
    }
  ],
  "types": [
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
      "name": "GenericExtraArgsV2",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "gasLimit",
            "type": "u128"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "SVMExtraArgsV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "computeUnits",
            "type": "u32"
          },
          {
            "name": "accountIsWritableBitmap",
            "type": "u64"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          },
          {
            "name": "tokenReceiver",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "accounts",
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
      "name": "SVM2AnyMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "receiver",
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
          },
          {
            "name": "feeToken",
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
      "name": "TokenTransferAdditionalData",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destBytesOverhead",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "GetFeeResult",
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
          },
          {
            "name": "juels",
            "type": "u128"
          },
          {
            "name": "tokenTransferAdditionalData",
            "type": {
              "vec": {
                "defined": "TokenTransferAdditionalData"
              }
            }
          },
          {
            "name": "processedExtraArgs",
            "type": {
              "defined": "ProcessedExtraArgs"
            }
          }
        ]
      }
    },
    {
      "name": "ProcessedExtraArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "bytes",
            "type": "bytes"
          },
          {
            "name": "gasLimit",
            "type": "u128"
          },
          {
            "name": "allowOutOfOrderExecution",
            "type": "bool"
          },
          {
            "name": "tokenReceiver",
            "type": {
              "option": "bytes"
            }
          }
        ]
      }
    },
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "maxNumberOfTokensPerMsg",
            "type": "u16"
          },
          {
            "name": "maxDataBytes",
            "type": "u32"
          },
          {
            "name": "maxPerMsgGasLimit",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteBase",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteHigh",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteThreshold",
            "type": "u32"
          },
          {
            "name": "destDataAvailabilityOverheadGas",
            "type": "u32"
          },
          {
            "name": "destGasPerDataAvailabilityByte",
            "type": "u16"
          },
          {
            "name": "destDataAvailabilityMultiplierBps",
            "type": "u16"
          },
          {
            "name": "defaultTokenFeeUsdcents",
            "type": "u16"
          },
          {
            "name": "defaultTokenDestGasOverhead",
            "type": "u32"
          },
          {
            "name": "defaultTxGasLimit",
            "type": "u32"
          },
          {
            "name": "gasMultiplierWeiPerEth",
            "type": "u64"
          },
          {
            "name": "networkFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "gasPriceStalenessThreshold",
            "type": "u32"
          },
          {
            "name": "enforceOutOfOrder",
            "type": "bool"
          },
          {
            "name": "chainFamilySelector",
            "type": {
              "array": [
                "u8",
                4
              ]
            }
          }
        ]
      }
    },
    {
      "name": "DestChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "usdPerUnitGas",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          }
        ]
      }
    },
    {
      "name": "TimestampedPackedU224",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          }
        ]
      }
    },
    {
      "name": "BillingTokenConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "enabled",
            "type": "bool"
          },
          {
            "name": "mint",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          },
          {
            "name": "premiumMultiplierWeiPerEth",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "TokenTransferFeeConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "maxFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "deciBps",
            "type": "u16"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destBytesOverhead",
            "type": "u32"
          },
          {
            "name": "isEnabled",
            "type": "bool"
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
      "name": "ConfigSet",
      "fields": [
        {
          "name": "maxFeeJuelsPerMsg",
          "type": "u128",
          "index": false
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "linkTokenLocalDecimals",
          "type": "u8",
          "index": false
        },
        {
          "name": "onramp",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "defaultCodeVersion",
          "type": {
            "defined": "CodeVersion"
          },
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenAdded",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "enabled",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenEnabled",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenDisabled",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "FeeTokenRemoved",
      "fields": [
        {
          "name": "feeToken",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "DestChainAdded",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "DestChainConfigUpdated",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "destChainConfig",
          "type": {
            "defined": "DestChainConfig"
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
      "name": "UsdPerUnitGasUpdated",
      "fields": [
        {
          "name": "destChain",
          "type": "u64",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "UsdPerTokenUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    },
    {
      "name": "TokenPriceUpdateIgnored",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "TokenTransferFeeConfigUpdated",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "tokenTransferFeeConfig",
          "type": {
            "defined": "TokenTransferFeeConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "PremiumMultiplierWeiPerEthUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "premiumMultiplierWeiPerEth",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "PriceUpdaterAdded",
      "fields": [
        {
          "name": "priceUpdater",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "PriceUpdaterRemoved",
      "fields": [
        {
          "name": "priceUpdater",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 8000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 8001,
      "name": "InvalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 8002,
      "name": "ZeroGasLimit",
      "msg": "Gas limit is zero"
    },
    {
      "code": 8003,
      "name": "DefaultGasLimitExceedsMaximum",
      "msg": "Default gas limit exceeds the maximum"
    },
    {
      "code": 8004,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 8005,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 8006,
      "name": "InvalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 8007,
      "name": "InvalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 8008,
      "name": "InvalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 8009,
      "name": "InvalidInputsMintOwner",
      "msg": "Mint account input has an invalid owner"
    },
    {
      "code": 8010,
      "name": "InvalidInputsTokenConfigAccount",
      "msg": "Token config account is invalid"
    },
    {
      "code": 8011,
      "name": "InvalidInputsMissingExtraArgs",
      "msg": "Missing extra args in message to SVM receiver"
    },
    {
      "code": 8012,
      "name": "InvalidInputsMissingDataAfterExtraArgs",
      "msg": "Missing data after extra args tag"
    },
    {
      "code": 8013,
      "name": "InvalidInputsDestChainStateAccount",
      "msg": "Destination chain state account is invalid"
    },
    {
      "code": 8014,
      "name": "InvalidInputsPerChainPerTokenConfig",
      "msg": "Per chain per token config account is invalid"
    },
    {
      "code": 8015,
      "name": "InvalidInputsBillingTokenConfig",
      "msg": "Billing token config account is invalid"
    },
    {
      "code": 8016,
      "name": "InvalidInputsAccountCount",
      "msg": "Number of accounts provided is incorrect"
    },
    {
      "code": 8017,
      "name": "InvalidInputsNoUpdates",
      "msg": "No price or gas update provided"
    },
    {
      "code": 8018,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 8019,
      "name": "DestinationChainDisabled",
      "msg": "Destination chain disabled"
    },
    {
      "code": 8020,
      "name": "FeeTokenDisabled",
      "msg": "Fee token disabled"
    },
    {
      "code": 8021,
      "name": "MessageTooLarge",
      "msg": "Message exceeds maximum data size"
    },
    {
      "code": 8022,
      "name": "UnsupportedNumberOfTokens",
      "msg": "Message contains an unsupported number of tokens"
    },
    {
      "code": 8023,
      "name": "InvalidEVMAddress",
      "msg": "Invalid EVM address"
    },
    {
      "code": 8024,
      "name": "InvalidEncoding",
      "msg": "Invalid encoding"
    },
    {
      "code": 8025,
      "name": "InvalidTokenPrice",
      "msg": "Invalid token price"
    },
    {
      "code": 8026,
      "name": "StaleGasPrice",
      "msg": "Stale gas price"
    },
    {
      "code": 8027,
      "name": "InvalidInputsMissingTokenConfig",
      "msg": "Inputs are missing token configuration"
    },
    {
      "code": 8028,
      "name": "MessageFeeTooHigh",
      "msg": "Message fee is too high"
    },
    {
      "code": 8029,
      "name": "MessageGasLimitTooHigh",
      "msg": "Message gas limit too high"
    },
    {
      "code": 8030,
      "name": "ExtraArgOutOfOrderExecutionMustBeTrue",
      "msg": "Extra arg out of order execution must be true"
    },
    {
      "code": 8031,
      "name": "InvalidExtraArgsTag",
      "msg": "Invalid extra args tag"
    },
    {
      "code": 8032,
      "name": "InvalidExtraArgsAccounts",
      "msg": "Invalid amount of accounts in extra args"
    },
    {
      "code": 8033,
      "name": "InvalidExtraArgsWritabilityBitmap",
      "msg": "Invalid writability bitmap in extra args"
    },
    {
      "code": 8034,
      "name": "InvalidChainFamilySelector",
      "msg": "Invalid chain family selector"
    },
    {
      "code": 8035,
      "name": "InvalidTokenReceiver",
      "msg": "Invalid token receiver"
    },
    {
      "code": 8036,
      "name": "InvalidSVMAddress",
      "msg": "Invalid SVM address"
    },
    {
      "code": 8037,
      "name": "UnauthorizedPriceUpdater",
      "msg": "The caller is not an authorized price updater"
    },
    {
      "code": 8038,
      "name": "InvalidLinkDecimals",
      "msg": "The LINK mint uses an unsupported number of decimals"
    },
    {
      "code": 8039,
      "name": "InvalidTokenTransferFeeMaxMin",
      "msg": "Minimum token transfer fee exceeds maximum"
    },
    {
      "code": 8040,
      "name": "InvalidTokenTransferFeeDestBytesOverhead",
      "msg": "Insufficient dest bytes overhead on transfer fee config"
    },
    {
      "code": 8041,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    }
  ]
};
