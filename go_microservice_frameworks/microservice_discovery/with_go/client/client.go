package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {
	lookupServiceWithConsul()

	fmt.Println("Starting Go microservice client.")
	var client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       10*time.Second,
	}

	callTheGreetingFunctionForeverEvery(5*time.Second, client)
}

func lookupServiceWithConsul() {

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
	body,_ := ioutil.ReadAll(response.Body)
	fmt.Printf("Received %s at time %v", body, t)
}
