# Builder stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api ./cmd/api

FROM alpine:latest

COPY --from=builder /app/api /api

EXPOSE 80

ENTRYPOINT ["/api"]
