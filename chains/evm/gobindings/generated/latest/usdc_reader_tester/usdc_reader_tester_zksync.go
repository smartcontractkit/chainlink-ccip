// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package usdc_reader_tester

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/types"
)

func DeployUSDCReaderTesterZk(deployOpts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ...interface{}) (common.Address, *types.Receipt, *USDCReaderTester, error) {
	var calldata []byte
	if len(args) > 0 {
		abi, err := USDCReaderTesterMetaData.GetAbi()
		if err != nil {
			return common.Address{}, nil, nil, err
		}
		calldata, err = abi.Pack("", args...)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	salt := make([]byte, 32)
	n, err := rand.Read(salt)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if n != len(salt) {
		return common.Address{}, nil, nil, fmt.Errorf("failed to read random bytes: expected %d, got %d", len(salt), n)
	}

	txHash, err := wallet.Deploy(deployOpts, accounts.Create2Transaction{
		Bytecode: ZkBytecode,
		Calldata: calldata,
		Salt:     salt,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	receipt, err := client.WaitMined(context.Background(), txHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address := receipt.ContractAddress
	contract, err := NewUSDCReaderTester(address, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return address, receipt, contract, nil
}

var ZkBytecode = common.Hex2Bytes("0000008003000039000000400030043f0000000100200190000000b20000c13d00000060021002700000003302200197000000040020008c000000ba0000413d000000000301043b0000003503300197000000360030009c000000ba0000c13d000001040020008c000000ba0000413d0000000003000416000000000003004b000000ba0000c13d0000000403100370000000000303043b000000330030009c000000ba0000213d0000002404100370000000000404043b000000330040009c000000ba0000213d0000004405100370000000000505043b000000330050009c000000ba0000213d000000c406100370000000000606043b000000370060009c000000ba0000213d000000e407100370000000000907043b000000370090009c000000ba0000213d0000002307900039000000000027004b000000ba0000813d0000000408900039000000000781034f000000000707043b000000370070009c000000ba0000213d00000000097900190000002409900039000000000029004b000000ba0000213d0000001f097000390000003f099001970000003f099000390000003f09900197000000380090009c000000bc0000813d0000008009900039000000400090043f0000002008800039000000000981034f000000800070043f0000003f0a7001980000001f0b70018f000000a008a00039000000460000613d000000a00c000039000000000d09034f00000000de0d043c000000000cec043600000000008c004b000000420000c13d00000000000b004b000000530000613d0000000009a9034f000000030ab00210000000000b080433000000000bab01cf000000000bab022f000000000909043b000001000aa000890000000009a9022f0000000009a901cf0000000009b9019f0000000000980435000000a0077000390000000000070435000000e008300210000000400700043d00000020037000390000000000830435000000e00440021000000024087000390000000000480435000000e00450021000000028057000390000000000450435000000c0046002100000002c057000390000000000450435000000a404100370000000000404043b000000340570003900000000004504350000006404100370000000000404043b000000540570003900000000004504350000008401100370000000000101043b000000740470003900000000001404350000009404700039000000800100043d000000000001004b0000007a0000613d00000000050000190000000006450019000000a008500039000000000808043300000000008604350000002005500039000000000015004b000000730000413d0000000004410019000000000004043500000074041000390000000000470435000000b3011000390000003f041001970000000001740019000000000041004b00000000040000390000000104004039000000370010009c000000bc0000213d0000000100400190000000bc0000c13d000000400010043f00000020040000390000000005410436000000000407043300000000004504350000004005100039000000000004004b000000980000613d000000000600001900000000075600190000000008360019000000000808043300000000008704350000002006600039000000000046004b000000910000413d0000001f034000390000003f023001970000000003540019000000000003043500000040022000390000006003200210000000390020009c0000003a03008041000000330010009c00000033010080410000004001100210000000000113019f0000000002000414000000330020009c0000003302008041000000c00220021000000000012100190000003b0110009a0000800d0200003900000001030000390000003c0400004100c700c20000040f0000000100200190000000ba0000613d0000000001000019000000c80001042e0000000001000416000000000001004b000000ba0000c13d0000002001000039000001000010044300000120000004430000003401000041000000c80001042e0000000001000019000000c9000104300000003d01000041000000000010043f0000004101000039000000040010043f0000003e01000041000000c900010430000000c5002104210000000102000039000000000001042d0000000002000019000000000001042d000000c700000432000000c80001042e000000c9000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000200000000000000000000000000000040000001000000000000000000ffffffff0000000000000000000000000000000000000000000000000000000062826f1800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffff80000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000ffffffff000000000000000000000000fe000000000000000000000000000000000000000000000000000000000000008c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b0364e487b71000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01411d572348fd0239e3447438a1c67a7ef5cee13fe794f6d8340350d432628d5")
