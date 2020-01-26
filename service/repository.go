package service

import "github.com/TenaHub/api/entity"

type ServiceRepository interface {
	Service(id uint) (*entity.Service, []error)
	PendingService(id uint) ([]entity.Service, []error)
	Services(id uint) ([]entity.Service, []error)
	UpdateService(user *entity.Service) (*entity.Service, []error)
	StoreService(user *entity.Service) (*entity.Service, []error)
	DeleteService(id uint) (*entity.Service, []error)
}