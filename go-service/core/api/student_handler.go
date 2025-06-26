package api

import (
	"go-service/core/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type StudentHandler struct {
	reportService service.StudentReportService
}

func NewStudentHandler(rs service.StudentReportService) *StudentHandler {
	return &StudentHandler{reportService: rs}
}

func (h *StudentHandler) GetStudentReport(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	id := mux.Vars(r)["id"]

	logger := logrus.WithFields(logrus.Fields{
		"method":     r.Method,
		"path":       r.URL.Path,
		"student_id": id,
		"user_agent": r.UserAgent(),
		"remote_ip":  r.RemoteAddr,
	})

	logger.Info("Generating student report")

	pdfBytes, err := h.reportService.GenerateReport(id)
	if err != nil {
		logger.WithError(err).Error("Failed to generate PDF report")
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)

	logger.WithField("duration_ms", time.Since(start).Milliseconds()).
		WithField("pdf_size_bytes", len(pdfBytes)).
		Info("Successfully generated and sent PDF report")
}
