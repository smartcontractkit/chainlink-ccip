import {
  createAssociatedTokenAccountInstruction,
  getAssociatedTokenAddressSync,
} from "@solana/spl-token";
import {
  Keypair,
  PublicKey,
} from "@solana/web3.js";
import { readFileSync } from "fs";
import { tokenAdminRegistry } from "../staging";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import { sendTransaction } from "./program-instructions";

async function main() {
  const keypair = Keypair.fromSecretKey(
    new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
  );

  const config = getCCIPSendConfig("devnet");
  const connection = config.connection;

  const mint = tokenAdminRegistry.mint;
  const pool_program = tokenAdminRegistry.token_pool_program;
  const token_program = tokenAdminRegistry.token_program;


  const [pool_signer] = PublicKey.findProgramAddressSync(
    [Buffer.from("ccip_tokenpool_signer"), mint.toBuffer()],
    pool_program
  );

  const ata = getAssociatedTokenAddressSync(
    mint,
    pool_signer,
    true,
    token_program
  );

  console.log("ℹ️ ATA to create:", ata.toBase58());

  const ix = createAssociatedTokenAccountInstruction(
    keypair.publicKey, // payer
    ata,               // ata address
    pool_signer,       // owner of the ATA
    mint,              // token mint
    token_program,
  );

  await sendTransaction(connection, keypair, ix);
}

main()
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });
