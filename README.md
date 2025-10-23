# GoFrame

GoFrame is a minimalist Go API framework built with Gin, GORM, MySQL, and Redis. It provides a clean structure for building RESTful APIs with database migrations, CLI commands, and modular architecture.

## Features

- **RESTful API**: Built-in CRUD operations for users with structured JSON responses.
- **Database Support**: MySQL integration with GORM ORM.
- **Caching**: Redis support for performance.
- **Migrations**: Database schema management with golang-migrate.
- **CLI Commands**: Easy-to-use commands for serving, migrating, and creating migrations.
- **Modular Structure**: Organized into controllers, services, models, and routes.
- **Environment Configuration**: Flexible config loading from .env files.
- **Validation**: Input validation with go-playground/validator.
- **Logging**: Structured logging with Zap.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/rezhasanjaya/goframe.git
   cd goframe
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your environment variables by creating a `.env` file:

   ```env
   APP_PORT=8080
   DB_HOST=127.0.0.1
   DB_USER=root
   DB_PASS=your_password
   DB_NAME=goframe
   DB_PORT=3306
   REDIS_HOST=127.0.0.1:6379
   ```

4. Run database migrations:
   ```bash
   go run main.go migrate
   ```

## Usage

### Running the Server

Start the HTTP server:

```bash
go run main.go serve
```

The server will run at `http://localhost:8080`.

### CLI Commands

- **Serve**: Start the server

  ```bash
  go run main.go serve
  ```

- **Migrate**: Run all pending migrations

  ```bash
  go run main.go migrate
  ```

- **Rollback**: Rollback the last migration

  ```bash
  go run main.go migrate:rollback
  ```

- **Create Migration**: Create new migration files
  ```bash
  go run main.go create:migration create_example_table
  ```

## API Documentation

### Base URL

```
http://localhost:8080/api
```

### Users Endpoints

#### Get All Users

- **GET** `/api/users`
- **Response**: List of users

#### Get User by ID

- **GET** `/api/users/:id`
- **Response**: Single user object

#### Create User

- **POST** `/api/users`
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response**: Created user object

#### Update User

- **PUT** `/api/users/:id`
- **Body**: Partial user data
- **Response**: Updated user data

#### Delete User

- **DELETE** `/api/users/:id`
- **Response**: No content

### Response Format

All API responses follow a consistent structure:

**Success Response**:

```json
{
  "code": 200,
  "status": true,
  "message": "Users retrieved",
  "data": [...]
}
```

**Error Response**:

```json
{
  "code": 400,
  "status": false,
  "message": "Failed to fetch users",
  "error": "error details"
}
```

**Validation Error Response**:

```json
{
  "code": 422,
  "status": false,
  "message": "Validation failed",
  "errors": "validation details"
}
```

## Project Structure

```
goframe/
├── cmd/                    # CLI commands
│   ├── root.go
│   ├── serve.go
│   └── migrate.go
├── config/                 # Configuration (legacy)
├── internal/
│   ├── app/
│   │   ├── http/
│   │   │   ├── controllers/    # HTTP controllers
│   │   │   │   ├── api/        # API controllers
│   │   │   │   ├── base_controller.go
│   │   │   │   └── welcome_controller.go
│   │   │   ├── middleware/     # Middleware
│   │   │   └── routes/         # Route definitions
│   │   ├── models/             # GORM models
│   │   ├── services/           # Business logic
│   │   └── validators/         # Input validators
│   └── core/
│       ├── bootstrap/          # App initialization
│       ├── config/             # Configuration management
│       └── utils/              # Utility functions
├── migrations/             # Database migrations
├── public/                 # Static assets
├── tests/                  # Test files
├── main.go                 # Entry point
├── go.mod
└── README.md
```

## Configuration

Configuration is loaded from environment variables or a `.env` file. See `internal/core/config/config.go` for available options.

## Development

### Adding New Features

1. Create models in `internal/app/models/`
2. Implement services in `internal/app/services/`
3. Add controllers in `internal/app/http/controllers/`
4. Register routes in `internal/app/routes/`

### Running Tests

```bash
go test ./...
```

## Upcoming Features

- Soft deletes for models
- JWT authentication middleware
- Enhanced input validation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

[Rezha Sanjaya](https://github.com/rezhasanjaya)
