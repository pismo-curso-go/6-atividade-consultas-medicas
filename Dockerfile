# Etapa 1: Build da aplicação
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api

# Etapa 2: Imagem final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY .env .

EXPOSE 8000

CMD ["./server"]
