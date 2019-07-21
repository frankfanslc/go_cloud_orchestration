FROM golang:1.12-alpine

# Install git and clone the Confluent Kafka go api sources
RUN apk update && apk upgrade && apk add --no-cache bash git &&
    go get github.com/confluentinc/confluent-kafka-go/kafka

ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/kafka/consumer/
COPY . ${SOURCES}



RUN cd ${SOURCES}consumer/ && CGO_ENABLED=0 go build -o consumer

ENV BROKER_ADDR localhost:9092

WORKDIR ${SOURCES}consumer/
CMD ${SOURCES}consumer/consumer
