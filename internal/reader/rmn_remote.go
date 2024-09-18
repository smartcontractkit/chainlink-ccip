package reader

import (
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type RMNRemote interface {
	GetMinSigners() uint64
	GetSignersInfo() []rmntypes.RMNRemoteSignerInfo
	GetRmnReportVersion() string
	GetRmnRemoteContractAddress() string
	GetRmnHomeConfigDigest() cciptypes.Bytes32
}

type RmnRemotePoller struct {
	rmnRemoteConfig rmntypes.RMNRemoteConfig
}

func NewRMNRemotePoller() RMNRemote {
	return &RmnRemotePoller{
		rmnRemoteConfig: rmntypes.RMNRemoteConfig{},
	}
}

func (r *RmnRemotePoller) GetMinSigners() uint64 {
	panic("implement me")
}

func (r *RmnRemotePoller) GetSignersInfo() []rmntypes.RMNRemoteSignerInfo {
	panic("implement me")
}

func (r *RmnRemotePoller) GetRmnReportVersion() string {
	panic("implement me")
}

func (r *RmnRemotePoller) GetRmnRemoteContractAddress() string {
	panic("implement me")
}

func (r *RmnRemotePoller) GetRmnHomeConfigDigest() cciptypes.Bytes32 {
	panic("implement me")
}

var _ RMNRemote = (*RmnRemotePoller)(nil)
