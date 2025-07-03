# CCIP Solidity Test Pattern Documentation for AI Agents

## Overview

This document provides comprehensive guidance for AI agents to write Solidity tests following the CCIP (Chainlink Cross-Chain Interoperability Protocol) test pattern. The pattern follows a one-test-file-per-function approach, where each contract function gets its own dedicated test file containing all success and failure test cases.

## Core Principles

1. **One File Per Function**: Each function in a contract has its own dedicated test file
2. **Comprehensive Coverage**: Each test file contains all success and failure cases for that function
3. **Predictable Naming**: Consistent naming conventions throughout
4. **Clear Organization**: Success cases first, then failure cases marked with `// Reverts` comment
5. **Foundry Framework**: All tests use Foundry's testing framework

## Directory Structure

The test directory structure mirrors the source contract file structure. Tests are placed under `contracts/test/` following the same path structure as the source contracts.

### Example: FeeQuoter

```
contracts/
├── FeeQuoter.sol                                    # Source contract
└── test/
    ├── BaseTest.t.sol                               # Base test utilities
    ├── TokenSetup.t.sol                             # Token test utilities
    └── feeQuoter/                                   # Test directory (lowercase)
        ├── FeeQuoterSetup.t.sol                     # Setup file for FeeQuoter tests
        ├── FeeQuoter.constructor.t.sol              # Constructor tests
        ├── FeeQuoter.applyDestChainConfigUpdates.t.sol
        ├── FeeQuoter.getTokenPrice.t.sol
        ├── FeeQuoter.updatePrices.t.sol
        └── FeeQuoter.[functionName].t.sol          # One file per function

### Example: Nested Contract (OffRamp)
```

contracts/
├── offRamp/
│ └── OffRamp.sol # Source contract in subdirectory
└── test/
└── offRamp/
└── OffRamp/ # Note: additional OffRamp directory
├── OffRampSetup.t.sol # Setup file
├── OffRamp.constructor.t.sol # Constructor tests
├── OffRamp.execute.t.sol # Execute function tests
├── OffRamp.commit.t.sol # Commit function tests
└── OffRamp.[functionName].t.sol # One file per function

````

### Key Points:
- Test directory structure mirrors the source contract structure
- Test directories use camelCase (e.g., `feeQuoter/`, `offRamp/`)
- For nested source contracts, tests may have an additional directory level
- Setup files are named `{ContractName}Setup.t.sol`
- Function test files are named `{ContractName}.{functionName}.t.sol`
- All test files end with `.t.sol`

## File Naming Convention

### Pattern: `ContractName.functionName.t.sol`

Examples:
- `OffRamp.constructor.t.sol` - Tests for constructor
- `OffRamp.execute.t.sol` - Tests for execute function
- `FeeQuoter.applyDestChainConfigUpdates.t.sol` - Tests for applyDestChainConfigUpdates
- `TokenPool.setChainRateLimiterConfig.t.sol` - Tests for setChainRateLimiterConfig

## Contract Naming Convention

### Pattern: `ContractName_functionName`

