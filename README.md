# Todo App

A RESTful API for managing todo lists and items, built with Go and Gin framework.

## Project Overview

This Todo App provides a complete backend solution for managing todo lists and items. It includes user authentication, todo list management, and item management within lists.

## Architecture

The project follows a clean architecture approach with the following components:

- **Models**: Data structures representing the domain entities
- **Repository**: Data access layer for interacting with the database
- **Services**: Business logic layer that implements the application's use cases
- **Routes**: HTTP handlers for processing API requests
- **Configs**: Application configuration management

### Directory Structure

```
todo-app/
├── cmd/            # Application entry points
├── configs/        # Configuration files
├── internal/       # Private application code
│   ├── models/     # Domain models
│   ├── repository/ # Data access layer
│   ├── routes/     # HTTP handlers
│   ├── services/   # Business logic
│   └── requests/   # Request models
├── pkg/            # Public libraries
│   ├── db/         # Database utilities
│   └── handler/    # HTTP server utilities
└── scripts/        # Utility scripts
```

## Features

- User registration and authentication with JWT
- Create, read, update, and delete todo lists
- Create, read, update, and delete todo items within lists
- Secure API endpoints with authentication middleware

## API Endpoints

### Authentication

- `POST /auth/sign-up` - Register a new user
- `POST /auth/sign-in` - Login and get JWT token

### Todo Lists

- `POST /api/lists` - Create a new todo list
- `GET /api/lists` - Get all todo lists for the authenticated user
- `GET /api/lists/:id` - Get a specific todo list by ID
- `PUT /api/lists/:id` - Update a todo list
- `DELETE /api/lists/:id` - Delete a todo list

### Todo Items

- `POST /api/lists/:id/items` - Create a new item in a todo list
- `GET /api/lists/:id/items` - Get all items in a todo list
- `GET /api/lists/:id/items/:item_id` - Get a specific item by ID
- `PUT /api/lists/:id/items/:item_id` - Update an item
- `DELETE /api/lists/:id/items/:item_id` - Delete an item

## Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **PostgreSQL** - Database
- **JWT** - Authentication
- **Viper** - Configuration management
- **godotenv** - Environment variable management

## Setup Instructions

### Prerequisites

- Go 1.16+
- PostgreSQL
- Git

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:
   ```
   DB_PASSWORD=your_db_password
   HASH_SALT=your_hash_salt
   JWT_KEY=your_jwt_secret
   ```

4. Configure the application in `configs/config.yaml`:
   ```yaml
   port: "8082"

   db:
     host: "localhost"
     port: "5432"
     username: "localhost"
     password: "root"  # It's recommended to use DB_PASSWORD from .env instead
     DBName: "mydb"
     SSLMode: "disable"
     DBMaxCons: 30  # Note: In config.yaml it's "DBMAxCons" but code uses "DBMaxCons"
     DBMinCons: 5
     DBMaxLifetime: "12"
     DBMaxIdTime: "2"  # Note: In config.yaml it's "DBMinLifetime" but code uses "DBMaxIdTime"
   ```

5. Run the application:
   ```
   go run cmd/main.go
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
