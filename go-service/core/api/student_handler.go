package api

import (
	"go-service/core/service"
	"net/http"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	reportService service.StudentReportService
}

func NewStudentHandler(rs service.StudentReportService) *StudentHandler {
	return &StudentHandler{reportService: rs}
}

func (h *StudentHandler) GetStudentReport(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	pdfBytes, err := h.reportService.GenerateReport(id)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
