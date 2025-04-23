import {
    Connection,
    Keypair,
    PublicKey,
    TransactionMessage,
    VersionedTransaction,
    AddressLookupTableAccount,
  } from "@solana/web3.js";
  import fs from "fs";
  import { ownerProposeAdministrator } from "./../solana-ccip-send/bindings/instructions/ownerProposeAdministrator";
import { getCCIPSendConfig } from "../solana-ccip-send/SolanaCCIPSendConfig";
import { SolanaCCIPPDAs } from "../solana-ccip-send/SolanaCCIPPDAs";
import { tokenAdminRegistry } from "../staging";

  export class SolanaTokenAdmin {
    readonly connection: Connection;
    readonly keypair: Keypair;

    constructor(readonly config = getCCIPSendConfig("devnet")) {
      const keypairPath = tokenAdminRegistry.key_pair_path;
      const secret = Uint8Array.from(
        JSON.parse(fs.readFileSync(keypairPath, "utf-8")),
      );
      this.keypair = Keypair.fromSecretKey(secret);
      this.connection = config.connection;
    }

    async proposeAdmin(
      mint: PublicKey,
      newAdmin: PublicKey,
      lookupTables: AddressLookupTableAccount[] = [],
    ) {
      const [configPDA] = SolanaCCIPPDAs.findConfigPDA(
        this.config.ccipRouterProgramId,
      );

      const tokenAdminRegistry =
        await this.getTokenAdminRegistryFromMint(mint);

      const instruction = ownerProposeAdministrator(
        { tokenAdminRegistryAdmin: newAdmin },
        {
          config: configPDA,
          tokenAdminRegistry: tokenAdminRegistry.address,
          mint,
          authority: this.keypair.publicKey,
          systemProgram: this.config.systemProgramId,
        },
      );

      instruction.programId = this.config.ccipRouterProgramId;

      const { blockhash } = await this.connection.getLatestBlockhash();

      const messageV0 = new TransactionMessage({
        payerKey: this.keypair.publicKey,
        recentBlockhash: blockhash,
        instructions: [instruction],
      }).compileToV0Message(lookupTables);

      const tx = new VersionedTransaction(messageV0);
      tx.sign([this.keypair]);

      console.log("📤 Sending transaction...");
      const sig = await this.connection.sendTransaction(tx, {
        skipPreflight: false,
        maxRetries: 3,
      });
      console.log("✅ Sent. Signature:", sig);
    }

    private async getTokenAdminRegistryFromMint(mint: PublicKey): Promise<{
      address: PublicKey;
      // Add other fields if needed
    }> {
      const [tokenAdminRegistry] =
        SolanaCCIPPDAs.findTokenAdminRegistryPDA(mint, this.config.ccipRouterProgramId);
      return {
        address: tokenAdminRegistry,
      };
    }
  }
