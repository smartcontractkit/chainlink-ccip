import * as anchor from "@coral-xyz/anchor";
import { Program, web3 } from "@coral-xyz/anchor";

import { CcipCommon } from "../../../contracts/target/types/ccip_common";

const ccipCommon = anchor.workspace.CcipCommon as Program<CcipCommon>;

// const [tokenAdminRegistryAddress] = web3.PublicKey.findProgramAddressSync(
//   [Buffer.from("token_admin_registry"), mint.toBuffer()],
//   routerAddress,
// );

// getTokenAdminRegistryAccount(ccipCommon, tokenAdminRegistryAddress).catch(
//   console.error,
// );

export async function getTokenAdminRegistryAccount(
  program: Program<CcipCommon>,
  tokenAdminRegistry: anchor.web3.PublicKey,
): Promise<{
  version: number;
  administrator: anchor.web3.PublicKey;
  pendingAdministrator: anchor.web3.PublicKey;
  lookupTable: anchor.web3.PublicKey;
  writableIndexes: any;
  mint: anchor.web3.PublicKey;
}> {
  const account =
    await program.account.tokenAdminRegistry.fetch(tokenAdminRegistry);
  console.log(account);
  return {
    version: account.version,
    administrator: account.administrator,
    pendingAdministrator: account.pendingAdministrator,
    lookupTable: account.lookupTable,
    writableIndexes: account.writableIndexes,
    mint: account.mint,
  };
}
