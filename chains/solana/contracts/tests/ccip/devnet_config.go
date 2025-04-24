package contracts

import (
	_ "embed"
	"testing"

	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/test-go/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

//go:embed devnet.config.yaml
var devnetInfoBuffer []byte

type DevnetInfo struct {
	Offramp                  string `yaml:"offramp"`
	RedirectingReceiver      string `yaml:"redirecting_receiver"`
	TokenMintForRedirectTest string `yaml:"token_mint_for_redirect_test"`
	FinalReceiverForRedirect string `yaml:"final_receiver_for_redirect"`
	PrivateKeys              struct {
		Admin []byte `yaml:"admin"`
	} `yaml:"private_keys"`
	ChainSelectors struct {
		Sepolia uint64 `yaml:"sepolia"`
		Fuji    uint64 `yaml:"fuji"`
		Svm     uint64 `yaml:"svm"`
	} `yaml:"chain_selectors"`
	RPC string `yaml:"rpc"`
}

func getDevnetInfo() (DevnetInfo, error) {
	var devnetInfo DevnetInfo
	if err := yaml.Unmarshal(devnetInfoBuffer, &devnetInfo); err != nil {
		return DevnetInfo{}, err
	}
	fmt.Printf("Devnet info: %+v\n", devnetInfo)
	return devnetInfo, nil
}

type offrampPDAs struct {
	config             solana.PublicKey
	referenceAddresses solana.PublicKey
	billingSigner      solana.PublicKey
	state              solana.PublicKey
}

func getOfframpPDAs(offrampAddress solana.PublicKey) (offrampPDAs, error) {
	offrampReferenceAddressesPDA, _, err := state.FindOfframpReferenceAddressesPDA(offrampAddress)
	if err != nil {
		return offrampPDAs{}, err
	}

	offrampConfigPDA, _, err := state.FindOfframpConfigPDA(offrampAddress)
	if err != nil {
		return offrampPDAs{}, err
	}

	offrampBillingSignerPDA, _, err := state.FindOfframpBillingSignerPDA(offrampAddress)
	if err != nil {
		return offrampPDAs{}, err
	}

	offrampStatePDA, _, err := state.FindOfframpStatePDA(offrampAddress)
	if err != nil {
		return offrampPDAs{}, err
	}

	return offrampPDAs{
		config:             offrampConfigPDA,
		referenceAddresses: offrampReferenceAddressesPDA,
		billingSigner:      offrampBillingSignerPDA,
		state:              offrampStatePDA,
	}, nil
}

type senderPDAs struct {
	program solana.PublicKey
	state   solana.PublicKey
}

func (s *senderPDAs) FindChainConfig(t *testing.T, chainSelector uint64) solana.PublicKey {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	pda, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_chain_config"), chainSelectorLE}, s.program)
	require.NoError(t, err)
	return pda
}

func getSenderPDAs(senderProgram solana.PublicKey) (senderPDAs, error) {
	senderStatePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, senderProgram)
	if err != nil {
		return senderPDAs{}, err
	}

	return senderPDAs{
		program: senderProgram,
		state:   senderStatePDA,
	}, nil
}
