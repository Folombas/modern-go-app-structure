# Билд приложения
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o delivery-app ./cmd/app

# Финальный образ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/delivery-app .
COPY migrations /app/migrations
COPY scripts /app/scripts
RUN chmod +x /app/scripts/entrypoint.sh
CMD ["/app/scripts/entrypoint.sh"]