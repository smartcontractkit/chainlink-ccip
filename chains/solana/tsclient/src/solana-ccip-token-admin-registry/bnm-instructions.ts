import { TransactionInstruction, PublicKey } from "@solana/web3.js";
import * as borsh from "@coral-xyz/borsh";
import { Buffer } from "buffer";
import crypto from "crypto";

// --------------------
// Type Definitions
// --------------------

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

export interface RemoteAddress {
  address: Uint8Array; // 32 bytes
}

export interface RemoteConfig {
  poolAddresses: RemoteAddress[];
  tokenAddress: RemoteAddress;
  decimals: number; // uint8
}

export interface InitChainRemoteConfigArgs {
  remoteChainSelector: bigint;
  mint: PublicKey;
  cfg: RemoteConfig;
}

export interface InitChainRemoteConfigAccounts {
  state: PublicKey;
  chainConfig: PublicKey;
  authority: PublicKey;
  systemProgram: PublicKey;
}

// --------------------
// Layouts and Discriminators
// --------------------

// Layout for initialize
const initializeLayout = borsh.struct([
  borsh.publicKey("router"),
  borsh.publicKey("rmnRemote"),
]);

// Layout for setRouter
const setRouterLayout = borsh.struct([
  borsh.publicKey("newRouter"),
]);

// Layout for RemoteAddress
const remoteAddressLayout = borsh.struct([
  borsh.vecU8("address"), // Array of bytes (expected to be 32 bytes)
]);

// Layout for RemoteConfig
const remoteConfigLayout = borsh.struct([
  borsh.vec(remoteAddressLayout, "poolAddresses"), // Array of RemoteAddress structs
  remoteAddressLayout.replicate("tokenAddress"),   // A single RemoteAddress
  borsh.u8("decimals"),                             // A uint8
]);

// Full layout for initChainRemoteConfig
const initChainRemoteConfigLayout = borsh.struct([
  borsh.u64("remoteChainSelector"),
  borsh.publicKey("mint"),
  remoteConfigLayout.replicate("cfg"),
]);

// Discriminators
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

const INIT_CHAIN_REMOTE_CONFIG_DISCRIMINATOR = crypto
  .createHash("sha256")
  .update("global:init_chain_remote_config")
  .digest()
  .subarray(0, 8);

// --------------------
// Instruction Builders
// --------------------

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

  return new TransactionInstruction({ programId, keys, data: encodedArgs });
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

  return new TransactionInstruction({ programId, keys, data: encodedArgs });
}

export function createInitChainRemoteConfigInstruction(
  accounts: InitChainRemoteConfigAccounts,
  args: InitChainRemoteConfigArgs,
  programId: PublicKey
): TransactionInstruction {
  const data = Buffer.alloc(2000);
  const len = initChainRemoteConfigLayout.encode(args, data);
  const encodedArgs = Buffer.concat([
    INIT_CHAIN_REMOTE_CONFIG_DISCRIMINATOR,
    data.subarray(0, len),
  ]);

  const keys = [
    { pubkey: accounts.state, isWritable: false, isSigner: false },
    { pubkey: accounts.chainConfig, isWritable: true, isSigner: false },
    { pubkey: accounts.authority, isWritable: true, isSigner: true },
    { pubkey: accounts.systemProgram, isWritable: false, isSigner: false },
  ];

  return new TransactionInstruction({ programId, keys, data: encodedArgs });
}
