// otherservice's handler implementation is trivial, and so can stay as it is
package otherservice

import (
	"log"
)

type Handler struct{}

func NewHandler() *Handler {
	h := Handler{}
	return &h
}

func (h *Handler) Noop() error {
	log.Println("noop")
	return nil
}
