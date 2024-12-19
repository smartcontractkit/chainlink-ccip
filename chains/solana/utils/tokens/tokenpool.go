package tokens

import (
	"context"
	"encoding/binary"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

type TokenPool struct {
	// token details
	Program solana.PublicKey
	Mint    solana.PrivateKey

	// admin registry PDA
	AdminRegistry solana.PublicKey

	// pool details
	PoolProgram, PoolConfig, PoolSigner, PoolTokenAccount solana.PublicKey
	PoolLookupTable                                       solana.PublicKey

	AdditionalAccounts solana.PublicKeySlice

	// AssociatedTokenAddress Lookups
	User map[solana.PublicKey]solana.PublicKey

	// remote chain config lookup
	Chain map[uint64]solana.PublicKey

	// billing config lookup
	Billing map[uint64]solana.PublicKey
}

func (tp TokenPool) ToTokenPoolEntries() []solana.PublicKey {
	list := solana.PublicKeySlice{
		tp.PoolLookupTable,
		tp.AdminRegistry,
		tp.PoolProgram,
		tp.PoolConfig,
		tp.PoolTokenAccount,
		tp.PoolSigner,
		tp.Program,
		tp.Mint.PublicKey(),
	}
	return append(list, tp.AdditionalAccounts...)
}

// NewTokenPool returns token + pool addresses. however, the token still needs to be deployed
func NewTokenPool(program solana.PublicKey) (TokenPool, error) {
	mint, err := solana.NewRandomPrivateKey()
	if err != nil {
		return TokenPool{}, err
	}
	tokenAdminRegistryPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.PublicKey().Bytes()}, config.CcipRouterProgram)
	if err != nil {
		return TokenPool{}, err
	}

	// preload with defined config.EvmChainSelector
	chainPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_chainconfig"), binary.LittleEndian.AppendUint64([]byte{}, config.EvmChainSelector), mint.PublicKey().Bytes()}, config.CcipTokenPoolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	billingPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_billing"), binary.LittleEndian.AppendUint64([]byte{}, config.EvmChainSelector), mint.PublicKey().Bytes()}, config.CcipRouterProgram)
	if err != nil {
		return TokenPool{}, err
	}

	p := TokenPool{
		Program:         program,
		Mint:            mint,
		AdminRegistry:   tokenAdminRegistryPDA,
		PoolLookupTable: solana.PublicKey{},
		User:            map[solana.PublicKey]solana.PublicKey{},
		Chain:           map[uint64]solana.PublicKey{},
		Billing:         map[uint64]solana.PublicKey{},
	}
	p.Chain[config.EvmChainSelector] = chainPDA
	p.Billing[config.EvmChainSelector] = billingPDA
	return p, nil
}

func (tp *TokenPool) SetupLookupTable(ctx context.Context, client *rpc.Client, admin solana.PrivateKey) error {
	table, err := common.CreateLookupTable(ctx, client, admin)
	if err != nil {
		return err
	}
	tp.PoolLookupTable = table // the LUT entries will include this, so set it before adding the addresses to the LUT
	if err = common.ExtendLookupTable(ctx, client, table, admin, tp.ToTokenPoolEntries()); err != nil {
		return err
	}
	return common.AwaitSlotChange(ctx, client)
}

func TokenPoolConfigAddress(token solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_config"), token.Bytes()}, config.CcipTokenPoolProgram)
	return addr, err
}

func TokenPoolSignerAddress(token solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_signer"), token.Bytes()}, config.CcipTokenPoolProgram)
	return addr, err
}

type EventBurnLock struct {
	Discriminator [8]byte
	Sender        solana.PublicKey
	Amount        uint64
}

type EventMintRelease struct {
	Discriminator [8]byte
	Sender        solana.PublicKey
	Recipient     solana.PublicKey
	Amount        uint64
}

type EventChainConfigured struct {
	Discriminator       [8]byte
	ChainSelector       uint64
	Token               []byte
	PreviousToken       []byte
	PoolAddress         []byte
	PreviousPoolAddress []byte
}

type EventRateLimitConfigured struct {
	Discriminator     [8]byte
	ChainSelector     uint64
	OutboundRateLimit token_pool.RateLimitConfig
	InboundRateLimit  token_pool.RateLimitConfig
}

type EventChainRemoved struct {
	Discriminator [8]byte
	ChainSelector uint64
}

type EventRouterUpdated struct {
	Discriminator [8]byte
	OldAuthority  solana.PublicKey
	NewAuthority  solana.PublicKey
}

func MethodToEvent(m string) string {
	mapping := map[string]string{
		"lock":    "Locked",
		"release": "Released",
		"burn":    "Burned",
		"mint":    "Minted",
	}
	return mapping[strings.ToLower(m)]
}

func ParseTokenLookupTable(ctx context.Context, client *rpc.Client, token TokenPool, userTokenAccount solana.PublicKey) (solana.AccountMetaSlice, map[solana.PublicKey]solana.PublicKeySlice, error) {
	tokenBillingConfig := token.Billing[config.EvmChainSelector]
	poolChainConfig := token.Chain[config.EvmChainSelector]

	lookupTableEntries, err := common.GetAddressLookupTable(ctx, client, token.PoolLookupTable)
	if err != nil {
		return nil, nil, err
	}

	list := []*solana.AccountMeta{
		solana.Meta(userTokenAccount).WRITE(),
		solana.Meta(tokenBillingConfig),
		solana.Meta(poolChainConfig).WRITE(),
		solana.Meta(lookupTableEntries[0]),         // lookup table
		solana.Meta(lookupTableEntries[1]),         // token admin registry
		solana.Meta(lookupTableEntries[2]),         // PoolProgram
		solana.Meta(lookupTableEntries[3]).WRITE(), // PoolConfig
		solana.Meta(lookupTableEntries[4]).WRITE(), // PoolTokenAccount
		solana.Meta(lookupTableEntries[5]),         // PoolSigner
		solana.Meta(lookupTableEntries[6]),         // TokenProgram
		solana.Meta(lookupTableEntries[7]).WRITE(), // Mint
	}

	for _, v := range token.AdditionalAccounts {
		list = append(list, solana.Meta(v))
	}

	addressTables := make(map[solana.PublicKey]solana.PublicKeySlice)
	addressTables[token.PoolLookupTable] = lookupTableEntries

	return list, addressTables, nil
}
