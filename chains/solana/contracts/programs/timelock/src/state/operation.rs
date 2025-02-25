//! # Timelock Operation Management
//!
//! This module implements the core data structures and functions for managing timelock operations.
//! Operations represent batches of instructions that can be scheduled for delayed execution,
//! providing governance safeguards for critical administrative actions.
//!
//! The timelock system uses a state machine approach where operations progress through
//! several defined states, with cryptographic verification to ensure integrity.

use anchor_lang::prelude::*;
use anchor_lang::solana_program::instruction::Instruction;
use anchor_lang::solana_program::keccak::{hashv, HASH_BYTES};

use crate::constants::ANCHOR_DISCRIMINATOR;

/// Represents the current state of a timelock operation in its lifecycle.
///
/// Operations move through these states sequentially as they are prepared,
/// scheduled, and eventually executed.
#[derive(AnchorSerialize, AnchorDeserialize, Clone, Debug, PartialEq, PartialOrd)]
pub enum OperationState {
    /// Operation is created but not yet finalized.
    Initialized,
    /// Operation is finalized (instructions are locked in) but not yet scheduled.
    Finalized,
    /// Operation is scheduled (ready to be executed when its time comes).
    Scheduled,
    /// Operation has been executed and is complete.
    Done,
}

/// Represents a batch of instructions that can be scheduled for delayed execution.
///
/// Operations are the core data structure of the timelock system. Each operation contains
/// a set of instructions that will be executed atomically once the operation is ready.
/// Operations include cryptographic safeguards (ID verification) and dependency tracking
/// to ensure proper sequencing.
#[account]
pub struct Operation {
    pub state: OperationState,
    pub timestamp: u64,                     // scheduled timestamp in unix time
    pub id: [u8; 32],                       // hashed operation id
    pub predecessor: [u8; 32],              // hash of the previous operation
    pub salt: [u8; 32],                     // random salt for the operation
    pub total_instructions: u32,            // total number of instructions in the operation
    pub instructions: Vec<InstructionData>, // list of instructions
}

impl Space for Operation {
    // state(u8) + timestamp + id + predecessor + salt + total_ixs + vec prefix for instructions
    const INIT_SPACE: usize = 1 + 8 + 32 + 32 + 32 + 4 + 4;
}

impl Operation {
    pub fn space_after_init_instruction(&self, new_accounts: &[InstructionAccount]) -> usize {
        ANCHOR_DISCRIMINATOR + Self::INIT_SPACE
            // space for existing instructions
            + self.instructions.iter().map(|ix| ix.space()).sum::<usize>()
            // space for new instruction(program_id + data vec prefix + accounts vec prefix + account metas(pubkey + is_signer + is_writable))
            + (32 + 4 + 4 + new_accounts.len() * (32 + 1 + 1))
    }

    pub fn space_after_append_instruction(&self, new_data: &[u8]) -> usize {
        ANCHOR_DISCRIMINATOR + Self::INIT_SPACE
            // space for existing instructions
            + self.instructions.iter().map(|ix| ix.space()).sum::<usize>()
            // space for new instruction data
            + new_data.len()
    }

    // before scheduling, timestamp should be 0
    pub fn is_scheduled(&self) -> bool {
        self.state >= OperationState::Scheduled
    }

    pub fn is_finalized(&self) -> bool {
        self.state == OperationState::Finalized
    }

    pub fn is_pending(&self) -> bool {
        // scheduled but not executed
        self.state == OperationState::Scheduled
    }

    pub fn is_ready(&self, current_timestamp: u64) -> bool {
        // scheduled and timestamp is in the past
        self.state == OperationState::Scheduled && self.timestamp <= current_timestamp
    }

    pub fn is_done(&self) -> bool {
        self.state == OperationState::Done
    }

    pub fn finalize(&mut self) {
        self.state = OperationState::Finalized;
    }

    pub fn schedule(&mut self, scheduled_time: u64) {
        self.timestamp = scheduled_time;
        self.state = OperationState::Scheduled;
    }

    pub fn mark_done(&mut self) {
        self.state = OperationState::Done;
    }

    pub fn hash_instructions(&self, salt: [u8; HASH_BYTES]) -> [u8; HASH_BYTES] {
        let total_size = self.instructions.iter().map(|ix| ix.space()).sum::<usize>()
            + HASH_BYTES * 2 // predecessor and salt
            + 4; // instruction array vector prefix

        let mut encoded_data = Vec::with_capacity(total_size);

        // add length prefix for instruction array
        encoded_data.extend_from_slice(&(self.instructions.len() as u32).to_le_bytes());

        // encode each instruction
        for ix in &self.instructions {
            encoded_data.extend_from_slice(&ix.program_id.to_bytes());

            // add length prefix for accounts array
            encoded_data.extend_from_slice(&(ix.accounts.len() as u32).to_le_bytes());

            for acc in &ix.accounts {
                encoded_data.extend_from_slice(&acc.pubkey.to_bytes());
                encoded_data.push(acc.is_signer as u8);
                encoded_data.push(acc.is_writable as u8);
            }

            // add length prefix for instruction data
            encoded_data.extend_from_slice(&(ix.data.len() as u32).to_le_bytes());
            encoded_data.extend_from_slice(&ix.data);
        }

        encoded_data.extend_from_slice(&self.predecessor);
        encoded_data.extend_from_slice(&salt);

        hashv(&[&encoded_data]).to_bytes()
    }

