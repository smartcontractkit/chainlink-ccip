import fs from "fs";
import { Program, Idl, AnchorProvider, Wallet, Provider } from '@coral-xyz/anchor'
import { readFileSync } from "fs";
import accounts, { tokenAdminRegistry } from "../staging";
import { CcipRouter } from '../../../contracts/target/types/ccip_router';

import { Connection, PublicKey, Keypair, Commitment } from "@solana/web3.js";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";


export async function loadCcipRouterProgram(): Promise<CcipRouterTokenAdminRegistryProgramContext> {

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
    const routerProgramId = accounts.addresses.router;

    const ccipRouterIdlFile = fs.readFileSync(tokenAdminRegistry.ccip_router_idl_path);
    const ccipRouterIdl: Idl = JSON.parse(ccipRouterIdlFile.toString());
    const ccipRouterProgram = new Program<CcipRouter>(ccipRouterIdl as CcipRouter, routerProgramId, provider);

    const [configPDA] = PublicKey.findProgramAddressSync(
        [Buffer.from("config")],
        routerProgramId
    );

    const [tokenAdminRegistryPDA] = PublicKey.findProgramAddressSync(
        [Buffer.from("token_admin_registry"), tokenAdminRegistry.mint.toBuffer()],
        routerProgramId
        );


    return {
        idl: ccipRouterIdl,
        program: ccipRouterProgram,
        connection,
        provider,
        keypair,
        tokenPoolProgramId: tokenAdminRegistry.token_pool_program,
        mint: tokenAdminRegistry.mint,
        router: accounts.addresses.router,
        rmnRemote: accounts.addresses.rmnRemote,
        feeQuoter: accounts.addresses.feeQuoter,
        configPDA,
        tokenAdminRegistryPDA,
        lookupTable: tokenAdminRegistry.lookup_table,
    }
}

export type CcipRouterTokenAdminRegistryProgramContext = {
    idl: Idl,
    program: Program<CcipRouter>,
    connection: Connection;
    provider: Provider;
    keypair: Keypair;
    tokenPoolProgramId: PublicKey;
    mint: PublicKey;
    router: PublicKey;
    rmnRemote: PublicKey;
    feeQuoter: PublicKey;
    configPDA: PublicKey;
    tokenAdminRegistryPDA: PublicKey;
    lookupTable: PublicKey;
};
