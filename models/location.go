package models

import (
	"time"
)

type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name      string    `json:"name" gorm:"type:varchar(128)" example:"Austin, TX"`
	CompanyID uint      `json:"company_id" gorm:"type:integer" example:"1"`
	Company   Company   `json:"company" gorm:"foreignKey:CompanyID"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetLocations() ([]Location, error) {
	var locations []Location
	if err := DB.Preload("Company").Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func GetLocation(id uint) (*Location, error) {
	var location Location
	if err := DB.Preload("Company").First(&location, id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func GetLocationByCompanyID(companyID uint) (*Location, error) {
	var location Location
	if err := DB.Where("company_id = ?", companyID).Preload("Company").First(&location).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (l *Location) CreateLocation() error {
	err := DB.Create(&l).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *Location) UpdateLocation() (*Location, error) {
	err := DB.Save(&l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Location) DeleteLocation() error {
	err := DB.Delete(&l).Error
	if err != nil {
		return err
	}
	return nil
}
