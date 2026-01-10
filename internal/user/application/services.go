package application

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"gin-quickstart/config"
	authDomain "gin-quickstart/internal/auth/domain"
	"gin-quickstart/internal/user/domain"
	"gin-quickstart/pkg/auth"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userServices struct {
	userRepo domain.UserRepository
	refreshTokenRepo authDomain.RefreshTokenRepository
	jwtCnf   *config.JWTConfig
}

func NewUserServices(userRepo domain.UserRepository, refreshTokenRepo authDomain.RefreshTokenRepository, jwtCnf *config.JWTConfig) *userServices {
    return &userServices{
        userRepo: userRepo,
				refreshTokenRepo: refreshTokenRepo,
				jwtCnf: jwtCnf,
    }
}

// get all users
func (s *userServices) GetAllUsers(params *domain.QueryParams) (*domain.QueryResult, error) {
	users, err := s.userRepo.GetAll(params)
	
	if err != nil {
		return nil, err
	}

	return users, nil
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

	// Generate JWT token pair
	jwtService := auth.NewJWTServices(s.jwtCnf.SecretKey, s.jwtCnf.AccessTokenExpireMinutes, s.jwtCnf.RefreshTokenExpireDays)
	tokenPair, err := jwtService.GenerateTokenPair(user.ID, user.Name, user.Email, user.UserName)
	if err != nil {
		return nil, err
	}

	// Store refresh token in database
	tokenHash := sha256.Sum256([]byte(tokenPair.RefreshToken))
	refreshToken := &authDomain.RefreshToken{
		UserID:    user.ID,
		TokenHash: hex.EncodeToString(tokenHash[:]),
		ExpiresAt: time.Now().Add(time.Duration(s.jwtCnf.RefreshTokenExpireDays) * 24 * time.Hour),
	}
	
	err = s.refreshTokenRepo.Create(refreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: tokenPair,
		User:  user,
	}, nil
}

// get user by email
func (s *userServices) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
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

	req.Password = string(hashedPassword)

	// Create user
	createdUser, err := s.userRepo.Create(req)
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

// helper function 
func updateUserFields(user *domain.User, req *domain.UpdateUserRequest) *domain.UpdateUserRequest {
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

	return req
}