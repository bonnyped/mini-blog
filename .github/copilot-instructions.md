# AI Coding Guidelines for mini-blog

## Architecture Overview
- **Framework**: Go web server using Chi router for HTTP handling
- **Storage**: PostgreSQL with `database/sql` and `lib/pq` driver
- **Config**: YAML-based with `cleanenv` for loading, supports env overrides
- **Logging**: Structured logging with `slog` in JSON format
- **Migrations**: Database schema managed with Goose

## Project Structure
- `cmd/main.go`: Application entry point with router setup
- `internal/`: Core business logic
  - `config/`: Configuration loading from `config/config.yaml`
  - `handlers/`: HTTP handlers as functions returning `http.HandlerFunc`
  - `models/domain/`: Domain models (User, Note structs)
  - `logger/sl/`: Simple error wrapping utility
- `storage/postgres/`: Database operations implementing storage interfaces
- `migrations/`: Goose migration files for schema changes

## Key Patterns
- **Handlers**: Accept logger and storage interface, return `http.HandlerFunc`. Use `middleware.GetReqID(r.Context())` for request tracing.
- **Storage Interfaces**: Define methods like `CreateUser(username string, creationTime time.Time) error` in handler packages.
- **Error Handling**: Wrap errors with `sl.Err(op, err)` where `op` is a constant like `"handlers.create_user.New"`.
- **Logging**: Enrich logger with operation and request_id for traceability.
- **Config Path**: Hardcoded relative path `../config/config.yaml` from `internal/config/config.go`.

## Development Workflows
- **Run Server**: `go run cmd/main.go` (listens on hardcoded `127.0.0.1:8082`, config partially unused)
- **Database Migrations**: Use Goose CLI with connection string from config (e.g., `goose postgres "host=localhost port=5432 user=postgres password=123 dbname=blog" up`)
- **Debugging**: VS Code launch config targets `${workspaceFolder}/cmd`

## Conventions
- Use `const op = "package.function"` for operation identifiers in error/logging contexts
- Form data parsing: `r.FormValue("key")` for simple inputs
- Database: Prepared statements for inserts, no transactions shown yet
- Models: JSON tags for serialization, basic validation tags (e.g., `validate:"required"`)

## Incomplete Areas
- `create_note` handler is stubbed; implement similar to `create_user`
- Storage missing `CreateNote` method; add to `postgres.go`
- Config loading ignores `HTTPServer` settings; update main.go to use `config.HttpServer`
- No tests or build scripts defined</content>
<parameter name="filePath">/home/bonnyped/projects/mini-blog/.github/copilot-instructions.md