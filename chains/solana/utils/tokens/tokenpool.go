package tokens

import (
	"context"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

type TokenPool struct {
	// token details
	Program        solana.PublicKey
	Mint           solana.PrivateKey
	FeeTokenConfig solana.PublicKey

	// admin registry PDA
	AdminRegistryPDA solana.PublicKey

	// pool details
	PoolProgram, PoolConfig, PoolSigner, PoolTokenAccount solana.PublicKey
	PoolLookupTable                                       solana.PublicKey
	WritableIndexes                                       []uint8

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
		tp.PoolLookupTable,  // 0
		tp.AdminRegistryPDA, // 1
		tp.PoolProgram,      // 2
		tp.PoolConfig,       // 3 - writable
		tp.PoolTokenAccount, // 4 - writable
		tp.PoolSigner,       // 5
		tp.Program,          // 6
		tp.Mint.PublicKey(), // 7 - writable
		tp.FeeTokenConfig,   // 8
	}
	return append(list, tp.AdditionalAccounts...)
}

// NewTokenPool returns token + pool addresses. however, the token still needs to be deployed
func NewTokenPool(program solana.PublicKey) (TokenPool, error) {
	mint, err := solana.NewRandomPrivateKey()
	if err != nil {
		return TokenPool{}, err
	}
	tokenAdminRegistryPDA, _, err := state.FindTokenAdminRegistryPDA(mint.PublicKey(), config.CcipRouterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	// preload with defined config.EvmChainSelector
	chainPDA, _, err := TokenPoolChainConfigPDA(config.EvmChainSelector, mint.PublicKey(), config.CcipTokenPoolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	billingPDA, _, err := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, mint.PublicKey(), config.FeeQuoterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	tokenConfigPda, _, err := state.FindFqBillingTokenConfigPDA(mint.PublicKey(), config.FeeQuoterProgram)
	if err != nil {
		return TokenPool{}, err
	}

	p := TokenPool{
		Program:          program,
		Mint:             mint,
		FeeTokenConfig:   tokenConfigPda,
		AdminRegistryPDA: tokenAdminRegistryPDA,
		PoolProgram:      config.CcipTokenPoolProgram,
		PoolLookupTable:  solana.PublicKey{},
		WritableIndexes:  []uint8{3, 4, 7}, // see ToTokenPoolEntries for writable indexes
		User:             map[solana.PublicKey]solana.PublicKey{},
		Chain:            map[uint64]solana.PublicKey{},
		Billing:          map[uint64]solana.PublicKey{},
	}
	p.Chain[config.EvmChainSelector] = chainPDA
	p.Billing[config.EvmChainSelector] = billingPDA
	p.PoolConfig, err = TokenPoolConfigAddress(p.Mint.PublicKey(), config.CcipTokenPoolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	p.PoolSigner, err = TokenPoolSignerAddress(p.Mint.PublicKey(), config.CcipTokenPoolProgram)
	if err != nil {
		return TokenPool{}, err
	}
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

func TokenPoolConfigAddress(token, programID solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_config"), token.Bytes()}, programID)
	return addr, err
}

func TokenPoolSignerAddress(token, programID solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_signer"), token.Bytes()}, programID)
	return addr, err
}

func TokenPoolChainConfigPDA(chainSelector uint64, mint, programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_chainconfig"), chainSelectorLE, mint.Bytes()}, programID)
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
	Discriminator         [8]byte
	ChainSelector         uint64
	Token                 token_pool.RemoteAddress
	PreviousToken         token_pool.RemoteAddress
	PoolAddresses         []token_pool.RemoteAddress
	PreviousPoolAddresses []token_pool.RemoteAddress
}

type EventRemotePoolsAppended struct {
	Discriminator         [8]byte
	ChainSelector         uint64
	PoolAddresses         []token_pool.RemoteAddress
	PreviousPoolAddresses []token_pool.RemoteAddress
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

	tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
	err := common.GetAccountDataBorshInto(ctx, client, token.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
	if err != nil {
		return nil, nil, err
	}

	lookupTableEntries, err := common.GetAddressLookupTable(ctx, client, token.PoolLookupTable)
	if err != nil {
		return nil, nil, err
	}

	writableBytes := append(tokenAdminRegistry.WritableIndexes[0].Bytes(), tokenAdminRegistry.WritableIndexes[1].Bytes()...)
	writableBits := ""
	for _, b := range writableBytes {
		writableBits += fmt.Sprintf("%08b", b)
	}

	lookupTableMeta := []*solana.AccountMeta{}
	for i := range lookupTableEntries {
		meta := solana.Meta(lookupTableEntries[i])

		if string(writableBits[i]) == "1" {
			meta = meta.WRITE()
		}
		lookupTableMeta = append(lookupTableMeta, meta)
	}

	list := []*solana.AccountMeta{
		solana.Meta(userTokenAccount).WRITE(),
		solana.Meta(tokenBillingConfig),
		solana.Meta(poolChainConfig).WRITE(),
	}
	list = append(list, lookupTableMeta...)

	addressTables := make(map[solana.PublicKey]solana.PublicKeySlice)
	addressTables[token.PoolLookupTable] = lookupTableEntries

	return list, addressTables, nil
}