```solidity
contract OffRamp_constructor is OffRampSetup {
    // constructor tests
}

contract OffRamp_execute is OffRampSetup {
    // execute function tests
}
````

## Function Naming Convention

### Pattern: `test{_fuzz_}{DescriptiveName}{Success|Reverts}`

Components:

- `test` - Required prefix for all test functions
- `_fuzz_` - Optional, indicates fuzz test
- `{DescriptiveName}` - Describes what the test does (PascalCase)
- `{Success|Reverts}` - Optional suffix, omit for success cases

### Examples:

#### Success Test Cases (no suffix needed):

```solidity
function test_Constructor() public { }
function test_SingleReport() public { }
function test_MultipleReports() public { }
function test_LargeBatch() public { }
function test_applyDestChainConfigUpdates() public { }
function testFuzz_SetChainRateLimiterConfig_Success(uint128 capacity, uint128 rate) public { }
```

#### Revert Test Cases:

```solidity
function test_RevertWhen_ZeroOnRampAddress() public { }
function test_RevertWhen_UnauthorizedTransmitter() public { }
function test_RevertWhen_OnlyOwnerOrRateLimitAdmin() public { }
function test_RevertWhen_InvalidDestChainConfig() public { }
```

## Test File Structure Template

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

// Import the contract being tested
import {ContractName} from "../../../path/to/ContractName.sol";
// Import setup file
import {ContractNameSetup} from "./ContractNameSetup.t.sol";
// Import other dependencies as needed

contract ContractName_functionName is ContractNameSetup {
    // Optional: Override setUp if specific setup needed
    function setUp() public virtual override {
        super.setUp();
        // Additional setup specific to this function's tests
    }

    // Success test cases first (no "Success" suffix needed)
    function test_BasicFunctionality() public {
        // Test implementation
    }

    function test_WithSpecificScenario() public {
        // Test implementation
    }

    function testFuzz_WithVariableInputs(uint256 amount, address user) public {
        // Fuzz test implementation
    }

    // Reverts section - clearly marked with comment
    // Reverts

    function test_RevertWhen_InvalidInput() public {
        vm.expectRevert(ContractName.CustomError.selector);
        // Call that should revert
    }

    function test_RevertWhen_UnauthorizedCaller() public {
        vm.startPrank(STRANGER);
        vm.expectRevert(abi.encodeWithSelector(ContractName.Unauthorized.selector, STRANGER));
        // Call that should revert
    }
}
```

## Setup File Pattern

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseTest} from "../../BaseTest.t.sol";
// Other imports

contract ContractNameSetup is BaseTest {
    // State variables for contract instances
    ContractName internal s_contract;

    // Test constants
    uint256 internal constant TEST_VALUE = 100;

    function setUp() public virtual override {
        super.setUp();

        // Deploy contracts
        s_contract = new ContractName(/* constructor args */);

        // Common setup operations
    }

    // Helper functions used across multiple test files
    function _generateTestData() internal pure returns (bytes memory) {
        // Implementation
    }

    // Assertion helpers
    function _assertConfigEqual(Config memory a, Config memory b) internal pure {
        assertEq(a.field1, b.field1);
        assertEq(a.field2, b.field2);
    }
}
```

## Common Test Patterns

### 1. Event Testing

```solidity
vm.expectEmit();
emit ContractName.EventName(param1, param2);
// Perform action that emits event
```

### 2. Access Control Testing

```solidity
function test_RevertWhen_OnlyOwner() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable.OnlyCallableByOwner.selector);
    s_contract.restrictedFunction();
}
```

### 3. Fuzz Testing

```solidity
function testFuzz_FunctionName_Success(uint256 amount) public {
    vm.assume(amount > 0 && amount < type(uint128).max);
    // Test with fuzzed input
}
```

### 4. Multiple Scenarios in One Test

```solidity
function test_MultipleScenarios() public {
    // Scenario 1
    _performAction1();
    _assertState1();

    // Scenario 2
    _performAction2();
    _assertState2();
}
```

### 5. State Verification Pattern

```solidity
function test_StateChanges() public {
    // Capture initial state
    uint256 initialValue = s_contract.getValue();

    // Perform action
    s_contract.updateValue(newValue);

    // Assert state changes
    assertEq(s_contract.getValue(), newValue);
    assertTrue(s_contract.getValue() != initialValue);
}
```

## Key Conventions

### 1. Variable Naming

- State variables prefix: `s_` (e.g., `s_offRamp`, `s_tokenPool`)
- Internal functions prefix: `_` (e.g., `_generateReport()`, `_assertState()`)
- Constants: `UPPER_SNAKE_CASE` (e.g., `DEST_CHAIN_SELECTOR`, `GAS_LIMIT`)

### 2. Test Data Generation

```solidity
// Use helper functions for complex data
Internal.Any2EVMRampMessage[] memory messages = _generateSingleBasicMessage(SOURCE_CHAIN_SELECTOR_1, ON_RAMP_ADDRESS_1);

