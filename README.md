# payment-gateway

### Para rodar o projeto execute seguinte comando na pasta raiz:
  ```bash
    $ docker-compose up -d
  ```

## Desafios
### 1) Golang
Nesse desafio você criará uma aplicação desenvolvida na linguagem **Go** que quando executada fará a inserção dos seguintes dados em uma tabela de transações em um banco de dados **SQLite**:
> ID da transação, AccountID, Amount e CreatedAt.

Compile sua aplicação e a inclua em uma imagem **Docker** juntamente com o banco de dados **SQLite**, logo, quando executarmos o container os dados deverão ser inseridos no banco e o container deverá continuar rodando.

Suba a imagem docker gerada no **DockerHub**.

> Dica: Para manter o container rodando você poderá utilizar `CMD ['tail','-f','/dev/null']` como instrução ao final do seu Dockerfile.

----
### 2) Nest.js
Nesse desafio você deverá criar uma aplicação **Nest.js** que possua 2 endpoints REST: um para a criação de uma transação e outro para a listagem de transações. O padrão da url deverá ser: `/transactions`.
O mecanismo de persistência será o banco de dados **SQLite3** e a aplicação deverá ser executada na porta `3000`.

Crie uma imagem **Docker**, faça o push para o **DockerHub**.
Os campos da tabela de transactions serão:
> id, account_id, amount, created_at, updated_at.

----

### 3) Kubernetes
Nesse desafio você deverá criar 2 arquivos de manifesto **Kubernetes** que colocará no ar o servidor web **nginx**. Os manifestos que deverão ser criados serão:
> deployment e o service.

Caso nunca tenha trabalhado localmente com o **Kubernetes**, recomendamos utilizar o **[Kind](https://kind.sigs.k8s.io/)**. Ele é uma ferramenta que cria seu cluster de forma local utilizando o próprio  **Docker** de seu computador.

Dicas:
> A imagem a ser utilizado do nginx é: nginx:latest.