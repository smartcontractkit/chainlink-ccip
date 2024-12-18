use anchor_lang::prelude::*;
use anchor_lang::solana_program::instruction::Instruction;
use anchor_lang::solana_program::keccak::{hashv, HASH_BYTES};

use crate::constants::DONE_TIMESTAMP;

#[account]
pub struct Operation {
    pub timestamp: u64,
    pub id: [u8; 32],
    pub predecessor: [u8; 32],
    pub salt: [u8; 32],

    pub authority: Pubkey,
    pub is_finalized: bool,
    pub total_instructions: u32,
    pub instructions: Vec<InstructionData>,
}

impl Operation {
    pub fn is_scheduled(&self) -> bool {
        self.timestamp > 0
    }

    pub fn is_pending(&self) -> bool {
        self.timestamp > DONE_TIMESTAMP
    }

    pub fn is_ready(&self, current_timestamp: u64) -> bool {
        self.timestamp > DONE_TIMESTAMP && self.timestamp <= current_timestamp
    }

    pub fn is_done(&self) -> bool {
        self.timestamp == DONE_TIMESTAMP
    }

    pub fn mark_done(&mut self) {
        self.timestamp = DONE_TIMESTAMP;
    }

    pub fn hash_instructions(&self, salt: [u8; HASH_BYTES]) -> [u8; HASH_BYTES] {
        let total_size = self
            .instructions
            .iter()
            .map(|ix_data: &InstructionData| ix_data.space())
            .sum::<usize>()
            + HASH_BYTES * 2; // add predecessor and salt

        let mut encoded_data = Vec::with_capacity(total_size);

        // encode each instruction
        for ix in &self.instructions {
            encoded_data.extend_from_slice(&ix.program_id.to_bytes());

            for acc in &ix.accounts {
                encoded_data.extend_from_slice(&acc.pubkey.to_bytes());
                encoded_data.push(acc.is_signer as u8);
                encoded_data.push(acc.is_writable as u8);
            }

            encoded_data.extend_from_slice(&ix.data);
        }

        encoded_data.extend_from_slice(&self.predecessor);
        encoded_data.extend_from_slice(&salt);

        // hash everything with keccak256
        hashv(&[&encoded_data]).to_bytes()
    }

    pub fn verify_id(&self) -> bool {
        self.hash_instructions(self.salt) == self.id
    }
}

impl Space for Operation {
    // timestamp + id + predecessor + salt + total_ixs + is_finalized + authority + vec prefix for instructions
    const INIT_SPACE: usize = 8 + 32 + 32 + 32 + 4 + 1 + 32 + 4;
}

// The native Solana's Instruction type from solana_program doesn't implement the AnchorSerialize trait.
// This is a wrapper that provides serialization capabilities while maintaining the same functionality
#[derive(AnchorSerialize, AnchorDeserialize, Clone, Default, Debug)]
pub struct InstructionData {
    pub program_id: Pubkey,
    pub data: Vec<u8>,
    pub accounts: Vec<InstructionAccount>,
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

// NOTE: space for InstructionAccount is calculated with InitSpace trait since it's static
#[derive(InitSpace, AnchorSerialize, AnchorDeserialize, Clone, Default, Debug)]
pub struct InstructionAccount {
    pub pubkey: Pubkey,
    pub is_signer: bool,
    pub is_writable: bool,
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
            timestamp: 0,
            id: [0u8; 32],
            predecessor,
            salt: [0u8; 32],
            authority: Pubkey::new_unique(),
            is_finalized: false,
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
}
