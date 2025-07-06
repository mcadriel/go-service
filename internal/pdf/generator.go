package pdf

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/mcadriel/go-service/internal/model"
)

func GenerateStudentReport(student *model.Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, fmt.Sprintf("Student Report for %s", student.Name))
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("ID: %d", student.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.Name))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
