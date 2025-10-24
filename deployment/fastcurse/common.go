package fastcurse

import (
	"encoding/binary"
)

type Subject = [16]byte

type CurseInput struct {
	Subjects      []Subject
	ChainSelector uint64
}

// GlobalCurseSubject is defined here - https://github.com/smartcontractkit/chainlink-ccip/blob/main/chains/evm/contracts/rmn/RMNRemote.sol#L12
// An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject
// for issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of
// using the local chain selector as a subject.
func GlobalCurseSubject() Subject {
	return Subject{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
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
