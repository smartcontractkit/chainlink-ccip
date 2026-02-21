// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IBurnMintERC20} from "../../../interfaces/IBurnMintERC20.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";
import {PolicyEngine, Policy, VolumePolicy} from "../../helpers/ACEPolicyEngineFlat.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {AdvancedPoolHooksExtractor} from "../../../pools/extractors/AdvancedPoolHooksExtractor.sol";
import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {Client} from "../../../libraries/Client.sol";
import {OnRampSetup} from "../../onRamp/OnRamp/OnRampSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts@5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract AdvancedPoolHooksE2ESetup is OnRampSetup {
  PolicyEngine internal s_policyEngine;
  AdvancedPoolHooksExtractor internal s_extractor;
  VolumePolicy internal s_volumePolicy;

  AdvancedPoolHooks internal s_advancedPoolHooks;
  BurnMintTokenPool internal s_burnMintPool;
  BurnMintERC20 internal s_aceToken;

  uint256 internal constant VOLUME_MIN = 10e18;
  uint256 internal constant VOLUME_MAX = 1000e18;
  uint256 internal constant VALID_AMOUNT = 100e18;
  uint256 internal constant TOO_HIGH_AMOUNT = 1001e18;
  uint256 internal constant TOO_LOW_AMOUNT = 5e18;

  function setUp() public virtual override {
    super.setUp();

    // Deploy the ERC20 token
    s_aceToken = new BurnMintERC20("ACE Test Token", "ACET", 18, 0, 0);

    // Deploy ACE PolicyEngine via ERC1967Proxy
    PolicyEngine policyEngineImpl = new PolicyEngine();
    bytes memory policyEngineData = abi.encodeWithSelector(
      PolicyEngine.initialize.selector,
      true, // defaultAllow
      OWNER
    );
    s_policyEngine = PolicyEngine(address(new ERC1967Proxy(address(policyEngineImpl), policyEngineData)));

    // Deploy and register the AdvancedPoolHooksExtractor
    s_extractor = new AdvancedPoolHooksExtractor();

    s_policyEngine.setExtractor(
      IAdvancedPoolHooks.preflightCheck.selector,
      address(s_extractor)
    );
    s_policyEngine.setExtractor(
      IAdvancedPoolHooks.postflightCheck.selector,
      address(s_extractor)
    );

    // Deploy VolumePolicy via ERC1967Proxy
    VolumePolicy volumePolicyImpl = new VolumePolicy();
    bytes memory policyData = abi.encodeWithSelector(
      Policy.initialize.selector,
      address(s_policyEngine),
      OWNER,
      abi.encode(VOLUME_MIN, VOLUME_MAX)
    );
    s_volumePolicy = VolumePolicy(address(new ERC1967Proxy(address(volumePolicyImpl), policyData)));

    // Deploy AdvancedPoolHooks with no allowlist, no threshold, with PolicyEngine, no authorized callers
    s_advancedPoolHooks = new AdvancedPoolHooks(
      new address[](0),
      0,
      address(s_policyEngine),
      new address[](0)
    );

    // Add VolumePolicy for both selectors
    bytes32[] memory volumeParams = new bytes32[](1);
    volumeParams[0] = s_extractor.PARAM_AMOUNT();

    s_policyEngine.addPolicy(
      address(s_advancedPoolHooks),
      IAdvancedPoolHooks.preflightCheck.selector,
      address(s_volumePolicy),
      volumeParams
    );

    s_policyEngine.addPolicy(
      address(s_advancedPoolHooks),
      IAdvancedPoolHooks.postflightCheck.selector,
      address(s_volumePolicy),
      volumeParams
    );

    // Deploy BurnMintTokenPool with hooks
    s_burnMintPool = new BurnMintTokenPool(
      IBurnMintERC20(address(s_aceToken)),
      18,
      address(s_advancedPoolHooks),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );

    // Grant minter/burner roles to the pool
    s_aceToken.grantMintAndBurnRoles(address(s_burnMintPool));

    // Register token pool in TokenAdminRegistry
    s_tokenAdminRegistry.proposeAdministrator(address(s_aceToken), OWNER);
    s_tokenAdminRegistry.acceptAdminRole(address(s_aceToken));
    s_tokenAdminRegistry.setPool(address(s_aceToken), address(s_burnMintPool));

    // Configure pool for DEST_CHAIN_SELECTOR
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_burnMintPool));
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(s_aceToken)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_burnMintPool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Configure pool for SOURCE_CHAIN_SELECTOR for dest chain tests
    bytes[] memory remotePoolAddressesSrc = new bytes[](1);
    remotePoolAddressesSrc[0] = abi.encode(address(s_burnMintPool));
    TokenPool.ChainUpdate[] memory chainUpdateSrc = new TokenPool.ChainUpdate[](1);
    chainUpdateSrc[0] = TokenPool.ChainUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddressesSrc,
      remoteTokenAddress: abi.encode(address(s_aceToken)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_burnMintPool.applyChainUpdates(new uint64[](0), chainUpdateSrc);

    // Register ACE token price in FeeQuoter so ccipSend can calculate fees
    s_feeQuoter.updatePrices(_getSingleTokenPriceUpdateStruct(address(s_aceToken), 1e18));
  }

  /// @notice Builds a Client.EVM2AnyMessage for ccipSend with a token transfer.
  function _buildCCIPMessage(
    address receiver,
    uint256 tokenAmount
  ) internal view returns (Client.EVM2AnyMessage memory) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({
      token: address(s_aceToken),
      amount: tokenAmount
    });
    return message;
  }
}
