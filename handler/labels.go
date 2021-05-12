package handler

import (
	"encoding/json"
	"n_labels/controller"
	"n_labels/entity"
	"n_labels/gateway"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// LabelHandler handles label endpoints
type LabelHandler interface {
	CreateLabel(w http.ResponseWriter, r *http.Request)
	DeleteLabel(w http.ResponseWriter, r *http.Request)
	ListLabel(w http.ResponseWriter, r *http.Request)
	AttachLabel(w http.ResponseWriter, r *http.Request)
	GetLabels(w http.ResponseWriter, r *http.Request)
	GetEntities(w http.ResponseWriter, r *http.Request)
	NewLabelRouter() http.Handler
}

type labelHandler struct {
	LabelService controller.LabelService
}

// NewLabelHandler creates new object of LabelHandler
func NewLabelHandler() LabelHandler {
	url := os.Getenv("MONGO_URL_VALUE")
	db := gateway.New(url)
	return &labelHandler{LabelService: controller.New(db)}
}

// NewLabelRouter constructs new router for label endpoints
func (h *labelHandler) NewLabelRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.CreateLabel)
	r.Delete("/{LabelID}", h.DeleteLabel)
	r.Put("/{LabelID}/_attach", h.AttachLabel)
	r.Put("/{LabelID}/_detach", h.DetachLabel)
	r.Get("/{LabelID}", h.ListLabel)
	r.Get("/{LabelID}/_entities", h.GetEntities)
	r.Get("/fetch/{EntityID}", h.GetLabels)
	return r
}

func (h *labelHandler) CreateLabel(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var createLabelRequest entity.CreateLabelRequest
	err := decoder.Decode(&createLabelRequest)

	if err != nil {
		e := entity.NewError("invalid create label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	status, err := h.LabelService.Create(createLabelRequest.Name, createLabelRequest.Namespace)

	if err != nil {
		// TODO: always return error in json format
		e := entity.NewError("error processing create label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	e := entity.SuccessResponse{Status: strconv.FormatBool(status)}
	res, _ := json.Marshal(e)
	w.Write(res)
}

func (h *labelHandler) DeleteLabel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "LabelID")

	namespace := "global"
	if strings.Contains(id, ":") {
		tokens := strings.Split(id, ":")
		namespace = tokens[0]
		id = tokens[1]
	}

	status, err := h.LabelService.Delete(id, namespace)

	if err != nil {
		e := entity.NewError("error processing delete label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	e := entity.SuccessResponse{Status: strconv.FormatBool(status)}
	res, _ := json.Marshal(e)
	w.Write(res)
}

func (h *labelHandler) ListLabel(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "LabelID")
	field := "name"
	namespace := "global"
	if strings.Contains(query, ":") {
		tokens := strings.Split(query, ":")
		namespace = tokens[0]
		query = tokens[1]
	}

	labels, err := h.LabelService.List(field, query, namespace)

	if err != nil {
		e := entity.NewError("error processing list label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(labels)
	w.Write(res)
}

func (h *labelHandler) AttachLabel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "LabelID")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var attachLabelRequest entity.AttachLabelRequest
	err := decoder.Decode(&attachLabelRequest)

	if err != nil {
		e := entity.NewError("invalid attach label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	status, err := h.LabelService.Attach(id, attachLabelRequest.EntityID, attachLabelRequest.Namespace)

	if err != nil {
		e := entity.NewError("error processing attach label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	e := entity.SuccessResponse{Status: strconv.FormatBool(status)}
	res, _ := json.Marshal(e)
	w.Write(res)
}

func (h *labelHandler) DetachLabel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "LabelID")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var detachLabelRequest entity.DetachLabelRequest
	err := decoder.Decode(&detachLabelRequest)

	if err != nil {
		e := entity.NewError("invalid detach label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	status, err := h.LabelService.Detach(id, detachLabelRequest.EntityID, detachLabelRequest.Namespace)

	if err != nil {
		e := entity.NewError("error processing detach label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	e := entity.SuccessResponse{Status: strconv.FormatBool(status)}
	res, _ := json.Marshal(e)
	w.Write(res)
}

func (h *labelHandler) GetLabels(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "EntityID")

	namespace := "global"
	if strings.Contains(id, ":") {
		tokens := strings.Split(id, ":")
		namespace = tokens[0]
		id = tokens[1]
	}

	labels, err := h.LabelService.GetLabels(id, namespace)

	if err != nil {
		e := entity.NewError("error processing get labels request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(labels)
	w.Write(res)
}

func (h *labelHandler) GetEntities(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "LabelID")

	namespace := "global"
	if strings.Contains(id, ":") {
		tokens := strings.Split(id, ":")
		namespace = tokens[0]
		id = tokens[1]
	}

	entities, err := h.LabelService.GetEntities(id, namespace)

	if err != nil {
		e := entity.NewError("error processing getEntities for label request")
		res, _ := json.Marshal(e)
		w.Write(res)
		return
	}

	res, _ := json.Marshal(entities)
	w.Write(res)
}
