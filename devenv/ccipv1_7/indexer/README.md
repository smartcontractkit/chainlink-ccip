# Indexer Service Boilerplate
Build dev image (hot-reload)
```
just docker-dev
```
Build production image
```
just docker
```

See [devenv](../services/indexer.go) wrapper and [environment](../environment.go) integration, run it with
```
ccip u
```

# Development

Since the product is in early stages we run all the services in dev mode by default, just change the code and it'd be reloaded.
In case of incompatible changes in other services that do not automatically reload run
```
ccip r
```