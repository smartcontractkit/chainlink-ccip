/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/rmn_remote.json`.
 */
export type RmnRemote = {
  "address": "RmnXLft1mSEwDgMKu2okYuHkiazxntFFcZFrrcXxYg7",
  "metadata": {
    "name": "rmnRemote",
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
        "Shared func signature with other programs.",
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
      "name": "curse",
      "docs": [
        "Curses an abstract subject. If the subject is CurseSubject::GLOBAL,",
        "the entire chain will be cursed.",
        "",
        "Only the CCIP Admin may perform this operation",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for adding a new curse.",
        "* `subject` - The subject to curse."
      ],
      "discriminator": [
        10,
        127,
        226,
        227,
        138,
        3,
        192,
        73
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "curses",
          "writable": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "subject",
          "type": {
            "defined": {
              "name": "curseSubject"
            }
          }
        }
      ]
    },
    {
      "name": "initialize",
      "docs": [
        "Initializes the Rmn Remote contract.",
        "",
        "The initialization is responsibility of Admin, nothing more than calling this method should be done first.",
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
          "name": "config",
          "writable": true
        },
        {
          "name": "curses",
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
      "name": "setDefaultCodeVersion",
      "docs": [
        "Sets the default code version to be used. This is then used by the slim routing layer to determine",
        "which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.",
        "",
        "Shared func signature with other programs.",
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
          "name": "curses"
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
      "name": "transferOwnership",
      "docs": [
        "Transfers the ownership of the fee quoter to a new proposed owner.",
        "",
        "Shared func signature with other programs.",
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
          "name": "curses"
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
      "name": "uncurse",
      "docs": [
        "Uncurses an abstract subject. If the subject is CurseSubject::GLOBAL,",
        "the entire chain curse will be lifted. (note that any other specific",
        "subject curses will remain active.)",
        "",
        "Only the CCIP Admin may perform this operation",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for removing a curse.",
        "* `subject` - The subject to uncurse."
      ],
      "discriminator": [
        235,
        227,
        61,
        167,
        15,
        109,
        129,
        79
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        },
        {
          "name": "curses",
          "writable": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": [
        {
          "name": "subject",
          "type": {
            "defined": {
              "name": "curseSubject"
            }
          }
        }
      ]
    },
    {
      "name": "verifyNotCursed",
      "docs": [
        "Verifies that the subject is not cursed AND that this chain is not globally cursed.",
        "In case either of those assumptions fail, the instruction reverts.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required to inspect curses.",
        "* `subject` - The subject to verify. Note that this instruction will revert if the chain",
        "is globally cursed too, even if the provided subject is not explicitly cursed."
      ],
      "discriminator": [
        86,
        200,
        58,
        143,
        7,
        109,
        155,
        125
      ],
      "accounts": [
        {
          "name": "curses"
        },
        {
          "name": "config"
        }
      ],
      "args": [
        {
          "name": "subject",
          "type": {
            "defined": {
              "name": "curseSubject"
            }
          }
        }
      ]
    }
  ],
  "accounts": [
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
      "name": "curses",
      "discriminator": [
        129,
        28,
        49,
        58,
        74,
        237,
        146,
        202
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
      "name": "subjectCursed",
      "discriminator": [
        64,
        234,
        236,
        62,
        237,
        179,
        9,
        192
      ]
    },
    {
      "name": "subjectUncursed",
      "discriminator": [
        238,
        50,
        186,
        246,
        156,
        119,
        251,
        250
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9001,
      "name": "subjectIsAlreadyCursed",
      "msg": "Subject is already cursed"
    },
    {
      "code": 9002,
      "name": "subjectWasNotCursed",
      "msg": "Subject was not cursed"
    },
    {
      "code": 9003,
      "name": "redundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9004,
      "name": "invalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9005,
      "name": "subjectCursed",
      "msg": "The subject is actively cursed"
    },
    {
      "code": 9006,
      "name": "globallyCursed",
      "msg": "This chain is globally cursed"
    },
    {
      "code": 9007,
      "name": "invalidCodeVersion",
      "msg": "Invalid code version"
    }
  ],
  "types": [
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
      "name": "curseSubject",
      "docs": [
        "Abstract curse subject.",
        "",
        "In particular, a curse subject can be constructed from a chain",
        "selector to signify that any lane involving that chain as `destination` or `source` is",
        "cursed.",
        "",
        "The above is not exhaustive: there may be other ways to define subjects."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": {
              "array": [
                "u8",
                16
              ]
            }
          }
        ]
      }
    },
    {
      "name": "curses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "cursedSubjects",
            "type": {
              "vec": {
                "defined": {
                  "name": "curseSubject"
                }
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
      "name": "subjectCursed",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "subject",
            "type": {
              "defined": {
                "name": "curseSubject"
              }
            }
          }
        ]
      }
    },
    {
      "name": "subjectUncursed",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "subject",
            "type": {
              "defined": {
                "name": "curseSubject"
              }
            }
          }
        ]
      }
    }
  ]
};
