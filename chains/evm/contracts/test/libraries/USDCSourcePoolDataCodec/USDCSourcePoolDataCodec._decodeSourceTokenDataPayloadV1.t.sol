// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCSourcePoolDataCodec} from "../../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCSourcePoolDataCodecHelper} from "../../helpers/USDCSourcePoolDataCodecHelper.sol";
import {Test} from "forge-std/Test.sol";

contract USDCSourcePoolDataCodec__decodeSourceTokenDataPayloadV1 is Test {
  USDCSourcePoolDataCodecHelper internal s_helper;

  uint64 internal constant NONCE = 12345;
  uint32 internal constant SOURCE_DOMAIN = 1553252;

  function setUp() public {
    s_helper = new USDCSourcePoolDataCodecHelper();
  }

  function test__decodeSourceTokenDataPayloadV1_CCTPV1() public pure {
    // Encode using the V1 function
    bytes memory payload = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: NONCE, sourceDomain: SOURCE_DOMAIN})
    );

    // Decode the payload
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory decoded =
      USDCSourcePoolDataCodec._decodeSourceTokenDataPayloadV1(payload);

    // Compare individual fields
    assertEq(decoded.nonce, NONCE, "Nonce mismatch");
    assertEq(decoded.sourceDomain, SOURCE_DOMAIN, "Source domain mismatch");
  }

  // Reverts

  function test__decodeSourceTokenDataPayloadV1_RevertWhen_InvalidVersionV2() public {
    bytes memory invalidPayload = abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_TAG, NONCE, SOURCE_DOMAIN);

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCSourcePoolDataCodec.InvalidVersion.selector, USDCSourcePoolDataCodec.CCTP_VERSION_2_TAG
      )
    );

    s_helper.decodeSourceTokenDataPayloadV1(invalidPayload);
  }

  function test__decodeSourceTokenDataPayloadV1_RevertWhen_InvalidVersionUnknown() public {
    bytes4 unknownVersion = 0x12345678;
    bytes memory invalidPayload = abi.encodePacked(unknownVersion, NONCE, SOURCE_DOMAIN);

    vm.expectRevert(abi.encodeWithSelector(USDCSourcePoolDataCodec.InvalidVersion.selector, unknownVersion));

    s_helper.decodeSourceTokenDataPayloadV1(invalidPayload);
  }
}
