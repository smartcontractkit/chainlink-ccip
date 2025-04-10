import BN from 'bn.js';

import { chains } from "../staging"

export type ChainFamily = 'EVM' | 'Solana';

export type ChainInfo = {
  name: string;
  selector: BN;
  family: ChainFamily;
  rpcURL: string;
};

export class ChainRegistry {

  static getChainSelector(name: string): BN {
    const chain = chains.get(name);
    if (!chain) throw new Error(`Unknown chain name: "${name}"`);
    return chain.selector;
  }

  static getChainInfo(name: string): ChainInfo {
    const chain = chains.get(name);
    if (!chain) throw new Error(`Unknown chain name: "${name}"`);
    return chain;
  }

  static getChainsByFamily(family: ChainFamily): ChainInfo[] {
    return [...chains.values()].filter(c => c.family === family);
  }

  static getAllChainNames(): string[] {
    return [...chains.keys()];
  }

  static getChainInfoBySelector(selector: BN): ChainInfo {
    for (const chain of chains.values()) {
      if (chain.selector.eq(selector)) return chain;
    }
    throw new Error(`Unknown chain selector: ${selector.toString()}`);
  }
}
