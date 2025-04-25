import {
  createAssociatedTokenAccountInstruction,
  getAssociatedTokenAddressSync,
} from "@solana/spl-token";
import {
  Keypair,
  PublicKey,
  TransactionMessage,
  VersionedTransaction,
} from "@solana/web3.js";
import { readFileSync } from "fs";
import { tokenAdminRegistry } from "../staging";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";

async function createATA() {
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

  const { blockhash } = await connection.getLatestBlockhash();

  const message = new TransactionMessage({
    payerKey: keypair.publicKey,
    recentBlockhash: blockhash,
    instructions: [ix],
  }).compileToV0Message([]);

  const tx = new VersionedTransaction(message);
  tx.sign([keypair]);

  console.log("🔍 Simulating...");
  const sim = await connection.simulateTransaction(tx);
  if (sim.value.err) {
    console.error("❌ Simulation failed:", sim.value.logs);
    throw new Error("Simulation failed");
  }

  console.log("✅ Simulation passed");
  const sig = await connection.sendTransaction(tx);
  console.log("📤 Tx sent:", sig);
  console.log("🏦 Created ATA:", ata.toBase58());
}

createATA()
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });
