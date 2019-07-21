# Microservice communication

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implement an RPC system with the ProtoBuf binary protocol](#2-implement-an-rpc-system-with-the-protobuf-binary-protocol)
3. [Implement resiliency in asynchronous calls using a circuit breaker](#3-implement-resiliency-in-asynchronous-calls-using-a-circuit-breaker)
4. [Implement async message queuing with RabbitMQ](#4-implement-async-message-queuing-with-rabbitmq)
5. [Implement publish-subscribe Kafka clients](#5-implement-publish-subscribe-kafka-clients)

## 1 Document objective

In this block we are going to:

* Implement an RPC system with the ProtoBuf binary protocol
* Implement resiliency in asynchronous calls using a circuit breaker
* Implement async message queuing with RabbitMQ
* Implement publish-subscribe Kafka clients

## 2 Implement an RPC system with the ProtoBuf binary protocol

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




## 3 Implement resiliency in asynchronous calls using a circuit breaker


## 4 Implement async message queuing with RabbitMQ

## 5 Implement publish-subscribe Kafka clients