package tokenprice

import (
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	bi100        = big.NewInt(100)
	bi200        = big.NewInt(200)
	tokenA       = cciptypes.UnknownEncodedAddress("0xAAAAAAAAAAAAAAAa75C1216873Ec4F88C11E57E3")
	cbi100       = cciptypes.BigInt{Int: bi100}
	cbi200       = cciptypes.BigInt{Int: bi200}
	tokenB       = cciptypes.UnknownEncodedAddress("0xBBBBBBBBBBBBBBBb75C1216873Ec4F88C11E57E3")
	tokenC       = cciptypes.UnknownEncodedAddress("0xCCCCCCCCCCCCCCCc75C1216873Ec4F88C11E57E3")
	tokenD       = cciptypes.UnknownEncodedAddress("0xDDDDDDDDDDDDDDDd75C1216873Ec4F88C11E57E3")
	feedChainSel = cciptypes.ChainSelector(1)
	destChainSel = cciptypes.ChainSelector(2)
	f            = 1
)

func bi(i int) *big.Int {
	return big.NewInt(int64(i))
}
func cbi(i int) cciptypes.BigInt {
	return cciptypes.BigInt{Int: bi(i)}
}
