/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/fee_quoter.json`.
 */
export type FeeQuoter = {
  "address": "FeeQPGkKDeRV1MgoYfMH6L8o3KeuYjwUZrgn4LRKfjHi",
  "metadata": {
    "name": "feeQuoter",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
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
      "discriminator": [
        63,
        156,
        254,
        216,
        227,
        53,
        0,
        69
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "billingTokenConfig",
          "writable": true
        },
        {
          "name": "tokenProgram"
        },
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenReceiver",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "feeBillingSigner"
        },
        {
          "name": "associatedTokenProgram"
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "config",
          "type": {
            "defined": {
              "name": "billingTokenConfig"
            }
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
      "discriminator": [
        122,
        202,
        174,
        155,
        55,
        100,
        102,
        36
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "destChain",
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
          "name": "destChainConfig",
          "type": {
            "defined": {
              "name": "destChainConfig"
            }
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
      "discriminator": [
        200,
        26,
        13,
        120,
        226,
        182,
        64,
        16
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
          "name": "priceUpdater",
          "type": "pubkey"
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
      "discriminator": [
        200,
        195,
        114,
        206,
        152,
        86,
        50,
        41
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "destChain",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
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
      "discriminator": [
        115,
        195,
        235,
        161,
        25,
        219,
        60,
        29
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "destChain"
        },
        {
          "name": "billingTokenConfig"
        },
        {
          "name": "linkTokenConfig"
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
            "defined": {
              "name": "svm2AnyMessage"
            }
          }
        }
      ],
      "returns": {
        "defined": {
          "name": "getFeeResult"
        }
      }
    },
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
        "* `max_fee_juels_per_msg` - The maximum fee in juels that can be charged per message.",
        "* `onramp` - The public key of the onramp.",
        "",
        "The function also uses the link_token_mint account from the context."
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
          "name": "config",
          "writable": true
        },
        {
          "name": "linkTokenMint"
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
          "name": "maxFeeJuelsPerMsg",
          "type": "u128"
        },
        {
          "name": "onramp",
          "type": "pubkey"
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
      "discriminator": [
        10,
        61,
        172,
        48,
        110,
        8,
        162,
        198
      ],
      "accounts": [
        {
          "name": "allowedPriceUpdater",
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
          "name": "priceUpdater",
          "type": "pubkey"
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
      "name": "setLinkTokenMint",
      "docs": [
        "Sets the link_token_mint and updates the link_token_local_decimals.",
        "",
        "Only the admin may set this.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration."
      ],
      "discriminator": [
        190,
        216,
        49,
        254,
        200,
        81,
        12,
        17
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "linkTokenMint"
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": []
    },
    {
      "name": "setMaxFeeJuelsPerMsg",
      "docs": [
        "Sets the max_fee_juels_per_msg, which is an upper bound on how much can be billed for any message.",
        "(1 juels = 1e-18 LINK)",
        "",
        "Only the admin may set this.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `max_fee_juels_per_msg` - The new value for the max_feel_juels_per_msg config."
      ],
      "discriminator": [
        50,
        235,
        110,
        147,
        169,
        199,
        69,
        46
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
          "name": "maxFeeJuelsPerMsg",
          "type": "u128"
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
      "discriminator": [
        76,
        243,
        16,
        214,
        126,
        11,
        254,
        77
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "perChainPerTokenConfig",
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
          "name": "mint",
          "type": "pubkey"
        },
        {
          "name": "cfg",
          "type": {
            "defined": {
              "name": "tokenTransferFeeConfig"
            }
          }
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
          "name": "newOwner",
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
      "discriminator": [
        140,
        184,
        124,
        146,
        204,
        62,
        244,
        79
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "billingTokenConfig",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        }
      ],
      "args": [
        {
          "name": "config",
          "type": {
            "defined": {
              "name": "billingTokenConfig"
            }
          }
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
      "discriminator": [
        215,
        122,
        81,
        22,
        190,
        58,
        219,
        13
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "destChain",
          "writable": true
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
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
            "defined": {
              "name": "destChainConfig"
            }
          }
        }
      ]
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
      "discriminator": [
        62,
        161,
        234,
        136,
        106,
        26,
        18,
        160
      ],
      "accounts": [
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "allowedPriceUpdater",
          "docs": [
            "was added by the owner as an allowed price updater. The constraints enforced guarantee that it is the right PDA",
            "and that it was initialized."
          ]
        },
        {
          "name": "config"
        }
      ],
      "args": [
        {
          "name": "tokenUpdates",
          "type": {
            "vec": {
              "defined": {
                "name": "tokenPriceUpdate"
              }
            }
          }
        },
        {
          "name": "gasUpdates",
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
  ],
  "accounts": [
    {
      "name": "allowedPriceUpdater",
      "discriminator": [
        51,
        136,
        222,
        161,
        38,
        7,
        184,
        190
      ]
    },
    {
      "name": "billingTokenConfigWrapper",
      "discriminator": [
        63,
        178,
        72,
        57,
        171,
        66,
        44,
        151
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
      "name": "destChain",
      "discriminator": [
        77,
        18,
        241,
        132,
        212,
        54,
        218,
        16
      ]
    },
    {
      "name": "perChainPerTokenConfig",
      "discriminator": [
        183,
        88,
        20,
        99,
        246,
        46,
        51,
        230
      ]
    }
  ],
  "events": [
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
      "name": "destChainAdded",
      "discriminator": [
        59,
        154,
        48,
        81,
        230,
        41,
        80,
        200
      ]
    },
    {
      "name": "destChainConfigUpdated",
      "discriminator": [
        3,
        141,
        73,
        190,
        73,
        231,
        51,
        80
      ]
    },
    {
      "name": "feeTokenAdded",
      "discriminator": [
        181,
        180,
        252,
        21,
        215,
        79,
        93,
        237
      ]
    },
    {
      "name": "feeTokenDisabled",
      "discriminator": [
        34,
        139,
        66,
        75,
        30,
        17,
        45,
        151
      ]
    },
    {
      "name": "feeTokenEnabled",
      "discriminator": [
        106,
        180,
        145,
        189,
        113,
        180,
        21,
        15
      ]
    },
    {
      "name": "feeTokenRemoved",
      "discriminator": [
        40,
        31,
        230,
        252,
        183,
        150,
        147,
        201
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
      "name": "premiumMultiplierWeiPerEthUpdated",
      "discriminator": [
        151,
        5,
        223,
        182,
        215,
        187,
        249,
        225
      ]
    },
    {
      "name": "priceUpdaterAdded",
      "discriminator": [
        87,
        31,
        151,
        133,
        151,
        187,
        97,
        186
      ]
    },
    {
      "name": "priceUpdaterRemoved",
      "discriminator": [
        225,
        194,
        40,
        213,
        212,
        39,
        76,
        148
      ]
    },
    {
      "name": "tokenPriceUpdateIgnored",
      "discriminator": [
        68,
        119,
        161,
        131,
        128,
        65,
        69,
        201
      ]
    },
    {
      "name": "tokenTransferFeeConfigUpdated",
      "discriminator": [
        253,
        199,
        166,
        1,
        178,
        150,
        242,
        253
      ]
    },
    {
      "name": "usdPerTokenUpdated",
      "discriminator": [
        67,
        154,
        252,
        56,
        104,
        14,
        192,
        219
      ]
    },
    {
      "name": "usdPerUnitGasUpdated",
      "discriminator": [
        174,
        255,
        2,
        41,
        197,
        110,
        31,
        40
      ]
    }
  ],
  "errors": [
    {
      "code": 8000,
      "name": "unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 8001,
      "name": "invalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 8002,
      "name": "zeroGasLimit",
      "msg": "Gas limit is zero"
    },
    {
      "code": 8003,
      "name": "defaultGasLimitExceedsMaximum",
      "msg": "Default gas limit exceeds the maximum"
    },
    {
      "code": 8004,
      "name": "invalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 8005,
      "name": "redundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 8006,
      "name": "invalidInputsMissingWritable",
      "msg": "Account should be writable"
    },
    {
      "code": 8007,
      "name": "invalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 8008,
      "name": "invalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 8009,
      "name": "invalidInputsMintOwner",
      "msg": "Mint account input has an invalid owner"
    },
    {
      "code": 8010,
      "name": "invalidInputsTokenConfigAccount",
      "msg": "Token config account is invalid"
    },
    {
      "code": 8011,
      "name": "invalidInputsMissingExtraArgs",
      "msg": "Missing extra args in message to SVM receiver"
    },
    {
      "code": 8012,
      "name": "invalidInputsMissingDataAfterExtraArgs",
      "msg": "Missing data after extra args tag"
    },
    {
      "code": 8013,
      "name": "invalidInputsDestChainStateAccount",
      "msg": "Destination chain state account is invalid"
    },
    {
      "code": 8014,
      "name": "invalidInputsPerChainPerTokenConfig",
      "msg": "Per chain per token config account is invalid"
    },
    {
      "code": 8015,
      "name": "invalidInputsBillingTokenConfig",
      "msg": "Billing token config account is invalid"
    },
    {
      "code": 8016,
      "name": "invalidInputsAccountCount",
      "msg": "Number of accounts provided is incorrect"
    },
    {
      "code": 8017,
      "name": "invalidInputsNoUpdates",
      "msg": "No price or gas update provided"
    },
    {
      "code": 8018,
      "name": "invalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 8019,
      "name": "destinationChainDisabled",
      "msg": "Destination chain disabled"
    },
    {
      "code": 8020,
      "name": "feeTokenDisabled",
      "msg": "Fee token disabled"
    },
    {
      "code": 8021,
      "name": "messageTooLarge",
      "msg": "Message exceeds maximum data size"
    },
    {
      "code": 8022,
      "name": "unsupportedNumberOfTokens",
      "msg": "Message contains an unsupported number of tokens"
    },
    {
      "code": 8023,
      "name": "invalidTokenPrice",
      "msg": "Invalid token price"
    },
    {
      "code": 8024,
      "name": "staleGasPrice",
      "msg": "Stale gas price"
    },
    {
      "code": 8025,
      "name": "invalidInputsMissingTokenConfig",
      "msg": "Inputs are missing token configuration"
    },
    {
      "code": 8026,
      "name": "messageFeeTooHigh",
      "msg": "Message fee is too high"
    },
    {
      "code": 8027,
      "name": "messageGasLimitTooHigh",
      "msg": "Message gas limit too high"
    },
    {
      "code": 8028,
      "name": "extraArgOutOfOrderExecutionMustBeTrue",
      "msg": "Extra arg out of order execution must be true"
    },
    {
      "code": 8029,
      "name": "invalidExtraArgsTag",
      "msg": "Invalid extra args tag"
    },
    {
      "code": 8030,
      "name": "invalidExtraArgsAccounts",
      "msg": "Invalid amount of accounts in extra args"
    },
    {
      "code": 8031,
      "name": "invalidExtraArgsWritabilityBitmap",
      "msg": "Invalid writability bitmap in extra args"
    },
    {
      "code": 8032,
      "name": "invalidTokenReceiver",
      "msg": "Invalid token receiver"
    },
    {
      "code": 8033,
      "name": "unauthorizedPriceUpdater",
      "msg": "The caller is not an authorized price updater"
    },
    {
      "code": 8034,
      "name": "invalidTokenTransferFeeMaxMin",
      "msg": "Minimum token transfer fee exceeds maximum"
    },
    {
      "code": 8035,
      "name": "invalidTokenTransferFeeDestBytesOverhead",
      "msg": "Insufficient dest bytes overhead on transfer fee config"
    },
    {
      "code": 8036,
      "name": "invalidCodeVersion",
      "msg": "Invalid code version"
    }
  ],
  "types": [
    {
      "name": "allowedPriceUpdater",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "billingTokenConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "enabled",
            "type": "bool"
          },
          {
            "name": "mint",
            "type": "pubkey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "defined": {
                "name": "timestampedPackedU224"
              }
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
              "defined": {
                "name": "billingTokenConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "codeVersion",
      "repr": {
        "kind": "rust"
      },
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "default"
          },
          {
            "name": "v1"
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
            "name": "version",
            "type": "u8"
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
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "pubkey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "pubkey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": {
                "name": "codeVersion"
              }
            }
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
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "pubkey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "pubkey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": {
                "name": "codeVersion"
              }
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
              "defined": {
                "name": "destChainState"
              }
            }
          },
          {
            "name": "config",
            "type": {
              "defined": {
                "name": "destChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "destChainAdded",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "destChainConfig",
            "type": {
              "defined": {
                "name": "destChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "destChainConfig",
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
              "defined": {
                "name": "codeVersion"
              }
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
      "name": "destChainConfigUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "destChainConfig",
            "type": {
              "defined": {
                "name": "destChainConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "destChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "usdPerUnitGas",
            "type": {
              "defined": {
                "name": "timestampedPackedU224"
              }
            }
          }
        ]
      }
    },
    {
      "name": "feeTokenAdded",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "feeToken",
            "type": "pubkey"
          },
          {
            "name": "enabled",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "feeTokenDisabled",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "feeToken",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "feeTokenEnabled",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "feeToken",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "feeTokenRemoved",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "feeToken",
            "type": "pubkey"
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
      "name": "getFeeResult",
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
          },
          {
            "name": "juels",
            "type": "u128"
          },
          {
            "name": "tokenTransferAdditionalData",
            "type": {
              "vec": {
                "defined": {
                  "name": "tokenTransferAdditionalData"
                }
              }
            }
          },
          {
            "name": "processedExtraArgs",
            "type": {
              "defined": {
                "name": "processedExtraArgs"
              }
            }
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
            "type": "pubkey"
          },
          {
            "name": "tokenTransferConfig",
            "type": {
              "defined": {
                "name": "tokenTransferFeeConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "premiumMultiplierWeiPerEthUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "premiumMultiplierWeiPerEth",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "priceUpdaterAdded",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "priceUpdater",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "priceUpdaterRemoved",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "priceUpdater",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "processedExtraArgs",
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
      "name": "svm2AnyMessage",
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
                "defined": {
                  "name": "svmTokenAmount"
                }
              }
            }
          },
          {
            "name": "feeToken",
            "type": "pubkey"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
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
    },
    {
      "name": "timestampedPackedU224",
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
      "name": "tokenPriceUpdateIgnored",
      "docs": [
        "Emitted when the price update corresponds to a token that isn't registered",
        "for price tracking."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "value",
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
      "name": "tokenTransferAdditionalData",
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
      "name": "tokenTransferFeeConfig",
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
      "name": "tokenTransferFeeConfigUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "tokenTransferFeeConfig",
            "type": {
              "defined": {
                "name": "tokenTransferFeeConfig"
              }
            }
          }
        ]
      }
    },
    {
      "name": "usdPerTokenUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
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
      "name": "usdPerUnitGasUpdated",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChain",
            "type": "u64"
          },
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
    }
  ]
};
