# Microservice discovery

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Run Consul](#2-run-consul)
3. [Register services with Consul using REST API](#3-register-services-with-consul-using-rest-api)
4. [Lookup services using the Consul UI and REST API](#4-lookup-services-using-the-consul-ui-and-rest-api)
5. [Go microservices registration with Consul](#5-go-microservices-registration-with-consul)
6. [Go microservices lookup with Consul](#6-go-microservices-lookup-with-consul)
7. [Go microservices discovery just with Kubernetes](#7-go-microservices-discovery-just-with-kubernetes)

## 1 Document objective

In this block we are going to:
 
* Use Consul for service discovery
* Perform service endpoint registration using Consul
* Implement microservice discovery with Go

## 2 Run Consul

Deploy Consul and two 'microservice-in-go' microservices in Docker:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
15:10:31 $ docker-compose up -d
Pulling consul (consul:latest)...
latest: Pulling from library/consul
e7c96db7181b: Pull complete
2967157a1cec: Pull complete
89eac26c7594: Pull complete
fed432a284a5: Pull complete
eff914b7f5d7: Pull complete
0c1d0a78f0c3: Pull complete
Digest: sha256:b31edc821d5e3deae8ce9f9a2070dc3fbaa72f5e1746a85a91ebe551ed8fb17f
Status: Downloaded newer image for consul:latest
Creating microservicediscovery_microservice-in-go-02_1 ... done
Creating microservicediscovery_microservice-in-go-01_1 ... done
Creating microservicediscovery_consul_1                ... done
```

Opening a browser and navigating to my Consul UI:

![alt text](images/image01.png "My Consul UI")

Consul's services catalog starts empty:

```
15:15:19 $ curl http://localhost:8500/v1/catalog/services
{
    "consul": []
}
```

## 3 Register services with Consul using REST API

Register Consul agent 'microservice-in-go-01':

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:36:41 $ cat register_consul_agent_microservice-in-go-01.json 
{
  "ID": "microservice-in-go-01",
  "Name": "microservice-in-go",
  "Tags": [
    "cloud-native-go",
    "v1"
  ],
  "Address": "localhost",
  "Port": 8080,
  "EnableTagOverride": false,
  "check": {
    "id": "ping",
    "name": "HTTP API on port 8080",
    "http": "http://microservice-in-go-01:8080/ping",
    "interval": "5s",
    "timeout": "1s"
  }
}

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:36:47 $ cat register_consul_agent_microservice-in-go-02.json 
{
  "ID": "microservice-in-go-02",
  "Name": "microservice-in-go",
  "Tags": [
    "cloud-native-go",
    "v1"
  ],
  "Address": "localhost",
  "Port": 9090,
  "EnableTagOverride": false,
  "check": {
    "id": "ping",
    "name": "HTTP API on port 9090",
    "http": "http://microservice-in-go-02:9090/ping",
    "interval": "5s",
    "timeout": "1s"
  }
}

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:38:37 $ curl -d "@register_consul_agent_microservice-in-go-01.json" -H "Content-Type: application/json" -X PUT http://localhost:8500/v1/agent/service/register

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
19:01:20 $ curl -d "@register_consul_agent_microservice-in-go-02.json" -H "Content-Type: application/json" -X PUT http://localhost:8500/v1/agent/service/register
```
We can unregister a service calling this url:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
08:30:59 $ curl -H "Content-Type: application/json" -X PUT http://localhost:8500/v1/agent/service/deregister/microservice-in-go-02
```


## 4 Lookup services using the Consul UI and REST API

Look at the registered services:

![alt text](images/image02.png "Service is registered")

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:39:10 $ curl http://localhost:8500/v1/catalog/services
{
    "consul": [],
    "microservice-in-go": [
        "cloud-native-go",
        "v1"
    ]
}
(base) 
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:39:25 $ curl http://localhost:8500/v1/agent/services
{
    "microservice-in-go-01": {
        "ID": "microservice-in-go-01",
        "Service": "microservice-in-go",
        "Tags": [
            "cloud-native-go",
            "v1"
        ],
        "Meta": {},
        "Port": 8080,
        "Address": "localhost",
        "Weights": {
            "Passing": 1,
            "Warning": 1
        },
        "EnableTagOverride": false
    },
    "microservice-in-go-02": {
        "ID": "microservice-in-go-02",
        "Service": "microservice-in-go",
        "Tags": [
            "cloud-native-go",
            "v1"
        ],
        "Meta": {},
        "Port": 9090,
        "Address": "localhost",
        "Weights": {
            "Passing": 1,
            "Warning": 1
        },
        "EnableTagOverride": false
    }
}
```

Get all the healthchecks information:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
19:09:34 $ curl http://localhost:8500/v1/health/service/microservice-in-go
[
    {
        "Node": {
            "ID": "4f3777fd-01da-afbd-235b-8f946911ed41",
            "Node": "819d33de1440",
            "Address": "127.0.0.1",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "127.0.0.1",
                "wan": "127.0.0.1"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 9,
            "ModifyIndex": 10
        },
        "Service": {
            "ID": "microservice-in-go-01",
            "Service": "microservice-in-go",
            "Tags": [
                "cloud-native-go",
                "v1"
            ],
            "Address": "localhost",
            "Meta": null,
            "Port": 8080,
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "ProxyDestination": "",
            "Proxy": {},
            "Connect": {},
            "CreateIndex": 407,
            "ModifyIndex": 407
        },
        "Checks": [
            {
                "Node": "819d33de1440",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Definition": {},
                "CreateIndex": 9,
                "ModifyIndex": 9
            },
            {
                "Node": "819d33de1440",
                "CheckID": "service:microservice-in-go-01",
                "Name": "HTTP API on port 8080",
                "Status": "passing",
                "Notes": "",
                "Output": "HTTP GET http://microservice-in-go-01:8080/ping: 200 OK Output: pong",
                "ServiceID": "microservice-in-go-01",
                "ServiceName": "microservice-in-go",
                "ServiceTags": [
                    "cloud-native-go",
                    "v1"
                ],
                "Definition": {},
                "CreateIndex": 407,
                "ModifyIndex": 408
            }
        ]
    },
    {
        "Node": {
            "ID": "4f3777fd-01da-afbd-235b-8f946911ed41",
            "Node": "819d33de1440",
            "Address": "127.0.0.1",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "127.0.0.1",
                "wan": "127.0.0.1"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 9,
            "ModifyIndex": 10
        },
        "Service": {
            "ID": "microservice-in-go-02",
            "Service": "microservice-in-go",
            "Tags": [
                "cloud-native-go",
                "v1"
            ],
            "Address": "localhost",
            "Meta": null,
            "Port": 9090,
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "ProxyDestination": "",
            "Proxy": {},
            "Connect": {},
            "CreateIndex": 412,
            "ModifyIndex": 412
        },
        "Checks": [
            {
                "Node": "819d33de1440",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Definition": {},
                "CreateIndex": 9,
                "ModifyIndex": 9
            },
            {
                "Node": "819d33de1440",
                "CheckID": "service:microservice-in-go-02",
                "Name": "HTTP API on port 9090",
                "Status": "passing",
                "Notes": "",
                "Output": "HTTP GET http://microservice-in-go-02:9090/ping: 200 OK Output: pong",
                "ServiceID": "microservice-in-go-02",
                "ServiceName": "microservice-in-go",
                "ServiceTags": [
                    "cloud-native-go",
                    "v1"
                ],
                "Definition": {},
                "CreateIndex": 412,
                "ModifyIndex": 413
            }
        ]
    }
]
```

Get the passing api information:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
19:11:26 $ curl http://localhost:8500/v1/health/service/microservice-in-go?passing
[
    {
        "Node": {
            "ID": "4f3777fd-01da-afbd-235b-8f946911ed41",
            "Node": "819d33de1440",
            "Address": "127.0.0.1",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "127.0.0.1",
                "wan": "127.0.0.1"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 9,
            "ModifyIndex": 10
        },
        "Service": {
            "ID": "microservice-in-go-01",
            "Service": "microservice-in-go",
            "Tags": [
                "cloud-native-go",
                "v1"
            ],
            "Address": "localhost",
            "Meta": null,
            "Port": 8080,
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "ProxyDestination": "",
            "Proxy": {},
            "Connect": {},
            "CreateIndex": 407,
            "ModifyIndex": 407
        },
        "Checks": [
            {
                "Node": "819d33de1440",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Definition": {},
                "CreateIndex": 9,
                "ModifyIndex": 9
            },
            {
                "Node": "819d33de1440",
                "CheckID": "service:microservice-in-go-01",
                "Name": "HTTP API on port 8080",
                "Status": "passing",
                "Notes": "",
                "Output": "HTTP GET http://microservice-in-go-01:8080/ping: 200 OK Output: pong",
                "ServiceID": "microservice-in-go-01",
                "ServiceName": "microservice-in-go",
                "ServiceTags": [
                    "cloud-native-go",
                    "v1"
                ],
                "Definition": {},
                "CreateIndex": 407,
                "ModifyIndex": 408
            }
        ]
    },
    {
        "Node": {
            "ID": "4f3777fd-01da-afbd-235b-8f946911ed41",
            "Node": "819d33de1440",
            "Address": "127.0.0.1",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "127.0.0.1",
                "wan": "127.0.0.1"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 9,
            "ModifyIndex": 10
        },
        "Service": {
            "ID": "microservice-in-go-02",
            "Service": "microservice-in-go",
            "Tags": [
                "cloud-native-go",
                "v1"
            ],
            "Address": "localhost",
            "Meta": null,
            "Port": 9090,
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "ProxyDestination": "",
            "Proxy": {},
            "Connect": {},
            "CreateIndex": 412,
            "ModifyIndex": 412
        },
        "Checks": [
            {
                "Node": "819d33de1440",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Definition": {},
                "CreateIndex": 9,
                "ModifyIndex": 9
            },
            {
                "Node": "819d33de1440",
                "CheckID": "service:microservice-in-go-02",
                "Name": "HTTP API on port 9090",
                "Status": "passing",
                "Notes": "",
                "Output": "HTTP GET http://microservice-in-go-02:9090/ping: 200 OK Output: pong",
                "ServiceID": "microservice-in-go-02",
                "ServiceName": "microservice-in-go",
                "ServiceTags": [
                    "cloud-native-go",
                    "v1"
                ],
                "Definition": {},
                "CreateIndex": 412,
                "ModifyIndex": 413
            }
        ]
    }
]
```

Get the critical api information:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
19:13:06 $ curl http://localhost:8500/v1/health/state/critical
[]

```
## 5 Go microservices registration with Consul

In this section we will:

* Implement with Go a service endpoint registration to Consul for our Go microservice
* Implement and register health check endpoint
* Run Consul and our Go microservice and see that our development works as expected

First, download the Hashicorp Consul Go libraries:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:47:28 $ go get github.com/hashicorp/consul/api
go get: warning: modules disabled by GO111MODULE=auto in GOPATH/src;
	ignoring ../../../go.mod;
	see 'go help modules'

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:47:59 $ ls /home/arturotarin/go/src/github.com/hashicorp/
consul

```

Next, build the server.go code provided:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:48:01 $ ls
server.go

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:50:34 $ go build -o server

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:50:57 $ ls -lrth
total 7,6M
-rw-rw-r-- 1 arturotarin arturotarin 1,5K jul 20 06:49 server.go
-rwxrwxr-x 1 arturotarin arturotarin 7,6M jul 20 06:50 server

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
06:51:47 $ ./server 
Starting Go microservice server
```

Next, create the Docker image based in this 'server' runnable file:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
08:11:49 $ docker-compose build
consul uses an image, skipping
Building go-microservice-server
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&     go get -u github.com/hashicorp/consul/api
 ---> Using cache
 ---> ec267151e2cc
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/server/
 ---> Using cache
 ---> 62e21093e1db
Step 4/8 : COPY . ${SOURCES}
 ---> fccb03e7d736
Removing intermediate container 1293f5130649
Step 5/8 : RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build -o go-microservice-server
 ---> Running in 73e8bb412a27
 ---> fc4f4a05c9d6
Removing intermediate container 73e8bb412a27
Step 6/8 : ENV CONSUL_HTTP_ADDR localhost:8500
 ---> Running in cdbc9e92f9e4
 ---> 8f19a59ca03d
Removing intermediate container cdbc9e92f9e4
Step 7/8 : WORKDIR ${SOURCES}server/
 ---> 8507e8415582
Removing intermediate container aabf2f804123
Step 8/8 : CMD ${SOURCES}server/go-microservice-server
 ---> Running in 2e5eb99ad157
 ---> ca7bbf8dffbf
Removing intermediate container 2e5eb99ad157
Successfully built ca7bbf8dffbf
Successfully tagged go-microservice-server:1.0.0
```

List the Docker images:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
07:52:03 $ docker images
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
go-microservice-server             1.0.0               b0864ef7cdbb        47 seconds ago      584MB
go-microservice           1.0.0               6e044057cc9a        17 hours ago        487MB
arturot/go-microservice   1.0.0               d47bd5746538        20 hours ago        477MB
golang                    1.12-alpine         6b21b4c6e7a3        8 days ago          350MB
consul                    latest              7d52b83f718f        3 weeks ago         115MB
```

Start Consul and the Go microservice:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
07:52:49 $ docker-compose up
Starting withgo_consul_1 ... done
Recreating withgo_go-microservice-server_1 ... done
Attaching to withgo_consul_1, withgo_go-microservice-server_1
go-microservice-server_1  | Starting Go microservice server.
consul_1         | ==> Starting Consul agent...
consul_1         |            Version: 'v1.5.2'
consul_1         |            Node ID: 'ddfbaeb0-1926-b72f-deee-a446df74fc97'
consul_1         |          Node name: '72f13483a857'
consul_1         |         Datacenter: 'dc1' (Segment: '<all>')
consul_1         |             Server: true (Bootstrap: false)
consul_1         |        Client Addr: [0.0.0.0] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
consul_1         |       Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
consul_1         |            Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false, Auto-Encrypt-TLS: false
consul_1         | 
consul_1         | ==> Log data will now stream in as it occurs:
consul_1         | 
consul_1         |     2019/07/20 06:13:39 [DEBUG] tlsutil: Update with version 1
consul_1         |     2019/07/20 06:13:39 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
consul_1         |     2019/07/20 06:13:39 [INFO]  raft: Initial configuration (index=1): [{Suffrage:Voter ID:ddfbaeb0-1926-b72f-deee-a446df74fc97 Address:127.0.0.1:8300}]
consul_1         |     2019/07/20 06:13:39 [INFO]  raft: Node at 127.0.0.1:8300 [Follower] entering Follower state (Leader: "")
consul_1         |     2019/07/20 06:13:39 [INFO] serf: EventMemberJoin: 72f13483a857.dc1 127.0.0.1
consul_1         |     2019/07/20 06:13:39 [INFO] serf: EventMemberJoin: 72f13483a857 127.0.0.1
consul_1         |     2019/07/20 06:13:39 [INFO] consul: Handled member-join event for server "72f13483a857.dc1" in area "wan"
consul_1         |     2019/07/20 06:13:39 [INFO] consul: Adding LAN server 72f13483a857 (Addr: tcp/127.0.0.1:8300) (DC: dc1)
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent: restored service definition "go-microservice-server" from "/consul/data/services/7a3514b8a6c6c9c4b5175d2e945565bb"
consul_1         |     2019/07/20 06:13:39 [DEBUG] tlsutil: OutgoingTLSConfigForCheck with version 1
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent: restored health check "service:go-microservice-server" from "/consul/data/checks/5cb16499c28118418f6f31578cf466c6"
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent/proxy: managed Connect proxy manager started
consul_1         |     2019/07/20 06:13:39 [INFO] agent: Started DNS server 0.0.0.0:8600 (udp)
consul_1         |     2019/07/20 06:13:39 [INFO] agent: Started DNS server 0.0.0.0:8600 (tcp)
consul_1         |     2019/07/20 06:13:39 [INFO] agent: Started HTTP server on [::]:8500 (tcp)
consul_1         |     2019/07/20 06:13:39 [INFO] agent: started state syncer
consul_1         | ==> Consul agent running!
consul_1         |     2019/07/20 06:13:39 [INFO] agent: Started gRPC server on [::]:8502 (tcp)
consul_1         |     2019/07/20 06:13:39 [WARN]  raft: Heartbeat timeout from "" reached, starting election
consul_1         |     2019/07/20 06:13:39 [INFO]  raft: Node at 127.0.0.1:8300 [Candidate] entering Candidate state in term 2
consul_1         |     2019/07/20 06:13:39 [DEBUG] raft: Votes needed: 1
consul_1         |     2019/07/20 06:13:39 [DEBUG] raft: Vote granted from ddfbaeb0-1926-b72f-deee-a446df74fc97 in term 2. Tally: 1
consul_1         |     2019/07/20 06:13:39 [INFO]  raft: Election won. Tally: 1
consul_1         |     2019/07/20 06:13:39 [INFO]  raft: Node at 127.0.0.1:8300 [Leader] entering Leader state
consul_1         |     2019/07/20 06:13:39 [INFO] consul: cluster leadership acquired
consul_1         |     2019/07/20 06:13:39 [INFO] consul: New leader elected: 72f13483a857
consul_1         |     2019/07/20 06:13:39 [INFO] connect: initialized primary datacenter CA with provider "consul"
consul_1         |     2019/07/20 06:13:39 [DEBUG] consul: Skipping self join check for "72f13483a857" since the cluster is too small
consul_1         |     2019/07/20 06:13:39 [INFO] consul: member '72f13483a857' joined, marking health alive
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent: Skipping remote check "serfHealth" since it is managed automatically
consul_1         |     2019/07/20 06:13:39 [INFO] agent: Synced service "go-microservice-server"
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:39 [DEBUG] agent: Node info in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Skipping remote check "serfHealth" since it is managed automatically
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Node info in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:40 [DEBUG] agent: Node info in sync
consul_1         |     2019/07/20 06:13:41 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
consul_1         |     2019/07/20 06:13:43 [DEBUG] tlsutil: OutgoingTLSConfigForCheck with version 1
consul_1         |     2019/07/20 06:13:43 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://726d6c0a6298:80/info: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
consul_1         |     2019/07/20 06:13:43 [INFO] agent: Synced service "go-microservice-server"
consul_1         |     2019/07/20 06:13:43 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:43 [DEBUG] agent: Node info in sync
consul_1         |     2019/07/20 06:13:43 [DEBUG] http: Request PUT /v1/agent/service/register (188.300145ms) from=172.29.0.3:49550
consul_1         |     2019/07/20 06:13:43 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:43 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:43 [DEBUG] agent: Node info in sync
go-microservice-server_1  | The /info endpoint is being called...
consul_1         |     2019/07/20 06:13:48 [DEBUG] agent: Check "service:go-microservice-server" is passing
consul_1         |     2019/07/20 06:13:48 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1         |     2019/07/20 06:13:48 [INFO] agent: Synced check "service:go-microservice-server"
consul_1         |     2019/07/20 06:13:48 [DEBUG] agent: Node info in sync
go-microservice-server_1  | The /info endpoint is being called...
consul_1         |     2019/07/20 06:13:53 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
consul_1         |     2019/07/20 06:13:58 [DEBUG] agent: Check "service:go-microservice-server" is passing
```

## 6 Go microservices lookup with Consul

In this section we will:

* Lookup the Go server microservice using Consul UI and Consul REST API calls
* Implement a client microservice application with Go
* Implement a service endpoint lookup via Consul
* Run Consul and our client and server Go microservices all together and see that our development works as expected

So this is what we can see this from either the Consul UI or from command line Consul REST API calls:

![alt text](images/image03.png "go-microservice-server has been registered using Go")

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:39:10 $ curl http://localhost:8500/v1/catalog/services
{
    "consul": [],
    "go-microservice-server": []
}

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/just_consul
18:39:25 $ curl http://localhost:8500/v1/agent/services
{
    "go-microservice-server": {
        "ID": "go-microservice-server",
        "Service": "go-microservice-server",
        "Tags": [],
        "Meta": {},
        "Port": 8080,
        "Address": "10cdd5bfcd08",
        "Weights": {
            "Passing": 1,
            "Warning": 1
        },
        "EnableTagOverride": false
    }
}
```

But the good thing is that a Go microservice client like the provided in our example can find the 
Go microservice server example using the Consul APIs as well. Let's see how.

First, build the go-microservice-client and test it locally:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
18:38:26 $ go build -o go-microservice-client

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
18:38:48 $ ls -lrth
total 15M
-rw-rw-r-- 1 arturotarin arturotarin  570 jul 20 16:54 Dockerfile
-rw-rw-r-- 1 arturotarin arturotarin 1,3K jul 20 18:38 client.go
-rwxrwxr-x 1 arturotarin arturotarin 7,2M jul 20 18:38 go-microservice-client
```

Test it locally:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
18:38:50 $ ./go-microservice-client 
Starting Go microservice client.
```

And see what is going on in Consul log:

```
consul_1                  |     2019/07/20 16:40:47 [DEBUG] http: Request GET /v1/agent/services (1.167346ms) from=172.29.0.1:48308
```

Now let's add this block in the docker-compose.yml file (already done):

```
  go-microservice-client:
    build:
      context: .
      dockerfile: client/Dockerfile
    image: go-microservice-client:1.0.0
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
    depends_on:
      - consul
      - go-microservice-server
    networks:
      - my-net
```

And run the three all together, Consul, go-microservice-server and go-microservice-client:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go
19:51:33 $ docker-compose up
Creating withgo_consul_1 ... done
Creating withgo_go-microservice-server_1 ... done
Creating withgo_go-microservice-client_1 ... done
Attaching to withgo_consul_1, withgo_go-microservice-server_1, withgo_go-microservice-client_1
consul_1                  | ==> Starting Consul agent...
consul_1                  |            Version: 'v1.5.2'
consul_1                  |            Node ID: 'a8a393e2-f416-f7e0-6945-622760ea7d1a'
consul_1                  |          Node name: 'b0aa3906d740'
consul_1                  |         Datacenter: 'dc1' (Segment: '<all>')
consul_1                  |             Server: true (Bootstrap: false)
consul_1                  |        Client Addr: [0.0.0.0] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
go-microservice-server_1  | Starting Go microservice server.
go-microservice-client_1  | Starting Go microservice client.
consul_1                  |       Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
consul_1                  |            Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false, Auto-Encrypt-TLS: false
consul_1                  | 
consul_1                  | ==> Log data will now stream in as it occurs:
consul_1                  | 
consul_1                  |     2019/07/20 17:51:42 [DEBUG] agent: Using random ID "a8a393e2-f416-f7e0-6945-622760ea7d1a" as node ID
consul_1                  |     2019/07/20 17:51:42 [DEBUG] tlsutil: Update with version 1
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:51:42 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
consul_1                  |     2019/07/20 17:51:42 [INFO]  raft: Initial configuration (index=1): [{Suffrage:Voter ID:a8a393e2-f416-f7e0-6945-622760ea7d1a Address:127.0.0.1:8300}]
consul_1                  |     2019/07/20 17:51:42 [INFO]  raft: Node at 127.0.0.1:8300 [Follower] entering Follower state (Leader: "")
consul_1                  |     2019/07/20 17:51:42 [INFO] serf: EventMemberJoin: b0aa3906d740.dc1 127.0.0.1
consul_1                  |     2019/07/20 17:51:42 [INFO] serf: EventMemberJoin: b0aa3906d740 127.0.0.1
consul_1                  |     2019/07/20 17:51:42 [INFO] consul: Adding LAN server b0aa3906d740 (Addr: tcp/127.0.0.1:8300) (DC: dc1)
consul_1                  |     2019/07/20 17:51:42 [INFO] consul: Handled member-join event for server "b0aa3906d740.dc1" in area "wan"
consul_1                  |     2019/07/20 17:51:42 [DEBUG] agent/proxy: managed Connect proxy manager started
consul_1                  |     2019/07/20 17:51:42 [INFO] agent: Started DNS server 0.0.0.0:8600 (tcp)
consul_1                  |     2019/07/20 17:51:42 [INFO] agent: Started DNS server 0.0.0.0:8600 (udp)
consul_1                  |     2019/07/20 17:51:42 [INFO] agent: Started HTTP server on [::]:8500 (tcp)
consul_1                  |     2019/07/20 17:51:42 [INFO] agent: Started gRPC server on [::]:8502 (tcp)
consul_1                  |     2019/07/20 17:51:42 [INFO] agent: started state syncer
consul_1                  | ==> Consul agent running!
consul_1                  |     2019/07/20 17:51:42 [WARN]  raft: Heartbeat timeout from "" reached, starting election
consul_1                  |     2019/07/20 17:51:42 [INFO]  raft: Node at 127.0.0.1:8300 [Candidate] entering Candidate state in term 2
consul_1                  |     2019/07/20 17:51:42 [DEBUG] raft: Votes needed: 1
consul_1                  |     2019/07/20 17:51:42 [DEBUG] raft: Vote granted from a8a393e2-f416-f7e0-6945-622760ea7d1a in term 2. Tally: 1
consul_1                  |     2019/07/20 17:51:42 [INFO]  raft: Election won. Tally: 1
consul_1                  |     2019/07/20 17:51:42 [INFO]  raft: Node at 127.0.0.1:8300 [Leader] entering Leader state
consul_1                  |     2019/07/20 17:51:42 [INFO] consul: cluster leadership acquired
consul_1                  |     2019/07/20 17:51:42 [INFO] consul: New leader elected: b0aa3906d740
consul_1                  |     2019/07/20 17:51:42 [INFO] connect: initialized primary datacenter CA with provider "consul"
consul_1                  |     2019/07/20 17:51:42 [DEBUG] consul: Skipping self join check for "b0aa3906d740" since the cluster is too small
consul_1                  |     2019/07/20 17:51:42 [INFO] consul: member 'b0aa3906d740' joined, marking health alive
consul_1                  |     2019/07/20 17:51:43 [DEBUG] agent: Skipping remote check "serfHealth" since it is managed automatically
consul_1                  |     2019/07/20 17:51:43 [INFO] agent: Synced node info
consul_1                  |     2019/07/20 17:51:43 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:51:44 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
consul_1                  |     2019/07/20 17:51:45 [DEBUG] agent: Skipping remote check "serfHealth" since it is managed automatically
consul_1                  |     2019/07/20 17:51:45 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:51:46 [DEBUG] tlsutil: OutgoingTLSConfigForCheck with version 1
consul_1                  |     2019/07/20 17:51:46 [INFO] agent: Synced service "go-microservice-server"
consul_1                  |     2019/07/20 17:51:46 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1                  |     2019/07/20 17:51:46 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:51:46 [DEBUG] http: Request PUT /v1/agent/service/register (404.693622ms) from=172.29.0.3:42478
consul_1                  |     2019/07/20 17:51:46 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1                  |     2019/07/20 17:51:46 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1                  |     2019/07/20 17:51:46 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:51:49 [DEBUG] agent: Check "service:go-microservice-server" is passing
consul_1                  |     2019/07/20 17:51:49 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1                  |     2019/07/20 17:51:49 [INFO] agent: Synced check "service:go-microservice-server"
consul_1                  |     2019/07/20 17:51:49 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:51:49 [DEBUG] http: Request GET /v1/agent/services (1.182676ms) from=172.29.0.4:39862
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:51:54 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:51:54.839648149 +0000 UTC m=+5.003828300
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:51:59 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:51:59.839703835 +0000 UTC m=+10.003884007
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:04 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:04.839611529 +0000 UTC m=+15.003791742
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:09 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:09.839656327 +0000 UTC m=+20.003836519
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:14 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:14.839707265 +0000 UTC m=+25.003887513
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:19 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:19.83978128 +0000 UTC m=+30.003961452
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:24 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:24.839706182 +0000 UTC m=+35.003886374
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:29 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:29.839661362 +0000 UTC m=+40.003841441
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:52:34 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:34.839644966 +0000 UTC m=+45.003825159
```

You can also look at the same information if you trace the go-microservice-client logs:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
19:43:24 $ docker ps
CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS              PORTS                                                                                                                      NAMES
eb126a18524d        go-microservice-client:1.0.0   "/bin/sh -c ${SOUR..."   20 seconds ago      Up 16 seconds                                                                                                                                  withgo_go-microservice-client_1
9da7180f545a        go-microservice-server:1.0.0   "/bin/sh -c ${SOUR..."   24 seconds ago      Up 20 seconds                                                                                                                                  withgo_go-microservice-server_1
b0aa3906d740        consul:latest                  "docker-entrypoint..."   29 seconds ago      Up 24 seconds       0.0.0.0:8300->8300/tcp, 0.0.0.0:8400->8400/tcp, 8301-8302/tcp, 8301-8302/udp, 0.0.0.0:8500->8500/tcp, 8600/tcp, 8600/udp   withgo_consul_1

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
19:52:26 $ docker logs eb126a18524d
Starting Go microservice client.
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:51:54.839648149 +0000 UTC m=+5.003828300
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:51:59.839703835 +0000 UTC m=+10.003884007
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:04.839611529 +0000 UTC m=+15.003791742
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:09.839656327 +0000 UTC m=+20.003836519
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:14.839707265 +0000 UTC m=+25.003887513
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:19.83978128 +0000 UTC m=+30.003961452
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:24.839706182 +0000 UTC m=+35.003886374
Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:52:29.839661362 +0000 UTC m=+40.003841441
```


Let's check how does lookup behave when things get complicated, i.e. what happens in the go-microservice-client 
in case the go-microservice-server goes down? Also, what happens in Consul?

To check it, lets stop the go-microservice-server container manually in Docker:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
19:52:29 $ docker ps
CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS              PORTS                                                                                                                      NAMES
eb126a18524d        go-microservice-client:1.0.0   "/bin/sh -c ${SOUR..."   6 minutes ago       Up 4 seconds                                                                                                                                   withgo_go-microservice-client_1
9da7180f545a        go-microservice-server:1.0.0   "/bin/sh -c ${SOUR..."   6 minutes ago       Up 6 minutes                                                                                                                                   withgo_go-microservice-server_1
b0aa3906d740        consul:latest                  "docker-entrypoint..."   6 minutes ago       Up 6 minutes        0.0.0.0:8300->8300/tcp, 0.0.0.0:8400->8400/tcp, 8301-8302/tcp, 8301-8302/udp, 0.0.0.0:8500->8500/tcp, 8600/tcp, 8600/udp   withgo_consul_1

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
19:58:21 $ docker stop 9da7180f545a
9da7180f545a
```

What we can see is this log in the go-microservice-client and Consul logs:

```
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:58:51.147124403 +0000 UTC m=+35.002978711
consul_1                  |     2019/07/20 17:58:54 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
consul_1                  |     2019/07/20 17:58:54 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1                  |     2019/07/20 17:58:54 [INFO] agent: Synced check "service:go-microservice-server"
consul_1                  |     2019/07/20 17:58:54 [DEBUG] agent: Node info in sync
withgo_go-microservice-server_1 exited with code 2
consul_1                  |     2019/07/20 17:59:02 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
go-microservice-client_1  | Get http://9da7180f545a:8080/info: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
go-microservice-client_1  | Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: read udp 127.0.0.1:45400->127.0.0.11:53: i/o timeout
consul_1                  |     2019/07/20 17:59:07 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
go-microservice-client_1  | Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
consul_1                  |     2019/07/20 17:59:12 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
go-microservice-client_1  | Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
consul_1                  |     2019/07/20 17:59:17 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: dial tcp: lookup 9da7180f545a on 127.0.0.11:53: no such host
consul_1                  |     2019/07/20 17:59:20 [DEBUG] tlsutil: OutgoingTLSConfigForCheck with version 1
consul_1                  |     2019/07/20 17:59:20 [WARN] agent: Check "service:go-microservice-server" HTTP request failed: Get http://9da7180f545a:8080/info: dial tcp 172.29.0.3:8080: connect: connection refused
go-microservice-client_1  | Get http://9da7180f545a:8080/info: dial tcp 172.29.0.3:8080: connect: connection refused
```

And if we start up again the go-microservice-server, what happens then?

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client
19:58:54 $ docker start 9da7180f545a
9da7180f545a
```

First, what it happens is that the go-microservice-server log is written:

```
go-microservice-server_1  | Starting Go microservice server.
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-server_1  | The /info endpoint is being called...
```

Next, Consul does the service sinchronization:
```
consul_1                  |     2019/07/20 17:59:21 [INFO] agent: Synced service "go-microservice-server"
consul_1                  |     2019/07/20 17:59:21 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1                  |     2019/07/20 17:59:21 [DEBUG] agent: Node info in sync
consul_1                  |     2019/07/20 17:59:21 [DEBUG] http: Request PUT /v1/agent/service/register (1.291463943s) from=172.29.0.3:42744
go-microservice-server_1  | Starting Go microservice server.
consul_1                  |     2019/07/20 17:59:21 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1                  |     2019/07/20 17:59:21 [DEBUG] agent: Check "service:go-microservice-server" in sync
consul_1                  |     2019/07/20 17:59:21 [DEBUG] agent: Node info in sync
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:25 [DEBUG] agent: Check "service:go-microservice-server" is passing
consul_1                  |     2019/07/20 17:59:25 [DEBUG] agent: Service "go-microservice-server" in sync
consul_1                  |     2019/07/20 17:59:25 [INFO] agent: Synced check "service:go-microservice-server"
consul_1                  |     2019/07/20 17:59:25 [DEBUG] agent: Node info in sync
go-microservice-server_1  | The /info endpoint is being called...
```

When it is done, the go-microservice-client can look it up again, and starts using it:
```
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:26.147175567 +0000 UTC m=+70.003029998
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:30 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:31.147199075 +0000 UTC m=+75.003053507
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:35 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:36.147180431 +0000 UTC m=+80.003034825
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:40 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:41.147174934 +0000 UTC m=+85.003029391
consul_1                  |     2019/07/20 17:59:42 [DEBUG] consul: Skipping self join check for "b0aa3906d740" since the cluster is too small
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:45 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:46.14717465 +0000 UTC m=+90.003029107
go-microservice-server_1  | The /info endpoint is being called...
consul_1                  |     2019/07/20 17:59:50 [DEBUG] agent: Check "service:go-microservice-server" is passing
go-microservice-server_1  | The /info endpoint is being called...
go-microservice-client_1  | Congratulations: you have obtained a bunch of really valuable information after you've called the /info endpoint. Time is 2019-07-20 17:59:51.147096826 +0000 UTC m=+95.002951202
```

## 7 Go microservices discovery just with Kubernetes

In this section we will:

* Deploy a Kubernetes service, and config map definitions 
* Implement the configuration of a client microservice application using the config map, with Go
* Run the Go microservice client and server with Kubernetes

So here no Consul configuration will be involved, we will use just Kubernetes.

First, we need to start Minikube and check our ip:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:04:57 $ minikube start
Starting local Kubernetes v1.10.0 cluster...
Starting VM...
Getting VM IP address...
Moving files into cluster...
Setting up certs...
Connecting to cluster...
Setting up kubeconfig...
Starting cluster components...
Kubectl is now configured to use the cluster.
Loading cached images from config file.

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:06:28 $ minikube ip
192.168.99.100

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:06:48 $ kubectl cluster-info
Kubernetes master is running at https://192.168.99.100:8443
KubeDNS is running at https://192.168.99.100:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

See the Minikube Docker environment:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:08:40 $ minikube docker-env
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.100:2376"
export DOCKER_CERT_PATH="/home/arturotarin/.minikube/certs"
export DOCKER_API_VERSION="1.35"
# Run this command to configure your shell:
# eval $(minikube docker-env)
```

Load Minikube Docker environment credentials:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:08:28 $ eval $(minikube docker-env)
```

Once we have done this, we can build the Docker images for the go-microservice-k8s-server and go-microservice-k8s-client applicatiions in the Minikube Docker environment:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
20:55:24 $ docker-compose build 
Building go-microservice-k8s-server
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&     go get -u github.com/hashicorp/consul/api
 ---> Running in b7961f2d690c
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/community/x86_64/APKINDEX.tar.gz
v3.10.1-11-g89d0862481 [http://dl-cdn.alpinelinux.org/alpine/v3.10/main]
v3.10.1-12-ga885fe876c [http://dl-cdn.alpinelinux.org/alpine/v3.10/community]
OK: 10327 distinct packages available
OK: 6 MiB in 15 packages
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/community/x86_64/APKINDEX.tar.gz
(1/10) Installing ncurses-terminfo-base (6.1_p20190518-r0)
(2/10) Installing ncurses-terminfo (6.1_p20190518-r0)
(3/10) Installing ncurses-libs (6.1_p20190518-r0)
(4/10) Installing readline (8.0.0-r0)
(5/10) Installing bash (5.0.0-r0)
Executing bash-5.0.0-r0.post-install
(6/10) Installing nghttp2-libs (1.38.0-r0)
(7/10) Installing libcurl (7.65.1-r0)
(8/10) Installing expat (2.2.7-r0)
(9/10) Installing pcre2 (10.33-r0)
(10/10) Installing git (2.22.0-r0)
Executing busybox-1.30.1-r2.trigger
OK: 30 MiB in 25 packages
Removing intermediate container b7961f2d690c
 ---> 07f629f1664c
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/server/
 ---> Running in f67902716099
Removing intermediate container f67902716099
 ---> a86b2d6a6451
Step 4/8 : COPY . ${SOURCES}
 ---> 197f77f42ca7
Step 5/8 : RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build -o go-microservice-server
 ---> Running in c7dcf121c0c3
Removing intermediate container c7dcf121c0c3
 ---> 2bfd85a87716
Step 6/8 : ENV CONSUL_HTTP_ADDR localhost:8500
 ---> Running in 16a945e448b4
Removing intermediate container 16a945e448b4
 ---> 2dc945a2ee5c
Step 7/8 : WORKDIR ${SOURCES}server/
Removing intermediate container deac3f2b199b
 ---> d33ccb25064e
Step 8/8 : CMD ${SOURCES}server/go-microservice-server
 ---> Running in ef1b04c082e2
Removing intermediate container ef1b04c082e2
 ---> 4a24c202d685
Successfully built 4a24c202d685
Successfully tagged go-microservice-k8s-server:1.0.0
Building go-microservice-k8s-client
Step 1/8 : FROM golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&     go get -u github.com/hashicorp/consul/api
 ---> Using cache
 ---> 07f629f1664c
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_go/client/
 ---> Running in bf10a15fede8
Removing intermediate container bf10a15fede8
 ---> c38af46c395a
Step 4/8 : COPY . ${SOURCES}
 ---> abc2bf9db624
Step 5/8 : RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build -o go-microservice-client
 ---> Running in f883a8c6ca85
Removing intermediate container f883a8c6ca85
 ---> c3b8e9e02549
Step 6/8 : ENV CONSUL_HTTP_ADDR localhost:8500
 ---> Running in 09f0d81971aa
Removing intermediate container 09f0d81971aa
 ---> 34a26c2b45a1
Step 7/8 : WORKDIR ${SOURCES}client/
Removing intermediate container 50ee8c46ac81
 ---> 1479b3965977
Step 8/8 : CMD ${SOURCES}client/go-microservice-client
 ---> Running in 0fb6ba99fcd6
Removing intermediate container 0fb6ba99fcd6
 ---> 1990f0237f33
Successfully built 1990f0237f33
Successfully tagged go-microservice-k8s-client:1.0.0
```

The images are built:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
20:58:23 $ docker images
REPOSITORY                                      TAG                  IMAGE ID            CREATED             SIZE
go-microservice-k8s-client                      1.0.0                1990f0237f33        2 minutes ago       558MB
go-microservice-k8s-server                      1.0.0                4a24c202d685        2 minutes ago       573MB
go-microservice                                 1.0.0                59a50ecec705        31 hours ago        477MB
nginx                                           latest               98ebf73aba75        2 days ago          109MB
golang                                          1.12-alpine          6b21b4c6e7a3        8 days ago          350MB
consul                                          latest               7d52b83f718f        3 weeks ago         115MB
tomcat                                          9.0                  449eebab16a3        8 months ago        662MB
dummy-dicom-validator                           latest               496aa2ba71c0        9 months ago        333MB
tomcat                                          8.0                  ef6a7c98d192        10 months ago       356MB
perl                                            latest               c58a7ea6dfc4        10 months ago       885MB
goldcar-alpakka-kafka-microservice              latest               810834e8a24b        11 months ago       344MB
openjdk                                         10.0.1-10-jre-slim   6cf6acb97a09        12 months ago       288MB
maven                                           3.5.3-jdk-10-slim    276091e24d4f        13 months ago       596MB
k8s.gcr.io/kube-proxy-amd64                     v1.10.0              bfc21aadc7d3        16 months ago       97MB
k8s.gcr.io/kube-controller-manager-amd64        v1.10.0              ad86dbed1555        16 months ago       148MB
k8s.gcr.io/kube-scheduler-amd64                 v1.10.0              704ba848e69a        16 months ago       50.4MB
k8s.gcr.io/kube-apiserver-amd64                 v1.10.0              af20925d51a3        16 months ago       225MB
k8s.gcr.io/etcd-amd64                           3.1.12               52920ad46f5b        16 months ago       193MB
k8s.gcr.io/kube-addon-manager                   v8.6                 9c16409588eb        17 months ago       78.4MB
prom/prometheus                                 v2.1.0               c8ecf7c719c1        18 months ago       112MB
k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64          1.14.8               c2ce1ffb51ed        18 months ago       41MB
k8s.gcr.io/k8s-dns-sidecar-amd64                1.14.8               6f7f2dc7fab5        18 months ago       42.2MB
k8s.gcr.io/k8s-dns-kube-dns-amd64               1.14.8               80cc5ea4b547        18 months ago       50.5MB
k8s.gcr.io/pause-amd64                          3.1                  da86e6ba6ca1        19 months ago       742kB
gcr.io/google_containers/metrics-server-amd64   v0.2.1               9801395070f3        19 months ago       42.5MB
k8s.gcr.io/kubernetes-dashboard-amd64           v1.8.1               e94d2f21bc0c        19 months ago       121MB
quay.io/coreos/k8s-prometheus-adapter-amd64     v0.2.0               2c0f732478d1        20 months ago       51.9MB
gcr.io/k8s-minikube/storage-provisioner         v1.8.1               4689081edb10        20 months ago       80.8MB
gcr.io/google_containers/kubernetes-kafka       1.0-10.2.1           f9da8ff94c0d        2 years ago         388MB
gcr.io/google_containers/kubernetes-zookeeper   1.0-3.4.10           5586da414c9c        2 years ago         273MB
```

Next, we can deploy the go-microservice-k8s-server application and service in Kubernetes:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:03:43 $ kubectl apply -f go-microservice-k8s-server-deployment.yaml 
deployment "go-microservice-k8s-server" created

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:03:51 $ kubectl apply -f go-microservice-k8s-server-service.yaml 
service "go-microservice-k8s-server" created
```

Have a look at the list of the pod, our server must be there:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:04:34 $ kubectl get pods
NAME                                                 READY     STATUS    RESTARTS   AGE
go-microservice-k8s-server-6d4c6b5475-zrhtp          1/1       Running   0          50s
goldcar-alpakka-kafka-microservice-dc8dbcb9f-bqz9d   1/1       Running   10         332d
kafka-0                                              1/1       Running   17         298d
microservice-in-go-54d84cb66d-4hpbc                  0/1       Pending   0          1d
microservice-in-go-54d84cb66d-ksg68                  0/1       Pending   0          1d
microservice-in-go-54d84cb66d-rchwj                  1/1       Running   1          1d
philips-microservice-6568ccdbd5-ckrsq                1/1       Running   3          292d
task-pv-pod                                          1/1       Running   5          293d
tomcat-56ff5c79c5-lf8fl                              1/1       Running   2          91d
zk-0                                                 1/1       Running   11         332d
```

Make sure that the service is also there:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:04:41 $ kubectl get services
NAME                                         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
go-microservice-k8s-server                   NodePort    10.96.74.139    <none>        9090:32463/TCP      1m
goldcar-alpakka-kafka-microservice-service   NodePort    10.99.105.212   <none>        8080:32000/TCP      332d
kafka-hs                                     ClusterIP   None            <none>        9093/TCP            332d
kubernetes                                   ClusterIP   10.96.0.1       <none>        443/TCP             332d
microservice-in-go                           NodePort    10.103.95.38    <none>        9090:31489/TCP      1d
philips-microservice-service                 NodePort    10.106.207.39   <none>        8083:31000/TCP      292d
tomcat                                       NodePort    10.105.69.168   <none>        8080:31723/TCP      91d
tomcat-service                               NodePort    10.103.30.100   <none>        8080:31234/TCP      91d
zk-cs                                        ClusterIP   10.99.168.219   <none>        2181/TCP            332d
zk-hs                                        ClusterIP   None            <none>        2888/TCP,3888/TCP   332d
```

So we have our go-microservice-k8s-server running and it is accessible via the DNS name.

Next, specify the config map, where we will define the  service.url as "http://go-microservice-k8s-server:9090/info".

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:05:42 $ kubectl apply -f go-microservice-k8s-configmap.yaml 
configmap "go-microservice-k8s-config" created

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:20:17 $ kubectl get configmap
NAME                         DATA      AGE
go-microservice-k8s-config   1         30s

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:20:21 $ kubectl describe configmap go-microservice-k8s-config
Name:         go-microservice-k8s-config
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"v1","data":{"service.url":"http://go-microservice-k8s-server:9090/info"},"kind":"ConfigMap","metadata":{"annotations":{},"name":"go-micr...

Data
====
service.url:
----
http://go-microservice-k8s-server:9090/info
Events:  <none>
```

This variable will be reacheable later on by the go-microservice-k8s-client appilcation via the
SERVICE_URL environment variable.

Next, deploy the go-microservice-k8s-client application:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:20:45 $ kubectl apply -f go-microservice-k8s-client-deployment.yaml 
deployment "go-microservice-k8s-client" created
```

As we can see with this command, the client is using the SERVICE_URL environment variable:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_discovery/with_kubernetes
21:23:35 $ kubectl describe deployment go-microservice-k8s-client
Name:                   go-microservice-k8s-client
Namespace:              default
CreationTimestamp:      Sat, 20 Jul 2019 21:22:40 +0200
Labels:                 io.kompose.service=go-microservice-k8s-client
Annotations:            deployment.kubernetes.io/revision=1
                        kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"extensions/v1beta1","kind":"Deployment","metadata":{"annotations":{},"name":"go-microservice-k8s-client","namespace":"default"},"spec":{...
Selector:               io.kompose.service=go-microservice-k8s-client
Replicas:               1 desired | 1 updated | 1 total | 0 available | 1 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  1 max unavailable, 1 max surge
Pod Template:
  Labels:  io.kompose.service=go-microservice-k8s-client
  Containers:
   go-microservice-k8s-client:
    Image:  go-microservice-k8s-client:1.0.0
    Port:   <none>
    Environment:
      SERVICE_URL:  <set to the key 'service.url' of config map 'go-microservice-k8s-config'>  Optional: false
    Mounts:         <none>
  Volumes:          <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    ReplicaSetUpdated
OldReplicaSets:  <none>
NewReplicaSet:   go-microservice-k8s-client-57f9974b9 (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  1m    deployment-controller  Scaled up replica set go-microservice-k8s-client-57f9974b9 to 1
```


