module github.com/smartcontractkit/chainlink-ccip

go 1.21.7

require (
	github.com/deckarep/golang-set/v2 v2.6.0
	github.com/smartcontractkit/chainlink-common v0.2.1-0.20240724105851-fc66051bcb6e
	github.com/smartcontractkit/libocr v0.0.0-20240419185742-fd3cab206b2c
	github.com/stretchr/testify v1.9.0
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.24.0
	golang.org/x/sync v0.7.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/invopop/jsonschema v0.12.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replicating the replace directive on cosmos SDK
replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// until merged upstream: https://github.com/mitchellh/mapstructure/pull/343
	github.com/mitchellh/mapstructure v1.5.0 => github.com/nolag/mapstructure v1.5.2-0.20240625151721-90ea83a3f479
)
