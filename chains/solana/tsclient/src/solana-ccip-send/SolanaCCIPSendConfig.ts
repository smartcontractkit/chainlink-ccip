import { Connection, PublicKey, clusterApiUrl } from '@solana/web3.js';
import accounts from '../staging';

export class SolanaCCIPSendConfig {
  readonly connection: Connection;
  readonly systemProgramId: PublicKey;
  readonly ccipRouterProgramId: PublicKey;
  readonly feeQuoterProgramId: PublicKey;
  readonly rmnRemoteProgramId: PublicKey;
  readonly linkTokenMint: PublicKey;

  constructor(params: {
    connection: Connection;
    systemProgramId: PublicKey;
    ccipRouterProgramId: PublicKey;
    feeQuoterProgramId: PublicKey;
    rmnRemoteProgramId: PublicKey;
    linkTokenMint: PublicKey;
  }) {
    this.connection = params.connection;
    this.systemProgramId = params.systemProgramId;
    this.ccipRouterProgramId = params.ccipRouterProgramId;
    this.feeQuoterProgramId = params.feeQuoterProgramId;
    this.rmnRemoteProgramId = params.rmnRemoteProgramId;
    this.linkTokenMint = params.linkTokenMint;
  }
}

export function getCCIPSendConfig(env: 'devnet' | 'mainnet'): SolanaCCIPSendConfig {
  switch (env) {
    case 'devnet':
      return new SolanaCCIPSendConfig({
        connection: new Connection(clusterApiUrl('devnet'), 'confirmed'),
        systemProgramId: new PublicKey('11111111111111111111111111111111'),
        ccipRouterProgramId: accounts.addresses.router,
        feeQuoterProgramId: accounts.addresses.feeQuoter,
        rmnRemoteProgramId: accounts.addresses.rmnRemote,
        linkTokenMint: accounts.addresses.linkMint,
      });

    case 'mainnet':
      throw new Error(`"mainnet" is a valid environment, but it's not configured yet. Please update CCIPSendConfig with proper mainnet program IDs and mint addresses.`);

    default:
      throw new Error(`Unsupported environment: ${env}`);
  }
}
