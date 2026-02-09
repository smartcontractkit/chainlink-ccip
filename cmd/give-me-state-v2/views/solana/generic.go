package solana

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sync"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/mr-tron/base58"

	"give-me-state-v2/views"
)

// Uint128 for large numbers
type Uint128 struct {
	Lo uint64
	Hi uint64
}

func (u Uint128) String() string {
	if u.Hi == 0 {
		return fmt.Sprintf("%d", u.Lo)
	}
	// For large values, show as hex or approximate
	return fmt.Sprintf("0x%x%016x", u.Hi, u.Lo)
}

// CodeVersion enum values
const (
	CodeVersionDefault = 0
	CodeVersionV1      = 1
)

func codeVersionString(v uint8) string {
	switch v {
	case CodeVersionDefault:
		return "Default"
	case CodeVersionV1:
		return "V1"
	default:
		return fmt.Sprintf("Unknown(%d)", v)
	}
}

// ===== Account Data Fetching =====

// getAccountDataRaw fetches raw account data for a Solana address (as base58 string)
func getAccountDataRaw(ctx *views.ViewContext, address string) ([]byte, map[string]any, error) {
	result := make(map[string]any)
	result["address"] = address
	result["chainSelector"] = ctx.ChainSelector

	// Convert base58 address to bytes for the call
	addrBytes := []byte(address)

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  addrBytes,
		Data:    nil,
	}

	callResult := ctx.TypedOrchestrator.Execute(call)
	if callResult.Error != nil {
		result["error"] = callResult.Error.Error()
		return nil, result, callResult.Error
	}

	// Parse the JSON response from getAccountInfo
	var accountInfo map[string]any
	if err := json.Unmarshal(callResult.Data, &accountInfo); err != nil {
		result["error"] = err.Error()
		return nil, result, err
	}

	// Check for null account (doesn't exist)
	if accountInfo["owner"] == nil {
		result["error"] = "account does not exist"
		return nil, result, fmt.Errorf("account does not exist")
	}

	// Add raw account info
	result["programOwner"] = accountInfo["owner"]
	result["lamports"] = accountInfo["lamports"]
	result["executable"] = accountInfo["executable"]

	// Get the base64 data
	dataStr, ok := accountInfo["data"].(string)
	if !ok || dataStr == "" {
		return nil, result, nil
	}

	// Decode base64
	rawData, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		result["decodeError"] = err.Error()
		return nil, result, err
	}

	return rawData, result, nil
}

// getAccountData fetches raw account data using the context's address
func getAccountData(ctx *views.ViewContext) ([]byte, map[string]any, error) {
	return getAccountDataRaw(ctx, ctx.AddressHex)
}

// fetchPDAData fetches and returns raw data for a PDA
func fetchPDAData(ctx *views.ViewContext, pda PublicKey) ([]byte, error) {
	pdaAddr := pda.String()
	addrBytes := []byte(pdaAddr)

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  addrBytes,
		Data:    nil,
	}

	callResult := ctx.TypedOrchestrator.Execute(call)
	if callResult.Error != nil {
		return nil, callResult.Error
	}

	// Parse JSON
	var accountInfo map[string]any
	if err := json.Unmarshal(callResult.Data, &accountInfo); err != nil {
		return nil, err
	}

	// Check for null
	if accountInfo["owner"] == nil {
		return nil, fmt.Errorf("PDA account does not exist: %s", pdaAddr)
	}

	// Get base64 data
	dataStr, ok := accountInfo["data"].(string)
	if !ok || dataStr == "" {
		return nil, fmt.Errorf("no data in PDA account: %s", pdaAddr)
	}

	// Decode base64
	return base64.StdEncoding.DecodeString(dataStr)
}

// ===== Borsh Decoding Helpers =====