// Use descriptive variable names
RateLimiter.Config memory newOutboundConfig = RateLimiter.Config({
    isEnabled: true,
    capacity: capacity,
    rate: rate
});
```

### 3. Assertion Patterns

```solidity
// Direct assertions for simple values
assertEq(actualValue, expectedValue);
assertTrue(condition);
assertFalse(condition);

// Custom assertion helpers for complex types
_assertConfigEqual(actualConfig, expectedConfig);
_assertExecutionStateChangedEventLogs(/* params */);
```

### 4. VM Cheatcodes Usage

```solidity
// Pranking
vm.startPrank(USER);
// ... actions as USER
vm.stopPrank();

// Expecting reverts
vm.expectRevert(CustomError.selector);
vm.expectRevert(abi.encodeWithSelector(ErrorWithParams.selector, param1, param2));

// Time manipulation
vm.warp(block.timestamp + 1 hours);

// Event expectations
vm.expectEmit();
emit EventName(param1, param2);
```

## Test Organization Best Practices

1. **Group Related Tests**: Keep similar test scenarios close together
2. **Clear Section Separation**: Use `// Reverts` comment to separate success and failure cases
3. **Descriptive Test Names**: Test names should clearly indicate what is being tested
4. **One Assertion Focus**: Each test should focus on testing one specific behavior
5. **Setup Reuse**: Use setup files to avoid duplication across test files

## Example Test Implementation

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Internal} from "../../libraries/Internal.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_applyDestChainConfigUpdates is FeeQuoterSetup {
    function test_applyDestChainConfigUpdates() public {
        FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = new FeeQuoter.DestChainConfigArgs[](2);
        destChainConfigArgs[0] = _generateFeeQuoterDestChainConfigArgs()[0];
        destChainConfigArgs[1] = _generateFeeQuoterDestChainConfigArgs()[0];
        destChainConfigArgs[1].destChainSelector = DEST_CHAIN_SELECTOR + 1;

        vm.expectEmit();
        emit FeeQuoter.DestChainConfigUpdated(DEST_CHAIN_SELECTOR, destChainConfigArgs[0].destChainConfig);
        vm.expectEmit();
        emit FeeQuoter.DestChainAdded(DEST_CHAIN_SELECTOR + 1, destChainConfigArgs[1].destChainConfig);

        s_feeQuoter.applyDestChainConfigUpdates(destChainConfigArgs);

        // Verify state
        FeeQuoter.DestChainConfig memory gotConfig = s_feeQuoter.getDestChainConfig(DEST_CHAIN_SELECTOR);
        _assertFeeQuoterDestChainConfigsEqual(destChainConfigArgs[0].destChainConfig, gotConfig);
    }

    // Reverts

    function test_RevertWhen_InvalidChainFamilySelector() public {
        FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = _generateFeeQuoterDestChainConfigArgs();
        destChainConfigArgs[0].destChainConfig.chainFamilySelector = bytes4(uint32(1));

        vm.expectRevert(
            abi.encodeWithSelector(FeeQuoter.InvalidDestChainConfig.selector, destChainConfigArgs[0].destChainSelector)
        );
        s_feeQuoter.applyDestChainConfigUpdates(destChainConfigArgs);
    }
}
```

## Checklist for Writing Tests

- [ ] Create separate test file for each function: `ContractName.functionName.t.sol`
- [ ] Name test contract: `ContractName_functionName`
- [ ] Inherit from appropriate setup contract
- [ ] Write success test cases first (without "Success" suffix)
- [ ] Add `// Reverts` comment before failure test cases
- [ ] Use descriptive test function names following the pattern
- [ ] Include appropriate event emission tests
- [ ] Test access control where applicable
- [ ] Add fuzz tests for functions with numeric inputs
- [ ] Verify all state changes
- [ ] Use helper functions from setup files
- [ ] Follow variable naming conventions (s\_ prefix for state variables)
- [ ] Include edge cases and boundary conditions

