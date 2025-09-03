# Chainlink CCV

## Getting started

Run all the services in local environment, install [Nix](https://github.com/DeterminateSystems/nix-installer) and run
```bash
nix develop
just clean-docker-dev # needed in case you have old JD image
just build-docker-dev
```
Then enter the `ccv` shell and run `up`
```bash
ccv sh
```