package service

import "github.com/natnaelawel/tenahubapi/entity"

type ServicesService interface {
	Service(id uint) (*entity.Service, []error)
	PendingService(id uint) ([]entity.Service, []error)
	Services(id uint) ([]entity.Service, []error)
	UpdateService(user *entity.Service) (*entity.Service, []error)
	StoreService(user *entity.Service) (*entity.Service, []error)
	DeleteService(id uint) (*entity.Service, []error)
}