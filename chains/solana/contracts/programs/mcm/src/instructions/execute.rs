use anchor_lang::prelude::*;
use anchor_lang::solana_program::instruction::Instruction;
use anchor_lang::solana_program::keccak::{hashv, HASH_BYTES};
use anchor_lang::solana_program::program::invoke_signed;

use crate::constant::*;
use crate::error::*;
use crate::eth_utils::*;
use crate::event::*;
use crate::state::config::*;
use crate::state::root::*;

pub fn execute<'info>(
    ctx: Context<'_, '_, '_, 'info, Execute<'info>>,
    multisig_id: [u8; MULTISIG_ID_PADDED],
    chain_id: u64,
    nonce: u64,
    data: Vec<u8>,
    proof: Vec<[u8; HASH_BYTES]>,
) -> Result<()> {
    let Execute {
        root_metadata,
        expiring_root_and_op_count,
        multisig_config,
        multisig_signer,
        ..
    } = &ctx.accounts;

    require!(
        root_metadata.post_op_count > expiring_root_and_op_count.op_count,
        McmError::PostOpCountReached
    );

    require!(chain_id == multisig_config.chain_id, McmError::WrongChainId);

    // use unix_timestamp from Clock which provides seconds precision.
    // while SVM's PoH provides cryptographic proofs of time passage,
    // clock drift can still occur if slot times exceed the ideal 400ms
    // (e.g., network had 30min drift in June 2022 when slots took 730ms)
    // ref: https://docs.solana.com/architecture/synchronization
    let now_ts = Clock::get()?.unix_timestamp;
    require!(
        now_ts <= expiring_root_and_op_count.valid_until.into(),
        McmError::RootExpired
    );

    require!(
        nonce == expiring_root_and_op_count.op_count,
        McmError::WrongNonce
    );

    let acc_infos = ctx.remaining_accounts;
    let acc_metas: Vec<AccountMeta> = acc_infos
        .iter()
        .flat_map(|acc_info| acc_info.to_account_metas(None))
        .collect();

    let op = Op {
        chain_id: chain_id.to_owned(),
        multisig: multisig_config.key(),
        nonce: nonce.to_owned(),
        data: data.to_owned(),
        to: ctx.accounts.to.key(),

        // remaining account metas, as received. No change made here to signers,
        // so that the proofs can be checked prior to editing the signer.
        remaining_accounts: acc_metas,
    };

    {
        let calculated_root = calculate_merkle_root(proof, &op.hash_leaf());
        require!(
            calculated_root == ctx.accounts.expiring_root_and_op_count.root,
            McmError::ProofCannotBeVerified
        );
    };

    // All validations past, now actually execute the op and count it
    let instruction = Instruction {
        program_id: op.to,
        accounts: op.cpi_remaining_accounts(multisig_signer.key()), // account metas enforcing our PDA as signer
        data: op.data,
    };

    let seeds = &[
        SIGNER_SEED,
        multisig_id.as_ref(),
        &[ctx.bumps.multisig_signer],
    ];
    let signer = &[&seeds[..]];

    ctx.accounts.expiring_root_and_op_count.op_count += 1;

    invoke_signed(&instruction, acc_infos, signer)?;

    emit!(OpExecuted {
        nonce,
        to: instruction.program_id,
        data: instruction.data,
    });

    Ok(())
}

#[derive(Accounts)]
#[instruction(multisig_id: [u8; MULTISIG_ID_PADDED])]
pub struct Execute<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_id.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(seeds = [ROOT_METADATA_SEED, multisig_id.as_ref()], bump)]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(mut, seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_id.as_ref()], bump)]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,

    /// CHECK: This is just used to invoke it through the CPI, it's value is not accessed directly
    pub to: UncheckedAccount<'info>,

    /// CHECK: program signer PDA that can hold balance
    #[account(
        seeds = [SIGNER_SEED, multisig_id.as_ref()],
        bump
    )]
    pub multisig_signer: UncheckedAccount<'info>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

#[derive(Debug)]
pub struct Op {
    pub chain_id: u64,    // network identification
    pub multisig: Pubkey, // multisig instance's PDA(config)
    pub nonce: u64,
    pub data: Vec<u8>,
    pub to: Pubkey,
    pub remaining_accounts: Vec<AccountMeta>,
}

impl Op {
    fn hash_leaf(&self) -> [u8; HASH_BYTES] {
        let chain_id = left_pad_vec(&self.chain_id.to_le_bytes());
        let nonce = left_pad_vec(&self.nonce.to_le_bytes());
        let data_len = left_pad_vec(&self.data.len().to_le_bytes());
        let remaining_accounts_len = left_pad_vec(&self.remaining_accounts.len().to_le_bytes());

        hashv(&[
            MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP,
            chain_id.as_slice(),
            &self.multisig.to_bytes(),
            nonce.as_slice(),
            &self.to.to_bytes(),
            data_len.as_slice(),
            self.data.as_slice(),
            remaining_accounts_len.as_slice(),
            self.serialized_remaining_accounts().as_slice(),
        ])
        .to_bytes()
    }

