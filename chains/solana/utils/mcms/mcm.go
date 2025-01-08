package mcms

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

// mcm signer dataless pda
func GetSignerPDA(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_signer"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func GetConfigPDA(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_config"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func GetConfigSignersPDA(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("multisig_config_signers"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func GetRootMetadataPDA(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("root_metadata"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

func GetExpiringRootAndOpCountPDA(msigName [32]byte) solana.PublicKey {
	pda, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("expiring_root_and_op_count"),
		msigName[:],
	}, config.McmProgram)
	return pda
}

// get address of the root_signatures pda
func GetRootSignaturesPDA(msigName [32]byte, root [32]byte, validUntil uint32) solana.PublicKey {
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
func GetSeenSignedHashesPDA(msigName [32]byte, root [32]byte, validUntil uint32) solana.PublicKey {
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

// utils for padding mcm id
func PadString32(input string) ([32]byte, error) {
	var result [32]byte
	inputBytes := []byte(input)
	inputLen := len(inputBytes)
	if inputLen > 32 {
		return result, errors.New("input string exceeds 32 bytes")
	}
	startPos := 32 - inputLen
	copy(result[startPos:], inputBytes)
	return result, nil
}

func UnpadString32(input [32]byte) string {
	startPos := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			startPos = i
			break
		}
	}
	return string(input[startPos:])
}

type McmConfigArgs struct {
	MultisigName    [32]uint8
	SignerAddresses [][20]uint8
	SignerGroups    []byte
	GroupQuorums    [32]uint8
	GroupParents    [32]uint8
	ClearRoot       bool
}

func NewValidMcmConfig(msigName [32]byte, signerPrivateKeys []string, signerGroups []byte, quorums []uint8, parents []uint8, clearRoot bool) (*McmConfigArgs, error) {
	if len(signerGroups) == 0 {
		return nil, fmt.Errorf("signerGroups cannot be empty")
	}

	signers, err := eth.GetEvmSigners(signerPrivateKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to get test EVM signers: %w", err)
	}

	if len(signers) != len(signerGroups) {
		return nil, fmt.Errorf("number of signers (%d) does not match length of signerGroups (%d)", len(signers), len(signerGroups))
	}

	signerAddresses := make([][20]uint8, len(signers))
	for i, signer := range signers {
		signerAddresses[i] = signer.Address
	}

	var groupQuorums [32]uint8
	var groupParents [32]uint8

	copy(groupQuorums[:], quorums)
	copy(groupParents[:], parents)

	// Create new config vars to ensure atomic test configs
	newSignerAddresses := make([][20]uint8, len(signerAddresses))
	copy(newSignerAddresses, signerAddresses)

	newSignerGroups := make([]byte, len(signerGroups))
	copy(newSignerGroups, signerGroups)

	newGroupQuorums := groupQuorums
	newGroupParents := groupParents
	newClearRoot := clearRoot

	config := &McmConfigArgs{
		MultisigName: msigName,
	}
	config.SignerAddresses = newSignerAddresses
	config.SignerGroups = newSignerGroups
	config.GroupQuorums = newGroupQuorums
	config.GroupParents = newGroupParents
	config.ClearRoot = newClearRoot
	return config, nil
}

func GetNewMcmMultisig(name [32]byte) Multisig {
	return Multisig{
		PaddedName:                name,
		SignerPDA:                 GetSignerPDA(name),
		ConfigPDA:                 GetConfigPDA(name),
		RootMetadataPDA:           GetRootMetadataPDA(name),
		ExpiringRootAndOpCountPDA: GetExpiringRootAndOpCountPDA(name),
		ConfigSignersPDA:          GetConfigSignersPDA(name),
		RootSignaturesPDA: func(root [32]byte, validUntil uint32) solana.PublicKey {
			return GetRootSignaturesPDA(name, root, validUntil)
		},
		SeenSignedHashesPDA: func(root [32]byte, validUntil uint32) solana.PublicKey {
			return GetSeenSignedHashesPDA(name, root, validUntil)
		},
	}
}

// instructions builder for preloading signers
func GetPreloadSignersIxs(signerAddresses [][20]uint8, msigName [32]byte, multisigCfgPDA solana.PublicKey, cfgSignersPDA solana.PublicKey, authority solana.PublicKey, appendChunkSize int) ([]solana.Instruction, error) {
	ixs := make([]solana.Instruction, 0)

	initSignersIx, isErr := mcm.NewInitSignersInstruction(
		msigName,
		//nolint:gosec
		uint8(len(signerAddresses)),
		multisigCfgPDA,
		cfgSignersPDA,
		authority,
		solana.SystemProgramID,
	).ValidateAndBuild()
	if isErr != nil {
		return nil, isErr
	}
	ixs = append(ixs, initSignersIx)

	appendSignersIxs, asErr := GetAppendSignersIxs(signerAddresses, msigName, multisigCfgPDA, cfgSignersPDA, authority, appendChunkSize)
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
func GetAppendSignersIxs(signerAddresses [][20]uint8, msigName [32]byte, multisigCfgPDA solana.PublicKey, cfgSignersPDA solana.PublicKey, authority solana.PublicKey, chunkSize int) ([]solana.Instruction, error) {
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
func GetMcmPreloadSignaturesIxs(signatures []mcm.Signature, msigName [32]byte, root [32]uint8, validUntil uint32, signaturesPDA solana.PublicKey, authority solana.PublicKey, appendChunkSize int) ([]solana.Instruction, error) {
	ixs := make([]solana.Instruction, 0)

	initSigsIx, isErr := mcm.NewInitSignaturesInstruction(
		msigName,
		root,
		validUntil,
		//nolint:gosec
		uint8(len(signatures)),
		signaturesPDA,
		authority,
		solana.SystemProgramID,
	).ValidateAndBuild()
	if isErr != nil {
		return nil, isErr
	}
	ixs = append(ixs, initSigsIx)

	appendSigsIxs, asErr := GetAppendSignaturesIxs(signatures, msigName, root, validUntil, signaturesPDA, authority, appendChunkSize)
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
func GetAppendSignaturesIxs(signatures []mcm.Signature, msigName [32]byte, root [32]uint8, validUntil uint32, signaturesPDA solana.PublicKey, authority solana.PublicKey, chunkSize int) ([]solana.Instruction, error) {
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

// simple salt generator that uses the current Unix timestamp(in mills)
func SimpleSalt() ([32]byte, error) {
	var salt [32]byte
	now := time.Now().UnixMilli()
	if now < 0 {
		return salt, fmt.Errorf("negative timestamp: %d", now)
	}
	// unix timestamp in millseconds
	binary.BigEndian.PutUint64(salt[:8], uint64(now))
	// Next 8 bytes: Crypto random
	randBytes := make([]byte, 8)
	if _, err := crypto_rand.Read(randBytes); err != nil {
		return salt, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	copy(salt[8:16], randBytes)
	return salt, nil
}
