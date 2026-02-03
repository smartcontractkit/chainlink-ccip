package evm

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"

	// Import version packages to trigger their init() functions
	_ "call-orchestrator-demo/views/evm/mcms"
	_ "call-orchestrator-demo/views/evm/v1_0"
	_ "call-orchestrator-demo/views/evm/v1_2"
	_ "call-orchestrator-demo/views/evm/v1_5"
	_ "call-orchestrator-demo/views/evm/v1_5_1"
	_ "call-orchestrator-demo/views/evm/v1_6"
	_ "call-orchestrator-demo/views/evm/v1_6_1"
	_ "call-orchestrator-demo/views/evm/v1_7"

	// Import Solana views
	_ "call-orchestrator-demo/views/solana"
)

func init() {
	// Register generic views for common contract types that don't have specific implementations
	// These provide basic owner/typeAndVersion for unsupported specific versions

	// ARMProxy - 1.0.0 is in v1_0 folder
	// RMNRemote - 1.6.0 is in v1_6 folder
	// RegistryModuleOwnerCustom - 1.5.0 and 1.6.0 are in their respective folders (no owner() method)
	// BurnMintTokenPool - 1.6.1 is in v1_6_1 folder
	// LockReleaseTokenPool - 1.6.1 is in v1_6_1 folder
	// TokenPoolFactory - 1.5.1 is in v1_5_1 folder

	// CommitteeVerifierResolver
	views.Register("evm", "CommitteeVerifierResolver", "1.7.0", ViewGenericContract)

	// Executor
	views.Register("evm", "Executor", "1.7.0", ViewGenericContract)

	// ExecutorProxy
	views.Register("evm", "ExecutorProxy", "1.7.0", ViewGenericContract)

	// MockReceiver
	views.Register("evm", "MockReceiver", "1.7.0", ViewGenericContract)

	// Token-related contracts (generic for now)
	views.Register("evm", "BurnMintERC20WithDrip", "1.5.0", ViewGenericContract)
	views.Register("evm", "BurnMintTokenPool", "1.7.0", ViewGenericContract)
	views.Register("evm", "AdvancedPoolHooks", "1.7.0", ViewGenericContract)
}

// ViewGenericContract provides a basic view for any EVM contract.
// Fetches owner and typeAndVersion only.
func ViewGenericContract(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	// Basic info
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Get owner
	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	// Get typeAndVersion
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	return result, nil
}
