#!/usr/bin/env bash
set -euo pipefail

echo "[build] running go mod tidy"
go mod tidy

echo "[build] building ordersystem"
go build -trimpath -o /app/ordersystem .
echo "[build] done -> /app/ordersystem"
