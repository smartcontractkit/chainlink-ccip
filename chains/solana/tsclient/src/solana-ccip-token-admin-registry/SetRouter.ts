import {
    Keypair,
    PublicKey,
    TransactionMessage,
    VersionedTransaction,
  } from "@solana/web3.js";
  import { readFileSync } from "fs";
  import { tokenAdminRegistry } from "../staging";
  import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import { createSetRouterInstruction } from "./bnm-instructions";

const POOL_STATE_SEED = Buffer.from("ccip_tokenpool_config");
  const keypair = Keypair.fromSecretKey(
    new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
  );
  const config = getCCIPSendConfig("devnet");
  const connection = config.connection;

  async function main() {
    const PROGRAM_ID = tokenAdminRegistry.token_pool_program;
    const MINT = tokenAdminRegistry.mint;

    const [statePda] = PublicKey.findProgramAddressSync(
      [POOL_STATE_SEED, MINT.toBuffer()],
      PROGRAM_ID
    );

    const ix = createSetRouterInstruction(
      {
        state: statePda,
        mint: MINT,
        authority: keypair.publicKey,
      },
      {
        newRouter: config.ccipRouterProgramId,
      },
      PROGRAM_ID
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
    console.log("📦 State PDA:", statePda.toBase58());
  }

  main()
    .then(() => console.log("✅ Done"))
    .catch((err) => console.error("❌ Error:", err));
