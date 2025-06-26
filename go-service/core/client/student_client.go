package client

import (
	"encoding/json"
	"fmt"
	"go-service/core/model"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type StudentClient interface {
	FetchStudent(id string) (*model.Student, error)
}

type studentClient struct {
	baseURL string
}

func NewStudentClient(url string) StudentClient {
	return &studentClient{baseURL: url}
}

func (c *studentClient) FetchStudent(id string) (*model.Student, error) {
	start := time.Now()
	url := fmt.Sprintf("%s/api/v1/students/%s", c.baseURL, id)

	logger := logrus.WithFields(logrus.Fields{
		"student_id": id,
		"url":        url,
	})

	logger.Debug("Fetching student data from backend")

	resp, err := http.Get(url)
	if err != nil {
		logger.WithError(err).WithField("duration_ms", time.Since(start).Milliseconds()).
			Error("HTTP request failed")
		return nil, fmt.Errorf("failed to fetch student from %s: %w", url, err)
	}
	defer resp.Body.Close()

	logger.WithFields(logrus.Fields{
		"status_code": resp.StatusCode,
		"duration_ms": time.Since(start).Milliseconds(),
	}).Debug("HTTP response received")

	if resp.StatusCode != http.StatusOK {
		logger.WithField("status_code", resp.StatusCode).Error("Backend returned non-OK status")
		return nil, fmt.Errorf("backend returned status %d for student %s", resp.StatusCode, id)
	}

	var student model.Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		logger.WithError(err).Error("Failed to decode JSON response")
		return nil, fmt.Errorf("failed to decode student response: %w", err)
	}

	logger.WithField("student_name", student.Name).Debug("Student data decoded successfully")
	return &student, nil
}
