package ocrimpls_test

import (
	"crypto/rand"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccipcapability/ccipevm"
	"github.com/smartcontractkit/chainlink-ccip/ccipcapability/ccipsolana"

	_ "github.com/smartcontractkit/chainlink-ccip/ccipcapability/ccipevm"    // Register EVM plugin config factories
	_ "github.com/smartcontractkit/chainlink-ccip/ccipcapability/ccipsolana" // Register Solana plugin config factories
	ccipcommon "github.com/smartcontractkit/chainlink-ccip/ccipcapability/common"
	"github.com/smartcontractkit/chainlink-ccip/ccipcapability/ocrimpls"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func randomReport(t *testing.T, len int) []byte {
	report := make([]byte, len)
	_, err := rand.Reader.Read(report)
	require.NoError(t, err, "failed to read random bytes")
	return report
}

func abiEncodeUint32(data uint32) ([]byte, error) {
	return utils.ABIEncode(`[{ "type": "uint32" }]`, data)
}

// Test EVM -> SVM extra data decoding in contract transmitter
func TestSVMExecCallDataFuncExtraDataDecoding(t *testing.T) {
	extraDataCodec := ccipcommon.ExtraDataCodec(map[string]ccipcommon.SourceChainExtraDataCodec{
		chainsel.FamilyEVM:    ccipevm.ExtraDataDecoder{},
		chainsel.FamilySolana: ccipsolana.ExtraDataDecoder{},
	})
	t.Run("fails when multiple reports are included", func(t *testing.T) {
		reports := []ccipocr3.ExecutePluginReportSingleChain{{}, {}}
		reportWithInfo := ccipocr3.ExecuteReportInfo{
			AbstractReports: reports,
		}

		encodedExecReport, err := reportWithInfo.Encode()
		require.NoError(t, err)

		rwi := ocr3types.ReportWithInfo[[]byte]{
			Report: randomReport(t, 96),
			Info:   encodedExecReport,
		}
		_, _, _, err = ocrimpls.SVMExecCalldataFunc([2][32]byte{}, rwi, nil, nil, [32]byte{}, extraDataCodec)
		require.Contains(t, err.Error(), "unexpected report length, expected 1, got 2")
	})
	t.Run("fails when multiple report contains multiple messages", func(t *testing.T) {
		reports := []ccipocr3.ExecutePluginReportSingleChain{{
			Messages: []ccipocr3.Message{{}, {}},
		}}
		reportWithInfo := ccipocr3.ExecuteReportInfo{
			AbstractReports: reports,
		}

		encodedExecReport, err := reportWithInfo.Encode()
		require.NoError(t, err)

		rwi := ocr3types.ReportWithInfo[[]byte]{
			Report: randomReport(t, 96),
			Info:   encodedExecReport,
		}
		_, _, _, err = ocrimpls.SVMExecCalldataFunc([2][32]byte{}, rwi, nil, nil, [32]byte{}, extraDataCodec)
		require.Contains(t, err.Error(), "unexpected message length, expected 1, got 2")
	})
	t.Run("fails with invalid extra args", func(t *testing.T) {
		// invalid encoded extra args
		encoded := []byte{1, 2, 3, 4}

		report := ccipocr3.ExecutePluginReportSingleChain{
			SourceChainSelector: 5009297550715157269,
			Messages: []ccipocr3.Message{{
				Header: ccipocr3.RampMessageHeader{
					// EVM
					SourceChainSelector: 5009297550715157269,
					// to SOL
					DestChainSelector: 124615329519749607,
				},
				ExtraArgs: encoded,
			}},
		}

		reportWithInfo := ccipocr3.ExecuteReportInfo{
			AbstractReports: []ccipocr3.ExecutePluginReportSingleChain{report},
		}

		encodedExecReport, err := reportWithInfo.Encode()
		require.NoError(t, err)

		rwi := ocr3types.ReportWithInfo[[]byte]{
			Report: randomReport(t, 96),
			Info:   encodedExecReport,
		}

		_, _, _, err = ocrimpls.SVMExecCalldataFunc([2][32]byte{}, rwi, nil, nil, [32]byte{}, extraDataCodec)
		require.Contains(t, err.Error(), "unknown extra args tag")
	})
	t.Run("fails with invalid extra exec data", func(t *testing.T) {
		// invalid encoded extra args
		encoded := []byte{31, 59, 58, 186, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 39, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170}
		encodedExecData := []byte{1, 2, 3, 4}

		report := ccipocr3.ExecutePluginReportSingleChain{
			SourceChainSelector: 5009297550715157269,
			Messages: []ccipocr3.Message{{
				Header: ccipocr3.RampMessageHeader{
					// EVM
					SourceChainSelector: 5009297550715157269,
					// to SOL
					DestChainSelector: 124615329519749607,
				},
				ExtraArgs: encoded,
				TokenAmounts: []ccipocr3.RampTokenAmount{{
					DestExecData: encodedExecData,
				}},
			}},
		}

		reportWithInfo := ccipocr3.ExecuteReportInfo{
			AbstractReports: []ccipocr3.ExecutePluginReportSingleChain{report},
		}

		encodedExecReport, err := reportWithInfo.Encode()
		require.NoError(t, err)

		rwi := ocr3types.ReportWithInfo[[]byte]{
			Report: randomReport(t, 96),
			Info:   encodedExecReport,
		}

		_, _, _, err = ocrimpls.SVMExecCalldataFunc([2][32]byte{}, rwi, nil, nil, [32]byte{}, extraDataCodec)
		require.Contains(t, err.Error(), "failed to decode token amount dest exec data: decode dest gas amount: abi decode uint32: abi: cannot marshal in to go type: length insufficient 4 require 32")
	})
	t.Run("Successfully decodes valid EVM -> SOL report", func(t *testing.T) {
		// hardcode abi encoded extra args for simplicity
		encoded := []byte{31, 59, 58, 186, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 39, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170}
		destGasAmount := uint32(10000)
		encodedExecData, err := abiEncodeUint32(destGasAmount)
		require.NoError(t, err)

		report := ccipocr3.ExecutePluginReportSingleChain{
			SourceChainSelector: 5009297550715157269,
			Messages: []ccipocr3.Message{{
				Header: ccipocr3.RampMessageHeader{
					// EVM
					SourceChainSelector: 5009297550715157269,
					// to SOL
					DestChainSelector: 124615329519749607,
				},
				ExtraArgs: encoded,
				TokenAmounts: []ccipocr3.RampTokenAmount{{
					DestExecData: encodedExecData,
				}},
			}},
		}

		reportWithInfo := ccipocr3.ExecuteReportInfo{
			AbstractReports: []ccipocr3.ExecutePluginReportSingleChain{report},
		}

		encodedExecReport, err := reportWithInfo.Encode()
		require.NoError(t, err)

		rwi := ocr3types.ReportWithInfo[[]byte]{
			Report: randomReport(t, 96),
			Info:   encodedExecReport,
		}

		_, _, args, err := ocrimpls.SVMExecCalldataFunc([2][32]byte{}, rwi, nil, nil, [32]byte{}, extraDataCodec)
		require.NoError(t, err)

		expectedArgs, ok := args.(ocrimpls.SVMExecCallArgs)
		require.True(t, ok)

		require.Equal(t, uint64(0x4), expectedArgs.ExtraData.ExtraArgsDecoded["accountIsWritableBitmap"])
		require.Equal(t, [32]uint8{44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170}, expectedArgs.ExtraData.ExtraArgsDecoded["accounts"].([][32]byte)[0])
		require.False(t, expectedArgs.ExtraData.ExtraArgsDecoded["allowOutOfOrderExecution"].(bool))
		require.Equal(t, destGasAmount, expectedArgs.ExtraData.ExtraArgsDecoded["computeUnits"])
		require.Equal(t, [32]uint8{44, 230, 105, 156, 244, 184, 196, 235, 30, 58, 209, 82, 8, 202, 25, 73, 167, 169, 34, 150, 141, 129, 169, 150, 219, 160, 186, 44, 72, 156, 50, 170}, expectedArgs.ExtraData.ExtraArgsDecoded["tokenReceiver"])
		require.Equal(t, destGasAmount, expectedArgs.ExtraData.DestExecDataDecoded[0]["destGasAmount"])
	})
}
