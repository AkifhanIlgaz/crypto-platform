# Build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /workspace

COPY shared/ shared/
COPY services/crypto-service/ services/crypto-service/

# Workspace rootunda go modüllerini indir, build et
WORKDIR /workspace/services/crypto-service
RUN go mod tidy
RUN go build -o crypto-service cmd/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /workspace/services/crypto-service/crypto-service .
COPY --from=builder /workspace/shared/config ./shared/config
EXPOSE 7777
CMD ["./crypto-service"]
