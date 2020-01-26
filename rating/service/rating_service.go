package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/rating"
)

// HcRatingService implements rating.RatingService
type HcRatingService struct{
	ratingRepo rating.RatingRepository
}

// NewHcRatingService creates object of HcRatingService
func NewHcRatingService(rRepo rating.RatingRepository) *HcRatingService {
	return &HcRatingService{ratingRepo: rRepo}
}

// Rating returns healthcenters rating
func (hs *HcRatingService) Rating(id uint) (float64, []error) {
	rating, errs := hs.ratingRepo.Rating(id)

	if len(errs) > 0 {
		return 0.0, errs
	}

	return rating, nil
}

// StoreRating stores health center rating
func (hs *HcRatingService) StoreRating(rating *entity.Comment) (*entity.Comment, []error) {
	r, errs := hs.ratingRepo.StoreRating(rating)
	if len(errs) > 0{
		return nil, errs
	}
	return r, nil
}