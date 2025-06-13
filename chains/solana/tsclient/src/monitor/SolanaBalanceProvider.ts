import { ChainInfo } from '../common/ChainRegistry';
import { IBalanceProvider } from './IBalanceProvider';
import { BalanceInfo } from './types/BalanceInfo';
import { TokenAccount } from './types/TokenAccount';
import { NativeAccount } from './types/NativeAccount';
import { Connection, PublicKey } from '@solana/web3.js';

export class SolanaBalanceProvider implements IBalanceProvider {
  getNativeBalances(chainInfo: ChainInfo, accounts: NativeAccount[]): Promise<BalanceInfo[]> {
    const connection = new Connection(chainInfo.rpcURL);

    return Promise.all(
      accounts.map(async ({ owner, name }) => {
        const lamports = await connection.getBalance(new PublicKey(owner));
        const sol = lamports / 1e9;

        return {
          chain: chainInfo.name,
          name,
          owner,
          balance: `${sol} SOL`,
        };
      })
    );
  }

  getTokenBalances(chainInfo: ChainInfo, accounts: TokenAccount[]): Promise<BalanceInfo[]> {
    const connection = new Connection(chainInfo.rpcURL);

    return Promise.all(
      accounts.map(async ({ mint, owner, name }) => {
        try {
          const mintPubkey = new PublicKey(mint);
          const ownerPubkey = new PublicKey(owner);
          const tokenAccount = await connection.getParsedTokenAccountsByOwner(ownerPubkey, {
            mint: mintPubkey,
          });

          if (tokenAccount.value.length > 0) {
            const balance = tokenAccount.value[0].account.data.parsed.info.tokenAmount.uiAmount;
            return {
              chain: chainInfo.name,
              name,
              owner,
              balance: `${balance} tokens`,
            };
          } else {
            return {
              chain: chainInfo.name,
              name,
              owner,
              balance: '0 tokens',
            };
          }
        } catch (err) {
          console.warn(`[Solana] Failed to fetch token balance for ${name} (${owner} @ ${mint}): ${err}`);
          return {
            chain: chainInfo.name,
            name,
            owner,
            balance: '0 tokens',
          };
        }
      })
    );
  }
}
