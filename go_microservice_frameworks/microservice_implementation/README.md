# Microservice implementation

#### Table Of Contents
1. [Document objective](#1-document-objective)
2. [Implemented endpoints](#2-implemented-endpoints)
3. [Containerizing the application with Docker](#3-containerizing-the-application-with-docker)
4. [Exercises](#4-exercises)

## 1 Document objective

In this block we are going to use Golang and the Gin-Gonic framework to:
 
* Implement a basic HTTP microservice service with configurable port
* Implement a basic routing logic for different paths and verbs
* Implement JSON request and response processing

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

Get all the cars:

```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
08:45:14 $ curl localhost:8080/api/cars
[{"id":"0345391802","manufacturer":"Ford","model":"Galaxy"},{"id":"0000000000","manufacturer":"Porsche","model":"Carrera"}]
```

Add a car and check the car has been added: 
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

Modify a car:
```
arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:02:56 $ curl -d '{"id":"00001","manufacturer":"Renault","model":"6"}' -H "Content-Type: application/json" -X PUT curl localhost:8080/api/cars/00001

arturotarin@QOSMIO-X70B:~/go/src/github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_implementation
11:03:32 $ curl localhost:8080/api/cars/00001
{"id":"00001","manufacturer":"Renault","model":"6"}
```

Delete a car:
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

