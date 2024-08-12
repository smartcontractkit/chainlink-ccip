# chainlink-ccip

This is the repo that implements the OCR3 CCIP plugins. This includes the commit and execution plugins.

## Development Cycle

1. Create a PR on chainlink-ccip with the changes you want to make.
2. CI will run the integration test in the CCIP repo after applying your changes.
3. If the integration test fails, make sure to fix it first before merging your changes into
the `ccip-develop` branch of chainlink-ccip. You can do this by:
    - Creating a branch in the CCIP repo and running `go get github.com/smartcontractkit/chainlink-ccip@<your-branch-commit-sha>`.
    - Fixing the build/tests.
4. Once your ccip PR is approved, merge it.
5. Go back to your chainlink-ccip PR and re-run the integration test workflow.
6. Once the integration test passes, merge your chainlink-ccip PR into `ccip-develop`, however do not delete the branch on the remote.
7. Create a new PR in ccip that points to the newly merged commit in the `ccip-develop` tree and merge that.
