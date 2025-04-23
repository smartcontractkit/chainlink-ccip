package tokens

import (
	"context"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

type TokenPool struct {
	// token details
	Program        solana.PublicKey
	Mint           solana.PublicKey
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

	// CCIP CPI signers
	RouterSigner  solana.PublicKey
	OfframpSigner solana.PublicKey
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
		tp.Mint,             // 7 - writable
		tp.FeeTokenConfig,   // 8
		tp.RouterSigner,     // 9
	}
	return append(list, tp.AdditionalAccounts...)
}

// NewTokenPool returns token + pool addresses. however, the token still needs to be deployed
func NewTokenPool(tokenProgram solana.PublicKey, poolProgram solana.PublicKey, mint solana.PublicKey) (TokenPool, error) {
	tokenAdminRegistryPDA, _, err := state.FindTokenAdminRegistryPDA(mint, config.CcipRouterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	// preload with defined config.EvmChainSelector
	evmChainPDA, _, err := TokenPoolChainConfigPDA(config.EvmChainSelector, mint, poolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	svmChainPDA, _, err := TokenPoolChainConfigPDA(config.SvmChainSelector, mint, poolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	evmBillingPDA, _, err := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, mint, config.FeeQuoterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	svmBillingPDA, _, err := state.FindFqPerChainPerTokenConfigPDA(config.SvmChainSelector, mint, config.FeeQuoterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	tokenConfigPda, _, err := state.FindFqBillingTokenConfigPDA(mint, config.FeeQuoterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	routerSignerPDA, _, err := state.FindExternalTokenPoolsSignerPDA(poolProgram, config.CcipRouterProgram)
	if err != nil {
		return TokenPool{}, err
	}
	offrampSignerPDA, _, err := state.FindExternalTokenPoolsSignerPDA(poolProgram, config.CcipOfframpProgram)
	if err != nil {
		return TokenPool{}, err
	}

	p := TokenPool{
		Program:          tokenProgram,
		Mint:             mint,
		FeeTokenConfig:   tokenConfigPda,
		AdminRegistryPDA: tokenAdminRegistryPDA,
		PoolProgram:      poolProgram,
		PoolLookupTable:  solana.PublicKey{},
		WritableIndexes:  []uint8{3, 4, 7}, // see ToTokenPoolEntries for writable indexes
		User:             map[solana.PublicKey]solana.PublicKey{},
		Chain:            map[uint64]solana.PublicKey{},
		Billing:          map[uint64]solana.PublicKey{},
		RouterSigner:     routerSignerPDA,
		OfframpSigner:    offrampSignerPDA,
	}
	p.Chain[config.EvmChainSelector] = evmChainPDA
	p.Chain[config.SvmChainSelector] = svmChainPDA
	p.Billing[config.EvmChainSelector] = evmBillingPDA
	p.Billing[config.SvmChainSelector] = svmBillingPDA
	p.PoolConfig, err = TokenPoolConfigAddress(p.Mint, poolProgram)
	if err != nil {
		return TokenPool{}, err
	}
	p.PoolSigner, err = TokenPoolSignerAddress(p.Mint, poolProgram)
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
	Mint          solana.PublicKey
}

type EventMintRelease struct {
	Discriminator [8]byte
	Sender        solana.PublicKey
	Recipient     solana.PublicKey
	Amount        uint64
	Mint          solana.PublicKey
}

type EventChainConfigured struct {
	Discriminator         [8]byte
	ChainSelector         uint64
	Token                 test_token_pool.RemoteAddress
	PreviousToken         test_token_pool.RemoteAddress
	PoolAddresses         []test_token_pool.RemoteAddress
	PreviousPoolAddresses []test_token_pool.RemoteAddress
	Mint                  solana.PublicKey
}

type EventRemotePoolsAppended struct {
	Discriminator         [8]byte
	ChainSelector         uint64
	PoolAddresses         []test_token_pool.RemoteAddress
	PreviousPoolAddresses []test_token_pool.RemoteAddress
	Mint                  solana.PublicKey
}

type EventRateLimitConfigured struct {
	Discriminator     [8]byte
	ChainSelector     uint64
	OutboundRateLimit test_token_pool.RateLimitConfig
	InboundRateLimit  test_token_pool.RateLimitConfig
	Mint              solana.PublicKey
}

type EventChainRemoved struct {
	Discriminator [8]byte
	ChainSelector uint64
	Mint          solana.PublicKey
}

type EventRouterUpdated struct {
	Discriminator [8]byte
	OldRouter     solana.PublicKey
	NewRouter     solana.PublicKey
	Mint          solana.PublicKey
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
	return ParseTokenLookupTableWithChain(ctx, client, token, userTokenAccount, config.EvmChainSelector)
}

func ParseTokenLookupTableWithChain(ctx context.Context, client *rpc.Client, token TokenPool, userTokenAccount solana.PublicKey, chainSelector uint64) (solana.AccountMetaSlice, map[solana.PublicKey]solana.PublicKeySlice, error) {
	tokenBillingConfig := token.Billing[chainSelector]
	poolChainConfig := token.Chain[chainSelector]

	tokenAdminRegistry := ccip_common.TokenAdminRegistry{}
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
