package application

import "gin-quickstart/internal/settings/domain"

type SettingsServices interface {
	GetAllSettings(req *domain.QueryParams) (*domain.QueryResult, error)
	GetSettingsByID(id int) (*domain.Settings, error)
	CreateSettings(req *domain.CreateSettingsRequest) (*domain.Settings, error)
	UpdateSettings(id int,req *domain.UpdateSettingsRequest) (*domain.Settings,error)
	DeleteSettings(id int) error
}