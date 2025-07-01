# 🚀 Go gRPC Demo Project

This is a simple Go project demonstrating all four types of gRPC communication:
- Unary RPC
- Server-side streaming
- Client-side streaming
- Bidirectional streaming

---

## ⚙️ Prerequisites

- Go 1.20+
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

Install them (if not already):

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

## Project Structure 

```
.
├── client
│   ├── bi_stream.go
│   ├── client_stream.go
│   ├── main.go
│   ├── server_stream.go
│   └── unary.go
├── go.mod
├── go.sum
├── proto
│   ├── greet_grpc.pb.go
│   ├── greet.pb.go
│   └── greet.proto
├── Readme.md
└── server
    ├── bi_stream.go
    ├── client_stream.go
    ├── main.go
    ├── server_stream.go
    └── unary.go

```

## Generate gRPC Code

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## Running the server

```bash
go run server/*.go
```

## Running the client

```bash
go run client/*.go
```

## gRPC Models Explained

| Mode                 | Client File        | Server File        |
| -------------------- | ------------------ | ------------------ |
| Unary RPC            | `unary.go`         | `unary.go`         |
| Server Streaming     | `server_stream.go` | `server_stream.go` |
| Client Streaming     | `client_stream.go` | `client_stream.go` |
| Bidirectional Stream | `bi_stream.go`     | `bi_stream.go`     |
