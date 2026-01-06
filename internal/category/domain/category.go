package domain

type Category struct {
	ID              int    `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Slug            string `db:"slug" json:"slug"`
	Description     string `db:"description" json:"description"`
	ImageUrl        string `db:"image_url" json:"image_url"`
	SortOrder       int    `db:"sort_order" json:"sort_order"`
	IsActive        bool   `db:"is_active" json:"is_active"`
	MetaTitle       string `db:"meta_title" json:"meta_title"`
	MetaDescription string `db:"meta_description" json:"meta_description"`
	MetaKeywords    string `db:"meta_keywords" json:"meta_keywords"`
	CreatedAt       string `db:"created_at" json:"created_at"`
	UpdatedAt       string `db:"updated_at" json:"updated_at"`
}

// base struct for command filed
type categoryRequest struct {
	Name            string `json:"name" binding:"required"`
	Slug            string `json:"slug" binding:"required"`
	Description     string `json:"description"`
	ImageUrl        string `json:"image_url"`
	SortOrder       int    `json:"sort_order"`
	IsActive        bool   `json:"is_active"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
}

// Create with validation
type CreateCategoryRequest struct {
	categoryRequest
}

// Update with validation
type UpdateCategoryRequest struct {
	categoryRequest
}