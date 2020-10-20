FROM golang:1.15.2

ENV export GO111MODULE=on
ENV export PATH="$PATH:$(go env GOPATH)/bin"

# Set the time zone
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/panyuenlau/mygrpc-server/proto

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]

EXPOSE 50051