package application

import (
	"fmt"
	"gin-quickstart/internal/category/domain"
)

type categoryServices struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryServices(categoryRepo domain.CategoryRepository) *categoryServices {
	return &categoryServices{
		categoryRepo: categoryRepo,
	}
}

// get all
func (s *categoryServices) GetAllCategories(params *domain.QueryParams) (*domain.QueryResponse, error) {
	categories, err := s.categoryRepo.GetAll(params)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// get by id
func (s *categoryServices) GetCategoryByID(id int) (*domain.Category, error) {
	category, err := s.categoryRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// create
func (s *categoryServices) CreateCategory(req *domain.CreateCategoryRequest) (*domain.Category, error){
	
	createdCategory, err := s.categoryRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return createdCategory, nil
}

// update 
func (s *categoryServices) UpdateCategory(id int, req *domain.UpdateCategoryRequest) (*domain.Category, error) {
	// Get existing category
	category, err := s.categoryRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	updateInfo := updateCategoryFields(category, req)

	// Update other fields as needed
	updateCategory, err := s.categoryRepo.Update(id, updateInfo)

	if err != nil {
		return nil, err
	}

	return updateCategory, nil
}

// delete
func (s *categoryServices) DeleteCategory(id int) error {
	err := s.categoryRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// helper function 
func updateCategoryFields(category *domain.Category, req *domain.UpdateCategoryRequest) *domain.UpdateCategoryRequest {
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Slug != "" {
		category.Slug = req.Slug
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.ImageUrl != "" {
		category.ImageUrl = req.ImageUrl
	}
	if req.SortOrder != 0 {
		category.SortOrder = req.SortOrder
	}
	if req.IsActive != false {
		category.IsActive = req.IsActive
	}
	if req.MetaTitle != "" {
		category.MetaTitle = req.MetaTitle
	}
	if req.MetaDescription != "" {
		category.MetaDescription = req.MetaDescription
	}

	return req
}