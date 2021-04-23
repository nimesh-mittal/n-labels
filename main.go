package main

import (
	"n_labels/handler"
	"n_labels/server"
)

func main() {
	h := handler.New()

	s := server.New()
	s.Mount("/labels", h.NewLabelHandler())

	s.StartServer(":8084")
}
