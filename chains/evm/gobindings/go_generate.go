package ccip

//go:generate go run ./wrap ccip OnRamp onramp latest
//go:generate go run ./wrap ccip OffRamp offramp latest
//go:generate go run ./wrap ccip Proxy proxy latest
//go:generate go run ./wrap ccip CREATE2Factory create2_factory latest
//go:generate go run ./wrap ccip CommitteeVerifier committee_verifier latest
//go:generate go run ./wrap ccip CCTPVerifier cctp_verifier latest
//go:generate go run ./wrap ccip LombardVerifier lombard_verifier latest
//go:generate go run ./wrap ccip VersionedVerifierResolver versioned_verifier_resolver latest
//go:generate go run ./wrap ccip Executor executor latest

//go:generate go run ./wrap ccip Router router latest
//go:generate go run ./wrap ccip FeeQuoter fee_quoter latest
//go:generate go run ./wrap ccip TokenAdminRegistry token_admin_registry latest
//go:generate go run ./wrap ccip TokenPoolFactory token_pool_factory latest
//go:generate go run ./wrap ccip FactoryBurnMintERC20 factory_burn_mint_erc20 latest
//go:generate go run ./wrap ccip RegistryModuleOwnerCustom registry_module_owner_custom latest
//go:generate go run ./wrap ccip RMNProxy rmn_proxy_contract latest
//go:generate go run ./wrap ccip RMNRemote rmn_remote latest
//go:generate go run ./wrap ccip HyperLiquidCompatibleERC20 hyper_liquid_compatible_erc20 latest
//go:generate go run ./wrap ccip EtherSenderReceiver ether_sender_receiver latest

// Pools
//go:generate go run ./wrap ccip TokenPool token_pool latest
//go:generate go run ./wrap ccip AdvancedPoolHooks advanced_pool_hooks latest

//go:generate go run ./wrap ccip BurnMintTokenPool burn_mint_token_pool latest
//go:generate go run ./wrap ccip BurnFromMintTokenPool burn_from_mint_token_pool latest
//go:generate go run ./wrap ccip BurnWithFromMintTokenPool burn_with_from_mint_token_pool latest
//go:generate go run ./wrap ccip BurnToAddressMintTokenPool burn_to_address_mint_token_pool latest

//go:generate go run ./wrap ccip LockReleaseTokenPool lock_release_token_pool latest
//go:generate go run ./wrap ccip SiloedLockReleaseTokenPool siloed_lock_release_token_pool latest
//go:generate go run ./wrap ccip ERC20LockBox erc20_lock_box latest
//go:generate go run ./wrap ccip USDCTokenPoolProxy usdc_token_pool_proxy latest
//go:generate go run ./wrap ccip CCTPThroughCCVTokenPool cctp_through_ccv_token_pool latest
//go:generate go run ./wrap ccip CCTPMessageTransmitterProxy cctp_message_transmitter_proxy latest
//go:generate go run ./wrap ccip SiloedUSDCTokenPool siloed_usdc_token_pool latest
//go:generate go run ./wrap ccip BurnMintWithLockReleaseFlagTokenPool burn_mint_with_lock_release_flag_token_pool latest
//go:generate go run ./wrap ccip LombardTokenPool lombard_token_pool latest

// Helpers
//go:generate go run ./wrap ccip MaybeRevertMessageReceiver maybe_revert_message_receiver latest
//go:generate go run ./wrap ccip LogMessageDataReceiver log_message_data_receiver latest
//go:generate go run ./wrap ccip PingPongDemo ping_pong_demo latest
//go:generate go run ./wrap ccip MessageHasher message_hasher latest
//go:generate go run ./wrap ccip USDCReaderTester usdc_reader_tester latest
//go:generate go run ./wrap ccip MockE2EUSDCTokenMessenger mock_usdc_token_messenger latest
//go:generate go run ./wrap ccip MockReceiverV2 mock_receiver_v2 latest
//go:generate go run ./wrap ccip MockE2EUSDCTransmitter mock_usdc_token_transmitter latest
//go:generate go run ./wrap ccip MockE2ELBTCTokenPool mock_lbtc_token_pool latest
