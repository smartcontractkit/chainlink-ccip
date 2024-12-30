package mcms

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

// mcm signer dataless pda
func McmSignerAddress(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_signer"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func McmConfigAddress(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_config"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func McmConfigSignersAddress(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_config_signers"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func RootMetadataAddress(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("root_metadata"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func ExpiringRootAndOpCountAddress(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("expiring_root_and_op_count"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

// get address of the root_signatures pda
func RootSignaturesAddress(msigName [32]byte, root [32]byte, validUntil uint32) solana.PublicKey {
	validUntilBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(validUntilBytes, validUntil)

	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("root_signatures"),
		msigName[:],
		root[:],
		validUntilBytes,
	}, config.McmProgram)
	return pda
}

// get address of the seen_signed_hashes pda
func SeenSignedHashesAddress(msigName [32]byte, root [32]byte, validUntil uint32) solana.PublicKey {
	validUntilBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(validUntilBytes, validUntil)
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("seen_signed_hashes"),
		msigName[:],
		root[:],
		validUntilBytes,
	}, config.McmProgram)
	return pda
}

func NewMcmMultisig(name [32]byte) Multisig {
	return Multisig{
		PaddedName:                name,
		SignerPDA:                 McmSignerAddress(name),
		ConfigPDA:                 McmConfigAddress(name),
		RootMetadataPDA:           RootMetadataAddress(name),
		ExpiringRootAndOpCountPDA: ExpiringRootAndOpCountAddress(name),
		ConfigSignersPDA:          McmConfigSignersAddress(name),
		RootSignaturesPDA: func(root [32]byte, validUntil uint32) solana.PublicKey {
			return RootSignaturesAddress(name, root, validUntil)
		},
		SeenSignedHashesPDA: func(root [32]byte, validUntil uint32) solana.PublicKey {
			return SeenSignedHashesAddress(name, root, validUntil)
		},
	}
}

// instructions builder for preloading signers
func McmPreloadSignersIxs(signerAddresses [][20]uint8, msigName [32]byte, multisigCfgPDA solana.PublicKey, cfgSignersPDA solana.PublicKey, authority solana.PublicKey, appendChunkSize int) ([]solana.Instruction, error) {
	ixs := make([]solana.Instruction, 0)

	parsedTotalSigners, pErr := SafeToUint8(len(signerAddresses))
	if pErr != nil {
		return nil, pErr
	}
	initSignersIx, isErr := mcm.NewInitSignersInstruction(
		msigName,
		parsedTotalSigners,
		multisigCfgPDA,
		cfgSignersPDA,
		authority,
		solana.SystemProgramID,
	).ValidateAndBuild()
	if isErr != nil {
		return nil, isErr
	}
	ixs = append(ixs, initSignersIx)

	appendSignersIxs, asErr := AppendSignersIxs(signerAddresses, msigName, multisigCfgPDA, cfgSignersPDA, authority, appendChunkSize)
	if asErr != nil {
		return nil, asErr
	}
	ixs = append(ixs, appendSignersIxs...)

	finalizeSignersIx, fsErr := mcm.NewFinalizeSignersInstruction(
		msigName,
		multisigCfgPDA,
		cfgSignersPDA,
		authority,
	).ValidateAndBuild()
	if fsErr != nil {
		return nil, fsErr
	}
	ixs = append(ixs, finalizeSignersIx)

	return ixs, nil
}

// get chunked append instructions to preload signers to pda, required before set_config
func AppendSignersIxs(signerAddresses [][20]uint8, msigName [32]byte, multisigCfgPDA solana.PublicKey, cfgSignersPDA solana.PublicKey, authority solana.PublicKey, chunkSize int) ([]solana.Instruction, error) {
	if chunkSize > config.MaxAppendSignerBatchSize {
		return nil, errors.New("chunkSize exceeds max signers chunk size")
	}
	ixs := make([]solana.Instruction, 0)
	for i := 0; i < len(signerAddresses); i += chunkSize {
		end := i + chunkSize
		if end > len(signerAddresses) {
			end = len(signerAddresses)
		}
		appendIx, appendErr := mcm.NewAppendSignersInstruction(
			msigName,
			signerAddresses[i:end],
			multisigCfgPDA,
			cfgSignersPDA,
			authority,
		).ValidateAndBuild()
		if appendErr != nil {
			return nil, appendErr
		}
		ixs = append(ixs, appendIx)
	}
	return ixs, nil
}

// instructions builder for preloading signatures
func McmPreloadSignaturesIxs(signatures []mcm.Signature, msigName [32]byte, root [32]uint8, validUntil uint32, signaturesPDA solana.PublicKey, authority solana.PublicKey, appendChunkSize int) ([]solana.Instruction, error) {
	ixs := make([]solana.Instruction, 0)

	parsedTotalSigs, pErr := SafeToUint8(len(signatures))
	if pErr != nil {
		return nil, pErr
	}

	initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
		msigName,
		root,
		validUntil,
		parsedTotalSigs,
		signaturesPDA,
		authority,
		solana.SystemProgramID,
	).ValidateAndBuild()
	if isErr != nil {
		return nil, isErr
	}
	ixs = append(ixs, initSigsIx)

	appendSigsIxs, asErr := AppendSignaturesIxs(signatures, msigName, root, validUntil, signaturesPDA, authority, appendChunkSize)
	if asErr != nil {
		return nil, asErr
	}

	ixs = append(ixs, appendSigsIxs...)

	finalizeSigsIx, fsErr := mcm.NewFinalizeSignaturesInstruction(
		msigName,
		root,
		validUntil,
		signaturesPDA,
		authority,
	).ValidateAndBuild()
	if fsErr != nil {
		return nil, fsErr
	}
	ixs = append(ixs, finalizeSigsIx)

	return ixs, nil
}

