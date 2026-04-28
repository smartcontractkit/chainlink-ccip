// Package registrations registers all known EVM token contract strategies
// with the singleton strategy.Registry at init time. Adapters that need
// per-token-type behavior pull in this package via a blank import; all
// known token types become available in one line.
//
// Convention: each strategy lives in its own file with its own init() that
// registers itself. Adding a new EVM token contract type means dropping in
// one new file (strategy struct + init); no central registry list to amend.
package registrations