## Common Errors to Avoid

1. **Don't include "Success" in success test names** - it's redundant
2. **Don't forget the `// Reverts` comment** - it's required for organization
3. **Don't test multiple unrelated behaviors in one test** - keep tests focused
4. **Don't duplicate setup code** - use setup files
5. **Don't use generic test names** - be descriptive about what's being tested
6. **Don't forget to test events** - use `vm.expectEmit()`
7. **Don't hardcode values** - use constants or helper functions
8. **Don't skip edge cases** - test boundaries and special values

## Test Case Generation Guidelines for AI Agents

### Analyzing a Function for Test Coverage

When presented with a new function to test, follow this systematic approach to achieve high branch coverage:

#### 1. Function Analysis Checklist

- [ ] **Input Parameters**: List all parameters and their types
- [ ] **Return Values**: Note what the function returns
- [ ] **State Changes**: Identify all state variables that are modified
- [ ] **External Calls**: List any external contract calls
- [ ] **Events**: Note all events that can be emitted
- [ ] **Modifiers**: List all modifiers (access control, state checks)
- [ ] **Conditionals**: Map all if/else branches
- [ ] **Loops**: Identify any loops and their bounds
- [ ] **Validations**: List all require/revert conditions

#### 2. Test Case Categories

##### A. Input Validation Tests

For each parameter, test:

- **Zero/Empty Values**: `address(0)`, `0`, empty arrays `new Type[](0)`, empty strings `""`
- **Boundary Values**:
  - Minimum valid values (e.g., `1` for non-zero checks)
  - Maximum values (e.g., `type(uint256).max`)
  - Values just above/below boundaries
- **Invalid Formats**: Wrong array lengths, invalid selectors, malformed data

##### B. State Transition Tests

- **Initial State**: Test with clean/default state
- **Modified State**: Test after previous operations
- **Edge State**: Test at capacity limits, after maximum operations
- **State Consistency**: Verify state remains valid after operations

##### C. Access Control Tests

For each modifier/role check:

- **Authorized Caller**: Test with correct permissions
- **Unauthorized Caller**: Test with `STRANGER` address
- **Role Transitions**: Test after role changes

##### D. Branch Coverage Tests

For each conditional:

- **True Branch**: Conditions that make the if statement true
- **False Branch**: Conditions that make the if statement false
- **Boundary Conditions**: Values at the exact decision point

##### E. Loop Tests

- **Zero Iterations**: Empty array or zero count
- **Single Iteration**: One element
- **Multiple Iterations**: 2-3 elements (typical case)
- **Maximum Iterations**: Test with maximum allowed elements

##### F. Integration Tests

- **External Call Success**: Mock successful responses
- **External Call Failure**: Mock reverts and error responses
- **Reentrancy**: Test reentrancy scenarios if applicable

### 3. Test Priority Matrix

| Scenario Type          | Priority     | When to Include                        |
| ---------------------- | ------------ | -------------------------------------- |
| Happy Path             | **CRITICAL** | Always - basic functionality must work |
| Zero/Empty Inputs      | **CRITICAL** | Always - common edge case              |
| Access Control         | **CRITICAL** | If function has modifiers              |
| Invalid Parameters     | **HIGH**     | Always - input validation              |
| Boundary Values        | **HIGH**     | For numeric inputs                     |
| State Transitions      | **HIGH**     | If function modifies state             |
| Events                 | **HIGH**     | If function emits events               |
| Max Values             | **MEDIUM**   | For arrays/loops                       |
| External Call Failures | **MEDIUM**   | If function makes external calls       |
| Gas Optimization       | **LOW**      | For gas-critical functions             |

### 4. Systematic Test Generation Process

```solidity
contract ContractName_functionName is ContractNameSetup {
    // Step 1: Basic success case - simplest valid inputs
    function test_BasicSuccess() public {
        // Test with minimal valid inputs
    }

    // Step 2: Complex success cases - real-world scenarios
    function test_WithMultipleElements() public {
        // Test with typical production data
    }

    // Step 3: Edge cases - boundaries and limits
    function test_MaximumValues() public {
        // Test at upper bounds
    }

    function test_MinimumValues() public {
        // Test at lower bounds
    }

    // Step 4: State variations
    function test_AfterMultipleOperations() public {
        // Test with modified state
    }

    // Step 5: Fuzz testing for numeric inputs
    function testFuzz_VariableInputs(uint256 amount) public {
        vm.assume(amount > 0 && amount < MAX_AMOUNT);
        // Test with random valid inputs
    }

    // Reverts

    // Step 6: Access control
    function test_RevertWhen_CallerUnauthorized() public {
        vm.startPrank(STRANGER);
        vm.expectRevert(/* error */);
        // Unauthorized call
    }

    // Step 7: Invalid inputs
    function test_RevertWhen_ZeroAddress() public {
        vm.expectRevert(/* error */);
        // Call with address(0)
    }

    // Step 8: State validation
    function test_RevertWhen_InvalidState() public {
        // Setup invalid state
        vm.expectRevert(/* error */);
        // Call that should fail
    }
}
```

### 5. Common Patterns by Function Type

#### Constructor Tests

- Valid deployment with different parameter combinations
- Zero address checks for each address parameter
- Invalid configuration combinations
- Event emissions during deployment
- Initial state verification

#### Setter Function Tests

- Setting new values (first time)
- Updating existing values
- Setting to zero/empty (removal)
- Setting same value twice (idempotency)
- Access control validation
- Event emission verification

#### Getter Function Tests

- Empty state returns
- Single element returns
- Multiple element returns
- Pagination tests (if applicable)
- View function gas optimization

#### Transfer/Payment Function Tests

- Zero amount transfers
- Minimum amount transfers (1 wei)
- Maximum balance transfers
- Insufficient balance scenarios
- Reentrancy protection
- Event emissions

#### Array/Mapping Function Tests

- Empty array operations
- Single element operations
- Multiple element operations
- Duplicate handling
- Out of bounds access
- Maximum array size limits

### 6. Branch Coverage Strategies

#### For Complex Conditionals

```solidity
// Function with: if (a && b || c)
function test_ConditionA_True_B_True() public { } // a=true, b=true, c=any
function test_ConditionA_True_B_False_C_True() public { } // a=true, b=false, c=true
function test_ConditionA_False_C_True() public { } // a=false, b=any, c=true
function test_AllConditionsFalse() public { } // a=false, b=any, c=false
```

#### For Nested Conditions

```solidity
// Test each path through nested ifs
function test_OuterTrue_InnerTrue() public { }
function test_OuterTrue_InnerFalse() public { }
function test_OuterFalse() public { }
```

### 7. Example: Comprehensive Test Generation

For a function like:

```solidity
function updateConfig(
    uint256 chainId,
    address admin,
    uint256[] memory limits
) external onlyOwner {
    require(chainId != 0, "Invalid chain");
    require(admin != address(0), "Invalid admin");
    require(limits.length > 0 && limits.length <= 10, "Invalid limits");

    for (uint i = 0; i < limits.length; i++) {
        require(limits[i] <= MAX_LIMIT, "Limit too high");
        configs[chainId].limits[i] = limits[i];
    }

    configs[chainId].admin = admin;
    emit ConfigUpdated(chainId, admin, limits);
}
```

Generate these tests:

