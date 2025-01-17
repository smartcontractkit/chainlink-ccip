package contracts

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func mustRandomPubkey() solana.PublicKey {
	k, err := solana.NewRandomPrivateKey()
	if err != nil {
		panic(err)
	}
	return k.PublicKey()
}

// NOTE: this test does not execute or validate transaction inputs, it simply builds transactions to calculate the size of each transaction with signers
func TestTransactionSizing(t *testing.T) {
	ccip_router.SetProgramID(config.CcipRouterProgram)

	auth, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	// mocked router lookup table for constant accounts
	// chain specific configs are not constant but are a small set relative to the number of users
	routerTable := map[string]solana.PublicKey{
		"routerConfig":          mustRandomPubkey(),
		"destChainConfig":       mustRandomPubkey(),
		"systemProgram":         solana.SystemProgramID,
		"billingTokenProgram":   solana.TokenProgramID,
		"billingTokenMint":      mustRandomPubkey(),
		"billingTokenConfig":    mustRandomPubkey(),
		"routerBillingTokenATA": mustRandomPubkey(),
		"routerBillingSigner":   mustRandomPubkey(),
		"routerTokenPoolSigner": mustRandomPubkey(),
		"sysVarInstruction":     solana.SysVarInstructionsPubkey,
		"originChainConfig":     mustRandomPubkey(),
		"arbMessagingSigner":    mustRandomPubkey(),
	}

	tokenTable := map[string]solana.PublicKey{
		"tokenAdminRegistryPDA": mustRandomPubkey(),
		"poolLookupTable":       mustRandomPubkey(),
		"poolProgram":           config.CcipTokenPoolProgram,
		"poolConfig":            mustRandomPubkey(),
		"poolTokenAccount":      mustRandomPubkey(),
		"poolSigner":            mustRandomPubkey(),
		"tokenProgram":          config.Token2022Program,
		"mint":                  mustRandomPubkey(),
	}

	run := func(name string, ix solana.Instruction, tables map[solana.PublicKey]solana.PublicKeySlice, opts ...common.TxModifier) string {
		tx, err := solana.NewTransaction([]solana.Instruction{ix}, solana.Hash{1}, solana.TransactionAddressTables(tables))
		require.NoError(t, err)

		for _, o := range opts {
			require.NoError(t, o(tx, nil))
		}

		_, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
			return &auth
		})
		require.NoError(t, err)

		bz, err := tx.MarshalBinary()
		require.NoError(t, err)
		l := len(bz)
		require.LessOrEqual(t, l, 1238)
		return fmt.Sprintf("%-55s: %-4d - remaining: %d", name, l, 1232-l)
	}

	// ccipSend test messages + instruction ---------------------------------
	sendNoTokens := ccip_router.SVM2AnyMessage{
		Receiver:     make([]byte, 20), // EVM address
		Data:         []byte{},
		TokenAmounts: []ccip_router.SVMTokenAmount{}, // no tokens
		FeeToken:     [32]byte{},                     // solana fee token
		ExtraArgs:    ccip_router.ExtraArgsInput{},   // default options
	}
	sendSingleMinimalToken := ccip_router.SVM2AnyMessage{
		Receiver: make([]byte, 20),
		Data:     []byte{},
		TokenAmounts: []ccip_router.SVMTokenAmount{ccip_router.SVMTokenAmount{
			Token:  [32]byte{},
			Amount: 0,
		}}, // one token
		FeeToken:  [32]byte{},
		ExtraArgs: ccip_router.ExtraArgsInput{}, // default options
	}
	ixCcipSend := func(msg ccip_router.SVM2AnyMessage, tokenIndexes []byte, addAccounts solana.PublicKeySlice) solana.Instruction {
		base := ccip_router.NewCcipSendInstruction(
			1,
			msg,
			tokenIndexes,
			routerTable["routerConfig"],
			routerTable["destChainConfig"],
			mustRandomPubkey(), // user nonce PDA
			auth.PublicKey(),   // sender/authority
			routerTable["systemProgram"],
			routerTable["billingTokenProgram"],
			routerTable["billingTokenMint"],
			routerTable["billingTokenConfig"],
			mustRandomPubkey(), // link billing config
			mustRandomPubkey(), // user billing token ATA
			routerTable["routerBillingTokenATA"],
			routerTable["routerBillingSigner"],
			routerTable["routerTokenPoolSigner"],
		)

		for _, v := range addAccounts {
			base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(v))
		}
		ix, err := base.ValidateAndBuild()
		require.NoError(t, err)
		return ix
	}

	// ccip commit test messages + instruction ------------------------
	commitNoPrices := ccip_router.CommitInput{
		MerkleRoot: ccip_router.MerkleRoot{
			SourceChainSelector: 0,
			OnRampAddress:       make([]byte, 20), // EVM onramp
			MinSeqNr:            0,
			MaxSeqNr:            0,
			MerkleRoot:          [32]uint8{},
		},
	}
	commitWithPrices := ccip_router.CommitInput{
		PriceUpdates: ccip_router.PriceUpdates{
			TokenPriceUpdates: make([]ccip_router.TokenPriceUpdate, 1),
			GasPriceUpdates:   make([]ccip_router.GasPriceUpdate, 1),
		},
		MerkleRoot: ccip_router.MerkleRoot{
			SourceChainSelector: 0,
			OnRampAddress:       make([]byte, 20),
			MinSeqNr:            0,
			MaxSeqNr:            0,
			MerkleRoot:          [32]uint8{},
		},
	}
	ixCommit := func(input ccip_router.CommitInput, addAccounts solana.PublicKeySlice) solana.Instruction {
		base := ccip_router.NewCommitInstruction(
			[3][32]byte{}, // report context
			input,
			make([][65]byte, 6), // f = 5, estimating f+1 signatures
			routerTable["routerConfig"],
			routerTable["originChainConfig"],
			mustRandomPubkey(), // commit report PDA
			auth.PublicKey(),
			routerTable["systemProgram"],
			routerTable["sysVarInstruction"],
		)

		for _, v := range addAccounts {
			base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(v))
		}
		ix, err := base.ValidateAndBuild()
		require.NoError(t, err)
		return ix
	}

	// ccip execute test messages + instruction -----------------------
	executeEmpty := ccip_router.ExecutionReportSingleChain{
		SourceChainSelector: 0,
		Message: ccip_router.Any2SVMRampMessage{
			Header: ccip_router.RampMessageHeader{
				MessageId:           [32]uint8{},
				SourceChainSelector: 0,
				DestChainSelector:   0,
				SequenceNumber:      0,
				Nonce:               0,
			},
			Sender:        make([]byte, 20), // EVM sender
			Data:          []byte{},
			TokenReceiver: [32]byte{},
			LogicReceiver: [32]byte{},
			TokenAmounts:  []ccip_router.Any2SVMTokenTransfer{},
			ExtraArgs: ccip_router.SVMExtraArgs{
				ComputeUnits:     0,
				IsWritableBitmap: 0,
				Accounts:         []solana.PublicKey{},
			},
		},
		OffchainTokenData: [][]byte{},
		Root:              [32]uint8{},
		Proofs:            [][32]uint8{}, // single message merkle root (added roots consume 32 bytes)
	}
	executeSingleToken := ccip_router.ExecutionReportSingleChain{
		SourceChainSelector: 0,
		Message: ccip_router.Any2SVMRampMessage{
			Header: ccip_router.RampMessageHeader{
				MessageId:           [32]uint8{},
				SourceChainSelector: 0,
				DestChainSelector:   0,
				SequenceNumber:      0,
				Nonce:               0,
			},
			Sender:        make([]byte, 20), // EVM sender
			Data:          []byte{},
			TokenReceiver: [32]byte{},
			LogicReceiver: [32]byte{},
			TokenAmounts: []ccip_router.Any2SVMTokenTransfer{{
				SourcePoolAddress: make([]byte, 20), // EVM origin token pool
				DestTokenAddress:  [32]byte{},
				DestGasAmount:     0,
				ExtraData:         []byte{},
				Amount:            [32]uint8{},
			}},
			ExtraArgs: ccip_router.SVMExtraArgs{
				ComputeUnits:     0,
				IsWritableBitmap: 0,
				Accounts:         []solana.PublicKey{},
			},
		},
		OffchainTokenData: [][]byte{},
		Root:              [32]uint8{},
		Proofs:            [][32]uint8{}, // single message merkle root (added roots consume 32 bytes)
	}

	ixExecute := func(report ccip_router.ExecutionReportSingleChain, tokenIndexes []byte, addAccounts solana.PublicKeySlice) solana.Instruction {
		base := ccip_router.NewExecuteInstruction(
			report,
			[3][32]byte{}, // report context
			tokenIndexes,
			routerTable["routerConfig"],
			routerTable["originChainConfig"],
			mustRandomPubkey(), // commit report PDA
			routerTable["arbMessagingSigner"],
			auth.PublicKey(),
			routerTable["systemProgram"],
			routerTable["sysVarInstruction"],
			routerTable["routerTokenPoolSigner"],
		)

		for _, v := range addAccounts {
			base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(v))
		}
		ix, err := base.ValidateAndBuild()
		require.NoError(t, err)
		return ix
	}

	// runner ---------------------------------------------------------
	params := []struct {
		name   string
		ix     solana.Instruction
		tables map[solana.PublicKey]solana.PublicKeySlice
	}{
		{
			"ccipSend:noToken",
			ixCcipSend(sendNoTokens, []byte{}, nil),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey(): maps.Values(routerTable),
			},
		},
		{
			"ccipSend:singleToken",
			ixCcipSend(sendSingleMinimalToken, []byte{0}, append([]solana.PublicKey{
				mustRandomPubkey(), // user ATA
				mustRandomPubkey(), // token billing config
				mustRandomPubkey(), // token pool chain config
			}, maps.Values(tokenTable)...)),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey():            maps.Values(routerTable),
				tokenTable["poolLookupTable"]: maps.Values(tokenTable),
			},
		},
		{
			"commit:noPrices",
			ixCommit(commitNoPrices, nil),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey(): maps.Values(routerTable),
			},
		},
		{
			"commit:withPrices",
			ixCommit(commitWithPrices, solana.PublicKeySlice{
				routerTable["billingTokenConfig"], // token price update
				routerTable["destChainConfig"],    // gas price update
			}),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey(): maps.Values(routerTable),
			},
		},
		{
			"execute:noToken",
			ixExecute(executeEmpty, []byte{}, nil),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey(): maps.Values(routerTable),
			},
		},
		{
			"execute:singleToken",
			ixExecute(executeSingleToken, []byte{0}, append([]solana.PublicKey{
				mustRandomPubkey(), // user ATA
				mustRandomPubkey(), // token billing config
				mustRandomPubkey(), // token pool chain config
			}, maps.Values(tokenTable)...)),
			map[solana.PublicKey]solana.PublicKeySlice{
				mustRandomPubkey():            maps.Values(routerTable),
				tokenTable["poolLookupTable"]: maps.Values(tokenTable),
			},
		},
	}

	divider := strings.Repeat("-", 78)
	outputs := []string{"TX SIZE ANALYSIS", divider}
	for _, p := range params {
		for _, l := range []string{"", " +lookupTable"} {
			var tables map[solana.PublicKey]solana.PublicKeySlice
			if strings.Contains(l, "+lookupTable") {
				tables = p.tables
			}

			outputs = append(outputs,
				run(p.name+l, p.ix, tables),
				run(p.name+l+" +cuLimit", p.ix, tables, common.AddComputeUnitLimit(0)),
				run(p.name+l+" +cuPrice", p.ix, tables, common.AddComputeUnitPrice(0)),
				run(p.name+l+" +cuPrice +cuLimit", p.ix, tables, common.AddComputeUnitLimit(0), common.AddComputeUnitPrice(0)),
				divider,
			)
		}
	}
	t.Logf("\n%s\n", strings.Join(outputs, "\n"))
}
