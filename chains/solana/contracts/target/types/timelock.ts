/**
 * Program IDL in camelCase format in order to be used in JS/TS.
 *
 * Note that this is only a type helper and is not the actual IDL. The original
 * IDL can be found at `target/idl/timelock.json`.
 */
export type Timelock = {
  "address": "DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA",
  "metadata": {
    "name": "timelock",
    "version": "0.0.1-dev",
    "spec": "0.1.0",
    "description": "SVM implementation of RBAC Timelock"
  },
  "docs": [
    "The `timelock` program provides a mechanism to schedule, execute, and (if needed) cancel",
    "operations in a delayed fashion. It supports both standard operations (which enforce delays and dependencies)",
    "and bypass operations (for emergency cases).",
    "",
    "Operation management for timelock system, handling both standard (timelock-enforced)",
    "and bypass (emergency) operations.",
    "",
    "Standard Operation Flow:",
    "- initialize -> append(init_ix, append_ix_data) -> finalize -> schedule -> execute_batch",
    "- Enforces timelock delays and predecessor dependencies",
    "",
    "Bypass Operation Flow:",
    "- initialize -> append(init_ix, append_ix_data) -> finalize -> bypass_execute_batch",
    "- No required delay or additional checks, closes operation account after execution",
    "",
    "Implementation uses separate code paths and PDAs for each operation type",
    "to maintain clear security boundaries and audit trails, despite similar logic.",
    "All operations enforce state transitions, size limits, and role-based access."
  ],
  "instructions": [
    {
      "name": "acceptOwnership",
      "docs": [
        "Accept ownership of the timelock config.",
        "",
        "The proposed new owner must call this function to assume ownership.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_timelock_id`: The timelock identifier."
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
          "name": "timelockId",
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
      "name": "appendBypasserInstructionData",
      "docs": [
        "Append additional data to an instruction of a bypasser operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the bypasser operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `_id`: The operation identifier.",
        "- `ix_index`: The index of the instruction.",
        "- `ix_data_chunk`: The data to append."
      ],
      "discriminator": [
        184,
        232,
        151,
        222,
        111,
        117,
        215,
        197
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "ixIndex",
          "type": "u32"
        },
        {
          "name": "ixDataChunk",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "appendInstructionData",
      "docs": [
        "Append additional instruction data to an instruction of an existing standard operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the operation account.",
        "- `_timelock_id`: The timelock identifier (for PDA derivation).",
        "- `_id`: The operation identifier.",
        "- `ix_index`: The index of the instruction to which the data will be appended.",
        "- `ix_data_chunk`: A chunk of data to be appended."
      ],
      "discriminator": [
        76,
        77,
        102,
        131,
        136,
        12,
        45,
        5
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "ixIndex",
          "type": "u32"
        },
        {
          "name": "ixDataChunk",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "batchAddAccess",
      "docs": [
        "Add a new access role in batch. Only the admin is allowed to perform this operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts required for batch adding access.",
        "- `timelock_id`: A unique, padded identifier for this timelock instance.",
        "- `role`: The role to be added."
      ],
      "discriminator": [
        73,
        141,
        223,
        79,
        66,
        154,
        226,
        67
      ],
      "accounts": [
        {
          "name": "config"
        },
        {
          "name": "accessControllerProgram"
        },
        {
          "name": "roleAccessController",
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "role",
          "type": {
            "defined": {
              "name": "role"
            }
          }
        }
      ]
    },
    {
      "name": "blockFunctionSelector",
      "docs": [
        "Block a function selector from being called.",
        "",
        "Only the admin can block function selectors.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `selector`: The 8-byte function selector(Anchor discriminator) to block."
      ],
      "discriminator": [
        119,
        89,
        101,
        41,
        72,
        143,
        218,
        185
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          }
        }
      ]
    },
    {
      "name": "bypasserExecuteBatch",
      "docs": [
        "Execute operations immediately using the bypasser flow, bypassing time delays",
        "and predecessor checks.",
        "",
        "This function provides an emergency execution mechanism that:",
        "1. Skips the timelock waiting period required for standard operations",
        "2. Does not enforce predecessor dependencies",
        "3. Closes the operation account after execution",
        "",
        "# Emergency Use Only",
        "",
        "The bypasser flow is intended strictly for emergency situations where",
        "waiting for the standard timelock delay would cause harm. Access to this",
        "function is tightly controlled through the Bypasser role.",
        "",
        "# Parameters",
        "",
        "- `ctx`: Context containing operation accounts and signer information",
        "- `timelock_id`: Identifier for the timelock instance",
        "- `_id`: Operation ID (used for PDA derivation)"
      ],
      "discriminator": [
        90,
        62,
        66,
        6,
        227,
        174,
        30,
        194
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "timelockSigner"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "cancel",
      "docs": [
        "Cancel a scheduled operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts for cancellation.",
        "- `timelock_id`: The timelock identifier.",
        "- `id`: The operation identifier (precalculated)."
      ],
      "discriminator": [
        232,
        219,
        223,
        41,
        219,
        236,
        220,
        190
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "clearBypasserOperation",
      "docs": [
        "Clear a finalized bypasser operation.",
        "",
        "Closes the bypasser operation account.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the bypasser operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `_id`: The operation identifier."
      ],
      "discriminator": [
        200,
        21,
        249,
        130,
        56,
        13,
        128,
        32
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "clearOperation",
      "docs": [
        "Clear an operation that has been finalized.",
        "",
        "This effectively closes the operation account so that it can be reinitialized later.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `_id`: The operation identifier."
      ],
      "discriminator": [
        111,
        217,
        62,
        240,
        224,
        75,
        60,
        58
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "executeBatch",
      "docs": [
        "Executes a scheduled batch of operations after validating readiness and predecessor dependencies.",
        "",
        "This function:",
        "1. Verifies the operation is ready for execution (delay period has passed)",
        "2. Validates that any predecessor operation has been completed",
        "3. Executes each instruction in the operation using the timelock signer PDA",
        "4. Emits events for each executed instruction",
        "",
        "# Parameters",
        "",
        "- `ctx`: Context containing operation accounts and signer information",
        "- `timelock_id`: Identifier for the timelock instance",
        "- `_id`: Operation ID (used for PDA derivation)",
        "",
        "# Security Considerations",
        "",
        "This instruction uses PDA signing to create a trusted execution environment.",
        "The timelock's signer PDA will replace any account marked as a signer in the",
        "original instructions, providing the necessary privileges while maintaining",
        "security through program derivation."
      ],
      "discriminator": [
        112,
        159,
        211,
        51,
        238,
        70,
        212,
        60
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "predecessorOperation"
        },
        {
          "name": "config"
        },
        {
          "name": "timelockSigner"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "finalizeBypasserOperation",
      "docs": [
        "Finalize a bypasser operation.",
        "",
        "Marks the bypasser operation as finalized, ready for execution.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the bypasser operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `_id`: The operation identifier."
      ],
      "discriminator": [
        45,
        55,
        198,
        51,
        124,
        24,
        169,
        250
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "finalizeOperation",
      "docs": [
        "Finalize a standard operation.",
        "",
        "Finalizing an operation marks it as ready for scheduling.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the operation account.",
        "- `_timelock_id`: The timelock identifier (for PDA derivation).",
        "- `_id`: The operation identifier."
      ],
      "discriminator": [
        63,
        208,
        32,
        98,
        85,
        182,
        236,
        140
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
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
      "name": "initialize",
      "docs": [
        "Initialize the timelock configuration.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts required for initialization.",
        "- `timelock_id`: A unique, padded identifier for this timelock instance.",
        "- `min_delay`: The minimum delay (in seconds) required for scheduled operations."
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
        },
        {
          "name": "accessControllerProgram"
        },
        {
          "name": "proposerRoleAccessController"
        },
        {
          "name": "executorRoleAccessController"
        },
        {
          "name": "cancellerRoleAccessController"
        },
        {
          "name": "bypasserRoleAccessController"
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "minDelay",
          "type": "u64"
        }
      ]
    },
    {
      "name": "initializeBypasserInstruction",
      "docs": [
        "Initialize an instruction for a bypasser operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the bypasser operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `_id`: The operation identifier.",
        "- `program_id`: The target program for the instruction.",
        "- `accounts`: The list of accounts required for the instruction."
      ],
      "discriminator": [
        50,
        17,
        205,
        172,
        175,
        140,
        195,
        39
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "programId",
          "type": "pubkey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": {
                "name": "instructionAccount"
              }
            }
          }
        }
      ]
    },
    {
      "name": "initializeBypasserOperation",
      "docs": [
        "Initialize a bypasser operation.",
        "",
        "Bypasser operations have no predecessor and can be executed without delay.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the bypasser operation account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `id`: The operation identifier.",
        "- `salt`: A salt value for PDA derivation.",
        "- `instruction_count`: The number of instructions to be added."
      ],
      "discriminator": [
        58,
        27,
        48,
        204,
        19,
        197,
        63,
        26
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "salt",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "instructionCount",
          "type": "u32"
        }
      ]
    },
    {
      "name": "initializeInstruction",
      "docs": [
        "Append a new instruction to an existing standard operation.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the operation account.",
        "- `_timelock_id`: The timelock identifier (for PDA derivation).",
        "- `_id`: The operation identifier.",
        "- `program_id`: The target program for the instruction.",
        "- `accounts`: The list of accounts required for the instruction."
      ],
      "discriminator": [
        195,
        230,
        213,
        135,
        144,
        148,
        142,
        85
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "programId",
          "type": "pubkey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": {
                "name": "instructionAccount"
              }
            }
          }
        }
      ]
    },
    {
      "name": "initializeOperation",
      "docs": [
        "Initialize a new standard timelock operation.",
        "",
        "This sets up a new operation with the given ID, predecessor, salt, and expected number of instructions.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the operation account.",
        "- `_timelock_id`: A padded identifier for the timelock (unused here but required for PDA derivation).",
        "- `id`: The unique identifier for the operation.",
        "- `predecessor`: The identifier of the predecessor operation.",
        "- `salt`: A salt value to help create unique PDAs.",
        "- `instruction_count`: The total number of instructions that will be added to this operation."
      ],
      "discriminator": [
        15,
        96,
        217,
        171,
        124,
        4,
        113,
        243
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "predecessor",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "salt",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "instructionCount",
          "type": "u32"
        }
      ]
    },
    {
      "name": "scheduleBatch",
      "docs": [
        "Schedule a finalized operation to be executed after a delay.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the accounts for scheduling.",
        "- `timelock_id`: The timelock identifier.",
        "- `id`: The operation identifier.",
        "- `delay`: The delay (in seconds) before the operation can be executed."
      ],
      "discriminator": [
        242,
        140,
        87,
        106,
        71,
        226,
        86,
        32
      ],
      "accounts": [
        {
          "name": "operation",
          "writable": true
        },
        {
          "name": "config"
        },
        {
          "name": "roleAccessController"
        },
        {
          "name": "authority",
          "writable": true,
          "signer": true
        }
      ],
      "args": [
        {
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "delay",
          "type": "u64"
        }
      ]
    },
    {
      "name": "transferOwnership",
      "docs": [
        "Propose a new owner for the timelock instance config.",
        "",
        "Only the current owner (admin) can propose a new owner.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_timelock_id`: The timelock identifier.",
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
          "name": "timelockId",
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
    },
    {
      "name": "unblockFunctionSelector",
      "docs": [
        "Unblock a previously blocked function selector.",
        "",
        "Only the admin can unblock function selectors.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `selector`: The function selector to unblock."
      ],
      "discriminator": [
        53,
        84,
        245,
        196,
        149,
        52,
        30,
        57
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          }
        }
      ]
    },
    {
      "name": "updateDelay",
      "docs": [
        "Update the minimum delay required for scheduled operations.",
        "",
        "Only the admin can update the delay.",
        "",
        "# Parameters",
        "",
        "- `ctx`: The context containing the configuration account.",
        "- `_timelock_id`: The timelock identifier.",
        "- `delay`: The new minimum delay value."
      ],
      "discriminator": [
        164,
        186,
        80,
        62,
        85,
        88,
        182,
        147
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
          }
        },
        {
          "name": "delay",
          "type": "u64"
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
      "name": "operation",
      "discriminator": [
        171,
        150,
        196,
        17,
        229,
        166,
        58,
        44
      ]
    }
  ],
  "events": [
    {
      "name": "bypasserCallExecuted",
      "discriminator": [
        61,
        41,
        96,
        207,
        16,
        173,
        99,
        75
      ]
    },
    {
      "name": "callExecuted",
      "discriminator": [
        237,
        120,
        238,
        142,
        189,
        37,
        65,
        128
      ]
    },
    {
      "name": "callScheduled",
      "discriminator": [
        191,
        85,
        90,
        167,
        132,
        223,
        184,
        57
      ]
    },
    {
      "name": "cancelled",
      "discriminator": [
        136,
        23,
        42,
        65,
        143,
        233,
        234,
        46
      ]
    },
    {
      "name": "functionSelectorBlocked",
      "discriminator": [
        67,
        101,
        36,
        217,
        222,
        85,
        191,
        71
      ]
    },
    {
      "name": "functionSelectorUnblocked",
      "discriminator": [
        189,
        124,
        164,
        141,
        141,
        189,
        0,
        218
      ]
    },
    {
      "name": "minDelayChange",
      "discriminator": [
        186,
        71,
        244,
        116,
        244,
        76,
        230,
        254
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "invalidInput",
      "msg": "Invalid inputs"
    },
    {
      "code": 6001,
      "name": "overflow",
      "msg": "overflow"
    },
    {
      "code": 6002,
      "name": "invalidId",
      "msg": "Provided ID is invalid"
    },
    {
      "code": 6003,
      "name": "operationNotFinalized",
      "msg": "operation not finalized"
    },
    {
      "code": 6004,
      "name": "operationAlreadyFinalized",
      "msg": "operation is already finalized"
    },
    {
      "code": 6005,
      "name": "tooManyInstructions",
      "msg": "too many instructions in the operation"
    },
    {
      "code": 6006,
      "name": "operationAlreadyScheduled",
      "msg": "operation already scheduled"
    },
    {
      "code": 6007,
      "name": "delayInsufficient",
      "msg": "insufficient delay"
    },
    {
      "code": 6008,
      "name": "operationNotCancellable",
      "msg": "operation cannot be cancelled"
    },
    {
      "code": 6009,
      "name": "operationNotReady",
      "msg": "operation is not ready"
    },
    {
      "code": 6010,
      "name": "operationAlreadyExecuted",
      "msg": "operation is already executed"
    },
    {
      "code": 6011,
      "name": "missingDependency",
      "msg": "Predecessor operation is not found"
    },
    {
      "code": 6012,
      "name": "invalidAccessController",
      "msg": "Provided access controller is invalid"
    },
    {
      "code": 6013,
      "name": "blockedSelector",
      "msg": "selector is blocked"
    },
    {
      "code": 6014,
      "name": "alreadyBlocked",
      "msg": "selector is already blocked"
    },
    {
      "code": 6015,
      "name": "selectorNotFound",
      "msg": "selector not found"
    },
    {
      "code": 6016,
      "name": "maxCapacityReached",
      "msg": "maximum capacity reached for function blocker"
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
    },
    {
      "name": "blockedSelectors",
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
                {
                  "array": [
                    "u8",
                    8
                  ]
                },
                128
              ]
            }
          },
          {
            "name": "len",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "bypasserCallExecuted",
      "docs": [
        "@dev Emitted when a call is performed via bypasser."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "index",
            "type": "u64"
          },
          {
            "name": "target",
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
      "name": "callExecuted",
      "docs": [
        "@dev Emitted when a call is performed as part of operation `id`."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "id",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "index",
            "type": "u64"
          },
          {
            "name": "target",
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
      "name": "callScheduled",
      "docs": [
        "@dev Emitted when a call is scheduled as part of operation `id`."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "id",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "index",
            "type": "u64"
          },
          {
            "name": "target",
            "type": "pubkey"
          },
          {
            "name": "predecessor",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "salt",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "delay",
            "type": "u64"
          },
          {
            "name": "data",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "cancelled",
      "docs": [
        "@dev Emitted when operation `id` is cancelled."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "id",
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
      "name": "config",
      "serialization": "bytemuck",
      "repr": {
        "kind": "c"
      },
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "timelockId",
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
            "name": "proposerRoleAccessController",
            "type": "pubkey"
          },
          {
            "name": "executorRoleAccessController",
            "type": "pubkey"
          },
          {
            "name": "cancellerRoleAccessController",
            "type": "pubkey"
          },
          {
            "name": "bypasserRoleAccessController",
            "type": "pubkey"
          },
          {
            "name": "minDelay",
            "type": "u64"
          },
          {
            "name": "blockedSelectors",
            "type": {
              "defined": {
                "name": "blockedSelectors"
              }
            }
          }
        ]
      }
    },
    {
      "name": "functionSelectorBlocked",
      "docs": [
        "@dev Emitted when a function selector is blocked."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "selector",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          }
        ]
      }
    },
    {
      "name": "functionSelectorUnblocked",
      "docs": [
        "@dev Emitted when a function selector is unblocked."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "selector",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          }
        ]
      }
    },
    {
      "name": "instructionAccount",
      "docs": [
        "Represents an account used in an instruction, including its metadata.",
        "",
        "This structure mirrors the AccountMeta used in Solana instructions",
        "but implements Anchor's serialization traits."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pubkey",
            "type": "pubkey"
          },
          {
            "name": "isSigner",
            "type": "bool"
          },
          {
            "name": "isWritable",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "instructionData",
      "docs": [
        "A serializable representation of a Solana instruction.",
        "",
        "The native SVM's Instruction type from solana_program doesn't implement the AnchorSerialize trait.",
        "This wrapper provides serialization capabilities while maintaining the same functionality."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "programId",
            "type": "pubkey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "accounts",
            "type": {
              "vec": {
                "defined": {
                  "name": "instructionAccount"
                }
              }
            }
          }
        ]
      }
    },
    {
      "name": "minDelayChange",
      "docs": [
        "@dev Emitted when the minimum delay for future operations is modified."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "oldDuration",
            "type": "u64"
          },
          {
            "name": "newDuration",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "operation",
      "docs": [
        "Represents a batch of instructions that can be scheduled for delayed execution.",
        "",
        "Operations are the core data structure of the timelock system. Each operation contains",
        "a set of instructions that will be executed atomically once the operation is ready.",
        "Operations include cryptographic safeguards (ID verification) and dependency tracking",
        "to ensure proper sequencing."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "state",
            "type": {
              "defined": {
                "name": "operationState"
              }
            }
          },
          {
            "name": "timestamp",
            "type": "u64"
          },
          {
            "name": "id",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "predecessor",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "salt",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "totalInstructions",
            "type": "u32"
          },
          {
            "name": "instructions",
            "type": {
              "vec": {
                "defined": {
                  "name": "instructionData"
                }
              }
            }
          }
        ]
      }
    },
    {
      "name": "operationState",
      "docs": [
        "Represents the current state of a timelock operation in its lifecycle.",
        "",
        "Operations move through these states sequentially as they are prepared,",
        "scheduled, and eventually executed."
      ],
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "initialized"
          },
          {
            "name": "finalized"
          },
          {
            "name": "scheduled"
          },
          {
            "name": "done"
          }
        ]
      }
    },
    {
      "name": "role",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "admin"
          },
          {
            "name": "proposer"
          },
          {
            "name": "executor"
          },
          {
            "name": "canceller"
          },
          {
            "name": "bypasser"
          }
        ]
      }
    }
  ]
};
