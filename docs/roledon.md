# Role DON

The "Role DON" feature for Chainlink CCIP DONs introduces a flexible and scalable way 
of organizing oracle responsibilities across multiple blockchains.

Role DON enables distinct groups of oracles to be assigned specific roles, such as observing particular 
chains or performing specialized tasks.

## Purpose

Traditionally, all oracles in a DON were required to support all chains, 
which limited scalability and flexibility. With Role DON, each node (oracle) 
is assigned only to the chains it needs to observe, so it's no longer necessary 
for every node to support every chain. This enables more efficient resource allocation 
and supports scaling the network to more chains and roles.

## Commit Plugin Role DON

The Commit DON serves as a specialized DON responsible for computing Merkle roots and reporting prices.

As of version `e2013c5c97d60023be5c71bf002c97d02a3fb36f` it works in the following way:

- **Chain Assignment:** Each oracle is assigned to observe specific source and/or destination chains.
This assignment defines node’s role by configuring the chains it supports in the `Home Chain`.

- **Reading Sequence Numbers:**
    - The Commit DON, to compute Merkle roots, 
must read message sequence numbers from both source and destination chains.
    - **Off-Ramp Sequence Numbers:** These are read using destination chain readers.
The off-ramp contract exists on the destination chain and is responsible for storing the committed message sequence numbers.
    - **On-Ramp Sequence Numbers:** These are read using source chain readers.
The on-ramp contract exists on the source chains and is responsible for storing the latest pending message sequence numbers.
- **Observation Separation:** Not all oracles read from all chains. 
Each oracle only observes the chains relevant to its assigned role, 
reducing unnecessary workload and improving efficiency.
- **Consensus and Merkle Roots:** The Commit DON collects all relevant sequence number data, 
and then computes Merkle roots for message batches, and reaches consensus on these roots using OCR-based protocols.
The merkle roots are computed by source chain readers since computing them requires 
source chain support, since messages exist on the source chain.
- **Thresholds and Fault Tolerance:**
    - **FRoleDON:** The fault tolerance parameter at the DON (OCR instance) level, 
determining how many faulty oracles can be tolerated for DON-wide consensus.
    - **FChain:** The fault tolerance parameter at the chain level, 
tuned per chain, determining how many faulty oracles can be tolerated for chain-specific observations.
- **Reporting:** Once consensus is achieved, the Commit 
DON can commit Merkle root results for use in cross-chain message verification and settlement. This is a responsibility
of destination chain supporting oracles.

Similarly, token prices, gas prices, and execution reports are produced by the DON based on oracle assigned chains.
Oracles supporting source chains observe data from those source chains, while oracles supporting the destination chain 
observe data from the destination chain. This information is then propagated throughout the DON, where consensus is 
reached in the OCR3 Outcome phase, ensuring that all participants share a consistent view of the data.

## Benefits

- Improved scalability by allowing DONs to specialize in specific chains or roles.
- Enhanced resource efficiency, as nodes are only required to operate on chains relevant to their assigned role.
- Cost efficiency, since it is now possible for some chains to be operating with less amount of oracles.
- Greater flexibility for cross-chain applications and DON setups.
