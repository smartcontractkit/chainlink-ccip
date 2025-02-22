export type Mcm = {
  "version": "0.1.0-dev",
  "name": "mcm",
  "docs": [
    "This is mcm program supporting multiple instances of multisig configuration",
    "A single deployed program manages multiple multisig states(configurations) identified by multisig_id"
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "initialize a new multisig configuration, store the chain_id and multisig_id",
        "multisig_id is a unique identifier for the multisig configuration(32 bytes, left-padded)"
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
            "name": "signatures",
            "type": {
              "vec": {
                "defined": "Signature"
              }
            }
          },
          {
            "name": "isFinalized",
            "type": "bool"
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
    "This is mcm program supporting multiple instances of multisig configuration",
    "A single deployed program manages multiple multisig states(configurations) identified by multisig_id"
  ],
  "instructions": [
    {
      "name": "initialize",
      "docs": [
        "initialize a new multisig configuration, store the chain_id and multisig_id",
        "multisig_id is a unique identifier for the multisig configuration(32 bytes, left-padded)"
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
            "name": "signatures",
            "type": {
              "vec": {
                "defined": "Signature"
              }
            }
          },
          {
            "name": "isFinalized",
            "type": "bool"
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
