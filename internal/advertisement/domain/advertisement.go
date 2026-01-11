package domain

import "time"

type Advertisement struct {
	ID               int    			`db:"id" json:"id"`
	Title            string 			`db:"title" json:"title"`
	ImageUrl         string 			`db:"image_url" json:"image_url"`
	Link             string 			`db:"link" json:"link"`
	Position         string 			`db:"position" json:"position"`
	IsActive         bool   			`db:"is_active" json:"is_active"`
	Priority         int    			`db:"priority" json:"priority"`
	ClickCount       int     			`db:"click_count" json:"click_count"`
	ImpressionCount  int     			`db:"impression_count" json:"impression_count"`
	TargetAudience   string  			`db:"target_audience" json:"target_audience"`
	CreatedBy        int     			`db:"created_by" json:"created_by"`
	StartDate        *time.Time 	`db:"start_date" json:"start_date"`
	EndDate          *time.Time 	`db:"end_date" json:"end_date"`
	CreatedAt        time.Time 	 	`db:"created_at" json:"created_at"`
	UpdatedAt        time.Time 	 	`db:"updated_at" json:"updated_at"`
}

type adsRequest struct {
	Title            string 		`json:"title" binding:"required"`
	ImageUrl         string 		`json:"image_url" binding:"required"`
	Link             string 		`json:"link" binding:"required"`
	Position         string 		`json:"position" binding:"required"`
	IsActive         bool   		`json:"is_active"`
	Priority         int    		`json:"priority"`
	TargetAudience   string 		`json:"target_audience"`
	CreatedBy        int    		`json:"created_by"`
	StartDate        *time.Time `json:"start_date" binding:"required"`
	EndDate          *time.Time `json:"end_date" binding:"required"`
}

type CreateAdsRequest struct {
	adsRequest
}

type UpdateAdsRequest struct {
	adsRequest
	ID int `json:"id"`
}