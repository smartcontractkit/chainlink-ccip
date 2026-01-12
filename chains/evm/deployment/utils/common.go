package utils

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/type_and_version"
)

func TypeAndVersion(addr common.Address, client bind.ContractBackend) (string, *semver.Version, error) {
	tv, err := type_and_version.NewITypeAndVersion(addr, client)
	if err != nil {
		return "", nil, err
	}
	tvStr, err := tv.TypeAndVersion(nil)
	if err != nil {
		return "", nil, fmt.Errorf("error calling typeAndVersion on addr: %s %w", addr.String(), err)
	}

	contractType, versionStr, err := ParseTypeAndVersion(tvStr)
	if err != nil {
		return "", nil, err
	}
	v, err := semver.NewVersion(versionStr)
	if err != nil {
		return "", nil, fmt.Errorf("failed parsing version %s: %w", versionStr, err)
	}
	return contractType, v, nil
}

func ParseTypeAndVersion(tvStr string) (string, string, error) {
	if tvStr == "" {
		return "", "", fmt.Errorf("type and version string is empty")
	}
	typeAndVersionValues := strings.Split(tvStr, " ")

	if len(typeAndVersionValues) < 2 {
		return "", "", fmt.Errorf("invalid type and version %s", tvStr)
	}
	return typeAndVersionValues[0], typeAndVersionValues[1], nil
}
