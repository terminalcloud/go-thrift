package service

import (
	types "github.com/terminalcloud/go-thrift/example"
	"log"
)

type Handler struct {
	count int16
}

func NewHandler() *Handler {
	h := Handler{0}
	return &h
}

func (h *Handler) Ping() error {
	log.Println("ping")
	return nil
}

func (h *Handler) Count() (count int16, err error) {
	log.Println("count")
	h.count++
	count = h.count
	return
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
