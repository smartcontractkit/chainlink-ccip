package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	bin "github.com/gagliardetto/binary"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
)

func IsEvent(event string, data []byte) bool {
	if len(data) < 8 {
		return false
	}
	d := Discriminator("event", event)
	return bytes.Equal(d, data[:8])
}

func NewEventNotFoundError(event string) error {
	return fmt.Errorf("%s: event not found", event)
}

// Note: This will not work for any events containing `Option`. This will be fixed when the version
// of anchor-go is updated to support events. In the meantime, events containing `Option` require
// bespoke parsing (see ParseEventCommitReportAccepted)
func ParseEvent(logs []string, event string, obj interface{}, shouldPrint ...bool) error {
	for _, v := range logs {
		if strings.Contains(v, "Program data:") {
			encodedData := strings.TrimSpace(strings.TrimPrefix(v, "Program data:"))
			data, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return err
			}
			if IsEvent(event, data) {
				if err := bin.UnmarshalBorsh(obj, data); err != nil {
					return err
				}

				if len(shouldPrint) > 0 && shouldPrint[0] {
					fmt.Printf("%s: %+v\n", event, obj)
				}
				return nil
			}
		}
	}
	return NewEventNotFoundError(event)
}

func ParseMultipleEvents[T any](logs []string, event string, shouldPrint bool) ([]T, error) {
	var results []T
	for _, v := range logs {
		if strings.Contains(v, "Program data:") {
			encodedData := strings.TrimSpace(strings.TrimPrefix(v, "Program data:"))
			data, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return nil, err
			}
			if IsEvent(event, data) {
				var obj T
				if err := bin.UnmarshalBorsh(&obj, data); err != nil {
					return nil, err
				}

				if shouldPrint {
					fmt.Printf("%s: %+v\n", event, obj)
				}

				results = append(results, obj)
			}
		}
	}
	if len(results) == 0 {
		return nil, NewEventNotFoundError(event)
	}

	return results, nil
}

// Redefined to avoid a circular dependency
type EventCommitReportAccepted struct {
	Discriminator [8]byte
	Report        *ccip_offramp.MerkleRoot `bin:"optional"`
	PriceUpdates  ccip_offramp.PriceUpdates
}

// This event uses bespoke parsing, as it contains an `Optional` field which cannot be unmarshalled through `bin.UnmarshalBorsh`
func ParseEventCommitReportAccepted(logs []string, event string, obj *EventCommitReportAccepted) error {
	for _, v := range logs {
		if strings.Contains(v, "Program data:") {
			encodedData := strings.TrimSpace(strings.TrimPrefix(v, "Program data:"))
			data, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return err
			}
			if IsEvent(event, data) {
				decoder := bin.NewBorshDecoder(data)

				// Deserialize `Discriminator`:
				err = decoder.Decode(&obj.Discriminator)
				if err != nil {
					return err
				}

				// Deserialize `Report` (optional):
				{
					ok, dErr := decoder.ReadBool()
					if dErr != nil {
						return dErr
					}
					if ok {
						dErr = decoder.Decode(&obj.Report)
						if dErr != nil {
							return dErr
						}
					}
				}
				// Deserialize `PriceUpdates`:
				err = decoder.Decode(&obj.PriceUpdates)
				if err != nil {
					return err
				}

				return nil
			}
		}
	}
	return NewEventNotFoundError(event)
}
