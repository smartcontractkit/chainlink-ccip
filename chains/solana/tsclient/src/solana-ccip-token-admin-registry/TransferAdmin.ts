import {
    TransactionMessage,
    VersionedTransaction,
    PublicKey,
    Keypair,
  } from "@solana/web3.js";
  import { readFileSync } from "fs";
  import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
  import { createTransferAdminRoleInstruction } from "./setup-token-admin-registry";
  import { tokenAdminRegistry } from "../staging";

  async function transferAdmin() {
    const keypair = Keypair.fromSecretKey(
      new Uint8Array(JSON.parse(readFileSync("../contracts/id.json", "utf-8")))
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


    const [pool_signer] = PublicKey.findProgramAddressSync(
        [Buffer.from("ccip_tokenpool_signer"), tokenAdminRegistry.mint.toBuffer()],
        tokenAdminRegistry.token_pool_program
    );

    const ix = createTransferAdminRoleInstruction(
      {
        config: CONFIG_PDA,
        tokenAdminRegistry: tokenAdminRegistryPDA,
        mint: MINT,
        authority: keypair.publicKey,
      },
      {
        newAdmin: pool_signer,
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
  }

  transferAdmin().catch(console.error);
