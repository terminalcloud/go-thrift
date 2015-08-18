package service

import (
	types "github.com/terminalcloud/go-thrift/example"
	"log"
)

type Handler struct {
	count int16

	ichan chan<- interface{}
	ochan <-chan *int16
}

func NewHandler() *Handler {
	ichan := make(chan interface{})
	ochan := make(chan *int16)

	h := Handler{
		count: 0,
		ichan: ichan,
		ochan: ochan,
	}

	// Processing loop to avoid race conditions from concurrent service function calls
	go func() {
		for {
			select {
			// Larger implementations should be moved to their own files
			case <-ichan:
				h.count++
				ochan <- &h.count
			}
		}
	}()

	return &h
}

func (h *Handler) Ping() error {
	log.Println("ping")
	return nil
}

func (h *Handler) Count() (count int16, err error) {
	log.Println("count")
	// Count relies on synchronous data access, so it is delegated to a chan
	h.ichan <- nil
	return *<-h.ochan, nil
}

func (h *Handler) Echo(str string) (string, error) {
	log.Println("echo")
	return str, nil
}

func (h *Handler) Flip(orig *types.Flop) (*types.Flop, error) {
	log.Println("flip")
	f := types.Flop{A: orig.B, B: orig.A}
	return &f, nil
}

func (h *Handler) Fail() error {
	log.Println("fail")
	e := types.ExampleException{}
	return &e
}
