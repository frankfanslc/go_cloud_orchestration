package main

import (
	"fmt"
	"github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// The Greeting api
type Greeting struct {

}

// This is a Greeting api method
func (g * Greeting) Hello(ctx context.Context, req *proto_definition.HelloRequest, rsp *proto_definition.HelloResponse) error {

	rsp.Greeting = req.Name
	fmt.Printf("Responding with the '%s' ProtoBuf service\n", rsp.Greeting)
	return nil
}
func main()  {
	// Create a new service with default flags
	service := micro.NewService(
		micro.Name("greeting"),
		micro.Version("1.0.0"),
		micro.Metadata(map[string]string{
			"type": "hello",
		}),
	)

	// Init parses the command line flags, which will override the default flags
	// written in the NewService initialization
	service.Init()

	// Register handler

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
