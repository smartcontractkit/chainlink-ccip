package shared

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	chainsel "github.com/smartcontractkit/chain-selectors"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

// ProtoChainTypeToFamily maps JD protobuf ChainType to chain-selectors family string.
var ProtoChainTypeToFamily = map[nodev1.ChainType]string{
	nodev1.ChainType_CHAIN_TYPE_EVM:      chainsel.FamilyEVM,
	nodev1.ChainType_CHAIN_TYPE_SOLANA:   chainsel.FamilySolana,
	nodev1.ChainType_CHAIN_TYPE_STARKNET: chainsel.FamilyStarknet,
	nodev1.ChainType_CHAIN_TYPE_APTOS:    chainsel.FamilyAptos,
}

type NOPAlias string

// ChainSupportByNOP maps NOP alias to the chain selectors that node supports.
type ChainSupportByNOP map[string][]uint64

// NOPJobSpecs maps NOP alias -> job ID -> job spec TOML.
// Used for intermediate representation in sequences/operations.
// For persistence, convert to NOPJobs.
type NOPJobSpecs map[NOPAlias]map[JobID]string

// JobProposalStatus represents the state of a job proposal.
type JobProposalStatus string

const (
	JobProposalStatusPending  JobProposalStatus = "pending"
	JobProposalStatusApproved JobProposalStatus = "approved"
	JobProposalStatusRevoked  JobProposalStatus = "revoked"
	JobProposalStatusRejected JobProposalStatus = "rejected"
)

// NOPMode defines how a Node Operator runs its jobs.
type NOPMode string

const (
	NOPModeCL         NOPMode = "cl"         // Managed via JD/Chainlink node
	NOPModeStandalone NOPMode = "standalone" // Running as standalone binary
)

// ProposalRevision tracks a single proposal revision from JD.
type ProposalRevision struct {
	ProposalID string            `json:"proposalId"`
	Revision   int64             `json:"revision"`
	Status     JobProposalStatus `json:"status"`
	Spec       string            `json:"spec"`
}

// JobInfo contains all metadata about a job and its proposal history.
type JobInfo struct {
	JobID         JobID    `json:"jobId"`
	JDJobID       string   `json:"jdJobId"`
	ExternalJobID string   `json:"externalJobId"`
	NOPAlias      NOPAlias `json:"nopAlias"`
	NodeID        string   `json:"nodeId,omitempty"`
	Mode          NOPMode  `json:"mode,omitempty"`

	Spec             string `json:"spec"`
	ActiveProposalID string `json:"activeProposalId,omitempty"`

	Proposals map[string]ProposalRevision `json:"proposals,omitempty"`
}

// LatestProposal returns the proposal with the highest revision number.
// Returns nil if no proposals exist.
func (j *JobInfo) LatestProposal() *ProposalRevision {
	if len(j.Proposals) == 0 {
		return nil
	}
	var latest *ProposalRevision
	for _, p := range j.Proposals {
		if latest == nil || p.Revision > latest.Revision {
			latest = &p
		}
	}
	return latest
}

// IsRunning returns true if there's an active (approved) proposal.
func (j *JobInfo) IsRunning() bool {
	return j.ActiveProposalID != ""
}

// LatestStatus returns the status of the highest revision proposal.
// Returns empty string if no proposals exist.
func (j *JobInfo) LatestStatus() JobProposalStatus {
	latest := j.LatestProposal()
	if latest == nil {
		return ""
	}
	return latest.Status
}

// NOPJobs maps NOP alias -> job ID -> job info.
type NOPJobs map[NOPAlias]map[JobID]JobInfo

// MonitoringInput defines the monitoring configuration.
type MonitoringInput struct {
	Enabled bool
	Type    string
	Beholder BeholderInput
}

// BeholderInput defines the Beholder monitoring configuration.
type BeholderInput struct {
	InsecureConnection       bool
	CACertFile               string
	OtelExporterGRPCEndpoint string
	OtelExporterHTTPEndpoint string
	LogStreamingEnabled      bool
	MetricReaderInterval     int64
	TraceSampleRatio         float64
	TraceBatchTimeout        int64
}

type JobID string

// CCVJobNamespace is a UUID v5 namespace for generating deterministic external job IDs.
var CCVJobNamespace = uuid.MustParse("a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d")

func (id JobID) ToExternalJobID() string {
	return uuid.NewSHA1(CCVJobNamespace, []byte(id)).String()
}

type JobScope interface {
	IsJobInScope(jobID JobID) bool
}

type ExecutorJobID struct {
	NOPAlias NOPAlias
	Scope    ExecutorJobScope
}

type ExecutorJobScope struct {
	ExecutorQualifier string
}

func NewExecutorJobID(nopAlias NOPAlias, scope ExecutorJobScope) ExecutorJobID {
	return ExecutorJobID{
		NOPAlias: nopAlias,
		Scope:    scope,
	}
}

func (id ExecutorJobID) ToJobID() JobID {
	return JobID(fmt.Sprintf("%s-%s-executor", string(id.NOPAlias), id.Scope.ExecutorQualifier))
}

func (id ExecutorJobID) GetExecutorID() string {
	return string(id.NOPAlias)
}

func (scope ExecutorJobScope) IsJobInScope(jobID JobID) bool {
	return strings.HasSuffix(string(jobID), fmt.Sprintf("-%s-executor", scope.ExecutorQualifier))
}

type VerifierJobScope struct {
	CommitteeQualifier string
}

type VerifierJobID struct {
	NOPAlias           NOPAlias
	CommitteeQualifier string
	AggregatorName     string
}

func NewVerifierJobID(nopAlias NOPAlias, aggregatorName string, scope VerifierJobScope) VerifierJobID {
	return VerifierJobID{
		NOPAlias:           nopAlias,
		CommitteeQualifier: scope.CommitteeQualifier,
		AggregatorName:     aggregatorName,
	}
}

func (id VerifierJobID) GetVerifierID() string {
	return fmt.Sprintf("%s-%s-verifier", id.AggregatorName, id.CommitteeQualifier)
}

func (id VerifierJobID) ToJobID() JobID {
	return JobID(fmt.Sprintf("%s-%s-%s-verifier", string(id.NOPAlias), id.AggregatorName, id.CommitteeQualifier))
}

func (scope VerifierJobScope) IsJobInScope(jobID JobID) bool {
	return strings.HasSuffix(string(jobID), fmt.Sprintf("-%s-verifier", scope.CommitteeQualifier))
}

func ConvertNopAliasToString(aliases []NOPAlias) []string {
	str := make([]string, len(aliases))
	for i, alias := range aliases {
		str[i] = string(alias)
	}
	return str
}

func ConvertStringToNopAliases(strs []string) []NOPAlias {
	aliases := make([]NOPAlias, len(strs))
	for i, alias := range strs {
		aliases[i] = NOPAlias(alias)
	}
	return aliases
}

func IsProductionEnvironment(env string) bool {
	return env == "mainnet"
}

func NOPAliasSliceToSet(slice []NOPAlias) map[NOPAlias]bool {
	if len(slice) == 0 {
		return nil
	}
	set := make(map[NOPAlias]bool, len(slice))
	for _, v := range slice {
		set[v] = true
	}
	return set
}

func (j *JobInfo) AddProposal(p ProposalRevision) {
	if j.Proposals == nil {
		j.Proposals = make(map[string]ProposalRevision)
	}
	j.Proposals[p.ProposalID] = p
}

func (j *JobInfo) SetActiveProposal(proposalID string) {
	j.ActiveProposalID = proposalID
}
