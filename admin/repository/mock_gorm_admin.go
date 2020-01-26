package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/natnaelawel/tenahubapi/admin"
	"github.com/natnaelawel/tenahubapi/entity"
)

type MockAdminGormRepo struct {
	conn *gorm.DB
}

func NewMockAdminGormRepo(db *gorm.DB) admin.AdminRepository{
	return &MockAdminGormRepo{conn:db}
}

func (adm *MockAdminGormRepo) Admin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := entity.MockAdmin
	return &admin, nil

}
func (adm *MockAdminGormRepo) AdminById(id uint) (*entity.Admin, []error) {
	admin := entity.MockAdmin
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &admin, nil
}

func (adm *MockAdminGormRepo) UpdateAdmin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := adminData
	return admin, []error{nil}
}
func (adm *MockAdminGormRepo) StoreAdmin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := adminData
	return admin, nil
}
func (adm *MockAdminGormRepo) DeleteAdmin(id uint) (*entity.Admin, []error) {
	admin, errs := adm.AdminById(id)
	return admin, errs
}


