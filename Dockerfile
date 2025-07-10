# Estágio de build
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Otimização de cache: baixa as dependências apenas se go.mod/go.sum mudarem.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila um binário estático para uma imagem final mínima.
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/main ./cmd/main.go

# Estágio final
FROM alpine:latest
WORKDIR /app

# Copia o executável do estágio de build.
COPY --from=builder /app/main .

EXPOSE 8080
ENTRYPOINT ["./main"]