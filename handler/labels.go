package handler

import (
  "n_labels/controller"
  "net/http"
  "github.com/go-chi/chi/v5"
)

type LabelHandler interface{
  CreateLabel(w http.ResponseWriter, r *http.Request)
  DeleteLabel(w http.ResponseWriter, r *http.Request)
  ListLabel(w http.ResponseWriter, r *http.Request)
  AttachLabel(w http.ResponseWriter, r *http.Request)
  GetLabels(w http.ResponseWriter, r *http.Request)
  GetEntities(w http.ResponseWriter, r *http.Request)
  NewLabelHandler() http.Handler
}

type labelHandler struct{
  LabelService controller.LabelService
}

func New() LabelHandler{
  return &labelHandler{LabelService: controller.New()}
}

func (h *labelHandler) NewLabelHandler() http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.CreateLabel)
  r.Delete("/{LabelID}", h.DeleteLabel)
  r.Get("/{Keyword}", h.ListLabel)
  r.Put("/{LabelID}/_attach", h.AttachLabel)

  r.Get("/{LabelID}/_getEntities", h.GetLabels)
  // TODO: review api path for get entities
  r.Get("/fetch/{EntityID}", h.GetEntities)
	return r
}

func (h *labelHandler) CreateLabel(w http.ResponseWriter, r *http.Request) {
  h.LabelService.Create("name", "namespace")
  w.Write([]byte("create label"))
}

func (h *labelHandler) DeleteLabel(w http.ResponseWriter, r *http.Request) {
  h.LabelService.Delete("name", "namespace")
  w.Write([]byte("delete label"))
}

func (h *labelHandler) ListLabel(w http.ResponseWriter, r *http.Request) {
  h.LabelService.List("keyword", "namespace")
  w.Write([]byte("list labels by keyword"))
}

func (h *labelHandler) AttachLabel(w http.ResponseWriter, r *http.Request) {
  h.LabelService.Attach("name", "entityId", "namespace")
  w.Write([]byte("attach label to an entity"))
}

func (h *labelHandler) GetLabels(w http.ResponseWriter, r *http.Request) {
  h.LabelService.GetLabels("entityId", "namespace")
  w.Write([]byte("get labels attached to an entity"))
}

func (h *labelHandler) GetEntities(w http.ResponseWriter, r *http.Request) {
  h.LabelService.GetEntities("labelName", "namespace")
  w.Write([]byte("get entities attached to a label"))
}