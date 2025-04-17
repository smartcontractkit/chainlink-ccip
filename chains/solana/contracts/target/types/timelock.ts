export type Timelock = {
  "version": "0.0.1-dev",
  "name": "timelock",
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
        },
        {
          "name": "accessControllerProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "proposerRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "executorRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cancellerRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "bypasserRoleAccessController",
          "isMut": false,
          "isSigner": false
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "accessControllerProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
            "defined": "Role"
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": "InstructionAccount"
            }
          }
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "predecessorOperation",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "timelockSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": "InstructionAccount"
            }
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "timelockSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        }
      ]
    },
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
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
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "proposerRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "executorRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "cancellerRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "bypasserRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "minDelay",
            "type": "u64"
          },
          {
            "name": "blockedSelectors",
            "type": {
              "defined": "BlockedSelectors"
            }
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
              "defined": "OperationState"
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
                "defined": "InstructionData"
              }
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "BlockedSelectors",
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
      "name": "InstructionData",
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
            "type": "publicKey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "accounts",
            "type": {
              "vec": {
                "defined": "InstructionAccount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "InstructionAccount",
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
            "type": "publicKey"
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
      "name": "TimelockError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "InvalidInput"
          },
          {
            "name": "Overflow"
          },
          {
            "name": "InvalidId"
          },
          {
            "name": "OperationNotFinalized"
          },
          {
            "name": "OperationAlreadyFinalized"
          },
          {
            "name": "TooManyInstructions"
          },
          {
            "name": "OperationAlreadyScheduled"
          },
          {
            "name": "DelayInsufficient"
          },
          {
            "name": "OperationNotCancellable"
          },
          {
            "name": "OperationNotReady"
          },
          {
            "name": "OperationAlreadyExecuted"
          },
          {
            "name": "MissingDependency"
          },
          {
            "name": "InvalidAccessController"
          },
          {
            "name": "BlockedSelector"
          },
          {
            "name": "AlreadyBlocked"
          },
          {
            "name": "SelectorNotFound"
          },
          {
            "name": "MaxCapacityReached"
          }
        ]
      }
    },
    {
      "name": "Role",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Admin"
          },
          {
            "name": "Proposer"
          },
          {
            "name": "Executor"
          },
          {
            "name": "Canceller"
          },
          {
            "name": "Bypasser"
          }
        ]
      }
    },
    {
      "name": "OperationState",
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
            "name": "Initialized"
          },
          {
            "name": "Finalized"
          },
          {
            "name": "Scheduled"
          },
          {
            "name": "Done"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CallScheduled",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "predecessor",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "salt",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "delay",
          "type": "u64",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "CallExecuted",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "BypasserCallExecuted",
      "fields": [
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "Cancelled",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "MinDelayChange",
      "fields": [
        {
          "name": "oldDuration",
          "type": "u64",
          "index": false
        },
        {
          "name": "newDuration",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "FunctionSelectorBlocked",
      "fields": [
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "FunctionSelectorUnblocked",
      "fields": [
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          },
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

export const IDL: Timelock = {
  "version": "0.0.1-dev",
  "name": "timelock",
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
        },
        {
          "name": "accessControllerProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "proposerRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "executorRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "cancellerRoleAccessController",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "bypasserRoleAccessController",
          "isMut": false,
          "isSigner": false
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
      "accounts": [
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "accessControllerProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
            "defined": "Role"
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": "InstructionAccount"
            }
          }
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "predecessorOperation",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "timelockSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        },
        {
          "name": "accounts",
          "type": {
            "vec": {
              "defined": "InstructionAccount"
            }
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
      "accounts": [
        {
          "name": "operation",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "config",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "timelockSigner",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "roleAccessController",
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
          "type": "publicKey"
        }
      ]
    },
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
          "name": "timelockId",
          "type": {
            "array": [
              "u8",
              32
            ]
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
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "proposerRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "executorRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "cancellerRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "bypasserRoleAccessController",
            "type": "publicKey"
          },
          {
            "name": "minDelay",
            "type": "u64"
          },
          {
            "name": "blockedSelectors",
            "type": {
              "defined": "BlockedSelectors"
            }
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
              "defined": "OperationState"
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
                "defined": "InstructionData"
              }
            }
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "BlockedSelectors",
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
      "name": "InstructionData",
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
            "type": "publicKey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "accounts",
            "type": {
              "vec": {
                "defined": "InstructionAccount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "InstructionAccount",
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
            "type": "publicKey"
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
      "name": "TimelockError",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "InvalidInput"
          },
          {
            "name": "Overflow"
          },
          {
            "name": "InvalidId"
          },
          {
            "name": "OperationNotFinalized"
          },
          {
            "name": "OperationAlreadyFinalized"
          },
          {
            "name": "TooManyInstructions"
          },
          {
            "name": "OperationAlreadyScheduled"
          },
          {
            "name": "DelayInsufficient"
          },
          {
            "name": "OperationNotCancellable"
          },
          {
            "name": "OperationNotReady"
          },
          {
            "name": "OperationAlreadyExecuted"
          },
          {
            "name": "MissingDependency"
          },
          {
            "name": "InvalidAccessController"
          },
          {
            "name": "BlockedSelector"
          },
          {
            "name": "AlreadyBlocked"
          },
          {
            "name": "SelectorNotFound"
          },
          {
            "name": "MaxCapacityReached"
          }
        ]
      }
    },
    {
      "name": "Role",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Admin"
          },
          {
            "name": "Proposer"
          },
          {
            "name": "Executor"
          },
          {
            "name": "Canceller"
          },
          {
            "name": "Bypasser"
          }
        ]
      }
    },
    {
      "name": "OperationState",
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
            "name": "Initialized"
          },
          {
            "name": "Finalized"
          },
          {
            "name": "Scheduled"
          },
          {
            "name": "Done"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CallScheduled",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "predecessor",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "salt",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "delay",
          "type": "u64",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "CallExecuted",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "BypasserCallExecuted",
      "fields": [
        {
          "name": "index",
          "type": "u64",
          "index": false
        },
        {
          "name": "target",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "data",
          "type": "bytes",
          "index": false
        }
      ]
    },
    {
      "name": "Cancelled",
      "fields": [
        {
          "name": "id",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "MinDelayChange",
      "fields": [
        {
          "name": "oldDuration",
          "type": "u64",
          "index": false
        },
        {
          "name": "newDuration",
          "type": "u64",
          "index": false
        }
      ]
    },
    {
      "name": "FunctionSelectorBlocked",
      "fields": [
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          },
          "index": false
        }
      ]
    },
    {
      "name": "FunctionSelectorUnblocked",
      "fields": [
        {
          "name": "selector",
          "type": {
            "array": [
              "u8",
              8
            ]
          },
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
