module github.com/smartcontractkit/chainlink-ccip/chains/solana

go 1.24.2

replace (
	github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings => ./gobindings
	github.com/smartcontractkit/chainlink-ccip/ocr3 => ../../ocr3
)

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/gagliardetto/binary v0.8.0
	github.com/gagliardetto/solana-go v1.13.0
	github.com/pelletier/go-toml/v2 v2.2.4
	github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings v0.0.0-00010101000000-000000000000
	github.com/smartcontractkit/chainlink-ccip/ocr3 v0.0.0-00010101000000-000000000000
	github.com/smartcontractkit/chainlink-common v0.9.1-0.20250813154823-1652ea0f0500
	github.com/stretchr/testify v1.10.0
	github.com/test-go/testify v1.1.4
	golang.org/x/crypto v0.40.0
	golang.org/x/exp v0.0.0-20250711185948-6ae5c78190dc
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cloudevents/sdk-go/binding/format/protobuf/v2 v2.16.1 // indirect
	github.com/cloudevents/sdk-go/v2 v2.16.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gagliardetto/treeout v0.1.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/smartcontractkit/chain-selectors v1.0.62 // indirect
	github.com/smartcontractkit/chainlink-common/pkg/chipingress v0.0.1 // indirect
	github.com/smartcontractkit/freeport v0.1.1 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.13.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.34.0 // indirect
	go.opentelemetry.io/otel/log v0.13.0 // indirect
	go.opentelemetry.io/otel/sdk v1.37.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.13.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.37.0 // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/grpc v1.73.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/BurntSushi/toml v1.4.0
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.63.0 // indirect
	github.com/prometheus/procfs v0.16.0 // indirect
	github.com/smartcontractkit/libocr v0.0.0-20250408131511-c90716988ee0 // indirect
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091 // indirect
	go.mongodb.org/mongo-driver v1.12.2 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/term v0.33.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250519155744-55703ea1f237 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
