## Introduction

This repository demonstrates a benchmark comparison between PostgreSQL (using JSONB) and MongoDB for storing and querying semi-structured data in Go. It provides:

- Dockerized instances of PostgreSQL and MongoDB for local testing.
- Go code to generate dummy JSON records, perform CRUD operations, and measure timings.
- Configuration via environment variables.
- Instructions to scale tests (e.g., 100, 1,000, 10,000 records).
- Guidance on indexing (GIN, expression indexes) to optimize PostgreSQL JSONB queries.

Use this README to set up your environment, run containers, configure the Go application, execute benchmarks, and interpret results.

---

## Prerequisites

Ensure the following tools are installed and available in your PATH. If not installed, follow the download links.

- **Go (1.20+ recommended)**
  - Download: https://go.dev/dl/
  - After installation, verify with:
    ```bash
    go version
    ```
- **Docker**
  - Download (Docker Desktop for Windows/macOS, or Docker Engine for Linux): https://www.docker.com/get-started
  - Verify:
    ```bash
    docker version
    ```
- **Git**
  - Download: https://git-scm.com/downloads
  - Verify:
    ```bash
    git --version
    ```
- **psql client (optional)**
  - To inspect PostgreSQL from host (outside Docker), install psql:
    - PostgreSQL downloads: https://www.postgresql.org/download/
    - Or use Docker’s psql via `docker exec`, so local install is optional.
- **Mongo Shell (optional)**
  - For inspecting MongoDB: https://www.mongodb.com/try/download/community
  - Or use Docker’s `mongosh` via `docker exec`, so local install is optional.
- **Environment**
  - Windows, macOS, or Linux with internet access to download above tools.
  - A terminal/PowerShell (on Windows) or bash/zsh (on macOS/Linux).

---
- `cmd/main.go`: application entrypoint. Loads .env, connects to DBs, invokes benchmark or CRUD tests.
- `internal/db/postgresql.go`: PostgreSQL connection, AutoMigrate, CRUD functions using `gorm.io/datatypes.JSON`.
- `internal/db/mongo.go`: MongoDB connection and CRUD functions using `mongo-driver`.
- `data/generator.go`: generates dummy JSON records (`DummyData` struct).
- `benchmark/runner.go`: runs Insert/Read/Update/Delete benchmarks for both DBs.
- `.env`: environment variables (DB URIs, credentials).
- `go.mod` / `go.sum`: Go module configuration.

---

## Environment Configuration

Create a `.env` file in the project root. Example contents:

```env
# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=pass
POSTGRES_DB=testdb

# MongoDB
MONGO_URI=mongodb://localhost:27017
MONGO_DB=testdb
