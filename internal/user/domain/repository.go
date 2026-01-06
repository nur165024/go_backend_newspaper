package domain

type QueryParams struct {
	// pagination
	Page     int `json:"page" form:"page" query:"page"`
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`
	// sorting
	SortBy string `json:"sort_by" form:"sort_by" query:"sort_by"`
	Order  string `json:"order" form:"order" query:"order"`
	// filtering
	Search string            `json:"search" form:"search" query:"search"`
	Filter map[string]string `json:"filter" form:"filter" query:"filter"`
}

type QueryResult struct {
	Data       []*User `json:"data"`
	TotalItem  int64   `json:"total_item"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
	TotalPages int     `json:"total_pages"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	LoginUser(email, password string) (*User, error)
	GetByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
	GetAll(params *QueryParams) (*QueryResult, error)
}