//go:generate thrift --gen go -out . example.thrift

package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/terminalcloud/go-thrift/example"
	"github.com/terminalcloud/go-thrift/example/service"
	"log"
)

func main() {
	// Instantiate the thrift server components
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket("localhost:9090")
	if err != nil {
		log.Fatalf("Error: ", err)
	}

	// Instantiate the Dnat service handler
	handler := service.NewHandler()
	processor := example.NewServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	log.Println("Serving on localhost:9090")
	server.Serve()
}
