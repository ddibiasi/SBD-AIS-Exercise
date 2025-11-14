#!/bin/sh
set -eu

echo "==> Running go mod tidy"
go mod tidy

echo "==> Running tests"
go test ./...

echo "==> Building orderservice binary"
go build -o ordersystem .
