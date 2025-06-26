# Student PDF Report Microservice (Go)

This microservice generates downloadable PDF reports for students by consuming data from an existing Node.js backend API. It is written in Go and follows clean, modular architecture.

---

## Features

- Consumes `/api/v1/students/:id` from Node.js backend
- Generates and returns PDF reports via `/api/v1/students/:id/report`
- Clean architecture using interface-driven design
- Code formatting and commit hooks for quality enforcement
- Structured logging with Logrus for observability

---

## Logging Strategy

The service uses Logrus for structured JSON logging with the following approach:

### Log Levels by Layer:
- **HTTP Handler**: Info level for request/response, Error for failures
- **Service Layer**: Debug for successful operations, Error for business failures  
- **Client Layer**: Debug for API calls, Error for network/HTTP failures
- **PDF Generator**: Error only for critical PDF generation failures

### Key Principles:
- **No Duplication**: Each layer logs only what it's responsible for
- **Contextual Fields**: Student ID, request details, timing, and sizes are logged
- **Error Wrapping**: Errors are wrapped with context and logged once
- **Performance Metrics**: Request duration and response sizes are tracked

### Example Log Output:
```json
{
  "level": "info",
  "msg": "Generating student report",
  "method": "GET",
  "path": "/api/v1/students/1/report",
  "student_id": "1",
  "time": "2024-01-15T10:30:00Z"
}
```

## Error Wrapping Strategy

This service uses Go's error wrapping (`fmt.Errorf("context: %w", err)`) to provide context as errors propagate through the application. This makes debugging easier and preserves the original error for inspection.

### Where errors are wrapped:
- **Service Layer:** Wraps errors from dependencies (e.g., client, PDF generator) with business context.
- **Client Layer:** Wraps network/decoding errors with operation context (e.g., which URL, which student ID).

### Where errors are not wrapped:
- **HTTP Handler:** Converts errors to HTTP responses, logs them, but does not wrap further.
- **Utility Functions:** Usually return raw errors, unless additional context is critical.

---

## Project Structure

```
go-service/
├── cmd/                    # Entry point (main.go)
├── core/
│   ├── api/               # HTTP handler (StudentHandler)
│   ├── client/            # Node.js API client
│   ├── pdf/               # PDF generator
│   ├── model/             # Student struct
│   └── service/           # Business logic (ReportService)
```

---

## Setup

1. **Clone this repo**

2. **Install Go dependencies**
   ```bash
   cd go-service
   go mod tidy
   ```

3. **Set up PostgreSQL and seed the DB**
   ```bash
   psql -U postgres -f seed_db/seed-db.sql
   ```

4. **Run the Node.js backend**
   ```bash
   cd backend
   npm install
   npm run dev
   ```

5. **Run the Go service**
   ```bash
   cd go-service
   make run
   ```

---

## API Usage

**Endpoint:** `GET /api/v1/students/:id/report`

**Example:**
```bash
curl http://localhost:8080/api/v1/students/1/report --output report.pdf
```

You should receive a downloadable PDF report for the student.

---

## Makefile

This project includes a simple `makefile` to streamline common development tasks. The available targets are:

| Target   | Description                                      |
|----------|--------------------------------------------------|
| `build`  | Compiles the Go service and outputs the binary to `bin/go-service`. |
| `run`    | Runs the Go service directly using `go run ./cmd`. |
| `fmt`    | Formats the codebase using `goimports` and `gofmt`. |
| `lint`   | Runs `golangci-lint` to check for linting issues.  |

### Usage

From the `go-service` directory, you can run:

```bash
make build   # Build the binary
make run     # Run the service
make fmt     # Format the code
make lint    # Lint the code
```

This helps ensure code quality and makes development more efficient.