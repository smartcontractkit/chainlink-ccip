# cre-dev deployment

- cluster-services
  - [x] otel-gateway
  - minio-operator
  - minio-tenant
- blockchain
  - [x] anvil
  - [x] anvil-blockscout
- chainlink
  - [x] cre don: 1 boot / 1 gateway / 4 nodes
  - [x] job-distributor
- chainlink-config
  - [ ] pre: onchain / contracts
  - [ ] post: jobs
  - [ ] cre-cli

## development

- login to stage to pull from sdlc

```sh
aws sso login --profile stage
aws ecr get-login-password --region us-west-2 --profile stage \
  | helm registry login --username AWS --password-stdin 804282218731.dkr.ecr.us-west-2.amazonaws.com
```

- fetch crib cli

```sh
task fetch-cli
asdf reshim
```

- deploy into single namespace

```sh
cd deployments/cre-dev

# copy env file
cp .env.example .env

# cluster-services
devspace run cluster-services

# blockchain: anvil / anvil-blockscout
devspace run blockchain
devspace run blockchain-reload

# chainlink: cre don / jd
devspace run chainlink

# nuke
helm list -q | xargs -r helm uninstall
kg pvc --no-headers | awk '{print $1}' | xargs -r kubectl delete pvc
```

### capabilities

- docker image: `804282218731.dkr.ecr.us-west-2.amazonaws.com/chainlink-develop:sha-34af645425-bcm-swift-poc`

- custom binaries in docker image

```sh
$ ls /usr/local/bin
attestaccount  chainlink-cosmos     chainlink-ocr3-capability  cron                 detectunlock         lock               streams
batchkvread    chainlink-feeds      chainlink-solana           detectattestaccount  kvstore              log-event-trigger  unlock
batchkvwrite   chainlink-medianpoc  chainlink-starknet         detectcreateaccount  libs                 readcontract       workflowevent
chainlink      chainlink-mercury    createaccount              detectlock           loadtestwritetarget  sign
```
