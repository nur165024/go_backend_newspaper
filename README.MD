# Gin Quickstart API

A RESTful API built with Go, Gin framework, and PostgreSQL.

## Features

- User management (CRUD operations)
- Authentication & Login
- Database migrations
- Clean Architecture
- SQLx for database operations

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Run migrations: `go run main.go migrate`
4. Start server: `go run main.go`

## API Endpoints

- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create user
- `POST /api/v1/users/login` - Login user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## Environment Variables

Create `.env` file with:

```
https://dbdiagram.io/d/NEWS-PAPER-68b91b1c61a46d388e5f124a
`
```
