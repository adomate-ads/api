package models

import "time"

type Keyword struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Keyword string `json:"keyword" gorm:"type:varchar(128)" example:"Dentistry"`
	//TODO - Add Location
	CPC       float64   `json:"cpc" gorm:"type:float" example:"9.25"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetKeywords() ([]Keyword, error) {
	var keywords []Keyword
	if err := DB.Find(&keywords).Error; err != nil {
		return nil, err
	}
	return keywords, nil
}

func GetKeyword(id uint) (*Keyword, error) {
	var keyword Keyword
	if err := DB.First(&keyword, id).Error; err != nil {
		return nil, err
	}
	return &keyword, nil
}

func GetKeywordByName(name string) (*Keyword, error) {
	var keyword Keyword
	if err := DB.Where("keyword = ?", name).First(&keyword).Error; err != nil {
		return nil, err
	}
	return &keyword, nil
}

func (k *Keyword) CreateKeyword() error {
	err := DB.Create(&k).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *Keyword) UpdateKeyword() (*Keyword, error) {
	err := DB.Save(&k).Error
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (k *Keyword) DeleteKeyword() error {
	err := DB.Delete(&k).Error
	if err != nil {
		return err
	}
	return nil
}
