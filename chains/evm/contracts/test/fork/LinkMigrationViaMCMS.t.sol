pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import "forge-std/console.sol";
import {IOwner} from "../../interfaces/IOwner.sol";
import {Router} from "../../Router.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {IERC20} from "@chainlink/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {Client} from "../../libraries/Client.sol";
import {Pool} from "../../libraries/Pool.sol";

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

interface ITypeAndVersion {
    function typeAndVersion() external view returns (string memory);
}

interface ILegacyPool {
    function isOffRamp(address offRamp) external view returns (bool);
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

    ChainConfig private ethereum;
    ChainConfig[12] private dests;

    address private constant OLD_LINK_POOL_ON_ETH = 0xC2291992A08eBFDfedfE248F2CCD34Da63570DF4;
    address private constant EXPECTED_ETH_REBALANCER = 0x0000000000000000000000000000000000000000;
    address private constant ETH_TIMELOCK = 0x44835bBBA9D40DEDa9b64858095EcFB2693c9449;

    function setUp() public {
        ethereum = ChainConfig(
            vm.createFork(vm.envString("ETHEREUM_RPC_URL")),
            vm.envBytes("ETHEREUM_PAYLOAD"),
            0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D,
            0x514910771AF9Ca656af840dff83E8264EcF986CA,
            5009297550715157269,
            0xb22764f98dD05c789929716D677382Df22C05Cb6
        );
        dests[0] = ChainConfig(
            vm.createFork(vm.envString("SONIC_RPC_URL")),
            vm.envBytes("SONIC_PAYLOAD"),
            0xB4e1Ff7882474BB93042be9AD5E1fA387949B860,
            0x71052BAe71C25C78E37fD12E5ff1101A71d9018F,
            1673871237479749969,
            0x2961Cb47b5111F38d75f415c21ceB4120ddd1b69
        );
        dests[1] = ChainConfig(
            vm.createFork(vm.envString("SONEIUM_RPC_URL")),
            vm.envBytes("SONEIUM_PAYLOAD"),
            0x8C8B88d827Fe14Df2bc6392947d513C86afD6977,
            0x32D8F819C8080ae44375F8d383Ffd39FC642f3Ec,
            12505351618335765396,
            0x5ba21F6824400B91F232952CA6d7c8875C1755a4
        );
        dests[2] = ChainConfig(
            vm.createFork(vm.envString("SCROLL_RPC_URL")),
            vm.envBytes("SCROLL_PAYLOAD"),
            0x9a55E8Cab6564eb7bbd7124238932963B8Af71DC,
            0x548C6944cba02B9D1C0570102c89de64D258d3Ac,
            13204309965629103672,
            0x846dEA1c1706FC35b4aa78B32d31F1599DAA47b4
        );
        dests[3] = ChainConfig(
            vm.createFork(vm.envString("CELO_RPC_URL")),
            vm.envBytes("CELO_PAYLOAD"),
            0xfB48f15480926A4ADf9116Dca468bDd2EE6C5F62,
            0xd07294e6E917e07dfDcee882dd1e2565085C2ae0,
            1346049177634351622,
            0xf19e0555fAA9051e277eeD5A0DcdB13CDaca39a9
        );
        dests[4] = ChainConfig(
            vm.createFork(vm.envString("BASE_RPC_URL")),
            vm.envBytes("BASE_PAYLOAD"),
            0x881e3A65B4d4a04dD529061dd0071cf975F58bCD,
            0x88Fb150BDc53A65fe94Dea0c9BA0a6dAf8C6e196,
            15971525489660198786,
            0x6f6C373d09C07425BaAE72317863d7F6bb731e37
        );
        dests[5] = ChainConfig(
            vm.createFork(vm.envString("BLAST_RPC_URL")),
            vm.envBytes("BLAST_PAYLOAD"),
            0x12e0B8E349C6fb7E6E40713E8125C3cF1127ea8C,
            0x93202eC683288a9EA75BB829c6baCFb2BfeA9013,
            4411394078118774322,
            0x846Fccd01D4115FD1E81267495773aeB33bF1dC7
        );
        dests[6] = ChainConfig(
            vm.createFork(vm.envString("ASTAR_RPC_URL")),
            vm.envBytes("ASTAR_PAYLOAD"),
            0x8D5c5CB8ec58285B424C93436189fB865e437feF,
            0x31EFB841d5e0b4082F7E1267dab8De1b853f2A9d,
            6422105447186081193,
            0xB98eEd70e3cE8E342B0f770589769E3A6bc20A09
        );
        dests[7] = ChainConfig(
            vm.createFork(vm.envString("RONIN_RPC_URL")),
            vm.envBytes("RONIN_PAYLOAD"),
            0x46527571D5D1B68eE7Eb60B18A32e6C60DcEAf99,
            0x3902228D6A3d2Dc44731fD9d45FeE6a61c722D0b,
            6916147374840168594,
            0x90e83d532A4aD13940139c8ACE0B93b0DdbD323a
        );
        dests[8] = ChainConfig(
            vm.createFork(vm.envString("MODE_RPC_URL")),
            vm.envBytes("MODE_PAYLOAD"),
            0x24C40f13E77De2aFf37c280BA06c333531589bf1,
            0x183E3691EfF3524B2315D3703D94F922CbE51F54,
            7264351850409363825,
            0xB4b40c010A547dff6A22d94bC2C1c1e745b62aB2
        );
        dests[9] = ChainConfig(
            vm.createFork(vm.envString("KROMA_RPC_URL")),
            vm.envBytes("KROMA_PAYLOAD"),
            0xE93E8B0d1b1CEB44350C8758ed1E2799CCee31aB,
            0xC1F6f7622ad37C3f46cDF6F8AA0344ADE80BF450,
            3719320017875267166,
            0x447066676A5413682a881c63aed0F03f8ACf7E45
        );
        dests[10] = ChainConfig(
            vm.createFork(vm.envString("WEMIX_RPC_URL")),
            vm.envBytes("WEMIX_PAYLOAD"),
            0x7798b795Fde864f4Cd1b124a38Ba9619B7F8A442,
            0x80f1FcdC96B55e459BF52b998aBBE2c364935d69,
            9284632837123596123,
            0xE993e046AC50659800a91Bab0bd2daBF59CbD171
        );
        dests[11] = ChainConfig(
            vm.createFork(vm.envString("METIS_RPC_URL")),
            vm.envBytes("METIS_PAYLOAD"),
            0x7b9FB8717D306e2e08ce2e1Efa81F026bf9AD13c,
            0xd2FE54D1E5F568eB710ba9d898Bf4bD02C7c0353,
            8805746078405598895,
            0x3af897541eB03927c7431bF68884A6C2C23b683f
        );

        /*
        zkSync fork isn't working for me, reverting at IOwner(call.target).owner()

        dests[12] = ChainConfig(
            vm.createFork(vm.envString("ZKSYNC_RPC_URL")),
            vm.envBytes("ZKSYNC_PAYLOAD"),
            0x748Fd769d81F5D94752bf8B0875E9301d0ba71bB,
            0x52869bae3E091e36b0915941577F2D47d8d8B534,
            1562403441176082196,
            0x100a47C9DB342884E3314B91cec076BbAC8e619c
        );
        */
    }

