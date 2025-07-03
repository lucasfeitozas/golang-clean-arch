# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar go mod e sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem

# Final stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/

# Copiar o binário da aplicação
COPY --from=builder /app/main .

# Copiar arquivos de configuração
COPY --from=builder /app/.env .

# Copiar scripts SQL
COPY --from=builder /app/sql ./sql

# Tornar o script executável
RUN chmod +x sql/init.sh

# Expor portas
EXPOSE 8000 50051 8080

# Comando para executar a aplicação
CMD ["./main"]
