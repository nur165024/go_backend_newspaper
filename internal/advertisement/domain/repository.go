package domain

type QueryParams struct {
	// pagination
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	// sorting
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
	// searching
	Search string `json:"search"`
}

type QueryResponse struct {
	Data       []*Advertisement `json:"data"`
	TotalItem  int64            `json:"total_item"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}

type AdsRepository interface {
	GetAll(params *QueryParams) (*QueryResponse, error)
	GetByID(id int) (*Advertisement, error)
	Create(ads *CreateAdsRequest) (*Advertisement, error)
	Update(id int, ads *UpdateAdsRequest) (*Advertisement, error)
	Delete(id int) error
}