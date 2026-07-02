package changesets

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/operations/fetch_signing_keys"
)

// deriveFamiliesFromSelectors returns the unique chain families for the given selectors.
func deriveFamiliesFromSelectors(selectors []uint64) []string {
	seen := make(map[string]struct{})
	for _, sel := range selectors {
		if family, err := chainsel.GetSelectorFamily(sel); err == nil {
			seen[family] = struct{}{}
		}
	}
	families := make([]string, 0, len(seen))
	for f := range seen {
		families = append(families, f)
	}
	return families
}

// fetchSigningKeysForNOPsByFamilies fetches JD signing keys for NOPs that are
// missing a configured signer address for any of the given families.
func fetchSigningKeysForNOPsByFamilies(e deployment.Environment, nops []offchain.NOPConfig, families []string) (fetch_signing_keys.SigningKeysByNOP, map[string]string, error) {
	return fetchSigningKeysForNOPsFiltered(e, nops, func(nop offchain.NOPConfig) bool {
		for _, family := range families {
			if nop.SignerAddressByFamily == nil || nop.SignerAddressByFamily[family] == "" {
				return true
			}
		}
		return false
	})
}

func fetchSigningKeysForNOPsFiltered(
	e deployment.Environment,
	nops []offchain.NOPConfig,
	include func(offchain.NOPConfig) bool,
) (fetch_signing_keys.SigningKeysByNOP, map[string]string, error) {
	if e.Offchain == nil {
		return nil, nil, nil
	}

	aliases := make([]string, 0, len(nops))
	for _, nop := range nops {
		if include(nop) {
			aliases = append(aliases, nop.Alias)
		}
	}

	if len(aliases) == 0 {
		return nil, nil, nil
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		fetch_signing_keys.FetchNOPSigningKeys,
		fetch_signing_keys.FetchSigningKeysDeps{
			JDClient: e.Offchain,
			Logger:   e.Logger,
			NodeIDs:  e.NodeIDs,
		},
		fetch_signing_keys.FetchSigningKeysInput{
			NOPAliases: aliases,
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch signing keys from JD for NOPs %v: %w", aliases, err)
	}

	return report.Output.SigningKeysByNOP, report.Output.RawPubKeyByNOP, nil
}

// getSignatureConfigForLane builds the committee verifier signature quorum
// config (signers + threshold) for a local→remote lane from topology + JD keys.
func getSignatureConfigForLane(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	committeeQualifier string,
	localSelector uint64,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
	rawPubKeyByNOP map[string]string,
) (*adapters.CommitteeVerifierSignatureQuorumConfig, error) {
	committee, ok := topology.NOPTopology.Committees[committeeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found", committeeQualifier)
	}

	chainCfg, ok := committee.ChainConfigs[strconv.FormatUint(remoteSelector, 10)]
	if !ok {
		return nil, fmt.Errorf("chain selector %d not found in committee %q", remoteSelector, committeeQualifier)
	}

	localFamily, err := chainsel.GetSelectorFamily(localSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get selector family for selector %d: %w", localSelector, err)
	}

	signers := make([]string, 0, len(chainCfg.NOPAliases))
	for _, alias := range chainCfg.NOPAliases {
		signer, err := signerAddressForNOPAlias(e, topology, alias, localFamily, committeeQualifier, remoteSelector, signingKeysByNOP, rawPubKeyByNOP[alias])
		if err != nil {
			return nil, err
		}
		signers = append(signers, signer)
	}

	return &adapters.CommitteeVerifierSignatureQuorumConfig{
		Threshold: chainCfg.Threshold,
		Signers:   signers,
	}, nil
}

func signerAddressForNOPAlias(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	alias string,
	localFamily string,
	committeeQualifier string,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
	rawPubKey string,
) (string, error) {
	nop, ok := topology.NOPTopology.GetNOP(alias)
	if !ok {
		return "", fmt.Errorf(
			"NOP alias %q not found for committee %q chain %d",
			alias, committeeQualifier, remoteSelector,
		)
	}

	if nop.SignerAddressByFamily != nil {
		if addr := nop.SignerAddressByFamily[localFamily]; addr != "" {
			return addr, nil
		}
	}

	if signer, ok := signerFromJDIfMissing(
		nop.SignerAddressByFamily,
		alias,
		localFamily,
		signingKeysByNOP,
	); ok {
		e.Logger.Debugw("Using signing address from JD",
			"nopAlias", alias,
			"chainFamily", localFamily,
			"signerAddress", signer,
		)
		return signer, nil
	}

	// No exact match for localFamily. A standalone (bootstrap-based) NOP signs with a
	// secp256k1 key that may be registered under a different family.
	// This key can often be translated into localFamily's representation without any new data.
	if signer, sourceFamily, translateErr := translateFromKnownFamily(
		nop.SignerAddressByFamily, signingKeysByNOP[alias], rawPubKey, localFamily,
	); signer != "" {
		e.Logger.Debugw("Translated signing address from a different chain family",
			"nopAlias", alias,
			"sourceFamily", sourceFamily,
			"targetFamily", localFamily,
			"signerAddress", signer,
		)
		return signer, nil
	} else if translateErr != nil {
		return "", fmt.Errorf(
			"NOP %q missing signer_address for family %s on committee %q chain %d: %w",
			alias, localFamily, committeeQualifier, remoteSelector, translateErr,
		)
	}

	return "", fmt.Errorf(
		"NOP %q missing signer_address for family %s on committee %q chain %d",
		alias, localFamily, committeeQualifier, remoteSelector,
	)
}

// addressClassFamilies are chain families whose signer identity is a 20-byte address
// derived via secp256k1 -> keccak256 (EVM-style, e.g. ecrecover). Members encode the
// identical bytes and differ only in string formatting, so they are freely
// interconvertible in both directions.
var addressClassFamilies = map[string]bool{
	chainsel.FamilyEVM:    true,
	chainsel.FamilySolana: true,
}

// rawPubKeySourceFamily is a synthetic "family" tag used only as a translation source,
// standing for a NOP's raw public key fetched directly from JD rather than derived from
// any one chain family's own address representation (see FetchSigningKeysOutput.RawPubKeyByNOP).
// It is never a real chain family and must never be passed as a translation target.
const rawPubKeySourceFamily = "xxx_notarealfamily_raw_pubkey"

// rawPubKeyClassFamilies are chain families (plus the synthetic rawPubKeySourceFamily)
// whose signer identity is the full uncompressed secp256k1 public key, hex-encoded with
// no per-family formatting differences. Members are interchangeable as-is.
var rawPubKeyClassFamilies = map[string]bool{
	chainsel.FamilyAptos:   true,
	chainsel.FamilyStellar: true,
	chainsel.FamilyCanton:  true,
	rawPubKeySourceFamily:  true,
}

// translateFromKnownFamily tries to derive localFamily's signer address from whichever
// family(ies) the NOP does have a signer address for (static topology, JD, and/or the
// NOP's raw public key), preferring a deterministic order when more than one is
// available.
//
// It returns the derived address and the family it came from on success. If a
// candidate was found but translation into localFamily isn't possible (e.g. deriving a
// raw public key back out of an already-hashed address), err explains why, so the caller
// can surface something more actionable than a generic "missing" error.
func translateFromKnownFamily(
	staticByFamily map[string]string,
	jdByFamily map[string]string,
	rawPubKey string,
	localFamily string,
) (address string, sourceFamily string, err error) {
	candidates := make(map[string]string, len(staticByFamily)+len(jdByFamily)+1)
	for family, addr := range staticByFamily {
		if addr != "" {
			candidates[family] = addr
		}
	}
	for family, addr := range jdByFamily {
		if addr != "" {
			if _, exists := candidates[family]; !exists {
				candidates[family] = addr
			}
		}
	}
	if rawPubKey != "" {
		candidates[rawPubKeySourceFamily] = rawPubKey
	}
	delete(candidates, localFamily)

	families := make([]string, 0, len(candidates))
	for family := range candidates {
		families = append(families, family)
	}
	sort.Strings(families) // deterministic when a NOP declares more than one other family

	var lastErr error
	for _, family := range families {
		translated, tErr := translateSignerAddress(family, candidates[family], localFamily)
		if tErr == nil {
			return translated, family, nil
		}
		lastErr = tErr
	}
	return "", "", lastErr
}

// translateSignerAddress converts a signer identity stored under sourceFamily into the
// representation targetFamily expects, valid only when both families sign with the same
// underlying secp256k1 key.
//
// Address-class families (evm, solana) store the same 20 bytes and translate in both
// directions. Raw-pubkey-class families (aptos, stellar, canton) store the same
// hex-encoded public key and translate in both directions. Deriving an address-class
// value from a raw public key is one-directional: an address is a keccak256 hash of the
// public key, so recovering the public key from an address is not possible — that
// direction returns an error instead of a wrong answer.
func translateSignerAddress(sourceFamily, value, targetFamily string) (string, error) {
	switch {
	case addressClassFamilies[sourceFamily] && addressClassFamilies[targetFamily]:
		addrBytes, err := decodeAddressHex(value)
		if err != nil {
			return "", fmt.Errorf("decode %s signer address: %w", sourceFamily, err)
		}
		return formatAddressForFamily(addrBytes, targetFamily)
	case rawPubKeyClassFamilies[sourceFamily] && rawPubKeyClassFamilies[targetFamily]:
		return strings.ToLower(strings.TrimPrefix(value, "0x")), nil
	case rawPubKeyClassFamilies[sourceFamily] && addressClassFamilies[targetFamily]:
		pubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(strings.ToLower(value), "0x"))
		if err != nil {
			return "", fmt.Errorf("decode %s raw public key: %w", sourceFamily, err)
		}
		pubKey, err := gethcrypto.UnmarshalPubkey(pubKeyBytes)
		if err != nil {
			return "", fmt.Errorf("unmarshal %s public key: %w", sourceFamily, err)
		}
		return formatAddressForFamily(gethcrypto.PubkeyToAddress(*pubKey).Bytes(), targetFamily)
	case addressClassFamilies[sourceFamily] && rawPubKeyClassFamilies[targetFamily]:
		return "", fmt.Errorf(
			"cannot derive a %s raw public key from a %s address: address derivation (keccak256) is not "+
				"reversible; register the node's raw public key for family %s directly",
			targetFamily, sourceFamily, targetFamily)
	default:
		return "", fmt.Errorf("no signer address translation defined from family %s to family %s", sourceFamily, targetFamily)
	}
}

