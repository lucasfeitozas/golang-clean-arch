#!/bin/bash
# Script para executar migrações do banco de dados

# Aguarda o MySQL estar pronto
echo "Aguardando MySQL estar pronto..."
until mysql -h localhost -P 3306 -u root -proot -e "SELECT 1" >/dev/null 2>&1; do
  echo "MySQL ainda não está pronto. Aguardando..."
  sleep 2
done

echo "MySQL está pronto. Executando migrações..."

# Executa as migrações
mysql -h localhost -P 3306 -u root -proot orders < /sql/migrations/001_create_orders_table.sql

echo "Migrações executadas com sucesso!"
