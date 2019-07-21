# Microservice communication

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implement an RPC client and server system with the ProtoBuf binary protocol](#2-implement-an-rpc-client-and-server-system-with-the-protobuf-binary-protocol)
3. [Add synchronous Hystrix circuit breaker and monitor to the RPC client and server](#3-add-synchronous-hystrix-circuit-breaker-and-monitor-to-the-rpc-client-and-server)
4. [Implement async message queuing with RabbitMQ](#4-implement-async-message-queuing-with-rabbitmq)
5. [Implement publish-subscribe Kafka clients](#5-implement-publish-subscribe-kafka-clients)

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

## 3 Add synchronous Hystrix circuit breaker and monitor to the RPC client and server

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

## 5 Implement publish-subscribe Kafka clients