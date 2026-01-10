package domain

import "time"

type Settings struct {
	ID          int    `db:"id" json:"id"`
	Key         string `db:"key" json:"key"`
	Value       string `db:"value" json:"value"`
	Description string `db:"description" json:"description"`
	DataType    string `db:"data_type" json:"data_type"`
	Category    string `db:"category" json:"category"`
	IsPublic    bool   `db:"is_public" json:"is_public"`
	IsEditable  bool   `db:"is_editable" json:"is_editable"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type settingsRequest struct {
	Key         string `json:"key" binding:"required"`
	Value       string `json:"value" binding:"required"`
	Description string `json:"description"`
	DataType    string `json:"data_type"`
	Category    string `json:"category"`
	IsPublic    bool   `json:"is_public"`
	IsEditable  bool   `json:"is_editable"`
}

// Create with validation
type CreateSettingsRequest struct {
	settingsRequest
}

// Update with validation
type UpdateSettingsRequest struct {
	settingsRequest
	ID int `json:"id"`
}