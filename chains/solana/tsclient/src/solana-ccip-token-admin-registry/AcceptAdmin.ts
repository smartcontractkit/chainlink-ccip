import { sendTransaction } from "./program-instructions";
import { loadCcipRouterProgram } from "./token-admin-registry-instructions";


async function main() {

  const ccipRouterTokenAdminRegistryContext = await loadCcipRouterProgram();

  const ix = await ccipRouterTokenAdminRegistryContext.program.methods.acceptAdminRoleTokenAdminRegistry(
  ).accounts(
    {
      config: ccipRouterTokenAdminRegistryContext.configPDA,
      tokenAdminRegistry: ccipRouterTokenAdminRegistryContext.tokenAdminRegistryPDA,
      mint: ccipRouterTokenAdminRegistryContext.mint,
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
