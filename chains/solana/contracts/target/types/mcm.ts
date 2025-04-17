export type Mcm = {
  "version": "0.1.0-dev",
  "name": "mcm",
  "docs": [
    "A multi-signature contract system that supports signing many transactions targeting",
    "multiple chains with a single set of signatures. This program manages multiple multisig",
    "configurations through Program Derived Accounts (PDAs), each identified by a unique",
    "multisig_id.",
    "",
    "Key Features:",
    "- Multiple Configurations: A single deployed program instance can manage",
    "multiple independent multisig configurations",
    "- Hierarchical Signature Groups: Supports complex approval structures with",
    "nested groups and customizable quorum requirements",
    "- Merkle Tree Operations: Batches signed operations in Merkle trees for",
    "efficient verification and execution",
    "",
    "Usage Flow:",
    "1. Initialize multisig configuration with a unique ID",
    "2. Set up signer hierarchy and group structure",
    "3. Set a Merkle root with authenticated metadata and signatures",
    "4. Execute operations by providing Merkle proofs"
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initialize a new multisig configuration.",
        "",
        "Creates the foundation for a new multisig instance by initializing the core configuration",
        "PDAs and registering the multisig_id and chain_id. This is the first step in setting up",
        "a new multisig configuration.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts required for initialization:",
        "- `multisig_config`: PDA that will store the core configuration",
        "- `root_metadata`: PDA that will store the current root's metadata",
        "- `expiring_root_and_op_count`: PDA that tracks the current root and operation count",
        "- `authority`: The deployer who becomes the initial owner",
        "- `program_data`: Used to validate that the caller is the program's upgrade authority",
        "- `chain_id`: Network identifier for the chain this configuration is targeting",
        "- `multisig_id`: A unique, 32-byte identifier (left-padded) for this multisig instance",
        "",
        "# Access Control",
        "",
        "This instruction can only be called by the program's upgrade authority (typically the deployer).",
        "",
        "# Note",
        "",
        "After initialization, the owner can transfer ownership through the two-step",
        "transfer_ownership/accept_ownership process."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
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
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "chainId",
          "type": "u64"
        },
        {
          "name": "multisigId",
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
      "name": "transferOwnership",
      "docs": [
        "Propose a new owner for the multisig instance config.",
        "",
        "Only the current owner (admin) can propose a new owner.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_multisig_id`: The multisig identifier.",
        "- `proposed_owner`: The public key of the proposed new owner."
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "proposedOwner",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "docs": [
        "Accept ownership of the multisig config.",
        "",
        "The proposed new owner must call this function to assume ownership.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_multisig_id`: The multisig identifier."
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
          "name": "multisigId",
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
      "name": "setConfig",
      "docs": [
        "Set up the configuration for the multisig instance.",
        "",
        "Validates and establishes the signer hierarchy, group structure, and quorum requirements.",
        "If `clear_root` is true, it also invalidates the current Merkle root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the multisig configuration account.",
        "- `multisig_id`: The unique identifier for this multisig instance.",
        "- `signer_groups`: Vector assigning each signer to a specific group (must match signers length).",
        "- `group_quorums`: Array defining the required signatures for each group. A group with quorum=0 is disabled.",
        "- `group_parents`: Array defining the hierarchical relationship between groups, forming a tree structure.",
        "- `clear_root`: If true, invalidates the current root to prevent further operations from being executed.",
        "",
        "# Example",
        "",
        "A group structure like this:",
        "",
        "```text",
        "┌──────┐",
        "┌─►│2-of-3│◄───────┐",
        "│  └──────┘        │",
        "│        ▲         │",
        "│        │         │",
        "┌──┴───┐ ┌──┴───┐ ┌───┴────┐",
        "┌──►│1-of-2│ │2-of-2│ │signer A│",
        "│   └──────┘ └──────┘ └────────┘",
        "│       ▲      ▲  ▲",
        "│       │      │  │",
        "┌───────┴┐ ┌────┴───┐ ┌┴───────┐",
        "│signer B│ │signer C│ │signer D│",
        "└────────┘ └────────┘ └────────┘",
        "```",
        "",
        "Would be configured with:",
        "- group_quorums = [2, 1, 2, ...] (root: 2-of-3, group1: 1-of-2, group2: 2-of-2)",
        "- group_parents = [0, 0, 0, ...] (all groups under root)"
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "configSigners",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "signerGroups",
          "type": "bytes"
        },
        {
          "name": "groupQuorums",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "groupParents",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "clearRoot",
          "type": "bool"
        }
      ]
    },
    {
      "name": "setRoot",
      "docs": [
        "Set a new Merkle root that defines approved operations.",
        "",
        "This function updates the active Merkle root after verifying ECDSA signatures and validating",
        "the provided metadata against a Merkle proof.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The new Merkle root to set.",
        "- `valid_until`: timestamp until which the root remains valid.",
        "- `metadata`: Structured input containing chain_id, multisig, and operation counters.",
        "- `metadata_proof`: Merkle proof validating the metadata."
      ],
      "accounts": [
        {
          "name": "rootSignatures",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "seenSignedHashes",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "multisigConfig",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "metadata",
          "type": {
            "defined": "RootMetadataInput"
          }
        },
        {
          "name": "metadataProof",
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
    },
    {
      "name": "execute",
      "docs": [
        "Executes an operation after verifying it's authorized in the current Merkle root.",
        "",
        "This function:",
        "1. Performs extensive validation checks on the operation",
        "- Ensures the operation is within the allowed count range",
        "- Verifies chain ID matches the configured chain",
        "- Checks the root has not expired",
        "- Validates the operation's nonce against current state",
        "2. Verifies the operation's inclusion in the Merkle tree",
        "3. Executes the cross-program invocation with the multisig signer PDA",
        "",
        "# Parameters",
        "",
        "- `ctx`: Context containing operation accounts and signer information",
        "- `multisig_id`: Identifier for the multisig instance",
        "- `chain_id`: Network identifier that must match configuration",
        "- `nonce`: Operation counter that must match current state",
        "- `data`: Instruction data to be executed",
        "- `proof`: Merkle proof for operation verification",
        "",
        "# Security Considerations",
        "",
        "This instruction implements secure privilege delegation through PDA signing.",
        "The multisig's signer PDA becomes the authoritative signer for the operation,",
        "allowing controlled execution of privileged actions while maintaining the",
        "security guarantees of the Merkle root validation."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "to",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "multisigSigner",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "chainId",
          "type": "u64"
        },
        {
          "name": "nonce",
          "type": "u64"
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "proof",
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
    },
    {
      "name": "initSigners",
      "docs": [
        "Initialize the storage for signer addresses.",
        "",
        "Creates a temporary account to hold signer addresses during the multisig configuration process.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `total_signers`: The total number of signers to be added."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "totalSigners",
          "type": "u8"
        }
      ]
    },
    {
      "name": "appendSigners",
      "docs": [
        "Append a batch of signer addresses to the temporary storage.",
        "",
        "Allows adding multiple signer addresses in batches to overcome transaction size limits.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `signers_batch`: A batch of Ethereum addresses (20 bytes each) to be added as signers."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "signersBatch",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          }
        }
      ]
    },
    {
      "name": "clearSigners",
      "docs": [
        "Clear the temporary signer storage.",
        "",
        "Closes the account storing signer addresses, allowing it to be reinitialized if needed.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
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
      "name": "finalizeSigners",
      "docs": [
        "Finalize the signer configuration.",
        "",
        "Marks the signer list as complete and ready for incorporation into the multisig configuration.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
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
      "name": "initSignatures",
      "docs": [
        "Initialize storage for ECDSA signatures.",
        "",
        "Creates a temporary account to hold signatures that will validate a new Merkle root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The new Merkle root these signatures will approve.",
        "- `valid_until`: Timestamp until which the root will remain valid.",
        "- `total_signatures`: The total number of signatures to be added."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "totalSignatures",
          "type": "u8"
        }
      ]
    },
    {
      "name": "appendSignatures",
      "docs": [
        "Append a batch of ECDSA signatures to the temporary storage.",
        "",
        "Allows adding multiple signatures in batches to overcome transaction size limits.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root being approved.",
        "- `valid_until`: Timestamp until which the root will remain valid.",
        "- `signatures_batch`: A batch of ECDSA signatures to be verified."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "signaturesBatch",
          "type": {
            "vec": {
              "defined": "Signature"
            }
          }
        }
      ]
    },
    {
      "name": "clearSignatures",
      "docs": [
        "Clear the temporary signature storage.",
        "",
        "Closes the account storing signatures, allowing it to be reinitialized if needed.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root associated with the signatures.",
        "- `valid_until`: Timestamp until which the root would remain valid."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        }
      ]
    },
    {
      "name": "finalizeSignatures",
      "docs": [
        "Finalize the signature configuration.",
        "",
        "Marks the signature list as finalized and ready for verification when setting a new root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root associated with the signatures.",
        "- `valid_until`: Timestamp until which the root will remain valid."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "configSigners",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "signerAddresses",
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
            "name": "totalSigners",
            "type": "u8"
          },
          {
            "name": "isFinalized",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "multisigConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisigId",
            "type": {
              "array": [
                "u8",
                32
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
            "name": "groupQuorums",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "groupParents",
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
                "defined": "McmSigner"
              }
            }
          }
        ]
      }
    },
    {
      "name": "rootSignatures",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "totalSignatures",
            "type": "u8"
          },
          {
            "name": "isFinalized",
            "type": "bool"
          },
          {
            "name": "signatures",
            "type": {
              "vec": {
                "defined": "Signature"
              }
            }
          }
        ]
      }
    },
    {
      "name": "rootMetadata",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisig",
            "type": "publicKey"
          },
          {
            "name": "preOpCount",
            "type": "u64"
          },
          {
            "name": "postOpCount",
            "type": "u64"
          },
          {
            "name": "overridePreviousRoot",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "expiringRootAndOpCount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "root",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "validUntil",
            "type": "u32"
          },
          {
            "name": "opCount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "seenSignedHash",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "seen",
            "type": "bool"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "Signature",
      "docs": [
        "ECDSA signature with components used in Ethereum signature verification."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "v",
            "type": "u8"
          },
          {
            "name": "r",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "s",
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
      "name": "McmSigner",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "evmAddress",
            "type": {
              "array": [
                "u8",
                20
              ]
            }
          },
          {
            "name": "index",
            "type": "u8"
          },
          {
            "name": "group",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "RootMetadataInput",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisig",
            "type": "publicKey"
          },
          {
            "name": "preOpCount",
            "type": "u64"
          },
          {
            "name": "postOpCount",
            "type": "u64"
          },
          {
            "name": "overridePreviousRoot",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "McmError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "InvalidInputs"
          },
          {
            "name": "Overflow"
          },
          {
            "name": "WrongMultiSig"
          },
          {
            "name": "WrongChainId"
          },
          {
            "name": "FailedEcdsaRecover"
          },
          {
            "name": "SignersNotFinalized"
          },
          {
            "name": "SignersAlreadyFinalized"
          },
          {
            "name": "SignaturesAlreadyFinalized"
          },
          {
            "name": "SignatureCountMismatch"
          },
          {
            "name": "TooManySignatures"
          },
          {
            "name": "SignaturesNotFinalized"
          },
          {
            "name": "MismatchedInputSignerVectorsLength"
          },
          {
            "name": "OutOfBoundsNumOfSigners"
          },
          {
            "name": "MismatchedInputGroupArraysLength"
          },
          {
            "name": "GroupTreeNotWellFormed"
          },
          {
            "name": "SignerInDisabledGroup"
          },
          {
            "name": "OutOfBoundsGroupQuorum"
          },
          {
            "name": "SignersAddressesMustBeStrictlyIncreasing"
          },
          {
            "name": "SignedHashAlreadySeen"
          },
          {
            "name": "InvalidSigner"
          },
          {
            "name": "MissingConfig"
          },
          {
            "name": "InsufficientSigners"
          },
          {
            "name": "ValidUntilHasAlreadyPassed"
          },
          {
            "name": "ProofCannotBeVerified"
          },
          {
            "name": "PendingOps"
          },
          {
            "name": "WrongPreOpCount"
          },
          {
            "name": "WrongPostOpCount"
          },
          {
            "name": "PostOpCountReached"
          },
          {
            "name": "RootExpired"
          },
          {
            "name": "WrongNonce"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "NewRoot",
      "fields": [
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "validUntil",
          "type": "u32",
          "index": false
        },
        {
          "name": "metadataChainId",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataMultisig",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "metadataPreOpCount",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataPostOpCount",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataOverridePreviousRoot",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "groupParents",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "groupQuorums",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "isRootCleared",
          "type": "bool",
          "index": false
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "defined": "McmSigner"
            }
          },
          "index": false
        }
      ]
    },
    {
      "name": "OpExecuted",
      "fields": [
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    }
  ]
};

