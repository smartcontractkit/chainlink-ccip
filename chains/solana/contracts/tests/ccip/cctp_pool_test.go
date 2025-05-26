package contracts

// import (
// 	"testing"

// 	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"

// 	cctp_message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_message_transmitter"
// 	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/rmn_remote"
// )

// func TestCctpPool(t *testing.T) {
// 	t.Parallel()

// 	cctp_message_transmitter.SetProgramID(config.CctpMessageTransmitter)
// 	rmn_remote.SetProgramID(config.RMNRemoteProgram)

// 	t.Run("CctpPool", func(t *testing.T) {
// 		t.Parallel()

// 		cctp_message_transmitter.NewReceiveMessageInstruction(
// 			cctp_message_transmitter.ReceiveMessageParams{},
// 		)
// 	})

// }
