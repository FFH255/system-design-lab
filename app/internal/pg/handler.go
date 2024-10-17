package pg

import (
	"github.com/gorilla/mux"
	"net/http"
	"sdl/app/pkg/api"
	"strconv"
)

type Handler struct {
	repo *Repository
}

func NewHandler(
	repo *Repository,
) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	attendanceList, err := h.repo.GetAll(r.Context())

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, attendanceList)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[CreateAttendanceDTO](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	attendance, err := h.repo.Create(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, attendance)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := api.ReadJSON[UpdateAttendanceDTO](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	attendance, err := h.repo.Update(r.Context(), dto)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, attendance)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	uidStr := mux.Vars(r)["id"]

	if uidStr == "" {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	uid, err := strconv.Atoi(uidStr)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, "uid is required")
		return
	}

	err = h.repo.Delete(r.Context(), DeleteAttendanceDTO{ID: uid})

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, nil)
}
