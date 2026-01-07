package application

import "gin-quickstart/internal/user/domain"

type UserServices interface {
	CreateUser(req *domain.CreateUserRequest) (*domain.User, error)
	UpdateUser(id int, req *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(id int) error
	GetUserByID(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetAllUsers(params *domain.QueryParams) (*domain.QueryResult, error)
	LoginUser(email, password string) (*domain.LoginResponse, error)
}
