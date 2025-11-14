# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder

# Let Go download a newer toolchain (1.25.x) automatically
ENV GOTOOLCHAIN=auto

# Create working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies (better cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the orderservice binary from the main package in this folder
# (no ./... here, we build exactly one main package)
RUN go build -o ordersystem .

# ========= Runtime stage =========
FROM alpine:3.20 AS run

# Create working directory for the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ordersystem /app/ordersystem

# (Optional) set the port the service listens on
ENV PORT=3000

# Expose the port for documentation purposes
EXPOSE 3000

# Start the service
CMD ["/app/ordersystem"]