package models

import "time"

type Service struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name      string    `json:"name" gorm:"type:varchar(255)" example:"Digital Marketing"`
	CompanyID uint      `json:"company_id" gorm:"type:integer" example:"1"`
	Company   Company   `json:"company" gorm:"foreignKey:CompanyID"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetServices() ([]Service, error) {
	var services []Service
	if err := DB.Preload("Company").Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func GetService(id uint) (*Service, error) {
	var service Service
	if err := DB.Preload("Company").First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func GetServicesByCompanyID(id uint) ([]Service, error) {
	var services []Service
	if err := DB.Preload("Company").Where("company_id = ?", id).Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (s *Service) CreateService() error {
	err := DB.Create(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateService() (*Service, error) {
	err := DB.Updates(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Service) DeleteService() error {
	err := DB.Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}
