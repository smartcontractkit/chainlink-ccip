/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/mcm.json`.
 */
export type Mcm = {
  "address": "5vNJx78mz7KVMjhuipyr9jKBKcMrKYGdjGkgE4LUmjKk",
  "metadata": {
    "name": "mcm",
    "version": "0.1.0-dev",
    "spec": "0.1.0",
    "description": "SVM implementation of ManyChainMultiSig"
  },
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
      "discriminator": [
        195,
        112,
        164,
        69,
        37,
        137,
        198,
        54
      ],
      "accounts": [
        {
          "name": "signatures",
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
              "defined": {
                "name": "signature"
              }
            }
          }
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
      "discriminator": [
        238,
        209,
        251,
        39,
        41,
        241,
        146,
        25
      ],
      "accounts": [
        {
          "name": "multisigConfig"
        },
        {
          "name": "configSigners",
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
      "discriminator": [
        80,
        0,
        39,
        255,
        46,
        165,
        193,
        109
      ],
      "accounts": [
        {
          "name": "signatures",
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
      "discriminator": [
        90,
        140,
        170,
        146,
        128,
        75,
        100,
        175
      ],
      "accounts": [
        {
          "name": "multisigConfig"
        },
        {
          "name": "configSigners",
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
          "name": "multisigConfig",
          "writable": true
        },
        {
          "name": "rootMetadata"
        },
        {
          "name": "expiringRootAndOpCount",
          "writable": true
        },
        {
          "name": "to"
        },
        {
          "name": "multisigSigner"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
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
      "discriminator": [
        77,
        138,
        152,
        199,
        37,
        141,
        189,
        159
      ],
      "accounts": [
        {
          "name": "signatures",
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
      "discriminator": [
        49,
        254,
        154,
        226,
        137,
        199,
        120,
        63
      ],
      "accounts": [
        {
          "name": "multisigConfig"
        },
        {
          "name": "configSigners",
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
      "discriminator": [
        190,
        120,
        207,
        36,
        26,
        58,
        196,
        13
      ],
      "accounts": [
        {
          "name": "signatures",
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
      "discriminator": [
        102,
        182,
        129,
        16,
        138,
        142,
        223,
        196
      ],
      "accounts": [
        {
          "name": "multisigConfig"
        },
        {
          "name": "configSigners",
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
          "name": "multisigConfig",
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
        },
        {
          "name": "rootMetadata",
          "writable": true
        },
        {
          "name": "expiringRootAndOpCount",
          "writable": true
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
      "discriminator": [
        108,
        158,
        154,
        175,
        212,
        98,
        52,
        66
      ],
      "accounts": [
        {
          "name": "multisigConfig",
          "writable": true
        },
        {
          "name": "configSigners",
          "writable": true
        },
        {
          "name": "rootMetadata",
          "writable": true
        },
        {
          "name": "expiringRootAndOpCount",
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
      "discriminator": [
        183,
        49,
        10,
        206,
        168,
        183,
        131,
        67
      ],
      "accounts": [
        {
          "name": "rootSignatures",
          "writable": true
        },
        {
          "name": "rootMetadata",
          "writable": true
        },
        {
          "name": "seenSignedHashes",
          "writable": true
        },
        {
          "name": "expiringRootAndOpCount",
          "writable": true
        },
        {
          "name": "multisigConfig"
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
            "defined": {
              "name": "rootMetadataInput"
            }
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
          "type": "pubkey"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "configSigners",
      "discriminator": [
        147,
        137,
        80,
        98,
        50,
        225,
        190,
        163
      ]
    },
    {
      "name": "expiringRootAndOpCount",
      "discriminator": [
        196,
        176,
        71,
        210,
        134,
        228,
        202,
        75
      ]
    },
    {
      "name": "multisigConfig",
      "discriminator": [
        44,
        62,
        172,
        225,
        246,
        3,
        178,
        33
      ]
    },
    {
      "name": "rootMetadata",
      "discriminator": [
        125,
        211,
        89,
        150,
        221,
        6,
        141,
        205
      ]
    },
    {
      "name": "rootSignatures",
      "discriminator": [
        21,
        186,
        10,
        33,
        117,
        215,
        246,
        76
      ]
    },
    {
      "name": "seenSignedHash",
      "discriminator": [
        229,
        115,
        10,
        185,
        39,
        100,
        210,
        151
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
      "name": "newRoot",
      "discriminator": [
        210,
        25,
        187,
        118,
        40,
        42,
        61,
        119
      ]
    },
    {
      "name": "opExecuted",
      "discriminator": [
        221,
        15,
        212,
        29,
        35,
        252,
        255,
        78
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "invalidInputs",
      "msg": "Invalid inputs"
    },
    {
      "code": 6001,
      "name": "overflow",
      "msg": "overflow occurred."
    },
    {
      "code": 6002,
      "name": "wrongMultiSig",
      "msg": "Invalid multisig"
    },
    {
      "code": 6003,
      "name": "wrongChainId",
      "msg": "Invalid chainID"
    },
    {
      "code": 6004,
      "name": "failedEcdsaRecover",
      "msg": "Failed ECDSA recover"
    },
    {
      "code": 6005,
      "name": "signersNotFinalized",
      "msg": "Config signers not finalized"
    },
    {
      "code": 6006,
      "name": "signersAlreadyFinalized",
      "msg": "Config signers already finalized"
    },
    {
      "code": 6007,
      "name": "signaturesAlreadyFinalized",
      "msg": "Signatures already finalized"
    },
    {
      "code": 6008,
      "name": "signatureCountMismatch",
      "msg": "Uploaded signatures count mismatch"
    },
    {
      "code": 6009,
      "name": "tooManySignatures",
      "msg": "Too many signatures"
    },
    {
      "code": 6010,
      "name": "signaturesNotFinalized",
      "msg": "Signatures not finalized"
    },
    {
      "code": 6200,
      "name": "mismatchedInputSignerVectorsLength",
      "msg": "The input vectors for signer addresses and signer groups must have the same length"
    },
    {
      "code": 6201,
      "name": "outOfBoundsNumOfSigners",
      "msg": "The number of signers is 0 or greater than MAX_NUM_SIGNERS"
    },
    {
      "code": 6202,
      "name": "mismatchedInputGroupArraysLength",
      "msg": "The input arrays for group parents and group quorums must be of length NUM_GROUPS"
    },
    {
      "code": 6203,
      "name": "groupTreeNotWellFormed",
      "msg": "the group tree isn't well-formed."
    },
    {
      "code": 6204,
      "name": "signerInDisabledGroup",
      "msg": "a disabled group contains a signer."
    },
    {
      "code": 6205,
      "name": "outOfBoundsGroupQuorum",
      "msg": "the quorum of some group is larger than the number of signers in it."
    },
    {
      "code": 6206,
      "name": "signersAddressesMustBeStrictlyIncreasing",
      "msg": "the signers' addresses are not a strictly increasing monotone sequence."
    },
    {
      "code": 6207,
      "name": "signedHashAlreadySeen",
      "msg": "The combination of signature and valid_until has already been seen"
    },
    {
      "code": 6208,
      "name": "invalidSigner",
      "msg": "Invalid signer"
    },
    {
      "code": 6209,
      "name": "missingConfig",
      "msg": "Missing configuration"
    },
    {
      "code": 6210,
      "name": "insufficientSigners",
      "msg": "Insufficient signers"
    },
    {
      "code": 6211,
      "name": "validUntilHasAlreadyPassed",
      "msg": "Valid until has already passed"
    },
    {
      "code": 6212,
      "name": "proofCannotBeVerified",
      "msg": "Proof cannot be verified"
    },
    {
      "code": 6213,
      "name": "pendingOps",
      "msg": "Pending operations"
    },
    {
      "code": 6214,
      "name": "wrongPreOpCount",
      "msg": "Wrong pre-operation count"
    },
    {
      "code": 6215,
      "name": "wrongPostOpCount",
      "msg": "Wrong post-operation count"
    },
    {
      "code": 6216,
      "name": "postOpCountReached",
      "msg": "Post-operation count reached"
    },
    {
      "code": 6217,
      "name": "rootExpired",
      "msg": "Root expired"
    },
    {
      "code": 6218,
      "name": "wrongNonce",
      "msg": "Wrong nonce"
    }
  ],
  "types": [
    {
      "name": "configSet",
      "docs": [
        "@dev Emitted when a new config is set."
      ],
      "type": {
        "kind": "struct",
        "fields": [
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
            "name": "groupQuorums",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "isRootCleared",
            "type": "bool"
          },
          {
            "name": "signers",
            "type": {
              "vec": {
                "defined": {
                  "name": "mcmSigner"
                }
              }
            }
          }
        ]
      }
    },
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
      "name": "mcmSigner",
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
            "type": "pubkey"
          },
          {
            "name": "proposedOwner",
            "type": "pubkey"
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
                "defined": {
                  "name": "mcmSigner"
                }
              }
            }
          }
        ]
      }
    },
    {
      "name": "newRoot",
      "docs": [
        "@dev Emitted when a new root is set."
      ],
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
            "name": "metadataChainId",
            "type": "u64"
          },
          {
            "name": "metadataMultisig",
            "type": "pubkey"
          },
          {
            "name": "metadataPreOpCount",
            "type": "u64"
          },
          {
            "name": "metadataPostOpCount",
            "type": "u64"
          },
          {
            "name": "metadataOverridePreviousRoot",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "opExecuted",
      "docs": [
        "@dev Emitted when an op gets successfully executed."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "nonce",
            "type": "u64"
          },
          {
            "name": "to",
            "type": "pubkey"
          },
          {
            "name": "data",
            "type": "bytes"
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
            "type": "pubkey"
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
      "name": "rootMetadataInput",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainId",
            "type": "u64"
          },
          {
            "name": "multisig",
            "type": "pubkey"
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
                "defined": {
                  "name": "signature"
                }
              }
            }
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
    },
    {
      "name": "signature",
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
    }
  ]
};
