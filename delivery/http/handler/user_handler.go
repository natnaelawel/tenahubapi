package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/user"
	"github.com/julienschmidt/httprouter"
)

// type response struct{
// 	Status string
// 	Content interface{}
// }

// UserHandler handles User related http requests
type UserHandler struct {
	userService user.UserService
}

// NewUserHander creates and returns new UserHandler object
func NewUserHander(us user.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

// GetUsers handles GET /v1/users request
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	users, errs := uh.userService.Users()

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// GetUser handles POST /v1/user
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Println(email,password)

	usr := entity.User{Email: email, Password: password}

	user, errs := uh.userService.User(&usr)
	fmt.Println(errs)
	if len(errs) > 0 {
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

// GetSingleUser handler GET /v1/users/:id
func (uh *UserHandler) GetSingleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.UserByID(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// PutUser handles PUT /v1/users/:id request
func (uh *UserHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.UserByID(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	err = json.Unmarshal(body, user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs = uh.userService.UpdateUser(user)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&user, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// DeleteUser Handler DELETE /v1/users/:id
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.DeleteUser(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = json.MarshalIndent(user, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}

// PostUser handles POST /v1/users requests
func (uh *UserHandler) PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	user := entity.User{}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	err := json.Unmarshal(body, &user)
	fmt.Printf("unmarshaling: %s",err)

	if err != nil {
		fmt.Println("here")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	u, errs := uh.userService.StoreUser(&user)
	fmt.Printf("storing: %s",errs)
	if len(errs) > 0 {
		fmt.Println("here")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println("everything okay")
	p := fmt.Sprintf("/v1/users/%d", u.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
