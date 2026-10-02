package shared

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// This file carries the four symbols the Solana changesets need from
// chainlink/deployment/ccip/shared's helpers.go, token_pools.go and token_info.go.
//
// Those three files are deliberately not copied. Between them they import
// chainlink/v2/core/utils/typeandversion (the node core module),
// chainlink-evm/gethwrappers (erc20, aggregator_v3_interface, capabilities_registry) and
// chains/evm/deployment — none of which the Solana code touches, and pulling them in would
// give this module a dependency on the chainlink node, which is the thing keeping the port
// out of chainlink-ccip in the first place.
//
// The symbols below are reproduced verbatim; if the originals change, these must follow
// until chainlink's copies are deleted.

// Versions of contracts this package refers to. Inlined from chainlink/deployment/version.go
// rather than imported, for the same reason: that package lives in the chainlink module.
var (
	Version1_5_0    = *semver.MustParse("1.5.0")
	Version1_5_1    = *semver.MustParse("1.5.1")
	Version1_6_0    = *semver.MustParse("1.6.0")
	Version1_6_1    = *semver.MustParse("1.6.1")
	Version1_6_2    = *semver.MustParse("1.6.2")
	Version1_6_3Dev = *semver.MustParse("1.6.3-dev")
	Version2_0_0    = *semver.MustParse("2.0.0")
)

// FastTransferTokenPoolVersion is the version fast-transfer token pools are deployed at.
var FastTransferTokenPoolVersion = Version1_6_3Dev

// TokenSymbol is the symbol of a token, e.g. "USDC". It is the token-scoped datastore
// qualifier for the rows that belong to one token.
type TokenSymbol string

// QualifierFromParts joins the parts that identify one instance into a datastore qualifier.
//
// Each part is quoted rather than joined on a bare separator: parts are often caller-supplied
// free-form text, so a plain join would let ("A", "B/C") and ("A/B", "C") produce the same
// qualifier and collide. Quoting escapes any separator inside a part, so distinct inputs always
// yield distinct qualifiers.
func QualifierFromParts(parts ...string) string {
	quoted := make([]string, 0, len(parts))
	for _, p := range parts {
		quoted = append(quoted, fmt.Sprintf("%q", p))
	}

	return strings.Join(quoted, "/")
}

// TokenPoolLookupTableQualifier returns the datastore qualifier for a Solana token-pool lookup
// table, which is uniquely identified by (token mint, pool type, metadata).
//
// Each component is quoted rather than joined on a bare separator. metadata is caller-supplied
// free-form text, so a plain "a/b/c" join would let ("A", "B/C") and ("A/B", "C") produce the same
// qualifier and collide in the datastore. Quoting escapes any separator inside a component, so
// distinct inputs always yield distinct qualifiers.
func TokenPoolLookupTableQualifier(tokenPubKey, poolType, metadata string) string {
	return QualifierFromParts(tokenPubKey, poolType, metadata)
}

// TokenPoolTypes is the set of contract types that are token pools.
var TokenPoolTypes = map[cldf.ContractType]struct{}{
	BurnMintFastTransferTokenPool:                   {},
	BurnMintTokenPool:                               {},
	BurnWithFromMintTokenPool:                       {},
	BurnFromMintTokenPool:                           {},
	LockReleaseTokenPool:                            {},
	USDCTokenPool:                                   {},
	HybridLockReleaseUSDCTokenPool:                  {},
	BurnMintWithExternalMinterFastTransferTokenPool: {},
	HybridWithExternalMinterFastTransferTokenPool:   {},
	BurnMintWithExternalMinterTokenPool:             {},
	HybridWithExternalMinterTokenPool:               {},
	USDCTokenPoolProxy:                              {},
}

// TokenPoolVersions is the set of versions token pools are deployed at.
var TokenPoolVersions = map[semver.Version]struct{}{
	Version1_5_0:                 {},
	Version1_5_1:                 {},
	FastTransferTokenPoolVersion: {},
	Version1_6_0:                 {},
	Version1_6_1:                 {},
	Version1_6_2:                 {},
	Version2_0_0:                 {},
}

// LinkToken is the burn/mint LINK token. Reproduced from
// chainlink/deployment/common/types, which is the only symbol the Solana state view needed from
// that package.
const LinkToken cldf.ContractType = "LinkToken"
