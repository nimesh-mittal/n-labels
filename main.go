package main

import (
  "fmt"
  "n_labels/entity"
)

func main(){
  fmt.Println("hello labels!!")
  l1 := entity.Label{Namespace: "ns01", Name: "label1", Active: true}
  fmt.Println(l1)
}