# FynroGo API

A REST API service built with Go, Gin, MySQL, and OpenAI integration.

## Features

- Question processing with OpenAI integration
- MySQL database storage
- RESTful API endpoints
- UUID generation for questions
- Proper error handling and logging
- CORS support

## Project Structure

```
fyrnogo/
├── cmd/server/main.go          # Application entry point
├── internal/
│   ├── config/config.go        # Configuration management
│   ├── database/database.go    # Database connection and setup
│   ├── handlers/question.go    # HTTP handlers
│   ├── models/models.go        # Data models
│   ├── services/               # Business logic
│   └── middleware/             # HTTP middleware
├── pkg/utils/                  # Utility functions
├── migrations/                 # Database migrations
└── go.mod                      # Go module file
```

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd fyrnogo
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Setup MySQL database**
   ```bash
   # Run the migration script
   mysql -u root -p < migrations/001_init.sql
   ```

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### POST /api/v1/question
Create a new question and get AI response.

**Request:**
```json
{
    "question": "What is the capital of France?",
    "username": "john_doe"  // optional
}
```

**Response:**
```json
{
    "success": true,
    "message": "Question processed successfully",
    "data": {
        "id": 1,
        "questionid": "uuid-here",
        "question": "What is the capital of France?",
        "username": "john_doe",
        "answer": "The capital of France is Paris.",
        "created_at": "2023-01-01T00:00:00Z"
    }
}
```

### GET /api/v1/question/:questionid
Retrieve a question and its response by questionid.

**Response:**
```json
{
    "success": true,
    "message": "Question retrieved successfully",
    "data": {
        "id": 1,
        "questionid": "uuid-here",
        "question": "What is the capital of France?",
        "username": "john_doe",
        "answer": "The capital of France is Paris.",
        "created_at": "2023-01-01T00:00:00Z"
    }
}
```

### GET /api/v1/health
Health check endpoint.

## Environment Variables

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 3306)
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `SERVER_PORT`: Server port (default: 8080)
- `OPENAI_API_KEY`: OpenAI API key

## Running in Development

```bash
# Install air for hot reloading (optional)
go install github.com/cosmtrek/air@latest

# Run with hot reloading
air
```

## Database Schema

### Questions Table
- `id`: INT AUTO_INCREMENT PRIMARY KEY
- `questionid`: VARCHAR(36) UNIQUE NOT NULL (UUID)
- `question`: TEXT NOT NULL
- `username`: VARCHAR(255)
- `created_at`: TIMESTAMP
- `updated_at`: TIMESTAMP

### Responses Table
- `id`: INT AUTO_INCREMENT PRIMARY KEY
- `questionid`: VARCHAR(36) NOT NULL (Foreign Key)
- `response`: TEXT NOT NULL
- `created_at`: TIMESTAMP
- `updated_at`: TIMESTAMP

## License

MIT License