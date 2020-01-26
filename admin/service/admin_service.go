package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/admin"
)

type AdminService struct {
	adminRepo admin.AdminRepository
}
func NewAdminService(serv admin.AdminRepository)(admin *AdminService){
	return &AdminService{adminRepo:serv}
}

func (adm *AdminService) AdminById(id uint) (*entity.Admin, []error) {
	adminData, errs := adm.adminRepo.AdminById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return adminData, errs
}
func (adm *AdminService) Admin(admin *entity.Admin) (*entity.Admin, []error) {
	adminData, errs := adm.adminRepo.Admin(admin)
	if len(errs) > 0 {
		return nil, errs
	}
	return adminData, errs
}
func (adm *AdminService) UpdateAdmin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin, errs := adm.adminRepo.UpdateAdmin(adminData)
	if len(errs) > 0 {
		return nil, errs
	}
	return admin, errs
}
