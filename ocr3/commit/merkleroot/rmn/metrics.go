package rmn

const (
	RmnMethodObservation     = "observation"
	RmnMethodReportSignature = "report_signature"
)

type MetricsReporter interface {
	TrackRmnRequest(method string, latency float64, nodeID uint64, err string)
}

type NoopMetrics struct{}

func (n NoopMetrics) TrackRmnRequest(string, float64, uint64, string) {}
