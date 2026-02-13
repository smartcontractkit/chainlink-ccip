// Package jd provides a typed orchestrator for the Job Distributor gRPC service.
// It manages its own gRPC connection and exposes read-only query methods for
// fetching node operator data (nodes + chain configs).
package jd

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

// JDConfig holds the configuration for connecting to a Job Distributor instance.
type JDConfig struct {
	GRPCURL string        // gRPC endpoint address (e.g. "jd.example.com:443")
	TLS     bool          // Whether to use TLS for the connection
	Auth    *JDAuthConfig // Cognito auth config; nil = no auth (insecure)
}

// JDAuthConfig holds Cognito OAuth2 credentials for JD authentication.
type JDAuthConfig struct {
	CognitoClientID     string
	CognitoClientSecret string
	Username            string
	Password            string
	AWSRegion           string
}

// JDOrchestrator manages a gRPC connection to the Job Distributor and provides
// read-only query methods for node operator data.
type JDOrchestrator struct {
	nodeClient nodev1.NodeServiceClient
	jobClient  jobv1.JobServiceClient
	conn       *grpc.ClientConn
}

// NewJDOrchestrator creates a new JD orchestrator with a gRPC connection.
// If auth is configured, it sets up a Cognito token source and attaches
// a bearer-token interceptor to all outgoing RPCs.
func NewJDOrchestrator(ctx context.Context, cfg JDConfig) (*JDOrchestrator, error) {
	var opts []grpc.DialOption
	var interceptors []grpc.UnaryClientInterceptor

	// TLS or insecure transport.
	if cfg.TLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Auth interceptor (Cognito OAuth2).
	if cfg.Auth != nil {
		tokenSource, err := NewCognitoTokenSource(ctx,
			cfg.Auth.CognitoClientID,
			cfg.Auth.CognitoClientSecret,
			cfg.Auth.Username,
			cfg.Auth.Password,
			cfg.Auth.AWSRegion,
		)
		if err != nil {
			return nil, fmt.Errorf("jd: failed to create cognito token source: %w", err)
		}
		interceptors = append(interceptors, authTokenInterceptor(tokenSource))
	}

	if len(interceptors) > 0 {
		opts = append(opts, grpc.WithChainUnaryInterceptor(interceptors...))
	}

	// Raise the default max receive message size (4 MB) to 50 MB to handle
	// large responses from ListProposals, ListJobs, etc.
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(50*1024*1024),
	))

	conn, err := grpc.NewClient(cfg.GRPCURL, opts...)
	if err != nil {
		return nil, fmt.Errorf("jd: failed to connect to %s: %w", cfg.GRPCURL, err)
	}

	return &JDOrchestrator{
		nodeClient: nodev1.NewNodeServiceClient(conn),
		jobClient:  jobv1.NewJobServiceClient(conn),
		conn:       conn,
	}, nil
}

// Close closes the underlying gRPC connection.
func (j *JDOrchestrator) Close() error {
	if j.conn != nil {
		return j.conn.Close()
	}
	return nil
}

