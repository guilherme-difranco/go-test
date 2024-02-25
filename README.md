# API de Gerenciamento de Tarefas

Esta API fornece um sistema simplificado para o gerenciamento de tarefas, permitindo criar, listar, atualizar e deletar tarefas. Desenvolvida com foco em performance e escalabilidade, suporta alto volume de requisições, demonstrando práticas de clean architecture, caching, e otimizações específicas para o MongoDB.

## Iniciando

Instruções para configurar a API localmente para desenvolvimento e testes.

### Pré-requisitos

Para executar este projeto, você precisará ter o Go instalado em sua máquina:

```bash
go version
```

Além disso, é necessário ter o MongoDB rodando localmente ou em um ambiente de nuvem.

### Instalação

Clone o projeto para sua máquina local:

```bash
git clone https://github.com/guilherme-difranco/go-test
```

Instale as dependências do projeto:

```bash
cd go-test
go mod tidy
```

Execute o servidor localmente:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

## Testes

Para executar os testes automatizados:

```bash
go test ./...
```

## Construído Com

- [Gin](https://github.com/gin-gonic/gin) - Framework web para Go
- [MongoDB](https://www.mongodb.com/) - Banco de dados NoSQL
- [Go](https://golang.org/) - Linguagem de programação

## Contribuindo

Contribuições são bem-vindas! Para contribuir com o projeto, siga estas etapas:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b <feature-branch>`)
3. Faça commit de suas mudanças (`git commit -m 'Add some <feature-branch>'`)
4. Push para a branch (`git push origin <feature-branch>`)
5. Abra um Pull Request

## Autores

- **Seu Nome** - Desenvolvimento inicial - [guilherme-difranco](https://github.com/guilherme-difranco)

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo `LICENSE` para mais detalhes.



- Inspirado nas melhores práticas de desenvolvimento de software.

