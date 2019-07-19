# Microservice implementation

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implemented endpoints](#2-implemented-endpoints)
3. [Containerizing the application with Docker](#3-containerizing-the-application-with-docker)
4. [Containerizing the application with Docker Compose](#4-containerizing-the-application-with-docker-compose)
5. [Pushing the Docker image to Docker Hub](#5-pushing-the-docker-image-to-docker-hub)
6. [Kubernetes orchestration](#6-kubernetes-orchestration)

## 1 Document objective

In this block we are going to use Golang and the Gin-Gonic framework to:
 
* Implement a basic HTTP microservice service with configurable port
* Implement a basic routing logic for different paths and verbs
* Implement JSON request and response processing
* Create Docker images based on these microservices
* Publish them into a Docker hub
* Run Docker containers
* Kubernetes orchestration of our containerized microservices

## 2 Implemented endpoints

We are importing the Gin Gonic framework to do so:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/main
19:44:05 $ go get github.com/gin-gonic/gin
```

Then, we build our Gin server:
 
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:04:42 $ go build

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:08:30 $ ls -rlht
total 15M
-rw-rw-r-- 1 arturotarin arturotarin 1,8K jul 19 08:56 main.go
-rw-rw-r-- 1 arturotarin arturotarin 1,2K jul 19 08:56 cars.go
-rwxrwxr-x 1 arturotarin arturotarin  15M jul 19 09:00 microservice_implementation
-rw-rw-r-- 1 arturotarin arturotarin 2,6K jul 19 09:00 README.md

```
 
Execute the released binary in debug mode:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:09:13 $ ./microservice_implementation 
GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] GET    /hello                    --> main.main.func2 (3 handlers)
[GIN-debug] GET    /api/cars                 --> main.main.func3 (3 handlers)
[GIN-debug] POST   /api/cars                 --> main.main.func4 (3 handlers)
[GIN-debug] GET    /api/cars/:carid          --> main.main.func5 (3 handlers)
[GIN-debug] PUT    /api/cars/:carid          --> main.main.func6 (3 handlers)
[GIN-debug] DELETE /api/cars/:carid          --> main.main.func7 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080

```

Or execute it in release mode if you want to:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:03:30 $ export GIN_MODE=release

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:03:36 $ ./microservice_implementation 
```

Then we can start sending requests to our Gin's /ping endpoint:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
06:50:40 $ curl localhost:8080/ping
pong
```

We can see how they're being logged in our Gin server:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
07:03:36 $ ./main 
[GIN] 2019/07/19 - 07:04:21 | 200 |      49.348µs |       127.0.0.1 | GET      /ping
```

Getting all the cars:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
08:45:14 $ curl localhost:8080/api/cars
[{"id":"0345391802","manufacturer":"Ford","model":"Galaxy"},{"id":"0000000000","manufacturer":"Porsche","model":"Carrera"}]
```

Adding a car and check the car has been added: 
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
09:04:59 $ curl -d '{"id":"00001","manufacturer":"Renault","model":"4"}' -H "Content-Type: application/json" -X POST curl localhost:8080/api/cars

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
09:07:44 $ curl localhost:8080/api/cars
[{"id":"0345391802","manufacturer":"Ford","model":"Galaxy"},{"id":"0000000000","manufacturer":"Porsche","model":"Carrera"},{"id":"00001","manufacturer":"Renault","model":"4"}]

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
09:07:55 $ curl localhost:8080/api/cars/00001
{"id":"00001","manufacturer":"Renault","model":"4"}

```

Modifying a car:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:02:56 $ curl -d '{"id":"00001","manufacturer":"Renault","model":"6"}' -H "Content-Type: application/json" -X PUT curl localhost:8080/api/cars/00001

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:03:32 $ curl localhost:8080/api/cars/00001
{"id":"00001","manufacturer":"Renault","model":"6"}
```

Deleting a car:
```
oservice_implementation
11:04:54 $ curl -X DELETE curl localhost:8080/api/cars/00001curl: (6) Could not resolve host: curl

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:05:45 $ curl localhost:8080/api/cars/00001

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:05:49 $ curl localhost:8080/api/cars
[{"id":"0345391802","manufacturer":"Ford","model":"Galaxy"},{"id":"0000000000","manufacturer":"Porsche","model":"Carrera"}]
```

## 3 Containerizing the application with Docker

Building and tagging the image

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:44:35 $ docker build -t go-microservice:1.0.0 .
Sending build context to Docker daemon  15.65MB
Step 1/8 : FROM golang:1.12-alpine
1.12-alpine: Pulling from library/golang
050382585609: Pull complete 
0bb4ee3360d7: Pull complete 
893f09c2afb0: Pull complete 
db25f79b026e: Pull complete 
4387e72e4ead: Pull complete 
Digest: sha256:1121c345b1489bb5e8a9a65b612c8fed53c175ce72ac1c76cf12bbfc35211310
Status: Downloaded newer image for golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&     go get -u github.com/gin-gonic/gin
 ---> Running in 8a5ed72fc68a
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/community/x86_64/APKINDEX.tar.gz
v3.10.1-11-g89d0862481 [http://dl-cdn.alpinelinux.org/alpine/v3.10/main]
v3.10.1-9-gebc7b05d9e [http://dl-cdn.alpinelinux.org/alpine/v3.10/community]
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
 ---> 752ce4079ad6
Removing intermediate container 8a5ed72fc68a
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation/
 ---> Running in 0acb7d5cdd18
 ---> 58bf1c51b655
Removing intermediate container 0acb7d5cdd18
Step 4/8 : COPY . ${SOURCES}
 ---> 0f7f499eade1
Removing intermediate container 4103d9ec0dca
Step 5/8 : RUN cd ${SOURCES} && CGO_ENABLED=0 go build
 ---> Running in b16c92a89a7f
 ---> 77e705be0ac2
Removing intermediate container b16c92a89a7f
Step 6/8 : WORKDIR ${SOURCES}
 ---> aac714ec0077
Removing intermediate container eebecc28bb4d
Step 7/8 : CMD ${SOURCES}microservice_implementation
 ---> Running in b080b0880f32
 ---> e14cd1edaabb
Removing intermediate container b080b0880f32
Step 8/8 : EXPOSE 8080
 ---> Running in 8229893c58fe
 ---> 7870aec6a444
Removing intermediate container 8229893c58fe
Successfully built 7870aec6a444
Successfully tagged go-microservice:1.0.0

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:47:05 $ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED              SIZE
go-microservice     1.0.0               7870aec6a444        About a minute ago   477MB
golang              1.12-alpine         6b21b4c6e7a3        7 days ago           350MB

```

Running a container based in our Docker image:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:51:56 $ docker run -d --name microservice-in-go go-microservice:1.0.0 
cf1db92e4c2cd04edbb413faa13df43f46456ee3bb8e8b298d262dc23215100c

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:53:15 $ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED              STATUS              PORTS               NAMES
cf1db92e4c2c        go-microservice:1.0.0   "/bin/sh -c ${SOUR..."   About a minute ago   Up About a minute   8080/tcp            microservice-in-go

```

Looking at the container logs:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:55:08 $ docker logs  microservice-in-go --tail 100 -f
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] GET    /hello                    --> main.main.func2 (3 handlers)
[GIN-debug] GET    /api/cars                 --> main.main.func3 (3 handlers)
[GIN-debug] POST   /api/cars                 --> main.main.func4 (3 handlers)
[GIN-debug] Loaded HTML Templates (2): 
	- 
	- index.html

[GIN-debug] GET    /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
[GIN-debug] HEAD   /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
[GIN-debug] GET    /                         --> main.main.func5 (3 handlers)
[GIN-debug] GET    /api/cars/:carid          --> main.main.func6 (3 handlers)
[GIN-debug] PUT    /api/cars/:carid          --> main.main.func7 (3 handlers)
[GIN-debug] DELETE /api/cars/:carid          --> main.main.func8 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080

```

SSHing to the container:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:58:31 $ docker exec -it microservice-in-go bash
bash-5.0# 

bash-5.0# echo $GOLANG_VERSION 
1.12.7

bash-5.0# echo $GOPATH/
/go/

bash-5.0# ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation/microservice_implement
   12 root      0:00 bash
   23 root      0:00 ps -ef

```

Knowing the container ip:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:01:53 $ docker inspect microservice-in-go | grep "IPAddress"
            "IPAddress": "172.17.0.3",
```

Calling the microservice:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:02:06 $ curl 172.17.0.3:8080/api/cars
[{"id":"0345391802","manufacturer":"Ford","model":"Galaxy"},{"id":"0000000000","manufacturer":"Porsche","model":"Carrera"}]
```

Stopping and removing it all:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:03:25 $ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS               NAMES
cf1db92e4c2c        go-microservice:1.0.0   "/bin/sh -c ${SOUR..."   12 minutes ago      Up 12 minutes       8080/tcp            microservice-in-go

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:04:58 $ docker rm -f cf1db92e4c2c
cf1db92e4c2c

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:05:12 $ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
go-microservice     1.0.0               7870aec6a444        18 minutes ago      477MB
golang              1.12-alpine         6b21b4c6e7a3        7 days ago          350MB

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:05:15 $ docker rmi 7870aec6a444
Untagged: go-microservice:1.0.0
Deleted: sha256:7870aec6a44491acd7a279c4f4efd74958c2946eb37a5eae8fb1abaa0b52cb36
Deleted: sha256:e14cd1edaabb53924fbcfdc50816aaa18effdbd3adf4d7184d15d15b0945e718
Deleted: sha256:aac714ec0077c587a3802938972457ac5c56aba970edd6882dc8979e2039c1a9
Deleted: sha256:77e705be0ac285c49f636c62709c5ccc9dd9002f02ef9a296b25beb2db2950ee
Deleted: sha256:205d43464f530f2aef6f521f94ae5a0c1e594cfcd461d27ee8aed7b7d4fd1bc2
Deleted: sha256:0f7f499eade135f850b7a2c846ee3199de03aef3a120f847755d0b715f64172a
Deleted: sha256:6274aa25700050accd3505452b9b2dc1716eca0bbd57f49c6163d72c8577871d
Deleted: sha256:58bf1c51b65512eda48356325f61d9361ac9a14c39419e998673080f95b58774
Deleted: sha256:752ce4079ad66768fe26470bad60b05cece32aa86212b3592b2bcd9094aa67d3
Deleted: sha256:e4ea2e51ccfa7545f541efa3064d333892155fe54f38d5489ad9384ed23a7e62
```

## 4 Containerizing the application with Docker Compose

Building the docker-compose.yml file:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:09:54 $ docker-compose build
Building microservice-in-go
Step 1/8 : FROM golang:1.12-alpine
1.12-alpine: Pulling from library/golang
050382585609: Pull complete
0bb4ee3360d7: Pull complete
893f09c2afb0: Pull complete
db25f79b026e: Pull complete
4387e72e4ead: Pull complete
Digest: sha256:1121c345b1489bb5e8a9a65b612c8fed53c175ce72ac1c76cf12bbfc35211310
Status: Downloaded newer image for golang:1.12-alpine
 ---> 6b21b4c6e7a3
Step 2/8 : RUN apk update && apk upgrade && apk add --no-cache bash git &&     go get -u github.com/gin-gonic/gin
 ---> Running in a0616bba6de5
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.10/community/x86_64/APKINDEX.tar.gz
v3.10.1-11-g89d0862481 [http://dl-cdn.alpinelinux.org/alpine/v3.10/main]
v3.10.1-9-gebc7b05d9e [http://dl-cdn.alpinelinux.org/alpine/v3.10/community]
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
 ---> fbe814ec2773
Removing intermediate container a0616bba6de5
Step 3/8 : ENV SOURCES /go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation/
 ---> Running in ed8122f2894c
 ---> ba43cacfe5e5
Removing intermediate container ed8122f2894c
Step 4/8 : COPY . ${SOURCES}
 ---> 3aa558a253d7
Removing intermediate container a275ae837600
Step 5/8 : RUN cd ${SOURCES} && CGO_ENABLED=0 go build
 ---> Running in 15dfa337c93c
 ---> 3b8bcb09ee38
Removing intermediate container 15dfa337c93c
Step 6/8 : WORKDIR ${SOURCES}
 ---> 6efda4b5dec4
Removing intermediate container 9886e657a219
Step 7/8 : CMD ${SOURCES}microservice_implementation
 ---> Running in e2911828c9bb
 ---> 77aa7ba4b918
Removing intermediate container e2911828c9bb
Step 8/8 : EXPOSE 8080
 ---> Running in 3681d2ecf269
 ---> d47bd5746538
Removing intermediate container 3681d2ecf269
Successfully built d47bd5746538
Successfully tagged go-microservice:1.0.1

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:16:17 $ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
go-microservice     1.0.1               d47bd5746538        4 minutes ago       477MB
golang              1.12-alpine         6b21b4c6e7a3        7 days ago          350MB
```

Running the container:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:16:27 $ docker-compose up -d microservice-in-go
Creating network "microserviceimplementation_default" with the default driver
Creating microserviceimplementation_microservice-in-go_1 ... done

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:21:31 $ docker-compose ps
                     Name                                    Command               State           Ports         
-----------------------------------------------------------------------------------------------------------------
microserviceimplementation_microservice-in-go_1   /bin/sh -c ${SOURCES}micro ...   Up      0.0.0.0:8080->8080/tcp
```

Viewing the container logs:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:22:29 $ docker-compose logs microservice-in-go 
Attaching to microserviceimplementation_microservice-in-go_1
microservice-in-go_1  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
microservice-in-go_1  | 
microservice-in-go_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
microservice-in-go_1  |  - using env:	export GIN_MODE=release
microservice-in-go_1  |  - using code:	gin.SetMode(gin.ReleaseMode)
microservice-in-go_1  | 
microservice-in-go_1  | [GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
microservice-in-go_1  | [GIN-debug] GET    /hello                    --> main.main.func2 (3 handlers)
microservice-in-go_1  | [GIN-debug] GET    /api/cars                 --> main.main.func3 (3 handlers)
microservice-in-go_1  | [GIN-debug] POST   /api/cars                 --> main.main.func4 (3 handlers)
microservice-in-go_1  | [GIN-debug] Loaded HTML Templates (2): 
microservice-in-go_1  | 	- 
microservice-in-go_1  | 	- index.html
microservice-in-go_1  | 
microservice-in-go_1  | [GIN-debug] GET    /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
microservice-in-go_1  | [GIN-debug] HEAD   /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
microservice-in-go_1  | [GIN-debug] GET    /                         --> main.main.func5 (3 handlers)
microservice-in-go_1  | [GIN-debug] GET    /api/cars/:carid          --> main.main.func6 (3 handlers)
microservice-in-go_1  | [GIN-debug] PUT    /api/cars/:carid          --> main.main.func7 (3 handlers)
microservice-in-go_1  | [GIN-debug] DELETE /api/cars/:carid          --> main.main.func8 (3 handlers)
microservice-in-go_1  | [GIN-debug] Listening and serving HTTP on :8080
```

As we have exposed the Docker port 8080 to our localhost, we can still use the same application urls as before.

Deleting everything:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:27:54 $ docker-compose stop
Stopping microserviceimplementation_microservice-in-go_1 ... done

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:28:02 $ docker-compose rm
Going to remove microserviceimplementation_microservice-in-go_1
Are you sure? [yN] y
Removing microserviceimplementation_microservice-in-go_1 ... done
```

## 5 Pushing the Docker image to Docker Hub

Tagging the image:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:35:56 $ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
go-microservice     1.0.1               d47bd5746538        23 minutes ago      477MB
golang              1.12-alpine         6b21b4c6e7a3        7 days ago          350MB

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:35:58 $ docker tag go-microservice:1.0.1 arturot/go-microservice:1.0.1

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:36:35 $ docker images
REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
go-microservice           1.0.1               d47bd5746538        24 minutes ago      477MB
arturot/go-microservice   1.0.1               d47bd5746538        24 minutes ago      477MB
golang                    1.12-alpine         6b21b4c6e7a3        7 days ago          350MB
```

Pushing the image to my Docker Hub account:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
12:36:37 $ docker push arturot/go-microservice
The push refers to a repository [docker.io/arturot/go-microservice]
2a768907face: Pushed 
7fa4dd082742: Pushed 
5f1968c1ed81: Pushed 
95f148bce382: Mounted from library/golang 
58de2ae4cd62: Mounted from library/golang 
e0a42524f665: Mounted from library/golang 
05540d8bb3fd: Mounted from library/golang 
1bfeebd65323: Mounted from library/golang 
1.0.1: digest: sha256:6a44e7dedc8c9fc462de739d4e971232d0186a37512ba2d9edf3d15cae33b031 size: 2000
```

The image has been uploaded to my Docker Hub account, arturot:


![alt text](images/image01.png "My Docker hub")


## 6 Kubernetes orchestration

Start Minikube and check our ip:

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
Check which images do I have there:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:09:20 $ docker images
REPOSITORY                                      TAG                  IMAGE ID            CREATED             SIZE
nginx                                           latest               98ebf73aba75        36 hours ago        109MB
nginx                                           <none>               27a188018e18        3 months ago        109MB
tomcat                                          9.0                  449eebab16a3        8 months ago        662MB
dummy-dicom-validator                           latest               496aa2ba71c0        9 months ago        333MB
<none>                                          <none>               1474e47af2a8        9 months ago        333MB
<none>                                          <none>               cbc5c5695bcd        9 months ago        333MB
<none>                                          <none>               b14931750f7d        9 months ago        288MB
<none>                                          <none>               51bcf101f911        9 months ago        288MB
<none>                                          <none>               39daf8f1a880        9 months ago        333MB
<none>                                          <none>               72313623dfca        9 months ago        718MB
<none>                                          <none>               3e425f5d2241        9 months ago        333MB
<none>                                          <none>               7a026463e021        9 months ago        718MB
<none>                                          <none>               4e03acded027        9 months ago        333MB
<none>                                          <none>               9e2c116693c7        9 months ago        717MB
<none>                                          <none>               e0a02a0ee4ce        9 months ago        333MB
<none>                                          <none>               a65ab992d795        9 months ago        717MB
<none>                                          <none>               b24c4b3c4df5        9 months ago        333MB
<none>                                          <none>               d468ec1d7e60        9 months ago        717MB
<none>                                          <none>               bec84c485ef8        9 months ago        717MB
<none>                                          <none>               e1f1c0942efd        9 months ago        333MB
<none>                                          <none>               af0c78ecef5d        9 months ago        333MB
<none>                                          <none>               ba28e4eefcd5        9 months ago        718MB
<none>                                          <none>               026bf60b3641        9 months ago        333MB
<none>                                          <none>               dfad3ed5f86f        9 months ago        718MB
nginx                                           <none>               bc26f1ed35cf        9 months ago        109MB
<none>                                          <none>               e30ec50258a8        9 months ago        333MB
<none>                                          <none>               05444eb9f831        9 months ago        718MB
<none>                                          <none>               3590abde12d0        9 months ago        596MB
<none>                                          <none>               2dade5299937        9 months ago        718MB
<none>                                          <none>               b659c5f7e59d        9 months ago        596MB
tomcat                                          8.0                  ef6a7c98d192        10 months ago       356MB
perl                                            latest               c58a7ea6dfc4        10 months ago       885MB
nginx                                           <none>               06144b287844        10 months ago       109MB
goldcar-alpakka-kafka-microservice              latest               810834e8a24b        11 months ago       344MB
<none>                                          <none>               31fe42030366        11 months ago       739MB
<none>                                          <none>               55efa18f365a        11 months ago       344MB
<none>                                          <none>               1d3f106a0913        11 months ago       739MB
<none>                                          <none>               2ca0750d2abf        11 months ago       344MB
<none>                                          <none>               c4a9284e2e8e        11 months ago       739MB
<none>                                          <none>               9fde9189faeb        11 months ago       344MB
openjdk                                         10.0.1-10-jre-slim   6cf6acb97a09        12 months ago       288MB
maven                                           3.5.3-jdk-10-slim    276091e24d4f        13 months ago       596MB
k8s.gcr.io/kube-proxy-amd64                     v1.10.0              bfc21aadc7d3        15 months ago       97MB
k8s.gcr.io/kube-scheduler-amd64                 v1.10.0              704ba848e69a        15 months ago       50.4MB
k8s.gcr.io/kube-controller-manager-amd64        v1.10.0              ad86dbed1555        15 months ago       148MB
k8s.gcr.io/kube-apiserver-amd64                 v1.10.0              af20925d51a3        15 months ago       225MB
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

Check what is running there:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:10:21 $ docker ps
CONTAINER ID        IMAGE                                           COMMAND                  CREATED             STATUS              PORTS               NAMES
4f36c5b66857        gcr.io/google_containers/metrics-server-amd64   "/metrics-server -..."   2 minutes ago       Up 2 minutes                            k8s_metrics-server_metrics-server-6fbfb84cdd-zxdl4_kube-system_a7dbd7ca-a5cb-11e8-9b6b-08002789fefc_20
2094e306d1cb        2c0f732478d1                                    "/adapter /adapter..."   2 minutes ago       Up 2 minutes                            k8s_custom-metrics-apiserver_custom-metrics-apiserver-7dd968d85-jn2st_monitoring_2c7b47de-a5cc-11e8-9b6b-08002789fefc_22
4a83020d32a2        e94d2f21bc0c                                    "/dashboard --inse..."   2 minutes ago       Up 2 minutes                            k8s_kubernetes-dashboard_kubernetes-dashboard-5498ccf677-9ml9p_kube-system_da0b4a48-a5c9-11e8-9b6b-08002789fefc_23
693db9958173        4689081edb10                                    "/storage-provisioner"   2 minutes ago       Up 2 minutes                            k8s_storage-provisioner_storage-provisioner_kube-system_da1c47cb-a5c9-11e8-9b6b-08002789fefc_22
15be3463f15e        80cc5ea4b547                                    "/kube-dns --domai..."   3 minutes ago       Up 3 minutes                            k8s_kubedns_kube-dns-86f4d74b45-7nhjk_kube-system_d93588fd-a5c9-11e8-9b6b-08002789fefc_18
45fb9d42adb6        gcr.io/google_containers/kubernetes-kafka       "sh -c 'exec kafka..."   3 minutes ago       Up 3 minutes                            k8s_k8skafka_kafka-0_default_02d95247-c07a-11e8-bfaa-08002789fefc_14
c78388c89c5d        bfc21aadc7d3                                    "/usr/local/bin/ku..."   3 minutes ago       Up 3 minutes                            k8s_kube-proxy_kube-proxy-s8mqd_kube-system_5ef116aa-aa15-11e9-a91a-08002789fefc_0
7ef95d1dd2ca        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 3 minutes ago       Up 3 minutes                            k8s_POD_kube-proxy-s8mqd_kube-system_5ef116aa-aa15-11e9-a91a-08002789fefc_0
e15043eb7522        nginx                                           "nginx -g 'daemon ..."   3 minutes ago       Up 3 minutes                            k8s_task-pv-container_task-pv-pod_default_e6f0a219-c484-11e8-91a9-08002789fefc_4
7134ab17431b        496aa2ba71c0                                    "java -jar dicom-v..."   4 minutes ago       Up 4 minutes                            k8s_philips-microservice_philips-microservice-6568ccdbd5-7b9cn_default_3b6775f3-c53e-11e8-9acb-08002789fefc_2
4e8e04617759        449eebab16a3                                    "catalina.sh run"        4 minutes ago       Up 4 minutes                            k8s_tomcat_tomcat-56ff5c79c5-znl24_default_b6a854b7-633a-11e9-8624-08002789fefc_1
7ea12830cfd2        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_philips-microservice-6568ccdbd5-7b9cn_default_3b6775f3-c53e-11e8-9acb-08002789fefc_2
4d49a7845059        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_tomcat-56ff5c79c5-znl24_default_b6a854b7-633a-11e9-8624-08002789fefc_1
9a8af4d4c026        496aa2ba71c0                                    "java -jar dicom-v..."   4 minutes ago       Up 4 minutes                            k8s_philips-microservice_philips-microservice-6568ccdbd5-ff2bf_default_3b6686da-c53e-11e8-9acb-08002789fefc_2
49f8ab057a4a        gcr.io/google_containers/kubernetes-zookeeper   "sh -c 'start-zook..."   4 minutes ago       Up 4 minutes                            k8s_kubernetes-zookeeper_zk-0_default_92098a98-a5cf-11e8-9b6b-08002789fefc_10
1a1db654b0b3        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_task-pv-pod_default_e6f0a219-c484-11e8-91a9-08002789fefc_4
5a475887990f        449eebab16a3                                    "catalina.sh run"        4 minutes ago       Up 4 minutes                            k8s_tomcat_tomcat-56ff5c79c5-lf8fl_default_b6a93644-633a-11e9-8624-08002789fefc_1
29b86d03f51f        prom/prometheus                                 "prometheus --conf..."   4 minutes ago       Up 4 minutes                            k8s_prometheus_prometheus-7dff795b9f-l2xdg_monitoring_1e77dde3-a5cc-11e8-9b6b-08002789fefc_9
da462895ff52        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kafka-0_default_02d95247-c07a-11e8-bfaa-08002789fefc_6
12adbc4ac4e1        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_philips-microservice-6568ccdbd5-ff2bf_default_3b6686da-c53e-11e8-9acb-08002789fefc_2
d8696a2fcc8a        496aa2ba71c0                                    "java -jar dicom-v..."   4 minutes ago       Up 4 minutes                            k8s_philips-microservice_philips-microservice-6568ccdbd5-ckrsq_default_3b677acc-c53e-11e8-9acb-08002789fefc_2
7a50550643d5        810834e8a24b                                    "java -jar goldcar..."   4 minutes ago       Up 4 minutes                            k8s_goldcar-alpakka-kafka-microservice_goldcar-alpakka-kafka-microservice-dc8dbcb9f-mhf2d_default_04266ca2-a5fb-11e8-9b6b-08002789fefc_9
182efd30d5de        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_tomcat-56ff5c79c5-lf8fl_default_b6a93644-633a-11e9-8624-08002789fefc_1
86e28fe71dbb        810834e8a24b                                    "java -jar goldcar..."   4 minutes ago       Up 4 minutes                            k8s_goldcar-alpakka-kafka-microservice_goldcar-alpakka-kafka-microservice-dc8dbcb9f-2sxw2_default_041b7ef6-a5fb-11e8-9b6b-08002789fefc_9
2446725d5a2a        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_philips-microservice-6568ccdbd5-ckrsq_default_3b677acc-c53e-11e8-9acb-08002789fefc_2
5449ea5fb0b8        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_zk-0_default_92098a98-a5cf-11e8-9b6b-08002789fefc_9
b13d676e9a5d        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_goldcar-alpakka-kafka-microservice-dc8dbcb9f-mhf2d_default_04266ca2-a5fb-11e8-9b6b-08002789fefc_9
df90235484ae        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_goldcar-alpakka-kafka-microservice-dc8dbcb9f-2sxw2_default_041b7ef6-a5fb-11e8-9b6b-08002789fefc_9
00fb9ac129c7        6f7f2dc7fab5                                    "/sidecar --v=2 --..."   4 minutes ago       Up 4 minutes                            k8s_sidecar_kube-dns-86f4d74b45-7nhjk_kube-system_d93588fd-a5c9-11e8-9b6b-08002789fefc_10
8425470ae489        810834e8a24b                                    "java -jar goldcar..."   4 minutes ago       Up 4 minutes                            k8s_goldcar-alpakka-kafka-microservice_goldcar-alpakka-kafka-microservice-dc8dbcb9f-bqz9d_default_0425a088-a5fb-11e8-9b6b-08002789fefc_9
f4958f7e9637        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_storage-provisioner_kube-system_da1c47cb-a5c9-11e8-9b6b-08002789fefc_9
99e35930926a        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_custom-metrics-apiserver-7dd968d85-jn2st_monitoring_2c7b47de-a5cc-11e8-9b6b-08002789fefc_9
e74b858df46c        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_prometheus-7dff795b9f-l2xdg_monitoring_1e77dde3-a5cc-11e8-9b6b-08002789fefc_9
55b32df04098        c2ce1ffb51ed                                    "/dnsmasq-nanny -v..."   4 minutes ago       Up 4 minutes                            k8s_dnsmasq_kube-dns-86f4d74b45-7nhjk_kube-system_d93588fd-a5c9-11e8-9b6b-08002789fefc_10
160e7bfec303        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_goldcar-alpakka-kafka-microservice-dc8dbcb9f-bqz9d_default_0425a088-a5fb-11e8-9b6b-08002789fefc_9
7b0af6e6a898        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_metrics-server-6fbfb84cdd-zxdl4_kube-system_a7dbd7ca-a5cb-11e8-9b6b-08002789fefc_9
91acf725c297        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kubernetes-dashboard-5498ccf677-9ml9p_kube-system_da0b4a48-a5c9-11e8-9b6b-08002789fefc_9
b0b983c5e716        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-dns-86f4d74b45-7nhjk_kube-system_d93588fd-a5c9-11e8-9b6b-08002789fefc_9
afa6b6991b6e        ad86dbed1555                                    "kube-controller-m..."   4 minutes ago       Up 4 minutes                            k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_eb389aa32974e264d14fcc4deed664ed_0
7a21118650f9        704ba848e69a                                    "kube-scheduler --..."   4 minutes ago       Up 4 minutes                            k8s_kube-scheduler_kube-scheduler-minikube_kube-system_31cf0ccbee286239d451edb6fb511513_0
fbb564b82d79        9c16409588eb                                    "/opt/kube-addons.sh"    4 minutes ago       Up 4 minutes                            k8s_kube-addon-manager_kube-addon-manager-minikube_kube-system_3afaf06535cc3b85be93c31632b765da_9
811eb6da2e91        af20925d51a3                                    "kube-apiserver --..."   4 minutes ago       Up 4 minutes                            k8s_kube-apiserver_kube-apiserver-minikube_kube-system_7ae7fb7d367808dd5e0f3be3047b2b7a_0
bc820d50c3d5        52920ad46f5b                                    "etcd --peer-key-f..."   4 minutes ago       Up 4 minutes                            k8s_etcd_etcd-minikube_kube-system_b0aca6fc6a477d1a0666c5a4c2697f5d_0
bb8c34798f88        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-controller-manager-minikube_kube-system_eb389aa32974e264d14fcc4deed664ed_0
af9df598bd19        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-scheduler-minikube_kube-system_31cf0ccbee286239d451edb6fb511513_0
42a598dbb5ce        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-apiserver-minikube_kube-system_7ae7fb7d367808dd5e0f3be3047b2b7a_0
b2c99aa7f3ff        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_etcd-minikube_kube-system_b0aca6fc6a477d1a0666c5a4c2697f5d_0
1cb79dd2ddaf        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-addon-manager-minikube_kube-system_3afaf06535cc3b85be93c31632b765da_9
```

Open Minikube UI:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
13:12:31 $ minikube dashboard
Opening kubernetes dashboard in default browser...
```

![alt text](images/image02.png "My Minikube UI")

 