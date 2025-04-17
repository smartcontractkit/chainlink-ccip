export type ExternalProgramCpiStub = {
  "version": "0.0.0-dev",
  "name": "external_program_cpi_stub",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "stubCaller",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "empty",
      "accounts": [],
      "args": []
    },
    {
      "name": "u8InstructionData",
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": "u8"
        }
      ]
    },
    {
      "name": "structInstructionData",
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": {
            "defined": "Value"
          }
        }
      ]
    },
    {
      "name": "accountRead",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "accountMut",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "stubCaller",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "bigInstructionData",
      "docs": [
        "instruction that accepts arbitrarily large instruction data."
      ],
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "noOp",
      "docs": [
        "no-op instruction that does nothing, also can be used to test maximum account references(remaining_accounts)"
      ],
      "accounts": [],
      "args": []
    },
    {
      "name": "computeHeavy",
      "accounts": [],
      "args": [
        {
          "name": "iterations",
          "type": "u32"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "value",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": "u8"
          }
        ]
      }
    }
  ]
};

export const IDL: ExternalProgramCpiStub = {
  "version": "0.0.0-dev",
  "name": "external_program_cpi_stub",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "stubCaller",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "empty",
      "accounts": [],
      "args": []
    },
    {
      "name": "u8InstructionData",
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": "u8"
        }
      ]
    },
    {
      "name": "structInstructionData",
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": {
            "defined": "Value"
          }
        }
      ]
    },
    {
      "name": "accountRead",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "accountMut",
      "accounts": [
        {
          "name": "u8Value",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "stubCaller",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "bigInstructionData",
      "docs": [
        "instruction that accepts arbitrarily large instruction data."
      ],
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "noOp",
      "docs": [
        "no-op instruction that does nothing, also can be used to test maximum account references(remaining_accounts)"
      ],
      "accounts": [],
      "args": []
    },
    {
      "name": "computeHeavy",
      "accounts": [],
      "args": [
        {
          "name": "iterations",
          "type": "u32"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "value",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": "u8"
          }
        ]
      }
    }
  ]
};
