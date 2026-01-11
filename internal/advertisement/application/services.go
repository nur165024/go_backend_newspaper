package application

import "gin-quickstart/internal/advertisement/domain"

type adsServices struct {
	adsRepo domain.AdsRepository
}

func NewAdsServices(adsRepo domain.AdsRepository) *adsServices {
	return &adsServices{
		adsRepo: adsRepo,
	}
}

// get all
func (s *adsServices) GetAllAds(params *domain.QueryParams) (*domain.QueryResponse, error) {
	ads, err := s.adsRepo.GetAll(params)

	if err != nil {
		return nil, err
	}
	return ads, nil
}

// get by id
func (s *adsServices) GetAdsByID(id int) (*domain.Advertisement, error) {
	ads, err := s.adsRepo.GetByID(id)

	if err != nil {
		return nil, err
	}
	return ads, nil
}

// create
func (s *adsServices) CreateAds(req *domain.CreateAdsRequest) (*domain.Advertisement, error) {
	ads, err := s.adsRepo.Create(req)

	if err != nil {
		return nil, err
	}
	return ads, nil
}

// update
func (s *adsServices) UpdateAds(id int, req *domain.UpdateAdsRequest) (*domain.Advertisement, error) {
	ads, err := s.adsRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	updateInfo := updateCategoryFields(ads, req)

	// Update other fields as needed
	updateAds, err := s.adsRepo.Update(id, updateInfo)

	if err != nil {
		return nil, err
	}

	return updateAds, nil
}

// delete
func (s *adsServices) DeleteAds(id int) error {
	err := s.adsRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// helper function 
func updateCategoryFields(ads *domain.Advertisement, req *domain.UpdateAdsRequest) *domain.UpdateAdsRequest {
	if req.Title != "" {
		ads.Title = req.Title
	}
	if req.ImageUrl != "" {
		ads.ImageUrl = req.ImageUrl
	}
	if req.Link != "" {
		ads.Link = req.Link
	}
	if req.Position != "" {
		ads.Position = req.Position
	}
	if req.TargetAudience != "" {
		ads.TargetAudience = req.TargetAudience
	}
	
	// Boolean field - always update
	ads.IsActive = req.IsActive
	
	// Integer fields - check if not zero
	if req.Priority != 0 {
		ads.Priority = req.Priority
	}
	if req.CreatedBy != 0 {
		ads.CreatedBy = req.CreatedBy
	}
	
	// Time fields - check if not zero time
	if !req.StartDate.IsZero() {
		ads.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		ads.EndDate = req.EndDate
	}

	return req
}