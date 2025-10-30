# Paddle Traffic API

Backend API server for Paddle Traffic application.

## Requirements

- Go 1.25.1 or higher
- PostgreSQL database
- Docker (optional, for running PostgreSQL locally)

## Setup

### 1. Start PostgreSQL Database

#### Option A: Using Docker

```bash
docker run --name paddletraffic-postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_USER=paddletraffic \
  -e POSTGRES_DB=paddletraffic \
  -p 5432:5432 \
  -d postgres:latest
```

#### Option B: Use existing PostgreSQL instance

Ensure you have a PostgreSQL database running and accessible.

### 2. Set Environment Variables

Set the `DATABASE_URL` environment variable with your PostgreSQL connection string:

```bash
export DATABASE_URL="postgresql://paddletraffic:password@localhost:5432/paddletraffic"
```

For Windows (PowerShell):
```powershell
$env:DATABASE_URL="postgresql://paddletraffic:password@localhost:5432/paddletraffic"
```

For Windows (Command Prompt):
```cmd
set DATABASE_URL=postgresql://paddletraffic:password@localhost:5432/paddletraffic?sslmode=disable
```

### 3. Install sqlc (if not already installed)

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Database Migrations

Migrations are located in `internal/database/migrations/` and use [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

**Migrations run automatically** when the server starts. The server will apply any pending migrations before accepting requests.

### Creating a New Migration

1. Create two new files in `internal/database/migrations/`:
   - `00X_description.up.sql` - for applying the migration
   - `00X_description.down.sql` - for rollback (recommended)

2. Write your SQL schema changes in the `.up.sql` file
3. Write the reverse operations in the `.down.sql` file

Migration files are numbered sequentially (001, 002, 003, etc.).

### Manual Migration Management (Optional)

If you need to manually control migrations, you can install the migrate CLI tool:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Then run migrations manually:

```bash
# Apply all pending migrations
migrate -path internal/database/migrations -database $DATABASE_URL up

# Rollback the last migration
migrate -path internal/database/migrations -database $DATABASE_URL down 1

# Check migration version
migrate -path internal/database/migrations -database $DATABASE_URL version
```

## Generating sqlc Code

After modifying SQL queries in `internal/database/queries/` or schema in `internal/database/migrations/`, regenerate the Go code:

```bash
sqlc generate
```

This will update the generated code in `internal/database/`.

## Running the Server

### Development

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`.

### Build and Run

```bash
go build -o server cmd/server/main.go
./server
```

On Windows:
```cmd
go build -o server.exe cmd/server/main.go
server.exe
```

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── controller/      # HTTP handlers
│   ├── service/         # Business logic
│   ├── repository/      # Data access layer
│   └── database/        # Database queries and migrations
│       ├── migrations/  # SQL migration files
│       └── queries/     # SQL query files for sqlc
├── go.mod
├── go.sum
└── sqlc.yaml           # sqlc configuration
```

## Useful Commands

```bash
# Download dependencies (optional - Go will auto-download when running/building)
# Useful for CI/CD or pre-populating the module cache
go mod download

# Run tests
go test ./...

# Format code
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Build for production
go build -o server cmd/server/main.go

# Generate sqlc code
sqlc generate
```
