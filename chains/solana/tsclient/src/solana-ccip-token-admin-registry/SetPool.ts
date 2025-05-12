import { Buffer } from "buffer";
import { sendTransaction } from "./program-instructions";
import { loadCcipRouterProgram } from "./token-admin-registry-instructions";


async function main() {

  const ccipRouterTokenAdminRegistryContext = await loadCcipRouterProgram();

  const ix = await ccipRouterTokenAdminRegistryContext.program.methods.setPool(
    Buffer.from([3, 4, 7])
  ).accounts(
    {
      config: ccipRouterTokenAdminRegistryContext.configPDA,
      tokenAdminRegistry: ccipRouterTokenAdminRegistryContext.tokenAdminRegistryPDA,
      mint: ccipRouterTokenAdminRegistryContext.mint,
      poolLookuptable: ccipRouterTokenAdminRegistryContext.lookupTable,
      authority: ccipRouterTokenAdminRegistryContext.keypair.publicKey,
    },
  ).instruction();

  await sendTransaction(ccipRouterTokenAdminRegistryContext.connection, ccipRouterTokenAdminRegistryContext.keypair, ix);
}

main()
  .then(() => {
    console.log("✅ Done");
  })
  .catch((err) => {
    console.error("❌ Error:", err);
  });
