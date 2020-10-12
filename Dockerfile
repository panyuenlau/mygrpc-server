FROM golang:1.15.2

ENV export GO111MODULE=on
ENV export PATH="$PATH:$(go env GOPATH)/bin"

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get -u google.golang.org/grpc

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]

EXPOSE 50051