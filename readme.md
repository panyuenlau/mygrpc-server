# Prerequisite
- Install docker for your operating system: https://docs.docker.com/get-docker/
-   Install Golang and protobuf\
    `brew install go`\
    `brew install protobuf`

# gRPC
## Install the protocol compiler plugin
`go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0`

## Generate grpc client code from .proto service definition.
At the project directory, run:\
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto`


# Docker
## Build docker image
At the project directory, run:\
`docker build -t [image-name] .`

Check to see if the image was built successfully\
`docker images`

## Run docker image
Run the docker image with port 50051 and port-forward to the host machine at port 50051. To run the image in interactive mode:\
`docker run -p 50051:50051 -it [image-name]`

The simple grpc server is now running and waiting for requests from client!