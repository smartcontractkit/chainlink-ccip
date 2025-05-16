import {
    PublicKey,
    TransactionMessage,
    VersionedTransaction,
} from "@solana/web3.js";
import { tokenAdminRegistry } from "../staging";
import { Buffer } from "buffer";

import { loadTokenPoolProgram } from "./bnm-instructions";
import BN from "bn.js";


async function main() {

    const bnMProgramContext = await loadTokenPoolProgram();

    const remoteChainSelector = tokenAdminRegistry.remoteChainSelector;
    const remoteChainSelector64 = remoteChainSelector.toArrayLike(Buffer, "le", 8);

    const capacity = new BN(100);
    const rate = new BN(1);

    const [chainConfigPda] = PublicKey.findProgramAddressSync(
        [
            Buffer.from("ccip_tokenpool_chainconfig"),
            remoteChainSelector64,
            bnMProgramContext.mint.toBuffer(),
        ],
        bnMProgramContext.programId
    );

    const ix = await bnMProgramContext.program.methods.setChainRateLimit(
        remoteChainSelector,
        tokenAdminRegistry.mint,
        { // inbound
            enabled: true,
            capacity,
            rate,
        },
        { // outbound
            enabled: true,
            capacity,
            rate,
        },
    ).accounts({
        state: bnMProgramContext.statePda,
        chainConfig: chainConfigPda,
        authority: bnMProgramContext.keypair.publicKey,
    }).instruction();

    const { blockhash } = await bnMProgramContext.connection.getLatestBlockhash();

    const message = new TransactionMessage({
        payerKey: bnMProgramContext.keypair.publicKey,
        recentBlockhash: blockhash,
        instructions: [ix],
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

}

main()
    .then(() => {
        console.log("✅ Done");
    })
    .catch((err) => {
        console.error("❌ Error:", err);
    });
