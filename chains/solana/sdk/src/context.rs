use solana_sdk::pubkey::Pubkey;
use solana_program::account_info::AccountInfo;
use std::fmt::Debug;


pub struct Context<'a, 'b, 'c, 'info, T> {
    /// Currently executing program id.
    pub program_id: &'a Pubkey,
    /// Deserialized accounts.
    pub accounts: &'b mut T,
    /// Remaining accounts given but not deserialized or validated.
    /// Be very careful when using this directly.
    pub remaining_accounts: &'c [AccountInfo<'info>],
}


pub struct CcipSendAccounts<'a> {
    // ---- local program state
    pub state: &'a AccountInfo<'a>,
    pub chain_config: &'a AccountInfo<'a>,
    pub ccip_sender: &'a AccountInfo<'a>,

    // ---- fee payer/authority
    pub authority_fee_token_ata: &'a AccountInfo<'a>,
    pub authority: &'a AccountInfo<'a>,
    pub system_program: &'a AccountInfo<'a>,

    // ---- required CCIP accounts (validated by router CPI)
    pub ccip_router: &'a AccountInfo<'a>,
    pub ccip_config: &'a AccountInfo<'a>,
    pub ccip_dest_chain_state: &'a AccountInfo<'a>,
    pub ccip_sender_nonce: &'a AccountInfo<'a>,
    pub ccip_fee_token_program: &'a AccountInfo<'a>,
    pub ccip_fee_token_mint: &'a AccountInfo<'a>,
    pub ccip_fee_token_user_ata: &'a AccountInfo<'a>,
    pub ccip_fee_token_receiver: &'a AccountInfo<'a>,
    pub ccip_fee_billing_signer: &'a AccountInfo<'a>,
    pub ccip_fee_quoter: &'a AccountInfo<'a>,
    pub ccip_fee_quoter_config: &'a AccountInfo<'a>,
    pub ccip_fee_quoter_dest_chain: &'a AccountInfo<'a>,
    pub ccip_fee_quoter_billing_token_config: &'a AccountInfo<'a>,
    pub ccip_fee_quoter_link_token_config: &'a AccountInfo<'a>,
    pub ccip_rmn_remote: &'a AccountInfo<'a>,
    pub ccip_rmn_remote_curses: &'a AccountInfo<'a>,
    pub ccip_rmn_remote_config: &'a AccountInfo<'a>,
}
