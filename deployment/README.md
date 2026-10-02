# CCIP Deployment Tooling API

Chain-agnostic changesets and adapter interfaces for deploying and operating CCIP across chain families. Each chain family (EVM, Solana, …) implements the interfaces and registers them from `init()`. Changesets dispatch to them by chain selector.

The API has two generations:

- **v1** (CCIP 1.6): `deploy/`, `lanes/`, `fees/`, …
- **v2** (CCIP 2.0): `v2_0_0/`

Tokens, MCMS, ownership, curse, and fees are shared between them.

Documentation lives in [`docs/`](docs/README.md):

- [Architecture](docs/architecture.md)
- [Consuming the API](docs/consuming.md)
- [Implementing 1.6](docs/implementing-1.6.md) and [Implementing 2.0](docs/implementing-2.0.md)
- [Interfaces reference](docs/interfaces.md)
- [Changeset style guide](docs/style-guide.md)

[Original design doc](https://docs.google.com/document/d/1axlWfAzINa_g01DrmbRMncgncKqgeKRMOzgch2S16Cs/edit?tab=t.0#heading=h.4e7ng0tekbkc) (historical; the docs above describe the current code).
