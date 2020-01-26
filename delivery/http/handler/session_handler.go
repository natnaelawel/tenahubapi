package handler

import (
	"github.com/TenaHub/api/session"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/TenaHub/api/entity"
	"fmt"
)

type SessionHandler struct {
	sessionService session.SessionService
}

func NewSessionHandler(sserv session.SessionService)*SessionHandler{
	return &SessionHandler{sessionService:sserv}
}

// GetSession handles GET on /v1/session
func (sh *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-type", "application/json")

	uuid := r.URL.Query().Get("uuid")

	session, errs := sh.sessionService.Session(uuid)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&session, "", "\n")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// PostSession handles POST on /v1/session
func (sh *SessionHandler) PostSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-type", "application/json")

	l := r.ContentLength
	data := make([]byte, l)
	r.Body.Read(data)

	session := entity.Session{}

	err := json.Unmarshal(data, &session)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	ses, errs := sh.sessionService.StoreSession(&session)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/session/?uuid=%s", ses.UUID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// DeleteSession handles DELETE v1/session/:uuid
func (sh *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-type", "application/json")

	uuid := ps.ByName("uuid")

	session, errs := sh.sessionService.DeleteSession(uuid)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err := json.MarshalIndent(session, "", "\n")


	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}