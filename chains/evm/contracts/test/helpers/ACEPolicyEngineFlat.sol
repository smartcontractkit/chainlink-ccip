// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.20;

// node_modules/@openzeppelin/contracts/access/IAccessControl.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/IAccessControl.sol)

/**
 * @dev External interface of AccessControl declared to support ERC165 detection.
 */
interface IAccessControl {
    /**
     * @dev The `account` is missing a role.
     */
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    /**
     * @dev The caller of a function is not the expected one.
     *
     * NOTE: Don't confuse with {AccessControlUnauthorizedAccount}.
     */
    error AccessControlBadConfirmation();

    /**
     * @dev Emitted when `newAdminRole` is set as ``role``'s admin role, replacing `previousAdminRole`
     *
     * `DEFAULT_ADMIN_ROLE` is the starting admin for all roles, despite
     * {RoleAdminChanged} not being emitted signaling this.
     */
    event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole);

    /**
     * @dev Emitted when `account` is granted `role`.
     *
     * `sender` is the account that originated the contract call, an admin role
     * bearer except when using {AccessControl-_setupRole}.
     */
    event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Emitted when `account` is revoked `role`.
     *
     * `sender` is the account that originated the contract call:
     *   - if using `revokeRole`, it is the admin role bearer
     *   - if using `renounceRole`, it is the role bearer (i.e. `account`)
     */
    event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) external view returns (bool);

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {AccessControl-_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) external view returns (bytes32);

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function grantRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function revokeRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been granted `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `callerConfirmation`.
     */
    function renounceRole(bytes32 role, address callerConfirmation) external;
}

// node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/IERC165.sol)

/**
 * @dev Interface of the ERC165 standard, as defined in the
 * https://eips.ethereum.org/EIPS/eip-165[EIP].
 *
 * Implementers can declare support of contract interfaces, which can then be
 * queried by others ({ERC165Checker}).
 *
 * For an implementation, see {ERC165}.
 */
interface IERC165 {
    /**
     * @dev Returns true if this contract implements the interface defined by
     * `interfaceId`. See the corresponding
     * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section]
     * to learn more about how these ids are created.
     *
     * This function call must use less than 30 000 gas.
     */
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}

// src/interfaces/IPolicyEngine.sol

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
  function setExtractor(bytes4 selector, address extractor) external;

  /**
   * @notice Assigns an extractor to the specified selectors, enabling policies to utilize it for parameter extraction.
   * @param selectors The selectors of the policies.
   * @param extractor The extractor address.
   */
  function setExtractors(bytes4[] calldata selectors, address extractor) external;

  /**
   * @notice Gets the extractor for a given selector.
   * @param selector The selector.
   * @return The extractor for the selector.
   */
  function getExtractor(bytes4 selector) external view returns (address);

  /**
   * @notice Sets the custom policy parameter mapper for a policy.
   * @param policy The policy address.
   * @param mapper The mapper address, address(0) to use the default mapper.
   */
  function setPolicyMapper(address policy, address mapper) external;

  /**
   * @notice Gets the policy parameter mapper for a given policy.
   * @param policy The policy address.
   * @return The custom policy parameter mapper for the policy, address(0) if the policy uses the default mapper.
   */
  function getPolicyMapper(address policy) external view returns (address);

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
  function addPolicy(address target, bytes4 selector, address policy, bytes32[] calldata policyParameterNames) external;

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
  )
    external;

  /**
   * @notice Removes a policy from the policy engine.
   * @param target The address of the target contract for which the policy was configured.
   * @param selector The selector of the policy.
   * @param policy The policy address.
   */
  function removePolicy(address target, bytes4 selector, address policy) external;

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
  function getPolicies(address target, bytes4 selector) external view returns (address[] memory);

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
  )
    external;

  /**
   * @notice Gets the current configuration version for a policy.
   * @param policy The address of the policy.
   * @return The current configuration version for the policy.
   */
  function getPolicyConfigVersion(address policy) external view returns (uint256);

  /**
   * @notice Sets whether to allow or reject the transaction if no policy explicitly returns an Allow or a Reject.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  function setDefaultPolicyAllow(bool defaultAllow) external;

  /**
   * @notice Sets whether to allow or reject the transaction if no policy explicitly returns an Allow or a Reject
   * for a specific target.
   * @param target The address of the target contract.
   * @param defaultAllow Indicates whether to allow or reject a transaction if no policy explicitly returns an Allow
   * or a Reject. True to allow, false to reject.
   */
  function setTargetDefaultPolicyAllow(address target, bool defaultAllow) external;

  /**
   * @notice Runs the policies for a given payload for offchain pre-validation. MUST revert on policy rejection/failure.
   * @param payload The payload to run the policies on.
   */
  function check(Payload calldata payload) external view;

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
  function run(Payload calldata payload) external;
}

// node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol

// OpenZeppelin Contracts (last updated v5.0.0) (proxy/utils/Initializable.sol)

/**
 * @dev This is a base contract to aid in writing upgradeable contracts, or any kind of contract that will be deployed
 * behind a proxy. Since proxied contracts do not make use of a constructor, it's common to move constructor logic to an
 * external initializer function, usually called `initialize`. It then becomes necessary to protect this initializer
 * function so it can only be called once. The {initializer} modifier provided by this contract will have this effect.
 *
 * The initialization functions use a version number. Once a version number is used, it is consumed and cannot be
 * reused. This mechanism prevents re-execution of each "step" but allows the creation of new initialization steps in
 * case an upgrade adds a module that needs to be initialized.
 *
 * For example:
 *
 * [.hljs-theme-light.nopadding]
 * ```solidity
 * contract MyToken is ERC20Upgradeable {
 *     function initialize() initializer public {
 *         __ERC20_init("MyToken", "MTK");
 *     }
 * }
 *
 * contract MyTokenV2 is MyToken, ERC20PermitUpgradeable {
 *     function initializeV2() reinitializer(2) public {
 *         __ERC20Permit_init("MyToken");
 *     }
 * }
 * ```
 *
 * TIP: To avoid leaving the proxy in an uninitialized state, the initializer function should be called as early as
 * possible by providing the encoded function call as the `_data` argument to {ERC1967Proxy-constructor}.
 *
 * CAUTION: When used with inheritance, manual care must be taken to not invoke a parent initializer twice, or to ensure
 * that all initializers are idempotent. This is not verified automatically as constructors are by Solidity.
 *
 * [CAUTION]
 * ====
 * Avoid leaving a contract uninitialized.
 *
 * An uninitialized contract can be taken over by an attacker. This applies to both a proxy and its implementation
 * contract, which may impact the proxy. To prevent the implementation contract from being used, you should invoke
 * the {_disableInitializers} function in the constructor to automatically lock it when it is deployed:
 *
 * [.hljs-theme-light.nopadding]
 * ```
 * /// @custom:oz-upgrades-unsafe-allow constructor
 * constructor() {
 *     _disableInitializers();
 * }
 * ```
 * ====
 */
abstract contract Initializable {
    /**
     * @dev Storage of the initializable contract.
     *
     * It's implemented on a custom ERC-7201 namespace to reduce the risk of storage collisions
     * when using with upgradeable contracts.
     *
     * @custom:storage-location erc7201:openzeppelin.storage.Initializable
     */
    struct InitializableStorage {
        /**
         * @dev Indicates that the contract has been initialized.
         */
        uint64 _initialized;
        /**
         * @dev Indicates that the contract is in the process of being initialized.
         */
        bool _initializing;
    }

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.Initializable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant INITIALIZABLE_STORAGE = 0xf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00;

    /**
     * @dev The contract is already initialized.
     */
    error InvalidInitialization();

    /**
     * @dev The contract is not initializing.
     */
    error NotInitializing();

    /**
     * @dev Triggered when the contract has been initialized or reinitialized.
     */
    event Initialized(uint64 version);

    /**
     * @dev A modifier that defines a protected initializer function that can be invoked at most once. In its scope,
     * `onlyInitializing` functions can be used to initialize parent contracts.
     *
     * Similar to `reinitializer(1)`, except that in the context of a constructor an `initializer` may be invoked any
     * number of times. This behavior in the constructor can be useful during testing and is not expected to be used in
     * production.
     *
     * Emits an {Initialized} event.
     */
    modifier initializer() {
        // solhint-disable-next-line var-name-mixedcase
        InitializableStorage storage $ = _getInitializableStorage();

        // Cache values to avoid duplicated sloads
        bool isTopLevelCall = !$._initializing;
        uint64 initialized = $._initialized;

        // Allowed calls:
        // - initialSetup: the contract is not in the initializing state and no previous version was
        //                 initialized
        // - construction: the contract is initialized at version 1 (no reininitialization) and the
        //                 current contract is just being deployed
        bool initialSetup = initialized == 0 && isTopLevelCall;
        bool construction = initialized == 1 && address(this).code.length == 0;

        if (!initialSetup && !construction) {
            revert InvalidInitialization();
        }
        $._initialized = 1;
        if (isTopLevelCall) {
            $._initializing = true;
        }
        _;
        if (isTopLevelCall) {
            $._initializing = false;
            emit Initialized(1);
        }
    }

    /**
     * @dev A modifier that defines a protected reinitializer function that can be invoked at most once, and only if the
     * contract hasn't been initialized to a greater version before. In its scope, `onlyInitializing` functions can be
     * used to initialize parent contracts.
     *
     * A reinitializer may be used after the original initialization step. This is essential to configure modules that
     * are added through upgrades and that require initialization.
     *
     * When `version` is 1, this modifier is similar to `initializer`, except that functions marked with `reinitializer`
     * cannot be nested. If one is invoked in the context of another, execution will revert.
     *
     * Note that versions can jump in increments greater than 1; this implies that if multiple reinitializers coexist in
     * a contract, executing them in the right order is up to the developer or operator.
     *
     * WARNING: Setting the version to 2**64 - 1 will prevent any future reinitialization.
     *
     * Emits an {Initialized} event.
     */
    modifier reinitializer(uint64 version) {
        // solhint-disable-next-line var-name-mixedcase
        InitializableStorage storage $ = _getInitializableStorage();

        if ($._initializing || $._initialized >= version) {
            revert InvalidInitialization();
        }
        $._initialized = version;
        $._initializing = true;
        _;
        $._initializing = false;
        emit Initialized(version);
    }

    /**
     * @dev Modifier to protect an initialization function so that it can only be invoked by functions with the
     * {initializer} and {reinitializer} modifiers, directly or indirectly.
     */
    modifier onlyInitializing() {
        _checkInitializing();
        _;
    }

    /**
     * @dev Reverts if the contract is not in an initializing state. See {onlyInitializing}.
     */
    function _checkInitializing() internal view virtual {
        if (!_isInitializing()) {
            revert NotInitializing();
        }
    }

    /**
     * @dev Locks the contract, preventing any future reinitialization. This cannot be part of an initializer call.
     * Calling this in the constructor of a contract will prevent that contract from being initialized or reinitialized
     * to any version. It is recommended to use this to lock implementation contracts that are designed to be called
     * through proxies.
     *
     * Emits an {Initialized} event the first time it is successfully executed.
     */
    function _disableInitializers() internal virtual {
        // solhint-disable-next-line var-name-mixedcase
        InitializableStorage storage $ = _getInitializableStorage();

        if ($._initializing) {
            revert InvalidInitialization();
        }
        if ($._initialized != type(uint64).max) {
            $._initialized = type(uint64).max;
            emit Initialized(type(uint64).max);
        }
    }

    /**
     * @dev Returns the highest version that has been initialized. See {reinitializer}.
     */
    function _getInitializedVersion() internal view returns (uint64) {
        return _getInitializableStorage()._initialized;
    }

    /**
     * @dev Returns `true` if the contract is currently initializing. See {onlyInitializing}.
     */
    function _isInitializing() internal view returns (bool) {
        return _getInitializableStorage()._initializing;
    }

    /**
     * @dev Returns a pointer to the storage namespace.
     */
    // solhint-disable-next-line var-name-mixedcase
    function _getInitializableStorage() private pure returns (InitializableStorage storage $) {
        assembly {
            $.slot := INITIALIZABLE_STORAGE
        }
    }
}

// node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol

// OpenZeppelin Contracts (last updated v5.0.1) (utils/Context.sol)

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract ContextUpgradeable is Initializable {
    function __Context_init() internal onlyInitializing {
    }

    function __Context_init_unchained() internal onlyInitializing {
    }
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    function _contextSuffixLength() internal view virtual returns (uint256) {
        return 0;
    }
}

// src/interfaces/IExtractor.sol

/**
 * @title IExtractor
 * @dev Interface for extracting parameters from a payload.
 */
interface IExtractor {
  /**
   * @notice Returns the type and version of the extractor.
   * @return A string representing the type and version of the extractor.
   */
  function typeAndVersion() external pure returns (string memory);

  /**
   * @notice Extracts parameters from a payload.
   * @param payload The payload to extract parameters from.
   * @return The extracted parameters.
   */
  function extract(IPolicyEngine.Payload calldata payload) external view returns (IPolicyEngine.Parameter[] memory);
}

// src/interfaces/IMapper.sol

/**
 * @title IMapper
 * @dev Interface for mapping extracted parameters to a list of policy parameters.
 */
interface IMapper {
  /**
   * @notice Returns the type and version of the mapper.
   * @return A string representing the type and version of the mapper.
   */
  function typeAndVersion() external pure returns (string memory);

  /**
   * @notice Maps extracted parameters to a list of policy parameters.
   * @param extractedParameters The extracted parameters.
   * @return The mapped parameters.
   */
  function map(IPolicyEngine.Parameter[] calldata extractedParameters) external view returns (bytes[] memory);
}

// node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/ERC165.sol)

/**
 * @dev Implementation of the {IERC165} interface.
 *
 * Contracts that want to implement ERC165 should inherit from this contract and override {supportsInterface} to check
 * for the additional interface id that will be supported. For example:
 *
 * ```solidity
 * function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
 *     return interfaceId == type(MyInterface).interfaceId || super.supportsInterface(interfaceId);
 * }
 * ```
 */
abstract contract ERC165Upgradeable is Initializable, IERC165 {
    function __ERC165_init() internal onlyInitializing {
    }

    function __ERC165_init_unchained() internal onlyInitializing {
    }
    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual returns (bool) {
        return interfaceId == type(IERC165).interfaceId;
    }
}

// src/interfaces/IPolicy.sol

/**
 * @title IPolicy
 * @dev Interface for running a policy.
 */
interface IPolicy is IERC165 {
  /**
   * @notice Returns the type and version of the policy.
   * @return A string representing the type and version of the policy.
   */
  function typeAndVersion() external pure returns (string memory);

  /**
   * @notice Hook called upon installation of the policy.
   * @param selector The selector of the policy.
   */
  function onInstall(bytes4 selector) external;

  /**
   * @notice Hook called upon uninstallation of the policy.
   * @param selector The selector of the policy.
   */
  function onUninstall(bytes4 selector) external;

  /**
   * @notice Runs the policy.
   * @param caller The address of the account which is calling the subject protected by the policy engine.
   * @param subject The address of the contract which is being protected by the policy engine.
   * @param selector The selector of the method being called on the protected contract.
   * @param parameters The parameters to use for running the policy.
   * @param context Additional information or authorization to perform the operation.
   * @return The result of running the policy.
   */
  function run(
    address caller,
    address subject,
    bytes4 selector,
    bytes[] calldata parameters,
    bytes calldata context
  )
    external
    view
    returns (IPolicyEngine.PolicyResult);

  /**
   * @notice Runs after the policy check if the check was successful, and MAY mutate state. State mutations SHOULD
   * consider state relative to the target, as the policy MAY be shared across multiple targets.
   * @param caller The address of the account which is calling the subject protected by the policy engine.
   * @param subject The address of the contract which is being protected by the policy engine.
   * @param selector The selector of the method being called on the protected contract.
   * @param parameters The parameters to use for running the policy.
   * @param context Additional information or authorization to perform the operation.
   */
  function postRun(
    address caller,
    address subject,
    bytes4 selector,
    bytes[] calldata parameters,
    bytes calldata context
  )
    external;
}

// node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/Ownable.sol)

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * The initial owner is set to the address provided by the deployer. This can
 * later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract OwnableUpgradeable is Initializable, ContextUpgradeable {
    /// @custom:storage-location erc7201:openzeppelin.storage.Ownable
    struct OwnableStorage {
        address _owner;
    }

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.Ownable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant OwnableStorageLocation = 0x9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300;

    function _getOwnableStorage() private pure returns (OwnableStorage storage $) {
        assembly {
            $.slot := OwnableStorageLocation
        }
    }

    /**
     * @dev The caller account is not authorized to perform an operation.
     */
    error OwnableUnauthorizedAccount(address account);

    /**
     * @dev The owner is not a valid owner account. (eg. `address(0)`)
     */
    error OwnableInvalidOwner(address owner);

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the address provided by the deployer as the initial owner.
     */
    function __Ownable_init(address initialOwner) internal onlyInitializing {
        __Ownable_init_unchained(initialOwner);
    }

    function __Ownable_init_unchained(address initialOwner) internal onlyInitializing {
        if (initialOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(initialOwner);
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        OwnableStorage storage $ = _getOwnableStorage();
        return $._owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        if (owner() != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        if (newOwner == address(0)) {
            revert OwnableInvalidOwner(address(0));
        }
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        OwnableStorage storage $ = _getOwnableStorage();
        address oldOwner = $._owner;
        $._owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}

// node_modules/@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/AccessControl.sol)

/**
 * @dev Contract module that allows children to implement role-based access
 * control mechanisms. This is a lightweight version that doesn't allow enumerating role
 * members except through off-chain means by accessing the contract event logs. Some
 * applications may benefit from on-chain enumerability, for those cases see
 * {AccessControlEnumerable}.
 *
 * Roles are referred to by their `bytes32` identifier. These should be exposed
 * in the external API and be unique. The best way to achieve this is by
 * using `public constant` hash digests:
 *
 * ```solidity
 * bytes32 public constant MY_ROLE = keccak256("MY_ROLE");
 * ```
 *
 * Roles can be used to represent a set of permissions. To restrict access to a
 * function call, use {hasRole}:
 *
 * ```solidity
 * function foo() public {
 *     require(hasRole(MY_ROLE, msg.sender));
 *     ...
 * }
 * ```
 *
 * Roles can be granted and revoked dynamically via the {grantRole} and
 * {revokeRole} functions. Each role has an associated admin role, and only
 * accounts that have a role's admin role can call {grantRole} and {revokeRole}.
 *
 * By default, the admin role for all roles is `DEFAULT_ADMIN_ROLE`, which means
 * that only accounts with this role will be able to grant or revoke other
 * roles. More complex role relationships can be created by using
 * {_setRoleAdmin}.
 *
 * WARNING: The `DEFAULT_ADMIN_ROLE` is also its own admin: it has permission to
 * grant and revoke this role. Extra precautions should be taken to secure
 * accounts that have been granted it. We recommend using {AccessControlDefaultAdminRules}
 * to enforce additional security measures for this role.
 */
abstract contract AccessControlUpgradeable is Initializable, ContextUpgradeable, IAccessControl, ERC165Upgradeable {
    struct RoleData {
        mapping(address account => bool) hasRole;
        bytes32 adminRole;
    }

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    /// @custom:storage-location erc7201:openzeppelin.storage.AccessControl
    struct AccessControlStorage {
        mapping(bytes32 role => RoleData) _roles;
    }

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.AccessControl")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant AccessControlStorageLocation = 0x02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800;

    function _getAccessControlStorage() private pure returns (AccessControlStorage storage $) {
        assembly {
            $.slot := AccessControlStorageLocation
        }
    }

    /**
     * @dev Modifier that checks that an account has a specific role. Reverts
     * with an {AccessControlUnauthorizedAccount} error including the required role.
     */
    modifier onlyRole(bytes32 role) {
        _checkRole(role);
        _;
    }

    function __AccessControl_init() internal onlyInitializing {
    }

    function __AccessControl_init_unchained() internal onlyInitializing {
    }
    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IAccessControl).interfaceId || super.supportsInterface(interfaceId);
    }

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) public view virtual returns (bool) {
        AccessControlStorage storage $ = _getAccessControlStorage();
        return $._roles[role].hasRole[account];
    }

    /**
     * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `_msgSender()`
     * is missing `role`. Overriding this function changes the behavior of the {onlyRole} modifier.
     */
    function _checkRole(bytes32 role) internal view virtual {
        _checkRole(role, _msgSender());
    }

    /**
     * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `account`
     * is missing `role`.
     */
    function _checkRole(bytes32 role, address account) internal view virtual {
        if (!hasRole(role, account)) {
            revert AccessControlUnauthorizedAccount(account, role);
        }
    }

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) public view virtual returns (bytes32) {
        AccessControlStorage storage $ = _getAccessControlStorage();
        return $._roles[role].adminRole;
    }

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleGranted} event.
     */
    function grantRole(bytes32 role, address account) public virtual onlyRole(getRoleAdmin(role)) {
        _grantRole(role, account);
    }

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleRevoked} event.
     */
    function revokeRole(bytes32 role, address account) public virtual onlyRole(getRoleAdmin(role)) {
        _revokeRole(role, account);
    }

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been revoked `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `callerConfirmation`.
     *
     * May emit a {RoleRevoked} event.
     */
    function renounceRole(bytes32 role, address callerConfirmation) public virtual {
        if (callerConfirmation != _msgSender()) {
            revert AccessControlBadConfirmation();
        }

        _revokeRole(role, callerConfirmation);
    }

    /**
     * @dev Sets `adminRole` as ``role``'s admin role.
     *
     * Emits a {RoleAdminChanged} event.
     */
    function _setRoleAdmin(bytes32 role, bytes32 adminRole) internal virtual {
        AccessControlStorage storage $ = _getAccessControlStorage();
        bytes32 previousAdminRole = getRoleAdmin(role);
        $._roles[role].adminRole = adminRole;
        emit RoleAdminChanged(role, previousAdminRole, adminRole);
    }

    /**
     * @dev Attempts to grant `role` to `account` and returns a boolean indicating if `role` was granted.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleGranted} event.
     */
    function _grantRole(bytes32 role, address account) internal virtual returns (bool) {
        AccessControlStorage storage $ = _getAccessControlStorage();
        if (!hasRole(role, account)) {
            $._roles[role].hasRole[account] = true;
            emit RoleGranted(role, account, _msgSender());
            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Attempts to revoke `role` to `account` and returns a boolean indicating if `role` was revoked.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleRevoked} event.
     */
    function _revokeRole(bytes32 role, address account) internal virtual returns (bool) {
        AccessControlStorage storage $ = _getAccessControlStorage();
        if (hasRole(role, account)) {
            $._roles[role].hasRole[account] = false;
            emit RoleRevoked(role, account, _msgSender());
            return true;
        } else {
            return false;
        }
    }
}

