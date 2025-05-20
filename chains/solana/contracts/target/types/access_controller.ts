/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/access_controller.json`.
 */
export type AccessController = {
  "address": "6KsN58MTnRQ8FfPaXHiFPPFGDRioikj9CdPvPxZJdCjb",
  "metadata": {
    "name": "accessController",
    "version": "1.0.1",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
    {
      "name": "acceptOwnership",
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
          "name": "state",
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
      "name": "addAccess",
      "discriminator": [
        151,
        189,
        105,
        24,
        113,
        60,
        99,
        138
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "owner",
          "signer": true
        },
        {
          "name": "address"
        }
      ],
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
          "name": "state",
          "writable": true
        },
        {
          "name": "owner",
          "signer": true
        }
      ],
      "args": []
    },
    {
      "name": "removeAccess",
      "discriminator": [
        92,
        172,
        70,
        124,
        83,
        45,
        88,
        22
      ],
      "accounts": [
        {
          "name": "state",
          "writable": true
        },
        {
          "name": "owner",
          "signer": true
        },
        {
          "name": "address"
        }
      ],
      "args": []
    },
    {
      "name": "transferOwnership",
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
          "name": "state",
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
    }
  ],
  "accounts": [
    {
      "name": "accessController",
      "discriminator": [
        143,
        45,
        12,
        204,
        220,
        20,
        114,
        87
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "unauthorized",
      "msg": "unauthorized"
    },
    {
      "code": 6001,
      "name": "invalidInput",
      "msg": "Invalid input"
    },
    {
      "code": 6002,
      "name": "full",
      "msg": "Access list is full"
    }
  ],
  "types": [
    {
      "name": "accessController",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "pubkey"
          },
          {
            "name": "proposedOwner",
            "type": "pubkey"
          },
          {
            "name": "accessList",
            "type": {
              "defined": {
                "name": "accessList"
              }
            }
          }
        ]
      }
    },
    {
      "name": "accessList",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "xs",
            "type": {
              "array": [
                "pubkey",
                64
              ]
            }
          },
          {
            "name": "len",
            "type": "u64"
          }
        ]
      }
    }
  ]
};
