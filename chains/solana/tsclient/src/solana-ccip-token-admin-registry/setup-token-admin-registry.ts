import { TransactionInstruction, PublicKey } from "@solana/web3.js";
import * as borsh from "@coral-xyz/borsh";
import { Buffer } from "buffer";
import crypto from "crypto";

export interface InitializeArgs {
  router: PublicKey;
  rmnRemote: PublicKey;
}

export interface InitializeAccounts {
  state: PublicKey;
  mint: PublicKey;
  authority: PublicKey;
  systemProgram: PublicKey;
  program: PublicKey;
  programData: PublicKey;
}

export interface SetRouterArgs {
  newRouter: PublicKey;
}

export interface SetRouterAccounts {
  state: PublicKey;
  mint: PublicKey;
  authority: PublicKey;
}

export interface OwnerProposeAdministratorArgs {
  tokenAdminRegistryAdmin: PublicKey;
}

export interface OwnerProposeAdministratorAccounts {
  config: PublicKey;
  tokenAdminRegistry: PublicKey;
  mint: PublicKey;
  authority: PublicKey;
  systemProgram: PublicKey;
}

const ownerProposeAdministratorLayout = borsh.struct([
  borsh.publicKey("tokenAdminRegistryAdmin"),
]);

const OWNER_PROPOSE_ADMINISTRATOR_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:owner_propose_administrator")
  .digest()
  .subarray(0, 8);

export function createOwnerProposeAdministratorInstruction(
  accounts: OwnerProposeAdministratorAccounts,
  args: OwnerProposeAdministratorArgs,
  programId: PublicKey
): TransactionInstruction {
  const data = Buffer.alloc(1000);
  const len = ownerProposeAdministratorLayout.encode(args, data);
  const encodedArgs = Buffer.concat([
    OWNER_PROPOSE_ADMINISTRATOR_DISCRIMINATOR,
    data.subarray(0, len),
  ]);

  const keys = [
    { pubkey: accounts.config, isWritable: false, isSigner: false },
    { pubkey: accounts.tokenAdminRegistry, isWritable: true, isSigner: false },
    { pubkey: accounts.mint, isWritable: false, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
    { pubkey: accounts.systemProgram, isWritable: false, isSigner: false },
  ];

  return new TransactionInstruction({
    programId,
    keys,
    data: encodedArgs,
  });
}

export interface AcceptAdminRoleAccounts {
  config: PublicKey;
  tokenAdminRegistry: PublicKey;
  mint: PublicKey;
  authority: PublicKey;
}

const ACCEPT_ADMIN_ROLE_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:accept_admin_role_token_admin_registry")
  .digest()
  .subarray(0, 8);

export function createAcceptAdminRoleInstruction(
  accounts: AcceptAdminRoleAccounts,
  programId: PublicKey
): TransactionInstruction {
  const encodedArgs = Buffer.from(ACCEPT_ADMIN_ROLE_DISCRIMINATOR);

  const keys = [
    { pubkey: accounts.config, isWritable: false, isSigner: false },
    { pubkey: accounts.tokenAdminRegistry, isWritable: true, isSigner: false },
    { pubkey: accounts.mint, isWritable: false, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
  ];

  return new TransactionInstruction({
    programId,
    keys,
    data: encodedArgs,
  });
}

export interface TransferAdminRoleArgs {
  newAdmin: PublicKey;
}

export interface TransferAdminRoleAccounts {
  config: PublicKey;
  tokenAdminRegistry: PublicKey;
  mint: PublicKey;
  authority: PublicKey;
}

const transferAdminRoleLayout = borsh.struct([
  borsh.publicKey("newAdmin"),
]);

const TRANSFER_ADMIN_ROLE_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:transfer_admin_role_token_admin_registry")
  .digest()
  .subarray(0, 8);

export function createTransferAdminRoleInstruction(
  accounts: TransferAdminRoleAccounts,
  args: TransferAdminRoleArgs,
  programId: PublicKey
): TransactionInstruction {
  const data = Buffer.alloc(1000);
  const len = transferAdminRoleLayout.encode(args, data);
  const encodedArgs = Buffer.concat([
    TRANSFER_ADMIN_ROLE_DISCRIMINATOR,
    data.subarray(0, len),
  ]);

  const keys = [
    { pubkey: accounts.config, isWritable: false, isSigner: false },
    { pubkey: accounts.tokenAdminRegistry, isWritable: true, isSigner: false },
    { pubkey: accounts.mint, isWritable: false, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
  ];

  return new TransactionInstruction({
    programId,
    keys,
    data: encodedArgs,
  });
}
