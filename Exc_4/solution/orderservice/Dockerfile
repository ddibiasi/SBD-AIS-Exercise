FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ordersystem ./...

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/ordersystem .
EXPOSE 3000
CMD ["./ordersystem"]
