import { BalanceInfo } from './types/BalanceInfo';
import { ChainAccounts } from './types/ChainAccounts';
import { ChainRegistry } from '../common/ChainRegistry';
import { SolanaBalanceProvider } from './SolanaBalanceProvider';
import { EvmBalanceProvider } from './EvmBalanceProvider';
import { IBalanceProvider } from './IBalanceProvider';

export class BalanceFetcher {
  private solanaProvider: IBalanceProvider;
  private evmProvider: IBalanceProvider;

  constructor() {
    this.solanaProvider = new SolanaBalanceProvider();
    this.evmProvider = new EvmBalanceProvider();
  }

  async getBalances(accounts: ChainAccounts): Promise<BalanceInfo[]> {
    const results: BalanceInfo[] = [];

    for (const [chainName, { native, tokens }] of Object.entries(accounts)) {
      let chainInfo;
      try {
        chainInfo = ChainRegistry.getChainInfo(chainName);
      } catch {
        console.warn(`[BalanceFetcher] Unknown chain "${chainName}". Skipping.`);
        continue;
      }

      const provider = this.getProvider(chainInfo.family);
      const nativeBalances = await provider.getNativeBalances(chainInfo, native);
      const tokenBalances = await provider.getTokenBalances(chainInfo, tokens);
      results.push(...nativeBalances, ...tokenBalances);
    }

    return results;
  }

  private getProvider(family: string): IBalanceProvider {
    switch (family) {
      case 'Solana':
        return this.solanaProvider;
      case 'EVM':
        return this.evmProvider;
      default:
        throw new Error(`Unsupported chain family: ${family}`);
    }
  }
}