    // Validate instruction data integrity by computing a salted hash of the instruction data
    // and comparing it against the stored operation ID. This ensures the uploaded
    // instructions remain unaltered between stored account and execution
    pub fn verify_id(&self) -> bool {
        self.hash_instructions(self.salt) == self.id
    }
}

/// A serializable representation of a Solana instruction.
///
/// The native SVM's Instruction type from solana_program doesn't implement the AnchorSerialize trait.
/// This wrapper provides serialization capabilities while maintaining the same functionality.
#[derive(AnchorSerialize, AnchorDeserialize, Clone, Default, Debug)]
pub struct InstructionData {
    pub program_id: Pubkey, // Target program that will process this instruction
    pub data: Vec<u8>,      // Instruction data passed to the program
    pub accounts: Vec<InstructionAccount>, // Accounts required for this instruction
}

impl InstructionData {
    pub fn space(&self) -> usize {
        // program id + vector prefix(data) + data + vector prefix(accounts) + accounts
        32 + 4 + self.data.len() + 4 + self.accounts.len() * InstructionAccount::INIT_SPACE
    }
}

impl From<&InstructionData> for Instruction {
    fn from(tx: &InstructionData) -> Instruction {
        Instruction {
            program_id: tx.program_id,
            accounts: tx.accounts.iter().map(Into::into).collect(),
            data: tx.data.clone(),
        }
    }
}

/// Represents an account used in an instruction, including its metadata.
///
/// This structure mirrors the AccountMeta used in Solana instructions
/// but implements Anchor's serialization traits.
// NOTE: space for InstructionAccount is calculated with InitSpace trait since it's static
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize, Clone, Default, Debug)]
pub struct InstructionAccount {
    pub pubkey: Pubkey,    // The account's public key
    pub is_signer: bool,   // Whether the account is a signer for the instruction
    pub is_writable: bool, // Whether the account is writable in the instruction
}

impl From<&InstructionAccount> for AccountMeta {
    fn from(account: &InstructionAccount) -> AccountMeta {
        match account.is_writable {
            false => AccountMeta::new_readonly(account.pubkey, account.is_signer),
            true => AccountMeta::new(account.pubkey, account.is_signer),
        }
    }
}

