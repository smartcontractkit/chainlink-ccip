package tokens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"

	lrtp161bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/lock_release_token_pool"
	siloed161bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/siloed_lock_release_token_pool"
	aphbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	cctbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cross_chain_token"
	lockboxbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	lrtp170bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lock_release_token_pool"
	siloed170bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_lock_release_token_pool"
	tpbinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
)

func appendWrite(ops []evm_contract.WriteOutput, w ops2contract.WriteOutput) []evm_contract.WriteOutput {
	return append(ops, writeOutputOps2ToLegacy(w))
}

func bindCrossChainToken(addr common.Address, chain evm.Chain) (cctbind.CrossChainTokenInterface, error) {
	c, err := cctbind.NewCrossChainToken(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CrossChainToken at %s: %w", addr.Hex(), err)
	}
	return c, nil
}

func bindLRTP170(addr common.Address, chain evm.Chain) (lrtp170bind.LockReleaseTokenPoolInterface, error) {
	p, err := lrtp170bind.NewLockReleaseTokenPool(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind LockReleaseTokenPool v2 at %s: %w", addr.Hex(), err)
	}
	return p, nil
}

func bindLRTP161(addr common.Address, chain evm.Chain) (lrtp161bind.LockReleaseTokenPoolInterface, error) {
	p, err := lrtp161bind.NewLockReleaseTokenPool(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind LockReleaseTokenPool v1.6.1 at %s: %w", addr.Hex(), err)
	}
	return p, nil
}

func bindSiloedLRTP161(addr common.Address, chain evm.Chain) (siloed161bind.SiloedLockReleaseTokenPoolInterface, error) {
	p, err := siloed161bind.NewSiloedLockReleaseTokenPool(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind SiloedLockReleaseTokenPool v1.6.1 at %s: %w", addr.Hex(), err)
	}
	return p, nil
}

func bindSiloedLRTP170(addr common.Address, chain evm.Chain) (siloed170bind.SiloedLockReleaseTokenPoolInterface, error) {
	p, err := siloed170bind.NewSiloedLockReleaseTokenPool(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind SiloedLockReleaseTokenPool v2 at %s: %w", addr.Hex(), err)
	}
	return p, nil
}

func bindLockBox(addr common.Address, chain evm.Chain) (lockboxbind.ERC20LockBoxInterface, error) {
	lb, err := lockboxbind.NewERC20LockBox(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind ERC20LockBox at %s: %w", addr.Hex(), err)
	}
	return lb, nil
}

func writeOutputOps2ToLegacy(w ops2contract.WriteOutput) evm_contract.WriteOutput {
	return WriteOutputOps2ToLegacy(w)
}

// WriteOutputOps2ToLegacy converts ops2 write output to the legacy contract batch format.
func WriteOutputOps2ToLegacy(w ops2contract.WriteOutput) evm_contract.WriteOutput {
	var ei *evm_contract.ExecInfo
	if w.ExecInfo != nil {
		ei = &evm_contract.ExecInfo{Hash: w.ExecInfo.Hash}
	}
	return evm_contract.WriteOutput{
		ChainSelector: w.ChainSelector,
		Tx:            w.Tx,
		ExecInfo:      ei,
	}
}

func bindTokenPool(addr common.Address, chain evm.Chain) (tpbinding.TokenPoolInterface, error) {
	return BindTokenPool(addr, chain)
}

// BindTokenPool binds a v2.0 TokenPool at the given address.
func BindTokenPool(addr common.Address, chain evm.Chain) (tpbinding.TokenPoolInterface, error) {
	tp, err := tpbinding.NewTokenPool(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind TokenPool at %s: %w", addr.Hex(), err)
	}
	return tp, nil
}

func bindAdvancedPoolHooks(addr common.Address, chain evm.Chain) (aphbind.AdvancedPoolHooksInterface, error) {
	return BindAdvancedPoolHooks(addr, chain)
}

// BindAdvancedPoolHooks binds AdvancedPoolHooks at the given address.
func BindAdvancedPoolHooks(addr common.Address, chain evm.Chain) (aphbind.AdvancedPoolHooksInterface, error) {
	aph, err := aphbind.NewAdvancedPoolHooks(addr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind AdvancedPoolHooks at %s: %w", addr.Hex(), err)
	}
	return aph, nil
}
