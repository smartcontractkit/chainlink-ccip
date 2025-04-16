import {
  AccountMeta,
  AddressLookupTableAccount,
  Connection,
  Keypair,
  PublicKey,
  sendAndConfirmTransaction,
  Transaction,
  TransactionMessage,
  VersionedTransaction,
} from "@solana/web3.js";
import {
  getAssociatedTokenAddress,
  TOKEN_PROGRAM_ID,
  NATIVE_MINT,
  TOKEN_2022_PROGRAM_ID,
  createAssociatedTokenAccountInstruction,
  ASSOCIATED_TOKEN_PROGRAM_ID,
} from "@solana/spl-token";

import { ccipSend } from "./bindings/instructions/ccipSend";
import { SolanaCCIPPDAs } from "./SolanaCCIPPDAs";
import {
  CcipSendAccounts,
  CcipSendArgs,
} from "./bindings/instructions/ccipSend";
import { SolanaCCIPSendConfig } from "./SolanaCCIPSendConfig";
import { SolanaCCIPSendRequest } from "./SolanaCCIPSendRequest";
import { CcipCommonClient } from "./SolanaCCIPReadAccounts";
import BN from "bn.js";

export class SolanaCCIPSender {
  private readonly config: SolanaCCIPSendConfig;
  private readonly ccipCommonClient: CcipCommonClient;

  constructor(config: SolanaCCIPSendConfig) {
    this.config = config;
    this.ccipCommonClient = new CcipCommonClient(
      "./ccip-keypair.json",
      "devnet",
    );
  }

