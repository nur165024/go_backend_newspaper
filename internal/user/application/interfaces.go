package application

import "gin-quickstart/internal/user/domain"

type UserService interface {
	CreateUser(req *domain.CreateUserRequest) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	LoginUser(email, password string) (*domain.User, error)
	UpdateUser(id int, req *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(id int) error
	GetAllUsers(params *domain.QueryParams) (*domain.QueryResult, error)
}
