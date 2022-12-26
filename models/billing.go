package models

import "time"

type Billing struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	CompanyID uint `json:"name"`
	Company   Company
	Amount    float64 `json:"amount" gorm:"type:float" example:"900.25"`
	// Available options: paid, unpaid, pending
	Status    string    `json:"status" gorm:"type:varchar(10)" example:"paid"`
	Comments  string    `json:"comments" gorm:"type:varchar(1000)" example:"Something about the invoice..."`
	DueAt     time.Time `json:"due_at" example:"2020-01-01T00:00:00Z"`
	IssuedAt  time.Time `json:"issued_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetBillings() ([]Billing, error) {
	var billings []Billing
	if err := DB.Find(&billings).Error; err != nil {
		return nil, err
	}
	return billings, nil
}

func GetBilling(id uint) (*Billing, error) {
	var billing Billing
	if err := DB.First(&billing, id).Error; err != nil {
		return nil, err
	}
	return &billing, nil
}

func GetBillingsByCompanyID(id uint) ([]Billing, error) {
	var billings []Billing
	if err := DB.Where("company_id = ?", id).Find(&billings).Error; err != nil {
		return nil, err
	}
	return billings, nil
}

func (b *Billing) CreateBilling() error {
	err := DB.Create(&b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Billing) UpdateBilling() (*Billing, error) {
	err := DB.Save(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Billing) DeleteBilling() error {
	err := DB.Delete(&b).Error
	if err != nil {
		return err
	}
	return nil
}
