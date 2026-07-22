module github.com/parbhat-cpp/koko-go/sdk/koko-scheduler

go 1.25.10

require (
	github.com/google/uuid v1.6.0
	github.com/parbhat-cpp/koko-go/proto v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.21.0
	google.golang.org/grpc v1.82.1
	google.golang.org/protobuf v1.36.11
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260709200747-435963d16310.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/parbhat-cpp/koko-go/proto => ../../proto
