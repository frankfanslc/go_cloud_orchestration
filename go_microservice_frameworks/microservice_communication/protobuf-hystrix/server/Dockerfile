FROM golang:1.12-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get -u github.com/micro/micro && \
    go get github.com/micro/protobuf/proto && \
    go get -u github.com/micro/protobuf/protoc-gen-go

ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/server
