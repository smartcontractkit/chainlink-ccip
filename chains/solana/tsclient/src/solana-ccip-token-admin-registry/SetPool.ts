import {
    TransactionMessage,
    VersionedTransaction,
    PublicKey,
    Keypair,
  } from "@solana/web3.js";
  import { readFileSync } from "fs";
  import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
  import { tokenAdminRegistry } from "../staging";
  import * as borsh from "@coral-xyz/borsh";
  import crypto from "crypto";
  import { Buffer } from "buffer";

  // Layout & discriminator
  const setPoolLayout = borsh.struct([borsh.vecU8("writableIndexes")]);

  const SET_POOL_DISCRIMINATOR = crypto
    .createHash("sha256")
    .update("global:set_pool")
    .digest()
    .subarray(0, 8);

  // Main function
  async function setPool() {
    const keypair = Keypair.fromSecretKey(
      new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
    );

    const config = getCCIPSendConfig("devnet");
    const connection = config.connection;
    const PROGRAM_ID = config.ccipRouterProgramId;
    const MINT = tokenAdminRegistry.mint;
    const LOOKUP_TABLE = tokenAdminRegistry.lookup_table;

    const [CONFIG_PDA] = PublicKey.findProgramAddressSync(
      [Buffer.from("config")],
      PROGRAM_ID
    );

    const [tokenAdminRegistryPDA] = PublicKey.findProgramAddressSync(
      [Buffer.from("token_admin_registry"), MINT.toBuffer()],
      PROGRAM_ID
    );

    // Define writableIndexes bitmap: 10 accounts, 3 writable: [1, 4, 5]
    const bitmap = new Uint8Array(2); // up to 16 bits
    bitmap[0] = 0b00110010; // bit 1, 4, 5 set
    const args = { writableIndexes: Buffer.from(bitmap) };

    const data = Buffer.alloc(100);
    const len = setPoolLayout.encode(args, data);
    const encodedArgs = Buffer.concat([
      SET_POOL_DISCRIMINATOR,
      data.subarray(0, len),
    ]);

    const keys = [
      { pubkey: CONFIG_PDA, isWritable: false, isSigner: false },
      { pubkey: tokenAdminRegistryPDA, isWritable: true, isSigner: false },
      { pubkey: MINT, isWritable: false, isSigner: false },
      { pubkey: LOOKUP_TABLE, isWritable: false, isSigner: false },
      { pubkey: keypair.publicKey, isWritable: true, isSigner: true },
    ];

    const ix = {
      programId: PROGRAM_ID,
      keys,
      data: encodedArgs,
    };

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

  setPool().catch(console.error);
