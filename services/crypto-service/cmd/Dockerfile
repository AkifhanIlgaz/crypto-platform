# Build stage
FROM golang:1.25.3-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod files and vendor
COPY go.mod go.sum ./
COPY vendor/ ./vendor/

# Copy source files
COPY shared/ ./shared/
COPY services/crypto-service/ ./services/crypto-service/

# Build using vendor (no network needed)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -ldflags="-w -s" -a -installsuffix cgo \
    -o crypto-service ./services/crypto-service/cmd/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/crypto-service .
COPY --from=builder /app/shared/config/config.yaml ./shared/config/config.yaml
EXPOSE 50051
CMD ["./crypto-service"]
