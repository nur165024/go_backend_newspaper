package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/settings/domain"
	"math"

	"github.com/jmoiron/sqlx"
)

type settingsRepository struct {
	db *sqlx.DB
}

func NewSettingsRepository(db *sqlx.DB) *settingsRepository {
	return &settingsRepository{db: db}
}

// get all
func (r *settingsRepository) GetAll(params *domain.QueryParams) (*domain.QueryResult, error) {
	// validation sort parameters
	validSortFields := map[string]bool{"id": true, "key": true, "value": true, "category": true, "data_type": true}
	// validation sort parameters
	validOrders := map[string]bool{"ASC": true, "DESC": true}

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
	countQuery := `SELECT COUNT(*) FROM settings WHERE 1=1 ` + whereClause

	var total int64
	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
			return nil, err
	}

	// Build main query with pagination - NOW SAFE
	offset := (params.Page - 1) * params.PageSize
	query := fmt.Sprintf(`
		SELECT id, key, value, category, data_type
			FROM settings
			WHERE 1=1 %s
			ORDER BY %s %s
			LIMIT %d OFFSET %d`, 
		whereClause, sortBy, order, len(args)+1, len(args)+2)

	args = append(args, params.PageSize, offset)

	// Execute query
	var settings []*domain.Settings

	err = r.db.Select(&settings, query, args...)
	if err != nil {
			return nil, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &domain.QueryResult{
		Data:       settings,
		TotalItem:  total,
		TotalPages: totalPages,
		Page:       params.Page,
		PageSize:   params.PageSize,
	}, nil
}

// get by id
func (r *settingsRepository) GetByID(id int) (*domain.Settings, error) {
	var setting domain.Settings
	err := r.db.Get(&setting, "SELECT id, key, value, category, data_type, category, is_public, is_editable FROM settings WHERE id=$1", id)
	
	if err != nil {
		return nil, err
	}

	return &setting, nil
}

// create
func (r *settingsRepository) Create(settings *domain.Settings) (*domain.Settings, error) {
	query := `
		INSERT INTO settings (key, value, category, data_type, is_public, is_editable)
		VALUES 
		(:key, :value, :category, :data_type, :is_public, :is_editable)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, settings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		// Scan RETURNING values
		err = rows.Scan(&settings.ID, &settings.CreatedAt, &settings.UpdatedAt)
		if err != nil {
				return nil, fmt.Errorf("failed to scan returned values: %w", err)
		}
		return settings, nil
	}

	return nil, fmt.Errorf("no rows returned after insert")
}

// update
func (r *settingsRepository) Update(id int, settings *domain.Settings) (*domain.Settings, error) {
	query := `
		UPDATE settings
		SET 
		key=:key, 
		value=:value, 
		category=:category, 
		data_type=:data_type, 
		is_public=:is_public, 
		is_editable=:is_editable,
		updated_at=NOW()
		WHERE id=:id
		RETURNING updated_at
	`
	
	rows, err := r.db.NamedQuery(query, settings)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&settings.UpdatedAt)
		if err != nil {
				return nil, err
		}
		return settings, nil
	}
	
	return nil, fmt.Errorf("Settings with id %d not found", id)
}

// delete
func (r *settingsRepository) Delete(id int) error {
	query := `DELETE FROM settings WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Build where clause for search
func (r *settingsRepository) buildWhereClause(params *domain.QueryParams) (string, []interface{}) {
	whereClause := ""
	args := []interface{}{}

	if params.Search != "" {
		whereClause += " AND (key ILIKE $1 OR value ILIKE $2 OR category ILIKE $3 OR data_type ILIKE $4)"
		args = append(args, "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%" , "%"+params.Search+"%")
	}

	return whereClause, args
}
