# Cloud orchestration development with Golang

An example of how we can write cloud native applications with Golang.

We will cover:

* [Microservices Implementation](go_cloud_orchestration/microservice_implementation/README.md)
* [Service discovery](go_cloud_orchestration/microservice_discovery/README.md)
* [Configuration](go_cloud_orchestration/microservice_configuration/README.md)
* [Comunication and coordination](go_cloud_orchestration/microservice_communication/README.md)
* Api gateway
* Diagnostics and monitoring 

# Key technologies used in a Cloud Native Application platform


<img src="http://yuml.me/diagram/scruffy/class/[API Gateway]<-[Service Discovery],[Service Discovery]<->[Microservices chassis service client],[Service Discovery]<-[Diagnostics and monitoring],[Service Discovery]<-[Configuration Coordination],[Microservices chassis service client]<-[Configuration Coordination],[Diagnostics and monitoring]<-[Configuration Coordination],[Diagnostics and monitoring]->[Microservices chassis service client],[API Gateway]<-[Configuration Coordination]"/>

<img src="http://yuml.me/diagram/scruffy/class/[API Gateway    How to access endpoints from the outside{bg:cornsilk}],[Configuration and coordination    How to provide cluster whide configuration and consensus{bg:cornsilk}],[Service discovery    How to expose and find service endpoints{bg:cornsilk}],[Microservice chassis    How to execute an ops component{bg:cornsilk}],[Microservice chassis    How to call other services in a resilient and responsive way{bg:cornsilk}],[Diagnostics and monitoring    How to detect operational anomalies{bg:cornsilk}]"/>

# What can we use for our building blocks?

Do not recreate the weel, for reference see the last updated Cloud Native Landscape of the CNCF project (https://github.com/cncf/landscape)

<img src="https://landscape.cncf.io/images/landscape.png">

For the examples shown in this project we will be using Golang integrated with:

* Docker: for containerization
* Kubernetes: for service discovery, configuration and orchestration
* Consul: for service discovery and configuration
* Hystrix: for circuit breaking and monitoring
* Kafka and RabbitMQ: for coordination and communication


**Note**
```
The UML diagram has been written with yuml.me, which doesn't require a server like PlantUML

See also PlantUML at plantuml.com. Although it needs a server, the syntax is much more clear thatn yuml.me as we can see below:

@startuml
Bob->Alice : hello
@enduml

Also c4model.com deserves a try.
```