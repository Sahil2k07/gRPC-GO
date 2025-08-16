# GRPC + Go Inventory Management

This project demonstrates a full-stack inventory management system built using **Go** and **gRPC**. It includes a Go backend server and gRPC services for inventory operations.

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [gRPC Setup](#grpc-setup)
- [Running the Server](#running-the-server)
- [Building the Project](#building-the-project)

---

## Features

- Inventory management with stock updates and consumption
- gRPC APIs for communication
- Multi-platform build support (Linux, Windows, macOS)
- Secure gRPC with TLS certificates

---

## Prerequisites

- Go >= 1.20
- Protobuf Compiler (`protoc`)
- OpenSSL (for certificates)

---

## gRPC Setup

1. **Install protobuf compiler on Ubuntu**

   ```bash
   sudo apt install -y protobuf-compiler
   protoc --version # Ensure compiler version is 3+
   ```

2. **Install `grpc-protobuf-compiler` in your `ubuntu` machine**

   ```bash
   apt install -y protobuf-compiler
   protoc --version # Ensure compiler version is 3+
   ```

3. **Install these go-packages locally to generate grpc go code**

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

4. **Use this command to generate the the go code for grpc services**

   ```bash
   protoc \
   --go_out=internal/generated/stock --go_opt=paths=source_relative \
   --go-grpc_out=internal/generated/stock --go-grpc_opt=paths=source_relative \
   proto/stock.proto
   ```

5. **Command for creating certs and keys**

   ```bash
   openssl req -new -x509 -days 365 -nodes -out certs/server.crt -keyout certs/server.key -subj "/CN=localhost"
   ```

6. **Use this command to create your custom `Authorization-Token` for you grpc services**

   ```bash
   openssl rand -hex 32
   ```

---

## Running the Server

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Sahil2k07/gRPC-GO.git

   cd gRPC-GO
   ```

2. **Restore all the packages:**

   ```bash
   go mod download
   ```

3. **Optionally have a local copy of all the packages:**

   ```bash
   go mod vendor
   ```

4. \*\*Make a `dev.toml` file and give all the credentials. A `dev.example.toml` has been provided

   ```toml
   [server]
   server_port = ":5000"
   origins = ["http://localhost:3000", "https://example.com"]

   [grpc]
   grpc_port = ":6000"
   grpc_url = "localhost:6000"
   grpc_token = "5763121b0c2141eb73d3c1ddfe65a02f30e56adc2fd6f62d1a143f38dc1f3680"

   [database]
   db_host = "localhost"
   db_port = "5432"
   db_user = "postgres"
   db_password = "sahil"
   db_name = "grpc-go"

   [jwt]
   cookie_name = "gRPC-GO-cookie"
   secret = "K3#v@9$1!pZ^mL2&uQ7*rF4)gT8_W+oB"
   ```

5. **Run the migrations and restore the tables in your database. Refer to `cmd/migration/main.go`**

   ```bash
   go run cmd/migration/main.go
   ```

6. **Start the REST API Server:**

   ```bash
   go run cmd/server/main.go
   ```

7. **Start the gRPC Server:**

   ```bash
   go run cmd/grpc/main.go
   ```

---

## Building the Project

This project uses a Makefile to build server and gRPC binaries.

- **Build Linux server:**

  ```bash
  make server
  ```

- **Build Linux gRPC service:**

  ```bash
  make grpc
  ```

- **Build for Windows/macOS (cross-platform):**

  ```bash
  make server-windows
  make grpc-windows
  ```

  ```bash
  make server-macos
  make grpc-macos
  ```

- **Build everything at once:**

  ```bash
  make all
  ```

- **Clean build artifacts:**

  ```bash
  make clean
  ```

- **Running the Server Build:**

  ```bash
  ./build/server
  ```

- **Running the gRPC Build:**

  ```bash
  ./build/grpc
  ```

## .NET Service

> **Note:** This is a link to the separate .NET repository. The .NET service and the Go/gRPC services communicate with each other.

You can find the .NET service repository here: [https://github.com/Sahil2k07/OMS-gRPC](https://github.com/Sahil2k07/dotnet-inventory-service)
