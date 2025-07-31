use anchor_lang::prelude::*;

declare_id!("EGfB7iiotGoDVpQvByFD8AD11BhTpc9WMCyUL5q64smj");

#[program]
pub mod test_event_emitter {
    use super::*;

    pub fn emit_ccip_cctp_msg_sent(
        _ctx: Context<EventEmitter>,
        args: CcipCctpMessageSentEventArgs,
    ) -> Result<()> {
        emit!(CcipCctpMessageSentEvent {
            remote_chain_selector: args.remote_chain_selector,
            message_sent_bytes: args.message_sent_bytes,
            msg_total_nonce: args.msg_total_nonce,
            original_sender: args.original_sender,
            event_address: args.event_address,
            source_domain: args.source_domain,
            cctp_nonce: args.cctp_nonce,
        });
        Ok(())
    }
}

#[derive(Accounts, Debug)]
pub struct EventEmitter<'info> {
    // This is unused, but Anchor requires that there is at least one account in the context
    pub clock: Sysvar<'info, Clock>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Debug)]
pub struct CcipCctpMessageSentEventArgs {
    pub message_sent_bytes: Vec<u8>,
    pub remote_chain_selector: u64,
    pub original_sender: Pubkey,
    pub event_address: Pubkey,
    pub msg_total_nonce: u64,
    pub source_domain: u32,
    pub cctp_nonce: u64,
}

#[event]
pub struct CcipCctpMessageSentEvent {
    // Seeds for the CCTP message sent event account
    pub original_sender: Pubkey,
    pub remote_chain_selector: u64,
    pub msg_total_nonce: u64,

    // Actual event account address, derived from the seeds above
    pub event_address: Pubkey,

    // CCTP values identifying the message
    pub source_domain: u32, // The source chain domain ID, which for Solana is always 5
    pub cctp_nonce: u64,

    // CCTP message bytes, used to get the attestation offchain and receive the message on dest
    pub message_sent_bytes: Vec<u8>,
}
