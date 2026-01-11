package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/category/domain"
	"math"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

// get all categories with search, sorting, pagination
func (r *categoryRepository) GetAll(params *domain.QueryParams) (*domain.QueryResponse, error) {
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
			SELECT id, name, slug, description, image_url, sort_order, is_active, meta_title, meta_description, created_at, updated_at
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

	return &domain.QueryResponse{
		Data:       categories,
		TotalItem:  total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

// get by id
func (r *categoryRepository) GetByID(id int) (*domain.Category, error) {
	query := `
		SELECT id, name, slug, description, image_url, sort_order, is_active, meta_title, meta_description, created_at, updated_at
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

// create
func (r *categoryRepository) Create(category *domain.CreateCategoryRequest) (*domain.Category, error) {
	query := `
		INSERT INTO categories 
		(name, slug, description, image_url, sort_order, is_active, meta_title, meta_description) 
		VALUES 
		(:name, :slug, :description, :image_url, :sort_order, :is_active, :meta_title, :meta_description)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, category)

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
		
		result := &domain.Category{
			ID: id,
			Name: category.Name,
			Slug: category.Slug,
			Description: category.Description,
			ImageUrl: category.ImageUrl,
			SortOrder: category.SortOrder,
			IsActive: category.IsActive,
			MetaTitle: category.MetaTitle,
			MetaDescription: category.MetaDescription,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		return result, nil
	}

	return nil, fmt.Errorf("no rows returned after insert")
}

// update
func (r *categoryRepository) Update(id int, category *domain.UpdateCategoryRequest) (*domain.Category, error) {
	category.ID = id

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
		var updatedAt time.Time

		err = rows.Scan(&updatedAt)
		if err != nil {
				return nil, err
		}

		result := &domain.Category{
			ID: id,
			Name: category.Name,
			Slug: category.Slug,
			Description: category.Description,
			ImageUrl: category.ImageUrl,
			SortOrder: category.SortOrder,
			IsActive: category.IsActive,
			MetaTitle: category.MetaTitle,
			MetaDescription: category.MetaDescription,
			UpdatedAt: updatedAt,
		}

		return result, nil
	}
	
	return nil, fmt.Errorf("category with id %d not found", id)
}

// delete
func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
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

