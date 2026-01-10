package interfaces

import (
	"gin-quickstart/internal/settings/application"
	"gin-quickstart/internal/settings/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SettingsHandler struct {
	settingsServices application.SettingsServices
}

func NewSettingsHandler(settingsServices application.SettingsServices) *SettingsHandler {
	return &SettingsHandler{settingsServices: settingsServices}
}

// GET /settings - settings all
func (h *SettingsHandler) GetAllSettings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "DESC")
	search := c.Query("search") 

	params := &domain.QueryParams{
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
		Order:    order,
		Search:   search,
	}

	settings, err := h.settingsServices.GetAllSettings(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": settings})
}

// GET /settings/:id - settings
func (h *SettingsHandler) GetSettingsByID(c *gin.Context) {
	settingsId := c.Param("id")
	id, _ := strconv.Atoi(settingsId)

	settings, err := h.settingsServices.GetSettingsByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": settings})
}

// POST /settings - settings
func (h *SettingsHandler) CreateSettings(c *gin.Context) {
	var req *domain.CreateSettingsRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := h.settingsServices.CreateSettings(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Settings created successfully",
		"data":    settings,
	})
}

// PUT /settings/:id - settings
func (h *SettingsHandler) UpdateSettings(c *gin.Context) {
	settingsId := c.Param("id")
	id, _ := strconv.Atoi(settingsId)

	var req *domain.UpdateSettingsRequest

	err := c.ShouldBindJSON(&req)

		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	settings, err := h.settingsServices.UpdateSettings(id, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": settings})
}

// DELETE /settings/:id - settings
func (h *SettingsHandler) DeleteSettings(c *gin.Context) {
	settingsId := c.Param("id")
	id, _ := strconv.Atoi(settingsId)

	err := h.settingsServices.DeleteSettings(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": "Settings deleted successfully"})
}


