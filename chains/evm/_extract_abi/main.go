package main

import (
	"fmt"
	"os"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/siloed_lock_release_token_pool"
)

func main() {
	os.MkdirAll("abi/v1_6_1", 0755)
	os.MkdirAll("bytecode/v1_6_1", 0755)

	os.WriteFile("abi/v1_6_1/lock_release_token_pool.json", []byte(lock_release_token_pool.LockReleaseTokenPoolABI), 0644)
	os.WriteFile("bytecode/v1_6_1/lock_release_token_pool.bin", []byte(lock_release_token_pool.LockReleaseTokenPoolBin), 0644)
	os.WriteFile("abi/v1_6_1/siloed_lock_release_token_pool.json", []byte(siloed_lock_release_token_pool.SiloedLockReleaseTokenPoolABI), 0644)
	os.WriteFile("bytecode/v1_6_1/siloed_lock_release_token_pool.bin", []byte(siloed_lock_release_token_pool.SiloedLockReleaseTokenPoolBin), 0644)

	fmt.Println("Extracted ABIs and bytecode for v1.6.1 LockReleaseTokenPool and SiloedLockReleaseTokenPool")
}
