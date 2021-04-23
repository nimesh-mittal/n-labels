package controller

import (
  "fmt"
  "n_labels/entity"
)

// Service represent API interface for labels
type LabelService interface{
  Create(name string, namespace string) (bool, error)
  Delete(name string, namespace string) (bool, error)
  List(keyword string, namespace string) ([]entity.Label, error)
  Attach(labelName string, entityId string, namespace string) (bool, error)
  GetEntities(labelName string, namespace string) ([]string, error)
  GetLabels(entityId string, namespace string) ([]entity.Label, error)
}

func New() LabelService {
  return &service{}
}

type service struct{
  // TODO: add gateway
}

func (s *service) Create(name string, namespace string)(bool, error){
  fmt.Println("created label")
  return true, nil
}

func (s *service) Delete(name string, namespace string)(bool, error){
  fmt.Println("deleted label")
  return true, nil
}

func (s *service) List(keyword string, namespace string)([]entity.Label, error){
  fmt.Println("list label")
  return []entity.Label{}, nil
}

func (s *service) Attach(labelName string, entityId string, namespace string)(bool, error){
  fmt.Println("attach label")
  return true, nil
}

func (s *service) GetEntities(labelName string, namespace string)([]string, error){
  fmt.Println("list entities associated with label")
  return []string{}, nil
}

func (s *service) GetLabels(entityId string, namespace string)([]entity.Label, error){
  fmt.Println("list labels associated with entity")
  return []entity.Label{}, nil
}