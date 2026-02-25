package token_verifier_config_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	onrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	dsutil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/token_verifier_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

const (
	testCCTPQualifier    = "test-cctp"
	testLombardQualifier = "test-lombard"
)

func TestBuildConfig(t *testing.T) {
	t.Run("all contracts present", func(t *testing.T) {
		// Test chain selectors
		chainSelector1 := uint64(123456789)
		chainSelector2 := uint64(987654321)
		chainSelectors := []uint64{chainSelector1, chainSelector2}

		// Test qualifiers
		cctpQualifier := testCCTPQualifier
		lombardQualifier := testLombardQualifier

		// Create test environment with mutable datastore
		env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, chainSelectors)

		// Add test data to datastore using the correct API pattern
		// OnRamp addresses
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Qualifier:     "",
			Address:       "0x1111111111111111111111111111111111111111",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Qualifier:     "",
			Address:       "0x1111111111111111111111111111111111111112",
		}))

		// RMN Remote addresses
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Qualifier:     "",
			Address:       "0x2222222222222222222222222222222222222221",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Qualifier:     "",
			Address:       "0x2222222222222222222222222222222222222222",
		}))

		// CCTP Verifier addresses (ContractType)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ContractType),
			Qualifier:     cctpQualifier,
			Address:       "0x3333333333333333333333333333333333333331",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(cctp_verifier.ContractType),
			Qualifier:     cctpQualifier,
			Address:       "0x3333333333333333333333333333333333333332",
		}))

		// CCTP Verifier Resolver addresses (ResolverType)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ResolverType),
			Qualifier:     cctpQualifier,
			Address:       "0x4444444444444444444444444444444444444441",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(cctp_verifier.ResolverType),
			Qualifier:     cctpQualifier,
			Address:       "0x4444444444444444444444444444444444444442",
		}))

		// Lombard Verifier Resolver addresses (ResolverType)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(lombard_verifier.ResolverType),
			Qualifier:     lombardQualifier,
			Address:       "0x5555555555555555555555555555555555555551",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(lombard_verifier.ResolverType),
			Qualifier:     lombardQualifier,
			Address:       "0x5555555555555555555555555555555555555552",
		}))

		// Seal the datastore and set it in the environment
		env.DataStore = ds.Seal()

		// Build the config using the BuildConfig operation
		input := token_verifier_config.BuildConfigInput{
			CCTPQualifier:    cctpQualifier,
			LombardQualifier: lombardQualifier,
			ChainSelectors:   chainSelectors,
		}

		deps := token_verifier_config.BuildConfigDeps{Env: env}

		// Call the BuildConfig operation
		ds2 := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		onRampAddresses := make(map[string]string)
		rmnRemoteAddresses := make(map[string]string)
		cctpVerifierAddresses := make(map[string]string)
		cctpVerifierResolverAddresses := make(map[string]string)
		lombardVerifierResolverAddresses := make(map[string]string)

		for _, chainSelector := range input.ChainSelectors {
			chainSelectorStr := strconv.FormatUint(chainSelector, 10)

			// Get OnRamp address
			onRampAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			onRampAddresses[chainSelectorStr] = onRampAddr

			// Get RMN Remote address
			rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			rmnRemoteAddresses[chainSelectorStr] = rmnRemoteAddr

			// Get CCTP Verifier address (ContractType) - optional
			cctpVerifierAddr, cctpVerifierErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Get CCTP Verifier Resolver address (ResolverType) - optional
			cctpVerifierResolverAddr, cctpResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Both CCTP addresses must be present or both absent
			require.Equal(t, cctpVerifierErr == nil, cctpResolverErr == nil,
				"CCTP verifier and resolver must both exist or both be absent")

			// If both CCTP addresses exist, add them to the maps
			if cctpVerifierErr == nil && cctpResolverErr == nil {
				cctpVerifierAddresses[chainSelectorStr] = cctpVerifierAddr
				cctpVerifierResolverAddresses[chainSelectorStr] = cctpVerifierResolverAddr
			}

			// Get Lombard Verifier Resolver address (ResolverType) - optional
			lombardVerifierResolverAddr, lombardResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(lombard_verifier.ResolverType),
				Qualifier: input.LombardQualifier,
			}, chainSelector, toAddress)

			// If Lombard address exists, add it to the map
			if lombardResolverErr == nil {
				lombardVerifierResolverAddresses[chainSelectorStr] = lombardVerifierResolverAddr
			}
		}

		config := &token_verifier_config.VerifierGeneratedConfig{
			OnRampAddresses:                  onRampAddresses,
			RMNRemoteAddresses:               rmnRemoteAddresses,
			CCTPVerifierAddresses:            cctpVerifierAddresses,
			CCTPVerifierResolverAddresses:    cctpVerifierResolverAddresses,
			LombardVerifierResolverAddresses: lombardVerifierResolverAddresses,
		}

		// Assertions
		require.NotNil(t, config)

		// Verify OnRamp addresses
		assert.Len(t, config.OnRampAddresses, 2)
		assert.Equal(t, "0x1111111111111111111111111111111111111111", config.OnRampAddresses["123456789"])
		assert.Equal(t, "0x1111111111111111111111111111111111111112", config.OnRampAddresses["987654321"])

		// Verify RMN Remote addresses
		assert.Len(t, config.RMNRemoteAddresses, 2)
		assert.Equal(t, "0x2222222222222222222222222222222222222221", config.RMNRemoteAddresses["123456789"])
		assert.Equal(t, "0x2222222222222222222222222222222222222222", config.RMNRemoteAddresses["987654321"])

		// Verify CCTP Verifier addresses
		assert.Len(t, config.CCTPVerifierAddresses, 2)
		assert.Equal(t, "0x3333333333333333333333333333333333333331", config.CCTPVerifierAddresses["123456789"])
		assert.Equal(t, "0x3333333333333333333333333333333333333332", config.CCTPVerifierAddresses["987654321"])

		// Verify CCTP Verifier Resolver addresses
		assert.Len(t, config.CCTPVerifierResolverAddresses, 2)
		assert.Equal(t, "0x4444444444444444444444444444444444444441", config.CCTPVerifierResolverAddresses["123456789"])
		assert.Equal(t, "0x4444444444444444444444444444444444444442", config.CCTPVerifierResolverAddresses["987654321"])

		// Verify Lombard Verifier Resolver addresses
		assert.Len(t, config.LombardVerifierResolverAddresses, 2)
		assert.Equal(t, "0x5555555555555555555555555555555555555551", config.LombardVerifierResolverAddresses["123456789"])
		assert.Equal(t, "0x5555555555555555555555555555555555555552", config.LombardVerifierResolverAddresses["987654321"])
	})

	t.Run("missing CCTP contracts on both chains", func(t *testing.T) {
		chainSelector1 := uint64(123456789)
		chainSelector2 := uint64(987654321)
		chainSelectors := []uint64{chainSelector1, chainSelector2}

		cctpQualifier := testCCTPQualifier
		lombardQualifier := testLombardQualifier

		env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, chainSelectors)

		// Add only required contracts (OnRamp and RMN Remote)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111111",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111112",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222221",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222222",
		}))

		// Add Lombard but not CCTP
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(lombard_verifier.ResolverType),
			Qualifier:     lombardQualifier,
			Address:       "0x5555555555555555555555555555555555555551",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(lombard_verifier.ResolverType),
			Qualifier:     lombardQualifier,
			Address:       "0x5555555555555555555555555555555555555552",
		}))

		env.DataStore = ds.Seal()

		input := token_verifier_config.BuildConfigInput{
			CCTPQualifier:    cctpQualifier,
			LombardQualifier: lombardQualifier,
			ChainSelectors:   chainSelectors,
		}

		deps := token_verifier_config.BuildConfigDeps{Env: env}
		ds2 := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		onRampAddresses := make(map[string]string)
		rmnRemoteAddresses := make(map[string]string)
		cctpVerifierAddresses := make(map[string]string)
		cctpVerifierResolverAddresses := make(map[string]string)
		lombardVerifierResolverAddresses := make(map[string]string)

		for _, chainSelector := range input.ChainSelectors {
			chainSelectorStr := strconv.FormatUint(chainSelector, 10)

			onRampAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			onRampAddresses[chainSelectorStr] = onRampAddr

			rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			rmnRemoteAddresses[chainSelectorStr] = rmnRemoteAddr

			cctpVerifierAddr, cctpVerifierErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			cctpVerifierResolverAddr, cctpResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			require.Equal(t, cctpVerifierErr == nil, cctpResolverErr == nil)

			if cctpVerifierErr == nil && cctpResolverErr == nil {
				cctpVerifierAddresses[chainSelectorStr] = cctpVerifierAddr
				cctpVerifierResolverAddresses[chainSelectorStr] = cctpVerifierResolverAddr
			}

			lombardVerifierResolverAddr, lombardResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(lombard_verifier.ResolverType),
				Qualifier: input.LombardQualifier,
			}, chainSelector, toAddress)

			if lombardResolverErr == nil {
				lombardVerifierResolverAddresses[chainSelectorStr] = lombardVerifierResolverAddr
			}
		}

		config := &token_verifier_config.VerifierGeneratedConfig{
			CCTPVerifierAddresses:            cctpVerifierAddresses,
			CCTPVerifierResolverAddresses:    cctpVerifierResolverAddresses,
			LombardVerifierResolverAddresses: lombardVerifierResolverAddresses,
		}

		// CCTP addresses should be empty
		assert.Empty(t, config.CCTPVerifierAddresses)
		assert.Empty(t, config.CCTPVerifierResolverAddresses)

		// Lombard addresses should be present
		assert.Len(t, config.LombardVerifierResolverAddresses, 2)
		assert.Equal(t, "0x5555555555555555555555555555555555555551", config.LombardVerifierResolverAddresses["123456789"])
	})

	t.Run("missing Lombard contracts on both chains", func(t *testing.T) {
		chainSelector1 := uint64(123456789)
		chainSelector2 := uint64(987654321)
		chainSelectors := []uint64{chainSelector1, chainSelector2}

		cctpQualifier := testCCTPQualifier
		lombardQualifier := testLombardQualifier

		env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, chainSelectors)

		// Add required contracts
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111111",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111112",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222221",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222222",
		}))

		// Add CCTP but not Lombard
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ContractType),
			Qualifier:     cctpQualifier,
			Address:       "0x3333333333333333333333333333333333333331",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(cctp_verifier.ContractType),
			Qualifier:     cctpQualifier,
			Address:       "0x3333333333333333333333333333333333333332",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ResolverType),
			Qualifier:     cctpQualifier,
			Address:       "0x4444444444444444444444444444444444444441",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector2,
			Type:          datastore.ContractType(cctp_verifier.ResolverType),
			Qualifier:     cctpQualifier,
			Address:       "0x4444444444444444444444444444444444444442",
		}))

		env.DataStore = ds.Seal()

		input := token_verifier_config.BuildConfigInput{
			CCTPQualifier:    cctpQualifier,
			LombardQualifier: lombardQualifier,
			ChainSelectors:   chainSelectors,
		}

		deps := token_verifier_config.BuildConfigDeps{Env: env}
		ds2 := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		onRampAddresses := make(map[string]string)
		rmnRemoteAddresses := make(map[string]string)
		cctpVerifierAddresses := make(map[string]string)
		cctpVerifierResolverAddresses := make(map[string]string)
		lombardVerifierResolverAddresses := make(map[string]string)

		for _, chainSelector := range input.ChainSelectors {
			chainSelectorStr := strconv.FormatUint(chainSelector, 10)

			onRampAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			onRampAddresses[chainSelectorStr] = onRampAddr

			rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)
			require.NoError(t, err)
			rmnRemoteAddresses[chainSelectorStr] = rmnRemoteAddr

			cctpVerifierAddr, cctpVerifierErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			cctpVerifierResolverAddr, cctpResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			require.Equal(t, cctpVerifierErr == nil, cctpResolverErr == nil)

			if cctpVerifierErr == nil && cctpResolverErr == nil {
				cctpVerifierAddresses[chainSelectorStr] = cctpVerifierAddr
				cctpVerifierResolverAddresses[chainSelectorStr] = cctpVerifierResolverAddr
			}

			lombardVerifierResolverAddr, lombardResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(lombard_verifier.ResolverType),
				Qualifier: input.LombardQualifier,
			}, chainSelector, toAddress)

			if lombardResolverErr == nil {
				lombardVerifierResolverAddresses[chainSelectorStr] = lombardVerifierResolverAddr
			}
		}

		config := &token_verifier_config.VerifierGeneratedConfig{
			CCTPVerifierAddresses:            cctpVerifierAddresses,
			LombardVerifierResolverAddresses: lombardVerifierResolverAddresses,
		}

		// CCTP addresses should be present
		assert.Len(t, config.CCTPVerifierAddresses, 2)
		assert.Equal(t, "0x3333333333333333333333333333333333333331", config.CCTPVerifierAddresses["123456789"])

		// Lombard addresses should be empty
		assert.Empty(t, config.LombardVerifierResolverAddresses)
	})

	t.Run("error when only CCTP verifier present without resolver", func(t *testing.T) {
		chainSelector1 := uint64(123456789)
		chainSelectors := []uint64{chainSelector1}

		cctpQualifier := testCCTPQualifier

		env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, chainSelectors)

		// Add required contracts
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111111",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222221",
		}))

		// Add CCTP Verifier but NOT the resolver (inconsistent state)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ContractType),
			Qualifier:     cctpQualifier,
			Address:       "0x3333333333333333333333333333333333333331",
		}))

		env.DataStore = ds.Seal()

		input := token_verifier_config.BuildConfigInput{
			CCTPQualifier:  cctpQualifier,
			ChainSelectors: chainSelectors,
		}

		deps := token_verifier_config.BuildConfigDeps{Env: env}
		ds2 := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		for _, chainSelector := range input.ChainSelectors {
			// Skip required contracts
			_, _ = dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)

			_, _ = dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)

			// Test CCTP contracts - should be inconsistent
			_, cctpVerifierErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			_, cctpResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Verify they are inconsistent
			assert.NotEqual(t, cctpVerifierErr == nil, cctpResolverErr == nil,
				"Expected CCTP verifier and resolver to be inconsistent")
			assert.NoError(t, cctpVerifierErr, "CCTP verifier should exist")
			assert.Error(t, cctpResolverErr, "CCTP resolver should not exist")
		}
	})

	t.Run("error when only CCTP resolver present without verifier", func(t *testing.T) {
		chainSelector1 := uint64(123456789)
		chainSelectors := []uint64{chainSelector1}

		cctpQualifier := testCCTPQualifier

		env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, chainSelectors)

		// Add required contracts
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       "0x1111111111111111111111111111111111111111",
		}))

		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       "0x2222222222222222222222222222222222222221",
		}))

		// Add CCTP Resolver but NOT the verifier (inconsistent state)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector1,
			Type:          datastore.ContractType(cctp_verifier.ResolverType),
			Qualifier:     cctpQualifier,
			Address:       "0x4444444444444444444444444444444444444441",
		}))

		env.DataStore = ds.Seal()

		input := token_verifier_config.BuildConfigInput{
			CCTPQualifier:  cctpQualifier,
			ChainSelectors: chainSelectors,
		}

		deps := token_verifier_config.BuildConfigDeps{Env: env}
		ds2 := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		for _, chainSelector := range input.ChainSelectors {
			// Skip required contracts
			_, _ = dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)

			_, _ = dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)

			// Test CCTP contracts - should be inconsistent
			_, cctpVerifierErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			_, cctpResolverErr := dsutil.FindAndFormatRef(ds2, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Verify they are inconsistent
			assert.NotEqual(t, cctpVerifierErr == nil, cctpResolverErr == nil,
				"Expected CCTP verifier and resolver to be inconsistent")
			assert.Error(t, cctpVerifierErr, "CCTP verifier should not exist")
			assert.NoError(t, cctpResolverErr, "CCTP resolver should exist")
		}
	})
}
