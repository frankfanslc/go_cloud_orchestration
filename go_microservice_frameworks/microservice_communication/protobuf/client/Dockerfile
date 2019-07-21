FROM golang:1.12-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get -u github.com/micro/micro && \
    go get github.com/micro/protobuf/proto && \
    go get -u github.com/micro/protobuf/protoc-gen-go && \
    go get github.com/micro/go-plugins/wrapper/breaker/hystrix && \
    go get github.com/afex/hystrix-go/hystrix

ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/
COPY . ${SOURCES}

RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build -o client

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client
