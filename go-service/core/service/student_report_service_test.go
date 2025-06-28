package service

import (
	"errors"
	"testing"

	"go-service/core/model"
)

type mockStudentClient struct {
	student *model.Student
	err     error
}

func (m *mockStudentClient) FetchStudent(id string) (*model.Student, error) {
	return m.student, m.err
}

type mockPDFGenerator struct {
	pdfBytes []byte
	err      error
}

func (m *mockPDFGenerator) Generate(student *model.Student) ([]byte, error) {
	return m.pdfBytes, m.err
}

func TestNewStudentReportService(t *testing.T) {
	mockClient := &mockStudentClient{}
	mockPDFGen := &mockPDFGenerator{}

	service := NewStudentReportService(mockClient, mockPDFGen)

	if service == nil {
		t.Fatal("Expected service to be created, got nil")
	}

	_, ok := service.(*studentReportService)
	if !ok {
		t.Fatal("Expected service to be of type *studentReportService")
	}
}

func TestStudentReportService_GenerateReport_Success(t *testing.T) {
	expectedStudent := &model.Student{
		ID:                 1,
		Name:               "John Doe",
		Email:              "john.doe@example.com",
		SystemAccess:       true,
		Phone:              "1234567890",
		Gender:             "Male",
		Dob:                "2000-01-01T00:00:00Z",
		Class:              "10",
		Section:            "A",
		Roll:               101,
		FatherName:         "John Doe Sr",
		FatherPhone:        "1234567891",
		MotherName:         "Jane Doe",
		MotherPhone:        "1234567892",
		GuardianName:       "John Doe Sr",
		GuardianPhone:      "1234567891",
		RelationOfGuardian: "Father",
		CurrentAddress:     "123 Main St",
		PermanentAddress:   "123 Main St",
		AdmissionDate:      "2020-01-01",
		ReporterName:       "Admin",
	}

	expectedPDFBytes := []byte("mock pdf content")

	mockClient := &mockStudentClient{
		student: expectedStudent,
		err:     nil,
	}

	mockPDFGen := &mockPDFGenerator{
		pdfBytes: expectedPDFBytes,
		err:      nil,
	}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pdfBytes == nil {
		t.Fatal("Expected PDF bytes to be returned, got nil")
	}

	if len(pdfBytes) != len(expectedPDFBytes) {
		t.Errorf("Expected PDF size %d, got %d", len(expectedPDFBytes), len(pdfBytes))
	}

	for i, b := range pdfBytes {
		if b != expectedPDFBytes[i] {
			t.Errorf("Expected PDF byte at index %d to be %d, got %d", i, expectedPDFBytes[i], b)
		}
	}
}

func TestStudentReportService_GenerateReport_ClientError(t *testing.T) {
	expectedError := errors.New("failed to fetch student")

	mockClient := &mockStudentClient{
		student: nil,
		err:     expectedError,
	}

	mockPDFGen := &mockPDFGenerator{}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("1")
	if err == nil {
		t.Fatal("Expected error when client fails, got nil")
	}

	if pdfBytes != nil {
		t.Fatal("Expected PDF bytes to be nil when client fails")
	}

	expectedErrorMsg := "failed to fetch student 1: failed to fetch student"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, err.Error())
	}
}

func TestStudentReportService_GenerateReport_PDFGenerationError(t *testing.T) {
	expectedStudent := &model.Student{
		ID:   1,
		Name: "John Doe",
	}

	expectedError := errors.New("failed to generate PDF")

	mockClient := &mockStudentClient{
		student: expectedStudent,
		err:     nil,
	}

	mockPDFGen := &mockPDFGenerator{
		pdfBytes: nil,
		err:      expectedError,
	}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("1")
	if err == nil {
		t.Fatal("Expected error when PDF generation fails, got nil")
	}

	if pdfBytes != nil {
		t.Fatal("Expected PDF bytes to be nil when PDF generation fails")
	}

	expectedErrorMsg := "failed to generate PDF for student 1: failed to generate PDF"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, err.Error())
	}
}

func TestStudentReportService_GenerateReport_EmptyStudentID(t *testing.T) {
	mockClient := &mockStudentClient{}
	mockPDFGen := &mockPDFGenerator{}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("")
	if err == nil {
		t.Fatal("Expected error for empty student ID, got nil")
	}

	if pdfBytes != nil {
		t.Fatal("Expected PDF bytes to be nil for empty student ID")
	}
}

func TestStudentReportService_GenerateReport_NilStudent(t *testing.T) {
	mockClient := &mockStudentClient{
		student: nil,
		err:     nil,
	}

	mockPDFGen := &mockPDFGenerator{}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("1")
	if err == nil {
		t.Fatal("Expected error when student is nil, got nil")
	}

	if pdfBytes != nil {
		t.Fatal("Expected PDF bytes to be nil when student is nil")
	}
}

func TestStudentReportService_GenerateReport_EmptyPDFBytes(t *testing.T) {
	expectedStudent := &model.Student{
		ID:   1,
		Name: "John Doe",
	}

	mockClient := &mockStudentClient{
		student: expectedStudent,
		err:     nil,
	}

	mockPDFGen := &mockPDFGenerator{
		pdfBytes: []byte{},
		err:      nil,
	}

	service := NewStudentReportService(mockClient, mockPDFGen)

	pdfBytes, err := service.GenerateReport("1")
	if err != nil {
		t.Fatalf("Expected no error for empty PDF bytes, got %v", err)
	}

	if pdfBytes == nil {
		t.Fatal("Expected empty PDF bytes to be returned, got nil")
	}

	if len(pdfBytes) != 0 {
		t.Errorf("Expected empty PDF bytes, got %d bytes", len(pdfBytes))
	}
}
