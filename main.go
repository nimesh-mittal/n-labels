package main

import (
	"go.uber.org/zap"
	"n_labels/handler"
	"n_labels/server"
)

func initLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}

func main() {
	initLogger()

	h := handler.New()

	s := server.New()
	s.Mount("/labels", h.NewLabelHandler())

	s.StartServer(":8084")
}
