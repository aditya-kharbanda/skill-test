package pdf

import (
	"bytes"
	"fmt"
	"go-service/core/model"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
)

type Generator interface {
	Generate(student *model.Student) ([]byte, error)
}

type generator struct{}

func NewGenerator() Generator {
	return &generator{}
}

func (g *generator) Generate(student *model.Student) ([]byte, error) {
	if student == nil {
		return nil, fmt.Errorf("student cannot be nil")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(60, 10, fmt.Sprintf("ID: %d", student.ID))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Name: %s", student.Name))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Email: %s", student.Email))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("System Access: %t", student.SystemAccess))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Phone: %s", student.Phone))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Gender: %s", student.Gender))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("DOB: %s", formatDOB(student.Dob)))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Class: %s", student.Class))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Section: %s", student.Section))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Roll: %d", student.Roll))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Father Name: %s", student.FatherName))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Father Phone: %s", student.FatherPhone))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Mother Name: %s", student.MotherName))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Mother Phone: %s", student.MotherPhone))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Guardian Name: %s", student.GuardianName))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Guardian Phone: %s", student.GuardianPhone))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Relation Of Guardian: %s", student.RelationOfGuardian))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Current Address: %s", student.CurrentAddress))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Permanent Address: %s", student.PermanentAddress))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Admission Date: %s", student.AdmissionDate))
	pdf.Ln(8)
	pdf.Cell(60, 10, fmt.Sprintf("Reporter Name: %s", student.ReporterName))
	pdf.Ln(8)
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		logrus.WithError(err).WithField("student_id", student.ID).Error("Failed to generate PDF output")
		return nil, err
	}
	return buf.Bytes(), err
}

func formatDOB(dob string) string {
	t, err := time.Parse(time.RFC3339, dob)
	if err != nil {
		logrus.WithError(err).WithField("dob", dob).Warn("Invalid date format for student DOB")
		return dob
	}
	return t.Format("02 Jan 2006")
}
