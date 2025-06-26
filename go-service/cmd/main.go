package main

import (
	"go-service/core/api"
	"go-service/core/client"
	"go-service/core/pdf"
	"go-service/core/service"
	"log"
	"net/http"
)

func main() {
	backendClient := client.NewStudentClient("http://localhost:5007")
	pdfGen := pdf.NewGenerator()
	reportService := service.NewStudentReportService(backendClient, pdfGen)
	router := api.NewRouter(reportService)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
