package domain

import "time"

type User struct {
    ID                   int        `db:"id" json:"id"`
    Name                 string     `db:"name" json:"name"`
    UserName             string     `db:"user_name" json:"user_name"`
    Email                string     `db:"email" json:"email"`
    Password             string     `db:"password" json:"-"` // Hidden from JSON
    Designation          string     `db:"designation" json:"designation"`
    Bio                  string     `db:"bio" json:"bio"`
    ProfilePicture       string     `db:"profile_picture" json:"profile_picture"`
    IsActive             bool       `db:"is_active" json:"is_active"`
    IsVerified           bool       `db:"is_verified" json:"is_verified"`
    VerificationToken    string     `db:"verification_token" json:"-"`
    ResetPasswordToken   string     `db:"reset_password_token" json:"-"`
    ResetPasswordExpires *time.Time `db:"reset_password_expires" json:"-"` // Nullable
    LastLogin            *time.Time `db:"last_login" json:"last_login"`     // Nullable
    CreatedAt            time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt            time.Time  `db:"updated_at" json:"updated_at"`
}

type CreateUserRequest struct {
    Name     string `json:"name" binding:"required"`
    UserName string `json:"user_name"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    Designation string `json:"designation"`
    Bio string `json:"bio"`
    ProfilePicture string `json:"profile_picture"`
}

type LoginUserRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
    Name     string `json:"name"`
    UserName string `json:"user_name"`
    Email    string `json:"email"`
    Designation string `json:"designation"`
    Bio string `json:"bio"`
    ProfilePicture string `json:"profile_picture"`
    IsActive bool `json:"is_active"`
    IsVerified bool `json:"is_verified"`
}