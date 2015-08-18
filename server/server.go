package server

import "git.apache.org/thrift.git/lib/go/thrift"

type Server struct {
	Addr   string
	wait   chan error
	Wait   <-chan error
	server thrift.TServer
}

func New(addr string, processor *thrift.TProcessor) (*Server, error) {
	// Instantiate the thrift server components
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return nil, err
	}

	wait := make(chan error)

	s := Server{Addr: addr, wait: wait, Wait: wait}
	s.server = thrift.NewTSimpleServer4(*processor, transport, transportFactory, protocolFactory)

	return &s, nil
}

func (s *Server) Run() error {
	return s.server.Serve()
}

func (s *Server) Start() {
	go func() {
		s.wait <- s.server.Serve()
	}()
}
