package client

import (
	"encoding/json"
	"fmt"
	"go-service/core/model"
	"net/http"
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
	url := fmt.Sprintf("%s/api/v1/students/%s", c.baseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned status: %d", resp.StatusCode)
	}

	var student model.Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, err
	}
	return &student, nil
}
