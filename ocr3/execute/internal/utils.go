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
	if i < 0 || i >= len(slice) {
		return slice // Return the original slice if index is out of bounds
	}
	newSlice := make([]T, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	return append(newSlice, slice[i+1:]...)
}