1. `test_UpdateConfig_SingleLimit` - Basic success
2. `test_UpdateConfig_MultipleLimits` - Array with 3 elements
3. `test_UpdateConfig_MaximumLimits` - Array with 10 elements
4. `test_UpdateConfig_UpdateExisting` - Overwrite previous config
5. `testFuzz_UpdateConfig_VariableLimits` - Fuzz test limits
6. `test_RevertWhen_NotOwner` - Access control
7. `test_RevertWhen_ChainIdZero` - Invalid chain
8. `test_RevertWhen_AdminZeroAddress` - Invalid admin
9. `test_RevertWhen_EmptyLimits` - Empty array
10. `test_RevertWhen_TooManyLimits` - Array with 11 elements
11. `test_RevertWhen_LimitTooHigh` - Element exceeds MAX_LIMIT

### 8. Complex Test Scenarios

#### Testing Functions with Multiple External Calls

```solidity
function test_MultipleExternalCalls() public {
    // Mock all external calls
    vm.mockCall(
        address(tokenContract),
        abi.encodeWithSelector(IERC20.balanceOf.selector),
        abi.encode(1000e18)
    );
    vm.mockCall(
        address(priceOracle),
        abi.encodeWithSelector(IPriceOracle.getPrice.selector),
        abi.encode(2000e8)
    );

    // Execute function
    uint256 result = s_contract.calculateValue(tokenAddress);

    // Verify all calls were made
    vm.expectCall(address(tokenContract), abi.encodeWithSelector(IERC20.balanceOf.selector));
    vm.expectCall(address(priceOracle), abi.encodeWithSelector(IPriceOracle.getPrice.selector));

    assertEq(result, expectedValue);
}
```

#### Testing Complex State Machine Transitions

```solidity
function test_StateTransitions() public {
    // Initial state
    assertEq(s_contract.getState(), State.IDLE);

    // Transition to ACTIVE
    s_contract.activate();
    assertEq(s_contract.getState(), State.ACTIVE);

    // Perform operations only valid in ACTIVE state
    s_contract.process(data);

    // Transition to PAUSED
    s_contract.pause();
    assertEq(s_contract.getState(), State.PAUSED);

    // Verify operations fail in PAUSED state
    vm.expectRevert("InvalidState");
    s_contract.process(data);
}
```

#### Testing Time-Dependent Logic

```solidity
function test_TimeDependentBehavior() public {
    uint256 startTime = block.timestamp;

    // Set initial state
    s_contract.startProcess();

    // Fast forward time
    vm.warp(startTime + 1 hours);

    // Operation should succeed after delay
    s_contract.executeAfterDelay();

    // Test expiration
    vm.warp(startTime + 25 hours);
    vm.expectRevert("Expired");
    s_contract.executeAfterDelay();
}
```

#### Testing Complex Permission Hierarchies

```solidity
function test_PermissionHierarchy() public {
    address admin = makeAddr("admin");
    address operator = makeAddr("operator");
    address user = makeAddr("user");

    // Setup roles
    s_contract.grantRole(ADMIN_ROLE, admin);
    s_contract.grantRole(OPERATOR_ROLE, operator);

    // Test admin can do everything
    vm.startPrank(admin);
    s_contract.adminFunction();
    s_contract.operatorFunction();
    vm.stopPrank();

    // Test operator has limited access
    vm.startPrank(operator);
    s_contract.operatorFunction();
    vm.expectRevert("AccessControl: account is missing role");
    s_contract.adminFunction();
    vm.stopPrank();

    // Test user has no access
    vm.startPrank(user);
    vm.expectRevert("AccessControl: account is missing role");
    s_contract.operatorFunction();
    vm.stopPrank();
}
```

#### Testing Reentrancy Protection

```solidity
contract ReentrantAttacker {
    TargetContract target;

    function attack() external {
        target.withdraw();
    }

    receive() external payable {
        if (address(target).balance > 0) {
            target.withdraw(); // Attempt reentrancy
        }
    }
}

function test_ReentrantProtection() public {
    ReentrantAttacker attacker = new ReentrantAttacker();

    // Fund the contract
    deal(address(s_contract), 10 ether);

    // Attempt reentrancy attack
    vm.expectRevert("ReentrancyGuard: reentrant call");
    attacker.attack();

    // Verify funds are safe
    assertEq(address(s_contract).balance, 10 ether);
}
```

