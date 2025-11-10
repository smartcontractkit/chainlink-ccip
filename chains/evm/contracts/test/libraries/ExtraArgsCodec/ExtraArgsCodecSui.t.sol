// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";

/// forge-config: default.allow_internal_expect_revert = true
contract ExtraArgsCodecSui_Test is BaseTest {
  ExtraArgsCodecHelper internal s_helper;

  function setUp() public override {
    super.setUp();
    s_helper = new ExtraArgsCodecHelper();
  }

  function test__ecodeSuiExecutorArgsV1_NoObjectIds() public view {
    ExtraArgsCodec.SuiExecutorArgsV1 memory args =
      ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    assertEq(encoded.length, ExtraArgsCodec.SUI_EXECUTOR_ARGS_V1_BASE_SIZE);

    ExtraArgsCodec.SuiExecutorArgsV1 memory decoded = s_helper._decodeSuiExecutorArgsV1(encoded);

    assertEq(decoded.receiverObjectIds.length, 0);
  }

  function test__decodeSuiExecutorArgsV1_WithObjectIds() public view {
    bytes32[] memory objectIds = new bytes32[](2);
    objectIds[0] = keccak256("object1");
    objectIds[1] = keccak256("object2");

    ExtraArgsCodec.SuiExecutorArgsV1 memory args = ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: objectIds});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    ExtraArgsCodec.SuiExecutorArgsV1 memory decoded = s_helper._decodeSuiExecutorArgsV1(encoded);

    assertEq(decoded.receiverObjectIds.length, 2);
    assertEq(decoded.receiverObjectIds[0], keccak256("object1"));
    assertEq(decoded.receiverObjectIds[1], keccak256("object2"));
  }

  function test__decodeSuiExecutorArgsV1_RevertWhen_EXTRA_ARGS_STATIC_LENGTH_FIELDS() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector,
        ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS,
        3
      )
    );
    s_helper._decodeSuiExecutorArgsV1(new bytes(3));
  }

  function test__decodeSuiExecutorArgsV1_RevertWhen_SUI_EXECUTOR_FINAL_OFFSET() public {
    ExtraArgsCodec.SuiExecutorArgsV1 memory args =
      ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    bytes memory withExtra = bytes.concat(encoded, bytes("extra"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.SUI_EXECUTOR_FINAL_OFFSET, 5
      )
    );
    s_helper._decodeSuiExecutorArgsV1(withExtra);
  }

  function test__decodeSuiExecutorArgsV1_RevertWhen_SUI_EXECUTOR_OBJECT_IDS_CONTENT() public {
    bytes memory invalidData = abi.encodePacked(
      ExtraArgsCodec.SUI_EXECUTOR_ARGS_V1_TAG,
      uint8(2) // Claims 2 object IDs but doesn't provide them.
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector,
        ExtraArgsCodec.EncodingErrorLocation.SUI_EXECUTOR_OBJECT_IDS_CONTENT,
        5
      )
    );
    s_helper._decodeSuiExecutorArgsV1(invalidData);
  }

  function test__encodeSuiExecutorArgsV1_RevertWhen_ENCODE_SUI_OBJECT_IDS_LENGTH() public {
    bytes32[] memory objectIds = new bytes32[](257);

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.ENCODE_SUI_OBJECT_IDS_LENGTH, 0
      )
    );
    ExtraArgsCodec._encodeSuiExecutorArgsV1(ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: objectIds}));
  }
}
