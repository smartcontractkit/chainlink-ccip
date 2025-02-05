# @chainlink/crib

## 2.3.0

### Minor Changes

- [#390](https://github.com/smartcontractkit/crib/pull/390) [`fba81a8`](https://github.com/smartcontractkit/crib/commit/fba81a8f852eac75e441d390c742013b87b6ff95) Thanks [@scheibinger](https://github.com/scheibinger)! - Downgrade postgres to v16 which is closer to what we have in production

- [#376](https://github.com/smartcontractkit/crib/pull/376) [`3cd6bf8`](https://github.com/smartcontractkit/crib/commit/3cd6bf897bc222b94b04c62de838f012e03a5507) Thanks [@scheibinger](https://github.com/scheibinger)! - Deploy geth and job-distributor on crib nodes, for more reliability

- [#372](https://github.com/smartcontractkit/crib/pull/372) [`033a503`](https://github.com/smartcontractkit/crib/commit/033a5037da1736cfcfcd5a3396368d1d034e4507) Thanks [@scheibinger](https://github.com/scheibinger)! - Add profile to enable service-monitor for prometheus

- [#370](https://github.com/smartcontractkit/crib/pull/370) [`a468bf9`](https://github.com/smartcontractkit/crib/commit/a468bf9d862197779aff2a51e0798e918a6eeffa) Thanks [@scheibinger](https://github.com/scheibinger)! - Make dashboard management avaialable via CRIB CLI

- [#375](https://github.com/smartcontractkit/crib/pull/375) [`605ac1b`](https://github.com/smartcontractkit/crib/commit/605ac1b720b04b1cb836f5fc00154f916128688d) Thanks [@scheibinger](https://github.com/scheibinger)! - Enable prometheus exporter by default in crib staging

- [#388](https://github.com/smartcontractkit/crib/pull/388) [`6ec0eba`](https://github.com/smartcontractkit/crib/commit/6ec0eba30f1277673cfce1c22bc1bcfc8a9b079e) Thanks [@scheibinger](https://github.com/scheibinger)! - Added geth-v2 dependency

## 2.2.0

### Minor Changes

- [#366](https://github.com/smartcontractkit/crib/pull/366) [`69d0372`](https://github.com/smartcontractkit/crib/commit/69d0372b8455dc7a24c2bbc77908a698d65f543a) Thanks [@scheibinger](https://github.com/scheibinger)! - Sync ccip-v2 setup with updated code for Load testing. Null out GAPConfig, as we'll rely on proxy"

## 2.1.0

### Minor Changes

- [#352](https://github.com/smartcontractkit/crib/pull/352) [`e8e4f75`](https://github.com/smartcontractkit/crib/commit/e8e4f754cc5142e50136ca345a643222faa6a266) Thanks [@scheibinger](https://github.com/scheibinger)! - Configure ccip-v2-scripts dependency to use ADDITIONAL_CHAINS_COUNT property for additional chains.

- [#356](https://github.com/smartcontractkit/crib/pull/356) [`e1affa7`](https://github.com/smartcontractkit/crib/commit/e1affa79dffc908aa1d6b3cb3e3586aee6439d8d) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - chainlink product now deploys pre-built image as default instead of building from `CHAINLINK_CODE_DIR`

- [#358](https://github.com/smartcontractkit/crib/pull/358) [`8efd8a2`](https://github.com/smartcontractkit/crib/commit/8efd8a2be8eb20e246a53aef9093fc19aaa3bdf0) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - Improved initial CRIB setup user prompt logic (PROVIDER is asked first w/ valid options presented, if `kind` is used `DEVSPACE_NAMESPACE` gets automatically set to `crib-local`)

- [#350](https://github.com/smartcontractkit/crib/pull/350) [`fdbddec`](https://github.com/smartcontractkit/crib/commit/fdbddec023f02fccf53df239cc0ac63c5d0b5b96) Thanks [@scheibinger](https://github.com/scheibinger)! - Update donut and geth dependencies to deploy variable number of geth chains. In donut add option to pass image URI and sha from .env file.

## 2.0.0

### Major Changes

- [#336](https://github.com/smartcontractkit/crib/pull/336) [`9343940`](https://github.com/smartcontractkit/crib/commit/93439405b1b5ffe855563eecc866fa1711d38c54) Thanks [@rafaelfelix](https://github.com/rafaelfelix)! - Changed all ingresses from class `alb` to Infra-Platform provided `nginx-internal`, for cost savings purposes. Versioned as major as it's a somewhat big change in the underlying infra of every CRIB setup, even though every component has been tested and proven to work with the new setup in a backwards-compatible way.

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
