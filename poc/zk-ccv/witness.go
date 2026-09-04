package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
)

// Witness is the verifierResults payload checked by SuccinctZKVerifier.
type Witness struct {
	AnchorBlockNumber uint64
	// Headers are RLP block headers, anchor block first, message block last.
	Headers    [][]byte
	TxIndex    uint16
	LogIndex   uint8
	ProofNodes [][]byte
}

// Encode produces the byte layout documented in SuccinctZKVerifier._decodeWitness.
func (w *Witness) Encode(versionTag [4]byte) ([]byte, error) {
	if len(w.Headers) > 0xffff {
		return nil, fmt.Errorf("too many headers: %d", len(w.Headers))
	}
	if len(w.ProofNodes) > 0xff {
		return nil, fmt.Errorf("too many proof nodes: %d", len(w.ProofNodes))
	}
	var buf bytes.Buffer
	buf.Write(versionTag[:])
	_ = binary.Write(&buf, binary.BigEndian, w.AnchorBlockNumber)
	_ = binary.Write(&buf, binary.BigEndian, uint16(len(w.Headers)))
	for _, h := range w.Headers {
		if err := writeLengthPrefixed(&buf, h); err != nil {
			return nil, err
		}
	}
	_ = binary.Write(&buf, binary.BigEndian, w.TxIndex)
	buf.WriteByte(w.LogIndex)
	buf.WriteByte(uint8(len(w.ProofNodes)))
	for _, n := range w.ProofNodes {
		if err := writeLengthPrefixed(&buf, n); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func writeLengthPrefixed(buf *bytes.Buffer, item []byte) error {
	if len(item) > 0xffff {
		return fmt.Errorf("item too long: %d bytes", len(item))
	}
	_ = binary.Write(buf, binary.BigEndian, uint16(len(item)))
	buf.Write(item)
	return nil
}

// buildHeaderChain returns RLP headers from anchorBlock down to messageBlock and checks the hash chain locally.
func buildHeaderChain(ctx context.Context, src *ethclient.Client, anchorBlock, messageBlock uint64, anchorBlockHash common.Hash) ([][]byte, error) {
	if anchorBlock < messageBlock {
		return nil, fmt.Errorf("anchor block %d is below message block %d", anchorBlock, messageBlock)
	}
	if anchorBlock == messageBlock {
		return nil, nil
	}
	var headers [][]byte
	expected := anchorBlockHash
	for n := anchorBlock; n >= messageBlock; n-- {
		header, err := src.HeaderByNumber(ctx, new(big.Int).SetUint64(n))
		if err != nil {
			return nil, fmt.Errorf("header %d: %w", n, err)
		}
		enc, err := rlp.EncodeToBytes(header)
		if err != nil {
			return nil, err
		}
		actual := crypto.Keccak256Hash(enc)
		if actual != header.Hash() {
			return nil, fmt.Errorf("header %d: local RLP hash %s differs from node hash %s", n, actual, header.Hash())
		}
		if actual != expected {
			return nil, fmt.Errorf("header %d: hash %s does not match expected %s", n, actual, expected)
		}
		headers = append(headers, enc)
		expected = header.ParentHash
	}
	return headers, nil
}

// buildReceiptProof rebuilds the receipts trie of a block and returns the proof nodes for txIndex and the root.
func buildReceiptProof(ctx context.Context, src *ethclient.Client, blockNumber uint64, txIndex uint) ([][]byte, common.Hash, error) {
	receipts, err := src.BlockReceipts(ctx, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber)))
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("block receipts %d: %w", blockNumber, err)
	}
	if int(txIndex) >= len(receipts) {
		return nil, common.Hash{}, fmt.Errorf("tx index %d out of range, block has %d receipts", txIndex, len(receipts))
	}

	tr := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	var buf bytes.Buffer
	for i := range receipts {
		key, err := rlp.EncodeToBytes(uint64(i))
		if err != nil {
			return nil, common.Hash{}, err
		}
		buf.Reset()
		types.Receipts(receipts).EncodeIndex(i, &buf)
		if err := tr.Update(key, common.CopyBytes(buf.Bytes())); err != nil {
			return nil, common.Hash{}, err
		}
	}
	root := tr.Hash()

	key, err := rlp.EncodeToBytes(uint64(txIndex))
	if err != nil {
		return nil, common.Hash{}, err
	}
	proofDB := memorydb.New()
	if err := tr.Prove(key, proofDB); err != nil {
		return nil, common.Hash{}, fmt.Errorf("prove: %w", err)
	}
	nodes, err := orderProofNodes(proofDB, root, key)
	if err != nil {
		return nil, common.Hash{}, err
	}
	return nodes, root, nil
}

