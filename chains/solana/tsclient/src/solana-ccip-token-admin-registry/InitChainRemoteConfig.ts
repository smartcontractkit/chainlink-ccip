import {
  PublicKey,
  TransactionMessage,
  VersionedTransaction,
  SystemProgram,
} from "@solana/web3.js";
import { EventParser, BorshCoder } from '@coral-xyz/anchor'
import { tokenAdminRegistry } from "../staging";
import { Buffer } from "buffer";

import { loadTokenPoolProgram } from "./bnm-instructions";


async function main() {

  const bnMProgramContext = await loadTokenPoolProgram();

  const remoteChainSelector = tokenAdminRegistry.remoteChainSelector;
  const remoteChainSelector64 = remoteChainSelector.toArrayLike(Buffer, "le", 8);

  const [chainConfigPda] = PublicKey.findProgramAddressSync(
    [
      Buffer.from("ccip_tokenpool_chainconfig"),
      remoteChainSelector64,
      bnMProgramContext.mint.toBuffer(),
    ],
    bnMProgramContext.programId
  );

  const ix = await bnMProgramContext.program.methods.initChainRemoteConfig(
    remoteChainSelector,
    tokenAdminRegistry.mint,
    {
      poolAddresses: [],
      tokenAddress: Uint8Array.from(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex")),
      decimals: 18,
    },
  ).accounts({
    state: bnMProgramContext.statePda,
    authority: bnMProgramContext.keypair.publicKey,
    systemProgram: SystemProgram.programId,
    chainConfig: chainConfigPda,
  }).instruction();

  const remoteAddressBytes = Buffer.from(tokenAdminRegistry.remoteAddress.replace(/^0x/, ""), "hex");
  const tokenAddressBytes = padTo32Bytes(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex"));

  const ix2 = await bnMProgramContext.program.methods.editChainRemoteConfig(
    remoteChainSelector,
    tokenAdminRegistry.mint,
    {
      poolAddresses: [{ address: remoteAddressBytes }],
      tokenAddress: { address: tokenAddressBytes },
      decimals: 18,
    },
  ).accounts({
    state: bnMProgramContext.statePda,
    authority: bnMProgramContext.keypair.publicKey,
    systemProgram: SystemProgram.programId,
    chainConfig: chainConfigPda,
  }).instruction();

  const { blockhash } = await bnMProgramContext.connection.getLatestBlockhash();

  const message = new TransactionMessage({
    payerKey: bnMProgramContext.keypair.publicKey,
    recentBlockhash: blockhash,
    instructions: [ix, ix2],
  }).compileToV0Message([]);

  const vtx = new VersionedTransaction(message);
  vtx.sign([bnMProgramContext.keypair]);

  console.log("🔍 Simulating...");
  const sim = await bnMProgramContext.connection.simulateTransaction(vtx);
  if (sim.value.err) {
    console.error("❌ Simulation failed:", sim.value.logs);
    throw new Error("Simulation failed");
  }

  console.log("✅ Simulation passed");
  const sig = await bnMProgramContext.connection.sendTransaction(vtx, {
    skipPreflight: false,
    preflightCommitment: "confirmed",
  });
  console.log("📤 Tx sent:", sig);
  console.log("🌐 Chain Config PDA:", chainConfigPda.toBase58());

  await new Promise((resolve) => setTimeout(resolve, 5000));
  // Verify that the chain config PDA contains the expected data
  const tx = await bnMProgramContext.provider.connection.getTransaction(sig, {
    commitment: "confirmed",
    maxSupportedTransactionVersion: 0,
  });

  const eventParser = new EventParser(bnMProgramContext.programId, new BorshCoder(bnMProgramContext.idl));
  const events = eventParser.parseLogs(tx?.meta?.logMessages!);
  for (let event of events) {
    console.log(event);
  }

}

main()
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });

function padTo32Bytes(buffer: Buffer): Buffer {
  if (buffer.length >= 32) {
    return buffer;
  }

  // Create a new buffer of 32 bytes
  const paddedBuffer = Buffer.alloc(32, 0); // Initialize with zeros

  // Copy the original buffer data to the end of the new buffer (right-aligned)
  // This is the standard Ethereum-style padding
  buffer.copy(paddedBuffer, 32 - buffer.length);

  return paddedBuffer;
}