// decodePublicKey decodes a 32-byte public key from Borsh decoder
func decodePublicKey(decoder *ag_binary.Decoder) (PublicKey, error) {
	var pk PublicKey
	if err := decoder.Decode(&pk); err != nil {
		return pk, err
	}
	return pk, nil
}

// decodeUint128 decodes a Uint128 from Borsh decoder
func decodeUint128(decoder *ag_binary.Decoder) (Uint128, error) {
	var u Uint128
	if err := decoder.Decode(&u.Lo); err != nil {
		return u, err
	}
	if err := decoder.Decode(&u.Hi); err != nil {
		return u, err
	}
	return u, nil
}

// ===== View Functions =====

// ViewGenericAccount provides a basic view for any Solana account (fallback)
func ViewGenericAccount(ctx *views.ViewContext) (map[string]any, error) {
	rawData, result, err := getAccountData(ctx)
	if err != nil {
		return result, nil
	}

	if rawData != nil {
		result["dataLength"] = len(rawData)
	}

	return result, nil
}

// ViewFeeQuoter decodes a FeeQuoter program's config PDA
func ViewFeeQuoter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "FeeQuoter"

	// Parse program address
	programPK, err := PublicKeyFromBase58(ctx.AddressHex)
	if err != nil {
		result["error"] = fmt.Sprintf("invalid program address: %v", err)
		return result, nil
	}

	// Derive the config PDA
	configPDA, _, err := FindFeeQuoterConfigPDA(programPK)
	if err != nil {
		result["error"] = fmt.Sprintf("failed to derive config PDA: %v", err)
		return result, nil
	}
	result["pda"] = configPDA.String()

	// Fetch the config PDA data
	rawData, err := fetchPDAData(ctx, configPDA)
	if err != nil {
		result["pdaError"] = err.Error()
		return result, nil
	}

	// Skip 8-byte discriminator
	if len(rawData) < 8 {
		result["decodeError"] = "data too short"
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	// Decode FeeQuoter Config fields
	var version uint8
	if err := decoder.Decode(&version); err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["version"] = version

	owner, err := decodePublicKey(decoder)
	if err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["owner"] = owner.String()

	proposedOwner, err := decodePublicKey(decoder)
	if err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["proposedOwner"] = proposedOwner.String()

	maxFeeJuels, err := decodeUint128(decoder)
	if err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["maxFeeJuelsPerMsg"] = maxFeeJuels.String()

	linkTokenMint, err := decodePublicKey(decoder)
	if err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["linkTokenMint"] = linkTokenMint.String()

	var linkTokenLocalDecimals uint8
	if err := decoder.Decode(&linkTokenLocalDecimals); err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["linkTokenLocalDecimals"] = linkTokenLocalDecimals

	onRamp, err := decodePublicKey(decoder)
	if err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["onRamp"] = onRamp.String()

	var defaultCodeVersion uint8
	if err := decoder.Decode(&defaultCodeVersion); err != nil {
		result["decodeError"] = err.Error()
		return result, nil
	}
	result["defaultCodeVersion"] = codeVersionString(defaultCodeVersion)

	// Fetch destination chain configs in parallel
	if len(ctx.AllChainSelectors) > 0 {
		destChainConfigs := fetchFeeQuoterDestChainConfigs(ctx, programPK)
		if len(destChainConfigs) > 0 {
			result["destinationChainConfig"] = destChainConfigs
		}
	}

	return result, nil
}