// orderProofNodes walks the proof from the root along the key, which is the order the onchain library expects.
func orderProofNodes(db *memorydb.Database, root common.Hash, key []byte) ([][]byte, error) {
	nibbles := make([]byte, 0, len(key)*2)
	for _, b := range key {
		nibbles = append(nibbles, b>>4, b&0x0f)
	}
	var out [][]byte
	next := root.Bytes()
	pos := 0
	for {
		node, err := db.Get(next)
		if err != nil {
			return nil, fmt.Errorf("proof node %x missing: %w", next, err)
		}
		out = append(out, node)
		ref, done, err := walkNode(node, nibbles, &pos)
		if err != nil {
			return nil, err
		}
		if done {
			return out, nil
		}
		next = ref
	}
}

// walkNode follows the key through one node and returns the next node hash, or done at the value.
func walkNode(node []byte, nibbles []byte, pos *int) ([]byte, bool, error) {
	elems, _, err := rlp.SplitList(node)
	if err != nil {
		return nil, false, err
	}
	count, err := rlp.CountValues(elems)
	if err != nil {
		return nil, false, err
	}
	switch count {
	case 17:
		if *pos == len(nibbles) {
			return nil, true, nil
		}
		child, err := listItem(elems, int(nibbles[*pos]))
		if err != nil {
			return nil, false, err
		}
		*pos++
		return followChild(child, nibbles, pos)
	case 2:
		pathItem, err := listItem(elems, 0)
		if err != nil {
			return nil, false, err
		}
		_, path, _, err := rlp.Split(pathItem)
		if err != nil {
			return nil, false, err
		}
		flag := path[0] >> 4
		var pathNibbles []byte
		if flag&1 == 1 {
			pathNibbles = append(pathNibbles, path[0]&0x0f)
		}
		for _, b := range path[1:] {
			pathNibbles = append(pathNibbles, b>>4, b&0x0f)
		}
		if !bytes.HasPrefix(nibbles[*pos:], pathNibbles) {
			return nil, false, fmt.Errorf("key path diverges from proof")
		}
		*pos += len(pathNibbles)
		if flag >= 2 {
			return nil, true, nil
		}
		child, err := listItem(elems, 1)
		if err != nil {
			return nil, false, err
		}
		return followChild(child, nibbles, pos)
	default:
		return nil, false, fmt.Errorf("unexpected node with %d items", count)
	}
}

func followChild(child []byte, nibbles []byte, pos *int) ([]byte, bool, error) {
	kind, content, _, err := rlp.Split(child)
	if err != nil {
		return nil, false, err
	}
	if kind == rlp.List {
		return walkNode(child, nibbles, pos)
	}
	if len(content) == 0 {
		return nil, false, fmt.Errorf("key not present in trie")
	}
	if len(content) != 32 {
		return nil, false, fmt.Errorf("unexpected child reference of %d bytes", len(content))
	}
	return content, false, nil
}

// listItem returns the raw RLP bytes of the i-th item, header included.
func listItem(elems []byte, i int) ([]byte, error) {
	rest := elems
	for {
		_, _, next, err := rlp.Split(rest)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			return rest[:len(rest)-len(next)], nil
		}
		rest = next
		i--
	}
}
