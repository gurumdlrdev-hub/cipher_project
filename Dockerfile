# Stage 1: Builder
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cipher_project .
FROM alpine:latest
WORKDIR /app


COPY --from=builder /app/cipher_project .
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/database ./database

EXPOSE 7777
CMD ["./cipher_project"]