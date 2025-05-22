pragma solidity ^0.8.24;

import {BurnMintTokenPool} from "./BurnMintTokenPool.sol";
import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";

contract BurnMintFastTransferTokenPool is BurnMintTokenPool, FastTransferTokenPoolAbstract {
    constructor(
        address token,
        uint8 localTokenDecimals,
        address[] memory allowlist,
        address rmnProxy,
        address router
    ) BurnMintTokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {} 

    function _handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) internal override {
        if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(lockOrBurnIn.remoteChainSelector)))) revert CursedByRMN();
        _checkAllowList(sender);
        if (!isSupportedChain(destinationChainSelector)) revert ChainNotAllowed(destinationChainSelector);
        _consumeOutboundRateLimit(destinationChainSelector, amount);
        _burn(amount);
    }

    function _transferFromFiller(address filler, address receiver, uint256 amount) internal override {
        _consumeInboundRateLimit(filler, amount);
        getToken().safeTransferFrom(filler, receiver, amount);
    }
    function _settle(uint64 sourceChainSelector, address receiver, uint256 amount, bool shouldConsumeRateLimit) internal override {
        if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(sourceChainSelector)))) revert CursedByRMN();
        // Validates that the source pool address is configured on this pool.
        if (!isRemotePool(sourceChainSelector, sourcePoolAddress)) {
            revert InvalidSourcePoolAddress(sourcePoolAddress);
        }
        if (shouldConsumeRateLimit) {
            _consumeInboundRateLimit(sourceChainSelector, amount);
        }
        IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, localAmount);
    }

    function _checkAdmin() internal view onlyOwner override {
    }
}