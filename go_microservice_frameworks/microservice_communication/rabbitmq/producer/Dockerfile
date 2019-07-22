FROM golang:1.12-alpine

# Install git and clone the AMQP RabbitMQ Go api sources
RUN apk update && apk upgrade && apk add --no-cache bash git &&go get github.com/streadway/amqp

ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/rabbitmq/
COPY . ${SOURCES}

RUN cd ${SOURCES}producer/ && CGO_ENABLED=0 go build -o producer

ENV BROKER_ADDR amqp://guest:guest@localhost:5672/

WORKDIR ${SOURCES}producer/
CMD ${SOURCES}producer/producer
