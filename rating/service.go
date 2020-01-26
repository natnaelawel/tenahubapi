package rating

import (
	"github.com/TenaHub/api/entity"
)

// RatingService is
type RatingService interface{
	Rating(id uint) (float64, []error)
	StoreRating(rating *entity.Comment) (*entity.Comment, []error)
}
