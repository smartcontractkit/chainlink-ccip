pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_generateNewReleaseOrMintIn is USDCTokenPoolProxySetup {
  function test_generateNewReleaseOrMintIn() public {
    // Arrange: Prepare test data
    uint64 nonce = 12345;
    uint32 sourceDomain = 67890;

    bytes memory originalSender = abi.encode(makeAddr("sender"));
    bytes memory sourcePoolAddress = abi.encode(makeAddr("sourcePoolAddress"));
    bytes memory offchainTokenData = "";

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      originalSender: originalSender,
      receiver: makeAddr("receiver"),
      sourceDenominatedAmount: 1e6,
      localToken: address(s_USDCToken),
      sourcePoolData: abi.encode(
        USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: nonce, sourceDomain: sourceDomain})
      ),
      sourcePoolAddress: sourcePoolAddress,
      offchainTokenData: offchainTokenData
    });

    Pool.ReleaseOrMintInV1 memory returnedReleaseOrMintIn =
      s_usdcTokenPoolProxy.generateNewReleaseOrMintIn(releaseOrMintIn);

    bytes memory expectedSourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV1({nonce: nonce, sourceDomain: sourceDomain})
    );

    assertEq(returnedReleaseOrMintIn.originalSender, releaseOrMintIn.originalSender, "originalSender does not match");
    assertEq(
      returnedReleaseOrMintIn.remoteChainSelector,
      releaseOrMintIn.remoteChainSelector,
      "remoteChainSelector does not match"
    );
    assertEq(returnedReleaseOrMintIn.receiver, releaseOrMintIn.receiver, "receiver does not match");
    assertEq(
      returnedReleaseOrMintIn.sourceDenominatedAmount,
      releaseOrMintIn.sourceDenominatedAmount,
      "sourceDenominatedAmount does not match"
    );
    assertEq(returnedReleaseOrMintIn.localToken, releaseOrMintIn.localToken, "localToken does not match");
    assertEq(
      returnedReleaseOrMintIn.sourcePoolAddress, releaseOrMintIn.sourcePoolAddress, "sourcePoolAddress does not match"
    );
    assertEq(returnedReleaseOrMintIn.sourcePoolData, expectedSourcePoolData, "sourcePoolData does not match");
    assertEq(
      returnedReleaseOrMintIn.offchainTokenData, releaseOrMintIn.offchainTokenData, "offchainTokenData does not match"
    );
  }
}
