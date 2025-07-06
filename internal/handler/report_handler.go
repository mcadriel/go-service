package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcadriel/go-service/internal/service"
)

func RegisterRoutes(r *gin.RouterGroup, svc service.Service) {
	r.GET("/students/:id/report", func(c *gin.Context) {
		id := c.Param("id")

		// Call the service function to generate the student's PDF report
		studentReport, err := svc.GenerateStudentReport(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
			return
		}
		// Set the HTTP header to indicate that this is a downloadable file with the given filename
		c.Header("Content-Disposition", "attachment; filename=student_report.pdf")
		// Send the PDF data with a 200 OK status and content type set to application/pdf
		c.Data(http.StatusOK, "application/pdf", studentReport)
	})
}
