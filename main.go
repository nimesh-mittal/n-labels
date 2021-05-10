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

  s := server.New()
	
  hh := handler.NewHealthHandler()
	s.Mount("/", hh.NewHealthRouter())

  lh := handler.NewLabelHandler()
  s.Mount("/labels", lh.NewLabelRouter())

	s.StartServer(":8084")
}
