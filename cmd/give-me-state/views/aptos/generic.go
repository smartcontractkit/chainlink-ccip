package aptos

import (
	"call-orchestrator-demo/views"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// ===== Resource Fetching =====

// AptosResource represents a single Aptos resource from the REST API
type AptosResource struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// getAccountResources fetches all resources for an Aptos account
func getAccountResources(ctx *views.ViewContext, address string) ([]AptosResource, error) {
	// Normalize address format
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  []byte(address),
		Data:    nil, // nil = fetch all resources
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, result.Error
	}

	var resources []AptosResource
	if err := json.Unmarshal(result.Data, &resources); err != nil {
		return nil, fmt.Errorf("failed to parse resources: %w", err)
	}

	return resources, nil
}

// getAccountResource fetches a specific resource for an Aptos account
func getAccountResource(ctx *views.ViewContext, address, resourceType string) (map[string]any, error) {
	// Normalize address format
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  []byte(address),
		Data:    []byte(resourceType),
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, result.Error
	}

	var resource AptosResource
	if err := json.Unmarshal(result.Data, &resource); err != nil {
		return nil, fmt.Errorf("failed to parse resource: %w", err)
	}

	var data map[string]any
	if err := json.Unmarshal(resource.Data, &data); err != nil {
		return nil, fmt.Errorf("failed to parse resource data: %w", err)
	}

	return data, nil
}

// findResourceByTypeSuffix finds a resource matching a type suffix (e.g., "::fee_quoter::Config")
func findResourceByTypeSuffix(resources []AptosResource, suffix string) (map[string]any, bool) {
	for _, res := range resources {
		if strings.HasSuffix(res.Type, suffix) {
			var data map[string]any
			if err := json.Unmarshal(res.Data, &data); err == nil {
				return data, true
			}
		}
	}
	return nil, false
}

// findResourceByTypePrefix finds a resource matching a type prefix
func findResourceByTypePrefix(resources []AptosResource, prefix string) (map[string]any, string, bool) {
	for _, res := range resources {
		if strings.Contains(res.Type, prefix) {
			var data map[string]any
			if err := json.Unmarshal(res.Data, &data); err == nil {
				return data, res.Type, true
			}
		}
	}
	return nil, "", false
}

// ===== View Functions =====

// ViewGenericAccount provides a basic view for any Aptos account
func ViewGenericAccount(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosGeneric"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	result["resourceCount"] = len(resources)

	// List resource types
	types := make([]string, len(resources))
	for i, res := range resources {
		types[i] = res.Type
	}
	result["resourceTypes"] = types

	return result, nil
}

// ViewCCIP decodes the main CCIP module state
// Note: The CCIP address is primarily a code deployment. The actual state resources
// (fee_quoter, rmn_remote, etc.) are stored at module-specific addresses.
func ViewCCIP(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosCCIP"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	result["resourceCount"] = len(resources)

	// List all resource types for debugging/discovery
	resourceTypes := make([]string, 0, len(resources))
	for _, res := range resources {
		resourceTypes = append(resourceTypes, res.Type)
	}
	result["resourceTypes"] = resourceTypes

	// Look for ObjectCore (common Aptos pattern)
	if objectCore, ok := findResourceByTypeSuffix(resources, "::object::ObjectCore"); ok {
		if owner, ok := objectCore["owner"].(string); ok {
			result["owner"] = owner
		}
	}

	// Look for PackageRegistry (code deployment info)
	if packageRegistry, ok := findResourceByTypeSuffix(resources, "::code::PackageRegistry"); ok {
		if packages, ok := packageRegistry["packages"].([]any); ok && len(packages) > 0 {
			result["packageCount"] = len(packages)
			// Extract package names if available
			packageNames := make([]string, 0)
			for _, pkg := range packages {
				if pkgMap, ok := pkg.(map[string]any); ok {
					if name, ok := pkgMap["name"].(string); ok {
						packageNames = append(packageNames, name)
					}
				}
			}
			if len(packageNames) > 0 {
				result["packages"] = packageNames
			}
		}
	}

	return result, nil
}

