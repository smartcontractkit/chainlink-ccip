package verification

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldverification "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification"

	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

const verificationHookEnv = "hooktest"

func testEVMVerifier() *EVMContractInputsProvider {
	return &EVMContractInputsProvider{}
}

// registerMyContractV1 seeds global contract metadata for the standard test contract type/version.
func registerMyContractV1(t *testing.T) {
	t.Helper()
	solidityJSON := `{"language":"Solidity","sources":{"MyContract.sol":{"content":"pragma solidity 0.8.19; contract MyContract {}"}},"settings":{}}`
	RegisterContractMetadata(cldf.ContractType("MyContract"), semver.MustParse("1.0.0"), solidityJSON, "0x00", "MyContract")
}

func newDomainWithExplorerNetwork(t *testing.T, chainSelector uint64, explorerURL, blockExplorerType string) domain.Domain {
	t.Helper()

	dom := domain.NewDomain(t.TempDir(), "test")
	require.NoError(t, os.MkdirAll(dom.ConfigNetworksDirPath(), 0o755))

	networkYAML := fmt.Sprintf(`networks:
  - type: mainnet
    chain_selector: %d
    block_explorer:
      type: %s
      url: %q
    rpcs:
      - http_url: http://127.0.0.1:8545
`, chainSelector, blockExplorerType, explorerURL)
	require.NoError(t, os.WriteFile(dom.ConfigNetworksFilePath("networks.yaml"), []byte(networkYAML), 0o600))

	domainYAML := fmt.Sprintf(`environments:
  %s:
    network_types:
      - mainnet
`, verificationHookEnv)
	require.NoError(t, os.WriteFile(dom.ConfigDomainFilePath(), []byte(domainYAML), 0o600))

	return dom
}

// newDomainWithTwoEVMSourcifyNetworks writes domain config with two mainnet EVM networks (e.g. Ethereum
// and Polygon) sharing the same Sourcify explorer URL (typically an httptest server).
func newDomainWithTwoEVMSourcifyNetworks(t *testing.T, explorerURL string, chainSelectorA, chainSelectorB uint64) domain.Domain {
	t.Helper()

	dom := domain.NewDomain(t.TempDir(), "test")
	require.NoError(t, os.MkdirAll(dom.ConfigNetworksDirPath(), 0o755))

	networkYAML := fmt.Sprintf(`networks:
  - type: mainnet
    chain_selector: %d
    block_explorer:
      type: sourcify
      url: %q
    rpcs:
      - http_url: http://127.0.0.1:8545
  - type: mainnet
    chain_selector: %d
    block_explorer:
      type: sourcify
      url: %q
    rpcs:
      - http_url: http://127.0.0.1:8545
`, chainSelectorA, explorerURL, chainSelectorB, explorerURL)
	require.NoError(t, os.WriteFile(dom.ConfigNetworksFilePath("networks.yaml"), []byte(networkYAML), 0o600))

	domainYAML := fmt.Sprintf(`environments:
  %s:
    network_types:
      - mainnet
`, verificationHookEnv)
	require.NoError(t, os.WriteFile(dom.ConfigDomainFilePath(), []byte(domainYAML), 0o600))

	return dom
}

func writeEnvDatastoreWithRefs(t *testing.T, dom domain.Domain, refs []datastore.AddressRef) {
	t.Helper()

	envDir := dom.EnvDir(verificationHookEnv)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))

	refsJSON, err := json.Marshal(refs)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), refsJSON, 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
}

func TestVerifyDeployedContractsPostHook_Definition(t *testing.T) {
	registerMyContractV1(t)
	h := hooks.NewVerifyDeployedContractsPostHook(domain.Domain{}, testEVMVerifier())
	require.Equal(t, hooks.VerifyDeployedContractsHookName, h.Name)
	require.Equal(t, changeset.Warn, h.FailurePolicy)
	require.NotNil(t, h.Func)
}

func TestRequireVerifiedEnvContractsPreHook_Definition(t *testing.T) {
	registerMyContractV1(t)
	h := hooks.NewRequireVerifiedEnvContractsPreHook(domain.Domain{}, testEVMVerifier(), nil, []uint64{})
	require.Equal(t, hooks.RequireVerifiedEnvContractsHookName, h.Name)
	require.Equal(t, changeset.Abort, h.FailurePolicy)
	require.NotNil(t, h.Func)
}

func TestVerifyDeployed_PostHook_SkipsWhenApplyFailed(t *testing.T) {
	registerMyContractV1(t)
	h := hooks.NewVerifyDeployedContractsPostHook(domain.Domain{}, testEVMVerifier())
	err := h.Func(t.Context(), changeset.PostHookParams{
		Err: errors.New("apply failed"),
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.NoError(t, err)
}

func TestVerifyDeployed_PostHook_SkipsWhenDataStoreNil(t *testing.T) {
	registerMyContractV1(t)
	h := hooks.NewVerifyDeployedContractsPostHook(domain.Domain{}, testEVMVerifier())
	err := h.Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: nil,
		},
	})
	require.NoError(t, err)
}

