package domain

type QueryParams struct {
	// pagination
	Page     int `json:"page" form:"page" query:"page"`
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`
	// sorting
	SortBy string `json:"sort_by" form:"sort_by" query:"sort_by"`
	Order  string `json:"order" form:"order" query:"order"`
	// searching
	Search string `json:"search" form:"search" query:"search"`
}

type QueryResult struct {
	Data       []*Category `json:"data"`
	TotalItem  int64       `json:"total_item"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

type CategoryRepository interface {
	GetAll(params *QueryParams) (*QueryResult, error)
	GetByID(id int) (*Category, error)
	Create(category *CreateCategoryRequest) (*Category, error)
	Update(id int, category *UpdateCategoryRequest) (*Category, error)
	Delete(id int) error
}