  async send(user: Keypair, msg: SolanaCCIPSendRequest): Promise<void> {
    const connection = this.config.connection;

    if (msg.feeToken.equals(PublicKey.default)) {
      msg.feeToken = NATIVE_MINT;
    }

    const selectorBigInt = BigInt(msg.destChainSelector.toString());

    const ccipSendAccounts: CcipSendAccounts = await buildCCipSendAccounts(
      this.config,
      selectorBigInt,
      user,
      msg,
    );

    let tokenIndexes: number[] = [];
    let remainingAccounts: Array<AccountMeta> = [];

    let lookupTableList: Array<AddressLookupTableAccount> = [];

    let lastIndex = 0; // TODO: this should be the amount of accounts for the arbitrary message execution

    for (const tm of msg.tokenAmounts) {
      const tokenMint = tm.token;
      const tokenProgram = tm.tokenProgram || TOKEN_2022_PROGRAM_ID;

      let tokenAccounts: Array<AccountMeta> = await buildTokenAccounts(
        this.config,
        this.ccipCommonClient,
        tokenMint,
        user,
        new PublicKey(tokenProgram),
        connection,
        lookupTableList,
        msg,
      );

      tokenIndexes.push(lastIndex);
      let currentLen = tokenAccounts.length; // the first 3 are userTokenAccount, tokenBillingConfig, poolChainConfig (not part of the lookuptable)

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

    const instruction = ccipSend(args, ccipSendAccounts);

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

async function buildCCipSendAccounts(
  config: SolanaCCIPSendConfig,
  selectorBigInt: bigint,
  user: Keypair,
  msg: SolanaCCIPSendRequest,
) {
  const [configPDA] = SolanaCCIPPDAs.findConfigPDA(config.ccipRouterProgramId);
  const [destChainState] = SolanaCCIPPDAs.findDestChainStatePDA(
    selectorBigInt,
    config.ccipRouterProgramId,
  );
  const [nonce] = SolanaCCIPPDAs.findNoncePDA(
    selectorBigInt,
    user.publicKey,
    config.ccipRouterProgramId,
  );
  const [feeBillingSigner] = SolanaCCIPPDAs.findFeeBillingSignerPDA(
    config.ccipRouterProgramId,
  );
  const [feeQuoterConfig] = SolanaCCIPPDAs.findFqConfigPDA(
    config.feeQuoterProgramId,
  );
  const [fqDestChain] = SolanaCCIPPDAs.findFqDestChainPDA(
    selectorBigInt,
    config.feeQuoterProgramId,
  );
  const [fqBillingTokenConfig] = SolanaCCIPPDAs.findFqBillingTokenConfigPDA(
    msg.feeToken,
    config.feeQuoterProgramId,
  );
  const [fqLinkBillingTokenConfig] = SolanaCCIPPDAs.findFqBillingTokenConfigPDA(
    config.linkTokenMint,
    config.feeQuoterProgramId,
  );
  const [rmnRemoteCurses] = SolanaCCIPPDAs.findRMNRemoteCursesPDA(
    config.rmnRemoteProgramId,
  );
  const [rmnRemoteConfig] = SolanaCCIPPDAs.findRMNRemoteConfigPDA(
    config.rmnRemoteProgramId,
  );
  const [externalTokenPoolsSigner] =
    SolanaCCIPPDAs.findExternalTokenPoolsSignerPDA(config.ccipRouterProgramId);

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
    config: configPDA,
    destChainState,
    nonce,
    authority: user.publicKey,
    systemProgram: config.systemProgramId,
    feeTokenProgram: TOKEN_PROGRAM_ID,
    feeTokenMint: msg.feeToken,
    feeTokenUserAssociatedAccount: userATA,
    feeTokenReceiver: feeTokenReceiverATA,
    feeBillingSigner,
    feeQuoter: config.feeQuoterProgramId,
    feeQuoterConfig: feeQuoterConfig,
    feeQuoterDestChain: fqDestChain,
    feeQuoterBillingTokenConfig: fqBillingTokenConfig,
    feeQuoterLinkTokenConfig: fqLinkBillingTokenConfig,
    rmnRemote: config.rmnRemoteProgramId,
    rmnRemoteCurses,
    rmnRemoteConfig,
    tokenPoolsSigner: externalTokenPoolsSigner, // TODO: Remove this
  };
  return accounts;
}

async function buildTokenAccounts(
  config: SolanaCCIPSendConfig,
  ccipCommonClient: CcipCommonClient,
  tokenMint: PublicKey,
  user: Keypair,
  tokenProgram: PublicKey,
  connection: Connection,
  lookupTableList: AddressLookupTableAccount[],
  msg: SolanaCCIPSendRequest,
) {
  const userTokenAccount = await getAssociatedTokenAddress(
    tokenMint,
    user.publicKey,
    true,
    tokenProgram,
    ASSOCIATED_TOKEN_PROGRAM_ID,
  );
  // init token account if needed
  const accountInfo = await connection.getAccountInfo(userTokenAccount);
  if (!accountInfo) {
    const ataIx = createAssociatedTokenAccountInstruction(
      user.publicKey,        // funding address (usually the payer)
      userTokenAccount,       // ATA to create (calculated before)
      user.publicKey,         // owner of the token account
      tokenMint,
      tokenProgram,
      ASSOCIATED_TOKEN_PROGRAM_ID,
    );
    const tx = new Transaction().add(ataIx);
    const sig = await sendAndConfirmTransaction(connection, tx, [user], {
      skipPreflight: false,
      commitment: "confirmed",
    });

    // Confirm it explicitly at the highest commitment level
    await connection.confirmTransaction(sig, "finalized");
  }

  const tokenAdminRegistry =
    await ccipCommonClient.getTokenAdminRegistry(tokenMint);
  const lookupTable = await getLookupTableAccount(
    connection,
    tokenAdminRegistry.lookupTable,
  );
  lookupTableList.push(lookupTable);
  const lookupTableAccounts = lookupTable.state.addresses;
  const poolProgram = getPoolProgram(lookupTableAccounts);

  const [tokenBillingConfig] = SolanaCCIPPDAs.findFqPerChainPerTokenConfigPDA(
    BigInt(msg.destChainSelector.toString()),
    tokenMint,
    config.feeQuoterProgramId,
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

  return tokenAccounts;
}

function toTokenAccounts(
  userTokenAccount: PublicKey,
  tokenBillingConfig: PublicKey,
  poolChainConfig: PublicKey,
  lookupTableEntries: Array<PublicKey>,
  writableIndexes: BN[],
): Array<AccountMeta> {
  return [
    { pubkey: userTokenAccount, isSigner: false, isWritable: true },
    { pubkey: tokenBillingConfig, isSigner: false, isWritable: false },
    { pubkey: poolChainConfig, isSigner: false, isWritable: true },

    // List of accounts from the lookup table
    {
      pubkey: lookupTableEntries[0], // poolProgram
      isSigner: false,
      isWritable: isWritable(0, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[1], // poolConfig
      isSigner: false,
      isWritable: isWritable(1, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[2], // poolTokenAccount
      isSigner: false,
      isWritable: isWritable(2, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[3], // poolSigner
      isSigner: false,
      isWritable: isWritable(3, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[4], // tokenProgram
      isSigner: false,
      isWritable: isWritable(4, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[5], // mint
      isSigner: false,
      isWritable: isWritable(5, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[6], // feeTokenConfig
      isSigner: false,
      isWritable: isWritable(6, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[7], // ccip_router_pool_signer
      isSigner: false,
      isWritable: isWritable(7, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[8], // ccip_router_pool_signer_bump
      isSigner: false,
      isWritable: isWritable(8, writableIndexes),
    },
    {
      pubkey: lookupTableEntries[9], // ccip_offramp_pool_signer_bump
      isSigner: false,
      isWritable: isWritable(9, writableIndexes),
    },
    //remainingAccounts: [], // TODO: Check if there are
  ];
}

function isWritable(index: number, writableIndexes: BN[]): boolean {
  return false;
  // const indexBigInt = BigInt(index);
  // return (
  //   (indexBigInt >= writableIndexes[0] && indexBigInt < writableIndexes[1]) ||
  //   (indexBigInt >= writableIndexes[1] && indexBigInt < writableIndexes[0])
  // );
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
  return lookupTableAccounts[0];
}
