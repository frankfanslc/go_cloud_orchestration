package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	registerServiceInConsul()

	fmt.Println("Starting Go microservice server.")
	http.HandleFunc("/info", info)
	http.ListenAndServe(port(), nil)
}

func registerServiceInConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	var registration = new(consulapi.AgentServiceRegistration)

	registration.ID = "go-microservice-server"
	registration.Name = "go-microservice-server"

	address := hostname()
	registration.Address = address
	port, _ := strconv.Atoi(port()[1:len(port())])
	registration.Port = port

	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/info", address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"

	consul.Agent().ServiceRegister(registration)
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The /info endpoint is being called...")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Congratulations: you have obtained a bunch of " +
		"really valuable information after you've called the /info endpoint")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
