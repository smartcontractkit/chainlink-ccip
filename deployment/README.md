# CCIP Deployments

CCIP deployments often target multiple chain families (and product versions, in the case of token pools). Each of these deployment stories have a singular changeset capable of interfacing with different logic per chain family. This document describes how these changesets are designed, how to run them, and how to extend support to new chain families.

## Design

This diagram provides an overview of all components and their locations. Connections are only made explicit made for the token path to avoid polluting the diagram.

```mermaid
flowchart LR
    subgraph chainlink-deployments
        dp[Durable Pipelines]
        ds[DataStore]
    end

    subgraph Chain Family Repo
        cfmr[MCMSReader]
        subgraph Sequences
            subgraph v1.6.0
                clls1_6[ConfigureLaneLegAsSource]
                clld1_6[ConfigureLaneLegAsDest]
                cllb1_6[ConfigureLaneLegBidirectionally]
                ctft1_6[ConfigureTokenForTransfers]
            end
            subgraph v1.7.0
                clls1_7[ConfigureChainForLanes]
                ctft1_7[ConfigureTokenForTransfers]
            end
            subgraph v1.5.0
                ctft1_5[ConfigureTokenForTransfers]
            end
        end
        subgraph Helpers
            artb[AddressRefToBytes]
            dtfp[DeriveTokenFromPool]
        end
    end

    subgraph chainlink-ccip/deployment
        subgraph changesets
            ctft[ConfigureTokensForTransfers]
            cc1_6[v1_6.ConnectChains]
            cc1_7[v1_7.ConnectChains]
        end
        subgraph interfaces
            mr[MCMSReader]
            ta[TokensAdapter]
            ca1_6[v1_6.ChainAdapter]
            ca1_7[v1_7.ChainAdapter]
        end
        subgraph registries
            tar[TokenAdapterRegistry]
            car1_6[v1_6.ChainAdapterRegistry]
            car1_7[v1_7.ChainAdapterRegistry]
            mrr[MCMSReaderRegistry]
        end
    end

    ctft --> tar
    ctft --> mrr

    tar --> ta
    mrr --> mr

    ta ----> artb
    ta ----> dtfp
    ta ----> ctft1_5
    ta ----> ctft1_6
    ta ----> ctft1_7
    mr ----> cfmr

    dp --init--> registries
    dp --run--> changesets
```

## Usage

Within the `PipelinesRegistryProvider.Init()` function in `chainlink-deployments/domains/<CCIP_DOMAIN>/<ENVIRONMENT>/pipelines.go`, you simply need to construct the required registries and initialize the changesets with them. Your pipelines are then ready for execution. For the most up-to-date information on how to execute a durable pipeline, see the [official documentation](https://docs.cld.cldev.sh/guides/pipelines/).

```golang
// NOTE: This is still example code - types are not guaranteed to exist as described here.
// TODO: Update when types and names are solidified.

mrr := changesets.NewMCMSReaderRegistry()
tar := tokens.NewTokenAdapterRegistry()

mrr.RegisterMCMSReader("evm", evm.MCMSReader{})
mrr.RegisterMCMSReader("solana", solana.MCMSReader{})
// ... and so on

tar.RegisterTokenAdapter("evm", semver.MustParse("1.7.0"), evm1_7_0.TokensAdapter{})
tar.RegisterTokenAdapter("solana", semver.MustParse("1.6.2"), evm1_7_0.TokensAdapter{})
// ... and so on

registry.Add("configure-tokens-for-transfers",
		changeset.Configure(tokens.ConfigureTokensForTransfers(mrr, tar)).WithEnvInput())
```

## New Chain Families

To implement a new chain family, implement each adapter. i.e.

- MCMSReader
- TokensAdapter (1 for each token pool version supported on the family)
- All ChainAdapters that must be supported by the family

Then, simply register your adapters in `chainlink-deployments/domains/<CCIP_DOMAIN>/<ENVIRONMENT>/pipelines.go`.
