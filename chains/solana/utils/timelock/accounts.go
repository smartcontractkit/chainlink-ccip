package timelock

import (
	crypto_rand "crypto/rand"
	"math/big"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"
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

func TestRoleAccounts(numAccounts int) ([]RoleAccounts, RoleMap) {
	roles := []RoleAccounts{
		{
			Role:             timelock.Proposer_Role,
			Accounts:         mustCreateRoleAccounts(numAccounts),
			AccessController: mustPrivateKey(),
		},
		{
			Role:             timelock.Executor_Role,
			Accounts:         mustCreateRoleAccounts(numAccounts),
			AccessController: mustPrivateKey(),
		},
		{
			Role:             timelock.Canceller_Role,
			Accounts:         mustCreateRoleAccounts(numAccounts),
			AccessController: mustPrivateKey(),
		},
		{
			Role:             timelock.Bypasser_Role,
			Accounts:         mustCreateRoleAccounts(numAccounts),
			AccessController: mustPrivateKey(),
		},
	}

	roleMap := make(RoleMap)
	for _, role := range roles {
		roleMap[role.Role] = role
	}
	return roles, roleMap
}

func mustCreateRoleAccounts(num int) []solana.PrivateKey {
	if num < 1 || num > 64 {
		panic("num should be between 1 and 64")
	}
	accounts := make([]solana.PrivateKey, num)
	for i := 0; i < num; i++ {
		accounts[i] = mustPrivateKey()
	}
	return accounts
}

func mustPrivateKey() solana.PrivateKey {
	key, err := solana.NewRandomPrivateKey()
	if err != nil {
		panic(err)
	}
	return key
}
