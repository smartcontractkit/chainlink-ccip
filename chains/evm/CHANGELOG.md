[NOTES.md](../../../chainlink/contracts/release/ccip/NOTES.md)# @chainlink/contracts-ccip

## 1.6.2

CCIP 1.6.2 is a minor release that focuses on adding support for new networks in miscellaneous token pools and deployable
ERC20-tokens.

Key improvements include: 
- **USDC Token Pool Support for Solana** - To enable sending USDC to and from Solana the USDC token pool has added a 
  `mintRecipientOverride` field for its CCTP messages. Since the PDA on Solana which will be the recipient of the minted
  USDC will not be the final recipient of the tokens, this field enables better support with CCTP for the Solana-CCIP 
  contracts

- **Hyperliquid Support for FactoryBurnMintERC20** - Tokens which are deployed using the Chainlink Token Manager interface
and associated contracts are now able to be deployed on Hyperliquid by adding support for the spot balance precompile 
system on Hypercore.

### Patch Changes

- [#1212](https://github.com/smartcontractkit/chainlink-ccip/pull/1212) - Fixes a bug in the FactoryBurnMintERC20 constructor
that prevents using a `preMint` and a `maxSupply` of `0`

- [#1072](https://github.com/smartcontractkit/chainlink-ccip/pull/1072) - Adds the CCTP Message Transmitter Proxy gobindings
wrapper

- [#1116](https://github.com/smartcontractkit/chainlink-ccip/pull/1116) - Removes the granting of mint and burn permissions from the FactoryBurnMintERC20 constructor to save gas on deplotyment.

### Minor Changes

- [#1145](https://github.com/smartcontractkit/chainlink-ccip/pull/1145) [`4f2c735`](https://github.com/smartcontractkit/chainlink-ccip/commit/4f2c735bf252ec69be1424647671b956dd159c55) - Add a new BurnMintERC20 compatible with Hyperliquid and changes to the FactoryBurnMintERC20:
  - Adds a check to the constructor that the `preMint` is less than the `maxSupply` 

- [#1009](https://github.com/smartcontractkit/chainlink-ccip/pull/1009) [`d22a78e`](https://github.com/smartcontractkit/chainlink-ccip/pull/1009/commits/d22a78e6de780b6bb1f8259a6ea7753ed92892ce) - Adds support for USDC messages on Solana using a mint-recipient
override field for CCTP messages.

## 1.6.1

CCIP 1.6.1 is a minor release that focuses on token pools, adding overall pool improvements.

Key improvements include:

- **Rate Limit With Local Denomination** - Token pools now use local token denomination for rate limiting during `releaseOrMint`, making it easier to handle tokens with different decimals cross-chain.
  - `ReleaseOrMintInV1.amount` has been changed to `ReleaseOrMintInV1.sourceDenominatedAmount`.
  - Conversion to local token denomination happens inside `releaseOrMint`.

### Minor Changes

- [#1006](https://github.com/smartcontractkit/chainlink-ccip/pull/1006) [`4e7de54`](https://github.com/smartcontractkit/chainlink-ccip/commit/4e7de54014) - #feature Change rate limiting to use local token denomination inside `releaseOrMint`

### Patch Changes

- [#950](https://github.com/smartcontractkit/chainlink-ccip/pull/950) [`6d82d44`](https://github.com/smartcontractkit/chainlink-ccip/commit/6d82d44e86) - Series of TokenPool improvements:
  - Remove `acceptLiquidity`from the LockRelease pool
  - Add a virtual `_releaseOrMint` and `_lockOrBurn` for easy pool implementations
  - Merge `Locked` and `Burned` event into `LockedOrBurned`, merge `Released` and `Minted` into `ReleasedOrMinted`
  - Remove `ILiquidityContainer`
  - When calling transferLiquidity with uint256.max as amount, it will transfer the current amount in the pool
  - Setting 0x0 as rebalancer is now allowed in siloed pools
  - Remove the aggregate rate limiter logic in RateLimiter
  - Allow setting 0 as rate and/or capacity
  - Add `OutboundRateLimitConsumed` and `InboundRateLimitConsumed` events
  - Use OZ 5 for non-token related imports
  - Add `RebalancerSet` event to the LockRelease pool

## 1.6.0

CCIP 1.6 has a new home: [github.com/smartcontractkit/chainlink-ccip](https://github.com/smartcontractkit/chainlink-ccip)!
This is the new long-term home for not only the EVM contracts, but also SVM (Solana), Aptos and all future non-EVM chains.

v1.6 is mostly an invisible release for end users, but features a full rework under the hood. Some of the features include

- More gas efficient code, leading to lower fees.
- Offchain RMN blessings.
  - The Commit DON now directly communicated with the RMN and submits pre-blessed roots.
  - Reduces gas cost of blessing significantly.
  - Reduces the number of onchain transactions, lowering the time it takes to execute a CCIP message.
- Merged the CommitStore with the OffRamp.
- A single OnRamp and OffRamp contract per chain instead of one per lane.
  - This greatly reduces the number of contracts deployed on each chain, and increases batching opportunities.
  - Faster chain expansion due to the reduced number of contracts.
  - Cross-source message execution in a single transaction.
- OnRamps have moved a lot of the dest chain specific logic to the FeeQuoter.
  - The FeeQuoter is an evolution of the PriceRegistry contract.
  - The FeeQuoter is upgradable without having to replace all On/OffRamps, meaning it's easier to add future non-EVM chains.
- Full non-EVM support, which will be rolled out in the coming months.
  - To support non-EVM without ballooning the number of extraArg tags we have renamed `EVMExtraArgsV2` to `GenericExtraArgsV2`. The tag bytes remain unchanged, meaning `EVMExtraArgsV2` and its encoding is still supported.
- A new way to permissionlessly onboard tokens to CCIP
  - In addition to the already existing owner() and getCCIPAdmin() functions, we introduce registering for OZ AccessControl.
  - A user that has the `DEFAULT_ADMIN_ROLE` on their token can now register their token on the TokenAdminRegistry.
- Upgraded the Solidity compiler to 0.8.26, with most contracts now requiring ^0.8.24.

### Minor Changes

- [#804](https://github.com/smartcontractkit/chainlink-ccip/pull/804) [`875e982`](https://github.com/smartcontractkit/chainlink-ccip/commit/875e982e6437dc126710d8224dd7c792a197bea6) - #feature move contract to chainlink-ccip

- [#15804](https://github.com/smartcontractkit/chainlink/pull/15804) [`46ef625`](https://github.com/smartcontractkit/chainlink/commit/46ef62537ea0389a86de03465253a8629766c2c9) - #feature Add two new pool types: Siloed-LockRelease and BurnToAddress and fix bug in HybridUSDCTokenPool for transferLiqudity #bugfix

- [#15811](https://github.com/smartcontractkit/chainlink/pull/15811) [`4e28497`](https://github.com/smartcontractkit/chainlink/commit/4e284976ea8ca7c0e355efce6336742d70918ac2) - Update FeeQuoter to support Solana chain families #feature

- [#14924](https://github.com/smartcontractkit/chainlink/pull/14924) [`161d227`](https://github.com/smartcontractkit/chainlink/commit/161d227575bbeca0119e5eee8e5a54cf3b4df677) - New contract for deploying, CCIP-compatible token pools and configuring with the tokenAdminRegistry, and a new ERC20 with constructor compatible with the factory's deployment pattern. #internal

- [#15099](https://github.com/smartcontractkit/chainlink/pull/15099) [`9c79488`](https://github.com/smartcontractkit/chainlink/commit/9c79488e4259bd59aa4d25b2be9c2ffd9390333d) - Add new event in setRateLimitAdmin for Atlas

- [#15737](https://github.com/smartcontractkit/chainlink/pull/15737) [`631cd8f`](https://github.com/smartcontractkit/chainlink/commit/631cd8fae34c73801e223f9ef96f23a032cf407f) - Account for tokenTransferBytesOverhead in exec cost

- [#16298](https://github.com/smartcontractkit/chainlink/pull/16298) [`ab06bf0`](https://github.com/smartcontractkit/chainlink/commit/ab06bf0277b5641c24596d05351dd23df544a72c) - Release CCIP 1.6, remove -dev suffix

- [#15123](https://github.com/smartcontractkit/chainlink/pull/15123) [`72da397`](https://github.com/smartcontractkit/chainlink/commit/72da397b9b1a8baa95129ffe635e5d60852e9ebc) - Add a new contract, BurnMintERC20, which is basically just our ERC677 implementation without the transferAndCall function. #internal

- [#14696](https://github.com/smartcontractkit/chainlink/pull/14696) [`072bfb6`](https://github.com/smartcontractkit/chainlink/commit/072bfb667a4e1f7cc0c874409ebfe6ef7f7b6cbe) - Add getNextDonId() and getNodes(bytes32[] calldata p2pIds) in CapabilitiesRegistry and define interface for node info

- [#14981](https://github.com/smartcontractkit/chainlink/pull/14981) [`a0309c9`](https://github.com/smartcontractkit/chainlink/commit/a0309c9a87ca7bcbe50db2fa272f9fab024bef13) - Skip stale price update from keystone instead of reverting

- [#15165](https://github.com/smartcontractkit/chainlink/pull/15165) [`03827b9`](https://github.com/smartcontractkit/chainlink/commit/03827b9d291cb11501162f9eafa1eba5619cabc9) - CCIP test restructuring

- [#14817](https://github.com/smartcontractkit/chainlink/pull/14817) [`974def5`](https://github.com/smartcontractkit/chainlink/commit/974def52d97ee548b7568cf2facbc556dfa0e797) - Change minSigners to f in RMNRemote/RMNHome

- [#15448](https://github.com/smartcontractkit/chainlink/pull/15448) [`c1ee7ab`](https://github.com/smartcontractkit/chainlink/commit/c1ee7ab715b524df6b580593e18f51637bd1500d) - Add supportsInterface to FeeQuoter for Keystone

- [#14877](https://github.com/smartcontractkit/chainlink/pull/14877) [`317f930`](https://github.com/smartcontractkit/chainlink/commit/317f93014d9b5deb76e5b54685b020f44be9b46e) - Fix sender encoding and comments in CCIP Any2EVMMEssage and corrected comments

- [#15504](https://github.com/smartcontractkit/chainlink/pull/15504) [`437ef64`](https://github.com/smartcontractkit/chainlink/commit/437ef640db1d4455e7b4d90868c9e5c4d62054df) - Fix gas estimation by adding a reverting clause

- [#14951](https://github.com/smartcontractkit/chainlink/pull/14951) [`2fab939`](https://github.com/smartcontractkit/chainlink/commit/2fab939ad3fbd49f79e96c177a9ffb11387f397e) - Add a missing condition for the Execution plugin in the \_afterOCR3ConfigSet function. Now, the function correctly reverts if signature verification is enabled for the Execution plugin

- [#14918](https://github.com/smartcontractkit/chainlink/pull/14918) [`1c53ec2`](https://github.com/smartcontractkit/chainlink/commit/1c53ec25ed6c6bfee37e9052236fe595eeef8d89) - ApplyRateLimiterConfigUpdates validation

- [#15983](https://github.com/smartcontractkit/chainlink/pull/15983) [`78d548d`](https://github.com/smartcontractkit/chainlink/commit/78d548d6101afd7fc82c2568a3665b4b57d5b1eb) - Add EVM extraArgs encode & decode to MessageHasher

- [#15575](https://github.com/smartcontractkit/chainlink/pull/15575) [`ef0dd1c`](https://github.com/smartcontractkit/chainlink/commit/ef0dd1c1b9c9650a3dba7b5c3d5f2e81db90a555) - Make gas for call exact check immutable

### Patch Changes

- [#804](https://github.com/smartcontractkit/chainlink-ccip/pull/804) [`875e982`](https://github.com/smartcontractkit/chainlink-ccip/commit/875e982e6437dc126710d8224dd7c792a197bea6) - Fix bug in BurnMintWithLockReleaseFlagTokenPool::releaseOrMint that would prevent token delivery #bugfix

- [#804](https://github.com/smartcontractkit/chainlink-ccip/pull/804) [`875e982`](https://github.com/smartcontractkit/chainlink-ccip/commit/875e982e6437dc126710d8224dd7c792a197bea6) - #feature add aptos support

- [#854](https://github.com/smartcontractkit/chainlink-ccip/pull/854) [`2551e92`](https://github.com/smartcontractkit/chainlink-ccip/commit/2551e92d1fda6748873b33e8a327237d6a7f0bf0) - Update chainlink-evm reference

- [#15422](https://github.com/smartcontractkit/chainlink/pull/15422) [`3cecd5f`](https://github.com/smartcontractkit/chainlink/commit/3cecd5f7dd5f8eaa624eafd8db701475facb8617) - Add legacy fallback to RMN

- [#14798](https://github.com/smartcontractkit/chainlink/pull/14798) [`26e22eb`](https://github.com/smartcontractkit/chainlink/commit/26e22eb6cfc320d981c91b1bfeb87c3c645f10d6) - Fixing comments and data availability bytes length calculation

- [#15829](https://github.com/smartcontractkit/chainlink/pull/15829) [`6e65dee`](https://github.com/smartcontractkit/chainlink/commit/6e65deecae053ee1e885da7ce6d1d308364ced1d) - Generate gethwrappers through Foundry instead of solc-select via python

- [#16201](https://github.com/smartcontractkit/chainlink/pull/16201) [`8ab88c2`](https://github.com/smartcontractkit/chainlink/commit/8ab88c2d9d5e5df4f8496763e74b0bb9c4c183a9) - Enable both blessed and unblessed roots in a single commit report

- [#15523](https://github.com/smartcontractkit/chainlink/pull/15523) [`baa225e`](https://github.com/smartcontractkit/chainlink/commit/baa225e7614f504a70cb51a8b0d0a98402971268) - Remove legacy curse check from RMNRemote isCursed() method #bugfix

- [#16226](https://github.com/smartcontractkit/chainlink/pull/16226) [`b92a304`](https://github.com/smartcontractkit/chainlink/commit/b92a304b55570fdeb87a4f3e840819d1cf33f043) - Add Support for SVM ATA to USDCTokenPool and send correct tokenReceiver to token pools

- [#16017](https://github.com/smartcontractkit/chainlink/pull/16017) [`17a9e2a`](https://github.com/smartcontractkit/chainlink/commit/17a9e2af202a313ea734a556af3f4460d3e8c795) - Decouple LiquidityManager tests with LockReleaseTokenPool + Test Rename

- [#15458](https://github.com/smartcontractkit/chainlink/pull/15458) [`5a3a99b`](https://github.com/smartcontractkit/chainlink/commit/5a3a99b7982dbf0f8aa57654dac356419736bd30) - Add token address to TokenHandlingError

- [#16141](https://github.com/smartcontractkit/chainlink/pull/16141) [`16a7985`](https://github.com/smartcontractkit/chainlink/commit/16a7985836d2055aca62d4cd331f2d374792d0d1) - Move multiple weth9 implementations to vendor

- [#14805](https://github.com/smartcontractkit/chainlink/pull/14805) [`b17c09d`](https://github.com/smartcontractkit/chainlink/commit/b17c09d252f75071f6ec54b3389257c1f27df9b2) - Gas optimizations and comment cleanup

- [#15904](https://github.com/smartcontractkit/chainlink/pull/15904) [`5314b41`](https://github.com/smartcontractkit/chainlink/commit/5314b4127404aa2f402cd396c3088923136ac9d0) - Add EIP-7623 support

- [#15357](https://github.com/smartcontractkit/chainlink/pull/15357) [`18cb44e`](https://github.com/smartcontractkit/chainlink/commit/18cb44e891a00edff7486640ffc8e0c9275a04f8) - New function to CCIPReaderTester getLatestPriceSequenceNumber

- [#15067](https://github.com/smartcontractkit/chainlink/pull/15067) [`eeb58e2`](https://github.com/smartcontractkit/chainlink/commit/eeb58e2d3ae5d84826b31eaf805b30f722a8e87d) - Adds OZ AccessControl support to the registry module

- [#14972](https://github.com/smartcontractkit/chainlink/pull/14972) [`6db71d3`](https://github.com/smartcontractkit/chainlink/commit/6db71d32d756eba147ec69f385252aa51589d517) - Minor nits, allow Router updates even when the offRamp has been used. Remove getRouter from onRamp

- [#15293](https://github.com/smartcontractkit/chainlink/pull/15293) [`4665863`](https://github.com/smartcontractkit/chainlink/commit/466586309a8cbbfc1c793ff1021b7fcd3522dd3e) - Allow multiple remote pools per chain selector

- [#14922](https://github.com/smartcontractkit/chainlink/pull/14922) [`42db9fd`](https://github.com/smartcontractkit/chainlink/commit/42db9fd17ca32d554f8bc9d0c6aab01e5c0e8c81) - Minor fixes to formatting, pragma, imports, etc. for Hybrid USDC Token Pools #bugfix

- [#14969](https://github.com/smartcontractkit/chainlink/pull/14969) [`ccd9956`](https://github.com/smartcontractkit/chainlink/commit/ccd9956ac6cec25770c93e57e283aa2a5ebb6737) - Remove rawVs from RMNRemote

- [#14845](https://github.com/smartcontractkit/chainlink/pull/14845) [`3f955bf`](https://github.com/smartcontractkit/chainlink/commit/3f955bfde18bbad19f7195da1da179d04275a873) - Minor fixes and changing the order of removes/adds in feeToken config

- [#14809](https://github.com/smartcontractkit/chainlink/pull/14809) [`082e6fc`](https://github.com/smartcontractkit/chainlink/commit/082e6fc8f918dc69e9a5a2acbff6644dca74f2b1) - Modified TokenPriceFeedConfig to support tokens with zero decimals #bugfix

- [#15006](https://github.com/smartcontractkit/chainlink/pull/15006) [`33dba89`](https://github.com/smartcontractkit/chainlink/commit/33dba89a3beda1138fe7332a4947411e748fd99f) - Minor gas optimizations and input sanity checks for CCIPHome #bugfix

- [#16102](https://github.com/smartcontractkit/chainlink/pull/16102) [`57ca0fb`](https://github.com/smartcontractkit/chainlink/commit/57ca0fb8f3c74fec461e2168d362b62374e26b63) - Comment and parameter validation fixes and remove outstandingTokens from BurnToAddressMintTokenPool #bugfix

- [#16090](https://github.com/smartcontractkit/chainlink/pull/16090) [`2efb46c`](https://github.com/smartcontractkit/chainlink/commit/2efb46c470867127671d39b214297f9419b24a48) - Minor FeeQuoter audit fixes

- [#16175](https://github.com/smartcontractkit/chainlink/pull/16175) [`1c76b30`](https://github.com/smartcontractkit/chainlink/commit/1c76b30f78fdb542d9f5b7ee4c7238e6b8f408d2) - Cap max accounts in svm extra args

- [#15386](https://github.com/smartcontractkit/chainlink/pull/15386) [`62c2376`](https://github.com/smartcontractkit/chainlink/commit/62c23768cd483b179301625603a785dd773f2c78) - Modify TokenPool.sol function setChainRateLimiterConfig to now accept an array of configs and set sequentially. Requested by front-end. PR issue CCIP-4329 #bugfix

- [#15984](https://github.com/smartcontractkit/chainlink/pull/15984) [`86e9119`](https://github.com/smartcontractkit/chainlink/commit/86e9119d69b15c5e83ed38edc4debe1e6ce87674) - Fix missing case in gas estimation logic

- [#15605](https://github.com/smartcontractkit/chainlink/pull/15605) [`8c65527`](https://github.com/smartcontractkit/chainlink/commit/8c65527c82a20c74b2a4707221ef496802b21804) - Replace f with fObserve in RMNHome and RMNRemote and update all tests CCIP-4058

- [#15405](https://github.com/smartcontractkit/chainlink/pull/15405) [`16f1529`](https://github.com/smartcontractkit/chainlink/commit/16f1529856a575c8d2091a16033a8d0371408d96) - Enable via-ir in CCIP compilation

- [#14960](https://github.com/smartcontractkit/chainlink/pull/14960) [`1b1dc3b`](https://github.com/smartcontractkit/chainlink/commit/1b1dc3b8ee058901828963a0b9e59fc239444e93) - CCIP-3789 Add check on MultiAggregateRateLimiter:UpdateRateLimitTokens that remote token is not abi.encode(address(0)) #bugfix

- [#15020](https://github.com/smartcontractkit/chainlink/pull/15020) [`9ec788e`](https://github.com/smartcontractkit/chainlink/commit/9ec788e78b4fcf3266b5e0a9ed2e166cea51388f) - More efficient ownership usage

- [#14734](https://github.com/smartcontractkit/chainlink/pull/14734) [`ca71878`](https://github.com/smartcontractkit/chainlink/commit/ca71878aa5a55fe239a456d7b564ffeba9bc84d7) - Make stalenessThreshold per dest chain and have 0 mean no staleness check.

- [#16299](https://github.com/smartcontractkit/chainlink/pull/16299) [`3b89c46`](https://github.com/smartcontractkit/chainlink/commit/3b89c464872609a87a172d93175f3314cb8558f6) - Additional security and parameter checks and comment fixes

- [#15743](https://github.com/smartcontractkit/chainlink/pull/15743) [`4a19318`](https://github.com/smartcontractkit/chainlink/commit/4a19318efa56c079da53777f82404a1fbb24479a) - Create a new version of the ERC165Checker library which checks for sufficient gas before making an external call to prevent message delivery issues. #bugfix

- [#15301](https://github.com/smartcontractkit/chainlink/pull/15301) [`6c4f1b9`](https://github.com/smartcontractkit/chainlink/commit/6c4f1b920c64b9c066b19119bb5990f0bb0714b0) - Refactor MockCCIPRouter to support EVMExtraArgsV2

- [#14863](https://github.com/smartcontractkit/chainlink/pull/14863) [`84bcbe0`](https://github.com/smartcontractkit/chainlink/commit/84bcbe03ebfa3ab7f58b97897eb0c55b45191859) - Change else if to else..if..

- [#15570](https://github.com/smartcontractkit/chainlink/pull/15570) [`c1341a5`](https://github.com/smartcontractkit/chainlink/commit/c1341a5081d098bce04a7564a6525a91f2beeecf) - Add getChainConfig to ccipHome

## 1.5.0

### Minor Changes

- [#14266](https://github.com/smartcontractkit/chainlink/pull/14266) [`c323e0d`](https://github.com/smartcontractkit/chainlink/commit/c323e0d600c659a4ea584dbae0a0db187afd51eb) Thanks [@asoliman92](https://github.com/asoliman92)! - Move latest ccip contracts code from ccip repo to chainlink repo
- [#13941](https://github.com/smartcontractkit/chainlink/pull/13941) [`9e74eee`](https://github.com/smartcontractkit/chainlink/commit/9e74eee9d415b386db33bdf2dd44facc82cd3551) Thanks [@RensR](https://github.com/RensR)! - Add ccip contracts to the repo

### Patch Changes

- [#14345](https://github.com/smartcontractkit/chainlink/pull/14345) [`c83c687`](https://github.com/smartcontractkit/chainlink/commit/c83c68735bdee6bbd8510733b7415797cd08ecbd) Thanks [@makramkd](https://github.com/makramkd)! - Merge ccip contracts
- [#14516](https://github.com/smartcontractkit/chainlink/pull/14516) [`0e32c07`](https://github.com/smartcontractkit/chainlink/commit/0e32c07d22973343e722a228ff1c3b1e8f9bc04e) Thanks [@mateusz-sekara](https://github.com/mateusz-sekara)! - Adding USDCReaderTester contract for CCIP integration tests #internal
- [#14739](https://github.com/smartcontractkit/chainlink/pull/14739) [`4842271`](https://github.com/smartcontractkit/chainlink/commit/4842271b0f7054f5f1364c59d3d9da534c5d4f25) Thanks [@RensR](https://github.com/RensR)! - remove CCIP 1.5
