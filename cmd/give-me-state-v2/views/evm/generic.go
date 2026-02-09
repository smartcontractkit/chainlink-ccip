package evm

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	_ "give-me-state-v2/views/evm/mcms"
	_ "give-me-state-v2/views/evm/v1_0"
	_ "give-me-state-v2/views/evm/v1_2"
	_ "give-me-state-v2/views/evm/v1_5"
	_ "give-me-state-v2/views/evm/v1_5_1"
	_ "give-me-state-v2/views/evm/v1_6"
	_ "give-me-state-v2/views/evm/v1_6_1"
	_ "give-me-state-v2/views/evm/v1_7"
)

func init() {
	views.Register("evm", "CommitteeVerifierResolver", "1.7.0", ViewGenericContract)
	views.Register("evm", "Executor", "1.7.0", ViewGenericContract)
	views.Register("evm", "ExecutorProxy", "1.7.0", ViewGenericContract)
	views.Register("evm", "MockReceiver", "1.7.0", ViewGenericContract)
	views.Register("evm", "BurnMintERC20WithDrip", "1.5.0", ViewGenericContract)
	views.Register("evm", "BurnMintTokenPool", "1.7.0", ViewGenericContract)
	views.Register("evm", "AdvancedPoolHooks", "1.7.0", ViewGenericContract)
}

func ViewGenericContract(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}
	return result, nil
}
