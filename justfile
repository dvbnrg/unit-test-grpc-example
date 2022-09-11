# List all recipes (you just called this)
default: dev
# run a local dev build
dev: build run
# Build Macro
build: protoc compile
# Build Client and server binaries
compile:
    echo "build called"
    go build -o bin/client client/main.go
    go build -o bin/server server/main.go
# Run all unit tests
test:
    echo "test called"
# Run Local versions of binaries
run:
    sleep 2
    goreman start
    echo "run called"
# Build protoc artifacts
protoc:
    protoc --go-grpc_out=proto --go_out=proto proto/hello.proto
# Clean all build artifacts
clean:
    rm proto/*.pb.go
    rm bin/*