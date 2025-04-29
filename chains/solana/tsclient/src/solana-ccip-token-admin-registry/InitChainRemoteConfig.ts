import fs from "fs";
import {
  Keypair,
  PublicKey,
  TransactionMessage,
  VersionedTransaction,
  SystemProgram,
} from "@solana/web3.js";
import { Program, Idl } from '@coral-xyz/anchor'
import { readFileSync } from "fs";
import { tokenAdminRegistry } from "../staging";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import { Buffer } from "buffer";
import { BurnmintTokenPool } from '../../../contracts/target/types/burnmint_token_pool';

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

  const remoteChainSelector = tokenAdminRegistry.remoteChainSelector;
  const remoteChainSelector64 = remoteChainSelector.toArrayLike(Buffer, "le", 8);

  const [chainConfigPda] = PublicKey.findProgramAddressSync(
    [
      Buffer.from("ccip_tokenpool_chainconfig"),
      remoteChainSelector64,
      MINT.toBuffer(),
    ],
    PROGRAM_ID
  );

  const idlFile = fs.readFileSync(tokenAdminRegistry.burnmint_token_pool_idl_path);
  const idlJson: Idl = JSON.parse(idlFile.toString());
  const BnMTokenPoolProgram = new Program<BurnmintTokenPool>(idlJson as BurnmintTokenPool, PROGRAM_ID);

  const ix = await BnMTokenPoolProgram.methods.initChainRemoteConfig(
    remoteChainSelector,
    tokenAdminRegistry.mint,
    {
      poolAddresses: [],
      tokenAddress: Uint8Array.from(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex")),
      decimals: 8,
    },
  ).accounts({
    state: statePda,
    authority: keypair.publicKey,
    systemProgram: SystemProgram.programId,
    chainConfig: chainConfigPda,
  }).instruction();

  const ix2 = await BnMTokenPoolProgram.methods.editChainRemoteConfig(
    remoteChainSelector,
    tokenAdminRegistry.mint,
    {
      poolAddresses: [Uint8Array.from(Buffer.from(tokenAdminRegistry.remoteAddress.replace(/^0x/, ""), "hex"))],
      tokenAddress: Uint8Array.from(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex")),
      decimals: 8,
    },
  ).accounts({
    state: statePda,
    authority: keypair.publicKey,
    systemProgram: SystemProgram.programId,
    chainConfig: chainConfigPda,
  }).instruction();

  const { blockhash } = await connection.getLatestBlockhash();

  const message = new TransactionMessage({
    payerKey: keypair.publicKey,
    recentBlockhash: blockhash,
    instructions: [ix, ix2],
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
