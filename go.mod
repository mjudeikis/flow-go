module github.com/dapperlabs/flow-go

go 1.13

require (
	cloud.google.com/go/storage v1.6.0
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/dapperlabs/flow-go/crypto v0.3.2-0.20200312195452-df4550a863b7
	github.com/dapperlabs/flow-go/integration v0.0.0-20200430042019-c68be52e6e34 // indirect
	github.com/dchest/siphash v1.2.1
	github.com/dgraph-io/badger/v2 v2.0.2
	github.com/ethereum/go-ethereum v1.9.9
	github.com/gammazero/deque v0.0.0-20200310222745-50fa758af896
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/golang-lru v0.5.3
	github.com/ipfs/go-log v0.0.1
	github.com/jrick/bitset v1.0.0
	github.com/json-iterator/go v1.1.8
	github.com/libp2p/go-libp2p v0.5.1
	github.com/libp2p/go-libp2p-core v0.3.0
	github.com/libp2p/go-libp2p-pubsub v0.2.5
	github.com/libp2p/go-libp2p-swarm v0.2.2
	github.com/libp2p/go-libp2p-transport-upgrader v0.1.1
	github.com/libp2p/go-tcp-transport v0.1.1
	github.com/magiconair/properties v1.8.1
	github.com/multiformats/go-multiaddr v0.2.0
	github.com/onflow/cadence v0.1.0
	github.com/onflow/flow/protobuf/go/flow v0.1.4-0.20200427172538-223aae761ecd
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.3.0
	github.com/rs/zerolog v1.15.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/vmihailenco/msgpack/v4 v4.3.11
	github.com/whyrusleeping/go-logging v0.0.1
	go.uber.org/atomic v1.5.1
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
	golang.org/x/tools v0.0.0-20200413161937-250b2131eb8b // indirect
	google.golang.org/api v0.18.0
	google.golang.org/grpc v1.28.0
)

replace mellium.im/sasl => github.com/mellium/sasl v0.2.1

replace github.com/dapperlabs/flow-go/crypto => ./crypto

replace github.com/dapperlabs/flow-go/protobuf => ./protobuf
