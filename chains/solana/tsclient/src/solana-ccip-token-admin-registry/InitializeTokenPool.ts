import {
    Keypair,
    PublicKey,
    TransactionMessage,
    VersionedTransaction,
  } from "@solana/web3.js";
import { readFileSync } from "fs";
import { tokenAdminRegistry } from "../staging";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import { createInitializeInstruction } from "./bnm-instructions";

const POOL_STATE_SEED = Buffer.from("ccip_tokenpool_config");

const keypair = Keypair.fromSecretKey(
new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
);

const config = getCCIPSendConfig("devnet");
const connection = config.connection;

const BPF_LOADER_UPGRADEABLE_PROGRAM_ID = new PublicKey("BPFLoaderUpgradeab1e11111111111111111111111");

export function getProgramDataAddress(programId: PublicKey): PublicKey {
  const [programData] = PublicKey.findProgramAddressSync(
    [programId.toBuffer()],
    BPF_LOADER_UPGRADEABLE_PROGRAM_ID
  );
  return programData;
}

async function main() {
  const PROGRAM_ID = tokenAdminRegistry.token_pool_program;
  const MINT = tokenAdminRegistry.mint;

  const [statePda] = PublicKey.findProgramAddressSync(
    [POOL_STATE_SEED, MINT.toBuffer()],
    PROGRAM_ID
  );
  console.log("Using PROGRAM_ID:", PROGRAM_ID.toBase58());

  const ix = createInitializeInstruction(
    {
      state: statePda,
      mint: MINT,
      authority: keypair.publicKey,
      systemProgram: config.systemProgramId,
      program: PROGRAM_ID,
      programData: getProgramDataAddress(PROGRAM_ID),
    },
    {
      router: config.ccipRouterProgramId,
      rmnRemote: config.rmnRemoteProgramId,
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
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });
