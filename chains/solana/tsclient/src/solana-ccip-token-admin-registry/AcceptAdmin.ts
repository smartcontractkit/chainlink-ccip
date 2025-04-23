import {
    TransactionMessage,
    VersionedTransaction,
    PublicKey,
    Keypair,
  } from "@solana/web3.js";
  import { readFileSync } from "fs";
  import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
  import { createAcceptAdminRoleInstruction } from "./setup-token-admin-registry";
  import { tokenAdminRegistry } from "../staging";

  async function acceptAdmin() {
    const keypair = Keypair.fromSecretKey(
      new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
    );

    const config = getCCIPSendConfig("devnet");
    const connection = config.connection;
    const PROGRAM_ID = config.ccipRouterProgramId;
    const MINT = tokenAdminRegistry.mint;

    const [CONFIG_PDA] = PublicKey.findProgramAddressSync(
      [Buffer.from("config")],
      PROGRAM_ID
    );

    const [tokenAdminRegistryPDA] = PublicKey.findProgramAddressSync(
      [Buffer.from("token_admin_registry"), MINT.toBuffer()],
      PROGRAM_ID
    );

    const ix = createAcceptAdminRoleInstruction(
      {
        config: CONFIG_PDA,
        tokenAdminRegistry: tokenAdminRegistryPDA,
        mint: MINT,
        authority: keypair.publicKey,
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

    console.log("\uD83D\uDD0D Simulating...");
    const sim = await connection.simulateTransaction(tx);
    if (sim.value.err) {
      console.error("\u274C Simulation failed:", sim.value.logs);
      throw new Error("Simulation failed");
    }

    console.log("\u2705 Simulation passed");
    const sig = await connection.sendTransaction(tx);
    console.log("\uD83D\uDCE4 Tx sent:", sig);
  }

  acceptAdmin().catch(console.error);