func TestVerifyDeployed_PostHook_LoadNetworksError(t *testing.T) {
	registerMyContractV1(t)
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewVerifyDeployedContractsPostHook(dom, testEVMVerifier())
	err := h.Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "verify hook: load networks: failed to load network config")
}

func TestRequireVerified_PreHook_DataStoreError(t *testing.T) {
	registerMyContractV1(t)
	// datastore dir is missing, so loading it should error
	// existing datastore is needed for the pre-hook
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), nil, []uint64{})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "require verified pre-hook: load datastore: failed to load datastore: failed to read address_refs")
}

func TestRequireVerified_PreHook_LoadNetworksError(t *testing.T) {
	registerMyContractV1(t)
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), nil, nil)
	envDir := dom.EnvDir("staging")
	// now there is a datastore, but no network config, so loading networks should error
	require.NoError(t, mkdirAllAndWrite(t, envDir.AddressRefsFilePath()))
	require.NoError(t, mkdirAllAndWrite(t, envDir.ChainMetadataFilePath()))
	require.NoError(t, mkdirAllAndWrite(t, envDir.ContractMetadataFilePath()))
	require.NoError(t, mkdirAllAndWrite(t, envDir.EnvMetadataFilePath()))

	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "require verified pre-hook: load networks")
}

func TestIterateContractVerifiers_EmptyNetworks(t *testing.T) {
	registerMyContractV1(t)
	ds := datastore.NewMemoryDataStore().Seal()
	cfg := cfgnet.NewConfig(nil)
	err := hooks.IterateVerifiers(t.Context(), ds, cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			t.Fatal("step should not run")
			return nil
		},
	)
	require.NoError(t, err)
}

func TestIterateContractVerifiers_SkipsUnsupportedNetwork(t *testing.T) {
	registerMyContractV1(t)
	ds := datastore.NewMemoryDataStore().Seal()
	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chainsel.APTOS_MAINNET.Selector, // not an EVM chain
		RPCs:          []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})
	verifier := testEVMVerifier()
	err := hooks.IterateVerifiers(t.Context(), ds, cfg, logger.Test(t), "test", verifier,
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			t.Fatal("step should not run")
			return nil
		},
	)
	require.NoError(t, err)
}

func TestIterateContractVerifiers_NoAPIKeyRecordsVerifierError(t *testing.T) {
	registerMyContractV1(t)
	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chain.Selector,
		BlockExplorer: cfgnet.BlockExplorer{
			Type: "etherscan",
			URL:  "https://api.etherscan.io/v2/api",
		},
		RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})

	err := hooks.IterateVerifiers(t.Context(), mds.Seal(), cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			t.Fatal("step should not run when verifier construction fails")
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorContains(t, err, `API key not configured`)
}

func TestIterateContractVerifiers_BuildInputsError(t *testing.T) {
	RegisterContractMetadata(cldf.ContractType("BadJSONContract"), semver.MustParse("1.0.0"), "{not-json", "0x", "BadJSONContract")

	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "BadJSONContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chain.Selector,
		BlockExplorer: cfgnet.BlockExplorer{
			Type:   "etherscan",
			URL:    "https://api.etherscan.io/v2/api",
			APIKey: "k"},
		RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})

	err := hooks.IterateVerifiers(t.Context(), mds.Seal(), cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			t.Fatal("step should not run when GetInputs fails")
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorContains(t, err, "build verifier")
	require.ErrorContains(t, err, "failed to unmarshal solidity standard JSON input")
}

func TestIterateContractVerifiers_SkipsNilVersion(t *testing.T) {
	registerMyContractV1(t)
	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       nil,
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chain.Selector,
		BlockExplorer: cfgnet.BlockExplorer{
			Type:   "etherscan",
			URL:    "https://api.etherscan.io/v2/api",
			APIKey: "k",
		},
		RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})

	err := hooks.IterateVerifiers(t.Context(), mds.Seal(), cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			t.Fatal("step should not run when version is nil")
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid address ref")
}

func TestIterateContractVerifiers_StepError(t *testing.T) {
	registerMyContractV1(t)
	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chain.Selector,
		BlockExplorer: cfgnet.BlockExplorer{
			Type:   "etherscan",
			URL:    "https://api.etherscan.io/v2/api",
			APIKey: "k",
		},
		RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})

	stepErr := errors.New("step failed")
	err := hooks.IterateVerifiers(t.Context(), mds.Seal(), cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			return stepErr
		},
	)
	require.ErrorIs(t, err, stepErr)
}

