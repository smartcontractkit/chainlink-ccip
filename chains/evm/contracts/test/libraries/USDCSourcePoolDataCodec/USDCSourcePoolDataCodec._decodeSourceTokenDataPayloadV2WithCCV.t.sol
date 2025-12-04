// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCSourcePoolDataCodec} from "../../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCSourcePoolDataCodecHelper} from "../../helpers/USDCSourcePoolDataCodecHelper.sol";
import {Test} from "forge-std/Test.sol";

contract USDCSourcePoolDataCodec__decodeSourceTokenDataPayloadV2 is Test {
  USDCSourcePoolDataCodecHelper internal s_helper;

  bytes32 internal constant DEPOSIT_HASH = keccak256("test deposit hash");
  uint32 internal constant SOURCE_DOMAIN = 1553252;

  function setUp() public {
    s_helper = new USDCSourcePoolDataCodecHelper();
  }

  function test__decodeSourceTokenDataPayloadV2WithCCV() public pure {
    // Encode using the V2 CCV function
    bytes memory payload = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2WithCCV();

    // Decode the payload
    bytes4 poolVersionTag = USDCSourcePoolDataCodec._decodeSourceTokenDataPayloadV2WithCCV(payload);

    // Compare pool version tag
    assertEq(poolVersionTag, USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG, "Pool version tag mismatch");
  }

  // Reverts

  function test__decodeSourceTokenDataPayloadV2WithCCV_RevertWhen_InvalidVersionV1() public {
    bytes memory invalidPayload =
      abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_1_TAG, SOURCE_DOMAIN, DEPOSIT_HASH);

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCSourcePoolDataCodec.InvalidVersion.selector, USDCSourcePoolDataCodec.CCTP_VERSION_1_TAG
      )
    );

    s_helper.decodeSourceTokenDataPayloadV2WithCCV(invalidPayload);
  }

  function test__decodeSourceTokenDataPayloadV2WithCCV_RevertWhen_InvalidVersionUnknown() public {
    bytes4 unknownVersion = 0x12345678;
    bytes memory invalidPayload = abi.encodePacked(unknownVersion, SOURCE_DOMAIN, DEPOSIT_HASH);

    vm.expectRevert(abi.encodeWithSelector(USDCSourcePoolDataCodec.InvalidVersion.selector, unknownVersion));

    s_helper.decodeSourceTokenDataPayloadV2WithCCV(invalidPayload);
  }
}
