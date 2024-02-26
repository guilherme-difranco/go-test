# API de Gerenciamento de Tarefas com Frontend em React

Este projeto inclui uma API para o gerenciamento de tarefas, juntamente com um frontend desenvolvido em React. A solução completa permite criar, listar, atualizar e deletar tarefas, destacando-se pela performance e escalabilidade, suportando alto volume de requisições. O projeto demonstra práticas de clean architecture, caching, otimizações específicas para o MongoDB, e uma interface de usuário amigável.

## Iniciando

Instruções para configurar o projeto localmente para desenvolvimento e testes.

### Pré-requisitos

Para executar este projeto, você precisará ter instalado:

- Go
- MongoDB
- Node.js e npm

```bash
go version
node --version
npm --version
```

### Instalação

1. Clone o projeto para sua máquina local:

    ```bash
    git clone https://github.com/guilherme-difranco/go-test
    ```

2. Backend:

    a. Acesse a pasta do backend e instale as dependências:

    ```bash
    cd go-test/backend
    go mod tidy
    ```

    b. Execute o servidor localmente:

    ```bash
    go run cmd/main.go
    ```

    A API estará disponível em `http://localhost:8080`.

3. Frontend:

    a. Acesse a pasta do frontend e instale as dependências:

    ```bash
    cd ../frontend
    npm install
    ```

    b. Execute o projeto React localmente:

    ```bash
    npm start
    ```

    O frontend estará disponível em `http://localhost:3000`.

## Testes

Para executar os testes automatizados no backend:

```bash
cd backend
go test ./...
```

Para rodar o teste de integração localmente, defina a variável de ambiente `$env:PROJECT_ROOT` apontando para o diretório raiz do projeto backend:

```bash
$env:PROJECT_ROOT="<caminho_para_o_diretório_backend>"
go test ./... -tags=integration
```

## Construído Com

- [Gin](https://github.com/gin-gonic/gin) - Framework web para Go
- [MongoDB](https://www.mongodb.com/) - Banco de dados NoSQL
- [Go](https://golang.org/) - Linguagem de programação
- [React](https://reactjs.org/) - Biblioteca JavaScript para construção de interfaces de usuário

## Contribuindo

Contribuições são bem-vindas! Para contribuir com o projeto, siga as etapas descritas na seção anterior.

## Autores

- **Guilherme Di Franco** - Desenvolvimento inicial - [guilherme-difranco](https://github.com/guilherme-difranco)

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo `LICENSE` para mais detalhes.
```

