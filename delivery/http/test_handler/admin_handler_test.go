package test_handler


import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	adminRepo "github.com/natnaelawel/tenahubapi/api/admin/repository"
	adminServ "github.com/natnaelawel/tenahubapi/api/admin/service"
	"github.com/natnaelawel/tenahubapi/api/delivery/http/handler"
	"github.com/natnaelawel/tenahubapi/api/entity"
	"reflect"
)


func TestAdmin(t *testing.T) {

	adminRepo := adminRepo.NewMockAdminGormRepo(nil)
	adminServ := adminServ.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminServ)

	mux := httprouter.New()
	mux.GET("/v1/admin", adminHandler.GetAdmin)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin")
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
	var mockAdmin entity.Admin
	var admin entity.Admin
	_ = json.Unmarshal(body, &admin )

	if !reflect.DeepEqual(mockAdmin, admin) {
		t.Errorf("want body to contain \n%q, but\n%q",mockAdmin, admin)
	}
}


