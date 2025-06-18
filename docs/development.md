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

### Running Integration Tests Locally Using Docker

Integration tests on Mac with ARM64 architecture have been reported to take over 10 minutes.
As a temporary workaround, you can use the following Docker-based setup to run them more efficiently.

One-Time Container Setup
```bash
# Run the following command once to set up the container
docker run --platform linux/amd64 -it --name amd64-ubuntu -v $HOME/amd64-root:/root ubuntu:22.04
```

You’ll now be inside the container:
```bash
# Install essential packages
apt update && apt install -y build-essential wget vim

# Install Go
wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

# Install PostgreSQL
apt install postgresql-14
```

Now configure PostgreSQL:
```bash
vim /etc/postgresql/14/main/pg_hba.conf  # Change 'peer' or 'scram-sha-256' to 'trust'
service postgresql restart
psql -U postgres
```

Inside the `psql` prompt:
```bash
create database cl_test;
\q
```

Then:
```bash
echo 'export CL_DATABASE_URL="postgres://postgres@localhost:5432/cl_test"' >> ~/.bashrc
```

**Running the Tests**
```bash
# Start and attach to the container (if stopped, i.e. you exited)
docker start -ai amd64-ubuntu
service postgresql start

# From your local machine, copy the Chainlink repo to the container
docker exec amd64-ubuntu rm -rf /app && docker cp ./chainlink amd64-ubuntu:/app
```

Inside the container, navigate to the Chainlink repo and prepare the test DB:
```bash
cd /app
go run ./core/store/cmd/preparetest
```

Then, run the integration tests:
```bash
cd integration-tests  # This must be done — it's a separate Go module
go test -v -timeout 5m -run "Test_CCIPTopologies_EVM2EVM_RoleDON_AllSupportSource_SomeSupportDest" ./...
```

You can now make changes changes in your localhost and run tests on docker using:
```bash
# copy the changes to docker (localhost)
docker exec amd64-ubuntu rm -rf /app && docker cp ./chainlink amd64-ubuntu:/app

# re-run the tests (docker container)
go test -v -timeout 5m -run "Test_CCIPTopologies_EVM2EVM_RoleDON_AllSupportSource_SomeSupportDest" ./...
```
