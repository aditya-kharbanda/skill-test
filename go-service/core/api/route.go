package api

import (
	"go-service/core/service"

	"github.com/gorilla/mux"
)

func NewRouter(reportService service.StudentReportService) *mux.Router {
	handler := NewStudentHandler(reportService)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/students/{id}/report", handler.GetStudentReport).Methods("GET")
	return r
}
