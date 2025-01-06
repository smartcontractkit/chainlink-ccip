package mcms

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

// Note: mcms tree reference can be found in https://github.com/smartcontractkit/mcms/blob/main/internal/core/merkle/tree.go

type MerkleNode interface {
	Hash() [32]byte
	Buffers() [][]byte
	Print(padding int)
	Size() int
	Parent() *OpMerkleTree
	SetParent(t *OpMerkleTree)
	Proofs() ([][32]byte, error)
	Leaves() []MerkleNode
}

type BaseNode struct {
	parent *OpMerkleTree
	size   int
}

func (b *BaseNode) SetParent(p *OpMerkleTree) { b.parent = p }

func (b *BaseNode) Parent() *OpMerkleTree {
	return b.parent
}

func (b *BaseNode) Size() int {
	return 1
}

type McmOpNode struct {
	BaseNode
	Nonce             uint64
	Data              []byte
	Multisig          solana.PublicKey // this is config PDA
	To                solana.PublicKey
	RemainingAccounts []*solana.AccountMeta
}

func (t *McmOpNode) Buffers() [][]byte {
	domainSeparatorHashBytes := eth.Keccak256([]byte("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP_SOLANA"))
	buffers := [][]byte{
		domainSeparatorHashBytes[:],
		config.TestChainIDPaddedBuffer[:],
		t.Multisig.Bytes(),
		numToU64LePaddedEncoding(t.Nonce),
		t.To.Bytes(),
		numToU64LePaddedEncoding(uint64(len(t.Data))),
		t.Data,
		numToU64LePaddedEncoding(uint64(len(t.RemainingAccounts))),
	}

	for _, account := range t.RemainingAccounts {
		buffers = append(buffers, serializeAccountMeta(account))
	}

	return buffers
}

func (t *McmOpNode) Hash() [32]byte {
	return CalculateHash(t.Buffers())
}

func (t *McmOpNode) Leaves() []MerkleNode {
	return []MerkleNode{t}
}

func (t *McmOpNode) Proofs() ([][32]byte, error) {
	if t.Parent() == nil {
		return nil, fmt.Errorf("cannot generate proof: TestOp is not part of a tree")
	}
	return GenerateProof(t)
}

func (t *McmOpNode) Print(padding int) {
	for i := 0; i < padding; i++ {
		fmt.Print("| ")
	}
	h := fmt.Sprintf("%x", t.Hash())
	parentHash := "ROOT"
	if t.parent != nil {
		parentHash = fmt.Sprintf("%x", t.parent.Hash())
	}
	fmt.Printf("-> Op #%d %s - Child of %s\n", t.Nonce, h, parentHash)
}

type RootMetadataNode struct {
	BaseNode
	PreOpCount           uint64
	PostOpCount          uint64
	Multisig             solana.PublicKey
	OverridePreviousRoot bool
}

func (rm *RootMetadataNode) Buffers() [][]byte {
	domainSeparatorHashBytes := eth.Keccak256([]byte("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA_SOLANA"))
	return [][]byte{
		domainSeparatorHashBytes[:],
		config.TestChainIDPaddedBuffer[:],
		rm.Multisig.Bytes(),
		numToU64LePaddedEncoding(rm.PreOpCount),
		numToU64LePaddedEncoding(rm.PostOpCount),
		boolToPaddedEncoding(rm.OverridePreviousRoot),
	}
}

func (rm *RootMetadataNode) Hash() [32]byte {
	return CalculateHash(rm.Buffers())
}

func (rm *RootMetadataNode) Leaves() []MerkleNode {
	return []MerkleNode{rm}
}

func (rm *RootMetadataNode) Proofs() ([][32]byte, error) {
	if rm.Parent() == nil {
		return nil, fmt.Errorf("cannot generate proof: RootMetadata is not part of a tree")
	}
	return GenerateProof(rm)
}

func (rm *RootMetadataNode) Print(padding int) {
	for i := 0; i < padding; i++ {
		fmt.Print("| ")
	}
	h := fmt.Sprintf("%x", rm.Hash())
	parentHash := "ROOT"
	if rm.parent != nil {
		parentHash = fmt.Sprintf("%x", rm.parent.Hash())
	}
	fmt.Printf("-> Metadata %s - Child of %s\n", h, parentHash)
}

