package handler

import (
	"github.com/TenaHub/api/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TenaHub/api/rating"
	"github.com/julienschmidt/httprouter"
)

// RatingHandler handles rating related http requests
type RatingHandler struct {
	ratingService rating.RatingService
}

// NewRatingHandler creates RatingHandler object
func NewRatingHandler(rserv rating.RatingService) *RatingHandler {
	return &RatingHandler{ratingService: rserv}
}

// GetRating handles GET /v1/rating/:id
func (rh *RatingHandler) GetRating(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Contenty-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	rating, errs := rh.ratingService.Rating(uint(id))

	fmt.Println(errs)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	ra := struct {
		Rating float64
	}{
		Rating: rating,
	}

	output, err := json.MarshalIndent(ra, "", "\n")

	fmt.Println(err)


	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}


// PostRating handles POST /v1/rating
func (rh *RatingHandler) PostRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	rating := entity.Comment{}

	l := r.ContentLength
	data := make([]byte, l)

	r.Body.Read(data)

	err := json.Unmarshal(data, &rating)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println(rating)

	rat, errs := rh.ratingService.StoreRating(&rating)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	
	fmt.Println(rat)

	output, err := json.MarshalIndent(rat, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}