# CRUD Users API

API REST desenvolvida em Go para gerenciamento de usuários, utilizando PostgreSQL como banco de dados e Docker para executar o banco.

## Tecnologias

- Go
- Chi Router
- PostgreSQL
- pgx
- Docker
- Docker Compose

## Funcionalidades

- Criar usuário
- Listar todos os usuários
- Buscar usuário por ID
- Atualizar usuário
- Excluir usuário

## Estrutura do usuário

```json
{
  "id": "uuid",
  "first_name": "Renato",
  "last_name": "Pereira",
  "biography": "Biografia do usuário com pelo menos vinte caracteres."
}
```

O ID de cada usuário é um UUID gerado automaticamente.

## Como executar

### 1. Clone o repositório

```bash
git clone https://github.com/renatopgn/go-api-users.git
cd go-api-users
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto baseado no `.env.example`:

```env
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=crud_users

DATABASE_URL=postgres://your_user:your_password@localhost:5432/crud_users
```

O arquivo `.env` não deve ser enviado para o GitHub.

### 3. Inicie o PostgreSQL

```bash
docker compose up -d
```

### 4. Execute a API

```bash
go run .
```

A API estará disponível em:

```text
http://localhost:8080
```

## Endpoints

| Método | Endpoint | Descrição |
|---|---|---|
| GET | `/api/users` | Lista todos os usuários |
| GET | `/api/users/{id}` | Busca um usuário pelo ID |
| POST | `/api/users` | Cria um usuário |
| PUT | `/api/users/{id}` | Atualiza um usuário |
| DELETE | `/api/users/{id}` | Exclui um usuário |

## Validações

- `first_name`: entre 2 e 20 caracteres
- `last_name`: entre 2 e 20 caracteres
- `biography`: entre 20 e 450 caracteres