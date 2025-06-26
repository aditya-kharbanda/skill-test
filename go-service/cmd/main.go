package main

import (
	"go-service/core/api"
	"go-service/core/client"
	"go-service/core/pdf"
	"go-service/core/service"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// Configure Logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	logger := logrus.WithField("service", "go-service")

	backendClient := client.NewStudentClient("http://localhost:5007")
	pdfGen := pdf.NewGenerator()
	reportService := service.NewStudentReportService(backendClient, pdfGen)
	router := api.NewRouter(reportService)

	logger.Info("Server starting on :8080")
	logger.Fatal(http.ListenAndServe(":8080", router))
}
