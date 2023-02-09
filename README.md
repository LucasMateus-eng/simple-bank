<h1 align="center"> Desafio Técnico </h1>

## :scroll: Descrição do projeto

API, para simular a operação de movimentação financeira entre correntistas de um banco, escrita em Go e estruturada segundo a arquitetura hexagonal e DDD

## :wrench: Status do projeto

<p align=center> :sparkles: Concluído :sparkles: </p>

## :hammer: Funcionalidades do projeto

1. Endpoints para o domínio Wallet (carteira):

- recurso para retorno de uma carteira em específico;
- recurso para cadastro de uma carteira;
- recurso para edição de uma carteira;
- recurso para exclusão de uma carteira;
- recurso para realizar um depósito em uma carteira;
- recurso para realizar transferência entre carteiras.

## :unlock: Como executar e consumir o projeto

1. Clone o repositório:

- Clonando o repositório:

```bash
git git@github.com:LucasMateus-eng/simple-bank.git
cd simple-bank
```

2. Crie um arquivo .env:

```
APP_PORT=8080
PG_USER=root
PG_PASSWORD=secretbank
PG_DATABASE=simple_bank
PG_PORT=5432
PG_HOST=localhost
```

- NOTA: pode existir algum problema na conexão entre a imagem docker e a aplicação, uma possível solução é definir o host em env como "host.docker.internal" (sem as aspas :D).

3. Executando a aplicação:

- NOTA: certifique-se de ter o Docker instalado em sua máquina. Veja em: [https://docs.docker.com/engine/install/]

- A inicialização de parte do sistema acontece através de comandos no Makefile.

  3.1 Para a criação do container PostgreSQL execute:

  ```bash
  make postgres
  ```

  3.2 Para a criação do banco de dados execute:

  ```bash
  make createdb
  ```

  3.3 Para a execução do arquivo de migration (inicialização do schema) execute (antes veja instrução abaixo):

  ```bash
  make migrateup
  ```

  3.4 Para a execução da aplicação em modo de desenvolvimento execute no terminal:

  ```bash
  go run main.go
  ```

- Para executar o comando em 3.3 você irá precisar de ter o CLI Migrate em sua máquina (confome SO):

  Veja a documentação em: [https://github.com/golang-migrate/migrate/tree/master/cmd/migrate]

- Se você for um usuário Windows poderá ter problemas na execução dos comandos Make. Se for o seu caso, ou se tiver tendo problemas para executar os comandos Make; independetemente do SO, recomendo executar, no seu terminal, os comandos existentes no arquivo Makefile para cada target dos passos acima (3.1, 3.2, 3.3).

- OBS.: Se você possuir o PostgreSQL instalado em sua máquina poderá ocorrer falhas na conexão da aplicação com o banco.
- Uma possícel solução é:

  1. coloque as configurações do banco existente em sua máquina nas variáveis de ambiente da aplicação (com a excessão do HOST);

4. Escolha algum cliente http de sua preferência e aproveite para testar os endpoints disponíveis:

- Uma outra opção é testar via a página swagger da API, localizada em: [http://localhost:8080/api/v1/docs/swagger/index.html]

5. Cuidado

- Se você estiver recebendo algum erro relacionado a conexão pode ser o seu anti-vírus proibindo o acesso a alguns recursos do seu sistema. Minha recomendação: desligue o seu anti-vírus enquanto executa os endpoints.

## :computer: Tecnologias utilizadas

- Go (1.18.0)
- Echo Framework
- PostgreSQL
- GORM
- Migrate CLI
- Docker
