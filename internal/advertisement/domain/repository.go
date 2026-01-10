package domain

type AdvertisementRepository interface {
	GetAll()
	GetByID(id int) (*Advertisement, error)
	Create()
	Update()
	Delete()
}