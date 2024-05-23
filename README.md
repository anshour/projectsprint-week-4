# BeliMang Backend Service

This is a backend service that provides a RESTful API.

## Getting Started

### Prerequisites

- Go 1.22
- Postgres
- Golang Migrate CLI

### Installation

1. Migrate the database schema

```sh
migrate -database "postgres://postgres:password@localhost:5432/belimangdb?sslmode=disable" -path ./db/migrations -verbose up
```

2. Run the application

```sh
go run main.go
```
