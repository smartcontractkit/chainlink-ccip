package executor

import "github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"

// TransmitterPayload is a struct that holds a message and its associated attestations
// We use this struct inside the main executor loop as a trigger for the contract transmitter to send messages
type TransmitterPayload struct {
	Message      modsectypes.Message
	Attestations []modsectypes.Attestation // List of attestations for the message
}
