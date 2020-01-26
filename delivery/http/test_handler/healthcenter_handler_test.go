package test_handler

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	healthCenterRepo "github.com/TenaHub/api/healthcenter/repository"
	healthCenterServ "github.com/TenaHub/api/healthcenter/service"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"reflect"
)

func TestHealthCenters(t *testing.T) {

	hcRepo := healthCenterRepo.NewMockHealthCenterGormRepo(nil)
	hcServ := healthCenterServ.NewHealthCenterService(hcRepo)
	hcHandler := handler.NewHealthCenterHandler(hcServ)

	mux := httprouter.New()
	mux.GET("/v1/healthcenters", hcHandler.GetHealthCenters)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/healthcenters")
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
	var mockHealthcenter []entity.HealthCenter
	var healthcenter []entity.HealthCenter
	_ = json.Unmarshal(body, &healthcenter)
	mockHealthcenter = append(mockHealthcenter, entity.MockHealthCenter, entity.MockHealthCenter)

	if !reflect.DeepEqual(mockHealthcenter, healthcenter) {
		t.Errorf("want body to contain \n%q, but\n%q", mockHealthcenter, healthcenter)

	}
}

func TestHealthCenter(t *testing.T) {
	hcRepo := healthCenterRepo.NewMockHealthCenterGormRepo(nil)
	hcServ := healthCenterServ.NewHealthCenterService(hcRepo)
	hcHandler := handler.NewHealthCenterHandler(hcServ)

	mux := httprouter.New()
	mux.GET("/v1/healthcenters/:id", hcHandler.GetHealthCenters)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/healthcenters/1")
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
	var mockHealthcenter entity.HealthCenter
	var healthcenter entity.HealthCenter
	_ = json.Unmarshal(body, &healthcenter)
	mockHealthcenter = entity.MockHealthCenter
	if !reflect.DeepEqual(mockHealthcenter, healthcenter) {
		t.Errorf("want body to contain \n%q, but\n%q", mockHealthcenter, healthcenter)

	}
}