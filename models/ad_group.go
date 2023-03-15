package models

import "time"

type AdGroup struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Name         string    `json:"name" gorm:"type:varchar(128)" example:"Primary Monthly"`
	ResourceName string    `json:"resource_name" gorm:"type:varchar(128)" example:"/customers/1234567890/adGroups/1234567890"`
	CompanyID    uint      `json:"company_id" gorm:"type:integer" example:"1"`
	Company      Company   `json:"company" gorm:"foreignKey:CompanyID"`
	CampaignID   uint      `json:"campaign_id" gorm:"type:integer" example:"1"`
	Campaign     Campaign  `json:"campaign" gorm:"foreignKey:CampaignID"`
	GoogleID     uint      `json:"google_id" gorm:"type:integer" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetAdGroups() ([]AdGroup, error) {
	var adGroups []AdGroup
	if err := DB.Preload("Company").Preload("Campaign").Find(&adGroups).Error; err != nil {
		return nil, err
	}
	return adGroups, nil
}

func GetAdGroup(id uint) (*AdGroup, error) {
	var adGroup AdGroup
	if err := DB.Preload("Company").Preload("Campaign").First(&adGroup, id).Error; err != nil {
		return nil, err
	}
	return &adGroup, nil
}

func GetAdGroupsByCompanyID(id uint) ([]AdGroup, error) {
	var adGroups []AdGroup
	if err := DB.Where("company_id = ?", id).Preload("Company").Preload("Campaign").Find(&adGroups).Error; err != nil {
		return nil, err
	}
	return adGroups, nil
}

func GetAdGroupsByCampaignID(id uint) ([]AdGroup, error) {
	var adGroups []AdGroup
	if err := DB.Where("campaign_id = ?", id).Preload("Company").Preload("Campaign").Find(&adGroups).Error; err != nil {
		return nil, err
	}
	return adGroups, nil
}

func GetAdGroupByGoogleID(id uint) (*AdGroup, error) {
	var adGroup AdGroup
	if err := DB.Where("google_id = ?", id).Preload("Company").Preload("Campaign").First(&adGroup).Error; err != nil {
		return nil, err
	}
	return &adGroup, nil
}

func (a *AdGroup) CreateAdGroup() error {
	err := DB.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AdGroup) UpdateAdGroup() (*AdGroup, error) {
	err := DB.Save(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AdGroup) DeleteAdGroup() error {
	err := DB.Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}
