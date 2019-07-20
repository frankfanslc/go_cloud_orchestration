package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"io/ioutil"
	"net/http"
	"time"
)

var url string

func main()  {
	lookupServiceInConsul()

	fmt.Println("Starting Go microservice client.")
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	callTheGreetingFunctionForeverEvery(5*time.Second, client)
}

func lookupServiceInConsul() {
	config := consulapi.DefaultConfig()

	consul, error := consulapi.NewClient(config)
	if error != nil {
		fmt.Println(error)
	}

	services, error := consul.Agent().Services()
	if error != nil {
		fmt.Println(error)
	}

	service := services["go-microservice-server"]
	address := service.Address
	port := service.Port

	url = fmt.Sprintf("http://%s:%v/info", address, port)
}

func callTheGreetingFunctionForeverEvery(duration time.Duration, client *http.Client) {
	for t := range time.Tick(duration) {
		greeting(t, client)
	}
}

func greeting(t time.Time, client *http.Client) {
	// Call the greeter server endpoint
	response,err := client.Get(url)

	// Exit with an error message in case of problems
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the endpoint response
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s. Time is %v\n", body, t)
	fmt.Sprintf("%s. at the time %v\n", body, t)
}
