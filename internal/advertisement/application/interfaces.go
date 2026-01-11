package application

import "gin-quickstart/internal/advertisement/domain"

type AdvertisementServices interface {
	GetAllAds(params *domain.QueryParams) (*domain.QueryResponse, error)
	GetAdsByID(id int) (*domain.Advertisement, error)
	CreateAds(ads *domain.CreateAdsRequest) (*domain.Advertisement, error)
	UpdateAds(id int, ads *domain.UpdateAdsRequest) (*domain.Advertisement, error)
	DeleteAds(id int) error
}