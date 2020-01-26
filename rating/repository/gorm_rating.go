package repository

import (
	"fmt"

	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
)

// RatingRepository implements rating.Repository
type GormRatingRepository struct {
	conn *gorm.DB
}

// NewGormRatingRepository creates object of RatingRepository
func NewGormRatingRepository(conn *gorm.DB) *GormRatingRepository {
	return &GormRatingRepository{conn: conn}
}

// Rating returns healthcenters rating from database
func (gr *GormRatingRepository) Rating(id uint) (float64, []error) {
	rating := struct {
		Rating float64
	}{}
	errs := gr.conn.Raw("select avg(rating) as rating from comments where health_center_id = ? group by health_center_id", id).Scan(&rating).GetErrors()

	fmt.Println("rating:", rating)

	if len(errs) > 0 {
		fmt.Println(errs)
	}

	return rating.Rating, nil
}

// StoreRating stores rating to database
func (gr *GormRatingRepository) StoreRating(rating *entity.Comment) (*entity.Comment, []error) {
	r := rating
	errs := gr.conn.Create(rating).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return r, nil
}
