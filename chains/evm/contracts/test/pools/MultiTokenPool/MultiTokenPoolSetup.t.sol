// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {BurnMintMultiTokenPool} from "../../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {MockPoolConfigManager} from "../../mocks/MockPoolConfigManager.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract MultiTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  MockPoolConfigManager internal s_poolConfigManager;
  BurnMintMultiTokenPool internal s_pool;

  address internal s_allowedOffRamp = makeAddr("allowed_offRamp");
  address internal s_allowedOnRamp = makeAddr("allowed_onRamp");

  address internal s_remotePool = makeAddr("remotePool");
  address internal s_remoteToken = makeAddr("remoteToken");
  address internal s_receiver = makeAddr("receiver");

  /// @dev The pool passes the pool data through without reading it.
  bytes internal constant DEST_POOL_DATA = "destPoolData";

  function setUp() public virtual override {
    super.setUp();

    s_token = IERC20(
      address(
        new CrossChainToken(
          BaseERC20.ConstructorParams({
            name: "LINK",
            symbol: "LNK",
            decimals: 18,
            maxSupply: 0,
            preMint: 0,
            preMintRecipient: address(0),
            ccipAdmin: OWNER
          }),
          OWNER,
          OWNER
        )
      )
    );

    s_poolConfigManager = new MockPoolConfigManager();
    s_pool = new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(s_sourceRouter), address(s_poolConfigManager));
    CrossChainToken(address(s_token)).grantMintAndBurnRoles(address(s_pool));

    // The default resolved config: a registered Burn token with a lane to the destination chain.
    s_poolConfigManager.setTokenConfig(_defaultTokenConfig());
    s_poolConfigManager.setOutboundLane(_defaultOutboundLane());
    s_poolConfigManager.setInboundTransfer(_defaultInboundTransfer(0));
    s_poolConfigManager.setLaneConnection(_defaultLaneConnection());

    _setMockRouterOnRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOnRamp);
    _setMockRouterOffRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOffRamp, true);
  }

  function _defaultTokenConfig() internal pure returns (IPoolConfigManager.TokenConfig memory config) {
    config.isRegistered = true;
    config.localDecimals = DEFAULT_TOKEN_DECIMALS;
    config.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Burn;
    return config;
  }

  function _defaultLaneConnection() internal view returns (IPoolConfigManager.LaneConnection memory connection) {
    connection.hasConnection = true;
    connection.remoteToken = abi.encode(s_remoteToken);
    return connection;
  }

  function _defaultOutboundLane() internal view returns (IPoolConfigManager.OutboundLane memory lane) {
    lane.remoteToken = abi.encode(s_remoteToken);
    lane.destPoolData = DEST_POOL_DATA;
    lane.localDecimals = DEFAULT_TOKEN_DECIMALS;
    lane.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Burn;
    return lane;
  }

  function _defaultInboundTransfer(
    uint256 localAmount
  ) internal pure returns (IPoolConfigManager.ResolvedInboundTransfer memory transfer) {
    transfer.localAmount = localAmount;
    transfer.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Burn;
    return transfer;
  }

  function _lockOrBurnIn(
    uint256 amount
  ) internal view returns (Pool.LockOrBurnInV1 memory) {
    return Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: bytes(""),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });
  }

  function _releaseOrMintIn(
    uint256 amount
  ) internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(OWNER),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: s_receiver,
      sourceDenominatedAmount: amount,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_remotePool),
      sourcePoolData: "",
      offchainTokenData: ""
    });
  }
}