impl From<&AccountMeta> for InstructionAccount {
    fn from(account_meta: &AccountMeta) -> InstructionAccount {
        InstructionAccount {
            pubkey: account_meta.pubkey,
            is_signer: account_meta.is_signer,
            is_writable: account_meta.is_writable,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use anchor_lang::solana_program::{keccak::HASH_BYTES, pubkey::Pubkey};

    fn create_test_operation(
        instructions: Vec<InstructionData>,
        predecessor: [u8; HASH_BYTES],
    ) -> Operation {
        Operation {
            state: OperationState::Initialized,
            timestamp: 0,
            id: [0u8; 32],
            predecessor,
            salt: [0u8; 32],
            total_instructions: instructions.len() as u32,
            instructions,
        }
    }

    #[test]
    fn test_hash_operation_batch() {
        let program_id = Pubkey::new_unique();
        let account1 = Pubkey::new_unique();
        let account2 = Pubkey::new_unique();

        let tx1 = InstructionData {
            program_id,
            accounts: vec![
                InstructionAccount {
                    pubkey: account1,
                    is_signer: true,
                    is_writable: true,
                },
                InstructionAccount {
                    pubkey: account2,
                    is_signer: false,
                    is_writable: true,
                },
            ],
            data: vec![1, 2, 3],
        };

        let tx2 = InstructionData {
            program_id,
            accounts: vec![InstructionAccount {
                pubkey: account2,
                is_signer: false,
                is_writable: false,
            }],
            data: vec![4, 5, 6],
        };

        let predecessor = [1u8; HASH_BYTES];
        let salt = [2u8; HASH_BYTES];

        // Test single instruction
        let single_op = create_test_operation(vec![tx1.clone()], predecessor);
        let result1 = single_op.hash_instructions(salt);
        assert_eq!(result1.len(), HASH_BYTES);

        // Test multiple instructions
        let multiple_op = create_test_operation(vec![tx1, tx2], predecessor);
        let result2 = multiple_op.hash_instructions(salt);
        assert_eq!(result2.len(), HASH_BYTES);

        // Results should be different
        assert_ne!(result1, result2);
    }

    #[test]
    fn test_empty_instruction_list() {
        let predecessor = [1u8; HASH_BYTES];
        let salt = [2u8; HASH_BYTES];

        let empty_op = create_test_operation(vec![], predecessor);
        let result = empty_op.hash_instructions(salt);
        assert_eq!(result.len(), HASH_BYTES);

        let different_salt = [3u8; HASH_BYTES];
        let result2 = empty_op.hash_instructions(different_salt);
        assert_ne!(result, result2);
    }

    #[test]
    fn test_different_predecessors() {
        let program_id = Pubkey::new_unique();
        let account = Pubkey::new_unique();

        let tx = InstructionData {
            program_id,
            accounts: vec![InstructionAccount {
                pubkey: account,
                is_signer: true,
                is_writable: true,
            }],
            data: vec![1, 2, 3],
        };

        let predecessor1 = [1u8; HASH_BYTES];
        let predecessor2 = [3u8; HASH_BYTES];
        let salt = [2u8; HASH_BYTES];

        let op1 = create_test_operation(vec![tx.clone()], predecessor1);
        let result1 = op1.hash_instructions(salt);

        let op2 = create_test_operation(vec![tx], predecessor2);
        let result2 = op2.hash_instructions(salt);

        assert_ne!(result1, result2);
    }

    #[test]
    fn test_deterministic_output() {
        let program_id = Pubkey::new_unique();
        let account = Pubkey::new_unique();

        let tx = InstructionData {
            program_id,
            accounts: vec![InstructionAccount {
                pubkey: account,
                is_signer: true,
                is_writable: true,
            }],
            data: vec![1, 2, 3],
        };

        let predecessor = [1u8; HASH_BYTES];
        let salt = [2u8; HASH_BYTES];

        let op1 = create_test_operation(vec![tx.clone()], predecessor);
        let result1 = op1.hash_instructions(salt);

        let op2 = create_test_operation(vec![tx], predecessor);
        let result2 = op2.hash_instructions(salt);

        assert_eq!(result1, result2);
    }

    #[test]
    fn test_different_instructions_different_hash() {
        let program_id = Pubkey::new_unique();
        let account1 = Pubkey::new_unique();
        let account2 = Pubkey::new_unique();
        let predecessor = [1u8; HASH_BYTES];
        let salt = [2u8; HASH_BYTES];

        let op1 = create_test_operation(
            vec![InstructionData {
                program_id,
                accounts: vec![InstructionAccount {
                    pubkey: account1,
                    is_signer: true,
                    is_writable: false,
                }],
                data: vec![1, 2],
            }],
            predecessor,
        );

        let op2 = create_test_operation(
            vec![InstructionData {
                program_id,
                accounts: vec![InstructionAccount {
                    pubkey: account2,
                    is_signer: true,
                    is_writable: false,
                }],
                data: vec![1, 2],
            }],
            predecessor,
        );

        let hash1 = op1.hash_instructions(salt);
        let hash2 = op2.hash_instructions(salt);

        // even though data is the same, account differs → hash must differ
        assert_ne!(hash1, hash2);
    }

    #[test]
    fn test_collision_prevention() {
        let program_id = Pubkey::default(); // use default to minimize random differences
        let salt = [0u8; HASH_BYTES];
        let predecessor = [0u8; HASH_BYTES];

        // [ [1,2], [3] ] vs [ [1], [2,3] ]
        let case1_ix1 = InstructionData {
            program_id,
            accounts: vec![],
            data: vec![1, 2],
        };
        let case1_ix2 = InstructionData {
            program_id,
            accounts: vec![],
            data: vec![3],
        };
        let op1 = create_test_operation(vec![case1_ix1, case1_ix2], predecessor);

        let case1_ix3 = InstructionData {
            program_id,
            accounts: vec![],
            data: vec![1],
        };
        let case1_ix4 = InstructionData {
            program_id,
            accounts: vec![],
            data: vec![2, 3],
        };
        let op2 = create_test_operation(vec![case1_ix3, case1_ix4], predecessor);

        let hash1 = op1.hash_instructions(salt);
        let hash2 = op2.hash_instructions(salt);

        assert_ne!(hash1, hash2);

        // single instruction with 2 accounts vs. 2 instructions with 1 account each
        let account = Pubkey::default();
        let case2_ix1 = InstructionData {
            program_id,
            accounts: vec![
                InstructionAccount {
                    pubkey: account,
                    is_signer: true,
                    is_writable: true,
                },
                InstructionAccount {
                    pubkey: account,
                    is_signer: false,
                    is_writable: false,
                },
            ],
            data: vec![],
        };
        let op3 = create_test_operation(vec![case2_ix1], predecessor);

        let ix2_a = InstructionData {
            program_id,
            accounts: vec![InstructionAccount {
                pubkey: account,
                is_signer: true,
                is_writable: true,
            }],
            data: vec![],
        };
        let ix2_b = InstructionData {
            program_id,
            accounts: vec![InstructionAccount {
                pubkey: account,
                is_signer: false,
                is_writable: false,
            }],
            data: vec![],
        };
        let op4 = create_test_operation(vec![ix2_a, ix2_b], predecessor);

        let hash3 = op3.hash_instructions(salt);
        let hash4 = op4.hash_instructions(salt);

        assert_ne!(hash3, hash4);
    }
}
