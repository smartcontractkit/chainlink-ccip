export type CcipRouter = {
  "version": "0.1.0-dev",
  "name": "ccip_router",
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
          "name": "feeAggregator",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey"
        },
        {
          "name": "rmnRemote",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "linkTokenMint",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "feeAggregator",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "newChainSelector",
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
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
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
        }
      ]
    },
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        }
      ],
      "args": []
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "newAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolLookuptable",
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
          "name": "writableIndexes",
          "type": "bytes"
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
      "accounts": [
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAccum",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "recipient",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
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
          "name": "transferAll",
          "type": "bool"
        },
        {
          "name": "desiredAmount",
          "type": "u64"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChainState",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "nonce",
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
          "name": "feeTokenUserAssociatedAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "If paying with native SOL, this must be the zero address."
          ]
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
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterLinkTokenConfig",
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
          "name": "destChainSelector",
          "type": "u64"
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyMessage"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChainState",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterLinkTokenConfig",
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
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
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
              "defined": "CodeVersion"
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
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
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "feeAggregator",
            "type": "publicKey"
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
    }
  ],
  "types": [
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
      "name": "SVM2AnyRampMessage",
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
            "type": "publicKey"
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
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVM2AnyTokenTransfer"
              }
            }
          },
          {
            "name": "feeTokenAmount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "feeValueJuels",
            "type": {
              "defined": "CrossChainAmount"
            }
          }
        ]
      }
    },
    {
      "name": "SVM2AnyTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "publicKey"
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
              "defined": "CrossChainAmount"
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
      "name": "GetFeeResult",
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
            "type": "publicKey"
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
              "defined": "RestoreOnAction"
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
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "allowedSenders",
            "type": {
              "vec": "publicKey"
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
    },
    {
      "name": "RestoreOnAction",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Upgrade"
          },
          {
            "name": "Rollback"
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
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "feeQuoter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "rmnRemote",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "feeAggregator",
          "type": "publicKey",
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
      "name": "OfframpAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "offramp",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OfframpRemoved",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "offramp",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "CcipVersionForDestChainVersionBumped",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "previousSequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "newSequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "CcipVersionForDestChainVersionRolledBack",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "previousSequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "newSequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "CCIPMessageSent",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyRampMessage"
          },
          "index": false
        }
      ]
    },
    {
      "name": "PoolSet",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "previousPoolLookupTable",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newPoolLookupTable",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AdministratorTransferRequested",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "currentAdmin",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAdmin",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AdministratorTransferred",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAdmin",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 7000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 7001,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 7002,
      "name": "InvalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 7003,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 7004,
      "name": "FeeTokenMismatch",
      "msg": "Fee token doesn't match transfer token"
    },
    {
      "code": 7005,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 7006,
      "name": "ReachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 7007,
      "name": "InvalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 7008,
      "name": "InvalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 7009,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 7010,
      "name": "InvalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 7011,
      "name": "InvalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 7012,
      "name": "InvalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 7013,
      "name": "InvalidInputsTokenAmount",
      "msg": "Cannot send zero tokens"
    },
    {
      "code": 7014,
      "name": "InvalidInputsTransferAllAmount",
      "msg": "Must specify zero amount to send alongside transfer_all"
    },
    {
      "code": 7015,
      "name": "InvalidInputsAtaAddress",
      "msg": "Invalid Associated Token Account address"
    },
    {
      "code": 7016,
      "name": "InvalidInputsAtaWritable",
      "msg": "Invalid Associated Token Account writable flag"
    },
    {
      "code": 7017,
      "name": "InvalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 7018,
      "name": "InsufficientLamports",
      "msg": "Insufficient lamports"
    },
    {
      "code": 7019,
      "name": "InsufficientFunds",
      "msg": "Insufficient funds"
    },
    {
      "code": 7020,
      "name": "SourceTokenDataTooLarge",
      "msg": "Source token data is too large"
    },
    {
      "code": 7021,
      "name": "InvalidTokenAdminRegistryInputsZeroAddress",
      "msg": "New Admin can not be zero address"
    },
    {
      "code": 7022,
      "name": "InvalidTokenAdminRegistryProposedAdmin",
      "msg": "An already owned registry can not be proposed"
    },
    {
      "code": 7023,
      "name": "SenderNotAllowed",
      "msg": "Sender not allowed for that destination chain"
    },
    {
      "code": 7024,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 7025,
      "name": "InvalidCcipVersionRollback",
      "msg": "Invalid rollback attempt on the CCIP version of the onramp to the destination chain"
    }
  ]
};

