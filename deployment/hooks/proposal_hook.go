package hooks

import (
	"context"
	"time"

	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
)

func GlobalPostProposalCCIPSendHookForEVM(dom domain.Domain) cldf_changeset.PostProposalHook {
	return cldf_changeset.PostProposalHook{
		HookDefinition: cldf_changeset.HookDefinition{
			Name:          "verify-ccip-send",
			FailurePolicy: cldf_changeset.Abort,
			Timeout:       5 * time.Minute,
		},
		Func: verifyCCIPSend,
	}
}

func verifyCCIPSend(ctx context.Context, params cldf_changeset.PostProposalHookParams) error {
	env := params.Env
	
}
