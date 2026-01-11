package interfaces

import (
	"gin-quickstart/internal/advertisement/application"
	"gin-quickstart/internal/advertisement/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdsHandler struct {
	adsServices application.AdvertisementServices
}

func NewAdsHandler(adsServices application.AdvertisementServices) *AdsHandler {
	return &AdsHandler{adsServices: adsServices}
}

// GET /ads - get all ads
func (h *AdsHandler) GetAllAds(c *gin.Context) {
	// Parse query parameters with defaults
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

	ads, err := h.adsServices.GetAllAds(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ads})
}

// GET /ads/:id - get by id
func (h *AdsHandler) GetAdsByID(c *gin.Context) {
	adsId := c.Param("id")
	id, _ := strconv.Atoi(adsId)

	ads, err := h.adsServices.GetAdsByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ads})
}

// POST /ads - create ads
func (h *AdsHandler) CreateAds(c *gin.Context) {
	var req *domain.CreateAdsRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ads, err := h.adsServices.CreateAds(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Advertisement created successfully",
		"data":    ads,
	})
}

// PUT /ads/:id - update ads
func (h *AdsHandler) UpdateAds(c *gin.Context) {
	adsId := c.Param("id")
	id, _ := strconv.Atoi(adsId)

	var req *domain.UpdateAdsRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ads, err := h.adsServices.UpdateAds(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ads})
}

// DELETE /ads/:id - delete ads
func (h *AdsHandler) DeleteAds(c *gin.Context) {
	adsId := c.Param("id")
	id, _ := strconv.Atoi(adsId)

	err := h.adsServices.DeleteAds(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Advertisement deleted successfully"})
}