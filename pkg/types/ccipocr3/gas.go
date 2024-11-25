package ccipocr3

// EstimateProvider is used to estimate the gas cost of a message or a merkle tree.
type EstimateProvider interface {
	CalculateMerkleTreeGas(numRequests int) uint64
	CalculateMessageMaxGas(msg Message) uint64
}
