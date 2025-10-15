package utils

import (
	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func MaybeDeployContract(
	b operations.Bundle,
	chain cldf_solana.Chain,
	input []datastore.AddressRef,
	contractType cldf_deployment.ContractType,
	contractVersion *semver.Version,
	contractQualifier string,
	programName string,
	programSize int,
) (datastore.AddressRef, error) {
	for _, ref := range input {
		if ref.Type == datastore.ContractType(contractType) &&
			ref.Version.Equal(contractVersion) {
			if contractQualifier != "" {
				if ref.Qualifier == contractQualifier {
					return ref, nil
				}
			} else {
				return ref, nil
			}
		}
	}
	programID, err := chain.DeployProgram(b.Logger, cldf_solana.ProgramInfo{
		Name:  programName,
		Bytes: programSize,
	}, false, true)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	// validate deployed programID
	_ = solana.MustPublicKeyFromBase58(programID)

	return datastore.AddressRef{
		Address:       programID,
		ChainSelector: chain.Selector,
		Type:          datastore.ContractType(contractType),
		Version:       contractVersion,
	}, nil
}
