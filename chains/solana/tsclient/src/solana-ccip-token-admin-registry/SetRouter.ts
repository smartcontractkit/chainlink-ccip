import { loadTokenPoolProgram } from "./bnm-instructions";
import { sendTransaction } from "./program-instructions";

async function main() {

  const bnMProgramContext = await loadTokenPoolProgram();

  const ix = await bnMProgramContext.program.methods.setRouter(
    bnMProgramContext.router
  ).accounts(
    {
      state: bnMProgramContext.statePda,
      mint: bnMProgramContext.mint,
      authority: bnMProgramContext.keypair.publicKey,
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
