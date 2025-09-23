package ccip

// Mod sec
//go:generate go run generation/generate/wrap.go ccip VerifierProxy verifier_proxy latest
//go:generate go run generation/generate/wrap.go ccip CommitteeVerifier committee_verifier latest
//go:generate go run generation/generate/wrap.go ccip CCVProxy ccv_proxy latest
//go:generate go run generation/generate/wrap.go ccip ExecutorOnRamp executor_onramp latest
//go:generate go run generation/generate/wrap.go ccip CCVAggregator ccv_aggregator latest

//go:generate go run generation/generate/wrap.go ccip Router router latest
//go:generate go run generation/generate/wrap.go ccip CCIPHome ccip_home latest
//go:generate go run generation/generate/wrap.go ccip OnRamp onramp latest
//go:generate go run generation/generate/wrap.go ccip OffRamp offramp latest
//go:generate go run generation/generate/wrap.go ccip OnRampWithMessageTransformer onramp_with_message_transformer latest
//go:generate go run generation/generate/wrap.go ccip OffRampWithMessageTransformer offramp_with_message_transformer latest
//go:generate go run generation/generate/wrap.go ccip FeeQuoter fee_quoter latest
//go:generate go run generation/generate/wrap.go ccip FeeQuoterV2 fee_quoter_v2 latest
//go:generate go run generation/generate/wrap.go ccip NonceManager nonce_manager latest
//go:generate go run generation/generate/wrap.go ccip TokenAdminRegistry token_admin_registry latest
//go:generate go run generation/generate/wrap.go ccip TokenPoolFactory token_pool_factory latest
//go:generate go run generation/generate/wrap.go ccip FactoryBurnMintERC20 factory_burn_mint_erc20 latest
//go:generate go run generation/generate/wrap.go ccip RegistryModuleOwnerCustom registry_module_owner_custom latest
//go:generate go run generation/generate/wrap.go ccip RMNProxy rmn_proxy_contract latest
//go:generate go run generation/generate/wrap.go ccip RMNRemote rmn_remote latest
//go:generate go run generation/generate/wrap.go ccip RMNHome rmn_home latest
//go:generate go run generation/generate/wrap.go ccip HyperLiquidCompatibleERC20 hyper_liquid_compatible_erc20 latest

// Pools
//go:generate go run generation/generate/wrap.go ccip BurnMintTokenPool burn_mint_token_pool latest
//go:generate go run generation/generate/wrap.go ccip BurnFromMintTokenPool burn_from_mint_token_pool latest
//go:generate go run generation/generate/wrap.go ccip BurnWithFromMintTokenPool burn_with_from_mint_token_pool latest
//go:generate go run generation/generate/wrap.go ccip LockReleaseTokenPool lock_release_token_pool latest
//go:generate go run generation/generate/wrap.go ccip TokenPool token_pool latest
//go:generate go run generation/generate/wrap.go ccip USDCTokenPool usdc_token_pool latest
//go:generate go run generation/generate/wrap.go ccip SiloedLockReleaseTokenPool siloed_lock_release_token_pool latest
//go:generate go run generation/generate/wrap.go ccip BurnToAddressMintTokenPool burn_to_address_mint_token_pool latest
//go:generate go run generation/generate/wrap.go ccip BurnMintFastTransferTokenPool fast_transfer_token_pool latest
//go:generate go run generation/generate/wrap.go ccip CCTPMessageTransmitterProxy cctp_message_transmitter_proxy latest
//go:generate go run generation/generate/wrap.go ccip ERC20LockBox erc20_lock_box latest
//go:generate go run generation/generate/wrap.go ccip SiloedUSDCTokenPool siloed_usdc_token_pool latest
//go:generate go run generation/generate/wrap.go ccip USDCTokenPoolCCTPV2 usdc_token_pool_cctp_v2 latest
//go:generate go run generation/generate/wrap.go ccip BurnMintWithLockReleaseFlagTokenPool burn_mint_with_lock_release_flag_token_pool latest

// Helpers
//go:generate go run generation/generate/wrap.go ccip MaybeRevertMessageReceiver maybe_revert_message_receiver latest
//go:generate go run generation/generate/wrap.go ccip LogMessageDataReceiver log_message_data_receiver latest
//go:generate go run generation/generate/wrap.go ccip PingPongDemo ping_pong_demo latest
//go:generate go run generation/generate/wrap.go ccip MessageHasher message_hasher latest
//go:generate go run generation/generate/wrap.go ccip MultiOCR3Helper multi_ocr3_helper latest
//go:generate go run generation/generate/wrap.go ccip USDCReaderTester usdc_reader_tester latest
//go:generate go run generation/generate/wrap.go ccip ReportCodec report_codec latest
//go:generate go run generation/generate/wrap.go ccip EtherSenderReceiver ether_sender_receiver latest
//go:generate go run generation/generate/wrap.go ccip MockE2EUSDCTokenMessenger mock_usdc_token_messenger latest
//go:generate go run generation/generate/wrap.go ccip MockReceiverV2 mock_receiver_v2 latest
//go:generate go run generation/generate/wrap.go ccip MockE2EUSDCTransmitter mock_usdc_token_transmitter latest
//go:generate go run generation/generate/wrap.go ccip MockE2ELBTCTokenPool mock_lbtc_token_pool latest
//go:generate go run generation/generate/wrap.go ccip CCIPReaderTester ccip_reader_tester latest

// EncodingUtils
//go:generate go run generation/generate/wrap.go ccip EncodingUtils ccip_encoding_utils latest

// Superchain Interop
//go:generate go run generation/generate/wrap.go ccip OnRampOverSuperchainInterop onramp_over_superchain_interop latest
//go:generate go run generation/generate/wrap.go ccip OffRampOverSuperchainInterop offramp_over_superchain_interop latest
