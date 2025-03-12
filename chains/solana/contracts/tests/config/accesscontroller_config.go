package config

import (
	"github.com/gagliardetto/solana-go"
)

var (
	AccessControllerProgram = solana.MustPublicKeyFromBase58("6KsN58MTnRQ8FfPaXHiFPPFGDRioikj9CdPvPxZJdCjb")
	AccSpace                = uint64(8 + 32 + 32 + ((32 * 64) + 8)) // discriminator + owner + proposed owner + access_list (64 max addresses + length)
)
