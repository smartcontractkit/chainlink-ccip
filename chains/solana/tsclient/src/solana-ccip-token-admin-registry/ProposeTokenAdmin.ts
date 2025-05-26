import { tokenAdminRegistry } from "../staging";
import { loadCcipRouterProgram } from "./token-admin-registry-instructions";
import { sendTransaction, SYSTEM_PROGRAM_ID } from "./program-instructions";

async function main() {

    const ccipRouterTokenAdminRegistryContext = await loadCcipRouterProgram();

    const ix = await ccipRouterTokenAdminRegistryContext.program.methods.ownerProposeAdministrator(
        tokenAdminRegistry.user
    ).accounts(
        {
            config: ccipRouterTokenAdminRegistryContext.configPDA,
            tokenAdminRegistry: ccipRouterTokenAdminRegistryContext.tokenAdminRegistryPDA,
            mint: ccipRouterTokenAdminRegistryContext.mint,
            authority: ccipRouterTokenAdminRegistryContext.keypair.publicKey,
            systemProgram: SYSTEM_PROGRAM_ID,
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
