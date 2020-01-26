package healthcenter

import "github.com/TenaHub/api/entity"

type HealthCenterService interface {
	SingleHealthCenter(id uint) (*entity.HealthCenter, []error)
	SearchHealthCenters(value string, column string) ([]entity.Hcrating, []error)
	Top(amount uint)([]entity.Hcrating, []error)
	HealthCenterById(id uint) (*entity.HealthCenter, []error)
	HealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)
	HealthCenters() ([]entity.HealthCenter, []error)
	DeleteHealthCenter(id uint) (*entity.HealthCenter, []error)
	UpdateHealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)
	StoreHealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)
	HealthCenterByAgentId(id uint) ([]entity.HealthCenter, []error)

}
