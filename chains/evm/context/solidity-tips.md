## Solidity Contribution Guidelines

### 1. General Principles

- **Think first, code second**: Minimize the number of lines changed and consider ripple effects across the codebase.
- **Prefer simplicity**: Fewer moving parts ➜ fewer bugs and lower audit overhead.

### 2. Assembly Usage

| Rule                                                                                                      | Rationale                                                             |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| Use assembly only when essential.                                                                         | Keeps code readable and auditable.                                    |
| Assembly is mandatory for low-level external calls.                                                       | Gives full control over call parameters & return data, and saves gas. |
| Precede every assembly block with: • A brief justification (1-2 lines). • Equivalent Solidity pseudocode. | Documents intent for reviewers.                                       |
| Mark assembly blocks memory-safe when the Solidity docs' criteria are met.                                | Enables compiler optimizations.                                       |

### 3. Gas Optimization

- Keep a dedicated **Gas Optimization** section in the PR description; justify any measurable gas deltas.
- Prefer `calldata` over `memory`.
- Limit storage (`sstore`, `sload`) operations; cache in memory wherever possible.
- Use forge snapshot, forge test --match-test "benchmark", and npm scripts:
  ```bash
  npm run snapshot:main   # captures gas baseline from main
  npm run diff:main       # compares your branch vs. main
  ```
- Large regressions must be explained.

### 4. Handling "Stack Too Deep"

- **Struct hack (tests only)**: Bundle local variables into a temporary struct declared above the test.
- **Scoped blocks**: Wrap code in `{ ... }` to drop unused vars from the stack.
- **Internal helper functions**: Encapsulate logic to shorten call frames.
- **Refactor / delete unnecessary variables before other tricks**.

### 5. Security Checklist

- Review every change with an adversarial mindset.
- Favor the simplest design that meets requirements.
- After coding, ask: "What new attack surface did I introduce?"
- Reject any change that raises security risk without strong justification.

### 6. Verification Workflow

```bash
export FOUNDRY_PROFILE=ccip
forge build                    # compile
forge test                     # full test suite
forge snapshot                 # gas snapshot (local)
forge test --match-test bench  # run benchmarks
npm run snapshot:main          # baseline gas (main)
npm run diff:main              # gas diff vs. main
```

### 7. Continuous Learning

- Consult official Solidity docs and relevant project references when uncertain.
- Borrow battle-tested patterns from audited codebases.

Apply these rules rigorously before opening a PR.

### Error Handling Style

Always use custom errors with the revert pattern instead of require statements:

```solidity
// ❌ Don't use require with string messages
require(amount > 0, "Amount must be positive");
require(to != address(0), "Cannot transfer to zero address");

// ✅ Do use custom errors with if/revert pattern
error AmountMustBePositive();
error CannotTransferToZeroAddress();

if (amount == 0) revert AmountMustBePositive();
if (to == address(0)) revert CannotTransferToZeroAddress();
```

**Benefits of custom errors**:

- More gas efficient than require strings
- Better error identification in tests and debugging
- Cleaner, more professional code
- Consistent with modern Solidity best practices

This applies to all Solidity code including contracts, libraries, and scripts.
