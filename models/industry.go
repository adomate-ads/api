package models

import "time"

type Industry struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Industry  string    `json:"Industry" gorm:"type:varchar(128)" example:"Health Care"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetIndustries() ([]Industry, error) {
	var industries []Industry
	if err := DB.Find(&industries).Error; err != nil {
		return nil, err
	}
	return industries, nil
}

func GetIndustry(id uint) (*Industry, error) {
	var industry Industry
	if err := DB.First(&industry, id).Error; err != nil {
		return nil, err
	}
	return &industry, nil
}

func GetIndustryByName(name string) (*Industry, error) {
	var industry Industry
	if err := DB.Where("industry = ?", name).First(&industry).Error; err != nil {
		return nil, err
	}
	return &industry, nil
}

func (i *Industry) CreateIndustry() error {
	err := DB.Create(&i).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *Industry) UpdateIndustry() (*Industry, error) {
	err := DB.Save(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *Industry) DeleteIndustry() error {
	err := DB.Delete(&i).Error
	if err != nil {
		return err
	}
	return nil
}
