# Microservice communication

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implement an RPC client and server system with the ProtoBuf binary protocol](#2-implement-an-rpc-client-and-server-communication-system-with-the-protobuf-binary-protocol)
3. [Add synchronous Hystrix circuit breaker and monitor the RPC client and server](#3-add-synchronous-hystrix-circuit-breaker-and-monitor-the-rpc-client-and-server)
4. [Implement async message queuing with RabbitMQ](#4-implement-async-message-queuing-with-rabbitmq)
5. [Implement publish-subscribe Kafka clients](#5-implement-publish-subscribe-kafka-clients)

## 1 Document objective

In this block we are going to:

* Implement an RPC server system with the ProtoBuf binary protocol
* Implement an RPC client connecting to this server
* Implement async message queuing with RabbitMQ
* Implement publish-subscribe Kafka clients

## 2 Implement an RPC client and server communication system with the ProtoBuf binary protocol

To do this, first we need to download the ProtoBuf Golang sources:
 
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
 
Then we install the protoc-gen-go application:

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

## 3 Add synchronous Hystrix circuit breaker and monitor the RPC client and server communication

Now Let's make things more interesting: let's implement synchronous call using the Hystrix circuit breaker and add a monitoring dashboard.

The Hystrix diagram states provided by Netflix is this:

![alt text](https://raw.githubusercontent.com/wiki/Netflix/Hystrix/images/hystrix-command-flow-chart.png "Hystrix states diagram")

Build all the images:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix
20:38:27 $ docker-compose build 
consul uses an image, skipping
Building protobuf-hystrix-server
Step 1/9 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/9 : RUN apk update && apk upgrade && apk add --no-cache bash git
 ---> Using cache
 ---> 983050f56abb
Step 3/9 : RUN go get -u github.com/micro/micro &&     go get github.com/micro/protobuf/proto &&     go get -u github.com/micro/protobuf/protoc-gen-go
 ---> Using cache
 ---> adb78c434529
Step 4/9 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/
 ---> Running in 83da1679a3a7
 ---> 09bb0f9d4f02
Removing intermediate container 83da1679a3a7
Step 5/9 : COPY . ${SOURCES}
 ---> daa056c612f4
Removing intermediate container 6de7bb4db61e
Step 6/9 : RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build
 ---> Running in b571880d6d14
 ---> 42bf6f05cd56
Removing intermediate container b571880d6d14
Step 7/9 : ENV CONSUL_HTTP_ADDR localhost:8500
 ---> Running in eb727233d489
 ---> 74a32fd2ef94
Removing intermediate container eb727233d489
Step 8/9 : WORKDIR ${SOURCES}server/
 ---> 2a6a2c195bb1
Removing intermediate container 2ffabd6da0a3
Step 9/9 : CMD ${SOURCES}server/server
 ---> Running in 5d998f15772e
 ---> e7d80c961225
Removing intermediate container 5d998f15772e
Successfully built e7d80c961225
Successfully tagged protobuf-hystrix-server:1.0.0
Building protobuf-hystrix-client
Step 1/9 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/9 : RUN apk update && apk upgrade && apk add --no-cache bash git
 ---> Using cache
 ---> 983050f56abb
Step 3/9 : RUN go get -u github.com/micro/micro &&     go get github.com/micro/protobuf/proto &&     go get -u github.com/micro/protobuf/protoc-gen-go &&     go get github.com/micro/go-plugins/wrapper/breaker/hystrix &&     go get github.com/afex/hystrix-go/hystrix
 ---> Using cache
 ---> 753b723e6e1d
Step 4/9 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/
 ---> Running in 9f3074cff144
 ---> 349323257f33
Removing intermediate container 9f3074cff144
Step 5/9 : COPY . ${SOURCES}
 ---> 32e9f28284d3
Removing intermediate container 92ab937d89e5
Step 6/9 : RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build
 ---> Running in cc881d26c203
 ---> 3886a192d812
Removing intermediate container cc881d26c203
Step 7/9 : ENV CONSUL_HTTP_ADDR localhost:8500
 ---> Running in 9dcbe5946fba
 ---> 37ab8c443dc7
Removing intermediate container 9dcbe5946fba
Step 8/9 : WORKDIR ${SOURCES}client/
 ---> 0fec1ab1223f
Removing intermediate container 992a4231016f
Step 9/9 : CMD ${SOURCES}client/client
 ---> Running in 773e97d9dd32
 ---> fc3f4416b440
Removing intermediate container 773e97d9dd32
Successfully built fc3f4416b440
Successfully tagged protobuf-hystrix-client:1.0.0
hystrix-dashboard uses an image, skipping
```

Start all the machines:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix
20:52:15 $ docker-compose up
Starting protobufhystrix_consul_1            ... done
Starting protobufhystrix_hystrix-dashboard_1       ... done
Starting protobufhystrix_protobuf-hystrix-server_1 ... done
Starting protobufhystrix_protobuf-hystrix-client_1 ... done
Attaching to protobufhystrix_consul_1, protobufhystrix_hystrix-dashboard_1, protobufhystrix_protobuf-hystrix-server_1, protobufhystrix_protobuf-hystrix-client_1
consul_1                   | ==> Starting Consul agent...
protobuf-hystrix-server_1  | 2019/07/21 18:54:28 Transport [http] Listening on [::]:38757
protobuf-hystrix-server_1  | 2019/07/21 18:54:28 Broker [http] Connected to [::]:36373
hystrix-dashboard_1        | 2019-07-21 18:54:27.826  INFO 1 --- [           main] c.l.HysterixDashboardApplication         : Starting HysterixDashboardApplication v0.0.1-SNAPSHOT on 1531c6966340 with PID 1 (/hysterix-dashboard.jar started by root in /)
consul_1                   | ==> Consul agent running!
consul_1                   |            Version: 'v0.8.3'
consul_1                   |            Node ID: '70cda15c-7b05-8a5f-3972-89b85bdcf817'
consul_1                   |          Node name: '005a0d262b01'
consul_1                   |         Datacenter: 'dc1'
consul_1                   |             Server: true (bootstrap: false)
consul_1                   |        Client Addr: 0.0.0.0 (HTTP: 8500, HTTPS: -1, DNS: 8600)
consul_1                   |       Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
consul_1                   |     Gossip encrypt: false, RPC-TLS: false, TLS-Incoming: false
consul_1                   |              Atlas: <disabled>
consul_1                   | 
consul_1                   | ==> Log data will now stream in as it occurs:
protobuf-hystrix-server_1  | 2019/07/21 18:54:28 Registry [mdns] Registering node: greeter-efaafc8c-68da-4945-b548-29c89891b9ef
consul_1                   | 
consul_1                   |     2019/07/21 18:54:26 [DEBUG] Using unique ID "70cda15c-7b05-8a5f-3972-89b85bdcf817" from host as node ID
consul_1                   |     2019/07/21 18:54:26 [INFO] raft: Initial configuration (index=1): [{Suffrage:Voter ID:127.0.0.1:8300 Address:127.0.0.1:8300}]
consul_1                   |     2019/07/21 18:54:26 [INFO] raft: Node at 127.0.0.1:8300 [Follower] entering Follower state (Leader: "")
consul_1                   |     2019/07/21 18:54:26 [INFO] serf: EventMemberJoin: 005a0d262b01 127.0.0.1
consul_1                   |     2019/07/21 18:54:26 [INFO] consul: Adding LAN server 005a0d262b01 (Addr: tcp/127.0.0.1:8300) (DC: dc1)
consul_1                   |     2019/07/21 18:54:26 [INFO] serf: EventMemberJoin: 005a0d262b01.dc1 127.0.0.1
consul_1                   |     2019/07/21 18:54:26 [INFO] consul: Handled member-join event for server "005a0d262b01.dc1" in area "wan"
hystrix-dashboard_1        | 2019-07-21 18:54:27.829  INFO 1 --- [           main] c.l.HysterixDashboardApplication         : No active profile set, falling back to default profiles: default
consul_1                   |     2019/07/21 18:54:26 [WARN] raft: Heartbeat timeout from "" reached, starting election
consul_1                   |     2019/07/21 18:54:26 [INFO] raft: Node at 127.0.0.1:8300 [Candidate] entering Candidate state in term 2
hystrix-dashboard_1        | 2019-07-21 18:54:27.874  INFO 1 --- [           main] s.c.a.AnnotationConfigApplicationContext : Refreshing org.springframework.context.annotation.AnnotationConfigApplicationContext@513710c5: startup date [Sun Jul 21 18:54:27 UTC 2019]; root of context hierarchy
consul_1                   |     2019/07/21 18:54:26 [DEBUG] raft: Votes needed: 1
hystrix-dashboard_1        | 2019-07-21 18:54:28.098  INFO 1 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'configurationPropertiesRebinderAutoConfiguration' of type [class org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration$$EnhancerBySpringCGLIB$$dc8d3d20] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
consul_1                   |     2019/07/21 18:54:26 [DEBUG] raft: Vote granted from 127.0.0.1:8300 in term 2. Tally: 1
consul_1                   |     2019/07/21 18:54:26 [INFO] raft: Election won. Tally: 1
consul_1                   |     2019/07/21 18:54:26 [INFO] raft: Node at 127.0.0.1:8300 [Leader] entering Leader state
hystrix-dashboard_1        | 2019-07-21 18:54:28.289  INFO 1 --- [           main] c.l.HysterixDashboardApplication         : Started HysterixDashboardApplication in 0.721 seconds (JVM running for 1.246)
consul_1                   |     2019/07/21 18:54:26 [INFO] consul: cluster leadership acquired
consul_1                   |     2019/07/21 18:54:26 [INFO] consul: New leader elected: 005a0d262b01
consul_1                   |     2019/07/21 18:54:26 [DEBUG] consul: reset tombstone GC to index 3
consul_1                   |     2019/07/21 18:54:26 [INFO] consul: member '005a0d262b01' joined, marking health alive
consul_1                   |     2019/07/21 18:54:26 [INFO] agent: Synced service 'consul'
consul_1                   |     2019/07/21 18:54:26 [DEBUG] agent: Node info in sync
hystrix-dashboard_1        | 
hystrix-dashboard_1        |   .   ____          _            __ _ _
hystrix-dashboard_1        |  /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
hystrix-dashboard_1        | ( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
hystrix-dashboard_1        |  \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
hystrix-dashboard_1        |   '  |____| .__|_| |_|_| |_\__, | / / / /
hystrix-dashboard_1        |  =========|_|==============|___/=/_/_/_/
hystrix-dashboard_1        |  :: Spring Boot ::        (v1.3.3.RELEASE)
hystrix-dashboard_1        | 
hystrix-dashboard_1        | 2019-07-21 18:54:28.371  INFO 1 --- [           main] c.l.HysterixDashboardApplication         : No active profile set, falling back to default profiles: default
hystrix-dashboard_1        | 2019-07-21 18:54:28.384  INFO 1 --- [           main] ationConfigEmbeddedWebApplicationContext : Refreshing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@3b0a60b2: startup date [Sun Jul 21 18:54:28 UTC 2019]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@513710c5
hystrix-dashboard_1        | 2019-07-21 18:54:28.994  INFO 1 --- [           main] o.s.b.f.s.DefaultListableBeanFactory     : Overriding bean definition for bean 'beanNameViewResolver' with a different definition: replacing [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration.class]] with [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter.class]]
hystrix-dashboard_1        | 2019-07-21 18:54:29.131  INFO 1 --- [           main] o.s.cloud.context.scope.GenericScope     : BeanFactory id=a278db08-7256-320d-9dae-8b02e415eff0
hystrix-dashboard_1        | 2019-07-21 18:54:29.176  INFO 1 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration' of type [class org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration$$EnhancerBySpringCGLIB$$dc8d3d20] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
hystrix-dashboard_1        | 2019-07-21 18:54:29.471  INFO 1 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat initialized with port(s): 9002 (http)
hystrix-dashboard_1        | 2019-07-21 18:54:29.482  INFO 1 --- [           main] o.apache.catalina.core.StandardService   : Starting service Tomcat
hystrix-dashboard_1        | 2019-07-21 18:54:29.483  INFO 1 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.0.32
hystrix-dashboard_1        | 2019-07-21 18:54:29.565  INFO 1 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
hystrix-dashboard_1        | 2019-07-21 18:54:29.565  INFO 1 --- [ost-startStop-1] o.s.web.context.ContextLoader            : Root WebApplicationContext: initialization completed in 1181 ms
hystrix-dashboard_1        | 2019-07-21 18:54:29.922  INFO 1 --- [ost-startStop-1] o.s.b.c.e.ServletRegistrationBean        : Mapping servlet: 'proxyStreamServlet' to [/proxy.stream]
hystrix-dashboard_1        | 2019-07-21 18:54:29.923  INFO 1 --- [ost-startStop-1] o.s.b.c.e.ServletRegistrationBean        : Mapping servlet: 'dispatcherServlet' to [/]
hystrix-dashboard_1        | 2019-07-21 18:54:29.928  INFO 1 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'characterEncodingFilter' to: [/*]
hystrix-dashboard_1        | 2019-07-21 18:54:29.928  INFO 1 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'hiddenHttpMethodFilter' to: [/*]
hystrix-dashboard_1        | 2019-07-21 18:54:29.929  INFO 1 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'httpPutFormContentFilter' to: [/*]
hystrix-dashboard_1        | 2019-07-21 18:54:29.929  INFO 1 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'requestContextFilter' to: [/*]
hystrix-dashboard_1        | 2019-07-21 18:54:30.080  INFO 1 --- [           main] o.s.ui.freemarker.SpringTemplateLoader   : SpringTemplateLoader for FreeMarker: using resource loader [org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@3b0a60b2: startup date [Sun Jul 21 18:54:28 UTC 2019]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@513710c5] and template loader path [classpath:/templates/]
hystrix-dashboard_1        | 2019-07-21 18:54:30.081  INFO 1 --- [           main] o.s.w.s.v.f.FreeMarkerConfigurer         : ClassTemplateLoader for Spring macros added to FreeMarker configuration
hystrix-dashboard_1        | 2019-07-21 18:54:30.359  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@3b0a60b2: startup date [Sun Jul 21 18:54:28 UTC 2019]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@513710c5
hystrix-dashboard_1        | 2019-07-21 18:54:30.418  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/hystrix]}" onto public java.lang.String org.springframework.cloud.netflix.hystrix.dashboard.HystrixDashboardController.home(org.springframework.ui.Model,org.springframework.web.context.request.WebRequest)
hystrix-dashboard_1        | 2019-07-21 18:54:30.419  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/hystrix/{path}]}" onto public java.lang.String org.springframework.cloud.netflix.hystrix.dashboard.HystrixDashboardController.monitor(java.lang.String,org.springframework.ui.Model,org.springframework.web.context.request.WebRequest)
hystrix-dashboard_1        | 2019-07-21 18:54:30.421  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error],produces=[text/html]}" onto public org.springframework.web.servlet.ModelAndView org.springframework.boot.autoconfigure.web.BasicErrorController.errorHtml(javax.servlet.http.HttpServletRequest,javax.servlet.http.HttpServletResponse)
hystrix-dashboard_1        | 2019-07-21 18:54:30.422  INFO 1 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error]}" onto public org.springframework.http.ResponseEntity<java.util.Map<java.lang.String, java.lang.Object>> org.springframework.boot.autoconfigure.web.BasicErrorController.error(javax.servlet.http.HttpServletRequest)
hystrix-dashboard_1        | 2019-07-21 18:54:30.454  INFO 1 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/webjars/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
hystrix-dashboard_1        | 2019-07-21 18:54:30.455  INFO 1 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
hystrix-dashboard_1        | 2019-07-21 18:54:30.492  INFO 1 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**/favicon.ico] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
hystrix-dashboard_1        | 2019-07-21 18:54:30.639  WARN 1 --- [           main] o.s.c.n.a.ArchaiusAutoConfiguration      : No spring.application.name found, defaulting to 'application'
hystrix-dashboard_1        | 2019-07-21 18:54:30.643  WARN 1 --- [           main] c.n.c.sources.URLConfigurationSource     : No URLs will be polled as dynamic configuration sources.
hystrix-dashboard_1        | 2019-07-21 18:54:30.644  INFO 1 --- [           main] c.n.c.sources.URLConfigurationSource     : To enable URLs as dynamic configuration sources, define System property archaius.configurationSource.additionalUrls or make config.properties available on classpath.
hystrix-dashboard_1        | 2019-07-21 18:54:30.652  WARN 1 --- [           main] c.n.c.sources.URLConfigurationSource     : No URLs will be polled as dynamic configuration sources.
hystrix-dashboard_1        | 2019-07-21 18:54:30.652  INFO 1 --- [           main] c.n.c.sources.URLConfigurationSource     : To enable URLs as dynamic configuration sources, define System property archaius.configurationSource.additionalUrls or make config.properties available on classpath.
hystrix-dashboard_1        | 2019-07-21 18:54:30.710  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on startup
hystrix-dashboard_1        | 2019-07-21 18:54:30.721  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Bean with name 'configurationPropertiesRebinder' has been autodetected for JMX exposure
hystrix-dashboard_1        | 2019-07-21 18:54:30.722  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Bean with name 'refreshScope' has been autodetected for JMX exposure
hystrix-dashboard_1        | 2019-07-21 18:54:30.723  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Bean with name 'environmentManager' has been autodetected for JMX exposure
hystrix-dashboard_1        | 2019-07-21 18:54:30.727  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Located managed bean 'environmentManager': registering with JMX server as MBean [org.springframework.cloud.context.environment:name=environmentManager,type=EnvironmentManager]
hystrix-dashboard_1        | 2019-07-21 18:54:30.747  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Located managed bean 'refreshScope': registering with JMX server as MBean [org.springframework.cloud.context.scope.refresh:name=refreshScope,type=RefreshScope]
hystrix-dashboard_1        | 2019-07-21 18:54:30.756  INFO 1 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Located managed bean 'configurationPropertiesRebinder': registering with JMX server as MBean [org.springframework.cloud.context.properties:name=configurationPropertiesRebinder,context=3b0a60b2,type=ConfigurationPropertiesRebinder]
hystrix-dashboard_1        | 2019-07-21 18:54:30.867  INFO 1 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat started on port(s): 9002 (http)
hystrix-dashboard_1        | 2019-07-21 18:54:30.869  INFO 1 --- [           main] c.l.HysterixDashboardApplication         : Started HysterixDashboardApplication in 3.392 seconds (JVM running for 3.825)
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:33.461533748 +0000 UTC m=+3.004515542
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:33.461533748 +0000 UTC m=+3.004515542
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:36.461618086 +0000 UTC m=+6.004599825
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:36.461618086 +0000 UTC m=+6.004599825
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:39.461620781 +0000 UTC m=+9.004602593
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:39.461620781 +0000 UTC m=+9.004602593
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:42.461547649 +0000 UTC m=+12.004529374
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:42.461547649 +0000 UTC m=+12.004529374
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:45.461623069 +0000 UTC m=+15.004604836
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:45.461623069 +0000 UTC m=+15.004604836
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:48.461596119 +0000 UTC m=+18.004577874
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:48.461596119 +0000 UTC m=+18.004577874
consul_1                   | ==> Newer Consul version available: 1.5.2 (currently running: 0.8.3)
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:51.461619119 +0000 UTC m=+21.004600912
protobuf-hystrix-client_1  | Hello Arturo, calling at 2019-07-21 18:54:51.461619119 +0000 UTC m=+21.004600912
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:54.461567077 +0000 UTC m=+24.004548854
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:54:57.461593139 +0000 UTC m=+27.004574898
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:55:00.461601189 +0000 UTC m=+30.004582944
protobuf-hystrix-client_1  | hystrix: circuit open
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:55:06.461628569 +0000 UTC m=+36.004610367
protobuf-hystrix-client_1  | hystrix: circuit open
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:55:12.461600103 +0000 UTC m=+42.004581878
protobuf-hystrix-client_1  | hystrix: circuit open
protobuf-hystrix-client_1  | hystrix: timeout. Insert fallback logic here.
protobuf-hystrix-server_1  | Responding with Hello Arturo, calling at 2019-07-21 18:55:18.461622161 +0000 UTC m=+48.004603931
protobuf-hystrix-client_1  | hystrix: circuit open
```

Access the Hystrix dashboard

Open a browser at the following URL: http://localhost:9002/hystrix/

Add the following URL to monitor: http://protobuf-hystrix-client:8081/hystrix.stream

![alt text](images/image01.png "Hystrix config")

As you can appreciate in the docker-compose log and the images, when the server is not reacheable, Hystrix disables the access from the client to the server,

![alt text](images/image03.png "Hystrix closed")


and after some seconds, Hystrix opens the circuit again 

![alt text](images/image02.png "Hystrix open")

## 4 Implement async message queuing with RabbitMQ

In this section we've implemented a message producer and a message consumer for RabbitMQ in Go

Let's build the images 

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/rabbitmq
07:11:56 $ docker-compose build 
rabbitmq uses an image, skipping
Building rabbitmq-producer
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&go get github.com/streadway/amqp
 ---> Using cache
 ---> eff6c80ce5fe
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/rabbitmq/
 ---> Using cache
 ---> c8bf6de8f9b5
Step 4/8 : COPY . ${SOURCES}
 ---> Using cache
 ---> dbc06b745cd6
Step 5/8 : RUN cd ${SOURCES}producer/ && CGO_ENABLED=0 go build -o producer
 ---> Using cache
 ---> dbe8860a5ee8
Step 6/8 : ENV BROKER_ADDR amqp://guest:guest@localhost:5672/
 ---> Using cache
 ---> c949ef5a187d
Step 7/8 : WORKDIR ${SOURCES}producer/
 ---> Using cache
 ---> 3ff96d74762e
Step 8/8 : CMD ${SOURCES}producer/producer
 ---> Using cache
 ---> 5e5add805f3f
Successfully built 5e5add805f3f
Successfully tagged rabbitmq-producer:1.0.0
Building rabbitmq-consumer
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git && go get github.com/streadway/amqp
 ---> Using cache
 ---> c8edb457e0f7
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/rabbitmq/
 ---> Using cache
 ---> 77baf45332fa
Step 4/8 : COPY . ${SOURCES}
 ---> Using cache
 ---> f7cdcfcbdbd8
Step 5/8 : RUN cd ${SOURCES}consumer/ && CGO_ENABLED=0 go build -o consumer
 ---> Using cache
 ---> 33c987cf4942
Step 6/8 : ENV BROKER_ADDR amqp://guest:guest@localhost:5672/
 ---> Using cache
 ---> b36fc500bad0
Step 7/8 : WORKDIR ${SOURCES}consumer/
 ---> Using cache
 ---> 7eb34a5df96e
Step 8/8 : CMD ${SOURCES}consumer/consumer
 ---> Using cache
 ---> ec646f2a0070
Successfully built ec646f2a0070
Successfully tagged rabbitmq-consumer:1.0.0
```

and run the containers:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/rabbitmq
07:17:49 $ docker-compose up
Removing rabbitmq_rabbitmq_1
Recreating 8728c95a75d8_rabbitmq_rabbitmq_1 ... done
Creating rabbitmq_rabbitmq-consumer_1       ... done
Creating rabbitmq_rabbitmq-producer_1       ... done
Attaching to rabbitmq_rabbitmq_1, rabbitmq_rabbitmq-producer_1, rabbitmq_rabbitmq-consumer_1
rabbitmq_1           | 
rabbitmq_1           |               RabbitMQ 3.6.9. Copyright (C) 2007-2016 Pivotal Software, Inc.
rabbitmq_1           |   ##  ##      Licensed under the MPL.  See http://www.rabbitmq.com/
rabbitmq_1           |   ##  ##
rabbitmq_1           |   ##########  Logs: tty
rabbitmq_1           |   ######  ##        tty
rabbitmq_1           |   ##########
rabbitmq_1           |               Starting broker...
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:07 ===
rabbitmq_1           | Starting RabbitMQ 3.6.9 on Erlang 19.1
rabbitmq_1           | Copyright (C) 2007-2016 Pivotal Software, Inc.
rabbitmq_1           | Licensed under the MPL.  See http://www.rabbitmq.com/
rabbitmq-producer_1  | Producing to RabbitMQ ...
rabbitmq-consumer_1  | Consuming from RabbitMQ ...
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:07 ===
rabbitmq_1           | node           : rabbit@b3714cb7907a
rabbitmq_1           | home dir       : /var/lib/rabbitmq
rabbitmq_1           | config file(s) : /etc/rabbitmq/rabbitmq.config
rabbitmq_1           | cookie hash    : 3uGdESrhdBaEwjwXsVxCWw==
rabbitmq_1           | log            : tty
rabbitmq_1           | sasl log       : tty
rabbitmq_1           | database dir   : /var/lib/rabbitmq/mnesia/rabbit@b3714cb7907a
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           | Memory limit set to 6382MB of 15956MB total.
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           | Disk free limit set to 50MB
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           | Limiting to approx 1048476 file handles (943626 sockets)
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           | FHC read buffering:  OFF
rabbitmq_1           | FHC write buffering: ON
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           | Database directory at /var/lib/rabbitmq/mnesia/rabbit@b3714cb7907a is empty. Initialising from scratch...
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:09 ===
rabbitmq_1           |     application: mnesia
rabbitmq_1           |     exited: stopped
rabbitmq_1           |     type: temporary
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Waiting for Mnesia tables for 30000 ms, 9 retries left
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Waiting for Mnesia tables for 30000 ms, 9 retries left
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Waiting for Mnesia tables for 30000 ms, 9 retries left
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Priority queues enabled, real BQ is rabbit_variable_queue
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Starting rabbit_node_monitor
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Management plugin: using rates mode 'basic'
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | msg_store_transient: using rabbit_msg_store_ets_index to provide index
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | msg_store_persistent: using rabbit_msg_store_ets_index to provide index
rabbitmq_1           | 
rabbitmq_1           | =WARNING REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | msg_store_persistent: rebuilding indices from scratch
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:12 ===
rabbitmq_1           | Adding vhost '/'
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Creating user 'guest'
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Setting user tags for user 'guest' to [administrator]
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Setting permissions for 'guest' in '/' to '.*', '.*', '.*'
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | started TCP Listener on [::]:5672
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Management plugin started. Port: 15672
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Statistics database started.
rabbitmq_1           |  completed with 6 plugins.
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:13 ===
rabbitmq_1           | Server startup complete; 6 plugins started.
rabbitmq_1           |  * rabbitmq_management
rabbitmq_1           |  * rabbitmq_management_agent
rabbitmq_1           |  * rabbitmq_web_dispatch
rabbitmq_1           |  * amqp_client
rabbitmq_1           |  * cowboy
rabbitmq_1           |  * cowlib
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:18 ===
rabbitmq_1           | accepting AMQP connection <0.547.0> (192.168.48.3:52428 -> 192.168.48.2:5672)
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:18 ===
rabbitmq_1           | connection <0.547.0> (192.168.48.3:52428 -> 192.168.48.2:5672): user 'guest' authenticated and granted access to vhost '/'
rabbitmq-producer_1  | 2019/07/22 05:18:18 'Message #1' message successfully published in RabbitMQ
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:19 ===
rabbitmq_1           | accepting AMQP connection <0.565.0> (192.168.48.4:39232 -> 192.168.48.2:5672)
rabbitmq_1           | 
rabbitmq_1           | =INFO REPORT==== 22-Jul-2019::05:18:19 ===
rabbitmq_1           | connection <0.565.0> (192.168.48.4:39232 -> 192.168.48.2:5672): user 'guest' authenticated and granted access to vhost '/'
rabbitmq-consumer_1  | 2019/07/22 05:18:19 Waiting for messages ...
rabbitmq-consumer_1  | 2019/07/22 05:18:19 Received the 'Message #1' message
rabbitmq-producer_1  | 2019/07/22 05:18:23 'Message #2' message successfully published in RabbitMQ
rabbitmq-consumer_1  | 2019/07/22 05:18:23 Received the 'Message #2' message
rabbitmq-producer_1  | 2019/07/22 05:18:28 'Message #3' message successfully published in RabbitMQ
rabbitmq-consumer_1  | 2019/07/22 05:18:28 Received the 'Message #3' message
rabbitmq-producer_1  | 2019/07/22 05:18:33 'Message #4' message successfully published in RabbitMQ
rabbitmq-consumer_1  | 2019/07/22 05:18:33 Received the 'Message #4' message
^CGracefully stopping... (press Ctrl+C again to force)
Stopping rabbitmq_rabbitmq-producer_1       ... done
Stopping rabbitmq_rabbitmq-consumer_1       ... done
Stopping rabbitmq_rabbitmq_1                ... done
```

## 5 Implement publish-subscribe Kafka clients

In this section we've implemented a message producer and a message consumer for Kafka in Go

Let's build the images 

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/kafka
08:57:32 $ docker-compose build
zookeeper uses an image, skipping
kafka uses an image, skipping
Building kafka-producer
Step 1/8 : FROM golang:1.9
 ---> ef89ef5c42a9
Step 2/8 : RUN go get github.com/Shopify/sarama
 ---> Using cache
 ---> a4595c6bcd01
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/kafka/
 ---> Using cache
 ---> f13486f198f1
Step 4/8 : COPY . ${SOURCES}
 ---> 379db99c2710
Removing intermediate container 2c19687add5f
Step 5/8 : RUN cd ${SOURCES}producer/ && CGO_ENABLED=0 go build -o producer
 ---> Running in a3a2a8b65f1f
 ---> 7bcaa02cf4d4
Removing intermediate container a3a2a8b65f1f
Step 6/8 : ENV BROKER_ADDR localhost:9092
 ---> Running in ba94e540ed99
 ---> 3d15625bb3bf
Removing intermediate container ba94e540ed99
Step 7/8 : WORKDIR ${SOURCES}producer/
 ---> c5cfee61bf56
Removing intermediate container 871b1d8f9858
Step 8/8 : CMD ${SOURCES}producer/producer
 ---> Running in 11ee9ef4bc6e
 ---> 28eb1f4d47db
Removing intermediate container 11ee9ef4bc6e
Successfully built 28eb1f4d47db
Successfully tagged kafka-producer:1.0.0
Building kafka-consumer
Step 1/8 : FROM golang:1.9
 ---> ef89ef5c42a9
Step 2/8 : RUN go get github.com/Shopify/sarama
 ---> Using cache
 ---> a4595c6bcd01
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_communication/kafka/
 ---> Using cache
 ---> f13486f198f1
Step 4/8 : COPY . ${SOURCES}
 ---> Using cache
 ---> 379db99c2710
Step 5/8 : RUN cd ${SOURCES}consumer/ && CGO_ENABLED=0 go build -o consumer
 ---> Running in 2dec83680839
 ---> 7fce44e94494
Removing intermediate container 2dec83680839
Step 6/8 : ENV BROKER_ADDR localhost:9092
 ---> Running in c926c315a884
 ---> 89bca60789be
Removing intermediate container c926c315a884
Step 7/8 : WORKDIR ${SOURCES}consumer/
 ---> 316a90c1f0a0
Removing intermediate container c6be46502444
Step 8/8 : CMD ${SOURCES}consumer/consumer
 ---> Running in 8a6b649956e4
 ---> ad9d7d7c2351
Removing intermediate container 8a6b649956e4
Successfully built ad9d7d7c2351
Successfully tagged kafka-consumer:1.0.0
```

and run the containers. Notice the producer sending messages to a Kafka topic and the consumer consuming them:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/kafka
08:58:40 $ docker-compose up
Creating kafka_zookeeper_1 ... done
Creating kafka_kafka_1     ... done
Creating kafka_kafka-producer_1 ... done
Creating kafka_kafka-consumer_1 ... done
Attaching to kafka_zookeeper_1, kafka_kafka_1, kafka_kafka-consumer_1, kafka_kafka-producer_1
zookeeper_1       | JMX enabled by default
zookeeper_1       | Using config: /opt/zookeeper-3.4.6/bin/../conf/zoo.cfg
zookeeper_1       | 2019-07-22 07:02:05,180 [myid:] - INFO  [main:QuorumPeerConfig@103] - Reading configuration from: /opt/zookeeper-3.4.6/bin/../conf/zoo.cfg
kafka_1           | [2019-07-22 07:02:12,200] INFO KafkaConfig values: 
kafka_1           | 	advertised.host.name = null
kafka_1           | 	metric.reporters = []
kafka_1           | 	quota.producer.default = 9223372036854775807
kafka_1           | 	offsets.topic.num.partitions = 50
kafka_1           | 	log.flush.interval.messages = 9223372036854775807
kafka_1           | 	auto.create.topics.enable = true
kafka_1           | 	controller.socket.timeout.ms = 30000
kafka_1           | 	log.flush.interval.ms = null
kafka_1           | 	principal.builder.class = class org.apache.kafka.common.security.auth.DefaultPrincipalBuilder
kafka_1           | 	replica.socket.receive.buffer.bytes = 65536
kafka_1           | 	min.insync.replicas = 1
kafka_1           | 	replica.fetch.wait.max.ms = 500
kafka_1           | 	num.recovery.threads.per.data.dir = 1
kafka_1           | 	ssl.keystore.type = JKS
kafka_1           | 	default.replication.factor = 1
kafka_1           | 	ssl.truststore.password = null
kafka_1           | 	log.preallocate = false
kafka_1           | 	sasl.kerberos.principal.to.local.rules = [DEFAULT]
kafka_1           | 	fetch.purgatory.purge.interval.requests = 1000
kafka_1           | 	ssl.endpoint.identification.algorithm = null
kafka_1           | 	replica.socket.timeout.ms = 30000
kafka_1           | 	message.max.bytes = 1000012
kafka_1           | 	num.io.threads = 8
kafka_1           | 	offsets.commit.required.acks = -1
kafka_1           | 	log.flush.offset.checkpoint.interval.ms = 60000
kafka_1           | 	delete.topic.enable = false
kafka_1           | 	quota.window.size.seconds = 1
kafka_1           | 	ssl.truststore.type = JKS
kafka_1           | 	offsets.commit.timeout.ms = 5000
kafka_1           | 	quota.window.num = 11
kafka_1           | 	zookeeper.connect = zookeeper:2181
kafka_1           | 	authorizer.class.name = 
kafka_1           | 	num.replica.fetchers = 1
kafka_1           | 	log.retention.ms = null
kafka_1           | 	log.roll.jitter.hours = 0
kafka_1           | 	log.cleaner.enable = false
kafka_1           | 	offsets.load.buffer.size = 5242880
kafka_1           | 	log.cleaner.delete.retention.ms = 86400000
kafka_1           | 	ssl.client.auth = none
kafka_1           | 	controlled.shutdown.max.retries = 3
kafka_1           | 	queued.max.requests = 500
kafka_1           | 	offsets.topic.replication.factor = 3
kafka_1           | 	log.cleaner.threads = 1
kafka_1           | 	sasl.kerberos.service.name = null
kafka_1           | 	sasl.kerberos.ticket.renew.jitter = 0.05
kafka_1           | 	socket.request.max.bytes = 104857600
kafka_1           | 	ssl.trustmanager.algorithm = PKIX
kafka_1           | 	zookeeper.session.timeout.ms = 6000
kafka_1           | 	log.retention.bytes = -1
kafka_1           | 	sasl.kerberos.min.time.before.relogin = 60000
kafka_1           | 	zookeeper.set.acl = false
kafka_1           | 	connections.max.idle.ms = 600000
kafka_1           | 	offsets.retention.minutes = 1440
kafka_1           | 	replica.fetch.backoff.ms = 1000
kafka_1           | 	inter.broker.protocol.version = 0.9.0.X
kafka_1           | 	log.retention.hours = 168
kafka_1           | 	num.partitions = 1
kafka_1           | 	broker.id.generation.enable = true
kafka_1           | 	listeners = null
kafka_1           | 	ssl.provider = null
kafka_1           | 	ssl.enabled.protocols = [TLSv1.2, TLSv1.1, TLSv1]
kafka_1           | 	log.roll.ms = null
kafka_1           | 	log.flush.scheduler.interval.ms = 9223372036854775807
kafka_1           | 	ssl.cipher.suites = null
kafka_1           | 	log.index.size.max.bytes = 10485760
kafka_1           | 	ssl.keymanager.algorithm = SunX509
kafka_1           | 	security.inter.broker.protocol = PLAINTEXT
kafka_1           | 	replica.fetch.max.bytes = 1048576
kafka_1           | 	advertised.port = null
kafka_1           | 	log.cleaner.dedupe.buffer.size = 134217728
kafka_1           | 	replica.high.watermark.checkpoint.interval.ms = 5000
kafka_1           | 	log.cleaner.io.buffer.size = 524288
kafka_1           | 	sasl.kerberos.ticket.renew.window.factor = 0.8
kafka_1           | 	zookeeper.connection.timeout.ms = 6000
kafka_1           | 	controlled.shutdown.retry.backoff.ms = 5000
kafka_1           | 	log.roll.hours = 168
kafka_1           | 	log.cleanup.policy = delete
kafka_1           | 	host.name = 
kafka_1           | 	log.roll.jitter.ms = null
kafka_1           | 	max.connections.per.ip = 2147483647
kafka_1           | 	offsets.topic.segment.bytes = 104857600
kafka_1           | 	background.threads = 10
kafka_1           | 	quota.consumer.default = 9223372036854775807
kafka_1           | 	request.timeout.ms = 30000
kafka_1           | 	log.index.interval.bytes = 4096
kafka_1           | 	log.dir = /tmp/kafka-logs
kafka_1           | 	log.segment.bytes = 1073741824
kafka_1           | 	log.cleaner.backoff.ms = 15000
kafka_1           | 	offset.metadata.max.bytes = 4096
kafka_1           | 	ssl.truststore.location = null
kafka_1           | 	group.max.session.timeout.ms = 30000
kafka_1           | 	ssl.keystore.password = null
kafka_1           | 	zookeeper.sync.time.ms = 2000
kafka_1           | 	port = 9092
kafka_1           | 	log.retention.minutes = null
kafka_1           | 	log.segment.delete.delay.ms = 60000
kafka_1           | 	log.dirs = /tmp/kafka-logs
kafka_1           | 	controlled.shutdown.enable = true
kafka_1           | 	compression.type = producer
kafka_1           | 	max.connections.per.ip.overrides = 
kafka_1           | 	sasl.kerberos.kinit.cmd = /usr/bin/kinit
kafka_1           | 	log.cleaner.io.max.bytes.per.second = 1.7976931348623157E308
kafka_1           | 	auto.leader.rebalance.enable = true
kafka_1           | 	leader.imbalance.check.interval.seconds = 300
kafka_1           | 	log.cleaner.min.cleanable.ratio = 0.5
kafka_1           | 	replica.lag.time.max.ms = 10000
kafka_1           | 	num.network.threads = 3
kafka_1           | 	ssl.key.password = null
kafka_1           | 	reserved.broker.max.id = 10000
kafka_1           | 	metrics.num.samples = 2
kafka_1           | 	socket.send.buffer.bytes = 102400
kafka_1           | 	ssl.protocol = TLS
kafka_1           | 	socket.receive.buffer.bytes = 102400
kafka_1           | 	ssl.keystore.location = null
kafka_1           | 	replica.fetch.min.bytes = 1
kafka_1           | 	unclean.leader.election.enable = true
kafka_1           | 	group.min.session.timeout.ms = 6000
kafka_1           | 	log.cleaner.io.buffer.load.factor = 0.9
kafka_1           | 	offsets.retention.check.interval.ms = 600000
kafka_1           | 	producer.purgatory.purge.interval.requests = 1000
kafka_1           | 	metrics.sample.window.ms = 30000
kafka_1           | 	broker.id = -1
kafka_1           | 	offsets.topic.compression.codec = 0
kafka_1           | 	log.retention.check.interval.ms = 300000
kafka_1           | 	advertised.listeners = null
kafka_1           | 	leader.imbalance.per.broker.percentage = 10
kafka_1           |  (kafka.server.KafkaConfig)
zookeeper_1       | 2019-07-22 07:02:05,185 [myid:] - INFO  [main:DatadirCleanupManager@78] - autopurge.snapRetainCount set to 3
kafka_1           | [2019-07-22 07:02:12,250] INFO starting (kafka.server.KafkaServer)
kafka_1           | [2019-07-22 07:02:12,254] INFO Connecting to zookeeper on zookeeper:2181 (kafka.server.KafkaServer)
kafka_1           | [2019-07-22 07:02:12,263] INFO Starting ZkClient event thread. (org.I0Itec.zkclient.ZkEventThread)
kafka_1           | [2019-07-22 07:02:12,267] INFO Client environment:zookeeper.version=3.4.6-1569965, built on 02/20/2014 09:09 GMT (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,267] INFO Client environment:host.name=00eded330f35 (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,185 [myid:] - INFO  [main:DatadirCleanupManager@79] - autopurge.purgeInterval set to 1
kafka-consumer_1  | Starting synchronous Kafka subscriber...
kafka-producer_1  | Starting synchronous Kafka producer...
kafka_1           | [2019-07-22 07:02:12,267] INFO Client environment:java.version=1.8.0_72-internal (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,267] INFO Client environment:java.vendor=Oracle Corporation (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,186 [myid:] - WARN  [main:QuorumPeerMain@113] - Either no config or no quorum defined in config, running  in standalone mode
zookeeper_1       | 2019-07-22 07:02:05,186 [myid:] - INFO  [PurgeTask:DatadirCleanupManager$PurgeTask@138] - Purge task started.
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:java.home=/usr/lib/jvm/java-8-openjdk-amd64/jre (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,195 [myid:] - INFO  [PurgeTask:DatadirCleanupManager$PurgeTask@144] - Purge task completed.
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:java.class.path=:/opt/kafka/bin/../libs/javax.ws.rs-api-2.0.1.jar:/opt/kafka/bin/../libs/jetty-server-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/kafka_2.11-0.9.0.1-scaladoc.jar:/opt/kafka/bin/../libs/connect-file-0.9.0.1.jar:/opt/kafka/bin/../libs/jackson-databind-2.5.4.jar:/opt/kafka/bin/../libs/metrics-core-2.2.0.jar:/opt/kafka/bin/../libs/lz4-1.2.0.jar:/opt/kafka/bin/../libs/scala-library-2.11.7.jar:/opt/kafka/bin/../libs/argparse4j-0.5.0.jar:/opt/kafka/bin/../libs/kafka-tools-0.9.0.1.jar:/opt/kafka/bin/../libs/aopalliance-repackaged-2.4.0-b31.jar:/opt/kafka/bin/../libs/hk2-locator-2.4.0-b31.jar:/opt/kafka/bin/../libs/snappy-java-1.1.1.7.jar:/opt/kafka/bin/../libs/jersey-client-2.22.1.jar:/opt/kafka/bin/../libs/jersey-container-servlet-core-2.22.1.jar:/opt/kafka/bin/../libs/jetty-security-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/kafka_2.11-0.9.0.1-sources.jar:/opt/kafka/bin/../libs/jersey-container-servlet-2.22.1.jar:/opt/kafka/bin/../libs/zookeeper-3.4.6.jar:/opt/kafka/bin/../libs/jopt-simple-3.2.jar:/opt/kafka/bin/../libs/javax.inject-2.4.0-b31.jar:/opt/kafka/bin/../libs/javax.annotation-api-1.2.jar:/opt/kafka/bin/../libs/scala-parser-combinators_2.11-1.0.4.jar:/opt/kafka/bin/../libs/kafka_2.11-0.9.0.1-javadoc.jar:/opt/kafka/bin/../libs/javax.servlet-api-3.1.0.jar:/opt/kafka/bin/../libs/osgi-resource-locator-1.0.1.jar:/opt/kafka/bin/../libs/scala-xml_2.11-1.0.4.jar:/opt/kafka/bin/../libs/connect-api-0.9.0.1.jar:/opt/kafka/bin/../libs/jetty-io-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/javassist-3.18.1-GA.jar:/opt/kafka/bin/../libs/jackson-core-2.5.4.jar:/opt/kafka/bin/../libs/hk2-api-2.4.0-b31.jar:/opt/kafka/bin/../libs/jetty-servlet-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/connect-json-0.9.0.1.jar:/opt/kafka/bin/../libs/validation-api-1.1.0.Final.jar:/opt/kafka/bin/../libs/zkclient-0.7.jar:/opt/kafka/bin/../libs/slf4j-api-1.7.6.jar:/opt/kafka/bin/../libs/jackson-jaxrs-base-2.5.4.jar:/opt/kafka/bin/../libs/jersey-guava-2.22.1.jar:/opt/kafka/bin/../libs/jetty-util-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/kafka_2.11-0.9.0.1.jar:/opt/kafka/bin/../libs/kafka_2.11-0.9.0.1-test.jar:/opt/kafka/bin/../libs/hk2-utils-2.4.0-b31.jar:/opt/kafka/bin/../libs/jersey-common-2.22.1.jar:/opt/kafka/bin/../libs/javax.inject-1.jar:/opt/kafka/bin/../libs/kafka-log4j-appender-0.9.0.1.jar:/opt/kafka/bin/../libs/kafka-clients-0.9.0.1.jar:/opt/kafka/bin/../libs/jersey-server-2.22.1.jar:/opt/kafka/bin/../libs/jackson-annotations-2.5.0.jar:/opt/kafka/bin/../libs/jackson-jaxrs-json-provider-2.5.4.jar:/opt/kafka/bin/../libs/slf4j-log4j12-1.7.6.jar:/opt/kafka/bin/../libs/jetty-http-9.2.12.v20150709.jar:/opt/kafka/bin/../libs/jackson-module-jaxb-annotations-2.5.4.jar:/opt/kafka/bin/../libs/connect-runtime-0.9.0.1.jar:/opt/kafka/bin/../libs/jersey-media-jaxb-2.22.1.jar:/opt/kafka/bin/../libs/log4j-1.2.17.jar (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,201 [myid:] - INFO  [main:QuorumPeerConfig@103] - Reading configuration from: /opt/zookeeper-3.4.6/bin/../conf/zoo.cfg
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:java.library.path=/usr/java/packages/lib/amd64:/usr/lib/x86_64-linux-gnu/jni:/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/lib/jni:/lib:/usr/lib (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:java.io.tmpdir=/tmp (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,202 [myid:] - INFO  [main:ZooKeeperServerMain@95] - Starting server
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:java.compiler=<NA> (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,207 [myid:] - INFO  [main:Environment@100] - Server environment:zookeeper.version=3.4.6-1569965, built on 02/20/2014 09:09 GMT
zookeeper_1       | 2019-07-22 07:02:05,207 [myid:] - INFO  [main:Environment@100] - Server environment:host.name=6d3da24fc0a6
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:os.name=Linux (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,207 [myid:] - INFO  [main:Environment@100] - Server environment:java.version=1.8.0_66-internal
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:os.arch=amd64 (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:os.version=4.15.0-54-generic (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,207 [myid:] - INFO  [main:Environment@100] - Server environment:java.vendor=Oracle Corporation
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:user.name=root (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:user.home=/root (org.apache.zookeeper.ZooKeeper)
zookeeper_1       | 2019-07-22 07:02:05,208 [myid:] - INFO  [main:Environment@100] - Server environment:java.home=/usr/lib/jvm/java-8-openjdk-amd64/jre
kafka_1           | [2019-07-22 07:02:12,268] INFO Client environment:user.dir=/ (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,269] INFO Initiating client connection, connectString=zookeeper:2181 sessionTimeout=6000 watcher=org.I0Itec.zkclient.ZkClient@6ce139a4 (org.apache.zookeeper.ZooKeeper)
kafka_1           | [2019-07-22 07:02:12,280] INFO Waiting for keeper state SyncConnected (org.I0Itec.zkclient.ZkClient)
zookeeper_1       | 2019-07-22 07:02:05,208 [myid:] - INFO  [main:Environment@100] - Server environment:java.class.path=/opt/zookeeper-3.4.6/bin/../build/classes:/opt/zookeeper-3.4.6/bin/../build/lib/*.jar:/opt/zookeeper-3.4.6/bin/../lib/slf4j-log4j12-1.6.1.jar:/opt/zookeeper-3.4.6/bin/../lib/slf4j-api-1.6.1.jar:/opt/zookeeper-3.4.6/bin/../lib/netty-3.7.0.Final.jar:/opt/zookeeper-3.4.6/bin/../lib/log4j-1.2.16.jar:/opt/zookeeper-3.4.6/bin/../lib/jline-0.9.94.jar:/opt/zookeeper-3.4.6/bin/../zookeeper-3.4.6.jar:/opt/zookeeper-3.4.6/bin/../src/java/lib/*.jar:/opt/zookeeper-3.4.6/bin/../conf:
kafka_1           | [2019-07-22 07:02:12,284] INFO Opening socket connection to server kafka_zookeeper_1.kafka_my-net/192.168.64.2:2181. Will not attempt to authenticate using SASL (unknown error) (org.apache.zookeeper.ClientCnxn)
kafka_1           | [2019-07-22 07:02:12,335] INFO Socket connection established to kafka_zookeeper_1.kafka_my-net/192.168.64.2:2181, initiating session (org.apache.zookeeper.ClientCnxn)
kafka_1           | [2019-07-22 07:02:13,133] INFO Session establishment complete on server kafka_zookeeper_1.kafka_my-net/192.168.64.2:2181, sessionid = 0x16c187cc2b00000, negotiated timeout = 6000 (org.apache.zookeeper.ClientCnxn)
zookeeper_1       | 2019-07-22 07:02:05,208 [myid:] - INFO  [main:Environment@100] - Server environment:java.library.path=/usr/java/packages/lib/amd64:/usr/lib/x86_64-linux-gnu/jni:/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/lib/jni:/lib:/usr/lib
kafka_1           | [2019-07-22 07:02:13,135] INFO zookeeper state changed (SyncConnected) (org.I0Itec.zkclient.ZkClient)
kafka_1           | [2019-07-22 07:02:15,214] INFO Log directory '/tmp/kafka-logs' not found, creating it. (kafka.log.LogManager)
kafka_1           | [2019-07-22 07:02:15,219] INFO Loading logs. (kafka.log.LogManager)
zookeeper_1       | 2019-07-22 07:02:05,208 [myid:] - INFO  [main:Environment@100] - Server environment:java.io.tmpdir=/tmp
kafka_1           | [2019-07-22 07:02:15,223] INFO Logs loading complete. (kafka.log.LogManager)
kafka_1           | [2019-07-22 07:02:15,224] INFO Starting log cleanup with a period of 300000 ms. (kafka.log.LogManager)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:java.compiler=<NA>
kafka_1           | [2019-07-22 07:02:15,226] INFO Starting log flusher with a default period of 9223372036854775807 ms. (kafka.log.LogManager)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:os.name=Linux
kafka_1           | [2019-07-22 07:02:15,228] WARN No meta.properties file under dir /tmp/kafka-logs/meta.properties (kafka.server.BrokerMetadataCheckpoint)
kafka_1           | [2019-07-22 07:02:15,701] INFO Awaiting socket connections on 0.0.0.0:9092. (kafka.network.Acceptor)
kafka_1           | [2019-07-22 07:02:15,704] INFO [Socket Server on Broker 10001], Started 1 acceptor threads (kafka.network.SocketServer)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:os.arch=amd64
kafka_1           | [2019-07-22 07:02:15,721] INFO [ExpirationReaper-10001], Starting  (kafka.server.DelayedOperationPurgatory$ExpiredOperationReaper)
kafka_1           | [2019-07-22 07:02:15,722] INFO [ExpirationReaper-10001], Starting  (kafka.server.DelayedOperationPurgatory$ExpiredOperationReaper)
kafka_1           | [2019-07-22 07:02:15,771] INFO Creating /controller (is it secure? false) (kafka.utils.ZKCheckedEphemeral)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:os.version=4.15.0-54-generic
kafka_1           | [2019-07-22 07:02:15,966] INFO Result of znode creation is: OK (kafka.utils.ZKCheckedEphemeral)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:user.name=root
kafka_1           | [2019-07-22 07:02:15,966] INFO 10001 successfully elected as leader (kafka.server.ZookeeperLeaderElector)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:user.home=/root
kafka_1           | [2019-07-22 07:02:16,465] INFO [GroupCoordinator 10001]: Starting up. (kafka.coordinator.GroupCoordinator)
zookeeper_1       | 2019-07-22 07:02:05,210 [myid:] - INFO  [main:Environment@100] - Server environment:user.dir=/
kafka_1           | [2019-07-22 07:02:16,467] INFO [ExpirationReaper-10001], Starting  (kafka.server.DelayedOperationPurgatory$ExpiredOperationReaper)
zookeeper_1       | 2019-07-22 07:02:05,212 [myid:] - INFO  [main:ZooKeeperServer@755] - tickTime set to 2000
kafka_1           | [2019-07-22 07:02:16,469] INFO [ExpirationReaper-10001], Starting  (kafka.server.DelayedOperationPurgatory$ExpiredOperationReaper)
zookeeper_1       | 2019-07-22 07:02:05,212 [myid:] - INFO  [main:ZooKeeperServer@764] - minSessionTimeout set to -1
kafka_1           | [2019-07-22 07:02:16,469] INFO [GroupCoordinator 10001]: Startup complete. (kafka.coordinator.GroupCoordinator)
zookeeper_1       | 2019-07-22 07:02:05,212 [myid:] - INFO  [main:ZooKeeperServer@773] - maxSessionTimeout set to -1
kafka_1           | [2019-07-22 07:02:16,471] INFO [Group Metadata Manager on Broker 10001]: Removed 0 expired offsets in 7 milliseconds. (kafka.coordinator.GroupMetadataManager)
zookeeper_1       | 2019-07-22 07:02:05,222 [myid:] - INFO  [main:NIOServerCnxnFactory@94] - binding to port 0.0.0.0/0.0.0.0:2181
kafka_1           | [2019-07-22 07:02:16,495] INFO [ThrottledRequestReaper-Produce], Starting  (kafka.server.ClientQuotaManager$ThrottledRequestReaper)
zookeeper_1       | 2019-07-22 07:02:12,337 [myid:] - INFO  [NIOServerCxn.Factory:0.0.0.0/0.0.0.0:2181:NIOServerCnxnFactory@197] - Accepted socket connection from /192.168.64.3:53996
kafka_1           | [2019-07-22 07:02:16,496] INFO [ThrottledRequestReaper-Fetch], Starting  (kafka.server.ClientQuotaManager$ThrottledRequestReaper)
zookeeper_1       | 2019-07-22 07:02:12,393 [myid:] - INFO  [NIOServerCxn.Factory:0.0.0.0/0.0.0.0:2181:ZooKeeperServer@868] - Client attempting to establish new session at /192.168.64.3:53996
kafka_1           | [2019-07-22 07:02:16,504] INFO Will not load MX4J, mx4j-tools.jar is not in the classpath (kafka.utils.Mx4jLoader$)
zookeeper_1       | 2019-07-22 07:02:12,395 [myid:] - INFO  [SyncThread:0:FileTxnLog@199] - Creating new log file: log.1
zookeeper_1       | 2019-07-22 07:02:13,131 [myid:] - INFO  [SyncThread:0:ZooKeeperServer@617] - Established session 0x16c187cc2b00000 with negotiated timeout 6000 for client /192.168.64.3:53996
kafka_1           | [2019-07-22 07:02:16,516] INFO Creating /brokers/ids/10001 (is it secure? false) (kafka.utils.ZKCheckedEphemeral)
kafka_1           | [2019-07-22 07:02:16,536] INFO New leader is 10001 (kafka.server.ZookeeperLeaderElector$LeaderChangeListener)
zookeeper_1       | 2019-07-22 07:02:13,190 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x5 zxid:0x3 txntype:-1 reqpath:n/a Error Path:/brokers Error:KeeperErrorCode = NoNode for /brokers
zookeeper_1       | 2019-07-22 07:02:14,136 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0xb zxid:0x7 txntype:-1 reqpath:n/a Error Path:/config Error:KeeperErrorCode = NoNode for /config
kafka_1           | [2019-07-22 07:02:16,943] INFO Result of znode creation is: OK (kafka.utils.ZKCheckedEphemeral)
kafka_1           | [2019-07-22 07:02:16,945] INFO Registered broker 10001 at path /brokers/ids/10001 with addresses: PLAINTEXT -> EndPoint(00eded330f35,9092,PLAINTEXT) (kafka.utils.ZkUtils)
zookeeper_1       | 2019-07-22 07:02:14,455 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x13 zxid:0xc txntype:-1 reqpath:n/a Error Path:/admin Error:KeeperErrorCode = NoNode for /admin
kafka_1           | [2019-07-22 07:02:16,953] INFO Kafka version : 0.9.0.1 (org.apache.kafka.common.utils.AppInfoParser)
kafka_1           | [2019-07-22 07:02:16,953] INFO Kafka commitId : 23c69d62a0cabf06 (org.apache.kafka.common.utils.AppInfoParser)
zookeeper_1       | 2019-07-22 07:02:15,968 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:setData cxid:0x22 zxid:0x13 txntype:-1 reqpath:n/a Error Path:/controller_epoch Error:KeeperErrorCode = NoNode for /controller_epoch
zookeeper_1       | 2019-07-22 07:02:16,408 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:delete cxid:0x31 zxid:0x15 txntype:-1 reqpath:n/a Error Path:/admin/preferred_replica_election Error:KeeperErrorCode = NoNode for /admin/preferred_replica_election
kafka_1           | [2019-07-22 07:02:16,954] INFO [Kafka Server 10001], started (kafka.server.KafkaServer)
zookeeper_1       | 2019-07-22 07:02:16,517 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x38 zxid:0x16 txntype:-1 reqpath:n/a Error Path:/brokers Error:KeeperErrorCode = NodeExists for /brokers
zookeeper_1       | 2019-07-22 07:02:16,521 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x39 zxid:0x17 txntype:-1 reqpath:n/a Error Path:/brokers/ids Error:KeeperErrorCode = NodeExists for /brokers/ids
zookeeper_1       | 2019-07-22 07:02:24,106 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:setData cxid:0x41 zxid:0x19 txntype:-1 reqpath:n/a Error Path:/config/topics/default-topic Error:KeeperErrorCode = NoNode for /config/topics/default-topic
zookeeper_1       | 2019-07-22 07:02:24,806 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x42 zxid:0x1a txntype:-1 reqpath:n/a Error Path:/config/topics Error:KeeperErrorCode = NodeExists for /config/topics
kafka_1           | [2019-07-22 07:02:24,839] INFO Topic creation {"version":1,"partitions":{"0":[10001]}} (kafka.admin.AdminUtils$)
kafka_1           | [2019-07-22 07:02:24,855] INFO [KafkaApi-10001] Auto creation of topic default-topic with 1 partitions and replication factor 1 is successful! (kafka.server.KafkaApis)
zookeeper_1       | 2019-07-22 07:02:24,879 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x4a zxid:0x1d txntype:-1 reqpath:n/a Error Path:/brokers/topics/default-topic/partitions/0 Error:KeeperErrorCode = NoNode for /brokers/topics/default-topic/partitions/0
zookeeper_1       | 2019-07-22 07:02:24,894 [myid:] - INFO  [ProcessThread(sid:0 cport:-1)::PrepRequestProcessor@645] - Got user-level KeeperException when processing sessionid:0x16c187cc2b00000 type:create cxid:0x4b zxid:0x1e txntype:-1 reqpath:n/a Error Path:/brokers/topics/default-topic/partitions Error:KeeperErrorCode = NoNode for /brokers/topics/default-topic/partitions
kafka_1           | [2019-07-22 07:02:24,983] INFO [ReplicaFetcherManager on broker 10001] Removed fetcher for partitions [default-topic,0] (kafka.server.ReplicaFetcherManager)
kafka_1           | [2019-07-22 07:02:25,005] INFO Completed load of log default-topic-0 with log end offset 0 (kafka.log.Log)
kafka_1           | [2019-07-22 07:02:25,009] INFO Created log for partition [default-topic,0] in /tmp/kafka-logs with properties {compression.type -> producer, file.delete.delay.ms -> 60000, max.message.bytes -> 1000012, min.insync.replicas -> 1, segment.jitter.ms -> 0, preallocate -> false, min.cleanable.dirty.ratio -> 0.5, index.interval.bytes -> 4096, unclean.leader.election.enable -> true, retention.bytes -> -1, delete.retention.ms -> 86400000, cleanup.policy -> delete, flush.ms -> 9223372036854775807, segment.ms -> 604800000, segment.bytes -> 1073741824, retention.ms -> 604800000, segment.index.bytes -> 10485760, flush.messages -> 9223372036854775807}. (kafka.log.LogManager)
kafka_1           | [2019-07-22 07:02:25,009] INFO Partition [default-topic,0] on broker 10001: No checkpointed highwatermark is found for partition [default-topic,0] (kafka.cluster.Partition)
kafka-consumer_1  | Received messages  Hello Kafka 1
kafka-producer_1  | Message is stored in topic(default-topic)/partition(0)/offset(0)
kafka-producer_1  | Message is stored in topic(default-topic)/partition(0)/offset(1)
kafka-consumer_1  | Received messages  Hello Kafka 2
kafka-consumer_1  | Received messages  Hello Kafka 3
kafka-producer_1  | Message is stored in topic(default-topic)/partition(0)/offset(2)
kafka-consumer_1  | Received messages  Hello Kafka 4
kafka-producer_1  | Message is stored in topic(default-topic)/partition(0)/offset(3)
kafka-producer_1  | Message is stored in topic(default-topic)/partition(0)/offset(4)
kafka-consumer_1  | Received messages  Hello Kafka 5
^CGracefully stopping... (press Ctrl+C again to force)
Stopping kafka_kafka-producer_1 ... done
Stopping kafka_kafka-consumer_1 ... done
Stopping kafka_kafka_1          ... done
Stopping kafka_zookeeper_1      ... done
```
