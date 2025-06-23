# gRPC Report Service with Cron

This project is a gRPC-based report generation system in Go, with a cron job that simulates scheduled requests.

## Features

- gRPC service to:
  - Generate reports for a given user ID
  - Health check endpoint
- Cron job that triggers report generation every 10 seconds for a fixed list of users
- In-memory report store

## Usage

### 1. Generate Go code from proto(only if changes have been done to 'report.proto' file)

```sh
protoc --go_out=. --go-grpc_out=. --proto_path=proto proto/report.proto
```
### 2. Download the dependencies
```sh
go mod tidy
```
### 3. Run the server

```sh
go run main.go
```


