// Package evmimpls implements the EVM-specific implementations of the modsec types
// using go-ethereum primitives.
//
// The intention of this package is to be able to test the modsec implementation outside
// of the chainlink node execution environment.
//
// It's possible that some types are usable in non-chainlink contexts, such as the codecs.
package evmimpls
