package executor

import (
	"hash/fnv"
	"slices"
	"sync"
)

// LeaderElection provides a way to determine if an executor should be currently executing a given message.
// This is used to prevent multiple executors from executing the same message concurrently.
type LeaderElection interface {
	IsLeader(msgId [32]byte, destChainSelector uint64, offset uint8) bool
	SelfParticipations() []uint64 // Think we need a better name for this?
	IsParticipant(destChainSelector uint64) bool
	GetParticipants(destChainSelector uint64) []string
	SetParticipants(destChainSelector uint64, participants []string)
}

type SimpleLeaderElection struct {
	mu                 *sync.RWMutex
	selfParticipations []uint64
	chainParticipants  map[uint64][]string
	self               string
}

func NewSimpleLeaderElection(self string, chainParticipants map[uint64][]string) LeaderElection {
	le := &SimpleLeaderElection{
		mu:                &sync.RWMutex{},
		chainParticipants: chainParticipants,
		self:              self,
	}

	// Initialize self participations
	le.writeSelfParticipations()

	return le
}

func (s *SimpleLeaderElection) SelfParticipations() []uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.selfParticipations
}

func (s *SimpleLeaderElection) IsLeader(msgId [32]byte, destChainSelector uint64, offset uint8) bool {
	// If the executor is not a participant, it cannot be the leader
	if !s.IsParticipant(destChainSelector) {
		return false
	}

	participants := s.GetParticipants(destChainSelector)

	// If the offset is greater than the number of participants, it cannot be the leader
	if len(participants) < int(offset)+1 {
		return false
	}

	// Concatenates each participant with the msgId and the participantId together and hashes the resulting value into a uint64
	// Given multiple participants, this will have an equal chance of resulting in the offset value, and will therefore be the leader

	// Currently this is using FNV-1a which is a non-cryptographic hash function, but it's fast and has a good distribution
	// However there is a non-zero chance of collisions

	hashes := []uint64{}
	selfHash := fnv.New64a()
	selfHash.Write([]byte(s.self + string(msgId[:])))
	selfHashValue := selfHash.Sum64()

	for _, participant := range participants {
		h := fnv.New64a()
		h.Write([]byte(participant + string(msgId[:])))
		hashes = append(hashes, h.Sum64())
	}

	// Sort the numerical values outputted by the hash function
	slices.Sort(hashes)

	// If my hash matches the offset, I am the leader
	return selfHashValue == hashes[offset]
}

func (s *SimpleLeaderElection) IsParticipant(destChainSelector uint64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Contains(s.selfParticipations, destChainSelector)
}

func (s *SimpleLeaderElection) GetParticipants(destChainSelector uint64) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.chainParticipants[destChainSelector]
}

func (s *SimpleLeaderElection) SetParticipants(destChainSelector uint64, participants []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.chainParticipants[destChainSelector] = participants
	s.writeSelfParticipations()
}

func (s *SimpleLeaderElection) writeSelfParticipations() {
	participations := []uint64{}
	for chainSelector, participants := range s.chainParticipants {
		for _, participant := range participants {
			if participant == s.self {
				participations = append(participations, chainSelector)
			}
		}
	}

	s.selfParticipations = participations
}
