  # todo
# -------- builder --------
FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/ordersystem

# -------- runtime --------
FROM alpine:3.20
RUN adduser -D -H app
USER app
WORKDIR /app
COPY --from=builder /app/ordersystem /app/ordersystem
EXPOSE 3000
CMD ["/app/ordersystem"]
