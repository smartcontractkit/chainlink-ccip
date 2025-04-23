import {
  AddressLookupTableProgram,
  Keypair,
  PublicKey,
  sendAndConfirmTransaction,
  Transaction,
} from "@solana/web3.js";
import { readFileSync } from "fs";
import {
  getAssociatedTokenAddressSync,
} from "@solana/spl-token";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import accounts, { tokenAdminRegistry } from "../staging";

async function createLookupTable() {
  const keypair = Keypair.fromSecretKey(
    new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
  );

  const config = getCCIPSendConfig("devnet");
  const connection = config.connection;

  const payer = keypair;
  const router = config.ccipRouterProgramId;
  const mint = tokenAdminRegistry.mint;
  const pool_program = tokenAdminRegistry.token_pool_program;
  const token_program = tokenAdminRegistry.token_program;
  const fee_quoter = accounts.addresses.feeQuoter;

  // Derive addresses
  const [token_admin_registry] = PublicKey.findProgramAddressSync(
    [Buffer.from("token_admin_registry"), mint.toBuffer()],
    router
  );

  const [pool_config] = PublicKey.findProgramAddressSync(
    [Buffer.from("ccip_tokenpool_config"), mint.toBuffer()],
    pool_program
  );

  const [pool_signer] = PublicKey.findProgramAddressSync(
    [Buffer.from("ccip_tokenpool_signer"), mint.toBuffer()],
    pool_program
  );

  const pool_token_account = getAssociatedTokenAddressSync(
    mint,
    pool_signer,
    true,
    token_program
  );

  const [fee_token_config] = PublicKey.findProgramAddressSync(
    [Buffer.from("fee_billing_token_config"), mint.toBuffer()],
    fee_quoter
  );

  const [ccip_router_pool_signer] = PublicKey.findProgramAddressSync(
    [Buffer.from("external_token_pools_signer"), pool_program.toBuffer()],
    router
  );

  // Step 1: Create LUT
  const slot = await connection.getSlot("finalized");
  const [createIx, lookupTableAddress] =
    AddressLookupTableProgram.createLookupTable({
      authority: payer.publicKey,
      payer: payer.publicKey,
      recentSlot: slot,
    });

  // Step 2: Extend LUT
  const extendIx = AddressLookupTableProgram.extendLookupTable({
    lookupTable: lookupTableAddress,
    authority: payer.publicKey,
    payer: payer.publicKey,
    addresses: [
      lookupTableAddress,
      token_admin_registry,
      pool_program,
      pool_config,
      pool_token_account,
      pool_signer,
      token_program,
      mint,
      fee_token_config,
      ccip_router_pool_signer,
    ],
  });

  const tx = new Transaction().add(createIx, extendIx);
  const sig = await sendAndConfirmTransaction(connection, tx, [payer]);

  console.log("✅ Lookup Table created and extended");
  console.log("🔑 LUT Address:", lookupTableAddress.toBase58());
  console.log("📤 Tx:", sig);
}

createLookupTable().catch(console.error);
