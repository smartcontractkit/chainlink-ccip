//go:build devnet
// +build devnet

package contracts

import (
	_ "embed"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/test-go/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

//go:embed devnet.config.yaml
var devnetInfoBuffer []byte

type DevnetInfo struct {
	Offramp       string `yaml:"offramp"`
	TestReceiver  string `yaml:"test_receiver"`
	ExampleSender string `yaml:"example_sender"`
	LinkMint      string `yaml:"link_mint"`
	PrivateKeys   struct {
		Admin    []byte `yaml:"admin"`
		Deployer []byte `yaml:"deployer"`
	} `yaml:"private_keys"`
	ChainSelectors struct {
		Sepolia uint64 `yaml:"sepolia"`
		Svm     uint64 `yaml:"svm"`
	} `yaml:"chain_selectors"`
	CCTP struct {
		TokenPool            string `yaml:"token_pool"`
		UsdcMint             string `yaml:"usdc_mint"`
		MessageTransmitter   string `yaml:"message_transmitter"`
		TokenMessengerMinter string `yaml:"token_messenger_minter"`
		Message              struct {
			MessageBytesHex     string `yaml:"message_bytes_hex"`
			AttestationBytesHex string `yaml:"attestation_bytes_hex"`
		} `yaml:"message"`
		Sepolia struct {
			RecipientBase58 string `yaml:"recipient_base58"`
			ReceiverAddress string `yaml:"receiver_address"`
			TokenPool       string `yaml:"token_pool"`
			AllowedCaller   string `yaml:"allowed_caller"`
		} `yaml:"sepolia"`
	} `yaml:"cctp"`
	RPC string `yaml:"rpc"`
}

func getDevnetInfo() (DevnetInfo, error) {
	var devnetInfo DevnetInfo
	if err := yaml.Unmarshal(devnetInfoBuffer, &devnetInfo); err != nil {
		return DevnetInfo{}, err
	}
	// fmt.Printf("Devnet info: %+v\n", devnetInfo)
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
