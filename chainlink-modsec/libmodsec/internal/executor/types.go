package executor

import "github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"

type VerifierId string

// Attestation represents a signature of a message by a verifier.
// TODO: Should these be called "Proofs" instead?
type Attestation struct {
	Proof      []byte   // Proof of the message
	VerifierId [32]byte // Identifier of the verifier that signed the message
}

// TransmitterPayload is a struct that holds a message and its associated attestations
// We use this struct inside the main executor loop as a trigger for the contract transmitter to send messages
type TransmitterPayload struct {
	Message      modsectypes.Message
	Attestations []Attestation // List of attestations for the message
}
