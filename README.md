# Mapeamento objeto relacional

## Grupo

- Lucas Ricieri
- Ruan Azevedo

## Como executar

Certifique-se de ter o postgres na versão v1.6.0 ou posterior configurado.

1. Clone o repositório e abra a pasta
```bash
git clone https://github.com/lucas-ricieri/lbd-gorm.git
cd lbd-gorm
```
2. Crie o arquivo `.env` na pasta root, com base no seguinte modelo:
```text
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=lbd_gorm
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=America/Sao_Paulo
```

3. Altere as variáveis do banco de dados no `.env`.
4. Execute o comando para instalar todas as dependências.
```bash
go mod tidy
```
5. Execute a aplicação:
```bash
go run ./cmd/api
```

A API será iniciada em `http://localhost:8080`.
