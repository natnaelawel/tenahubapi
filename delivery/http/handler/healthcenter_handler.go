package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/healthcenter"
	"github.com/TenaHub/api/entity"
	"fmt"
)

type HealthCenterHandler struct {
	healthCenterService healthcenter.HealthCenterService
}
func NewHealthCenterHandler(adm healthcenter.HealthCenterService) *HealthCenterHandler {
	return &HealthCenterHandler{healthCenterService: adm}
}

func (adm *HealthCenterHandler) GetSingleHealthCenter(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	healthcenter, errs := adm.healthCenterService.HealthCenterById(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(healthcenter, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (uh *HealthCenterHandler) GetHealthCenter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	healthcenter := entity.HealthCenter{Email: email, Password: password}
	user, _ := uh.healthCenterService.HealthCenter(&healthcenter)

	if user == nil {
		data, err := json.MarshalIndent(&response{Status:"error", Content:nil},"", "\t")
		if err != nil {

		}
		http.Error(w, string(data) , http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(response{Status:"success", Content:&user}, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Write(output)
	return
}


func (adm *HealthCenterHandler) GetHealthCenters(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {

	agents, errs := adm.healthCenterService.HealthCenters()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agents, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (adm *HealthCenterHandler) DeleteHealthCenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.healthCenterService.DeleteHealthCenter(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (adm *HealthCenterHandler) PutHealthCenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	healthCenterData, errs := adm.healthCenterService.HealthCenterById(uint(id))
	//
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	json.Unmarshal(body, &healthCenterData)
	healthCenterData.ID = uint(id)
	healthCenterData, errs = adm.healthCenterService.UpdateHealthCenter(healthCenterData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(healthCenterData, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(output)
	return
}

func (adm *HealthCenterHandler) PostHealthCenter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	header.Add("Access-Control-Max-Age","86400")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	healthcenterData := &entity.HealthCenter{}

	err := json.Unmarshal(body, healthcenterData)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.healthCenterService.StoreHealthCenter(healthcenterData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

// GetHealthcenters handles GET /healthcenters
func (hh *HealthCenterHandler) SearchHealthcenters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	searchKey := r.URL.Query().Get("search-key")
	column := r.URL.Query().Get("column")
	fmt.Println(searchKey)

	hcs, errs := hh.healthCenterService.SearchHealthCenters(searchKey, column)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&hcs, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// GetTop handles Get /healthcenters/top/:amount
func (hh *HealthCenterHandler) GetTop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	amount, err := strconv.Atoi(ps.ByName("amount"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	results, errs := hh.healthCenterService.Top(uint(amount))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(results, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

func (adm *HealthCenterHandler) GetHealthCentersByAgentId(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	agents, errs := adm.healthCenterService.HealthCenterByAgentId(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agents, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}