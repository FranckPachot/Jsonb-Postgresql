## Introduction

This repository demonstrates a benchmark comparison between PostgreSQL (using JSONB) and MongoDB for storing and querying semi-structured data in Go. It provides Dockerized instances, Go code for CRUD + benchmarks, indexing guidance, and analysis of real-world use cases.

---

## Prerequisites

Ensure these tools are installed (download links included):

- **Go (1.20+ recommended)**  
  Download: https://go.dev/dl/  
  Verify: `go version`
- **Docker**  
  Download: https://www.docker.com/get-started  
  Verify: `docker version`
- **Git**  
  Download: https://git-scm.com/downloads  
  Verify: `git --version`
- **psql client (optional)**  
  Download: https://www.postgresql.org/download/  
  Or use `docker exec` to run psql inside the container.
- **Mongo Shell (optional)**  
  Download: https://www.mongodb.com/try/download/community  
  Or use Docker’s `mongosh` via `docker exec mongo-json-test mongosh`.

---

## Repository Structure

```
Jsonb-Postgresql/
├── cmd/
│   └── main.go
├── internal/
│   └── db/
│       ├── postgresql.go
│       └── mongo.go
├── data/
│   └── generator.go
├── benchmark/
│   └── runner.go
├── .env
├── go.mod
├── go.sum
└── README.md
```

- `cmd/main.go`: entrypoint; loads `.env`, connects to DBs, runs benchmarks or CRUD tests.
- `internal/db/postgresql.go`: PostgreSQL connection, AutoMigrate, CRUD with `gorm.io/datatypes.JSON` and index creation.
- `internal/db/mongo.go`: MongoDB connection and CRUD.
- `data/generator.go`: dummy JSON data generator.
- `benchmark/runner.go`: runs Insert/Read/Update/Delete benchmarks for both DBs.
- `.env`: environment variables.
- `go.mod`/`go.sum`: Go module files.

---

## Environment Configuration

Create `.env` in project root with contents like:

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
```

Adjust values if using different ports, credentials, or hosted services.

---

## Docker Setup

### PostgreSQL

1. Pull image:
   ```bash
   docker pull postgres
   ```
2. Run container:
   ```bash
   docker run --name pg-json-test \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=pass \
     -e POSTGRES_DB=testdb \
     -p 5432:5432 \
     -d postgres
   ```
3. Verify:
   ```bash
   docker ps
   ```
4. Optional psql:
   ```bash
   docker exec -it pg-json-test psql -U postgres -d testdb
   ```
5. Stop/remove when done:
   ```bash
   docker stop pg-json-test
   docker rm pg-json-test
   ```

### MongoDB

1. Run container:
   ```bash
   docker run --name mongo-json-test \
     -p 27017:27017 \
     -d mongo
   ```
2. Verify:
   ```bash
   docker ps
   ```
3. Optional shell:
   ```bash
   docker exec -it mongo-json-test mongosh
   ```
4. Stop/remove when done:
   ```bash
   docker stop mongo-json-test
   docker rm mongo-json-test
   ```

---

## Go Module Setup

1. In project root, initialize module (if not already):
   ```bash
   go mod init your/module/path
   ```
   - Replace `your/module/path` with actual module path (e.g., github.com/username/Jsonb-Postgresql).
2. Install dependencies:

   ```bash
   go get github.com/joho/godotenv
   go get gorm.io/gorm
   go get gorm.io/driver/postgres
   go get gorm.io/datatypes
   go get go.mongodb.org/mongo-driver/mongo
   go get go.mongodb.org/mongo-driver/mongo/options

   or

   go mod tidy
   ```

3. Ensure `go.mod` and `go.sum` are updated.

---

## Running the Application

From project root:

```bash
go run cmd/main.go
```

---

## Sample Benchmark Output for 10000 records

```
✅ .env loaded
✅ GIN index on JSONB 'data' created (or already exists)
✅ Connected and AutoMigrated PostgreSQL
✅ Connected to PostgreSQL in 111.2811ms
✅ Connected to MongoDB
✅ Connected to MongoDB in 15.1681ms
🔗 Total DB connection setup time: 127.5404ms
✅ Generated 10000 dummy records for benchmarking
📝 PostgreSQL Insert Time: 45.3893459s -> Too Slow
📝 MongoDB Insert Time: 65.7657ms
🔍 PostgreSQL Read Time: 1.8192ms
🔍 MongoDB Read Time: 1.2427ms
✏️ PostgreSQL Update Time: 305.8399ms
✏️ MongoDB Update Time: 73.6471ms
🗑️ PostgreSQL Delete Time: 6.5021ms
🗑️ MongoDB Delete Time: 7.8429ms
🏁 Benchmark tests completed!
```

---

## Use Case Assessment

Below are key use cases for JSONB (Postgres) vs. NoSQL, with evaluation:

### 1. User Preferences/Settings

- **Use Case**: Store per-user flexible settings instead of many columns.
- **Assessment**: Valid. JSONB persists arbitrary structures; GIN or expression indexes allow filtering. Suitable when writes are moderate.

### 2. Event Logs / Audit Trails

- **Use Case**: Capture events with varying payloads in one table.
- **Assessment**: Valid for moderate volume. JSONB handles different schemas; index on `event_type`, `occurred_at`, or specific JSON keys makes queries efficient. For extremely high ingest rates, consider batching, unlogged tables, or a write-optimized store, but JSONB works if SQL analytics or joins are needed.

### 3. Product Catalogs with Dynamic Attributes

- **Use Case**: Store product-specific fields (phones vs. books vs. shoes) in JSONB.
- **Assessment**: Correct. Use JSONB for variable attributes; index “hot” fields (e.g., brand, size). Keep core columns (name, price) relational for type safety and fast aggregations.

### 4. API Response Caching / Form Submissions

- **Use Case**: Save third-party API payloads or arbitrary form data without rigid schema.
- **Assessment**: Reasonable. JSONB stores raw blobs. Querying inside may be rare, but having data in Postgres simplifies backup and consistency. For heavy writes, use batching, unlogged tables, or partitioning.

### 5. Feature Flags & Configuration

- **Use Case**: Store per-feature configs with nested fields.
- **Assessment**: Valid. JSONB holds nested structures; index keys for fast lookups. For very high read volume, ensure caching or indexes.

### 6. Hybrid Relational + Flexible Fields

- **Use Case**: Mix real columns (FKs, totals) with JSONB extras (e.g., orders with line_items in JSON).
- **Assessment**: Strongly valid. Core columns ensure integrity and fast aggregates; JSONB holds variable details. Common in many systems.

### 7. Analytics on Semi-Structured Data

- **Use Case**: Run SQL aggregates on JSON fields (SUM, AVG after casting).
- **Assessment**: True. Extract fields via `->>` or generated/stored columns, index or aggregate. Performance depends on indexing and data volume but is workable.

### 8. Full-Text Search Inside JSON

- **Use Case**: Create GIN index on `to_tsvector(data->>'content')`.
- **Assessment**: Correct. Enables SQL full-text search on JSON fields.

### 9. Migration from MongoDB

- **Use Case**: Export from MongoDB into a Postgres JSONB table.
- **Assessment**: Feasible. Many consolidate to Postgres. Must handle write performance and index creation after import.

### 10. Cheat-Sheet Operators

- **Use Case**: Use containment (`@>`), existence (`?`), path (`#>`), concatenation (`||`), deletion (`-`, `#-`), etc.
- **Assessment**: All correct. These operators power JSON manipulation in SQL.

