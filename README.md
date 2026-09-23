# Go Users API

API RESTful desenvolvida em Go para gerenciamento de usuários, utilizando armazenamento em memória com `map`.

O projeto foi desenvolvido com o objetivo de praticar conceitos de HTTP, APIs REST, JSON e CRUD em Go.

## Tecnologias

- Go
- Chi Router
- Google UUID

## Funcionalidades

- Criar usuários
- Listar usuários
- Buscar usuário por ID
- Atualizar usuários
- Deletar usuários
- Validação dos dados
- Tratamento de erros em JSON

## Endpoints

| Método | Endpoint | Descrição |
|---|---|---|
| POST | `/api/users` | Cria um usuário |
| GET | `/api/users` | Lista todos os usuários |
| GET | `/api/users/{id}` | Busca um usuário pelo ID |
| PUT | `/api/users/{id}` | Atualiza um usuário |
| DELETE | `/api/users/{id}` | Remove um usuário |

## Executando o projeto

```bash
go mod download
go run .
```

O servidor será iniciado em:

```text
http://localhost:8080
```

## Observação

Os dados são armazenados apenas em memória. Portanto, todos os usuários cadastrados são perdidos quando o servidor é reiniciado.