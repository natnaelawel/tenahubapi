package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/admin"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"fmt"
)

type AdminGormRepo struct {
	conn *gorm.DB
}

func NewAdminGormRepo(db *gorm.DB) admin.AdminRepository{
	return &AdminGormRepo{conn:db}
}

func (adm *AdminGormRepo) Admin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := entity.Admin{}
	//errs := adm.conn.Where("email = ? AND password = ?", adminData.Email, adminData.Password).First(&admin).GetErrors()
	//if len(errs) > 0 {
	//	return nil, errs
	//}
	//return &admin, errs
	errs := adm.conn.Select("password").Where("email = ? ", adminData.Email).First(&admin).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	same := handler.VerifyPassword(adminData.Password, admin.Password)
	fmt.Println("is same",same)
	if same {
		errs := adm.conn.Where("email = ?", adminData.Email).First(&admin).GetErrors()
		return &admin, errs
	}
	return nil, errs

}
func (adm *AdminGormRepo) AdminById(id uint) (*entity.Admin, []error) {
	admin := entity.Admin{}
	errs := adm.conn.First(&admin, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &admin, errs
}

func (adm *AdminGormRepo) UpdateAdmin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := adminData
	//errs := adm.conn.Save(admin).GetErrors()
	//if len(errs) > 0 {
	//	return nil, errs
	//}
	//return admin, errs
	data := entity.Admin{}
	if adminData.Password != "" {
		admin.Password,_ = handler.HashPassword(adminData.Password)
	}
	//errs := adm.conn.Save(healthcenter).GetErrors()
	errs := adm.conn.Model(&data).Updates(admin).Error
	if errs != nil {
		return nil, []error{errs}
	}
	return &data, []error{errs}

}


