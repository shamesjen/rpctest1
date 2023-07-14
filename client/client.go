package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shamesjen/rpctest1/gen-go/calculator"
)

var defaultCtx = context.Background()

func handleClient(client *calculator.CalculatorClient) (err error) {
	res, err := client.Add(defaultCtx, 1500, 20)
	if err != nil {
		fmt.Println("Error calling Add:", err)
		return err
	}
	fmt.Println("Result from server:", res)
	return nil
}

func runClient(addr string) error {
	transportSocket, err := thrift.NewTSocket(addr)
	if err != nil {
		return err
	}
	transportFactory := thrift.NewTTransportFactory()
	transport, err := transportFactory.GetTransport(transportSocket)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := calculator.NewCalculatorClientFactory(transport, protocolFactory)

	return handleClient(client)
}


func main() {
	if err := runClient("localhost:9090"); err != nil {
		fmt.Println("Error in client:", err)
	}
}