// ViewRouter decodes an Aptos Router state
func ViewRouter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosRouter"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	// Look for router config
	if routerConfig, ok := findResourceByTypeSuffix(resources, "::router::Config"); ok {
		if owner, ok := routerConfig["owner"].(string); ok {
			result["owner"] = owner
		}
		if destChains, ok := routerConfig["dest_chains"].(map[string]any); ok {
			result["destChains"] = destChains
		}
	}

	// Check for OnRamps
	if onRamps, ok := findResourceByTypeSuffix(resources, "::router::OnRamps"); ok {
		result["onRamps"] = onRamps
	}

	// Fetch destination chain configs in parallel
	if len(ctx.AllChainSelectors) > 0 {
		destConfigs := fetchAptosRouterDestConfigs(ctx, resources)
		if len(destConfigs) > 0 {
			result["destinationChainConfigs"] = destConfigs
		}
	}

	return result, nil
}

// fetchAptosRouterDestConfigs extracts destination chain configs from router resources
func fetchAptosRouterDestConfigs(ctx *views.ViewContext, resources []AptosResource) map[uint64]map[string]any {
	configs := make(map[uint64]map[string]any)

	// Look for dest chain entries in resources
	for _, res := range resources {
		if strings.Contains(res.Type, "::router::DestChain") {
			var data map[string]any
			if err := json.Unmarshal(res.Data, &data); err == nil {
				if selector, ok := data["chain_selector"].(float64); ok {
					configs[uint64(selector)] = data
				}
			}
		}
	}

	return configs
}

// ViewOnRamp decodes an Aptos OnRamp state
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosOnRamp"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	// Look for onramp config
	if onRampConfig, ok := findResourceByTypeSuffix(resources, "::onramp::Config"); ok {
		if owner, ok := onRampConfig["owner"].(string); ok {
			result["owner"] = owner
		}
		if staticConfig, ok := onRampConfig["static_config"].(map[string]any); ok {
			result["staticConfig"] = staticConfig
		}
		if dynamicConfig, ok := onRampConfig["dynamic_config"].(map[string]any); ok {
			result["dynamicConfig"] = dynamicConfig
		}
	}

	// Look for dest chain configs
	destChainConfigs := make(map[uint64]map[string]any)
	for _, res := range resources {
		if strings.Contains(res.Type, "::onramp::DestChain") {
			var data map[string]any
			if err := json.Unmarshal(res.Data, &data); err == nil {
				if selector, ok := data["chain_selector"].(float64); ok {
					destChainConfigs[uint64(selector)] = data
				}
			}
		}
	}
	if len(destChainConfigs) > 0 {
		result["destChainConfigs"] = destChainConfigs
	}

	return result, nil
}

// ViewOffRamp decodes an Aptos OffRamp state
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosOffRamp"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	// Look for offramp config
	if offRampConfig, ok := findResourceByTypeSuffix(resources, "::offramp::Config"); ok {
		if owner, ok := offRampConfig["owner"].(string); ok {
			result["owner"] = owner
		}
		if staticConfig, ok := offRampConfig["static_config"].(map[string]any); ok {
			result["staticConfig"] = staticConfig
		}
		if dynamicConfig, ok := offRampConfig["dynamic_config"].(map[string]any); ok {
			result["dynamicConfig"] = dynamicConfig
		}
	}

	// Look for source chain configs
	sourceChainConfigs := make(map[uint64]map[string]any)
	for _, res := range resources {
		if strings.Contains(res.Type, "::offramp::SourceChain") {
			var data map[string]any
			if err := json.Unmarshal(res.Data, &data); err == nil {
				if selector, ok := data["chain_selector"].(float64); ok {
					sourceChainConfigs[uint64(selector)] = data
				}
			}
		}
	}
	if len(sourceChainConfigs) > 0 {
		result["sourceChainConfigs"] = sourceChainConfigs
	}

	return result, nil
}

// ViewTokenPool decodes an Aptos TokenPool state
func ViewTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosTokenPool"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	result["resourceCount"] = len(resources)

	// List all resource types for debugging
	resourceTypes := make([]string, 0, len(resources))
	for _, res := range resources {
		resourceTypes = append(resourceTypes, res.Type)
	}
	result["resourceTypes"] = resourceTypes

	// Look for TokenPoolRegistration
	if registration, ok := findResourceByTypeSuffix(resources, "::token_admin_registry::TokenPoolRegistration"); ok {
		result["tokenPoolRegistration"] = registration
	}

	// Look for ObjectCore (to get owner)
	if objectCore, ok := findResourceByTypeSuffix(resources, "::object::ObjectCore"); ok {
		if owner, ok := objectCore["owner"].(string); ok {
			result["owner"] = owner
		}
	}

	// Look for PackageRegistry (code info)
	if packageRegistry, ok := findResourceByTypeSuffix(resources, "::code::PackageRegistry"); ok {
		if packages, ok := packageRegistry["packages"].([]any); ok && len(packages) > 0 {
			packageNames := make([]string, 0)
			for _, pkg := range packages {
				if pkgMap, ok := pkg.(map[string]any); ok {
					if name, ok := pkgMap["name"].(string); ok {
						packageNames = append(packageNames, name)
					}
				}
			}
			if len(packageNames) > 0 {
				result["packages"] = packageNames
			}
		}
	}

	return result, nil
}

