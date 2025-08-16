# gRPC - GO

## GRPC Setup

1. Install `grpc-protobuf-compiler` in your `ubuntu` machine

   ```bash
   apt install -y protobuf-compiler
   protoc --version # Ensure compiler version is 3+
   ```

2. Install these go-packages locally to generate grpc go code

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. Use this command to generate the the go code for grpc services

   ```bash
   protoc \
   --go_out=internal/generated/stock --go_opt=paths=source_relative \
   --go-grpc_out=internal/generated/stock --go-grpc_opt=paths=source_relative \
   proto/stock.proto
   ```

4. Command for creating certs and keys

   ```bash
   openssl req -new -x509 -days 365 -nodes -out certs/server.crt -keyout certs/server.key -subj "/CN=localhost"
   ```

5. Use this command to create your custom `Authorization-Token` for you grpc services

   ```bash
   openssl rand -hex 32
   ```