export const IDL: Mcm = {
  "version": "0.1.0-dev",
  "name": "mcm",
  "docs": [
    "A multi-signature contract system that supports signing many transactions targeting",
    "multiple chains with a single set of signatures. This program manages multiple multisig",
    "configurations through Program Derived Accounts (PDAs), each identified by a unique",
    "multisig_id.",
    "",
    "Key Features:",
    "- Multiple Configurations: A single deployed program instance can manage",
    "multiple independent multisig configurations",
    "- Hierarchical Signature Groups: Supports complex approval structures with",
    "nested groups and customizable quorum requirements",
    "- Merkle Tree Operations: Batches signed operations in Merkle trees for",
    "efficient verification and execution",
    "",
    "Usage Flow:",
    "1. Initialize multisig configuration with a unique ID",
    "2. Set up signer hierarchy and group structure",
    "3. Set a Merkle root with authenticated metadata and signatures",
    "4. Execute operations by providing Merkle proofs"
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "Initialize a new multisig configuration.",
        "",
        "Creates the foundation for a new multisig instance by initializing the core configuration",
        "PDAs and registering the multisig_id and chain_id. This is the first step in setting up",
        "a new multisig configuration.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts required for initialization:",
        "- `multisig_config`: PDA that will store the core configuration",
        "- `root_metadata`: PDA that will store the current root's metadata",
        "- `expiring_root_and_op_count`: PDA that tracks the current root and operation count",
        "- `authority`: The deployer who becomes the initial owner",
        "- `program_data`: Used to validate that the caller is the program's upgrade authority",
        "- `chain_id`: Network identifier for the chain this configuration is targeting",
        "- `multisig_id`: A unique, 32-byte identifier (left-padded) for this multisig instance",
        "",
        "# Access Control",
        "",
        "This instruction can only be called by the program's upgrade authority (typically the deployer).",
        "",
        "# Note",
        "",
        "After initialization, the owner can transfer ownership through the two-step",
        "transfer_ownership/accept_ownership process."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
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
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "chainId",
          "type": "u64"
        },
        {
          "name": "multisigId",
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
      "name": "transferOwnership",
      "docs": [
        "Propose a new owner for the multisig instance config.",
        "",
        "Only the current owner (admin) can propose a new owner.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_multisig_id`: The multisig identifier.",
        "- `proposed_owner`: The public key of the proposed new owner."
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "proposedOwner",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "acceptOwnership",
      "docs": [
        "Accept ownership of the multisig config.",
        "",
        "The proposed new owner must call this function to assume ownership.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_multisig_id`: The multisig identifier."
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
          "name": "multisigId",
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
      "name": "setConfig",
      "docs": [
        "Set up the configuration for the multisig instance.",
        "",
        "Validates and establishes the signer hierarchy, group structure, and quorum requirements.",
        "If `clear_root` is true, it also invalidates the current Merkle root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the multisig configuration account.",
        "- `multisig_id`: The unique identifier for this multisig instance.",
        "- `signer_groups`: Vector assigning each signer to a specific group (must match signers length).",
        "- `group_quorums`: Array defining the required signatures for each group. A group with quorum=0 is disabled.",
        "- `group_parents`: Array defining the hierarchical relationship between groups, forming a tree structure.",
        "- `clear_root`: If true, invalidates the current root to prevent further operations from being executed.",
        "",
        "# Example",
        "",
        "A group structure like this:",
        "",
        "```text",
        "┌──────┐",
        "┌─►│2-of-3│◄───────┐",
        "│  └──────┘        │",
        "│        ▲         │",
        "│        │         │",
        "┌──┴───┐ ┌──┴───┐ ┌───┴────┐",
        "┌──►│1-of-2│ │2-of-2│ │signer A│",
        "│   └──────┘ └──────┘ └────────┘",
        "│       ▲      ▲  ▲",
        "│       │      │  │",
        "┌───────┴┐ ┌────┴───┐ ┌┴───────┐",
        "│signer B│ │signer C│ │signer D│",
        "└────────┘ └────────┘ └────────┘",
        "```",
        "",
        "Would be configured with:",
        "- group_quorums = [2, 1, 2, ...] (root: 2-of-3, group1: 1-of-2, group2: 2-of-2)",
        "- group_parents = [0, 0, 0, ...] (all groups under root)"
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "configSigners",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "signerGroups",
          "type": "bytes"
        },
        {
          "name": "groupQuorums",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "groupParents",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "clearRoot",
          "type": "bool"
        }
      ]
    },
    {
      "name": "setRoot",
      "docs": [
        "Set a new Merkle root that defines approved operations.",
        "",
        "This function updates the active Merkle root after verifying ECDSA signatures and validating",
        "the provided metadata against a Merkle proof.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The new Merkle root to set.",
        "- `valid_until`: timestamp until which the root remains valid.",
        "- `metadata`: Structured input containing chain_id, multisig, and operation counters.",
        "- `metadata_proof`: Merkle proof validating the metadata."
      ],
      "accounts": [
        {
          "name": "rootSignatures",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "seenSignedHashes",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "multisigConfig",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "metadata",
          "type": {
            "defined": "RootMetadataInput"
          }
        },
        {
          "name": "metadataProof",
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
    },
    {
      "name": "execute",
      "docs": [
        "Executes an operation after verifying it's authorized in the current Merkle root.",
        "",
        "This function:",
        "1. Performs extensive validation checks on the operation",
        "- Ensures the operation is within the allowed count range",
        "- Verifies chain ID matches the configured chain",
        "- Checks the root has not expired",
        "- Validates the operation's nonce against current state",
        "2. Verifies the operation's inclusion in the Merkle tree",
        "3. Executes the cross-program invocation with the multisig signer PDA",
        "",
        "# Parameters",
        "",
        "- `ctx`: Context containing operation accounts and signer information",
        "- `multisig_id`: Identifier for the multisig instance",
        "- `chain_id`: Network identifier that must match configuration",
        "- `nonce`: Operation counter that must match current state",
        "- `data`: Instruction data to be executed",
        "- `proof`: Merkle proof for operation verification",
        "",
        "# Security Considerations",
        "",
        "This instruction implements secure privilege delegation through PDA signing.",
        "The multisig's signer PDA becomes the authoritative signer for the operation,",
        "allowing controlled execution of privileged actions while maintaining the",
        "security guarantees of the Merkle root validation."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "rootMetadata",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "expiringRootAndOpCount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "to",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "multisigSigner",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "chainId",
          "type": "u64"
        },
        {
          "name": "nonce",
          "type": "u64"
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "proof",
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
    },
    {
      "name": "initSigners",
      "docs": [
        "Initialize the storage for signer addresses.",
        "",
        "Creates a temporary account to hold signer addresses during the multisig configuration process.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `total_signers`: The total number of signers to be added."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "totalSigners",
          "type": "u8"
        }
      ]
    },
    {
      "name": "appendSigners",
      "docs": [
        "Append a batch of signer addresses to the temporary storage.",
        "",
        "Allows adding multiple signer addresses in batches to overcome transaction size limits.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `signers_batch`: A batch of Ethereum addresses (20 bytes each) to be added as signers."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "signersBatch",
          "type": {
            "vec": {
              "array": [
                "u8",
                20
              ]
            }
          }
        }
      ]
    },
    {
      "name": "clearSigners",
      "docs": [
        "Clear the temporary signer storage.",
        "",
        "Closes the account storing signer addresses, allowing it to be reinitialized if needed.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
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
      "name": "finalizeSigners",
      "docs": [
        "Finalize the signer configuration.",
        "",
        "Marks the signer list as complete and ready for incorporation into the multisig configuration.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier."
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "configSigners",
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
          "name": "multisigId",
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
      "name": "initSignatures",
      "docs": [
        "Initialize storage for ECDSA signatures.",
        "",
        "Creates a temporary account to hold signatures that will validate a new Merkle root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The new Merkle root these signatures will approve.",
        "- `valid_until`: Timestamp until which the root will remain valid.",
        "- `total_signatures`: The total number of signatures to be added."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "totalSignatures",
          "type": "u8"
        }
      ]
    },
    {
      "name": "appendSignatures",
      "docs": [
        "Append a batch of ECDSA signatures to the temporary storage.",
        "",
        "Allows adding multiple signatures in batches to overcome transaction size limits.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root being approved.",
        "- `valid_until`: Timestamp until which the root will remain valid.",
        "- `signatures_batch`: A batch of ECDSA signatures to be verified."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        },
        {
          "name": "signaturesBatch",
          "type": {
            "vec": {
              "defined": "Signature"
            }
          }
        }
      ]
    },
    {
      "name": "clearSignatures",
      "docs": [
        "Clear the temporary signature storage.",
        "",
        "Closes the account storing signatures, allowing it to be reinitialized if needed.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root associated with the signatures.",
        "- `valid_until`: Timestamp until which the root would remain valid."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        }
      ]
    },
    {
      "name": "finalizeSignatures",
      "docs": [
        "Finalize the signature configuration.",
        "",
        "Marks the signature list as finalized and ready for verification when setting a new root.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing required accounts.",
        "- `multisig_id`: The multisig instance identifier.",
        "- `root`: The Merkle root associated with the signatures.",
        "- `valid_until`: Timestamp until which the root will remain valid."
      ],
      "accounts": [
        {
          "name": "signatures",
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
          "name": "multisigId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "validUntil",
          "type": "u32"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "configSigners",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "signerAddresses",
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
            "name": "totalSigners",
            "type": "u8"
          },
          {
            "name": "isFinalized",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "multisigConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisigId",
            "type": {
              "array": [
                "u8",
                32
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
            "name": "groupQuorums",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "groupParents",
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
                "defined": "McmSigner"
              }
            }
          }
        ]
      }
    },
    {
      "name": "rootSignatures",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "totalSignatures",
            "type": "u8"
          },
          {
            "name": "isFinalized",
            "type": "bool"
          },
          {
            "name": "signatures",
            "type": {
              "vec": {
                "defined": "Signature"
              }
            }
          }
        ]
      }
    },
    {
      "name": "rootMetadata",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisig",
            "type": "publicKey"
          },
          {
            "name": "preOpCount",
            "type": "u64"
          },
          {
            "name": "postOpCount",
            "type": "u64"
          },
          {
            "name": "overridePreviousRoot",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "expiringRootAndOpCount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "root",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "validUntil",
            "type": "u32"
          },
          {
            "name": "opCount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "seenSignedHash",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "seen",
            "type": "bool"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "Signature",
      "docs": [
        "ECDSA signature with components used in Ethereum signature verification."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "v",
            "type": "u8"
          },
          {
            "name": "r",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "s",
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
      "name": "McmSigner",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "evmAddress",
            "type": {
              "array": [
                "u8",
                20
              ]
            }
          },
          {
            "name": "index",
            "type": "u8"
          },
          {
            "name": "group",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "RootMetadataInput",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisig",
            "type": "publicKey"
          },
          {
            "name": "preOpCount",
            "type": "u64"
          },
          {
            "name": "postOpCount",
            "type": "u64"
          },
          {
            "name": "overridePreviousRoot",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "McmError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "InvalidInputs"
          },
          {
            "name": "Overflow"
          },
          {
            "name": "WrongMultiSig"
          },
          {
            "name": "WrongChainId"
          },
          {
            "name": "FailedEcdsaRecover"
          },
          {
            "name": "SignersNotFinalized"
          },
          {
            "name": "SignersAlreadyFinalized"
          },
          {
            "name": "SignaturesAlreadyFinalized"
          },
          {
            "name": "SignatureCountMismatch"
          },
          {
            "name": "TooManySignatures"
          },
          {
            "name": "SignaturesNotFinalized"
          },
          {
            "name": "MismatchedInputSignerVectorsLength"
          },
          {
            "name": "OutOfBoundsNumOfSigners"
          },
          {
            "name": "MismatchedInputGroupArraysLength"
          },
          {
            "name": "GroupTreeNotWellFormed"
          },
          {
            "name": "SignerInDisabledGroup"
          },
          {
            "name": "OutOfBoundsGroupQuorum"
          },
          {
            "name": "SignersAddressesMustBeStrictlyIncreasing"
          },
          {
            "name": "SignedHashAlreadySeen"
          },
          {
            "name": "InvalidSigner"
          },
          {
            "name": "MissingConfig"
          },
          {
            "name": "InsufficientSigners"
          },
          {
            "name": "ValidUntilHasAlreadyPassed"
          },
          {
            "name": "ProofCannotBeVerified"
          },
          {
            "name": "PendingOps"
          },
          {
            "name": "WrongPreOpCount"
          },
          {
            "name": "WrongPostOpCount"
          },
          {
            "name": "PostOpCountReached"
          },
          {
            "name": "RootExpired"
          },
          {
            "name": "WrongNonce"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "NewRoot",
      "fields": [
        {
          "name": "root",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "validUntil",
          "type": "u32",
          "index": false
        },
        {
          "name": "metadataChainId",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataMultisig",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "metadataPreOpCount",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataPostOpCount",
          "type": "u64",
          "index": false
        },
        {
          "name": "metadataOverridePreviousRoot",
          "type": "bool",
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "groupParents",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "groupQuorums",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "isRootCleared",
          "type": "bool",
          "index": false
        },
        {
          "name": "signers",
          "type": {
            "vec": {
              "defined": "McmSigner"
            }
          },
          "index": false
        }
      ]
    },
    {
      "name": "OpExecuted",
      "fields": [
        {
          "name": "nonce",
          "type": "u64",
          "index": false
        },
        {
          "name": "to",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    }
  ]
};
