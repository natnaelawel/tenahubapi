package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	// "github.com/TenaHub/api/healthcenter"
	"github.com/TenaHub/api/delivery/http/handler"
	"github.com/TenaHub/api/healthcenter"
	"fmt"
)

type HealthCenterGormRepo struct {
	conn *gorm.DB
}

func NewHealthCenterGormRepo(db *gorm.DB) healthcenter.HealthCenterRepository{
	return &HealthCenterGormRepo{conn:db}
}

func (adm HealthCenterGormRepo) HealthCenterById(id uint) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	errs := adm.conn.First(&healthcenter, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &healthcenter, errs
}
func (adm HealthCenterGormRepo) HealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	//errs := adm.conn.Where("email = ? AND password = ?", healthcenterData.Email, healthcenterData.Password).First(&healthcenter).GetErrors()
	//if len(errs) > 0 {
	//	return nil, errs
	//}
	errs := adm.conn.Select("password").Where("email = ? ", healthcenterData.Email).First(&healthcenter).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	same := handler.VerifyPassword(healthcenterData.Password, healthcenter.Password)
	if same {
		errs := adm.conn.Where("email = ?", healthcenterData.Email).First(&healthcenter).GetErrors()
		return &healthcenter, errs
	}
	return nil, errs
}

func (adm *HealthCenterGormRepo) HealthCenters() ([]entity.HealthCenter, []error) {
	var healthcenters []entity.HealthCenter
	errs := adm.conn.Find(&healthcenters).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenters, errs

}
func (adm *HealthCenterGormRepo) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.HealthCenterById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = adm.conn.Delete(healthcenter, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}

func (adm *HealthCenterGormRepo) UpdateHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := healthcenterData
	data := entity.HealthCenter{}
	healthcenter.Password,_ = handler.HashPassword(healthcenterData.Password)
	//errs := adm.conn.Save(healthcenter).GetErrors()
	errs := adm.conn.Model(&data).Updates(healthcenter).Error
	if errs != nil {
		return nil, []error{errs}
	}
	return healthcenter, []error{errs}
}

func (adm *HealthCenterGormRepo) SingleHealthCenter(id uint) (*entity.HealthCenter, []error) {
	hcs := entity.HealthCenter{}
	errs := adm.conn.Where("id = ?", id).First(&hcs).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return &hcs, nil
}
func (adm HealthCenterGormRepo) HealthCenterByAgentId(id uint) ([]entity.HealthCenter, []error) {
	var healthcenters []entity.HealthCenter

	//errs := adm.conn.Find(&healthcenters, id).GetErrors()
	errs := adm.conn.Where("agent_id = ?", id).Find(&healthcenters).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenters, errs
}

// HealthCenters returns all healthcenters data from database
func (adm *HealthCenterGormRepo) SearchHealthCenters(value string, column string) ([]entity.Hcrating, []error) {
	// hcs := []entity.HealthCenter{}
	hcsRating := []entity.Hcrating{}
	fmt.Println("value, column", value, column)
	switch column {
	case "name":
		fmt.Println("name")
		// errs := hcr.conn.Where("name ILIKE ?", "%"+value+"%").Find(&hcs).GetErrors()
		errs := adm.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.name ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()
		fmt.Println(hcsRating)
		if len(errs) > 0 {
			return nil, errs
		}
		return hcsRating, nil
	case "city":
		fmt.Println("city")
		// errs := hcr.conn.Where("city ILIKE ?", "%"+value+"%").Find(&hcs).GetErrors()
		errs := adm.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.city ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()

		if len(errs) > 0 {
			return nil, errs
		}
		return hcsRating, nil
	case "service":
		fmt.Println("service")
		// errs := hcr.conn.Raw("select * from health_centers where id in (?)", hcr.conn.Table("services").Select("health_center_id").Where("name ILIKE ?", "%"+value+"%").QueryExpr()).Find(&hcr).GetErrors()
		result := []struct {
			HealthCenterID int
		}{}
		errs := adm.conn.Table("services").Select("health_center_id").Where("name ILIKE ?", "%"+value+"%").Find(&result).GetErrors()
		fmt.Println(result)

		if len(errs) > 0 {
			return nil, errs
		}

		arr := make([]int, len(result))

		for _, hid := range result {
			arr = append(arr, hid.HealthCenterID)
		}
		// errs = hcr.conn.Table("health_centers").Select("health_centers.*, avg(comments.rating) as rating").Joins("left join comments on comments.health_center_id = health_centers.id").Where("health_centers.id in (?)", arr).Group("healch_centers.id").Scan(&hcsRating).GetErrors()
		errs = adm.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.id in (?) group by health_centers.id order by rating desc;", arr).Scan(&hcsRating).GetErrors()

		fmt.Println(errs)
		fmt.Println(hcsRating)

		if len(errs) > 0 {
			return nil, errs
		}
		return hcsRating, nil
	default:
		fmt.Println("default")
		errs := adm.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id where health_centers.name ILIKE ? group by health_centers.id order by rating desc;", "%"+value+"%").Scan(&hcsRating).GetErrors()
		if len(errs) > 0 {
			return nil, errs
		}
		return hcsRating, nil
	}
}

// Top returns healthcenters with rating from database
func (adm *HealthCenterGormRepo) Top(amount uint) ([]entity.Hcrating, []error) {
	result := []entity.Hcrating{}
	errs := adm.conn.Raw("select health_centers.*, avg(comments.rating) as rating from health_centers left join comments on health_centers.id = comments.health_center_id group by health_centers.id order by rating limit ?;", amount).Scan(&result).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println(result)
	return result, nil
}
func (adm *HealthCenterGormRepo) StoreHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	center := healthcenterData
	center.Password,_ = handler.HashPassword(healthcenterData.Password)
	errs := adm.conn.Create(&center).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return center, errs
}


