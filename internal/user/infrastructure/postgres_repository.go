package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/user/domain"
	"math"
	"strings"

	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

// get all users
func (r *postgresUserRepository) GetAll(params *domain.QueryParams) (*domain.QueryResult, error) {
	// set defaults
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}
	if params.SortBy == "" {
		params.SortBy = "id"
	}
	if params.Order == "" {
		params.Order = "DESC"
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

	// Build main query
	offset := (params.Page - 1) * params.PageSize
	// get all users query
	query := fmt.Sprintf(`%s %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
        r.getUserSelectQuery(), whereClause, params.SortBy, params.Order, len(args)+1, len(args)+2)

	
	args = append(args, params.PageSize, offset)

	// Execute query
	users := []*domain.User{}
	err = r.db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &domain.QueryResult{
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

	// Dynamic filters
	for key, value := range params.Filter {
		if value != "" {
			conditions = append(conditions, fmt.Sprintf("%s = $%d", key, argIndex))
			args = append(args, value)
			argIndex++
		}
	}

	if len(conditions) > 0 {
		return " WHERE " + strings.Join(conditions, " AND "), args
	}
	return "", args
}

// create user
func (r *postgresUserRepository) Create(user *domain.User) error {
	query := `
	INSERT INTO users (name, user_name, email, password, designation, bio, profile_picture, is_active, is_verified)
	VALUES (:name, :user_name, :email, :password, :designation, :bio, :profile_picture, :is_active, :is_verified)
	RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	}

	return nil
}

// login user
func (r *postgresUserRepository) LoginUser(email, password string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE email = $1 AND password = $2`

	var user domain.User
	err := r.db.Get(&user, query, email, password)
	return &user, err	
}

// get by email
func (r *postgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	query := r.getUserSelectQuery() + " WHERE email = $1"

	var user domain.User
	err := r.db.Get(&user, query, email)
	
	return &user, err
}

// get by id
func (r *postgresUserRepository) GetByID(id int) (*domain.User, error) {
	query := r.getUserSelectQuery() + " WHERE id = $1"

	var user domain.User
	err := r.db.Get(&user, query, id)
	return &user, err
}

// update user
func (r *postgresUserRepository) Update(user *domain.User) error {
	query := `
		UPDATE users SET 
			name = :name, user_name = :user_name, email = :email, 
			designation = :designation, bio = :bio, profile_picture = :profile_picture,
			is_active = :is_active, is_verified = :is_verified, updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at
		`

	_, err := r.db.NamedQuery(query, user)
	return err
}

// delete user
func (r *postgresUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
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
