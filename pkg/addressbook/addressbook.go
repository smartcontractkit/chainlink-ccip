package addressbook

import (
	"errors"
	"sync"

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

// IBook represents an address book operations for a Decentralized Oracle Network (DON),
// can be used either for commit or execute plugin. It ensures that all oracles within the network have a
// consistent and synchronized view of the address book.
type IBook interface {
	GetContractAddress(name ContractName, chain ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error)
	AppendState(addresses ContractAddresses) error
}

type Book struct {
	mem ContractAddresses
	mu  *sync.RWMutex
}

func NewBook() *Book {
	return &Book{
		mem: make(ContractAddresses),
		mu:  &sync.RWMutex{},
	}
}

func (b *Book) GetContractAddress(name ContractName, chain ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error) {
	b.mu.RLock()
	if len(b.mem) == 0 {
		return nil, ErrAddressBookNotInitialized
	}
	contractOnAllChains, ok := b.mem[name]
	b.mu.RUnlock()

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

func (b *Book) AppendState(addresses ContractAddresses) error {
	b.mu.Lock()
	for name, chains := range addresses {
		if _, ok := b.mem[name]; !ok {
			// if contract does not exist just set the state
			b.mem[name] = chains
			continue
		}

		// if contract exists, set or replace any existing address for each chain
		for chain, addr := range chains {
			b.mem[name][chain] = addr
		}
	}
	b.mu.Unlock()
	return nil
}
