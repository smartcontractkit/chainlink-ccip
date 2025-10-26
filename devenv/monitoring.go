package ccip

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
Loki labels.
*/
const (
	LokiCCIPMessageSentLabel       = "on-chain-sent"
	LokiExecutionStateChangedLabel = "on-chain-exec"
)

/*
This file includes common monitoring utilities that work with Loki/Prometheus/Tempo
This package should not define any particular product metrics but provide clients and common wrappers for products to use
*/

var (
	metricsServer *http.Server
	serverMutex   sync.Mutex
)

// ExposePrometheusMetricsFor temporarily exposes Prometheus endpoint so metrics can be scraped.
func ExposePrometheusMetricsFor(reg *prometheus.Registry, interval time.Duration) error {
	serverMutex.Lock()
	defer serverMutex.Unlock()
	if metricsServer != nil {
		Plog.Info().Msg("Shutting down previous metrics server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := metricsServer.Shutdown(ctx); err != nil {
			Plog.Warn().Err(err).Msg("Failed to gracefully shutdown previous metrics server")
		}
		metricsServer = nil
	}

	// Create new mux to avoid conflicts with global http.Handle and run
	mux := http.NewServeMux()
	mux.Handle("/on-chain-metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	metricsServer = &http.Server{
		Addr:    ":9112",
		Handler: mux,
	}
	go func() {
		Plog.Info().Msg("Starting new Prometheus metrics server on :9112")
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Plog.Error().Err(err).Msg("Metrics server error")
		}
	}()
	Plog.Info().Msgf("Exposing Prometheus metrics for %s seconds...", interval.String())
	time.Sleep(interval)
	return nil
}
