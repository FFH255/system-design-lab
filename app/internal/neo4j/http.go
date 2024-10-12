package neo4j

import "github.com/gorilla/mux"

func ServeHTTP(defaultRouter *mux.Router, handler *Handler) {
	r := defaultRouter.PathPrefix("/neo4j").Subrouter()

	r.HandleFunc("/student", handler.CreateStudent).Methods("POST")
	r.HandleFunc("/group", handler.CreateGroup).Methods("POST")
	r.HandleFunc("/course", handler.CreateCourse).Methods("POST")
	r.HandleFunc("/addStudentToGroup", handler.AddStudentToGroup).Methods("POST")
	r.HandleFunc("/enrollStudentInCourse", handler.EnrollStudentInGroup).Methods("POST")
}
