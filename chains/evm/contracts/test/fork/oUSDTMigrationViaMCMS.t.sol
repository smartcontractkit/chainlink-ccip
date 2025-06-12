pragma solidity ^0.8.24;


import {IOwner} from "../../interfaces/IOwner.sol";
import {Router} from "../../Router.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {Pool} from "../../libraries/Pool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";import {Client} from "../../libraries/Client.sol";

import {Test} from "forge-std/Test.sol";
import {console2 as console} from "forge-std/console2.sol";

contract MCMS {
    struct Call {
        address target;
        uint256 value;
        bytes data;
    }

    function scheduleBatch(
        Call[] calldata calls,
        bytes32 predecessor,
        bytes32 salt,
        uint256 delay
    ) external {}
}


struct StaticConfig {
    address linkToken; // ────────╮ Link token address
    uint64 chainSelector; // ─────╯ Source chainSelector
    uint64 destChainSelector; // ─╮ Destination chainSelector
    uint64 defaultTxGasLimit; //  │ Default gas limit for a tx
    uint96 maxNopFeesJuels; // ───╯ Max nop fee balance onramp can have
    address prevOnRamp; //          Address of previous-version OnRamp
    address rmnProxy; //            Address of RMN proxy
    address tokenAdminRegistry; //  Address of the token admin registry
}


interface IEVM2EVMOnRamp {
    function getStaticConfig() external view returns (StaticConfig memory);
}

contract LinkMigrationViaMCMSTest is Test {
    struct ChainConfig {
        uint256 forkId;
        bytes payload;
        address router;
        address linkToken;
        uint64 chainSelector;
        address tokenAdminRegistry;
    }

    event PoolSet(address indexed localToken, address indexed previousPool, address indexed newPool);

    ChainConfig private ethereum;

    address public constant USDT = 0xdAC17F958D2ee523a2206206994597C13D831ec7;

    address expectedNewPool = 0xa3532633401AbFfbd15e6be825a45FB7F141469B;
    address oldPool = 0xCe19f75BCE7Fb74c9e2328766Ebe50465df24CA3;

    function setUp() public {
        ethereum = ChainConfig(
            vm.createSelectFork(vm.envString("ETHEREUM_RPC_URL")),
            vm.envBytes("ETHEREUM_PAYLOAD"),
            0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D,
            0x514910771AF9Ca656af840dff83E8264EcF986CA,
            5009297550715157269,
            0xb22764f98dD05c789929716D677382Df22C05Cb6
        );
    }

    function testUSDTMigration() public {
        uint256 preBalance = IERC20(USDT).balanceOf(expectedNewPool);

        address to = 0x44835bBBA9D40DEDa9b64858095EcFB2693c9449;
        address multisig = 0xE53289F32c8E690b7173aA33affE9B6B0CB0012F;
        uint256 amount = 1000000000000;
        
        vm.startPrank(multisig);

        vm.expectEmit();
        emit TokenAdminRegistry.PoolSet(USDT, oldPool, expectedNewPool);

        (bool success, ) = to.call{value: 0}(ethereum.payload);
        require(success, "RBACTimelock: underlying transaction reverted");
        
        TokenAdminRegistry.TokenConfig memory mainnetCfg = TokenAdminRegistry(ethereum.tokenAdminRegistry).getTokenConfig(USDT);

        assertEq(mainnetCfg.tokenPool, expectedNewPool, "Token pool should be the expected new pool");

        console.log("mainnetCfg.tokenPool: %s", mainnetCfg.tokenPool);

        uint256 postBalance = IERC20(USDT).balanceOf(expectedNewPool);
        
        assertEq(postBalance, preBalance + amount, "USDT balance should be greater than pre-balance by 1M USDT");
    }
       
}