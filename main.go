package main

import (
	"n_labels/handler"
	"n_labels/server"
  "go.uber.org/zap"
)

func initLogger(){
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
