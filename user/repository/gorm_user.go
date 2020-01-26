package repository

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserGormRepo is repository implements user.UserRepository
type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates and returns new UserGormRepo object
func NewUserGormRepo(dbConn *gorm.DB) *UserGormRepo {
	return &UserGormRepo{conn: dbConn}
}

// Users returns all users from the database
func (ur *UserGormRepo) Users() ([]entity.User, []error) {
	users := []entity.User{}
	errs := ur.conn.Find(&users).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return users, nil
}

// User returns a single users from the database with user name and password
func (ur *UserGormRepo) User(user *entity.User) (*entity.User, []error) {
	lgusr := user
	usr := entity.User{}
	errs := ur.conn.Where("email = ?", user.Email).First(&usr).GetErrors()
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(lgusr.Password))
	fmt.Println(err)
	if err != nil {
		return nil, []error{err}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &usr, nil
}

// UserByID returns a single user by its id from the database
func (ur *UserGormRepo) UserByID(id uint) (*entity.User, []error) {
	usr := entity.User{}
	errs := ur.conn.First(&usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &usr, nil
}

// UpdateUser updates user from the database
func (ur *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := ur.conn.Save(usr).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// DeleteUser deletes a single user from the database by it's id
func (ur *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := ur.UserByID(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = ur.conn.Delete(usr, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// StoreUser will insert a new user to the database
func (ur *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := ur.conn.Create(usr).GetErrors()

	for _, err := range errs {
		pqerr := err.(*pq.Error)
		fmt.Println(pqerr)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}