type OpMerkleTree struct {
	BaseNode
	Left  MerkleNode
	Right MerkleNode
}

func (o *OpMerkleTree) Size() int {
	if o.Left == nil && o.Right == nil {
		return 1
	}
	return o.Left.Size() + o.Right.Size()
}

func (o *OpMerkleTree) Buffers() [][]byte {
	if o.Left == nil && o.Right == nil {
		fmt.Println("Warning: OpMerkleTree: both Left and Right are nil")
		return [][]byte{}
	}
	var leftHash, rightHash [32]byte
	if o.Left != nil {
		leftHash = o.Left.Hash()
	}
	if o.Right != nil {
		rightHash = o.Right.Hash()
	}
	return [][]byte{leftHash[:], rightHash[:]}
}

func (o *OpMerkleTree) Hash() [32]byte {
	return CalculateHash(o.Buffers())
}

func (o *OpMerkleTree) EthMsgHash(validUntil uint32) []byte {
	hash := o.Hash()
	hashedEncodedParams := eth.Keccak256(append(hash[:], numToU64BePaddedEncoding(uint64(validUntil))...))
	return eth.Keccak256(append([]byte("\x19Ethereum Signed Message:\n32"), hashedEncodedParams[:]...))
}

func (o *OpMerkleTree) Leaves() []MerkleNode {
	if o == nil {
		return []MerkleNode{}
	}
	if o.Left == nil && o.Right == nil {
		return []MerkleNode{o}
	}
	var leaves []MerkleNode
	if o.Left != nil {
		leaves = append(leaves, o.Left.Leaves()...)
	}
	if o.Right != nil {
		leaves = append(leaves, o.Right.Leaves()...)
	}
	return leaves
}

func (o *OpMerkleTree) Proofs() ([][32]byte, error) {
	var proofs [][32]byte
	for current := o; current.Parent() != nil; current = current.Parent() {
		sibling, err := current.Parent().FindSiblingOf(current)
		if err != nil {
			return nil, fmt.Errorf("failed to find sibling: %w", err)
		}
		proofs = append(proofs, sibling.Hash())
	}
	return proofs, nil
}

func (o *OpMerkleTree) FindSiblingOf(n MerkleNode) (MerkleNode, error) {
	if Eq(o.Left, n) {
		return o.Right, nil
	}
	if Eq(o.Right, n) {
		return o.Left, nil
	}
	return nil, fmt.Errorf("node is not a direct child of this tree")
}

func (o *OpMerkleTree) Print(padding int) {
	for i := 0; i < padding; i++ {
		fmt.Print("| ")
	}
	h := fmt.Sprintf("%x", o.Hash())
	parentHash := "ROOT"
	if o.parent != nil {
		parentHash = fmt.Sprintf("%x", o.parent.Hash())
	}
	fmt.Printf("Node %d %s - Child of %s\n", o.size, h, parentHash)
	if o.Left != nil {
		o.Left.Print(padding + 1)
	}
	if o.Right != nil {
		o.Right.Print(padding + 1)
	}
}

func NewOpMerkleTree(nodes []MerkleNode) (MerkleNode, error) {
	switch len(nodes) {
	case 0:
		return nil, fmt.Errorf("cannot create tree with no nodes")
	case 1:
		return nodes[0], nil
	default:
		pivot := (len(nodes) + 1) / 2

		left, err := NewOpMerkleTree(nodes[:pivot])
		if err != nil {
			return nil, fmt.Errorf("failed to create left subtree: %w", err)
		}

		right, err := NewOpMerkleTree(nodes[pivot:])
		if err != nil {
			return nil, fmt.Errorf("failed to create right subtree: %w", err)
		}

		tree := &OpMerkleTree{
			Left:  left,
			Right: right,
		}

		if !Lt(left, right) {
			tree.Left, tree.Right = right, left
		}

		tree.Left.SetParent(tree)
		tree.Right.SetParent(tree)

		return tree, nil
	}
}

func numToU64LePaddedEncoding(n uint64) []byte {
	b := make([]byte, 32)
	binary.LittleEndian.PutUint64(b[24:], n)
	return b
}

