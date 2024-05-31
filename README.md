# Desafio 3 - Clean Architecture

## Descrição
O objetivo deste desafio é criar um novo caso de uso para fazer a listagem de pedidos via: API Rest, GRPC e GraphQL.

## Requisitos Gerais
- Executar as migrations necessárias para criar as tabelas no banco de dados;
- Implementar um novo caso de uso para listar os pedidos;
- Implementar a listagem de pedidos via API Rest;
- Implementar a listagem de pedidos via GRPC;
- Implementar a listagem de pedidos via GraphQL;

## Execução
- Na pasta raíz do projeto, executar o comando:
```bash
docker-compose up
```

- Acessar a pasta `cmd/ordersystem` e executar o comando:
```bash
go run main.go wire_gen.go
```

# Validação

### API Rest
- Acessar a pasta `api` e executar o comando;
- Utilizar o arquivo `create_order.http` para criar um pedido;
- Utilizar o arquivo `list_orders.http` para listar os pedidos;

### GRPC
- Utilizar o comando `evans -r repl` para acessar o client GRPC;
- Utilizar o comando `call CreateOrder` para criar um pedido;
- Utilizar o comando `call ListOrders` para listar os pedidos;

### GraphQL
- Acessar a url `http://localhost:8080/playground`;
- Utilizar a mutation `createOrder` para criar um pedido;
- Utilizar a query `listOrders` para listar os pedidos;

## Observações
- As migrations são executadas automaticamente pelo flyway após a inicialização do container do banco de dados;
