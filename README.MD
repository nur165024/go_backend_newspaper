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

````
## Database Schema

```mermaid
erDiagram
    users {
        int id PK
        varchar name
        varchar user_name UK
        varchar email UK
        varchar password
        varchar designation
        text bio
        varchar profile_picture
        boolean is_active
        boolean is_verified
        varchar verification_token
        varchar reset_password_token
        timestamp reset_password_expires
        timestamp last_login
        int role_id FK
        timestamp created_at
        timestamp updated_at
    }

    categories {
        int id PK
        varchar name
        varchar slug UK
        text description
        varchar image_url
        int sort_order
        boolean is_active
        varchar meta_title
        text meta_description
        varchar meta_keywords
        timestamp created_at
        timestamp updated_at
    }

    news {
        int id PK
        varchar title
        varchar slug UK
        text content
        text excerpt
        varchar featured_image
        json images
        varchar language
        boolean is_premium
        varchar status
        boolean is_featured
        timestamp published_at
        int category_id FK
        int user_id FK
        timestamp created_at
        timestamp updated_at
    }

    news_comments {
        int id PK
        int news_id FK
        int user_id FK
        text comment_text
        varchar status
        timestamp created_at
        timestamp updated_at
    }

    news_likes {
        int id PK
        int news_id FK
        int user_id FK
        varchar like_type
        timestamp created_at
        timestamp updated_at
    }

    news_views {
        int id PK
        int news_id FK
        int user_id FK
        varchar ip_address
        timestamp viewed_at
        timestamp created_at
        timestamp updated_at
    }

    tags {
        int id PK
        varchar name
        varchar slug UK
        timestamp created_at
        timestamp updated_at
    }

    news_tags {
        int id PK
        int news_id FK
        int tag_id FK
        timestamp created_at
        timestamp updated_at
    }

    polls {
        int id PK
        varchar title
        text description
        text question
        varchar poll_type
        varchar status
        boolean is_featured
        boolean allow_anonymous
        int max_votes_per_user
        timestamp start_date
        timestamp end_date
        int created_by FK
        int total_votes
        timestamp created_at
        timestamp updated_at
    }

    poll_options {
        int id PK
        int poll_id FK
        varchar option_text
        varchar option_image
        int sort_order
        int vote_count
        timestamp created_at
        timestamp updated_at
    }

    votes {
        int id PK
        int poll_id FK
        int option_id FK
        int user_id FK
        varchar ip_address
        varchar user_agent
        int rating_value
        timestamp voted_at
        timestamp created_at
        timestamp updated_at
    }

    roles {
        int id PK
        varchar name
        text description
        timestamp created_at
        timestamp updated_at
    }

    permissions {
        int id PK
        varchar name
        text description
        varchar resource
        varchar action
        timestamp created_at
        timestamp updated_at
    }

    role_permissions {
        int id PK
        int role_id FK
        int permission_id FK
        timestamp created_at
        timestamp updated_at
    }

    %% Relationships
    users ||--o{ news : creates
    categories ||--o{ news : contains
    news ||--o{ news_comments : has
    news ||--o{ news_likes : has
    news ||--o{ news_views : has
    news ||--o{ news_tags : tagged
    tags ||--o{ news_tags : used_in
    users ||--o{ news_comments : writes
    users ||--o{ news_likes : gives
    users ||--o{ news_views : views
    users ||--o{ polls : creates
    polls ||--o{ poll_options : has
    polls ||--o{ votes : receives
    poll_options ||--o{ votes : gets
    users ||--o{ votes : casts
    roles ||--o{ users : assigned
    roles ||--o{ role_permissions : has
    permissions ||--o{ role_permissions : granted


`
````
