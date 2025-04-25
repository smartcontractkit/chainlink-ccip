import {
    Keypair,
    PublicKey,
    TransactionMessage,
    VersionedTransaction,
    SystemProgram,
  } from "@solana/web3.js";
  import { readFileSync } from "fs";
  import { tokenAdminRegistry } from "../staging";
  import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
  import { createInitChainRemoteConfigInstruction, RemoteAddress } from "./bnm-instructions";
  import { Buffer } from "buffer";

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
    console.log("Using PROGRAM_ID:", PROGRAM_ID.toBase58());

    const remoteChainSelectorBigInt = BigInt(tokenAdminRegistry.remoteChainSelector.toString());
    const remoteChainSelectorBuf = Buffer.alloc(8);
    remoteChainSelectorBuf.writeBigUInt64LE(remoteChainSelectorBigInt);


    const [chainConfigPda] = PublicKey.findProgramAddressSync(
      [
        Buffer.from("chain_config"),
        remoteChainSelectorBuf,
        MINT.toBuffer(),
      ],
      PROGRAM_ID
    );
    const poolAddresses: RemoteAddress[] = [
        { address: Uint8Array.from(Buffer.from(tokenAdminRegistry.remoteAddress.replace(/^0x/, ""), "hex")) }
      ];

      const tokenAddress: RemoteAddress = {
        address: Uint8Array.from(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex")),
      };


      const ix = createInitChainRemoteConfigInstruction(
        {
          state: statePda,
          chainConfig: chainConfigPda,
          authority: keypair.publicKey,
          systemProgram: SystemProgram.programId,
        },
        {
          remoteChainSelector: remoteChainSelectorBigInt,
          cfg: {
            poolAddresses,
            tokenAddress,
            decimals: 8,
          },
          mint: tokenAdminRegistry.mint,
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
    console.log("🌐 Chain Config PDA:", chainConfigPda.toBase58());
  }

  main()
    .then(() => {
      console.log("✅ Done");
    })
    .catch((err) => {
      console.error("❌ Error:", err);
    });
