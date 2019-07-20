FROM golang:1.12-alpine

# Install git and clone the Consul api sources
RUN apk update && apk upgrade && apk add --no-cache bash git && \
    go get -u github.com/hashicorp/consul/api

# Copy my sources on the Docker image
ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/server/
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build -o go-microservice-server

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/go-microservice-server
