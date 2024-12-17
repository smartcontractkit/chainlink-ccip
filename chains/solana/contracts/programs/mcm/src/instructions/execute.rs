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
    multisig_name: [u8; MULTISIG_NAME_PADDED],
    chain_id: u64,
    nonce: u64,
    data: Vec<u8>,
    proof: Vec<[u8; 32]>,
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
    // while Solana's PoH provides cryptographic proofs of time passage,
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
        multisig_name.as_ref(),
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
#[instruction(multisig_name: [u8; MULTISIG_NAME_PADDED])]
pub struct Execute<'info> {
    #[account(mut, seeds = [CONFIG_SEED, multisig_name.as_ref()], bump)]
    pub multisig_config: Account<'info, MultisigConfig>,

    #[account(seeds = [ROOT_METADATA_SEED, multisig_name.as_ref()], bump)]
    pub root_metadata: Account<'info, RootMetadata>,

    #[account(mut, seeds = [EXPIRING_ROOT_AND_OP_COUNT_SEED, multisig_name.as_ref()], bump)]
    pub expiring_root_and_op_count: Account<'info, ExpiringRootAndOpCount>,

    /// CHECK: This is just used to invoke it through the CPI, it's value is not accessed directly
    pub to: UncheckedAccount<'info>,

    /// CHECK: program signer PDA that can hold balance
    #[account(
        seeds = [SIGNER_SEED, multisig_name.as_ref()],
        bump
    )]
    pub multisig_signer: UncheckedAccount<'info>,

    #[account(mut)]
    pub authority: Signer<'info>,
}

const MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP: &[u8] = &[
    0x08, 0xd2, 0x75, 0x62, 0x20, 0x06, 0xc4, 0xca, 0x82, 0xd0, 0x3f, 0x49, 0x8e, 0x90, 0x16, 0x3c,
    0xaf, 0xd5, 0x3c, 0x66, 0x3a, 0x48, 0x47, 0x0c, 0x3b, 0x52, 0xac, 0x8b, 0xfb, 0xd9, 0xf5, 0x2c,
];

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
    let byte = match (b1, b2) {
        (false, false) => 0x00,
        (false, true) => 0x01,
        (true, false) => 0x02,
        (true, true) => 0x03,
    };
    vec![byte]
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
            // 08d275622006c4ca82d03f498e90163cafd53c663a48470c3b52ac8bfbd9f52c  separator
            // 0000000000000000000000000000000000000000000000001266a21317e30848  chain_id
            // eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13  multisig
            // 0000000000000000000000000000000000000000000000000000000000000000  nonce
            // 30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb  to
            // 0000000000000000000000000000000000000000000000000800000000000000  data_len
            // d62c04f70c29d96e                                                  data
            // 0000000000000000000000000000000000000000000000000000000000000000  remaining_accounts_len

            assert_eq!(
                op.hash_leaf(),
                decode32("98ab3312f47d75a0173f23f50ba6042748e07bea064ae297982efebae3ecec9b")
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
            // Buffer[0]: 08d275622006c4ca82d03f498e90163cafd53c663a48470c3b52ac8bfbd9f52c
            // Buffer[1]: 0000000000000000000000000000000000000000000000001266a21317e30848
            // Buffer[2]: eceeab9f961bbf0050babcbf22663ee37905749830f41cd5096fbc9b158f8b13
            // Buffer[3]: 0000000000000000000000000000000000000000000000000100000000000000
            // Buffer[4]: 30d721519466e7a7c60691f75f49734954bee02fdb6700dd7f02a19de972ccdb
            // Buffer[5]: 0000000000000000000000000000000000000000000000000900000000000000
            // Buffer[6]: 11af9cfd5bad1ae47b
            // Buffer[7]: 0000000000000000000000000000000000000000000000000000000000000000

            assert_eq!(
                op.hash_leaf(),
                decode32("a23f9edc9ab94e247f6273d36633dced27525727b6a50aee95f14fd92aadb6e0")
            );
        }
    }
}
