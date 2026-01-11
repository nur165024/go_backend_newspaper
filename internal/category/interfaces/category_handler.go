package interfaces

import (
	"gin-quickstart/internal/category/application"
	"gin-quickstart/internal/category/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryServices application.CategoryServices
}

func NewCategoryHandler(categoryServices application.CategoryServices) *CategoryHandler {
	return &CategoryHandler{categoryServices: categoryServices}
}

// GET /categories - get all categories
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
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

	categories, err := h.categoryServices.GetAllCategories(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GET /categories/:id - get category by id
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	categoryId := c.Param("id")
	id, _ := strconv.Atoi(categoryId)

	category, err := h.categoryServices.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// POST /categories - create category
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req *domain.CreateCategoryRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryServices.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})
}

// PUT /categories/:id - update category
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, _ := strconv.Atoi(categoryId)

	var req *domain.UpdateCategoryRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryServices.UpdateCategory(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DELETE /categories/:id - delete category
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, _ := strconv.Atoi(categoryId)

	err := h.categoryServices.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Category deleted successfully"})
}





