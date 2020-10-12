`export GO111MODULE=on  # Enable module mode`

`go get github.com/golang/protobuf/protoc-gen-go`

`go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0`

`export PATH="$PATH:$(go env GOPATH)/bin"`

## To install gRPC
`go get -u google.golang.org/grpc`


## To generate gRPC code:
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto`