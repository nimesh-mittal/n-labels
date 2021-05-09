package controller

import (
	"go.uber.org/zap"
	"n_labels/entity"
	"n_labels/gateway"
)

// LabelService represent API interface for labels
type LabelService interface {
	Create(name string, namespace string) (bool, error)
	Delete(name string, namespace string) (bool, error)
	List(filterField string, filterValue string, namespace string) ([]entity.Label, error)
	Attach(labelName string, entityID string, namespace string) (bool, error)
	Detach(labelName string, entityID string, namespace string) (bool, error)
	GetEntities(labelName string, namespace string) ([]string, error)
	GetLabels(entityID string, namespace string) ([]entity.Label, error)
}
const mongoDB = "labels_db"
const mongoColl = "label_col"
const mongoLEColl = "label_entity_col"

// New creates new object of LabelService
func New(db gateway.MongoClient) LabelService {
	return &service{mongoDB: db}
}

type service struct {
	mongoDB gateway.MongoClient
}

func (s *service) Create(name string, namespace string) (bool, error) {
	zap.L().Info("receive create label request",
		zap.String("name", name),
		zap.String("namespace", namespace))

	// TODO: add remaining fields
	record := entity.Label{
		Name:      name,
		Namespace: namespace,
	}

	err := s.mongoDB.InsertDoc(mongoDB, mongoColl, record)

	if err != nil {
		zap.L().Error("error processing created label request", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (s *service) Delete(name string, namespace string) (bool, error) {
	zap.L().Info("receive delete label request",
		zap.String("name", name),
		zap.String("namespace", namespace))
	filter := map[string]interface{}{
		"name":      name,
		"namespace": namespace,
	}
	return s.mongoDB.DeleteDocByID(mongoDB, mongoColl, filter)
}

func (s *service) List(filterField string, filterValue string, namespace string) ([]entity.Label, error) {
	labels := []entity.Label{}
	filter := map[string]interface{}{filterField: filterValue, "namespace": namespace}
	err := s.mongoDB.ListDocs(mongoDB, mongoColl, &labels, filter, 10, 0)
	return labels, err
}

func (s *service) Attach(labelName string, entityID string, namespace string) (bool, error) {
	zap.L().Info("receive attach label request",
		zap.String("name", labelName),
		zap.String("entityId", entityID),
		zap.String("namespace", namespace))
	record := entity.LabelEntity{Name: labelName, EntityID: entityID, Namespace: namespace}
	err := s.mongoDB.InsertDoc(mongoDB, mongoLEColl, record)
	if err != nil {
		zap.L().Error("error processing", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (s *service) Detach(labelName string, entityID string, namespace string) (bool, error) {
	zap.L().Info("receive detach label request",
		zap.String("name", labelName),
		zap.String("entityId", entityID),
		zap.String("namespace", namespace))

	filter := map[string]interface{}{
		"namespace": namespace,
		"name":      labelName,
		"entityid":  entityID,
	}

	status, err := s.mongoDB.DeleteDocByID(mongoDB, mongoLEColl, filter)

	if err != nil {
		zap.L().Error("error processing", zap.Error(err))
		return false, err
	}
	return status, nil
}

func (s *service) GetEntities(labelName string, namespace string) ([]string, error) {
	zap.L().Info("receive GetEntities reqyest",
		zap.String("name", labelName),
		zap.String("namespace", namespace))
	results := []entity.LabelEntity{}
	filter := map[string]interface{}{"name": labelName, "namespace": namespace}
	err := s.mongoDB.ListDocs(mongoDB, mongoLEColl, &results, filter, 10, 0)

	res := []string{}
	for _, v := range results {
		res = append(res, v.EntityID)
	}

	return res, err
}

func (s *service) GetLabels(entityID string, namespace string) ([]entity.Label, error) {
	zap.L().Info("receive GetLabels request",
		zap.String("entityId", entityID),
		zap.String("namespace", namespace))

	results := []entity.LabelEntity{}
	filter := map[string]interface{}{"entityid": entityID, "namespace": namespace}
	err := s.mongoDB.ListDocs(mongoDB, mongoLEColl, &results, filter, 10, 0)

	res := []entity.Label{}
	for _, v := range results {
		res = append(res, entity.Label{Name: v.Name, Namespace: v.Namespace})
	}

	return res, err
}
