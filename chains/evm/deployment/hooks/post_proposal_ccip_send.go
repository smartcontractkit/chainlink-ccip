package hooks

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/maps"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fq16 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	fq20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
	cciphooks "github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

var _ cciphooks.PostProposalCCIPSend = (*EVMPostProposalCCIPSend)(nil)

func init() {
	cciphooks.GetPostProposalCCIPSendRegistry().Register(chain_selectors.FamilyEVM, &EVMPostProposalCCIPSend{})
}

// EVMPostProposalCCIPSend provides EVM-specific discovery of CCIP lanes and fee tokens for post-proposal verify.
type EVMPostProposalCCIPSend struct{}

func (e *EVMPostProposalCCIPSend) SkipSend(env cldf.Environment) bool {
	isForked := strings.Contains(env.Name, "fork")
	// skip sending messages if the env is not forked
	// we only send messages in proposal forked env where there is no offchain and we have nodeIDs to execute the ccip-send through ccip plugin
	// skip if offchain is not nil or we have nodeIDs configured,
	// as that means we are in a non-forked environment where the send would actually be received
	return !isForked || env.Offchain != nil || len(env.NodeIDs) > 0
}

func (e *EVMPostProposalCCIPSend) PreSendValidation(env cldf.Environment, srcSel uint64) error {
	_, err := chain_selectors.GetSelectorFamily(srcSel)
	if err != nil {
		return fmt.Errorf("get selector family: %w", err)
	}
	_, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return fmt.Errorf("chain '%d' not found in chain selectors", srcSel)
	}
	return nil
}

func (e *EVMPostProposalCCIPSend) SupportedDestinations(env cldf.Environment, srcSel uint64) ([]uint64, error) {
	allDests, err := e.supportedRemoteChainsWithVersions(env, srcSel)
	if err != nil {
		return nil, err
	}
	return maps.Keys(allDests), nil
}

func (e *EVMPostProposalCCIPSend) AdapterVersionForLane(env cldf.Environment, srcSel, destSel uint64) (*semver.Version, error) {
	allDests, err := e.supportedRemoteChainsWithVersions(env, srcSel)
	if err != nil {
		return nil, err
	}
	for sel, version := range allDests {
		if sel == destSel {
			return version, nil
		}
	}
	return nil, fmt.Errorf("no adapter version found for src %d -> dest %d", srcSel, destSel)
}

func (e *EVMPostProposalCCIPSend) supportedRemoteChainsWithVersions(env cldf.Environment, srcSel uint64) (map[uint64]*semver.Version, error) {
	resolver := &adapters1_2.LaneVersionResolver{}
	alldests, _, err := resolver.DeriveLaneVersionsForChain(env, srcSel)
	if err != nil {
		return nil, fmt.Errorf("derive all lane versions from source %d: %w", srcSel, err)
	}
	return alldests, nil
}

func (e *EVMPostProposalCCIPSend) SupportedFeeTokens(env cldf.Environment, srcSel uint64) ([]string, error) {
	fqAddr, fqVer, err := sequences.GetFeeQuoterAddressAndVersionFromOnRamp(env.DataStore, srcSel, env.BlockChains)
	if err != nil {
		return nil, err
	}
	chain, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return nil, fmt.Errorf("chain %d not in environment EVM chains", srcSel)
	}

	var addrs []common.Address
	switch fqVer.Major() {
	case 1:
		fq, err := fq16.NewFeeQuoter(fqAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("fee quoter 1.x binding: %w", err)
		}
		addrs, err = fq.GetFeeTokens(nil)
		if err != nil {
			return nil, fmt.Errorf("getFeeTokens: %w", err)
		}
	case 2:
		fq, err := fq20.NewFeeQuoter(fqAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("fee quoter 2.x binding: %w", err)
		}
		addrs, err = fq.GetFeeTokens(nil)
		if err != nil {
			return nil, fmt.Errorf("getFeeTokens: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported fee quoter major version %d for chain %d", fqVer.Major(), srcSel)
	}

	out := make([]string, 0, len(addrs)+1)
	out = append(out, "") // native (empty encodes to wrapped native in adapter)
	for _, a := range addrs {
		out = append(out, a.Hex())
	}
	return out, nil
}
