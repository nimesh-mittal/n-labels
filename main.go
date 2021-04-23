package main

import (
  "n_labels/server"
  "n_labels/handler"
)

func main(){
  h := handler.New()

  s := server.New()
  s.Mount("/labels", h.NewLabelHandler())

  s.StartServer(":8084")
}