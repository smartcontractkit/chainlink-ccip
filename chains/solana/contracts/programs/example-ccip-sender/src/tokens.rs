use anchor_lang::prelude::*;
use anchor_spl::{
    token_2022::spl_token_2022::{
        self,
        instruction::{approve_checked, transfer_checked},
    },
    token_interface,
};
use ccip_router::messages::SVMTokenAmount;
use solana_program::{address_lookup_table::state::AddressLookupTable, program::invoke_signed};

#[allow(clippy::too_many_arguments)]
pub fn transfer_to_self_and_approve<'a>(
    token_program: &AccountInfo<'a>,
    token_mint: &AccountInfo<'a>,
    from_ata: &AccountInfo<'a>,
    self_ata: &AccountInfo<'a>,
    self_signer: &AccountInfo<'a>,
    to_signer: &AccountInfo<'a>,
    seeds: &[&[u8]],
    amount: u64,
    decimals: u8,
) -> Result<()> {
    // transfer to this program
    {
        // build transfer from token_2022 SDK, but set expected token program ID
        let mut ix = transfer_checked(
            &spl_token_2022::ID,
            &from_ata.key(),
            &token_mint.key(),
            &self_ata.key(),
            &self_signer.key(),
            &[],
            amount,
            decimals,
        )?;
        ix.program_id = token_program.key(); // allow any custom token program
        invoke_signed(
            &ix,
            &[
                from_ata.clone(),
                token_mint.clone(),
                self_ata.clone(),
                self_signer.clone(),
            ],
            &[seeds],
        )?;
    }

    // approve router to withdraw from this program
    {
        let mut ix = approve_checked(
            &spl_token_2022::ID,
            &self_ata.key(),
            &token_mint.key(),
            &to_signer.key(),
            &self_signer.key(),
            &[],
            amount,
            decimals,
        )?;
        ix.program_id = token_program.key(); // allow any custom token program
        invoke_signed(
            &ix,
            &[
                self_ata.clone(),
                token_mint.clone(),
                to_signer.clone(),
                self_signer.clone(),
            ],
            &[seeds],
        )?;
    }
    Ok(())
}

pub struct TokenAccounts<'a> {
    pub program: &'a AccountInfo<'a>,
    pub mint: &'a AccountInfo<'a>,
    pub decimals: u8,

    // ATAs
    pub from_ata: &'a AccountInfo<'a>,
    pub self_ata: &'a AccountInfo<'a>,

    // signers
    pub pool_signer: &'a AccountInfo<'a>,
    pub ccip_router_pool_signer: &'a AccountInfo<'a>,

    // billing configs
    pub token_billing_config: &'a AccountInfo<'a>,
    pub fee_token_config: &'a AccountInfo<'a>,

    // pool interaction accounts
    pub pool_accounts: &'a [AccountInfo<'a>],
}

pub fn parse_and_validate_token_pool_accounts<'a>(
    token_amounts: &[SVMTokenAmount],
    token_indexes: &[u8],
    all_token_accounts: &'a [AccountInfo<'a>],
) -> Result<(Vec<TokenAccounts<'a>>, Vec<u8>)> {
    // expected order of accounts
    // [ATAs for user] - equal to number of tokens transferred
    // [accounts for pool 0]
    // [accounts for pool 1]
    // ...

    let mut parsed: Vec<TokenAccounts> = vec![];
    let mut ccip_token_indexes: Vec<u8> = vec![];
    if !token_amounts.is_empty() {
        assert_eq!(token_amounts.len(), token_indexes.len());
        assert_eq!(token_indexes[0] as usize, token_amounts.len()); // check for correct staring index to leave space for contract ATAs
        for (i, _) in token_amounts.iter().enumerate() {
            let end: usize = if i == token_indexes.len() - 1 {
                all_token_accounts.len()
            } else {
                token_indexes[i + 1] as usize
            };
            let accounts = &all_token_accounts[token_indexes[i] as usize..end];
            // accounts based on user or chain
            // 0 - self_token_account
            // 1 - token_billing_config
            // 2 - pool_chain_config

            // constant accounts for any pool interaction
            // 3 - lookup_table
            // 4 - token_admin_registry
            // 5 - pool_program
            // 6 - pool_config
            // 7 - pool_token_account
            // 8 - pool_signer
            // 9 - token_program
            // 10 - mint
            // 11 - fee_token_config
            // 12 - ccip_router_pool_signer

            let mint_data = &mut &accounts[10].data.borrow()[..];
            let mint_account = token_interface::Mint::try_deserialize(mint_data).unwrap();

            parsed.push(TokenAccounts {
                program: &accounts[9],
                mint: &accounts[10],
                decimals: mint_account.decimals,
                from_ata: &all_token_accounts[i],
                self_ata: &accounts[0],
                pool_signer: &accounts[8],
                token_billing_config: &accounts[1],
                fee_token_config: &accounts[11],
                ccip_router_pool_signer: &accounts[12],
                pool_accounts: accounts,
            });

            // calculate token indexes for ccipSend CPI call
            let ccip_token_index: u8 = if i == 0 {
                0 // start case
            } else {
                // calculate start index based on previous index and previous set of accounts
                ccip_token_indexes[i - 1] + parsed[i - 1].pool_accounts.len() as u8
            };
            ccip_token_indexes.push(ccip_token_index);

            // validate pool entries match
            // router will conduct further validation - this makes sure that accounts passed match
            let lookup_table_data = &mut &accounts[3].data.borrow()[..];
            let lookup_table_account: AddressLookupTable =
                AddressLookupTable::deserialize(lookup_table_data).unwrap();
            let expected = accounts[3..]
                .iter()
                .map(|e| e.key())
                .collect::<Vec<Pubkey>>();
            assert_eq!(lookup_table_account.addresses.as_ref(), expected);
        }
    };

    Ok((parsed, ccip_token_indexes))
}