// src/core/Policy.sol

abstract contract Policy is Initializable, OwnableUpgradeable, ERC165Upgradeable, IPolicy {
  error Unauthorized();
  error InvalidParameters(string reason);

  /// @custom:storage-location erc7201:chainlink.ace.Policy
  struct PolicyStorage {
    address policyEngine;
  }

  // keccak256(abi.encode(uint256(keccak256("chainlink.ace.Policy")) - 1)) &
  // ~bytes32(uint256(0xff))
  // solhint-disable-next-line const-name-snakecase
  bytes32 private constant PolicyStorageLocation = 0x88b18ed68be9f5af7a0aa0e9a55256b17a6bcc168c9c257d2c5556789ebee900;

  function _getPolicyStorage() private pure returns (PolicyStorage storage $) {
    // solhint-disable-next-line no-inline-assembly
    assembly {
      $.slot := PolicyStorageLocation
    }
  }

  constructor() {
    _disableInitializers();
  }

  modifier onlyPolicyEngine() {
    if (msg.sender != _getPolicyStorage().policyEngine) {
      revert Unauthorized();
    }
    _;
  }

  /**
   * @notice Initializes the policy contract.
   * @dev This function MUST be called immediately after deployment (or via proxy initializer),
   * setting up the policy with the address of the policy engine, initial ownership, and
   * any configuration parameters required by the policy.
   * @param policyEngine The address of the policy engine that will invoke this policy.
   * @param initialOwner The address that will be assigned as the initial owner of the policy contract.
   * @param configParams Arbitrary encoded configuration parameters to initialize the policy logic.
   */
  function initialize(
    address policyEngine,
    address initialOwner,
    bytes calldata configParams
  )
    public
    virtual
    initializer
  {
    __Policy_init(policyEngine, initialOwner);
    configure(configParams);
  }

  // solhint-disable-next-line no-empty-blocks
  function configure(bytes calldata parameters) internal virtual onlyInitializing {}

  function __Policy_init(address policyEngine, address initialOwner) internal onlyInitializing {
    __Policy_init_unchained(policyEngine);
    __Ownable_init(initialOwner);
    __ERC165_init();
  }

  function __Policy_init_unchained(address policyEngine) internal onlyInitializing {
    _getPolicyStorage().policyEngine = policyEngine;
  }

  /// @inheritdoc IPolicy
  // solhint-disable-next-line no-empty-blocks
  function onInstall(bytes4 /*selector*/ ) public virtual override onlyPolicyEngine {}

  /// @inheritdoc IPolicy
  // solhint-disable-next-line no-empty-blocks
  function onUninstall(bytes4 /*selector*/ ) public virtual override onlyPolicyEngine {}

  function run(
    address caller,
    address subject,
    bytes4 selector,
    bytes[] calldata parameters,
    bytes calldata context
  )
    public
    view
    virtual
    override
    returns (IPolicyEngine.PolicyResult);

  function postRun(
    address, /*caller*/
    address, /*subject*/
    bytes4, /*selector*/
    bytes[] calldata, /*parameters*/
    bytes calldata /*context*/
  )
    public
    virtual
    override
    onlyPolicyEngine
  // solhint-disable-next-line no-empty-blocks
  {}

  /**
   * @dev See {ERC165-supportsInterface}.
   */
  function supportsInterface(bytes4 interfaceId)
    public
    view
    virtual
    override(ERC165Upgradeable, IERC165)
    returns (bool)
  {
    return interfaceId == type(IPolicy).interfaceId || super.supportsInterface(interfaceId);
  }
}