// get chunked append instructions to preload signatures to pda, required before set_root
func AppendSignaturesIxs(signatures []mcm.Signature, msigName [32]byte, root [32]uint8, validUntil uint32, signaturesPDA solana.PublicKey, authority solana.PublicKey, chunkSize int) ([]solana.Instruction, error) {
	if chunkSize > config.MaxAppendSignatureBatchSize {
		return nil, errors.New("chunkSize exceeds max signatures chunk size")
	}
	ixs := make([]solana.Instruction, 0)
	for i := 0; i < len(signatures); i += chunkSize {
		end := i + chunkSize
		if end > len(signatures) {
			end = len(signatures)
		}
		appendIx, appendErr := mcm.NewAppendSignaturesInstruction(
			msigName,
			root,
			validUntil,
			signatures[i:end],
			signaturesPDA,
			authority,
		).ValidateAndBuild()
		if appendErr != nil {
			return nil, appendErr
		}
		ixs = append(ixs, appendIx)
	}
	return ixs, nil
}

type McmRootInput struct {
	Multisig             solana.PublicKey
	Operations           []McmOpNode
	PreOpCount           uint64
	PostOpCount          uint64
	ValidUntil           uint32
	OverridePreviousRoot bool
}

type McmRootData struct {
	EthMsgHash    []byte
	Root          [32]byte
	Metadata      mcm.RootMetadataInput
	MetadataProof [][32]uint8
}

func CreateMcmRootData(input McmRootInput) (McmRootData, error) {
	numOps := len(input.Operations)

	// add 1 for the root metadata node
	nodes := make([]MerkleNode, numOps+1)
	for i := range input.Operations {
		nodes[i] = &input.Operations[i]
	}

	rootMetadata := RootMetadataNode{
		Multisig:             input.Multisig,
		PreOpCount:           input.PreOpCount,
		PostOpCount:          input.PostOpCount,
		OverridePreviousRoot: input.OverridePreviousRoot,
	}
	nodes[numOps] = &rootMetadata

	// construct the tree
	tree, err := NewOpMerkleTree(nodes)
	if err != nil {
		return McmRootData{}, fmt.Errorf("failed to create tree: %w", err)
	}

	metadata := mcm.RootMetadataInput{
		ChainId:              config.TestChainID,
		Multisig:             rootMetadata.Multisig,
		PreOpCount:           rootMetadata.PreOpCount,
		PostOpCount:          rootMetadata.PostOpCount,
		OverridePreviousRoot: rootMetadata.OverridePreviousRoot,
	}

	// convert root to 32 byte array
	root := tree.Hash()

	metadataProof, err := rootMetadata.Proofs()
	if err != nil {
		return McmRootData{}, fmt.Errorf("failed to get metadata proof: %w", err)
	}

	opTree, ok := tree.(*OpMerkleTree)
	if !ok {
		return McmRootData{}, fmt.Errorf("tree is not of type *OpMerkleTree")
	}
	ethMsgHash := opTree.EthMsgHash(input.ValidUntil)

	return McmRootData{
		Root:          root,
		EthMsgHash:    ethMsgHash,
		Metadata:      metadata,
		MetadataProof: metadataProof,
	}, nil
}

func BulkSignOnMsgHash(signers []eth.Signer, ethMsgHash []byte) ([]mcm.Signature, error) {
	signatures := make([]mcm.Signature, len(signers))
	for i, signer := range signers {
		signature, err := signer.Sign(ethMsgHash)
		if err != nil {
			return nil, err
		}
		signatures[i] = signature
	}
	return signatures, nil
}

func IxToMcmTestOpNode(multisig solana.PublicKey, msigSigner solana.PublicKey, ix solana.Instruction, nonce uint64) (McmOpNode, error) {
	ixData, err := ix.Data()
	if err != nil {
		return McmOpNode{}, err
	}
	// Create the accounts slice with the correct size
	accounts := make([]*solana.AccountMeta, 0, len(ix.Accounts()))

	for _, acc := range ix.Accounts() {
		accCopy := *acc
		// NOTE: this bypasses utils.sendTransaction signing part since it's PDA and it doesn't have private key
		if accCopy.PublicKey == msigSigner {
			accCopy.IsSigner = false
		}
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  accCopy.PublicKey,
			IsSigner:   accCopy.IsSigner,
			IsWritable: accCopy.IsWritable,
		})
	}

	node := McmOpNode{
		Multisig:          multisig,
		Nonce:             nonce,
		To:                ix.ProgramID(),
		Data:              ixData,
		RemainingAccounts: accounts,
	}

	return node, nil
}