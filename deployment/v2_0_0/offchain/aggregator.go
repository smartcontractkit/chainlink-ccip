package offchain

type SourceSelector = string

type DestinationSelector = string

type Signer struct {
	Address string `json:"address"`
}

type QuorumConfig struct {
	SourceVerifierAddress string   `json:"sourceVerifierAddress"`
	Signers               []Signer `json:"signers"`
	Threshold             uint8    `json:"threshold"`
}

type Committee struct {
	QuorumConfigs        map[SourceSelector]*QuorumConfig      `json:"quorumConfigs"`
	DestinationVerifiers map[DestinationSelector]string `json:"destinationVerifiers"`
}
