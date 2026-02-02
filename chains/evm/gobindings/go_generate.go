package ccip

//go:generate go run ./wrap ccip OnRamp onramp
//go:generate go run ./wrap ccip OffRamp offramp
//go:generate go run ./wrap ccip Proxy proxy
//go:generate go run ./wrap ccip CREATE2Factory create2_factory
//go:generate go run ./wrap ccip CommitteeVerifier committee_verifier
//go:generate go run ./wrap ccip CCTPVerifier cctp_verifier
//go:generate go run ./wrap ccip LombardVerifier lombard_verifier
//go:generate go run ./wrap ccip VersionedVerifierResolver versioned_verifier_resolver
//go:generate go run ./wrap ccip Executor executor

//go:generate go run ./wrap ccip Router router
//go:generate go run ./wrap ccip FeeQuoter fee_quoter
//go:generate go run ./wrap ccip TokenAdminRegistry token_admin_registry
//go:generate go run ./wrap ccip TokenPoolFactory token_pool_factory
//go:generate go run ./wrap ccip FactoryBurnMintERC20 factory_burn_mint_erc20
//go:generate go run ./wrap ccip RegistryModuleOwnerCustom registry_module_owner_custom
//go:generate go run ./wrap ccip RMNProxy rmn_proxy_contract
//go:generate go run ./wrap ccip RMNRemote rmn_remote
//go:generate go run ./wrap ccip HyperLiquidCompatibleERC20 hyper_liquid_compatible_erc20
//go:generate go run ./wrap ccip EtherSenderReceiver ether_sender_receiver

// Pools
//go:generate go run ./wrap ccip TokenPool token_pool
//go:generate go run ./wrap ccip AdvancedPoolHooks advanced_pool_hooks

//go:generate go run ./wrap ccip BurnMintTokenPool burn_mint_token_pool
//go:generate go run ./wrap ccip BurnFromMintTokenPool burn_from_mint_token_pool
//go:generate go run ./wrap ccip BurnWithFromMintTokenPool burn_with_from_mint_token_pool
//go:generate go run ./wrap ccip BurnToAddressMintTokenPool burn_to_address_mint_token_pool

//go:generate go run ./wrap ccip LockReleaseTokenPool lock_release_token_pool
//go:generate go run ./wrap ccip SiloedLockReleaseTokenPool siloed_lock_release_token_pool
//go:generate go run ./wrap ccip ERC20LockBox erc20_lock_box
//go:generate go run ./wrap ccip USDCTokenPoolProxy usdc_token_pool_proxy
//go:generate go run ./wrap ccip CCTPThroughCCVTokenPool cctp_through_ccv_token_pool
//go:generate go run ./wrap ccip CCTPMessageTransmitterProxy cctp_message_transmitter_proxy
//go:generate go run ./wrap ccip SiloedUSDCTokenPool siloed_usdc_token_pool
//go:generate go run ./wrap ccip BurnMintWithLockReleaseFlagTokenPool burn_mint_with_lock_release_flag_token_pool
//go:generate go run ./wrap ccip LombardTokenPool lombard_token_pool

// Helpers
//go:generate go run ./wrap ccip MaybeRevertMessageReceiver maybe_revert_message_receiver
//go:generate go run ./wrap ccip LogMessageDataReceiver log_message_data_receiver
//go:generate go run ./wrap ccip PingPongDemo ping_pong_demo
//go:generate go run ./wrap ccip MessageHasher message_hasher
//go:generate go run ./wrap ccip USDCReaderTester usdc_reader_tester
//go:generate go run ./wrap ccip MockE2EUSDCTokenMessenger mock_usdc_token_messenger
//go:generate go run ./wrap ccip MockReceiverV2 mock_receiver_v2
//go:generate go run ./wrap ccip MockE2EUSDCTransmitter mock_usdc_token_transmitter
//go:generate go run ./wrap ccip MockE2EUSDCTransmitterCCTPV2 mock_usdc_token_transmitter_v2
//go:generate go run ./wrap ccip MockE2ELBTCTokenPool mock_lbtc_token_pool
//go:generate go run ./wrap ccip MockLombardBridge mock_lombard_bridge
