# Backend API

Go-based REST API server for the Programming Learning App, handling authentication, content delivery, quiz management, code execution, and user progress.

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: JWT with bcrypt password hashing
- **Migrations**: golang-migrate
- **Testing**: Go's built-in testing with httptest
- **Deployment**: Docker

## Setup

### Prerequisites

- Go 1.21+
- PostgreSQL
- golang-migrate (install via `brew install golang-migrate` or download binary)

### Installation

1. Navigate to backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Setup environment:
   ```bash
   cp .env.example .env
   # Edit .env with your database URL, JWT secret, etc.
   ```

4. Run database migrations:
   ```bash
   migrate -path migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" up
   ```

5. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication
- `POST /signup` - User registration
- `POST /login` - User login (returns JWT)

### Lessons
- `GET /lessons` - List all lessons
- `GET /lessons/:slug` - Get specific lesson with Markdown content

### Quizzes
- `GET /lessons/:slug/quizzes` - Get quizzes for a lesson
- `POST /quiz/submit` - Submit quiz answers and get score

### Progress
- `GET /users/:id/progress` - Get user progress
- `POST /progress` - Update user progress

### Code Execution
- `POST /submissions` - Submit code for execution
- `GET /submissions/:id` - Get execution result

### Leaderboard
- `GET /leaderboard` - Get top users by XP

All protected routes require `Authorization: Bearer <jwt>` header.

## Testing

Run tests:
```bash
go test ./...
```

## Docker

Build and run with Docker:
```bash
docker build -t learnapp-backend .
docker run -p 8080:8080 learnapp-backend
```

## Deployment

1. Build Docker image
2. Deploy to cloud platform (Render, Fly.io)
3. Configure environment variables
4. Run migrations on production DB

## Development

- Server code in `cmd/server/main.go`
- Handlers in `internal/handlers/`
- Models in `internal/models/`
- Database layer in `internal/store/`
- Migrations in `migrations/`

See root [README.md](../README.md) for project overview.