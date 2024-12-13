package mcms

import (
	crypto_rand "crypto/rand"
	"math/big"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/generated/timelock"
)

// test helper with normal accounts
type RoleMap map[timelock.Role]RoleAccounts

type RoleAccounts struct {
	Role             timelock.Role
	Accounts         []solana.PrivateKey
	AccessController solana.PrivateKey
}

func (r RoleAccounts) RandomPick() solana.PrivateKey {
	if len(r.Accounts) == 0 {
		panic("no accounts to pick from")
	}

	maxN := big.NewInt(int64(len(r.Accounts)))
	n, err := crypto_rand.Int(crypto_rand.Reader, maxN)
	if err != nil {
		panic(err)
	}

	return r.Accounts[n.Int64()]
}

func TestRoleAccounts(t *testing.T, numAccounts int) ([]RoleAccounts, RoleMap) {
	roles := []RoleAccounts{
		{
			Role:             timelock.Proposer_Role,
			Accounts:         createRoleAccounts(t, numAccounts),
			AccessController: getPrivateKey(t),
		},
		{
			Role:             timelock.Executor_Role,
			Accounts:         createRoleAccounts(t, numAccounts),
			AccessController: getPrivateKey(t),
		},
		{
			Role:             timelock.Canceller_Role,
			Accounts:         createRoleAccounts(t, numAccounts),
			AccessController: getPrivateKey(t),
		},
		{
			Role:             timelock.Bypasser_Role,
			Accounts:         createRoleAccounts(t, numAccounts),
			AccessController: getPrivateKey(t),
		},
	}

	roleMap := make(RoleMap)
	for _, role := range roles {
		roleMap[role.Role] = role
	}
	return roles, roleMap
}

func createRoleAccounts(t *testing.T, num int) []solana.PrivateKey {
	if num < 1 || num > 64 {
		panic("num should be between 1 and 64")
	}
	accounts := make([]solana.PrivateKey, num)
	for i := 0; i < num; i++ {
		account, err := solana.NewRandomPrivateKey()
		require.NoError(t, err)
		accounts[i] = account
	}
	return accounts
}

func getPrivateKey(t *testing.T) solana.PrivateKey {
	key, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	return key
}
