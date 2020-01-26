package repository

import (
	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/user"
)

// UserGormRepo is repository implements user.UserRepository
type MockUserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates and returns new UserGormRepo object
func NewMockUserGormRepo(dbConn *gorm.DB) user.UserRepository {
	return &MockUserGormRepo{conn: dbConn}
}

// Users returns all users from the database
func (ur *MockUserGormRepo) Users() ([]entity.User, []error) {
	users := []entity.User{}
	users = append(users, entity.MockUser, entity.MockUser)
	return users, nil
}

// User returns a single users from the database with user name and password
func (ur *MockUserGormRepo) User(user *entity.User) (*entity.User, []error) {
	usr := entity.MockUser
	return &usr, nil
}

// UserByID returns a single user by its id from the database
func (ur *MockUserGormRepo) UserByID(id uint) (*entity.User, []error) {
	usr := entity.MockUser
	return &usr, nil
}

// UpdateUser updates user from the database
func (ur *MockUserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user

	return usr, nil
}

// DeleteUser deletes a single user from the database by it's id
func (ur *MockUserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, _ := ur.UserByID(id)

	return usr, nil
}

// StoreUser will insert a new user to the database
func (ur *MockUserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user

	return usr, nil
}
