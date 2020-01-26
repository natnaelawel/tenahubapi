package test_handler
//
import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	serviceRepo "github.com/TenaHub/api/hcservice/repository"
	serviceServ "github.com/TenaHub/api/hcservice/service"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"reflect"
)

func TestServices(t *testing.T) {

	serviceRepo := serviceRepo.NewMockServiceGormRepo(nil)
	serviceServ := serviceServ.NewServiceService(serviceRepo)
	serviceHandler := handler.NewServiceHandler(serviceServ)

	mux := httprouter.New()
	mux.GET("/v1/services/:id", serviceHandler.GetServices)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/services/1")
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
	var mockService []entity.Service
	var Service  []entity.Service
	_ = json.Unmarshal(body, &Service )
	mockService = append(mockService, entity.MockService, entity.MockService)

	if !reflect.DeepEqual(mockService, Service) {
		t.Errorf("want body to contain \n%q, but\n%q",mockService, Service)
	}

}

func TestPendingService(t *testing.T) {

	serviceRepo := serviceRepo.NewMockServiceGormRepo(nil)
	serviceServ := serviceServ.NewServiceService(serviceRepo)
	serviceHandler := handler.NewServiceHandler(serviceServ)

	mux := httprouter.New()
	mux.GET("/v1/pendingservice", serviceHandler.GetPendingServices)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/pendingservice")
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
	var mockService []entity.Service
	var Service  []entity.Service
	_ = json.Unmarshal(body, &Service )
	mockService = append(mockService, entity.MockService, entity.MockService)

	if !reflect.DeepEqual(mockService, Service) {
		t.Errorf("want body to contain \n%q, but\n%q",mockService, Service)
	}
}


