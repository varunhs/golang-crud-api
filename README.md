# Golang CRUD Application

A simple Go REST API for user registration and login with PostgreSQL, JWT authentication, and request validation.

## Features

- User registration endpoint (`POST /register`)
- User login endpoint (`POST /login`)
- Password hashing with `bcrypt`
- JWT token generation for authentication
- PostgreSQL database integration using GORM
- Input validation for user data
- Middleware support for protected routes under `/api`

## Tech stack

- Go
- Gorilla Mux
- GORM
- PostgreSQL
- JWT (`github.com/golang-jwt/jwt/v5`)
- bcrypt (`golang.org/x/crypto/bcrypt`)
- dotenv (`github.com/joho/godotenv`)

## Project structure

- `main.go` - app entry point and route registration
- `src/config/db.go` - database connection logic
- `src/controllers/loginController.go` - register and login handlers
- `src/models/user.go` - user model definition
- `src/middleware/authMiddleware.go` - JWT auth middleware
- `src/routes/routes.go` - route definitions
- `src/utils/jwt.go` - JWT token generation
- `src/validation/validate_user_request.go` - request validation

## Environment variables

Create a `.env` file in the project root with the following values:

```env
PORT=8080
DB_URL=postgres://username:password@host:port/dbname
```

## Installation

1. Install Go and PostgreSQL.
2. Clone the repo.
3. Create a `.env` file with the required variables.
4. Run `go mod tidy` to install dependencies.

## Run the application

```bash
go run main.go
```

The server will start on the port defined by `PORT` in `.env`.

## API endpoints

### Register a new user

- Method: `POST`
- URL: `/register`
- Request body:

```json
{
  "name": "Varun H S",
  "email": "varun@example.com",
  "password": "password123"
}
```

### Login

- Method: `POST`
- URL: `/login`
- Request body:

```json
{
  "email": "varun@example.com",
  "password": "password123"
}
```

- Successful response returns a JSON object with a JWT token.

### Protected routes

The app configures protected routes under the `/api` prefix using `AuthMiddleware`.
Requests to these routes must include a valid JWT in the `Authorization` header:

```
Authorization: Bearer <token>
```

## Notes

- The current JWT secret is hard-coded to `secret` in `src/utils/jwt.go` and `src/middleware/authMiddleware.go`.
- The token expiration is set to 2 minutes.
- The `login` route should be accessed using `POST /login`.
