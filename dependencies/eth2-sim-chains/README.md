# DevChains  
Devspace components to launch Simulated/Private Chains

## Usage
Include any of the components as a as dependency in your setup.

Example:
```yaml
version: v2beta1
name: chainlink

dependencies:
  besu-prysm:
    path: ../../dependencies/eth2-sim-chains/besu-prysm.yaml
vars:
  DEVSPACE_ENV_FILE: .env
```


## geth-prysm
Utilizes common chart from CTF repo to deploy:
* Single Geth Node (Execution layer)
* Single Prysm Beacon Node (Consensus layer)
* Single Prysm Validator (Consensus layer)

Each of them is deployed as STS and backed with PVC for persistence during restarts

## besu-prysm
Utilizes common chart from CTF repo to deploy:
* Single Besu Node (Execution layer)
* Single Prysm Beacon Node (Consensus layer)
* Single Prysm Validator (Consensus layer)

Each of them is deployed as STS and backed with PVC for persistence during restarts

## Limitations
Currently, the setup is not able to survive Node restarts, we have a plan to address it in [this ticket](https://smartcontract-it.atlassian.net/browse/CRIB-434)