package controller

import (
  "testing"
  "n_labels/mocks"
  "github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T){
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
  mockMongoClient.EXPECT().InsertDoc(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

  s := New(mockMongoClient)
  
  status, err := s.Create("label1", "namespace1")

  if err != nil{
    t.Errorf("expected %s", "no errors")
  }

  if status != true {
    t.Errorf("expected %s", "status true")
  }
}