package mcms

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
)

func TestMcmMerkle(t *testing.T) {
	multisig := config.McmProgram
	to := config.ExternalCpiStubProgram

	op := &McmOpNode{
		Nonce:    1,
		Data:     []byte("test data"),
		Multisig: multisig,
		To:       to,
		RemainingAccounts: []*solana.AccountMeta{
			{PublicKey: solana.NewWallet().PublicKey(), IsSigner: true, IsWritable: false},
		},
	}

	t.Run("Buffers", func(t *testing.T) {
		buffers := op.Buffers()
		assert.Equal(t, 9, len(buffers), "Expected 9 buffers")
	})

	t.Run("Hash", func(t *testing.T) {
		assert.Equal(t, 32, len(op.Hash()), "Expected hash length of 32 bytes")
	})

	t.Run("Leaves", func(t *testing.T) {
		leaves := op.Leaves()
		assert.Equal(t, 1, len(leaves), "Expected 1 leaf")
		assert.Equal(t, op, leaves[0], "Expected leaf to be the op itself")
	})
}

func TestMcmRootMetadata(t *testing.T) {
	multisig := config.McmProgram

	metadata := &RootMetadataNode{
		PreOpCount:           10,
		PostOpCount:          15,
		Multisig:             multisig,
		OverridePreviousRoot: true,
	}

	t.Run("Buffers", func(t *testing.T) {
		buffers := metadata.Buffers()
		assert.Equal(t, 6, len(buffers), "Expected 6 buffers")
	})

	t.Run("Hash", func(t *testing.T) {
		assert.Equal(t, 32, len(metadata.Hash()), "Expected hash length of 32 bytes")
	})
}

func TestMcmOpMerkleTree(t *testing.T) {
	multisig := config.McmProgram
	to := config.ExternalCpiStubProgram

	op1 := &McmOpNode{Nonce: 1, Data: []byte("data1"), Multisig: multisig, To: to}
	op2 := &McmOpNode{Nonce: 2, Data: []byte("data2"), Multisig: multisig, To: to}
	op3 := &McmOpNode{Nonce: 3, Data: []byte("data3"), Multisig: multisig, To: to}

	t.Run("NewOpMerkleTree", func(t *testing.T) {
		tree, err := NewOpMerkleTree([]MerkleNode{op1, op2, op3})
		require.NoError(t, err)
		assert.Equal(t, 3, tree.Size(), "Expected tree size of 3")
	})

	t.Run("EthMsgHash", func(t *testing.T) {
		tree, err := NewOpMerkleTree([]MerkleNode{op1, op2})
		require.NoError(t, err)
		hash := tree.(*OpMerkleTree).EthMsgHash(1234)
		assert.Equal(t, 32, len(hash), "Expected hash length of 32 bytes")
	})

	t.Run("Proofs", func(t *testing.T) {
		_, err := NewOpMerkleTree([]MerkleNode{op1, op2, op3})
		require.NoError(t, err)
		proofs, err := op1.Proofs()
		require.NoError(t, err)
		assert.Equal(t, 2, len(proofs), "Expected 2 proof hashes for a tree with 3 nodes")
	})
}

func TestMcmHelperFunctions(t *testing.T) {
	t.Run("numToU64LePaddedEncoding", func(t *testing.T) {
		result := numToU64LePaddedEncoding(42)
		assert.Equal(t, 32, len(result), "Expected 32-byte result")
		assert.Equal(t, byte(42), result[24], "Expected 42 in the correct position")
	})

	t.Run("boolToPaddedEncoding", func(t *testing.T) {
		trueResult := boolToPaddedEncoding(true)
		falseResult := boolToPaddedEncoding(false)
		assert.Equal(t, byte(1), trueResult[31], "Expected 1 for true")
		assert.Equal(t, byte(0), falseResult[31], "Expected 0 for false")
	})

	t.Run("serializeAccountMeta", func(t *testing.T) {
		account := solana.AccountMeta{
			PublicKey:  solana.NewWallet().PublicKey(),
			IsSigner:   true,
			IsWritable: false,
		}
		result := serializeAccountMeta(&account)
		assert.Equal(t, 33, len(result), "Expected 33-byte result")
		assert.Equal(t, byte(0b10), result[32], "Expected correct flags")
	})

	t.Run("CalculateHash", func(t *testing.T) {
		buffers := [][]byte{[]byte("test1"), []byte("test2")}
		assert.Equal(t, 32, len(CalculateHash(buffers)), "Expected 32-byte hash")
	})
}

func TestMcmMerkleNodeInterface(t *testing.T) {
	multisig := config.McmProgram
	to := config.ExternalCpiStubProgram

	op := &McmOpNode{Nonce: 1, Data: []byte("test"), Multisig: multisig, To: to}
	metadata := &RootMetadataNode{PreOpCount: 1, PostOpCount: 2, Multisig: multisig}

	nodes := []MerkleNode{op, metadata}

	for _, node := range nodes {
		t.Run("MerkleNode Interface", func(t *testing.T) {
			assert.NotNil(t, node.Hash(), "Hash should not be nil")
			assert.NotEmpty(t, node.Buffers(), "Buffers should not be empty")
			assert.Greater(t, node.Size(), 0, "Size should be greater than 0")
			assert.NotNil(t, node.Leaves(), "Leaves should not be nil")

			proofs, err := node.Proofs()
			assert.Error(t, err, "Proofs should return an error for standalone nodes")
			assert.Nil(t, proofs, "Proofs should be nil for standalone nodes")
		})
	}

	// Test proofs for nodes in a tree
	tree, err := NewOpMerkleTree(nodes)
	require.NoError(t, err, "Failed to create tree")

	for _, node := range tree.Leaves() {
		t.Run("MerkleNode in Tree", func(t *testing.T) {
			proofs, err := node.Proofs()
			assert.NoError(t, err, "Proofs should not return an error for nodes in a tree")
			assert.NotEmpty(t, proofs, "Proofs should not be empty for nodes in a tree")
		})
	}
}

func TestMcmMerkleTreeOrdering(t *testing.T) {
	multisig := config.McmProgram
	to := config.ExternalCpiStubProgram

	op1 := &McmOpNode{Nonce: 1, Data: []byte("data1"), Multisig: multisig, To: to}
	op2 := &McmOpNode{Nonce: 2, Data: []byte("data2"), Multisig: multisig, To: to}

	tree, err := NewOpMerkleTree([]MerkleNode{op1, op2})
	require.NoError(t, err)

	opTree, ok := tree.(*OpMerkleTree)
	require.True(t, ok, "Expected OpMerkleTree type")

	assert.True(t, Lt(opTree.Left, opTree.Right), "Left node should be less than right node")
	assert.False(t, Eq(opTree.Left, opTree.Right), "Left and right nodes should not be equal")
}
