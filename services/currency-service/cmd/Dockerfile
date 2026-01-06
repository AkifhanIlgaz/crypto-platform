# Build stage
FROM golang:1.25.3-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod files and vendor
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY shared/ ./shared/
COPY services/currency-service/ ./services/currency-service/

# Build using vendor (no network needed)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -p 1  -ldflags="-w -s" -a -installsuffix cgo \
    -o currency-service ./services/currency-service/cmd/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/currency-service .
COPY --from=builder /app/shared/config/config.yaml ./shared/config/config.yaml
EXPOSE 50052
CMD ["./currency-service"]
