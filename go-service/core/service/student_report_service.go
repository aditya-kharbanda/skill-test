package service

import (
	"go-service/core/client"
	"go-service/core/pdf"
)

type StudentReportService interface {
	GenerateReport(id string) ([]byte, error)
}

type studentReportService struct {
	client client.StudentClient
	pdfGen pdf.Generator
}

func NewStudentReportService(c client.StudentClient, p pdf.Generator) StudentReportService {
	return &studentReportService{client: c, pdfGen: p}
}

func (s *studentReportService) GenerateReport(id string) ([]byte, error) {
	student, err := s.client.FetchStudent(id)
	if err != nil {
		return nil, err
	}
	return s.pdfGen.Generate(student)
}
