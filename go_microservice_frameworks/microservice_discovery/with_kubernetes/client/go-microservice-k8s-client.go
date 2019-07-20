package main

import (

"fmt"
"io/ioutil"
"net/http"
	"os"
	"time"
)

var url string

func main()  {
	initServiceUrl()
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	callTheGreetingFunctionForeverEvery(5*time.Second, client)
}

func initServiceUrl() {
	// Initialize this service url from the SERVICE_URL environment variable
	url := os.Getenv("SERVICE_URL")
	if len(url) == 0 {
		url = "http://go-microservice-k8s-server:9090/info"
	}
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
	fmt.Sprintf("%s. at the time %v\n", body, t)
}
