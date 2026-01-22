// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.20;

/**
 * @title IPolicyEngine
 * @dev Interface for the policy engine.
 */
interface IPolicyEngine {
  /// @notice Error emitted when the target is not attached to the policy engine.
  error TargetNotAttached(address target);
  /// @notice Error emitted when the target is already attached to the policy engine.
  error TargetAlreadyAttached(address target);
  /// @notice Error emitted when the policy engine is missing or not present.
  error PolicyEngineUndefined();
  /// @notice Error emitted when the PolicyEngine run has been rejected by one of the polices.
  error PolicyRunRejected(address policy, string rejectReason, Payload payload);
  /// @notice Error emitted when a policy mapper results in an error.
  error PolicyMapperError(address policy, bytes errorReason, Payload payload);
  /// @notice Error emitted when an individual policy is rejecting a transaction.
  error PolicyRejected(string rejectReason);
  /// @notice Error emitted when the PolicyEngine run encounters an error while executing one of the policies.
  error PolicyRunError(address policy, bytes errorReason, Payload payload);
  /// @notice Error emitted when a policy run is unauthorized.
  error PolicyRunUnauthorizedError(address account);
  /// @notice Error emitted when a policy postRun results in an error.
  error PolicyPostRunError(address policy, bytes errorReason, Payload payload);
  /// @notice Error emitted when a policy extractor is run with an unsupported selector.
  error UnsupportedSelector(bytes4 selector);
  /// @notice Error emitted when a policy action results in an error.
  error PolicyActionError(address policy, bytes errorReason);
  /// @notice Error emitted when a policy configuration change results in an error.
  error PolicyConfigurationError(address policy, bytes errorReason);
  /// @notice Error emitted when a policy configuration version does not match the expected version.
  error PolicyConfigurationVersionError(address policy, uint256 expectedVersion, uint256 actualVersion);
  /// @notice Error emitted when an extraction of parameters results in an error.
  error ExtractorError(address extractor, bytes errorReason, Payload payload);

  /**
   * @notice Emitted when a target contract has attached to the policy engine.
   * @param target The target contract.
   */
  event TargetAttached(address indexed target);

  /**
   * @notice Emitted when a target contract has detached from the policy engine.
   * @param target The target contract.
   */
  event TargetDetached(address indexed target);

  /**
   * @notice Emitted when a policy configuration is performed.
   * @param policy The address of the policy.
   * @param configSelector The selector of the configuration function.
   * @param configVersion The version of the configuration.
   * @param configData The data of the configuration.
   */
  event PolicyConfigured(
    address indexed policy, uint256 indexed configVersion, bytes4 indexed configSelector, bytes configData
  );

  /**
   * @notice Emitted when a policy engine run has completed successfully.
   * @param sender The sender of the transaction.
   * @param target The target contract that invoked the method.
   * @param selector The selector of the method invoked on the target.
   * @param extractedParameters The parameters extracted from the payload for policy evaluation.
   * @param context Additional context data from the payload.
   */
  event PolicyRunComplete(
    address indexed sender,
    address indexed target,
    bytes4 indexed selector,
    Parameter[] extractedParameters,
    bytes context
  );

  /**
   * @notice Emitted when a policy is added to the policy engine.
   * @param target The address of the target contract for which the policy was configured.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   * @param position The position of the policy in the policy chain.
   * @param policyParameterNames The parameter names for the policy.
   */
  event PolicyAdded(
    address indexed target, bytes4 indexed selector, address policy, uint256 position, bytes32[] policyParameterNames
  );

  /**
   * @notice Emitted when a policy is added to the policy engine at a specific position.
   * @param target The address of the target contract for which the policy was configured.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   * @param position The position of the policy in the policy chain.
   * @param policyParameterNames The parameter names for the policy.
   * @param policies The complete ordered array of all policy addresses after the insertion.
   */
  event PolicyAddedAt(
    address indexed target,
    bytes4 indexed selector,
    address policy,
    uint256 position,
    bytes32[] policyParameterNames,
    address[] policies
  );

  /**
   * @notice Emitted when a policy is removed from the policy engine.
   * @param target The address of the target contract for which the policy was configured.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   */
  event PolicyRemoved(address indexed target, bytes4 indexed selector, address policy);

  /**
   * @notice Emitted when an extractor is set for a selector.
   * @param selector The selector.
   * @param extractor The extractor address.
   */
  event ExtractorSet(bytes4 indexed selector, address indexed extractor);

  /**
   * @notice Emitted when a policy mapper is set for a policy.
   * @param policy The policy address.
   * @param mapper The mapper address.
   */
  event PolicyMapperSet(address indexed policy, address indexed mapper);

  /**
   * @notice Emitted when policy parameters are set for a policy.
   * @param policy The policy address.
   * @param parameters The parameters for the policy.
   */
  event PolicyParametersSet(address indexed policy, bytes[] parameters);

  /**
   * @notice Emitted when the default policy action rule is set for the policy engine.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  event DefaultPolicyAllowSet(bool defaultAllow);

  /**
   * @notice Emitted when the default policy allow rule for a target is set.
   * @param target The target contract.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  event TargetDefaultPolicyAllowSet(address indexed target, bool defaultAllow);

  /**
   * @notice The PolicyResult enum represents the possible types of success results of a policy run. When a policy
   * should reject a transaction, it MUST revert using the `PolicyReject` error with a descriptive reject message.
   * @param None No specific policy result, typically used as a default or uninitialized state.
   * @param Allowed The policy allowed the run.
   * @param Continue The policy did not reject the run and processing should continue to the next policy.
   */
  enum PolicyResult {
    None,
    Allowed,
    Continue
  }

