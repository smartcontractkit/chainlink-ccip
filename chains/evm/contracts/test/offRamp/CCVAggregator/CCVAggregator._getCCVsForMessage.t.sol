// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

contract CCVAggregator__getCCVsForMessage is CCVAggregatorSetup {
  using EnumerableSet for EnumerableSet.AddressSet;

  EnumerableSet.AddressSet internal s_ccvs;

  function testFuzz__getCCVsForMessage_MultipleDefaultsHandledCorrectly(
    uint8[] memory defaults,
    uint8[] memory receiverRequired,
    uint8[] memory pool
  ) public {
    vm.assume(defaults.length > 0);

    address[] memory defaultCCVs = new address[](defaults.length);
    address[] memory receiverCCVs = new address[](receiverRequired.length);
    address[] memory poolCCVs = new address[](pool.length);

    for (uint256 i = 0; i < defaults.length; i++) {
      defaultCCVs[i] = address(uint160(defaults[i]) + 1);
    }

    for (uint256 i = 0; i < receiverRequired.length; i++) {
      receiverCCVs[i] = address(uint160(receiverRequired[i]));
    }

    for (uint256 i = 0; i < pool.length; i++) {
      poolCCVs[i] = address(uint160(pool[i]));
    }

    // Check for address(0) BEFORE removing duplicates
    bool hadZeroInReceiver = false;
    bool hadZeroInPool = false;
    for (uint256 i = 0; i < receiverCCVs.length; i++) {
      if (receiverCCVs[i] == address(0)) hadZeroInReceiver = true;
    }
    for (uint256 i = 0; i < poolCCVs.length; i++) {
      if (poolCCVs[i] == address(0)) hadZeroInPool = true;
    }

    defaultCCVs = _onlyUniques(defaultCCVs);
    receiverCCVs = _onlyUniquesWithZero(receiverCCVs);
    poolCCVs = _onlyUniquesWithZero(poolCCVs);

    _applySourceConfig(abi.encode(makeAddr("onRamp")), true, defaultCCVs, new address[](0));

    address receiver = makeAddr("receiver");

    _setGetCCVsReturnData(receiver, SOURCE_CHAIN_SELECTOR, receiverCCVs, new address[](0), 0);

    // Populate the expected set based on the function's behavior:

    // The function will include defaults in these cases:
    // 1. If address(0) is found in any arrays (replaced by defaults)
    // 2. If receiver returns empty CCVs (receiver fallback)
    // 3. If pool returns empty CCVs (pool fallback)

    // Note: We need to be careful about double-counting defaults
    // Defaults are added if:
    // - address(0) was in receiver/pool arrays before deduplication
    // - receiver returns empty (after deduplication)
    // - pool returns empty (after deduplication)
    bool defaultsWillBeAdded = hadZeroInReceiver || hadZeroInPool || receiverCCVs.length == 0 || poolCCVs.length == 0;

    if (defaultsWillBeAdded) {
      for (uint256 i = 0; i < defaultCCVs.length; i++) {
        s_ccvs.add(defaultCCVs[i]);
      }
    }

    // For receiver: if empty, _getCCVsFromReceiver returns defaults (already added above)
    // If not empty, it returns the receiver CCVs
    if (receiverCCVs.length > 0) {
      for (uint256 i = 0; i < receiverCCVs.length; i++) {
        s_ccvs.add(receiverCCVs[i]);
      }
    }

    // For pool: if empty, _getCCVsFromPool returns defaults (already added above)
    // If not empty, it returns the pool CCVs
    if (poolCCVs.length > 0) {
      for (uint256 i = 0; i < poolCCVs.length; i++) {
        s_ccvs.add(poolCCVs[i]);
      }
    }

    // Remove address(0) from the set as it will be replaced by the defaults.
    s_ccvs.remove(address(0));

    MessageV1Codec.TokenTransferV1[] memory tokenAmounts = _setCCVsOnPool(poolCCVs);

    (address[] memory requiredCCVs,,) = s_agg.__getCCVsForMessage(SOURCE_CHAIN_SELECTOR, receiver, tokenAmounts, 0);

    assertEq(_onlyUniques(requiredCCVs).length, requiredCCVs.length, "still has duplicates");

    for (uint256 i = 0; i < requiredCCVs.length; i++) {
      assertTrue(s_ccvs.contains(requiredCCVs[i]), "requiredCCVs[i] not found in s_ccvs");
    }
    assertEq(s_ccvs.length(), requiredCCVs.length, "ccvs length mismatch");
  }

  function _onlyUniques(
    address[] memory input
  ) internal pure returns (address[] memory output) {
    output = new address[](input.length);
    uint256 outputIndex = 0;

    for (uint256 i = 0; i < input.length; i++) {
      bool isUnique = true;
      for (uint256 j = 0; j < outputIndex; j++) {
        if (output[j] == input[i] || input[i] == address(0)) {
          isUnique = false;
          break;
        }
      }
      if (isUnique) {
        output[outputIndex++] = input[i];
      }
    }

    assembly {
      mstore(output, outputIndex)
    }
    return output;
  }

  function _onlyUniquesWithZero(
    address[] memory input
  ) internal pure returns (address[] memory output) {
    output = new address[](input.length);
    uint256 outputIndex = 0;

    for (uint256 i = 0; i < input.length; i++) {
      bool isUnique = true;
      for (uint256 j = 0; j < outputIndex; j++) {
        if (output[j] == input[i]) {
          isUnique = false;
          break;
        }
      }
      if (isUnique) {
        output[outputIndex++] = input[i];
      }
    }

    assembly {
      mstore(output, outputIndex)
    }
    return output;
  }

  function _setCCVsOnPool(
    address[] memory ccvs
  ) internal returns (MessageV1Codec.TokenTransferV1[] memory tokenAmounts) {
    address pool = makeAddr("pool");
    address sourceToken = makeAddr("sourceToken");
    address token = makeAddr("token");

    // Modify message with token transfer.
    tokenAmounts = new MessageV1Codec.TokenTransferV1[](1);
    tokenAmounts[0] = MessageV1Codec.TokenTransferV1({
      amount: 1,
      sourcePoolAddress: abi.encodePacked(pool),
      sourceTokenAddress: abi.encodePacked(sourceToken),
      destTokenAddress: abi.encodePacked(token),
      extraData: ""
    });

    // Mock token admin registry to return the pool.
    vm.mockCall(s_tokenAdminRegistry, abi.encodeCall(ITokenAdminRegistry.getPool, (token)), abi.encode(pool));

    // Mock pool supportsInterface for IPoolV2.
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool to require a specific CCV.
    vm.mockCall(
      pool,
      abi.encodeCall(
        IPoolV2.getRequiredInboundCCVs,
        (token, SOURCE_CHAIN_SELECTOR, tokenAmounts[0].amount, 0, tokenAmounts[0].extraData)
      ),
      abi.encode(ccvs)
    );

    return tokenAmounts;
  }
}
