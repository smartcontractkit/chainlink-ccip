import fs from "fs";
import {
  Keypair,
  PublicKey,
  TransactionMessage,
  VersionedTransaction,
  SystemProgram,
  Connection,
  Commitment,
} from "@solana/web3.js";
import { Program, Idl, AnchorProvider, Wallet, EventParser, BorshCoder } from '@coral-xyz/anchor'
import { readFileSync } from "fs";
import { tokenAdminRegistry } from "../staging";
import { Buffer } from "buffer";
import { BurnmintTokenPool } from '../../../contracts/target/types/burnmint_token_pool';

const POOL_STATE_SEED = Buffer.from("ccip_tokenpool_config");

const keypair = Keypair.fromSecretKey(
  new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
);

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
  const connection = new Connection(
    "https://api.devnet.solana.com",
    { commitment: 'confirmed' as Commitment }
  )
  const wallet = new Wallet(keypair)
  const provider = new AnchorProvider(connection, wallet, {
    preflightCommitment: 'confirmed' as Commitment,
    commitment: 'confirmed' as Commitment,
  })

  const BnMTokenPoolProgram = new Program<BurnmintTokenPool>(idlJson as BurnmintTokenPool, PROGRAM_ID, provider);

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

  const remoteAddressBytes = padTo32Bytes(Buffer.from(tokenAdminRegistry.remoteAddress.replace(/^0x/, ""), "hex"));
  const tokenAddressBytes = padTo32Bytes(Buffer.from(tokenAdminRegistry.tokenAddress.replace(/^0x/, ""), "hex"));

  const ix2 = await BnMTokenPoolProgram.methods.editChainRemoteConfig(
    remoteChainSelector,
    tokenAdminRegistry.mint,
    {
      poolAddresses: [{ address: remoteAddressBytes }],
      tokenAddress: { address: tokenAddressBytes },
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

  const vtx = new VersionedTransaction(message);
  vtx.sign([keypair]);

  console.log("🔍 Simulating...");
  const sim = await connection.simulateTransaction(vtx);
  if (sim.value.err) {
    console.error("❌ Simulation failed:", sim.value.logs);
    throw new Error("Simulation failed");
  }

  console.log("✅ Simulation passed");
  const sig = await connection.sendTransaction(vtx, {
    skipPreflight: false,
    preflightCommitment: "confirmed",
  });
  console.log("📤 Tx sent:", sig);
  console.log("🌐 Chain Config PDA:", chainConfigPda.toBase58());

  await new Promise((resolve) => setTimeout(resolve, 5000));
  // Verify that the chain config PDA contains the expected data
  const tx = await provider.connection.getTransaction(sig, {
    commitment: "confirmed",
    maxSupportedTransactionVersion: 0,
  });

  const eventParser = new EventParser(PROGRAM_ID, new BorshCoder(idlJson));
  const events = eventParser.parseLogs(tx?.meta?.logMessages!);
  for (let event of events) {
    console.log(event);
  }

  const tpstate = (await BnMTokenPoolProgram.account.state.fetch("4U55f3yBpWMwcgvE6MvGg7Eqe46gZdq9YmUKyqVi4x96", 'confirmed')) as any;
  console.log("📜 Chain Config State:", JSON.stringify(tpstate, null, "  "));

  const chainConfig = (await BnMTokenPoolProgram.account.chainConfig.fetch("4kKMcdVGKQ3idqjMc19LmsST3wXZrRQ6xMpf8w1QybJc", 'confirmed')) as any;
  console.log("📜 Whatever State:", JSON.stringify(chainConfig, null, "  "));

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