#### Testing Gas-Intensive Operations

```solidity
function test_GasIntensiveOperation() public {
    uint256 largeArraySize = 1000;
    uint256[] memory data = new uint256[](largeArraySize);

    // Fill array with data
    for (uint i = 0; i < largeArraySize; i++) {
        data[i] = i;
    }

    // Measure gas usage
    uint256 gasBefore = gasleft();
    s_contract.processLargeArray(data);
    uint256 gasUsed = gasBefore - gasleft();

    // Verify gas usage is within acceptable bounds
    assertLt(gasUsed, MAX_ACCEPTABLE_GAS);

    // Verify operation completed correctly
    assertEq(s_contract.getProcessedCount(), largeArraySize);
}
```

### 9. Special Testing Patterns

#### Testing Upgradeable Contracts

```solidity
function test_UpgradePattern() public {
    // Deploy V1
    ContractV1 v1 = new ContractV1();
    v1.initialize(INITIAL_VALUE);

    // Store some state
    v1.setValue(42);

    // Upgrade to V2
    ContractV2 v2 = new ContractV2();

    // Verify state preservation
    assertEq(v2.getValue(), 42);

    // Test new functionality
    v2.setNewFeature(true);
    assertTrue(v2.hasNewFeature());
}
```

#### Testing Circuit Breakers

```solidity
function test_CircuitBreaker() public {
    // Normal operation
    assertTrue(s_contract.processTransaction(100));

    // Trigger circuit breaker
    s_contract.emergencyStop();

    // Verify all operations are blocked
    vm.expectRevert("CircuitBreaker: stopped");
    s_contract.processTransaction(100);

    // Resume operations
    s_contract.resume();
    assertTrue(s_contract.processTransaction(100));
}
```

#### Testing Multi-Chain Scenarios

```solidity
function test_CrossChainMessage() public {
    uint64 destChainSelector = 12345;
    bytes memory message = abi.encode("data");

    // Mock chain validation
    vm.mockCall(
        address(router),
        abi.encodeWithSelector(IRouter.isChainSupported.selector, destChainSelector),
        abi.encode(true)
    );

    // Send cross-chain message
    vm.expectEmit();
    emit MessageSent(destChainSelector, message);

    bytes32 messageId = s_contract.sendMessage(destChainSelector, message);

    // Verify message state
    assertEq(s_contract.getMessageStatus(messageId), MessageStatus.SENT);
}
```

### 10. Verification Checklist

After generating tests, verify:

- [ ] All function parameters have validation tests
- [ ] All revert conditions are tested with specific error messages
- [ ] All events are verified with correct parameters using `vm.expectEmit()`
- [ ] All state changes are asserted with before/after comparisons
- [ ] Edge cases for numeric values are covered (0, 1, max-1, max)
- [ ] Access control is properly tested for each role
- [ ] External call scenarios are mocked for both success and failure
- [ ] Fuzz tests are added for appropriate inputs with proper bounds
- [ ] Time-dependent logic is tested with `vm.warp()`
- [ ] Gas consumption is within acceptable limits for gas-critical functions
- [ ] Reentrancy protection is verified where applicable
- [ ] State consistency is maintained across all operations

## Final Notes for AI Agents

When writing tests:

1. Start by analyzing the function systematically using the checklist
2. Generate test cases for each category relevant to the function
3. Prioritize critical paths and common failure modes
4. Ensure each test has a single, clear purpose
5. Use descriptive names that indicate the scenario being tested
6. Follow the established patterns and conventions exactly
7. Aim for complete branch coverage, not just line coverage
8. Consider both technical correctness and business logic validation