func numToU64BePaddedEncoding(n uint64) []byte {
	b := make([]byte, 32)
	binary.BigEndian.PutUint64(b[24:], n)
	return b
}

func boolToPaddedEncoding(b bool) []byte {
	result := make([]byte, 32)
	if b {
		result[31] = 1
	}
	return result
}

func serializeAccountMeta(a *solana.AccountMeta) []byte {
	var flags byte
	if a.IsSigner {
		flags |= 0b10
	}
	if a.IsWritable {
		flags |= 0b01
	}
	result := append(a.PublicKey.Bytes(), flags)
	return result
}

func Lt(a, b MerkleNode) bool {
	hashA := a.Hash()
	hashB := b.Hash()
	return bytes.Compare(hashA[:], hashB[:]) < 0
}

func Eq(a, b MerkleNode) bool {
	hashA := a.Hash()
	hashB := b.Hash()
	return bytes.Equal(hashA[:], hashB[:])
}

func CalculateHash(buffers [][]byte) [32]byte {
	hash := eth.Keccak256(bytes.Join(buffers, nil))
	var hash32 [32]byte
	copy(hash32[:], hash)
	return hash32
}

func GenerateProof(node MerkleNode) ([][32]byte, error) {
	var proofs [][32]byte
	for current := node; current.Parent() != nil; current = current.Parent() {
		sibling, err := current.Parent().FindSiblingOf(current)
		if err != nil {
			return nil, fmt.Errorf("failed to find sibling: %w", err)
		}
		proofs = append(proofs, sibling.Hash())
	}
	return proofs, nil
}

func ConvertProof(proof [][]byte) ([][32]uint8, error) {
	fixedProof := make([][32]uint8, len(proof))
	for i, p := range proof {
		if len(p) != 32 {
			return nil, fmt.Errorf("proof element %d is not 32 bytes long", i)
		}
		copy(fixedProof[i][:], p)
	}
	return fixedProof, nil
}

// dump function for debugging
func DumpOpDetails(op *McmOpNode) string {
	var sb strings.Builder

	sb.WriteString("Operation Details:\n")
	sb.WriteString("=================\n\n")

	// Print basic info
	sb.WriteString(fmt.Sprintf("Nonce: %d\n", op.Nonce))
	sb.WriteString(fmt.Sprintf("Multisig: %s\n", op.Multisig.String()))
	sb.WriteString(fmt.Sprintf("To: %s\n", op.To.String()))

	// Print raw buffers
	sb.WriteString("Raw Buffers:\n")
	buffers := op.Buffers()
	for i, buf := range buffers {
		sb.WriteString(fmt.Sprintf("Buffer[%d]: %s\n", i, hex.EncodeToString(buf)))
	}
	sb.WriteString("\n")

	// Print execution data
	sb.WriteString(fmt.Sprintf("Execution Data: %s\n\n", hex.EncodeToString(op.Data)))

	// Print hash
	hash := op.Hash()
	sb.WriteString(fmt.Sprintf("Leaf Hash: %s\n", hex.EncodeToString(hash[:])))

	// Print Merkle proofs if available
	if op.Parent() != nil {
		proofs, err := op.Proofs()
		if err == nil {
			sb.WriteString("\nMerkle Proofs:\n")
			for i, proof := range proofs {
				sb.WriteString(fmt.Sprintf("Proof[%d]: %s\n", i, hex.EncodeToString(proof[:])))
			}
		}
	}

	// Print remaining accounts if any
	if len(op.RemainingAccounts) > 0 {
		sb.WriteString("\nRemaining Accounts:\n")
		for i, acc := range op.RemainingAccounts {
			sb.WriteString(fmt.Sprintf("Account[%d]:\n", i))
			sb.WriteString(fmt.Sprintf("  Address: %s\n", acc.PublicKey.String()))
			sb.WriteString(fmt.Sprintf("  Is Signer: %v\n", acc.IsSigner))
			sb.WriteString(fmt.Sprintf("  Is Writable: %v\n", acc.IsWritable))
			// Print serialized form for verification
			sb.WriteString(fmt.Sprintf("  Serialized: %s\n",
				hex.EncodeToString(serializeAccountMeta(acc))))
		}
	}

	return sb.String()
}
