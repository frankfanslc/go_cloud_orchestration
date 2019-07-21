package main

import (
	"fmt"
	"time"

	proto "github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf/proto_definition"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// The Greeter API.
type Greeter struct{}

// Hello is a Greeter API method.
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}

func hello(t time.Time, greeter proto.GreeterClient) {
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Arturo, calling at " + t.String()})
	if err != nil {
   	    fmt.Println(err.Error())
		return
	}

	// Print response
	fmt.Printf("%s\n", rsp.Greeting)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings.
	// specify a Hystrix breaker client wrapper here
	service.Init(
	)

	// Create new greeter client and call hello
	greeter := proto.NewGreeterClient("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}