// fetchFeeQuoterDestChainConfigs fetches all destination chain configs for a FeeQuoter
func fetchFeeQuoterDestChainConfigs(ctx *views.ViewContext, program PublicKey) map[uint64]map[string]any {
	configs := make(map[uint64]map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSelector := range ctx.AllChainSelectors {
		// Skip self
		if chainSelector == ctx.ChainSelector {
			continue
		}

		wg.Add(1)
		go func(remote uint64) {
			defer wg.Done()

			pda, _, err := FindFeeQuoterDestChainPDA(remote, program)
			if err != nil {
				return
			}

			rawData, err := fetchPDAData(ctx, pda)
			if err != nil {
				return // Not configured
			}

			if len(rawData) < 8 {
				return
			}

			decoder := ag_binary.NewBorshDecoder(rawData[8:])

			config := make(map[string]any)
			config["pda"] = pda.String()

			// Decode DestChain struct
			var version uint8
			decoder.Decode(&version)

			var chainSel uint64
			decoder.Decode(&chainSel)

			// DestChainState
			var stateValue [28]byte
			decoder.Decode(&stateValue)
			var stateTimestamp int64
			decoder.Decode(&stateTimestamp)

			// DestChainConfig
			var isEnabled bool
			decoder.Decode(&isEnabled)
			config["isEnabled"] = isEnabled

			var laneCodeVersion uint8
			decoder.Decode(&laneCodeVersion)
			config["laneCodeVersion"] = codeVersionString(laneCodeVersion)

			var maxNumberOfTokensPerMsg uint16
			decoder.Decode(&maxNumberOfTokensPerMsg)
			config["maxNumberOfTokensPerMsg"] = maxNumberOfTokensPerMsg

			var maxDataBytes uint32
			decoder.Decode(&maxDataBytes)
			config["maxDataBytes"] = maxDataBytes

			var maxPerMsgGasLimit uint32
			decoder.Decode(&maxPerMsgGasLimit)
			config["maxPerMsgGasLimit"] = maxPerMsgGasLimit

			var destGasOverhead uint32
			decoder.Decode(&destGasOverhead)
			config["destGasOverhead"] = destGasOverhead

			var destGasPerPayloadByteBase uint32
			decoder.Decode(&destGasPerPayloadByteBase)
			config["destGasPerPayloadByteBase"] = destGasPerPayloadByteBase

			var destGasPerPayloadByteHigh uint32
			decoder.Decode(&destGasPerPayloadByteHigh)
			config["destGasPerPayloadByteHigh"] = destGasPerPayloadByteHigh

			var destGasPerPayloadByteThreshold uint32
			decoder.Decode(&destGasPerPayloadByteThreshold)
			config["destGasPerPayloadByteThreshold"] = destGasPerPayloadByteThreshold

			var destDataAvailabilityOverheadGas uint32
			decoder.Decode(&destDataAvailabilityOverheadGas)
			config["destDataAvailabilityOverheadGas"] = destDataAvailabilityOverheadGas

			var destGasPerDataAvailabilityByte uint16
			decoder.Decode(&destGasPerDataAvailabilityByte)
			config["destGasPerDataAvailabilityByte"] = destGasPerDataAvailabilityByte

			var destDataAvailabilityMultiplierBps uint16
			decoder.Decode(&destDataAvailabilityMultiplierBps)
			config["destDataAvailabilityMultiplierBps"] = destDataAvailabilityMultiplierBps

			var defaultTokenFeeUsdcents uint16
			decoder.Decode(&defaultTokenFeeUsdcents)
			config["defaultTokenFeeUSDCents"] = defaultTokenFeeUsdcents

			var defaultTokenDestGasOverhead uint32
			decoder.Decode(&defaultTokenDestGasOverhead)
			config["defaultTokenDestGasOverhead"] = defaultTokenDestGasOverhead

			var defaultTxGasLimit uint32
			decoder.Decode(&defaultTxGasLimit)
			config["defaultTxGasLimit"] = defaultTxGasLimit

			var gasMultiplierWeiPerEth uint64
			decoder.Decode(&gasMultiplierWeiPerEth)
			config["gasMultiplierWeiPerEth"] = gasMultiplierWeiPerEth

			var networkFeeUsdcents uint32
			decoder.Decode(&networkFeeUsdcents)
			config["networkFeeUSDCents"] = networkFeeUsdcents

			var gasPriceStalenessThreshold uint32
			decoder.Decode(&gasPriceStalenessThreshold)
			config["gasPriceStalenessThreshold"] = gasPriceStalenessThreshold

			var enforceOutOfOrder bool
			decoder.Decode(&enforceOutOfOrder)
			config["enforceOutOfOrder"] = enforceOutOfOrder

			var chainFamilySelector [4]byte
			decoder.Decode(&chainFamilySelector)
			config["chainFamilySelector"] = fmt.Sprintf("%x", chainFamilySelector)

			mu.Lock()
			configs[remote] = config
			mu.Unlock()
		}(chainSelector)
	}

	wg.Wait()
	return configs
}

// ViewRouter decodes a Router program's config PDA
func ViewRouter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "Router"

	programPK, err := PublicKeyFromBase58(ctx.AddressHex)
	if err != nil {
		result["error"] = fmt.Sprintf("invalid program address: %v", err)
		return result, nil
	}

	configPDA, _, err := FindRouterConfigPDA(programPK)
	if err != nil {
		result["error"] = fmt.Sprintf("failed to derive config PDA: %v", err)
		return result, nil
	}
	result["pda"] = configPDA.String()

	rawData, err := fetchPDAData(ctx, configPDA)
	if err != nil {
		result["pdaError"] = err.Error()
		return result, nil
	}

	if len(rawData) < 8 {
		result["decodeError"] = "data too short"
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	// Router Config fields
	var defaultCodeVersion uint8
	decoder.Decode(&defaultCodeVersion)
	result["defaultCodeVersion"] = codeVersionString(defaultCodeVersion)

	var svmChainSelector uint64
	decoder.Decode(&svmChainSelector)
	result["svmChainSelector"] = svmChainSelector

	owner, _ := decodePublicKey(decoder)
	result["owner"] = owner.String()

	proposedOwner, _ := decodePublicKey(decoder)
	result["proposedOwner"] = proposedOwner.String()

	feeQuoter, _ := decodePublicKey(decoder)
	result["feeQuoter"] = feeQuoter.String()

	rmnRemote, _ := decodePublicKey(decoder)
	result["rmnRemote"] = rmnRemote.String()

	linkTokenMint, _ := decodePublicKey(decoder)
	result["linkTokenMint"] = linkTokenMint.String()

	feeAggregator, _ := decodePublicKey(decoder)
	result["feeAggregator"] = feeAggregator.String()

	// Fetch destination chain configs
	if len(ctx.AllChainSelectors) > 0 {
		destChainConfigs := fetchRouterDestChainConfigs(ctx, programPK)
		if len(destChainConfigs) > 0 {
			result["destinationChainConfig"] = destChainConfigs
		}
	}

	return result, nil
}

// fetchRouterDestChainConfigs fetches destination chain configs for Router
func fetchRouterDestChainConfigs(ctx *views.ViewContext, program PublicKey) map[uint64]map[string]any {
	configs := make(map[uint64]map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSelector := range ctx.AllChainSelectors {
		if chainSelector == ctx.ChainSelector {
			continue
		}

		wg.Add(1)
		go func(remote uint64) {
			defer wg.Done()

			pda, _, err := FindDestChainStatePDA(remote, program)
			if err != nil {
				return
			}

			rawData, err := fetchPDAData(ctx, pda)
			if err != nil {
				return
			}

			if len(rawData) < 8 {
				return
			}

			decoder := ag_binary.NewBorshDecoder(rawData[8:])
			config := make(map[string]any)
			config["pda"] = pda.String()

			// DestChain struct
			var version uint8
			decoder.Decode(&version)

			var chainSel uint64
			decoder.Decode(&chainSel)

			// State fields...
			var sequenceNumber uint64
			decoder.Decode(&sequenceNumber)

			// Config
			var laneCodeVersion uint8
			decoder.Decode(&laneCodeVersion)
			config["laneCodeVersion"] = codeVersionString(laneCodeVersion)

			var allowListEnabled bool
			decoder.Decode(&allowListEnabled)
			config["allowListEnabled"] = allowListEnabled

			// AllowedSenders is a Vec<Pubkey> - skip for now
			config["allowedSenders"] = []string{}

			mu.Lock()
			configs[remote] = config
			mu.Unlock()
		}(chainSelector)
	}

	wg.Wait()
	return configs
}

// ViewOffRamp decodes an OffRamp program's config PDA
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "OffRamp"

	programPK, err := PublicKeyFromBase58(ctx.AddressHex)
	if err != nil {
		result["error"] = fmt.Sprintf("invalid program address: %v", err)
		return result, nil
	}

	configPDA, _, err := FindOffRampConfigPDA(programPK)
	if err != nil {
		result["error"] = fmt.Sprintf("failed to derive config PDA: %v", err)
		return result, nil
	}
	result["pda"] = configPDA.String()

	rawData, err := fetchPDAData(ctx, configPDA)
	if err != nil {
		result["pdaError"] = err.Error()
		return result, nil
	}

	if len(rawData) < 8 {
		result["decodeError"] = "data too short"
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	var defaultCodeVersion uint8
	decoder.Decode(&defaultCodeVersion)
	result["defaultCodeVersion"] = defaultCodeVersion

	var svmChainSelector uint64
	decoder.Decode(&svmChainSelector)
	result["svmChainSelector"] = svmChainSelector

	var enableManualExecutionAfter int64
	decoder.Decode(&enableManualExecutionAfter)
	result["enableManualExecutionAfter"] = enableManualExecutionAfter

	owner, _ := decodePublicKey(decoder)
	result["owner"] = owner.String()

	proposedOwner, _ := decodePublicKey(decoder)
	result["proposedOwner"] = proposedOwner.String()

	// Fetch reference addresses
	refAddrPDA, _, _ := FindOffRampReferenceAddressesPDA(programPK)
	if refData, err := fetchPDAData(ctx, refAddrPDA); err == nil && len(refData) > 8 {
		refDecoder := ag_binary.NewBorshDecoder(refData[8:])

		refAddrs := make(map[string]any)
		refAddrs["pda"] = refAddrPDA.String()

		var refVersion uint8
		refDecoder.Decode(&refVersion)
		refAddrs["version"] = refVersion

		router, _ := decodePublicKey(refDecoder)
		refAddrs["router"] = router.String()

		feeQuoter, _ := decodePublicKey(refDecoder)
		refAddrs["feeQuoter"] = feeQuoter.String()

		offrampLookupTable, _ := decodePublicKey(refDecoder)
		refAddrs["offrampLookupTable"] = offrampLookupTable.String()

		rmnRemote, _ := decodePublicKey(refDecoder)
		refAddrs["rmnRemote"] = rmnRemote.String()

		result["referenceAddresses"] = refAddrs
	}

	// Fetch source chain configs
	if len(ctx.AllChainSelectors) > 0 {
		sourceChainConfigs := fetchOffRampSourceChainConfigs(ctx, programPK)
		if len(sourceChainConfigs) > 0 {
			result["sourceChains"] = sourceChainConfigs
		}
	}

	return result, nil
}

// fetchOffRampSourceChainConfigs fetches source chain configs for OffRamp
func fetchOffRampSourceChainConfigs(ctx *views.ViewContext, program PublicKey) map[uint64]map[string]any {
	configs := make(map[uint64]map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSelector := range ctx.AllChainSelectors {
		if chainSelector == ctx.ChainSelector {
			continue
		}

		wg.Add(1)
		go func(remote uint64) {
			defer wg.Done()

			pda, _, err := FindOffRampSourceChainPDA(remote, program)
			if err != nil {
				return
			}

			rawData, err := fetchPDAData(ctx, pda)
			if err != nil {
				return
			}

			if len(rawData) < 8 {
				return
			}

			decoder := ag_binary.NewBorshDecoder(rawData[8:])
			config := make(map[string]any)
			config["pda"] = pda.String()

			// SourceChain struct
			var version uint8
			decoder.Decode(&version)

			var chainSel uint64
			decoder.Decode(&chainSel)

			// State fields
			var minSeqNr uint64
			decoder.Decode(&minSeqNr)

			// Config
			var isEnabled bool
			decoder.Decode(&isEnabled)
			config["isEnabled"] = isEnabled

			var isRmnVerificationDisabled bool
			decoder.Decode(&isRmnVerificationDisabled)
			config["isRmnVerificationDisabled"] = isRmnVerificationDisabled

			var laneCodeVersion uint8
			decoder.Decode(&laneCodeVersion)
			config["laneCodeVersion"] = codeVersionString(laneCodeVersion)

			// OnRamp is a variable-length address
			var onRampLen uint8
			decoder.Decode(&onRampLen)
			if onRampLen > 0 && onRampLen <= 64 {
				onRampBytes := make([]byte, onRampLen)
				decoder.Decode(&onRampBytes)
				// Convert to appropriate format based on chain family
				if onRampLen == 20 {
					// EVM address
					config["onRamp"] = fmt.Sprintf("0x%x", onRampBytes)
				} else if onRampLen == 32 {
					// Solana address
					config["onRamp"] = base58.Encode(onRampBytes)
				} else {
					config["onRamp"] = fmt.Sprintf("0x%x", onRampBytes)
				}
			}

			mu.Lock()
			configs[remote] = config
			mu.Unlock()
		}(chainSelector)
	}

	wg.Wait()
	return configs
}

// ViewRMNRemote decodes an RMNRemote config account
func ViewRMNRemote(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "RMNRemote"

	programPK, err := PublicKeyFromBase58(ctx.AddressHex)
	if err != nil {
		result["error"] = fmt.Sprintf("invalid program address: %v", err)
		return result, nil
	}

	configPDA, _, err := FindRMNRemoteConfigPDA(programPK)
	if err != nil {
		result["error"] = fmt.Sprintf("failed to derive config PDA: %v", err)
		return result, nil
	}
	result["pda"] = configPDA.String()

	rawData, err := fetchPDAData(ctx, configPDA)
	if err != nil {
		result["pdaError"] = err.Error()
		return result, nil
	}

	if len(rawData) < 8 {
		result["decodeError"] = "data too short"
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	var version uint8
	decoder.Decode(&version)
	result["version"] = version

	owner, _ := decodePublicKey(decoder)
	result["owner"] = owner.String()

	proposedOwner, _ := decodePublicKey(decoder)
	result["proposedOwner"] = proposedOwner.String()

	return result, nil
}

// ViewTokenPool decodes a TokenPool state account
func ViewTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "TokenPool"

	// For token pools, we need the mint address to derive the PDA
	// For now, return basic info
	rawData, baseResult, err := getAccountData(ctx)
	if err != nil {
		return baseResult, nil
	}

	for k, v := range baseResult {
		result[k] = v
	}

	if rawData != nil {
		result["dataLength"] = len(rawData)
	}

	return result, nil
}

// ViewMCMConfig decodes an MCM MultisigConfig account
func ViewMCMConfig(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "MCM"

	rawData, baseResult, err := getAccountData(ctx)
	if err != nil {
		return baseResult, nil
	}

	for k, v := range baseResult {
		result[k] = v
	}

	if rawData == nil || len(rawData) < 8 {
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	var chainId uint64
	decoder.Decode(&chainId)
	result["chainId"] = chainId

	var multisigId [32]byte
	decoder.Decode(&multisigId)
	result["multisigId"] = string(bytes.TrimRight(multisigId[:], "\x00"))

	owner, _ := decodePublicKey(decoder)
	result["owner"] = owner.String()

	proposedOwner, _ := decodePublicKey(decoder)
	result["proposedOwner"] = proposedOwner.String()

	var groupQuorums [32]uint8
	decoder.Decode(&groupQuorums)
	result["groupQuorums"] = groupQuorums

	var groupParents [32]uint8
	decoder.Decode(&groupParents)
	result["groupParents"] = groupParents

	return result, nil
}

// ViewTimelock decodes a Timelock config account
func ViewTimelock(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "Timelock"

	rawData, baseResult, err := getAccountData(ctx)
	if err != nil {
		return baseResult, nil
	}

	for k, v := range baseResult {
		result[k] = v
	}

	if rawData == nil || len(rawData) < 8 {
		return result, nil
	}

	decoder := ag_binary.NewBorshDecoder(rawData[8:])

	owner, _ := decodePublicKey(decoder)
	result["owner"] = owner.String()

	proposedOwner, _ := decodePublicKey(decoder)
	result["proposedOwner"] = proposedOwner.String()

	proposerRoleAC, _ := decodePublicKey(decoder)
	result["proposerRoleAccessController"] = proposerRoleAC.String()

	executorRoleAC, _ := decodePublicKey(decoder)
	result["executorRoleAccessController"] = executorRoleAC.String()

	cancellerRoleAC, _ := decodePublicKey(decoder)
	result["cancellerRoleAccessController"] = cancellerRoleAC.String()

	bypasserRoleAC, _ := decodePublicKey(decoder)
	result["bypasserRoleAccessController"] = bypasserRoleAC.String()

	var minDelay uint64
	decoder.Decode(&minDelay)
	result["minDelay"] = minDelay

	return result, nil
}

// ViewSPLToken is a generic view for SPL token accounts
func ViewSPLToken(ctx *views.ViewContext) (map[string]any, error) {
	rawData, result, err := getAccountData(ctx)
	if err != nil || rawData == nil {
		return result, nil
	}

	result["type"] = "SPLToken"
	result["dataLength"] = len(rawData)

	// SPL Token Mint has a specific 82-byte layout
	if len(rawData) >= 82 {
		// Mint authority (36 bytes: 4 byte option + 32 byte pubkey)
		if rawData[0] == 1 { // Some
			mintAuth := rawData[4:36]
			result["mintAuthority"] = base58.Encode(mintAuth)
		}

		// Supply (8 bytes, little endian)
		supply := binary.LittleEndian.Uint64(rawData[36:44])
		result["supply"] = supply

		// Decimals (1 byte)
		result["decimals"] = rawData[44]

		// Is initialized (1 byte)
		result["isInitialized"] = rawData[45] == 1

		// Freeze authority (36 bytes)
		if rawData[46] == 1 { // Some
			freezeAuth := rawData[50:82]
			result["freezeAuthority"] = base58.Encode(freezeAuth)
		}
	}

	return result, nil
}

// ViewLinkToken is the same as SPLToken for Solana
func ViewLinkToken(ctx *views.ViewContext) (map[string]any, error) {
	result, err := ViewSPLToken(ctx)
	if result != nil {
		result["type"] = "LinkToken"
	}
	return result, err
}

// ViewReceiver provides a basic view for receiver accounts
func ViewReceiver(ctx *views.ViewContext) (map[string]any, error) {
	return ViewGenericAccount(ctx)
}

// ViewRemoteSource provides a basic view for remote source config
func ViewRemoteSource(ctx *views.ViewContext) (map[string]any, error) {
	return ViewGenericAccount(ctx)
}

// ViewRemoteDest provides a basic view for remote dest config
func ViewRemoteDest(ctx *views.ViewContext) (map[string]any, error) {
	return ViewGenericAccount(ctx)
}

// ViewAccessController provides a basic view for access controller
func ViewAccessController(ctx *views.ViewContext) (map[string]any, error) {
	return ViewGenericAccount(ctx)
}
