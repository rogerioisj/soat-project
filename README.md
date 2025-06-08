# soat-project

### Autor: Rogerio Inacio
### RM: 364104

# Sobre o Projeto
Este projeto tem como objetivo desenvolver um sistema de vendas de produtos em uma lanchonete. Inicialmente, focando apenas na jornada do cliente, desde a identificação, até a retirada do pedido.

# Tecnologias Utilizadas
- Go
- PostgreSQL
- Docker

# Como Executar o Projeto
Para facilitar a execução do sistema, foi criado um arquivo `docker-compose.yml` que configura o ambiente com o banco de dados PostgreSQL e o serviço Go.
Os arquivos docker estão localizados na pasta `.docker`.

## Para execução com Docker:
1. Certifique-se de ter o Docker e o Docker Compose instalados ou o Docker Desktop.
2. Certifique-se de que o Docker está rodando.
3. Verifique se existe algum serviço ocupando as portas 8080 e 5432. Ou altere se preferir.
4. Abra o terminal na pasta `.docker` dentro do projeto.
5. Execute o comando:
   ```bash
   docker-compose up -d
   ```
    ou para Docker Desktop:
    ```bash
    docker compose up -d
    ```
6. Execute o comando para execução das migrations:
   ```bash
   go run cmd/db/main.go
   ```
   6.1. Em caso de falha, será necessário a execução de cada uma das migrations manualmente, na ordem correta, para criar as tabelas no banco de dados.
   ```bash
   go run cmd/db/main.go version 1
   go run cmd/db/main.go version 2
   go run cmd/db/main.go version 3
    ```
7. Verifique o swagger em `http://localhost:8080/` para testar as rotas.

## Para execução sem Docker:
1. Certifique-se de ter o Go instalado.
2. Certifique-se de ter o serviço PostgreSQL rodando.
3. Crie um banco de dados de sua escolha no PostgreSQL.
4. Configure as variáveis de ambiente no arquivo `.env` com as informações do banco de dados.
5. Abra o terminal na pasta raiz do projeto.
6. Execute o comando para execução das migrations:
   ```bash
   go run cmd/db/main.go
   ```
   6.1. Em caso de falha, será necessário a execução de cada uma das migrations manualmente, na ordem correta, para criar as tabelas no banco de dados.
   ```bash
   go run cmd/db/main.go version 1
   go run cmd/db/main.go version 2
   go run cmd/db/main.go version 3
   ```
7. Execute o comando para iniciar o servidor:
   ```bash
   go run cmd/api/main.go
   ```
8. Verifique o swagger em `http://localhost:8080/` para testar as rotas.

# Observações
1. O projeto utiliza o Swagger para documentação das rotas, facilitando o teste e a visualização das APIs.
2. A chamada das rotas pelo swagger pode vir a falhar. Nesse caso existe arquivos `.http` na pasta `requests` que podem ser utilizados para testar as rotas.