    function testLinkMigration() public {
        _executeProposalOnTimeLock(ethereum);
        _checkMainnetSetup();
        for (uint256 i = 0; i < dests.length; ++i) {
            _executeProposalOnTimeLock(dests[i]);
            _checkDestSetup(i);
        }
    }

    function _executeProposalOnTimeLock(ChainConfig memory chain) internal {
        vm.selectFork(chain.forkId);
        (MCMS.Call[] memory calls, , , ) = abi.decode(
            chain.payload,
            (MCMS.Call[], bytes32, bytes32, uint256)
        );
        for (uint256 i = 0; i < calls.length; ++i) {
            MCMS.Call memory call = calls[i];
            vm.startPrank(IOwner(call.target).owner());
            (bool success, ) = call.target.call{value: call.value}(call.data);
            require(success, "RBACTimelock: underlying transaction reverted");
            vm.stopPrank();
        }
    }

    function _checkMainnetSetup() internal {
        vm.selectFork(ethereum.forkId);
        TokenAdminRegistry.TokenConfig memory mainnetCfg = TokenAdminRegistry(ethereum.tokenAdminRegistry).getTokenConfig(ethereum.linkToken);
        assertEq(
            ITypeAndVersion(mainnetCfg.tokenPool).typeAndVersion(),
            "LockReleaseTokenPool 1.5.1",
            "Token pool type and version should be LockReleaseTokenPool 1.5.1"
        );
        assertEq(
            IERC20(ethereum.linkToken).balanceOf(mainnetCfg.tokenPool),
            100_000_000000000000000000,
            "Token pool balance should be 100000 LINK"
        );
        assertEq(
            LockReleaseTokenPool(mainnetCfg.tokenPool).getRebalancer(),
            EXPECTED_ETH_REBALANCER,
            "Token pool rebalancer should be expected ETH rebalancer"
        );
        assertEq(
            ILegacyPool(OLD_LINK_POOL_ON_ETH).isOffRamp(ETH_TIMELOCK),
            false,
            "ETH timelock should not be an off ramp by end of migration"
        );
        uint64[] memory remoteChainSelectors = TokenPool(mainnetCfg.tokenPool).getSupportedChains();
        for (uint256 j = 0; j < remoteChainSelectors.length; ++j) {
            assertEq(TokenPool(mainnetCfg.tokenPool).getRemotePools(remoteChainSelectors[j]).length, 2, "There should be two remote pools");
        }
    }

    function _checkDestSetup(uint256 i) internal {
        ChainConfig memory dest = dests[i];
        vm.selectFork(dest.forkId);
        TokenAdminRegistry.TokenConfig memory destCfg = TokenAdminRegistry(dest.tokenAdminRegistry).getTokenConfig(dest.linkToken);
        assertEq(
            ITypeAndVersion(destCfg.tokenPool).typeAndVersion(),
            "BurnMintTokenPool 1.5.1",
            "Token pool type and version should be BurnMintTokenPool 1.5.1"
        );
        uint64[] memory remoteChainSelectors = TokenPool(destCfg.tokenPool).getSupportedChains();
        for (uint256 j = 0; j < remoteChainSelectors.length; ++j) {
            assertEq(TokenPool(destCfg.tokenPool).getRemotePools(remoteChainSelectors[j]).length, 2, "There should be two remote pools");
        }
    }
}
