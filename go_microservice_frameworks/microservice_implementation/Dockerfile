FROM golang:1.12-alpine

# Install git and clone the Gin framework
RUN apk update && apk upgrade && apk add --no-cache bash git && \
    go get -u github.com/gin-gonic/gin

# Copy my sources on the Docker image
ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}microservice_implementation

EXPOSE 8080