package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	bin "github.com/gagliardetto/binary"
)

func IsEvent(event string, data []byte) bool {
	if len(data) < 8 {
		return false
	}
	d := Discriminator("event", event)
	return bytes.Equal(d, data[:8])
}

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
	return fmt.Errorf("%s: event not found", event)
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
		return nil, fmt.Errorf("%s: event not found", event)
	}

	return results, nil
}
