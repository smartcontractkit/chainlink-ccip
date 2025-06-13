import { AnchorProvider, Program, BN } from "@coral-xyz/anchor";
import { Keypair, PublicKey, Connection } from "@solana/web3.js";
import fs from "fs";
import NodeWallet from "@coral-xyz/anchor/dist/cjs/nodewallet";

import { CcipCommon } from "../../../contracts/target/types/ccip_common";
import ccipCommonIdl from "../../../contracts/target/idl/ccip_common.json";
import accounts from "../staging";
import { getCCIPSendConfig } from "./SolanaCCIPSendConfig";

export type TokenAdminRegistry = {
  version: number;
  administrator: PublicKey;
  pendingAdministrator: PublicKey;
  lookupTable: PublicKey;
  writableIndexes: BN[];
  mint: PublicKey;
};

export class CcipCommonClient {
  readonly connection: Connection;
  readonly provider: AnchorProvider;
  readonly wallet: NodeWallet;
  readonly program: Program<CcipCommon>;

  public constructor(
    readonly keypairPath: string,
    readonly network: "devnet" | "mainnet",
  ) {
    const config = getCCIPSendConfig(network);
    this.connection = config.connection;

    if (!fs.existsSync(keypairPath)) {
      throw new Error(`❌ Missing keypair file at ${keypairPath}`);
    }

    const secret = Uint8Array.from(
      JSON.parse(fs.readFileSync(keypairPath, "utf-8")),
    );
    const keypair = Keypair.fromSecretKey(secret);
    this.wallet = new NodeWallet(keypair);
    this.provider = new AnchorProvider(this.connection, this.wallet, {});

    this.program = new Program<CcipCommon>(
      ccipCommonIdl as CcipCommon,
      accounts.addresses.router,
      this.provider,
    );
  }

  private getTokenAdminRegistryAddress(mint: PublicKey): PublicKey {
    const [pda] = PublicKey.findProgramAddressSync(
      [Buffer.from("token_admin_registry"), mint.toBuffer()],
      this.program.programId,
    );
    return pda;
  }

  public async getTokenAdminRegistry(
    mint: PublicKey,
  ): Promise<TokenAdminRegistry> {
    const pda = this.getTokenAdminRegistryAddress(mint);
    const account = await this.program.account.tokenAdminRegistry.fetch(pda);
    return account;
  }
}
