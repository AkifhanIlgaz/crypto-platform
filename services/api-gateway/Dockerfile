# Build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /workspace

COPY shared/ shared/
COPY services/api-gateway/ services/api-gateway/


# Workspace rootunda go modüllerini indir, build et
WORKDIR /workspace/services/api-gateway
RUN go mod download
RUN go build -o api-gateway cmd/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /workspace/services/api-gateway/api-gateway .
COPY --from=builder /workspace/shared/config ./shared/config
EXPOSE 7777
CMD ["./api-gateway"]
