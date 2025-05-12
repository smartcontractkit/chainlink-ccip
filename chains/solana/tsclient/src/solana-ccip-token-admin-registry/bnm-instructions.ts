import fs from "fs";
import { Program, Idl, AnchorProvider, Wallet, Provider } from '@coral-xyz/anchor'
import { readFileSync } from "fs";
import accounts, { tokenAdminRegistry } from "../staging";
import { Buffer } from "buffer";
import { BurnmintTokenPool } from '../../../contracts/target/types/burnmint_token_pool';
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";

import { Connection, PublicKey, Keypair, Commitment } from "@solana/web3.js";
import { BPF_LOADER_UPGRADEABLE_PROGRAM_ID } from "./program-instructions";

export const CCIP_TOKENPOOL_CONFIG_SEED = Buffer.from("ccip_tokenpool_config");

export async function loadTokenPoolProgram(): Promise<TokenPoolProgramContext> {

    const keypair = Keypair.fromSecretKey(
        new Uint8Array(JSON.parse(readFileSync(tokenAdminRegistry.key_pair_path, "utf-8")))
    );

    const config = getCCIPSendConfig("devnet");
    const connection = config.connection;

    const wallet = new Wallet(keypair)
    const provider = new AnchorProvider(connection, wallet, {
        preflightCommitment: 'confirmed' as Commitment,
        commitment: 'confirmed' as Commitment,
    })

    const tokenPoolProgramId = tokenAdminRegistry.token_pool_program;

    const bnMTokenPoolIdlFile = fs.readFileSync(tokenAdminRegistry.burnmint_token_pool_idl_path);
    const bnMTokenPoolIdl: Idl = JSON.parse(bnMTokenPoolIdlFile.toString());
    const bnMTokenPoolProgram = new Program<BurnmintTokenPool>(bnMTokenPoolIdl as BurnmintTokenPool, tokenPoolProgramId, provider);

    const [statePda] = PublicKey.findProgramAddressSync(
        [CCIP_TOKENPOOL_CONFIG_SEED, tokenAdminRegistry.mint.toBuffer()],
        tokenAdminRegistry.token_pool_program
    );

    return {
        idl: bnMTokenPoolIdl,
        program: bnMTokenPoolProgram,
        programId: tokenAdminRegistry.token_pool_program,
        connection,
        provider,
        keypair,
        tokenPoolProgramId: tokenPoolProgramId,
        mint: tokenAdminRegistry.mint,
        statePda,
        router: accounts.addresses.router,
        rmnRemote: accounts.addresses.rmnRemote,
        feeQuoter: accounts.addresses.feeQuoter,
    }
}

export type TokenPoolProgramContext = {
    idl: Idl,
    program: Program<BurnmintTokenPool>,
    programId: PublicKey;
    connection: Connection;
    provider: Provider;
    keypair: Keypair;
    tokenPoolProgramId: PublicKey;
    mint: PublicKey;
    statePda: PublicKey;
    router: PublicKey;
    rmnRemote: PublicKey;
    feeQuoter: PublicKey;
};

export function getProgramDataAddress(programId: PublicKey): PublicKey {
    const [programData] = PublicKey.findProgramAddressSync(
        [programId.toBuffer()],
        BPF_LOADER_UPGRADEABLE_PROGRAM_ID
    );
    return programData;
}
