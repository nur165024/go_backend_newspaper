package infrastructure

import (
	"fmt"
	"gin-quickstart/internal/advertisement/domain"
	"math"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type adsRepository struct {
	db *sqlx.DB
}

func NewAdsRepository(db *sqlx.DB) *adsRepository {
	return &adsRepository{db:db}
}

// get all ads
func (r *adsRepository) GetAll(params *domain.QueryParams) (*domain.QueryResponse, error) {
	// Validate sort parameters
	validSortFields := map[string]bool{
		"id": true, "title": true, "position": true, "created_at": true, "updated_at": true,
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
	countQuery := "SELECT COUNT(*) FROM advertisements" + whereClause
	var total int64
	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
			return nil, err
	}

	// Build main query with pagination - NOW SAFE
	offset := (params.Page - 1) * params.PageSize
	query := fmt.Sprintf(`
			SELECT id, title, image_url, position, is_active, start_date, end_date, priority, click_count, created_at, updated_at
			FROM advertisements %s 
			ORDER BY %s %s 
			LIMIT $%d OFFSET $%d`,
			whereClause, sortBy, order, len(args)+1, len(args)+2)
	
	args = append(args, params.PageSize, offset)

	// Execute query
	var ads []*domain.Advertisement

	err = r.db.Select(&ads, query, args...)
	if err != nil {
			return nil, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &domain.QueryResponse{
		Data:       ads,
		TotalItem:  total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

// get by id
func (r *adsRepository) GetByID(id int) (*domain.Advertisement, error) {
	query := `
		SELECT id, title ,image_url ,link ,position ,is_active ,start_date ,end_date ,priority ,click_count ,impression_count ,target_audience ,created_by ,created_at ,updated_at
		FROM advertisements
		WHERE id = $1
	`

	var ads domain.Advertisement
	err := r.db.Get(&ads, query, id)
 
	if err != nil {
		return nil, err
	}

	return &ads, nil
}

// create
func (r *adsRepository) Create(ads *domain.CreateAdsRequest) (*domain.Advertisement, error) {
	query := `
		INSERT INTO advertisements 
		(title, image_url, link, position, is_active, priority, target_audience, created_by, start_date, end_date) 
		VALUES 
		(:title, :image_url, :link, :position, :is_active, :priority, :target_audience, :created_by, :start_date, :end_date)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, ads)

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
		
		result := &domain.Advertisement{
			ID: id,
			Title: ads.Title,
			ImageUrl: ads.ImageUrl,
			Link: ads.Link,
			Position: ads.Position,
			IsActive: ads.IsActive,
			Priority: ads.Priority,
			TargetAudience: ads.TargetAudience,
			CreatedBy: ads.CreatedBy,
			StartDate: ads.StartDate,
			EndDate: ads.EndDate,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		return result, nil
	}

	return nil, fmt.Errorf("no rows returned after insert")
}

// update
func (r *adsRepository) Update(id int, ads *domain.UpdateAdsRequest) (*domain.Advertisement, error) {
	ads.ID = id

	query := `
		UPDATE categories 
		SET 
		title = :title, 
		image_url = :image_url, 
		link = :link,
		position = :position, 
		is_active = :is_active, 
		priority = :priority, 
		target_audience = :target_audience, 
		created_by = :created_by, 
		start_date = :start_date, 
		end_date = :end_date,
		updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at
	`

	rows, err := r.db.NamedQuery(query, ads)

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

		result := &domain.Advertisement{
			ID: id,
			Title: ads.Title,
			ImageUrl: ads.ImageUrl,
			Link: ads.Link,
			Position: ads.Position,
			IsActive: ads.IsActive,
			Priority: ads.Priority,
			TargetAudience: ads.TargetAudience,
			CreatedBy: ads.CreatedBy,
			StartDate: ads.StartDate,
			EndDate: ads.EndDate,
			UpdatedAt: updatedAt,
		}

		return result, nil
	}
	
	return nil, fmt.Errorf("Advertisement with id %d not found", id)
}

// delete
func (r *adsRepository) Delete(id int) error {
	query := `DELETE FROM advertisements WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Build WHERE clause for search
func (r *adsRepository) buildWhereClause(params *domain.QueryParams) (string, []interface{}) {
    var conditions []string
    var args []interface{}
    argIndex := 1

    // Search functionality (title, position only - not dates)
    if params.Search != "" {
        conditions = append(conditions, fmt.Sprintf("(title ILIKE $%d OR position ILIKE $%d)", argIndex, argIndex))
        args = append(args, "%"+params.Search+"%")
        argIndex++
    }

    if len(conditions) > 0 {
        return " WHERE " + strings.Join(conditions, " AND "), args
    }
    return "", args
}
