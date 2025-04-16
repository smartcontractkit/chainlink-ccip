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

const initializeLayout = borsh.struct([
  borsh.publicKey("router"),
  borsh.publicKey("rmnRemote"),
]);

const setRouterLayout = borsh.struct([
  borsh.publicKey("newRouter"),
]);

const INITIALIZE_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:initialize")
  .digest()
  .subarray(0, 8);

const SET_ROUTER_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:set_router")
  .digest()
  .subarray(0, 8);

export function createInitializeInstruction(
  accounts: InitializeAccounts,
  args: InitializeArgs,
  programId: PublicKey
): TransactionInstruction {
  const data = Buffer.alloc(1000);
  const len = initializeLayout.encode(args, data);
  const encodedArgs = Buffer.concat([
    INITIALIZE_DISCRIMINATOR,
    data.subarray(0, len),
  ]);

  const keys = [
    { pubkey: accounts.state, isWritable: true, isSigner: false },
    { pubkey: accounts.mint, isWritable: false, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
    { pubkey: accounts.systemProgram, isWritable: false, isSigner: false },
    { pubkey: accounts.program, isWritable: false, isSigner: false },
    { pubkey: accounts.programData, isWritable: false, isSigner: false },
  ];

  return new TransactionInstruction({
    programId,
    keys,
    data: encodedArgs,
  });
}

export function createSetRouterInstruction(
  accounts: SetRouterAccounts,
  args: SetRouterArgs,
  programId: PublicKey
): TransactionInstruction {
  const data = Buffer.alloc(1000);
  const len = setRouterLayout.encode(args, data);
  const encodedArgs = Buffer.concat([
    SET_ROUTER_DISCRIMINATOR,
    data.subarray(0, len),
  ]);

  const keys = [
    { pubkey: accounts.state, isWritable: true, isSigner: false },
    { pubkey: accounts.mint, isWritable: false, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
  ];

  return new TransactionInstruction({
    programId,
    keys,
    data: encodedArgs,
  });
}
