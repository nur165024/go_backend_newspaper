package interfaces

import (
	"fmt"
	"gin-quickstart/internal/user/application"
	"gin-quickstart/internal/user/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userServices application.UserServices
}

func NewUserHandler(userServices application.UserServices) *UserHandler {
	return &UserHandler{
		userServices: userServices,
	}
}

// POST /users - Create user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req domain.CreateUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Request: %+v\n", req)

	user, err := h.userServices.CreateUser(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// PUT /users/:id - update user by id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	var req domain.UpdateUserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userServices.UpdateUser(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id - delete user by id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	err := h.userServices.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}

// POST /users/login - Login user
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req domain.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userServices.LoginUser(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GET /users/:id - get user by id
func (h *UserHandler) GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	user, err := h.userServices.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users - get all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
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
		Search: search,
	}

	result, err := h.userServices.GetAllUsers(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GET /users/:email - get user by email
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.userServices.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}


