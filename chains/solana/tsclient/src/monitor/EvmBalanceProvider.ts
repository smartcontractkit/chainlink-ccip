import { IBalanceProvider } from './IBalanceProvider';
import { BalanceInfo } from './types/BalanceInfo';
import { TokenAccount } from './types/TokenAccount';
import { NativeAccount } from './types/NativeAccount';
import { ChainInfo } from '../common/ChainRegistry';
import { ethers } from 'ethers';
import ERC20 from '@openzeppelin/contracts/build/contracts/ERC20.json';

export class EvmBalanceProvider implements IBalanceProvider {
  getNativeBalances(chainInfo: ChainInfo, accounts: NativeAccount[]): Promise<BalanceInfo[]> {
    const connection = new ethers.JsonRpcProvider(chainInfo.rpcURL);
    return Promise.all(
      accounts.map(async ({ owner, name }) => {
        const balance = await connection.getBalance(owner);
        return {
          chain: chainInfo.name,
          name,
          owner,
          balance: `${ethers.formatEther(balance)} ETH`,
        };
      })
    );
  }

  async getTokenBalances(chainInfo: ChainInfo, accounts: TokenAccount[]): Promise<BalanceInfo[]> {
    const provider = new ethers.JsonRpcProvider(chainInfo.rpcURL);

    return Promise.all(
      accounts.map(async ({ type, mint, owner, name }) => {
        try {
          // Skip non-ERC20 tokens (explicitly typed)
          if (type && type.toUpperCase() !== 'ERC20') {
            console.warn(`[EVM] Skipping unsupported token type "${type}" for ${name} (${owner})`);
            return {
              chain: chainInfo.name,
              name,
              owner,
              balance: '0',
            };
          }
          
          const contract = new ethers.Contract(mint, ERC20.abi, provider);
          const rawBalance = await contract.balanceOf(owner);
          const decimals = await contract.decimals();
          const symbol = await contract.symbol();

          return {
            chain: chainInfo.name,
            name: `${name} (${symbol})`,
            owner,
            balance: `${ethers.formatUnits(rawBalance, decimals)} ${symbol}`
          };
        } catch (err) {
          console.warn(`[EVM] Failed to fetch token balance for ${name} (${owner} @ ${mint}): ${err}`);
          return {
            chain: chainInfo.name,
            name,
            owner,
            balance: '0',
          };
        }
      })
    );
  }
}
