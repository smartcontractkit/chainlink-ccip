# @chainlink/crib

## 1.4.0

### Minor Changes

- [#315](https://github.com/smartcontractkit/crib/pull/315) [`42e3ac0`](https://github.com/smartcontractkit/crib/commit/42e3ac0f5f6abf3425caebe49300e294835b1cb5) Thanks [@njegosrailic](https://github.com/njegosrailic)! - Adding support for labeling CRIB namespaces from the CLI for cost attribution

### Patch Changes

- [#339](https://github.com/smartcontractkit/crib/pull/339) [`0e527f8`](https://github.com/smartcontractkit/crib/commit/0e527f8f31e800df81707b4bbc82b21ca5200d74) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - grpcurl available on the nix shell

- [#340](https://github.com/smartcontractkit/crib/pull/340) [`4b96521`](https://github.com/smartcontractkit/crib/commit/4b96521d98f97dcddb4baebce61cae52e797759b) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - fix utils.SetupKubeConfig authInfo handling (prevents overwriting non-CRIB related entries)

## 1.3.0

### Minor Changes

- [#334](https://github.com/smartcontractkit/crib/pull/334) [`30e93f5`](https://github.com/smartcontractkit/crib/commit/30e93f57349176caad7ee52eb35aadb4a90feec6) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - cli command `devspace ingress-check` now supports Infra Platform managed ingress classes `nginx-internal` and `nginx-external`

## 1.2.0

### Minor Changes

- [#325](https://github.com/smartcontractkit/crib/pull/325) [`3ff989b`](https://github.com/smartcontractkit/crib/commit/3ff989b79a94fc97fbaccf4b9b163a79ae6b5ad3) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - removed cribbit.sh and references to it

### Patch Changes

- [#332](https://github.com/smartcontractkit/crib/pull/332) [`a2dc1c9`](https://github.com/smartcontractkit/crib/commit/a2dc1c906e424a1eec87490a909764d42843a593) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - setting a PriceMax on dev-simulated-core-ocr1 profile to prevent validation error

## 1.1.0

### Minor Changes

- [#298](https://github.com/smartcontractkit/crib/pull/298) [`1544e36`](https://github.com/smartcontractkit/crib/commit/1544e360c3309fcfddbbe33c574bcce7cd198e09) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - replaced ingress_check.sh with the new CLI command `crib devspace ingress-check`

## 1.0.0

### Major Changes

- [#253](https://github.com/smartcontractkit/crib/pull/253) [`f20e1db`](https://github.com/smartcontractkit/crib/commit/f20e1db369005617d92c583b61a49843e9a3b337) Thanks [@scheibinger](https://github.com/scheibinger)! - Deleted `go.work` to facilitate more flexible go module development within CRIB mono repo
  This is backward incompatible change.

### Patch Changes

- [#309](https://github.com/smartcontractkit/crib/pull/309) [`27397d9`](https://github.com/smartcontractkit/crib/commit/27397d99996399bdb3a7f4d28aba0f83ecdb6bd5) Thanks [@scheibinger](https://github.com/scheibinger)! - Fix CRIB nix config to use GOBIN from the go env provided by nix environment

## 0.2.1

### Patch Changes

- [#307](https://github.com/smartcontractkit/crib/pull/307) [`d029a3d`](https://github.com/smartcontractkit/crib/commit/d029a3d5e7deba507a13548bac14d76cb0ec1559) Thanks [@scheibinger](https://github.com/scheibinger)! - Fix docker login in kind provider. To be able to pull from sdlc or prod ECR registries we need login to them as well.

- [#310](https://github.com/smartcontractkit/crib/pull/310) [`f8a0afa`](https://github.com/smartcontractkit/crib/commit/f8a0afa2a2818295e81014cde1a8f67e89cae8f7) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - prevent errors from exiting the shell after setup is done

- [#311](https://github.com/smartcontractkit/crib/pull/311) [`3876989`](https://github.com/smartcontractkit/crib/commit/387698999ec43a0c82c349f2f5e7636c812c6b32) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - fixed kind ingress provisioning

## 0.2.0

### Minor Changes

- [#304](https://github.com/smartcontractkit/crib/pull/304) [`4620289`](https://github.com/smartcontractkit/crib/commit/46202896b97636c0ceed4ed3aeca5baf088d0e9a) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - tag checkouts download their corresponding crib CLI binary instead of latest in CI

## 0.1.1

### Patch Changes

- [#189](https://github.com/smartcontractkit/crib/pull/189) [`7382b00`](https://github.com/smartcontractkit/crib/commit/7382b00de78f4832a4fdf80d6eeade9db1bef160) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - first versioned release of CRIB
