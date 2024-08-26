# Monitoring, Logging, Observability, and Instrumentation

OpenTelemetry (OTEL) is an open-source framework for generating, collecting, and exporting telemetry data (traces, metrics, and logs) to help monitor applications.

As a CRIB user, we strongly recommend integrating with the existing observability stack provided by the platform team and avoiding the deployment of duplicate components in the CRIB environment. This approach helps reduce complexity in CRIB and allows for the reuse of existing platform components.

The platform team provides an `OTEL` stack running in the staging and production Ops Kubernetes clusters, which is already set up as a data source in the Grafana instance in production. Since CRIB is only running in the staging environment, we only need to focus on that.

## Monitoring

Metrics from CRIB environments deployed on the AWS EKS main staging cluster are collected by Prometheus and sent to a centralized monitoring store. You can then access and analyze these metrics using Thanos and Grafana.

- Grafana: https://grafana.ops.prod.cldev.sh/
- Thanos: https://thanos-querier.ops.prod.cldev.sh/graph

Note: For CRIB environments deployed on the local `Kind` cluster, metrics are not currently shipped, as support for `Kind` is in the alpha phase.

## Tracing

**Tempo** is Grafana's distributed tracing backend, designed to receive, store, and query traces. It integrates seamlessly with `OTEL` and Grafana for full observability.

The current data flows we foresee are:

- **Default use case:** `app -> gateway -> telemetry backend, Tempo`
- **Advanced use case:** `app -> sidecar -> gateway -> telemetry backend, Tempo`

For more details, please check the [Getting Started with Tracing / Tempo](https://smartcontract-it.atlassian.net/wiki/spaces/OBS/pages/823984555/OpenTelemetry+Collector+Getting+Started+with+Tracing+Tempo) documentation.
Also, you may want to use one of the existing [libraries](https://smartcontract-it.atlassian.net/wiki/spaces/OBS/pages/896369537/Quickstart#Tracing-Libraries).

### Standard Use Case - Use the OTLP Endpoint

To send spans to the observability infrastructure, you can point your application to one of the existing `OTLP` endpoints:

```yaml
open-telemetry-deployment.ops.stage.cldev.sh:4317 # OTLP/gRPC
open-telemetry-deployment.ops.stage.cldev.sh:4318 # HTTP/Jaeger
```

### Use OTEL Sidecar Collector

#### Default OTEL Sidecar

For simple use cases, add the `sidecar.opentelemetry.io/inject: "true"` annotation to the application pod. The [OTEL Kubernetes Operator](https://github.com/open-telemetry/opentelemetry-operator) will then automatically inject the sidecar into the application pod. Ensure that you add the annotation to the application `Pod` itself, rather than to the `Deployment` object's annotations.

#### Custom OTEL Sidecar

If you need a custom sidecar collector, follow these steps:

1. **Create a Custom Resource**: Define your custom configuration in your Helm chart.
2. **Annotate Your Deployment**: Use the sidecar injection annotation with your namespaced custom sidecar name.

The custom OTEL sidecar can then be automatically injected into your application's Kubernetes pods using annotations. For more details, please refer to the [documentation](https://smartcontract-it.atlassian.net/wiki/spaces/OBS/pages/823984555/OpenTelemetry+Collector+Getting+Started+with+Tracing+Tempo#Advanced).

### Configuring Tracing for DON

Here is an example of how to configure tracing for a DON node by overriding the default configuration in the Helm values. In this example, we assume that a sidecar is injected and that traces are sent to a local gRPC endpoint `localhost:4317` provided by the sidecar. If you prefer to send traces to a central `OTLP` endpoint, you will need to change the `CollectorTarget` configuration parameter to `open-telemetry-deployment.ops.stage.cldev.sh:4317`.

```ini
    [Tracing]
    Enabled = true
    SamplingRatio = 1.0
    CollectorTarget = 'localhost:4317'
    TLSCertPath = ''
    Mode = 'unencrypted'
```