// src/core/PolicyEngine.sol

contract PolicyEngine is Initializable, AccessControlUpgradeable, IPolicyEngine {
  string public constant override typeAndVersion = "PolicyEngine 1.0.0";

  uint256 private constant MAX_POLICIES = 8;
  bytes32 public constant POLICY_CONFIG_ADMIN_ROLE = keccak256("POLICY_CONFIG_ADMIN_ROLE");
  bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

  /// @custom:storage-location erc7201:chainlink.ace.PolicyEngine
  struct PolicyEngineStorage {
    bool defaultPolicyAllow;
    mapping(bytes4 selector => address extractor) extractorBySelector;
    mapping(address policy => address mapper) policyMappers;
    mapping(address policy => uint256 configVersion) policyConfigVersions;
    mapping(address target => bool attached) targetAttached;
    mapping(address target => mapping(bytes4 selector => address[] policies)) targetPolicies;
    mapping(address target => mapping(bytes4 selector => mapping(address policy => bytes32[] policyParameterNames)))
      targetPolicyParameters;
    mapping(address target => bool hasTargetDefault) targetHasDefault;
    mapping(address target => bool targetDefaultPolicyAllow) targetDefaultPolicyAllow;
  }

  // keccak256(abi.encode(uint256(keccak256("chainlink.ace.PolicyEngine")) - 1)) &
  // ~bytes32(uint256(0xff))
  // solhint-disable-next-line const-name-snakecase
  bytes32 private constant policyEngineStorageLocation =
    0x9876d26c639ec5f9246047c1a6b3d2d4c94a7f0dd7848b1a4f882f50fcb29f00;

  function _policyEngineStorage() private pure returns (PolicyEngineStorage storage $) {
    // solhint-disable-next-line no-inline-assembly
    assembly {
      $.slot := policyEngineStorageLocation
    }
  }

  constructor() {
    _disableInitializers();
  }

  /**
   * @dev Initializes the policy engine.
   * @param defaultAllow The default policy result. True to allow, false to reject.
   */
  function initialize(bool defaultAllow, address initialOwner) public virtual initializer {
    __PolicyEngine_init(defaultAllow, initialOwner);
  }

  function __PolicyEngine_init(bool defaultAllow, address initialOwner) internal onlyInitializing {
    __PolicyEngine_init_unchained(defaultAllow, initialOwner);
    __AccessControl_init_unchained();
  }

  function __PolicyEngine_init_unchained(bool defaultAllow, address initialOwner) internal onlyInitializing {
    _policyEngineStorage().defaultPolicyAllow = defaultAllow;
    emit DefaultPolicyAllowSet(defaultAllow);
    _grantRole(DEFAULT_ADMIN_ROLE, initialOwner);
    _grantRole(ADMIN_ROLE, initialOwner);
    _grantRole(POLICY_CONFIG_ADMIN_ROLE, initialOwner);
  }

  /// @inheritdoc IPolicyEngine
  function attach() public {
    _attachTarget(msg.sender);
  }

  function _attachTarget(address target) internal {
    if (_policyEngineStorage().targetAttached[target]) {
      revert IPolicyEngine.TargetAlreadyAttached(target);
    }
    _policyEngineStorage().targetAttached[target] = true;
    emit TargetAttached(target);
  }

  /// @inheritdoc IPolicyEngine
  function detach() public {
    _detachTarget(msg.sender);
  }

  function _detachTarget(address target) internal {
    if (!_policyEngineStorage().targetAttached[target]) {
      revert IPolicyEngine.TargetNotAttached(target);
    }
    _policyEngineStorage().targetAttached[target] = false;
    emit TargetDetached(target);
  }

  /// @inheritdoc IPolicyEngine
  function setDefaultPolicyAllow(bool defaultAllow) public onlyRole(ADMIN_ROLE) {
    _policyEngineStorage().defaultPolicyAllow = defaultAllow;
    emit DefaultPolicyAllowSet(defaultAllow);
  }

  /// @inheritdoc IPolicyEngine
  function setTargetDefaultPolicyAllow(address target, bool defaultAllow) public onlyRole(ADMIN_ROLE) {
    PolicyEngineStorage storage $ = _policyEngineStorage();
    $.targetHasDefault[target] = true;
    $.targetDefaultPolicyAllow[target] = defaultAllow;
    emit TargetDefaultPolicyAllowSet(target, defaultAllow);
  }

  /// @inheritdoc IPolicyEngine
  function setPolicyMapper(address policy, address mapper) public onlyRole(ADMIN_ROLE) {
    _policyEngineStorage().policyMappers[policy] = mapper;
    emit PolicyMapperSet(policy, mapper);
  }

  /// @inheritdoc IPolicyEngine
  function getPolicyMapper(address policy) external view returns (address) {
    return _policyEngineStorage().policyMappers[policy];
  }

  /// @inheritdoc IPolicyEngine
  function check(IPolicyEngine.Payload calldata payload) public view virtual override {
    address[] memory policies = _policyEngineStorage().targetPolicies[msg.sender][payload.selector];

    if (policies.length == 0) {
      _checkDefaultPolicyAllowRevert(msg.sender, payload);
      return;
    }

    IPolicyEngine.Parameter[] memory extractedParameters = _extractParameters(payload);
    for (uint256 i = 0; i < policies.length; i++) {
      address policy = policies[i];

      bytes[] memory policyParameterValues = _policyParameterValues(
        policy,
        _policyEngineStorage().targetPolicyParameters[msg.sender][payload.selector][policy],
        extractedParameters,
        payload
      );
      try IPolicy(policy).run(payload.sender, msg.sender, payload.selector, policyParameterValues, payload.context)
      returns (IPolicyEngine.PolicyResult policyResult) {
        if (policyResult == IPolicyEngine.PolicyResult.Allowed) {
          return;
        } // else continue to next policy
      } catch (bytes memory err) {
        _handlePolicyError(payload, policy, err);
      }
    }

    _checkDefaultPolicyAllowRevert(msg.sender, payload);
  }

  /// @inheritdoc IPolicyEngine
  function run(IPolicyEngine.Payload calldata payload) public virtual override {
    address[] memory policies = _policyEngineStorage().targetPolicies[msg.sender][payload.selector];
    IPolicyEngine.Parameter[] memory extractedParameters = _extractParameters(payload);
    if (policies.length == 0) {
      _checkDefaultPolicyAllowRevert(msg.sender, payload);
      emit PolicyRunComplete(payload.sender, msg.sender, payload.selector, extractedParameters, payload.context);
      return;
    }

    for (uint256 i = 0; i < policies.length; i++) {
      address policy = policies[i];

      bytes[] memory policyParameterValues = _policyParameterValues(
        policy,
        _policyEngineStorage().targetPolicyParameters[msg.sender][payload.selector][policy],
        extractedParameters,
        payload
      );
      try IPolicy(policy).run(payload.sender, msg.sender, payload.selector, policyParameterValues, payload.context)
      returns (IPolicyEngine.PolicyResult policyResult) {
        // solhint-disable-next-line no-empty-blocks
        try IPolicy(policy).postRun(
          payload.sender, msg.sender, payload.selector, policyParameterValues, payload.context
        ) {} catch (bytes memory err) {
          revert IPolicyEngine.PolicyPostRunError(policy, err, payload);
        }
        if (policyResult == IPolicyEngine.PolicyResult.Allowed) {
          emit PolicyRunComplete(payload.sender, msg.sender, payload.selector, extractedParameters, payload.context);
          return;
        }
      } catch (bytes memory err) {
        _handlePolicyError(payload, policy, err);
      }
    }

    _checkDefaultPolicyAllowRevert(msg.sender, payload);
    emit PolicyRunComplete(payload.sender, msg.sender, payload.selector, extractedParameters, payload.context);
  }

  /// @inheritdoc IPolicyEngine
  function setExtractor(bytes4 selector, address extractor) public virtual override onlyRole(ADMIN_ROLE) {
    _setExtractor(selector, extractor);
  }

  /// @inheritdoc IPolicyEngine
  function setExtractors(bytes4[] calldata selectors, address extractor) public virtual override onlyRole(ADMIN_ROLE) {
    for (uint256 i = 0; i < selectors.length; i++) {
      _setExtractor(selectors[i], extractor);
    }
  }

  function _setExtractor(bytes4 selector, address extractor) internal {
    _policyEngineStorage().extractorBySelector[selector] = extractor;
    emit ExtractorSet(selector, extractor);
  }

  /// @inheritdoc IPolicyEngine
  function getExtractor(bytes4 selector) public view virtual override returns (address) {
    return _policyEngineStorage().extractorBySelector[selector];
  }

  /// @inheritdoc IPolicyEngine
  function addPolicy(
    address target,
    bytes4 selector,
    address policy,
    bytes32[] calldata policyParameterNames
  )
    public
    virtual
    override
    onlyRole(ADMIN_ROLE)
  {
    _checkPolicyConfiguration(target, selector, policy);
    _policyEngineStorage().targetPolicies[target][selector].push(policy);
    _policyEngineStorage().targetPolicyParameters[target][selector][policy] = policyParameterNames;
    IPolicy(policy).onInstall(selector);
    emit PolicyAdded(
      target, selector, policy, _policyEngineStorage().targetPolicies[target][selector].length - 1, policyParameterNames
    );
  }

  /// @inheritdoc IPolicyEngine
  function addPolicyAt(
    address target,
    bytes4 selector,
    address policy,
    bytes32[] calldata policyParameterNames,
    uint256 position
  )
    public
    virtual
    override
    onlyRole(ADMIN_ROLE)
  {
    address[] storage policies = _policyEngineStorage().targetPolicies[target][selector];
    if (position > policies.length) {
      revert Policy.InvalidParameters("Position is greater than the number of policies");
    }
    _checkPolicyConfiguration(target, selector, policy);
    policies.push();
    for (uint256 i = policies.length - 1; i > position; i--) {
      policies[i] = policies[i - 1];
    }
    policies[position] = policy;
    _policyEngineStorage().targetPolicyParameters[target][selector][policy] = policyParameterNames;
    IPolicy(policy).onInstall(selector);
    emit PolicyAddedAt(target, selector, policy, position, policyParameterNames, policies);
  }

  /// @inheritdoc IPolicyEngine
  function removePolicy(address target, bytes4 selector, address policy) public virtual override onlyRole(ADMIN_ROLE) {
    address[] storage policies = _policyEngineStorage().targetPolicies[target][selector];
    address removedPolicy = address(0);
    for (uint256 i = 0; i < policies.length; i++) {
      if (policies[i] == policy) {
        removedPolicy = policies[i];
        for (uint256 j = i; j < policies.length - 1; j++) {
          policies[j] = policies[j + 1];
        }
        policies.pop();
        emit PolicyRemoved(target, selector, policy);
        break;
      }
    }
    if (removedPolicy != address(0)) {
      IPolicy(policy).onUninstall(selector);
    }
  }

  /// @inheritdoc IPolicyEngine
  function getPolicies(
    address target,
    bytes4 selector
  )
    public
    view
    virtual
    override
    returns (address[] memory policies)
  {
    return _policyEngineStorage().targetPolicies[target][selector];
  }

  function setPolicyConfiguration(
    address policy,
    uint256 configVersion,
    bytes4 configSelector,
    bytes calldata configData
  )
    public
    virtual
    override
    onlyRole(POLICY_CONFIG_ADMIN_ROLE)
  {
    if (_policyEngineStorage().policyConfigVersions[policy] != configVersion) {
      revert IPolicyEngine.PolicyConfigurationVersionError(
        policy, configVersion, _policyEngineStorage().policyConfigVersions[policy]
      );
    }
    _policyEngineStorage().policyConfigVersions[policy]++;
    (bool success, bytes memory result) = policy.call(abi.encodePacked(configSelector, configData));
    if (!success) {
      revert IPolicyEngine.PolicyConfigurationError(policy, result);
    }
    emit PolicyConfigured(policy, configVersion, configSelector, configData);
  }

  function getPolicyConfigVersion(address policy) public view virtual override returns (uint256) {
    return _policyEngineStorage().policyConfigVersions[policy];
  }

  function _handlePolicyError(Payload memory payload, address policy, bytes memory err) internal pure {
    (bytes4 errorSelector, bytes memory errorData) = _decodeError(err);
    if (errorSelector == IPolicyEngine.PolicyRejected.selector) {
      revert IPolicyEngine.PolicyRunRejected(policy, abi.decode(errorData, (string)), payload);
    } else {
      revert IPolicyEngine.PolicyRunError(policy, err, payload);
    }
  }

  function _checkDefaultPolicyAllowRevert(address target, IPolicyEngine.Payload memory payload) private view {
    PolicyEngineStorage storage $ = _policyEngineStorage();
    bool defaultAllow = $.defaultPolicyAllow;
    if ($.targetHasDefault[target]) {
      defaultAllow = $.targetDefaultPolicyAllow[target];
    }
    if (!defaultAllow) {
      revert IPolicyEngine.PolicyRunRejected(address(0), "no policy allowed the action and default is reject", payload);
    }
  }

  function _checkPolicyConfiguration(address target, bytes4 selector, address policy) private view {
    if (policy == address(0)) {
      revert Policy.InvalidParameters("Policy address cannot be zero");
    }
    if (_policyEngineStorage().targetPolicies[target][selector].length >= MAX_POLICIES) {
      revert Policy.InvalidParameters("Maximum policies reached");
    }
    address[] memory policies = _policyEngineStorage().targetPolicies[target][selector];
    for (uint256 i = 0; i < policies.length; i++) {
      if (policies[i] == policy) {
        revert Policy.InvalidParameters("Policy already added");
      }
    }
  }

  function _extractParameters(IPolicyEngine.Payload memory payload)
    private
    view
    returns (IPolicyEngine.Parameter[] memory)
  {
    IExtractor extractor = IExtractor(_policyEngineStorage().extractorBySelector[payload.selector]);
    IPolicyEngine.Parameter[] memory extractedParameters;

    if (address(extractor) == address(0)) {
      return extractedParameters;
    }

    try extractor.extract(payload) returns (IPolicyEngine.Parameter[] memory _extractedParameters) {
      extractedParameters = _extractedParameters;
    } catch (bytes memory err) {
      revert IPolicyEngine.ExtractorError(address(extractor), err, payload);
    }

    return extractedParameters;
  }

  function _policyParameterValues(
    address policy,
    bytes32[] memory policyParameterNames,
    IPolicyEngine.Parameter[] memory extractedParameters,
    IPolicyEngine.Payload memory payload
  )
    private
    view
    returns (bytes[] memory)
  {
    address mapper = _policyEngineStorage().policyMappers[policy];
    // use custom mapper if set
    if (mapper != address(0)) {
      try IMapper(mapper).map(extractedParameters) returns (bytes[] memory mappedParameters) {
        return mappedParameters;
      } catch (bytes memory err) {
        revert IPolicyEngine.PolicyMapperError(policy, err, payload);
      }
    }

    bytes[] memory policyParameterValues = new bytes[](policyParameterNames.length);

    uint256 parameterCount = policyParameterNames.length;
    if (parameterCount == 0) {
      return policyParameterValues;
    }

    uint256 mappedParameterCount = 0;
    for (uint256 i = 0; i < extractedParameters.length; i++) {
      for (uint256 j = 0; j < parameterCount; j++) {
        if (extractedParameters[i].name == policyParameterNames[j]) {
          policyParameterValues[j] = extractedParameters[i].value;
          mappedParameterCount++;
          break;
        }
      }
      if (mappedParameterCount == parameterCount) {
        return policyParameterValues;
      }
    }
    revert Policy.InvalidParameters("Missing policy parameters");
  }

  function _decodeError(bytes memory err) internal pure returns (bytes4, bytes memory) {
    // If the error length is less than 4, it is not a valid error
    if (err.length < 4) {
      return (0, err);
    }
    bytes4 selector = bytes4(err);
    bytes memory errorData = new bytes(err.length - 4);
    for (uint256 i = 0; i < err.length - 4; i++) {
      errorData[i] = err[i + 4];
    }
    return (selector, errorData);
  }
}

