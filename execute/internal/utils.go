package internal

import "encoding/json"

func EncodedSize[T any](obj T) int {
	enc, err := json.Marshal(obj)
	if err != nil {
		return 0
	}
	return len(enc)
}

func RemoveIthElement[T any](slice []T, i int) []T {
	newSlice := make([]T, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	return append(newSlice, slice[i+1:]...)
}
