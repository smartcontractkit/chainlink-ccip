package mcms

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"give-me-state-v2/views"

	gethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const rbacTimelockABIJson = `[{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"}],"name":"getRoleMemberCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getRoleMember","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`

var rbacTimelockABI abi.ABI

func init() {
	var err error
	rbacTimelockABI, err = abi.JSON(strings.NewReader(rbacTimelockABIJson))
	if err != nil {
		panic("Failed to parse RBACTimelock ABI: " + err.Error())
	}
}

// Role hashes (bytes32)
var (
	roleAdminRole, _        = hex.DecodeString("a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775")
	roleBypasserRole, _     = hex.DecodeString("a1b2b8005de234c4b8ce8cd0be058239056e0d54f6097825b5117101469d5a8d")
	roleCancellerRole, _    = hex.DecodeString("fd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783")
	roleDefaultAdminRole, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	roleExecutorRole, _     = hex.DecodeString("d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63")
	roleProposerRole, _     = hex.DecodeString("b09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1")
)

// roleDefinitions maps role name to its bytes32 hash
var roleDefinitions = map[string][]byte{
	"ADMIN_ROLE":         roleAdminRole,
	"BYPASSER_ROLE":      roleBypasserRole,
	"CANCELLER_ROLE":     roleCancellerRole,
	"DEFAULT_ADMIN_ROLE": roleDefaultAdminRole,
	"EXECUTOR_ROLE":      roleExecutorRole,
	"PROPOSER_ROLE":      roleProposerRole,
}

// executeRBACTimelockCall packs a call, executes it, and returns raw response bytes.
func executeRBACTimelockCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := rbacTimelockABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}

	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("%s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// ViewRBACTimelock generates a view of an RBACTimelock contract.
// Uses bespoke ABI JSON for proper decoding (no Go bindings available).
func ViewRBACTimelock(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	membersByRole := getMembersByRole(ctx)
	result["membersByRole"] = membersByRole

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

// getRoleMembers fetches all members for a specific role using ABI bindings.
func getRoleMembers(ctx *views.ViewContext, roleHash []byte) ([]string, error) {
	var role [32]byte
	copy(role[:], roleHash)

	count, err := getRoleMemberCount(ctx, role)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return []string{}, nil
	}

	members := make([]string, count)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var fetchErr error

	for i := uint64(0); i < count; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()

			addr, err := getRoleMember(ctx, role, index)
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

// getRoleMemberCount calls getRoleMemberCount(bytes32) using ABI bindings.
func getRoleMemberCount(ctx *views.ViewContext, role [32]byte) (uint64, error) {
	data, err := executeRBACTimelockCall(ctx, "getRoleMemberCount", role)
	if err != nil {
		return 0, err
	}

	results, err := rbacTimelockABI.Unpack("getRoleMemberCount", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getRoleMemberCount: %w", err)
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("no results from getRoleMemberCount call")
	}
	val, ok := results[0].(*big.Int)
	if !ok {
		return 0, fmt.Errorf("unexpected type for role member count: %T", results[0])
	}
	return val.Uint64(), nil
}

// getRoleMember calls getRoleMember(bytes32, uint256) using ABI bindings.
func getRoleMember(ctx *views.ViewContext, role [32]byte, index uint64) (string, error) {
	data, err := executeRBACTimelockCall(ctx, "getRoleMember", role, new(big.Int).SetUint64(index))
	if err != nil {
		return "", err
	}

	results, err := rbacTimelockABI.Unpack("getRoleMember", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getRoleMember: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from getRoleMember call")
	}
	addr, ok := results[0].(gethCommon.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for role member: %T", results[0])
	}
	return addr.Hex(), nil
}
