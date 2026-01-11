package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/user/domain"
	"math"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

// create user
func (r *postgresUserRepository) Create(user *domain.CreateUserRequest) (*domain.User, error) {
	params := map[string]interface{}{
		"name":            user.Name,
		"user_name":       user.UserName,
		"email":           user.Email,
		"password":        user.Password,
		"designation":     user.Designation,
		"bio":             user.Bio,
		"profile_picture": user.ProfilePicture,
		"is_active":       user.IsActive,
		"is_verified":     user.IsVerified,
	}

	query := `
	INSERT INTO users (name, user_name, email, password, designation, bio, profile_picture, is_active, is_verified)
	VALUES (:name, :user_name, :email, :password, :designation, :bio, :profile_picture, :is_active, :is_verified)
	RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var createdAt, updatedAt time.Time
		err = rows.Scan(&id, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan returned values: %w", err)
		}
		
		// Create User struct with all fields
		result := &domain.User{
			ID:             id,
			Name:           user.Name,
			UserName:       user.UserName,
			Email:          user.Email,
			Password:       user.Password,
			Designation:    user.Designation,
			Bio:            user.Bio,
			ProfilePicture: user.ProfilePicture,
			IsActive:       user.IsActive,
			IsVerified:     user.IsVerified,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		}
		return result, nil
	}


	return nil, fmt.Errorf("no rows returned after insert")
}

// update user
func (r *postgresUserRepository) Update(id int, user *domain.UpdateUserRequest) (*domain.User, error) {
	params := map[string]interface{}{
		"id":              id,
		"name":            user.Name,
		"user_name":       user.UserName,
		"email":           user.Email,
		"designation":     user.Designation,
		"bio":             user.Bio,
		"profile_picture": user.ProfilePicture,
		"is_active":       user.IsActive,
		"is_verified":     user.IsVerified,
	}

	query := `
		UPDATE users 
		SET name = :name, user_name = :user_name, email = :email, 
		    designation = :designation, bio = :bio, profile_picture = :profile_picture,
		    is_active = :is_active, is_verified = :is_verified, updated_at = NOW()
		WHERE id = :id
		RETURNING created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, params)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var UpdatedAt time.Time
		var createdAt, updatedAt time.Time
		err = rows.Scan(&createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		result := &domain.User{
			ID:             id,
			Name:           user.Name,
			UserName:       user.UserName,
			Email:          user.Email,
			Designation:    user.Designation,
			Bio:            user.Bio,
			ProfilePicture: user.ProfilePicture,
			IsActive:       user.IsActive,
			IsVerified:     user.IsVerified,
			UpdatedAt:      UpdatedAt,
		}	

		return result, nil
	}

	return nil, fmt.Errorf("user with id %d not found", id) 
}

// delete user
func (r *postgresUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// get by email
func (r *postgresUserRepository) GetByEmail(email string) (*domain.User, error) {
    query := r.getUserSelectQuery() + " WHERE email = $1"

    var user domain.User
    err := r.db.Get(&user, query, email)
    if err != nil {
        return nil, err // ✅ Return nil on error
    }
    
    return &user, nil
}

// get by id
func (r *postgresUserRepository) GetByID(id int) (*domain.User, error) {
	query := r.getUserSelectQuery() + " WHERE id = $1"

	var user domain.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// get all users
func (r *postgresUserRepository) GetAll(params *domain.QueryParams) (*domain.QueryResponse, error) {
    // Validate sort parameters
    validSortFields := map[string]bool{
        "id": true, "name": true, "email": true, "created_at": true, "updated_at": true,
    }
    validOrders := map[string]bool{
        "ASC": true, "DESC": true,
    }
    
    sortBy := "id" // default
    if validSortFields[params.SortBy] {
        sortBy = params.SortBy
    }
    
    order := "DESC" // default
    if validOrders[params.Order] {
        order = params.Order
    }

    // Build WHERE clause
    whereClause, args := r.buildWhereClause(params)
    
    // Count total
    countQuery := "SELECT COUNT(*) FROM users" + whereClause
    var total int64
    err := r.db.Get(&total, countQuery, args...)
    if err != nil {
        return nil, err
    }

    // Build main query - NOW SAFE
    offset := (params.Page - 1) * params.PageSize
    query := fmt.Sprintf(`%s %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
        r.getUserSelectQuery(), whereClause, sortBy, order, len(args)+1, len(args)+2)
    
    args = append(args, params.PageSize, offset)

    // Execute query
    users := []*domain.User{}
		
    err = r.db.Select(&users, query, args...)
    if err != nil {
        return nil, err
    }

    // Calculate total pages
    totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

    return &domain.QueryResponse{
        Data:       users,
        TotalItem:  total,
        Page:       params.Page,
        PageSize:   params.PageSize,
        TotalPages: totalPages,
    }, nil
}


// build where clause
func (r *postgresUserRepository) buildWhereClause(params *domain.QueryParams) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	// Search functionality
	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR email ILIKE $%d OR user_name ILIKE $%d)", argIndex, argIndex, argIndex))
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		return " WHERE " + strings.Join(conditions, " AND "), args
	}
	return "", args
}

// Helper function for consistent SELECT query (add this after NewPostgresUserRepository)
func (r *postgresUserRepository) getUserSelectQuery() string {
    return `
    SELECT 
        id, name, 
        COALESCE(user_name, '') as user_name,
        email, password,
        COALESCE(designation, '') as designation,
        COALESCE(bio, '') as bio,
        COALESCE(profile_picture, '') as profile_picture,
        is_active, is_verified,
        COALESCE(verification_token, '') as verification_token,
        COALESCE(reset_password_token, '') as reset_password_token,
        reset_password_expires, last_login,
        created_at, updated_at
    FROM users`
}
