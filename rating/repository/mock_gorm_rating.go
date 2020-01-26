package repository

import (
	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/rating"
)

// RatingRepository implements rating.Repository
type MockGormRatingRepository struct {
	conn *gorm.DB
}

// NewGormRatingRepository creates object of RatingRepository
func NewMockGormRatingRepository(conn *gorm.DB) rating.RatingRepository {
	return &MockGormRatingRepository{conn: conn}
}

// Rating returns healthcenters rating from database
func (gr *MockGormRatingRepository) Rating(id uint) (float64, []error) {
	rating := entity.MockRating
	return float64(rating.Value), nil
}

// StoreRating stores rating to database
func (gr *MockGormRatingRepository) StoreRating(rating *entity.Comment) (*entity.Comment, []error) {
	r := rating

	return r, nil
}
