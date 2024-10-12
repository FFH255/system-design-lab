package neo4j

import (
	"net/http"
	"sdl/app/pkg/api"
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

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	student, err := api.ReadJSON[Student](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, "cannot read student")
		return
	}

	err = h.repo.CreateStudent(r.Context(), student)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, student)
}

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	group, err := api.ReadJSON[Group](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, "cannot read group")
		return
	}

	err = h.repo.CreateGroup(r.Context(), group)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, group)
}

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	course, err := api.ReadJSON[Course](r)

	if err != nil {
		api.WriteError(w, http.StatusBadRequest, "cannot read course")
		return
	}

	err = h.repo.CreateCourse(r.Context(), course)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, course)
}

func (h *Handler) AddStudentToGroup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	studentID := query.Get("student_id")
	if studentID == "" {
		api.WriteError(w, http.StatusBadRequest, "student_id is required")
		return
	}

	groupID := query.Get("group_id")
	if groupID == "" {
		api.WriteError(w, http.StatusBadRequest, "group_id is required")
		return
	}

	err := h.repo.AddStudentToGroup(r.Context(), studentID, groupID)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) EnrollStudentInGroup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	studentID := query.Get("student_id")
	if studentID == "" {
		api.WriteError(w, http.StatusBadRequest, "student_id is required")
		return
	}

	courseID := query.Get("course_id")
	if courseID == "" {
		api.WriteError(w, http.StatusBadRequest, "course_id is required")
		return
	}

	err := h.repo.EnrollStudentInCourse(r.Context(), studentID, courseID)

	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, nil)
}
