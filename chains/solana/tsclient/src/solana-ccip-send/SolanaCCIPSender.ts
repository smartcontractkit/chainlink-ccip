import {
  AccountMeta,
  AddressLookupTableAccount,
  Connection,
  Keypair,
  PublicKey,
  TransactionMessage,
  VersionedTransaction,
} from "@solana/web3.js";
import {
  getAssociatedTokenAddress,
  TOKEN_PROGRAM_ID,
  NATIVE_MINT,
  TOKEN_2022_PROGRAM_ID,
} from "@solana/spl-token";

import { ccipSend } from "./bindings/instructions/ccipSend";
import { SolanaCCIPPDAs } from "./SolanaCCIPPDAs";
import {
  CcipSendAccounts,
  CcipSendArgs,
} from "./bindings/instructions/ccipSend";
import { SolanaCCIPSendConfig } from "./SolanaCCIPSendConfig";
import { SolanaCCIPSendRequest } from "./SolanaCCIPSendRequest";

export class SolanaCCIPSender {
  private readonly config: SolanaCCIPSendConfig;

  constructor(config: SolanaCCIPSendConfig) {
    this.config = config;
  }

  async send(user: Keypair, msg: SolanaCCIPSendRequest): Promise<void> {
    const connection = this.config.connection;

    if (msg.feeToken.equals(PublicKey.default)) {
      msg.feeToken = NATIVE_MINT;
    }

    const selectorBigInt = BigInt(msg.destChainSelector.toString());

    const [config] = SolanaCCIPPDAs.findConfigPDA(
      this.config.ccipRouterProgramId,
    );
    const [destChainState] = SolanaCCIPPDAs.findDestChainStatePDA(
      selectorBigInt,
      this.config.ccipRouterProgramId,
    );
    const [nonce] = SolanaCCIPPDAs.findNoncePDA(
      selectorBigInt,
      user.publicKey,
      this.config.ccipRouterProgramId,
    );
    const [feeBillingSigner] = SolanaCCIPPDAs.findFeeBillingSignerPDA(
      this.config.ccipRouterProgramId,
    );
    const [feeQuoterConfig] = SolanaCCIPPDAs.findFqConfigPDA(
      this.config.feeQuoterProgramId,
    );
    const [fqDestChain] = SolanaCCIPPDAs.findFqDestChainPDA(
      selectorBigInt,
      this.config.feeQuoterProgramId,
    );
    const [fqBillingTokenConfig] = SolanaCCIPPDAs.findFqBillingTokenConfigPDA(
      msg.feeToken,
      this.config.feeQuoterProgramId,
    );
    const [fqLinkBillingTokenConfig] =
      SolanaCCIPPDAs.findFqBillingTokenConfigPDA(
        this.config.linkTokenMint,
        this.config.feeQuoterProgramId,
      );
    const [rmnRemoteCurses] = SolanaCCIPPDAs.findRMNRemoteCursesPDA(
      this.config.rmnRemoteProgramId,
    );
    const [rmnRemoteConfig] = SolanaCCIPPDAs.findRMNRemoteConfigPDA(
      this.config.rmnRemoteProgramId,
    );
    const [externalTokenPoolsSigner] =
      SolanaCCIPPDAs.findExternalTokenPoolsSignerPDA(
        this.config.ccipRouterProgramId,
      );

    const userATA = await getAssociatedTokenAddress(
      NATIVE_MINT,
      user.publicKey,
      true,
      TOKEN_PROGRAM_ID,
    );

    // TODO only when using native fee. Update for other fee tokens
    const feeTokenReceiverATA = await getAssociatedTokenAddress(
      NATIVE_MINT,
      feeBillingSigner,
      true,
      TOKEN_PROGRAM_ID,
    );

    const accounts: CcipSendAccounts = {
      config,
      destChainState,
      nonce,
      authority: user.publicKey,
      systemProgram: this.config.systemProgramId,
      feeTokenProgram: TOKEN_PROGRAM_ID,
      feeTokenMint: msg.feeToken,
      feeTokenUserAssociatedAccount: userATA,
      feeTokenReceiver: feeTokenReceiverATA,
      feeBillingSigner,
      feeQuoter: this.config.feeQuoterProgramId,
      feeQuoterConfig: feeQuoterConfig,
      feeQuoterDestChain: fqDestChain,
      feeQuoterBillingTokenConfig: fqBillingTokenConfig,
      feeQuoterLinkTokenConfig: fqLinkBillingTokenConfig,
      rmnRemote: this.config.rmnRemoteProgramId,
      rmnRemoteCurses,
      rmnRemoteConfig,
      tokenPoolsSigner: externalTokenPoolsSigner, // TODO: Remove this
    };

    let tokenIndexes: number[] = [];
    let remainingAccounts: Array<AccountMeta> = [];

    let lookupTableList: Array<AddressLookupTableAccount> = [];

    let lastIndex = 0; // TODO: this should be the amount of accounts for the arbitrary message execution

    for (const tm of msg.tokenAmounts) {
      const tokenMint = tm.token;
      const tokenProgram = tm.tokenProgram || TOKEN_2022_PROGRAM_ID;

      const userTokenAccount = await getAssociatedTokenAddress(
        tokenMint,
        user.publicKey,
        true,
        tokenProgram,
      );

      const tokenAdminRegistry = getTokenAdminRegistry();
      const lookupTable = await getLookupTableAccount(
        connection,
        tokenAdminRegistry.lookupTable,
      );
      lookupTableList.push(lookupTable);
      const lookupTableAccounts = lookupTable.state.addresses;
      const poolProgram = getPoolProgram(lookupTableAccounts);

      const [tokenBillingConfig] =
        SolanaCCIPPDAs.findFqPerChainPerTokenConfigPDA(
          BigInt(msg.destChainSelector.toString()),
          tokenMint,
          this.config.feeQuoterProgramId,
        );

      const [poolChainConfig] = SolanaCCIPPDAs.findTokenPoolChainConfigPDA(
        BigInt(msg.destChainSelector.toString()),
        tokenMint,
        poolProgram,
      );

      let tokenAccounts: Array<AccountMeta> = toTokenAccounts(
        userTokenAccount,
        tokenBillingConfig,
        poolChainConfig,
        lookupTableAccounts,
        tokenAdminRegistry.writableIndexes,
      );

      tokenIndexes.push(lastIndex);
      let currentLen = 3 + tokenAccounts.length; // the first 3 are userTokenAccount, tokenBillingConfig, poolChainConfig (not part of the lookuptable)

      lastIndex += currentLen;

      remainingAccounts = remainingAccounts.concat(tokenAccounts);
    }

    const args: CcipSendArgs = {
      destChainSelector: msg.destChainSelector,
      message: {
        receiver: msg.receiver,
        data: msg.data,
        tokenAmounts: msg.tokenAmounts,
        feeToken: msg.feeToken,
        extraArgs: msg.extraArgs,
      },
      tokenIndexes: new Uint8Array(tokenIndexes),
    };

    const instruction = ccipSend(args, accounts);
    instruction.keys.push(...remainingAccounts);

    instruction.programId = this.config.ccipRouterProgramId;

    const instructions = [instruction];
    const { blockhash } = await connection.getLatestBlockhash();

    const messageV0 = new TransactionMessage({
      payerKey: user.publicKey,
      recentBlockhash: blockhash,
      instructions,
    }).compileToV0Message(lookupTableList);

    const tx = new VersionedTransaction(messageV0);

    tx.sign([user]);

    console.log("🔍 Simulating transaction...");
    const simulation = await connection.simulateTransaction(tx);
    if (simulation.value.err) {
      console.error("❌ Simulation failed:", simulation.value.logs);
      throw new Error("Simulation failed");
    }
    console.log("✅ Simulation passed");

    const signature = await connection.sendTransaction(tx, {
      skipPreflight: false,
      maxRetries: 3,
    });
    console.log("📤 Transaction sent:", signature);
  }
}

