package handler

import (
	"encoding/json"
	"fmt"
	"n_labels/controller"
	"n_labels/entity"
	"n_labels/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func GetCreateLabelRequest() *http.Request {
	data := entity.CreateLabelRequest{
		Name:      "label1",
		Namespace: "global",
	}
	b, _ := json.Marshal(data)
	payload := strings.NewReader(string(b))
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8084/", payload)
	return req
}

func GetDeleteLabelRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8084/label1", nil)
	return req
}

func GetListLabelRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8084/label1", nil)
	return req
}

func GetAttachLabelRequest() *http.Request {
	data := entity.AttachLabelRequest{
		EntityID:  "e1",
		Namespace: "global",
	}
	b, _ := json.Marshal(data)
	payload := strings.NewReader(string(b))
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8084/label1/_attach", payload)
	return req
}

func GetDetachLabelRequest() *http.Request {
	data := entity.DetachLabelRequest{
		EntityID:  "e1",
		Namespace: "global",
	}
	b, _ := json.Marshal(data)
	payload := strings.NewReader(string(b))
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8084/label1/_detach", payload)
	return req
}

func GetGetEntitiesLabelRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8084/label1/_entities", nil)
	return req
}

func GetGetLabelsLabelRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8084/fetch/e1", nil)
	return req
}

func GetMockCreateLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().InsertDoc(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockDeleteLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().DeleteDocByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(true, nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockListLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockAttachLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().InsertDoc(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockDetachLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().DeleteDocByID(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(true, nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockGetEntitiesLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func GetMockGetLabelsLabelHandler(t *testing.T) LabelHandler {
	mockCtrl := gomock.NewController(t)

	mockMongoClient := mocks.NewMockMongoClient(mockCtrl)
	mockMongoClient.EXPECT().ListDocs(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).
		Return(nil).
		Times(1)

	return &labelHandler{LabelService: controller.New(mockMongoClient)}
}

func TestCreateLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockCreateLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetCreateLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("create label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestDeleteLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockDeleteLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetDeleteLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("delete label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestListLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockListLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetListLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("list label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestAttachLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockAttachLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetAttachLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("attach label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestDetachLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockDetachLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetDetachLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("detach label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestGetEntitiesLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockGetEntitiesLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetGetEntitiesLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetEntities label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}

func TestGetLabelsLabel(t *testing.T) {

	w := httptest.NewRecorder()

	GetMockGetLabelsLabelHandler(t).NewLabelRouter().ServeHTTP(w, GetGetLabelsLabelRequest())
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetLabels label didn’t respond 200 OK: %s", resp.Status)
	}

	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Body)
}
