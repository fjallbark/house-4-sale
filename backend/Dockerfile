# Stage 1: Build
FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o house4sale ./cmd/api

# Stage 2: Minimal runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/house4sale .
CMD ["./house4sale"]