export const IDL: CcipRouter = {
  "version": "0.1.0-dev",
  "name": "ccip_router",
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
          "name": "feeAggregator",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey"
        },
        {
          "name": "rmnRemote",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "linkTokenMint",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "feeAggregator",
          "type": "publicKey"
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
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "rmnRemote",
          "type": "publicKey"
        }
      ]
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "newChainSelector",
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
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
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "allowedOfframp",
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
          "name": "sourceChainSelector",
          "type": "u64"
        },
        {
          "name": "offramp",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
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
      "accounts": [
        {
          "name": "destChainState",
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
          "name": "destChainSelector",
          "type": "u64"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "tokenAdminRegistryAdmin",
          "type": "publicKey"
        }
      ]
    },
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        }
      ],
      "args": []
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
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
          "name": "newAdmin",
          "type": "publicKey"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "tokenAdminRegistry",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "mint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "poolLookuptable",
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
          "name": "writableIndexes",
          "type": "bytes"
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
      "accounts": [
        {
          "name": "feeTokenMint",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeTokenAccum",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "recipient",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeBillingSigner",
          "isMut": false,
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
          "name": "transferAll",
          "type": "bool"
        },
        {
          "name": "desiredAmount",
          "type": "u64"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChainState",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "nonce",
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
          "name": "feeTokenUserAssociatedAccount",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "If paying with native SOL, this must be the zero address."
          ]
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
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterLinkTokenConfig",
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
          "name": "destChainSelector",
          "type": "u64"
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyMessage"
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "destChainState",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoter",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterDestChain",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterBillingTokenConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "feeQuoterLinkTokenConfig",
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
    }
  ],
  "accounts": [
    {
      "name": "allowedOfframp",
      "type": {
        "kind": "struct",
        "fields": []
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
              "defined": "CodeVersion"
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
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
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "feeAggregator",
            "type": "publicKey"
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
    }
  ],
  "types": [
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
      "name": "SVM2AnyRampMessage",
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
            "type": "publicKey"
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
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVM2AnyTokenTransfer"
              }
            }
          },
          {
            "name": "feeTokenAmount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "feeValueJuels",
            "type": {
              "defined": "CrossChainAmount"
            }
          }
        ]
      }
    },
    {
      "name": "SVM2AnyTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "publicKey"
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
              "defined": "CrossChainAmount"
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
      "name": "GetFeeResult",
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
            "type": "publicKey"
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
              "defined": "RestoreOnAction"
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
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "allowedSenders",
            "type": {
              "vec": "publicKey"
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
    },
    {
      "name": "RestoreOnAction",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Upgrade"
          },
          {
            "name": "Rollback"
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
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "feeQuoter",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "rmnRemote",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "linkTokenMint",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "feeAggregator",
          "type": "publicKey",
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
      "name": "OfframpAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "offramp",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "OfframpRemoved",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "offramp",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "CcipVersionForDestChainVersionBumped",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "previousSequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "newSequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "CcipVersionForDestChainVersionRolledBack",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "previousSequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "newSequenceNumber",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "CCIPMessageSent",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyRampMessage"
          },
          "index": false
        }
      ]
    },
    {
      "name": "PoolSet",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "previousPoolLookupTable",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newPoolLookupTable",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AdministratorTransferRequested",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "currentAdmin",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAdmin",
          "type": "publicKey",
          "index": false
        }
      ]
    },
    {
      "name": "AdministratorTransferred",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "newAdmin",
          "type": "publicKey",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 7000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 7001,
      "name": "InvalidRMNRemoteAddress",
      "msg": "Invalid RMN Remote Address"
    },
    {
      "code": 7002,
      "name": "InvalidInputsMint",
      "msg": "Mint account input is invalid"
    },
    {
      "code": 7003,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 7004,
      "name": "FeeTokenMismatch",
      "msg": "Fee token doesn't match transfer token"
    },
    {
      "code": 7005,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 7006,
      "name": "ReachedMaxSequenceNumber",
      "msg": "Reached max sequence number"
    },
    {
      "code": 7007,
      "name": "InvalidInputsTokenIndices",
      "msg": "Invalid pool account account indices"
    },
    {
      "code": 7008,
      "name": "InvalidInputsPoolAccounts",
      "msg": "Invalid pool accounts"
    },
    {
      "code": 7009,
      "name": "InvalidInputsTokenAccounts",
      "msg": "Invalid token accounts"
    },
    {
      "code": 7010,
      "name": "InvalidInputsTokenAdminRegistryAccounts",
      "msg": "Invalid Token Admin Registry account"
    },
    {
      "code": 7011,
      "name": "InvalidInputsLookupTableAccounts",
      "msg": "Invalid LookupTable account"
    },
    {
      "code": 7012,
      "name": "InvalidInputsLookupTableAccountWritable",
      "msg": "Invalid LookupTable account writable access"
    },
    {
      "code": 7013,
      "name": "InvalidInputsTokenAmount",
      "msg": "Cannot send zero tokens"
    },
    {
      "code": 7014,
      "name": "InvalidInputsTransferAllAmount",
      "msg": "Must specify zero amount to send alongside transfer_all"
    },
    {
      "code": 7015,
      "name": "InvalidInputsAtaAddress",
      "msg": "Invalid Associated Token Account address"
    },
    {
      "code": 7016,
      "name": "InvalidInputsAtaWritable",
      "msg": "Invalid Associated Token Account writable flag"
    },
    {
      "code": 7017,
      "name": "InvalidInputsChainSelector",
      "msg": "Chain selector is invalid"
    },
    {
      "code": 7018,
      "name": "InsufficientLamports",
      "msg": "Insufficient lamports"
    },
    {
      "code": 7019,
      "name": "InsufficientFunds",
      "msg": "Insufficient funds"
    },
    {
      "code": 7020,
      "name": "SourceTokenDataTooLarge",
      "msg": "Source token data is too large"
    },
    {
      "code": 7021,
      "name": "InvalidTokenAdminRegistryInputsZeroAddress",
      "msg": "New Admin can not be zero address"
    },
    {
      "code": 7022,
      "name": "InvalidTokenAdminRegistryProposedAdmin",
      "msg": "An already owned registry can not be proposed"
    },
    {
      "code": 7023,
      "name": "SenderNotAllowed",
      "msg": "Sender not allowed for that destination chain"
    },
    {
      "code": 7024,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 7025,
      "name": "InvalidCcipVersionRollback",
      "msg": "Invalid rollback attempt on the CCIP version of the onramp to the destination chain"
    }
  ]
};
