package domain

type QueryParams struct {
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"page_size" form:"page_size" query:"page_size"`
	Search   string `json:"search" form:"search" query:"search"`
	SortBy   string `json:"sort_by" form:"sort_by" query:"sort_by"`
	Order    string `json:"order" form:"order" query:"order"`
}

type QueryResult struct {
	Data       []*Settings `json:"total_count"`
	TotalItem  int64       `json:"total_item"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

type SettingsRepository interface {
	GetAll(params *QueryParams) (*QueryResult, error)
	GetByID(id int) (*Settings, error)
	Create(category *Settings) (*Settings, error)
	Update(id int, category *Settings) (*Settings, error)
	Delete(id int) error
}