package mongo

import (
	"github.com/gorilla/mux"
	"net/http"
	"sdl/app/pkg/api"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["id"]

	if uid == "" {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	group, err := h.service.Get(r.Context(), UID(uid))

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, group)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[Group](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Create(r.Context(), &dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[Group](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Update(r.Context(), &dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["id"]

	if uid == "" {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	err := h.service.Delete(r.Context(), UID(uid))

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, nil)
}
