package state

import (
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

//////////////////////
// CCIP Router PDAs //
//////////////////////

func FindConfigPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, ccipRouterProgram)
}

func FindFeeBillingSignerPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, ccipRouterProgram)
}

func FindTokenAdminRegistryPDA(mint solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.Bytes()}, ccipRouterProgram)
}

func FindDestChainStatePDA(chainSelector uint64, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), chainSelectorLE}, ccipRouterProgram)
	return p, err
}

func FindNoncePDA(chainSelector uint64, user solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, user.Bytes()}, ccipRouterProgram)
	return p, err
}

func FindApprovedSender(chainSelector uint64, sourceSender []byte, receiverProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("approved_ccip_sender"), chainSelectorLE, []byte{uint8(len(sourceSender))}, sourceSender}, receiverProgram) //nolint:gosec
	return p, err
}

func FindAllowedOfframpPDA(chainSelector uint64, offramp solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("allowed_offramp"), chainSelectorLE, offramp.Bytes()}, ccipRouterProgram)
	return p, err
}

//////////////////////////////////////////
// PDAs with same seeds across programs //
//////////////////////////////////////////

func FindExternalTokenPoolsSignerPDA(poolProgram solana.PublicKey, program solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer"), poolProgram.Bytes()}, program)
}

func FindExternalExecutionConfigPDA(logicReceiver solana.PublicKey, program solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_execution_config"), logicReceiver.Bytes()}, program)
}

/////////////////////
// Fee Quoter PDAs //
/////////////////////

func FindFqConfigPDA(feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, feeQuoterProgram)
}

func FindFqDestChainPDA(chainSelector uint64, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("dest_chain"), chainSelectorLE}, feeQuoterProgram)
}

func FindFqBillingTokenConfigPDA(mint solana.PublicKey, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, feeQuoterProgram)
}

func FindFqPerChainPerTokenConfigPDA(chainSelector uint64, mint solana.PublicKey, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("per_chain_per_token_config"), chainSelectorLE, mint.Bytes()}, feeQuoterProgram)
}

func FindFqAllowedPriceUpdaterPDA(priceUpdater solana.PublicKey, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("allowed_price_updater"), priceUpdater.Bytes()}, feeQuoterProgram)
}

//////////////////
// Offramp PDAs //
//////////////////

func FindOfframpConfigPDA(offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, offrampProgram)
}

func FindOfframpReferenceAddressesPDA(offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("reference_addresses")}, offrampProgram)
}

func FindOfframpSourceChainPDA(chainSelector uint64, offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), chainSelectorLE}, offrampProgram)
}

func FindOfframpBillingSignerPDA(offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, offrampProgram)
}

func FindOfframpStatePDA(offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("state")}, offrampProgram)
}

func FindOfframpCommitReportPDA(chainSelector uint64, root [32]byte, offrampProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("commit_report"), chainSelectorLE, root[:]}, offrampProgram)
	return p, err
}

/////////////////////
// RMN Remote PDAs //
/////////////////////

func FindRMNRemoteConfigPDA(rmnRemoteProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, rmnRemoteProgram)
}

func FindRMNRemoteCursesPDA(rmnRemoteProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("curses")}, rmnRemoteProgram)
}

////////////////////
// Ping Pong Demo //
////////////////////

func FindPingPongDemoConfigPDA(pingPongDemoProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, pingPongDemoProgram)
}

func FindPingPongCCIPSendSignerPDA(pingPongDemoProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("ccip_send_signer")}, pingPongDemoProgram)
}

/////////////////
// Shared PDAs //
/////////////////

func FindNameAndVersionPDA(program solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("name_version")}, program)
}
