export type RmnRemote = {
  "version": "0.1.0-dev",
  "name": "rmn_remote",
  "instructions": [
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
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "curses",
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
        "Transfers the ownership of the fee quoter to a new proposed owner.",
        "",
        "Shared func signature with other programs.",
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
          "name": "curses",
          "isMut": false,
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
        "Shared func signature with other programs.",
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
        "Shared func signature with other programs.",
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
          "name": "curses",
          "isMut": false,
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
      "accounts": [
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
          "name": "curses",
          "isMut": true,
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
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          }
        }
      ]
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
      "accounts": [
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
          "name": "curses",
          "isMut": true,
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
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
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
      "accounts": [
        {
          "name": "curses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
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
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
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
                "defined": "CurseSubject"
              }
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CurseSubject",
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
          "name": "defaultCodeVersion",
          "type": {
            "defined": "CodeVersion"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SubjectCursed",
      "fields": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SubjectUncursed",
      "fields": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9001,
      "name": "SubjectIsAlreadyCursed",
      "msg": "Subject is already cursed"
    },
    {
      "code": 9002,
      "name": "SubjectWasNotCursed",
      "msg": "Subject was not cursed"
    },
    {
      "code": 9003,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9004,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9005,
      "name": "SubjectCursed",
      "msg": "The subject is actively cursed"
    },
    {
      "code": 9006,
      "name": "GloballyCursed",
      "msg": "This chain is globally cursed"
    },
    {
      "code": 9007,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    }
  ]
};

export const IDL: RmnRemote = {
  "version": "0.1.0-dev",
  "name": "rmn_remote",
  "instructions": [
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
      "accounts": [
        {
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "curses",
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
        "Transfers the ownership of the fee quoter to a new proposed owner.",
        "",
        "Shared func signature with other programs.",
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
          "name": "curses",
          "isMut": false,
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
        "Shared func signature with other programs.",
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
        "Shared func signature with other programs.",
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
          "name": "curses",
          "isMut": false,
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
      "accounts": [
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
          "name": "curses",
          "isMut": true,
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
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          }
        }
      ]
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
      "accounts": [
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
          "name": "curses",
          "isMut": true,
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
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
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
      "accounts": [
        {
          "name": "curses",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
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
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
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
                "defined": "CurseSubject"
              }
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "CurseSubject",
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
          "name": "defaultCodeVersion",
          "type": {
            "defined": "CodeVersion"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SubjectCursed",
      "fields": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SubjectUncursed",
      "fields": [
        {
          "name": "subject",
          "type": {
            "defined": "CurseSubject"
          },
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 9000,
      "name": "Unauthorized",
      "msg": "The signer is unauthorized"
    },
    {
      "code": 9001,
      "name": "SubjectIsAlreadyCursed",
      "msg": "Subject is already cursed"
    },
    {
      "code": 9002,
      "name": "SubjectWasNotCursed",
      "msg": "Subject was not cursed"
    },
    {
      "code": 9003,
      "name": "RedundantOwnerProposal",
      "msg": "Proposed owner is the current owner"
    },
    {
      "code": 9004,
      "name": "InvalidVersion",
      "msg": "Invalid version of the onchain state"
    },
    {
      "code": 9005,
      "name": "SubjectCursed",
      "msg": "The subject is actively cursed"
    },
    {
      "code": 9006,
      "name": "GloballyCursed",
      "msg": "This chain is globally cursed"
    },
    {
      "code": 9007,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    }
  ]
};
