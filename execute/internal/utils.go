package internal

import "encoding/json"

func EncodedSize[T any](obj T) int {
	enc, err := json.Marshal(obj)
	if err != nil {
		return 0
	}
	return len(enc)
}
