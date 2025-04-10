import { BalanceFetcher } from './monitor/BalanceFetcher';
import { ChainAccounts } from './monitor/types/ChainAccounts';
import {balanceAccounts} from "./staging"


async function main() {
  const fetcher = new BalanceFetcher();
  const balances = await fetcher.getBalances(balanceAccounts as any as ChainAccounts);

  console.log('Staging accounts balances:');

  for (const { chain, name, owner, balance } of balances) {
    console.log(`[${chain}] ${name} (${owner}): ${balance}`);
  }
}

main().catch((err) => {
  console.error('Error fetching balances:', err);
});
