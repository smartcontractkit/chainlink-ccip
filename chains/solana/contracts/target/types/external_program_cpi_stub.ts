/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/external_program_cpi_stub.json`.
 */
export type ExternalProgramCpiStub = {
  "address": "2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm",
  "metadata": {
    "name": "externalProgramCpiStub",
    "version": "0.0.0-dev",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
    {
      "name": "accountMut",
      "discriminator": [
        12,
        2,
        137,
        19,
        22,
        235,
        144,
        70
      ],
      "accounts": [
        {
          "name": "u8Value",
          "writable": true
        },
        {
          "name": "stubCaller",
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": []
    },
    {
      "name": "accountRead",
      "discriminator": [
        79,
        53,
        80,
        124,
        182,
        81,
        206,
        85
      ],
      "accounts": [
        {
          "name": "u8Value"
        }
      ],
      "args": []
    },
    {
      "name": "bigInstructionData",
      "docs": [
        "instruction that accepts arbitrarily large instruction data."
      ],
      "discriminator": [
        250,
        215,
        200,
        174,
        42,
        217,
        129,
        182
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
      "name": "computeHeavy",
      "discriminator": [
        114,
        240,
        76,
        252,
        181,
        175,
        211,
        78
      ],
      "accounts": [],
      "args": [
        {
          "name": "iterations",
          "type": "u32"
        }
      ]
    },
    {
      "name": "empty",
      "discriminator": [
        214,
        44,
        4,
        247,
        12,
        41,
        217,
        110
      ],
      "accounts": [],
      "args": []
    },
    {
      "name": "initialize",
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
          "name": "u8Value",
          "writable": true
        },
        {
          "name": "stubCaller",
          "writable": true,
          "signer": true
        },
        {
          "name": "systemProgram"
        }
      ],
      "args": []
    },
    {
      "name": "noOp",
      "docs": [
        "no-op instruction that does nothing, also can be used to test maximum account references(remaining_accounts)"
      ],
      "discriminator": [
        36,
        122,
        159,
        43,
        166,
        240,
        121,
        88
      ],
      "accounts": [],
      "args": []
    },
    {
      "name": "structInstructionData",
      "discriminator": [
        132,
        84,
        80,
        47,
        117,
        198,
        94,
        67
      ],
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": {
            "defined": {
              "name": "value"
            }
          }
        }
      ]
    },
    {
      "name": "u8InstructionData",
      "discriminator": [
        17,
        175,
        156,
        253,
        91,
        173,
        26,
        228
      ],
      "accounts": [],
      "args": [
        {
          "name": "data",
          "type": "u8"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "value",
      "discriminator": [
        135,
        158,
        244,
        117,
        72,
        203,
        24,
        194
      ]
    }
  ],
  "types": [
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
