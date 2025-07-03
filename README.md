# Sistema de Orders - Clean Architecture

Este projeto implementa o desafio de Clean Architecture em Go, com suporte a três tipos de interfaces: REST API, gRPC e GraphQL.

## 📋 Funcionalidades

- ✅ Criar orders (REST, gRPC, GraphQL)
- ✅ Listar orders (REST, gRPC, GraphQL)
- ✅ Arquitetura limpa com separação de responsabilidades
- ✅ Eventos assíncronos com RabbitMQ
- ✅ Banco de dados MySQL
- ✅ Containerização com Docker

## 🛠 Tecnologias Utilizadas

- **Go 1.23**
- **MySQL 5.7**
- **RabbitMQ**
- **gRPC**
- **GraphQL (gqlgen)**
- **Chi Router** (REST API)
- **Docker & Docker Compose**

## 🚀 Como Executar

### Pré-requisitos

- Docker
- Docker Compose

### Executando o projeto

1. Clone o repositório:
```bash
git clone https://github.com/lucasfeitozas/golang-clean-arch.git
cd golang-clean-arch
```

2. Execute o projeto com Docker Compose:
```bash
docker-compose up --build
```

O comando acima irá:
- Subir o banco MySQL na porta 3306
- Subir o RabbitMQ na porta 5672 (management: 15672)
- Executar as migrações do banco automaticamente
- Subir a aplicação com todos os serviços

## 🌐 Portas dos Serviços

| Serviço | Porta | Descrição |
|---------|-------|-----------|
| REST API | 8000 | Endpoints HTTP |
| gRPC | 50051 | Serviços gRPC |
| GraphQL | 8080 | Playground e queries |
| MySQL | 3306 | Banco de dados |
| RabbitMQ | 5672 | Message broker |
| RabbitMQ Management | 15672 | Interface web (guest/guest) |

## 📡 Como Testar os Endpoints

### 1. REST API

#### Criar Order
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "order-123",
    "price": 100.50,
    "tax": 10.05
  }'
```

#### Listar Orders
```bash
curl -X GET http://localhost:8000/order
```

### 2. gRPC

Você pode usar ferramentas como `grpcurl` ou `evans` para testar:

```bash
# Listar serviços disponíveis
grpcurl -plaintext localhost:50051 list

# Criar order
grpcurl -plaintext -d '{
  "id": "grpc-order-1",
  "price": 150.75,
  "tax": 15.08
}' localhost:50051 pb.OrderService/CreateOrder

# Listar orders
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders
```

### 3. GraphQL

Acesse o GraphQL Playground em: http://localhost:8080

#### Criar Order (Mutation)
```graphql
mutation {
  createOrder(input: {
    id: "gql-order-1"
    Price: 200.00
    Tax: 20.00
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Listar Orders (Query)
```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## 📁 Estrutura do Projeto

```
.
├── cmd/ordersystem/          # Ponto de entrada da aplicação
├── configs/                  # Configurações
├── internal/
│   ├── entity/              # Entidades de domínio
│   ├── usecase/             # Casos de uso
│   ├── infra/
│   │   ├── database/        # Repositórios
│   │   ├── web/             # Handlers REST
│   │   ├── grpc/            # Serviços gRPC
│   │   └── graph/           # Resolvers GraphQL
│   └── event/               # Eventos e handlers
├── pkg/events/              # Package de eventos
├── sql/migrations/          # Migrações do banco
├── api/                     # Arquivos .http para testes
├── docker-compose.yaml      # Orquestração dos serviços
├── Dockerfile              # Imagem da aplicação
└── README.md               # Este arquivo
```

## 🗄 Banco de Dados

O projeto utiliza MySQL com a seguinte estrutura:

```sql
CREATE TABLE orders (
    id VARCHAR(255) PRIMARY KEY,
    price DECIMAL(10,2) NOT NULL,
    tax DECIMAL(10,2) NOT NULL,
    final_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

As migrações são executadas automaticamente quando o container sobe.

## 📨 RabbitMQ

O projeto utiliza RabbitMQ para eventos assíncronos:

- **Fila**: `orders`
- **Exchange**: `amq.direct`
- **Routing Key**: `order.created`

A fila é criada automaticamente quando a aplicação inicia. Você pode monitorar as mensagens através do RabbitMQ Management em http://localhost:15672 (guest/guest).

## 🔧 Desenvolvimento

### Executando localmente (sem Docker)

1. Configure as variáveis de ambiente no arquivo `.env`
2. Certifique-se de que MySQL e RabbitMQ estão rodando
3. Execute as migrações manualmente:
```bash
mysql -u root -p orders < sql/migrations/001_create_orders_table.sql
```
4. Execute a aplicação:
```bash
go run cmd/ordersystem/main.go
```

### Gerando código gRPC

```bash
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

### Gerando código GraphQL

```bash
go run github.com/99designs/gqlgen generate
```

## 📝 Arquivos de Teste

O projeto inclui arquivos `.http` na pasta `api/` para facilitar os testes:

- `api/create_order.http` - Criar order via REST
- `api/list_orders.http` - Listar orders via REST

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.