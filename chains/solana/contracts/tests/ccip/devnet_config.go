package contracts

import (
	_ "embed"

	"fmt"

	"github.com/gagliardetto/solana-go"
	"gopkg.in/yaml.v2"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

//go:embed devnet.config.yaml
var devnetInfoBuffer []byte

type DevnetInfo struct {
	Offramp     string `yaml:"offramp"`
	PrivateKeys struct {
		User        []byte   `yaml:"user"`
		Admin       []byte   `yaml:"admin"`
		Transmitter []byte   `yaml:"transmitter"`
		Signers     []string `yaml:"signers"` // hex
	} `yaml:"private_keys"`
	ChainSelectors struct {
		Sepolia uint64 `yaml:"sepolia"`
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
