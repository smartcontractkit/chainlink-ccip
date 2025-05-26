import { PublicKey } from '@solana/web3.js';

export class SolanaCCIPPDAs {
  // Utility: convert u64 to 8-byte little endian buffer
  static uint64ToLE(n: number | bigint): Buffer {
    const bn = BigInt(n);
    const buf = Buffer.alloc(8);
    buf.writeBigUInt64LE(bn);
    return buf;
  }

  // //////////////////////
  // CCIP Router PDAs
  // //////////////////////

  static findConfigPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('config')], programId);
  }

  static findFeeBillingSignerPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('fee_billing_signer')], programId);
  }

  static findTokenAdminRegistryPDA(mint: PublicKey, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('token_admin_registry'), mint.toBuffer()],
      programId
    );
  }

  static findDestChainStatePDA(chainSelector: bigint, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('dest_chain_state'), this.uint64ToLE(chainSelector)],
      programId
    );
  }

  static findNoncePDA(chainSelector: bigint, authority: PublicKey, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('nonce'), this.uint64ToLE(chainSelector), authority.toBuffer()],
      programId
    );
  }

  static findApprovedSenderPDA(chainSelector: bigint, sourceSender: Buffer, receiverProgram: PublicKey) {
    const lenPrefix = Buffer.from([sourceSender.length]);
    return PublicKey.findProgramAddressSync(
      [Buffer.from('approved_ccip_sender'), this.uint64ToLE(chainSelector), lenPrefix, sourceSender],
      receiverProgram
    );
  }

  static findAllowedOfframpPDA(chainSelector: bigint, offramp: PublicKey, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('allowed_offramp'), this.uint64ToLE(chainSelector), offramp.toBuffer()],
      programId
    );
  }

  static findTokenPoolChainConfigPDA(
    chainSelector: bigint,
    tokenMint: PublicKey,
    programId: PublicKey
  ) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('ccip_tokenpool_chainconfig'), this.uint64ToLE(chainSelector), tokenMint.toBuffer()],
      programId
    );
  }

  // //////////////////////////
  // External Shared PDAs
  // //////////////////////////

  static findExternalTokenPoolsSignerPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('external_token_pools_signer')], programId);
  }

  static findExternalExecutionConfigPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('external_execution_config')], programId);
  }

  // /////////////////////
  // Fee Quoter PDAs
  // /////////////////////

  static findFqConfigPDA(feeQuoter: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('config')], feeQuoter);
  }

  static findFqDestChainPDA(chainSelector: bigint, feeQuoter: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('dest_chain'), this.uint64ToLE(chainSelector)],
      feeQuoter
    );
  }

  static findFqBillingTokenConfigPDA(mint: PublicKey, feeQuoter: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('fee_billing_token_config'), mint.toBuffer()],
      feeQuoter
    );
  }

  static findFqPerChainPerTokenConfigPDA(chainSelector: bigint, mint: PublicKey, feeQuoter: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('per_chain_per_token_config'), this.uint64ToLE(chainSelector), mint.toBuffer()],
      feeQuoter
    );
  }

  static findFqAllowedPriceUpdaterPDA(priceUpdater: PublicKey, feeQuoter: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('allowed_price_updater'), priceUpdater.toBuffer()],
      feeQuoter
    );
  }

  // ////////////////////
  // Offramp PDAs
  // ////////////////////

  static findOfframpConfigPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('config')], programId);
  }

  static findOfframpReferenceAddressesPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('reference_addresses')], programId);
  }

  static findOfframpSourceChainPDA(chainSelector: bigint, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('source_chain_state'), this.uint64ToLE(chainSelector)],
      programId
    );
  }

  static findOfframpBillingSignerPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('fee_billing_signer')], programId);
  }

  static findOfframpStatePDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('state')], programId);
  }

  static findOfframpCommitReportPDA(chainSelector: bigint, root: Buffer, programId: PublicKey) {
    return PublicKey.findProgramAddressSync(
      [Buffer.from('commit_report'), this.uint64ToLE(chainSelector), root],
      programId
    );
  }

  // ////////////////////
  // RMN Remote PDAs
  // ////////////////////

  static findRMNRemoteConfigPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('config')], programId);
  }

  static findRMNRemoteCursesPDA(programId: PublicKey) {
    return PublicKey.findProgramAddressSync([Buffer.from('curses')], programId);
  }
}
