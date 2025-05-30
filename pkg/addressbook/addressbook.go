package addressbook

import (
	"errors"
	"sync"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	ErrAddressBookNotInitialized = errors.New("address book is not initialized yet")
	ErrContractNotRegistered     = errors.New("contract type not registered on the address book")
	ErrContractAddressNotFound   = errors.New("address not found for the specified contract")
	ErrContractAddressEmpty      = errors.New("contract address is empty")
)

type ContractName string
type ContractAddresses map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress

// DonRead represents an address book read operations for a Decentralized Oracle Network (DON),
// either for commit or execute plugin. It ensures that all oracles within the network have a
// consistent and synchronized view of the address book.
type DonRead interface {
	GetOnRampAddress(sourceChain ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error)
	GetOffRampAddress() (ccipocr3.UnknownAddress, error)
}

// DonWrite represents an address book write operations for a Decentralized Oracle Network (DON).
// It allows the address book to be updated with new contract addresses, ensuring that all oracles
// within the network can access the latest information.
type DonWrite interface {
	SetState(addresses ContractAddresses) error
}

type DON struct {
	mem       ContractAddresses
	destChain ccipocr3.ChainSelector
	mu        *sync.RWMutex
}

func NewDon(destChain ccipocr3.ChainSelector) *DON {
	return &DON{
		mem:       make(ContractAddresses),
		destChain: destChain,
		mu:        &sync.RWMutex{},
	}
}

func (d *DON) GetOnRampAddress(sourceChain ccipocr3.ChainSelector) ([]byte, error) {
	return d.getContractAddress(consts.ContractNameOnRamp, sourceChain)
}

func (d *DON) GetOffRampAddress() ([]byte, error) {
	return d.getContractAddress(consts.ContractNameOffRamp, d.destChain)
}

func (d *DON) getContractAddress(contractName ContractName, chain ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error) {
	d.mu.RLock()
	if len(d.mem) == 0 {
		return nil, ErrAddressBookNotInitialized
	}
	contractOnAllChains, ok := d.mem[contractName]
	d.mu.RUnlock()

	if !ok {
		return nil, ErrContractNotRegistered
	}

	contractAddr, ok := contractOnAllChains[chain]
	if !ok {
		return nil, ErrContractAddressNotFound
	}

	if contractAddr.IsZeroOrEmpty() {
		return nil, ErrContractAddressEmpty
	}

	return contractAddr, nil
}

func (d *DON) SetState(addresses ContractAddresses) error {
	d.mu.Lock()
	d.mem = addresses
	d.mu.Unlock()
	return nil
}
