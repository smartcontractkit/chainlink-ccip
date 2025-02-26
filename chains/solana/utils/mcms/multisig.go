package mcms

import (
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

type Multisig struct {
	PaddedID                  [32]byte
	SignerPDA                 solana.PublicKey
	ConfigPDA                 solana.PublicKey
	RootMetadataPDA           solana.PublicKey
	ExpiringRootAndOpCountPDA solana.PublicKey
	ConfigSignersPDA          solana.PublicKey
	RootSignaturesPDA         func(root [32]byte, validUntil uint32, authority solana.PublicKey) solana.PublicKey
	SeenSignedHashesPDA       func(root [32]byte, validUntil uint32) solana.PublicKey

	RawConfig McmConfigArgs
	Signers   []eth.Signer
}

type SignatureAnalyzer struct {
	signerGroups  []byte
	groupQuorums  [32]uint8
	groupParents  [32]uint8
	groupMembers  map[byte][]string
	visitedGroups map[byte]bool
	numGroups     int // Track actual number of groups
}

func NewSignatureAnalyzer(signerGroups []byte, groupQuorums, groupParents [32]uint8) *SignatureAnalyzer {
	// Find actual number of groups
	numGroups := 0
	for i, quorum := range groupQuorums {
		if quorum > 0 {
			if i+1 > numGroups {
				numGroups = i + 1
			}
		}
	}

	sa := &SignatureAnalyzer{
		signerGroups:  signerGroups,
		groupQuorums:  groupQuorums,
		groupParents:  groupParents,
		groupMembers:  make(map[byte][]string),
		visitedGroups: make(map[byte]bool),
		numGroups:     numGroups,
	}

	for i, group := range signerGroups {
		if int(group) < numGroups {
			sa.groupMembers[group] = append(sa.groupMembers[group], fmt.Sprintf("signer_%d", i))
		}
	}

	return sa
}

func (sa *SignatureAnalyzer) getChildGroups(groupID byte) []byte {
	var children []byte
	for i := 0; i < sa.numGroups; i++ {
		if sa.groupParents[i] == groupID {
			children = append(children, byte(i))
		}
	}
	return children
}

func (sa *SignatureAnalyzer) getRequiredSigners(groupID byte) map[string]bool {
	if sa.visitedGroups[groupID] {
		return make(map[string]bool)
	}
	sa.visitedGroups[groupID] = true

	possibleSigners := make(map[string]bool)

	// Add direct members
	for _, signer := range sa.groupMembers[groupID] {
		possibleSigners[signer] = true
	}

	// Add child group members
	childGroups := sa.getChildGroups(groupID)
	for _, child := range childGroups {
		childSigners := sa.getRequiredSigners(child)
		for signer := range childSigners {
			possibleSigners[signer] = true
		}
	}

	return possibleSigners
}

func (sa *SignatureAnalyzer) PrintGroupStructure() {
	fmt.Println("\nGroup Structure Analysis:")

	for groupID := 0; groupID < sa.numGroups; groupID++ {
		fmt.Printf("\nGroup %d:\n", groupID)
		fmt.Printf("Quorum required: %d\n", sa.groupQuorums[groupID])

		parentID := sa.groupParents[groupID]
		parentStr := "root"

		//nolint:gosec
		if parentID != uint8(groupID) {
			parentStr = fmt.Sprintf("%d", parentID)
		}
		fmt.Printf("Parent group: %s\n", parentStr)

		fmt.Printf("Direct members: %v\n", sa.groupMembers[byte(groupID)])
		fmt.Printf("Child groups: %v\n", sa.getChildGroups(byte(groupID)))
	}

	possibleSigners := sa.getRequiredSigners(0)
	signersList := make([]string, 0, len(possibleSigners))
	for signer := range possibleSigners {
		signersList = append(signersList, signer)
	}

	fmt.Printf("\nSignature Requirements:\n")
	fmt.Printf("Root group (Group 0) requires %d signatures\n", sa.groupQuorums[0])
	fmt.Printf("Possible signers: [%s]\n", strings.Join(signersList, ", "))
}
