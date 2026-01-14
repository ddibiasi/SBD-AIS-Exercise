<<<<<<< HEAD
#!/usr/bin/env bash
set -euo pipefail

echo "[build] running go mod tidy"
go mod tidy

echo "[build] building ordersystem"
go build -trimpath -o /app/ordersystem .
echo "[build] done -> /app/ordersystem"
=======
#!/bin/sh
# Exit if any command fails
set -e
cd /app
go mod download
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem
>>>>>>> ed41b50cc2fd92dbd0df12eafd134d95e2bbbd93
