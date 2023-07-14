package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shamesjen/rpctest1/gen-go/calculator"
)

type CalculatorHandler struct{}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}

func (p *CalculatorHandler) Add(ctx context.Context, num1 int32, num2 int32) (int32, error) {
	return num1 + num2, nil
}

func runServer(addr string) error {
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := NewCalculatorHandler()
	processor := calculator.NewCalculatorProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}

func main() {
	runServer("localhost:9090")
}
