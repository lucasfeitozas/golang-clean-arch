.PHONY: build run test clean docker-up docker-down proto graphql

# Build da aplicação
build:
	go build -o bin/ordersystem cmd/ordersystem/main.go

# Executar a aplicação localmente
run:
	go run cmd/ordersystem/main.go

# Executar testes
test:
	go test -v ./...

# Limpar binários
clean:
	rm -rf bin/

# Subir todos os serviços com Docker
docker-up:
	docker-compose up --build -d

# Parar todos os serviços
docker-down:
	docker-compose down

# Restart completo
docker-restart: docker-down docker-up

# Ver logs da aplicação
logs:
	docker-compose logs -f app

# Gerar código gRPC
proto:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

# Gerar código GraphQL
graphql:
	go run github.com/99designs/gqlgen generate

# Executar migrações manualmente
migrate:
	mysql -h localhost -P 3306 -u root -proot orders < sql/migrations/001_create_orders_table.sql

# Instalar dependências
deps:
	go mod download
	go mod tidy

# Help
help:
	@echo "Comandos disponíveis:"
	@echo "  build        - Compilar a aplicação"
	@echo "  run          - Executar a aplicação localmente"
	@echo "  test         - Executar testes"
	@echo "  clean        - Limpar binários"
	@echo "  docker-up    - Subir todos os serviços com Docker"
	@echo "  docker-down  - Parar todos os serviços"
	@echo "  docker-restart - Reiniciar todos os serviços"
	@echo "  logs         - Ver logs da aplicação"
	@echo "  proto        - Gerar código gRPC"
	@echo "  graphql      - Gerar código GraphQL"
	@echo "  migrate      - Executar migrações manualmente"
	@echo "  deps         - Instalar dependências"
