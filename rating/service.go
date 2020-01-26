package rating

import (
	"github.com/natnaelawel/tenahubapi/api/entity"
)

// RatingService is
type RatingService interface{
	Rating(id uint) (float64, []error)
	StoreRating(rating *entity.Comment) (*entity.Comment, []error)
}
