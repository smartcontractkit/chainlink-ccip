package solana

// SolChainView is the json-persistable view of one Solana chain's CCIP deployment.
//
// It lives here rather than in chainlink/deployment/ccip/view, where the EVM ChainView it sits
// beside drags in the v1_0..v1_6 gethwrappers plus the Aptos, Sui and TON view packages. Nothing
// in it is EVM-specific, so it moves with the Solana views.
type SolChainView struct {
	ChainSelector uint64 `json:"chainSelector,omitempty"`
	ChainID       string `json:"chainID,omitempty"`
	// v1.6
	FeeQuoter        map[string]FeeQuoterView `json:"feeQuoter,omitempty"`
	Router           map[string]RouterView    `json:"router,omitempty"`
	OffRamp          map[string]OffRampView   `json:"offRamp,omitempty"`
	RMNRemote        map[string]RMNRemoteView `json:"rmnRemote,omitempty"`
	TokenPool        map[string]TokenPoolView `json:"tokenPool,omitempty"`
	LinkToken        TokenView                `json:"linkToken"`
	Tokens           map[string]TokenView     `json:"tokens,omitempty"`
	MCMSWithTimelock MCMSWithTimelockView     `json:"mcmsWithTimelock"`
}

// NewSolChain returns a SolChainView with every map allocated, ready to be filled in.
func NewSolChain() SolChainView {
	return SolChainView{
		FeeQuoter:        make(map[string]FeeQuoterView),
		Router:           make(map[string]RouterView),
		OffRamp:          make(map[string]OffRampView),
		RMNRemote:        make(map[string]RMNRemoteView),
		TokenPool:        make(map[string]TokenPoolView),
		Tokens:           make(map[string]TokenView),
		MCMSWithTimelock: MCMSWithTimelockView{},
	}
}
