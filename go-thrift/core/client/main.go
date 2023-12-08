package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go-thrift/gen-go/thrift_demo"
)

func main() {
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		panic(err)
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := thrift_demo.NewHelloWorldClientFactory(transport, protocolFactory)
	result, err := client.SayHello(context.Background())
	if err != nil {
		println(err.Error())
		panic(err)
	}
	fmt.Println(result)
}
