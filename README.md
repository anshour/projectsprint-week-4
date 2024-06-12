# BeliMang Backend Service

This is a backend service that provides api for ride hailing food order online.
for calculation distance using haversine, NearestTSP below 3km

## Getting Started

### Prerequisites

- Go 1.22
- Postgres
- Golang Migrate CLI
- Docker

### Installation

1. Migrate the database schema

```sh
migrate -database "postgres://postgres:password@localhost:5432/belimangdb?sslmode=disable" -path ./db/migrations -verbose up
```

2. Run the application

```sh
go run main.go
```

Specification App
https://www.notion.so/dionfananie/BeliMang-3680606d0cfa4b1c875fc13dd27a2eab
