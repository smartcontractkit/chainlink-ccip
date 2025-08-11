// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";

import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampHelper} from "../../helpers/OnRampHelper.sol";
import {VerifierProxySetup} from "./VerifierProxySetup.t.sol";

contract OnRamp_constructor is VerifierProxySetup {}
