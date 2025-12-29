// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPoolV1} from "./IPool.sol";
import {IPoolV2} from "./IPoolV2.sol";

interface IPoolV1V2 is IPoolV1, IPoolV2 {}