    fn serialized_remaining_accounts(&self) -> Vec<u8> {
        self.remaining_accounts
            .iter()
            .flat_map(|meta| {
                let mut bytes_vec = meta.pubkey.to_bytes().to_vec();
                bytes_vec.append(&mut bools_to_byte(meta.is_signer, meta.is_writable));
                bytes_vec
            })
            .collect()
    }

    /// Prepares account metas for CPI execution with the multisig signer.
    ///
    /// This critical security function:
    /// 1. Creates a modified copy of the original account metas
    /// 2. Sets `is_signer = true` only for the designated multisig signer PDA
    /// 3. Ensures other accounts cannot act as signers
    ///
    /// This selective signer assignment guarantees that only operations
    /// validated through the MCM's Merkle verification process can
    /// exercise signing authority, protecting privileged operations from
    /// unauthorized execution.
    ///
    /// # Parameters
    ///
    /// - `multisig_signer`: The public key of the multisig signer PDA
    ///
    /// # Returns
    ///
    /// A modified vector of AccountMeta with the signer flag properly set
    fn cpi_remaining_accounts(&self, multisig_signer: Pubkey) -> Vec<AccountMeta> {
        self.remaining_accounts
            .iter()
            .map(|acc_meta| {
                let mut cpi_acc_meta = acc_meta.clone();
                cpi_acc_meta.is_signer = acc_meta.pubkey == multisig_signer;
                cpi_acc_meta
            })
            .collect()
    }
}

fn bools_to_byte(b1: bool, b2: bool) -> Vec<u8> {
    vec![(u8::from(b1) << 1) + u8::from(b2)]
}

#[cfg(test)]
mod tests {
    use super::*;

    // Last 8 bytes of keccak256("solana:localnet") as big-endian
    // This is 0x4808e31713a26612 --> in little-endian, it is "1266a21317e30848"
    const CHAIN_ID: u64 = 5190648258797659666;

    fn decode32(s: &str) -> [u8; 32] {
        hex::decode(s).unwrap().to_owned().try_into().unwrap()
    }

    mod test_op_hash_leaf {
        use super::*;

        #[test]
        fn empty_op() {
            let op = Op {
                chain_id: CHAIN_ID,
                multisig: Pubkey::from(decode32(
                    "eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13",
                )),
                nonce: 0,
                data: hex::decode("d62c04f70c29d96e").unwrap(),
                to: Pubkey::from(decode32(
                    "30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb",
                )),
                remaining_accounts: vec![],
            };

            // op in hex:
            // fb98816ff3c5138a68abfd40b8d8fbc22972fea1dd89757331327e6e0a9440b7  separator
            // 0000000000000000000000000000000000000000000000001266a21317e30848  chain_id
            // eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13  multisig
            // 0000000000000000000000000000000000000000000000000000000000000000  nonce
            // 30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb  to
            // 0000000000000000000000000000000000000000000000000800000000000000  data_len
            // d62c04f70c29d96e                                                  data
            // 0000000000000000000000000000000000000000000000000000000000000000  remaining_accounts_len
            assert_eq!(
                op.hash_leaf(),
                decode32("0145195c134f5cec64fba146648763c2a7ac9bdb2c2efbc31a72bf9fb5f4246c")
            );
        }

        #[test]
        fn u8_instruction_data() {
            let op = Op {
                chain_id: CHAIN_ID,
                multisig: Pubkey::from(decode32(
                    "eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13",
                )),
                nonce: 1,
                data: hex::decode("11af9cfd5bad1ae47b").unwrap(),
                to: Pubkey::from(decode32(
                    "30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb",
                )),
                remaining_accounts: vec![],
            };

            // Raw Buffers
            // Buffer[0]: fb98816ff3c5138a68abfd40b8d8fbc22972fea1dd89757331327e6e0a9440b7
            // Buffer[1]: 0000000000000000000000000000000000000000000000001266a21317e30848
            // Buffer[2]: eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13
            // Buffer[3]: 0000000000000000000000000000000000000000000000000100000000000000
            // Buffer[4]: 30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb
            // Buffer[5]: 0000000000000000000000000000000000000000000000000900000000000000
            // Buffer[6]: 11af9cfd5bad1ae47b
            // Buffer[7]: 0000000000000000000000000000000000000000000000000000000000000000

            assert_eq!(
                op.hash_leaf(),
                decode32("d0dade6731524fbf07bf67fb9d250bc10b6e22adcd5e04ad9d6f515c75ac4951")
            );
        }
    }
}
