module authentication

go 1.21.1

replace github.com/Codelax/scaleway-connect-go-helpers => ../../

require (
	buf.build/gen/go/scaleway/scaleway-apis/connectrpc/go v1.13.0-20231218112748-02e581403d9b.1
	connectrpc.com/connect v1.13.0
	github.com/Codelax/scaleway-connect-go-helpers v0.0.0-00010101000000-000000000000
)

require (
	buf.build/gen/go/scaleway/scaleway-apis/protocolbuffers/go v1.31.0-20231218112748-02e581403d9b.2 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.22 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
