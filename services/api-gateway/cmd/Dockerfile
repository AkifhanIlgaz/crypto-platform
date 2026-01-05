# Build stage
FROM golang:1.25.3-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first (caching için)
COPY go.mod go.sum ./
COPY vendor/ ./vendor/

# Copy only necessary files
COPY shared/ ./shared/
COPY services/api-gateway/ ./services/api-gateway/

# Build with optimization flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -ldflags="-w -s" -a -installsuffix cgo \
    -o api-gateway ./services/api-gateway/cmd/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/api-gateway .
COPY --from=builder /app/shared/config/config.yaml ./shared/config/config.yaml
EXPOSE 7777
CMD ["./api-gateway"]
