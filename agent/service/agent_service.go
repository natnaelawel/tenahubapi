package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/agent"
)

type AgentService struct {
	agentRepo agent.AgentRepository
}
func NewAgentService(serv agent.AgentRepository)(admin *AgentService){
	return &AgentService{agentRepo:serv}
}

func (adm *AgentService) AgentById(id uint) (*entity.Agent, []error) {
	agent, errs := adm.agentRepo.AgentById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
func (adm *AgentService) Agent(agent *entity.Agent) (*entity.Agent, []error) {
	agentData, errs := adm.agentRepo.Agent(agent)
	if len(errs) > 0 {
		return nil, errs
	}
	return agentData, errs
}
func (adm *AgentService) Agents() ([]entity.Agent, []error) {
	Agents, errs := adm.agentRepo.Agents()
	if len(errs) > 0 {
		return nil, errs
	}
	return Agents, errs
}
func (adm *AgentService) UpdateAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent, errs := adm.agentRepo.UpdateAgent(agentData)
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
func (adm *AgentService) StoreAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent, errs := adm.agentRepo.StoreAgent(agentData)
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
func (adm *AgentService) DeleteAgent(id uint) (*entity.Agent, []error) {
	agentData, errs := adm.agentRepo.DeleteAgent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return agentData, errs
}


