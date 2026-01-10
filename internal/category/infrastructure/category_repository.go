package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/category/domain"
	"math"
	"strings"

	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

// create
func (r *categoryRepository) Create(category *domain.Category) (*domain.Category, error) {
	query := `
		INSERT INTO categories 
		(name, slug, description, image_url, sort_order, is_active, meta_title, meta_description, meta_keywords) 
		VALUES 
		(:name, :slug, :description, :image_url, :sort_order, :is_active, :meta_title, :meta_description, :meta_keywords)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, category)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		// Scan RETURNING values
		err = rows.Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
				return nil, fmt.Errorf("failed to scan returned values: %w", err)
		}
		return category, nil
	}

	return nil, fmt.Errorf("no rows returned after insert")
}

// update
func (r *categoryRepository) Update(id int, category *domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories 
		SET 
		name = :name, 
		slug = :slug, 
		description = :description, 
		image_url = :image_url, 
		sort_order = :sort_order, 
		is_active = :is_active,
		meta_title = :meta_title, 
		meta_description = :meta_description, 
		meta_keywords = :meta_keywords,
		updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at
	`

	rows, err := r.db.NamedQuery(query, category)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&category.UpdatedAt)
		if err != nil {
				return nil, err
		}
		return category, nil
	}
	
	return nil, fmt.Errorf("category with id %d not found", id)
}

// delete
func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// get by id
func (r *categoryRepository) GetByID(id int) (*domain.Category, error) {
	query := `
		SELECT id, name, slug, description, image_url, sort_order, is_active, meta_title, meta_description, meta_keywords, created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	var category domain.Category
	err := r.db.Get(&category, query, id)
 
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// get all categories with search, sorting, pagination
func (r *categoryRepository) GetAll(params *domain.QueryParams) (*domain.QueryResult, error) {
	// Validate sort parameters
	validSortFields := map[string]bool{
		"id": true, "name": true, "slug": true, "created_at": true, "updated_at": true, "sort_order": true,
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
	
	// Count total records
	countQuery := "SELECT COUNT(*) FROM categories" + whereClause
	var total int64
	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
			return nil, err
	}

	// Build main query with pagination - NOW SAFE
	offset := (params.Page - 1) * params.PageSize
	query := fmt.Sprintf(`
			SELECT id, name, slug, description, image_url, sort_order, is_active, meta_title, meta_description, meta_keywords, created_at, updated_at
			FROM categories %s 
			ORDER BY %s %s 
			LIMIT $%d OFFSET $%d`,
			whereClause, sortBy, order, len(args)+1, len(args)+2)
	
	args = append(args, params.PageSize, offset)

	// Execute query
	var categories []*domain.Category

	err = r.db.Select(&categories, query, args...)
	if err != nil {
			return nil, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &domain.QueryResult{
		Data:       categories,
		TotalItem:  total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}


// Build WHERE clause for search
func (r *categoryRepository) buildWhereClause(params *domain.QueryParams) (string, []interface{}) {
    var conditions []string
    var args []interface{}
    argIndex := 1

    // Search functionality (name, slug, description)
    if params.Search != "" {
			conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR slug ILIKE $%d OR description ILIKE $%d)", argIndex, argIndex, argIndex))
			args = append(args, "%"+params.Search+"%")
			argIndex++
    }

    if len(conditions) > 0 {
        return " WHERE " + strings.Join(conditions, " AND "), args
    }
    return "", args
}

