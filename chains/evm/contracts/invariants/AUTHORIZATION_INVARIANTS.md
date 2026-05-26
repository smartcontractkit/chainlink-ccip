# Authorization & Cryptographic Safety Invariants

## 1. Signer/Key Uniqueness

- **INV-AUTH-1**: Any set of cryptographic keys used for quorum verification must enforce uniqueness. Duplicate keys in the signer set must be rejected at configuration time. A single compromised key must not be able to satisfy a multi-key threshold by appearing multiple times in the signer list.

---

## 2. Signature Domain Separation

- **INV-AUTH-2**: Cryptographic signatures must include domain separation that binds the signature to the specific verifier instance, chain, and context. The signed payload must include the verifier's own identity (address, instance ID, or equivalent). The same signature must not be valid across multiple verifier instances, even if they share signer keys.
