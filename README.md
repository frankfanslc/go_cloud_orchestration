# Cloud orchestration development with Golang

An example of how we can write cloud native applications with Golang.

We will cover:

* Microservices
* Service discovery
* Configuration
* Comunication and coordination
* Api gateway
* Diagnostics and monitoring 

# Key technologies used in a Cloud Native Application platform


<img src="http://yuml.me/diagram/scruffy/class/[API Gateway]<-[Service Discovery],[Service Discovery]<->[Microservices chassis service client],[Service Discovery]<-[Diagnostics and monitoring],[Service Discovery]<-[Configuration Coordination],[Microservices chassis service client]<-[Configuration Coordination],[Diagnostics and monitoring]<-[Configuration Coordination],[Diagnostics and monitoring]->[Microservices chassis service client],[API Gateway]<-[Configuration Coordination]"/>

<img src="http://yuml.me/diagram/scruffy/class/[API Gateway    How to access endpoints from the outside{bg:cornsilk}],[Configuration and coordination    How to provide cluster whide configuration and consensus{bg:cornsilk}],[Service discovery    How to expose and find service endpoints{bg:cornsilk}],[Microservice chassis    How to execute an ops component{bg:cornsilk}],[Microservice chassis    How to call other services in a resilient and responsive way{bg:cornsilk}],[Diagnostics and monitoring    How to detect operational anomalies{bg:cornsilk}]"/>

```
Diagrams written with yuml.me

See also:

plantuml.com/es

@startuml
Bob->Alice : hello
@enduml

Also c4model.com
```


# What can we use for our building blocks?

Do not recreate the weel, for reference see the last updated Cloud Native Landscape of the CNCF project (https://github.com/cncf/landscape)

<img src="https://landscape.cncf.io/images/landscape.png">

For this example pupose, we will be using Golang integrated with:

* Docker
* Kubernetes 
* Consul

