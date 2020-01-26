package test_handler

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	ratingRepo "github.com/TenaHub/api/rating/repository"
	ratingServ "github.com/TenaHub/api/rating/service"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"reflect"
)

func TestRating(t *testing.T) {
	ratingRepo := ratingRepo.NewMockGormRatingRepository(nil)
	ratingServ := ratingServ.NewHcRatingService(ratingRepo)
	ratingHandler := handler.NewRatingHandler(ratingServ)

	mux := httprouter.New()
	mux.GET("/v1/rating/:id", ratingHandler.GetRating)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/rating/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	var mockRating entity.Rating
	var Rating  entity.Rating
	_ = json.Unmarshal(body, &Rating )

	if !reflect.DeepEqual(mockRating, Rating) {
		t.Errorf("want body to contain \n%q, but\n%q",mockRating, Rating)
	}

}



