import { TransactionInstruction, PublicKey, AccountMeta } from "@solana/web3.js" // eslint-disable-line @typescript-eslint/no-unused-vars
import BN from "bn.js" // eslint-disable-line @typescript-eslint/no-unused-vars
import * as borsh from "@coral-xyz/borsh" // eslint-disable-line @typescript-eslint/no-unused-vars
import * as types from "../types" // eslint-disable-line @typescript-eslint/no-unused-vars
import { PROGRAM_ID } from "../programId"

export interface OwnerOverridePendingAdministratorArgs {
  tokenAdminRegistryAdmin: PublicKey
}

export interface OwnerOverridePendingAdministratorAccounts {
  config: PublicKey
  tokenAdminRegistry: PublicKey
  mint: PublicKey
  authority: PublicKey
  systemProgram: PublicKey
}

export const layout = borsh.struct([borsh.publicKey("tokenAdminRegistryAdmin")])

/**
 * Overrides the pending admin of the Token Admin Registry by the token owner
 *
 * # Arguments
 *
 * * `ctx` - The context containing the accounts required for registration.
 * * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
 */
export function ownerOverridePendingAdministrator(
  args: OwnerOverridePendingAdministratorArgs,
  accounts: OwnerOverridePendingAdministratorAccounts,
  programId: PublicKey = PROGRAM_ID
) {
  const keys: Array<AccountMeta> = [
    { pubkey: accounts.config, isSigner: false, isWritable: false },
    { pubkey: accounts.tokenAdminRegistry, isSigner: false, isWritable: true },
    { pubkey: accounts.mint, isSigner: false, isWritable: false },
    { pubkey: accounts.authority, isSigner: true, isWritable: true },
    { pubkey: accounts.systemProgram, isSigner: false, isWritable: false },
  ]
  const identifier = Buffer.from([230, 111, 134, 149, 203, 168, 118, 201])
  const buffer = Buffer.alloc(1000)
  const len = layout.encode(
    {
      tokenAdminRegistryAdmin: args.tokenAdminRegistryAdmin,
    },
    buffer
  )
  const data = Buffer.concat([identifier, buffer]).slice(0, 8 + len)
  const ix = new TransactionInstruction({ keys, programId, data })
  return ix
}
