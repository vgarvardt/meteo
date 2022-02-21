//go:build tools

package main

// Place any runtime dependencies as imports in this file.
// Go modules will be forced to download and install them.
import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
