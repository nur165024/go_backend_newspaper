package application

import (
	"fmt"
	"gin-quickstart/internal/settings/domain"
)

type settingsServices struct {
	settingsRepo domain.SettingsRepository
}

func NewSettingsServices(settingsRepo domain.SettingsRepository) *settingsServices {
	return &settingsServices{
		settingsRepo: settingsRepo,
	}
}

// get all
func (s *settingsServices) GetAllSettings(params *domain.QueryParams) (*domain.QueryResult, error) {
	settings, err := s.settingsRepo.GetAll(params)

	if err != err {
		return nil, err
	}

	return settings, nil
}

// get by id
func (s *settingsServices) GetSettingsByID(id int) (*domain.Settings, error) {
	settings, err := s.settingsRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return settings, nil
}

// create
func (s *settingsServices) CreateSettings(req *domain.CreateSettingsRequest) (*domain.Settings, error) {
	createSettings, err := s.settingsRepo.Create(req)

	if err != nil {
		return nil, fmt.Errorf("failed to create Settings: %w", err)
	}

	return createSettings, nil
}

// update
func (s *settingsServices) UpdateSettings(id int,req *domain.UpdateSettingsRequest) (*domain.Settings, error) {
	settings,err := s.settingsRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	// update fields
	updateInfo := updateCategoryFields(settings, req)

	// update other fields as needed
	updatedSettings, err := s.settingsRepo.Update(id, updateInfo)

	if err != nil {
		return nil, err
	}

	return updatedSettings, nil
}

// delete
func (s *settingsServices) DeleteSettings(id int) error {
	err := s.settingsRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// helper function 
func updateCategoryFields(settings *domain.Settings , req *domain.UpdateSettingsRequest) *domain.UpdateSettingsRequest   {
	if req.Key != "" {
		settings.Key = req.Key
	}
	if req.Value != "" {
		settings.Value = req.Value
	}
	if req.Description != "" {
		settings.Description = req.Description
	}
	if req.DataType != "" {
		settings.DataType = req.DataType
	}
	if req.Category != "" {
		settings.Category = req.Category
	}
	if req.IsPublic != false {
		settings.IsPublic = req.IsPublic
	}
	if req.IsEditable != false {
		settings.IsEditable = req.IsEditable
	}

	return req
}
