## How to add new services to your local environment

1. Create a new file `services/service.go`
2. Define [Input](indexer.go) and [Output](indexer.go) structs, implement `NewX(*Input) (*Output, error)`
3. Go to [environment.go](../environment.go) add config and connect your service
    ```
    type Cfg struct {
        Indexer         *services.IndexerInput `toml:"indexer"`
        ...
    }
    ```
    ```
    func localEnvironment() (*Cfg, error) {
        ...
        _, err = services.NewX(in.InputX)
        if err != nil {
            return nil, fmt.Errorf("failed to create an example service: %w", err)
        }
    ```

4. Define profiles in one or more `env.toml` files and run the environment: `ccip u` or `ccip -c custom.toml u`
    ```
    [service_example]
      image = "f4hrenh9it/df-fakes:latest"
      port = 8332
      [service_example.db]
        image = "postgres:16-alpine"
    ```
5. Add needed CLI commands [here](../cmd/ccip.go), for example to inspect the database, etc

## Good practices
- Do not expose `testcontainers.Container` or other implementation details, use basic types in `Input` and `Output`. Raw types makes easy to replace any component and reusing tests and tooling on any environment!
- Prefer static ports, it makes chaos testing and tooling easier, see [h.PortBindings](indexer.go)
- Embed components that are required only by your service (for example PostgreSQL). Expose only params that make sense for tests/users, otherwise use default constants.
- Keep it simple: if your component is mostly static and rarely change follow steps 1-3 to integrate it. If your component requires configurability go through steps 3-5, add `TOML` config and potentially implement CLI commands in our common [CLI](../cmd/ccip.go).
  Document both `TOML` and `CLI` commands so other developers can use it.

## How to make a service hot-reloadable

See an example [indexer](../indexer/README.md) boilerplate.

Since the product is in early stage we run all the services in dev mode with automatic reload.