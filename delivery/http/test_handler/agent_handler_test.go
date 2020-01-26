package test_handler

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	agentRepo "github.com/TenaHub/api/agent/repository"
	agentServ "github.com/TenaHub/api/agent/service"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"reflect"
)

func TestAgents(t *testing.T) {

	agentRepo := agentRepo.NewMockAgentGormRepo(nil)
	agentServ := agentServ.NewAgentService(agentRepo)
	agentHandler := handler.NewAgentHandler(agentServ)

	mux := httprouter.New()
	mux.GET("/v1/agent", agentHandler.GetAgents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/agent")
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
	var mockAgent []entity.Agent
	var agent []entity.Agent
	_ = json.Unmarshal(body, &agent)
	mockAgent = append(mockAgent, entity.MockAgent, entity.MockAgent)

	if !reflect.DeepEqual(mockAgent, agent) {
		t.Errorf("want body to contain \n%q, but\n%q", mockAgent, agent)

	}
}

func TestAgent(t *testing.T) {
	agentRepo := agentRepo.NewMockAgentGormRepo(nil)
	agentServ := agentServ.NewAgentService(agentRepo)
	agentHandler := handler.NewAgentHandler(agentServ)

	mux := httprouter.New()
	mux.GET("/v1/agent/:id", agentHandler.GetSingleAgent)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/agent/1")
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
	var mockAgent entity.Agent
	var agent entity.Agent
	_ = json.Unmarshal(body, &agent)
	mockAgent = entity.MockAgent

	if !reflect.DeepEqual(mockAgent, agent) {
		t.Errorf("want body to contain \n%q, but\n%q", mockAgent, agent)

	}
}