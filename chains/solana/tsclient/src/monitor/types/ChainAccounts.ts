import { NativeAccount } from './NativeAccount';
import { TokenAccount } from './TokenAccount';

export type ChainAccounts = {
    [chain: string]: {
      native: NativeAccount[];
      tokens: TokenAccount[];
    };
  };