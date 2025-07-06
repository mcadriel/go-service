# go-service
Student Management System - Developer Skill Test

This Go-based microservice generates PDF reports for students by consuming data from an existing Node.js backend service.

---

## 📦 Features

- Fetches student data from a remote API
- Generates downloadable PDF reports
- Clean layered architecture (`handler`, `service`, `client`, `pdf`)
- Environment-based configuration using `.env` (via `godotenv`)
- Follows Go best practices using a modular structure (`cmd/`, `internal/`, `pkg/`)

---

## 🏗️ Project Structure

```
cmd/server/           # Main application entrypoint
internal/
├── handler/          # HTTP route handlers (Gin)
├── service/          # Business logic (report generation)
├── client/           # External API integration (e.g. student service)
├── model/            # Data models
├── config/           # Configuration loader (uses godotenv)
├── pdf/              # PDF generation logic
└── auth/             # (Optional) Authentication helpers
```

---

## 🚀 How It Works

1. API endpoint: `GET /api/v1/students/:id/report`
2. The handler retrieves student data from the Node.js backend via the client module.
3. A PDF report is generated using the fetched data.
4. The PDF is sent as a downloadable file in the response.

---

## 🛠️ Setup Instructions

### Prerequisites

- Go 1.20+
- Node.js (for backend API)
- `.env` file in the root directory with:

```env
STUDENT_SERVICE_URL=http://localhost:5007/api/v1/students
```

---

### 🟢 Start the Node.js Backend

1. Clone the backend repository:
   ```
   git clone https://github.com/sevengit-wq/skill-test.git
   cd skill-test/backend
   ```

2. Update the controller function in the Node.js backend
   ``` /backend/src/modules/students/students-controller.js```
   to include:
   ```js
   const handleGetStudentDetail = asyncHandler(async (req, res) => {
       const studentId = req.params.id;
       const student = await getStudentDetail(studentId);
       res.status(200).json(student);
   });
   ```

3. Start the server:
   ```bash
   npm install
   npm start
   ```

---

### 🚀 Run the Go Microservice

In the root directory of the Go service:

```bash
go run cmd/server/main.go
```

---

### 📄 Example Request

Use `curl` to generate and download a student PDF report:

```bash
curl -o report.pdf http://localhost:8080/api/v1/students/1/report
```

This will download `report.pdf` for student with ID `1`.

---