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



![](https://yuml.me/diagram/scruffy/class/[API Gateway]<-[Service Discovery],[Service Discovery]<->[Microservices chassis service client],[Service Discovery]<-[Diagnostics],[Service Discovery]<-[Configuration Coordination],[Microservices chassis service client]<-[Configuration Coordination],[Diagnostics]<-[Configuration Coordination],[Diagnostics]->[Microservices chassis service client],[API Gateway]<-[Configuration Coordination] "yUML")

![](http://yuml.me/diagram/scruffy/class/[API Gateway    How to access endpoints from the outside{bg:cornsilk}],[Configuration and coordination    How to provide cluster whide configuration and consensus{bg:cornsilk}],[Service discovery    How to expose and find service endpoints{bg:cornsilk}],[Microservice chassis    How to execute an ops component{bg:cornsilk}],[Microservice chassis    How to call other services in a resilient and responsive way{bg:cornsilk}] "yUML")


````
Diagrams written with yuml.me

See also:

plantuml.com/es
c4model.com
````

![](http://yuml.me/diagram/scruffy/class/[Customer]<>1->*[Order] "yUML")

![](http://yuml.me/diagram/scruffy/class/[API Gateway]<-[Service Discovery],[Service Discovery]<->[Microservices chassis service client],[Service Discovery]<-[Diagnostics],[Service Discovery]<-[Configuration Coordination],[Microservices chassis service client]<-[Configuration Coordination],[Diagnostics]<-[Configuration Coordination],[Diagnostics]->[Microservices chassis service client],[API Gateway]<-[Configuration Coordination] "yUML")


<img src="http://yuml.me/diagram/scruffy/class/[Account]++owner-0..*>[Repository]"/>

<img src="http://yuml.me/diagram/scruffy/class/[API Gateway]<-[Service Discovery],[Service Discovery]<->[Microservices chassis service client],[Service Discovery]<-[Diagnostics],[Service Discovery]<-[Configuration Coordination],[Microservices chassis service client]<-[Configuration Coordination],[Diagnostics]<-[Configuration Coordination],[Diagnostics]->[Microservices chassis service client],[API Gateway]<-[Configuration Coordination]"/>