package application

import (
	"errors"
	"fmt"
	"gin-quickstart/config"
	"gin-quickstart/internal/user/domain"
	"gin-quickstart/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type userServices struct {
	userRepo domain.UserRepository
	jwtCnf   *config.JWTConfig
}

func NewUserServices(userRepo domain.UserRepository, jwtCnf *config.JWTConfig) *userServices {
    return &userServices{
        userRepo: userRepo,
				jwtCnf: jwtCnf,
    }
}

// create user
func (s *userServices) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("this email already exists!")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:           req.Name,
		UserName:       req.UserName,
		Email:          req.Email,
		Password:       string(hashedPassword),
		Designation:    req.Designation,
		Bio:            req.Bio,
		ProfilePicture: req.ProfilePicture,
		IsActive:       true,  // Set default
		IsVerified:     false, // Set default
	}

	// Create user
	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser, nil
}

// update user
func (s *userServices) UpdateUser(id int, req *domain.UpdateUserRequest) (*domain.User, error) {
	// Get existing user
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	updateInfo := updateUserFields(user, req)

	// Update in database
	updatedUser, err := s.userRepo.Update(id, updateInfo)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// delete user
func (s *userServices) DeleteUser(id int) error {
	err := s.userRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// get user by id
func (s *userServices) GetUserByID(id int) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// login user
// login user
func (s *userServices) LoginUser(email, password string) (*domain.LoginResponse, error) {
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

	// Generate JWT token
	token, err := auth.NewJWTServices(s.jwtCnf.SecretKey, s.jwtCnf.AccessTokenExpireMinutes, s.jwtCnf.RefreshTokenExpireDays).GenerateTokenPair(user.ID, user.Name, user.Email, user.UserName)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}


// get all users
func (s *userServices) GetAllUsers(params *domain.QueryParams) (*domain.QueryResult, error) {
	users, err := s.userRepo.GetAll(params)
	
	if err != nil {
		return nil, err
	}

	return users, nil
}

// get user by email
func (s *userServices) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// helper function 
func updateUserFields(user *domain.User, req *domain.UpdateUserRequest) *domain.User {
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
	user.IsActive = true
	user.IsVerified = true

	return user
}