package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"encoding/json"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/agent"
)

type AgentHandler struct {
	agentService agent.AgentService
}
func NewAgentHandler(adm agent.AgentService) *AgentHandler {
	return &AgentHandler{agentService: adm}
}


func (adm *AgentHandler) GetSingleAgent(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	agent, errs := adm.agentService.AgentById(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agent, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8282")

	w.Write(output)
	return
}
func (adm *AgentHandler) GetAgents(w http.ResponseWriter,r *http.Request, ps httprouter.Params) {

	agents, errs := adm.agentService.Agents()

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(output)
	return
}
func (adm *AgentHandler) DeleteAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.agentService.DeleteAgent(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusNoContent)
	return
}
func (adm *AgentHandler) PostAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	header.Add("Access-Control-Max-Age","86400")


	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	agentData := &entity.Agent{}

	err := json.Unmarshal(body, agentData)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_, errs := adm.agentService.StoreAgent(agentData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	//if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusCreated)
		//return
	//}

	return
}
func (adm *AgentHandler) PutAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	agentData, errs := adm.agentService.AgentById(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &agentData)
	agentData.ID = uint(id)
	agentData, errs = adm.agentService.UpdateAgent(agentData)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(agentData, "", "\t\t")

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
func (uh *AgentHandler) GetAgent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	agent := entity.Agent{Email: email, Password: password}
	user, _ := uh.agentService.Agent(&agent)

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


