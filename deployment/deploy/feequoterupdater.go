package deploy

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type FeeQuoterUpdateInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	ContractMeta      []datastore.ContractMetadata
}
