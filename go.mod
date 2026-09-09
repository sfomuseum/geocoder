module github.com/sfomuseum/geocoder

go 1.27.1

require (
	github.com/aaronland/go-http-maps/v2 v2.5.1
	github.com/aaronland/go-http/v4 v4.1.0
	github.com/aaronland/go-json-query v0.3.0
	github.com/aaronland/go-pagination v0.3.0
	github.com/aaronland/go-roster v1.0.0
	github.com/adrg/strutil v0.3.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgraph-io/ristretto/v2 v2.4.2
	github.com/netascode/xmldot v0.5.1
	github.com/parquet-go/parquet-go v0.32.0
	github.com/paulmach/orb v0.13.0
	github.com/sfomuseum/go-csvdict/v2 v2.0.1
	github.com/sfomuseum/go-database v0.0.22
	github.com/sfomuseum/go-edtf v1.3.0
	github.com/sfomuseum/go-embeddings v0.6.1
	github.com/sfomuseum/go-flags v0.12.1
	github.com/sfomuseum/go-parquet v0.3.1
	github.com/tidwall/gjson v1.19.0
	github.com/whosonfirst/go-openpois v0.0.1
	github.com/whosonfirst/go-reader/v2 v2.1.0
	github.com/whosonfirst/go-rfc-5646 v0.1.0
	github.com/whosonfirst/go-whosonfirst/v4 v4.13.0
	golang.org/x/text v0.42.0
	modernc.org/sqlite v1.58.0
)

require (
	cel.dev/expr v0.25.2 // indirect
	cloud.google.com/go v0.123.0 // indirect
	cloud.google.com/go/auth v0.18.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	cloud.google.com/go/iam v1.5.3 // indirect
	cloud.google.com/go/monitoring v1.24.3 // indirect
	cloud.google.com/go/storage v1.61.3 // indirect
	dario.cat/mergo v1.0.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.21.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.6.4 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.33.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.55.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.55.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/RoaringBitmap/roaring v1.5.0 // indirect
	github.com/aaronland/go-artisanal-integers v0.9.1 // indirect
	github.com/aaronland/go-aws/v3 v3.7.1 // indirect
	github.com/aaronland/go-brooklynintegers-api v1.2.10 // indirect
	github.com/aaronland/go-pagination-sql v0.2.0 // indirect
	github.com/aaronland/go-pool/v2 v2.0.0 // indirect
	github.com/aaronland/go-string v1.0.0 // indirect
	github.com/aaronland/go-uid v0.5.0 // indirect
	github.com/aaronland/go-uid-artisanal v0.0.5 // indirect
	github.com/aaronland/go-uid-proxy v0.4.1 // indirect
	github.com/aaronland/go-uid-whosonfirst v0.0.7 // indirect
	github.com/aaronland/gocloud v1.3.2 // indirect
	github.com/akrylysov/algnhsa v1.1.0 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/aws/aws-lambda-go v1.54.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.43.3 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.16 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.32.33 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.32 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.20.35 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.8.35 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.33 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager v0.2.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.36.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.62.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.32.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.56.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.106.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.73.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.33.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.38.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.45.2 // indirect
	github.com/aws/smithy-go v1.27.6 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.6.3 // indirect
	github.com/cncf/xds/go v0.0.0-20260202195803-dba9d589def2 // indirect
	github.com/cyphar/filepath-securejoin v0.6.1 // indirect
	github.com/dhconnelly/rtreego v1.2.0 // indirect
	github.com/dominikbraun/graph v0.23.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ebitengine/purego v0.10.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/envoyproxy/go-control-plane/envoy v1.37.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.3.3 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/g8rswimmer/error-chain v1.0.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.9.0 // indirect
	github.com/go-git/go-git/v5 v5.19.2 // indirect
	github.com/go-jose/go-jose/v4 v4.1.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/google/wire v0.7.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.14 // indirect
	github.com/googleapis/gax-go/v2 v2.19.0 // indirect
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.72 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.8.6 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-version v1.8.0 // indirect
	github.com/hybridgroup/yzma v1.26.1 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtacoma/uritemplates v1.0.0 // indirect
	github.com/jupiterrider/ffi v0.7.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.18.7 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/parquet-go/bitpack v1.0.0 // indirect
	github.com/parquet-go/jsonlite v1.0.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paulmach/protoscan v0.2.1 // indirect
	github.com/peterstace/simplefeatures v0.59.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pjbgf/sha1cd v0.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/protomaps/go-pmtiles v1.31.2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/schollz/progressbar/v3 v3.13.1 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/sfomuseum/go-encoderfile v0.0.1 // indirect
	github.com/sfomuseum/go-mobileclip v0.1.3 // indirect
	github.com/skeema/knownhosts v1.3.1 // indirect
	github.com/spiffe/go-spiffe/v2 v2.7.0 // indirect
	github.com/tidwall/geoindex v1.4.4 // indirect
	github.com/tidwall/geojson v1.4.6 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/rtree v1.3.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/twpayne/go-geom v1.6.1 // indirect
	github.com/ulikunitz/xz v0.5.15 // indirect
	github.com/whosonfirst/go-ioutil v1.0.2 // indirect
	github.com/whosonfirst/go-sanitize v0.1.0 // indirect
	github.com/whosonfirst/go-whosonfirst-sources v0.1.0 // indirect
	github.com/whosonfirst/go-whosonfirst-uri v1.3.0 // indirect
	github.com/whosonfirst/go-writer/v3 v3.1.1 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	go.mongodb.org/mongo-driver/v2 v2.5.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.44.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.67.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.67.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/sdk v1.44.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.uber.org/ratelimit v0.3.1 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	gocloud.dev v0.46.0 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/oauth2 v0.36.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/term v0.44.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/api v0.272.0 // indirect
	google.golang.org/genproto v0.0.0-20260316180232-0b37fe3546d5 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/grpc v1.83.1 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/ini.v1 v1.67.3 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	modernc.org/libc v1.75.6 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.12.1 // indirect
	zombiezen.com/go/sqlite v1.1.2 // indirect
)
