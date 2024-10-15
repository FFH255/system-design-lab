package elastic_search

import (
	"github.com/gorilla/mux"
	"net/http"
	"sdl/app/pkg/api"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{
		repo: repository,
	}
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")

	courses, err := h.repo.Search(r.Context(), search)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, courses)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[Course](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.repo.Create(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, dto)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[Course](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.repo.Update(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, dto)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["id"]

	if uid == "" {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	err := h.repo.Delete(r.Context(), uid)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, nil)
}
