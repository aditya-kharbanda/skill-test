package service

import (
	"fmt"
	"go-service/core/client"
	"go-service/core/pdf"

	"github.com/sirupsen/logrus"
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
	logger := logrus.WithField("student_id", id)

	student, err := s.client.FetchStudent(id)
	if err != nil {
		logger.WithError(err).Error("Failed to fetch student data")
		return nil, fmt.Errorf("failed to fetch student %s: %w", id, err)
	}

	if student == nil {
		logger.Error("Student data is nil")
		return nil, fmt.Errorf("student data is nil for student %s", id)
	}

	logger.WithField("student_name", student.Name).Debug("Student data fetched successfully")

	pdfBytes, err := s.pdfGen.Generate(student)
	if err != nil {
		logger.WithError(err).Error("Failed to generate PDF")
		return nil, fmt.Errorf("failed to generate PDF for student %s: %w", id, err)
	}

	logger.WithField("pdf_size_bytes", len(pdfBytes)).Debug("PDF generated successfully")
	return pdfBytes, nil
}
