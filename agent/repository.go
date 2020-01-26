package agent


import "github.com/natnaelawel/tenahubapi/api/entity"

type AgentRepository interface {
	AgentById(id uint) (*entity.Agent, []error)
	Agents() ([]entity.Agent, []error)
	Agent(agent *entity.Agent)(*entity.Agent, []error)
	UpdateAgent(user *entity.Agent) (*entity.Agent, []error)
	StoreAgent(user *entity.Agent) (*entity.Agent, []error)
	DeleteAgent(id uint) (*entity.Agent, []error)
}
