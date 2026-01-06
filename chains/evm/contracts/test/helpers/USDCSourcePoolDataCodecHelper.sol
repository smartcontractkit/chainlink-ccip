// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";

// Helper contract to expose library functions for testing reverts
contract USDCSourcePoolDataCodecHelper {
  function decodeSourceTokenDataPayloadV2(
    bytes memory sourcePoolData
  ) external pure returns (USDCSourcePoolDataCodec.SourceTokenDataPayloadV2 memory) {
    return USDCSourcePoolDataCodec._decodeSourceTokenDataPayloadV2(sourcePoolData);
  }
}
