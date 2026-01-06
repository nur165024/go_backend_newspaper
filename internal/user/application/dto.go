package application

import (
	"gin-quickstart/internal/user/domain"
	"time"
)

// create user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// login user
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


// update user
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password"`
}

// response DTOs
type UserResponse struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Designation string `json:"designation"`
	Bio string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	IsActive bool `json:"is_active"`
	IsVerified bool `json:"is_verified"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Data      []*UserResponse `json:"data"`
	TotalItems int64          `json:"total_items"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}

// conversion methods
func (r *CreateUserRequest) ToDomain() *domain.User {
	return &domain.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func ToUserResponse(user *domain.User) *domain.User {
	return &domain.User{
		ID:    user.ID,
		Name:  user.Name,
		UserName: user.UserName,
		Email: user.Email,
		Designation: user.Designation,
		Bio: user.Bio,
		ProfilePicture: user.ProfilePicture,
		IsActive: user.IsActive,
		IsVerified: user.IsVerified,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserListResponse(result *domain.QueryResult) *domain.QueryResult {
	users := make([]*domain.User, len(result.Data))

	for i, user := range result.Data {
		users[i] = &domain.User{
			ID:    user.ID,
			Name:  user.Name,
			UserName: user.UserName,
			Email: user.Email,
			Designation: user.Designation,
			Bio: user.Bio,
			ProfilePicture: user.ProfilePicture,
			IsActive: user.IsActive,
			IsVerified: user.IsVerified,
			LastLogin: user.LastLogin,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return &domain.QueryResult{
		Data:       users,
		TotalItem: result.TotalItem,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}
}
