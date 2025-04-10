import * as fs from 'fs';
import { Keypair, PublicKey } from '@solana/web3.js';
import { getCCIPSendConfig } from './solana-ccip-send/SolanaCCIPSendConfig';
import { SolanaCCIPSender } from './solana-ccip-send/SolanaCCIPSender';
import { SolanaCCIPSendRequest } from './solana-ccip-send/SolanaCCIPSendRequest';
import { ChainRegistry } from './common/ChainRegistry';
import { SolanaToEVMUtils } from './solana-ccip-send/SolanaToEVMUtils';
import BN from 'bn.js';
import accounts from './staging';

async function main() {
  const keypairPath = './ccip-keypair.json';
  if (!fs.existsSync(keypairPath)) {
    throw new Error(`❌ Missing keypair file at ${keypairPath}`);
  }

  const secret = Uint8Array.from(JSON.parse(fs.readFileSync(keypairPath, 'utf-8')));
  const keypair = Keypair.fromSecretKey(secret);

  console.log(`🔑 Using keypair with public key: ${keypair.publicKey.toBase58()}`);

  const args = process.argv.slice(2); // Ignore the first args (node & script path)
  const messageArg = args.find(arg => arg.startsWith('--message='));
  const message = messageArg ? messageArg.split('=')[1] : 'Hello from Solana! Agus';

  const tokens = args.includes('--tokens');

  let tokenAmounts: Array<{
      token: PublicKey;
      amount: BN;
      tokenProgram?: PublicKey;
    }> = [];

  if (tokens) {
    const token_mint = new PublicKey(accounts.addresses.tokenMint);
    const token_amount = new BN('1000000000000000000'); // 1 BnM token
    tokenAmounts = [{
      token: token_mint,
      amount: token_amount,
    }];
  }

  const config = getCCIPSendConfig('devnet');
  const ccip = new SolanaCCIPSender(config);

  const request: SolanaCCIPSendRequest = {
    destChainSelector: ChainRegistry.getChainSelector(accounts.remoteChain.chainName),
    receiver: SolanaToEVMUtils.evmAddressToSolanaBytes(accounts.remoteChain.receiver),
    data: Buffer.from(message),
    tokenAmounts,
    feeToken: PublicKey.default,
    extraArgs: Buffer.alloc(0),
  };
  await ccip.send(keypair, request);
}

main().catch(console.error);