export type TokenAdminRegistry = {
  version: number;
  administrator: PublicKey;
  pendingAdministrator: PublicKey;
  lookupTable: PublicKey;
  writableIndexes: [bigint, bigint];
  mint: PublicKey;
};

function toTokenAccounts(
  userTokenAccount: PublicKey,
  tokenBillingConfig: PublicKey,
  poolChainConfig: PublicKey,
  lookupTableEntries: Array<PublicKey>,
  writableIndexes: [bigint, bigint],
): Array<AccountMeta> {
  return [
    { pubkey: userTokenAccount, isSigner: false, isWritable: true },
    { pubkey: tokenBillingConfig, isSigner: false, isWritable: false },
    { pubkey: poolChainConfig, isSigner: false, isWritable: true },

    // List of accounts from the lookup table
    {
      pubkey: lookupTableEntries[0],
      isSigner: false,
      isWritable: isWritable(0, writableIndexes),
    }, // poolProgram
    {
      pubkey: lookupTableEntries[1],
      isSigner: false,
      isWritable: isWritable(1, writableIndexes),
    }, // poolConfig
    {
      pubkey: lookupTableEntries[2],
      isSigner: false,
      isWritable: isWritable(2, writableIndexes),
    }, // poolTokenAccount
    {
      pubkey: lookupTableEntries[3],
      isSigner: false,
      isWritable: isWritable(3, writableIndexes),
    }, // poolSigner
    {
      pubkey: lookupTableEntries[4],
      isSigner: false,
      isWritable: isWritable(4, writableIndexes),
    }, // tokenProgram
    {
      pubkey: lookupTableEntries[5],
      isSigner: false,
      isWritable: isWritable(5, writableIndexes),
    }, // mint
    {
      pubkey: lookupTableEntries[6],
      isSigner: false,
      isWritable: isWritable(6, writableIndexes),
    }, // feeTokenConfig

    // pub ccip_router_pool_signer: &'a AccountInfo<'a>,
    // pub ccip_router_pool_signer_bump: u8,
    // pub ccip_offramp_pool_signer_bump: u8,

    //remainingAccounts: [], // TODO: Check if there are
  ];
}

