# go-service
Student Management System - Developer Skill Test

This Go-based microservice generates PDF reports for students by consuming data from an existing backend service.

## 📦 Features

- Fetches student data from a remote API
- Generates downloadable PDF reports
- Clean layered architecture: handler, service, client, PDF generator
- Environment-based configuration using `.env` (via `godotenv`)
- Structured using Go best practices (`cmd/`, `internal/`, `pkg/`)

## 🏗️ Project Structure
```
cmd/server/           # Main application entrypoint
internal/
├── handler/        # HTTP route handlers (Gin)
├── service/        # Business logic (report generation)
├── client/         # External API integration (e.g. student service)
├── model/          # Data models
├── config/         # Configuration loader
├── pdf/            # PDF generation logic
└── auth/           # (Optional) Authentication helpers
```

## 🚀 How It Works

1. API call: `GET /students/:id/report`
2. The handler fetches student data from another service via `client`.
3. Generates a PDF using data and streams it as a download response.

## 🛠️ Setup Instructions

### Prerequisites

- Go 1.20+
- A `.env` file with the following:
```
STUDENT_SERVICE_URL=http://localhost:5007/api/v1/students
```
### Run nodejs backend
- clone this repository https://github.com/sevengit-wq/skill-test/tree/main/backend
- update controller to add this line of code
```
const handleGetStudentDetail = asyncHandler(async (req, res) => {
    // new added code here..
    const studentId = req.params.id;
    const student = await getStudentDetail(studentId);
    res.status(200).json(student);

});

```
```
npm install
npm start
```

### Run the server

```bash
go run cmd/server/main.go
```

### Example Request

```
curl -o report.pdf http://localhost:8080/api/v1/students/1/report
```