// ViewMCMS decodes an Aptos MCMS (Multi-Chain Multi-Sig) state
func ViewMCMS(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosManyChainMultisig"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	result["resourceCount"] = len(resources)

	// Look for MCMS MultisigState
	if multisigState, ok := findResourceByTypeSuffix(resources, "::mcms::MultisigState"); ok {
		result["multisigState"] = multisigState
	}

	// Look for MCMS Timelock
	if timelock, ok := findResourceByTypeSuffix(resources, "::mcms::Timelock"); ok {
		result["timelock"] = timelock
		// Extract min_delay if available
		if minDelay, ok := timelock["min_delay"]; ok {
			result["timelockMinDelay"] = minDelay
		}
	}

	// Look for MCMS AccountState
	if accountState, ok := findResourceByTypeSuffix(resources, "::mcms_account::AccountState"); ok {
		result["accountState"] = accountState
		// Extract owner if available
		if owner, ok := accountState["owner"].(string); ok {
			result["owner"] = owner
		}
	}

	// Look for MCMS RegistryState
	if registryState, ok := findResourceByTypeSuffix(resources, "::mcms_registry::RegistryState"); ok {
		result["registryState"] = registryState
	}

	return result, nil
}

// ViewToken decodes an Aptos Token (Managed Token / Fungible Asset) state
func ViewToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosToken"

	resources, err := getAccountResources(ctx, ctx.AddressHex)
	if err != nil {
		result["error"] = err.Error()
		return result, nil
	}

	result["resourceCount"] = len(resources)

	// Look for fungible asset metadata (standard Aptos FA)
	if metadata, ok := findResourceByTypeSuffix(resources, "::fungible_asset::Metadata"); ok {
		if name, ok := metadata["name"].(string); ok {
			result["name"] = name
		}
		if symbol, ok := metadata["symbol"].(string); ok {
			result["symbol"] = symbol
		}
		if decimals, ok := metadata["decimals"].(float64); ok {
			result["decimals"] = int(decimals)
		}
		if icon, ok := metadata["icon_uri"].(string); ok && icon != "" {
			result["iconUri"] = icon
		}
		if project, ok := metadata["project_uri"].(string); ok && project != "" {
			result["projectUri"] = project
		}
	}

	// Look for concurrent supply
	if supply, ok := findResourceByTypeSuffix(resources, "::fungible_asset::ConcurrentSupply"); ok {
		if current, ok := supply["current"].(map[string]any); ok {
			if value, ok := current["value"].(string); ok {
				result["supply"] = value
			}
		}
	}

	// Look for ObjectCore (to get owner)
	if objectCore, ok := findResourceByTypeSuffix(resources, "::object::ObjectCore"); ok {
		if owner, ok := objectCore["owner"].(string); ok {
			result["owner"] = owner
		}
	}

	// Look for managed_token metadata refs
	if tokenRefs, ok := findResourceByTypeSuffix(resources, "::managed_token::TokenMetadataRefs"); ok {
		result["managedTokenRefs"] = tokenRefs
	}

	return result, nil
}

// ===== Helper for parallel fetching =====

// fetchAptosDestChainConfigs fetches destination chain configs in parallel
func fetchAptosDestChainConfigs(ctx *views.ViewContext, baseAddress string, configType string) map[uint64]map[string]any {
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

			// Try to fetch the resource for this chain
			// Resource type would be something like {address}::{module}::DestChain<{selector}>
			// This is simplified - actual implementation would need proper resource type construction
			resourceType := fmt.Sprintf("%s::%s::DestChain", baseAddress, configType)

			data, err := getAccountResource(ctx, baseAddress, resourceType)
			if err != nil {
				return
			}

			mu.Lock()
			configs[remote] = data
			mu.Unlock()
		}(chainSelector)
	}

	wg.Wait()
	return configs
}
