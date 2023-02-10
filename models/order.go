package models

import (
	"time"
)

type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	CompanyID uint    `json:"company_id" gorm:"type:integer" example:"1"`
	Company   Company `json:"company" gorm:"foreignKey:CompanyID"`
	Amount    float64 `json:"amount" gorm:"type:float" example:"900.25"`
	// Available options: pending, active, cancelled
	Status string `json:"status" gorm:"type:varchar(10)" example:"pending"`
	// Available options: base, premium, enterprise, ads
	Type string `json:"type" gorm:"type:varchar(1000)" example:"base"`
	// TODO - Check Stripe transaction ID standards for better example and update varchar length
	StartAt   time.Time `json:"start_at" example:"2020-01-01T00:00:00Z"`
	EndAt     time.Time `json:"end_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetOrders() ([]Order, error) {
	var orders []Order
	if err := DB.Preload("Company").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrder(id uint) (*Order, error) {
	var order Order
	if err := DB.Preload("Company").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func GetOrdersByCompanyID(id uint) ([]Order, error) {
	var orders []Order
	if err := DB.Preload("Company").Where("company_id = ?", id).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *Order) CreateOrder() error {
	err := DB.Create(&o).Error
	if err != nil {
		return err
	}
	return nil
}

type UpdateOrder struct {
	Amount float64   `json:"amount"`
	Status string    `json:"status"`
	Type   string    `json:"type"`
	EndAt  time.Time `json:"end_at"`
}

func (o *Order) UpdateOrder(UpdatedOrder UpdateOrder) (*Order, error) {
	err := DB.Model(o).Updates(UpdatedOrder).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Order) DeleteOrder() error {
	err := DB.Delete(&o).Error
	if err != nil {
		return err
	}
	return nil
}
