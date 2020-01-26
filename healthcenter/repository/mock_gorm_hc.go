package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"

	"github.com/TenaHub/api/healthcenter"
)

type MockHealthCenterGormRepo struct {
	conn *gorm.DB
}

func NewMockHealthCenterGormRepo(db *gorm.DB) healthcenter.HealthCenterRepository{
	return &MockHealthCenterGormRepo{conn:db}
}

func (adm MockHealthCenterGormRepo) HealthCenterById(id uint) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	healthcenter = entity.MockHealthCenter
	return &healthcenter, nil
}
func (adm MockHealthCenterGormRepo) HealthCenterByAgentId(id uint) ([]entity.HealthCenter, []error) {
	var healthcenters []entity.HealthCenter
	healthcenters = append(healthcenters, entity.MockHealthCenter,entity.MockHealthCenter)
	return healthcenters, nil
}
func (adm MockHealthCenterGormRepo) HealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := entity.MockHealthCenter
	return &healthcenter, nil
}

func (adm *MockHealthCenterGormRepo) HealthCenters() ([]entity.HealthCenter, []error) {
	var healthcenters []entity.HealthCenter
	healthcenters = append(healthcenters, entity.MockHealthCenter, entity.MockHealthCenter)
	return healthcenters, nil

}
func (adm *MockHealthCenterGormRepo) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.HealthCenterById(id)
	return healthcenter, errs
}

func (adm *MockHealthCenterGormRepo) UpdateHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := healthcenterData

	return healthcenter, []error{nil}
}

func (adm *MockHealthCenterGormRepo) SingleHealthCenter(id uint) (*entity.HealthCenter, []error) {
	hcs := entity.HealthCenter{}
	hcs = entity.MockHealthCenter
	return &hcs, nil
}

// HealthCenters returns all healthcenters data from database
func (adm *MockHealthCenterGormRepo) SearchHealthCenters(value string, column string) ([]entity.Hcrating, []error) {
	// hcs := []entity.HealthCenter{}
	hcsRating := []entity.Hcrating{}
	hcsRating = append(hcsRating, entity.Hcrating{}, entity.Hcrating{})
	return hcsRating, nil
}

// Top returns healthcenters with rating from database
func (adm *MockHealthCenterGormRepo) Top(amount uint) ([]entity.Hcrating, []error) {
	result := []entity.Hcrating{}
	result = append(result, entity.Hcrating{}, entity.Hcrating{})
	return result, nil
}
func (adm *MockHealthCenterGormRepo) StoreHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	center := healthcenterData
	return center, nil
}


