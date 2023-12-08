package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go-thrift/gen-go/thrift_demo"
)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) SayHello(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}

func main() {
	handler := &HelloWorldHandler{}
	processor := thrift_demo.NewHelloWorldProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTBufferedTransportFactory(1000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Starting the server...")
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
