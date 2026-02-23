

use solana_program::{
    account_info::AccountInfo,
    instruction::{AccountMeta, Instruction},
    program::invoke_signed
};

use crate::{context::{CcipSendAccounts, Context}, discriminator::{CCIP_SENDER, CCIP_SEND_DISCRIMINATOR}};



pub fn ccip_send<'info>(
    program_id: AccountInfo<'info>,
    ctx: Context<'_, '_, 'info, 'info, CcipSendAccounts<'info>>,
) -> Result<()> {

        // minimum set of ccipSend accounts
        let mut acc_infos: Vec<AccountInfo> = [
            ctx.accounts.ccip_config.to_account_info(),
            ctx.accounts.ccip_dest_chain_state.to_account_info(),
            ctx.accounts.ccip_sender_nonce.to_account_info(),
            ctx.accounts.ccip_sender.to_account_info(),
            ctx.accounts.system_program.to_account_info(),
            ctx.accounts.ccip_fee_token_program.to_account_info(),
            ctx.accounts.ccip_fee_token_mint.to_account_info(),
            ctx.accounts.ccip_fee_token_user_ata.to_account_info(),
            ctx.accounts.ccip_fee_token_receiver.to_account_info(),
            ctx.accounts.ccip_fee_billing_signer.to_account_info(),
            ctx.accounts.ccip_fee_quoter.to_account_info(),
            ctx.accounts.ccip_fee_quoter_config.to_account_info(),
            ctx.accounts.ccip_fee_quoter_dest_chain.to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_billing_token_config
                .to_account_info(),
            ctx.accounts
                .ccip_fee_quoter_link_token_config
                .to_account_info(),
            ctx.accounts.ccip_rmn_remote.to_account_info(),
            ctx.accounts.ccip_rmn_remote_curses.to_account_info(),
            ctx.accounts.ccip_rmn_remote_config.to_account_info(),
        ]
        .to_vec();

    let acc_metas: Vec<AccountMeta> = acc_infos
        .to_vec()
        .iter()
        .flat_map(|acc_info| {
            // Check signer from PDA External Execution config
            let is_signer = acc_info.key() == ctx.accounts.ccip_sender.key();
            acc_info.to_account_metas(Some(is_signer))
        })
        .collect();

    let data = builder::instruction_with_token_indexes(
            &message,
            CCIP_SEND_DISCRIMINATOR,
            dest_chain_selector,
            &ccip_token_indexes,
        );

    let instruction = Instruction {
        program_id: *program_id.key,
        accounts: acc_metas,
        data: data.into_inner(),
    };

    let seeds = &[CCIP_SENDER, &[ctx.bumps.ccip_sender]];
    let signer = &[&seeds[..]];

    invoke_signed(&instruction, &acc_infos, ctx.signer_seeds,)
    .map_err(Into::into)
}