---

## Indexing Guidance

After AutoMigrate, create indexes manually in `ConnectPostgres()`:

```go
db.Exec(`CREATE INDEX IF NOT EXISTS idx_json_data_gin ON json_data USING GIN (data)`)
db.Exec(`CREATE INDEX IF NOT EXISTS idx_json_data_id ON json_data ((data->>'id'))`)
```

- GIN index speeds containment queries (`@>`).
- Expression index on `data->>'id'` speeds lookups by that JSON key.
- For other frequent keys, create similar expression indexes or generated columns.

---

## Troubleshooting

- **.env not found**: Run from project root (`go run cmd/main.go`) so `.env` is found.
- **Table missing**: Verify “Connected and AutoMigrated PostgreSQL” appears; in psql, `\dt` should list `json_data`.
- **Index not used**: Ensure query uses the same form as the index (e.g., `data->>'key' = ?`). Test in psql with `EXPLAIN ANALYZE`.
- **Port conflicts**: Change ports in `docker run -p` if 5432/27017 are in use.
- **High resource usage**: For large benchmarks, ensure sufficient RAM/CPU; consider batching or COPY for inserts.
- **MongoDB authentication**: For production-like, create users in Mongo and update `MONGO_URI`.

---

## Cleanup

```bash
docker stop pg-json-test mongo-json-test
docker rm pg-json-test mongo-json-test
docker rmi postgres mongo
```

Remove volumes if used:

```bash
docker volume rm pgdata
```

---

## Extending the Project

- Increase record count for stress tests (e.g., 10k, 100k).
- Add more complex JSON queries: nested arrays, full-text search on JSON content.
- Implement PostgreSQL partitioning for large event tables.
- Experiment with COPY or bulk insert optimizations for JSONB.
- Automate benchmarks in CI and record results over time.
- Test against hosted databases (Aiven, AWS RDS, MongoDB Atlas) to gauge real-world latency.
- Add TLS/SSL, authentication, monitoring (Prometheus/Grafana) for production readiness.

---

## References & Downloads

- Go: https://go.dev/dl/
- Docker: https://www.docker.com/get-started
- GORM: https://gorm.io/
- Mongo Go Driver: https://pkg.go.dev/go.mongodb.org/mongo-driver
- godotenv: https://github.com/joho/godotenv

---

## License

Choose a license (e.g., MIT) and include here. This project demonstrates JSONB vs. MongoDB benchmarking for educational or internal use.
