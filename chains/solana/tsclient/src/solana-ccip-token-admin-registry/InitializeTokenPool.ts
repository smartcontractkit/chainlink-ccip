import { getProgramDataAddress, loadTokenPoolProgram } from "./bnm-instructions";
import { sendTransaction, SYSTEM_PROGRAM_ID } from "./program-instructions";

async function main() {

  const bnMProgramContext = await loadTokenPoolProgram();

  const ix = await bnMProgramContext.program.methods.initialize(
    bnMProgramContext.router,
    bnMProgramContext.rmnRemote).accounts(
      {
        state: bnMProgramContext.statePda,
        mint: bnMProgramContext.mint,
        authority: bnMProgramContext.keypair.publicKey,
        systemProgram: SYSTEM_PROGRAM_ID,
        program: bnMProgramContext.programId,
        programData: getProgramDataAddress(bnMProgramContext.programId),
      }
    ).instruction();

  await sendTransaction(bnMProgramContext.connection, bnMProgramContext.keypair, ix);
}

main()
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });
