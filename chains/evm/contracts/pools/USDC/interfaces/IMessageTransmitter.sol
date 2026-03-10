// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright (c) 2022, Circle Internet Financial Limited.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
pragma solidity ^0.8.0;

interface IMessageTransmitter {
  /// @notice Unlocks USDC tokens on the destination chain through CCTPv2.
  /// @dev Message format.
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * sourceDomain               4          uint32     4
  ///     * destinationDomain          4          uint32     8
  ///     * nonce                      32         bytes32   12
  ///     * sender                     32         bytes32   44
  ///     * recipient                  32         bytes32   76
  ///     * destinationCaller          32         bytes32   108
  ///     * minFinalityThreshold       4          uint32    140
  ///     * finalityThresholdExecuted  4          uint32    144
  ///     * messageBody                dynamic    bytes     148
  /// @dev CCTP burn message body format.
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * burnToken                  32         bytes32    4
  ///     * mintRecipient              32         bytes32    36
  ///     * amount                     32         uint256    68
  ///     * messageSender              32         bytes32    100
  ///     * maxFee                     32         uint256    132
  ///     * feeExecuted                32         uint256    164
  ///     * expirationBlock            32         uint256    196
  ///     * hookData                   dynamic    bytes      228
  /// @dev Hook data format.
  ///     * Field                      Bytes      Type       Index
  ///     * verifierVersion            4          bytes4     0
  ///     * messageId                  32         bytes32    4
  function receiveMessage(
    bytes calldata message,
    bytes calldata attestation
  ) external returns (bool success);

  /// Returns domain of chain on which the contract is deployed.
  /// @dev immutable.
  function localDomain() external view returns (uint32);

  /// Returns message format version.
  /// @dev immutable.
  function version() external view returns (uint32);
}
