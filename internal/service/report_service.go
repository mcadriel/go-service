package service

import (
	"github.com/mcadriel/go-service/internal/client"
	"github.com/mcadriel/go-service/internal/pdf"
)

func GenerateStudentReport(studentID string) ([]byte, error) {
	student, err := client.FetchStudentByID(studentID)
	if err != nil {
		return nil, err
	}

	return pdf.GenerateStudentReport(student)
}
