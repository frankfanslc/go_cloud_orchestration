package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {

	http.HandleFunc("/info", info)
	http.ListenAndServe(port(), nil)
}

func info(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("The /info endpoint has been called ...")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "You're calling the go-microservice-k8s-server application," +
		" the discovery and configuration is now being managed by Kubernetes, no Consul involved anymore")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
