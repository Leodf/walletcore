# Walletcore

Este projeto exemplar ilustra a criação de dois microsserviços: o primeiro, denominado "CORE", realiza chamadas POST para criar as entidades de cliente, conta e transação. Ele opera com eventos e os envia para o Kafka. O segundo microsserviço, chamado "BALANCES", é capaz de receber os eventos gerados pelo microsserviço "CORE" através do Kafka e persistir os saldos atualizados para cada conta em um banco de dados MySQL. Adicionalmente, o projeto oferece um endpoint que possibilita a visualização do saldo atualizado de uma conta específica.

## Pré-Requisitos

Antes de executar o microsserviço, certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.

## Executando o Microsserviço

1. Clone o repositório deste microsserviço em sua máquina.

2. No diretório raiz do projeto, execute o seguinte comando para iniciar o microsserviço e todos os microsserviços relacionados:

```bash
  docker-compose up -d
```

Este comando inicia o microsserviço de Balances, o microsserviço Wallet Core e os serviços de Kafka e banco de dados.

Certifique-se de que as migrações do banco de dados foram executadas automaticamente durante o início do serviço.

Para acompanhar os logs dos microserviços rodar em terminais separadamente:

```bash
  docker logs core -f
  docker logs balances -f
```

## Documentação API dos Microserviços

### CORE

#### Salva um cliente no banco de dados

```http
POST http://localhost:8080/clients
Content-Type: application/json

{
  "name": "usuario",
  "email": "usuario@email.com"
}
```

#### Salva uma account no banco de dados

```http
POST http://localhost:8080/accounts
Content-Type: application/json

{
  "clientId": "{client_id}"
}
```

#### Salva uma transação no banco de dados

```http
POST http://localhost:8080/transactions
Content-Type: application/json

{
  "accountIdFrom":
  {account_id_from},
  "accountIdTo":
  {account_id_to},
  "amount": 1
}
```

### BALANCES

#### Retorna uma transação no banco de dados

```http
GET http://localhost:3003/api/balances/{account_id}
```

## Stack utilizada

**Back-end:** Go lang, chi router, go fiber, mysql, kafka.
