# SequenceFeeQuoterInputCreation Mapping Document

## From v1.6.0 (`CreateFeeQuoterUpdateInputFromV160`)

**Source Contract**: FeeQuoter v1.6.3  
**Target Contract**: FeeQuoter v1.7.0

### Field Mapping
| Target Field (FeeQuoter v1.7.0) | Source Field (FeeQuoter v1.6.3) | Notes |
|----------------------------------|----------------------------------|-------|
| `ConstructorArgs.StaticConfig.LinkToken` | `StaticCfg.LinkToken` | Direct copy |
| `ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg` | `StaticCfg.MaxFeeJuelsPerMsg` | Direct copy |
| `ConstructorArgs.PriceUpdaters` or `AuthorizedCallerUpdates.AddedCallers` | `PriceUpdaters` | Depends on new deployment vs update |
| `DestChainConfig.IsEnabled` | `RemoteChainCfgs[chain].DestChainCfg.IsEnabled` | Direct copy |
| `DestChainConfig.MaxDataBytes` | `RemoteChainCfgs[chain].DestChainCfg.MaxDataBytes` | Direct copy |
| `DestChainConfig.MaxPerMsgGasLimit` | `RemoteChainCfgs[chain].DestChainCfg.MaxPerMsgGasLimit` | Direct copy |
| `DestChainConfig.DestGasOverhead` | `RemoteChainCfgs[chain].DestChainCfg.DestGasOverhead` | Direct copy |
| `DestChainConfig.DestGasPerPayloadByteBase` | `RemoteChainCfgs[chain].DestChainCfg.DestGasPerPayloadByteBase` | Direct copy |
| `DestChainConfig.ChainFamilySelector` | `RemoteChainCfgs[chain].DestChainCfg.ChainFamilySelector` | Direct copy |
| `DestChainConfig.DefaultTokenFeeUSDCents` | `RemoteChainCfgs[chain].DestChainCfg.DefaultTokenFeeUSDCents` | Direct copy |
| `DestChainConfig.DefaultTokenDestGasOverhead` | `RemoteChainCfgs[chain].DestChainCfg.DefaultTokenDestGasOverhead` | Direct copy |
| `DestChainConfig.DefaultTxGasLimit` | `RemoteChainCfgs[chain].DestChainCfg.DefaultTxGasLimit` | Direct copy |
| `DestChainConfig.NetworkFeeUSDCents` | `RemoteChainCfgs[chain].DestChainCfg.NetworkFeeUSDCents` | Cast from `uint8` to `uint16` |
| `DestChainConfig.LinkFeeMultiplierPercent` | N/A | Hardcoded to `90` |
| `TokenTransferFeeConfig.FeeUSDCents` | `RemoteChainCfgs[chain].TokenTransferFeeCfgs[token].MinFeeUSDCents` | Direct copy |
| `TokenTransferFeeConfig.DestGasOverhead` | `RemoteChainCfgs[chain].TokenTransferFeeCfgs[token].DestGasOverhead` | Direct copy |
| `TokenTransferFeeConfig.DestBytesOverhead` | `RemoteChainCfgs[chain].TokenTransferFeeCfgs[token].DestBytesOverhead` | Direct copy |
| `TokenTransferFeeConfig.IsEnabled` | `RemoteChainCfgs[chain].TokenTransferFeeCfgs[token].IsEnabled` | Direct copy |

## From v1.5.0 (`CreateFeeQuoterUpdateInputFromV150`)

**Source Contract**: EVM2EVMOnRamp v1.5.0  
**Target Contract**: FeeQuoter v1.7.0

### Field Mapping
| Target Field (FeeQuoter v1.7.0) | Source Field (EVM2EVMOnRamp v1.5.0) | Notes |
|----------------------------------|--------------------------------------|-------|
| `ConstructorArgs.StaticConfig.LinkToken` | `OnRampCfg.StaticConfig.LinkToken` | From first OnRamp (if empty) |
| `ConstructorArgs.StaticConfig.MaxFeeJuelsPerMsg` | `OnRampCfg.StaticConfig.MaxNopFeesJuels` | From first OnRamp (if empty) |
| `ConstructorArgs.PriceUpdaters` | N/A | Empty array `[]` (TODO: what to do with price updaters for 1.5 if there is no 1.6 lanes here) |
| `DestChainConfig.IsEnabled` | N/A | Hardcoded to `true` (if chain is supported on OnRamp, enable it on FeeQuoter) |
| `DestChainConfig.MaxDataBytes` | `OnRampCfg.DynamicConfig.MaxDataBytes` | Direct copy |
| `DestChainConfig.MaxPerMsgGasLimit` | `OnRampCfg.DynamicConfig.MaxPerMsgGasLimit` | Direct copy |
| `DestChainConfig.DestGasOverhead` | `OnRampCfg.DynamicConfig.DestGasOverhead` | Direct copy |
| `DestChainConfig.DestGasPerPayloadByteBase` | `OnRampCfg.DynamicConfig.DestGasPerPayloadByte` | Cast from `uint8` |
| `DestChainConfig.ChainFamilySelector` | N/A | Hardcoded to EVM family selector `0x2812d52c` |
| `DestChainConfig.DefaultTokenFeeUSDCents` | `OnRampCfg.DynamicConfig.DefaultTokenFeeUSDCents` | Direct copy |
| `DestChainConfig.DefaultTokenDestGasOverhead` | `OnRampCfg.DynamicConfig.DefaultTokenDestGasOverhead` | Direct copy |
| `DestChainConfig.DefaultTxGasLimit` | `OnRampCfg.StaticConfig.DefaultTxGasLimit` | Cast to `uint32` |
| `DestChainConfig.NetworkFeeUSDCents` | `OnRampCfg.FeeTokenConfig[].NetworkFeeUSDCents` | From first non-zero value (same across all fee tokens) |
| `DestChainConfig.LinkFeeMultiplierPercent` | N/A | Hardcoded to `90` |
| `TokenTransferFeeConfig.FeeUSDCents` | `OnRampCfg.TokenTransferFeeConfig[token].MinFeeUSDCents` | Direct copy |
| `TokenTransferFeeConfig.DestGasOverhead` | `OnRampCfg.TokenTransferFeeConfig[token].DestGasOverhead` | Direct copy |
| `TokenTransferFeeConfig.DestBytesOverhead` | `OnRampCfg.TokenTransferFeeConfig[token].DestBytesOverhead` | Direct copy |
| `TokenTransferFeeConfig.IsEnabled` | `OnRampCfg.TokenTransferFeeConfig[token].IsEnabled` | Direct copy |
