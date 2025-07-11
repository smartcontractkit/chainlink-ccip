package solana

import "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

// These types are needed as special call args for the CCIP contract transmitter on
// Solana. These call args are passed to the ChainWriter and are used to derive additional
// fields to send txs on chain.
type ExtraDataDecoded struct {
	// ExtraArgsDecoded contain message specific extra args.
	ExtraArgsDecoded map[string]any
	// DestExecDataDecoded contain token transfer specific extra args.
	DestExecDataDecoded []map[string]any
}

type SVMExecCallArgs struct {
	ReportContext [2][32]byte                `mapstructure:"ReportContext"`
	Report        []byte                     `mapstructure:"Report"`
	Info          ccipocr3.ExecuteReportInfo `mapstructure:"Info"`
	ExtraData     ExtraDataDecoded           `mapstructure:"ExtraData"`
	TokenIndexes  []byte                     `mapstructure:"TokenIndexes"`
}

type SVMCommitCallArgs struct {
	ReportContext [2][32]byte               `mapstructure:"ReportContext"`
	Report        []byte                    `mapstructure:"Report"`
	Rs            [][32]byte                `mapstructure:"Rs"`
	Ss            [][32]byte                `mapstructure:"Ss"`
	RawVs         [32]byte                  `mapstructure:"RawVs"`
	Info          ccipocr3.CommitReportInfo `mapstructure:"Info"`
}
