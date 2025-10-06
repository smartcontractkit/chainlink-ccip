# Go bindings

After a version is released, the corresponding bindings will be moved to a versioned directory (e.g., `v1.0.0`).
Contracts in the versioned directory are guaranteed to be stable and audited.

Any contract bindings in the `latest` directory are not guaranteed to be stable or audited, and are only intended for testing and development, not for mainnet use.

Versions will be tagged in git using the format `contracts-ccip-v<x.y.z>`, e.g., `contracts-ccip-v1.0.0`.
Versioned directories will be named using the format `v<x_y_z>`, e.g., `v1_0_0`.