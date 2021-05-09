package controller

import (
  "n_labels/entity"
  "n_labels/gateway"
  "go.uber.org/zap"
  "os"
)

// Service represent API interface for labels
type LabelService interface{
  Create(name string, namespace string) (bool, error)
  Delete(name string, namespace string) (bool, error)
  List(filterField string, filterValue string, namespace string) ([]entity.Label, error)
  Attach(labelName string, entityId string, namespace string) (bool, error)
  Detach(labelName string, entityId string, namespace string) (bool, error)
  GetEntities(labelName string, namespace string) ([]string, error)
  GetLabels(entityId string, namespace string) ([]entity.Label, error)
}

const MONGO_DB = "labels_db"
const MONGO_COL = "label_col"
const MONGO_LE_COL = "label_entity_col"

func New() LabelService {
  url := os.Getenv("MONGO_URL_VALUE")
  db := gateway.New(url)
  return &service{MongoDB: db}
}

type service struct{
  MongoDB gateway.MongoClient
}

func (s *service) Create(name string, namespace string)(bool, error){
  zap.L().Info("receive create label request",
    zap.String("name", name),
    zap.String("namespace", namespace))

  // TODO: add remaining fields
  record := entity.Label{
    Name: name,
    Namespace: namespace,
  }

  err := s.MongoDB.InsertDoc(MONGO_DB, MONGO_COL, record)

  if err!= nil{
    zap.L().Error("error processing created label request", zap.Error(err))
    return false, err
  }

  return true, nil
}

func (s *service) Delete(name string, namespace string)(bool, error){
  zap.L().Info("receive delete label request",
    zap.String("name", name),
    zap.String("namespace", namespace))
  filter := map[string]interface{}{
    "name": name,
    "namespace": namespace,
  }
  return s.MongoDB.DeleteDocByID(MONGO_DB, MONGO_COL, filter)
}

func (s *service) List(filterField string, filterValue string, namespace string)([]entity.Label, error){
  labels := []entity.Label{}
  filter := map[string]interface{}{filterField:filterValue, "namespace": namespace}
  err := s.MongoDB.ListDocs(MONGO_DB, MONGO_COL, &labels, filter, 10, 0)
  return labels, err
}

func (s *service) Attach(labelName string, entityId string, namespace string)(bool, error){
  zap.L().Info("receive attach label request",
    zap.String("name", labelName),
    zap.String("entityId", entityId),
    zap.String("namespace", namespace))
  record := entity.LabelEntity{Name: labelName, EntityID: entityId, Namespace: namespace}
  err := s.MongoDB.InsertDoc(MONGO_DB, MONGO_LE_COL, record)
  if err != nil{
    zap.L().Error("error processing", zap.Error(err))
    return false, err
  }
  return true, nil
}

func (s *service) Detach(labelName string, entityId string, namespace string)(bool, error){
  zap.L().Info("receive detach label request",
    zap.String("name", labelName),
    zap.String("entityId", entityId),
    zap.String("namespace", namespace))

  filter := map[string]interface{}{
    "namespace": namespace,
    "name": labelName,
    "entityid": entityId,
  }

  status, err := s.MongoDB.DeleteDocByID(MONGO_DB, MONGO_LE_COL, filter)

  if err != nil{
    zap.L().Error("error processing", zap.Error(err))
    return false, err
  }
  return status, nil
}

func (s *service) GetEntities(labelName string, namespace string)([]string, error){
  zap.L().Info("receive GetEntities reqyest",
    zap.String("name", labelName),
    zap.String("namespace", namespace))
  results := []entity.LabelEntity{}
  filter := map[string]interface{}{"name":labelName, "namespace": namespace}
  err := s.MongoDB.ListDocs(MONGO_DB, MONGO_LE_COL, &results, filter, 10, 0)

  res := []string{}
  for _, v := range results{
    res = append(res, v.EntityID)
  }

  return res, err
}

func (s *service) GetLabels(entityId string, namespace string)([]entity.Label, error){
  zap.L().Info("receive GetLabels request",
    zap.String("entityId", entityId),
    zap.String("namespace", namespace))

    results := []entity.LabelEntity{}
  filter := map[string]interface{}{"entityid":entityId, "namespace": namespace}
  err := s.MongoDB.ListDocs(MONGO_DB, MONGO_LE_COL, &results, filter, 10, 0)

  res := []entity.Label{}
  for _, v := range results{
    res = append(res, entity.Label{Name:v.Name, Namespace:v.Namespace})
  }

  return res, err
}