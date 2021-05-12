package controller

import (
	"errors"
	"n_labels/entity"
	"n_labels/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().InsertDoc(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(nil).Times(1)

	s := New(mockMongoClient)

	status, err := s.Create("label1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if status != true {
		t.Errorf("expected true but got %v", status)
	}
}

func TestCreateError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().InsertDoc(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(errors.New("error")).Times(1)

	s := New(mockMongoClient)

	_, err := s.Create("label1", "namespace1")

	if err == nil {
		t.Errorf("expected %s", "errors")
	}
}

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().DeleteDocByID(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(true, nil).Times(1)

	s := New(mockMongoClient)

	status, err := s.Delete("label1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if status != true {
		t.Errorf("expected true but got %v", status)
	}
}

func TestDeleteError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().DeleteDocByID(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(false, errors.New("error")).Times(1)

	s := New(mockMongoClient)

	status, err := s.Delete("label1", "namespace1")

	if err == nil {
		t.Errorf("expected %s", "errors")
	}

	if status != false {
		t.Errorf("expected false but got %v", status)
	}
}

func TestList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	l := []entity.Label{}
	l = append(l, entity.Label{Name: "label1"})

	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		SetArg(2, l).
		Times(1)

	s := New(mockMongoClient)

	labels, err := s.List("name", "label1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if len(labels) != 1 {
		t.Errorf("expected 1 but got %d", len(labels))
	}
}

func TestAttach(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	mockMongoClient.EXPECT().InsertDoc(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		Times(1)

	s := New(mockMongoClient)

	status, err := s.Attach("label1", "entity1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if status != true {
		t.Errorf("expected true but got %v", status)
	}
}

func TestAttachError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	mockMongoClient.EXPECT().InsertDoc(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(errors.New("error")).
		Times(1)

	s := New(mockMongoClient)

	status, err := s.Attach("label1", "entity1", "namespace1")

	if err == nil {
		t.Errorf("expected %s", "errors")
	}

	if status != false {
		t.Errorf("expected false but got %v", status)
	}
}

func TestDetach(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	mockMongoClient.EXPECT().DeleteDocByID(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(true, nil).
		Times(1)

	s := New(mockMongoClient)

	status, err := s.Detach("label1", "entity1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if status != true {
		t.Errorf("expected true but got %v", status)
	}
}

func TestDetachError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	mockMongoClient.EXPECT().DeleteDocByID(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(false, errors.New("error")).
		Times(1)

	s := New(mockMongoClient)

	status, err := s.Detach("label1", "entity1", "namespace1")

	if err == nil {
		t.Errorf("expected %s", "errors")
	}

	if status != false {
		t.Errorf("expected false but got %v", status)
	}
}

func TestGetEntities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	l := []entity.LabelEntity{}
	l = append(l, entity.LabelEntity{Name: "label1"})

	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		SetArg(2, l).
		Times(1)

	s := New(mockMongoClient)

	entities, err := s.GetEntities("label1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if len(entities) != 1 {
		t.Errorf("expected 1 but got %d", len(entities))
	}
}

func TestGetLabels(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)

	l := []entity.LabelEntity{}
	l = append(l, entity.LabelEntity{Name: "label1"})

	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		SetArg(2, l).
		Times(1)

	s := New(mockMongoClient)

	labels, err := s.GetLabels("entity1", "namespace1")

	if err != nil {
		t.Errorf("expected %s", "no errors")
	}

	if len(labels) != 1 {
		t.Errorf("expected 1 but got %d", len(labels))
	}
}
