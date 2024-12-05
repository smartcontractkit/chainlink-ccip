# ccip-v2-scripts
Devspace component to integrate with the CCIP V2 deployment tooling 
> product agnostic set of environment abstractions used
to deploy and configure products including both on/offchain
dependencies.


## Runbooks
### Syncing go.mod with the upstream

* Grab the sha of the chainlink repo, that you want to sync with. `CHAINLINK_SHA=<paste>`
* Sync integration tests module `go get "github.com/smartcontractkit/chainlink/deployment@$CHAINLINK_SHA`
* Update `replace` directives in `go.mod` file
  * Compare them with deployment package `go.mod` and sync 
  * Update `github.com/smartcontractkit/chainlink/v2` references to use the same sha. For example when  
    the sha for deployment package is `v0.0.0-20241202095458-c8a829dc327e` the v2 should be the same, just 
    with different version in the beginning: `v2.0.0-20241202095458-c8a829dc327e`, `v2` instead of `v0`

There is WIP script that could help partly with that process [scripts/bump-chainlink-dep.sh](./../../scripts/bump-chainlink-dep.sh)