package v1_7

import (
	"fmt"

	"call-orchestrator-demo/views"

	"github.com/ethereum/go-ethereum/common"

	committee_verifier "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
)

// packCommitteeVerifierCall packs a method call using the CommitteeVerifier ABI and returns the calldata bytes.
func packCommitteeVerifierCall(method string, args ...interface{}) ([]byte, error) {
	return CommitteeVerifierABI.Pack(method, args...)
}

// executeCommitteeVerifierCall packs a call, executes it via the orchestrator, and returns raw response bytes.
func executeCommitteeVerifierCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packCommitteeVerifierCall(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("%s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getCommitteeVerifierOwner fetches the owner address using the CommitteeVerifier bindings.
func getCommitteeVerifierOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeCommitteeVerifierCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := CommitteeVerifierABI.Unpack("owner", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack owner response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from owner call")
	}

	owner, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for owner: %T", results[0])
	}

	return owner.Hex(), nil
}

// getCommitteeVerifierTypeAndVersion fetches the typeAndVersion string using the CommitteeVerifier bindings.
func getCommitteeVerifierTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeCommitteeVerifierCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := CommitteeVerifierABI.Unpack("typeAndVersion", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack typeAndVersion response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from typeAndVersion call")
	}

	typeAndVersion, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for typeAndVersion: %T", results[0])
	}

	return typeAndVersion, nil
}

// getCommitteeVerifierAllSignatureConfigs fetches all signature configs using the CommitteeVerifier bindings.
func getCommitteeVerifierAllSignatureConfigs(ctx *views.ViewContext) ([]committee_verifier.SignatureQuorumValidatorSignatureConfig, error) {
	data, err := executeCommitteeVerifierCall(ctx, "getAllSignatureConfigs")
	if err != nil {
		return nil, err
	}

	results, err := CommitteeVerifierABI.Unpack("getAllSignatureConfigs", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllSignatureConfigs response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getAllSignatureConfigs call")
	}

	configs, ok := results[0].([]struct {
		SourceChainSelector uint64           `json:"sourceChainSelector"`
		Threshold           uint8            `json:"threshold"`
		Signers             []common.Address `json:"signers"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for getAllSignatureConfigs: %T", results[0])
	}

	result := make([]committee_verifier.SignatureQuorumValidatorSignatureConfig, len(configs))
	for i, cfg := range configs {
		result[i] = committee_verifier.SignatureQuorumValidatorSignatureConfig{
			SourceChainSelector: cfg.SourceChainSelector,
			Threshold:           cfg.Threshold,
			Signers:             cfg.Signers,
		}
	}

	return result, nil
}

// collectSignatureConfigs fetches signature configs and formats them for the view.
func collectSignatureConfigs(ctx *views.ViewContext) ([]map[string]any, error) {
	configs, err := getCommitteeVerifierAllSignatureConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get signature configs: %w", err)
	}

	signatureConfigs := make([]map[string]any, 0, len(configs))

	for _, cfg := range configs {
		signersHex := make([]string, len(cfg.Signers))
		for i, signer := range cfg.Signers {
			signersHex[i] = signer.Hex()
		}

		configInfo := map[string]any{
			"sourceChainSelector": cfg.SourceChainSelector,
			"threshold":           cfg.Threshold,
			"signers":             signersHex,
		}

		signatureConfigs = append(signatureConfigs, configInfo)
	}

	return signatureConfigs, nil
}

// ViewCommitteeVerifier generates a view of the CommitteeVerifier contract (v1.7.0).
// Uses the generated bindings to pack/unpack calls.
func ViewCommitteeVerifier(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.7.0"

	owner, err := getCommitteeVerifierOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getCommitteeVerifierTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	signatureConfigs, err := collectSignatureConfigs(ctx)
	if err != nil {
		result["signatureConfigs_error"] = err.Error()
	} else {
		result["signatureConfigs"] = signatureConfigs
	}

	return result, nil
}
