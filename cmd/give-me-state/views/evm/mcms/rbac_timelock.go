package mcms

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
	"sync"
)

// Function selectors for RBACTimelock
var (
	// getRoleMemberCount(bytes32) returns (uint256)
	selectorGetRoleMemberCount = common.HexToSelector("ca15c873")
	// getRoleMember(bytes32, uint256) returns (address)
	selectorGetRoleMember = common.HexToSelector("9010d07c")
)

// Role hashes (bytes32)
var (
	roleAdminRole, _          = hex.DecodeString("a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775")
	roleBypasserRole, _       = hex.DecodeString("a1b2b8005de234c4b8ce8cd0be058239056e0d54f6097825b5117101469d5a8d")
	roleCancellerRole, _      = hex.DecodeString("fd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783")
	roleDefaultAdminRole, _   = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	roleExecutorRole, _       = hex.DecodeString("d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63")
	roleProposerRole, _       = hex.DecodeString("b09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1")
)

// roleDefinition maps role name to its bytes32 hash
var roleDefinitions = map[string][]byte{
	"ADMIN_ROLE":         roleAdminRole,
	"BYPASSER_ROLE":      roleBypasserRole,
	"CANCELLER_ROLE":     roleCancellerRole,
	"DEFAULT_ADMIN_ROLE": roleDefaultAdminRole,
	"EXECUTOR_ROLE":      roleExecutorRole,
	"PROPOSER_ROLE":      roleProposerRole,
}

// ViewRBACTimelock generates a view of an RBACTimelock contract.
func ViewRBACTimelock(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Get members for each role concurrently
	membersByRole := getMembersByRole(ctx)
	result["membersByRole"] = membersByRole

	// The "owner" is typically whoever has DEFAULT_ADMIN_ROLE
	if admins, ok := membersByRole["DEFAULT_ADMIN_ROLE"].([]string); ok && len(admins) > 0 {
		result["owner"] = admins[0]
	} else {
		result["owner"] = "0x0000000000000000000000000000000000000000"
	}

	return result, nil
}

// getMembersByRole fetches all role members concurrently.
func getMembersByRole(ctx *views.ViewContext) map[string]any {
	result := make(map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for roleName, roleHash := range roleDefinitions {
		wg.Add(1)
		go func(name string, hash []byte) {
			defer wg.Done()

			members, err := getRoleMembers(ctx, hash)
			if err == nil {
				mu.Lock()
				result[name] = members
				mu.Unlock()
			}
		}(roleName, roleHash)
	}

	wg.Wait()
	return result
}

// getRoleMembers fetches all members for a specific role.
func getRoleMembers(ctx *views.ViewContext, roleHash []byte) ([]string, error) {
	// Get member count
	count, err := getRoleMemberCount(ctx, roleHash)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return []string{}, nil
	}

	// Fetch all members concurrently
	members := make([]string, count)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var fetchErr error

	for i := uint64(0); i < count; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()

			addr, err := getRoleMember(ctx, roleHash, index)
			if err != nil {
				mu.Lock()
				if fetchErr == nil {
					fetchErr = err
				}
				mu.Unlock()
				return
			}

			mu.Lock()
			members[index] = addr
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if fetchErr != nil {
		return nil, fetchErr
	}

	return members, nil
}

// getRoleMemberCount calls getRoleMemberCount(bytes32) on the contract.
func getRoleMemberCount(ctx *views.ViewContext, roleHash []byte) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRoleMemberCount, roleHash)
	if err != nil {
		return 0, err
	}

	if len(data) < 32 {
		return 0, nil
	}

	return common.DecodeUint64FromBytes(data[0:32]), nil
}

// getRoleMember calls getRoleMember(bytes32, uint256) on the contract.
func getRoleMember(ctx *views.ViewContext, roleHash []byte, index uint64) (string, error) {
	indexBytes := common.EncodeUint64(index)
	data, err := common.ExecuteCall(ctx, selectorGetRoleMember, roleHash, indexBytes)
	if err != nil {
		return "", err
	}

	return common.DecodeAddress(data)
}
