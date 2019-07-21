# Microservice communication

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implement an RPC client and server system with the ProtoBuf binary protocol](#2-implement-an-rpc-client-and-server-system-with-the-protobuf-binary-protocol)
3. [Implement async message queuing with RabbitMQ](#3-implement-async-message-queuing-with-rabbitmq)
4. [Implement publish-subscribe Kafka clients](#4-implement-publish-subscribe-kafka-clients)

## 1 Document objective

In this block we are going to:

* Implement an RPC server system with the ProtoBuf binary protocol
* Implement an RPC client connecting to this server
* Implement async message queuing with RabbitMQ
* Implement publish-subscribe Kafka clients

## 2 Implement an RPC client and server system with the ProtoBuf binary protocol

To do this, first we need to download these Golang sources:
 
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition
07:53:42 $ go get github.com/micro/go-micro
go get: warning: modules disabled by GO111MODULE=auto in GOPATH/src;
	ignoring ../../../../go.mod;
	see 'go help modules'

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition
08:00:35 $ go get golang.org/x/net/context
go get: warning: modules disabled by GO111MODULE=auto in GOPATH/src;
	ignoring ../../../../go.mod;
	see 'go help modules'
```
 
Then we need to install the protoc-gen-go application:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition/
07:34:07 $ sudo apt install golang-goprotobuf-dev
```



Then use it with the provided ProtoBuf service definition file, service.proto to implement the code:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition/
07:41:34 $ protoc --go_out=. service.proto 
```

The service.pb.go file has been created:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition
07:46:37 $ ls -rlht
total 8,0K
-rw-rw-r-- 1 arturotarin arturotarin  213 jul 21 07:40 service.proto
-rw-rw-r-- 1 arturotarin arturotarin 1,4K jul 21 07:46 service.pb.go
```

The client will instantiate the ProtoBuf proto.HelloRequest function to call the server. Let's start both machines:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf
19:14:16 $ docker-compose build
Building protobuf-server
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git
 ---> Using cache
 ---> 2d7c606916eb
Step 3/8 : RUN go get -u github.com/micro/micro &&     go get github.com/micro/protobuf/proto &&     go get -u github.com/micro/protobuf/protoc-gen-go
 ---> Using cache
 ---> 355a843cbfca
Step 4/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/
 ---> Using cache
 ---> a8b68a89bf66
Step 5/8 : COPY . ${SOURCES}
 ---> b770af5ceda7
Removing intermediate container 51c58bc4f840
Step 6/8 : RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build -o server
 ---> Running in 79bcaa257dbb
 ---> 6671a0a08229
Removing intermediate container 79bcaa257dbb
Step 7/8 : WORKDIR ${SOURCES}server/
 ---> 1e4018f35a47
Removing intermediate container 034714e685d2
Step 8/8 : CMD ${SOURCES}server/server
 ---> Running in db46f8fbe190
 ---> 0e41c0601a47
Removing intermediate container db46f8fbe190
Successfully built 0e41c0601a47
Successfully tagged protobuf-server:1.0.0
Building protobuf-client
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git
 ---> Using cache
 ---> 2d7c606916eb
Step 3/8 : RUN go get -u github.com/micro/micro &&     go get github.com/micro/protobuf/proto &&     go get -u github.com/micro/protobuf/protoc-gen-go &&     go get github.com/micro/go-plugins/wrapper/breaker/hystrix &&     go get github.com/afex/hystrix-go/hystrix
 ---> Using cache
 ---> 7b28b3e8b685
Step 4/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/
 ---> Using cache
 ---> 7f254b2ac6a0
Step 5/8 : COPY . ${SOURCES}
 ---> a1c95ba9ae30
Removing intermediate container c2a61e4a7113
Step 6/8 : RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build -o client
 ---> Running in ef6033893f37
 ---> 2d582a280337
Removing intermediate container ef6033893f37
Step 7/8 : WORKDIR ${SOURCES}client/
 ---> a16dd6b38b98
Removing intermediate container 91d050d631dd
Step 8/8 : CMD ${SOURCES}client/client
 ---> Running in 431e12a3a9c0
 ---> 271e4497d5b4
Removing intermediate container 431e12a3a9c0
Successfully built 271e4497d5b4
Successfully tagged protobuf-client:1.0.0
```

Run the machines and see the logs:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf
19:15:29 $ docker-compose up --remove-orphans
Creating protobuf_protobuf-server_1 ... done
Creating protobuf_protobuf-client_1 ... done
Attaching to protobuf_protobuf-server_1, protobuf_protobuf-client_1
protobuf-server_1  | 2019/07/21 17:16:26 Transport [http] Listening on [::]:44261
protobuf-server_1  | 2019/07/21 17:16:26 Broker [http] Connected to [::]:42911
protobuf-server_1  | 2019/07/21 17:16:26 Registry [mdns] Registering node: greeter-a4ccac44-f609-4ff8-9d38-4d23d7d22f09
protobuf-server_1  | Responding with Hello Arturo, calling at 2019-07-21 17:16:35.442786841 +0000 UTC m=+3.002971305
protobuf-client_1  | Hello Arturo, calling at 2019-07-21 17:16:35.442786841 +0000 UTC m=+3.002971305
protobuf-server_1  | Responding with Hello Arturo, calling at 2019-07-21 17:16:38.442872465 +0000 UTC m=+6.003056862
protobuf-client_1  | Hello Arturo, calling at 2019-07-21 17:16:38.442872465 +0000 UTC m=+6.003056862
protobuf-server_1  | Responding with Hello Arturo, calling at 2019-07-21 17:16:41.442781763 +0000 UTC m=+9.002966256
protobuf-client_1  | Hello Arturo, calling at 2019-07-21 17:16:41.442781763 +0000 UTC m=+9.002966256
protobuf-server_1  | Responding with Hello Arturo, calling at 2019-07-21 17:16:44.442778208 +0000 UTC m=+12.002962630
protobuf-client_1  | Hello Arturo, calling at 2019-07-21 17:16:44.442778208 +0000 UTC m=+12.002962630
^CGracefully stopping... (press Ctrl+C again to force)
Stopping protobuf_protobuf-client_1   ... done
Stopping protobuf_protobuf-server_1   ... done
```

## 3 Implement async message queuing with RabbitMQ

## 4 Implement publish-subscribe Kafka clients