import { ChainInfo } from '../common/ChainRegistry';
import { BalanceInfo } from './types/BalanceInfo';
import { TokenAccount } from './types/TokenAccount';
import { NativeAccount } from './types/NativeAccount';

export interface IBalanceProvider {
  getNativeBalances(chainInfo: ChainInfo, accounts: NativeAccount[]): Promise<BalanceInfo[]>;
  getTokenBalances(chainInfo: ChainInfo, accounts: TokenAccount[]): Promise<BalanceInfo[]>;
}
