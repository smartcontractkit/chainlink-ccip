package fastcurse

import (
	"bytes"
	"encoding/binary"
)

type Subject = [16]byte

func IfSubjectEqual(s, other Subject) bool {
	byteSub := s[:]
	otherByteSub := other[:]
	return bytes.Equal(byteSub, otherByteSub)
}

type CurseInput struct {
	Subjects      []Subject
	ChainSelector uint64
	// MCMSQualifier is the qualifier of the MCMS stack that will execute the resulting
	// proposal. Families whose RMN curse entrypoint takes an explicit caller argument
	// (Stellar: curse(caller, subjects), where the caller must be the invoking timelock)
	// need it to resolve that timelock at proposal-build time. Empty on direct /
	// no-MCMS runs. Families that authorize by transaction sender ignore it.
	MCMSQualifier string
}

// GlobalCurseSubject is defined here - https://github.com/smartcontractkit/chainlink-ccip/blob/main/chains/evm/contracts/rmn/RMNRemote.sol#L12
// An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject
// for issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of
// using the local chain selector as a subject.
func GlobalCurseSubject() Subject {
	return Subject{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
}

// FiredrillSubject is a hardcoded, inert subject used only to exercise the
// curse/uncurse changeset flow end-to-end. It is not derived from any chain
// selector and is not the GlobalCurseSubject, so cursing/uncursing it has no
// effect on real lanes or the global curse state.
func FiredrillSubject() Subject {
	return Subject{0xFD, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFD}
}

func GenericSelectorToSubject(selector uint64) Subject {
	var b Subject
	binary.BigEndian.PutUint64(b[8:], selector)
	return b
}

func GenericSubjectToSelector(subject Subject) (uint64, error) {
	if subject == GlobalCurseSubject() {
		return 0, nil
	}

	return binary.BigEndian.Uint64(subject[8:]), nil
}
