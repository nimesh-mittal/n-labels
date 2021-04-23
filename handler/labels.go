package handler

import (
  "n_labels/controller"
  "n_labels/entity"
  "net/http"
  "github.com/go-chi/chi/v5"
  "encoding/json"
  "strconv"
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
  r.Get("/{LabelID}/_getEntities", h.GetEntities)
  // TODO: review api path for get entities
  r.Get("/fetch/{EntityID}", h.GetLabels)
	return r
}

func (h *labelHandler) CreateLabel(w http.ResponseWriter, r *http.Request) {
  
  decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var createLabelRequest entity.CreateLabelRequest
	err := decoder.Decode(&createLabelRequest)

  if err != nil{
    e := entity.NewError("invalid create label request")
    res, _ := json.Marshal(e)
    w.Write(res)  
    return
  }

  status, err := h.LabelService.Create(createLabelRequest.Name, createLabelRequest.Namespace)

  if err != nil{
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

  // TODO: also take namespace as input
  status, err := h.LabelService.Delete(id, "namespace")
  
  if err != nil{
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
  keyword := chi.URLParam(r, "Keyword")

  labels, err := h.LabelService.List(keyword, "namespace")
  
  if err != nil{
    e := entity.NewError("error processing list label request")
    res, _ := json.Marshal(e)
    w.Write(res)  
    return
  }

  res, _ := json.Marshal(labels)
  w.Write(res)
}

func (h *labelHandler) AttachLabel(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "LabelId")

  decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var attachLabelRequest entity.AttachLabelRequest
	err := decoder.Decode(&attachLabelRequest)

  if err != nil{
    e := entity.NewError("invalid attach label request")
    res, _ := json.Marshal(e)
    w.Write(res)  
    return
  }

  status, err := h.LabelService.Attach(id, attachLabelRequest.EntityId, attachLabelRequest.Namespace)

  if err != nil{
    e := entity.NewError("error processing attach label request")
    res, _ := json.Marshal(e)
    w.Write(res)  
    return
  }

  e := entity.SuccessResponse{Status: strconv.FormatBool(status)}
  res, _ := json.Marshal(e)
  w.Write(res)
}

func (h *labelHandler) GetLabels(w http.ResponseWriter, r *http.Request) {
  // TODO: complete the implementation
  w.Write([]byte("get labels attached to an entity"))
}

func (h *labelHandler) GetEntities(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "LabelId")

  entities, err := h.LabelService.GetEntities(id, "namespace")

  if err != nil{
    e := entity.NewError("error processing getEntities for label request")
    res, _ := json.Marshal(e)
    w.Write(res)  
    return
  }

  res, _ := json.Marshal(entities)
  w.Write(res)
}