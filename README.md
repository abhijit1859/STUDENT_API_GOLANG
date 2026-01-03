# STUDENT_API_GOLANG

This is a student API that I wrote while learning Golang.

Simple student management HTTP API (SQLite-backed).

Prerequisites
- Go 1.25.1 (see [go.mod](go.mod))

Quick start
- Provide configuration path via env or flag:
  - ENV: `CONFIG_PATH=config/local.yaml`
  - or flag: `-config config/local.yaml`
  (loader: [`config.MustLoad`](internal/config/config.go), sample: [config/local.yaml](config/local.yaml))

Build & run
```sh
# run directly
export CONFIG_PATH=config/local.yaml
go run ./cmd/student-api

# or with flag
go run ./cmd/student-api -config 

# build
go build -o bin/student-api ./cmd/student-api
./bin/student-api -config
```