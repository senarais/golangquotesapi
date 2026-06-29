# Golang Quotes API

Simple REST API that returns random inspirational quotes, built with Go and Gin.

## Run

```bash
go run cmd/api/main.go
```

Server starts at `http://localhost:5000`.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Health check |
| GET | `/quote` | Get a random quote |

### GET `/quote`

Returns a random quote in JSON:

```json
{
  "Quote": "Stay hungry, stay foolish.",
  "Author": "Steve Jobs"
}
```
