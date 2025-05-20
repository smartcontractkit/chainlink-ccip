/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/ccip_router.json`.
 */
export type CcipRouter = {
  "address": "Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C",
  "metadata": {
    "name": "ccipRouter",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "docs": [
    "The `ccip_router` module contains the implementation of the Cross-Chain Interoperability Protocol (CCIP) Router.",
    "",
    "This is the Collapsed Router Program for CCIP.",
    "As it's upgradable persisting the same program id, there is no need to have an indirection of a Proxy Program.",
    "This Router handles the OnRamp flow of the CCIP Messages.",
    "",
    "NOTE to devs: This file however should contain *no logic*, only the entrypoints to the different versioned modules,",
    "thus making it easier to ensure later on that logic can be changed during upgrades without affecting the interface."
  ],
  "instructions": [
    {
      "name": "acceptAdminRoleTokenAdminRegistry",
      "docs": [
        "Accepts the admin role of the token admin registry.",
        "",
        "The Pending Admin must call this function to accept the admin role of the Token Admin Registry.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for accepting the admin role.",
        "* `mint` - The public key of the token mint."
      ],
      "discriminator": [
        106,
        240,
        16,
        173,
        137,
        213,
        163,
        246
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": []
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
      "name": "addChainSelector",
      "docs": [
        "Adds a new chain selector to the router.",
        "",
        "The Admin needs to add any new chain supported (this means both OnRamp and OffRamp).",
        "When adding a new chain, the Admin needs to specify if it's enabled or not.",
        "They may enable only source, or only destination, or neither, or both.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding the chain selector.",
        "* `new_chain_selector` - The new chain selector to be added.",
        "* `source_chain_config` - The configuration for the chain as source.",
        "* `dest_chain_config` - The configuration for the chain as destination."
      ],
      "discriminator": [
        28,
        60,
        171,
        0,
        195,
        113,
        56,
        7
      ],
      "accounts": [
        {
          "name": "destChainState",
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
      "name": "addOfframp",
      "docs": [
        "Add an offramp address to the list of offramps allowed by the router, for a",
        "particular source chain. External users will check this list before accepting",
        "a `ccip_receive` CPI.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `source_chain_selector` - The source chain for the offramp's lane.",
        "* `offramp` - The offramp's address."
      ],
      "discriminator": [
        164,
        255,
        154,
        96,
        204,
        239,
        24,
        2
      ],
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "bumpCcipVersionForDestChain",
      "docs": [
        "Bumps the CCIP version for a destination chain.",
        "This effectively just resets the sequence number of the destination chain state.",
        "If there had been a previous rollback, on re-upgrade the sequence number will resume from where it was",
        "prior to the rollback.",
        "",
        "# Arguments",
        "* `ctx` - The context containing the accounts required for the bump.",
        "* `dest_chain_selector` - The destination chain selector to bump version for."
      ],
      "discriminator": [
        120,
        25,
        6,
        201,
        42,
        224,
        235,
        187
      ],
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "ccipAdminOverridePendingAdministrator",
      "docs": [
        "Overrides the pending admin of the Token Admin Registry",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for registration.",
        "* `token_admin_registry_admin` - The public key of the token admin registry admin to propose."
      ],
      "discriminator": [
        163,
        206,
        164,
        199,
        248,
        92,
        36,
        46
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
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
          "name": "tokenAdminRegistryAdmin",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "ccipAdminProposeAdministrator",
      "docs": [
        "Token Admin Registry //",
        "Registers the Token Admin Registry via the CCIP Admin",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for registration.",
        "* `token_admin_registry_admin` - The public key of the token admin registry admin to propose."
      ],
      "discriminator": [
        218,
        37,
        139,
        107,
        142,
        228,
        51,
        219
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
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
          "name": "tokenAdminRegistryAdmin",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "ccipSend",
      "docs": [
        "On Ramp Flow //",
        "Sends a message to the destination chain.",
        "",
        "Request a message to be sent to the destination chain.",
        "The method name needs to be ccip_send with Anchor encoding.",
        "This function is called by the CCIP Sender Contract (or final user) to send a message to the CCIP Router.",
        "The message will be sent to the receiver on the destination chain selector.",
        "This message emits the event CCIPMessageSent with all the necessary data to be retrieved by the OffChain Code",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for sending the message.",
        "* `dest_chain_selector` - The chain selector for the destination chain.",
        "* `message` - The message to be sent. The size limit of data is 256 bytes.",
        "* `token_indexes` - Indices into the remaining accounts vector where the subslice for a token begins."
      ],
      "discriminator": [
        108,
        216,
        134,
        191,
        249,
        234,
        33,
        84
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "destChainState",
          "writable": true
        },
        {
          "name": "nonce",
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
          "name": "feeTokenProgram"
        },
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenUserAssociatedAccount",
          "docs": [
            "If paying with native SOL, this must be the zero address."
          ]
        },
        {
          "name": "feeTokenReceiver",
          "writable": true
        },
        {
          "name": "feeBillingSigner"
        },
        {
          "name": "feeQuoter"
        },
        {
          "name": "feeQuoterConfig"
        },
        {
          "name": "feeQuoterDestChain"
        },
        {
          "name": "feeQuoterBillingTokenConfig"
        },
        {
          "name": "feeQuoterLinkTokenConfig"
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
        },
        {
          "name": "tokenIndexes",
          "type": "bytes"
        }
      ],
      "returns": {
        "array": [
          "u8",
          32
        ]
      }
    },
    {
      "name": "getFee",
      "docs": [
        "Queries the onramp for the fee required to send a message.",
        "",
        "This call is permissionless. Note it does not verify whether there's a curse active",
        "in order to avoid the RMN CPI overhead.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for obtaining the message fee.",
        "* `dest_chain_selector` - The chain selector for the destination chain.",
        "* `message` - The message to be sent. The size limit of data is 256 bytes."
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
          "name": "destChainState"
        },
        {
          "name": "feeQuoter"
        },
        {
          "name": "feeQuoterConfig"
        },
        {
          "name": "feeQuoterDestChain"
        },
        {
          "name": "feeQuoterBillingTokenConfig"
        },
        {
          "name": "feeQuoterLinkTokenConfig"
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
        "Initialization Flow //",
        "Initializes the CCIP Router.",
        "",
        "The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `fee_aggregator` - The public key of the fee aggregator.",
        "* `fee_quoter` - The public key of the fee quoter.",
        "* `link_token_mint` - The public key of the LINK token mint.",
        "* `rmn_remote` - The public key of the RMN remote."
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
          "name": "feeAggregator",
          "type": "pubkey"
        },
        {
          "name": "feeQuoter",
          "type": "pubkey"
        },
        {
          "name": "linkTokenMint",
          "type": "pubkey"
        },
        {
          "name": "rmnRemote",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "ownerOverridePendingAdministrator",
      "docs": [
        "Overrides the pending admin of the Token Admin Registry by the token owner",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for registration.",
        "* `token_admin_registry_admin` - The public key of the token admin registry admin to propose."
      ],
      "discriminator": [
        230,
        111,
        134,
        149,
        203,
        168,
        118,
        201
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
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
          "name": "tokenAdminRegistryAdmin",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "ownerProposeAdministrator",
      "docs": [
        "Registers the Token Admin Registry by the token owner.",
        "",
        "The Authority of the Mint Token can claim the registry of the token.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for registration.",
        "* `token_admin_registry_admin` - The public key of the token admin registry admin to propose."
      ],
      "discriminator": [
        175,
        81,
        160,
        246,
        206,
        132,
        18,
        22
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
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
          "name": "tokenAdminRegistryAdmin",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "removeOfframp",
      "docs": [
        "Remove an offramp address from the list of offramps allowed by the router, for a",
        "particular source chain. External users will check this list before accepting",
        "a `ccip_receive` CPI.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for this operation.",
        "* `source_chain_selector` - The source chain for the offramp's lane.",
        "* `offramp` - The offramp's address."
      ],
      "discriminator": [
        252,
        152,
        51,
        170,
        241,
        13,
        199,
        8
      ],
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "rollbackCcipVersionForDestChain",
      "docs": [
        "Rolls back the CCIP version for a destination chain.",
        "This effectively just restores the old version's sequence number of the destination chain state.",
        "We only support 1 consecutive rollback. If a rollback has occurred for that lane, the version can't",
        "be rolled back again without bumping the version first.",
        "",
        "# Arguments",
        "* `ctx` - The context containing the accounts required for the rollback.",
        "* `dest_chain_selector` - The destination chain selector to rollback the version for."
      ],
      "discriminator": [
        95,
        107,
        33,
        138,
        26,
        57,
        154,
        110
      ],
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "setDefaultCodeVersion",
      "docs": [
        "Config //",
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
        },
        {
          "name": "systemProgram"
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
        "Sets the address of the LINK token mint.",
        "The Admin is the only one able to set it.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `link_token_mint` - The new address of the LINK token mint."
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
          "name": "authority",
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "linkTokenMint",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "setPool",
      "docs": [
        "Sets the pool lookup table for a given token mint.",
        "",
        "The administrator of the token admin registry can set the pool lookup table for a given token mint.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for setting the pool.",
        "* `writable_indexes` - a bit map of the indexes of the accounts in lookup table that are writable"
      ],
      "discriminator": [
        119,
        30,
        14,
        180,
        115,
        225,
        167,
        238
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "poolLookuptable"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "writableIndexes",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "transferAdminRoleTokenAdminRegistry",
      "docs": [
        "Transfers the admin role of the token admin registry to a new admin.",
        "",
        "Only the Admin can transfer the Admin Role of the Token Admin Registry, this setups the Pending Admin and then it's their responsibility to accept the role.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the transfer.",
        "* `mint` - The public key of the token mint.",
        "* `new_admin` - The public key of the new admin."
      ],
      "discriminator": [
        178,
        98,
        203,
        181,
        203,
        107,
        106,
        14
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "tokenAdminRegistry",
          "writable": true
        },
        {
          "name": "mint"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "newAdmin",
          "type": "pubkey"
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
      "name": "updateDestChainConfig",
      "docs": [
        "Updates the configuration of the destination chain selector.",
        "",
        "The Admin is the only one able to update the destination chain config.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the chain selector.",
        "* `dest_chain_selector` - The destination chain selector to be updated.",
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
          "name": "destChainState",
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
    },
    {
      "name": "updateFeeAggregator",
      "docs": [
        "Updates the fee aggregator in the router configuration.",
        "The Admin is the only one able to update the fee aggregator.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `fee_aggregator` - The new fee aggregator address (ATAs will be derived for it for each token)."
      ],
      "discriminator": [
        85,
        112,
        115,
        60,
        22,
        95,
        230,
        56
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "feeAggregator",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "updateRmnRemote",
      "docs": [
        "Updates the RMN remote program in the router configuration.",
        "The Admin is the only one able to update the RMN remote program.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for updating the configuration.",
        "* `rmn_remote,` - The new RMN remote address."
      ],
      "discriminator": [
        66,
        12,
        215,
        147,
        14,
        176,
        55,
        214
      ],
      "accounts": [
        {
          "name": "config",
          "writable": true
        },
        {
          "name": "authority",
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "rmnRemote",
          "type": "pubkey"
        }
      ]
    },
    {
      "name": "updateSvmChainSelector",
      "docs": [
        "Updates the SVM chain selector in the router configuration.",
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
        },
        {
          "name": "systemProgram"
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
      "name": "withdrawBilledFunds",
      "docs": [
        "Billing //",
        "Transfers the accumulated billed fees in a particular token to an arbitrary token account.",
        "Only the CCIP Admin can withdraw billed funds.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for the transfer of billed fees.",
        "* `transfer_all` - A flag indicating whether to transfer all the accumulated fees in that token or not.",
        "* `desired_amount` - The amount to transfer. If `transfer_all` is true, this value must be 0."
      ],
      "discriminator": [
        16,
        116,
        73,
        38,
        77,
        232,
        6,
        28
      ],
      "accounts": [
        {
          "name": "feeTokenMint"
        },
        {
          "name": "feeTokenAccum",
          "writable": true
        },
        {
          "name": "recipient",
          "writable": true
        },
        {
          "name": "tokenProgram"
        },
        {
          "name": "feeBillingSigner"
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
          "name": "transferAll",
          "type": "bool"
        },
        {
          "name": "desiredAmount",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "discriminator": [
        247,
        97,
        179,
        16,
        207,
        36,
        236,
        132
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
      "name": "nonce",
      "discriminator": [
        143,
        197,
        147,
        95,
        106,
        165,
        50,
        43
      ]
    },
    {
      "name": "tokenAdminRegistry",
      "discriminator": [
        70,
        92,
        207,
        200,
        76,
        17,
        57,
        114
      ]
    }
  ],
  "events": [
    {
      "name": "administratorTransferRequested",
      "discriminator": [
        159,
        30,
        110,
        86,
        22,
        35,
        70,
        125
      ]
    },
    {
      "name": "administratorTransferred",
      "discriminator": [
        103,
        127,
        255,
        114,
        168,
        163,
        159,
        124
      ]
    },
    {
      "name": "ccipMessageSent",
      "discriminator": [
        23,
        77,
        73,
        183,
        123,
        185,
        115,
        57
      ]
    },
    {
      "name": "ccipVersionForDestChainVersionBumped",
      "discriminator": [
        81,
        97,
        90,
        70,
        154,
        163,
        255,
        78
      ]
    },
    {
      "name": "ccipVersionForDestChainVersionRolledBack",
      "discriminator": [
        50,
        79,
        44,
        175,
        232,
        241,
        225,
        171
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
      "name": "offrampAdded",
      "discriminator": [
        158,
        77,
        52,
        73,
        113,
        247,
        76,
        150
      ]
    },
    {
      "name": "offrampRemoved",
      "discriminator": [
        231,
        81,
        202,
        9,
        89,
        193,
        154,
        37
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
      "name": "poolSet",
      "discriminator": [
        135,
        203,
        185,
        106,
        113,
        87,
        177,
        32
      ]
    }
  ],
  "errors": [
    {
      "code": 7000,
      "name": "unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 7001,
      "name": "invalidRmnRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 7002,
      "name": "invalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 7003,
      "name": "invalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 7004,
      "name": "feeTokenMismatch",
      "msg": "Fee token doesn't match transfer token"
    },
    {
      "code": 7005,
      "name": "redundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 7006,
      "name": "reachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 7007,
      "name": "invalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 7008,
      "name": "invalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 7009,
      "name": "invalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 7010,
      "name": "invalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 7011,
      "name": "invalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 7012,
      "name": "invalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 7013,
      "name": "invalidInputsTokenAmount",
      "msg": "Cannot send zero tokens"
    },
    {
      "code": 7014,
      "name": "invalidInputsTransferAllAmount",
      "msg": "Must specify zero amount to send alongside transfer_all"
    },
    {
      "code": 7015,
      "name": "invalidInputsAtaAddress",
      "msg": "Invalid Associated Token Account address"
    },
    {
      "code": 7016,
      "name": "invalidInputsAtaWritable",
      "msg": "Invalid Associated Token Account writable flag"
    },
    {
      "code": 7017,
      "name": "invalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 7018,
      "name": "insufficientLamports",
      "msg": "Insufficient lamports"
    },
    {
      "code": 7019,
      "name": "insufficientFunds",
      "msg": "Insufficient funds"
    },
    {
      "code": 7020,
      "name": "sourceTokenDataTooLarge",
      "msg": "Source token data is too large"
    },
    {
      "code": 7021,
      "name": "invalidTokenAdminRegistryInputsZeroAddress",
      "msg": "New Admin can not be zero address"
    },
    {
      "code": 7022,
      "name": "invalidTokenAdminRegistryProposedAdmin",
      "msg": "An already owned registry can not be proposed"
    },
    {
      "code": 7023,
      "name": "senderNotAllowed",
      "msg": "Sender not allowed for that destination chain"
    },
    {
      "code": 7024,
      "name": "invalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 7025,
      "name": "invalidCcipVersionRollback",
      "msg": "Invalid rollback attempt on the CCIP version of the onramp to the destination chain"
    }
  ],
  "types": [
    {
      "name": "administratorTransferRequested",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "currentAdmin",
            "type": "pubkey"
          },
          {
            "name": "newAdmin",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "administratorTransferred",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "newAdmin",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
      }
    },
    {
      "name": "ccipMessageSent",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "message",
            "type": {
              "defined": {
                "name": "svm2AnyRampMessage"
              }
            }
          }
        ]
      }
    },
    {
      "name": "ccipVersionForDestChainVersionBumped",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "previousSequenceNumber",
            "type": "u64"
          },
          {
            "name": "newSequenceNumber",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "ccipVersionForDestChainVersionRolledBack",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "previousSequenceNumber",
            "type": "u64"
          },
          {
            "name": "newSequenceNumber",
            "type": "u64"
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
            "name": "defaultCodeVersion",
            "type": {
              "defined": {
                "name": "codeVersion"
              }
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
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
            "name": "feeQuoter",
            "type": "pubkey"
          },
          {
            "name": "rmnRemote",
            "type": "pubkey"
          },
          {
            "name": "linkTokenMint",
            "type": "pubkey"
          },
          {
            "name": "feeAggregator",
            "type": "pubkey"
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
            "name": "feeQuoter",
            "type": "pubkey"
          },
          {
            "name": "rmnRemote",
            "type": "pubkey"
          },
          {
            "name": "linkTokenMint",
            "type": "pubkey"
          },
          {
            "name": "feeAggregator",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "crossChainAmount",
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
            "name": "laneCodeVersion",
            "type": {
              "defined": {
                "name": "codeVersion"
              }
            }
          },
          {
            "name": "allowedSenders",
            "type": {
              "vec": "pubkey"
            }
          },
          {
            "name": "allowListEnabled",
            "type": "bool"
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
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "sequenceNumberToRestore",
            "type": "u64"
          },
          {
            "name": "restoreOnAction",
            "type": {
              "defined": {
                "name": "restoreOnAction"
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
      "name": "getFeeResult",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "amount",
            "type": "u64"
          },
          {
            "name": "juels",
            "type": "u128"
          },
          {
            "name": "token",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "nonce",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "counter",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "offrampAdded",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "offramp",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "offrampRemoved",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "offramp",
            "type": "pubkey"
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
      "name": "poolSet",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "pubkey"
          },
          {
            "name": "previousPoolLookupTable",
            "type": "pubkey"
          },
          {
            "name": "newPoolLookupTable",
            "type": "pubkey"
          }
        ]
      }
    },
    {
      "name": "rampMessageHeader",
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
      "name": "restoreOnAction",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "none"
          },
          {
            "name": "upgrade"
          },
          {
            "name": "rollback"
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
      "name": "svm2AnyRampMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "header",
            "type": {
              "defined": {
                "name": "rampMessageHeader"
              }
            }
          },
          {
            "name": "sender",
            "type": "pubkey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "receiver",
            "type": "bytes"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
          },
          {
            "name": "feeToken",
            "type": "pubkey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": {
                  "name": "svm2AnyTokenTransfer"
                }
              }
            }
          },
          {
            "name": "feeTokenAmount",
            "type": {
              "defined": {
                "name": "crossChainAmount"
              }
            }
          },
          {
            "name": "feeValueJuels",
            "type": {
              "defined": {
                "name": "crossChainAmount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "svm2AnyTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "pubkey"
          },
          {
            "name": "destTokenAddress",
            "type": "bytes"
          },
          {
            "name": "extraData",
            "type": "bytes"
          },
          {
            "name": "amount",
            "type": {
              "defined": {
                "name": "crossChainAmount"
              }
            }
          },
          {
            "name": "destExecData",
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
      "name": "tokenAdminRegistry",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "administrator",
            "type": "pubkey"
          },
          {
            "name": "pendingAdministrator",
            "type": "pubkey"
          },
          {
            "name": "lookupTable",
            "type": "pubkey"
          },
          {
            "name": "writableIndexes",
            "type": {
              "array": [
                "u128",
                2
              ]
            }
          },
          {
            "name": "mint",
            "type": "pubkey"
          }
        ]
      }
    }
  ]
};
