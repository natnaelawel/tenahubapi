package test_handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	userRepo "github.com/TenaHub/api/user/repository"
	userServ "github.com/TenaHub/api/user/service"
	"github.com/julienschmidt/httprouter"
	"reflect"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"encoding/json"
)

func TestUsers(t *testing.T) {

	userRepo := userRepo.NewMockUserGormRepo(nil)
	userServ := userServ.NewUserService(userRepo)
	userHandler := handler.NewUserHander(userServ)

	mux := httprouter.New()
	mux.GET("/v1/users", userHandler.GetUsers)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/users")
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
	var mockUsers []entity.User
	var users []entity.User
	_ = json.Unmarshal(body, &users)
	mockUsers = append(mockUsers, entity.MockUser, entity.MockUser)

	if !reflect.DeepEqual(mockUsers, users) {
		t.Errorf("want body to contain \n%q, but\n%q",mockUsers, users)
	}

}

func TestUser(t *testing.T) {

	userRepo := userRepo.NewMockUserGormRepo(nil)
	userServ := userServ.NewUserService(userRepo)
	userHandler :=handler.NewUserHander(userServ)

	mux := httprouter.New()
	mux.GET("/v1/users/:id", userHandler.GetSingleUser)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/users/1")
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
	var mockUsers entity.User
	var users entity.User
	_ = json.Unmarshal(body, &users)
	mockUsers = entity.MockUser
	if !reflect.DeepEqual(mockUsers, users) {
		t.Errorf("want body to contain \n%q, but\n%q",mockUsers, users)
	}

}

