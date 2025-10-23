# -------- Stage 1: Builder --------
FROM golang:1.25 AS builder
WORKDIR /app

# Copiem només els scripts necessaris
COPY skeleton/scripts ./scripts

# Fem executable i cridem el build
RUN chmod +x scripts/build-application.sh && bash scripts/build-application.sh

# -------- Stage 2: Runner --------
FROM alpine:latest
WORKDIR /app

# Copiem l'aplicació compilada des del builder
COPY --from=builder /app/ordersystem ./ordersystem

EXPOSE 8080

ENV POSTGRES_DB=order \
    POSTGRES_USER=docker \
    POSTGRES_PASSWORD=docker \
    POSTGRES_TCP_PORT=5432 \
    DB_HOST=127.0.0.1

CMD ["./ordersystem"]
