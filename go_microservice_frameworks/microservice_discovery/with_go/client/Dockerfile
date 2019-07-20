FROM golang:1.12-alpine

# Install git and clone the Consul api sources
RUN apk update && apk upgrade && apk add --no-cache bash git && \
    go get -u github.com/hashicorp/consul/api

# Copy my sources on the Docker image
ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client/
COPY .. ${SOURCES}

RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build -o go-microservice-client

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/go-microservice-client