package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/admin"
	"github.com/TenaHub/api/entity"
	"errors"
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


