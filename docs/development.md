# Development

## Getting Started

### Go Version

The Go version is specified in the project's [go.mod](../go.mod) file. You can install Go from the official [installation page](https://go.dev/doc/install).

### Running the Linter

We use `golangci-lint` as our linting tool. Run the linter with:

```sh
make lint
```

### Running Unit Tests

Run the unit tests with:

```sh
make test
```

### Generating Mocks

We use `mockery` to generate mocks, which are configured in the [mockery.yaml](../.mockery.yaml) file. Generate mocks with:

```sh
make generate
```

## Running Integration Tests

The E2E integration tests are maintained in the [Chainlink repository](https://github.com/smartcontractkit/chainlink). To ensure the `main` branch remains stable, we must verify that the integration tests there pass ✅.

The **Chainlink-CCIP** repository's CI runs a subset of these tests when you open a PR targeting `main`.

⚠️ Always follow the steps below to avoid unexpected failures.

### Integration Test Workflow

1. **Create a PR** in the **Chainlink-CCIP** repository with your proposed changes.
2. **CI Execution**: The CI will run a subset of integration tests from the [Chainlink repository](https://github.com/smartcontractkit/chainlink) after applying your changes.
3. **Handling Failures**:
    - If tests fail, investigate the cause:
        1. A bug in your changes.
        2. Flaky tests.
        3. Breaking changes introduced by your PR.
    - If the failure is due to flaky tests, re-run the failing tests once or twice.
    - If you introduced breaking changes:
        - Create a branch in the **Chainlink** repository.
        - Run:
          ```sh
          go get github.com/smartcontractkit/chainlink-ccip@<your-branch-commit-sha>
          ```
        - Migrate to your changes by fixing any errors or failing tests.
        - Update your **Chainlink-CCIP** PR by specifying the corresponding **Chainlink** commit hash in the PR description:
          ```
          core ref: <commit-sha>
          ```
          ⚠️ Avoid using single quotes (`'`) in your PR description if you include `core ref`.
4. **Merging the PR**:
    - Once your **Chainlink-CCIP** PR is approved, merge it. The E2E tests on `main` may temporarily fail—this is expected.
5. **Updating Chainlink**:
    - Return to your **Chainlink** PR and update the **Chainlink-CCIP** version to the latest `main` SHA.
6. **Final Integration Test Run**:
    - Once integration tests pass, merge your **Chainlink** PR into `develop`.
    - Re-run integration tests on **Chainlink-CCIP** to verify stability.