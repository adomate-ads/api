package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	Role      string    `json:"Role" gorm:"type:varchar(128)" example:"Administrator"`
	CreatedAt time.Time `json:"created_at" example:"2020-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-01-01T00:00:00Z"`
}

func GetRoles() ([]Role, error) {
	var roles []Role
	if err := DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRole(id uint) (*Role, error) {
	var role Role
	if err := DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func GetRoleByName(name string) (*Role, error) {
	var role Role
	if err := DB.Where("role = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Role) CreateRole() error {
	err := DB.Create(&r).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) UpdateRole() (*Role, error) {
	err := DB.Save(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Role) DeleteRole() error {
	err := DB.Delete(&r).Error
	if err != nil {
		return err
	}
	return nil
}
