package redis

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

	student, err := h.service.Get(r.Context(), UID(uid))

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, student)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[CreateStudentDTO](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	student, err := h.service.Create(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, student)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[UpdateStudentDTO](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	student, err := h.service.Update(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, student)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["id"]

	if uid == "" {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	h.service.Delete(r.Context(), UID(uid))
	api.WriteJSON(w, http.StatusOK, nil)
}