  /**
   * @notice The Payload struct combines the components on which policies operate.
   * @param selector The selector of the method being invoked on the target.
   * @param sender The sender of the transaction.
   * @param data The original calldata of the invoked transaction.
   * @param context Additional information or authorization to perform the operation.
   */
  struct Payload {
    bytes4 selector;
    address sender;
    bytes data;
    bytes context;
  }

  /**
   * @notice The Parameter struct contains the data of the parameters sent to policies.
   * @param name The name of the parameter.
   * @param value The value of the parameter.
   */
  struct Parameter {
    bytes32 name;
    bytes value;
  }

  /**
   * @notice Returns the type and version of the policy engine.
   * @return A string representing the type and version of the policy engine.
   */
  function typeAndVersion() external pure returns (string memory);

  /**
   * @notice Attaches the calling contract to the policy engine.
   */
  function attach() external;

  /**
   * @notice Detaches the calling contract from the policy engine.
   */
  function detach() external;

  /**
   * @notice Assigns an extractor to the specified selector, enabling policies to utilize it for parameter extraction.
   * @param selector The selector of the policy.
   * @param extractor The extractor address.
   */
  function setExtractor(
    bytes4 selector,
    address extractor
  ) external;

  /**
   * @notice Assigns an extractor to the specified selectors, enabling policies to utilize it for parameter extraction.
   * @param selectors The selectors of the policies.
   * @param extractor The extractor address.
   */
  function setExtractors(
    bytes4[] calldata selectors,
    address extractor
  ) external;

  /**
   * @notice Gets the extractor for a given selector.
   * @param selector The selector.
   * @return The extractor for the selector.
   */
  function getExtractor(
    bytes4 selector
  ) external view returns (address);

  /**
   * @notice Sets the custom policy parameter mapper for a policy.
   * @param policy The policy address.
   * @param mapper The mapper address, address(0) to use the default mapper.
   */
  function setPolicyMapper(
    address policy,
    address mapper
  ) external;

  /**
   * @notice Gets the policy parameter mapper for a given policy.
   * @param policy The policy address.
   * @return The custom policy parameter mapper for the policy, address(0) if the policy uses the default mapper.
   */
  function getPolicyMapper(
    address policy
  ) external view returns (address);

  /**
   * @notice Adds a policy to the policy engine.
   *
   * - Policy MUST be added to the end of the current policy list.
   *
   * @param target The address of the target contract for which the policy apply.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   * @param policyParameterNames The parameter names for the policy.
   */
  function addPolicy(
    address target,
    bytes4 selector,
    address policy,
    bytes32[] calldata policyParameterNames
  ) external;

  /**
   * @notice Adds a policy to the policy engine at a specific position.
   *
   * @param target The address of the target contract for which the policy apply.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   * @param policyParameterNames The parameter names for the policy.
   * @param position The position to add the policy at.
   */
  function addPolicyAt(
    address target,
    bytes4 selector,
    address policy,
    bytes32[] calldata policyParameterNames,
    uint256 position
  ) external;

  /**
   * @notice Removes a policy from the policy engine.
   * @param target The address of the target contract for which the policy was configured.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   */
  function removePolicy(
    address target,
    bytes4 selector,
    address policy
  ) external;

  /**
   * @notice Gets the policies for a given selector and target.
   *
   * - MUST return the policies in the order they will execute.
   * - MUST return an empty array if no policies are found.
   *
   * @param selector The selector of the policy.
   * @param target The address of the target contract for which the policies are configured.
   * @return The policies for the selector and target.
   */
  function getPolicies(
    address target,
    bytes4 selector
  ) external view returns (address[] memory);

  /**
   * @notice Sets the configuration for a policy.
   * @param policy The address of the policy to configure.
   * @param configVersion The version of the configuration.
   * @param configSelector The selector of the configuration function.
   * @param configData The calldata for the configuration function.
   */
  function setPolicyConfiguration(
    address policy,
    uint256 configVersion,
    bytes4 configSelector,
    bytes calldata configData
  ) external;

  /**
   * @notice Gets the current configuration version for a policy.
   * @param policy The address of the policy.
   * @return The current configuration version for the policy.
   */
  function getPolicyConfigVersion(
    address policy
  ) external view returns (uint256);

  /**
   * @notice Sets whether to allow or reject the transaction if no policy explicitly returns an Allow or a Reject.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  function setDefaultPolicyAllow(
    bool defaultAllow
  ) external;

  /**
   * @notice Sets whether to allow or reject the transaction if no policy explicitly returns an Allow or a Reject
   * for a specific target.
   * @param target The address of the target contract.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  function setTargetDefaultPolicyAllow(
    address target,
    bool defaultAllow
  ) external;

  /**
   * @notice Runs the policies for a given payload for offchain pre-validation. MUST revert on policy rejection/failure.
   * @param payload The payload to run the policies on.
   */
  function check(
    Payload calldata payload
  ) external view;

  /**
   * @notice Runs the policies for a given operation payload.
   *
   * - MUST revert on policy rejection/failure.
   * - MUST revert if the target contract that invoked the method is not allowed. Target contract address is
   *   obtained from the msg.sender global variable.
   * - MUST execute policies in the order they were added or that were specified using `addPolicyAt`.
   *
   * @param payload The payload to run the policies on.
   */
  function run(
    Payload calldata payload
  ) external;
}
