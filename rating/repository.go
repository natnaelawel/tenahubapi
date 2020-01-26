package rating

import (
	"github.com/natnaelawel/tenahubapi/api/entity"
)

// RatingRepository is
type RatingRepository interface{
	Rating(id uint) (float64, []error)
	StoreRating(rating *entity.Comment) (*entity.Comment, []error)
}
