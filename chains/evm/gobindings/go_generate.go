package ccip

//go:generate go run generation/generate/wrap.go ccip OnRamp onramp latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip OffRamp offramp latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip Proxy proxy latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip CommitteeVerifier committee_verifier latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip Executor executor latest ../../../ccv/chains/evm/gobindings

//go:generate go run generation/generate/wrap.go ccip Router router latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip FeeQuoter fee_quoter latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip TokenAdminRegistry token_admin_registry latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip TokenPoolFactory token_pool_factory latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip FactoryBurnMintERC20 factory_burn_mint_erc20 latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip RegistryModuleOwnerCustom registry_module_owner_custom latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip RMNProxy rmn_proxy_contract latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip RMNRemote rmn_remote latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip HyperLiquidCompatibleERC20 hyper_liquid_compatible_erc20 latest ../../../ccv/chains/evm/gobindings

// Pools
//go:generate go run generation/generate/wrap.go ccip BurnMintTokenPool burn_mint_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip BurnFromMintTokenPool burn_from_mint_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip BurnWithFromMintTokenPool burn_with_from_mint_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip LockReleaseTokenPool lock_release_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip TokenPool token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip USDCTokenPool usdc_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip SiloedLockReleaseTokenPool siloed_lock_release_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip BurnToAddressMintTokenPool burn_to_address_mint_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip CCTPMessageTransmitterProxy cctp_message_transmitter_proxy latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip ERC20LockBox erc20_lock_box latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip SiloedUSDCTokenPool siloed_usdc_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip USDCTokenPoolCCTPV2 usdc_token_pool_cctp_v2 latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip USDCTokenPoolProxy usdc_token_pool_proxy latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip BurnMintWithLockReleaseFlagTokenPool burn_mint_with_lock_release_flag_token_pool latest ../../../ccv/chains/evm/gobindings

// Helpers
//go:generate go run generation/generate/wrap.go ccip MaybeRevertMessageReceiver maybe_revert_message_receiver latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip LogMessageDataReceiver log_message_data_receiver latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip PingPongDemo ping_pong_demo latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MessageHasher message_hasher latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip USDCReaderTester usdc_reader_tester latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip EtherSenderReceiver ether_sender_receiver latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MockE2EUSDCTokenMessenger mock_usdc_token_messenger latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MockReceiverV2 mock_receiver_v2 latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MockE2EUSDCTransmitter mock_usdc_token_transmitter latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MockE2ELBTCTokenPool mock_lbtc_token_pool latest ../../../ccv/chains/evm/gobindings
//go:generate go run generation/generate/wrap.go ccip MessageHasher message_hasher latest ../../../ccv/chains/evm/gobindings
