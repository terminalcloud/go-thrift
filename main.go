//go:generate thrift --gen go -out . example.thrift

package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/terminalcloud/go-thrift/example"
	"github.com/terminalcloud/go-thrift/example/otherservice"
	"github.com/terminalcloud/go-thrift/example/service"
	"github.com/terminalcloud/go-thrift/server"
	"log"
)

func main() {
	var processor thrift.TProcessor

	// Start Service server
	processor = example.NewServiceProcessor(service.NewHandler())
	serviceServer, err := server.New("localhost:9090", &processor)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting %s server on %s", "Service", serviceServer.Addr)
	serviceServer.Start()

	// Start OtherService server
	processor = example.NewOtherServiceProcessor(otherservice.NewHandler())
	otherServer, err := server.New("localhost:9091", &processor)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting %s server on %s", "OtherService", otherServer.Addr)
	otherServer.Start()

	for {
		var service string
		var server *server.Server
		var err error

		select {
		case err = <-serviceServer.Wait:
			service = "Service"
			server = serviceServer
		case err = <-otherServer.Wait:
			service = "Other"
			server = otherServer
		}

		log.Printf("Restarting %s server on %s: %s", service, server.Addr, err)
		server.Start()
	}
}
