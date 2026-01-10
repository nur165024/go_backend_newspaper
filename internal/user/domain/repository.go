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
	Data       []*User `json:"data"`
	TotalItem  int64   `json:"total_item"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
	TotalPages int     `json:"total_pages"`
}

type UserRepository interface {
	GetAll(params *QueryParams) (*QueryResult, error)
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *CreateUserRequest) (*User, error)
	Update(id int, user *UpdateUserRequest) (*User, error)
	Delete(id int) error
}