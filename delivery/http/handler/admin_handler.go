package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/admin"
	"github.com/TenaHub/api/entity"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
	adminService admin.AdminService
}
func NewAdminHandler(adm admin.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adm}
}

func (adm *AdminHandler) PutAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	adminData, errs := adm.adminService.AdminById(uint(id))
	//
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	json.Unmarshal(body, &adminData)
	adminData.ID = uint(id)
	adminData, errs = adm.adminService.UpdateAdmin(adminData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(adminData, "", "\t\t")

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



func (adm *AdminHandler) GetSingleAdmin(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	admin, errs := adm.adminService.AdminById(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(admin, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

type response struct{
	Status string
	Content interface{}
}

func (uh *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	admin := entity.Admin{Email: email, Password: password}
	user, _ := uh.adminService.Admin(&admin)
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