// ListNodes fetches all nodes from the JD and returns them as JSON-serializable maps.
// Each map contains: id, name, publicKey (CSA), isEnabled, isConnected, labels, version.
func (j *JDOrchestrator) ListNodes(ctx context.Context) ([]map[string]any, error) {
	resp, err := j.nodeClient.ListNodes(ctx, &nodev1.ListNodesRequest{})
	if err != nil {
		return nil, fmt.Errorf("jd: ListNodes failed: %w", err)
	}

	nodes := make([]map[string]any, 0, len(resp.GetNodes()))
	for _, n := range resp.GetNodes() {
		labels := make([]map[string]any, 0, len(n.GetLabels()))
		for _, l := range n.GetLabels() {
			label := map[string]any{"key": l.GetKey()}
			if l.Value != nil {
				label["value"] = l.GetValue()
			}
			labels = append(labels, label)
		}

		node := map[string]any{
			"id":          n.GetId(),
			"name":        n.GetName(),
			"publicKey":   n.GetPublicKey(),
			"isEnabled":   n.GetIsEnabled(),
			"isConnected": n.GetIsConnected(),
			"labels":      labels,
			"version":     n.GetVersion(),
		}

		// Include workflow key if present.
		if wk := n.GetWorkflowKey(); wk != "" {
			node["workflowKey"] = wk
		}

		// Include NOP friendly name if present.
		if nfn := n.GetNopFriendlyName(); nfn != "" {
			node["nopFriendlyName"] = nfn
		}

		// Include P2P key bundles if present.
		if bundles := n.GetP2PKeyBundles(); len(bundles) > 0 {
			p2pBundles := make([]map[string]any, 0, len(bundles))
			for _, b := range bundles {
				p2pBundles = append(p2pBundles, map[string]any{
					"peerId":    b.GetPeerId(),
					"publicKey": b.GetPublicKey(),
				})
			}
			node["p2pKeyBundles"] = p2pBundles
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

// ListNodeChainConfigs fetches chain configurations for the given node IDs.
// Returns a map from nodeID to a slice of chain config maps.
// Each chain config map contains: chainId, chainType, accountAddress,
// adminAddress, and ocr2Config (if present).
func (j *JDOrchestrator) ListNodeChainConfigs(ctx context.Context, nodeIDs []string) (map[string][]map[string]any, error) {
	resp, err := j.nodeClient.ListNodeChainConfigs(ctx, &nodev1.ListNodeChainConfigsRequest{
		Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeIds: nodeIDs,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("jd: ListNodeChainConfigs failed: %w", err)
	}

	result := make(map[string][]map[string]any)
	for _, cc := range resp.GetChainConfigs() {
		nodeID := cc.GetNodeId()

		cfg := map[string]any{
			"accountAddress": cc.GetAccountAddress(),
			"adminAddress":   cc.GetAdminAddress(),
		}

		// Chain info.
		if chain := cc.GetChain(); chain != nil {
			cfg["chainId"] = chain.GetId()
			cfg["chainType"] = friendlyChainType(chain.GetType())
		}

		// OCR2 config.
		if ocr2 := cc.GetOcr2Config(); ocr2 != nil {
			ocr2Map := map[string]any{
				"enabled":     ocr2.GetEnabled(),
				"isBootstrap": ocr2.GetIsBootstrap(),
				"multiaddr":   ocr2.GetMultiaddr(),
			}
			if fa := ocr2.GetForwarderAddress(); fa != "" {
				ocr2Map["forwarderAddress"] = fa
			}

			// P2P key bundle.
			if p2p := ocr2.GetP2PKeyBundle(); p2p != nil {
				ocr2Map["p2pKeyBundle"] = map[string]any{
					"peerId":    p2p.GetPeerId(),
					"publicKey": p2p.GetPublicKey(),
				}
			}

			// OCR key bundle.
			if okb := ocr2.GetOcrKeyBundle(); okb != nil {
				ocr2Map["ocrKeyBundle"] = map[string]any{
					"bundleId":              okb.GetBundleId(),
					"configPublicKey":       okb.GetConfigPublicKey(),
					"offchainPublicKey":     okb.GetOffchainPublicKey(),
					"onchainSigningAddress": okb.GetOnchainSigningAddress(),
				}
			}

			// Plugins.
			if plugins := ocr2.GetPlugins(); plugins != nil {
				ocr2Map["plugins"] = map[string]any{
					"commit":     plugins.GetCommit(),
					"execute":    plugins.GetExecute(),
					"median":     plugins.GetMedian(),
					"mercury":    plugins.GetMercury(),
					"rebalancer": plugins.GetRebalancer(),
				}
			}

			cfg["ocr2Config"] = ocr2Map
		}

		result[nodeID] = append(result[nodeID], cfg)
	}

	return result, nil
}

// ListJobs fetches jobs from the JD for the given node IDs.
// Returns a map from nodeID to a slice of job maps.
// Each job map contains: id, uuid, proposalIds, labels.
func (j *JDOrchestrator) ListJobs(ctx context.Context, nodeIDs []string) (map[string][]map[string]any, error) {
	resp, err := j.jobClient.ListJobs(ctx, &jobv1.ListJobsRequest{
		Filter: &jobv1.ListJobsRequest_Filter{
			NodeIds: nodeIDs,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("jd: ListJobs failed: %w", err)
	}

	result := make(map[string][]map[string]any)
	for _, job := range resp.GetJobs() {
		nodeID := job.GetNodeId()
		jobMap := map[string]any{
			"id":          job.GetId(),
			"uuid":        job.GetUuid(),
			"proposalIds": job.GetProposalIds(),
		}

		if labels := job.GetLabels(); len(labels) > 0 {
			labelsSlice := make([]map[string]any, 0, len(labels))
			for _, l := range labels {
				label := map[string]any{"key": l.GetKey()}
				if l.Value != nil {
					label["value"] = l.GetValue()
				}
				labelsSlice = append(labelsSlice, label)
			}
			jobMap["labels"] = labelsSlice
		}

		result[nodeID] = append(result[nodeID], jobMap)
	}

	return result, nil
}

// ListProposals fetches proposals for the given job IDs in batches to avoid
// exceeding the gRPC max message size.
// Returns a map from jobID to the latest approved proposal map.
// Each proposal map contains: id, revision, status, spec, jobId.
func (j *JDOrchestrator) ListProposals(ctx context.Context, jobIDs []string) (map[string]map[string]any, error) {
	const batchSize = 100

	result := make(map[string]map[string]any)

	for i := 0; i < len(jobIDs); i += batchSize {
		end := i + batchSize
		if end > len(jobIDs) {
			end = len(jobIDs)
		}
		batch := jobIDs[i:end]

		resp, err := j.jobClient.ListProposals(ctx, &jobv1.ListProposalsRequest{
			Filter: &jobv1.ListProposalsRequest_Filter{
				JobIds: batch,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("jd: ListProposals failed (batch %d-%d): %w", i, end, err)
		}

		// For each job, keep only the approved proposal with the highest revision.
		for _, p := range resp.GetProposals() {
			if p.GetStatus() != jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED {
				continue
			}

			jobID := p.GetJobId()
			proposalMap := map[string]any{
				"id":       p.GetId(),
				"revision": p.GetRevision(),
				"status":   friendlyProposalStatus(p.GetStatus()),
				"spec":     p.GetSpec(),
				"jobId":    jobID,
			}

			existing, ok := result[jobID]
			if !ok {
				result[jobID] = proposalMap
			} else {
				// Keep the one with the higher revision.
				if existingRev, _ := existing["revision"].(int64); p.GetRevision() > existingRev {
					result[jobID] = proposalMap
				}
			}
		}
	}

	return result, nil
}

// friendlyChainType converts the proto ChainType enum to a short human-readable string.
func friendlyChainType(ct nodev1.ChainType) string {
	// e.g. "CHAIN_TYPE_EVM" -> "EVM"
	s := ct.String()
	s = strings.TrimPrefix(s, "CHAIN_TYPE_")
	if s == "UNSPECIFIED" {
		return "unknown"
	}
	return s
}

// friendlyProposalStatus converts the proto ProposalStatus enum to a short string.
func friendlyProposalStatus(ps jobv1.ProposalStatus) string {
	s := ps.String()
	s = strings.TrimPrefix(s, "PROPOSAL_STATUS_")
	return strings.ToLower(s)
}
