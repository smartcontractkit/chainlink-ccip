package tokenpools

import (
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ManualRegistrationInput struct {
	ChainSelector        uint64              `yaml:"chain-selector" json:"chainSelector"`
	MCMS                 mcms.Input          `yaml:"mcms,omitempty" json:"mcms,omitempty"`
	RegisterTokenConfigs RegisterTokenConfig `yaml:"register-token-configs" json:"registerTokenConfigs"`
	SVMExtraArgs         *SVMExtraArgs       `yaml:"svm-extra-args,omitempty" json:"svmExtraArgs,omitempty"`
}

type RegisterTokenConfig struct {
	TokenAddress  string `yaml:"token-address" json:"tokenAddress"`
	ProposedOwner string `yaml:"proposed-owner" json:"proposedOwner"`
	Metadata      string `yaml:"metadata" json:"metadata"`
	PoolType      string `yaml:"pool-type" json:"poolType"`
}

type SVMExtraArgs struct {
	CustomerMintAuthorities []solana.PublicKey `yaml:"customer-mint-authorities,omitempty" json:"customerMintAuthorities,omitempty"`
	SkipTokenPoolInit       bool               `yaml:"skip-token-pool-init" json:"skipTokenPoolInit"`
}
