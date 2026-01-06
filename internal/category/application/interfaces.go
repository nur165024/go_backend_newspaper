package application

import "gin-quickstart/internal/category/domain"

type CategoryServices interface {
	CreateCategory(req *domain.CreateCategoryRequest) (*domain.Category, error)
	UpdateCategory(id int, req *domain.UpdateCategoryRequest) (*domain.Category, error)
	DeleteCategory(id int) error
	GetCategoryByID(id int) (*domain.Category, error)
	GetAllCategories(params *domain.QueryParams) (*domain.QueryResult, error)
}