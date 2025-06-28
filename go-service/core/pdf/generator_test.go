package pdf

import (
	"go-service/core/model"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	generator := NewGenerator()

	if generator == nil {
		t.Fatal("Expected generator to be created, got nil")
	}
}

func TestGenerator_Generate_Success(t *testing.T) {
	student := &model.Student{
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

	generator := NewGenerator()

	pdfBytes, err := generator.Generate(student)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pdfBytes == nil {
		t.Fatal("Expected PDF bytes to be returned, got nil")
	}

	if len(pdfBytes) == 0 {
		t.Fatal("Expected non-empty PDF bytes")
	}

	if len(pdfBytes) < 4 || string(pdfBytes[:4]) != "%PDF" {
		t.Errorf("Expected PDF to start with %%PDF, got %s", string(pdfBytes[:min(10, len(pdfBytes))]))
	}
}

func TestGenerator_Generate_EmptyStudent(t *testing.T) {
	student := &model.Student{}

	generator := NewGenerator()

	pdfBytes, err := generator.Generate(student)
	if err != nil {
		t.Fatalf("Expected no error for empty student, got %v", err)
	}

	if pdfBytes == nil {
		t.Fatal("Expected PDF bytes to be returned for empty student, got nil")
	}

	if len(pdfBytes) == 0 {
		t.Fatal("Expected non-empty PDF bytes for empty student")
	}

	if len(pdfBytes) < 4 || string(pdfBytes[:4]) != "%PDF" {
		t.Errorf("Expected PDF to start with %%PDF, got %s", string(pdfBytes[:min(10, len(pdfBytes))]))
	}
}

func TestGenerator_Generate_NilStudent(t *testing.T) {
	generator := NewGenerator()

	pdfBytes, err := generator.Generate(nil)
	if err == nil {
		t.Fatal("Expected error for nil student, got nil")
	}

	if pdfBytes != nil {
		t.Fatal("Expected PDF bytes to be nil for nil student")
	}
}

func TestFormatDOB_ValidDate(t *testing.T) {
	validDOB := "2000-01-15T00:00:00Z"
	formatted := formatDOB(validDOB)

	expected := "15 Jan 2000"
	if formatted != expected {
		t.Errorf("Expected formatted date '%s', got '%s'", expected, formatted)
	}
}

func TestFormatDOB_InvalidDate(t *testing.T) {
	invalidDOB := "invalid-date"
	formatted := formatDOB(invalidDOB)

	if formatted != invalidDOB {
		t.Errorf("Expected original string '%s' for invalid date, got '%s'", invalidDOB, formatted)
	}
}

func TestFormatDOB_DifferentFormats(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"2000-01-01T00:00:00Z", "01 Jan 2000"},
		{"1995-12-25T00:00:00Z", "25 Dec 1995"},
		{"2010-06-15T00:00:00Z", "15 Jun 2010"},
		{"1985-03-08T00:00:00Z", "08 Mar 1985"},
	}

	for _, tc := range testCases {
		formatted := formatDOB(tc.input)
		if formatted != tc.expected {
			t.Errorf("For input '%s', expected '%s', got '%s'", tc.input, tc.expected, formatted)
		}
	}
}

func TestGenerator_Generate_MultipleStudents(t *testing.T) {
	generator := NewGenerator()

	students := []*model.Student{
		{
			ID:   1,
			Name: "Student 1",
			Dob:  "2000-01-01T00:00:00Z",
		},
		{
			ID:   2,
			Name: "Student 2",
			Dob:  "2001-02-02T00:00:00Z",
		},
		{
			ID:   3,
			Name: "Student 3",
			Dob:  "2002-03-03T00:00:00Z",
		},
	}

	for i, student := range students {
		pdfBytes, err := generator.Generate(student)
		if err != nil {
			t.Fatalf("Expected no error for student %d, got %v", i+1, err)
		}

		if pdfBytes == nil {
			t.Fatalf("Expected PDF bytes for student %d, got nil", i+1)
		}

		if len(pdfBytes) == 0 {
			t.Fatalf("Expected non-empty PDF bytes for student %d", i+1)
		}

		if len(pdfBytes) < 4 || string(pdfBytes[:4]) != "%PDF" {
			t.Errorf("Expected PDF to start with %%PDF for student %d, got %s", i+1, string(pdfBytes[:min(10, len(pdfBytes))]))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