function isWritable(index: number, writableIndexes: [bigint, bigint]): boolean {
  return false;
  // const indexBigInt = BigInt(index);
  // return (
  //   (indexBigInt >= writableIndexes[0] && indexBigInt < writableIndexes[1]) ||
  //   (indexBigInt >= writableIndexes[1] && indexBigInt < writableIndexes[0])
  // );
}

function getTokenAdminRegistry(
  // tokenAdminRegistryAddress: PublicKey,
): TokenAdminRegistry {
  return {
    version: 1,
    administrator: PublicKey.default,
    pendingAdministrator: PublicKey.default,
    lookupTable: new PublicKey(""),
    writableIndexes: [BigInt(0), BigInt(0)],
    mint: PublicKey.default,
  };
}

async function getLookupTableAccount(
  connection: Connection,
  lookupTableAddress: PublicKey,
): Promise<AddressLookupTableAccount> {
  const { value: lookupTableAccount } =
    await connection.getAddressLookupTable(lookupTableAddress);
  if (!lookupTableAccount) {
    throw new Error("Lookup table not found");
  }
  if (lookupTableAccount.state.addresses.length < 7) {
    throw new Error("Lookup table does not have enough accounts");
  }

  return lookupTableAccount;
}

function getPoolProgram(lookupTableAccounts: Array<PublicKey>): PublicKey {
  // TODO: lookupTableAccount.state.addresses[0] is the pool program
  return lookupTableAccounts[0];
}
