package application

import (
	"errors"
	"gin-quickstart/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// get all users
func (s *userService) GetAllUsers(params *domain.QueryParams) (*domain.QueryResult, error) {
	// Set defaults if params is nil
	if params == nil {
		params = &domain.QueryParams{
			Page:     1,
			PageSize: 10,
			SortBy:   "id",
			Order:    "DESC",
		}
	}

	result, err := s.userRepo.GetAll(params)
	
	if err != nil {
		return nil, err
	}

	return result, nil
}

// create user
func (s *userService) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Check if user already exists
	existingUser, err  := s.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {  // ✅ Fixed: Check both error and user
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Convert to domain
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		UserName: req.UserName,
		Password: string(hashedPassword),
		Designation: req.Designation,
		Bio: req.Bio,
		ProfilePicture: req.ProfilePicture,
		IsActive: true,
		IsVerified: true,
	}

	// Create user
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// get user by id
func (s *userService) GetUserByID(id int) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// find user by email
func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// login user
func (s *userService) LoginUser(email, password string) (*domain.User, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// update user
func (s *userService) UpdateUser(id int, req *domain.UpdateUserRequest) (*domain.User, error) {
	// Get existing user
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
        user.Name = req.Name
	}
	if req.UserName != "" {
		user.UserName = req.UserName
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Designation != "" {
		user.Designation = req.Designation
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = req.ProfilePicture
	}
	
	// Boolean fields (always update)
	user.IsActive = req.IsActive
	user.IsVerified = req.IsVerified


	// Update in database
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// delete user
func (s *userService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}

