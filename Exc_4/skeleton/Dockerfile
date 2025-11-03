# Stage 1 :: Builder for the golang REST API app
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN chmod +x scripts/build-application.sh
RUN ./scripts/build-application.sh

# Stage 2 :: Final lightweight image based on alpine
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/ordersystem .
RUN ls /app/ordersystem
EXPOSE 8080

CMD ["./ordersystem"]