// decodeAddressHex parses a 20-byte address from hex, tolerating an optional 0x prefix
// and either case (JD-fetched addresses are lowercased and 0x-prefixed unconditionally;
// statically-configured ones may be EIP-55 checksummed or, for solana, bare lowercase).
func decodeAddressHex(s string) ([]byte, error) {
	s = strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(b) != 20 {
		return nil, fmt.Errorf("expected 20-byte address, got %d bytes", len(b))
	}
	return b, nil
}

// formatAddressForFamily renders a 20-byte address per family convention: EIP-55
// checksummed with 0x for evm, lowercase without 0x for solana (matching the chainlink
// node's prior art in its Solana OCR2 keyring).
func formatAddressForFamily(addr []byte, family string) (string, error) {
	switch family {
	case chainsel.FamilyEVM:
		return gethcommon.BytesToAddress(addr).Hex(), nil
	case chainsel.FamilySolana:
		return strings.ToLower(hex.EncodeToString(addr)), nil
	default:
		return "", fmt.Errorf("unsupported address-class family %q", family)
	}
}

func signerFromJDIfMissing(
	signerAddresses map[string]string,
	nopAlias string,
	family string,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (string, bool) {
	if signerAddresses != nil && signerAddresses[family] != "" {
		return "", false
	}

	if signingKeysByNOP == nil {
		return "", false
	}

	if signer := signingKeysByNOP[nopAlias][family]; signer != "" {
		return signer, true
	}

	return "", false
}