func TestIterateContractVerifiers_StepSuccess(t *testing.T) {
	registerMyContractV1(t)
	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chain.Selector,
		BlockExplorer: cfgnet.BlockExplorer{
			Type:   "etherscan",
			URL:    "https://api.etherscan.io/v2/api",
			APIKey: "k",
		},
		RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})

	var stepCalls int
	err := hooks.IterateVerifiers(t.Context(), mds.Seal(), cfg, logger.Test(t), "test", testEVMVerifier(),
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			stepCalls++
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, 1, stepCalls)
}

func TestRequireVerified_PreHook_Sourcify_NotVerified(t *testing.T) {
	registerMyContractV1(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Files have not been found"))
	}))
	t.Cleanup(server.Close)

	dom := newDomainWithExplorerNetwork(t, chainsel.HEDERA_MAINNET.Selector, server.URL, "sourcify")
	writeEnvDatastoreWithRefs(t, dom, []datastore.AddressRef{{
		ChainSelector: chainsel.HEDERA_MAINNET.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}})

	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), nil, []uint64{chainsel.HEDERA_MAINNET.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "not verified on explorer")
}

func TestRequireVerified_PreHook_Sourcify_AlreadyVerified(t *testing.T) {
	registerMyContractV1(t)
	called := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "full"})
		called = true
	}))
	t.Cleanup(server.Close)

	chain, ok := chainsel.ChainBySelector(chainsel.HEDERA_MAINNET.Selector)
	require.True(t, ok)

	dom := newDomainWithExplorerNetwork(t, chain.Selector, server.URL, "sourcify")
	ref := datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}
	writeEnvDatastoreWithRefs(t, dom, []datastore.AddressRef{ref})
	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), []datastore.AddressRef{ref}, []uint64{chain.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
	require.True(t, called, "Explorer should be queried for verification status even if ref is in verified list, to avoid false positives")
	called = false
	h = hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), nil, []uint64{chain.Selector})
	err = h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
	require.True(t, called, "Explorer should be queried for verification status even if ref is in verified list, to avoid false positives")
}

func TestVerifyDeployed_PostHook_Blockscout_CallsVerifyWhenNotVerified(t *testing.T) {
	registerMyContractV1(t)
	var verifyPOSTs int
	var isVerifiedGETs int
	var txListGETs int
	var checkStatusGETs int
	// After the first successful verify POST, getabi should report the contract as verified so a
	// second hook run does not POST again (mirrors explorer state after verification).
	var verifyCompleted bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")
		switch {
		case r.Method == http.MethodGet && action == "getabi":
			isVerifiedGETs++
			w.Header().Set("Content-Type", "application/json")
			if verifyCompleted {
				// Blockscout: status "1" and non-empty result means verified (see blockscout_verifier.IsVerified).
				_ = json.NewEncoder(w).Encode(map[string]string{"status": "1", "result": `[{"type":"constructor"}]`})
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]string{"status": "0", "result": ""})
		case r.Method == http.MethodGet && action == "txlist":
			txListGETs++
			w.Header().Set("Content-Type", "application/json")
			// Return a creation tx whose input does not share the metadata bytecode prefix.
			// This makes constructor args empty and keeps the test focused on hook behavior.
			_ = json.NewEncoder(w).Encode(map[string]any{
				"status":  "1",
				"message": "OK",
				"result":  []map[string]string{{"input": "0xdeadbeef"}},
			})
		case r.Method == http.MethodPost && action == "verifysourcecode":
			verifyPOSTs++
			verifyCompleted = true
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "1",
				"message": "OK",
				"result":  "test-guid",
			})
		case r.Method == http.MethodGet && action == "checkverifystatus":
			checkStatusGETs++
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "1",
				"message": "OK",
				"result":  "Pass - Verified",
			})
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.RequestURI())
		}
	}))
	t.Cleanup(server.Close)

	chain, ok := chainsel.ChainBySelector(chainsel.ZORA_MAINNET.Selector)
	require.True(t, ok)

	dom := newDomainWithExplorerNetwork(t, chain.Selector, server.URL, "blockscout")

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chain.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))

	h := hooks.NewVerifyDeployedContractsPostHook(dom, testEVMVerifier())
	err := h.Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: mds,
		},
	})
	require.NoError(t, err)
	// Post-hook step calls IsVerified, then Verify (which calls IsVerified again), then POST verify.
	require.Equal(t, 1, verifyPOSTs, "first run should submit verification once")
	require.Equal(t, 2, isVerifiedGETs, "IsVerified: once in hook step, once inside blockscout Verify()")
	require.Equal(t, 1, txListGETs, "first run should fetch creation tx to derive constructor args")
	require.Equal(t, 1, checkStatusGETs, "first run should poll checkverifystatus once")

	err = h.Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: mds,
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, verifyPOSTs, "second run should not POST again when explorer reports verified")
	require.Equal(t, 3, isVerifiedGETs, "second run: one IsVerified in step; Verify returns early without another GET")
	require.Equal(t, 1, txListGETs, "second run should not fetch txlist again")
	require.Equal(t, 1, checkStatusGETs, "second run should not poll checkverifystatus again")
}

