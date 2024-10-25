# Cross-Chain Interoperability Protocol (CCIP)

A protocol designed to transfer data and value across blockchains in a trust
minimized fashion. This is done using a combination of checks on the source
chain(s), destination chains, and offchain security. This document describes
the most important security considerations along with the messaging protocol
at the data and event processing level. As this repository is for the off-chain
portion of CCIP, we'll focus on off-chain components and the on-chain
integration points.

The user interface is not described here. It is largely unchanged compared to
the earlier architecture and the [public docoumentation][public-docs] provides
an excellent guide for how to use CCIP.

## OffChain Reporting (OCR)

The [libOCR][ocr-repo] framework provides security to the system with a
Distributed Oracle Network (DON) utilizing a [byzantine fault tolerant][bft]
protocol. Security is assumed given the [Reporting Plugin][ocr-interface] built
within this framework ensures observation thresholds on the data used by the
system. See the [Consensus and Sub-consensus](consensus.md) page for more
information about how the CCIP plugins use and extend this the typical OCR
framework.

## Plugins

The protocol is implemented in two phases, each with a corresponding libOCR
plugin: commit and execute.

### Commit

In the commit phase, user created messages are collected from every source
chain, a merkle root is generated for messages from each source, and the roots
are put into a [Commit Report][commit-report-src]. The report is sent to the
on-chain code for final processing where a commit report log is written to the
blockchain.

An independant Risk Management Network (RMN) is used as a second layer of
security to ensure the merkle roots are valid. RMN can be thought of as a
validation network with veto power. In the previous version of CCIP RMN was
an independent component which provided a blessing or curse after the commit
report was written to the network. With this architecture the plugin queries
the RMN network directly so that the RMN signatures can be included in the
initial report.

The report includes additional gas and token price data required by the billing
algorithm. Strictly speaking, it is not part of the protocol and could be
implemented separately. For convenience it is included in the commit plugin.

More detail about the implementation can be found in the [README](commit#readme).

### Execute

In the execute phase, the plugin searches for commit reports with pending
executions. All messages for these commit reports are gathered and a
[special merkle proof][merklemulti] is computed for all messsages ready for
execution. These proofs are put into the [Execute Plugin Report][exec-report-src].

Due to the [Role DON](consensus.md#role-don) architecture, this
process has to be done across several rounds of consensus:

1. Destination readers look for commit reports.
2. Based on the commit reports, source readers fetch the messages.
3. Based on the messages, the destination reader determines execution order.

The commit report is sent to the on-chain code for final processing. This
includes all token transfers, data handling and user contract interaction.

[public-docs]: https://docs.chain.link/ccip
[ocr-repo]: https://github.com/smartcontractkit/libocr/tree/master
[bft]: https://en.wikipedia.org/wiki/Byzantine_fault
[ocr-interface]: https://github.com/smartcontractkit/libocr/blob/adbe57025f12b9958907bb203acba14360bf8ed2/offchainreporting2plus/ocr3types/plugin.go#L165
[commit-report-src]: https://github.com/smartcontractkit/chainlink-ccip/blob/0f6dce5d1fdb67b3127332ac729191f2c1c790ff/pkg/types/ccipocr3/plugin_commit_types.go#L19
[merklemulti]: https://github.com/smartcontractkit/chainlink-common/tree/main/pkg/merklemulti
[exec-report-src]: https://github.com/smartcontractkit/chainlink-common/tree/main/pkg/merklemulti
