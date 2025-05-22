import {
    Connection,
    Keypair,
    PublicKey,
    TransactionInstruction,
    TransactionMessage,
    VersionedTransaction,
} from "@solana/web3.js";

export const BPF_LOADER_UPGRADEABLE_PROGRAM_ID = new PublicKey("BPFLoaderUpgradeab1e11111111111111111111111");
export const SYSTEM_PROGRAM_ID = new PublicKey('11111111111111111111111111111111');

export async function sendTransaction(connection: Connection, keypair: Keypair, ix: TransactionInstruction) {
    const { blockhash } = await connection.getLatestBlockhash();

    const message = new TransactionMessage({
        payerKey: keypair.publicKey,
        recentBlockhash: blockhash,
        instructions: [ix],
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
}

