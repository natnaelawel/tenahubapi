package admin

import "github.com/natnaelawel/tenahubapi/api/entity"

type AdminService interface {
	AdminById(id uint)(*entity.Admin, []error)
	Admin(admin *entity.Admin)(*entity.Admin, []error)
	UpdateAdmin(user *entity.Admin) (*entity.Admin, []error)

}
