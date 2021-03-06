package admin
import "github.com/natnaelawel/tenahubapi/entity"

type AdminRepository interface {
	AdminById(id uint)(*entity.Admin, []error)
	Admin(admin *entity.Admin)(*entity.Admin, []error)
	UpdateAdmin(user *entity.Admin) (*entity.Admin, []error)
	StoreAdmin(user *entity.Admin) (*entity.Admin, []error)
	DeleteAdmin(id uint) (*entity.Admin, []error)
}

