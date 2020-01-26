package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/healthcenter"
	"fmt"
)

type HealthCenterService struct {
	healthCenterRepo healthcenter.HealthCenterRepository
}

func NewHealthCenterService(serv healthcenter.HealthCenterRepository)(admin *HealthCenterService){
	return &HealthCenterService{healthCenterRepo:serv}
}


func (adm *HealthCenterService) HealthCenterById(id uint) (*entity.HealthCenter, []error) {
	healthCenter, errs := adm.healthCenterRepo.HealthCenterById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenter, errs
}
func (adm *HealthCenterService) HealthCenterByAgentId(id uint) ([]entity.HealthCenter, []error) {
	healthCenter, errs := adm.healthCenterRepo.HealthCenterByAgentId(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenter, errs
}
func (adm *HealthCenterService) HealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthCenter, errs := adm.healthCenterRepo.HealthCenter(healthcenter)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenter, errs
}
func (adm *HealthCenterService) HealthCenters() ([]entity.HealthCenter, []error) {
	healthCenters, errs := adm.healthCenterRepo.HealthCenters()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenters, errs
}
func (adm *HealthCenterService) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.healthCenterRepo.DeleteHealthCenter(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}
func (adm *HealthCenterService) UpdateHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.healthCenterRepo.UpdateHealthCenter(healthcenterData)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}

// HealthCenter returns single healthcenter data
func (hcs *HealthCenterService) SingleHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := hcs.healthCenterRepo.SingleHealthCenter(uint(id))

	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, nil
}

// HealthCenters returns all healthcenters data
func (hcs *HealthCenterService) SearchHealthCenters(value string, column string) ([]entity.Hcrating, []error) {
	healthcenters, errs := hcs.healthCenterRepo.SearchHealthCenters(value, column)

	if errs != nil {
		return nil, errs
	}
	return healthcenters, nil
}

// Top returns healthcenters with rating from database
func (hcs *HealthCenterService) Top(amount uint) ([]entity.Hcrating, []error) {
	result, errs := hcs.healthCenterRepo.Top(amount)
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println(result)
	return result, nil
}

func (adm *HealthCenterService) StoreHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.healthCenterRepo.StoreHealthCenter(healthcenterData)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}