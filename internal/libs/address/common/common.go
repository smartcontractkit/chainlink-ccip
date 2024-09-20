package common

// Address is the specific address family which can encode itself.
type Address interface {
	Encode() (EncodedAddress, error)
}

// EncodedAddress is the specific encoded address family which can decode itself.
type EncodedAddress interface {
	Decode() (Address, error)
}
