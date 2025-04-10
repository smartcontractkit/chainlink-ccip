import { PublicKey } from '@solana/web3.js';
import BN from 'bn.js';

export type SolanaCCIPSendRequest = {
  destChainSelector: BN;
  receiver: Uint8Array;
  data: Buffer;
  tokenAmounts: Array<{
    token: PublicKey;
    amount: BN;
    tokenProgram?: PublicKey;
  }>;
  feeToken: PublicKey;
  feeTokenProgram?: PublicKey;
  extraArgs: Buffer;
};