contract VolumePolicy is Policy {
  string public constant override typeAndVersion = "VolumePolicy 1.0.0";

  /**
   * @notice Emitted when the maximum volume limit is set.
   * @param maxAmount The maximum amount parameter limit. If set to 0, there is no maximum limit.
   */
  event MaxVolumeSet(uint256 maxAmount);
  /**
   * @notice Emitted when the minimum volume limit is set.
   * @param minAmount The minimum amount parameter limit. If set to 0, there is no minimum limit.
   */
  event MinVolumeSet(uint256 minAmount);

  /// @custom:storage-location erc7201:chainlink.ace.VolumePolicy
  struct VolumePolicyStorage {
    /// @notice The maximum amount parameter limit. If set to 0, there is no maximum limit.
    uint256 maxAmount;
    /// @notice The minimum amount parameter limit. If set to 0, there is no minimum limit.
    uint256 minAmount;
  }

  // keccak256(abi.encode(uint256(keccak256("chainlink.ace.VolumePolicy")) - 1)) & ~bytes32(uint256(0xff))
  bytes32 private constant VolumePolicyStorageLocation =
    0xdd888d8665dad1bb1ff8a8ea6b67594646c08e9b44b1c5c650075d4887d09500;

  function _getVolumePolicyStorage() private pure returns (VolumePolicyStorage storage $) {
    assembly {
      $.slot := VolumePolicyStorageLocation
    }
  }

  /**
   * @notice Configures the policy by setting minimum and maximum amount parameter limits.
   * @param parameters ABI-encoded bytes containing two `uint256` values: the min and max amount limits.
   *      - `minAmount`: The minimum amount parameter limit, 0 for no minimum.
   *      - `maxAmount`: The maximum amount parameter limit, 0 for no maximum.
   * @dev These limits apply to function parameters, not actual token transfer amounts.
   */
  function configure(bytes calldata parameters) internal override {
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    ($.minAmount, $.maxAmount) = abi.decode(parameters, (uint256, uint256));
    require($.maxAmount > $.minAmount || $.maxAmount == 0, "maxAmount must be greater than minAmount");
  }

  /**
   * @notice Sets the maximum amount parameter limit for the policy.
   * @param maxAmount The maximum amount parameter limit.
   * @dev Reverts if the new max amount is less than or equal to the min amount (unless maxAmount is 0),
   * or if it is the same as the current max amount.
   */
  function setMax(uint256 maxAmount) public onlyOwner {
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    require(maxAmount > $.minAmount || maxAmount == 0, "maxAmount must be greater than minAmount");
    require(maxAmount != $.maxAmount, "maxAmount cannot be the same as current maxAmount");
    $.maxAmount = maxAmount;
    emit MaxVolumeSet(maxAmount);
  }

  /**
   * @notice Gets the current maximum amount parameter limit for the policy.
   * @return maxAmount The current maximum amount parameter limit. If 0, there is no maximum limit.
   */
  function getMax() public view returns (uint256) {
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    return $.maxAmount;
  }

  /**
   * @notice Sets the minimum amount parameter limit for the policy.
   * @param minAmount The minimum amount parameter limit.
   * @dev Reverts if the new min amount is greater than or equal to the max amount (unless maxAmount is 0),
   * or if it is the same as the current min amount.
   */
  function setMin(uint256 minAmount) public onlyOwner {
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    require(minAmount < $.maxAmount || $.maxAmount == 0, "minAmount must be less than maxAmount");
    require(minAmount != $.minAmount, "minAmount cannot be the same as current minAmount");
    $.minAmount = minAmount;
    emit MinVolumeSet(minAmount);
  }

  /**
   * @notice Gets the current minimum amount parameter limit for the policy.
   * @return minAmount The current minimum amount parameter limit. If 0, there is no minimum limit.
   */
  function getMin() public view returns (uint256) {
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    return $.minAmount;
  }

  /**
   * @notice Function called by the policy engine to validate amount parameters against configured limits.
   * @param parameters [amount(uint256)] The parameters of the called method.
   * @return result The result of the policy check.
   * @dev This policy validates the amount argument passed to a protected function, not actual token transfers.
   * For tokens with fees, rebasing, or other non-standard behaviors, the actual transferred amount may differ.
   */
  function run(
    address, /* caller */
    address, /* subject */
    bytes4, /*selector*/
    bytes[] calldata parameters,
    bytes calldata /* context */
  )
    public
    view
    override
    returns (IPolicyEngine.PolicyResult)
  {
    if (parameters.length != 1) {
      revert InvalidParameters("expected 1 parameter");
    }
    uint256 amount = abi.decode(parameters[0], (uint256));

    // Gas optimization: load storage reference once
    VolumePolicyStorage storage $ = _getVolumePolicyStorage();
    if (($.maxAmount != 0 && amount > $.maxAmount) || amount < $.minAmount) {
      revert IPolicyEngine.PolicyRejected("amount outside allowed volume limits");
    }

    return IPolicyEngine.PolicyResult.Continue;
  }
}
