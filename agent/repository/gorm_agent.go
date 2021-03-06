package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/natnaelawel/tenahubapi/entity"
	"github.com/natnaelawel/tenahubapi/agent"
	"github.com/natnaelawel/tenahubapi/delivery/http/handler"
)

type AgentGormRepo struct {
	conn *gorm.DB
}

func NewAgentGormRepo(db *gorm.DB) agent.AgentRepository{
	return &AgentGormRepo{conn:db}
}

func (adm *AgentGormRepo) Agent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := entity.Agent{}
	errs := adm.conn.Select("password").Where("email = ? ", agentData.Email).First(&agent).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	same := handler.VerifyPassword(agentData.Password, agent.Password)
	if same {
		errs := adm.conn.Where("email = ?", agentData.Email).First(&agent).GetErrors()
		return &agent, errs
	}
	return nil, errs

}
func (adm *AgentGormRepo) AgentById(id uint) (*entity.Agent, []error) {
	agent := entity.Agent{}
	errs := adm.conn.First(&agent, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &agent, errs
}
func (adm *AgentGormRepo) Agents() ([]entity.Agent, []error) {
	var agents []entity.Agent
	errs := adm.conn.Find(&agents).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agents, errs
}
func (adm *AgentGormRepo) UpdateAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	if agentData.Password != "" {
		agent.Password,_ = handler.HashPassword(agentData.Password)
	}
	errs := adm.conn.Model(&agentData).Updates(agent).Error
	if errs != nil {
		return nil, []error{errs}
	}

	return agent, nil

}
func (adm *AgentGormRepo) StoreAgent(agentData *entity.Agent) (*entity.Agent, []error) {
	agent := agentData
	agent.Password,_ = handler.HashPassword(agentData.Password)
	errs := adm.conn.Create(agent).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
func (adm *AgentGormRepo) DeleteAgent(id uint) (*entity.Agent, []error) {
	agent, errs := adm.AgentById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = adm.conn.Delete(agent, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return agent, errs
}
