# Project settings
BUILD_DIR := build

# Server settings
SERVER_MAIN := ./cmd/server/main.go
SERVER_BIN := $(BUILD_DIR)/server

# gRPC settings
GRPC_MAIN := ./cmd/grpc/main.go
GRPC_BIN := $(BUILD_DIR)/grpc

# Default target: build server for Linux
.PHONY: server
server:
	@echo "Building server for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(SERVER_BIN) $(SERVER_MAIN)
	@echo "✅ Linux server binary created at $(SERVER_BIN)"

# Cross-platform builds for server
.PHONY: server-windows
server-windows:
	@echo "Building server for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(SERVER_BIN).exe $(SERVER_MAIN)
	@echo "✅ Windows server binary created at $(SERVER_BIN).exe"

.PHONY: server-macos
server-macos:
	@echo "Building server for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(SERVER_BIN)-mac $(SERVER_MAIN)
	@echo "✅ macOS server binary created at $(SERVER_BIN)-mac"

# gRPC build for Linux
.PHONY: grpc
grpc:
	@echo "Building gRPC server for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(GRPC_BIN) $(GRPC_MAIN)
	@echo "✅ Linux gRPC binary created at $(GRPC_BIN)"

# Cross-platform builds for gRPC
.PHONY: grpc-windows
grpc-windows:
	@echo "Building gRPC server for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(GRPC_BIN).exe $(GRPC_MAIN)
	@echo "✅ Windows gRPC binary created at $(GRPC_BIN).exe"

.PHONY: grpc-macos
grpc-macos:
	@echo "Building gRPC server for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(GRPC_BIN)-mac $(GRPC_MAIN)
	@echo "✅ macOS gRPC binary created at $(GRPC_BIN)-mac"

# Build all binaries (server + gRPC) for Linux
.PHONY: all
all: server grpc
	@echo "✅ All binaries built for Linux"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	@echo "🧹 Clean complete"