func TestRequireVerified_PreHook_Sourcify_ExplorerError(t *testing.T) {
	registerMyContractV1(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("something went wrong"))
	}))
	t.Cleanup(server.Close)

	dom := newDomainWithExplorerNetwork(t, chainsel.HEDERA_MAINNET.Selector, server.URL, "sourcify")
	writeEnvDatastoreWithRefs(t, dom, []datastore.AddressRef{{
		ChainSelector: chainsel.HEDERA_MAINNET.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}})

	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, testEVMVerifier(), nil, []uint64{chainsel.HEDERA_MAINNET.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to check verification status")
}

// TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_EVMSourcifyEndToEnd registers the EVM
// verifier once, requests two EVM chain selectors in the same family, and asserts a single pre-hook
// verifies contracts on both chains against a mock Sourcify server (same pattern as
// TestRequireVerified_PreHook_Sourcify_AlreadyVerified).
func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_EVMSourcifyEndToEnd(t *testing.T) {
	registerMyContractV1(t)
	hooks.ResetContractVerificationRegistryForTest()
	hooks.GetContractVerificationRegistry().Register(chainsel.FamilyEVM, testEVMVerifier())

	var sourcifyCalls int
	var mu sync.Mutex
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		sourcifyCalls++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "full"})
	}))
	t.Cleanup(server.Close)

	eth := chainsel.ETHEREUM_MAINNET
	poly := chainsel.POLYGON_MAINNET
	dom := newDomainWithTwoEVMSourcifyNetworks(t, server.URL, eth.Selector, poly.Selector)

	refs := []datastore.AddressRef{
		{
			ChainSelector: eth.Selector,
			Type:          "MyContract",
			Version:       semver.MustParse("1.0.0"),
			Address:       "0x0000000000000000000000000000000000000001",
		},
		{
			ChainSelector: poly.Selector,
			Type:          "MyContract",
			Version:       semver.MustParse("1.0.0"),
			Address:       "0x0000000000000000000000000000000000000002",
		},
	}
	writeEnvDatastoreWithRefs(t, dom, refs)

	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom,
		[]uint64{eth.Selector, poly.Selector}, nil)
	require.Len(t, preHooks, 1, "two EVM selectors share one chain family → one pre-hook")
	require.Equal(t, hooks.RequireVerifiedEnvContractsHookName, preHooks[0].Name)

	err := preHooks[0].Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, sourcifyCalls, 2, "each chain’s contract should hit the explorer")
}

// TestVerifyDeployedContractsPostHookForMultipleChainFamilies_EVMSourcifyEndToEnd uses the registry
// multifamily post-hook with two EVM selectors; contracts are already reported verified by the mock
// Sourcify server so the hook completes without submission errors.
func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_EVMSourcifyEndToEnd(t *testing.T) {
	registerMyContractV1(t)
	hooks.ResetContractVerificationRegistryForTest()
	hooks.GetContractVerificationRegistry().Register(chainsel.FamilyEVM, testEVMVerifier())

	var sourcifyCalls int
	var mu sync.Mutex
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		sourcifyCalls++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "full"})
	}))
	t.Cleanup(server.Close)

	eth := chainsel.ETHEREUM_MAINNET
	poly := chainsel.POLYGON_MAINNET
	dom := newDomainWithTwoEVMSourcifyNetworks(t, server.URL, eth.Selector, poly.Selector)

	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: eth.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000001",
	}))
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: poly.Selector,
		Type:          "MyContract",
		Version:       semver.MustParse("1.0.0"),
		Address:       "0x0000000000000000000000000000000000000002",
	}))

	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom,
		[]string{chainsel.FamilyEVM})
	require.Len(t, postHooks, 1, "two EVM selectors share one chain family → one post-hook")
	require.Equal(t, hooks.VerifyDeployedContractsHookName, postHooks[0].Name)

	err := postHooks[0].Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: verificationHookEnv, Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: mds,
		},
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, sourcifyCalls, 2, "post-hook should check verification per chain")
}

func mkdirAllAndWrite(t *testing.T, path string) error {
	t.Helper()
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))

	return os.WriteFile(path, nil, 0o600)
}
